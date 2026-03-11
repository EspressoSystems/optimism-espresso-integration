package enclave_tools

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EnclaveMeasurements struct {
	PCR0 string `json:"PCR0"`
	PCR1 string `json:"PCR1"`
	PCR2 string `json:"PCR2"`
}

// Builds docker and enclaver EIF image for op-batcher and registers EIF's PCR0 with
// EspressoNitroTEEVerifier. args... are command-line arguments to op-batcher
// to be baked into the image.
func BuildBatcherImage(ctx context.Context, opRoot string, tag string, args ...string) (EnclaveMeasurements, error) {
	intermediateTag := tag + "intermediate"

	dockerCli := new(environment.DockerCli)
	err := dockerCli.Build(
		ctx,
		intermediateTag,
		filepath.Join(opRoot, "ops/docker/op-stack-go/Dockerfile"),
		"op-batcher-enclave-target",
		opRoot,
		environment.DockerBuildArg{
			Name:  "ENCLAVE_BATCHER_ARGS",
			Value: strings.Join(args, " "),
		},
	)
	if err != nil {
		return EnclaveMeasurements{}, fmt.Errorf("failed to build intermediate docker image: %w", err)
	}

	// Build EIF image based on the docker image we just built
	enclaverCli := new(EnclaverCli)
	manifest := DefaultManifest("op-batcher", tag, intermediateTag)
	measurements, err := enclaverCli.BuildEnclave(ctx, manifest)
	return measurements, err
}

// BuildEifFromImage builds an EIF image by wrapping a pre-built app Docker image with Enclaver.
func BuildEifFromImage(ctx context.Context, appImage string, eifTag string) (EnclaveMeasurements, error) {
	manifest := DefaultManifest("op-batcher", eifTag, appImage)
	return new(EnclaverCli).BuildEnclave(ctx, manifest)
}

// getNitroVerifier retrieves the Nitro TEE verifier instance and L1 client by traversing the contract chain.
func getNitroVerifier(ctx context.Context, authenticatorAddress common.Address, L1Url string) (*bindings.EspressoNitroTEEVerifier, *ethclient.Client, error) {
	l1Client, err := ethclient.DialContext(ctx, L1Url)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to L1 client: %w", err)
	}

	authenticator, err := bindings.NewBatchAuthenticator(authenticatorAddress, l1Client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create batch authenticator: %w", err)
	}

	verifierAddress, err := authenticator.EspressoTEEVerifier(&bind.CallOpts{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get verifier address: %w", err)
	}

	verifier, err := bindings.NewEspressoTEEVerifier(verifierAddress, l1Client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create verifier: %w", err)
	}

	nitroVerifierAddress, err := verifier.EspressoNitroTEEVerifier(&bind.CallOpts{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get nitro verifier address: %w", err)
	}

	nitroVerifier, err := bindings.NewEspressoNitroTEEVerifier(nitroVerifierAddress, l1Client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create nitro verifier: %w", err)
	}

	return nitroVerifier, l1Client, nil
}

// RegisterEnclaveHash registers the enclave PCR0 hash with the EspressoNitroTEEVerifier.
func RegisterEnclaveHash(ctx context.Context, authenticatorAddress common.Address, L1Url string, key *ecdsa.PrivateKey, pcr0Bytes []byte) error {
	nitroVerifier, l1Client, err := getNitroVerifier(ctx, authenticatorAddress, L1Url)
	if err != nil {
		return err
	}

	chainID, err := l1Client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}
	registrationTx, err := nitroVerifier.SetEnclaveHash(opts, crypto.Keccak256Hash(pcr0Bytes), true)
	if err != nil {
		return fmt.Errorf("failed to create registration transaction: %w", err)
	}

	receipt, err := geth.WaitForTransaction(registrationTx.Hash(), l1Client, 2*time.Minute)
	if err != nil {
		return fmt.Errorf("failed to wait for registration transaction: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("registration transaction failed")
	}

	return nil
}

// IsEnclaveHashRegistered checks if the given PCR0 hash is already registered with the EspressoNitroTEEVerifier
func IsEnclaveHashRegistered(ctx context.Context, authenticatorAddress common.Address, L1Url string, pcr0Bytes []byte) (bool, error) {
	nitroVerifier, _, err := getNitroVerifier(ctx, authenticatorAddress, L1Url)
	if err != nil {
		return false, fmt.Errorf("failed to get nitro verifier: %w", err)
	}

	isRegisteredTx, err := nitroVerifier.RegisteredEnclaveHash(&bind.CallOpts{}, crypto.Keccak256Hash(pcr0Bytes))
	if err != nil {
		return false, fmt.Errorf("failed to call registeredEnclaveHash function: %w", err)
	}

	return isRegisteredTx, nil
}
