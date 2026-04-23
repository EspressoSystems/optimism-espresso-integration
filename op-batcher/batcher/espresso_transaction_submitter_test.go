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
// the underlying conditions in the espressoTransactionSubmitter that may
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
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	submitter := batcher.NewEspressoTransactionSubmitter(
		batcher.WithEspressoClient(new(AlwaysFailingEspressoClient)),
		batcher.WithContext(ctx),
	)

	submitter.SpawnWorkers(4, 4)
	submitter.Start()

	for i := 0; i < 1024*4; i++ {
		txn := espressoCommon.Transaction{
			Namespace: 1,
			Payload:   make([]byte, 8),
		}
		binary.LittleEndian.PutUint64(txn.Payload, uint64(i))

		submitCtx, submitCancel := context.WithCancel(context.Background())
		go (func(cancel context.CancelFunc, txn espressoCommon.Transaction) {
			submitter.SubmitTransaction(&txn)
			cancel()
		})(submitCancel, txn)

		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

		select {
		case <-submitCtx.Done():
			// Things progressed without issue.
			timeoutCancel()
			submitCancel()
			continue
		case <-timeoutCtx.Done():
			// The call to SubmitTransaction did not return within the expected
			// time frame.

			timeoutCancel()
			submitCancel()
			t.Fatalf("SubmitTransaction has blocked for longer than 200 milliseconds, on transaction %d, for error: %s\n", i, timeoutCtx.Err())
		}
	}

	// If we've gotten here, we have passed without deadlocking.
}
