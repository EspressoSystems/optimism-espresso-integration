package espresso

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/EspressoSystems/espresso-network/sdks/go/types"
	espressoCommon "github.com/EspressoSystems/espresso-network/sdks/go/types"
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
	FetchNamespaceTransactionsInRange(ctx context.Context, fromHeight uint64, toHeight uint64, namespace uint64) ([]types.NamespaceTransactionsRangeData, error)
}

// L1Client is an interface that documents the methods we utilize for
// the L1 client.
type L1Client interface {
	HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error)
}

// espresso-network go sdk's HeaderInterface currently lacks a function to get this info,
// although it is present in all header versions
func GetFinalizedL1(header *espressoCommon.HeaderImpl) espressoCommon.L1BlockInfo {
	v0_1, ok := header.Header.(*espressoCommon.Header0_1)
	if ok {
		return *v0_1.L1Finalized
	}
	v0_2, ok := header.Header.(*espressoCommon.Header0_2)
	if ok {
		return *v0_2.L1Finalized
	}
	v0_3, ok := header.Header.(*espressoCommon.Header0_3)
	if ok {
		return *v0_3.L1Finalized
	}
	panic("Unsupported header version")
}

type BatchStreamer[B Batch] struct {
	// Namespace of the rollup we're interested in
	Namespace uint64

	L1Client                      L1Client
	RollupL1Client                L1Client
	EspressoClient                EspressoClient
	EspressoLightClient           LightClientCallerInterface
	Log                           log.Logger
	PollingHotShotPollingInterval time.Duration

	// Batch number we're to give out next
	BatchPos uint64
	// Position of the last safe batch, we can use it as the position to fallback when resetting
	fallbackBatchPos uint64
	// HotShot block that was visited last
	hotShotPos uint64
	// HotShot position that we can fallback to, guaranteeing not to skip any unsafe batches
	fallbackHotShotPos uint64
	// HotShot position we start reading from, exclusive
	originHotShotPos uint64
	// Latest finalized block on the L1.
	FinalizedL1 eth.L1BlockRef

	// Maintained in sorted order, but may be missing batches if we receive
	// any out of order.
	BatchBuffer BatchBuffer[B]

	// Manage the batches which origin is unfinalized
	RemainingBatches map[common.Hash]B

	unmarshalBatch func([]byte) (*B, error)
}

// Compile time assertion to ensure EspressoStreamer implements
// EspressoStreamerIFace
var _ EspressoStreamer[Batch] = (*BatchStreamer[Batch])(nil)

func NewEspressoStreamer[B Batch](
	namespace uint64,
	l1Client L1Client,
	rollupL1Client L1Client,
	espressoClient EspressoClient,
	lightClient LightClientCallerInterface,
	log log.Logger,
	unmarshalBatch func([]byte) (*B, error),
	pollingHotShotPollingInterval time.Duration,
	originHotShotPos uint64,
	originBatchPos uint64,
) *BatchStreamer[B] {
	return &BatchStreamer[B]{
		L1Client:            l1Client,
		RollupL1Client:      rollupL1Client,
		EspressoClient:      espressoClient,
		EspressoLightClient: lightClient,
		Log:                 log,
		Namespace:           namespace,
		// Internally, BatchPos is the position of the batch we are to give out next, hence the +1
		BatchPos:                      originBatchPos + 1,
		fallbackBatchPos:              originBatchPos + 1,
		BatchBuffer:                   NewBatchBuffer[B](),
		PollingHotShotPollingInterval: pollingHotShotPollingInterval,
		RemainingBatches:              make(map[common.Hash]B),
		unmarshalBatch:                unmarshalBatch,
		originHotShotPos:              originHotShotPos,
		fallbackHotShotPos:            originHotShotPos,
		hotShotPos:                    originHotShotPos,
	}
}

// Reset the state to the last safe batch
func (s *BatchStreamer[B]) Reset() {
	s.Log.Info("reset espresso streamer", "hotshot pos", s.fallbackHotShotPos, "batch pos", s.fallbackBatchPos)
	s.hotShotPos = s.fallbackHotShotPos
	s.BatchPos = s.fallbackBatchPos + 1
	s.BatchBuffer.Clear()
}

// RefreshSafeL1Origin is a convenience method that allows us to update the
// safe L1 origin of the Streamer. It will confirm the Espresso Block Height
// and reset the state if necessary.
func (s *BatchStreamer[B]) RefreshSafeL1Origin(safeL1Origin eth.BlockID) {
	shouldReset := s.confirmEspressoBlockHeight(safeL1Origin)
	if shouldReset {
		s.Reset()
	}
}

// Update streamer state based on L1 and L2 sync status
func (s *BatchStreamer[B]) Refresh(ctx context.Context, finalizedL1 eth.L1BlockRef, safeBatchNumber uint64, safeL1Origin eth.BlockID) error {
	s.FinalizedL1 = finalizedL1

	s.RefreshSafeL1Origin(safeL1Origin)

	// NOTE: be sure to update s.finalizedL1 before checking this condition and returning
	if s.fallbackBatchPos == safeBatchNumber {
		// This means everything is in sync, no state update needed
		return nil
	}

	shouldReset := safeBatchNumber < s.fallbackBatchPos

	// We should jump ahead if fallback position is higher than what we're currently reading from
	shouldReset = shouldReset && (s.fallbackBatchPos > s.hotShotPos)

	s.fallbackBatchPos = safeBatchNumber
	if shouldReset {
		s.Reset()
	}
	return nil
}

// CheckBatch checks the validity of the given batch against the finalized L1
// block and the safe L1 origin.
func (s *BatchStreamer[B]) CheckBatch(ctx context.Context, batch B) (BatchValidity, int) {

	// Make sure the finalized L1 block is initialized before checking the block number.
	if s.FinalizedL1 == (eth.L1BlockRef{}) {
		s.Log.Error("Finalized L1 block not initialized")
		return BatchUndecided, 0
	}
	origin := (batch).L1Origin()

	if origin.Number > s.FinalizedL1.Number {
		// Signal to resync to wait for the L1 finality.
		s.Log.Warn("L1 origin not finalized, pending resync", "finalized L1 block number", s.FinalizedL1.Number, "origin number", origin.Number)
		return BatchUndecided, 0
	}

	l1headerHash, err := s.RollupL1Client.HeaderHashByNumber(ctx, new(big.Int).SetUint64(origin.Number))
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

// HOTSHOT_BLOCK_STREAM_LIMIT is the maximum number of blocks to attempt to
// load from Espresso in a single process using streaming API.
// This helps to limit our block polling to a limited number of blocks within
// a single batched attempt.
const HOTSHOT_BLOCK_STREAM_LIMIT = 500

// HOTSHOT_BLOCK_FETCH_LIMIT is the maximum number of blocks to attempt to
// load from Espresso in a single process using fetch API.
// This helps to limit our block polling to a limited number of blocks within
// a single batched attempt.
const HOTSHOT_BLOCK_FETCH_LIMIT = 100

// computeEspressoBlockHeightsRange computes the range of block heights to fetch
// from Espresso. It starts from the last processed block and goes up to
// `limit` blocks ahead or the current block height, whichever
// is smaller.
func (s *BatchStreamer[B]) computeEspressoBlockHeightsRange(currentBlockHeight uint64, limit uint64) (start uint64, finish uint64) {
	start = s.hotShotPos
	if start > 0 {
		// We've already processed the block in hotShotPos.  In order to avoid
		// reprocessing the same block, we want to start from the next block.
		start++
	}
	// `FetchNamespaceTransactionsInRange` is exclusive to finish, so we add 1 to currentBlockHeight
	finish = min(start+limit, currentBlockHeight+1)

	return start, finish
}

// Update will update the `EspressoStreamer“ by attempting to ensure that the
// next call to the `Next` method will return a `Batch`.
//
// It attempts to ensure the existence of a next batch, provided no errors
// occur when communicating with HotShot, by processing Blocks retrieved from
// `HotShot` in discreet batches. If each processing of a batch of blocks will
// not yield a new `Batch`, then it will continue to process the next batch
// of blocks from HotShot until it runs out of blocks to process.
//
//	NOTE: this method is best effort.  It is unable to guarantee that the
//	next call to `Next` will return a batch.  However, the only things
//	that will prevent the next call to `Next` from returning a batch is if
//	there are no more HotShot blocks to process currently, or if an error
//	occurs when communicating with HotShot.
func (s *BatchStreamer[B]) Update(ctx context.Context) error {
	// Retrieve the current block height from Espresso.  We grab this reference
	// so we don't have to keep fetching it in a loop, and it informs us of
	// the current block height available to process.
	currentBlockHeight, err := s.EspressoClient.FetchLatestBlockHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch latest block height: %w", err)
	}

	// Fetch API implementation
	for i := 0; ; i++ {
		// Fetch more batches from HotShot if available.
		start, finish := s.computeEspressoBlockHeightsRange(currentBlockHeight, HOTSHOT_BLOCK_FETCH_LIMIT)

		if start >= finish || (start+1 == finish && i > 0) {
			// If start is one less than our finish, then that means we
			// already processed all of the blocks available to us.  We
			// should break out of the loop.  Sadly, this means that we
			// likely do not have any batches to process.
			//
			// NOTE: this also likely means that the following is true:
			// start + 1 == finish  == currentBlockHeight + 1
			//
			// NOTE: there is an edge case here if the only block available is
			// the initial block of Espresso, then we get stuck in a loop
			// repeatedly processing it again and again.   So to catch
			// this case, we check to see if start is equal to finish, after
			// an initial iteration.
			break
		}

		// Process the remaining batches
		s.processRemainingBatches(ctx)

		s.Log.Debug("Fetching hotshot blocks", "from", start, "upTo", finish)

		// Process the new batches fetched from Espresso
		if err := s.fetchHotShotRange(ctx, start, finish); err != nil {
			return fmt.Errorf("failed to process hotshot range: %w", err)
		}

		if s.HasNext(ctx) {
			// If we have a batch ready to be processed, we can exit the loop,
			// otherwise, we will want to continue to the next range of blocks
			// to fetch.
			//
			// The goal here is to try and provide our best effort to ensure
			// that we have the next batch available for processing.  We should
			// only fail to do this if there currently is no next batch
			// currently available (or if we error while attempting to retrieve
			// transactions from HotShot).
			break
		}
	}

	return nil
}

// fetchHotShotRange is a helper method that will load all of the blocks from
// Hotshot from start to finish, inclusive. It will process each block and
// update the batch buffer with any batches found in the block.
// It will also update the hotShotPos to the last block processed, in order
// to effectively keep track of the last block we have successfully fetched,
// and therefore processed from Hotshot.
func (s *BatchStreamer[B]) fetchHotShotRange(ctx context.Context, start, finish uint64) error {
	// Process the new batches fetched from Espresso
	s.Log.Trace("Fetching HotShot block range", "start", start, "finish", finish)

	// FetchNamespaceTransactionsInRange fetches transactions in [start, finish)
	namespaceRangeTransactions, err := s.EspressoClient.FetchNamespaceTransactionsInRange(ctx, start, finish, s.Namespace)
	if err != nil {
		return err
	}

	s.Log.Info("Fetched HotShot block range", "start", start, "finish", finish, "numNamespaceTransactions", len(namespaceRangeTransactions))

	// We want to keep track of the latest block we have processed.
	// This is essential for ensuring we don't unnecessarily keep
	// refetching the same blocks that we have already processed.
	// This should ensure that we keep moving forward and consuming
	// from the Espresso Blocks without missing any blocks.
	s.hotShotPos = finish - 1
	if len(namespaceRangeTransactions) == 0 {
		s.Log.Trace("No transactions in hotshot block range", "start", start, "finish", finish)
	}

	for _, namespaceTransaction := range namespaceRangeTransactions {
		for _, txn := range namespaceTransaction.Transactions {
			s.processEspressoTransaction(ctx, txn.Payload)
		}
	}

	return nil
}

// processRemainingBatches is a helper method that checks the remaining batches
// and prunes or adds them to the batch buffer as appropriate.
func (s *BatchStreamer[B]) processRemainingBatches(ctx context.Context) {
	// Collect keys to delete, without modifying the batch list during iteration.
	var keysToDelete []common.Hash

	// Process the remaining batches
	for k, batch := range s.RemainingBatches {
		validity, pos := s.CheckBatch(ctx, batch)

		switch validity {
		case BatchDrop:
			s.Log.Warn("Dropping batch", "batch", batch)
			keysToDelete = append(keysToDelete, k)
			continue

		case BatchPast:
			s.Log.Warn("Batch already processed. Skipping", "batch", batch)
			keysToDelete = append(keysToDelete, k)
			continue

		case BatchUndecided:
			s.Log.Warn("Batch is still undecided, keeping it in the remaining list", "batch", batch)
			continue

		case BatchAccept:
			s.Log.Info("Remaining list", "Recovered batch, inserting batch", batch)

		case BatchFuture:
			// The function CheckBatch is not expected to return BatchFuture so if we enter this case there is a problem.
			s.Log.Error("Remaining list", "BatchFuture validity not expected for batch", batch)
			continue
		}

		header := batch.Header()
		s.Log.Trace("Remaining list", "Inserting batch into buffer",
			"parentHash", header.ParentHash,
			"epochNum", header.Number,
			"timestamp", header.Time)
		s.BatchBuffer.Insert(batch, pos)
		keysToDelete = append(keysToDelete, k)
	}

	// Delete keys all at once.
	for _, k := range keysToDelete {
		delete(s.RemainingBatches, k)
	}
}

// processEspressoTransaction is a helper method that encapsulates the logic of
// processing batches from the transactions in a block fetched from Espresso.
func (s *BatchStreamer[B]) processEspressoTransaction(ctx context.Context, transaction espressoCommon.Bytes) {
	batch, err := s.UnmarshalBatch(transaction)
	if err != nil {
		s.Log.Warn("Dropping batch with invalid transaction data", "error", err)
		return
	}

	validity, pos := s.CheckBatch(ctx, *batch)

	switch validity {
	case BatchDrop:
		s.Log.Info("Dropping batch", batch)

	case BatchPast:
		s.Log.Info("Batch already processed. Skipping", "batch", (*batch).Number())

	case BatchUndecided:
		hash := (*batch).Hash()
		if existingBatch, ok := s.RemainingBatches[hash]; ok {
			s.Log.Warn("Batch already in buffer", "batch", existingBatch)
		}
		s.RemainingBatches[hash] = *batch

	case BatchAccept:
		header := (*batch).Header()
		s.Log.Info("Inserting batch into buffer",
			"parentHash", header.ParentHash,
			"epochNum", header.Number,
			"timestamp", header.Time)
		s.BatchBuffer.Insert(*batch, pos)

	case BatchFuture:
		// The function CheckBatch is not expected to return BatchFuture so if we enter this case there is a problem.
		s.Log.Error("Remaining list", "BatchFuture validity not expected for batch", batch)
	}

}

// UnmarshalBatch implements EspressoStreamerIFace
func (s *BatchStreamer[B]) Next(ctx context.Context) *B {
	// Is the next batch available?
	if s.HasNext(ctx) {
		// Current batch is going to be processed, update fallback batch position
		s.BatchPos += 1
		return s.BatchBuffer.Pop()
	}

	return nil
}

// HasNext implements EspressoStreamerIFace
func (s *BatchStreamer[B]) HasNext(ctx context.Context) bool {
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
//
// We do not propagate the error if Light Client is unreachable - this is not an essential
// operation and streamer can continue operation
func (s *BatchStreamer[B]) confirmEspressoBlockHeight(safeL1Origin eth.BlockID) (shouldReset bool) {
	shouldReset = false
	if s.EspressoLightClient == nil {
		s.Log.Warn("Espresso light client is not initialized")
		return false
	}

	hotshotState, err := s.EspressoLightClient.
		FinalizedState(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(safeL1Origin.Number)})

	if err != nil {
		// If we have already advanced our fallback position before, there's no need to roll it back
		s.fallbackHotShotPos = max(s.fallbackHotShotPos, s.originHotShotPos)
		s.Log.Warn("failed to get finalized state from light client", "err", err)
		return false
	}

	// If hotshot block height at L1 origin is lower than our
	// hotshot origin, we never want to update our fallback
	// position to this height, or we risk dipping below
	// hotshot origin on reset.
	if hotshotState.BlockHeight <= s.originHotShotPos {
		s.Log.Info("HotShot height at L1 Origin less than HotShot origin of the streamer, ignoring")
		return shouldReset
	}

	// If we assigned to fallback position from hotsthot height before
	// and now the light client reports a smaller height, there was an L1
	// reorg and we should reset our state
	shouldReset = hotshotState.BlockHeight < s.fallbackHotShotPos

	s.fallbackHotShotPos = hotshotState.BlockHeight

	return shouldReset
}

// UnmarshalBatch implements EspressoStreamerIFace
func (s *BatchStreamer[B]) UnmarshalBatch(b []byte) (*B, error) {
	return s.unmarshalBatch(b)
}
