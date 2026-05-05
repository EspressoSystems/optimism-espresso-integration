package batcher

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network/sdks/go/light-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hf/nitrite"

	"github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum-optimism/optimism/op-batcher/enclave"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	opcrypto "github.com/ethereum-optimism/optimism/op-service/crypto"
)

// EspressoBatcherConfig groups all Espresso-specific configuration the
// batcher consumes at runtime. It is embedded as a single field on
// BatcherConfig (see service.go) to keep the upstream Optimism
// BatcherConfig field block compact and minimize cherry-pick churn.
//
// Fields are populated from CLIConfig.Espresso (and a few RollupConfig
// fallbacks) by initEspresso below; ephemeral key material is generated
// by initKeyPair (TEE) or copied from CLIConfig.Espresso.TestingBatcherPrivateKey
// (devnet/test).
type EspressoBatcherConfig struct {
	Enabled                    bool
	PollInterval               time.Duration
	AttestationService         string
	CaffeinationHeightEspresso uint64
	// CaffeinationHeightL2 is the L2 batch position at which the Espresso
	// streamer should start emitting batches. Operational parameter for
	// restarting batchers mid-chain (e.g. after a fallback batcher event).
	// When zero, the driver falls back to
	// RollupConfig.EspressoOriginBatchPos().
	CaffeinationHeightL2 uint64

	// Receipt verification tuning for the Espresso transaction submitter.
	VerifyReceiptMaxBlocks     uint64
	VerifyReceiptSafetyTimeout time.Duration
	VerifyReceiptRetryDelay    time.Duration

	// FallbackAuthLeadTime is consulted only by the fallback (non-TEE)
	// batcher's EspressoEnforcement gate. It advances the fallback batcher's
	// switch to authenticated batches relative to the on-chain
	// EspressoEnforcementTime, absorbing the worst-case L1 inclusion delay
	// between batcher decision time (L1 tip) and verifier evaluation time
	// (containing L1 block). See isFallbackAuthRequired for details.
	FallbackAuthLeadTime time.Duration

	// BatcherPublicKey/BatcherPrivateKey is the batcher's identity for the
	// Espresso authentication path. In TEE deployments the private key is
	// generated inside the enclave (initKeyPair) and the public key is
	// attested to via Nitro Enclave PCR0; outside TEE (devnet/test), the
	// configured TestingBatcherPrivateKey overrides them in initEspresso.
	BatcherPublicKey  *ecdsa.PublicKey
	BatcherPrivateKey *ecdsa.PrivateKey
}

// EspressoStreamer returns the Espresso batch streamer driven by this batcher.
func (bs *BatcherService) EspressoStreamer() espresso.EspressoStreamer[derive.EspressoBatch] {
	return bs.driver.espressoStreamer
}

// initChainSigner asserts that the configured TxManager implements the
// ChainSigner interface and stores the embedded ChainSigner on the service.
// Espresso uses ChainSigner to sign batch authentication payloads sent to the
// BatchAuthenticator contract; the cast is required by every Espresso path.
func (bs *BatcherService) initChainSigner() error {
	cast, castOk := bs.TxManager.(opcrypto.ChainSigner)
	if !castOk {
		return fmt.Errorf("tx manager does not implement ChainSigner")
	}
	bs.ChainSigner = cast
	return nil
}

// applyEspressoDriverSetup writes the Espresso-specific fields onto a
// DriverSetup populated with upstream-Optimism fields. Kept separate from the
// main initDriver struct literal so that the upstream block stays in upstream
// shape — minimizing cherry-pick churn when upstream renames or reorders
// fields.
func (bs *BatcherService) applyEspressoDriverSetup(ds *DriverSetup) {
	ds.SequencerAddress = bs.TxManager.From()
	ds.ChainSigner = bs.ChainSigner
	ds.Espresso = bs.EspressoClient
	ds.EspressoLightClient = bs.EspressoLightClient
	ds.Attestation = bs.Attestation
}

// initKeyPair generates an ephemeral ECDSA key pair for the batcher's
// Espresso authentication path. In TEE deployments this key is attested
// to via Nitro Enclave PCR0; outside TEE (devnet/test), the configured
// TestingBatcherPrivateKey overrides this key in initEspresso.
func (bs *BatcherService) initKeyPair() error {
	key, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("failed to generate key pair for batcher: %w", err)
	}
	bs.Espresso.BatcherPrivateKey = key
	bs.Espresso.BatcherPublicKey = &key.PublicKey
	return nil
}

// initEspresso configures the Espresso (Caff Node) integration on the
// BatcherService. Always propagates FallbackAuthLeadTime (consulted by
// the fallback non-TEE batcher even when --espresso.enabled=false), and
// when enabled, wires up the Espresso query-service client, light client,
// ephemeral key pair, and Nitro Enclave attestation (if running in TEE).
func (bs *BatcherService) initEspresso(cfg *CLIConfig) error {
	// FallbackAuthLeadTime is consulted by the fallback (non-TEE) batcher and
	// must be propagated regardless of whether --espresso.enabled is set —
	// the fallback batcher runs with Enabled=false but still needs this knob
	// when a BatchAuthenticator is configured on the rollup.
	bs.Espresso.FallbackAuthLeadTime = cfg.Espresso.FallbackAuthLeadTime

	if !cfg.Espresso.Enabled {
		return nil
	}

	if cfg.Espresso.RollupL1URL == "" {
		cfg.Espresso.RollupL1URL = cfg.L1EthRpc
	}

	if cfg.Espresso.RollupL1URL != cfg.L1EthRpc {
		log.Warn("Espresso Rollup L1 URL differs from batcher's L1EthRpc")
	}

	if cfg.Espresso.L1URL == "" {
		log.Warn("Espresso L1 URL not provided, using batcher's L1EthRpc")
		cfg.Espresso.L1URL = cfg.L1EthRpc
	}
	if cfg.Espresso.Namespace == 0 {
		log.Info("Using L2 chain ID as namespace by default")
		cfg.Espresso.Namespace = bs.RollupConfig.L2ChainID.Uint64()
	}
	if cfg.Espresso.BatchAuthenticatorAddr == (common.Address{}) {
		cfg.Espresso.BatchAuthenticatorAddr = bs.RollupConfig.BatchAuthenticatorAddress
	}

	if err := cfg.Espresso.Check(); err != nil {
		return fmt.Errorf("invalid Espresso config: %w", err)
	}

	bs.Espresso.Enabled = true
	bs.Espresso.PollInterval = cfg.Espresso.PollInterval
	bs.Espresso.AttestationService = cfg.Espresso.EspressoAttestationService
	bs.Espresso.CaffeinationHeightEspresso = cfg.Espresso.CaffeinationHeightEspresso
	bs.Espresso.CaffeinationHeightL2 = cfg.Espresso.CaffeinationHeightL2
	bs.Espresso.VerifyReceiptMaxBlocks = cfg.Espresso.VerifyReceiptMaxBlocks
	bs.Espresso.VerifyReceiptSafetyTimeout = cfg.Espresso.VerifyReceiptSafetyTimeout
	bs.Espresso.VerifyReceiptRetryDelay = cfg.Espresso.VerifyReceiptRetryDelay

	client, err := espressoClient.NewMultipleNodesClient(cfg.Espresso.QueryServiceURLs)
	if err != nil {
		return fmt.Errorf("failed to create Espresso client: %w", err)
	}
	bs.EspressoClient = client

	lightClient, err := espressoLightClient.NewLightclientCaller(cfg.Espresso.LightClientAddr, bs.L1Client)
	if err != nil {
		return fmt.Errorf("failed to create Espresso light client: %w", err)
	}
	bs.EspressoLightClient = lightClient

	if err := bs.initKeyPair(); err != nil {
		return fmt.Errorf("failed to create key pair for batcher: %w", err)
	}

	// try to generate attestationBytes on public key when start batcher
	attestationBytes, err := enclave.AttestationWithPublicKey(bs.Espresso.BatcherPublicKey)
	if err != nil {
		bs.Log.Info("Not running in enclave, skipping attestation", "info", err)

		// Replace ephemeral keys with configured keys, as in devnet they'll be pre-approved for batching
		privateKey := cfg.Espresso.TestingBatcherPrivateKey
		if privateKey == nil {
			return fmt.Errorf("when not running in enclave, testing batcher private key should be set")
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error casting public key to ECDSA")
		}

		bs.Espresso.BatcherPrivateKey = privateKey
		bs.Espresso.BatcherPublicKey = publicKeyECDSA
	} else {
		// output length of attestation
		bs.Log.Info("Successfully got attestation. Attestation length", "length", len(attestationBytes))
		_, err := nitrite.Verify(attestationBytes, nitrite.VerifyOptions{})
		if err != nil {
			return fmt.Errorf("Couldn't verify attestation: %w", err)
		}
		bs.Attestation = attestationBytes
	}

	return nil
}
