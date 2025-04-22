package espresso

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network-go/light-client"
	espressoTypes "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	L1Client                      L1Client // TODO Philippe apparently not used yet
	EspressoClient                *espressoClient.Client
	EspressoLightClient           *espressoLightClient.LightClientReader
	Log                           log.Logger
	PollingHotShotPollingInterval time.Duration

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
	pollingHotShotPollingInterval time.Duration,
) EspressoStreamer[B] {
	return EspressoStreamer[B]{
		L1Client:            l1Client,
		EspressoClient:      espressoClient,
		EspressoLightClient: lightClient,
		Log:                 log,

		Namespace:                     namespace,
		BatchPos:                      1,
		BatchBuffer:                   NewBatchBuffer[B](),
		PollingHotShotPollingInterval: pollingHotShotPollingInterval,

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

	hotshotState, err := s.EspressoLightClient.LightClient.
		FinalizedState(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(syncStatus.SafeL2.L1Origin.Number)})
	if err != nil {
		return false, err
	}

	s.finalizedL1 = syncStatus.FinalizedL1
	s.confirmedBatchPos = syncStatus.SafeL2.Number
	s.confirmedHotShotPos = hotshotState.BlockHeight
	s.Reset()
	return true, nil
}

func (s *EspressoStreamer[B]) CheckBatch(batch B) (BatchValidity, int) {

	// TODO finality check
	//espressoFinalizedL1 := getFinalizedL1(&batch)
	//if espressoFinalizedL1 == nil {
	//	log.Error("Invalid batch: Unknown Espresso header version")
	//	return BatchDrop, 0
	//}

	//if uint64(batch.Batch.EpochNum) > espressoFinalizedL1.Number {
	//	// Enforce that we only deal with finalized deposits
	//	log.Warn("batch with unfinalized L1 origin",
	//		"batchEpochNum", batch.Batch.EpochNum, "espressoFinalizedL1Num", espressoFinalizedL1.Number,
	//	)
	//	return BatchUndecided, 0
	//} else {
	//	// make sure it's a valid L1 origin state by check the hash
	//	// TODO Adapt Sishan's logic described in
	//	// https: //github.com/EspressoSystems/optimism-espresso-integration/blob/40a52d5b334f5dca169dfc1b41d8d06a2a72470d/op-node/rollup/derive/espresso_streamer.go#L148
	//}

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

	// Find a slot to insert the batch
	i, batchRecorded := s.BatchBuffer.TryInsert(batch)

	// Batch already buffered/finalized
	if batch.Number() < s.BatchPos {
		s.Log.Error("Batch is older than current batchPos, skipping", "batchNr", batch.Number(), "batchPos", b.batchPos)
		return BatchPast, 0
	}

	if batchRecorded {
		// Duplicate batch found, skip it
		return BatchPast, i
	}

	// We can do this check earlier, but it's a more intensive one, so we do this last.
	// TODO as the batcher is considered honest does is this check needed?
	//for i, txBytes := range batch.Batch.Transactions {
	//	if len(txBytes) == 0 {
	//		b.Log.Error("Transaction data must not be empty, but found empty tx", "tx_index", i)
	//		return BatchDrop, 0
	//	}
	//	if txBytes[0] == types.DepositTxType {
	//		log.Error("sequencers may not embed any deposits into batch data, but found tx that has one", "tx_index", i)
	//		return BatchDrop, 0
	//	}
	//}

	return BatchAccept, i
}

func (s *EspressoStreamer[B]) Update(ctx context.Context) error {
	// Fetch more batches from HotShot if available.
	blockHeight, err := s.EspressoClient.FetchLatestBlockHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch HotShot block height: %w", err)
	}

	targetHeight := min(blockHeight, s.hotShotPos+100)
	s.Log.Debug("Fetching hotshot blocks", "from", s.hotShotPos, "upTo", targetHeight)

	for ; s.hotShotPos < targetHeight; s.hotShotPos += 1 {
		s.Log.Trace("Fetching HotShot block", "block", s.hotShotPos)

		txns, err := s.EspressoClient.FetchTransactionsInBlock(ctx, s.hotShotPos, s.Namespace)
		if err != nil {
			return fmt.Errorf("failed to fetch transactions in block: %w", err)
		}

		s.Log.Trace("Fetched HotShot block", "block", s.hotShotPos, "txns", len(txns.Transactions))

		if len(txns.Transactions) == 0 {
			s.Log.Trace("No transactions in hotshot block", "block", s.hotShotPos)
			continue
		}

		for _, transaction := range txns.Transactions {

			batch, err := s.UnmarshalBatch(transaction)
			if err != nil {
				// Invalid Batch
				s.Log.Warn("Invalid batch", "error", err)
				continue
			}

			s.Log.Trace("Inserting batch into buffer", "batch", batch)

			validity, i := s.CheckBatch(*batch)

			switch validity {

			case BatchDrop:
				//b.Log.Info("Dropping batch", batch)
				return nil

			case BatchPast:
				//b.Log.Info("Batch already processed. Skipping", batch)
				return nil

			case BatchUndecided: // Sishan TODO: remove if this is not needed
				// TODO Philippe logic of remaining list
				return nil

			case BatchAccept:
				//b.Log.Debug("Recovered batch, inserting", "batchnr", batch.Number())

			case BatchFuture:
				//b.Log.Info("Inserting batch for future processing")
			}

			s.BatchBuffer.Insert(*batch, i)
		}
	}

	return nil
}

func (s *EspressoStreamer[B]) Start(ctx context.Context, wg *sync.WaitGroup) {

	s.Log.Info("Starting espresso streamer")
	defer wg.Done()
	ticker := time.NewTicker(s.PollingHotShotPollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := s.Update(ctx)
			if err != nil {
				s.Log.Error("failed to update Espresso streamer", "err", err)
				continue
			}

		case <-ctx.Done():
			s.Log.Info("espressoBatchLoadingLoop returning")
			return
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
