// Steps to run these tests on a Nitro-enabled EC2 machine:
//
//   - Run `just op-batcher-enclave-image` in kurtosis-devnet/ folder
//     This is just to warm up the docker build cache, otherwise
//     tests may time out building the batcher image from scratch
//
//   - `export ESPRESSO_RUN_ENCLAVE_TESTS=true`
//     Enclave tests are skipped by default
//
//   - `go test ./espresso/enclave-tests/...`
//     Run the tests
package enclave_tests

import (
	"context"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
)

// TestE2eDevNetWithEspressoSimpleTransactions launches the e2e Dev Net with the Espresso Dev Node
// and runs a couple of simple transactions to it.
func TestE2eDevNetWithEspressoSimpleTransactions(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env.RunOnlyWithEnclave(t)

	launcher := new(env.EspressoDevNodeLauncherDocker)
	launcher.EnclaveBatcher = true

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Signal the testnet to shut down on exit
	defer env.Stop(t, espressoDevNode)
	defer env.Stop(t, system)
	// Send Transaction on L1, and wait for verification on the L2 Verifier
	env.RunSimpleL1TransferAndVerifier(ctx, t, system)

	// Submit a Transaction on the L2 Sequencer node, to a Burn Address
	env.RunSimpleL2Burn(ctx, t, system)

}
