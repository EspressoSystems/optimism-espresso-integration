package environment

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	gethNode "github.com/ethereum/go-ethereum/node"
)

// const ESPRESSO_DEV_NODE_DOCKER_IMAGE = "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:main"
// "const ESPRESSO_DEV_NODE_DOCKER_IMAGE = "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-newfoundland"
// "const ESPRESSO_DEV_NODE_DOCKER_IMAGE = "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-labrador"
// const ESPRESSO_DEV_NODE_DOCKER_IMAGE = "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-builder"
// const ESPRESSO_DEV_NODE_DOCKER_IMAGE = "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:20241115"
const ESPRESSO_DEV_NODE_DOCKER_IMAGE = "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-goldendoodle"

// deployed ESPRESSO_SEQUENCER_LIGHT_CLIENT_ADDRESS at 0x17435cce3d1b4fa2e5f8a08ed921d57c6762a180
// deployed ESPRESSO_SEQUENCER_PLONK_VERIFIER_ADDRESS at 0xb4b46bdaa835f8e4b4d8e208b6559cd267851051
const ESPRESSO_LIGHT_CLIENT_ADDRESS = "0x17435cce3d1b4fa2e5f8a08ed921d57c6762a180"

// This is the mnemonic that we use to create the private key for deploying
// contacts on the L1
const ESPRESSO_MNEMONIC = "giant issue aisle success illegal bike spike question tent bar rely arctic volcano long crawl hungry vocal artwork sniff fantasy very lucky have athlete"

// This is the Mnemonic Index that we use to create the private key for deploying
// contracts on the L1
const ESPRESSO_MNEMONIC_INDEX = "0"

// This is address that corresponds to the menmonic we pass to the espresso-dev-node
var ESPRESSO_CONTRACT_ACCOUNT = common.HexToAddress("0x8943545177806ed17b9f23f0a21ee5948ecaa776")

const ESPRESSO_BUILDER_PORT = "31003"
const ESPRESSO_SEQUENCER_API_PORT = "24000"
const ESPRESSO_DEV_NODE_PORT = "24002"

// ErrEspressoBlockHeightDidNotIncrease is a sentinel error that occurs when
// the Espresso Block Height does not increase within the alloted context
// allowance.
var ErrEspressoBlockHeightDidNotIncrease = errors.New("espresso block height did not increase")

// ErrFailedToParseNumber is a sentinel error that occurs when we are unable
// to parse a number from a string
var ErrFailedToParseNumber = errors.New("failed to parse number from string")

// WaitForEspressoBlockHeightToBePositive waits for the Espresso Block Height to
// increase beyond 0.
func WaitForEspressoBlockHeightToBePositive(ctx context.Context, url string) error {
	for {
		select {
		case <-ctx.Done():
			// We've timed out
			return ErrEspressoBlockHeightDidNotIncrease
		default:
		}

		time.Sleep(time.Millisecond * 10)

		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			// Service may not yet be available?
			continue
		}

		if response.StatusCode != http.StatusOK {
			// Service may not yet be available?
			continue
		}

		// Alright, presumably, we have a block height

		buf := new(bytes.Buffer)
		io.Copy(buf, response.Body)
		response.Body.Close()

		blockHeight, ok := new(big.Int).SetString(buf.String(), 10)
		if !ok {
			return ErrFailedToParseNumber
		}

		if blockHeight.Cmp(big.NewInt(0)) > 0 {
			// We have a positive block height! That means we're
			// committing blocks, and we're progressing.  We
			// **SHOULD** be good to continue"
			return nil
		}
	}
}

// EspressoDevNodeLauncherDocker is an implementation of EspressoDevNodeLauncher
// that uses Docker to launch the Espresso Dev Node
type EspressoDevNodeLauncherDocker struct{}

var _ EspressoDevNetLauncher = (*EspressoDevNodeLauncherDocker)(nil)

// FailedToDetermineL1RPCURL represents a class of errors that occur when we
// are unable to correctly form our L1 RPC URL
type FailedToDetermineL1RPCURL struct {
	Cause error
}

// Error implements error
func (f FailedToDetermineL1RPCURL) Error() string {
	return fmt.Sprintf("failed to determine the L1 RPC URL: %v", f.Cause)
}

// FailedToLoadEspressoAccount represents a class of errors that occur when we
// are unable to load the espresso account
type FailedToLoadEspressoAccount struct {
	Cause error
}

// Error implements error
func (f FailedToLoadEspressoAccount) Error() string {
	return fmt.Sprintf("failed to load the espresso account: %v", f.Cause)
}

// FailedToLaunchDockerContainer represents a class of errors that occur when
// we are unable to launch a docker container
type FailedToLaunchDockerContainer struct {
	Cause error
}

// Error implements error
func (f FailedToLaunchDockerContainer) Error() string {
	return fmt.Sprintf("failed to launch docker container: %v", f.Cause)
}

// EspressoNodeFailedToBecomeReady represents a class of errors that indicate
// that the espresso-dev-node failed to become ready.
type EspressoNodeFailedToBecomeReady struct {
	Cause error
}

// Error implements error
func (e EspressoNodeFailedToBecomeReady) Error() string {
	return fmt.Sprintf("espresso node failed to become ready: %v", e.Cause)
}

type EspressoDevNodeContainerInfo struct {
	ContainerInfo DockerContainerInfo
}

var _ EspressoDevNode = (*EspressoDevNodeContainerInfo)(nil)

func (e EspressoDevNodeContainerInfo) getPort(originalPort string) string {
	hosts := e.ContainerInfo.PortMap[originalPort]

	if len(hosts) == 0 {
		return ""
	}

	_, port, err := net.SplitHostPort(hosts[0])
	if err != nil {
		return ""
	}

	return port
}

// SequencerPort implements EspressoDevNode, by returning the relevant
// port for the sequencer API in the Espresso dev node
func (e EspressoDevNodeContainerInfo) SequencerPort() string {
	return e.getPort(ESPRESSO_SEQUENCER_API_PORT)
}

// BuilderPort implements EspressoDevNode, by returning the relevant
// port for the builder API in the Espresso dev node
func (e EspressoDevNodeContainerInfo) BuilderPort() string {
	return e.getPort(ESPRESSO_BUILDER_PORT)
}

// Stop implements EspressoDevNode, and is a convenience method to stop the
// running container.
//
// This is mostly unnecessary as the context that the container was launched
// in will govern the lifecycle of the container automatically, assuming that
// the context is following the recommended context usage patterns.
func (e EspressoDevNodeContainerInfo) Stop() error {
	cli := new(DockerCli)
	return cli.StopContainer(context.Background(), e.ContainerInfo.ContainerID)
}

// ErrUnableToDetermineEspressoDevNodeSequencerHost is a sentinel error that
// indicates that we were unable to determine what the Sequencer API host
// is meant to be.
var ErrUnableToDetermineEspressoDevNodeSequencerHost = errors.New("unable to determine the host for the espresso-dev-node sequencer api")

func (l *EspressoDevNodeLauncherDocker) StartDevNet(ctx context.Context, t *testing.T, options ...DevNetLauncherOption) (*e2esys.System, EspressoDevNode, error) {
	originalCtx := ctx

	sysConfig := e2esys.DefaultSystemConfig(t, e2esys.WithAllocType(config.AllocTypeStandard))
	sysConfig.DeployConfig.DeployCeloContracts = true
	// sysConfig.DeployConfig.DAChallengeWindow = 16
	// sysConfig.DeployConfig.DAResolveWindow = 16
	// sysConfig.DeployConfig.DABondSize = 1000000
	// sysConfig.DeployConfig.DAResolverRefundPercentage = 0
	// sysConfig.DeployConfig.RollupConfig()
	// sysConfig.DeployConfig.L2ChainID = params.CeloBaklavaChainID

	// Ensure that we fund the dev accounts
	sysConfig.DeployConfig.FundDevAccounts = true

	initialOptions := []DevNetLauncherOption{
		allowHostDockerInternalVirtualHost(),
		fundEspressoAccount(),
		launchEspressoDevNodeDocker(),
	}

	launchContext := DevNetLauncherContext{
		Ctx: originalCtx,
	}

	allOptions := append(initialOptions, options...)

	// getOptions := map[string][]geth.GethOption{}
	startOptions := []e2esys.StartOption{}

	for _, opt := range allOptions {
		options := opt(&launchContext)

		if gethOption := options.GethOptions; gethOption != nil {
			for k, v := range gethOption {
				sysConfig.GethOptions[k] = append(sysConfig.GethOptions[k], v...)
			}
		}

		if startOption := options.StartOptions; startOption != nil {
			startOptions = append(startOptions, startOption...)
		}
	}

	// We want to run the espresso-dev-node.  But we need it to be able to
	// access the L1 node.

	system, err := sysConfig.Start(
		t,

		startOptions...,
	)

	if err != nil {
		if system != nil {
			// We don't want the system running in a partial / incomplete
			// state. So we'll tell it to stop here, just in case.
			system.Close()
		}

		return system, nil, err
	}

	// Auto System Cleanup tied to the passed in context.
	{
		// We want to ensure that the lifecycle of the system node is tied to
		// the context we were given, just like the espresso-dev-node.  So if
		// the context is canceled, or otherwise closed, it will automatically
		// clean up the system.
		go (func(ctx context.Context) {
			<-ctx.Done()

			// The system is guaranteed to not be null here.
			system.Close()
		})(originalCtx)
	}

	return system, launchContext.EspressoDevNode, launchContext.Error
}

// EspressoDevNodeDockerContainerInfo is an implementation of
// EspressoDevNode that uses a Docker container to run the Espresso Dev Node
// and provides the relevant port information for the sequencer API and
type EspressoDevNodeDockerContainerInfo DockerContainerInfo

// EspressoDevNodeDockerContainerInfo is an implementation of
// EspressoDevNode.
var _ EspressoDevNode = (*EspressoDevNodeDockerContainerInfo)(nil)

// SequencerPort implements EspressoDevNode
func (e EspressoDevNodeDockerContainerInfo) SequencerPort() string {
	ports := e.PortMap[ESPRESSO_SEQUENCER_API_PORT]
	if len(ports) <= 0 {
		return ""
	}

	return ports[0]
}

// BuilderPort implements EspressoDevNode
func (e EspressoDevNodeDockerContainerInfo) BuilderPort() string {
	ports := e.PortMap[ESPRESSO_BUILDER_PORT]
	if len(ports) <= 0 {
		return ""
	}

	return ports[0]
}

// ContainerID implements EspressoDevNode
func (e EspressoDevNodeDockerContainerInfo) Stop() error {
	cli := new(DockerCli)
	return cli.StopContainer(context.Background(), e.ContainerID)
}

// allowHostDockerInternalVirtualHost is a convenience method that configures
// Geth instance to allow communication from a virtual host of
// "host.docker.internal".
//
// host.docker.internal is a special DNS name that allows docker containers
// to speak to ports hosted on the host node.
func allowHostDockerInternalVirtualHost() DevNetLauncherOption {
	return func(c *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			GethOptions: map[string][]geth.GethOption{
				e2esys.RoleL1: {
					func(thCfg *ethconfig.Config, nodeCfg *gethNode.Config) error {
						// We append the host machine address to the list of virtual hosts, so
						// that we do not get denied when attempting to access the host machine's
						// RPC API.
						nodeCfg.HTTPVirtualHosts = append(nodeCfg.HTTPVirtualHosts, "host.docker.internal", "localhost")

						return nil
					},
				},
			},
		}
	}
}

// fundEspressoAccount is a convenience method that funds the espresso
// account with an initial amount of ETH, so that it can deploy contracts
// on the L1.  This is necessary as the espresso-dev-node does not
func fundEspressoAccount() DevNetLauncherOption {
	return func(c *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			StartOptions: []e2esys.StartOption{
				{
					Key:  "afterRollupNodeStart",
					Role: e2esys.RoleVerif,
					Action: func(sysConfig *e2esys.SystemConfig, sys *e2esys.System) {
						if c.Error != nil {
							// Early Return if we already have an Error set
							return
						}

						c.System = sys

						ctx, cancel := context.WithCancel(c.Ctx)
						defer cancel()

						// Fund the Espresso Account, so it is able to deploy contracts
						l1Client := sys.NodeClient(e2esys.RoleL1)

						tx, err := SignTransaction(&types.DynamicFeeTx{
							ChainID:   sysConfig.L1ChainIDBig(),
							To:        &ESPRESSO_CONTRACT_ACCOUNT,
							Value:     big.NewInt(1_000_000_000_000_000_000),
							GasTipCap: big.NewInt(1_000_000_000),
							GasFeeCap: big.NewInt(1_000_000_000),
							Gas:       25000,
							Data:      nil,
						}, sysConfig.Secrets.Alice, sysConfig.L1ChainIDBig())
						if err != nil {
							c.Error = FailedToLoadEspressoAccount{Cause: err}
							return
						}

						startingBalance, err := l1Client.BalanceAt(ctx, ESPRESSO_CONTRACT_ACCOUNT, nil)
						if err != nil {
							c.Error = FailedToLoadEspressoAccount{Cause: err}
							return
						}

						err = l1Client.SendTransaction(ctx, tx)
						if err != nil {
							c.Error = FailedToLoadEspressoAccount{Cause: err}
							return
						}

						{
							ctx, cancel := context.WithTimeout(ctx, time.Second*30)
							defer cancel()
							if err := WaitForIncreasedBalance(ctx, l1Client, ESPRESSO_CONTRACT_ACCOUNT, startingBalance); err != nil {
								c.Error = FailedToLoadEspressoAccount{Cause: err}
								return
							}
						}
					},
				},
			},
		}
	}
}

// launchEspressoDevNodeDocker is DevNetLauncherOption that launches th
// Espresso Dev Node within a Docker container.  It also ensures that the
// Espresso Dev Node is actively producing blocks before returning.
func launchEspressoDevNodeDocker() DevNetLauncherOption {
	return func(ct *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			StartOptions: []e2esys.StartOption{
				{
					Role: "launch-espresso-dev-node",
					BatcherMod: func(c *batcher.CLIConfig) {
						if ct.Error != nil {
							// Early Return if we already have an Error set
							return
						}

						l1EthRpcURL, err := url.Parse(c.L1EthRpc)
						if err != nil {
							ct.Error = FailedToDetermineL1RPCURL{Cause: err}
							return
						}

						// Let's spin up the espresso-dev-node
						{

							// We need to know the port, so we can configure docker to
							// communicate with the L1 RPC node running on the host machine.
							_, port, err := net.SplitHostPort(l1EthRpcURL.Host)
							if err != nil {
								ct.Error = FailedToDetermineL1RPCURL{Cause: err}
								return
							}

							// We replace the host with host.docker.internal to inform
							// docker to communicate with the host system.

							if isRunningOnLinux {
								l1EthRpcURL.Host = net.JoinHostPort("localhost", port)
							} else {
								l1EthRpcURL.Host = net.JoinHostPort("host.docker.internal", port)
							}

							l1EthRpcURL.Scheme = "http"

							containerCli := new(DockerCli)

							dockerConfig := DockerContainerConfig{
								Image: ESPRESSO_DEV_NODE_DOCKER_IMAGE,
								Environment: map[string]string{
									"ESPRESSO_DEPLOYER_ACCOUNT_INDEX":             ESPRESSO_MNEMONIC_INDEX,
									"ESPRESSO_SEQUENCER_ETH_MNEMONIC":             ESPRESSO_MNEMONIC,
									"ESPRESSO_SEQUENCER_L1_PROVIDER":              l1EthRpcURL.String(),
									"ESPRESSO_SEQUENCER_DATABASE_MAX_CONNECTIONS": "25",
									"ESPRESSO_SEQUENCER_STORAGE_PATH":             "/data/espresso",
									"RUST_LOG":                                    "info",

									"ESPRESSO_BUILDER_PORT":       ESPRESSO_BUILDER_PORT,
									"ESPRESSO_SEQUENCER_API_PORT": ESPRESSO_SEQUENCER_API_PORT,
									"ESPRESSO_DEV_NODE_PORT":      ESPRESSO_DEV_NODE_PORT,
								},
								Ports: []string{
									ESPRESSO_BUILDER_PORT,
									ESPRESSO_SEQUENCER_API_PORT,
									ESPRESSO_DEV_NODE_PORT,
								},
							}

							if isRunningOnLinux {
								// We launch in network mode host on linux,
								// otherwise the container is not able to
								// communicate with the host system.
								// We use host.docker.internal to do this on
								// platforms that are not running natively on
								// linux, as this special address achieves the
								// same result.  But on linux, this does not
								// work, and we need to run on the host instead.
								dockerConfig.Network = "host"
							}
							espressoDevNodeContainerInfo, err := containerCli.LaunchContainer(ct.Ctx, dockerConfig)

							if err != nil {
								ct.Error = FailedToLaunchDockerContainer{Cause: err}
								return
							}

							ct.EspressoDevNode = EspressoDevNodeDockerContainerInfo(espressoDevNodeContainerInfo)

							// We have all of our ports.
							// Let's return all of the relevant port mapping information
							// for easy reference, and cancellation

							hosts := espressoDevNodeContainerInfo.PortMap[ESPRESSO_SEQUENCER_API_PORT]

							if len(hosts) == 0 {
								ct.Error = ErrUnableToDetermineEspressoDevNodeSequencerHost
								return
							}

							// We may have more than a single host, but we'll make do.

							host, port, err := net.SplitHostPort(hosts[0])
							if err != nil {
								ct.Error = ErrUnableToDetermineEspressoDevNodeSequencerHost
								return
							}

							var hostPort string
							switch host {
							case "0.0.0.0":
								// IPv4
								hostPort = net.JoinHostPort("localhost", port)
							case "[::]":
								// IPv6
								hostPort = net.JoinHostPort("localhost", port)
							default:
								hostPort = net.JoinHostPort(host, port)
							}

							currentBlockHeightURLString := "http://" + hostPort + "/status/block-height"

							// Wait for Espresso to be ready
							timeoutCtx, cancel := context.WithTimeout(ct.Ctx, 3*time.Minute)
							defer cancel()
							if err := WaitForEspressoBlockHeightToBePositive(timeoutCtx, currentBlockHeightURLString); err != nil {
								ct.Error = EspressoNodeFailedToBecomeReady{Cause: err}
								return
							}

							c.EspressoUrl = "http://" + hostPort
						}
					},
				},
			},
		}
	}
}
