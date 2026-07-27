//go:build !mips64

package rollup

import "github.com/ethereum-optimism/optimism/espresso"

// BatchAuthLookbackWindowOrDefault returns the configured lookback window,
// or espresso.DefaultBatchAuthLookbackWindow (100) when unset.
func (cfg *Config) BatchAuthLookbackWindowOrDefault() uint64 {
	if cfg.BatchAuthLookbackWindow == 0 {
		return espresso.DefaultBatchAuthLookbackWindow
	}
	return cfg.BatchAuthLookbackWindow
}
