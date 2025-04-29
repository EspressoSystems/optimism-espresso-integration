package espresso

import (
	"cmp"
	"slices"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/core/types"
)

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

type Batch interface {
	Number() uint64
	L1Origin() eth.BlockID
	Header() *types.Header
}

type BatchBuffer[B Batch] struct {
	batches []B
}

func NewBatchBuffer[B Batch]() BatchBuffer[B] {
	return BatchBuffer[B]{
		batches: []B{},
	}
}

func (b BatchBuffer[B]) Len() int {
	return len(b.batches)
}

func (b *BatchBuffer[B]) Clear() {
	b.batches = nil
}

func (b *BatchBuffer[B]) Insert(batch B, i int) {
	b.batches = slices.Insert(b.batches, i, batch)
}

func (b *BatchBuffer[B]) TryInsert(batch B) (int, bool) {
	pos, batchIsRecorded := slices.BinarySearchFunc(b.batches, batch, func(x, y B) int {
		return cmp.Compare(x.Number(), y.Number())
	})

	return pos, batchIsRecorded

}

func (b *BatchBuffer[B]) Peek() *B {
	if len(b.batches) == 0 {
		return nil
	}
	return &b.batches[0]
}

func (b *BatchBuffer[B]) Pop() *B {
	if len(b.batches) == 0 {
		return nil
	}

	batch := b.batches[0]
	b.batches = b.batches[1:]

	return &batch
}
