package environment_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// TestDeterministicDerivationExecutionStateWithInvalidTransaction is a test that
// attempts to make sure that the caff node can derive the same state as the
// original op-node (non caffeinated).
//
// This test is designed to evaluate Test 3.2 as outlined within the
// Espresso Celo Integration plan.  It has stated task definition as follows:
//
//	Arrange:
//		Running Sequencer, Batcher in Espresso mode, Caff node, and OP node.
//	Act:
//		Send some transactions from Bob to Alice and some regular L2 transactions.
//		While you send normal L2 tx to the sequencer and monitor the OP node as well as the Caff node,
//		you also send transactions to Espresso using an invalid batcher address, and transactions directly to L1 (e.g. transactions that were not previously posted to Espresso).
//	Assert:
//		Once a state of op-node is finalized on L1, it should match the state that was earlier reported by the caff-node for the same block.
//		Query the executive machine state when Caff node is on
//		Query the executive machine state when OP node is on
//		Make sure the states are the same

func TestDeterministicDerivationExecutionStateWithInvalidTransaction(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t)
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

	// Get caffNodeL2Client from caff node's engine state
	caffNodeL2Client := caffNode.OpNode.EngineState()

	// We want to setup our test
	addressAlice := system.Cfg.Secrets.Addresses().Alice
	espressoClient := espressoClient.NewClient(espressoDevNode.EspressoUrl())
	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	l2Seq := system.NodeClient(e2esys.RoleSeq)

	// We want to send some transactions from Bob to Alice
	{
		privateKey := system.Cfg.Secrets.Bob
		bobOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, system.Cfg.L1ChainIDBig())
		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to create transaction options for bob:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		mintAmount := new(big.Int).SetUint64(1)
		bobOptions.Value = mintAmount
		_ = helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, bobOptions, func(l2Opts *helpers.DepositTxOpts) {
			// Send from Bob to Alice
			l2Opts.ToAddr = addressAlice
		})
	}

	numIterations := 10
	attackRoundEspresso := 5 // the round where we send transaction directly to Espresso outside of the batcher
	attackRoundL1 := 7       // the round where we send transaction directly to batch inbox
	// Compare states between nodes for multiple latest blocks
	// We don't compare states for every individual block as any diff in block x will be reflected in block x + n
	for i := 0; i < numIterations; i++ {

		// Send some regular L2 transactions in each iteration
		tx := geth_types.MustSignNewTx(system.Cfg.Secrets.Bob, geth_types.LatestSignerForChainID(system.Cfg.L2ChainIDBig()), &geth_types.DynamicFeeTx{
			ChainID:   system.Cfg.L2ChainIDBig(),
			Nonce:     uint64(i + 1), // +1 because of the deposit transaction above
			To:        &addressAlice,
			Value:     big.NewInt(1),
			GasTipCap: big.NewInt(10),
			GasFeeCap: big.NewInt(200),
			Gas:       21_000,
		})
		err := l2Seq.SendTransaction(ctx, tx)
		if have, want := err, error(nil); have != want {
			t.Fatalf("Sending L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}
		// Wait for the receipt
		_, err = wait.ForReceiptOK(ctx, l2Seq, tx.Hash())
		if have, want := err, error(nil); have != want {
			t.Fatalf("Waiting for L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		// When it is the attack round, try to send some Espresso transactions with fakeBatcherPrivateKey directly to Espresso, outside of the batcher.
		// Use the same way as creating a real transaction but a fake batcher private key to create a fake Espresso transaction.
		if i == attackRoundEspresso {
			// Create a fake Espresso transaction
			fakeBatcherPrivateKey, err := forgeBatcherPrivateKey()
			if err != nil {
				t.Fatalf("Failed to get fake batcher private key:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
			}
			fakeEspressoTransaction, err := createEspressoTransaction(system.Cfg.L2ChainIDBig(), fakeBatcherPrivateKey)
			if err != nil {
				t.Fatalf("Failed to create fake Espresso transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
			}

			// Send transaction directly to Espresso to bypass the batcher
			_, err = espressoClient.SubmitTransaction(ctx, *fakeEspressoTransaction)
			if err != nil {
				t.Fatalf("Failed to submit transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
			}

		} else if i == attackRoundL1 {
			// create a transaction
			tx := geth_types.MustSignNewTx(system.Cfg.Secrets.Bob, system.RollupConfig.L1Signer(), &geth_types.DynamicFeeTx{
				ChainID:   system.Cfg.L1ChainIDBig(),
				Nonce:     1,
				To:        &system.RollupConfig.BatchInboxAddress,
				Value:     big.NewInt(1),
				GasTipCap: big.NewInt(10),
				GasFeeCap: big.NewInt(200),
				Gas:       21_000,
			})
			// Send a transaction directly to L1
			err = l1Client.SendTransaction(ctx, tx)
			if have, want := err, error(nil); have != want {
				t.Fatalf("failed to send transaction directly to L1:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}
		}

		// Get latest safe blocks from caff node first
		// as caff node usually lags behind the sequencer node on safe blocks due to submitting additionally to Espresso.
		// We use l2BlockRefByLabel to get the states as the engine state will be reflected in the block.
		caffBlock, err := caffNodeL2Client.L2BlockRefByLabel(ctx, eth.Safe)
		if err != nil {
			t.Fatalf("failed to get block from caff node: %v", err)
		}

		// Get the corresponding block from sequencer
		seqBlock, err := l2Seq.BlockByNumber(ctx, big.NewInt(0).SetUint64(caffBlock.Number))
		if err != nil {
			t.Fatalf("failed to get block from l2Seq: %v", err)
		}

		// Compare block states
		if have, want := caffBlock.Hash, seqBlock.Hash(); have != want {
			t.Errorf("block hash mismatch between sequencer and caff node at block %v\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", seqBlock.Number(), have, want)
		}
	}

}

// forgeBatcherPrivateKey is a helper function that forge a batcher private key
func forgeBatcherPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func realBatcherPrivateKey(system *e2esys.System) (*ecdsa.PrivateKey, error) {
	return system.Cfg.Secrets.Batcher, nil
}

// createEspressoTransaction creates a Espresso transaction with a FAKE or REAL batcher private key
func createEspressoTransaction(chainID *big.Int, batcherKey *ecdsa.PrivateKey) (*espressoCommon.Transaction, error) {
	// This is the genesis Espresso transaction that created by honest sequencer
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
// This test is a unit test to serve the correctness of TestDeterministicDerivationExecutionStateWithInvalidTransaction.
func TestValidEspressoTransactionCreation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// once this StartDevNet returns, we have a running Espresso Dev Node
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t)
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
	espressoClient := espressoClient.NewClient(espressoDevNode.EspressoUrl())

	// create a real Espresso transaction and make sure it can go through
	{
		// Create a real Espresso transaction
		realBatcherPrivateKey, err := realBatcherPrivateKey(system)
		if err != nil {
			t.Fatalf("Failed to get real batcher private key:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		}
		realEspressoTransaction, err := createEspressoTransaction(system.Cfg.L2ChainIDBig(), realBatcherPrivateKey)
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

		// Make sure the transaction will go through to caff node by checking the unmarshal works
		// The check can directly reflect whether the transaction is valid or not
		caffStreamer := caffNode.OpNode.EspressoStreamer()
		_, err = caffStreamer.UnmarshalBatch(realEspressoTransaction.Payload)
		if have, want := err, error(nil); have != want {
			t.Fatalf("Failed to unmarshal batch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		// Make sure the transaction will go through to op node by checking it will go through batch submitter's streamer
		batchSubmitter := system.BatchSubmitter
		_, err = batchSubmitter.EspressoStreamer().UnmarshalBatch(realEspressoTransaction.Payload)
		if have, want := err, error(nil); have != want {
			t.Fatalf("Failed to unmarshal batch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

	}

}
