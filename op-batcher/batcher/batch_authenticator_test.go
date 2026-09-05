package batcher

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/espresso/bindings"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
)

var testAuthAddr = common.HexToAddress("0x00000000000000000000000000000000000000aa")

// mockAuthBackend is a minimal bind.ContractCaller standing in for the L1
// client.
type mockAuthBackend struct {
	abi *abi.ABI

	// code is returned by CodeAt. Empty means "not deployed yet".
	code []byte
	// codeErr, when set, fails CodeAt instead of returning code.
	codeErr error

	activeIsEspresso bool

	codeAtCalls int
	callCalls   int
	// lastCallHadDeadline records whether the context reaching CallContract
	// carried a deadline, i.e. whether the reader applied NetworkTimeout.
	lastCallHadDeadline bool
}

func newMockAuthBackend(t *testing.T) *mockAuthBackend {
	t.Helper()
	parsed, err := bindings.BatchAuthenticatorMetaData.GetAbi()
	require.NoError(t, err)
	return &mockAuthBackend{abi: parsed, code: []byte{0x60, 0x00}}
}

func (m *mockAuthBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.codeAtCalls++
	if m.codeErr != nil {
		return nil, m.codeErr
	}
	return m.code, nil
}

func (m *mockAuthBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	m.callCalls++
	_, m.lastCallHadDeadline = ctx.Deadline()
	if len(call.Data) < 4 {
		return nil, errors.New("short calldata")
	}
	if string(call.Data[:4]) != string(m.abi.Methods["activeIsEspresso"].ID) {
		return nil, errors.New("unexpected method call")
	}
	return m.abi.Methods["activeIsEspresso"].Outputs.Pack(m.activeIsEspresso)
}

func newTestReader(t *testing.T, backend *mockAuthBackend) *batchAuthenticatorReader {
	t.Helper()
	r, err := newBatchAuthenticatorReader(testAuthAddr, backend, time.Second)
	require.NoError(t, err)
	return r
}

// TestBatchAuthenticatorReader_ProbeLatches is the core claim of the shared
// reader: the deployment probe is paid once, not once per publish tick. It also
// checks that the reader applies the network timeout itself, so no call site
// can forget it.
func TestBatchAuthenticatorReader_ProbeLatches(t *testing.T) {
	backend := newMockAuthBackend(t)
	backend.activeIsEspresso = true
	r := newTestReader(t, backend)

	for i := 0; i < 5; i++ {
		active, err := r.ActiveIsEspresso(context.Background())
		require.NoError(t, err)
		require.True(t, active)
	}

	require.Equal(t, 1, backend.codeAtCalls, "deployment probe should be paid once, not per read")
	require.Equal(t, 5, backend.callCalls, "each read is still one eth_call")
	require.True(t, backend.lastCallHadDeadline, "reader must bound reads by NetworkTimeout")
}

// TestBatchAuthenticatorReader_ProbeFailureIsNotLatched covers the reason the
// probe stays lazy: a batcher may legitimately start before the
// BatchAuthenticator is deployed (the pre-fork fallback batcher). Such a
// batcher must keep retrying and recover once the contract appears, rather
// than caching the failure for the life of the process.
func TestBatchAuthenticatorReader_ProbeFailureIsNotLatched(t *testing.T) {
	backend := newMockAuthBackend(t)
	backend.code = nil // not deployed yet
	r := newTestReader(t, backend)

	_, err := r.ActiveIsEspresso(context.Background())
	require.ErrorContains(t, err, "no contract code at BatchAuthenticator address")
	require.Zero(t, backend.callCalls, "should not read from an undeployed address")

	backend.codeErr = errors.New("rpc boom")
	_, err = r.ActiveIsEspresso(context.Background())
	require.ErrorContains(t, err, "rpc boom")

	// Contract shows up; the reader recovers without needing a restart.
	backend.codeErr = nil
	backend.code = []byte{0x60, 0x00}
	backend.activeIsEspresso = true
	active, err := r.ActiveIsEspresso(context.Background())
	require.NoError(t, err)
	require.True(t, active)

	require.Equal(t, 3, backend.codeAtCalls)
	require.Equal(t, 1, backend.callCalls)
}

// TestIsBatcherActive covers the publish gate's decision table: this process is
// the active batcher exactly when the on-chain flag agrees with its own
// --espresso.enabled setting.
func TestIsBatcherActive(t *testing.T) {
	tests := []struct {
		name             string
		activeIsEspresso bool
		espressoEnabled  bool
		want             bool
	}{
		{"espresso batcher while espresso is active", true, true, true},
		{"espresso batcher while fallback is active", false, true, false},
		{"fallback batcher while espresso is active", true, false, false},
		{"fallback batcher while fallback is active", false, false, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backend := newMockAuthBackend(t)
			backend.activeIsEspresso = test.activeIsEspresso

			l := &BatchSubmitter{}
			l.Log = testlog.Logger(t, log.LevelDebug)
			l.Txmgr = &testutils.FakeTxMgr{}
			l.Config.Espresso.Enabled = test.espressoEnabled
			l.batchAuth = newTestReader(t, backend)

			require.True(t, l.hasBatchAuthenticator())

			got, err := l.isBatcherActive(context.Background())
			require.NoError(t, err)
			require.Equal(t, test.want, got)
		})
	}
}

// TestIsBatcherActive_NoAuthenticator guards the nil reader case: without a
// configured BatchAuthenticator the gate must report an error rather than
// silently treating this batcher as active.
func TestIsBatcherActive_NoAuthenticator(t *testing.T) {
	l := &BatchSubmitter{}
	l.Log = testlog.Logger(t, log.LevelDebug)

	require.False(t, l.hasBatchAuthenticator())

	_, err := l.isBatcherActive(context.Background())
	require.ErrorContains(t, err, "no BatchAuthenticator configured")
}
