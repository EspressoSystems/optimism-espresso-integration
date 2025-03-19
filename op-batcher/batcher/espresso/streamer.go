package espresso

import (
	// #cgo darwin,arm64 LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
	"C"
	"cmp"
	"context"
	"fmt"
	"math/big"
	"slices"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network-go/light-client"
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
	// Position of ???
	confirmedBatchPos uint64
	// HotShot block we're to fetch next
	hotShotPos uint64

	// Maintained in sorted order, but may be missing batches if we receive
	// any out of order.
	batchBuffer []EspressoBatch
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
	s.BatchPos = s.confirmedBatchPos + 1
	s.hotShotPos = hotshotState.BlockHeight
	s.batchBuffer = nil
	return true, nil
}

func (s *EspressoStreamer) Update(ctx context.Context) error {
	// Fetch more batches from HotShot if available.
	hotShotHeight, err := s.EspressoClient.FetchLatestBlockHeight(ctx)

	if err != nil {
		return fmt.Errorf("failed to fetch HotShot block height: %w", err)
	}

	targetHeight := min(hotShotHeight, s.hotShotPos+100)

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

		header, err := s.EspressoClient.FetchHeaderByHeight(ctx, s.hotShotPos)
		if err != nil {
			return fmt.Errorf("failed to fetch HotShot header: %w", err)
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
			// TODO: Fails
			s.Log.Error("namespace verification failed for HS block", "blockNr", s.hotShotPos)
			// return fmt.Errorf("namespace verification failed")
		}

		for _, transaction := range txns.Transactions {
			batch, err := UnmarshalEspressoTransaction(transaction, s.BatcherAddress)
			if err != nil {
				s.Log.Info("Failed to unmarshal espresso transaction", "error", err)
				continue
			}

			if batch.BlockNum < s.BatchPos {
				// Batch already buffered/finalized
				s.Log.Debug("batch is older than current batchPos, skipping", "batchNr", batch.BlockNum, "batchPos", s.BatchPos)
				continue
			}

			if uint64(batch.Batch.EpochNum) > header.Header.GetL1Head() {
				// Enforce that we only deal with finalized deposits
				s.Log.Warn("batch with unfinalized L1 origin")
				continue
			}

			// Find a slot to insert the batch
			i, batchRecorded := slices.BinarySearchFunc(s.batchBuffer, batch, func(x, y EspressoBatch) int {
				return cmp.Compare(x.BlockNum, y.BlockNum)
			})

			if batchRecorded {
				// Duplicate batch found, skip it
				s.Log.Debug("duplicate batch, skipping", "batchNr", batch.BlockNum)
				continue
			}

			s.Log.Debug("recovered batch, buffering", "batchnr", batch.BlockNum)
			s.batchBuffer = slices.Insert(s.batchBuffer, i, batch)
		}
	}

	return nil
}

func (s *EspressoStreamer) Next(ctx context.Context) *EspressoBatch {
	// Is the next batch available?
	if len(s.batchBuffer) > 0 && s.batchBuffer[0].BlockNum == s.BatchPos {
		var batch EspressoBatch
		batch, s.batchBuffer = s.batchBuffer[0], s.batchBuffer[1:]
		s.BatchPos += 1
		return &batch
	}

	return nil

}
