package environment

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
)

// CeloE2eTestError is the error that indicates that the celo
// e2e test "test_weth_bridge" failed to run successfully.
type CeloE2eTestError struct {
	Cause error
}

// Error implements error
func (e CeloE2eTestError) Error() string {
	return fmt.Sprintf("celo e2e test failed with error: %v", e.Cause)
}

// runCeloE2eTest is a helper function that runs the celo e2e test
// script with the given name. It sets up the environment to point
// to the correct L1 and L2 RPC URLs, and runs the script in the
// op-e2e/celo directory.
func runCeloE2eTest(ct *DevNetLauncherContext, script string) func(c *batcher.CLIConfig) {
	return func(c *batcher.CLIConfig) {
		if ct.Error != nil {
			// Don't continue if we already have an error
			return
		}

		ctx, cancel := context.WithCancel(ct.Ctx)
		defer cancel()

		l1EthRpcURL, err := url.Parse(c.L1EthRpc)
		if err != nil {
			ct.Error = FailedToDetermineL1RPCURL{Cause: err}
			return
		}
		l2EthRpcURL, err := url.Parse(c.L2EthRpc)
		if err != nil {
			ct.Error = FailedToDetermineL1RPCURL{Cause: err}
			return
		}

		l1EthRpcURLCopy := new(url.URL)
		*l1EthRpcURLCopy = *l1EthRpcURL
		l1EthRpcURLCopy.Scheme = "http"

		l2EthRpcURLCopy := new(url.URL)
		*l2EthRpcURLCopy = *l2EthRpcURL
		l2EthRpcURLCopy.Scheme = "http"

		projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(strings.TrimSpace(os.Args[0]))))
		celoE2eRoot := filepath.Join(projectRoot, "op-e2e", "celo")

		// Launch the test_weth_bridge.sh script with the environment setup to
		// point to the correct L1 and L2 RPC URLs.
		cmd := exec.CommandContext(
			ctx,
			script,
		)
		cmd.Env = append(cmd.Env, os.Environ()...) // Copy the current environment
		cmd.Env = append(
			cmd.Env,
			"ETH_RPC_URL="+l2EthRpcURLCopy.String(),
			"ETH_RPC_URL_L1="+l1EthRpcURLCopy.String(),
			"FOUNDRY_DISABLE_NIGHTLY_WARNING=1",
		)
		cmd.Dir = celoE2eRoot
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the celo setup script
		if err := cmd.Run(); err != nil {
			ct.Error = CeloE2eTestError{Cause: err}
		}
	}
}

// celoE2eTestWethBridge is a function that runs the celo e2e test
// "test_weth_bridge" before the Batcher is started.
func celoE2eTestWethBridge(ct *DevNetLauncherContext) func(c *batcher.CLIConfig) {
	return runCeloE2eTest(ct, "./test_weth_bridge.sh")
}

// celoE2eTestNPM is a function that runs the celo e2e test
// "test_npm" before the Batcher is started.
func celoE2eTestNPM(ct *DevNetLauncherContext) func(c *batcher.CLIConfig) {
	return runCeloE2eTest(ct, "./test_npm.sh")
}

// WithCeloTestWethBridge is a DevNetLauncherOption that runs the celo
// e2e test "test_weth_bridge" before the Batcher is started.
func WithCeloTestWethBridge() DevNetLauncherOption {
	return func(c *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			StartOptions: []e2esys.StartOption{
				{
					Role:       "test_weth_bridge",
					BatcherMod: celoE2eTestWethBridge(c),
				},
			},
		}
	}
}

// WithCeloTestNPM is a DevNetLauncherOption that runs the celo
// e2e test "test_npm" before the Batcher is started.
func WithCeloTestNPM() DevNetLauncherOption {
	return func(c *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			StartOptions: []e2esys.StartOption{
				{
					Role:       "test_npm",
					BatcherMod: celoE2eTestNPM(c),
				},
			},
		}
	}
}
