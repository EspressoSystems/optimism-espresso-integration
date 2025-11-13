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
	var nvo opcm.DeployAWSNitroVerifierOutput
	nvo, err = opcm.DeployAWSNitroVerifier(env.L1ScriptHost, opcm.DeployAWSNitroVerifierInput{
		EnclaveHash: [32]byte{},
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
		Salt:                  st.Create2Salt,
		PreApprovedBatcherKey: chainIntent.PreApprovedBatcherKey,
		NitroTEEVerifier:      nvo.NitroTEEVerifierAddress,
	}, batchAuthenticatorOwnwerAddress)
	if err != nil {
		return fmt.Errorf("failed to deploy espresso contracts: %w", err)
	}

	chainState.BatchInboxAddress = eo.BatchInboxAddress
	chainState.BatchAuthenticatorAddress = eo.BatchAuthenticatorAddress
	lgr.Info("Espresso batch inbox contract deployed at", "address", eo.BatchInboxAddress)
	return nil
}
