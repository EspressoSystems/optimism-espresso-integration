package environment_test

import (
	"context"
	"math/big"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	geth_types "github.com/ethereum/go-ethereum/core/types"
)

// TestDeterministicDerivationExecutionState is a test that
// attempts to make sure that the caff node can derive the same state as the
// original op-node (non caffeinated).
//
// This test is designed to evaluate Test 3.2 as outlined within the
// Espresso Celo Integration plan.  It has stated task definition as follows:
//
//	Arrange:
//		Running Sequencer, Batcher in Espresso mode, Caff node, and OP node.
//	Act:
//		Send some transactions from Bob to Alice
//	Assert:
//		Once a state of op-node is finalized on L1, it should match the state that was earlier reported by the caff-node for the same block.
//		Query the executive machine state when Caff node is on
//		Query the executive machine state when OP node is on
//		Make sure the states are the same

func TestDeterministicDerivationExecutionState(t *testing.T) {
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

	// Get L2Client from caff node's engine state
	caffNodeL2Client := caffNode.OpNode.EngineState()

	numIterations := 10
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

		// Get latest blocks from caff node first as caff node usually has bigger overhead
		// We use l2BlockRefByLabel to get the states as the engine state will be reflected in the block
		caffBlock, err := caffNodeL2Client.L2BlockRefByLabel(ctx, eth.Finalized)
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
