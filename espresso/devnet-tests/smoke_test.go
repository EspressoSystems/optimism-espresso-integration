package devnet_tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSmokeWithoutTEE(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t, ComposeProfileNonTee)
	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())
}

func TestSmokeWithTEE(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t, ComposeProfileTee)
	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())
}
