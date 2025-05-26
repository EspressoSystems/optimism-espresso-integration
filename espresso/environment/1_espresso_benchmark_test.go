package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/espresso/environment/benchmark"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	geth_types "github.com/ethereum/go-ethereum/core/types"
)

// TestE2eDevNetWithEspressoFastConfirmationStability is a test that tests
// the benchmarking setup of the Espresso Caff Node's performance versus the
// L2 Verifier derived from the L1.
//
// This test is designed to evaluate Espresso's impact while under load on
// the Optimism stack.  The point of this test is to ensure that, even under
// heavy load, the Espresso Caff Node can maintain its performance and
// not introduce significant delays in the confirmation process.
//
// This test spins up the E2E dev net with the espresso-dev-node and the
// Caff Node.  It then runs a benchmarks that submits transactions to the
// L2 Sequencer and observes the time it takes to reach each stage of the
// confirmation process.  The test will run for a couple of minutes and
// then check the statistics to ensure that the performance is within
// acceptable limits.
//
// The acceptance criteria of this test is stated to be that the time
// taken between each stage of this process should not change significantly
// over time.
//
// It is difficult to meet this criteria as it is stated with vague terms
// and with the intention of a much longer runtime duration than what we'd
// want when evaluating this consistency.
//
// Instead this test will Run for 2 minutes, with Block Times set to the
// values of the typical L1 and L2 block times. It will place a load
// upon it, and it will check the standard deviation of the time taken
// between each stage of the confirmation process in order to make sure
// that they do not exceed a "reasonable" value.
//
// For the purposes of this test the "reasonable" value is defined to
// be 2 seconds.
func TestE2eDevNetWithEspressoFastConfirmationStability(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)
	system, espressoDevNode, err := launcher.StartDevNet(
		ctx,
		t,
		env.WithSequencerUseFinalized(true),
		env.WithL1BlockTime(12*time.Second),
		env.WithL2BlockTime(2*time.Second),
	)

	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	caffNode, err := env.LaunchDecaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer env.Stop(t, caffNode)

	keys := system.Cfg.Secrets
	addresses := keys.Addresses()

	signer := geth_types.LatestSignerForChainID(system.Cfg.L2ChainIDBig())

	// Submit Transactions to the Sequencer
	bencher := benchmark.CreateBenchmarker(
		ctx,
		benchmark.WithSeqClient(system.NodeClient(e2esys.RoleSeq)),
		benchmark.WithCaffClient(system.NodeClient(env.RoleCaffNode)),
		benchmark.WithVerifyClient(system.NodeClient(e2esys.RoleVerif)),
		benchmark.AddSubmitter(benchmark.BenchmarkSubmitterConfig{
			Interval: 100 * time.Millisecond,
			To:       &addresses.Bob,
			Value:    big.NewInt(1),
			Signer:   signer,
			ChainID:  system.RollupConfig.L2ChainID,
			Key:      keys.Alice,
		}),
		benchmark.AddSubmitter(benchmark.BenchmarkSubmitterConfig{
			Interval: 100 * time.Millisecond,
			To:       &addresses.Mallory,
			Value:    big.NewInt(1),
			Signer:   signer,
			ChainID:  system.RollupConfig.L2ChainID,
			Key:      keys.Bob,
		}),
		benchmark.AddSubmitter(benchmark.BenchmarkSubmitterConfig{
			Interval: 100 * time.Millisecond,
			To:       &addresses.Alice,
			Value:    big.NewInt(1),
			Signer:   signer,
			ChainID:  system.RollupConfig.L2ChainID,
			Key:      keys.Mallory,
		}),
	)

	// Alright, let's run the benchmark for a couple of minutes

	{
		ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()

		stats, err := bencher.RunWithContext(ctx)
		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to run benchmark:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		// Let's check the statistics
		metrics := benchmark.ComputeRawMetricsStatistics(stats)

		if have, want := metrics.SubmittedToReceipt.Count, 0; have <= want {
			t.Errorf("expected to have a positive count for receipts received:\nhave:\n\t\"%v\"\nwant:\n\t> \"%v\"\n", have, want)
		}

		if have, want := metrics.ReceiptToCaff.Count, 0; have <= want {
			t.Errorf("expected to have a positive count for caff headers:\nhave:\n\t\"%v\"\nwant:\n\t> \"%v\"\n", have, want)
		}

		if have, want := metrics.ReceiptToVerify.Count, 0; have <= want {
			t.Errorf("expected to have a positive count for verify headers:\nhave:\n\t\"%v\"\nwant:\n\t> \"%v\"\n", have, want)
		}

		// We do not expect a signification amount of variance or std deviation
		if have, want := metrics.SubmittedToReceipt.StdDev, 2*time.Second; have > want {
			t.Errorf("expected a small amount of variance in the submitted to receipt time:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		if have, want := metrics.ReceiptToCaff.StdDev, 2*time.Second; have > want {
			t.Errorf("expected a small amount of variance in the receipt to caff time:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		if have, want := metrics.ReceiptToVerify.StdDev, 2*time.Second; have > want {
			t.Errorf("expected a small amount of variance in the receipt to L1 time:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}
	}

}
