package devnet_tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// smokeTestTimeout: devnet bring-up on Mac ARM64 (emulated amd64 images) can exceed 20m; use 45m so Up() can complete.
const smokeTestTimeout = 45 * time.Minute

func TestSmokeWithoutTEE(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), smokeTestTimeout)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())
}

func TestSmokeWithTEE(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), smokeTestTimeout)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())
}
