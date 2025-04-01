package espresso

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	deriveTestutils "github.com/ethereum-optimism/optimism/op-node/rollup/derive/test"
	"github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/signer"
	"github.com/ethereum/go-ethereum/log"
)

var rollupCfg = &rollup.Config{
	Genesis:   rollup.Genesis{L2: eth.BlockID{Number: 0}},
	L2ChainID: big.NewInt(42),
}

const (
	mnemonic = "test test test test test test test test test test test junk"
	hdPath   = "m/44'/60'/0'/0/1"
)

func TestBatchRoundtrip(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	block, _ := deriveTestutils.RandomL2Block(rng, 10, time.Now())

	batch, err := BlockToEspressoBatch(rollupCfg, block)
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
	signer := signerFactory(rollupCfg.L2ChainID, batcherAddress)

	transaction, err := batch.ToEspressoTransaction(
		context.Background(),
		rollupCfg.L2ChainID.Uint64(),
		signer,
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = UnmarshalEspressoTransaction(transaction.Payload, batcherAddress)
	if err != nil {
		t.Fatal(err)
	}
}
