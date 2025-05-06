package environment_test

import (
	"context"
	"math/big"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// TestE2eDevNetWithEspressoWithCaffNodeDeterministicDerivation is a test that
// attempts to make sure that the caff node can derive the same state as the
// original op-node (non caffeinated).
//
// This tests is designed to evaluate Test 3.1 as outlined within the
// Espresso Celo Integration plan.  It has stated task definition as follows:
//
//	Arrange:
//		Running Sequencer, Batcher in Espresso mode, Caff node  OP node.
//		Balance of Alice is 0.
//		Check that this is the case querying both Caff and OP nodes
//	Act:
//		Send a single transaction that transfers 1 coin to Alice
//	Assert:
//		Query the Caff node to check that Alice balance has been increased by 1
//		Query the OP node to check that Alice balance has been increased by 1
//
// The actual tests is unable to make Alice's initial balance zero, and will
// instead just check Alice's starting balance against the rest of the cases.
func TestE2eDevNetWithEspressoWithCaffNodeDeterministicDerivation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, env.WithL1FinalizedDistance(0))
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
	var balanceAliceInitial *big.Int

	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	caffVerif := system.NodeClient(env.RoleCaffNode)

	// Retrieve Alice's starting Balance, and verify that they match between
	// the Verification Node, and the Caff Node
	{
		verifBalance, err := l2Verif.BalanceAt(ctx, addressAlice, nil)
		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to get alice's balance from verification node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}
		caffBalance, err := caffVerif.BalanceAt(ctx, addressAlice, nil)
		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to get alice's balance from caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}
		if have, want := verifBalance, caffBalance; have.Cmp(want) != 0 {
			t.Fatalf("alice's balance does not match between verification node and caff node:\nhave:\n\t\"%s\"\nwant:\n\t\"%s\"\n", have, want)
		}

		balanceAliceInitial = verifBalance
	}

	// Next We want to Increase Alice's balance by 1, and verify that the balance
	// matches between the verification node and the caff node
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

		verifBalanceNew, err := wait.ForBalanceChange(ctx, l2Verif, addressAlice, balanceAliceInitial)
		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to get alice's new balance from verification node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}
		caffBalanceNew, err := wait.ForBalanceChange(ctx, caffVerif, addressAlice, balanceAliceInitial)
		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to get alice's new balance from caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		if have, want := verifBalanceNew, caffBalanceNew; have.Cmp(want) != 0 {
			t.Fatalf("alice's new balance does not match between verification node and caff node:\nhave:\n\t\"%s\"\nwant:\n\t\"%s\"\n", have, want)
		}

		// We have a new balance, and it matches between the verification node
		// and the Caff Node.

		// Let's check to make sure that Alice's balance has increased by
		// exactly 1.

		diff := new(big.Int).Sub(verifBalanceNew, balanceAliceInitial)
		if have, want := diff, mintAmount; have.Cmp(want) != 0 {
			t.Fatalf("alice's balance did not increase by 1:\nhave:\n\t\"%s\"\nwant:\n\t\"%s\"\n", have, want)
		}

	}

}
