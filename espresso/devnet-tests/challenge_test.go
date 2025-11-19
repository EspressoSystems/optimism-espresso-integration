package devnet_tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestChallengeGame(t *testing.T) {
	t.Skip("Temporarily skipped: Re-enable once Succinct Integration is investigated.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Wait for the proposer to make a claim.
	var games []ChallengeGame
	for len(games) == 0 {
		var err error
		t.Logf("waiting for a challenge game")
		time.Sleep(5 * time.Second)
		games, err = d.ListChallengeGames()
		require.NoError(t, err)
	}
	t.Logf("game created: %v", games[0])
	require.Equal(t, uint64(1), games[0].Claims)

	// Attack the first claimed state.
	t.Logf("attacking game")
	require.NoError(t, d.OpChallenger("move", "--attack", "--game-address", games[0].Address.Hex()))

	// Check that the proposer correctly responds.
	CLAIMS_NUMBER := uint64(3) // First claim by the proposer + attack + response
	for {
		updatedGames, err := d.ListChallengeGames()
		require.NoError(t, err)
		if updatedGames[0].Claims == CLAIMS_NUMBER {
			require.Equal(t, updatedGames[0].OutputRoot, games[0].OutputRoot)
			break
		}

		t.Logf("waiting for a response")
		time.Sleep(time.Second)
	}
}
