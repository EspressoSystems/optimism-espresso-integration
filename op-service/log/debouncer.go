package log

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	// DebounceDuration is the time window during which duplicate messages are suppressed
	DebounceDuration = 100 * time.Millisecond
	// DebounceTickerInterval is how often we check and report debounced message counts
	DebounceTickerInterval = 5 * time.Second
	// DebounceWarningMessage is the message logged when messages have been debounced
	DebounceWarningMessage = "Some messages were debounced"
)

type DebounchingHandler struct {
	handler  slog.Handler
	messages *lru.Cache[string, time.Time]
	counter  atomic.Uint64
	ticker   *time.Ticker
}

func NewDebouncingHandler(handler slog.Handler) *DebounchingHandler {
	messages, _ := lru.New[string, time.Time](1024)
	return &DebounchingHandler{
		handler:  handler,
		messages: messages,
		ticker:   time.NewTicker(DebounceTickerInterval),
	}
}

func (h *DebounchingHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return h.handler.Enabled(ctx, lvl)
}

func (h *DebounchingHandler) Handle(ctx context.Context, record slog.Record) error {
	select {
	case <-h.ticker.C:
		cntr := h.counter.Load()
		h.counter.Store(0)

		if cntr > 0 {
			warningRecord := slog.NewRecord(time.Now(), slog.LevelWarn, DebounceWarningMessage, 0)
			warningRecord.Add("nDebounced", cntr)
			err := h.handler.Handle(ctx, warningRecord)
			if err != nil {
				return err
			}
		}

	default:
	}

	if last, ok := h.messages.Get(record.Message); ok && time.Since(last) < DebounceDuration {
		h.counter.Add(1)
		return nil
	}
	h.messages.Add(record.Message, time.Now())

	return h.handler.Handle(ctx, record)
}

func (h *DebounchingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewDebouncingHandler(h.handler.WithAttrs(attrs))
}

func (h *DebounchingHandler) WithGroup(name string) slog.Handler {
	return NewDebouncingHandler(h.handler.WithGroup(name))
}
