package script

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
)

var (
	// DeterministicDeployerAddress is the address of the deterministic deployer Forge uses
	// to provide deterministic contract addresses.
	DeterministicDeployerAddress = common.HexToAddress("0x4e59b44847b379578588920ca78fbf26c0b4956c")

	// Arachnid deterministic deployment proxy runtime bytecode (https://github.com/Arachnid/deterministic-deployment-proxy).
	// Seeded on fork so CREATE2 via this deployer works when using WithCreate2Deployer (e.g. bootstrap on anvil).
	create2DeployerBytecode = mustDecodeHex("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf3")
)

func mustDecodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// EnsureCreate2Deployer sets the Create2Deployer contract code at DeterministicDeployerAddress
// if the account has no code. Call after selecting a fork so CREATE2 via the deployer works
// (e.g. when using WithCreate2Deployer on a fresh anvil fork).
func (h *Host) EnsureCreate2Deployer() {
	if h.state.GetCodeSize(DeterministicDeployerAddress) == 0 {
		h.state.SetCode(DeterministicDeployerAddress, create2DeployerBytecode, tracing.CodeChangeUnspecified)
	}
}
