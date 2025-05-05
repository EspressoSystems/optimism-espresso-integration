package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
//	Act:
//		Send some transactions from Bob to Alice
//	Assert:
//		Query the executive machine state when Caff node is on
//		Query the executive machine state when OP node is on

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
	caffNodeL2Client := caffNode.OpNode.EngineState().L2Client

	// Compare states between nodes
	for i := 0; i < 10; i++ {
		// Get latest blocks from each node
		seqBlock, err := l2Seq.BlockByNumber(ctx, nil)
		if err != nil {
			t.Fatalf("failed to get block from l2Seq: %v", err)
		}

		caffBlock, err := caffNodeL2Client.InfoByLabel(ctx, eth.Unsafe)
		if err != nil {
			t.Fatalf("failed to get block from caff node: %v", err)
		}

		// Compare block states
		t.Logf("Block number %v:", seqBlock.Number())
		t.Logf("Block number of caffBlock %v:", caffBlock.NumberU64())
		t.Logf("  Sequencer    hash: %v", seqBlock.Hash())
		t.Logf("  Caff node    hash: %v", caffBlock.Hash())
		t.Logf("  State roots:")
		t.Logf("    Sequencer: %v", seqBlock.Root())
		t.Logf("    Caff node: %v", caffBlock.Root())

		// Verify state consistency
		if seqBlock.Hash() != caffBlock.Hash() {
			t.Errorf("block hash mismatch between sequencer and caff node at block %v", seqBlock.Number())
		}
		if seqBlock.Root() != caffBlock.Root() {
			t.Errorf("state root mismatch between sequencer and caff node at block %v", seqBlock.Number())
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
		}
	}
	// Sishan TODO: SendL2Tx instead of DepositTx

}
