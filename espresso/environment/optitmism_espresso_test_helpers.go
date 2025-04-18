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
	"strconv"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	gethNode "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

const ESPRESSO_DEV_NODE_DOCKER_IMAGE = "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:20250412-dev-node-pos-preview"

const ESPRESSO_LIGHT_CLIENT_ADDRESS = "0x703848f4c85f18e3acd8196c8ec91eb0b7bd0797"

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

// getPort is a helper function that takes the original port and returns
// the remapped port that the container is listening on.
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

	// Ensure that we fund the dev accounts
	sysConfig.DeployConfig.FundDevAccounts = true

	// Pre-fund Espresso acount with 1M Ether
	espressoPremine := new(big.Int).Mul(new(big.Int).SetUint64(1_000_000), new(big.Int).SetUint64(params.Ether))
	sysConfig.Premine[ESPRESSO_CONTRACT_ACCOUNT] = espressoPremine

	initialOptions := []DevNetLauncherOption{
		allowHostDockerInternalVirtualHost(),
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

// This code is adapted from a gist file:
// https://gist.github.com/sevkin/96bdae9274465b2d09191384f86ef39d
func determineFreePort() (port int, err error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer func() {
		err = listener.Close()
	}()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

// getContainerRemappedHostPort is a helper function that takes the
// containerListeningHostPort and returns the remapped host port
// that the container is listening on.
//
// By default the mapped hosts and ports are in the form of
// - 0.0.0.0:<port> for IPv4
// - [::]:<port> for IPv6
//
// So this function will replace the host with "localhost" to allow
// for communication with the host system.
func getContainerRemappedHostPort(containerListeningHostPort string) (string, error) {
	_, port, err := net.SplitHostPort(containerListeningHostPort)
	if err != nil {
		return "", ErrUnableToDetermineEspressoDevNodeSequencerHost
	}

	hostPort := net.JoinHostPort("localhost", port)

	return hostPort, nil
}

// waitForEspressoToFinishSpinningUp is a helper function that waits for the
// espresso dev node to finish spinning up.
// It checks the portMap of the DockerContainerInfo to retrieve the
// Espresso Dev Node Sequencer API port, and then waits for the block height
// to be greater than 0.
func waitForEspressoToFinishSpinningUp(ct *DevNetLauncherContext, espressoDevNodeContainerInfo DockerContainerInfo) error {
	// We have all of our ports.
	// Let's return all of the relevant port mapping information
	// for easy reference, and cancellation

	hosts := espressoDevNodeContainerInfo.PortMap[ESPRESSO_SEQUENCER_API_PORT]

	if len(hosts) == 0 {
		return ErrUnableToDetermineEspressoDevNodeSequencerHost
	}

	// We may have more than a single host, but we'll make do.
	hostPort, err := getContainerRemappedHostPort(hosts[0])
	if err != nil {
		return err
	}

	currentBlockHeightURLString := "http://" + hostPort + "/status/block-height"

	// Wait for Espresso to be ready
	timeoutCtx, cancel := context.WithTimeout(ct.Ctx, 3*time.Minute)
	defer cancel()
	return WaitForEspressoBlockHeightToBePositive(timeoutCtx, currentBlockHeightURLString)
}

// translateContainerToNodeURL is a helper function that translates the the
// given URL to be used by a container to a form that can be communicated with
// the host system.
//
//	Note:
//		if the network passed in is determined to be "host" we will assume that
//		the host machine can be accessed via "localhost".
//
//	Note:
//
//		The default way we assume this will work is with the Docker for X
//		platform, in which the reserved "host.docker.internal" domain name
//		will allow communication with the host system.  This does **NOT**
//		work on a native Linux platform.
func translateContainerToNodeURL(parsedURL url.URL, network string) (url.URL, error) {
	// We need to know the port, so we can configure docker to
	// communicate with the L1 RPC node running on the host machine.
	_, port, err := net.SplitHostPort(parsedURL.Host)
	if err != nil {
		return url.URL{}, FailedToDetermineL1RPCURL{Cause: err}
	}

	// We replace the host with host.docker.internal to inform
	// docker to communicate with the host system.
	if network == "host" {
		parsedURL.Host = net.JoinHostPort("localhost", port)
	} else {
		parsedURL.Host = net.JoinHostPort("host.docker.internal", port)
	}

	return parsedURL, nil
}

// determineEspressoDevNodeDockerContainerConfig will return an initial
// configuration for the docker cli command to launch the espresso-dev-node.
// It will also return a port mapping that will contain any remapped ports,
// should they be necessary.
func determineEspressoDevNodeDockerContainerConfig(l1EthRpcURL url.URL, network string) (containerConfig DockerContainerConfig, portMapping map[string]string, err error) {
	// These are the expected initial mappings for the ports.  This will
	// be fine when running in an isolated container, and these ports cannot
	// possibly overlap.
	portRemapping := map[string]string{
		ESPRESSO_BUILDER_PORT:       ESPRESSO_BUILDER_PORT,
		ESPRESSO_SEQUENCER_API_PORT: ESPRESSO_SEQUENCER_API_PORT,
		ESPRESSO_DEV_NODE_PORT:      ESPRESSO_DEV_NODE_PORT,
	}

	if network == "host" {
		// If we're running in host mode, we will can potentially have overlapping
		// port definitions, as we spin up nodes in parallel.
		// So we need to determine the free ports on the host system
		// to bind the espresso-dev-node to.
		for portKey := range portRemapping {
			// We need to determine a free port on the host system
			// to bind the espresso-dev-node to.
			freePort, err := determineFreePort()
			if err != nil {
				return DockerContainerConfig{}, nil, FailedToDetermineL1RPCURL{Cause: err}
			}
			portRemapping[portKey] = strconv.FormatInt(int64(freePort), 10)
		}
	}

	l1EthRpcURL.Scheme = "http"

	dockerConfig := DockerContainerConfig{
		Image:   ESPRESSO_DEV_NODE_DOCKER_IMAGE,
		Network: network,
		Environment: map[string]string{
			"ESPRESSO_DEPLOYER_ACCOUNT_INDEX":             ESPRESSO_MNEMONIC_INDEX,
			"ESPRESSO_SEQUENCER_ETH_MNEMONIC":             ESPRESSO_MNEMONIC,
			"ESPRESSO_SEQUENCER_L1_PROVIDER":              l1EthRpcURL.String(),
			"ESPRESSO_SEQUENCER_L1_POLLING_INTERVAL":      "30ms",
			"ESPRESSO_SEQUENCER_DATABASE_MAX_CONNECTIONS": "25",
			"ESPRESSO_SEQUENCER_STORAGE_PATH":             "/data/espresso",
			"RUST_LOG":                                    "info",

			"ESPRESSO_BUILDER_PORT":       portRemapping[ESPRESSO_BUILDER_PORT],
			"ESPRESSO_SEQUENCER_API_PORT": portRemapping[ESPRESSO_SEQUENCER_API_PORT],
			"ESPRESSO_DEV_NODE_PORT":      portRemapping[ESPRESSO_DEV_NODE_PORT],
		},
		Ports: []string{
			portRemapping[ESPRESSO_BUILDER_PORT],
			portRemapping[ESPRESSO_SEQUENCER_API_PORT],
			portRemapping[ESPRESSO_DEV_NODE_PORT],
		},
	}

	return dockerConfig, portRemapping, nil
}

// determineDockerNetworkMode is a helper function that determines the
// docker network mode to use for the container.
//
// We launch in network mode host on linux, otherwise the container is not able
// to communicate with the host system. We use host.docker.internal to do this
// on platforms that are not running natively on linux, as this special address
// achieves the same result.  But on linux, this does not work, and we need to
// run on the host instead.
func determineDockerNetworkMode() string {
	if isRunningOnLinux {
		return "host"
	}

	return ""
}

// ensureHardCodedPortsAreMappedFromTheirOriginalValues is a convenience
// function that makes sure that hard coded ports are associated with their
// remapped port values.  This is done for convenience in order to ensure that
// we can still reference the hard coded ports, even if they've been remapped
// from their original values.
func ensureHardCodedPortsAreMappedFromTheirOriginalValues(containerInfo *DockerContainerInfo, portRemapping map[string]string) {
	if _, ok := portRemapping[ESPRESSO_SEQUENCER_API_PORT]; !ok {
		// If we don't have the original port mapping for the hard
		// coded port, we will need to back fill them in, just
		// to make life easier for consumers.

		for portKey, portValue := range portRemapping {
			// We copy the port mapping information
			// so we know the original mapping again,
			// since we're hard-coding the ports to use.
			// This should allow us to run multiple
			// e2e test environments in parallel on
			// linux as well.
			containerInfo.PortMap[portKey] = containerInfo.PortMap[portValue]
		}
	}
}

// launchEspressoDevNodeDocker is DevNetLauncherOption that launches th
// Espresso Dev Node within a Docker container.  It also ensures that the
// Espresso Dev Node is actively producing blocks before returning.
func launchEspressoDevNodeStartOption(ct *DevNetLauncherContext) e2esys.StartOption {
	return e2esys.StartOption{
		Role: "launch-espresso-dev-node",
		BatcherMod: func(c *batcher.CLIConfig) {
			if ct.Error != nil {
				// Early Return if we already have an Error set
				return
			}

			l1EthRpcURLPtr, err := url.Parse(c.L1EthRpc)
			if err != nil {
				ct.Error = FailedToDetermineL1RPCURL{Cause: err}
				return
			}

			network := determineDockerNetworkMode()

			// Let's spin up the espresso-dev-node
			l1EthRpcURL, err := translateContainerToNodeURL(*l1EthRpcURLPtr, network)
			if err != nil {
				ct.Error = err
				return
			}

			dockerConfig, portRemapping, err := determineEspressoDevNodeDockerContainerConfig(l1EthRpcURL, network)
			if err != nil {
				ct.Error = err
				return
			}

			containerCli := new(DockerCli)

			espressoDevNodeContainerInfo, err := containerCli.LaunchContainer(ct.Ctx, dockerConfig)

			if err != nil {
				ct.Error = FailedToLaunchDockerContainer{Cause: err}
				return
			}

			ct.EspressoDevNode = EspressoDevNodeDockerContainerInfo(espressoDevNodeContainerInfo)
			ensureHardCodedPortsAreMappedFromTheirOriginalValues(&espressoDevNodeContainerInfo, portRemapping)

			if err := waitForEspressoToFinishSpinningUp(ct, espressoDevNodeContainerInfo); err != nil {
				ct.Error = err
				return
			}

			// This skip on error check **SHOULD** be safe as this was
			// already performed inside the `waitForEspressoToFinishSpinningUp`
			// call.
			hostPort, _ := getContainerRemappedHostPort(espressoDevNodeContainerInfo.PortMap[ESPRESSO_SEQUENCER_API_PORT][0])

			c.EspressoUrl = "http://" + hostPort
		},
	}

}

// launchEspressoDevNodeDocker is DevNetLauncherOption that launches th
// Espresso Dev Node within a Docker container.  It also ensures that the
// Espresso Dev Node is actively producing blocks before returning.
func launchEspressoDevNodeDocker() DevNetLauncherOption {
	return func(ct *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			StartOptions: []e2esys.StartOption{
				launchEspressoDevNodeStartOption(ct),
			},
		}
	}
}
