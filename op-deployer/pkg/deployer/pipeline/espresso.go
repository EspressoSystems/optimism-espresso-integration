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
		lgr.Info("espresso batch inbox contract deployment not needed")
		return nil
	}

	lgr.Info("deploying espresso contracts")
	// read the nitro enclaver verifier address from environment variable, fallback to empty address
	var nitroEnclaveVerifierAddress common.Address
	if envVar := os.Getenv("NITRO_ENCLAVE_VERIFIER_ADDRESS"); envVar != "" {
		nitroEnclaveVerifierAddress = common.HexToAddress(envVar)
		lgr.Info("Using nitro enclave verifier address from NITRO_ENCLAVE_VERIFIER_ADDRESS env var", "address", nitroEnclaveVerifierAddress.Hex())
	} else {
		lgr.Info("NITRO_ENCLAVE_VERIFIER_ADDRESS env var not set, using empty address")
		// this means we should deploy a mock verifier ( should only be used in dev / test environments
		nitroEnclaveVerifierAddress = common.Address{}
	}

	// get enclave hash from environment variable, fallback to zeroed hash
	var enclaveHash [32]byte
	if envVar := os.Getenv("ENCLAVE_HASH"); envVar != "" {
		copy(enclaveHash[:], common.FromHex(envVar))
		lgr.Info("Using enclave hash from ENCLAVE_HASH env var", "hash", common.Bytes2Hex(enclaveHash[:]))
	} else {
		lgr.Info("ENCLAVE_HASH env var not set, using zeroed hash")
	}

	var nvo opcm.DeployAWSNitroVerifierOutput
	nvo, err = opcm.DeployAWSNitroVerifier(env.L1ScriptHost, opcm.DeployAWSNitroVerifierInput{
		EnclaveHash:          enclaveHash,
		NitroEnclaveVerifier: nitroEnclaveVerifierAddress,
	})
	if err != nil {
		return fmt.Errorf("failed to deploy nitro verifier contracts: %w", err)
	}

	var eo opcm.DeployEspressoOutput
	// Read batch authenticator owner address from environment variable, fallback to env.Deployer
	var batchAuthenticatorOwnwerAddress common.Address
	if batchAuthenticatorOwnerEnv := os.Getenv("BATCH_AUTHENTICATOR_OWNER_ADDRESS"); batchAuthenticatorOwnerEnv != "" {
		batchAuthenticatorOwnwerAddress = common.HexToAddress(batchAuthenticatorOwnerEnv)
		lgr.Info("Using batch authenticator owner address from BATCH_AUTHENTICATOR_OWNER_ADDRESS env var", "address", batchAuthenticatorOwnwerAddress.Hex())
	} else {
		batchAuthenticatorOwnwerAddress = env.Deployer
		lgr.Info("Using deployer address from env.Deployer", "address", batchAuthenticatorOwnwerAddress.Hex())
	}

	eo, err = opcm.DeployEspresso(env.L1ScriptHost, opcm.DeployEspressoInput{
		Salt:             st.Create2Salt,
		NitroTEEVerifier: nvo.NitroTEEVerifierAddress,
		NonTeeBatcher:    chainIntent.NonTeeBatcher,
		TeeBatcher:       chainIntent.TeeBatcher,
	}, batchAuthenticatorOwnwerAddress)
	if err != nil {
		return fmt.Errorf("failed to deploy espresso contracts: %w", err)
	}

	chainState.BatchInboxAddress = eo.BatchInboxAddress
	chainState.BatchAuthenticatorAddress = eo.BatchAuthenticatorAddress
	lgr.Info("Espresso batch inbox contract deployed at", "address", eo.BatchInboxAddress)
	return nil
}
