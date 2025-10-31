package espresso

import (
	"crypto/ecdsa"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
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
	UseFetchApiFlagName              = espressoFlags("fetch-api")
	QueryServiceUrlsFlagName         = espressoFlags("urls")
	LightClientAddrFlagName          = espressoFlags("light-client-addr")
	L1UrlFlagName                    = espressoFlags("l1-url")
	TestingBatcherPrivateKeyFlagName = espressoFlags("testing-batcher-private-key")
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
		&cli.BoolFlag{
			Name:     UseFetchApiFlagName,
			Usage:    "Use fetch API for Espresso queries",
			Value:    false,
			EnvVars:  espressoEnvs(envPrefix, "FETCH_API"),
			Category: category,
		},
		&cli.StringSliceFlag{
			Name:     QueryServiceUrlsFlagName,
			Usage:    "Comma-separated list of Espresso query service URLs",
			EnvVars:  espressoEnvs(envPrefix, "URLS"),
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
	}
}

type CLIConfig struct {
	Enabled                  bool
	PollInterval             time.Duration
	UseFetchAPI              bool
	QueryServiceURLs         []string
	LightClientAddr          common.Address
	L1URL                    string
	TestingBatcherPrivateKey *ecdsa.PrivateKey
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
	}
	return nil
}

func ReadCLIConfig(c *cli.Context) CLIConfig {
	config := CLIConfig{
		Enabled:      c.Bool(EnabledFlagName),
		PollInterval: c.Duration(PollIntervalFlagName),
		UseFetchAPI:  c.Bool(UseFetchApiFlagName),
		L1URL:        c.String(L1UrlFlagName),
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
