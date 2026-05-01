//go:build mips64

package rollup

// defaultBatchAuthLookbackWindowMIPS64 mirrors espresso.DefaultBatchAuthLookbackWindow.
// The value is duplicated here because the espresso package is not buildable under
// mips64 (op-program-client-mips64), which prevents importing it from this package.
// Keep this in sync with espresso.DefaultBatchAuthLookbackWindow.
const defaultBatchAuthLookbackWindowMIPS64 uint64 = 100

// BatchAuthLookbackWindowOrDefault returns the configured lookback window,
// or the default (100) when unset. Mirrors the !mips64 implementation in
// espresso_config.go.
func (cfg *Config) BatchAuthLookbackWindowOrDefault() uint64 {
	if cfg.BatchAuthLookbackWindow == 0 {
		return defaultBatchAuthLookbackWindowMIPS64
	}
	return cfg.BatchAuthLookbackWindow
}
