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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

// TestEspressoEnforcementHardfork exercises the EspressoEnforcementTime hardfork
// transition on a single continuous chain, including the boundary itself.
//
// Pre-fork the chain runs in pure upstream Optimism mode: the fallback batcher
// posts plain BatchInbox txs signed by the SystemConfig batcher key, with no
// BatchAuthenticator interaction at all. The verifier accepts these batches via
// sender-based authorization (the upstream code path), as gated by
// DataSourceConfig.isEspressoEnforcement on L1 origin time.
//
// The test deliberately runs the *fallback* batcher across the fork boundary
// rather than swapping it out for the TEE batcher beforehand. This exercises
// the asymmetry between the batcher's fallback-auth gate (decided on L1 tip
// time) and the verifier's gate (decided on the L1 origin time of the block
// containing the batch tx): the batcher must start authenticating before the
// verifier requires it, otherwise batches submitted around the boundary
// arrive in post-fork L1 blocks without their BatchInfoAuthenticated events
// and the verifier silently drops them. The lead-time mechanism in
// `isFallbackAuthRequired` (op-batcher/batcher/espresso_active.go) prevents
// this.
//
// After the verifier safely crosses the boundary, the test stops the fallback
// batcher and starts the Espresso TEE batcher with caffeination heights set
// to the live head (avoiding a streamer backfill from genesis). The same
// continuous verifier then accepts post-fork event-authenticated batches
// from the TEE batcher.
//
// This complements TestBatcherSwitching, which only exercises the on-chain
// activeIsEspresso flip within the Espresso world. This test exercises the
// consensus hardfork itself.
func TestEspressoEnforcementHardfork(t *testing.T) {
	// Pre-fork window: needs to cover the time it takes the Espresso devnet
	// to spin up (Docker container start, L1 contract deploys, etc.) plus
	// the pre-fork test operations (deposit + L2 burn round-trips). The L1
	// RPC tip advances at near wall-clock speed (L1BlockTime=2s on a
	// freshly-mined L1), so a too-short offset means the fork has already
	// activated by the time the batcher does its first publish attempt.
	const enforcementOffset = 5 * time.Minute

	// Tighten the fallback-auth lead time so the batcher actually spends a
	// meaningful chunk of wall-clock in pre-fork (no-auth) mode before the
	// boundary, while still being large enough to absorb a devnet L1's tx
	// inclusion delay (L1BlockTime=2s). 30s is comfortably both: it leaves
	// roughly 4.5 minutes of pre-fork operation, and is ~15 L1 blocks of
	// safety margin against the verifier's L1-origin-time gate.
	const fallbackAuthLeadTime = 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Minute)
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// Capture the original Espresso TEE batcher CLI config so we can restart
	// it with caffeination heights post-fork.
	espressoBatcherConfig := &batcher.CLIConfig{}

	system, espressoDevNode, err := launcher.StartE2eDevnet(ctx, t,
		env.WithEspressoEnforcementOffset(enforcementOffset),
		env.WithFallbackAuthLeadTime(fallbackAuthLeadTime),
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
	// fallback-auth gate (`isFallbackAuthRequired`) returns false because
	// even with the lead-time added the predicted-future origin time is
	// still pre-fork, so the batcher behaves identically to an upstream
	// Optimism batcher: plain BatchInbox tx signed by
	// SystemConfig.batcherHash, no BatchAuthenticator interaction.
	require.NoError(t, system.FallbackBatchSubmitter.TestDriver().StartBatchSubmitting())

	env.RunSimpleL1TransferAndVerifier(ctx, t, system)
	env.RunSimpleL2Burn(ctx, t, system)

	status, err := verifRollup.SyncStatus(ctx)
	require.NoError(t, err)
	require.Less(t, status.SafeL2.Time, forkTime,
		"verifier should still be pre-fork after initial L2 burn")

	// Phase 2 (boundary): keep the fallback batcher running across the fork
	// boundary. As `tip.Time + fallbackAuthLeadTime` crosses
	// EspressoEnforcementTime, the batcher's gate flips and subsequent
	// batches are submitted via authenticateBatchInfo. Batches that land in
	// post-fork L1 blocks must have a corresponding BatchInfoAuthenticated
	// event in the lookback window, otherwise the verifier silently drops
	// them and the safe head stalls.
	//
	// Capture the L1 block range we're interested in for event-presence
	// assertions below. We snapshot the current L1 head before crossing the
	// boundary; any BatchInfoAuthenticated events emitted from this block
	// onward are attributable to the fallback batcher's lead-time-driven
	// switch to authenticated submission.
	preBoundaryL1Head, err := l1Client.BlockNumber(ctx)
	require.NoError(t, err)
	require.NoError(t, wait.ForBlockWithTimestamp(ctx, l1Client, forkTime))
	t.Logf("L1 reached fork timestamp; pre-boundary L1 head was %d", preBoundaryL1Head)

	// Wait for the verifier to cross the fork timestamp without dropped
	// batches. A stall here (tx in post-fork L1 block lacking auth event)
	// would manifest as a context timeout. We assert the L2 safe head moves
	// past forkTime, which can only happen if the fallback batcher's
	// authenticated post-fork batches are accepted by the verifier.
	require.NoError(t, wait.ForBlockWithTimestamp(ctx, verifClient, forkTime+30))
	require.NoError(t, wait.ForNextBlock(ctx, l1Client))
	require.NoError(t, wait.ForNextBlock(ctx, l1Client))

	status, err = verifRollup.SyncStatus(ctx)
	require.NoError(t, err)
	require.Greater(t, status.SafeL2.Time, forkTime,
		"verifier should have advanced strictly past forkTime — a stall here likely "+
			"indicates the fallback batcher posted an unauthenticated batch in a "+
			"post-fork L1 block (lead-time too short)")

	postBoundaryL1Head, err := l1Client.BlockNumber(ctx)
	require.NoError(t, err)

	// Verify the fallback batcher actually emitted at least one
	// BatchInfoAuthenticated event around the boundary (i.e., the new
	// fallback-auth code path actually executed). Without this, the test
	// could pass even if the lead time were so large that all sampled
	// batches happened to be post-fork already at submission time.
	batchAuth, err := bindings.NewBatchAuthenticator(
		system.RollupConfig.BatchAuthenticatorAddress, l1Client)
	require.NoError(t, err)
	filterOpts := &bind.FilterOpts{
		Context: ctx,
		Start:   preBoundaryL1Head,
		End:     &postBoundaryL1Head,
	}
	authIter, err := batchAuth.FilterBatchInfoAuthenticated(filterOpts, nil)
	require.NoError(t, err)
	defer authIter.Close()
	authEventCount := 0
	for authIter.Next() {
		authEventCount++
	}
	require.NoError(t, authIter.Error())
	require.Greater(t, authEventCount, 0,
		"fallback batcher should have emitted at least one BatchInfoAuthenticated event "+
			"around the fork boundary")
	t.Logf("fallback batcher emitted %d BatchInfoAuthenticated events around the boundary",
		authEventCount)

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
	// The captured config inherits Stopped=true from WithBatcherStoppedInitially.
	// Clear it so the new TEE batcher actually begins submitting on Start.
	espressoBatcherConfig.Stopped = false

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
