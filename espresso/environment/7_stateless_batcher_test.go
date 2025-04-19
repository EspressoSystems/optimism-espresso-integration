package environment_test

import (
	"context"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

// TODO Phlippe make the description of the test more detailed using the AAA pattern

// TestStatelessBatcher is a test that verifies a batcher can operate (especially restart) correctly and efficiently without persistent storage.
//
// This tests is designed to evaluate Test 7 as outlined within the
// Espresso Celo Integration plan.  It has stated task definition as follows:
// Run the rollup and randomly restart the batcher. Check the liveness of the rollup, and the consistency of Espresso confirmations and L1 confirmations.
// We don't need to clear persistent storage because the original Optimism code isn't and our integration work shouldn't use any.

func TestStatelessBatcher(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t)
	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer system.Close()
	//defer espressoDevNode.Stop()

	caffNode, err := env.LaunchDecaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer caffNode.Close(ctx)

	addressAlice := system.Cfg.Secrets.Addresses().Alice

	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	caffVerif := system.NodeClient(env.RoleCaffNode)

	balanceAliceInitial, err := l2Verif.BalanceAt(ctx, addressAlice, nil)
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

	//driver := system.BatchSubmitter.TestDriver()

	// We select a range of iterations when the batcher is turned off.
	var rangeBatcherDown [2]int
	rangeBatcherDown[0] = 2 //rand.IntN(5)     // Random number between 0 and 4
	rangeBatcherDown[1] = 4 //rand.IntN(5) + 5 // Random number between 5 and 9

	batcherIsUp := true
	for i := 0; i < 10; i++ {

		t.Log("******************* Iteration: ", i)
		// Let us stop the batcher
		//if i == rangeBatcherDown[0] {
		//
		//	err = driver.StopBatchSubmitting(ctx)
		//	require.NoError(t, err)
		//	time.Sleep(2 * time.Second)
		//	batcherIsUp = false
		//}
		//
		//// Let us start the batcher again
		//if i == rangeBatcherDown[1] {
		//	driver.StartBatchSubmitting()
		//	batcherIsUp = true
		//}

		// The batcher is up, we can send coins
		if batcherIsUp {
			_ = helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, bobOptions, func(l2Opts *helpers.DepositTxOpts) {
				// Send from Bob to Alice
				l2Opts.ToAddr = addressAlice
			})
			numDeposits++
		}

	}

	var numDepositsBigInt big.Int
	numDepositsBigInt.SetInt64(int64(numDeposits))

	expectedAmount := new(big.Int).Mul(new(big.Int).Add(balanceAliceInitial, &numDepositsBigInt), amount)

	// TODO this is not very robust. Should use functions like wait.ForBalanceChange
	caffBalanceNew, _ = caffVerif.BalanceAt(ctx, addressAlice, nil)
	l2BalanceNew, _ := l2Verif.BalanceAt(ctx, addressAlice, nil)

	assert.Equal(t, expectedAmount, caffBalanceNew)
	assert.Equal(t, expectedAmount, l2BalanceNew)

}
