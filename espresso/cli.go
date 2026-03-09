package espresso

import (
	"crypto/ecdsa"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network/sdks/go/light-client"
)

// espressoFlags returns the flag names for espresso
func espressoFlags(v string) string {
	return "espresso." + v
}

func espressoEnvs(envprefix, v string) []string {
	return []string{envprefix + "_ESPRESSO_" + v}
}

var (
	EnabledFlagName                  = espressoFlags("enabled")
	PollIntervalFlagName             = espressoFlags("poll-interval")
	QueryServiceUrlsFlagName         = espressoFlags("urls")
	BuilderUrlFlagName               = espressoFlags("builder-url")
	LightClientAddrFlagName          = espressoFlags("light-client-addr")
	L1UrlFlagName                    = espressoFlags("l1-url")
	TestingBatcherPrivateKeyFlagName = espressoFlags("testing-batcher-private-key")
	CaffeinationHeightEspresso       = espressoFlags("origin-height-espresso")
	CaffeinationHeightL2             = espressoFlags("origin-height-l2")
	NamespaceFlagName                = espressoFlags("namespace")
	RollupL1UrlFlagName              = espressoFlags("rollup-l1-url")
	AttestationServiceFlagName       = espressoFlags("espresso-attestation-service")
)

func CLIFlags(envPrefix string, category string) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:     EnabledFlagName,
			Usage:    "Enable Espresso mode",
			Value:    false,
			EnvVars:  espressoEnvs(envPrefix, "ENABLED"),
			Category: category,
		},
		&cli.DurationFlag{
			Name:     PollIntervalFlagName,
			Usage:    "Polling interval for Espresso queries",
			Value:    250 * time.Millisecond,
			EnvVars:  espressoEnvs(envPrefix, "POLL_INTERVAL"),
			Category: category,
		},
		&cli.StringSliceFlag{
			Name:     QueryServiceUrlsFlagName,
			Usage:    "Comma-separated list of Espresso query service URLs",
			EnvVars:  espressoEnvs(envPrefix, "URLS"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     BuilderUrlFlagName,
			Usage:    "Espresso builder URL",
			EnvVars:  espressoEnvs(envPrefix, "BUILDER_URL"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     LightClientAddrFlagName,
			Usage:    "Address of the Espresso light client",
			EnvVars:  espressoEnvs(envPrefix, "LIGHT_CLIENT_ADDR"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     L1UrlFlagName,
			Usage:    "L1 RPC URL Espresso contracts are deployed on",
			EnvVars:  espressoEnvs(envPrefix, "L1_URL"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     TestingBatcherPrivateKeyFlagName,
			Usage:    "Pre-approved batcher ephemeral key (testing only)",
			EnvVars:  espressoEnvs(envPrefix, "TESTING_BATCHER_PRIVATE_KEY"),
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     CaffeinationHeightEspresso,
			Usage:    "Espresso transactions below this height will not be considered",
			EnvVars:  espressoEnvs(envPrefix, "CAFFEINATION_HEIGHT_ESPRESSO"),
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     CaffeinationHeightL2,
			Usage:    "L2 height at which derivation pipeline of Caff node switches to Espresso",
			Value:    0,
			EnvVars:  espressoEnvs(envPrefix, "CAFFEINATION_HEIGHT_L2"),
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     NamespaceFlagName,
			Usage:    "Namespace of Espresso transactions",
			EnvVars:  espressoEnvs(envPrefix, "NAMESPACE"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     RollupL1UrlFlagName,
			Usage:    "RPC URL of L1 backing the Rollup we're streaming for",
			EnvVars:  espressoEnvs(envPrefix, "ROLLUP_L1_URL"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     AttestationServiceFlagName,
			Usage:    "URL of the Espresso attestation service",
			EnvVars:  espressoEnvs(envPrefix, "ESPRESSO_ATTESTATION_SERVICE"),
			Category: category,
		},
	}
}

type CLIConfig struct {
	Enabled                    bool
	PollInterval               time.Duration
	QueryServiceURLs           []string
	BuilderURL                 string
	LightClientAddr            common.Address
	L1URL                      string
	RollupL1URL                string
	TestingBatcherPrivateKey   *ecdsa.PrivateKey
	Namespace                  uint64
	CaffeinationHeightEspresso uint64
	CaffeinationHeightL2       uint64
	EspressoAttestationService string

	// Non directly configurable option
	allowEmptyAttestationService bool `json:"-"`
}

// AllowEmptyAttestationService allows the attestation service URL to be
// empty. This is set explicitly from a public method, and isn't derivable
// from serialization or any other form other than this method.  This allows
// this setting to be configured via the code, but not externally.
func (c *CLIConfig) AllowEmptyAttestationService() {
	c.allowEmptyAttestationService = true
}

func (c CLIConfig) Check() error {
	if c.Enabled {
		// Check required fields when Espresso is enabled
		if len(c.QueryServiceURLs) == 0 {
			return fmt.Errorf("query service URLs are required when Espresso is enabled")
		}
		if c.LightClientAddr == (common.Address{}) {
			return fmt.Errorf("light client address is required when Espresso is enabled")
		}
		if c.L1URL == "" {
			return fmt.Errorf("L1 URL is required when Espresso is enabled")
		}
		if c.RollupL1URL == "" {
			return fmt.Errorf("rollup L1 URL is required when Espresso is enabled")
		}
		if c.Namespace == 0 {
			return fmt.Errorf("namespace is required when Espresso is enabled")
		}
		if !c.allowEmptyAttestationService && c.EspressoAttestationService == "" {
			return fmt.Errorf("attestation service URL is required when Espresso is enabled")
		}
	}
	return nil
}

func ReadCLIConfig(c *cli.Context) CLIConfig {
	config := CLIConfig{
		Enabled:                    c.Bool(EnabledFlagName),
		PollInterval:               c.Duration(PollIntervalFlagName),
		L1URL:                      c.String(L1UrlFlagName),
		RollupL1URL:                c.String(RollupL1UrlFlagName),
		Namespace:                  c.Uint64(NamespaceFlagName),
		CaffeinationHeightEspresso: c.Uint64(CaffeinationHeightEspresso),
		CaffeinationHeightL2:       c.Uint64(CaffeinationHeightL2),
		EspressoAttestationService: c.String(AttestationServiceFlagName),
		BuilderURL:                 c.String(BuilderUrlFlagName),
	}

	config.QueryServiceURLs = c.StringSlice(QueryServiceUrlsFlagName)

	addrStr := c.String(LightClientAddrFlagName)
	config.LightClientAddr = common.HexToAddress(addrStr)

	pkStr := c.String(TestingBatcherPrivateKeyFlagName)
	pkStr = strings.TrimPrefix(pkStr, "0x")
	pk, err := crypto.HexToECDSA(pkStr)
	if err == nil {
		config.TestingBatcherPrivateKey = pk
	}

	return config
}

func BatchStreamerFromCLIConfig[B Batch](
	cfg CLIConfig,
	log log.Logger,
	unmarshalBatch func([]byte) (*B, error),
) (*BatchStreamer[B], error) {
	if !cfg.Enabled {
		return nil, fmt.Errorf("Espresso is not enabled")
	}

	l1Client, err := ethclient.Dial(cfg.L1URL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial L1 RPC at %s: %w", cfg.L1URL, err)
	}

	RollupL1Client, err := ethclient.Dial(cfg.RollupL1URL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial Rollup L1 RPC at %s: %w", cfg.RollupL1URL, err)
	}

	urlZero := cfg.QueryServiceURLs[0]
	espressoClient := espressoClient.NewClient(urlZero)

	espressoLightClient, err := espressoLightClient.NewLightclientCaller(cfg.LightClientAddr, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Espresso light client")
	}

	streamer := NewEspressoStreamer(
		cfg.Namespace,
		NewAdaptL1BlockRefClient(l1Client),
		NewAdaptL1BlockRefClient(RollupL1Client),
		espressoClient,
		espressoLightClient,
		log,
		unmarshalBatch,
		cfg.PollInterval,
		cfg.CaffeinationHeightEspresso,
		cfg.CaffeinationHeightL2,
	)

	return streamer, nil
}
