package batcher_test

import (
	"context"
	"encoding/binary"
	"testing"
	"time"

	espressoCommon "github.com/EspressoSystems/espresso-network/sdks/go/types"
	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
)

// TestEspressoTransactionSubmitterDeadlock is a test that is meant to test
// the underlying conditions in the espressoTransactionsubmitter that may
// result in a deadlock.
//
// The basic idea is to intentionally trigger the circumstances that can
// trigger the deadlock to occur in the espressoTransactionSubmitter. This
// test with this description **SHOULD** remain as a regression test.
//
// The specific suspected criteria for triggering a deadlock in the
// implementation of the espressoTransactionSubmitter  as of 2026-04-23 is
// as follows:
//
// - Assume that the Espresso Service is not ever returning successfully
// - Assume we are receiving a consistent stream of of new blocks coming in
//
// If we have both of these criteria, it should be possible to fill up the
// underlying channels of `espressoTransactionSubmitter` for both Submission
// jobs, and verification jobs.
//
// The way we *should* be able to detect the deadlock is if a call to
// SubmitTransaction blocks.
//
// NOTE: After this has been fixed, it is unlikely that this will fail in the
// future. This will primarily be due to `SubmitTransaction` being modified
// to longer explicitly block.
func TestEspressoTransactionSubmitterDeadlock(t *testing.T) {
	submitter := batcher.NewEspressoTransactionSubmitter(
		batcher.WithEspressoClient(new(AlwaysFailingEspressoClient)),
	)

	submitter.SpawnWorkers(4, 4)
	submitter.Start()

	binary.LittleEndian.PutUint64(make([]byte, 8), 1)

	for i := range 1024 * 4 {
		txn := espressoCommon.Transaction{
			Namespace: 1,
			Payload:   make([]byte, 8),
		}
		binary.LittleEndian.PutUint64(txn.Payload, uint64(i))

		ctx, cancel := context.WithCancel(context.Background())
		go (func(cancel context.CancelFunc, txn espressoCommon.Transaction) {
			submitter.SubmitTransaction(&txn)
			cancel()
		})(cancel, txn)

		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

		select {
		case <-ctx.Done():
			// Things progressed without issue.
			timeoutCancel()
			cancel()
			continue
		case <-timeoutCtx.Done():
			// The call to SubmitTransaction did not return within the expected
			// time frame.

			timeoutCancel()
			cancel()
			t.Errorf("SubmitTransaction has blocked for longer than 200 milliseconds, on transaction %d, for error: %s\n", i, timeoutCtx.Err())
			return
		}
	}

	// If we've gotten here, we have passed without deadlocking.
}
