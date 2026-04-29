//go:build !mips64

package rollup

import "github.com/ethereum-optimism/optimism/espresso"

// ToCLIConfig converts the local CaffNodeConfig to espresso.CLIConfig for use
// by the Espresso streamer and other Espresso-specific code paths.
func (c CaffNodeConfig) ToCLIConfig() espresso.CLIConfig {
	return espresso.CLIConfig{
		Enabled:                    c.Enabled,
		PollInterval:               c.PollInterval,
		QueryServiceURLs:           c.QueryServiceURLs,
		LightClientAddr:            c.LightClientAddr,
		BatchAuthenticatorAddr:     c.BatchAuthenticatorAddr,
		L1URL:                      c.L1URL,
		RollupL1URL:                c.RollupL1URL,
		Namespace:                  c.Namespace,
		CaffeinationHeightEspresso: c.CaffeinationHeightEspresso,
		CaffeinationHeightL2:       c.CaffeinationHeightL2,
		EspressoAttestationService: c.EspressoAttestationService,
		VerifyReceiptMaxBlocks:     c.VerifyReceiptMaxBlocks,
		VerifyReceiptSafetyTimeout: c.VerifyReceiptSafetyTimeout,
		VerifyReceiptRetryDelay:    c.VerifyReceiptRetryDelay,
	}
}

// BatchAuthLookbackWindowOrDefault returns the configured lookback window,
// or espresso.DefaultBatchAuthLookbackWindow (100) when unset.
func (cfg *Config) BatchAuthLookbackWindowOrDefault() uint64 {
	if cfg.BatchAuthLookbackWindow == 0 {
		return espresso.DefaultBatchAuthLookbackWindow
	}
	return cfg.BatchAuthLookbackWindow
}

// EspressoOriginBatchPos returns the L2 batch number at which the Espresso
// streamer should start emitting batches. It is derived from the rollup config's
// EspressoEnforcementTime: the streamer must align its origin with the L2 block
// at which Espresso enforcement activates. Returns 0 (i.e. start at genesis L2
// position) if EspressoEnforcementTime is unset.
func (c *Config) EspressoOriginBatchPos() uint64 {
	if c.EspressoEnforcementTime == nil {
		return 0
	}
	n, err := c.TargetBlockNumber(*c.EspressoEnforcementTime)
	if err != nil {
		// Fork timestamp is before genesis; nothing to skip.
		return 0
	}
	return n
}

// CaffNodeConfigFromCLIConfig converts an espresso.CLIConfig to a CaffNodeConfig
// for embedding in rollup.Config.
func CaffNodeConfigFromCLIConfig(c espresso.CLIConfig) CaffNodeConfig {
	return CaffNodeConfig{
		Enabled:                    c.Enabled,
		PollInterval:               c.PollInterval,
		QueryServiceURLs:           c.QueryServiceURLs,
		LightClientAddr:            c.LightClientAddr,
		BatchAuthenticatorAddr:     c.BatchAuthenticatorAddr,
		L1URL:                      c.L1URL,
		RollupL1URL:                c.RollupL1URL,
		Namespace:                  c.Namespace,
		CaffeinationHeightEspresso: c.CaffeinationHeightEspresso,
		CaffeinationHeightL2:       c.CaffeinationHeightL2,
		EspressoAttestationService: c.EspressoAttestationService,
		VerifyReceiptMaxBlocks:     c.VerifyReceiptMaxBlocks,
		VerifyReceiptSafetyTimeout: c.VerifyReceiptSafetyTimeout,
		VerifyReceiptRetryDelay:    c.VerifyReceiptRetryDelay,
	}
}
