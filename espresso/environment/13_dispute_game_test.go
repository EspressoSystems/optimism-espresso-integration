package environment_test

import (
	"context"
	"math/rand"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	op_e2e "github.com/ethereum-optimism/optimism/op-e2e"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// Sishan TODO: add comment and description
// Test the disput game still work as expected with Espresso
func TestOutputAlphabetGameWithEspresso_ChallengerWins(t *testing.T) {
	op_e2e.InitParallel(t)
	ctx := context.Background()

	/////// start of espresso code
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
	sys, espressoDevNode, err := launcher.StartDevNetWithFaultDisputeSystem(ctx, t, option)
	l1Client := sys.NodeClient("l1")

	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer sys.Close()
	defer func() {
		err = espressoDevNode.Stop()
		if err != nil {
			t.Fatalf("failed to stop espresso dev node: %v", err)
		}
	}()
	/////// end of espresso code

	// Pasted from `TestOutputAlphabetGame_ChallengerWins` in `op-e2e/faultproofs/output_alphabet_test.go`
	disputeGameFactory := disputegame.NewFactoryHelper(t, ctx, sys)
	game := disputeGameFactory.StartOutputAlphabetGame(ctx, "sequencer", 3, common.Hash{0xff})
	correctTrace := game.CreateHonestActor(ctx, "sequencer")
	game.LogGameData(ctx)

	opts := challenger.WithPrivKey(sys.Cfg.Secrets.Alice)
	game.StartChallenger(ctx, "sequencer", "Challenger", opts)
	game.LogGameData(ctx)

	// Challenger should post an output root to counter claims down to the leaf level of the top game
	claim := game.RootClaim(ctx)
	for claim.IsOutputRoot(ctx) && !claim.IsOutputRootLeaf(ctx) {
		if claim.AgreesWithOutputRoot() {
			// If the latest claim agrees with the output root, expect the honest challenger to counter it
			claim = claim.WaitForCounterClaim(ctx)
			game.LogGameData(ctx)
			claim.RequireCorrectOutputRoot(ctx)
		} else {
			// Otherwise we should counter
			claim = claim.Attack(ctx, common.Hash{0xaa})
			game.LogGameData(ctx)
		}
	}

	// Wait for the challenger to post the first claim in the cannon trace
	claim = claim.WaitForCounterClaim(ctx)
	game.LogGameData(ctx)

	// Attack the root of the alphabet trace subgame
	claim = correctTrace.AttackClaim(ctx, claim)
	for !claim.IsMaxDepth(ctx) {
		if claim.AgreesWithOutputRoot() {
			// If the latest claim supports the output root, wait for the honest challenger to respond
			claim = claim.WaitForCounterClaim(ctx)
			game.LogGameData(ctx)
		} else {
			// Otherwise we need to counter the honest claim
			claim = correctTrace.AttackClaim(ctx, claim)
			game.LogGameData(ctx)
		}
	}
	// Challenger should be able to call step and counter the leaf claim.
	claim.WaitForCountered(ctx)
	game.LogGameData(ctx)

	sys.TimeTravelClock.AdvanceTime(game.MaxClockDuration(ctx))
	require.NoError(t, wait.ForNextBlock(ctx, l1Client))
	game.WaitForGameStatus(ctx, types.GameStatusChallengerWon)
	game.LogGameData(ctx)
}
