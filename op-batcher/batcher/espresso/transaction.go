package espresso

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	opCrypto "github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// A SingularBatch with block number attached to restore ordering
// when fetching from Espresso
type EspressoBatch struct {
	Header types.Header
	Batch  derive.SingularBatch
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
	batch, _, err := derive.BlockToSingularBatch(rollupCfg, block)
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
func BatchToIncompleteBlock(rollupCfg *rollup.Config, espressoBatch *EspressoBatch) (*types.Block, error) {
	batch := espressoBatch.Batch

	FakeL1info, err := derive.L1InfoDeposit(
		rollupCfg,
		eth.SystemConfig{},
		espressoBatch.Batch.Epoch().Number,
		&testutils.MockBlockInfo{
			InfoHash:    batch.ParentHash,
			InfoBaseFee: big.NewInt(0),
		},
		espressoBatch.Header.Time,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create fake L1 info: %w", err)
	}
	// Insert a fake deposit transaction so that channel doesn't complain about empty blocks
	txs := []*types.Transaction{types.NewTx(FakeL1info)}
	for i, opaqueTx := range batch.Transactions {
		var tx types.Transaction
		err := tx.UnmarshalBinary(opaqueTx)
		if err != nil {
			return nil, fmt.Errorf("could not decode tx %d: %w", i, err)
		}
		txs = append(txs, &tx)
	}
	return types.NewBlockWithHeader(&espressoBatch.Header).WithBody(types.Body{
		Transactions: txs,
	}), nil
}
