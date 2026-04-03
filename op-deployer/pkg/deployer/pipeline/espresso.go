package pipeline

import (
	"fmt"
	"os"

	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/opcm"
	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/state"
	"github.com/ethereum/go-ethereum/common"
)

func DeployEspresso(env *Env, intent *state.Intent, st *state.State, chainID common.Hash) error {
	lgr := env.Logger.New("stage", "deploy-espresso")

	chainIntent, err := intent.Chain(chainID)
	if err != nil {
		return fmt.Errorf("failed to get chain intent: %w", err)
	}

	chainState, err := st.Chain(chainID)
	if err != nil {
		return fmt.Errorf("failed to get chain state: %w", err)
	}

	if !chainIntent.EspressoEnabled {
		lgr.Info("espresso not enabled, skipping BatchAuthenticator deployment")
		return nil
	}

	lgr.Info("deploying espresso contracts")

	// Read the underlying AWS NitroEnclaveVerifier address (from Automata).
	// If not set, address(0) triggers mock deployment — dev/test only.
	var nitroEnclaveVerifierAddress common.Address
	if envVar := os.Getenv("NITRO_ENCLAVE_VERIFIER_ADDRESS"); envVar != "" {
		nitroEnclaveVerifierAddress = common.HexToAddress(envVar)
		lgr.Info("using nitro enclave verifier from NITRO_ENCLAVE_VERIFIER_ADDRESS", "address", nitroEnclaveVerifierAddress.Hex())
	} else {
		lgr.Info("NITRO_ENCLAVE_VERIFIER_ADDRESS not set — deploying mock TEE verifiers")
	}

	var batchAuthOwner common.Address
	if envVar := os.Getenv("BATCH_AUTHENTICATOR_OWNER_ADDRESS"); envVar != "" {
		batchAuthOwner = common.HexToAddress(envVar)
		lgr.Info("using batch authenticator owner from BATCH_AUTHENTICATOR_OWNER_ADDRESS", "address", batchAuthOwner.Hex())
	} else {
		batchAuthOwner = env.Deployer
		lgr.Info("using deployer as batch authenticator owner", "address", batchAuthOwner.Hex())
	}

	eo, err := opcm.DeployEspresso(env.L1ScriptHost, opcm.DeployEspressoInput{
		NitroEnclaveVerifier: nitroEnclaveVerifierAddress,
		TeeBatcher:           chainIntent.TeeBatcher,
		SystemConfig:         chainState.SystemConfigProxy,
		ProxyAdminOwner:      batchAuthOwner,
	}, batchAuthOwner)
	if err != nil {
		return fmt.Errorf("failed to deploy espresso contracts: %w", err)
	}

	chainState.BatchAuthenticatorAddress = eo.BatchAuthenticatorAddress
	lgr.Info("espresso contracts deployed",
		"batchAuthenticator", eo.BatchAuthenticatorAddress,
		"teeVerifier", eo.TeeVerifierProxy,
		"nitroTEEVerifier", eo.NitroTEEVerifier,
	)
	return nil
}
