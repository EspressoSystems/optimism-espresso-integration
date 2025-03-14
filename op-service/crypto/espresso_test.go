package crypto

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// func TestSign(t *testing.T) {

// 	signerClient, err := opsigner.NewSignerClientFromConfig(l, signerConfig)
// 	if err != nil {
// 		l.Error("Unable to create Signer Client", "error", err)
// 		return nil, common.Address{}, fmt.Errorf("failed to create the signer client: %w", err)
// 	}
// 	signer := func(chainID *big.Int) ChainSigner {
// 		return &clientSigner{
// 			signerClient: signerClient,
// 			fromAddress:  common.HexToAddress("0x00a4FE4C6AaA0729d7699c387E7f281DD64aFA2a"),
// 			chainID:      chainID,
// 		}
// 	}

// 	sequencerBatchesByte, err := hex.DecodeString("1e7e580d65989969957450819e382bf27cd04eaf3d390f915b907091f5e50faa")
// 	require.NoError(t, err)
// 	signature, err := signer.Sign(context.Background(), common.HexToAddress("0x00a4FE4C6AaA0729d7699c387E7f281DD64aFA2a"), sequencerBatchesByte)
// 	require.NoError(t, err)

// 	batcherSignature, err := hex.DecodeString("39c969f723e8eefa9c367cd79e29a69dfc39084c9e46e929e3f6fc52e00fbb3b420e37e556434302dd971377d0a5d10b7da8062185eeb896352a952539133dc701")
// 	require.NoError(t, err)

// 	require.Equal(t, signature, batcherSignature)

// }

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
