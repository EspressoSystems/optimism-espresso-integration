package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// TestEspressoDockerDevNodeSmokeTest is a smoke test for the Espresso Dev Node
// Docker implementation. It starts the dev node and then stops it. And tries
// to ensure that the e2e system, and the docker container stop correctly.
func TestEspressoDockerDevNodeSmokeTest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer env.Stop(t, espressoDevNode, env.IgnoreStopErrors)
	defer env.Stop(t, system, env.IgnoreStopErrors)

	{
		// Stop the Docker Container
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		espressoClose := make(chan struct{})

		var err error

		go (func(ch chan struct{}) {
			err = espressoDevNode.Stop()
			close(ch)
		})(espressoClose)

		select {
		case <-ctx.Done():
			t.Errorf("espresso dev node failed to stop in the anticipated time given: %v", ctx.Err())
		case <-espressoClose:
			// Espresso Dev Node stopped in the anticipated time
			if err != nil {
				t.Fatalf("failed to stop espresso dev node: %v", err)
			}
		}

		// One last sanity check to ensure that the container is not still
		// running.

		err = espressoDevNode.Stop()
		if err == nil {
			t.Fatalf("espresso dev node should return an error indicating that it cannot be stopped, as it is not running")
		}

		if _, castOk := err.(env.DockerContainerNotRunningError); !castOk {
			t.Fatalf("espresso dev node should return a DockerContainerNotRunningError, but received: %v", err)
		}
	}

	{
		// Stop the e2e system
		sysClose := make(chan struct{})

		go (func(ch chan struct{}) {
			system.Close()
			close(ch)
		})(sysClose)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		select {
		case <-ctx.Done():
			t.Errorf("system failed to close in the anticipated time given: %v", ctx.Err())

		case <-sysClose:
			// System closed in the anticipated time
		}
	}
}

// runSimpleL1TransferAndVerifier runs a simple L1 transfer and verifies it on
// the L2 Verifier.
func runSimpleL1TransferAndVerifier(ctx context.Context, t *testing.T, system *e2esys.System) {
	privateKey := system.Cfg.Secrets.Bob

	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	fromAddress := system.Cfg.Secrets.Addresses().Bob

	// Send Transaction on L1, and wait for verification on the L2 Verifier
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	// Get the Starting Balance of the Address
	startBalance, err := l2Verif.BalanceAt(ctx, fromAddress, nil)
	if have, want := err, error(nil); have != want {
		t.Errorf("attempt to get starting balance for %s failed:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", fromAddress, have, want)
	}

	// Create a new Keyed Transaction
	options, err := bind.NewKeyedTransactorWithChainID(privateKey, system.Cfg.L1ChainIDBig())
	if have, want := err, error(nil); have != want {
		t.Errorf("attempt to get keyed transaction with chain ID %d failed:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", system.Cfg.L1ChainIDBig(), have, want)
	}

	if err == nil {
		// We can only continue with these tests if the error above was nil

		// Send a Deposit Transaction
		mintAmount := big.NewInt(1_000_000_000_000)
		options.Value = mintAmount
		_ = helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, options, nil)

		endBalance, err := wait.ForBalanceChange(ctx, l2Verif, fromAddress, startBalance)
		if have, want := err, error(nil); have != want {
			t.Errorf("waiting for balance change returned with error:\nhave:\n\t\"%v\"\nwant:\t\n\"%v\"\n", have, want)
		}

		diff := new(big.Int).Sub(endBalance, startBalance)
		if have, want := diff, mintAmount; have.Cmp(want) != 0 {
			t.Errorf("balance change does not match mint amount:\nhave;\n\t\"%s\"\nwant:\n\t\"%s\"\n", have, want)
		}
	}

	cancel()
}

// runSimpleL2Burn runs a simple L2 burn transaction and verifies it on the
// L2 Verifier.
func runSimpleL2Burn(ctx context.Context, t *testing.T, system *e2esys.System) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	privateKey := system.Cfg.Secrets.Bob

	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	amountToBurn := big.NewInt(500_000_000)
	burnAddress := common.Address{0xff, 0xff}
	_ = helpers.SendL2Tx(
		t,
		system.Cfg,
		l2Seq,
		privateKey,
		env.L2TxWithOptions(
			env.L2TxWithAmount(amountToBurn),
			env.L2TxWithNonce(1), // Already have deposit
			env.L2TxWithToAddress(&burnAddress),
			env.L2TxWithVerifyOnClients(l2Verif),
		),
	)

	// Check the balance of hte burn address using the L2 Verifier
	balanceBurned, err := wait.ForBalanceChange(ctx, l2Verif, burnAddress, big.NewInt(0))
	if have, want := err, error(nil); have != want {
		t.Errorf("wait for balance change for burn address %s failed:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", burnAddress, have, want)
	}

	// Make sure that these match
	if have, want := balanceBurned, amountToBurn; have.Cmp(want) != 0 {
		t.Errorf("balance of burn address does not match amount burned:\nhave:\n\t\"%s\"\nwant:\n\t\"%s\"\n", have, want)
	}

	cancel()
}

// TestE2eDevNetWithEspressoSimpleTransactions launches the e2e Dev Net with the Espresso Dev Node
// and runs a couple of simple transactions to it.
func TestE2eDevNetWithEspressoSimpleTransactions(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Signal the testnet to shut down on exit
	defer env.Stop(t, espressoDevNode)
	defer env.Stop(t, system)
	// Send Transaction on L1, and wait for verification on the L2 Verifier
	runSimpleL1TransferAndVerifier(ctx, t, system)

	// Submit a Transaction on the L2 Sequencer node, to a Burn Address
	runSimpleL2Burn(ctx, t, system)

}

// TestE2eDevNetWithoutEspressoSimpleTransactions launches the e2e Dev Net
// without the Espresso Dev Node and runs a couple of simple transactions to it.
func TestE2eDevNetWithoutEspressoSimpleTransaction(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sysConfig := e2esys.DefaultSystemConfig(t, e2esys.WithAllocType(config.AllocTypeStandard))

	system, err := sysConfig.Start(t)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start e2e dev environment:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}
	// Shut down the test net on exit
	defer env.Stop(t, system)

	// Send Transaction on L1, and wait for verification on the L2 Verifier
	runSimpleL1TransferAndVerifier(ctx, t, system)

	// Submit a Transaction on the L2 Sequencer node, to a Burn Address
	runSimpleL2Burn(ctx, t, system)
}
