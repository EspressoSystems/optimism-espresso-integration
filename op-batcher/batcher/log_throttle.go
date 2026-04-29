package batcher

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// repeatStateReminderInterval is how often a long-running degraded state
// re-emits a Warn so operators don't lose visibility while the state persists.
const repeatStateReminderInterval = 5 * time.Minute

// repeatStateLogger collapses warnings tied to a "degraded state" into a single
// log on entry, periodic reminders while the state persists, and a recovery
// log on exit. This avoids flooding the log debouncer when a tick-driven loop
// fires the same warning every poll interval.
//
// State is keyed by a free-form string supplied by the caller; entries with
// different keys are independent. Safe for concurrent use.
type repeatStateLogger struct {
	mu     sync.Mutex
	states map[string]*repeatStateEntry
}

type repeatStateEntry struct {
	firstSeen        time.Time
	lastLogged       time.Time
	suppressed       int
	totalOccurrences int
}

func newRepeatStateLogger() *repeatStateLogger {
	return &repeatStateLogger{
		states: make(map[string]*repeatStateEntry),
	}
}

// Warn reports an observation of a degraded state. The first observation since
// the most recent Clear (or first ever for the key) emits at warn level.
// Subsequent observations within repeatStateReminderInterval are silently
// counted; once the interval has elapsed a single reminder warn is emitted
// with the suppressed count and total duration.
func (r *repeatStateLogger) Warn(l log.Logger, key, msg string, ctx ...any) {
	now := time.Now()
	r.mu.Lock()
	e, active := r.states[key]
	if !active {
		r.states[key] = &repeatStateEntry{
			firstSeen:        now,
			lastLogged:       now,
			totalOccurrences: 1,
		}
		r.mu.Unlock()
		l.Warn(msg, ctx...)
		return
	}
	e.suppressed++
	e.totalOccurrences++
	if now.Sub(e.lastLogged) < repeatStateReminderInterval {
		r.mu.Unlock()
		return
	}
	suppressed := e.suppressed
	duration := now.Sub(e.firstSeen).Round(time.Second)
	e.suppressed = 0
	e.lastLogged = now
	r.mu.Unlock()

	args := append([]any{"suppressed", suppressed, "duration", duration}, ctx...)
	l.Warn(msg, args...)
}

// Clear marks the named state as resolved. If the state was active a single
// info-level recovery log is emitted summarising the duration and total
// occurrences. Calling Clear when the state is not active is a no-op, so it is
// safe to call on every successful tick of the loop.
func (r *repeatStateLogger) Clear(l log.Logger, key, recoveryMsg string, ctx ...any) {
	r.mu.Lock()
	e, active := r.states[key]
	if active {
		delete(r.states, key)
	}
	r.mu.Unlock()
	if !active {
		return
	}
	args := append([]any{
		"duration", time.Since(e.firstSeen).Round(time.Second),
		"occurrences", e.totalOccurrences,
	}, ctx...)
	l.Info(recoveryMsg, args...)
}
