package environment_test

import (
	"context"
	"log/slog"
	"math/big"
	"math/rand"
	"sync"
	"testing"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	"github.com/ethereum-optimism/optimism/espresso"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// TestE2eDevNetWithEspressoEspressoDegradedLiveness is a test that checks that
// the rollup will continue to make progress even in the event of intermittent
// Espresso system failures.
//
// The Criteria for this test is as follows:
//
//	Requirement: Resubmission to Espresso.
//		Randomly turn the Espresso builder off and on. Check that the rollup
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
	_, server, option := env.SetupQueryServiceIntercept(
		// This decider will randomly report successful submissions of
		// transactions to Espresso, but will not actually submit them.
		// This will approximately occur 10% of the time, given the
		// criteria to roll a number 0-9 and only to occur if the rolled
		// number is 0.
		env.SetDecider(env.NewRandomRollFakeSubmitTransactionSuccess(
			10,
			0,
			1,
			rand.New(rand.NewSource(0)),
		)),
	)

	defer server.Close()
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0, option)

	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer system.Close()
	defer espressoDevNode.Stop()

	addressAlice := system.Cfg.Secrets.Addresses().Alice

	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	balanceAliceInitial, err := l2Verif.BalanceAt(ctx, addressAlice, nil)
	if have, want := err, error(nil); have != want {
		t.Fatalf("Failed to fetch Alice's balance:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	const N = 10
	{
		var receipts []*geth_types.Receipt

		for i := 0; i < N; i++ {
			receipt := helpers.SendL2TxWithID(t, system.Cfg.L2ChainIDBig(), l2Seq, system.Cfg.Secrets.Bob, func(opts *helpers.TxOpts) {
				opts.Nonce = uint64(i)
				opts.ToAddr = &addressAlice
				opts.Value = big.NewInt(1)
			})

			receipts = append(receipts, receipt)
		}

		// Let's verify that all of our transactions came through successfully
		for _, receipt := range receipts {
			_, err := wait.ForReceiptOK(ctx, l2Verif, receipt.TxHash)
			if have, want := err, error(nil); have != want {
				t.Fatalf("Waiting for L2 tx on verification client:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}
		}

		// Alice's balance should have increased by N
		balanceAliceFinal, err := l2Verif.BalanceAt(ctx, addressAlice, nil)
		if have, want := err, error(nil); have != want {
			t.Fatalf("Failed to fetch Alice's balance:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		expectedBalance := new(big.Int).Add(balanceAliceInitial, big.NewInt(int64(N)))
		if balanceAliceFinal.Cmp(expectedBalance) != 0 {
			t.Fatalf("Alice's balance did not increase as expected:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", balanceAliceFinal, expectedBalance)
		}
	}
}

// TestE2eDevNetWithEspressoEspressoDegradedLivenessViaCaffNode is a test that
// checks that Espresso will return fast confirmations even when in a
// degraded state.
//
// The criteria for this test is as follows:
//	Requirement: Liveness:
//    The rollup should continue to run, [to] post Espresso confirmations
//    within 10 seconds of each rollup block produced by the sequencer.
//
// As a result, this test will submit a number of transactions to the sequencer,
// while also consuming the Espresso stream of blocks utilizing the Espresso
// streamer.  We **SHOULD** be able to match up the transactions submitted to
// the blocks being produced by the Espresso Streamer, and the time it takes
// from transaction submission to receiving the Block that contains that same
// transaction should be less than 10 seconds.
//
// More importantly, this **SHOULD** also continue to be the state even when
// Espresso is in a degraded state.
//
// Sadly, there does not seem to be an easy way to associate the Transaction
// submitted to the L2 with the Block being returned from Espresso.  However,
// in this test scenario, we know that the Batch number will correspond to the
// transaction number we submitted to the sequencer.  As a result, we should
// be able to match the Batch number to the transaction number in transaction
// order.
// Sadly, there does not seem to be an easy way to associate the Transaction
// submitted to the L2 with the Block being returned from Espresso.  However,
// in this test scenario, we know that the Batch number will correspond to the
// transaction number we submitted to the sequencer.  As a result, we should
// be able to match the Batch number to the transaction number in transaction
// order.
// Sadly, there does not seem to be an easy way to associate the Transaction
// submitted to the L2 with the Block being returned from Espresso.  However,
// in this test scenario, we know that the Batch number will correspond to the
// transaction number we submitted to the sequencer.  As a result, we should
// be able to match the Batch number to the transaction number in transaction
// order.
// Sadly, there does not seem to be an easy way to associate the Transaction
// submitted to the L2 with the Block being returned from Espresso.  However,
// in this test scenario, we know that the Batch number will correspond to the
// transaction number we submitted to the sequencer.  As a result, we should
// be able to match the Batch number to the transaction number in transaction
// order.

func TestE2eDevNetWithEspressoEspressoDegradedLivenessViaCaffNode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// Start a Server to proxy requests to Espresso, with a decider that will
	// simulate degraded liveness failures by reporting false successful
	// submissions 10% of the time, and 503 errors 10% of the time, with
	// actual proxied requests 80% of the time.
	_, server, option := env.SetupQueryServiceIntercept(
		env.SetDecider(env.NewRandomRollFakeSubmitTransactionSuccess(
			10,
			0,
			1,
			rand.New(rand.NewSource(0)),
		)),
	)

	defer env.Stop(t, server)
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0, option)

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

	addressAlice := system.Cfg.Secrets.Addresses().Alice

	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Seq := system.NodeClient(e2esys.RoleSeq)
	caffVerif := system.NodeClient(env.RoleCaffNode)

	balanceAliceInitial, err := caffVerif.BalanceAt(ctx, addressAlice, nil)
	if have, want := err, error(nil); have != want {
		t.Fatalf("Failed to fetch Alice's balance:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	type espressoReceived struct {
		batch    *derive.EspressoBatch
		block    *geth_types.Block
		received time.Time
	}

	espressoReceipts := map[uint64]espressoReceived{}

	streamBlocksCtx, streamBlocksCancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	defer streamBlocksCancel()
	{
		// Streamer Setup and Configuration
		l := log.NewLogger(slog.Default().Handler())
		streamer := espresso.NewEspressoStreamer(
			system.RollupConfig.L2ChainID.Uint64(),
			batcher.NewAdaptL1BlockRefClient(l1Client),
			espressoClient.NewClient(server.URL),
			nil, // TODO(AG)
			l,
			func(b []byte) (*derive.EspressoBatch, error) {
				return derive.UnmarshalEspressoTransaction(b, system.RollupConfig.Genesis.SystemConfig.BatcherAddr)
			},
			100*time.Millisecond,
		)

		l1Client, _ := client.NewRPC(streamBlocksCtx, l, system.NodeEndpoint(e2esys.RoleL1).RPC())
		l2Seq, _ := client.NewRPC(streamBlocksCtx, l, system.NodeEndpoint(e2esys.RoleSeq).RPC())

		l1RefClient, err := sources.NewL1Client(l1Client, l, nil, sources.L1ClientDefaultConfig(system.RollupConfig, true, sources.RPCKindStandard))
		require.NoError(t, err, "failed to create L1 Ref client")
		l2RefClient, err := sources.NewL2Client(l2Seq, l, nil, sources.L2ClientDefaultConfig(system.RollupConfig, true))
		require.NoError(t, err, "failed to create L2 Ref client")
		l2BlockRef, err := l2RefClient.L2BlockRefByLabel(streamBlocksCtx, eth.Safe)
		require.NoError(t, err, "failed to get safe L2 block ref")
		finalizedL1BlockRef, err := l1RefClient.L1BlockRefByLabel(streamBlocksCtx, eth.Finalized)
		require.NoError(t, err, "failed to get finalized L1 block ref")
		streamer.Refresh(streamBlocksCtx, finalizedL1BlockRef, l2BlockRef.Number)

		// Start consuming Batches from the Streamer
		// We cannot guarantee the order of these batches coming from the
		// streamer.  However, luckily, we can rely on the fact that the Batch
		// number will be stored within the Batch itself, and we can use that
		// to match up with the order that the transactions are being submitted
		// to the sequencer.
		wg.Add(1)
		go (func(ctx context.Context, wg *sync.WaitGroup, streamer espresso.EspressoStreamer[derive.EspressoBatch]) {
			cfg := system.RollupConfig
			defer wg.Done()
			for {
				select {
				default:
				case <-ctx.Done():
					// We are being told to exit, so we exit
					return
				}

				if !streamer.HasNext(ctx) {
					if err := streamer.Update(ctx); err != nil {
						// Try again?
						time.Sleep(100 * time.Millisecond)
					}
					continue
				}

				// consume all of the available batches
				for batch := streamer.Next(ctx); batch != nil; batch = streamer.Next(ctx) {
					block, err := batch.ToBlock(cfg)
					if err != nil {
						// Try again?
						time.Sleep(100 * time.Millisecond)
						continue
					}

					espressoReceipts[batch.Number()] = espressoReceived{
						batch:    batch,
						block:    block,
						received: time.Now(),
					}
				}
			}
		})(streamBlocksCtx, &wg, streamer)
	}

	type submission struct {
		receipt   *geth_types.Receipt
		created   time.Time
		submitted time.Time
		received  time.Time
	}
	var submissions []submission

	// The number of transaction we want to submit to the L2.
	// This will also correspond to the number of batches we expect to receive
	// from the Espresso streamer.
	const N = 10
	{

		for i := 0; i < N; i++ {
			// Create teh transaction
			tx := geth_types.MustSignNewTx(system.Cfg.Secrets.Bob, geth_types.LatestSignerForChainID(system.Cfg.L2ChainIDBig()), &geth_types.DynamicFeeTx{
				ChainID:   system.Cfg.L2ChainIDBig(),
				Nonce:     uint64(i),
				To:        &addressAlice,
				Value:     big.NewInt(1),
				GasTipCap: big.NewInt(10),
				GasFeeCap: big.NewInt(200),
				Gas:       21_000,
			})
			created := time.Now()

			// Send the transaction
			err := l2Seq.SendTransaction(ctx, tx)
			if have, want := err, error(nil); have != want {
				t.Fatalf("Sending L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}

			// We have submitted the transaction to the L2, successfully.
			submitted := time.Now()

			// Wait for the receive
			receipt, err := wait.ForReceiptOK(ctx, l2Seq, tx.Hash())
			if have, want := err, error(nil); have != want {
				t.Fatalf("Waiting for L2 tx:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}

			// We not have a receipt from the L2 Sequencer, indicating that
			// the transaction was successfully included in a block.
			received := time.Now()

			submissions = append(submissions, submission{
				receipt:   receipt,
				created:   created,
				submitted: submitted,
				received:  received,
			})
		}

		// Let's verify that all of our transactions came through successfully,
		// using our Caff Node as the verification client.
		for i, submission := range submissions {
			receipt, err := wait.ForReceiptOK(ctx, caffVerif, submission.receipt.TxHash)
			if have, want := err, error(nil); have != want {
				t.Fatalf("Waiting for L2 tx on verification client:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}

			// Transaction Hash should match
			if have, want := receipt.TxHash, submission.receipt.TxHash; have != want {
				t.Errorf("Receipt tx hash mismatch for submission %d:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", i, have, want)
			}

			// Block Hash should match
			if have, want := receipt.BlockHash, submission.receipt.BlockHash; have != want {
				t.Errorf("Receipt block hash mismatch for submission %d:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", i, have, want)
			}
		}

		// Alice's balance should have increased by N
		balanceAliceFinal, err := caffVerif.BalanceAt(ctx, addressAlice, nil)
		if have, want := err, error(nil); have != want {
			t.Fatalf("Failed to fetch Alice's balance:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		expectedBalance := new(big.Int).Add(balanceAliceInitial, big.NewInt(int64(N)))
		if balanceAliceFinal.Cmp(expectedBalance) != 0 {
			t.Errorf("Alice's balance did not increase as expected:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", balanceAliceFinal, expectedBalance)
		}
	}

	// Tell the Streamer to stop streaming.
	streamBlocksCancel()
	wg.Wait()

	// We'll check that our timings meet or exceed the requirements of the test.
	var totalDiff time.Duration
	for i, submission := range submissions {
		espressoReceived, ok := espressoReceipts[uint64(i+1)]

		if have, want := ok, true; have != want {
			t.Errorf("Failed to find espresso block for submission %d:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", i, have, want)
			continue
		}

		diff := espressoReceived.received.Sub(submission.submitted)
		totalDiff += diff

		if have, want := diff, 10*time.Second; have > want {
			t.Errorf("Submission %d was not confirmed in an espresso block within 10 seconds of submission:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", i, diff, want)
		}
	}

	averageDuration := totalDiff / N
	if have, want := averageDuration, 10*time.Second; have > want {
		t.Errorf("Average time to confirm transactions in espresso blocks exceeded 10 seconds:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", averageDuration, want)
	}
}
