package devnet_tests

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

// TestChallengeGame verifies that the succinct proposer creates dispute games
// and that games can be queried from the DisputeGameFactory contract.
// The succinct proposer needs finalized L2 blocks before creating games.
func TestChallengeGame(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Verify devnet is running and generate some L2 activity
	require.NoError(t, d.RunSimpleL2Burn())

	// Wait a bit for blocks to be produced and batched
	t.Log("Waiting for blocks to be produced and batched...")
	time.Sleep(10 * time.Second)

	// Wait for the succinct proposer to create a dispute game
	// The proposer creates games when safe L2 head >= anchor + proposal_interval (3 blocks)
	t.Log("Waiting for succinct-proposer to create a dispute game...")
	var games []ChallengeGame
	maxGameWait := 2 * time.Minute
	gameWaitStart := time.Now()

	for len(games) == 0 {
		if time.Since(gameWaitStart) > maxGameWait {
			t.Fatalf("timeout waiting for dispute game to be created (waited %v)", maxGameWait)
		}

		t.Logf("waiting for a challenge game to be created by succinct-proposer...")
		time.Sleep(5 * time.Second)

		var err error
		games, err = d.ListChallengeGames()
		if err != nil {
			t.Logf("error listing games (will retry): %v", err)
		}
	}

	t.Logf("game created: index=%d, address=%s, claims=%d",
		games[0].Index, games[0].Address.Hex(), games[0].Claims)

	// Verify the game has at least 1 claim (the root claim from proposer)
	require.GreaterOrEqual(t, games[0].Claims, uint64(1), "Game should have at least 1 claim")

	// Bind the dispute game contract and log its initial status.
	disputeGame, err := bindings.NewFaultDisputeGame(games[0].Address, d.L1)
	require.NoError(t, err)
	statusRaw, err := disputeGame.Status(&bind.CallOpts{})
	require.NoError(t, err)
	gameStatus, err := types.GameStatusFromUint8(statusRaw)
	require.NoError(t, err)
	t.Logf("dispute game initial status: %s (%d)", gameStatus.String(), statusRaw)
	require.Equal(t, types.GameStatusInProgress, gameStatus, "Dispute game should start InProgress")

	// Observe the dispute game for a limited time to see if it resolves.
	maxObservation := 15 * time.Minute
	pollInterval := 10 * time.Second
	waitStart := time.Now()
	finalStatus := gameStatus
	finalStatusRaw := statusRaw

	t.Logf("Observing dispute game %s for up to %s to see if it resolves...", games[0].Address.Hex(), maxObservation)

	for time.Since(waitStart) < maxObservation {
		statusRaw, err := disputeGame.Status(&bind.CallOpts{})
		require.NoError(t, err)
		status, err := types.GameStatusFromUint8(statusRaw)
		require.NoError(t, err)

		finalStatus = status
		finalStatusRaw = statusRaw

		if status != types.GameStatusInProgress {
			t.Logf("dispute game resolved during observation window: %s (%d)", status.String(), statusRaw)
			require.Equal(t, types.GameStatusDefenderWon, status, "Expected honest proposer/defender to win succinct dispute game")
			break
		}

		time.Sleep(pollInterval)
	}

	t.Logf("dispute game observed final status after %s: %s (%d)", time.Since(waitStart), finalStatus.String(), finalStatusRaw)
	require.Equal(t, finalStatus, types.GameStatusDefenderWon,
		"succinct dispute game final status must be DefenderWon, got %s (%d)",
		finalStatus.String(), finalStatusRaw,
	)

	t.Logf("TestChallengeGame passed: dispute game successfully created by succinct-proposer")
}
