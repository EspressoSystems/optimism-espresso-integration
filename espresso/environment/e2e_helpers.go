package environment

import (
	"math/big"

	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// L2TxWithAmount is a helper.TxOptsFn that sets the Amount of the transaction.
func L2TxWithAmount(amount *big.Int) helpers.TxOptsFn {
	return func(opts *helpers.TxOpts) {
		opts.Value = amount
	}
}

// L2TxWithNonce is a helper.TxOptsFn that sets the Nonce of the transaction.
func L2TxWithNonce(nonce uint64) helpers.TxOptsFn {
	return func(opts *helpers.TxOpts) {
		opts.Nonce = nonce
	}
}

// L2WithToAddress is a helper.TxOptsFn that sets the To address of the
// transaction.
func L2TxWithToAddress(toAddr *common.Address) helpers.TxOptsFn {
	return func(opts *helpers.TxOpts) {
		opts.ToAddr = toAddr
	}
}

// L2TxWithVerifyOnClients is a helper.TxOptsFn that sets the list of
// verification clients to verify the transaction on.
func L2TxWithVerifyOnClients(clients ...*ethclient.Client) helpers.TxOptsFn {
	return func(opts *helpers.TxOpts) {
		opts.VerifyOnClients(clients...)
	}
}

// L2TxWithOptions is a helper.TxOptsFn that sets multiple options for the
// transaction. By default the L2 transaction helper function is only able
// to accept a single helpers.TxOptsFn, so this function allows multiple
// to be passed as a single option, allowing for more granular configuration
// options.
func L2TxWithOptions(options ...helpers.TxOptsFn) helpers.TxOptsFn {
	return func(opts *helpers.TxOpts) {
		for _, option := range options {
			option(opts)
		}
	}
}

// WithSequencerUseFinalized is a DevNetLauncherOption that configures the sequencer's
// `SequencerUseFinalized` option to the provided value.
func WithSequencerUseFinalized(useFinalized bool) DevNetLauncherOption {
	return func(c *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			SysConfigOption: func(cfg *e2esys.SystemConfig) {
				seqConfig := cfg.Nodes[e2esys.RoleSeq]
				seqConfig.Driver.SequencerUseFinalized = useFinalized
			},
		}
	}
}

// WithNonFinalizedProposals is a DevNetLauncherOption that configures the system's
// `NonFinalizedProposals` option to the provided value.
func WithNonFinalizedProposals(useNonFinalized bool) DevNetLauncherOption {
	return func(c *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			SysConfigOption: func(cfg *e2esys.SystemConfig) {
				cfg.NonFinalizedProposals = useNonFinalized
			},
		}
	}
}

// WithL1FinalizedDistance is a DevNetLauncherOption that configures the system's
// `L1FinalizedDistance` option to the provided value.
func WithL1FinalizedDistance(distance uint64) DevNetLauncherOption {
	return func(c *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			SysConfigOption: func(cfg *e2esys.SystemConfig) {
				cfg.L1FinalizedDistance = distance
			},
		}
	}
}
