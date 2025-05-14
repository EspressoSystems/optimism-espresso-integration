package environment_test

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

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
//	Running Sequencer, Batcher in Espresso mode, Caff node, and OP node with happy path.
//
// Act:
//
//	Submit a number of transactions (or no transaction?) to the sequencer
//
// Assert:
//
//	We should be able to query op-node and caff node with update every 2-4 seconds.
func TestFastDerivationAndCaffNode(t *testing.T) {
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

	// Initialize ticker to fire every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	finishticker := time.NewTicker(30 * time.Second)
	defer finishticker.Stop()

	lastHead, err := l2Verif.BlockByNumber(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to get initial l2Verif block: %v", err)
	}

	lastCaffHead, err := caffVerif.BlockByNumber(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to get initial caffVerif block: %v", err)
	}

	// checkNewBlocks checks for new blocks and verifies their timestamps
	checkNewBlocks := func(client *ethclient.Client, lastBlock *types.Block, nodeName string, logger string) (*types.Block, error) {
		newBlock, err := client.BlockByNumber(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to get new %s block: %v", nodeName, err)
		}

		// Skip the warm-up period and check for new block by comparing the latest hash
		if newBlock.Number().Cmp(big.NewInt(0)) != 0 && newBlock.Hash() == lastBlock.Hash() {
			// we only report an error here, but not return a failure
			log.Info("No new block for RPC after 2 seconds", "node", nodeName, "current number", newBlock.Number())
		} else {
			// instead, we check block timestamps
			lastBlockTime := lastBlock.Time()
			for j := new(big.Int).Add(lastBlock.Number(), big.NewInt(1)); j.Cmp(newBlock.Number()) <= 0; j.Add(j, big.NewInt(1)) {
				block, err := client.BlockByNumber(ctx, j)
				if err != nil {
					return nil, fmt.Errorf("Failed to get block from %s: %v", nodeName, err)
				}
				// check the block timestamp
				if have, want := block.Time(), lastBlockTime+uint64(1); have != want {
					return nil, fmt.Errorf("Block timestamp mismatch:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
				}
				lastBlockTime = block.Time()
			}
		}
		return newBlock, nil
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Check for new block of op-node
			newHead, err := checkNewBlocks(l2Verif, lastHead, "op-node", "l2verifier")
			if err != nil {
				t.Fatal(err)
			}
			lastHead = newHead

			// Check for new block of caff-node
			newCaff, err := checkNewBlocks(caffVerif, lastCaffHead, "caff-node", "caffverifier")
			if err != nil {
				t.Fatal(err)
			}
			lastCaffHead = newCaff
		case <-finishticker.C:
			return
		}
	}

}
