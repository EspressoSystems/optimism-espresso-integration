package espresso

import (
	"context"
	"fmt"
	"math/big"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network-go/light-client"
	espressoTypes "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type L1Client interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

// espresso-network-go's HeaderInterface currently lacks a function to get this info,
// although it is present in all header versions
func getFinalizedL1(header *espressoTypes.HeaderImpl) espressoTypes.L1BlockInfo {
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

	L1Client            L1Client // TODO Philippe apparently not used yet
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
	finalizedL1         eth.L1BlockRef

	// Maintained in sorted order, but may be missing batches if we receive
	// any out of order.
	BatchBuffer BatchBuffer[B]

	UnmarshalBatch func([]byte) (*B, error)
}

func NewEspressoStreamer[B Batch](
	namespace uint64,
	l1Client L1Client,
	espressoClient *espressoClient.Client,
	lightClient *espressoLightClient.LightClientReader,
	log log.Logger,
	unmarshalBatch func([]byte) (*B, error),
) EspressoStreamer[B] {
	return EspressoStreamer[B]{
		L1Client:            l1Client,
		EspressoClient:      espressoClient,
		EspressoLightClient: lightClient,
		Log:                 log,

		Namespace:   namespace,
		BatchPos:    1,
		BatchBuffer: NewBatchBuffer[B](),

		UnmarshalBatch: unmarshalBatch,
	}
}

// Reset the state to the last safe batch
func (s *EspressoStreamer[B]) Reset() {
	s.BatchPos = s.confirmedBatchPos + 1
	s.hotShotPos = s.confirmedHotShotPos
	s.BatchBuffer.Clear()
}

// Handle both L1 reorgs and batcher restarts by updating our state in case it is
// not consistent with what's on the L1. Returns true if the state was updated.
func (s *EspressoStreamer[B]) Refresh(ctx context.Context, syncStatus *eth.SyncStatus) (bool, error) {
	if s.confirmedBatchPos == syncStatus.SafeL2.Number {
		return false, nil
	}

	//	hotshotState, err := s.EspressoLightClient.LightClient.
	//		FinalizedState(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(syncStatus.SafeL2.L1Origin.Number)})
	//	if err != nil {
	//		return false, err
	//	}

	s.finalizedL1 = syncStatus.FinalizedL1
	s.confirmedBatchPos = syncStatus.SafeL2.Number
	s.confirmedHotShotPos = 0
	s.Reset()
	return true, nil
}

func (s *EspressoStreamer[B]) Update(ctx context.Context) error {
	// Fetch more batches from HotShot if available.
	blockHeight, err := s.EspressoClient.FetchLatestBlockHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch HotShot block height: %w", err)
	}

	targetHeight := min(blockHeight, s.hotShotPos+100)
	s.Log.Info("Fetching", "from", s.hotShotPos, "upTo", targetHeight)

	for ; s.hotShotPos < targetHeight; s.hotShotPos += 1 {
		s.Log.Debug("fetching HotShot block", "blockNr", s.hotShotPos)

		txns, err := s.EspressoClient.FetchTransactionsInBlock(ctx, s.hotShotPos, s.Namespace)
		if err != nil {
			return fmt.Errorf("failed to fetch transactions in block: %w", err)
		}

		s.Log.Info(fmt.Sprint("Fetched HS block %w", rune(s.hotShotPos)))
		if len(txns.Transactions) == 0 {
			s.Log.Info("no transactions in hotshot block", "blockNr", s.hotShotPos)
			continue
		} else {
			s.Log.Warn("yes transactions in hotshot block", "blockNr", s.hotShotPos)
		}

		// rawHeader, err := s.EspressoClient.FetchRawHeaderByHeight(ctx, s.hotShotPos)
		// if err != nil {
		// 	return fmt.Errorf("failed to fetch raw HotShot header: %w", err)
		// }

		// var header espressoTypes.HeaderImpl
		// err = json.Unmarshal(rawHeader, &header)
		// if err != nil {
		// 	return fmt.Errorf("could not unmarshal header from bytes")
		// }

		// snapshot, err := s.EspressoLightClient.FetchMerkleRoot(s.hotShotPos, nil)
		// if err != nil {
		// 	return fmt.Errorf("failed to fetch Merkle root: %w", err)
		// }

		// if snapshot.Height <= s.hotShotPos {
		// 	return fmt.Errorf("snapshot height is less than or equal to the requested height")
		// }

		// TODO Philippe initialize when creating the streamer
		//s.BatchBuffer.SetBatcherAddress(s.BatcherAddress)
		for _, transaction := range txns.Transactions {

			batch, err := s.UnmarshalBatch(transaction)
			if err != nil {
				// Invalid Batch
				s.Log.Warn("Invalid batch", "error", err)
				continue
			}

			if (*batch).Number() < s.BatchPos {
				continue
			}

			// origin := (*batch).L1Origin()
			// if origin.Number > s.finalizedL1.Number {
			// 	break
			// }

			// l1header, err := s.L1Client.HeaderByNumber(ctx, new(big.Int).SetUint64(origin.Number))
			// if err != nil {
			// 	break
			// }

			// if l1header.Hash() != origin.Hash {
			// 	continue
			// }

			s.Log.Warn("Inserting batch into buffer", "batch", batch)
			s.BatchBuffer.Insert(*batch)
		}
	}

	return nil
}

func (s *EspressoStreamer[B]) Start(ctx context.Context) error {
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
}

// TODO this logic might be slightly different between batcher and derivation
func (s *EspressoStreamer[B]) Next(ctx context.Context) *B {
	// Is the next batch available?
	if s.BatchBuffer.Len() > 0 && (*s.BatchBuffer.Peek()).Number() == s.BatchPos {
		s.BatchPos += 1
		return s.BatchBuffer.Pop()
	}

	return nil
}
