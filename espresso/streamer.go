package espresso

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"math/big"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network-go/light-client"
	espressoTypes "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type L1Client interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

type EspressoBatchI interface {
	ToIncompleteBlock(rollupCfg *rollup.Config) (*types.Block, error)
}

type BatchBuffer interface {
	Empty()
	SetHeader(header espressoTypes.HeaderImpl)
	SetBatchPos(pos uint64)
	SetBatcherAddress(address common.Address)
	ParseAndInsert(data []byte)
	ReferenceL1BlockNumber() uint64
	RemoveFirst()
	Get(pos int) EspressoBatchI
	Len() int
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
	BatchBuffer BatchBuffer
}

// Reset the state to the last safe batch
func (s *EspressoStreamer) Reset() {
	s.BatchPos = s.confirmedBatchPos + 1
	s.hotShotPos = s.confirmedHotShotPos
	s.BatchBuffer.Empty()
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

		// TODO Philippe initialize when creating the streamer
		s.BatchBuffer.SetBatcherAddress(s.BatcherAddress)
		for _, transaction := range txns.Transactions {

			s.BatchBuffer.SetBatchPos(s.BatchPos)
			s.BatchBuffer.SetHeader(header)
			s.BatchBuffer.ParseAndInsert(transaction)
		}
	}

	// TODO iterate over the remaining list and possibly update the buffer

	return nil
}

func (s *EspressoStreamer) Start(ctx context.Context) error {

	s.Log.Info("In the function, Starting espresso streamer")
	bigTimeout := 2 * time.Minute
	timer := time.NewTimer(bigTimeout)
	defer timer.Stop()

	// Sishan TODO: maybe use better handler with dynamic interval in the future
	ticker := time.NewTicker(2) // TODO make it configurable
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := s.Update(ctx)
			if err != nil {
				s.Log.Error("Error while updating the batches: ", err)
			} else {
				s.Log.Info("Processing block", "block number", s.hotShotPos)
				// Successful execution: reset the timer to start the timeout period over.
				// Stop the timer and drain if needed.
				// TODO Here we need to build a L2 block from the new batch
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(bigTimeout)
			}
			timer.Reset(bigTimeout)

		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			return fmt.Errorf("timeout while queueing messages from hotshot")
		}
	}

	return nil
}

// TODO this logic might be slightly different between batcher and derivation
func (s *EspressoStreamer) Next(ctx context.Context) *EspressoBatchI {
	// Is the next batch available?
	if s.BatchBuffer.Len() > 0 && s.BatchBuffer.ReferenceL1BlockNumber() == s.BatchPos {
		var batch EspressoBatchI
		batch = s.BatchBuffer.Get(0)
		s.BatchBuffer.RemoveFirst()
		s.BatchPos += 1
		return &batch
	}

	return nil

}
