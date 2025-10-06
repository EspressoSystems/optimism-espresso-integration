package environment

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
)

// EspressoE2eDevnetLauncher is an interface for launching an E2E devnet with Espresso, and
// configuring it to run in a desired manner.
type EspressoE2eDevnetLauncher interface {

	// StartE2eDevnet will launch the devnet with the provided options. The returned system will be
	// a fully configured e2e system with the configured options.
	StartE2eDevnet(ctx context.Context, t *testing.T, options ...E2eDevnetLauncherOption) (*e2esys.System, EspressoDevNode, error)
}

// E2eDevnetLauncherContext is a struct that contains the context and any errors that may have
// occurred during the launch of the E2E devnet. It also contains the current system instance.
type E2eDevnetLauncherContext struct {
	// The launching Context
	Ctx context.Context

	// Any Current Error
	Error error

	// The Current System configuration
	SystemCfg *e2esys.SystemConfig

	// The Current System instance
	System *e2esys.System

	// EspressoDevNode represents the Espresso Dev Node that is being launched.
	EspressoDevNode
}

// E2eDevnetLauncherOption is a function that takes a E2eDevnetLauncherContext
// and returns an E2eSystemOption.
type E2eDevnetLauncherOption func(
	ctx *E2eDevnetLauncherContext,
) E2eSystemOption

// E2eSystemOption is a struct that contains the options for the
// e2e system that is being launched. It contains the GethOptions and
// any relevant StartOptions that may be needed for the system.
type E2eSystemOption struct {
	SysConfigOption func(*e2esys.SystemConfig)

	// The GethOptions to pass to the Geth Node.
	GethOptions map[string][]geth.GethOption

	// Any relevant StartOptions to pass to the e2e system.
	StartOptions []e2esys.StartOption
}

// EspressoDevNode is an interface that wraps the Espresso Dev Node
// to expose certain functionality, and information that may be needed
// to effectively interact with the Espresso Dev Node.
type EspressoDevNode interface {
	// SequencerPort returns the port that the sequencer is running on.
	SequencerPort() string

	// BuilderPort returns the port that the builder is running on.
	BuilderPort() string

	// EspressoUrls returns the URLs of the Espresso node
	EspressoUrls() []string

	// Shut Down the Espresso Dev Node
	Stop() error
}
