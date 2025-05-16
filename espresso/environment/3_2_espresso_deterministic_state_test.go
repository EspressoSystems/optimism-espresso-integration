package environment_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	geth_common "github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
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

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, env.WithL1FinalizedDistance(0), env.WithSequencerUseFinalized(true))
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
	// espressoClient := espressoClient.NewClient(espressoDevNode.EspressoUrl())
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
			// // Create a fake Espresso transaction
			// fakeBatcherPrivateKey, err := forgedBatcherPrivateKey()
			// if err != nil {
			// 	t.Fatalf("Failed to get fake batcher private key:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
			// }
			// _, fakeEspressoTransaction, err := createEspressoTransaction(ctx, l2Seq, system.RollupConfig, system.Cfg.L2ChainIDBig(), fakeBatcherPrivateKey, t)
			// if err != nil {
			// 	t.Fatalf("Failed to create fake Espresso transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
			// }

			// // Send transaction directly to Espresso to bypass the batcher
			// _, err = espressoClient.SubmitTransaction(ctx, *fakeEspressoTransaction)
			// if err != nil {
			// 	t.Fatalf("Failed to submit transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
			// }

		} else if i == attackRoundL1 {
			// create a transaction
			tx := geth_types.MustSignNewTx(system.Cfg.Secrets.Bob, system.RollupConfig.L1Signer(), &geth_types.DynamicFeeTx{
				ChainID:   system.Cfg.L1ChainIDBig(),
				Nonce:     1,
				To:        &system.RollupConfig.BatchInboxAddress,
				Value:     big.NewInt(1),
				GasTipCap: big.NewInt(1 * params.GWei),
				GasFeeCap: big.NewInt(10 * params.GWei),
				Gas:       5_000_000,
			})
			// Send a transaction directly to L1
			err = l1Client.SendTransaction(ctx, tx)
			if have, want := err, error(nil); have != want {
				t.Fatalf("failed to send transaction directly to L1:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}

			// Wait for the receipt to fail
			_, err = wait.ForReceiptFail(ctx, l1Client, tx.Hash())
			if have, want := err, error(nil); have != want {
				t.Fatalf("failed to get receipt for transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}
		}

		// Get latest safe blocks from op node first as op node usually lags behind.
		// We use BlockByNumber to get the states as the engine state will be reflected in the block.
		opBlock, err := l2Verif.BlockByNumber(ctx, big.NewInt(rpc.FinalizedBlockNumber.Int64()))
		if err != nil {
			t.Fatalf("failed to get block from opBlock: %v", err)
		}

		// Get the corresponding safe blocks from caff node
		// We use L2BlockRefByLabel to get the states as the engine state will be reflected in the block.
		caffBlock, err := caffNodeL2Client.L2BlockRefByNumber(ctx, opBlock.Number().Uint64())
		if err != nil {
			t.Fatalf("failed to get block from caff node: %v", err)
		}

		// Compare block states
		if have, want := caffBlock.Hash, opBlock.Hash(); have != want {
			t.Errorf("block hash mismatch between sequencer and caff node at block %v\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", opBlock.Number(), have, want)
		}
	}

}

// createValidEspressoBatch creates a valid Espresso batch by
// constructing a block with a deposit transaction. It uses the latest
// block from the sequencer to create a new block with a deposit
// transaction. The block is then converted to an Espresso batch using
// the derive.BlockToEspressoBatch function.
func createValidEspressoBatch(ctx context.Context, depositTx *geth_types.Transaction, tx *geth_types.Transaction, cli *ethclient.Client, rollupCfg *rollup.Config, hasher geth_types.TrieHasher, t *testing.T) (*derive.EspressoBatch, error) {
	// Determine what the latest block in the sequencer is, so we can
	// hope to create a valid transaction, to get something out of it.
	latestBlock, err := cli.BlockByNumber(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to get latest block:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		return nil, err
	}

	latestHeader := latestBlock.Header()
	body := &geth_types.Body{
		Transactions: []*geth_types.Transaction{
			depositTx,
			tx,
		},
	}

	espressoBatch, err := derive.BlockToEspressoBatch(
		rollupCfg,
		geth_types.NewBlock(
			&geth_types.Header{
				ParentHash: latestBlock.Hash(),
				UncleHash:  latestHeader.UncleHash,
				Coinbase:   latestHeader.Coinbase,
				Root:       latestHeader.Root,
				Bloom:      latestHeader.Bloom,
				Difficulty: latestHeader.Difficulty,
				Number:     new(big.Int).Add(latestBlock.Number(), big.NewInt(1)),
				GasLimit:   latestHeader.GasLimit,
				GasUsed:    latestHeader.GasUsed,
				Time:       latestHeader.Time + 1,
				Extra:      latestHeader.Extra,
				MixDigest:  latestHeader.MixDigest,
				Nonce:      latestHeader.Nonce,
			},
			body,
			nil,
			hasher,
			geth_types.DefaultBlockConfig,
		),
	)
	if err != nil {
		t.Fatalf("Failed to create valid Espresso batch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		return nil, err
	}
	return espressoBatch, nil
}

// forgeBatcherPrivateKey is a helper function that forge a batcher private key
func forgedBatcherPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func realBatcherPrivateKey(system *e2esys.System) (*ecdsa.PrivateKey, error) {
	return system.Cfg.Secrets.Batcher, nil
}

// createEspressoTransaction creates a Espresso transaction with a FAKE or REAL batcher private key
func createEspressoTransaction(ctx context.Context, depositTx *geth_types.Transaction, tx *geth_types.Transaction, l2Seq *ethclient.Client, rollupCfg *rollup.Config, chainID *big.Int, batcherKey *ecdsa.PrivateKey, t *testing.T) (*espressoCommon.Transaction, error) {
	// create a valid Espresso batch first
	stackTrie := trie.NewStackTrie(func(path []byte, hash geth_common.Hash, blob []byte) {})
	batch, err := createValidEspressoBatch(ctx, depositTx, tx, l2Seq, rollupCfg, stackTrie, t)
	if err != nil {
		t.Fatalf("Failed to create valid Espresso batch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		return nil, err
	}

	// encode the batch and sign with FAKE or REAL batcher private key
	buf := new(bytes.Buffer)
	err = rlp.Encode(buf, *batch)
	if err != nil {
		t.Fatalf("Failed to encode batch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		return nil, fmt.Errorf("failed to encode batch: %w", err)
	}

	// Sign the encoded batch with FAKE or REAL batcher private key
	batcherSignature, err := crypto.Sign(crypto.Keccak256(buf.Bytes()), batcherKey)
	if err != nil {
		t.Fatalf("Failed to create batcher signature:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
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
	l2Seq := system.NodeClient(e2esys.RoleSeq)

	// create a real Espresso transaction and make sure it can go through
	{
		depositTx := geth_types.NewTx(
			&geth_types.DepositTx{
				From:                system.Cfg.Secrets.Addresses().Alice,
				Value:               big.NewInt(params.Ether),
				Gas:                 1000001,
				Data:                []byte{},
				IsSystemTransaction: false,
			},
		)

		tx := geth_types.MustSignNewTx(system.Cfg.Secrets.Bob, system.RollupConfig.L1Signer(), &geth_types.DynamicFeeTx{
			ChainID:   system.Cfg.L1ChainIDBig(),
			Nonce:     1,
			To:        &system.RollupConfig.BatchInboxAddress,
			Value:     big.NewInt(1),
			GasTipCap: big.NewInt(1 * params.GWei),
			GasFeeCap: big.NewInt(10 * params.GWei),
			Gas:       5_000_000,
		})

		// Create a real Espresso transaction
		realBatcherPrivateKey, err := realBatcherPrivateKey(system)
		if err != nil {
			t.Fatalf("Failed to get real batcher private key:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		}
		realEspressoTransaction, err := createEspressoTransaction(ctx, depositTx, tx, l2Seq, system.RollupConfig, system.Cfg.L2ChainIDBig(), realBatcherPrivateKey, t)
		if err != nil {
			t.Fatalf("Failed to create real Espresso transaction:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", err, nil)
		}

		// Send transaction directly to Espresso
		espressoTxHash, err := espressoClient.SubmitTransaction(ctx, *realEspressoTransaction)
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

		transactionFound := false
		for !transactionFound {
			select {
			case <-ticker.C:
				_, err := espressoClient.FetchTransactionByHash(ctx, espressoTxHash)
				if err == nil {
					// test pass
					transactionFound = true
				}
			case <-timer.C:
				t.Fatalf("Failed to fetch transaction by hash after multiple attempts")
			case <-ctx.Done():
				t.Fatalf("Cancelling transaction publishing")
			}
		}

		_, err = wait.ForReceiptOK(ctx, l2Seq, tx.Hash())
		if have, want := err, error(nil); have != want {
			t.Fatalf("Waiting for L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		// // Make sure the transaction will go through to caff node by checking the unmarshal works
		// // The check can directly reflect whether the transaction is valid or not
		// caffStreamer := caffNode.OpNode.EspressoStreamer()
		// _, err = caffStreamer.UnmarshalBatch(realEspressoTransaction.Payload)
		// if have, want := err, error(nil); have != want {
		// 	t.Fatalf("Failed to unmarshal batch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		// }

		// // Make sure the transaction will go through to op node by checking it will go through batch submitter's streamer
		// batchSubmitter := system.BatchSubmitter
		// _, err = batchSubmitter.EspressoStreamer().UnmarshalBatch(realEspressoTransaction.Payload)
		// if have, want := err, error(nil); have != want {
		// 	t.Fatalf("Failed to unmarshal batch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		// }

	}

}
