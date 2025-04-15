package batcher

import (
	"fmt"
	"time"

	"context"
	"errors"
	"math/big"
	"sync"

	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
)

// Parameters for transaction fetching loop, which waits for transactions
// to be sequenced on Espresso
const (
	transactionFetchTimeout  = 4 * time.Second
	transactionFetchInterval = 100 * time.Millisecond
)

// Parameters for finality checking loop, which waits for merkle proof for
// Espresso transaction to be available from Light Client contract
const (
	finalityTimeout       = 2 * time.Minute
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

	err := l.registerBatcher(ctx)
	if err != nil {
		l.Log.Error("could not register with batch inbox contract", "err", err)
		return
	}

	streamer := espresso.NewEspressoStreamer(
		l.RollupConfig.L2ChainID.Uint64(),
		l.L1Client,
		l.Espresso,
		l.EspressoLightClient,
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

func (l *BatchSubmitter) registerBatcher(ctx context.Context) error {
	if l.Attestation == nil {
		l.Log.Warn("Attestation is nil, skipping registration")
		return nil
	}

	batchAuthenticator, err := bindings.NewBatchAuthenticator(l.RollupConfig.CaffNodeConfig.BatchAuthenticatorAddress, l.L1Client)
	if err != nil {
		return fmt.Errorf("failed to create batch verifier contract bindings: %w", err)
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
			Err: fmt.Errorf("failed to get batch verifier ABI: %w", err),
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
		To:     &l.RollupConfig.CaffNodeConfig.BatchAuthenticatorAddress,
	}

	l.Log.Debug(
		"Sending authenticateBatch transaction",
		"txRef", transactionReference,
		"commitment", hexutil.Encode(commitment[:]),
		"sig", hexutil.Encode(signature),
		"address", l.RollupConfig.CaffNodeConfig.BatchAuthenticatorAddress.String(),
	)
	_, err = l.Txmgr.Send(l.killCtx, verifyCandidate)
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to send authenticateBatch transaction: %w", err),
		}
		return
	}

	l.Log.Debug("Queueing transaction", "txRef", transactionReference)
	queue.Send(transactionReference, *candidate, receiptsCh)
}
