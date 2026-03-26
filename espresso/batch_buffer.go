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

var ErrDuplicateBatch = errors.New("duplicate batch")

type Batch interface {
	Number() uint64
	L1Origin() eth.BlockID
	Header() *types.Header
	Hash() common.Hash
}

type BatchBuffer[B Batch] struct {
	batches []B
}

func NewBatchBuffer[B Batch](capacity uint64) BatchBuffer[B] {
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

// Insert adds a batch to the buffer in sorted order by batch number.
// The caller (processEspressoTransaction) is responsible for ensuring the buffer
// stays bounded by dropping batches that are too far ahead of the current position.
func (b *BatchBuffer[B]) Insert(batch B) error {
	pos, alreadyExists := slices.BinarySearchFunc(b.batches, batch, func(a, t B) int {
		// Note: we use a custom comparison function that returns 0 only if the batches are actually
		// the same to ensure that newer batches with the same number are stored later in the buffer
		if a.Hash() == t.Hash() {
			return 0
		}

		if a.Number() > t.Number() {
			return 1
		} else {
			return -1
		}
	})

	if alreadyExists {
		return ErrDuplicateBatch
	}

	b.batches = slices.Insert(b.batches, pos, batch)
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
