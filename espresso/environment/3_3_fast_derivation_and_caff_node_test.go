package environment_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// checkNewBlocks is a helper function for TestFastDerivationAndCaffNode to check for new blocks by comparing the hash of new block and previous block
func checkNewBlocks(ctx context.Context, client *ethclient.Client, previousBlock *types.Block, nodeName string, tickerDuration time.Duration) (*types.Block, error) {
	newBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get new %s block: %w", nodeName, err)
	}

	// Make sure newBlock comes after previousBlock
	if have, want := newBlock.Number(), previousBlock.Number(); have.Cmp(want) <= 0 {
		return nil, fmt.Errorf("No new block for %s after %s\nhave:\n\t\"%v\"\nwant:\n\t> \"%v\"\n", nodeName, tickerDuration, have, want)
	}
	return newBlock, nil
}

// TestFastDerivationAndCaffNode is a test that
// checks the derivation pipeline is fast and the Caff node is working properly with the happy path.
//
// The criteria for this test is as follows:
//
//	Requirement:
//	   Make sure the node's RPC can be queried with update every 2-4 seconds.
//
// Arrange:
//
//	Running Sequencer, Batcher in Espresso mode, and Caff node with happy path.
//
// Act:
//
//	Submit a number of transactions (or no transaction?) to the sequencer
//
// Assert:
//
//	We should be able to query caff node with update every 2-4 seconds.  We use ticker to query the node every 4 seconds.
//
// checkNewBlocks checks for new blocks and verifies their timestamps
func TestFastDerivationAndCaffNode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, env.WithL1FinalizedDistance(0), env.WithSequencerUseFinalized(true))

	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	caffNode, err := env.LaunchCaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer env.Stop(t, caffNode)

	addressAlice := system.Cfg.Secrets.Addresses().Alice
	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	caffVerif := system.NodeClient(env.RoleCaffNode)

	// We want to send some transactions from Bob to Alice
	{
		privateKey := system.Cfg.Secrets.Bob
		bobOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, system.Cfg.L1ChainIDBig())
		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to create transaction options for bob:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		mintAmount := new(big.Int).SetUint64(1)
		bobOptions.Value = mintAmount
		_ = helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, bobOptions, func(l2Opts *helpers.DepositTxOpts) {
			// Send from Bob to Alice
			l2Opts.ToAddr = addressAlice
		})
	}

	// Initialize ticker to fire every 4 seconds
	tickerDuration := 4 * time.Second
	ticker := time.NewTicker(tickerDuration)
	defer ticker.Stop()

	finishTicker := time.NewTicker(30 * time.Second)
	defer finishTicker.Stop()

	lastCaffHead, err := caffVerif.BlockByNumber(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to get initial caffVerif block: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Check for new block of caff-node
			newCaff, err := checkNewBlocks(ctx, caffVerif, lastCaffHead, "caff-node", tickerDuration)
			if have, want := err, error(nil); have != want {
				t.Fatalf("failed to get new caff-node block:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}
			lastCaffHead = newCaff
		case <-finishTicker.C:
			return
		}
	}

}
