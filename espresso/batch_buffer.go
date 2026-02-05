package espresso

import (
	"errors"
	"slices"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
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
	// BatchPast indicates that the batch is from the past, i.e. its timestamp is smaller or equal
	// to the safe head's timestamp.
	BatchPast
)

var ErrAtCapacity = errors.New("batch buffer at capacity")

type Batch interface {
	Number() uint64
	L1Origin() eth.BlockID
	Header() *types.Header
	Hash() common.Hash
}

type BatchBuffer[B Batch] struct {
	batches  []B
	capacity uint64
}

func NewBatchBuffer[B Batch](capacity uint64) BatchBuffer[B] {
	return BatchBuffer[B]{
		capacity: capacity,
		batches:  []B{},
	}
}

func (b BatchBuffer[B]) Capacity() uint64 {
	return b.capacity
}

func (b BatchBuffer[B]) Len() int {
	return len(b.batches)
}

func (b *BatchBuffer[B]) Clear() {
	b.batches = nil
}

func (b *BatchBuffer[B]) Insert(batch B) error {
	if uint64(b.Len()) >= b.capacity {
		return ErrAtCapacity
	}

	pos, alreadyExists := slices.BinarySearchFunc(b.batches, batch, func(a, t B) int {
		// Note: we use a custom comparison function that returns 0 only if the batches are actually
		// the same to ensure that newer batches with the same number are stored later in the buffer
		if a.Number() > t.Number() {
			return 1
		} else if a.Number() == t.Number() && a.Hash() == t.Hash() {
			return 0
		} else {
			return -1
		}
	})

	if !alreadyExists {
		b.batches = slices.Insert(b.batches, pos, batch)
	}

	return nil
}

func (b *BatchBuffer[B]) Get(i int) *B {
	if i < b.Len() {
		return &b.batches[i]
	} else {
		return nil
	}
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
