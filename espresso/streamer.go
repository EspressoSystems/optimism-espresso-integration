package espresso

import (
	// #cgo darwin,arm64 LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
	"C"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"

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

type L1Client interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

/// Struct definitions copied here to avoid import cycles. TODO Philippe is this the right way?

// A SingularBatch with block number attached to restore ordering
// when fetching from Espresso
type EspressoBatch struct {
	Header types.Header
	Batch  SingularBatch
}

// SingularBatch is an implementation of Batch interface, containing the input to build one L2 block.
type SingularBatch struct {
	ParentHash   common.Hash  // parent L2 block hash
	EpochNum     rollup.Epoch // aka l1 num
	EpochHash    common.Hash  // l1 block hash
	Timestamp    uint64       // l2 block timestamp
	Transactions []hexutil.Bytes
}

type BatchBuffer interface {
	empty()
	setHeader(header espressoTypes.HeaderImpl)
	setBatchPos(pos uint64)
	setBatcherAddress(address common.Address)
	parseAndInsert(data []byte)
	referenceL1BlockNumber() uint64
	removeFirst()
	get(post int) EspressoBatch
	len() int
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
	batchBuffer BatchBuffer
}

// Reset the state to the last safe batch
func (s *EspressoStreamer) Reset() {
	s.BatchPos = s.confirmedBatchPos + 1
	s.hotShotPos = s.confirmedHotShotPos
	s.batchBuffer.empty()
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

		// TODO Philippe initialize when creating the streamer
		s.batchBuffer.setBatcherAddress(s.BatcherAddress)
		for _, transaction := range txns.Transactions {

			s.batchBuffer.setBatchPos(s.BatchPos)
			s.batchBuffer.setHeader(header)

			s.batchBuffer.parseAndInsert(transaction)
		}
	}

	return nil
}

func (s *EspressoStreamer) Next(ctx context.Context) *EspressoBatch {
	// Is the next batch available?
	if s.batchBuffer.len() > 0 && s.batchBuffer.referenceL1BlockNumber() == s.BatchPos {
		var batch EspressoBatch
		batch = s.batchBuffer.get(0)
		s.batchBuffer.removeFirst()
		s.BatchPos += 1
		return &batch
	}

	return nil

}
