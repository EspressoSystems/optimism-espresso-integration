package rollup

import "github.com/ethereum-optimism/optimism/espresso"

// BatchAuthLookbackWindowOrDefault returns the configured lookback window,
// or espresso.DefaultBatchAuthLookbackWindow (100) when unset.
//
// This file has no build tag so it can be referenced from mips64-reachable
// derivation code (op-node/rollup/derive). It imports only the espresso
// package's constants.go, which is the only file in that package without a
// build tag. The CLI-config conversion helpers (which depend on
// espresso.CLIConfig, an mips64-excluded symbol) live in espresso_cli_config.go.
func (cfg *Config) BatchAuthLookbackWindowOrDefault() uint64 {
	if cfg.BatchAuthLookbackWindow == 0 {
		return espresso.DefaultBatchAuthLookbackWindow
	}
	return cfg.BatchAuthLookbackWindow
}
