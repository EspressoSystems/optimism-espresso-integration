package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// checkBatcherSubmission checks if the batcher has submitted the L2 block corresponding to the given L1 block number.
func checkBatcherSubmission(ctx context.Context, l1Client *ethclient.Client, blockNum uint64) bool {
        _, err := l1Client.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
        if err != nil {
            return false
        }

        return true
}

// TestBatcherWaitForFinality is a test that attempts to make sure that the batcher waits for the
// derived L1 block to be finalized before submitting a new block.
//
// This tests is designed to evaluate Test 8.2 as outlined within the Espresso Celo Integration
// plan. It has stated task definition as follows:
//
//	Arrange:
//		Running the sequencer and the batcher in Espresso mode.
//	Act:
//		Send a single transaction derived from an unfinalized L1 block.
//	Assert:
//		The batcher doesn't submit the transaction to the L1 immediatly.
//		After the derived L1 block is finalized, the batcher submits the transaciton.
func TestBatcherWaitForFinality(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t)
	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer system.Close()
	defer espressoDevNode.Stop()

	caffNode, err := env.LaunchDecaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer caffNode.Close(ctx)

	// We want to setup our test condition
	addressAlice := system.Cfg.Secrets.Addresses().Alice

	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	// Increase Alice's balance by 1 via a deposit transaction
    privateKey := system.Cfg.Secrets.Bob
    bobOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, system.Cfg.L1ChainIDBig())
    if err != nil {
        t.Fatalf("failed to create transaction options for Bob: %v", err)
    }

    mintAmount := big.NewInt(1)
    bobOptions.Value = mintAmount

	_ = helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, bobOptions, func(l2Opts *helpers.DepositTxOpts) {
        l2Opts.ToAddr = addressAlice
    })

	depositReceipt2 := helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, bobOptions, func(l2Opts *helpers.DepositTxOpts) {
        l2Opts.ToAddr = addressAlice
    })

	depositBlockNumber2 := depositReceipt2.BlockNumber.Uint64()
	t.Logf("Deposit landed in L1: %d", depositBlockNumber2)

    // Assert that the batcher has not submitted the L2 block to L1 before L1 finalization
    if checkBatcherSubmission(ctx, l1Client, depositBlockNumber2) {
        t.Fatalf("batcher submitted the L2 block to L1 before L1 block was finalized")
    }

	// TODO (Keyao) Find a proper time
	time.Sleep(5 * time.Second)

    // Assert that the batcher has submitted the L2 block to L1 after L1 finalization
    if !checkBatcherSubmission(ctx, l1Client, depositBlockNumber2) {
        t.Fatalf("batcher did not submit the L2 block to L1 after L1 block was finalized")
    }
}
