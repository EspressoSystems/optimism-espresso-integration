package environment_test

import (
	"context"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
)

// TestCeloWethBridge is a test that runs the celo e2e test script
// "test_weth_bridge" against our local e2e devnet environment.
func TestCeloWethBridge(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	_, _, err := launcher.StartDevNet(ctx, t, env.WithCeloTestWethBridge())
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}
}
