package batcher

import (
	"fmt"
	"time"

	"context"
	"math/big"
	"sync"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
)

// espressoSubmitTransactionJob is a struct that holds the state required to
// submit a transaction to Espresso.
// It contains the transaction to be submitted itself, and a number to
// track the total number of attempts to submit this transaction to Espresso.
type espressoSubmitTransactionJob struct {
	attempts    int
	transaction *espressoCommon.Transaction
}

// espressoSubmitTransactionJobResponse is a struct that holds the
// response from the Espresso client after submitting a transaction.
// It contains the job that was submitted, the hash of the transaction
// that was submitted (if successful), and any error that occurred during the
// submission (if unsuccessful).
type espressoSubmitTransactionJobResponse struct {
	job  espressoSubmitTransactionJob
	hash *espressoCommon.TaggedBase64
	err  error
}

// espressoTransactionJobAttempt is a struct that holds the job and
// response channel for a transaction submission job.
//
// This is the unit of work that is submitted to the worker to process
// for transaction submissions.
type espressoTransactionJobAttempt struct {
	job  espressoSubmitTransactionJob
	resp chan espressoSubmitTransactionJobResponse
}

// espressoVerifyReceiptJob is a struct that holds the state required to
// verify a receipt for a transaction that was submitted to Espresso.
// It contains the transaction that was submitted, the hash of the
// transaction, and the number of attempts to verify the receipt.
type espressoVerifyReceiptJob struct {
	attempts    int
	start       time.Time
	transaction espressoSubmitTransactionJob
	hash        *espressoCommon.TaggedBase64
}

// espressoVerifyReceiptJobResponse is a struct that holds the
// response from the Espresso client after verifying a receipt.
// It contains the job that was submitted, and any error that occurred
// during the verification (if unsuccessful).
type espressoVerifyReceiptJobResponse struct {
	job espressoVerifyReceiptJob
	err error
}

// espressoVerifyReceiptJobAttempt is a struct that holds the job and
// response channel for a receipt verification job.
//
// This is the unit of work that is submitted to the worker to process
// for receipt verifications.
type espressoVerifyReceiptJobAttempt struct {
	job  espressoVerifyReceiptJob
	resp chan espressoVerifyReceiptJobResponse
}

// espressoTransactionSubmitter is a struct that holds the state that governs
// the worker queue processing details for submitting transactions to Espresso
// without spawning arbitrarily many goroutines.
type espressoTransactionSubmitter struct {
	ctx                      context.Context
	wg                       *sync.WaitGroup
	submitJobQueue           chan espressoSubmitTransactionJob
	submitRespQueue          chan espressoSubmitTransactionJobResponse
	submitWorkerQueue        chan chan espressoTransactionJobAttempt
	verifyReceiptJobQueue    chan espressoVerifyReceiptJob
	verifyReceiptRespQueue   chan espressoVerifyReceiptJobResponse
	verifyReceiptWorkerQueue chan chan espressoVerifyReceiptJobAttempt
	espresso                 espressoClient.EspressoClient
}

// EspressoTransactionSubmitterConfig is a configuration struct for the
// EspressoTransactionSubmitter. It contains the configurable details for
// creating the EspressoTransactionSubmitter.
type EspressoTransactionSubmitterConfig struct {
	Ctx                                context.Context
	EspressoClient                     espressoClient.EspressoClient
	Wg                                 *sync.WaitGroup
	SubmitJobQueueCapacity             int
	SubmitResponseQueueCapacity        int
	VerifyReceiptJobQueueCapacity      int
	VerifyReceiptResponseQueueCapacity int
}

// EspressoTransactionSubmitterOption is a function that can be used to
// configure the EspressoTransactionSubmitterConfig.
type EspressoTransactionSubmitterOption func(*EspressoTransactionSubmitterConfig)

// WithEspressoClient is an option that can be used to set the Espresso client
// for the EspressoTransactionSubmitterConfig.
func WithContext(ctx context.Context) EspressoTransactionSubmitterOption {
	return func(config *EspressoTransactionSubmitterConfig) {
		config.Ctx = ctx
	}
}

// WithEspressoClient is an option that can be used to set the Espresso client
// for the EspressoTransactionSubmitterConfig.
func WithEspressoClient(client espressoClient.EspressoClient) EspressoTransactionSubmitterOption {
	return func(config *EspressoTransactionSubmitterConfig) {
		config.EspressoClient = client
	}
}

// WithWaitGroup is an option that can be used to set the wait group
// for the EspressoTransactionSubmitterConfig.
func WithWaitGroup(wg *sync.WaitGroup) EspressoTransactionSubmitterOption {
	return func(config *EspressoTransactionSubmitterConfig) {
		config.Wg = wg
	}
}

// NewEspressoTransactionSubmitter creates a new EspressoTransactionSubmitter
// with the given context and espresso client.  It will create a new
func NewEspressoTransactionSubmitter(options ...EspressoTransactionSubmitterOption) *espressoTransactionSubmitter {
	config := EspressoTransactionSubmitterConfig{
		Ctx:                                context.Background(),
		Wg:                                 new(sync.WaitGroup),
		SubmitJobQueueCapacity:             1024,
		SubmitResponseQueueCapacity:        10,
		VerifyReceiptJobQueueCapacity:      1024,
		VerifyReceiptResponseQueueCapacity: 10,
	}

	for _, option := range options {
		option(&config)
	}

	if config.EspressoClient == nil {
		panic("Espresso client is required")
	}

	return &espressoTransactionSubmitter{
		ctx:                      config.Ctx,
		wg:                       config.Wg,
		submitJobQueue:           make(chan espressoSubmitTransactionJob, config.SubmitJobQueueCapacity),
		submitRespQueue:          make(chan espressoSubmitTransactionJobResponse, config.SubmitResponseQueueCapacity),
		submitWorkerQueue:        make(chan chan espressoTransactionJobAttempt),
		verifyReceiptJobQueue:    make(chan espressoVerifyReceiptJob, config.VerifyReceiptJobQueueCapacity),
		verifyReceiptRespQueue:   make(chan espressoVerifyReceiptJobResponse, config.VerifyReceiptResponseQueueCapacity),
		verifyReceiptWorkerQueue: make(chan chan espressoVerifyReceiptJobAttempt),
		espresso:                 config.EspressoClient,
	}

}

// SubmitTransaction will submit a transaction to the Job queue.
//
// NOTE: This submits to a channel, and as a result, if the channel is full,
// it will block execution until the channel is able to accept the job.
// If the channel is buffered with sufficient space, it should not cause
// any blocking issues.
func (s *espressoTransactionSubmitter) SubmitTransaction(job *espressoCommon.Transaction) {
	s.submitJobQueue <- espressoSubmitTransactionJob{
		transaction: job,
	}
}

// handleTransactionSubmitJobResponse is a function that is meant to be run in a
// goroutine.
//
// It handles the responses from the submit transaction jobs.  It will
// determine if the transaction was successfully submitted to Espresso, and
// if not, it will retry the transaction.  If the transaction was successfully
// submitted, it will then submit a job to the verify receipt job queue to
// verify the receipt of the transaction.
func (s *espressoTransactionSubmitter) handleTransactionSubmitJobResponse() {
	for {
		var jobResp espressoSubmitTransactionJobResponse
		var ok bool

		select {
		case <-s.ctx.Done():
			return
		case jobResp, ok = <-s.submitRespQueue:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		// TODO: Evaluate the specific error type, and determine if we
		// should retry
		if jobResp.err != nil {
			s.submitJobQueue <- jobResp.job
			continue
		}

		verifyJob := espressoVerifyReceiptJob{
			start:       time.Now(),
			transaction: jobResp.job,
			hash:        jobResp.hash,
		}

		select {
		case <-s.ctx.Done():
			return
		// Move to verifying the receipt
		case s.verifyReceiptJobQueue <- verifyJob:
		}
	}
}

// VERIFY_RECEIPT_TIMEOUT is the amount of time we will wait for a receipt to
// be verified before we requeue the job for another attempt.
const VERIFY_RECEIPT_TIMEOUT = 4 * time.Second

// VERIFY_RECEIPT_RETRY_DELAY is the amount of time we will wait before
// retrying a job that failed to verify the receipt.
const VERIFY_RECEIPT_RETRY_DELAY = 100 * time.Millisecond

// handleVerifyReceiptJobResponse is a function that is meant to be run in a
// goroutine.
//
// This function handles responses from the verify receipt job queue.  It will
// check the results for any errors, and if there are any errors that are
// applicable to retry, it will requeue the job for another attempt.
// If the the job is successful, no further processing is needed and it is
// considered complete.
// If the job has taken too long to verify, then it will re-submit the job
// back to the submit transaction queue for another attempt.
//
// NOTE: This function currently will loop forever if the transaction is
// never going to be available.
//
// TODO: we need to put some sensible limits on the number of times we will
// retry a job, depending on the type of the error we received.
func (s *espressoTransactionSubmitter) handleVerifyReceiptJobResponse() {
	for {
		var jobResp espressoVerifyReceiptJobResponse
		var ok bool

		select {
		case <-s.ctx.Done():
			return
		case jobResp, ok = <-s.verifyReceiptRespQueue:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		// TODO: Evaluate the specific error type, and determine if we
		// should retry
		if jobResp.err != nil {

			// Let's check our timeout
			if have := time.Now(); have.Sub(jobResp.job.start) > VERIFY_RECEIPT_TIMEOUT {
				// We were not able to verify the receipt in time.  So we will
				// degrade this transaction back to the submit transaction phase
				// and try again.
				submitJob := jobResp.job.transaction
				select {
				case <-s.ctx.Done():
					return
				case s.submitJobQueue <- submitJob:
				}

				continue
			}

			s.verifyReceiptJobQueue <- jobResp.job
			continue
		}

		// We're done with this job and transaction, we have successfully
		// confirmed that the transaction was submitted to Espresso
	}
}

// scheduleSubmitTransactionJobs is a function that is meant to be run in a
// goroutine.
//
// It handles the scheduling of submit transaction jobs so that the submit
// transaction workers can process them.
func (s *espressoTransactionSubmitter) scheduleSubmitTransactionJobs() {
	for {
		var ok bool

		// Get a worker from the worker queue
		var worker chan espressoTransactionJobAttempt
		select {
		case <-s.ctx.Done():
			return

		case worker, ok = <-s.submitWorkerQueue:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		// Get a job from the job queue
		var job espressoSubmitTransactionJob
		select {
		case <-s.ctx.Done():
			return
		case job, ok = <-s.submitJobQueue:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		// Submit the job to the worker
		select {
		case <-s.ctx.Done():
			return

		case worker <- espressoTransactionJobAttempt{job: job, resp: s.submitRespQueue}:
		}
	}
}

// scheduleVerifyReceiptJobs is a function that is meant to be run in a
// goroutine.
//
// It handles the scheduling of verify receipt jobs so that the verify receipt
// workers can process them.
func (s *espressoTransactionSubmitter) scheduleVerifyReceiptsJobs() {
	for {
		var ok bool

		// Get a worker from the worker queue
		var worker chan espressoVerifyReceiptJobAttempt
		select {
		case <-s.ctx.Done():
			return

		case worker, ok = <-s.verifyReceiptWorkerQueue:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		// Get a job from the job queue
		var job espressoVerifyReceiptJob
		select {
		case <-s.ctx.Done():
			return
		case job, ok = <-s.verifyReceiptJobQueue:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		// Submit the job to the worker
		select {
		case <-s.ctx.Done():
			return

		case worker <- espressoVerifyReceiptJobAttempt{job: job, resp: s.verifyReceiptRespQueue}:
		}
	}
}

// espressoSubmitTransactionWorker is a function that is meant to be run as a
// goroutine.  It will create a channel for it's job queue, and submit those to
// the worker queue in order to wait for work.  It will then take that job and
// attempt to submit the transaction contained within to espresso using the
// given espresso client. It will submit the response back to the channel
// contained within the job attempt it received.
//
// It's lifetime is governed by the context passed to it, and it will stop
// processing when that context is cancelled.
//
// NOTE: If the context is cancelled after a job has been received, but before
// it is able to submit the transaction, or report about it's result, the job
// may be lost.
func espressoSubmitTransactionWorker(
	ctx context.Context,
	wg *sync.WaitGroup,
	cli espressoClient.EspressoClient,
	workerQueue chan<- chan espressoTransactionJobAttempt,
) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer wg.Done()
	ch := make(chan espressoTransactionJobAttempt)
	defer close(ch)

	for {
		var ok bool
		select {
		case <-ctx.Done():
			return

			// Queue our job queue, asking for work
		case workerQueue <- ch:
		}

		// Wait for a job to run
		var jobAttempt espressoTransactionJobAttempt
		select {
		case <-ctx.Done():
			return
		case jobAttempt, ok = <-ch:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		// Submit the transaction to Espresso
		hash, err := cli.SubmitTransaction(ctx, *jobAttempt.job.transaction)

		jobAttempt.job.attempts++
		resp := espressoSubmitTransactionJobResponse{
			job:  jobAttempt.job,
			hash: hash,
			err:  err,
		}

		if jobAttempt.job.attempts > 3 {
			// We have submitted this job quite a bit.
		}

		select {
		case <-ctx.Done():
			return

		// Send the response back via the channel in the job attempt struct
		case jobAttempt.resp <- resp:
		}
	}
}

// espressoVerifyTransactionWorker is a function that is meant to be run as a
// goroutine.  It will create a channel for it's job queue, and submit those to
// the worker queue in order to wait for work.  It will then take that job and
// attempt to verify the transaction contained within to espresso using the
// given espresso client. It will submit the response back to the channel
// contained within the job attempt it received.
func espressoVerifyTransactionWorker(
	ctx context.Context,
	wg *sync.WaitGroup,
	cli espressoClient.EspressoClient,
	workerQueue chan<- chan espressoVerifyReceiptJobAttempt,
) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer wg.Done()
	ch := make(chan espressoVerifyReceiptJobAttempt)
	defer close(ch)

	for {
		var ok bool
		select {
		case <-ctx.Done():
			return

			// Queue our job queue, asking for work
		case workerQueue <- ch:
		}

		// Wait for a job to run
		var jobAttempt espressoVerifyReceiptJobAttempt
		select {
		case <-ctx.Done():
			return
		case jobAttempt, ok = <-ch:
			if !ok {
				// Our channel is closed, and we are done
				return
			}
		}

		if jobAttempt.job.attempts > 0 {
			// We have already attempted this job, so we will wait a bit
			// NOTE: this prevents this worker from being able to process
			// other jobs while we wait for this delay.
			time.Sleep(VERIFY_RECEIPT_RETRY_DELAY)
		}

		_, err := cli.FetchTransactionByHash(ctx, jobAttempt.job.hash)

		jobAttempt.job.attempts++
		resp := espressoVerifyReceiptJobResponse{
			job: jobAttempt.job,
			err: err,
		}

		select {
		case <-ctx.Done():
			return

		case jobAttempt.resp <- resp:
		}
	}
}

// SpawnWorkers spawns the given number of workers to process the
// submit transaction jobs and verify receipt jobs.
func (s *espressoTransactionSubmitter) SpawnWorkers(numSubmitTransactionWorkers, numVerifyReceiptWorkers int) {
	workersCtx := s.ctx

	for i := 0; i < numSubmitTransactionWorkers; i++ {
		s.wg.Add(1)
		go espressoSubmitTransactionWorker(workersCtx, s.wg, s.espresso, s.submitWorkerQueue)
	}

	for i := 0; i < numVerifyReceiptWorkers; i++ {
		s.wg.Add(1)
		go espressoVerifyTransactionWorker(workersCtx, s.wg, s.espresso, s.verifyReceiptWorkerQueue)
	}
}

func (s *espressoTransactionSubmitter) Start() {
	// Submit Transaction Jobs
	go s.scheduleSubmitTransactionJobs()
	go s.handleTransactionSubmitJobResponse()

	// Verify Receipt Jobs
	go s.scheduleVerifyReceiptsJobs()
	go s.handleVerifyReceiptJobResponse()
}

// Converts a block to an EspressoBatch and starts a goroutine that publishes it to Espresso
// Returns error only if batch conversion fails, otherwise it is infallible, as the goroutine
// will retry publishing until successful.
func (l *BatchSubmitter) queueBlockToEspresso(ctx context.Context, block *types.Block) error {
	espressoBatch, err := derive.BlockToEspressoBatch(l.RollupConfig, block)
	if err != nil {
		l.Log.Warn("Failed to derive batch from block", "err", err)
		return fmt.Errorf("failed to derive batch from block: %w", err)
	}

	transaction, err := espressoBatch.ToEspressoTransaction(ctx, l.RollupConfig.L2ChainID.Uint64(), l.ChainSigner)
	if err != nil {
		l.Log.Warn("Failed to create Espresso transaction from a batch", "err", err)
		return fmt.Errorf("failed to create Espresso transaction from a batch: %w", err)
	}

	l.submitter.SubmitTransaction(transaction)

	return nil
}

func (l *BatchSubmitter) espressoSyncAndRefresh(ctx context.Context, newSyncStatus *eth.SyncStatus) {
	shouldClearState, err := l.streamer.Refresh(ctx, newSyncStatus.FinalizedL1, newSyncStatus.SafeL2.Number, newSyncStatus.SafeL2.L1Origin)
	shouldClearState = shouldClearState || err != nil

	l.channelMgrMutex.Lock()
	defer l.channelMgrMutex.Unlock()
	syncActions, outOfSync := computeSyncActions(*newSyncStatus, l.prevCurrentL1, l.channelMgr.blocks, l.channelMgr.channelQueue, l.Log, l.Config.PreferLocalSafeL2)
	if outOfSync {
		l.Log.Warn("Sequencer is out of sync, retrying next tick.")
		return
	}
	l.prevCurrentL1 = newSyncStatus.CurrentL1
	if syncActions.clearState == nil && shouldClearState {
		l.channelMgr.Clear(newSyncStatus.SafeL2.L1Origin)
		l.streamer.Reset()
	} else if syncActions.clearState != nil {
		l.channelMgr.Clear(*syncActions.clearState)
		l.streamer.Reset()
	} else {
		l.channelMgr.PruneSafeBlocks(syncActions.blocksToPrune)
		l.channelMgr.PruneChannels(syncActions.channelsToPrune)
	}
}

// AdaptL1BlockRefClient is a wrapper around eth.L1BlockRef that implements the espresso.L1Client interface
type AdaptL1BlockRefClient struct {
	L1Client L1Client
}

// NewAdaptL1BlockRefClient creates a new L1BlockRefClient
func NewAdaptL1BlockRefClient(L1Client L1Client) *AdaptL1BlockRefClient {
	return &AdaptL1BlockRefClient{
		L1Client: L1Client,
	}
}

// HeaderHashByNumber implements the espresso.L1Client interface
func (c *AdaptL1BlockRefClient) HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error) {
	expectedL1BlockRef, err := c.L1Client.HeaderByNumber(ctx, number)
	if err != nil {
		return common.Hash{}, err
	}

	return expectedL1BlockRef.Hash(), nil
}

// Periodically refreshes the sync status and polls Espresso streamer for new batches
func (l *BatchSubmitter) espressoBatchLoadingLoop(ctx context.Context, wg *sync.WaitGroup, publishSignal chan struct{}) {
	l.Log.Info("Starting EspressoBatchLoadingLoop")

	defer wg.Done()
	ticker := time.NewTicker(l.Config.EspressoPollInterval)
	defer ticker.Stop()
	defer close(publishSignal)

	for {
		select {
		case <-ticker.C:
			newSyncStatus, err := l.getSyncStatus(ctx)
			if err != nil {
				l.Log.Error("failed to refresh sync status", "err", err)
				continue
			}

			l.espressoSyncAndRefresh(ctx, newSyncStatus)

			err = l.streamer.Update(ctx)
			remainingListLen := len(l.streamer.RemainingBatches)
			if remainingListLen > 0 {
				l.Log.Warn("Remaining list not empty.", "Number items", remainingListLen)
			}

			var batch *derive.EspressoBatch

			for {

				batch = l.streamer.Next(ctx)

				if batch == nil {
					break
				}

				// This should happen ONLY if the batch is malformed. ToBlock has to guarantee no
				// transient errors.
				block, err := batch.ToBlock(l.RollupConfig)
				if err != nil {
					l.Log.Error("failed to convert singular batch to block", "err", err)
					continue
				}

				l.Log.Trace(
					"Received block from Espresso",
					"blockNr", block.NumberU64(),
					"blockHash", block.Hash(),
					"parentHash", block.ParentHash(),
				)

				l.channelMgrMutex.Lock()
				err = l.channelMgr.AddL2Block(block)
				l.channelMgrMutex.Unlock()

				if err != nil {
					l.Log.Error("failed to add L2 block to channel manager", "err", err)
					l.clearState(ctx)
					l.streamer.Reset()
				}

				l.Log.Info("Added L2 block to channel manager")
			}

			trySignal(publishSignal)

			// A failure in the streamer Update can happen after the buffer has been partially filled
			if err != nil {
				l.Log.Error("failed to update Espresso streamer", "err", err)
				continue
			}

		case <-ctx.Done():
			l.Log.Info("espressoBatchLoadingLoop returning")
			return
		}
	}
}

type BlockLoader struct {
	queuedBlocks   []eth.L2BlockRef
	prevSyncStatus *eth.SyncStatus
	batcher        *BatchSubmitter
}

func (l *BlockLoader) reset(ctx context.Context) {
	l.prevSyncStatus = nil
	l.queuedBlocks = nil
	l.batcher.clearState(ctx)
	l.batcher.safeL1Origin(ctx)
}

func (l *BlockLoader) EnqueueBlocks(ctx context.Context, blocksToQueue inclusiveBlockRange) {
	l.batcher.Log.Info("Loading and queueing blocks", "range", blocksToQueue)
	for i := blocksToQueue.start; i <= blocksToQueue.end; i++ {
		block, err := l.batcher.fetchBlock(ctx, i)
		for _, txn := range block.Transactions() {
			l.batcher.Log.Info("tx hash before submitting to Espresso", "hash", txn.Hash().String())
		}

		if err != nil {
			l.batcher.Log.Warn("Failed to fetch block", "err", err)
			break
		}

		if len(l.queuedBlocks) > 0 && block.ParentHash() != l.queuedBlocks[len(l.queuedBlocks)-1].Hash {
			l.batcher.Log.Warn("Found L2 reorg", "block_number", i)
			l.reset(ctx)
			break
		}

		blockRef, err := derive.L2BlockToBlockRef(l.batcher.RollupConfig, block)
		if err != nil {
			continue
		}

		err = l.batcher.queueBlockToEspresso(ctx, block)
		if err != nil {
			continue
		}

		l.queuedBlocks = append(l.queuedBlocks, blockRef)
	}
}

type EnqueueBlockAction uint

const (
	ActionEnqueue = iota
	ActionRetry
	ActionReset
)

// This function is an analogue of `computeSyncActions` for Espresso batcher mode
//
// It computes the next block range to enqueue to Espresso based on new newSyncStatus and
// does a number of checks to ensure consistency of the chain.
//
// If reorg is detected, empty range and ActionReset is returned.
// If there isn't enough information or no blocks to load yet, empty range and ActionRetry is returned.
func (l *BlockLoader) nextBlockRange(newSyncStatus *eth.SyncStatus) (inclusiveBlockRange, EnqueueBlockAction) {
	if newSyncStatus.HeadL1 == (eth.L1BlockRef{}) {
		// empty sync status
		return inclusiveBlockRange{}, ActionRetry
	}

	if l.prevSyncStatus == nil {
		l.prevSyncStatus = newSyncStatus
	}

	if newSyncStatus.CurrentL1.Number < l.prevSyncStatus.CurrentL1.Number {
		// sequencer restarted and hasn't caught up yet
		l.batcher.Log.Warn("sequencer currentL1 reversed", "new currentL1", newSyncStatus.CurrentL1.Number, "previous currentL1", l.prevSyncStatus.CurrentL1)
		return inclusiveBlockRange{}, ActionRetry
	}

	var safeL2 eth.L2BlockRef
	if l.batcher.Config.PreferLocalSafeL2 {
		// This is preffered when running interop, but not yet enabled by default.
		safeL2 = newSyncStatus.LocalSafeL2
	} else {
		safeL2 = newSyncStatus.SafeL2
	}

	// State empty, just enqueue all unsafe blocks
	if len(l.queuedBlocks) == 0 {
		return inclusiveBlockRange{safeL2.Number + 1, newSyncStatus.UnsafeL2.Number}, ActionEnqueue
	}

	lastQueuedBlock := l.queuedBlocks[len(l.queuedBlocks)-1]
	firstQueuedBlock := l.queuedBlocks[0]
	nextSafeBlockNum := safeL2.Number + 1

	if lastQueuedBlock.Number >= newSyncStatus.UnsafeL2.Number {
		// nothing to enqueue, unsafe block number is not higher than safe
		return inclusiveBlockRange{}, ActionRetry
	}

	if lastQueuedBlock.Number < safeL2.Number {
		// derivation pipeline is somehow ahead of us, reset
		return inclusiveBlockRange{}, ActionReset
	}

	if nextSafeBlockNum < firstQueuedBlock.Number {
		l.batcher.Log.Warn("next safe block is below oldest block in state")
		return inclusiveBlockRange{}, ActionReset
	}

	numBlocksToEnqueue := nextSafeBlockNum - firstQueuedBlock.Number

	if numBlocksToEnqueue > uint64(len(l.queuedBlocks)) {
		l.batcher.Log.Warn("safe head above newest block in state, resetting loader")
		return inclusiveBlockRange{}, ActionReset
	}

	if numBlocksToEnqueue > 0 && l.queuedBlocks[numBlocksToEnqueue-1].Hash != safeL2.Hash {
		l.batcher.Log.Warn("safe chain reorg, resetting loader")
		return inclusiveBlockRange{}, ActionReset
	}

	if newSyncStatus.UnsafeL2.Number <= lastQueuedBlock.Number+1 {
		return inclusiveBlockRange{}, ActionRetry
	}

	if safeL2.Number > firstQueuedBlock.Number {
		numFinalizedBlocks := safeL2.Number - firstQueuedBlock.Number
		l.batcher.Log.Warn(
			"Removing finalized blocks from queued",
			"numFinalizedBlocks", numFinalizedBlocks,
			"safeL2", safeL2,
			"firstQueuedBlock", firstQueuedBlock)
		l.queuedBlocks = l.queuedBlocks[numFinalizedBlocks:]
	}

	return inclusiveBlockRange{lastQueuedBlock.Number + 1, newSyncStatus.UnsafeL2.Number}, ActionEnqueue
}

// blockLoadingLoop
// -  polls the sequencer,
// -  queues unsafe blocks from the sequencer to Espresso
func (l *BatchSubmitter) espressoBatchQueueingLoop(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(l.Config.PollInterval)
	defer ticker.Stop()
	defer wg.Done()

	var loader = BlockLoader{
		batcher: l,
	}

	for {
		select {
		case <-ticker.C:
			newSyncStatus, err := l.getSyncStatus(ctx)

			if err != nil {
				l.Log.Error("Couldn't get sync status", "error", err)
				continue
			}

			blocksToQueue, action := loader.nextBlockRange(newSyncStatus)

			if action == ActionEnqueue {
				loader.EnqueueBlocks(ctx, blocksToQueue)
			} else if action == ActionReset {
				loader.reset(ctx)
			}

		case <-ctx.Done():
			l.Log.Info("blockLoadingLoop returning")
			return
		}
	}
}

func (l *BatchSubmitter) fetchBlock(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	l2Client, err := l.EndpointProvider.EthClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting L2 client: %w", err)
	}

	cCtx, cancel := context.WithTimeout(ctx, l.Config.NetworkTimeout)
	defer cancel()

	block, err := l2Client.BlockByNumber(cCtx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("getting L2 block: %w", err)
	}

	return block, nil
}

func (l *BatchSubmitter) registerBatcher(ctx context.Context) error {
	if l.Attestation == nil {
		l.Log.Warn("Attestation is nil, skipping registration")
		return nil
	}

	batchAuthenticator, err := bindings.NewBatchAuthenticator(l.RollupConfig.BatchAuthenticatorAddress, l.L1Client)
	if err != nil {
		return fmt.Errorf("failed to create batch authenticator contract bindings: %w", err)
	}

	// Decode the attestation off-chain to conserve gas
	attestationTbs, signature, err := batchAuthenticator.DecodeAttestationTbs(&bind.CallOpts{}, l.Attestation)
	if err != nil {
		return fmt.Errorf("failed to decode attestation: %w", err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(l.Config.BatcherPrivateKey, l.RollupConfig.L1ChainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	// Submit decoded attestation to batch inbox contract
	tx, err := batchAuthenticator.RegisterSigner(txOpts, attestationTbs, signature)
	if err != nil {
		return fmt.Errorf("failed to create RegisterSigner transaction: %w", err)
	}

	candidate := txmgr.TxCandidate{
		TxData: tx.Data(),
		To:     tx.To(),
	}

	_, err = l.Txmgr.Send(ctx, candidate)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	l.Log.Info("Registered batcher with the batch inbox contract")

	return nil
}

// sendEspressoTx uses the txmgr queue to send the given transaction candidate after setting its
// gaslimit. It will block if the txmgr queue has reached its MaxPendingTransactions limit.
func (l *BatchSubmitter) sendEspressoTx(txdata txData, isCancel bool, candidate *txmgr.TxCandidate, queue TxSender[txRef], receiptsCh chan txmgr.TxReceipt[txRef]) {
	transactionReference := txRef{id: txdata.ID(), isCancel: isCancel, isBlob: txdata.daType == DaTypeBlob}
	l.Log.Debug("Sending Espresso-enabled L1 transaction", "txRef", transactionReference)

	var commitment [32]byte
	if len(candidate.Blobs) == 0 {
		commitment = crypto.Keccak256Hash(candidate.TxData)
		l.Log.Debug("Hashing calldata transaction", "txRef", transactionReference, "commitment", hexutil.Encode(commitment[:]))
	} else {
		contactenatedBlobHashes := make([]byte, 0)
		for _, blob := range candidate.Blobs {
			blobCommitment, err := blob.ComputeKZGCommitment()
			if err != nil {
				receiptsCh <- txmgr.TxReceipt[txRef]{
					ID:  transactionReference,
					Err: fmt.Errorf("failed to compute KZG commitment for blob: %w", err),
				}
				return
			}
			blobHash := eth.KZGToVersionedHash(blobCommitment)
			contactenatedBlobHashes = append(contactenatedBlobHashes, blobHash.Bytes()...)
		}
		commitment = crypto.Keccak256Hash(contactenatedBlobHashes)
		l.Log.Debug("Hashing blob transaction", "txRef", transactionReference, "commitment", hexutil.Encode(commitment[:]))
	}

	signature, err := crypto.Sign(commitment[:], l.Config.BatcherPrivateKey)
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to sign transaction: %w", err),
		}
		return
	}
	l.Log.Debug("Signed transaction", "txRef", transactionReference, "commitment", hexutil.Encode(commitment[:]), "sig", hexutil.Encode(signature))

	batchAuthenticatorAbi, err := bindings.BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to get batch authenticator ABI: %w", err),
		}
		return
	}

	authenticateBatchCalldata, err := batchAuthenticatorAbi.Pack("authenticateBatch", commitment, signature)
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to pack authenticateBatch calldata: %w", err),
		}
		return
	}

	verifyCandidate := txmgr.TxCandidate{
		TxData: authenticateBatchCalldata,
		To:     &l.RollupConfig.BatchAuthenticatorAddress,
	}

	l.Log.Debug(
		"Sending authenticateBatch transaction",
		"txRef", transactionReference,
		"commitment", hexutil.Encode(commitment[:]),
		"sig", hexutil.Encode(signature),
		"address", l.RollupConfig.BatchAuthenticatorAddress.String(),
	)
	_, err = l.Txmgr.Send(l.killCtx, verifyCandidate)
	if err != nil {
		l.Log.Error("Failed to send authenticateBatch transaction", "txRef", transactionReference, "err", err)
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to send authenticateBatch transaction: %w", err),
		}
		return
	}

	l.Log.Debug("Queueing transaction", "txRef", transactionReference)
	queue.Send(transactionReference, *candidate, receiptsCh)
}
