package environment_test

import (
	"context"
	"fmt"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"testing"
	"time"
)

const DUMMY_SERVER_HOST = "localhost"
const DUMMY_SERVER_PORT = "8888"

// This dummy Espresso node return "hello" to all requests
func startServer() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})
	go func() {
		// Optional: handle error from ListenAndServe
		addr := fmt.Sprintf(":%s", DUMMY_SERVER_PORT)
		if err := http.ListenAndServe(addr, handler); err != nil {
			fmt.Println("Server error:", err)
		}
	}()
}

func TestHelloServer(t *testing.T) {
	startServer()
	time.Sleep(100 * time.Millisecond) // Let server start

	url := fmt.Sprintf("http://%s:%s", DUMMY_SERVER_HOST, DUMMY_SERVER_PORT)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}

func launch(t *testing.T, numGoodUrls int, numBadUrls int, badServerUrl string, expectedError bool) {
	startServer()
	time.Sleep(100 * time.Millisecond) // Let server start

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, devNode, err := launcher.StartDevNet(ctx, t, env.SetEspressoUrls(numGoodUrls, numBadUrls, badServerUrl))
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	caffNode, err := env.LaunchDecaffNode(t, system, devNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer env.Stop(t, caffNode)

	caffClient := system.NodeClient(e2esys.RoleVerif)

	// Wait for batcher to start advancing L2 head
	_, err = geth.WaitForBlockToBeSafe(big.NewInt(2), caffClient, 30*time.Second)

	if expectedError {
		require.Error(t, err, "The L2 should not be progressing")
	} else {
		require.NoError(t, err, "The L2 should be progressing")
	}

}

// TestEnforceMajorityRule
func TestEnforceMajorityRule(t *testing.T) {

	badServerUrl := fmt.Sprintf("http://%s:%s", DUMMY_SERVER_HOST, DUMMY_SERVER_PORT)

	launch(t, 1, 0, badServerUrl, false)
	launch(t, 2, 1, badServerUrl, false)
	launch(t, 0, 2, badServerUrl, true)
	launch(t, 1, 1, badServerUrl, true)
}
