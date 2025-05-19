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
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// Time to wait for a transaction to be enforced after submission.
const WAIT_FORCED_TXN_TIME = 25 * time.Second

// Window small enough to guarantee that transactions are enforced before WAIT_FORCED_TXN_TIME.
const SMALL_SEQUENCER_WINDOW = 2 // Minimum possible value

// Window large enough to guarantee that transactions are not enforced before WAIT_FORCED_TXN_TIME.
const LARGER_SEQUENCER_WINDOW = 1000

// Get the appropriate sequencer window size for testing.
func sequencer_window_size(withSmallWindow bool) uint64 {
	if withSmallWindow {
		return SMALL_SEQUENCER_WINDOW
	}
	return LARGER_SEQUENCER_WINDOW
}

// ForcedTransaction attempts to verify that the forced transaction mechanism works with and
// without Espresso dev node.
//
// This function is designed to evaluate Test 11 as outlined within the Espresso Celo Integration
// plan. It has stated task definition as follows:
//
//	Arrange:
//		Set the sequencer window size small or large.
//		Start the devnet with the sequencer window setting, with or without the Espresso dev node.
//		Stop the sequencer.
//	Act:
//		Send a deposit and wait until the small window is passed.
//		Send a withdrawal and wait until the small window is passed.
//	Assert:
//		The balance reflects (or does not) reflect the deposit transaction and the withdrawal
//		transaction is (or is) succeeds, if the sequencer window is set small (or large,
//		respectively), regardless of whether launching with the Espresso dev node.
func ForcedTransaction(t *testing.T, withSmallSequencerWindow bool, withEspresso bool) {
	// Set up the test timeout condition.
	// Extended timeout to accommodate slower processing in test environments
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Launch the devnet with the given sequencer window size.
	var system *e2esys.System
	var err error
	if withEspresso {
		launcher := new(env.EspressoDevNodeLauncherDocker)
		systemWithEspresso, espressoDevNode, err := launcher.StartDevNet(ctx, t, env.WithSequencerWindowSize(sequencer_window_size(withSmallSequencerWindow)))
		system = systemWithEspresso
		require.NoError(t, err, "Failed to launch with the Espresso dev node")
		defer env.Stop(t, system)
		defer env.Stop(t, espressoDevNode)
	} else {
		sysConfig := e2esys.DefaultSystemConfig(t, e2esys.WithAllocType(config.AllocTypeStandard))
		sysConfig.DeployConfig.SequencerWindowSize = sequencer_window_size(withSmallSequencerWindow)
		system, err = sysConfig.Start(t)
		require.NoError(t, err, "failed to launch without the Espresso dev node")
		defer env.Stop(t, system)
	}

	// Retrieve L1 and L2 clients.
	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	// Set up Alice's address and record the initial balance.
	address := system.Cfg.Secrets.Addresses().Alice
	initialBalance, err := l2Verif.BalanceAt(ctx, address, nil)
	require.NoError(t, err, "Failed to get the initial balance")

	// Simulate sequencer downtime.
	err = system.RollupNodes["sequencer"].Stop(ctx)
	require.NoError(t, err, "Failed to stop the sequencer")

	// Send a deposit from Bob to Alice without waiting for the receipt.
	bobPrivateKey := system.Cfg.Secrets.Bob
	options, err := bind.NewKeyedTransactorWithChainID(bobPrivateKey, system.Cfg.L1ChainIDBig())
	require.NoError(t, err, "Failed to create deposit transaction options")
	depositAmount := new(big.Int).SetUint64(100000)
	options.Value = depositAmount
	env.SendDepositTxNoReceipt(t, system.Cfg, l1Client, l2Verif, options, func(l2Opts *helpers.DepositTxOpts) {
		l2Opts.ToAddr = address
	})

	// Wait and attempt to get the new balance after the deposit.
	time.Sleep(WAIT_FORCED_TXN_TIME)
	balanceAfterDeposit, err := wait.ForBalanceChange(ctx, l2Verif, address, initialBalance)

	if withSmallSequencerWindow {
		// Verify that Alice's balance increases as expected.
		require.NoError(t, err, "Failed to get the new balance")
		require.Equal(t, new(big.Int).Add(initialBalance, depositAmount), balanceAfterDeposit, "Incorrect balance after deposit")
	} else {
		// Verify that Alice's balance is inaccessible.
		require.Error(t, err, "Not expected to get the new balance")
	}

	// Send a withdrawal from Alice to L2ToL1MessagePasser.
	alicePrivateKey := system.Cfg.Secrets.Alice
	l2ToL1MessagePasserAddr := common.HexToAddress(predeploys.L2ToL1MessagePasser)
	gasPrice, err := l2Verif.SuggestGasPrice(ctx)
	require.NoError(t, err, "Failed to get gas price")
	nonce, err := l2Verif.PendingNonceAt(ctx, address)
	require.NoError(t, err, "Failed to get nonce")
	withdrawalAmount := new(big.Int).SetUint64(1000)
	tx := types.NewTransaction(
		nonce,
		l2ToL1MessagePasserAddr,
		withdrawalAmount,
		300000,
		gasPrice,
		nil,
	)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(system.Cfg.L2ChainIDBig()), alicePrivateKey)
	require.NoError(t, err, "Failed to sign the withdrawal transaction")
	err = l2Verif.SendTransaction(ctx, signedTx)
	require.NoError(t, err, "Failed to send the withdrawal transaction")

	// Wait and attempt to get the new balance after the withdrawal.
	// TODO (Keyao) The balance check below fails due to unchanged balance, not resolved by
	// increasing the wait time.
	time.Sleep(WAIT_FORCED_TXN_TIME)
	balanceAfterWithdrawal, err := wait.ForBalanceChange(ctx, l2Verif, address, balanceAfterDeposit)

	if withSmallSequencerWindow {
		// Verify that Alice's balance decreases as expected.
		require.NoError(t, err, "Failed to get the new balance")
		require.Equal(t, new(big.Int).Sub(balanceAfterDeposit, withdrawalAmount), balanceAfterWithdrawal, "Incorrect balance after withdrawal")
	} else {
		// Verify that Alice's balance is inaccessible.
		require.Error(t, err, "Not expected to get the new balance")
	}
}

// TestForcedTransactionsWithoutEspressoSmallWindow verifies that the deposit and the withdrawal
// transactions are enforced after the sequencer window is passed when launching without the
// Espressso dev node.
func TestForcedTransactionsWithoutEspressoSmallWindow(t *testing.T) {
	ForcedTransaction(t, true, false)
}

// TODO (Keyao) Restore the following tests once TestForcedTransactionsWithoutEspressoSmallWindow
// passes.

// // TestForcedTransactionsWithoutEspressoLargeWindow verifies that the deposit and the withdrawal
// // transactions are not enforced before the sequencer window is passed when launching without the
// // Espressso dev node.
// func TestForcedTransactionsWithoutEspressoLargeWindow(t *testing.T) {
// 	ForcedTransaction(t, false, false)
// }

// // TestForcedTransactionsWithEspressoSmallWindow verifies that the deposit and the withdrawal
// // transactions are enforced after the sequencer window is passed when launching with the Espressso
// // dev node.
// func TestForcedTransactionsWithEspressoSmallWindow(t *testing.T) {
// 	ForcedTransaction(t, true, true)
// }

// // TestForcedTransactionsWithEspressoLargeWindow verifies that the deposit and the withdrawal
// // transactions are not enforced before the sequencer window is passed when launching with the
// // Espressso dev node.
// func TestForcedTransactionsWithEspressoLargeWindow(t *testing.T) {
// 	ForcedTransaction(t, false, true)
// }
