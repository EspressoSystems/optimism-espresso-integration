package pipeline

import (
	"fmt"

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

	var eo opcm.DeployEspressoOutput
	lgr.Info("deploying espresso contracts")
	eo, err = opcm.DeployEspresso(env.L1ScriptHost, opcm.DeployEspressoInput{
		Salt:                  st.Create2Salt,
		PreApprovedBatcherKey: chainIntent.PreApprovedBatcherKey,
	})
	if err != nil {
		return fmt.Errorf("failed to deploy espresso contracts: %w", err)
	}

	chainState.BatchInboxAddress = eo.BatchInboxAddress
	lgr.Info("Espresso batch inbox contract deployed at", "address", eo.BatchInboxAddress)
	return nil
}
