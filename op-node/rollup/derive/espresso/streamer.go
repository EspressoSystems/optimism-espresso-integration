package espresso

import (
	// #cgo darwin,arm64 LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
	"C"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"slices"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network-go/light-client"
	espressoTypes "github.com/EspressoSystems/espresso-network-go/types"
	espressoVerification "github.com/EspressoSystems/espresso-network-go/verification"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// TODO Philippe Copied from batches.go to avoid circular dependencies

type BatchValidity uint8

const (
	// BatchDrop indicates that the batch is invalid, and will always be in the future, unless we reorg
	BatchDrop = iota
	// BatchAccept indicates that the batch is valid and should be processed
	BatchAccept
	// BatchUndecided indicates we are lacking L1 information until we can proceed batch filtering
	BatchUndecided
	// BatchFuture indicates that the batch may be valid, but cannot be processed yet and should be checked again later
	BatchFuture
	// BatchPast indicates that the batch is from the past, i.e. its timestamp is smaller or equal
	// to the safe head's timestamp.
	BatchPast
)

// espresso-network-go's HeaderInterface currently lacks a function to get this info,
// although it is present in all header versions
func getFinalizedL1(header *espressoTypes.HeaderImpl) *espressoTypes.L1BlockInfo {
	v0_1, ok := header.Header.(*espressoTypes.Header0_1)
	if ok {
		return v0_1.L1Finalized
	}
	v0_2, ok := header.Header.(*espressoTypes.Header0_2)
	if ok {
		return v0_2.L1Finalized
	}
	v0_3, ok := header.Header.(*espressoTypes.Header0_3)
	if ok {
		return v0_3.L1Finalized
	}
	return nil
}

type L1Client interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

type EspressoStreamer struct {
	// Namespace of the rollup we're interested in
	Namespace uint64
	// Address of the batcher, we expect transactions to
	// be signed by the corresponding private key
	BatcherAddress common.Address

	L1Client            L1Client
	EspressoClient      *espressoClient.Client
	EspressoLightClient *espressoLightClient.LightClientReader
	Log                 log.Logger

	// Batch number we're to give out next
	BatchPos uint64
	// HotShot block we're to fetch next
	hotShotPos uint64
	// Position of the last safe batch
	confirmedBatchPos uint64
	// Hotshot block corresponding to the last safe batch
	confirmedHotShotPos uint64

	// Maintained in sorted order, but may be missing batches if we receive
	// any out of order.
	batchBuffer []EspressoBatch
}

// Reset the state to the last safe batch
func (s *EspressoStreamer) Reset() {
	s.BatchPos = s.confirmedBatchPos + 1
	s.hotShotPos = s.confirmedHotShotPos
	s.batchBuffer = nil
}

// Handle both L1 reorgs and batcher restarts by updating our state in case it is
// not consistent with what's on the L1. Returns true if the state was updated.
func (s *EspressoStreamer) Refresh(ctx context.Context, syncStatus *eth.SyncStatus) (bool, error) {
	if s.confirmedBatchPos == syncStatus.SafeL2.Number {
		return false, nil
	}

	hotshotState, err := s.EspressoLightClient.LightClient.
		FinalizedState(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(syncStatus.SafeL2.L1Origin.Number)})
	if err != nil {
		return false, err
	}

	s.confirmedBatchPos = syncStatus.SafeL2.Number
	s.confirmedHotShotPos = hotshotState.BlockHeight
	s.Reset()
	return true, nil
}

func (s *EspressoStreamer) CheckBatch(batch EspressoBatch, espressoFinalizedL1 *espressoTypes.L1BlockInfo) (BatchValidity, int) {

	// TODO Philippe copying here verification in Sishan's Espresso streamer. As mentioned in the comment below, not sure this applies to the Caff node either. Let us discuss.
	// Sishan TODO: these checks are copy-pasted from OP's checkSingularBatch(), we should check whether these apply to caff node
	//nextTimestamp := l2SafeHead.Time + cfg.BlockTime
	//if batch.Timestamp > nextTimestamp {
	//	log.Trace("received out-of-order batch for future processing after next batch", "next_timestamp", nextTimestamp)
	//	return BatchFuture
	//}
	//if batch.Timestamp < nextTimestamp {
	//	log.Warn("dropping past batch with old timestamp", "min_timestamp", nextTimestamp)
	//	return BatchDrop
	//}
	//
	//// dependent on above timestamp check. If the timestamp is correct, then it must build on top of the safe head.
	//if batch.ParentHash != l2SafeHead.Hash {
	//	log.Warn("ignoring batch with mismatching parent hash", "current_safe_head", l2SafeHead.Hash)
	//	return BatchDrop
	//}

	if batch.Number() < s.BatchPos {
		// Batch already buffered/finalized
		s.Log.Debug("batch is older than current batchPos, skipping", "batchNr", batch.Number(), "batchPos", s.BatchPos)
		return BatchDrop, 0
	}

	if uint64(batch.Batch.EpochNum) > espressoFinalizedL1.Number {
		// Enforce that we only deal with finalized deposits
		s.Log.Warn("batch with unfinalized L1 origin",
			"batchEpochNum", batch.Batch.EpochNum, "espressoFinalizedL1Num", espressoFinalizedL1.Number,
		)
		return BatchFuture, 0
	}

	// TODO Philippe. Why do we want to ignore batches with these malformed transactions. Can't the EVM deal with these txs?
	// We can do this check earlier, but it's a more intensive one, so we do this last.
	for i, txBytes := range batch.Batch.Transactions {
		if len(txBytes) == 0 {
			log.Warn("transaction data must not be empty, but found empty tx", "tx_index", i)
			return BatchDrop, 0
		}
		if txBytes[0] == types.DepositTxType {
			log.Warn("sequencers may not embed any deposits into batch data, but found tx that has one", "tx_index", i)
			return BatchDrop, 0

		}
	}

	// Find a slot to insert the batch
	i, batchRecorded := slices.BinarySearchFunc(s.batchBuffer, batch, func(x, y EspressoBatch) int {
		return cmp.Compare(x.Number(), y.Number())
	})

	if batchRecorded {
		// Duplicate batch found, skip it
		s.Log.Debug("duplicate batch, skipping", "batchNr", batch.Number())
		return BatchDrop, 0
	}

	return BatchAccept, i

}

func (s *EspressoStreamer) Update(ctx context.Context) error {
	// Fetch more batches from HotShot if available.
	hotshotState, err := s.EspressoLightClient.LightClient.FinalizedState(&bind.CallOpts{})

	if err != nil {
		return fmt.Errorf("failed to fetch HotShot block height: %w", err)
	}

	s.Log.Debug("Updated finalized hotshot state", "hotshotState", hotshotState)

	targetHeight := min(hotshotState.BlockHeight, s.hotShotPos+100)

	for ; s.hotShotPos < targetHeight; s.hotShotPos += 1 {
		s.Log.Debug("fetching HotShot block", "blockNr", s.hotShotPos)

		txns, err := s.EspressoClient.FetchTransactionsInBlock(ctx, s.hotShotPos, s.Namespace)
		if err != nil {
			return fmt.Errorf("failed to fetch transactions in block: %w", err)
		}

		if len(txns.Transactions) == 0 {
			s.Log.Debug("no transactions in hotshot block", "blockNr", s.hotShotPos)
			continue
		}

		rawHeader, err := s.EspressoClient.FetchRawHeaderByHeight(ctx, s.hotShotPos)
		if err != nil {
			return fmt.Errorf("failed to fetch raw HotShot header: %w", err)
		}

		var header espressoTypes.HeaderImpl
		err = json.Unmarshal(rawHeader, &header)
		if err != nil {
			return fmt.Errorf("could not unmarshal header from bytes")
		}

		snapshot, err := s.EspressoLightClient.FetchMerkleRoot(s.hotShotPos, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch Merkle root: %w", err)
		}

		if snapshot.Height <= s.hotShotPos {
			return fmt.Errorf("snapshot height is less than or equal to the requested height")
		}

		nextHeader, err := s.EspressoClient.FetchHeaderByHeight(ctx, snapshot.Height)
		if err != nil {
			return fmt.Errorf("error fetching the snapshot header (height: %d): %w", snapshot.Height, err)
		}

		proof, err := s.EspressoClient.FetchBlockMerkleProof(ctx, snapshot.Height, s.hotShotPos)
		if err != nil {
			return fmt.Errorf("error fetching merkle proof")
		}

		blockMerkleTreeRoot := nextHeader.Header.GetBlockMerkleTreeRoot()

		log.Info("Verifying merkle proof", "height", s.hotShotPos)
		ok := espressoVerification.VerifyMerkleProof(proof.Proof, rawHeader, *blockMerkleTreeRoot, snapshot.Root)
		if !ok {
			return fmt.Errorf("error validating merkle proof (height: %d, snapshot height: %d)", s.hotShotPos, snapshot.Height)
		}

		namespaceOk := espressoVerification.VerifyNamespace(
			s.Namespace,
			txns.Proof,
			*header.Header.GetPayloadCommitment(),
			*header.Header.GetNsTable(),
			txns.Transactions,
			txns.VidCommon,
		)

		if !namespaceOk {
			s.Log.Error("namespace verification failed for HS block", "blockNr", s.hotShotPos)
			return fmt.Errorf("namespace verification failed")
		}

		for _, transaction := range txns.Transactions {
			batch, err := UnmarshalEspressoTransaction(transaction, s.BatcherAddress)
			if err != nil {
				s.Log.Info("Failed to unmarshal espresso transaction", "error", err)
				continue
			}

			espressoFinalizedL1 := getFinalizedL1(&header)
			if espressoFinalizedL1 == nil {
				return fmt.Errorf("unknown Espresso header version")
			}

			var action, i = s.CheckBatch(batch, espressoFinalizedL1)

			switch action {

			case BatchDrop:
				continue

			case BatchFuture:
				continue // TODO Philippe update Sishan's remaining list

			case BatchAccept:
				s.Log.Debug("recovered batch, buffering", "batchnr", batch.Number())
				s.batchBuffer = slices.Insert(s.batchBuffer, i, batch)

			}

		}
	}

	return nil
}

func (s *EspressoStreamer) Next(ctx context.Context) *EspressoBatch {
	// Is the next batch available?
	if len(s.batchBuffer) > 0 && s.batchBuffer[0].Number() == s.BatchPos {
		var batch EspressoBatch
		batch, s.batchBuffer = s.batchBuffer[0], s.batchBuffer[1:]
		s.BatchPos += 1
		return &batch
	}

	return nil

}
