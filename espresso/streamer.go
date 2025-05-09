package espresso

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoTypes "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/log"
)

// Espresso light client bindings don't have an explicit name for this struct,
// so we define it here to avoid spelling it out every time
type FinalizedState = struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}

// LightClientCallerInterface is an interface that documents the methods we utilize
// for the espresso light client
//
// We define this here locally in order to effectively document the methods
// we utilize.  This approach allows us to avoid importing the entire package
// and allows us to easily swap implementations for testing.
type LightClientCallerInterface interface {
	FinalizedState(opts *bind.CallOpts) (FinalizedState, error)
}

// EspressoClient is an interface that documents the methods we utilize for
// the espressoClient.Client.
//
// As a result we are able to easily swap implementations for testing, or
// for modification / wrapping.
type EspressoClient interface {
	FetchLatestBlockHeight(ctx context.Context) (uint64, error)
	FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (espressoClient.TransactionsInBlock, error)
}

// L1Client is an interface that documents the methods we utilize for
// the L1 client.
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
	EspressoClient                EspressoClient
	EspressoLightClient           LightClientCallerInterface
	Log                           log.Logger
	PollingHotShotPollingInterval time.Duration

	// Batch number we're to give out next
	BatchPos uint64
	// HotShot block that was visited last
	hotShotPos uint64
	// Position of the last safe batch, we can use it as the position to fallback when resetting
	fallbackBatchPos uint64
	// HotShot position that we can fallback to, guaranteeing not to skip any unsafe batches
	fallbackHotShotPos uint64
	// Latest finalized block on the L1.
	FinalizedL1 eth.L1BlockRef

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
	espressoClient EspressoClient,
	lightClient LightClientCallerInterface,
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
	s.hotShotPos = s.fallbackHotShotPos
	s.BatchPos = s.fallbackBatchPos + 1
	s.BatchBuffer.Clear()
}

// Handle both L1 reorgs and batcher restarts by updating our state in case it is
// not consistent with what's on the L1. Returns true if the state was updated.
func (s *EspressoStreamer[B]) Refresh(ctx context.Context, finalizedL1 eth.L1BlockRef, safeBatchNumber uint64, safeL1Origin eth.BlockID) (bool, error) {
	s.FinalizedL1 = finalizedL1

	err := s.confirmEspressoBlockHeight(safeL1Origin)
	if err != nil {
		return false, err
	}

	// NOTE: be sure to update s.finalizedL1 before checking this condition and returning
	if s.fallbackBatchPos == safeBatchNumber {
		// This means everything is in sync, no state update needed
		return false, nil
	}

	s.fallbackBatchPos = safeBatchNumber
	s.Reset()
	return true, nil
}

func (s *EspressoStreamer[B]) CheckBatch(ctx context.Context, batch B) (BatchValidity, int) {

	// Make sure the finalized L1 block is initialized before checking the block number.
	if s.FinalizedL1 == (eth.L1BlockRef{}) {
		s.Log.Error("Finalized L1 block not initialized")
		return BatchDrop, 0
	}
	origin := (batch).L1Origin()
	if origin.Number > s.FinalizedL1.Number {
		// Signal to resync to wait for the L1 finality.
		s.Log.Warn("L1 origin not finalized, pending resync", "finalized L1 block number", s.FinalizedL1.Number, "origin number", origin.Number)
		return BatchUndecided, 0
	}

	l1headerHash, err := s.L1Client.HeaderHashByNumber(ctx, new(big.Int).SetUint64(origin.Number))
	if err != nil {
		// Signal to resync to be able to fetch the L1 header.
		s.Log.Warn("Failed to fetch the L1 header, pending resync", "error", err)
		return BatchUndecided, 0
	} else {
		if l1headerHash != origin.Hash {
			s.Log.Warn("Dropping batch with invalid L1 origin hash")
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

// HOTSHOT_BLOCK_LOAD_LIMIT is the maximum number of blocks to attempt to
// load from Espresso in a single process. This helps to limit our block
// polling to a limited number of blocks within a single batched attempt.
const HOTSHOT_BLOCK_LOAD_LIMIT = 100

// computeEspressoBlockHeightsRange computes the range of block heights to fetch
// from Espresso. It starts from the last processed block and goes up to
// HOTSHOT_BLOCK_LOAD_LIMIT blocks ahead or the current block height, whichever
// is smaller.
func (s *EspressoStreamer[B]) computeEspressoBlockHeightsRange(currentBlockHeight uint64) (start uint64, finish uint64) {
	start = s.hotShotPos
	finish = min(start+HOTSHOT_BLOCK_LOAD_LIMIT, currentBlockHeight)

	return start, finish
}

// / Update the batch buffer by reading from the Espresso blocks
// / @param ctx context
// / @return error possible error
func (s *EspressoStreamer[B]) Update(ctx context.Context) error {
	currentBlockHeight, err := s.EspressoClient.FetchLatestBlockHeight(ctx)
	if err != nil {
		return err
	}

	s.Log.Info("Remaining list before", "Size", len(s.RemainingBatches))

	// Process the remaining batches
	s.processRemainingBatches(ctx)

	s.Log.Info("Remaining list after", "Size", len(s.RemainingBatches))

	// Fetch more batches from HotShot if available.
	start, finish := s.computeEspressoBlockHeightsRange(currentBlockHeight)
	s.Log.Info("Fetching hotshot blocks", "from", start, "upTo", finish)
	i := start

	// Process the new batches fetched from Espresso
	for ; i <= finish; i++ {
		s.Log.Trace("Fetching HotShot block", "block", i)

		txns, err := s.EspressoClient.FetchTransactionsInBlock(ctx, i, s.Namespace)
		if err != nil {
			return fmt.Errorf("failed to fetch transactions in block: %w", err)
		}

		s.Log.Trace("Fetched HotShot block", "block", i, "txns", len(txns.Transactions))

		if len(txns.Transactions) == 0 {
			s.Log.Trace("No transactions in hotshot block", "block", i)
			continue
		}

		s.processEspressoTransactions(ctx, i, txns)
	}

	return nil
}

// processRemainingBatches is a helper method that checks the remaining batches
// and prunes or adds them to the batch buffer as appropriate.
func (s *EspressoStreamer[B]) processRemainingBatches(ctx context.Context) {
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
}

// processEspressoTransactions is a helper method that encapsulates the logic of
// processing batches from the transactions in a block fetched from Espresso.
func (s *EspressoStreamer[B]) processEspressoTransactions(ctx context.Context, i uint64, txns espressoClient.TransactionsInBlock) {
	for _, transaction := range txns.Transactions {
		batch, err := s.UnmarshalBatch(transaction)
		if err != nil {
			s.Log.Warn("Dropping batch with invalid transaction data", "error", err)
			continue
		}

		s.Log.Info("Inserting batch into buffer", "batch", batch)

		validity, pos := s.CheckBatch(ctx, *batch)

		switch validity {

		case BatchDrop:
			s.Log.Info("Dropping batch", batch)
			continue

		case BatchPast:
			s.Log.Info("Batch already processed. Skipping", "batch", (*batch).Number())
			continue

		case BatchUndecided:
			hash := (*batch).Hash()
			if existingBatch, ok := s.RemainingBatches[hash]; ok {
				s.Log.Warn("Batch already in buffer", "batch", existingBatch)
			}
			s.RemainingBatches[hash] = *batch
			continue

		case BatchAccept:
			s.Log.Info("Inserting accepted batch")

		case BatchFuture:
			s.Log.Info("Inserting batch for future processing")
		}

		s.Log.Trace("Inserting batch into buffer", "batch", batch)
		s.BatchBuffer.Insert(*batch, pos)
	}
	s.hotShotPos = i
}

// TODO this logic might be slightly different between batcher and derivation
func (s *EspressoStreamer[B]) Next(ctx context.Context) *B {
	// Is the next batch available?
	if s.HasNext(ctx) {
		// Current batch is going to be processed, update fallback batch position
		s.fallbackBatchPos = s.BatchPos
		s.BatchPos += 1
		return s.BatchBuffer.Pop()
	}

	return nil
}

func (s *EspressoStreamer[B]) HasNext(ctx context.Context) bool {
	if s.BatchBuffer.Len() > 0 {
		return (*s.BatchBuffer.Peek()).Number() == s.BatchPos
	}

	return false
}

// This function allows to "pin" the Espresso block height that is guaranteed not to contain
// any batches that have origin >= safeL1Origin.
// We do this by reading block height from Light Client FinalizedState at safeL1Origin.
//
// For reference on why doing this guarantees we won't skip any unsafe blocks:
// https://eng-wiki.espressosys.com/mainch30.html#:Components:espresso%20streamer:initializing%20hotshot%20height
func (s *EspressoStreamer[B]) confirmEspressoBlockHeight(safeL1Origin eth.BlockID) error {
	hotshotState, err := s.EspressoLightClient.
		FinalizedState(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(safeL1Origin.Number)})
	if errors.Is(err, bind.ErrNoCode) {
		s.fallbackHotShotPos = 0
		return nil
	} else if err != nil {
		return err
	}

	s.fallbackHotShotPos = hotshotState.BlockHeight
	return nil
}
