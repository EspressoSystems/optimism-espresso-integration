package derive

import (
	"bytes"
	"cmp"
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"

	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	espresso "github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	opCrypto "github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

// espresso-network-go's HeaderInterface currently lacks a function to get this info,
// although it is present in all header versions
func getFinalizedL1(header *espressoCommon.HeaderImpl) *espressoCommon.L1BlockInfo {
	v0_1, ok := header.Header.(*espressoCommon.Header0_1)
	if ok {
		return v0_1.L1Finalized
	}
	v0_2, ok := header.Header.(*espressoCommon.Header0_2)
	if ok {
		return v0_2.L1Finalized
	}
	v0_3, ok := header.Header.(*espressoCommon.Header0_3)
	if ok {
		return v0_3.L1Finalized
	}
	return nil
}

// A SingularBatch with block number attached to restore ordering
// when fetching from Espresso
type EspressoBatch struct {
	Header types.Header
	Batch  SingularBatch
}

// TODO Philippe find better name
type EspressoBatchBuffer struct {
	batches        []EspressoBatch
	batchPos       uint64
	batcherAddress common.Address
	header         espressoCommon.HeaderImpl
	Log            log.Logger
}

func NewEspressoBatchBuffer(batcherAddress common.Address, log log.Logger) *EspressoBatchBuffer {

	bb := new(EspressoBatchBuffer)
	bb.Log = log
	bb.batcherAddress = batcherAddress
	bb.batchPos = 0 // TODO Philippe is this correct?

	return bb

}

func (b *EspressoBatchBuffer) Empty() {
	b.batches = nil
}

func (b *EspressoBatchBuffer) SetHeader(header espressoCommon.HeaderImpl) {
	b.header = header
}

func (b *EspressoBatchBuffer) SetBatchPos(pos uint64) {
	b.batchPos = pos
}

func (b *EspressoBatchBuffer) SetBatcherAddress(batcherAddress common.Address) {
	b.batcherAddress = batcherAddress
}

func (b *EspressoBatchBuffer) Len() int {
	return len(b.batches)
}

func (b *EspressoBatchBuffer) ReferenceL1BlockNumber() uint64 {
	return b.batches[0].Number()
}

func (b *EspressoBatchBuffer) RemoveFirst() {
	b.batches = b.batches[1:]
}

func (b *EspressoBatchBuffer) Get(pos int) espresso.EspressoBatchI {
	return &b.batches[pos]
}

func (b *EspressoBatchBuffer) checkBatch(batch EspressoBatch) (BatchValidity, int) {

	espressoFinalizedL1 := getFinalizedL1(&b.header)
	if espressoFinalizedL1 == nil {
		log.Error("Invalid batch: Unknown Espresso header version")
		return BatchDrop, 0
	}

	if uint64(batch.Batch.EpochNum) > espressoFinalizedL1.Number {
		// Enforce that we only deal with finalized deposits
		log.Warn("batch with unfinalized L1 origin",
			"batchEpochNum", batch.Batch.EpochNum, "espressoFinalizedL1Num", espressoFinalizedL1.Number,
		)
		return BatchUndecided, 0
	} else {
		// make sure it's a valid L1 origin state by check the hash
		// TODO Adapt Sishan's logic described in
		// https: //github.com/EspressoSystems/optimism-espresso-integration/blob/40a52d5b334f5dca169dfc1b41d8d06a2a72470d/op-node/rollup/derive/espresso_streamer.go#L148
	}

	// Find a slot to insert the batch
	i, batchRecorded := slices.BinarySearchFunc(b.batches, batch, func(x, y EspressoBatch) int {
		return cmp.Compare(x.Number(), y.Number())
	})

	// Batch already buffered/finalized
	if batch.Number() < b.batchPos {

		b.Log.Error("Batch is older than current batchPos, skipping", "batchNr", batch.Number(), "batchPos", b.batchPos)
		return BatchPast, 0
	}

	if batchRecorded {
		// Duplicate batch found, skip it
		return BatchPast, i
	}

	// We can do this check earlier, but it's a more intensive one, so we do this last.
	// TODO as the batcher is considered honest does is this check needed?
	for i, txBytes := range batch.Batch.Transactions {
		if len(txBytes) == 0 {
			b.Log.Error("Transaction data must not be empty, but found empty tx", "tx_index", i)
			return BatchDrop, 0
		}
		if txBytes[0] == types.DepositTxType {
			log.Error("sequencers may not embed any deposits into batch data, but found tx that has one", "tx_index", i)
			return BatchDrop, 0
		}
	}

	return BatchAccept, i
}

func (b *EspressoBatchBuffer) ParseAndInsert(data []byte) {
	batch, err := UnmarshalEspressoTransaction(data, b.batcherAddress)
	if err != nil {
		b.Log.Info("Failed to unmarshal espresso transaction", "error", err)
		return
	}

	var validity, i = b.checkBatch(batch)

	switch validity {

	case BatchDrop:
		b.Log.Info("Dropping batch", batch)
		return

	case BatchPast:
		b.Log.Info("Batch already processed. Skipping", batch)
		return

	case BatchUndecided: // Sishan TODO: remove if this is not needed
		// TODO Philippe logic of remaining list
		return

	case BatchAccept:
		b.Log.Debug("Recovered batch, inserting", "batchnr", batch.Number())

	case BatchFuture:
		b.Log.Info("Inserting batch for future processing")
	}

	// For both BatchAccept and BatchFuture we insert.
	b.batches = slices.Insert(b.batches, i, batch)

}

func (b *EspressoBatch) Number() uint64 {
	return b.Header.Number.Uint64()
}

func (b *EspressoBatch) ToEspressoTransaction(ctx context.Context, namespace uint64, signer opCrypto.ChainSigner) (*espressoCommon.Transaction, error) {
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, b)
	if err != nil {
		return nil, fmt.Errorf("failed to encode batch: %w", err)
	}

	batcherSignature, err := signer.Sign(ctx, crypto.Keccak256(buf.Bytes()))

	if err != nil {
		return nil, fmt.Errorf("failed to create batcher signature: %w", err)
	}

	payload := append(batcherSignature, buf.Bytes()...)

	return &espressoCommon.Transaction{Namespace: namespace, Payload: payload}, nil

}

func BlockToEspressoBatch(rollupCfg *rollup.Config, block *types.Block) (*EspressoBatch, error) {
	batch, _, err := BlockToSingularBatch(rollupCfg, block)
	if err != nil {
		return nil, err
	}

	return &EspressoBatch{
		Batch:  *batch,
		Header: *block.Header(),
	}, nil
}

func UnmarshalEspressoTransaction(data []byte, batcherAddress common.Address) (EspressoBatch, error) {
	signatureData, batchData := data[:crypto.SignatureLength], data[crypto.SignatureLength:]
	batchHash := crypto.Keccak256(batchData)

	signer, err := crypto.SigToPub(batchHash, signatureData)
	if err != nil {
		return EspressoBatch{}, err
	}
	if crypto.PubkeyToAddress(*signer) != batcherAddress {
		return EspressoBatch{}, errors.New("invalid signer")
	}

	var batch EspressoBatch
	if err := rlp.DecodeBytes(batchData, &batch); err != nil {
		return EspressoBatch{}, err
	}

	return batch, nil
}

// Deposit transactions obviously aren't recovered from the batch, so this doesn't return
// the original block, but we don't care for batcher purposes,as this incomplete block will be
// converted back to batch later on anyway. This double-conversion is done to avoid extensive
// modifications to channel manager that would be needed to allow it to accept batches directly
//
// NOTE: This function MUST guarantee no transient errors. It is allowed to fail only on
// invalid batches or in case of misconfiguration of the batcher, in which case it should fail
// for all batches.
func (b *EspressoBatch) ToIncompleteBlock(rollupCfg *rollup.Config) (*types.Block, error) {

	FakeL1info, err := L1InfoDeposit(
		rollupCfg,
		eth.SystemConfig{},
		b.Batch.Epoch().Number,
		&testutils.MockBlockInfo{
			InfoHash:    b.Batch.ParentHash,
			InfoBaseFee: big.NewInt(0),
		},
		b.Header.Time,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create fake L1 info: %w", err)
	}
	// Insert a fake deposit transaction so that channel doesn't complain about empty blocks
	txs := []*types.Transaction{types.NewTx(FakeL1info)}
	for i, opaqueTx := range b.Batch.Transactions {
		var tx types.Transaction
		err := tx.UnmarshalBinary(opaqueTx)
		if err != nil {
			return nil, fmt.Errorf("could not decode tx %d: %w", i, err)
		}
		txs = append(txs, &tx)
	}
	return types.NewBlockWithHeader(&b.Header).WithBody(types.Body{
		Transactions: txs,
	}), nil
}
