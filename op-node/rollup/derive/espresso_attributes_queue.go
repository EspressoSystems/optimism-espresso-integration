//go:build !mips64

package derive

import (
	"context"
	"fmt"

	op "github.com/EspressoSystems/espresso-streamers/op"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum-optimism/optimism/espresso/logmodule"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type espressoAttributesQueue struct {
	isCaffNode           bool
	caffeinationHeightL2 uint64
	espressoStreamer     *op.BatchStreamer[EspressoBatch]
}

func newEspressoAttributesQueue(logger log.Logger, cfg *rollup.Config) espressoAttributesQueue {
	return espressoAttributesQueue{
		isCaffNode:           cfg.CaffNodeConfig.Enabled,
		caffeinationHeightL2: cfg.CaffNodeConfig.CaffeinationHeightL2,
		espressoStreamer:     initEspressoStreamer(logger, cfg),
	}
}

func (e *espressoAttributesQueue) nextBatch(ctx context.Context, parent eth.L2BlockRef, blockTime uint64, l1Fetcher L1Fetcher, prev SingularBatchProvider, logger log.Logger) (*SingularBatch, bool, error) {
	if e.isCaffNode && parent.Number >= e.caffeinationHeightL2 {
		if e.espressoStreamer == nil {
			logger.Error("Espresso streamer not initialized as expected when isCaffNode is ON")
			return nil, false, ErrCritical
		}
		return CaffNextBatch(e.espressoStreamer, ctx, parent, blockTime, l1Fetcher)
	}
	return prev.NextBatch(ctx, parent)
}

func initEspressoStreamer(log log.Logger, cfg *rollup.Config) *op.BatchStreamer[EspressoBatch] {
	if !cfg.CaffNodeConfig.Enabled {
		log.Info("Espresso streamer not initialized: Caff node is not enabled")
		return nil
	}

	if cfg.CaffNodeConfig.Namespace == 0 {
		log.Info("Using L2 chain ID as namespace by default")
		cfg.CaffNodeConfig.Namespace = cfg.L2ChainID.Uint64()
	}
	if cfg.CaffNodeConfig.BatchAuthenticatorAddr == (common.Address{}) {
		cfg.CaffNodeConfig.BatchAuthenticatorAddr = cfg.BatchAuthenticatorAddress
	}

	streamer, err := espresso.BatchStreamerFromCLIConfig(cfg.CaffNodeConfig.ToCLIConfig(), log, func(data []byte) (*EspressoBatch, error) {
		return UnmarshalEspressoTransaction(data)
	})
	if err != nil {
		log.Error("Failed to initialize Espresso streamer", "err", err)
		return nil
	}

	log.Info("Espresso streamer initialized", "namespace", streamer.Namespace, "hotshot polling interval", cfg.CaffNodeConfig.PollInterval, "hotshot urls", cfg.CaffNodeConfig.QueryServiceURLs)
	return streamer
}

// CaffNextBatch fetches the next batch from the Espresso streamer for the caff node.
//
// It follows the flow: Refresh() -> Update() -> Next().
//
// This is similar to the batcher's flow: espressoBatchLoadingLoop -> getSyncStatus -> refresh -> Update -> Next,
// but with a few key differences:
// - It only calls Update() when needed and everytime only calls Next() once. While the batcher calls Next() in a loop.
// - It performs additional checks, such as validating the timestamp and parent hash, which does not apply to the batcher.
func CaffNextBatch(s *op.BatchStreamer[EspressoBatch], ctx context.Context, parent eth.L2BlockRef, blockTime uint64, l1Fetcher L1Fetcher) (*SingularBatch, bool, error) {
	// Get the L1 finalized block
	finalizedL1Block, err := l1Fetcher.L1BlockRefByLabel(ctx, eth.Finalized)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get the L1 finalized block: %w", err)
	}
	// Refresh the sync status
	if err := s.Refresh(ctx, finalizedL1Block, parent.Number, parent.L1Origin); err != nil {
		return nil, false, fmt.Errorf("failed to refresh Espresso streamer: %w", err)
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
			s.Log.Error(logmodule.DroppingBatch, "batch", espressoBatch.Number(), "timestamp", batch.Timestamp, "expected", nextTimestamp)
			return nil, false, ErrTemporary
		}

		// dependent on above timestamp check. If the timestamp is correct, then it must build on top of the safe head.
		if batch.ParentHash != parent.Hash {
			s.Log.Error("ignoring batch with mismatching parent hash", "current_safe_head", parent.Hash)
			return nil, false, ErrTemporary
		}
	}
	// For caff node, when we get a batch, we assign concluding to true to drive progress
	concluding := true
	return batch, concluding, nil
}
