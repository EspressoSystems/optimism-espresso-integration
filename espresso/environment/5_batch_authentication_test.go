package environment_test

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hf/nitrite"
)

// TestE2eDevNetWithInvalidAttestation verifies that the batcher correctly fails to register
// when provided with an invalid attestation. This test ensures that the batch inbox contract
// properly validates attestations
func TestE2eDevNetWithInvalidAttestation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate private key")
	}

	system, _, err :=
		launcher.StartDevNet(ctx, t,
			env.SetBatcherKey(*privateKey),
			env.Config(func(cfg *e2esys.SystemConfig) {
				cfg.DisableBatcher = true
			}),
		)

	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	batchDriver := system.BatchSubmitter.TestDriver()
	batchDriver.Attestation = &nitrite.Result{
		Document: &nitrite.Document{
			CABundle: [][]byte{[]byte{1, 2, 3, 4}},
		},
	}
	err = batchDriver.StartBatchSubmitting()

	if err == nil {
		t.Fatalf("batcher should've failed to register with invalid attestation but got nil error")
	}

	// Check for the key part of the error message
	expectedMsg := "could not register with batch inbox contract"
	errMsg := err.Error()
	if !strings.Contains(errMsg, expectedMsg) {
		t.Fatalf("error message does not contain expected message %q:\ngot: %q", expectedMsg, errMsg)
	}
}

// TestE2eDevNetWithUnattestedBatcherKey verifies that when a batcher key is not properly
// attested, the L2 chain can still produce unsafe blocks but cannot progress to safe L2 blocks.
func TestE2eDevNetWithUnattestedBatcherKey(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate private key")
	}

	system, _, err :=
		launcher.StartDevNet(ctx, t,
			env.SetBatcherKey(*privateKey),
		)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	l2Seq := system.NodeClient("sequencer")

	// Check that unsafe L2 is progressing...
	_, err = geth.WaitForBlock(big.NewInt(15), l2Seq)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to wait for block:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// ...but safe L2 is not
	_, err = geth.WaitForBlockToBeSafe(big.NewInt(1), l2Seq, 2*time.Minute)
	if err == nil {
		t.Fatalf("block shouldn't be safe")
	}

	_ = system
}
