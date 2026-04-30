package environment_test

import (
	"context"
	"testing"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	"github.com/ethereum-optimism/optimism/espresso/bindings"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/stretchr/testify/require"
)

// TestEspressoEnforcementHardfork exercises the EspressoEnforcementTime hardfork
// transition on a single continuous chain.
//
// Pre-fork the chain runs in pure upstream Optimism mode: the fallback batcher
// posts plain BatchInbox txs signed by the SystemConfig batcher key, with no
// BatchAuthenticator interaction at all. The verifier accepts these batches via
// sender-based authorization (the upstream code path), as gated by
// DataSourceConfig.isEspressoEnforcement on L1 origin time.
//
// After the fork timestamp passes, the test tears down the fallback batcher and
// starts the Espresso TEE batcher with caffeination heights set to the live
// head (avoiding a streamer backfill from genesis). The same continuous
// verifier then accepts post-fork event-authenticated batches, exercising the
// per-L1-block fork gate.
//
// This complements TestBatcherSwitching, which only exercises the on-chain
// activeIsEspresso flip within the Espresso world. This test exercises the
// consensus hardfork itself.
func TestEspressoEnforcementHardfork(t *testing.T) {
	// 60s pre-fork window: long enough to fit a few L2 transactions and a
	// settled L1 block before the fork; short enough to keep test runtime
	// reasonable. With L1BlockTime=2s and L2BlockTime=1s the verifier sees
	// ~30 L2 blocks pre-fork.
	const enforcementOffset = 60 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Minute)
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// Capture the original Espresso TEE batcher CLI config so we can restart
	// it with caffeination heights post-fork.
	espressoBatcherConfig := &batcher.CLIConfig{}

	system, espressoDevNode, err := launcher.StartE2eDevnet(ctx, t,
		env.WithEspressoEnforcementOffset(enforcementOffset),
		env.WithL1FinalizedDistance(0),
		env.WithSequencerUseFinalized(true),
		env.WithBatcherStoppedInitially(),
		env.GetBatcherConfig(espressoBatcherConfig),
	)
	require.NoError(t, err)
	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	require.NotNil(t, system.RollupConfig.EspressoEnforcementTime,
		"test requires EspressoEnforcementTime to be set on the rollup config")
	forkTime := *system.RollupConfig.EspressoEnforcementTime
	nowSec := uint64(time.Now().Unix())
	require.Greater(t, forkTime, nowSec,
		"fork timestamp must be in the future at test start")
	t.Logf("EspressoEnforcementTime = %d (now = %d, in %ds)",
		forkTime, nowSec, forkTime-nowSec)

	l1Client := system.NodeClient(e2esys.RoleL1)
	verifClient := system.NodeClient(e2esys.RoleVerif)
	verifRollup := system.RollupClient(e2esys.RoleVerif)
	espClient := espressoClient.NewClient(espressoDevNode.EspressoUrls()[0])

	// Phase 1 (pre-fork): run the fallback batcher. Pre-fork the batcher's
	// hardfork gate skips authenticateBatchInfo entirely, so this batcher
	// behaves identically to an upstream Optimism batcher: plain BatchInbox
	// tx signed by SystemConfig.batcherHash.
	require.NoError(t, system.FallbackBatchSubmitter.TestDriver().StartBatchSubmitting())

	env.RunSimpleL1TransferAndVerifier(ctx, t, system)
	env.RunSimpleL2Burn(ctx, t, system)

	status, err := verifRollup.SyncStatus(ctx)
	require.NoError(t, err)
	require.Less(t, status.SafeL2.Time, forkTime,
		"verifier should still be pre-fork after initial L2 burn")

	// Phase 2: wait for the verifier to cross the fork timestamp. The fork
	// gate in the data source uses L1 origin time, so wait for the L2
	// timestamp to cross then advance L1 a couple more blocks to ensure the
	// verifier has consumed at least one fully post-fork L1 origin.
	require.NoError(t, wait.ForBlockWithTimestamp(ctx, verifClient, forkTime))
	require.NoError(t, wait.ForNextBlock(ctx, l1Client))
	require.NoError(t, wait.ForNextBlock(ctx, l1Client))

	status, err = verifRollup.SyncStatus(ctx)
	require.NoError(t, err)
	require.GreaterOrEqual(t, status.SafeL2.Time, forkTime,
		"verifier should have crossed the fork timestamp")

	// Phase 3: stop the fallback batcher, capture live heads, and start the
	// Espresso TEE batcher.
	require.NoError(t, system.FallbackBatchSubmitter.TestDriver().StopBatchSubmitting(ctx))

	// Capture the current L2/Espresso heads so the new TEE batcher streams at
	// the live head rather than reprocessing the entire chain history.
	l2Height, err := waitForRollupToMovePastL1Block(ctx, verifRollup, status.CurrentL1.Number)
	require.NoError(t, err)
	espHeight, err := espClient.FetchLatestBlockHeight(ctx)
	require.NoError(t, err)

	// activeIsEspresso defaults to true after BatchAuthenticator initialization,
	// matching the TEE batcher we are about to start. No SwitchBatcher call
	// is needed here.

	// Sanity-check that the BatchAuthenticator is in the expected state.
	batchAuthenticator, err := bindings.NewBatchAuthenticator(
		system.RollupConfig.BatchAuthenticatorAddress, l1Client)
	require.NoError(t, err)
	activeIsEspresso, err := batchAuthenticator.ActiveIsEspresso(nil)
	require.NoError(t, err)
	require.True(t, activeIsEspresso,
		"activeIsEspresso should be true at fork transition (default contract state)")

	espressoBatcherConfig.MaxChannelDuration = 10
	espressoBatcherConfig.TargetNumFrames = 1
	espressoBatcherConfig.MaxL1TxSize = 120_000
	espressoBatcherConfig.Espresso.CaffeinationHeightEspresso = espHeight
	espressoBatcherConfig.Espresso.CaffeinationHeightL2 = l2Height

	teeCtx, teeCancel := context.WithCancelCause(ctx)
	defer teeCancel(nil)
	teeBatcher, err := batcher.BatcherServiceFromCLIConfig(
		teeCtx, teeCancel, "0.0.1", espressoBatcherConfig, system.BatchSubmitter.Log)
	require.NoError(t, err)
	require.NoError(t, teeBatcher.Start(teeCtx))

	// Phase 4 (post-fork): the same verifier accepts event-authenticated
	// batches without restart. Use a longer timeout while the new batcher
	// drains its first post-fork channel.
	env.RunSimpleL2BurnWithTimeout(ctx, t, system, 5*time.Minute)

	status, err = verifRollup.SyncStatus(ctx)
	require.NoError(t, err)
	require.Greater(t, status.SafeL2.Time, forkTime,
		"verifier should have advanced post-fork")
}
