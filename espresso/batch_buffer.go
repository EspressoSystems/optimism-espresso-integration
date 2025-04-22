package espresso

import (
	"cmp"
	"slices"

	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type BatchValidity uint8

type Batch interface {
	Number() uint64
	L1Origin() eth.BlockID
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

func (b *BatchBuffer[B]) Insert(batch B) {
	i, batchRecorded := slices.BinarySearchFunc(b.batches, batch, func(x, y B) int {
		return cmp.Compare(x.Number(), y.Number())
	})

	if batchRecorded {
		return
	}

	b.batches = slices.Insert(b.batches, i, batch)
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
