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

// TestE2eDevNetWithEspressoFastConfirmationStability is a test that attempts
// to verify that the Espresso pipeline is able to perform consistently
// under load without causing any significant delays in the confirmation
// / finalization process.
func TestE2eDevNetWithEspressoFastConfirmationStability(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)
	system, espressoDevNode, err := launcher.StartDevNet(
		ctx,
		t,
		env.WithSequencerUseFinalized(true),
		env.WithL1BlockTime(13*time.Second),
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
		individualTransactions := stats.IndividualTransactionMetrics()
		metrics := benchmark.ComputeCompletedTransactionStatistics(individualTransactions)

		// We do not expect a signification amount of variance or std deviation
		if have, want := metrics.SubmittedToReceipt.StdDev, metrics.SubmittedToReceipt.Mean/10; have > want {
			t.Errorf("expected a small amount of variance in the submitted to receipt time:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		if have, want := metrics.ReceiptToCaff.StdDev, metrics.ReceiptToCaff.Mean/10; have > want {
			t.Errorf("expected a small amount of variance in the receipt to caff time:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		if have, want := metrics.ReceiptToVerify.StdDev, metrics.ReceiptToVerify.Mean/10; have > want {
			t.Errorf("expected a small amount of variance in the receipt to L1 time:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}
	}

}
