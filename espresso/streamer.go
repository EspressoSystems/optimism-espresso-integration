package espresso

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network-go/light-client"
	espressoTypes "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/log"
)

type L1Client interface {
	HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error)
}

// espresso-network-go's HeaderInterface currently lacks a function to get this info,
// although it is present in all header versions
func GetFinalizedL1(header *espressoTypes.HeaderImpl) espressoTypes.L1BlockInfo {
	v0_1, ok := header.Header.(*espressoTypes.Header0_1)
	if ok {
		return *v0_1.L1Finalized
	}
	v0_2, ok := header.Header.(*espressoTypes.Header0_2)
	if ok {
		return *v0_2.L1Finalized
	}
	v0_3, ok := header.Header.(*espressoTypes.Header0_3)
	if ok {
		return *v0_3.L1Finalized
	}
	panic("Unsupported header version")
}

type EspressoStreamer[B Batch] struct {
	// Namespace of the rollup we're interested in
	Namespace uint64

	L1Client                      L1Client // TODO Philippe apparently not used yet
	EspressoClient                *espressoClient.Client
	EspressoLightClient           *espressoLightClient.LightClientReader
	Log                           log.Logger
	PollingHotShotPollingInterval time.Duration

	// Batch number we're to give out next
	BatchPos uint64
	// HotShot block that was visited last
	hotShotPos uint64
	// Position of the last safe batch
	confirmedBatchPos uint64
	// Hotshot block corresponding to the last safe batch
	confirmedHotShotPos uint64
	// Latest finalized block on the L1. Used by the batcher, not initialized by the Caff node
	// until it calls `Refresh`.
	finalizedL1 eth.L1BlockRef

	// Maintained in sorted order, but may be missing batches if we receive
	// any out of order.
	BatchBuffer BatchBuffer[B]

	// Manage the batches which origin is unfinalized
	RemainingBatches map[common.Hash]B

	UnmarshalBatch func([]byte) (*B, error)
}

func NewEspressoStreamer[B Batch](
	namespace uint64,
	l1Client L1Client,
	espressoClient *espressoClient.Client,
	lightClient *espressoLightClient.LightClientReader,
	log log.Logger,
	unmarshalBatch func([]byte) (*B, error),
	pollingHotShotPollingInterval time.Duration,
) EspressoStreamer[B] {
	return EspressoStreamer[B]{
		L1Client:                      l1Client,
		EspressoClient:                espressoClient,
		EspressoLightClient:           lightClient,
		Log:                           log,
		Namespace:                     namespace,
		BatchPos:                      1,
		BatchBuffer:                   NewBatchBuffer[B](),
		PollingHotShotPollingInterval: pollingHotShotPollingInterval,
		RemainingBatches:              make(map[common.Hash]B),
		UnmarshalBatch:                unmarshalBatch,
	}
}

// Reset the state to the last safe batch
func (s *EspressoStreamer[B]) Reset() {
	s.BatchPos = s.confirmedBatchPos + 1
	s.BatchBuffer.Clear()
	s.confirmEspressoBlockHeight()
}

// Handle both L1 reorgs and batcher restarts by updating our state in case it is
// not consistent with what's on the L1. Returns true if the state was updated.
func (s *EspressoStreamer[B]) Refresh(ctx context.Context, syncStatus *eth.SyncStatus) (bool, error) {
	s.Log.Info("Refreshing streamer...")
	s.Log.Info("L2 ", "safe block number", syncStatus.SafeL2.Number)
	s.Log.Info("L1 ", "finalized block number", syncStatus.FinalizedL1.Number, "safe block number", syncStatus.SafeL1.Number)
	s.finalizedL1 = syncStatus.FinalizedL1

	// NOTE: be sure to update s.finalizedL1 before checking this condition and returning
	if s.confirmedBatchPos == syncStatus.SafeL2.Number {
		s.BatchPos = s.confirmedBatchPos + 1
		s.confirmedHotShotPos = s.hotShotPos
		return false, nil
	}

	s.confirmedBatchPos = syncStatus.SafeL2.Number
	s.Reset()
	return true, nil
}

// Sishan TODO: this refresh() is needed before CaffNextBatch, but it is not guaranteed to deal with restarting caff node
func (s *EspressoStreamer[B]) CaffRefresh(ctx context.Context, parent eth.L2BlockRef, l1Finalized func() (eth.L1BlockRef, error)) error {
	s.BatchPos = s.confirmedBatchPos + 1
	s.confirmedBatchPos = parent.Number
	s.confirmedHotShotPos = s.hotShotPos
	finalizedL1Block, err := l1Finalized()
	if err != nil {
		s.Log.Error("failed to get the L1 finalized block", "err", err)
		return err
	}
	s.finalizedL1 = finalizedL1Block
	return nil
}

func (s *EspressoStreamer[B]) CheckBatch(ctx context.Context, batch B) (BatchValidity, int) {

	// Make sure the finalized L1 block is initialized before checking the block number.
	if s.finalizedL1 == (eth.L1BlockRef{}) {
		s.Log.Error("Finalized L1 block not initialized")
		return BatchDrop, 0
	}
	origin := (batch).L1Origin()
	if origin.Number > s.finalizedL1.Number {
		// Signal to resync to wait for the L1 finality.
		s.Log.Warn("L1 origin not finalized, pending resync", "finalized L1 block number", s.finalizedL1.Number, "origin number", origin.Number)
		return BatchUndecided, 0
	}

	l1headerHash, err := s.L1Client.HeaderHashByNumber(ctx, new(big.Int).SetUint64(origin.Number))
	if err != nil {
		// Signal to resync to be able to fetch the L1 header.
		s.Log.Warn("Failed to fetch the L1 header, pending resync", "error", err)
		return BatchUndecided, 0
	} else {
		if l1headerHash != origin.Hash {
			s.Log.Warn("Dropping batch with invalid L1 origin hash", "error", err)
			return BatchDrop, 0
		}
	}

	// Find a slot to insert the batch
	i, batchRecorded := s.BatchBuffer.TryInsert(batch)

	// Batch already buffered/finalized
	if batch.Number() < s.BatchPos {
		s.Log.Warn("Batch is older than current batchPos, skipping", "batchNr", batch.Number(), "batchPos", s.BatchPos)
		return BatchPast, 0
	}

	if batchRecorded {
		// Duplicate batch found, skip it
		return BatchPast, i
	}

	return BatchAccept, i
}

func (s *EspressoStreamer[B]) computeEspressoBlockHeightsRange(ctx context.Context) (uint64, uint64, error) {
	currentBlockHeight, err := s.EspressoClient.FetchLatestBlockHeight(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to fetch HotShot block height: %w", err)
	}
	start := s.confirmedHotShotPos
	finish := min(start+100, currentBlockHeight)

	return start, finish, nil
}

// / Update the batch buffer by reading from the Espresso blocks
// / @param ctx context
// / @return error possible error
func (s *EspressoStreamer[B]) Update(ctx context.Context) error {
	// s.BatchBuffer.Mu.Lock()
	// defer s.BatchBuffer.Mu.Unlock()

	// Fetch more batches from HotShot if available.
	start, finish, err := s.computeEspressoBlockHeightsRange(ctx)
	if err != nil {
		return err
	}

	s.Log.Info("Fetching hotshot blocks", "from", start, "upTo", finish)

	i := start

	s.Log.Info("Remaining list before", "Size", len(s.RemainingBatches))

	// Process the remaining batches
	for k, batch := range s.RemainingBatches {

		validity, pos := s.CheckBatch(ctx, batch)

		switch validity {

		case BatchDrop:
			s.Log.Warn("Dropping batch", "batch", batch)
			delete(s.RemainingBatches, k)
			continue

		case BatchPast:
			s.Log.Warn("Batch already processed. Skipping", "batch", batch)
			delete(s.RemainingBatches, k)
			continue

		case BatchUndecided:
			s.Log.Warn("Batch is still undecided, keeping it in the remaining list", "batch", batch)
			continue

		case BatchAccept:
			s.Log.Info("Remaining list", "Recovered batch, inserting batch", batch)

		case BatchFuture:
			s.Log.Info("Remaining list", "Inserting batch for future processing", batch)
		}

		s.Log.Trace("Remaining list", "Inserting batch into buffer", "batch", batch)
		s.BatchBuffer.Insert(batch, pos)
		delete(s.RemainingBatches, k)

	}

	s.Log.Info("Remaining list after", "Size", len(s.RemainingBatches))

	// Process the new batches fetched from Espresso
	for ; i <= finish; i++ {
		s.Log.Trace("Fetching HotShot block", "block", i)

		txns, err := s.EspressoClient.FetchTransactionsInBlock(ctx, i, s.Namespace)
		if err != nil {
			return fmt.Errorf("Failed to fetch transactions in block: %w", err)
		}

		s.Log.Trace("Fetched HotShot block", "block", i, "txns", len(txns.Transactions))

		if len(txns.Transactions) == 0 {
			s.Log.Trace("No transactions in hotshot block", "block", i)
			continue
		}

		for _, transaction := range txns.Transactions {

			batch, err := s.UnmarshalBatch(transaction)
			if err != nil {
				s.Log.Warn("Dropping batch with invalid transaction data", "error", err)
				continue
			}

			s.Log.Info("Inserting batch into buffer", "batch", batch)

			validity, pos := s.CheckBatch(ctx, *batch)
			if pos == 0 {
				s.hotShotPos = i
			}

			switch validity {

			case BatchDrop:
				s.Log.Info("Dropping batch", batch)
				continue

			case BatchPast:
				s.Log.Info("Batch already processed. Skipping", "batch", batch)
				continue

			case BatchUndecided:
				hash := (*batch).Header().Hash()
				s.RemainingBatches[hash] = *batch
				continue

			case BatchAccept:
				s.Log.Info("Recovered batch, inserting")

			case BatchFuture:
				s.Log.Info("Inserting batch for future processing")
			}

			s.Log.Trace("Inserting batch into buffer", "batch", batch)
			s.BatchBuffer.Insert(*batch, pos)
		}

	}

	return nil
}

func (s *EspressoStreamer[B]) Start(ctx context.Context) {

	s.Log.Info("Starting espresso streamer")
	ticker := time.NewTicker(s.PollingHotShotPollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := s.Update(ctx)
			if err != nil {
				s.Log.Error("failed to update Espresso streamer", "err", err)
				continue
			}

		case <-ctx.Done():
			s.Log.Info("espressoStreamerLoop returning")
			return
		}
	}

}

// TODO this logic might be slightly different between batcher and derivation
func (s *EspressoStreamer[B]) Next(ctx context.Context) *B {
	// s.BatchBuffer.Mu.Lock()
	// defer s.BatchBuffer.Mu.Unlock()

	// Is the next batch available?
	if s.BatchBuffer.Len() > 0 && (*s.BatchBuffer.Peek()).Number() == s.BatchPos {
		s.BatchPos += 1
		return s.BatchBuffer.Pop()
	}

	return nil
}

// This function allows to "pin" the Espresso block height corresponding to the last safe batch
// Note that this function can be called
func (s *EspressoStreamer[B]) confirmEspressoBlockHeight() {
	s.confirmedHotShotPos = s.hotShotPos
}
