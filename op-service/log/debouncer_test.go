package log

import (
	"context"
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/log"
)

// safeTestRecorder is a thread-safe version of testRecorder for concurrent tests
type safeTestRecorder struct {
	mu      sync.Mutex
	records []slog.Record
}

func (r *safeTestRecorder) Enabled(context.Context, slog.Level) bool {
	return true
}

func (r *safeTestRecorder) Handle(_ context.Context, rec slog.Record) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.records = append(r.records, rec)
	return nil
}

func (r *safeTestRecorder) WithAttrs([]slog.Attr) slog.Handler { return r }
func (r *safeTestRecorder) WithGroup(string) slog.Handler      { return r }

func (r *safeTestRecorder) GetRecords() []slog.Record {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Return a copy to avoid race conditions
	result := make([]slog.Record, len(r.records))
	copy(result, r.records)
	return result
}

func (r *safeTestRecorder) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.records)
}

func TestDebouncingHandler_Basic(t *testing.T) {
	h := new(testRecorder)
	d := NewDebouncingHandler(h)
	logger := log.NewLogger(d)

	// First message should go through
	logger.Info("hello world")
	require.Len(t, h.records, 1)
	require.Equal(t, "hello world", h.records[0].Message)

	// Same message within 100ms should be dropped
	logger.Info("hello world")
	require.Len(t, h.records, 1)

	// Different message should go through
	logger.Info("different message")
	require.Len(t, h.records, 2)
	require.Equal(t, "different message", h.records[1].Message)

	// Wait for debounce period to expire
	time.Sleep(DebounceDuration + 1*time.Millisecond)

	// Same message should now go through again
	logger.Info("hello world")
	require.Len(t, h.records, 3)
	require.Equal(t, "hello world", h.records[2].Message)
}

func TestDebouncingHandler_MultipleMessages(t *testing.T) {
	h := new(testRecorder)
	d := NewDebouncingHandler(h)
	logger := log.NewLogger(d)

	// Send multiple different messages
	messages := []string{"msg1", "msg2", "msg3", "msg4", "msg5"}
	for _, msg := range messages {
		logger.Info(msg)
	}
	require.Len(t, h.records, len(messages))

	// Try to resend them immediately - all should be dropped
	for _, msg := range messages {
		logger.Info(msg)
	}
	require.Len(t, h.records, len(messages))

	// Wait for debounce period
	time.Sleep(DebounceDuration + 1*time.Millisecond)

	// Now they should all go through again
	for _, msg := range messages {
		logger.Info(msg)
	}
	require.Len(t, h.records, 2*len(messages))
}

func TestDebouncingHandler_CacheEviction(t *testing.T) {
	h := new(testRecorder)
	d := NewDebouncingHandler(h)
	logger := log.NewLogger(d)

	// Generate more than 1024 unique messages to trigger LRU eviction
	const numMessages = 1100
	for i := range numMessages {
		logger.Info(slog.IntValue(i).String())
	}
	require.Len(t, h.records, numMessages)

	// The earliest messages should have been evicted from cache
	// So they should go through again without waiting
	logger.Info(slog.IntValue(0).String())
	require.Len(t, h.records, numMessages+1)

	// Recent messages should still be debounced
	logger.Info(slog.IntValue(numMessages - 1).String())
	require.Len(t, h.records, numMessages+1)
}

func TestDebouncingHandler_SameMessageDifferentAttrs(t *testing.T) {
	h := new(testRecorder)
	d := NewDebouncingHandler(h)
	logger := log.NewLogger(d)

	// Log message with one set of attributes
	logger.Info("same message", "key1", "value1", "key2", "value2")
	require.Len(t, h.records, 1)
	require.Equal(t, "same message", h.records[0].Message)

	// Same message with different attributes should still be debounced
	logger.Info("same message", "key3", "value3", "key4", "value4")
	require.Len(t, h.records, 1)

	// Same message with no attributes should still be debounced
	logger.Info("same message")
	require.Len(t, h.records, 1)

	// Same message with partially overlapping attributes should still be debounced
	logger.Info("same message", "key1", "different_value", "key5", "value5")
	require.Len(t, h.records, 1)

	// Wait for debounce period
	time.Sleep(DebounceDuration + 1*time.Millisecond)

	// Now the same message with any attributes should go through
	logger.Info("same message", "totally", "new", "attrs", "here")
	require.Len(t, h.records, 2)
	require.Equal(t, "same message", h.records[1].Message)
}

func TestDebouncingHandler_TickerWarning(t *testing.T) {
	h := new(testRecorder)
	d := NewDebouncingHandler(h)
	logger := log.NewLogger(d)

	// Send initial message
	logger.Info("test message 1")
	require.Len(t, h.records, 1)
	require.Equal(t, "test message 1", h.records[0].Message)

	// Trigger several debounced messages
	for i := 0; i < 10; i++ {
		logger.Info("test message 1")
	}
	// Still only the first message
	require.Len(t, h.records, 1)

	// Send another unique message and debounce it
	logger.Info("test message 2")
	require.Len(t, h.records, 2)
	for i := 0; i < 5; i++ {
		logger.Info("test message 2")
	}

	// Wait for ticker to fire (5 seconds)
	time.Sleep(DebounceTickerInterval + 100*time.Millisecond)

	// Send a new message to trigger the ticker check
	logger.Info("trigger ticker check")

	// Should have: original 2 messages, warning about debounced messages, and the trigger message
	require.Len(t, h.records, 4)
	require.Equal(t, "test message 1", h.records[0].Message)
	require.Equal(t, "test message 2", h.records[1].Message)
	require.Equal(t, DebounceWarningMessage, h.records[2].Message)
	require.Equal(t, "trigger ticker check", h.records[3].Message)

	// Check that the warning record has the debounced count
	warningRecord := h.records[2]
	hasDebounceCount := false
	warningRecord.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "nDebounced" {
			require.Equal(t, uint64(15), attr.Value.Uint64()) // 10 + 5 debounced messages
			hasDebounceCount = true
		}
		return true
	})
	require.True(t, hasDebounceCount, "Warning should contain nDebounced attribute")

	// Counter should be reset, so debouncing more messages starts fresh
	for i := 0; i < 3; i++ {
		logger.Info("trigger ticker check")
	}
	// No new messages should be logged (they're debounced)
	require.Len(t, h.records, 4)

	// Wait for ticker again
	time.Sleep(DebounceTickerInterval + 100*time.Millisecond)

	// Trigger ticker check
	logger.Info("final message")

	// Should have another warning for the 3 newly debounced messages
	require.Len(t, h.records, 6)
	require.Equal(t, DebounceWarningMessage, h.records[4].Message)
	require.Equal(t, "final message", h.records[5].Message)

	// Check the second warning has count of 3
	secondWarning := h.records[4]
	secondWarning.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "nDebounced" {
			require.Equal(t, uint64(3), attr.Value.Uint64())
		}
		return true
	})
}

func TestDebouncingHandler_Concurrent(t *testing.T) {
	h := new(safeTestRecorder)
	d := NewDebouncingHandler(h)
	logger := log.NewLogger(d)

	const numGoroutines = 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := range numGoroutines {
		go func(id int) {
			defer wg.Done()
			logger.Info("hello")
		}(i)
	}

	wg.Wait()

	require.Equal(t, h.Len(), 1, "Should have debounced duplicate messages")
}
