package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

func TestForcedTransaction(t *testing.T) {
	// Basic setup.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	launcher := new(env.EspressoDevNodeLauncherDocker)

	// Set a small WithSeqWindowSize for faster testing.
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, env.WithSequencerWindowSize(2))
	require.NoError(t, err, "failed to start dev environment with espresso dev node")
	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	// Retrieve L1 and L2 clients.
	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	// Set up addresses and record initial balance of Alice on the L2.
	addressAlice := system.Cfg.Secrets.Addresses().Alice
	initialBalance, err := l2Verif.BalanceAt(ctx, addressAlice, nil)
	require.NoError(t, err, "failed to get the initial balance")

	// Simulate sequencer downtime.
	err = system.RollupNodes["sequencer"].Stop(ctx)
	require.NoError(t, err, "failed to stop the sequencer")

	// Send the deposit without waiting for the receipt because there will not be immediate receipt.
	privateKey := system.Cfg.Secrets.Bob
	bobOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, system.Cfg.L1ChainIDBig())
	require.NoError(t, err, "failed to create transaction options")
	amount := new(big.Int).SetUint64(1)
	bobOptions.Value = amount
	env.SendDepositTxNoReceipt(t, system.Cfg, l1Client, l2Verif, bobOptions, func(l2Opts *helpers.DepositTxOpts) {
		// Send from Bob to Alice
		l2Opts.ToAddr = addressAlice
	})

	// Wait for the L2 chain to process the deposit
	time.Sleep(10 * time.Second)

	// Record the new balance of Alice on L2
	newBalance, err := wait.ForBalanceChange(ctx, l2Verif, addressAlice, initialBalance)
	require.NoError(t, err, "failed to get the new balance")

	// Verify that Alice's balance increased by the deposited amount
	require.Equal(t, new(big.Int).Add(initialBalance, amount), newBalance, "Alice's balance should have increased by the deposited amount")
}
