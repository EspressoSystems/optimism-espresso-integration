package crypto

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestVerify(t *testing.T) {
	// logger := testlog.Logger(t, log.LevelDebug)

	// batcherSignature, err := hex.DecodeString("39c969f723e8eefa9c367cd79e29a69dfc39084c9e46e929e3f6fc52e00fbb3b420e37e556434302dd971377d0a5d10b7da8062185eeb896352a952539133dc701")
	// require.NoError(t, err)

	// sequencerBatchesByte, err := hex.DecodeString("1e7e580d65989969957450819e382bf27cd04eaf3d390f915b907091f5e50faa")
	// require.NoError(t, err)

	// expected := common.HexToAddress("0x00a4FE4C6AaA0729d7699c387E7f281DD64aFA2a")

	batcherSignature := []byte{
		109, 206, 105, 108, 152, 110, 156, 111, 239, 153, 224, 182, 140, 49, 105, 120,
		153, 163, 162, 47, 119, 34, 68, 128, 118, 33, 143, 79, 101, 212, 75, 161,
		124, 77, 236, 159, 70, 167, 95, 51, 92, 127, 236, 253, 4, 211, 222, 117,
		54, 27, 214, 232, 135, 87, 33, 77, 16, 155, 164, 116, 220, 116, 31, 208, 1,
	}
	sequencerBatchesByte := []byte{
		166, 136, 91, 55, 49, 112, 45, 166,
		46, 142, 74, 143, 88, 74, 196, 106,
		127, 104, 34, 244, 226, 186, 80, 251,
		169, 2, 246, 123, 21, 136, 210, 59,
	}

	expected := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	err := Verify(sequencerBatchesByte, batcherSignature, expected)
	require.NoError(t, err)
}
