package environment_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	espresso "github.com/EspressoSystems/espresso-network-go/client"
	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/common/hexutil"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// TestDeterministicDerivationExecutionState is a test that
// attempts to make sure that the caff node can derive the same state as the
// original op-node (non caffeinated).
//
// This tests is designed to evaluate Test 3.2 as outlined within the
// Espresso Celo Integration plan.  It has stated task definition as follows:
//
//	Arrange:
//		Running Sequencer, Batcher in Espresso mode, Caff node, and OP node.
//		Once a state of op-node is finalized on L1, it should match the state that was earlier reported by the caff-node for the same block.
//		Manually sending some “invalid” transactions to Espresso (e.g. transactions that look valid but are sent from outside the batcher).
//	Act:
//		Form a transaction that looks valid, and send it directly to Espresso to bypass the batcher.
//	Assert:
//		Query the OP node and make sure the invalid transaction didn't go through.
//		Query the Caff node and make sure the invalid transaction didn't go through.

// forgeBatcherPrivateKey is a helper function that forge a batcher private key
func forgeBatcherPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func realBatcherPrivateKey(system *e2esys.System) (*ecdsa.PrivateKey, error) {
	return system.Cfg.Secrets.Batcher, nil
}

// createEspressoTransaction creates a Espresso transaction with a FAKE or REAL batcher private key
func createEspressoTransaction(tx *geth_types.Transaction, chainID *big.Int, batcherKey *ecdsa.PrivateKey) (*espressoCommon.Transaction, error) {
	// Create a mock block header for the batch
	header := &geth_types.Header{
		Number: big.NewInt(1), // Mock block number
		Time:   uint64(time.Now().Unix()),
	}

	// Create a SingularBatch with the transaction
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to encode transaction: %w", err)
	}

	batch := derive.SingularBatch{
		EpochNum:     rollup.Epoch(1), // Mock epoch number
		Timestamp:    header.Time,
		Transactions: []hexutil.Bytes{txBytes},
	}

	// Create EspressoBatch
	espressoBatch := &derive.EspressoBatch{
		BatchHeader:   header,
		Batch:         batch,
		L1InfoDeposit: tx,
	}

	// Encode the EspressoBatch using RLP
	buf := new(bytes.Buffer)
	err = rlp.Encode(buf, espressoBatch)
	if err != nil {
		return nil, fmt.Errorf("failed to encode batch: %w", err)
	}

	// Sign the encoded batch with FAKE or REAL batcher private key
	batcherSignature, err := crypto.Sign(crypto.Keccak256(buf.Bytes()), batcherKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create batcher signature: %w", err)
	}

	// Combine signature and batch data
	payload := append(batcherSignature, buf.Bytes()...)

	// Create and return Espresso Transaction
	return &espressoCommon.Transaction{
		Namespace: chainID.Uint64(),
		Payload:   payload,
	}, nil
}

func TestInvalidTransaction(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0)
	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	caffNode, err := env.LaunchDecaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer env.Stop(t, caffNode)

	// We want to setup our test
	addressAlice := system.Cfg.Secrets.Addresses().Alice

	l2Verif := system.NodeClient(e2esys.RoleVerif)
	caffVerif := system.NodeClient(env.RoleCaffNode)
	espressoClient := espresso.NewClient("http://op-espresso-devnode:24000")

	// Form a transaction that looks valid
	tx := geth_types.MustSignNewTx(system.Cfg.Secrets.Bob, geth_types.LatestSignerForChainID(system.Cfg.L2ChainIDBig()), &geth_types.DynamicFeeTx{
		ChainID:   system.Cfg.L2ChainIDBig(),
		Nonce:     uint64(0),
		To:        &addressAlice,
		Value:     big.NewInt(1),
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21_000,
	})

	// Create a real Espresso transaction
	realBatcherPrivateKey, err := realBatcherPrivateKey(system)
	if err != nil {
		t.Fatalf("Failed to get real batcher private key:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
	}
	realEspressoTransaction, err := createEspressoTransaction(tx, system.Cfg.L2ChainIDBig(), realBatcherPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create real Espresso transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
	}

	// Send transaction directly to Espresso
	_, err = espressoClient.SubmitTransaction(ctx, *realEspressoTransaction)
	if err != nil {
		t.Fatalf("Failed to submit transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
	}

	// Check the transaction go through l2Verif
	_, err = wait.ForReceiptOK(ctx, l2Verif, tx.Hash())
	if have, want := err, error(nil); have != want {
		t.Fatalf("Waiting for L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}
	// Check the transaction go through caff node
	_, err = wait.ForReceiptOK(ctx, caffVerif, tx.Hash())
	if have, want := err, error(nil); have != want {
		t.Fatalf("Waiting for L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Create a fake Espresso transaction
	fakeBatcherPrivateKey, err := forgeBatcherPrivateKey()
	if err != nil {
		t.Fatalf("Failed to get fake batcher private key:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
	}
	fakeEspressoTransaction, err := createEspressoTransaction(tx, system.Cfg.L2ChainIDBig(), fakeBatcherPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create fake Espresso transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
	}

	// Send transaction directly to Espresso to bypass the batcher
	_, err = espressoClient.SubmitTransaction(ctx, *fakeEspressoTransaction)
	if err != nil {
		t.Fatalf("Failed to submit transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
	}

	// Check the transaction never go through to l2Verif
	_, err = wait.ForReceiptOK(ctx, l2Verif, tx.Hash())
	if have, notwant := err, error(nil); have == notwant {
		t.Fatalf("Waiting for L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, notwant)
	}

	// Check the transaction never go through to caff node
	_, err = wait.ForReceiptOK(ctx, caffVerif, tx.Hash())
	if have, notwant := err, error(nil); have == notwant {
		t.Fatalf("Waiting for L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, notwant)
	}

}
