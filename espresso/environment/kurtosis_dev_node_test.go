package environment_test

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	env "github.com/ethereum-optimism/optimism/espresso/environment"
// )

// func TestSmokeKurtosisEspressoDevNet(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	name := "espresso-devnet-test1"
// 	kurtosisEnclave, err := env.ConfigureKurtosisEspressoDevNet(name)
// 	if have, want := err, error(nil); have != want {
// 		t.Fatalf("failed to spawn kurtosis devnet:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}

// 	// Spin up the Kurtosis Devnet enclave
// 	if have, want := kurtosisEnclave.Spawn(), error(nil); have != want {
// 		t.Fatalf("failed to spawn kurtosis devnet:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}

// 	// Let's check that our L1, L2, and Espresso and making progress

// 	l1RPC, err := kurtosisEnclave.L1ExecutionLayerRPC(ctx)
// 	if have, want := err, error(nil); have != want {
// 		t.Fatalf("failed to get L1 RPC:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}

// 	l2RPC, err := kurtosisEnclave.L2ExecutionLayerRPC(ctx)
// 	if have, want := err, error(nil); have != want {
// 		t.Fatalf("failed to get L2 RPC:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}

// 	var header gethTypes.Header
// 	if have, want := l1RPC.Call(&header, "eth_getBlockByNumber", "latest", true), error(nil); have != want {
// 		t.Fatalf("failed to get L1 header:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}
// 	l1Height0 := new(big.Int)
// 	l1Height0.FillBytes(header.Number.Bytes())

// 	if have, want := l2RPC.Call(&header, "eth_getBlockByNumber", "latest", true), error(nil); have != want {
// 		t.Fatalf("failed to get L1 header:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}
// 	l2Height0 := new(big.Int)
// 	l2Height0.FillBytes(header.Number.Bytes())

// 	// Wait for some time to pass
// 	time.Sleep(time.Second * 5)

// 	if have, want := l1RPC.Call(&header, "eth_getBlockByNumber", "latest", true), error(nil); have != want {
// 		t.Fatalf("failed to get L1 header:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}
// 	l1Height1 := new(big.Int)
// 	l1Height1.FillBytes(header.Number.Bytes())

// 	if have, want := l2RPC.Call(&header, "eth_getBlockByNumber", "latest", true), error(nil); have != want {
// 		t.Fatalf("failed to get L1 header:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}
// 	l2Height1 := new(big.Int)
// 	l2Height1.FillBytes(header.Number.Bytes())

// 	if have, want := l1Height1.Cmp(l1Height0), 1; have != want {
// 		t.Fatalf("L1 height did not increase:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}

// 	if have, want := l2Height1.Cmp(l2Height0), 1; have != want {
// 		t.Fatalf("L2 height did not increase:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}

// 	// espressoDevNode, err := kurtosisEnclave.EspressoDevNode()
// 	// if have, want := err, error(nil); have != want {
// 	// 	t.Fatalf("failed to get Espresso Dev Node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	// }

// 	// Stop the enclave after starting it
// 	if have, want := kurtosisEnclave.Stop(), error(nil); have != want {
// 		t.Fatalf("failed to stop kurtosis devnet:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}

// 	if have, want := kurtosisEnclave.CleanAll(), error(nil); have != want {
// 		t.Fatalf("failed to clean kurtosis devnet:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
// 	}
// }
