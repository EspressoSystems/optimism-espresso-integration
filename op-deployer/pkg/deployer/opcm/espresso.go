package opcm

import (
	"fmt"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
	"github.com/ethereum/go-ethereum/common"
)

type DeployEspressoInput struct {
	NitroEnclaveVerifier common.Address
	EspressoBatcher      common.Address
	SystemConfig         common.Address
	ProxyAdminOwner      common.Address
}

type DeployEspressoOutput struct {
	BatchAuthenticatorAddress common.Address
	TeeVerifierProxy          common.Address
	TeeVerifierProxyAdmin     common.Address
	NitroTEEVerifier          common.Address
}

type DeployEspressoScript struct {
	Run func(input, output, deployerAddress common.Address) error
}

func DeployEspresso(
	host *script.Host,
	input DeployEspressoInput,
	deployerAddress common.Address,
) (DeployEspressoOutput, error) {
	var output DeployEspressoOutput
	inputAddr := host.NewScriptAddress()
	outputAddr := host.NewScriptAddress()

	cleanupInput, err := script.WithPrecompileAtAddress(host, inputAddr, &input)
	if err != nil {
		return output, fmt.Errorf("failed to insert DeployEspressoInput precompile: %w", err)
	}
	defer cleanupInput()

	cleanupOutput, err := script.WithPrecompileAtAddress(host, outputAddr, &output,
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

	if err := deployScript.Run(inputAddr, outputAddr, deployerAddress); err != nil {
		return output, fmt.Errorf("failed to run %s script: %w", implContract, err)
	}

	return output, nil
}
