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
const ERROR_EXPECTED = true
const NO_ERROR_EXPECTED = false

// This dummy Espresso Query Service node returns "hello" to all requests
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

// A small test to ensure the dummy server runs properly
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

// runWithMultiClient spins up the sequencer, L2 verifier and batcher in Espresso mode.
// Moreover, a dummy Espresso Query Service (EQS) is run on port DUMMY_SERVER_PORT.
// The batcher is initialized with M good Espresso urls and N bad ones (using the dummy EQS url)
// @param numGoodUrls M as mentioned in the above description
// @param numBadUrls N as mentioned in the above description
// @param expectedError if set to true, we expect a timeout error as the L2 cannot make progress. Otherwise, we expect no error at all.
func runWithMultiClient(t *testing.T, numGoodUrls int, numBadUrls int, expectedError bool) {

	badServerUrl := fmt.Sprintf("http://%s:%s", DUMMY_SERVER_HOST, DUMMY_SERVER_PORT)
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

// TestEnforceMajorityRule allows to check that the batcher uses the multiclient for fetching information from Espresso and that this multiclient enforces the majority rule.
// This test is designed to evaluate Test 12 as outlined within the Espresso Celo Integration plan.
// Its concrete description is as follows:
// Arrange:
//
//	Running Sequencer, Batcher in Espresso mode and OP node.
//	Set up the batcher with a list of M "good" urls and N "bad" urls
//
// Act:
//
//	Just wait for the batcher to submits batches and the L2 to make progress.
//
// Assert:
//
//	If M>N, the chain should make progress, otherwise it should not.
func TestEnforceMajorityRule(t *testing.T) {

	runWithMultiClient(t, 1, 0, NO_ERROR_EXPECTED)
	runWithMultiClient(t, 2, 1, NO_ERROR_EXPECTED)
	runWithMultiClient(t, 0, 2, ERROR_EXPECTED)
	runWithMultiClient(t, 1, 1, ERROR_EXPECTED)
}
