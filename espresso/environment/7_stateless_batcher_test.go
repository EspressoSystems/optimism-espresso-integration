package environment_test

import (
	"context"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"math/rand/v2"
	"testing"
	"time"
)

// TestStatelessBatcher is a test that verifies a batcher can operate (especially restart) correctly and efficiently without persistent storage.
//
// This tests is designed to evaluate Test 7 as outlined within the
// Espresso Celo Integration plan.  It has stated task definition as follows:
// Run the rollup and randomly restart the batcher. Check the liveness of the rollup, and the consistency of Espresso confirmations and L1 confirmations.
// We don't need to clear persistent storage because the original Optimism code isn't and our integration work shouldn't use any.
// More specifically the test is defined as follows
//	Arrange:
//		Running Sequencer, Batcher in Espresso mode, Caff node  OP node.
//	Act:
//		Loop over n iterations
//      Randomly pick one iteration to stop the batcher and another to start the batcher
//      For all the other iterations send one coin to Alice.
//	Assert:
//		Query the Caff node to check that Alice balance has been increased by n-2
//		Query the OP node to check that Alice balance has been increased by n-2

func TestStatelessBatcher(t *testing.T) {
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

	addressAlice := system.Cfg.Secrets.Addresses().Alice

	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	caffVerif := system.NodeClient(env.RoleCaffNode)

	balanceAliceInitial, err := l2Verif.BalanceAt(ctx, addressAlice, nil)
	if have, want := err, error(nil); have != want {
		t.Fatalf("Failed to fetch Alice's balance:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Setup Bob for sending coins to Alice
	privateKey := system.Cfg.Secrets.Bob
	bobOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, system.Cfg.L1ChainIDBig())
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to create transaction options for bob:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	amount := new(big.Int).SetUint64(1)
	numDeposits := 0
	bobOptions.Value = amount

	var caffBalanceNew *big.Int

	driver := system.BatchSubmitter.TestDriver()
	numIterations := 10

	// We select a range of iterations when the batcher is turned off.
	turnBatcherOffIteration := rand.IntN(numIterations / 2)
	turnBatcherOnIteration := rand.IntN(numIterations/2) + numIterations/2

	batcherIsUp := true
	for i := 0; i < numIterations; i++ {

		t.Log("******************* Iteration: ", i)
		//Let us stop the batcher
		if i == turnBatcherOffIteration {
			err = driver.StopBatchSubmitting(ctx)
			require.NoError(t, err)
			time.Sleep(2 * time.Second)
			batcherIsUp = false
		}

		// Let us start the batcher again
		if i == turnBatcherOnIteration {
			err = driver.StartBatchSubmitting()
			require.NoError(t, err)
			batcherIsUp = true
		}

		// The batcher is up, we can send coins
		if batcherIsUp {
			receipt := helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, bobOptions, func(l2Opts *helpers.DepositTxOpts) {
				// Send from Bob to Alice
				l2Opts.ToAddr = addressAlice
			})
			t.Log("Deposit transaction receipt", "receipt", receipt)
			numDeposits++
		}

	}

	var numDepositsBigInt big.Int
	numDepositsBigInt.SetInt64(int64(numDeposits))

	expectedAmount := new(big.Int).Mul(new(big.Int).Add(balanceAliceInitial, &numDepositsBigInt), amount)

	caffBalanceNew, _ = caffVerif.BalanceAt(ctx, addressAlice, nil)
	l2BalanceNew, _ := l2Verif.BalanceAt(ctx, addressAlice, nil)

	assert.Equal(t, expectedAmount, l2BalanceNew)
	assert.Equal(t, expectedAmount, caffBalanceNew)
}
