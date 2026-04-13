package batcher

import (
	"context"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	op "github.com/EspressoSystems/espresso-streamers/op"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

const mapStreamerBatchSize uint64 = 100

type MapStreamer[B op.Batch] struct {
	client    *espressoClient.MultipleNodesClient
	namespace uint64
	unmarshal func([]byte) (*B, error)
	log       log.Logger

	batches map[uint64]map[common.Hash]*B

	nextBatchPos uint64

	safeBatchNumber uint64

	hotShotPos uint64

	originHotShotPos uint64

	originBatchPos uint64
}

func NewMapStreamer[B op.Batch](
	client *espressoClient.MultipleNodesClient,
	namespace uint64,
	unmarshal func([]byte) (*B, error),
	logger log.Logger,
	originHotShotPos uint64,
	originBatchPos uint64,
) *MapStreamer[B] {
	return &MapStreamer[B]{
		client:           client,
		namespace:        namespace,
		unmarshal:        unmarshal,
		log:              logger,
		batches:          make(map[uint64]map[common.Hash]*B),
		nextBatchPos:     originBatchPos + 1,
		safeBatchNumber:  originBatchPos,
		hotShotPos:       originHotShotPos,
		originHotShotPos: originHotShotPos,
		originBatchPos:   originBatchPos,
	}
}

func (s *MapStreamer[B]) Update(ctx context.Context) error {
	latest, err := s.client.FetchLatestBlockHeight(ctx)
	if err != nil {
		return err
	}
	if s.hotShotPos >= latest {
		return nil
	}

	end := s.hotShotPos + mapStreamerBatchSize
	if end > latest+1 {
		end = latest + 1
	}

	blocks, err := s.client.FetchNamespaceTransactionsInRange(ctx, s.hotShotPos, end, s.namespace)
	if err != nil {
		return err
	}

	s.log.Info("MapStreamer fetched HotShot range", "start", s.hotShotPos, "end", end, "blocks", len(blocks))

	for i, block := range blocks {
		hotShotHeight := s.hotShotPos + uint64(i)
		for _, txn := range block.Transactions {
			batch, err := s.unmarshal(txn.Payload)
			if err != nil {
				s.log.Warn("MapStreamer: failed to unmarshal batch", "hotShotHeight", hotShotHeight, "err", err)
				continue
			}
			blockNum := (*batch).Number()
			parentHash := (*batch).Header().ParentHash

			if s.batches[blockNum] == nil {
				s.batches[blockNum] = make(map[common.Hash]*B)
			}
			if _, exists := s.batches[blockNum][parentHash]; !exists {
				s.batches[blockNum][parentHash] = batch
				s.log.Info("MapStreamer: stored batch",
					"blockNr", blockNum,
					"hash", (*batch).Hash(),
					"parentHash", parentHash,
					"blockHash", (*batch).Header().Hash(),
					"hotShotHeight", hotShotHeight,
					"forks", len(s.batches[blockNum]),
				)
			} else {
				s.log.Info("MapStreamer: ignoring duplicate batch",
					"blockNr", blockNum,
					"hash", (*batch).Hash(),
					"parentHash", parentHash,
					"hotShotHeight", hotShotHeight,
				)
			}
		}
	}

	s.hotShotPos = end
	return nil
}

func (s *MapStreamer[B]) Refresh(_ context.Context, _ eth.L1BlockRef, safeBatchNumber uint64, _ eth.BlockID) error {
	if safeBatchNumber <= s.safeBatchNumber {
		return nil
	}
	for n := s.safeBatchNumber + 1; n <= safeBatchNumber; n++ {
		delete(s.batches, n)
	}
	s.log.Info("MapStreamer: pruned safe batches", "oldSafe", s.safeBatchNumber, "newSafe", safeBatchNumber)
	s.safeBatchNumber = safeBatchNumber
	return nil
}

func (s *MapStreamer[B]) RefreshSafeL1Origin(_ eth.BlockID) {}

func (s *MapStreamer[B]) Reset() {
	prev := s.nextBatchPos
	s.nextBatchPos = s.safeBatchNumber + 1
	s.log.Info("MapStreamer: reset", "prevNextBatchPos", prev, "safeBatchNumber", s.safeBatchNumber, "newNextBatchPos", s.nextBatchPos)
}

func (s *MapStreamer[B]) UnmarshalBatch(b []byte) (*B, error) {
	return s.unmarshal(b)
}

func (s *MapStreamer[B]) HasNext(_ context.Context) bool {
	return len(s.batches[s.nextBatchPos]) > 0
}

func (s *MapStreamer[B]) Peek(parentHash common.Hash) *B {
	forks := s.batches[s.nextBatchPos]
	if len(forks) == 0 {
		return nil
	}
	// Zero hash means the channel manager has no tip yet (empty state); accept any fork.
	if parentHash == (common.Hash{}) {
		for _, batch := range forks {
			return batch
		}
	}
	batch, ok := forks[parentHash]
	if !ok {
		s.log.Info("MapStreamer: no fork matches tip",
			"blockNr", s.nextBatchPos,
			"tip", parentHash,
			"forks", len(forks),
		)
		return nil
	}
	return batch
}

func (s *MapStreamer[B]) Next(_ context.Context) *B {
	forks := s.batches[s.nextBatchPos]
	if len(forks) == 0 {
		return nil
	}

	var batch *B
	for _, b := range forks {
		batch = b
		break
	}
	s.log.Info("MapStreamer: advancing past batch",
		"blockNr", s.nextBatchPos,
		"forks", len(forks),
		"nextBatchPos", s.nextBatchPos+1,
	)
	s.nextBatchPos++
	return batch
}
