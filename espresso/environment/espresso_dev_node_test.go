package environment_test

import (
	"context"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
)

// TestEspressoDockerDevNodeSmokeTest is a smoke test for the Espresso Dev Node
// Docker implementation. It starts the dev node and then stops it. And tries
// to ensure that the e2e system, and the docker container stop correctly.
func TestEspressoDockerDevNodeSmokeTest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	{
		// Stop the Docker Container
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		espressoClose := make(chan struct{})

		var err error

		go (func(ch chan struct{}) {
			err = espressoDevNode.Stop()
			close(ch)
		})(espressoClose)

		select {
		case <-ctx.Done():
			t.Errorf("espresso dev node failed to stop in the anticipated time given: %v", ctx.Err())
		case <-espressoClose:
			// Espresso Dev Node stopped in the anticipated time
			if err != nil {
				t.Fatalf("failed to stop espresso dev node: %v", err)
			}
		}

		// One last sanity check to ensure that the container is not still
		// running.

		err = espressoDevNode.Stop()
		if err == nil {
			t.Fatalf("espresso dev node should return an error indicating that it cannot be stopped, as it is not running")
		}

		if _, castOk := err.(env.DockerContainerNotRunningError); !castOk {
			t.Fatalf("espresso dev node should return a DockerContainerNotRunningError, but received: %v", err)
		}
	}

	{
		// Stop the e2e system
		sysClose := make(chan struct{})

		go (func(ch chan struct{}) {
			system.Close()
			close(ch)
		})(sysClose)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		select {
		case <-ctx.Done():
			t.Errorf("system failed to close in the anticipated time given: %v", ctx.Err())

		case <-sysClose:
			// System closed in the anticipated time
		}
	}
}

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
