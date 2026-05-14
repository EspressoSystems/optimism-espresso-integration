//go:build !mips64

package rollup

import "github.com/ethereum-optimism/optimism/espresso"

// ToCLIConfig converts the local CaffNodeConfig to espresso.CLIConfig for use
// by the Espresso streamer and other Espresso-specific code paths.
//
// This file is tagged //go:build !mips64 because espresso.CLIConfig is defined
// in espresso/cli.go which depends on the espresso-streamers and
// espresso-network/sdks Go modules — neither of which compiles on mips64
// (the op-program fault-proof target). The constant-only
// BatchAuthLookbackWindowOrDefault helper lives untagged in espresso_config.go.
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
