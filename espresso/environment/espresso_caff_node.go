package environment

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/opnode"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/chaincfg"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
)

const (
	RoleCaffNode = "caff-node"
)

// ErrorFailedToParseSequencerPort is returned when the sequencer port
// cannot be parsed from the espresso dev node.
type ErrorFailedToParseSequencerPort struct {
	Have string
}

// Error implements error
func (e ErrorFailedToParseSequencerPort) Error() string {
	return fmt.Sprintf("failed to parse sequencer port URL: \"%s\"", e.Have)
}

// ErrorFailedToStartCaffNodeGeth is returned when the caff node geth
// instance fails to start.
type ErrorFailedToStartCaffNodeGeth struct {
	Cause error
}

// Error implements error
func (e ErrorFailedToStartCaffNodeGeth) Error() string {
	return fmt.Sprintf("failed to start caff node geth instance: %v", e.Cause)
}

// Unwrap allows for the root cause of the error to be extracted.
func (e ErrorFailedToStartCaffNodeGeth) Unwrap() error {
	return e.Cause
}

// ErrorFailedToStartCaffNodeOpNode is returned when the caff node op
// node instance fails to start.
type ErrorFailedToStartCaffNodeOpNode struct {
	Cause error
}

// Error implements error
func (e ErrorFailedToStartCaffNodeOpNode) Error() string {
	return fmt.Sprintf("failed to start caff node op node instance: %v", e.Cause)
}

// Unwrap allows for the root cause of the error to be extracted.
func (e ErrorFailedToStartCaffNodeOpNode) Unwrap() error {
	return e.Cause
}

// CaffNodeInstance is a wrapper around the caff node geth instance and op node
// instance, for the Caff Node. It is used to interact with the caff node.
type CaffNodeInstance struct {
	OpNode *opnode.Opnode
	Geth   *geth.GethInstance
}

// Close closes the caff node geth instance and op node instance.
func (c *CaffNodeInstance) Close(ctx context.Context) error {
	return errors.Join(c.OpNode.Stop(ctx), c.Geth.Close())
}

// LaunchCaffNode launches a caff node in the given system. It will
// configure the caff node to connect to the given espresso dev node.
func LaunchCaffNode(t *testing.T, system *e2esys.System, espressoDevNode EspressoDevNode) (*CaffNodeInstance, error) {
	sequencerHostAndPort := espressoDevNode.SequencerPort()
	_, sequencerPort, err := net.SplitHostPort(sequencerHostAndPort)
	if have, want := err, error(nil); have != want {
		return nil, ErrorFailedToParseSequencerPort{Have: sequencerHostAndPort}
	}

	u := url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort("localhost", sequencerPort),
		Path:   "/",
	}

	// Let's start the Caff Node now.
	// Configure our caff-node geth instance
	caffNodeGeth, err := geth.InitL2(RoleCaffNode, system.L2GenesisCfg, system.Cfg.JWTFilePath)
	if have, want := err, error(nil); have != want {
		return nil, ErrorFailedToStartCaffNodeGeth{Cause: have}
	}

	// start our caff-node geth instance
	if have, want := caffNodeGeth.Node.Start(), error(nil); have != want {
		return nil, ErrorFailedToStartCaffNodeGeth{Cause: have}
	}

	system.EthInstances[RoleCaffNode] = caffNodeGeth
	system.Cfg.Loggers[RoleCaffNode] = testlog.Logger(t, slog.LevelInfo).New("role", RoleCaffNode)

	// Make a copy

	caffNodeConfig := *system.Cfg.Nodes[e2esys.RoleVerif]
	caffNodeConfig.Rollup = *system.RollupConfig
	caffNodeConfig.Rollup.CaffNodeConfig = rollup.CaffNodeConfig{
		IsCaffNode:                    true,
		PollingHotShotPollingInterval: 30 * time.Millisecond,
		HotShotUrls:                   []string{u.String()},
		L1EthRpc:                      system.L1.UserRPC().RPC(),
		EspressoLightClientAddr:       ESPRESSO_LIGHT_CLIENT_ADDRESS,
	}

	// Configure
	e2esys.ConfigureL1(&caffNodeConfig, system.EthInstances[e2esys.RoleL1], system.L1BeaconEndpoint())
	e2esys.ConfigureL2(&caffNodeConfig, caffNodeGeth, system.Cfg.JWTSecret)

	// Create the Op Node Now
	caffNodeConfig.Rollup.LogDescription(system.Cfg.Loggers[RoleCaffNode], chaincfg.L2ChainIDToNetworkDisplayName)
	l := system.Cfg.Loggers[RoleCaffNode]

	var opNodeError error
	caffNode, err := opnode.NewOpnode(l, &caffNodeConfig, func(e error) {
		opNodeError = e
	})
	if have, want := err, error(nil); have != want {
		// Clean up the Caff Node Geth instance
		caffNodeGeth.Close()
		return nil, ErrorFailedToStartCaffNodeOpNode{Cause: have}
	}

	if have, want := opNodeError, error(nil); have != want {
		caffNodeGeth.Close()
		return nil, ErrorFailedToStartCaffNodeOpNode{Cause: have}
	}

	// Alright, we should have our Caff Node Launched now.

	return &CaffNodeInstance{
		OpNode: caffNode,
		Geth:   caffNodeGeth,
	}, nil
}
