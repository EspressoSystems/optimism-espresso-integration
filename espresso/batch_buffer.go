package espresso

import (
	"bytes"
	"cmp"
	"encoding/binary"
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
	Hash() common.Hash
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

type DummyBatch struct {
	number   uint64
	l1Origin eth.BlockID
}

func (b DummyBatch) Number() uint64 {
	return b.number
}

func (b DummyBatch) L1Origin() eth.BlockID {
	return b.l1Origin
}

func (b DummyBatch) Hash() common.Hash {
	return common.Hash{}
}

func (b DummyBatch) Header() *types.Header {
	return nil
}

var b = NewBatchBuffer[DummyBatch]()

func Fuzz(data []byte) int {
	var num uint64
	err := binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &num)

	if err != nil {
		b.Insert(DummyBatch{number: num, l1Origin: eth.BlockID{Number: num}}, 0)
	} else {
		b.Peek()
		b.Pop()
	}

	return 0
}
