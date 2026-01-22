package devnet_tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSmokeWithoutTEE(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(DevnetProfileNonTee))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())
}

func TestSmokeWithTEE(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(DevnetProfileTee))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// STOP HERE
	fmt.Println("FLAG TestSmokeWithTEE: Sleep START 10 min")
	time.Sleep(5 * time.Minute)
	fmt.Println("FLAG TestSmokeWithTEE: Sleep FINISHED")

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())
}
