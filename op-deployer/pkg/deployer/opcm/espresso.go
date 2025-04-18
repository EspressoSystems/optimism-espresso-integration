package opcm

import (
	"fmt"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
	"github.com/ethereum/go-ethereum/common"
)

type DeployAWSNitroVerifierInput struct {
	EnclaveHash [32]byte
}

type DeployAWSNitroVerifierOutput struct {
	NitroTEEVerifierAddress common.Address
}

type DeployEspressoInput struct {
	Salt                  common.Hash
	PreApprovedBatcherKey common.Address
	NitroTEEVerifier      common.Address
}

type DeployEspressoOutput struct {
	BatchAuthenticatorAddress common.Address
	BatchInboxAddress         common.Address
}

type DeployEspressoScript struct {
	Run func(input, output common.Address) error
}

type DeployAWSNitroVerifierScript struct {
	Run func(input, output common.Address) error
}

func DeployAWSNitroVerifier(
	host *script.Host,
	input DeployAWSNitroVerifierInput,
) (DeployAWSNitroVerifierOutput, error) {
	var output DeployAWSNitroVerifierOutput
	inputAddr := host.NewScriptAddress()
	outputAddr := host.NewScriptAddress()

	cleanupInput, err := script.WithPrecompileAtAddress[*DeployAWSNitroVerifierInput](host, inputAddr, &input)
	if err != nil {
		return output, fmt.Errorf("failed to insert DeployAWSNitroVerifierInput precompile: %w", err)
	}
	defer cleanupInput()

	cleanupOutput, err := script.WithPrecompileAtAddress[*DeployAWSNitroVerifierOutput](host, outputAddr, &output,
		script.WithFieldSetter[*DeployAWSNitroVerifierOutput])
	if err != nil {
		return output, fmt.Errorf("failed to insert DeployAWSNitroVerifierOutput precompile: %w", err)
	}
	defer cleanupOutput()

	implContract := "DeployAWSNitroVerifier"
	deployScript, cleanupDeploy, err := script.WithScript[DeployAWSNitroVerifierScript](host, "DeployAWSNitroVerifier.s.sol", implContract)
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
