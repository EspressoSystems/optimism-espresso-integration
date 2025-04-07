package batcher

import (
	"bytes"
	"fmt"
	"time"

	"context"
	"encoding/json"
	"errors"
	"math/big"
	"sync"

	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	espressoVerification "github.com/EspressoSystems/espresso-network-go/verification"
	"github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// Parameters for transaction fetching loop, which waits for transactions
// to be sequenced on Espresso
const (
	transactionFetchTimeout  = 10 * time.Second
	transactionFetchInterval = 100 * time.Millisecond
)

// Parameters for finality checking loop, which waits for merkle proof for
// Espresso transaction to be available from Light Client contract
const (
	finalityTimeout       = 20 * time.Minute
	finalityCheckInterval = 100 * time.Millisecond
)

func (l *BatchSubmitter) tryPublishBatchToEspresso(ctx context.Context, transaction espressoCommon.Transaction) error {
	txHash, err := l.Espresso.SubmitTransaction(ctx, transaction)
	if err != nil {
		l.Log.Error("Failed to submit transaction", "transaction", transaction, "error", err)
		return fmt.Errorf("failed to submit transaction: %w", err)
	}

	timer := time.NewTimer(transactionFetchTimeout)
	defer timer.Stop()

	ticker := time.NewTicker(transactionFetchInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_, err := l.Espresso.FetchTransactionByHash(ctx, txHash)
			if err == nil {
				return nil
				//return l.waitForFinality(submittedTxInfo)
			}
		case <-timer.C:
			l.Log.Error("Failed to fetch transaction by hash after multiple attempts", "txHash", txHash)
			return fmt.Errorf("failed to fetch transaction by hash: %w", err)
		case <-ctx.Done():
			l.Log.Info("Cancelling transaction publishing", "txHash", txHash)
			return nil
		}
	}
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

	go func() {
		// We will retry publishing until successful
		for {
			err := l.tryPublishBatchToEspresso(ctx, *transaction)
			if err == nil {
				l.Log.Info(fmt.Sprintf("Published block %s to Espresso", eth.ToBlockID(block)))
				break
			}
		}
	}()

	return nil
}

func (l *BatchSubmitter) espressoSyncAndRefresh(ctx context.Context, newSyncStatus *eth.SyncStatus, streamer *espresso.EspressoStreamer[derive.EspressoBatch]) {
	shouldClearState, err := streamer.Refresh(ctx, newSyncStatus)
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
		streamer.Reset()
	} else if syncActions.clearState != nil {
		l.channelMgr.Clear(*syncActions.clearState)
		streamer.Reset()
	} else {
		l.channelMgr.PruneSafeBlocks(syncActions.blocksToPrune)
		l.channelMgr.PruneChannels(syncActions.channelsToPrune)
	}
}

// Periodically refreshes the sync status and polls Espresso streamer for new batches
func (l *BatchSubmitter) espressoBatchLoadingLoop(ctx context.Context, wg *sync.WaitGroup, publishSignal chan struct{}) {
	l.Log.Info("Starting EspressoBatchLoadingLoop")

	defer wg.Done()
	ticker := time.NewTicker(l.Config.PollInterval)
	defer ticker.Stop()
	defer close(publishSignal)

	streamer := espresso.NewEspressoStreamer(
		l.RollupConfig.L2ChainID.Uint64(),
		l.L1Client,
		l.Espresso,
		l.EspressoLightClient, // TODO (Keyao) BatchSubmitter doesn't have field EspressoLightClient.
		l.Log,
		func(data []byte) (*derive.EspressoBatch, error) {
			return derive.UnmarshalEspressoTransaction(data, l.SequencerAddress)
		},
		2*time.Second,
	)

	for {
		select {
		case <-ticker.C:
			newSyncStatus, err := l.getSyncStatus(ctx)
			if err != nil {
				l.Log.Error("failed to refresh sync status", "err", err)
				continue
			}

			l.espressoSyncAndRefresh(ctx, newSyncStatus, &streamer)

			err = streamer.Update(ctx)
			if err != nil {
				l.Log.Error("failed to update Espresso streamer", "err", err)
				continue
			}

			var batch *derive.EspressoBatch

			for {

				batch = streamer.Next(ctx)

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
					streamer.Reset()
				}

				l.Log.Info("Added L2 block to channel manager")
			}
			trySignal(publishSignal)

		case <-ctx.Done():
			l.Log.Info("espressoBatchLoadingLoop returning")
			return
		}
	}
}

type BlockLoader struct {
	prevSyncStatus  *eth.SyncStatus
	lastQueuedBlock *eth.L2BlockRef
	batcher         *BatchSubmitter
}

func (l *BlockLoader) Reset(ctx context.Context) {
	l.prevSyncStatus = nil
	l.lastQueuedBlock = nil
	l.batcher.clearState(ctx)
}

func (l *BlockLoader) EnqueueBlocks(ctx context.Context, blocksToQueue inclusiveBlockRange) {
	for i := blocksToQueue.start; i <= blocksToQueue.end; i++ {
		block, err := l.batcher.fetchBlock(ctx, i)
		if errors.Is(err, ErrReorg) {
			l.batcher.Log.Warn("Found L2 reorg", "block_number", i)
			l.Reset(ctx)
			break
		} else if err != nil {
			l.batcher.Log.Warn("Failed to fetch block", "err", err)
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

		l.lastQueuedBlock = &blockRef
	}
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

			if newSyncStatus.HeadL1 == (eth.L1BlockRef{}) {
				// empty sync status
				continue
			}

			if loader.prevSyncStatus == nil {
				loader.prevSyncStatus = newSyncStatus
			}

			if newSyncStatus.CurrentL1.Number < loader.prevSyncStatus.CurrentL1.Number {
				// sequencer restarted and hasn't caught up yet
				continue
			}

			var safeL2 eth.L2BlockRef
			if l.Config.PreferLocalSafeL2 {
				// This is preffered when running interop, but not yet enabled by default.
				safeL2 = newSyncStatus.LocalSafeL2
			} else {
				safeL2 = newSyncStatus.SafeL2
			}

			if loader.lastQueuedBlock == nil {
				loader.lastQueuedBlock = &safeL2
			}

			if loader.lastQueuedBlock.Number >= newSyncStatus.UnsafeL2.Number {
				// nothing to enqueue, unsafe block number is not higher than safe
				continue
			}

			if loader.lastQueuedBlock.Number < safeL2.Number {
				// derivation pipeline is somehow ahead of us, reset
				loader.Reset(ctx)
				continue
			}

			blocksToQueue := inclusiveBlockRange{loader.lastQueuedBlock.Number + 1, newSyncStatus.UnsafeL2.Number}

			loader.EnqueueBlocks(ctx, blocksToQueue)

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

func (l *BatchSubmitter) waitForFinality(submittedTxInfo espressoCommon.TransactionQueryData) error {
	rawHeader, err := l.Espresso.FetchRawHeaderByHeight(l.shutdownCtx, submittedTxInfo.BlockHeight)
	if err != nil {
		return err
	}

	var header espressoCommon.HeaderImpl
	err = json.Unmarshal(rawHeader, &header)
	if err != nil {
		return fmt.Errorf("could not unmarshal header from bytes")
	}

	height := header.Header.GetBlockHeight()

	timer := time.NewTimer(finalityTimeout)
	defer timer.Stop()

	ticker := time.NewTicker(finalityCheckInterval)
	defer ticker.Stop()

	var snapshot espressoCommon.BlockMerkleSnapshot

Loop:
	for {
		select {
		case <-ticker.C:
			res, err := l.EspressoLightClient.FetchMerkleRoot(height, nil)
			if err == nil {
				snapshot = res
				break Loop
			}
		case <-timer.C:
			return fmt.Errorf("failed to fetch merkle root")
		}
	}

	if snapshot.Height <= height {
		return fmt.Errorf("snapshot height is less than or equal to the requested height")
	}

	nextHeader, err := l.Espresso.FetchHeaderByHeight(l.shutdownCtx, snapshot.Height)
	if err != nil {
		return fmt.Errorf("error fetching the snapshot header (height: %d): %w", snapshot.Height, err)
	}

	proof, err := l.Espresso.FetchBlockMerkleProof(l.shutdownCtx, snapshot.Height, height)
	if err != nil {
		return fmt.Errorf("error fetching merkle proof")
	}

	blockMerkleTreeRoot := nextHeader.Header.GetBlockMerkleTreeRoot()

	log.Info("Verifying merkle proof", "height", height)
	ok := espressoVerification.VerifyMerkleProof(proof.Proof, rawHeader, *blockMerkleTreeRoot, snapshot.Root)
	if !ok {
		return fmt.Errorf("error validating merkle proof (height: %d, snapshot height: %d)", height, snapshot.Height)
	}

	log.Info("Verifying namespace proof", "height", height)
	resp, err := l.Espresso.FetchTransactionsInBlock(l.shutdownCtx, height, l.RollupConfig.L2ChainID.Uint64())
	if err != nil {
		return fmt.Errorf("failed to fetch the transactions in block")
	}

	txIncluded := false
	for _, includedTx := range resp.Transactions {
		if bytes.Equal(submittedTxInfo.Transaction.Payload, includedTx) {
			txIncluded = true
			break
		}
	}

	if !txIncluded {
		return fmt.Errorf("transaction not included in block")
	}

	// Verify the namespace proof
	namespaceOk := espressoVerification.VerifyNamespace(
		l.RollupConfig.L2ChainID.Uint64(),
		resp.Proof,
		*header.Header.GetPayloadCommitment(),
		*header.Header.GetNsTable(),
		resp.Transactions,
		resp.VidCommon,
	)

	if !namespaceOk {
		return fmt.Errorf("error validating namespace proof (height: %d)", height)
	}

	return nil
}
