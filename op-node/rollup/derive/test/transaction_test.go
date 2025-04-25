package test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	espresso_batch "github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/signer"
	"github.com/ethereum/go-ethereum/log"
)

var rollupCfgTest = &rollup.Config{
	Genesis:   rollup.Genesis{L2: eth.BlockID{Number: 0}},
	L2ChainID: big.NewInt(42),
}

const (
	mnemonic = "test test test test test test test test test test test junk"
	hdPath   = "m/44'/60'/0'/0/1"
)

func TestBatchRoundtrip(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	block, _ := RandomL2Block(rng, 10, time.Now())

	batch, err := espresso_batch.BlockToEspressoBatch(rollupCfgTest, block)
	if err != nil {
		t.Fatal(err)
	}

	signerFactory, batcherAddress, err := crypto.ChainSignerFactoryFromConfig(
		log.New(context.Background()),
		"",
		mnemonic,
		hdPath,
		signer.NewCLIConfig(),
	)
	if err != nil {
		t.Fatal(err)
	}
	signer := signerFactory(rollupCfgTest.L2ChainID, batcherAddress)

	transaction, err := batch.ToEspressoTransaction(
		context.Background(),
		rollupCfgTest.L2ChainID.Uint64(),
		signer,
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = espresso_batch.UnmarshalEspressoTransaction(transaction.Payload, batcherAddress)
	if err != nil {
		t.Fatal(err)
	}
}
