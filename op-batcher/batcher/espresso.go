package batcher

import (
	// #cgo darwin,arm64 LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
	"C"
	"encoding/json"
	"fmt"
	"time"

	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	espressoVerification "github.com/EspressoSystems/espresso-network-go/verification"
	"github.com/ethereum/go-ethereum/log"
)
import (
	"bytes"
	"cmp"
	"context"
	"errors"
	"math/big"
	"slices"
	"sync"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network-go/light-client"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type EspressoBatch struct {
	BlockNum uint64
	Batch    derive.SingularBatch
}

// TODO: Pull out to be re-used in op-node for derivation from Espresso
type Transaction struct {
	Namespace        uint64
	BatcherSignature []byte
	// RLP-encoded EspressoBatch
	Batch []byte
}

// Parameters for transaction fetching loop, which waits for transactions
// to be sequenced on Espresso
const (
	transactionFetchTimeout  = 2 * time.Minute
	transactionFetchInterval = 100 * time.Millisecond
)

// Parameters for finality checking loop, which waits for merkle proof for
// Espresso transaction to be available from Light Client contract
const (
	finalityTimeout       = 2 * time.Minute
	finalityCheckInterval = 100 * time.Millisecond
)

func (t Transaction) toEspresso() espressoCommon.Transaction {
	payload := append(t.BatcherSignature, t.Batch...)
	return espressoCommon.Transaction{
		Namespace: t.Namespace,
		Payload:   payload,
	}
}

func UnmarshalEspressoTransaction(data []byte, batcherAddress common.Address) (EspressoBatch, error) {
	signatureData, batchData := data[:crypto.SignatureLength], data[crypto.SignatureLength:]
	batchHash := crypto.Keccak256(batchData)

	signer, err := crypto.SigToPub(batchHash, signatureData)
	if err != nil {
		return EspressoBatch{}, err
	}
	if crypto.PubkeyToAddress(*signer) != batcherAddress {
		return EspressoBatch{}, errors.New("invalid signer")
	}
	if !crypto.VerifySignature(crypto.FromECDSAPub(signer), batchHash, signatureData) {
		// TODO: ???
		// return EspressoBatch{}, errors.New("invalid signature")
	}

	var batch EspressoBatch
	if err := rlp.DecodeBytes(batchData, &batch); err != nil {
		return EspressoBatch{}, err
	}

	return batch, nil
}

type EspressoStreamer struct {
	// Batch number we're to give out next
	batchPos          uint64
	confirmedBatchPos uint64
	// HotShot block we're to fetch next
	hotShotPos uint64

	l1Client            *L1Client
	EndpointProvider    *dial.L2EndpointProvider
	Espresso            *espressoClient.Client
	EspressoLightClient *espressoLightClient.LightClientReader

	bs *BatchSubmitter

	// Maintained in sorted order, but may be missing batches if we receive
	// any out of order.
	batchBuffer []EspressoBatch
}

// Handle both L1 reorgs and batcher restarts by updating our state in case it is
// not consistent with what's on the L1. Returns true if the state was updated.
func (s *EspressoStreamer) Refresh(ctx context.Context, syncStatus *eth.SyncStatus) (bool, error) {
	hotshotState, err := s.EspressoLightClient.LightClient.
		FinalizedState(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(syncStatus.SafeL2.L1Origin.Number)})
	if err != nil {
		return true, err
	}

	s.confirmedBatchPos = syncStatus.SafeL2.Number
	s.batchPos = s.confirmedBatchPos + 1
	s.hotShotPos = hotshotState.BlockHeight
	s.batchBuffer = nil
	return true, nil
}

func (s *EspressoStreamer) Update(ctx context.Context) error {
	// Fetch more batches from HotShot if available.
	hotShotHeight, err := s.Espresso.FetchLatestBlockHeight(ctx)

	namespace := s.bs.RollupConfig.L2ChainID.Uint64()

	if err != nil {
		return fmt.Errorf("failed to fetch HotShot block height: %w", err)
	}

	for ; s.hotShotPos < hotShotHeight; s.hotShotPos += 1 {
		s.bs.Log.Info("Fetching HS block", "blockNr", s.hotShotPos)

		txns, err := s.Espresso.FetchTransactionsInBlock(ctx, s.hotShotPos, namespace)
		if err != nil {
			return fmt.Errorf("failed to fetch transactions in block: %w", err)
		}

		if len(txns.Transactions) == 0 {
			// No transactions in block
			s.bs.Log.Info("No transactions in HS block")
			continue
		}

		header, err := s.Espresso.FetchHeaderByHeight(ctx, s.hotShotPos)
		if err != nil {
			return fmt.Errorf("failed to fetch HotShot header: %w", err)
		}

		namespaceOk := espressoVerification.VerifyNamespace(
			namespace,
			txns.Proof,
			*header.Header.GetPayloadCommitment(),
			*header.Header.GetNsTable(),
			txns.Transactions,
			txns.VidCommon,
		)

		if !namespaceOk {
			// TODO: Fails
			// return fmt.Errorf("namespace verification failed")
		}

		for _, transaction := range txns.Transactions {
			batch, err := UnmarshalEspressoTransaction(transaction, s.bs.BatcherAddress)
			if err != nil {
				s.bs.Log.Info("Failed to unmarshal espresso transaction", "error", err)
				continue
			}

			if batch.BlockNum < s.batchPos {
				// Batch already buffered/finalized
				continue
			}

			if uint64(batch.Batch.EpochNum) > header.Header.GetL1Head() {
				// Enforce that we only deal with finalized deposits
				s.bs.Log.Info("Unfinalized???")
				continue
			}

			// Find a slot to insert the batch
			i, batchRecorded := slices.BinarySearchFunc(s.batchBuffer, batch, func(x, y EspressoBatch) int {
				return cmp.Compare(x.BlockNum, y.BlockNum)
			})

			if batchRecorded {
				// Duplicate batch found, skip it
				continue
			}

			s.batchBuffer = slices.Insert(s.batchBuffer, i, batch)
		}
	}

	return nil
}

func (s *EspressoStreamer) Next(ctx context.Context) *derive.SingularBatch {
	// Is the next batch available?
	if len(s.batchBuffer) > 0 && s.batchBuffer[0].BlockNum == s.batchPos {
		var batch EspressoBatch
		batch, s.batchBuffer = s.batchBuffer[0], s.batchBuffer[1:]
		s.batchPos += 1
		return &batch.Batch
	} else {
		return nil
	}
}

func (l *BatchSubmitter) waitForFinality(height uint64, rawHeader json.RawMessage, header *espressoCommon.HeaderImpl) error {
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

	// Verify the namespace proof
	log.Info("Verifying namespace proof", "height", height)
	resp, err := l.Espresso.FetchTransactionsInBlock(l.shutdownCtx, height, 42)
	if err != nil {
		return fmt.Errorf("failed to fetch the transactions in block")
	}

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

func (l *BatchSubmitter) tryPublishBatchToEspresso(transaction espressoCommon.Transaction) error {
	txHash, err := l.Espresso.SubmitTransaction(l.shutdownCtx, transaction)
	if err != nil {
		l.Log.Error("Failed to submit transaction", "transaction", transaction, "error", err)
		return fmt.Errorf("failed to submit transaction: %w", err)
	}

	timer := time.NewTimer(transactionFetchTimeout)
	defer timer.Stop()

	ticker := time.NewTicker(transactionFetchInterval)
	defer ticker.Stop()

Loop:
	for {
		select {
		case <-ticker.C:
			_, err = l.Espresso.FetchTransactionByHash(l.shutdownCtx, txHash)
			if err == nil {
				break Loop
			}
			l.Log.Warn("Retry fetching transaction by hash", "txHash", txHash, "error", err)
		case <-timer.C:
			l.Log.Error("Failed to fetch transaction by hash after multiple attempts", "txHash", txHash)
			return fmt.Errorf("failed to fetch transaction by hash: %w", err)
		}
	}

	return nil
}

func (l *BatchSubmitter) queueBlockToEspreso(block *types.Block) error {
	batch, _, err := derive.BlockToSingularBatch(l.RollupConfig, block)
	if err != nil {
		return fmt.Errorf("failed to derive batch from block: %w", err)
	}

	buf := new(bytes.Buffer)
	err = rlp.Encode(buf, EspressoBatch{
		BlockNum: block.NumberU64(),
		Batch:    *batch,
	})
	if err != nil {
		return fmt.Errorf("failed to encode batch: %w", err)
	}

	batcherSignature, err := l.ChainSigner.Sign(l.shutdownCtx, l.BatcherAddress, crypto.Keccak256(buf.Bytes()))

	if err != nil {
		return fmt.Errorf("failed to create batcher signature: %w", err)
	}

	transaction := Transaction{
		Namespace:        l.RollupConfig.L2ChainID.Uint64(),
		BatcherSignature: batcherSignature,
		Batch:            buf.Bytes(),
	}.toEspresso()

	go func() {
		for {
			err := l.tryPublishBatchToEspresso(transaction)
			if err == nil {
				l.Log.Info(fmt.Sprintf("Published block %s to Espresso", eth.ToBlockID(block)))
				break
			}
		}
	}()

	return nil
}

// Deposit transactions obviously aren't recovered from the batch, so this doesn't return
// the original block, but we don't care for batcher purposes,as this incomplete block will be
// converted back to batch later on anyway. This double-conversion is done to avoid extensive
// modifications to channel manager that would be needed to allow it to accept batches directly
func batchToIncompleteBlock(batch *derive.SingularBatch) (*types.Block, error) {
	txs := []*types.Transaction{}
	for i, opaqueTx := range batch.Transactions {
		var tx types.Transaction
		err := tx.UnmarshalBinary(opaqueTx)
		if err != nil {
			return nil, fmt.Errorf("could not decode tx %d: %w", i, err)
		}
		txs = append(txs, &tx)
	}
	return types.NewBlockWithHeader(&types.Header{
		ParentHash: batch.ParentHash,
		Time:       batch.Timestamp,
	}).WithBody(types.Body{
		Transactions: txs,
	}), nil
}

func (l *BatchSubmitter) espressoBatchLoadingLoop(ctx context.Context, wg *sync.WaitGroup, publishSignal chan struct{}) {
	l.Log.Info("Starting EspressoBatchLoadingLoop")

	defer wg.Done()
	ticker := time.NewTicker(l.Config.PollInterval)
	defer ticker.Stop()

	streamer := EspressoStreamer{
		l1Client:            &l.L1Client,
		EndpointProvider:    &l.EndpointProvider,
		Espresso:            l.Espresso,
		EspressoLightClient: l.EspressoLightClient,
		bs:                  l,
	}

	newSyncStatus, err := l.getSyncStatus(ctx)
	if err != nil {
		l.Log.Error("failed to refresh sync status", "err", err)
		return
	}
	streamer.Refresh(ctx, newSyncStatus)

	l.Log.Info("Looping...")
	for {
		select {
		case <-ticker.C:
			newSyncStatus, err := l.getSyncStatus(ctx)
			if err != nil {
				l.Log.Error("failed to refresh sync status", "err", err)
				continue
			}

			l.Log.Info("Syncing and and pruning...")
			func() {
				l.channelMgrMutex.Lock()
				defer l.channelMgrMutex.Unlock()
				syncActions, outOfSync := computeSyncActions(*newSyncStatus, l.prevCurrentL1, l.channelMgr.blocks, l.channelMgr.channelQueue, l.Log, l.Config.PreferLocalSafeL2)
				if outOfSync {
					l.Log.Warn("Sequencer is out of sync, retrying next tick.")
					return
				}
				l.prevCurrentL1 = newSyncStatus.CurrentL1
				if syncActions.clearState != nil {
					streamer.Refresh(ctx, newSyncStatus)
					l.channelMgr.Clear(*syncActions.clearState)
				} else {
					l.channelMgr.PruneSafeBlocks(syncActions.blocksToPrune)
					l.channelMgr.PruneChannels(syncActions.channelsToPrune)
				}
			}()
			l.Log.Info("Synced and pruned")

			l.Log.Info("Updating streamer...")
			err = streamer.Update(ctx)
			l.Log.Info("Updated!..")
			if err != nil {
				l.Log.Error("failed to update Espresso streamer", "err", err)
				continue
			}

			var batch *derive.SingularBatch
			for {
				l.Log.Info("Next batch...")
				batch = streamer.Next(ctx)
				if batch == nil {
					l.Log.Info("Empty!..")
					break
				}
				l.Log.Info("Some!")

				block, err := batchToIncompleteBlock(batch)
				if err != nil {
					l.Log.Error("failed to convert singular batch to block", "err", err)
					continue
				}

				l.channelMgrMutex.Lock()
				err = l.channelMgr.AddL2Block(block)
				l.channelMgrMutex.Unlock()

				if err != nil {
					l.Log.Error("failed to add L2 block to channel manager", "err", err)
				}

				l.Log.Info("Added L2 block to channel manager")
			}
			trySignal(publishSignal)

		case <-ctx.Done():
			l.Log.Info("espressoLoop returning")
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

		err = l.batcher.queueBlockToEspreso(block)
		if err != nil {
			continue
		}

		l.lastQueuedBlock = &blockRef
	}
}

// blockLoadingLoop
// -  polls the sequencer,
// -  loads unsafe blocks from the sequencer
func (l *BatchSubmitter) espressoBlockLoadingLoop(ctx context.Context, wg *sync.WaitGroup) {
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

// loadBlockIntoState fetches & stores a single block into `state`. It returns the block it loaded.
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
