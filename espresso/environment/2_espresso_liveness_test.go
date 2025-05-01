package environment_test

import (
	"context"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
)

// TestE2eDevNetWithEspressoEspressoDegradedLiveness is a test that checks that
// the rollup will continue to make progress even in the event of intermittent
// Espresso system failures.
//
// The Criteria for this test is as follows:
//
//	Requirement: Resubmission to Espresso.
//		Randomy turn the Espresso builder off and on. Check that the rollup
//		continues to make progress, including progressing settlement on the
//		base layer.
//
// We don't have any direct way of turning the Espresso builder off and on via
// the Dev node API at the moment.  However, we do have the ability to turn
// the consensus layer on and off via turning hotshot on and off.
//
// This is **NOT** the same thing, nor would it result in the same behavior as
// turning the Builder off and on. For the following reasons:
//
//	1 HotShot being off means no new blocks are being produced
//	2 The Builder being off means that only empty blocks are being produced
//	3 Turning the Builder off potentially means losing pool information,
//	  requiring re-submission so that the builder can include the transaction
//	  in the next block.
//
// With these caveats in mind, we may be able to simulate the behavior of 2
// at the very least, if we intercept the client submitting transactions to
// Espresso, and simulating the client being unable to submit transactions.
// Likewise, we might be able to simulate 3 by falsely reporting to the
// submitter that the transaction was submitted successfully, and withholding
// the submission itself.
func TestE2eDevNetWithEspressoEspressoDegradedLiveness(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// Start a Server to proxy requests to Espresso
	proxy, server, option := env.SetupQueryServiceIntercept()
	defer server.Close()
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0, option)

	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer system.Close()
	defer espressoDevNode.Stop()
	proxy.TurnOnBehavior(env.BehaviorRandomRollTxnSubmissionFailure)

	const N = 10
	{
		for i := 0; i < N; i++ {
			runSimpleL1TransferAndVerifier(ctx, t, system)
		}
	}
}
