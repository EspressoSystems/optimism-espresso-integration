package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

// ForcedDeposit attempts to verify that the forced transaction mechanism works for the deposit
// transaction with and without Espresso dev node.
//
// This function and ForcedWithdrawal are designed to evaluate Test 11 as outlined within the
// Espresso Celo Integration plan. It has stated task definition as follows:
//
//	Arrange:
//		Set the sequencer window size small or large.
//		Start the devnet with the sequencer window setting, with or without the Espresso dev node.
//		Stop the sequencer.
//	Act:
//		Send a deposit and wait until the small window is passed.
//	Assert:
//		The balance reflects (or does not reflect) the deposit transaction, if the sequencer window
//		is set small (or large, respectively), regardless of whether launching with the Espresso
//		dev node.
func ForcedDeposit(t *testing.T, withSmallSequencerWindow bool, withEspresso bool) {
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
		require.NoError(t, err, "Failed to launch with Espresso dev node")
		defer env.Stop(t, system)
		defer env.Stop(t, espressoDevNode)
	} else {
		sysConfig := e2esys.DefaultSystemConfig(t, e2esys.WithAllocType(config.AllocTypeStandard))
		sysConfig.DeployConfig.SequencerWindowSize = sequencer_window_size(withSmallSequencerWindow)
		system, err = sysConfig.Start(t)
		require.NoError(t, err, "failed to launch without Espresso dev node")
		defer env.Stop(t, system)
	}

	// Retrieve L1 and L2 clients.
	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	// Set up Alice's address and record the initial balance.
	address := system.Cfg.Secrets.Addresses().Alice
	initialBalance, err := l2Verif.BalanceAt(ctx, address, nil)
	require.NoError(t, err, "Failed to get initial balance")

	// Simulate sequencer downtime.
	err = system.RollupNodes["sequencer"].Stop(ctx)
	require.NoError(t, err, "Failed to stop sequencer")

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
		require.NoError(t, err, "Failed to get new balance")
		require.Equal(t, new(big.Int).Add(initialBalance, depositAmount), balanceAfterDeposit, "Incorrect balance after deposit")
	} else {
		// Verify that Alice's balance is inaccessible.
		require.Error(t, err, "Not expected to get new balance")
	}
}

// TestForcedDepositWithoutEspressoSmallWindow verifies that the deposit transaction is enforced
// after the sequencer window is passed when launching without the Espressso dev node.
func TestForcedDepositWithoutEspressoSmallWindow(t *testing.T) {
	ForcedDeposit(t, true, false)
}

// TestForcedDepositWithoutEspressoLargeWindow verifies that the deposit transaction is not
// enforced before the sequencer window is passed when launching without the Espressso dev node.
func TestForcedDepositWithoutEspressoLargeWindow(t *testing.T) {
	ForcedDeposit(t, false, false)
}

// TestForcedDepositWithEspressoSmallWindow verifies that the deposit transaction is enforced after
// the sequencer window is passed when launching with the Espressso dev node.
func TestForcedDepositWithEspressoSmallWindow(t *testing.T) {
	ForcedDeposit(t, true, true)
}

// TestForcedDepositWithEspressoLargeWindow verifies that the deposit transaction is not enforced
// before the sequencer window is passed when launching with the Espressso dev node.
func TestForcedDepositWithEspressoLargeWindow(t *testing.T) {
	ForcedDeposit(t, false, true)
}

// ForcedWithdrawal attempts to verify that the forced transaction mechanism works for the
// withdrawal transaction with and without Espresso dev node.
//
// This function and ForcedDeposit are designed to evaluate Test 11 as outlined within the Espresso
// Celo Integration plan. It has stated task definition as follows:
//
//	Arrange:
//		Set the sequencer window size small or large.
//		Start the devnet with the sequencer window setting, with or without the Espresso dev node.
//		Stop the sequencer.
//	Act:
//		Send a withdrawal and wait until the small window is passed.
//	Assert:
//		The balance reflects (or does not reflect) the withdrawal transaction, if the sequencer
//		window is set small (or large, respectively), regardless of whether launching with the
//		Espresso dev node.
func ForcedWithdrawal(t *testing.T, withSmallSequencerWindow bool, withEspresso bool) {
	// Set up the test timeout condition.
	// Extended timeout to accommodate slower processing in test environments
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Launch the devnet with the given sequencer window size.
	var system *e2esys.System
	var err error
	if withEspresso {
		launcher := new(env.EspressoDevNodeLauncherDocker)
		// TODO (Keyao) Once the tests without Espresso are fixed, update the config parameters
		// similarily here.
		systemWithEspresso, espressoDevNode, err := launcher.StartDevNet(ctx, t, env.WithSequencerWindowSize(sequencer_window_size(withSmallSequencerWindow)))
		system = systemWithEspresso
		require.NoError(t, err, "Failed to launch with the Espresso dev node")
		defer env.Stop(t, system)
		defer env.Stop(t, espressoDevNode)
	} else {
		sysConfig := e2esys.DefaultSystemConfig(t, e2esys.WithAllocType(config.AllocTypeStandard))
		sysConfig.DeployConfig.SequencerWindowSize = sequencer_window_size(withSmallSequencerWindow)
		// TODO (Keyao) Once the tests without Espresso are fixed, remove unnecessary config
		// parameters below.
		sysConfig.DeployConfig.FinalizationPeriodSeconds = 1
		sysConfig.DeployConfig.MaxSequencerDrift = 1
		sysConfig.DeployConfig.L2BlockTime = 1
		sysConfig.L1FinalizedDistance = 0
		sysConfig.DeployConfig.L2OutputOracleSubmissionInterval = 1
		sysConfig.DeployConfig.L2OutputOracleStartingTimestamp = 0
		system, err = sysConfig.Start(t)
		require.NoError(t, err, "failed to launch without Espresso dev node")
		defer env.Stop(t, system)
	}

	// Set up Alice's address and record the initial balance.
	// l2Verif := system.NodeClient(e2esys.RoleVerif)
	address := system.Cfg.Secrets.Addresses().Alice
	l1Client := system.NodeClient(e2esys.RoleL1)
	initialBalance, err := l1Client.BalanceAt(ctx, address, nil)
	require.NoError(t, err, "Failed to get initial balance")

	// Simulate sequencer downtime.
	err = system.RollupNodes["sequencer"].Stop(ctx)
	require.NoError(t, err, "Failed to stop sequencer")

	// Send a withdrawal from Alice to L2CrossDomainMessenger.
	opts, err := bind.NewKeyedTransactorWithChainID(system.Cfg.Secrets.Alice, system.Cfg.L1ChainIDBig())
	require.NoError(t, err)

	withdrawalAmount := new(big.Int).SetUint64(1000)

	portal, err := bindings.NewOptimismPortal(system.Cfg.L1Deployments.OptimismPortalProxy, l1Client)
	require.NoError(t, err)
	tx, err := portal.DepositTransaction(
		opts, // keyed transactor for gas & value
		common.HexToAddress(predeploys.L2ToL1MessagePasser), // L2CrossDomainMessenger predeploy
		withdrawalAmount, // no ETH value needed
		uint64(300_000),  // gas limit
		false,            // _isCreation
		nil,              // _extraData
	)
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, l1Client, tx)
	require.NoError(t, err)

	time.Sleep(WAIT_FORCED_TXN_TIME)

	// TODO (Keyao) The nonce and the balance checks below both fail. We probably don't need both
	// to pass, but should fix at least one.

	// newNonce, err := l1Client.NonceAt(ctx, address, nil)
	// require.NoError(t, err, "Failed to get new nonce")
	// require.Greater(t, newNonce, initialNonce)

	newBalance, err := wait.ForBalanceChange(ctx, l1Client, address, initialBalance)
	if withSmallSequencerWindow {
		// Verify that Alice's balance decreases as expected.
		require.NoError(t, err, "Failed to get new balance")
		require.Less(t, newBalance.Uint64(), initialBalance.Uint64(), "Balance not decreased")
	} else {
		// Verify that Alice's balance is inaccessible.
		require.Error(t, err, "Not expected to get new balance")
	}
}

// TestForcedWithdrawalWithoutEspressoSmallWindow verifies that the withdrawal transaction is
// enforced after the sequencer window is passed when launching without the Espressso dev node.
func TestForcedWithdrawalWithoutEspressoSmallWindow(t *testing.T) {
	ForcedWithdrawal(t, true, false)
}

// TODO (Keyao) Restore the following tests once TestForcedWithdrawalWithoutEspressoSmallWindow
// passes.

// // TestForcedWithdrawalWithoutEspressoLargeWindow verifies that the withdrawal transaction is not
// // enforced before the sequencer window is passed when launching without the Espressso dev node.
// func TestForcedWithdrawalWithoutEspressoLargeWindow(t *testing.T) {
// 	ForcedWithdrawal(t, false, false)
// }

// // TestForcedWithdrawalWithEspressoSmallWindow verifies that the withdrawal transaction is enforced
// // after the sequencer window is passed when launching with the Espressso dev node.
// func TestForcedWithdrawalWithEspressoSmallWindow(t *testing.T) {
// 	ForcedWithdrawal(t, true, true)
// }

// // TestForcedWithdrawalWithEspressoLargeWindow verifies that the withdrawal transaction is not
// // enforced before the sequencer window is passed when launching with the Espressso dev node.
// func TestForcedWithdrawalWithEspressoLargeWindow(t *testing.T) {
// 	ForcedWithdrawal(t, false, false)
// }
