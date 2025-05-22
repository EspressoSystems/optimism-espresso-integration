package derive

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/ethereum-optimism/optimism/espresso"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	lightclient "github.com/EspressoSystems/espresso-network-go/light-client"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// The attributes queue sits after the batch queue.
// It transforms batches into payload attributes. The outputted payload
// attributes cannot be buffered because each batch->attributes transformation
// pulls in data about the current L2 safe head.
//
// It also buffers batches that have been output because multiple batches can
// be created at once.
//
// This stage can be reset by clearing its batch buffer.
// This stage does not need to retain any references to L1 blocks.

type AttributesBuilder interface {
	PreparePayloadAttributes(ctx context.Context, l2Parent eth.L2BlockRef, epoch eth.BlockID) (attrs *eth.PayloadAttributes, err error)
}

type AttributesWithParent struct {
	Attributes *eth.PayloadAttributes
	Parent     eth.L2BlockRef
	Concluding bool // Concluding indicates that the attributes conclude the pending safe phase

	DerivedFrom eth.L1BlockRef
}

// WithDepositsOnly return a shallow clone with all non-Deposit transactions
// stripped from the transactions of its attributes. The order is preserved.
func (a *AttributesWithParent) WithDepositsOnly() *AttributesWithParent {
	clone := *a
	clone.Attributes = clone.Attributes.WithDepositsOnly()
	return &clone
}

func (a *AttributesWithParent) IsDerived() bool {
	return a.DerivedFrom != (eth.L1BlockRef{})
}

type AttributesQueue struct {
	log     log.Logger
	config  *rollup.Config
	builder AttributesBuilder
	prev    SingularBatchProvider

	batch       *SingularBatch
	concluding  bool
	lastAttribs *AttributesWithParent

	isCaffNode       bool
	espressoStreamer *espresso.EspressoStreamer[EspressoBatch]
}

type SingularBatchProvider interface {
	ResettableStage
	ChannelFlusher
	Origin() eth.L1BlockRef
	NextBatch(context.Context, eth.L2BlockRef) (*SingularBatch, bool, error)
}

func initEspressoStreamer(log log.Logger, cfg *rollup.Config, l1Fetcher L1Fetcher) *espresso.EspressoStreamer[EspressoBatch] {

	if !cfg.CaffNodeConfig.IsCaffNode {
		return nil
	}

	// Create an adapter that implements espresso.L1Client
	l1BlockRefClient := NewL1BlockRefClient(l1Fetcher)

	l1Client, err := ethclient.Dial(cfg.CaffNodeConfig.L1EthRpc)
	if err != nil {
		return nil
	}

	lightClient, err := lightclient.NewLightclientCaller(common.HexToAddress(cfg.CaffNodeConfig.EspressoLightClientAddr), l1Client)
	if err != nil {
		return nil
	}

	streamer := espresso.NewEspressoStreamer(
		cfg.L2ChainID.Uint64(),
		l1BlockRefClient,
		espressoClient.NewMultipleNodesClient(cfg.CaffNodeConfig.HotShotUrls),
		lightClient,
		log,
		func(data []byte) (*EspressoBatch, error) {
			return UnmarshalEspressoTransaction(data, cfg.Genesis.SystemConfig.BatcherAddr)
		},
		cfg.CaffNodeConfig.PollingHotShotPollingInterval,
	)

	log.Debug("Espresso Streamer namespace:", streamer.Namespace)

	log.Info("Espresso streamer initialized", "namespace", cfg.L2ChainID.Uint64(), "next hotshot block num", cfg.CaffNodeConfig.NextHotShotBlockNum, "polling hotshot polling interval", cfg.CaffNodeConfig.PollingHotShotPollingInterval, "hotshot urls", cfg.CaffNodeConfig.HotShotUrls)
	return &streamer
}

func NewAttributesQueue(log log.Logger, cfg *rollup.Config, builder AttributesBuilder, prev SingularBatchProvider, l1Fetcher L1Fetcher) *AttributesQueue {
	return &AttributesQueue{
		log:              log,
		config:           cfg,
		builder:          builder,
		prev:             prev,
		isCaffNode:       cfg.CaffNodeConfig.IsCaffNode,
		espressoStreamer: initEspressoStreamer(log, cfg, l1Fetcher),
	}
}

func (aq *AttributesQueue) Origin() eth.L1BlockRef {
	return aq.prev.Origin()
}

// CaffNextBatch fetches the next batch from the Espresso streamer for the caff node.
//
// It follows the flow: Refresh() -> Update() -> Next().
//
// This is similar to the batcher's flow: espressoBatchLoadingLoop -> getSyncStatus -> refresh -> Update -> Next,
// but with a few key differences:
// - CaffNextBatch obtains sync state differently from the batcher, it treated parent.Number() as the latest safe batch number.
// - It only calls Update() when needed and everytime only calls Next() once. While the batcher calls Next() in a loop.
// - It performs additional checks, such as validating the timestamp and parent hash, which does not apply to the batcher.
func CaffNextBatch(s *espresso.EspressoStreamer[EspressoBatch], ctx context.Context, parent eth.L2BlockRef, blockTime uint64, l1Fetcher L1Fetcher) (*SingularBatch, bool, error) {
	// Get the L1 finalized block
	finalizedL1Block, err := l1Fetcher.L1BlockRefByLabel(ctx, eth.Finalized)
	if err != nil {
		s.Log.Error("failed to get the L1 finalized block", "err", err)
		return nil, false, err
	}
	// Refresh the sync status
	if _, err := s.Refresh(ctx, finalizedL1Block, parent.Number, parent.L1Origin); err != nil {
		return nil, false, err
	}

	// Update the streamer if needed
	if !s.HasNext(ctx) {
		err := s.Update(ctx)
		if err != nil {
			s.Log.Error("failed to update Espresso streamer", "err", err)
		}
	}

	// Get the next batch
	var espressoBatch = s.Next(ctx)

	if espressoBatch == nil {
		return nil, true, NotEnoughData
	}

	batch := &espressoBatch.Batch
	s.Log.Info("espressoBatch", "batch", espressoBatch.Batch)

	// These batch checks are retained because they add minimal latency (O(1) per batch).
	// They're primarily a safeguard for cases where the streamer fails to emit batches correctly,
	// which should only happen if there's a bug.
	{
		// check the batch is valid regarding given parent
		nextTimestamp := parent.Time + blockTime

		if batch.Timestamp != nextTimestamp {
			s.Log.Warn("Dropping batch", "batch", espressoBatch.Number(), "timestamp", batch.Timestamp, "expected", nextTimestamp)
			return nil, false, ErrTemporary
		}

		// dependent on above timestamp check. If the timestamp is correct, then it must build on top of the safe head.
		if batch.ParentHash != parent.Hash {
			s.Log.Warn("ignoring batch with mismatching parent hash", "current_safe_head", parent.Hash)
			return nil, false, ErrTemporary
		}
	}
	// For caff node, when we get a batch, we assign concluding to true to drive progress
	concluding := true
	return batch, concluding, nil
}

func (aq *AttributesQueue) NextAttributes(ctx context.Context, parent eth.L2BlockRef, l1Fetcher L1Fetcher) (*AttributesWithParent, error) {
	// Get a batch if we need it
	if aq.batch == nil {
		var batch *SingularBatch
		var concluding bool
		var err error
		if aq.isCaffNode {
			batch, concluding, err = CaffNextBatch(aq.espressoStreamer, ctx, parent, aq.config.BlockTime, l1Fetcher)
		} else {
			batch, concluding, err = aq.prev.NextBatch(ctx, parent)
		}

		if err != nil {
			return nil, err
		}
		aq.batch = batch
		aq.concluding = concluding
		aq.log.Info("singular batch from op-node is ", "batch", aq.batch, "concluding", concluding)
	}

	// Actually generate the next attributes
	if attrs, err := aq.createNextAttributes(ctx, aq.batch, parent); err != nil {
		return nil, err
	} else {
		// Clear out the local state once we will succeed
		attr := AttributesWithParent{
			Attributes:  attrs,
			Parent:      parent,
			Concluding:  aq.concluding,
			DerivedFrom: aq.Origin(),
		}
		aq.lastAttribs = &attr
		aq.batch = nil
		aq.concluding = false
		return &attr, nil
	}
}

// createNextAttributes transforms a batch into a payload attributes. This sets `NoTxPool` and appends the batched transactions
// to the attributes transaction list
func (aq *AttributesQueue) createNextAttributes(ctx context.Context, batch *SingularBatch, l2SafeHead eth.L2BlockRef) (*eth.PayloadAttributes, error) {
	// sanity check parent hash
	if batch.ParentHash != l2SafeHead.Hash {
		return nil, NewResetError(fmt.Errorf("valid batch has bad parent hash %s, expected %s", batch.ParentHash, l2SafeHead.Hash))
	}
	// sanity check timestamp
	if expected := l2SafeHead.Time + aq.config.BlockTime; expected != batch.Timestamp {
		return nil, NewResetError(fmt.Errorf("valid batch has bad timestamp %d, expected %d", batch.Timestamp, expected))
	}
	fetchCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	attrs, err := aq.builder.PreparePayloadAttributes(fetchCtx, l2SafeHead, batch.Epoch())
	if err != nil {
		return nil, err
	}

	// we are verifying, not sequencing, we've got all transactions and do not pull from the tx-pool
	// (that would make the block derivation non-deterministic)
	attrs.NoTxPool = true
	attrs.Transactions = append(attrs.Transactions, batch.Transactions...)

	aq.log.Info("generated attributes in payload queue", "txs", len(attrs.Transactions), "timestamp", batch.Timestamp)

	return attrs, nil
}

func (aq *AttributesQueue) reset() {
	aq.batch = nil
	aq.concluding = false // overwritten later, but set for consistency
	aq.lastAttribs = nil
}

func (aq *AttributesQueue) Reset(ctx context.Context, _ eth.L1BlockRef, _ eth.SystemConfig) error {
	aq.reset()
	return io.EOF
}

func (aq *AttributesQueue) DepositsOnlyAttributes(parent eth.BlockID, derivedFrom eth.L1BlockRef) (*AttributesWithParent, error) {
	// Sanity checks - these cannot happen with correct deriver implementations.
	if aq.batch != nil {
		return nil, fmt.Errorf("unexpected buffered batch, parent hash: %s, epoch: %s", aq.batch.ParentHash, aq.batch.Epoch())
	} else if aq.lastAttribs == nil {
		return nil, errors.New("no attributes generated yet")
	} else if derivedFrom != aq.lastAttribs.DerivedFrom {
		return nil, fmt.Errorf(
			"unexpected derivation origin, last_origin: %s, invalid_origin: %s",
			aq.lastAttribs.DerivedFrom, derivedFrom)
	} else if parent != aq.lastAttribs.Parent.ID() {
		return nil, fmt.Errorf(
			"unexpected parent: last_parent: %s, invalid_parent: %s",
			aq.lastAttribs.Parent.ID(), parent)
	}

	aq.prev.FlushChannel() // flush all channel data in previous stages
	attrs := aq.lastAttribs.WithDepositsOnly()
	aq.lastAttribs = attrs
	return attrs, nil
}
