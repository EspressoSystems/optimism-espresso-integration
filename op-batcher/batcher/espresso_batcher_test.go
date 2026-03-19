package batcher

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// nonEmptyHeadL1 returns a minimal L1BlockRef that is not the zero value, so that
// getSyncStatus does not enter its empty-sync-status backoff loop.
var nonEmptyHeadL1 = eth.L1BlockRef{Number: 1}

func syncStatusAt(safeL2, unsafeL2 uint64) *eth.SyncStatus {
	return &eth.SyncStatus{
		HeadL1:   nonEmptyHeadL1,
		SafeL2:   eth.L2BlockRef{Number: safeL2},
		UnsafeL2: eth.L2BlockRef{Number: unsafeL2},
	}
}

// setupCaffTest returns a BatchSubmitter pre-wired for caffeination-height tests:
// killCtx is live, PollInterval is short, NetworkTimeout is short.
func setupCaffTest(t *testing.T) (*BatchSubmitter, *mockL2EndpointProvider) {
	t.Helper()
	bs, ep := setup(t)
	bs.Config.PollInterval = 10 * time.Millisecond
	bs.Config.NetworkTimeout = 100 * time.Millisecond
	bs.killCtx, bs.cancelKillCtx = context.WithCancel(context.Background())
	t.Cleanup(bs.cancelKillCtx)
	return bs, ep
}

// ── waitForCaffeinationHeight ─────────────────────────────────────────────────

// TestWaitForCaffeinationHeight_ZeroHeight verifies that CaffeinationHeightL2 == 0
// is treated as "disabled" and the function returns immediately without polling.
func TestWaitForCaffeinationHeight_ZeroHeight(t *testing.T) {
	bs, _ := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 0

	err := bs.waitForCaffeinationHeight()
	require.NoError(t, err)
	// If the mock received any SyncStatus calls, testify's mock would
	// complain about unexpected calls at test teardown – no assertion needed here.
}

// TestWaitForCaffeinationHeight_AlreadyAtHeight verifies that the function returns
// nil on the very first poll when UnsafeL2 is already at caffHeight.
func TestWaitForCaffeinationHeight_AlreadyAtHeight(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100

	ep.rollupClient.ExpectSyncStatus(syncStatusAt(80, 100), nil)

	err := bs.waitForCaffeinationHeight()
	require.NoError(t, err)
}

// TestWaitForCaffeinationHeight_AboveHeight verifies that the function returns nil
// when UnsafeL2 is strictly above caffHeight on the first poll.
func TestWaitForCaffeinationHeight_AboveHeight(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100

	ep.rollupClient.ExpectSyncStatus(syncStatusAt(80, 150), nil)

	err := bs.waitForCaffeinationHeight()
	require.NoError(t, err)
}

// TestWaitForCaffeinationHeight_WaitsForTick verifies that when the first poll
// returns a status below caffHeight, the function waits for the ticker and
// eventually returns nil once the height is reached.
func TestWaitForCaffeinationHeight_WaitsForTick(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100

	// First poll: below height → should wait for ticker
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(50, 50), nil)
	// Second poll: at height → should return
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(80, 100), nil)

	err := bs.waitForCaffeinationHeight()
	require.NoError(t, err)
}

// TestWaitForCaffeinationHeight_ContextCancelled verifies that cancelling killCtx
// while the function is waiting for the ticker causes it to return an error.
func TestWaitForCaffeinationHeight_ContextCancelled(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100
	// Use a very long PollInterval so the select is still waiting when we cancel.
	bs.Config.PollInterval = 10 * time.Second

	// Replace killCtx with one we control independently of the cleanup registered
	// by setupCaffTest.
	killCtx, cancelKill := context.WithCancel(context.Background())
	bs.killCtx = killCtx
	bs.cancelKillCtx = cancelKill

	// First poll returns below the caffeination height.
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(50, 50), nil)

	// Cancel before calling – the select inside waitForCaffeinationHeight will
	// immediately pick killCtx.Done() after the first failed poll.
	cancelKill()

	err := bs.waitForCaffeinationHeight()
	require.Error(t, err)
	require.ErrorContains(t, err, "batcher stopped while waiting for caffeination height")
}

// TestWaitForCaffeinationHeight_ErrorThenSuccess verifies that a transient RPC
// error on the first poll is logged and the function retries, eventually returning
// nil when the height is reached.
func TestWaitForCaffeinationHeight_ErrorThenSuccess(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100

	// First call: RPC error
	ep.rollupClient.ExpectSyncStatus(&eth.SyncStatus{}, errors.New("rpc unavailable"))
	// Second call: at height
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(80, 100), nil)

	err := bs.waitForCaffeinationHeight()
	require.NoError(t, err)
}

// ── BlockLoader.nextBlockRange ────────────────────────────────────────────────

// newBlockLoader builds the minimal BlockLoader needed for nextBlockRange tests.
func newBlockLoader(t *testing.T, caffHeight uint64) *BlockLoader {
	t.Helper()
	return &BlockLoader{
		batcher: &BatchSubmitter{
			DriverSetup: DriverSetup{
				Log: testlog.Logger(t, log.LevelDebug),
				Config: BatcherConfig{
					CaffeinationHeightL2: caffHeight,
				},
			},
		},
	}
}

// TestNextBlockRange_EmptySyncStatus verifies that a zero HeadL1 causes ActionRetry
// (the sync status is not yet meaningful).
func TestNextBlockRange_EmptySyncStatus(t *testing.T) {
	loader := newBlockLoader(t, 0)

	_, action := loader.nextBlockRange(&eth.SyncStatus{})
	require.Equal(t, EnqueueBlockAction(ActionRetry), action)
}

// TestNextBlockRange_CaffHeightDisabled verifies that when CaffeinationHeightL2 == 0
// the start of the first range is safeL2+1 (no clamping applied).
func TestNextBlockRange_CaffHeightDisabled(t *testing.T) {
	loader := newBlockLoader(t, 0)

	r, action := loader.nextBlockRange(syncStatusAt(50, 150))
	require.Equal(t, EnqueueBlockAction(ActionEnqueue), action)
	require.Equal(t, uint64(51), r.start)
	require.Equal(t, uint64(150), r.end)
}

// TestNextBlockRange_CaffHeightClampsStart verifies that when CaffeinationHeightL2
// is above safeL2+1, the first range starts at caffHeight, not safeL2+1.
func TestNextBlockRange_CaffHeightClampsStart(t *testing.T) {
	const caffHeight = uint64(100)
	loader := newBlockLoader(t, caffHeight)

	// safeL2=50 → safeL2+1=51, caffHeight=100: max(51,100)=100
	r, action := loader.nextBlockRange(syncStatusAt(50, 150))
	require.Equal(t, EnqueueBlockAction(ActionEnqueue), action)
	require.Equal(t, caffHeight, r.start)
	require.Equal(t, uint64(150), r.end)
}

// TestNextBlockRange_CaffHeightAboveUnsafe verifies that when CaffeinationHeightL2
// is above UnsafeL2 (chain hasn't caught up yet) the function returns ActionRetry.
func TestNextBlockRange_CaffHeightAboveUnsafe(t *testing.T) {
	loader := newBlockLoader(t, 100)

	// unsafeL2=80 < caffHeight=100 → start(100) > end(80) → ActionRetry
	_, action := loader.nextBlockRange(syncStatusAt(50, 80))
	require.Equal(t, EnqueueBlockAction(ActionRetry), action)
}

// TestNextBlockRange_CaffHeightBelowSafeL2 verifies that when safeL2+1 already
// exceeds caffHeight, the clamp has no effect and start = safeL2+1.
func TestNextBlockRange_CaffHeightBelowSafeL2(t *testing.T) {
	loader := newBlockLoader(t, 100)

	// safeL2=120 → safeL2+1=121, caffHeight=100: max(121,100)=121
	r, action := loader.nextBlockRange(syncStatusAt(120, 150))
	require.Equal(t, EnqueueBlockAction(ActionEnqueue), action)
	require.Equal(t, uint64(121), r.start)
	require.Equal(t, uint64(150), r.end)
}

// TestNextBlockRange_CaffHeightExactlyAtSafeL2Plus1 verifies the boundary case
// where caffHeight == safeL2+1: max() is a no-op and start == caffHeight == safeL2+1.
func TestNextBlockRange_CaffHeightExactlyAtSafeL2Plus1(t *testing.T) {
	loader := newBlockLoader(t, 100)

	// safeL2=99 → safeL2+1=100, caffHeight=100: max(100,100)=100
	r, action := loader.nextBlockRange(syncStatusAt(99, 150))
	require.Equal(t, EnqueueBlockAction(ActionEnqueue), action)
	require.Equal(t, uint64(100), r.start)
	require.Equal(t, uint64(150), r.end)
}

// TestNextBlockRange_QueueNotEmpty verifies that once blocks are queued the
// caffeination height has no further effect – the next range starts at
// lastQueuedBlock+1 regardless of caffHeight.
func TestNextBlockRange_QueueNotEmpty(t *testing.T) {
	loader := newBlockLoader(t, 100)

	// Seed the loader as if blocks 100-101 were already queued.
	// safeL2=99 keeps the queue consistent: nextSafeBlockNum=100=firstQueuedBlock,
	// so numBlocksToEnqueue=0 and no hash check is required.
	loader.queuedBlocks = []eth.L2BlockRef{
		{Number: 100},
		{Number: 101},
	}

	// safeL2=99, unsafeL2=150: lastQueuedBlock=101, next range is 102-150.
	// caffHeight=100 is ignored because the queue is not empty.
	r, action := loader.nextBlockRange(syncStatusAt(99, 150))
	require.Equal(t, EnqueueBlockAction(ActionEnqueue), action)
	require.Equal(t, uint64(102), r.start)
	require.Equal(t, uint64(150), r.end)
}
