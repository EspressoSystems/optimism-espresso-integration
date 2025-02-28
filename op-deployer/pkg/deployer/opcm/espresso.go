package opcm

import (
	"fmt"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
	"github.com/ethereum/go-ethereum/common"
)

type DeployNitroVerifierInput struct {
	EnclaveHash [32]byte
}

type DeployNitroVerifierOutput struct {
	NitroTEEVerifierAddress common.Address
}

type DeployEspressoInput struct {
	Salt                  common.Hash
	PreApprovedBatcherKey common.Address
	NitroTeeVerifier      common.Address
}

type DeployEspressoOutput struct {
	BatchVerifierAddress common.Address
	BatchInboxAddress    common.Address
}

type DeployEspressoScript struct {
	Run func(input, output common.Address) error
}

type DeployNitroVerifierScript struct {
	Run func(input, output common.Address) error
}

func DeployNitroVerifier(
	host *script.Host,
	input DeployNitroVerifierInput,
) (DeployNitroVerifierOutput, error) {
	var output DeployNitroVerifierOutput
	inputAddr := host.NewScriptAddress()
	outputAddr := host.NewScriptAddress()

	cleanupInput, err := script.WithPrecompileAtAddress[*DeployNitroVerifierInput](host, inputAddr, &input)
	if err != nil {
		return output, fmt.Errorf("failed to insert DeployNitroVerifierInput precompile: %w", err)
	}
	defer cleanupInput()

	cleanupOutput, err := script.WithPrecompileAtAddress[*DeployNitroVerifierOutput](host, outputAddr, &output,
		script.WithFieldSetter[*DeployNitroVerifierOutput])
	if err != nil {
		return output, fmt.Errorf("failed to insert DeployNitroVerifierOutput precompile: %w", err)
	}
	defer cleanupOutput()

	implContract := "DeployNitroVerifier"
	deployScript, cleanupDeploy, err := script.WithScript[DeployNitroVerifierScript](host, "DeployEspresso.s.sol", implContract)
	if err != nil {
		return output, fmt.Errorf("failed to load %s script: %w", implContract, err)
	}
	defer cleanupDeploy()

	if err := deployScript.Run(inputAddr, outputAddr); err != nil {
		return output, fmt.Errorf("failed to run %s script: %w", implContract, err)
	}

	return output, nil
}

func DeployEspresso(
	host *script.Host,
	input DeployEspressoInput,
) (DeployEspressoOutput, error) {
	var output DeployEspressoOutput
	inputAddr := host.NewScriptAddress()
	outputAddr := host.NewScriptAddress()

	cleanupInput, err := script.WithPrecompileAtAddress[*DeployEspressoInput](host, inputAddr, &input)
	if err != nil {
		return output, fmt.Errorf("failed to insert DeployEspressoInput precompile: %w", err)
	}
	defer cleanupInput()

	cleanupOutput, err := script.WithPrecompileAtAddress[*DeployEspressoOutput](host, outputAddr, &output,
		script.WithFieldSetter[*DeployEspressoOutput])
	if err != nil {
		return output, fmt.Errorf("failed to insert DeployEspressoOutput precompile: %w", err)
	}
	defer cleanupOutput()

	implContract := "DeployEspresso"
	deployScript, cleanupDeploy, err := script.WithScript[DeployEspressoScript](host, "DeployEspresso.s.sol", implContract)
	if err != nil {
		return output, fmt.Errorf("failed to load %s script: %w", implContract, err)
	}
	defer cleanupDeploy()

	if err := deployScript.Run(inputAddr, outputAddr); err != nil {
		return output, fmt.Errorf("failed to run %s script: %w", implContract, err)
	}

	return output, nil
}
