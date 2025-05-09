package environment_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/log"

	espresso "github.com/EspressoSystems/espresso-network-go/client"
	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum/go-ethereum/common/hexutil"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// forgeBatcherPrivateKey is a helper function that forge a batcher private key
func forgeBatcherPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func realBatcherPrivateKey(system *e2esys.System) (*ecdsa.PrivateKey, error) {
	return system.Cfg.Secrets.Batcher, nil
}

// createEspressoTransaction creates a Espresso transaction with a FAKE or REAL batcher private key
func createEspressoTransaction(tx *geth_types.Transaction, chainID *big.Int, batcherKey *ecdsa.PrivateKey) (*espressoCommon.Transaction, error) {
	bufData, err := hexutil.Decode("0xf90388f9023da00d68b82fa254b7d23a8584bcaa67be241a269c86aac05a2a6fc805a672bb910ea01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347944200000000000000000000000000000000000011a0d6cc9c002bc6a8d1c8501c57301b6b2f037494e1e0f61e417411e17f4e80b5afa028881bc4fc4c5fa67f26462837f88937961b6667ae4af043218a0c1b72a5f53ca0d8056577b8ef8e580c0ebc96def906b3699ddc8d91e15abf9c7a7e7bb4f85c96b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080018401c9c380830272ca84681d98b780a0000000000000000000000000000000000000000000000000000000000000000088000000000000000001a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b4218080a00000000000000000000000000000000000000000000000000000000000000000f849a00d68b82fa254b7d23a8584bcaa67be241a269c86aac05a2a6fc805a672bb910e80a0d7d069186bed40982ca7e7747d61c78718d0dda165d74e62164c1bba165001f784681d98b7c0b8fb7ef8f8a07a2aa57f213dfe5e61ceaebcd45c61252157b4e3c1e82e1ec0dca455b1173ad894deaddeaddeaddeaddeaddeaddeaddeaddead00019442000000000000000000000000000000000000158080830f424080b8a4440a5e20000f424000000000000000000000000100000000681d98b60000000000000000000000000000000000000000000000000000000000000000000000003b9aca000000000000000000000000000000000000000000000000000000000000000001d7d069186bed40982ca7e7747d61c78718d0dda165d74e62164c1bba165001f70000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc")
	if err != nil {
		log.Error("failed to decode Espresso transaction in the test", "error", err)
		return nil, err
	}
	buf := bytes.NewBuffer(bufData)

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

// TestValidEspressoTransactionCreation is a test that
// make sure we have correct way to create a Espresso transaction.
// This test is to serve TestInvalidTransactionOutsideBatcher
func TestValidEspressoTransactionCreation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// once this StartDevNet returns, we have a running Espresso Dev Node, and should be able to fetch batcherCLIConfig.EspressoUrl
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

	// l2Verif := system.NodeClient(e2esys.RoleVerif)
	// caffVerif := system.NodeClient(env.RoleCaffNode)
	espressoClient := espresso.NewClient(espressoDevNode.EspressoUrl())

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

	// create a real Espresso transaction and make sure it can go through
	{
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
		txHash, err := espressoClient.SubmitTransaction(ctx, *realEspressoTransaction)
		if err != nil {
			t.Fatalf("Failed to submit transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		}

		// Parameters for transaction fetching loop, which waits for transactions
		// to be sequenced on Espresso
		transactionFetchTimeout := 4 * time.Second
		transactionFetchInterval := 100 * time.Millisecond

		// Check Espresso will accept the transaction
		timer := time.NewTimer(transactionFetchTimeout)
		defer timer.Stop()

		ticker := time.NewTicker(transactionFetchInterval)
		defer ticker.Stop()

	fetchLoop:
		for {
			select {
			case <-ticker.C:
				_, err := espressoClient.FetchTransactionByHash(ctx, txHash)
				if err == nil {
					// test pass
					break fetchLoop
				}
			case <-timer.C:
				t.Fatalf("Failed to fetch transaction by hash after multiple attempts")
			case <-ctx.Done():
				t.Fatalf("Cancelling transaction publishing")
			}
		}
	}

}

// TestInvalidEspressoTransactionOutsideBatcher is a test that
// attempts to make sure that invalid transaction doesn't go through.
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
func TestInvalidEspressoTransactionOutsideBatcher(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// once this StartDevNet returns, we have a running Espresso Dev Node, and should be able to fetch batcherCLIConfig.EspressoUrl
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
	espressoClient := espresso.NewClient(espressoDevNode.EspressoUrl())

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

	// use the same way as creating a real transaction but a fake batcher private key to create a fake Espresso transaction, and make sure it cannot go through
	{
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
}
