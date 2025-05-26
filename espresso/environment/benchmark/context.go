package benchmark

import (
	"context"
	"sync"
	"time"
)

// waitGroupContext is a helpful struct that facilitates the
// pieces of a worker that are important to keep track of in order
// to facilitate their cancellation, shutdown, and synchronization.
type waitGroupContext struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// Deadline implements context.Context.
func (w *waitGroupContext) Deadline() (deadline time.Time, ok bool) {
	return w.ctx.Deadline()
}

// Done implements context.Context.
func (w *waitGroupContext) Done() <-chan struct{} {
	return w.ctx.Done()
}

// Err implements context.Context.
func (w *waitGroupContext) Err() error {
	return w.ctx.Err()
}

// Value implements context.Context.
func (w *waitGroupContext) Value(key any) any {
	return w.ctx.Value(key)
}

var _ context.Context = (*waitGroupContext)(nil)

// NewCancelContext creates a new context that is cancellable and
// has a wait group associated with it.
func NewCancelContext(ctx context.Context) waitGroupContext {
	ctx, cancel := context.WithCancel(ctx)
	return waitGroupContext{
		wg:     sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
	}
}
