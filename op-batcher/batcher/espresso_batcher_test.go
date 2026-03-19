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

// nonEmptyHeadL1 is a minimal L1BlockRef that prevents getSyncStatus from
// entering its empty-sync-status backoff loop.
var nonEmptyHeadL1 = eth.L1BlockRef{Number: 1}

func syncStatusAt(safeL2, unsafeL2 uint64) *eth.SyncStatus {
	return &eth.SyncStatus{
		HeadL1:   nonEmptyHeadL1,
		SafeL2:   eth.L2BlockRef{Number: safeL2},
		UnsafeL2: eth.L2BlockRef{Number: unsafeL2},
	}
}

func setupCaffTest(t *testing.T) (*BatchSubmitter, *mockL2EndpointProvider) {
	t.Helper()
	bs, ep := setup(t)
	bs.Config.PollInterval = 10 * time.Millisecond
	bs.Config.NetworkTimeout = 100 * time.Millisecond
	bs.killCtx, bs.cancelKillCtx = context.WithCancel(context.Background())
	t.Cleanup(bs.cancelKillCtx)
	return bs, ep
}

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

// ── waitForCaffeinationHeight ─────────────────────────────────────────────────

func TestWaitForCaffeinationHeight_ZeroHeight(t *testing.T) {
	bs, _ := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 0
	require.NoError(t, bs.waitForCaffeinationHeight())
}

// TestWaitForCaffeinationHeight_HeightReached covers both "exactly at" and "above" caffHeight.
func TestWaitForCaffeinationHeight_HeightReached(t *testing.T) {
	for _, unsafeL2 := range []uint64{100, 150} {
		bs, ep := setupCaffTest(t)
		bs.Config.CaffeinationHeightL2 = 100
		ep.rollupClient.ExpectSyncStatus(syncStatusAt(80, unsafeL2), nil)
		require.NoError(t, bs.waitForCaffeinationHeight())
	}
}

func TestWaitForCaffeinationHeight_WaitsForTick(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(50, 50), nil)  // below → wait
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(80, 100), nil) // at height → return
	require.NoError(t, bs.waitForCaffeinationHeight())
}

func TestWaitForCaffeinationHeight_ContextCancelled(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100
	bs.Config.PollInterval = 10 * time.Second // long enough that ctx fires first
	killCtx, cancelKill := context.WithCancel(context.Background())
	bs.killCtx, bs.cancelKillCtx = killCtx, cancelKill
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(50, 50), nil)
	cancelKill()
	err := bs.waitForCaffeinationHeight()
	require.ErrorContains(t, err, "batcher stopped while waiting for caffeination height")
}

func TestWaitForCaffeinationHeight_ErrorThenSuccess(t *testing.T) {
	bs, ep := setupCaffTest(t)
	bs.Config.CaffeinationHeightL2 = 100
	ep.rollupClient.ExpectSyncStatus(&eth.SyncStatus{}, errors.New("rpc unavailable"))
	ep.rollupClient.ExpectSyncStatus(syncStatusAt(80, 100), nil)
	require.NoError(t, bs.waitForCaffeinationHeight())
}

// ── BlockLoader.nextBlockRange ────────────────────────────────────────────────

func TestNextBlockRange(t *testing.T) {
	tests := []struct {
		name         string
		caffHeight   uint64
		safeL2       uint64
		unsafeL2     uint64
		queuedBlocks []eth.L2BlockRef
		emptySS      bool
		wantAction   EnqueueBlockAction
		wantStart    uint64
		wantEnd      uint64
	}{
		{
			name:       "empty sync status",
			emptySS:    true,
			wantAction: ActionRetry,
		},
		{
			name:       "caffHeight disabled",
			safeL2: 50, unsafeL2: 150,
			wantAction: ActionEnqueue, wantStart: 51, wantEnd: 150,
		},
		{
			name:       "caffHeight clamps start",
			caffHeight: 100, safeL2: 50, unsafeL2: 150,
			wantAction: ActionEnqueue, wantStart: 100, wantEnd: 150,
		},
		{
			name:       "caffHeight above unsafe",
			caffHeight: 100, safeL2: 50, unsafeL2: 80,
			wantAction: ActionRetry,
		},
		{
			name:       "caffHeight below safeL2",
			caffHeight: 100, safeL2: 120, unsafeL2: 150,
			wantAction: ActionEnqueue, wantStart: 121, wantEnd: 150,
		},
		{
			name:       "caffHeight exactly at safeL2+1",
			caffHeight: 100, safeL2: 99, unsafeL2: 150,
			wantAction: ActionEnqueue, wantStart: 100, wantEnd: 150,
		},
		{
			name:         "queue not empty ignores caffHeight",
			caffHeight:   100, safeL2: 99, unsafeL2: 150,
			queuedBlocks: []eth.L2BlockRef{{Number: 100}, {Number: 101}},
			wantAction:   ActionEnqueue, wantStart: 102, wantEnd: 150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loader := newBlockLoader(t, tt.caffHeight)
			loader.queuedBlocks = tt.queuedBlocks
			ss := syncStatusAt(tt.safeL2, tt.unsafeL2)
			if tt.emptySS {
				ss = &eth.SyncStatus{}
			}
			r, action := loader.nextBlockRange(ss)
			require.Equal(t, tt.wantAction, action)
			if tt.wantAction == ActionEnqueue {
				require.Equal(t, tt.wantStart, r.start)
				require.Equal(t, tt.wantEnd, r.end)
			}
		})
	}
}
