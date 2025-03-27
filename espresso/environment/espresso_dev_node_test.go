package environment_test

import (
	"context"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
)

// TestEspressoDockerDevNode tests to ensure that the Espresso Dev Node can be
// launched without error.
func TestEspressoDockerDevNode(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, devNodeInfo, err := launcher.StartDevNet(ctx, t)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	_ = system
	_ = devNodeInfo

	cancel()

	time.Sleep(time.Second)
}
