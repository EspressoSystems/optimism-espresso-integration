package crypto

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestVerify(t *testing.T) {
	// logger := testlog.Logger(t, log.LevelDebug)

	batcherSignature, err := hex.DecodeString("39c969f723e8eefa9c367cd79e29a69dfc39084c9e46e929e3f6fc52e00fbb3b420e37e556434302dd971377d0a5d10b7da8062185eeb896352a952539133dc701")
	require.NoError(t, err)

	sequencerBatchesByte, err := hex.DecodeString("1e7e580d65989969957450819e382bf27cd04eaf3d390f915b907091f5e50faa")
	require.NoError(t, err)

	expected := common.HexToAddress("0x00a4FE4C6AaA0729d7699c387E7f281DD64aFA2a")

	err = Verify(sequencerBatchesByte, batcherSignature, expected)
	require.NoError(t, err)
}
