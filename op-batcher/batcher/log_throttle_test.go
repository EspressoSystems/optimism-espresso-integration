package batcher

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// fakeClock returns the time stored at *t, allowing the test to advance time
// deterministically without sleeping.
func fakeClock(now *time.Time) func() time.Time {
	return func() time.Time { return *now }
}

func TestRepeatStateLogger_FirstWarnEmitsThenSuppresses(t *testing.T) {
	lgr, capt := testlog.CaptureLogger(t, log.LevelTrace)
	now := time.Unix(0, 0)
	r := newRepeatStateLogger()
	r.clock = fakeClock(&now)

	r.Warn(lgr, "k1", "degraded", "err", "boom")
	now = now.Add(1 * time.Second)
	r.Warn(lgr, "k1", "degraded", "err", "boom")
	now = now.Add(1 * time.Second)
	r.Warn(lgr, "k1", "degraded", "err", "boom")

	warns := capt.FindLogs(testlog.NewLevelFilter(log.LevelWarn), testlog.NewMessageFilter("degraded"))
	require.Len(t, warns, 1, "only the first observation should emit a log")

	require.Equal(t, "boom", warns[0].AttrValue("err"))
	require.Nil(t, warns[0].AttrValue("suppressed"), "first emission should not carry a suppressed count")
}

func TestRepeatStateLogger_ReminderAfterInterval(t *testing.T) {
	lgr, capt := testlog.CaptureLogger(t, log.LevelTrace)
	now := time.Unix(0, 0)
	r := newRepeatStateLogger()
	r.clock = fakeClock(&now)

	// Initial emission.
	r.Warn(lgr, "k1", "degraded")

	// 9 silent observations, every 30s. Cumulative 4m30s, still under the 5m threshold.
	for i := 0; i < 9; i++ {
		now = now.Add(30 * time.Second)
		r.Warn(lgr, "k1", "degraded")
	}
	warns := capt.FindLogs(testlog.NewLevelFilter(log.LevelWarn), testlog.NewMessageFilter("degraded"))
	require.Len(t, warns, 1, "no reminder before the interval has elapsed")

	// Cross the threshold.
	now = now.Add(31 * time.Second)
	r.Warn(lgr, "k1", "degraded")

	warns = capt.FindLogs(testlog.NewLevelFilter(log.LevelWarn), testlog.NewMessageFilter("degraded"))
	require.Len(t, warns, 2, "reminder should fire once the interval has elapsed")

	// Reminder includes 10 suppressed observations (9 silent + the one that triggered the reminder).
	require.EqualValues(t, 10, warns[1].AttrValue("suppressed"))
	// Duration since firstSeen is 9*30s + 31s = 5m1s, rounded to nearest second.
	require.Equal(t, 5*time.Minute+1*time.Second, warns[1].AttrValue("duration"))
}

func TestRepeatStateLogger_KeysAreIndependent(t *testing.T) {
	lgr, capt := testlog.CaptureLogger(t, log.LevelTrace)
	now := time.Unix(0, 0)
	r := newRepeatStateLogger()
	r.clock = fakeClock(&now)

	r.Warn(lgr, "k1", "first state")
	r.Warn(lgr, "k2", "second state")
	// Both keys are now active. Repeats for either should be suppressed.
	r.Warn(lgr, "k1", "first state")
	r.Warn(lgr, "k2", "second state")

	require.Len(t, capt.FindLogs(testlog.NewMessageFilter("first state")), 1)
	require.Len(t, capt.FindLogs(testlog.NewMessageFilter("second state")), 1)

	// Clearing one key must not affect the other.
	r.Clear(lgr, "k1", "first recovered")
	r.Warn(lgr, "k2", "second state") // still suppressed
	require.Len(t, capt.FindLogs(testlog.NewMessageFilter("second state")), 1)
}

func TestRepeatStateLogger_ClearEmitsRecoveryWhenActive(t *testing.T) {
	lgr, capt := testlog.CaptureLogger(t, log.LevelTrace)
	now := time.Unix(0, 0)
	r := newRepeatStateLogger()
	r.clock = fakeClock(&now)

	r.Warn(lgr, "k1", "degraded")
	now = now.Add(2 * time.Second)
	r.Warn(lgr, "k1", "degraded") // suppressed, but increments totalOccurrences
	now = now.Add(3 * time.Second)
	r.Clear(lgr, "k1", "recovered", "extra", "ctx")

	infos := capt.FindLogs(testlog.NewLevelFilter(log.LevelInfo), testlog.NewMessageFilter("recovered"))
	require.Len(t, infos, 1)
	require.Equal(t, 5*time.Second, infos[0].AttrValue("duration"))
	require.EqualValues(t, 2, infos[0].AttrValue("occurrences"))
	require.Equal(t, "ctx", infos[0].AttrValue("extra"))
}

func TestRepeatStateLogger_ClearWhenInactiveIsNoop(t *testing.T) {
	lgr, capt := testlog.CaptureLogger(t, log.LevelTrace)
	now := time.Unix(0, 0)
	r := newRepeatStateLogger()
	r.clock = fakeClock(&now)

	r.Clear(lgr, "k1", "recovered")

	require.Empty(t, capt.FindLogs(testlog.NewMessageFilter("recovered")))
}

func TestRepeatStateLogger_FreshAfterClear(t *testing.T) {
	lgr, capt := testlog.CaptureLogger(t, log.LevelTrace)
	now := time.Unix(0, 0)
	r := newRepeatStateLogger()
	r.clock = fakeClock(&now)

	r.Warn(lgr, "k1", "degraded")
	r.Warn(lgr, "k1", "degraded") // suppressed
	r.Clear(lgr, "k1", "recovered")

	// After Clear, the next Warn should emit again as a fresh first observation.
	r.Warn(lgr, "k1", "degraded")

	warns := capt.FindLogs(testlog.NewLevelFilter(log.LevelWarn), testlog.NewMessageFilter("degraded"))
	require.Len(t, warns, 2, "first Warn after Clear should emit")
	require.Nil(t, warns[1].AttrValue("suppressed"), "fresh emission should not carry a suppressed count")
}

func TestRepeatStateLogger_ConcurrentCallersDoNotRace(t *testing.T) {
	lgr := testlog.Logger(t, log.LevelTrace)
	r := newRepeatStateLogger()

	const goroutines = 32
	const callsPerG = 200

	var wg sync.WaitGroup
	wg.Add(goroutines)
	var clears atomic.Int64
	for g := 0; g < goroutines; g++ {
		go func(id int) {
			defer wg.Done()
			key := "k" + string(rune('a'+(id%4)))
			for i := 0; i < callsPerG; i++ {
				r.Warn(lgr, key, "degraded", "g", id)
				if i%50 == 0 {
					r.Clear(lgr, key, "recovered")
					clears.Add(1)
				}
			}
		}(g)
	}
	wg.Wait()
	// No assertion on counts — this test exists to flush out races under -race
	// and to assert the logger does not panic under concurrent Warn/Clear.
	require.Greater(t, clears.Load(), int64(0))
}
