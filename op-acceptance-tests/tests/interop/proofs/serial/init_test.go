package serial

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-devstack/presets"
	"github.com/ethereum-optimism/optimism/op-devstack/stack"
	"github.com/ethereum-optimism/optimism/op-devstack/sysgo"
)

func TestMain(m *testing.M) {
	presets.DoMain(m,
		presets.WithSuperInteropSupernode(),
		stack.MakeCommon(sysgo.WithChallengerCannonKonaEnabled()),
		// celo: skip kona-host tests until Rust RollupConfig supports cel2_time
		presets.WithCompatibleTypes("non-existent-type"),
	)
}
