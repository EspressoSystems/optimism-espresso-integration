package environment_test

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/config"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestBatcherSwitching(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// We will need this config to start a new instance of "TEE" batcher
	// with parameters tweaked.
	batcherConfig := &batcher.CLIConfig{}
	system, espressoDevNode, err := launcher.StartE2eDevnet(ctx, t, env.WithSequencerUseFinalized(true), env.GetBatcherConfig(batcherConfig))
	require.NoError(t, err)

	l1Client := system.NodeClient(e2esys.RoleL1)
	verifClient := system.NodeClient(e2esys.RoleVerif)
	espClient := espressoClient.NewClient(espressoDevNode.EspressoUrls()[0])

	deployerTransactor, err := bind.NewKeyedTransactorWithChainID(system.Config().Secrets.Deployer, system.Cfg.L1ChainIDBig())
	require.NoError(t, err)

	batchAuthenticator, err := bindings.NewBatchAuthenticator(system.RollupConfig.BatchAuthenticatorAddress, l1Client)
	require.NoError(t, err)

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	// Send Transaction on L1, and wait for verification on the L2 Verifier
	env.RunSimpleL1TransferAndVerifier(ctx, t, system)

	// Verify everything works
	env.RunSimpleL2Burn(ctx, t, system)

	// Stop the "TEE" batcher
	err = system.BatchSubmitter.TestDriver().StopBatchSubmitting(ctx)
	require.NoError(t, err)

	// Switch active batcher
	tx, err := batchAuthenticator.SwitchBatcher(deployerTransactor)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, l1Client, tx.Hash())
	require.NoError(t, err)

	// Start the fallback batcher
	err = system.FallbackBatchSubmitter.TestDriver().StartBatchSubmitting()
	require.NoError(t, err)

	// Everything should still work
	env.RunSimpleL2Burn(ctx, t, system)

	// Stop the fallback batcher
	err = system.FallbackBatchSubmitter.TestDriver().StopBatchSubmitting(ctx)
	require.NoError(t, err)

	// Switch batcher back to the "TEE" batcher
	tx, err = batchAuthenticator.SwitchBatcher(deployerTransactor)
	require.NoError(t, err)
	switchReceipt, err := wait.ForReceiptOK(ctx, l1Client, tx.Hash())
	require.NoError(t, err)

	// Give things time to settle
	var l2Height uint64

	ticker := time.NewTicker(100 * time.Millisecond)
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

Loop:
	for {
		select {
		case <-timeoutCtx.Done():
			panic("Timeout waiting for verifier derivation pipeline to advance past the fallback batcher switchoff point")
		case <-ticker.C:
			status, err := system.RollupClient(e2esys.RoleVerif).SyncStatus(ctx)
			require.NoError(t, err)
			if status.CurrentL1.Number > switchReceipt.BlockNumber.Uint64() {
				l2Height = status.LocalSafeL2.Number
				break Loop
			}
		}
	}

	espHeight, err := espClient.FetchLatestBlockHeight(ctx)
	require.NoError(t, err)

	// Start a new "TEE" batcher
	batcherConfig.Espresso.CaffeinationHeightEspresso = espHeight
	batcherConfig.Espresso.CaffeinationHeightL2 = l2Height
	newBatcher, err := batcher.BatcherServiceFromCLIConfig(ctx, "0.0.1", batcherConfig, system.BatchSubmitter.Log)
	require.NoError(t, err)
	err = newBatcher.Start(ctx)
	require.NoError(t, err)

	// Everything should still work
	env.RunSimpleL2Burn(ctx, t, system)

	caffNode, err := env.LaunchCaffNode(t, system, espressoDevNode, func(c *config.Config) {
		c.Rollup.CaffNodeConfig.CaffeinationHeightEspresso = espHeight
		c.Rollup.CaffNodeConfig.CaffeinationHeightL2 = l2Height
	})
	require.NoError(t, err)
	defer env.Stop(t, caffNode)

	caffClient := system.NodeClient(env.RoleCaffNode)

	verifHeight, err := verifClient.BlockNumber(ctx)
	require.NoError(t, err)
	verifBlock, err := verifClient.BlockByNumber(ctx, new(big.Int).SetUint64(verifHeight))
	require.NoError(t, err)

	err = wait.ForBlock(ctx, caffClient, verifHeight)
	require.NoError(t, err)

	caffBlock, err := caffClient.BlockByNumber(ctx, new(big.Int).SetUint64(verifHeight))
	require.NoError(t, err)

	require.Equal(t, verifBlock.Hash(), caffBlock.Hash())
}

// TxManagerIntercept is a txmgr.TxManager that wraps another txmgr.TxManager
// and intercepts calls to Send and SendAsync.
//
// The purpose of this intercept is to simulate a failure in the tx submission
// process such that when activated a single frame of a multi-frame channel is
// sent to L1, and the remaining frames fail to be sent, triggering the fallback
// batcher to take over.
type TxManagerIntercept struct {
	txmgr.TxManager

	// shouldFail indicates whether to simulate a failure on Send/SendAsync.
	shouldFail bool

	// triggerAfterOne indicates whether to start failing after a single
	// successful Send/SendAsync.
	triggerAfterOne bool

	// failureCount tracks the number of failures that have occurred.
	failureCount int
}

// ErrSimulatedTxSubmissionFailure is the sentinel error returned when a
// simulated tx submission failure is triggered.
//
// We utilize this as a placeholder error to indicate that the tx submission
// failure was intentional for testing purposes.
var ErrSimulatedTxSubmissionFailure = errors.New("simulated tx submission failure")

// Send implements txmgr.TxManager.
//
// Send is overridden to simulate a failure when shouldFail is true, and to
// allow for one final transaction to be sent before failures begin when
// triggerAfterOne is true.
func (t *TxManagerIntercept) Send(ctx context.Context, candidate txmgr.TxCandidate) (*types.Receipt, error) {
	if t.shouldFail {
		t.failureCount++
		time.Sleep(50 * time.Millisecond) // Simulate some delay
		return nil, ErrSimulatedTxSubmissionFailure
	}

	if t.triggerAfterOne {
		t.shouldFail = true
	}

	return t.TxManager.Send(ctx, candidate)
}

// SendAsync implements txmgr.TxManager.
//
// SendAsync is overridden to simulate a failure when shouldFail is true, and
// to allow for one final transaction to be sent before failures begin when
// triggerAfterOne is true.
func (t *TxManagerIntercept) SendAsync(ctx context.Context, candidate txmgr.TxCandidate, ch chan txmgr.SendResponse) {
	if t.shouldFail {
		t.failureCount++
		time.Sleep(50 * time.Millisecond) // Simulate some delay
		ch <- txmgr.SendResponse{Err: ErrSimulatedTxSubmissionFailure}
		return
	}

	if t.triggerAfterOne {
		t.shouldFail = true
	}

	t.TxManager.SendAsync(ctx, candidate, ch)
}

// Compile time assertion to ensure TxManagerIntercept implements
// txmgr.TxManager.
var _ txmgr.TxManager = (*TxManagerIntercept)(nil)

// TestFallbackMechanismIntegrationTestChannelNotClosed is a test case that is
// meant to verify the correct expected behavior in the event that the Espresso
// Batcher encounters an error mid L1 Batch submission that prevents the full
// channel from being submitted to the L1.
//
// In this scenario this issue is expected to send a single frame of a
// multi-frame channel to the contract. At this point the batch should be
// switched to the fallback and the fallback batcher should continue
// submitting the remaining frames of the channel without any issues.
func TestFallbackMechanismIntegrationTestChannelNotClosed(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartE2eDevnet(
		ctx,
		t,
		// We want a Max L1 Number of frames larger than 1 to ensure we can
		// trigger the multi-frame channel scenario.
		env.WithBatcherTargetNumFrames(3),

		// We want a small Max L1 Tx Size to ensure that even a small L2
		// transaction will result in multiple L1 Transactions.
		env.WithBatcherMaxL1TxSize(5000),
	)

	require.NoError(t, err)

	// We create an intercept around the existing tx manager so we have
	// control over when our failures start to occur.

	interceptTxManager := &TxManagerIntercept{
		TxManager: system.BatchSubmitter.TxManager,
	}
	system.BatchSubmitter.TxManager = interceptTxManager
	system.BatchSubmitter.TestDriver().DriverSetup.Txmgr = interceptTxManager

	l1Client := system.NodeClient(e2esys.RoleL1)

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	// Send Transaction on L1, and wait for verification on the L2 Verifier
	env.RunSimpleL1TransferAndVerifier(ctx, t, system)

	// Verify everything works
	env.RunSimpleL2Burn(ctx, t, system)

	// We want to trigger the failure mode now.
	interceptTxManager.triggerAfterOne = true

	// Now we need to submit a multi-frame channel to L1 to trigger the
	// failure. We can do this by adjusting the batcher config to use a very
	// small MaxL1TxSize such that even a small L2 transaction will result in
	// multiple frames.

	// We want enough L2 Transactions to ensure we have multiple frames.
	n := 10

	receipts := env.RunSimpleMultiTransactions(ctx, t, system, n)

	// We want to wait until we know that the intercept tx manager has
	// trigger the failure mode successfully, and that all n transactions
	// have been attempted.

	err = wait.For(ctx, 10*time.Second, func() (bool, error) {
		return interceptTxManager.failureCount >= 1, nil
	})
	require.NoError(t, err)

	if have, want := interceptTxManager.failureCount, 1; have < want {
		t.Fatalf("tx submission failure not triggered enough times:\nhave:\n\t%d\nwant at least:\n\t%d", have, want)
	}

	// Make sure that the verifier doesn't see any of the transactions.

	l2Verif := system.NodeClient(e2esys.RoleVerif)

	for _, receipt := range receipts {
		_, err := l2Verif.TransactionReceipt(ctx, receipt.TxHash)
		if have, doNotWant := err, error(nil); have == doNotWant {
			t.Errorf("receipt for tx %s found on L2 Verifier when not expected:\nhave:\n\t%v\nwant:\n\t%v", receipt.TxHash, have, doNotWant)
		}
	}

	// Stop the "TEE" batcher
	err = system.BatchSubmitter.TestDriver().StopBatchSubmitting(ctx)
	require.NoError(t, err)

	// Switch active batcher
	options, err := bind.NewKeyedTransactorWithChainID(system.Config().Secrets.Deployer, system.Cfg.L1ChainIDBig())
	require.NoError(t, err)

	batchAuthenticator, err := bindings.NewBatchAuthenticator(system.RollupConfig.BatchAuthenticatorAddress, l1Client)
	require.NoError(t, err)

	tx, err := batchAuthenticator.SwitchBatcher(options)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, l1Client, tx.Hash())
	require.NoError(t, err)

	// Start the fallback batcher
	err = system.FallbackBatchSubmitter.TestDriver().StartBatchSubmitting()
	require.NoError(t, err)

	// Everything should still work
	env.RunSimpleL2Burn(ctx, t, system)

	// Verify that our previous receipts were also recorded on L1
	for i, receipt := range receipts {
		_, err := wait.ForReceiptOK(ctx, l2Verif, receipt.TxHash)
		if have, want := err, error(nil); have != want {
			t.Errorf("receipt %d for tx %s not found on L2 Verifier:\nhave:\n\t%v\nwant:\n\t%v", i, receipt.TxHash, have, want)
		}
	}
}
