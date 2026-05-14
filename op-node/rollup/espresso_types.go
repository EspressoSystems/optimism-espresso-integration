package rollup

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// CaffNodeConfig holds Espresso (Caff Node) configuration embedded in rollup.Config.
// Fields mirror espresso.CLIConfig without importing the espresso package, keeping
// the fault-proof program (op-program) import graph free of Espresso network dependencies.
//
// CaffeinationHeightL2 is an operational parameter (the L2 batch position at which the
// Caff streamer should start emitting batches) and is independent of the
// EspressoEnforcementTime hardfork on the surrounding rollup.Config: the fork timestamp
// gates derivation semantics consensus-wide, while CaffeinationHeightL2 controls where
// a specific Caff node deployment begins streaming. When zero, callers fall back to
// Config.EspressoOriginBatchPos() so that fresh deployments at genesis still work
// without explicit configuration.
type CaffNodeConfig struct {
	Enabled                    bool
	PollInterval               time.Duration
	QueryServiceURLs           []string
	LightClientAddr            common.Address
	BatchAuthenticatorAddr     common.Address
	L1URL                      string
	RollupL1URL                string
	Namespace                  uint64
	CaffeinationHeightEspresso uint64
	CaffeinationHeightL2       uint64
	EspressoAttestationService string
	VerifyReceiptMaxBlocks     uint64
	VerifyReceiptSafetyTimeout time.Duration
	VerifyReceiptRetryDelay    time.Duration
}

// IsEspressoEnforcement returns true if the Espresso enforcement upgrade is active at or past
// the given L2 block timestamp. When active, the derivation pipeline runs all Espresso-specific
// semantics (event-based batch authentication, Caff node HotShot derivation). When inactive, the
// pipeline behaves exactly as upstream Optimism.
func (c *Config) IsEspressoEnforcement(timestamp uint64) bool {
	return c.EspressoEnforcementTime != nil && timestamp >= *c.EspressoEnforcementTime
}
