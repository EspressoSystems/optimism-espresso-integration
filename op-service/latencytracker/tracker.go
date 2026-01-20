package latencytracker

import (
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// Tracker tracks latency samples and computes avg/p95 statistics.
// It is designed for temporary debugging of latency issues.
type Tracker struct {
	name    string
	mu      sync.Mutex
	samples []time.Duration
	maxSize int

	// For periodic logging
	lastLogTime     time.Time
	logInterval     time.Duration
	samplesSinceLog int
}

// New creates a new latency tracker with the given name.
// It will log stats every logInterval if there are new samples.
func New(name string, maxSamples int, logInterval time.Duration) *Tracker {
	return &Tracker{
		name:        name,
		samples:     make([]time.Duration, 0, maxSamples),
		maxSize:     maxSamples,
		logInterval: logInterval,
		lastLogTime: time.Now(),
	}
}

// Record records a latency sample and potentially logs statistics.
func (lt *Tracker) Record(d time.Duration) {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	// Add sample, evicting oldest if at capacity
	if len(lt.samples) >= lt.maxSize {
		lt.samples = lt.samples[1:]
	}
	lt.samples = append(lt.samples, d)
	lt.samplesSinceLog++

	// Log periodically
	if time.Since(lt.lastLogTime) >= lt.logInterval && lt.samplesSinceLog > 0 {
		lt.logStatsLocked()
	}
}

// RecordSince records the duration since the given start time.
func (lt *Tracker) RecordSince(start time.Time) {
	lt.Record(time.Since(start))
}

// ForceLog forces logging of current statistics.
func (lt *Tracker) ForceLog() {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	lt.logStatsLocked()
}

func (lt *Tracker) logStatsLocked() {
	if len(lt.samples) == 0 {
		return
	}

	// Compute statistics
	sorted := make([]time.Duration, len(lt.samples))
	copy(sorted, lt.samples)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	var sum time.Duration
	for _, s := range sorted {
		sum += s
	}
	avg := sum / time.Duration(len(sorted))

	p50Idx := len(sorted) * 50 / 100
	p95Idx := len(sorted) * 95 / 100
	p99Idx := len(sorted) * 99 / 100
	if p95Idx >= len(sorted) {
		p95Idx = len(sorted) - 1
	}
	if p99Idx >= len(sorted) {
		p99Idx = len(sorted) - 1
	}

	log.Warn("LATENCY_TRACKER",
		"metric", lt.name,
		"samples", len(sorted),
		"avg", avg,
		"p50", sorted[p50Idx],
		"p95", sorted[p95Idx],
		"p99", sorted[p99Idx],
		"min", sorted[0],
		"max", sorted[len(sorted)-1],
	)

	lt.lastLogTime = time.Now()
	lt.samplesSinceLog = 0
}

// Global latency trackers for espresso components
var (
	// Streamer latency trackers
	StreamerFetchLatestHeightLatency = New("streamer_fetch_latest_height", 1000, 30*time.Second)
	StreamerFetchBlockLatency        = New("streamer_fetch_block_range", 1000, 30*time.Second)
	StreamerCheckBatchLatency        = New("streamer_check_batch", 1000, 30*time.Second)
	StreamerL1HeaderLookupLatency    = New("streamer_l1_header_lookup", 1000, 30*time.Second)
	StreamerUpdateTotalLatency       = New("streamer_update_total", 1000, 30*time.Second)
	StreamerProcessTxLatency         = New("streamer_process_tx", 1000, 30*time.Second)
	StreamerConfirmHeightLatency     = New("streamer_confirm_espresso_height", 1000, 30*time.Second)
	StreamerRefreshLatency           = New("streamer_refresh", 1000, 30*time.Second)

	// Caff node derivation latency trackers
	CaffNextBatchTotalLatency   = New("caff_next_batch_total", 1000, 30*time.Second)
	CaffL1FinalizedFetchLatency = New("caff_l1_finalized_fetch", 1000, 30*time.Second)
	CaffStreamerRefreshLatency  = New("caff_streamer_refresh", 1000, 30*time.Second)
	CaffStreamerUpdateLatency   = New("caff_streamer_update", 1000, 30*time.Second)
	CaffStreamerNextLatency     = New("caff_streamer_next", 1000, 30*time.Second)
	CaffBatchValidationLatency  = New("caff_batch_validation", 1000, 30*time.Second)

	// Batcher submission latency trackers
	BatcherQueueToSubmitLatency        = New("batcher_queue_to_submit", 1000, 30*time.Second)
	BatcherSubmitTxLatency             = New("batcher_submit_tx", 1000, 30*time.Second)
	BatcherVerifyReceiptLatency        = New("batcher_verify_receipt", 1000, 30*time.Second)
	BatcherTotalConfirmLatency         = New("batcher_total_confirm", 1000, 30*time.Second)
	BatcherBlockToEspressoBatchLatency = New("batcher_block_to_espresso_batch", 1000, 30*time.Second)
)
