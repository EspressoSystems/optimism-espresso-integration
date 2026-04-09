// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BatchAuthenticatorMetaData contains all meta data concerning the BatchAuthenticator contract.
var BatchAuthenticatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsEspresso\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"espressoBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGuardians\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guardianCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_espressoBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"contractISystemConfig\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isGuardian\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProxyAdmin\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdminOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"verificationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEspressoBatcher\",\"inputs\":[{\"name\":\"_newEspressoBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISystemConfig\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BatchInfoAuthenticated\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatcherSwitched\",\"inputs\":[{\"name\":\"activeIsEspresso\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EspressoBatcherUpdated\",\"inputs\":[{\"name\":\"oldEspressoBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEspressoBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianAdded\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianRemoved\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerRegistrationInitiated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BatchAuthenticator_Paused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"contract_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidGuardianAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotGuardian\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotGuardianOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnerCantBeGuardian\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOrProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotResolvedDelegateProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotSharedProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_ProxyAdminNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReinitializableBase_ZeroInitVersion\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060015f8160ff160361004e576040517f9b01afed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060ff1660808160ff16815250505061006b61007060201b60201c565b6101eb565b5f61007f61016e60201b60201c565b9050805f0160089054906101000a900460ff16156100c9576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461016b5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161016291906101d2565b60405180910390a15b50565b5f8061017e61018760201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b6101cc816101b0565b82525050565b5f6020820190506101e55f8301846101c3565b92915050565b60805161261b6102035f395f61065d015261261b5ff3fe608060405234801561000f575f80fd5b5060043610610171575f3560e01c806379ba5097116100dc578063dad544e011610095578063f2fde38b1161006f578063f2fde38b146103b9578063f8c8765e146103d5578063fa14fe6d146103f1578063fc619e411461040f57610171565b8063dad544e01461035f578063e30c39781461037d578063eca919df1461039b57610171565b806379ba5097146102d757806388da3bb7146102e15780638da5cb5b146102ff578063a526d83b1461031d578063ba58e82a14610339578063bc347f471461035557610171565b80633e47158c1161012e5780633e47158c1461023957806354387ad71461025757806354fd4d50146102755780635c975abb1461029357806371404156146102b1578063715018a6146102cd57610171565b80630665f04b146101755780630c68ba21146101935780631b076a4c146101c35780632ce53247146101e157806333d7e2bd146101fd57806338d38c971461021b575b5f80fd5b61017d61042b565b60405161018a9190611d08565b60405180910390f35b6101ad60048036038101906101a89190611d5a565b610444565b6040516101ba9190611d9f565b60405180910390f35b6101cb610468565b6040516101d89190611dc7565b60405180910390f35b6101fb60048036038101906101f69190611d5a565b6104fc565b005b610205610635565b6040516102129190611e3b565b60405180910390f35b61022361065a565b6040516102309190611e6f565b60405180910390f35b610241610681565b60405161024e9190611ea8565b60405180910390f35b61025f6108d0565b60405161026c9190611ed9565b60405180910390f35b61027d6108e8565b60405161028a9190611f62565b60405180910390f35b61029b610921565b6040516102a89190611d9f565b60405180910390f35b6102cb60048036038101906102c69190611d5a565b6109b5565b005b6102d5610a2f565b005b6102df610a42565b005b6102e9610ad0565b6040516102f69190611dc7565b60405180910390f35b610307610af3565b6040516103149190611dc7565b60405180910390f35b61033760048036038101906103329190611d5a565b610b01565b005b610353600480360381019061034e9190611fe3565b610c85565b005b61035d610d9f565b005b610367610ea7565b6040516103749190611dc7565b60405180910390f35b610385610f21565b6040516103929190611dc7565b60405180910390f35b6103a3610f56565b6040516103b09190611d9f565b60405180910390f35b6103d360048036038101906103ce9190611d5a565b610f69565b005b6103ef60048036038101906103ea91906120d7565b610fd0565b005b6103f961133c565b604051610406919061215b565b60405180910390f35b610429600480360381019061042491906121a7565b611361565b005b606061043f610438611475565b5f0161149c565b905090565b5f61046182610451611475565b5f016114bb90919063ffffffff16565b9050919050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d80a4c286040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104f7919061223f565b905090565b6105046114e8565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361057457806040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815260040161056b9190611dc7565b60405180910390fd5b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fc7bef7b97a10ef514a01fa4d5552f5c57e72a37aa901f567bb49bbf9ea449f9c60405160405180910390a35050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f806106ae7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035f1b61156f565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146106ec57809150506108cd565b60026040518060400160405280601a81526020017f4f564d5f4c3143726f7373446f6d61696e4d657373656e6765720000000000008152505161072f9190612297565b7f4f564d5f4c3143726f7373446f6d61696e4d657373656e6765720000000000005f1c175f1b610786305f60405160200161076b9291906122d8565b60405160208183030381529060405280519060200120611579565b146107bd576040517f54e433cd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6107f03060016040516020016107d59291906122d8565b6040516020818303038152906040528051906020012061156f565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461089b578073ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561086e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108929190612313565b925050506108cd565b6040517f332144db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b90565b5f6108e36108dc611475565b5f01611583565b905090565b6040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561098c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b09190612368565b905090565b6109bd6114e8565b5f6109c6611475565b90506109dd82825f0161159690919063ffffffff16565b6109e75750610a2c565b8173ffffffffffffffffffffffffffffffffffffffff167fb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c5260405160405180910390a2505b50565b610a376114e8565b610a405f6115c3565b565b5f610a4b611600565b90508073ffffffffffffffffffffffffffffffffffffffff16610a6c610f21565b73ffffffffffffffffffffffffffffffffffffffff1614610ac457806040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610abb9190611dc7565b60405180910390fd5b610acd816115c3565b50565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f610afc611607565b905090565b610b096114e8565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b6e576040517f1b08105400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b76610af3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610be15750610bb2610f21565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15610c18576040517f3af3c41c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610c21611475565b9050610c3882825f0161163c90919063ffffffff16565b15610c81578173ffffffffffffffffffffffffffffffffffffffff167f038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f96960405160405180910390a25b5050565b610c8d610921565b15610cc4576040517fb3a266c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637f82ea6c8585858560015f6040518763ffffffff1660e01b8152600401610d2996959493929190612496565b5f604051808303815f87803b158015610d40575f80fd5b505af1158015610d52573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff167f665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c60405160405180910390a250505050565b610dbb33610dab611475565b5f016114bb90919063ffffffff16565b158015610dfb5750610dcb610af3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610e3d57336040517fd53780c4000000000000000000000000000000000000000000000000000000008152600401610e349190611dc7565b60405180910390fd5b600160149054906101000a900460ff1615600160146101000a81548160ff021916908315150217905550600160149054906101000a900460ff1615157fb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad360405160405180910390a2565b5f610eb0610681565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ef8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f1c9190612313565b905090565b5f80610f2b611669565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b600160149054906101000a900460ff1681565b610f716114e8565b610f8d81610f7d611475565b5f016114bb90919063ffffffff16565b15610fc4576040517f3af3c41c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fcd81611690565b50565b610fd861065a565b60ff165f610fe4611749565b9050805f0160089054906101000a900460ff168061102c57508167ffffffffffffffff16815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1610155b15611063576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001815f0160086101000a81548160ff0219169083151502179055506110b061175c565b6110b98361180b565b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361112957846040517f8e4c8aa60000000000000000000000000000000000000000000000000000000081526004016111209190611dc7565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff160361119957836040517f8e4c8aa60000000000000000000000000000000000000000000000000000000081526004016111909190611dc7565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff160361120957856040517f8e4c8aa60000000000000000000000000000000000000000000000000000000081526004016112009190611dc7565b60405180910390fd5b8560015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550845f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060146101000a81548160ff0219169083151502179055505f815f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d28260405161132c919061250d565b60405180910390a1505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611369610921565b156113a0576040517fb3a266c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355ddfa0683838660015f6040518663ffffffff1660e01b8152600401611403959493929190612535565b602060405180830381865afa15801561141e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114429190612368565b50827fee0d07d204d979d28885955e59a46f754c4db7378b7df1a95123525aac6e3f8060405160405180910390a2505050565b5f7f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f00905090565b60605f6114aa835f0161181f565b905060608190508092505050919050565b5f6114e0835f018373ffffffffffffffffffffffffffffffffffffffff165f1b611878565b905092915050565b6114f0611600565b73ffffffffffffffffffffffffffffffffffffffff1661150e610af3565b73ffffffffffffffffffffffffffffffffffffffff161461156d57611531611600565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016115649190611dc7565b60405180910390fd5b565b5f81549050919050565b5f81549050919050565b5f61158f825f01611898565b9050919050565b5f6115bb835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6118a7565b905092915050565b5f6115cc611669565b9050805f015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556115fc826119a3565b5050565b5f33905090565b5f80611611611a74565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b5f611661835f018373ffffffffffffffffffffffffffffffffffffffff165f1b611a9b565b905092915050565b5f7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00905090565b6116986114e8565b5f6116a1611669565b905081815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff16611703610af3565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b5f80611753611b02565b90508091505090565b3373ffffffffffffffffffffffffffffffffffffffff1661177b610681565b73ffffffffffffffffffffffffffffffffffffffff16141580156117d257503373ffffffffffffffffffffffffffffffffffffffff166117b9610ea7565b73ffffffffffffffffffffffffffffffffffffffff1614155b15611809576040517fc4050a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611813611b2b565b61181c81611b6b565b50565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561186c57602002820191905f5260205f20905b815481526020019060010190808311611858575b50505050509050919050565b5f80836001015f8481526020019081526020015f20541415905092915050565b5f815f01805490509050919050565b5f80836001015f8481526020019081526020015f205490505f8114611998575f6001826118d49190612581565b90505f6001865f01805490506118ea9190612581565b9050808214611950575f865f018281548110611909576119086125b4565b5b905f5260205f200154905080875f01848154811061192a576119296125b4565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f01805480611963576119626125e1565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061199d565b5f9150505b92915050565b5f6119ac611a74565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f611aa68383611878565b611af857825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f208190555060019050611afc565b5f90505b92915050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b611b33611b7f565b611b69576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611b73611b2b565b611b7c81611b9d565b50565b5f611b88611749565b5f0160089054906101000a900460ff16905090565b611ba5611b2b565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611c15575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401611c0c9190611dc7565b60405180910390fd5b611c1e816115c3565b50565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611c7382611c4a565b9050919050565b611c8381611c69565b82525050565b5f611c948383611c7a565b60208301905092915050565b5f602082019050919050565b5f611cb682611c21565b611cc08185611c2b565b9350611ccb83611c3b565b805f5b83811015611cfb578151611ce28882611c89565b9750611ced83611ca0565b925050600181019050611cce565b5085935050505092915050565b5f6020820190508181035f830152611d208184611cac565b905092915050565b5f80fd5b5f80fd5b611d3981611c69565b8114611d43575f80fd5b50565b5f81359050611d5481611d30565b92915050565b5f60208284031215611d6f57611d6e611d28565b5b5f611d7c84828501611d46565b91505092915050565b5f8115159050919050565b611d9981611d85565b82525050565b5f602082019050611db25f830184611d90565b92915050565b611dc181611c69565b82525050565b5f602082019050611dda5f830184611db8565b92915050565b5f819050919050565b5f611e03611dfe611df984611c4a565b611de0565b611c4a565b9050919050565b5f611e1482611de9565b9050919050565b5f611e2582611e0a565b9050919050565b611e3581611e1b565b82525050565b5f602082019050611e4e5f830184611e2c565b92915050565b5f60ff82169050919050565b611e6981611e54565b82525050565b5f602082019050611e825f830184611e60565b92915050565b5f611e9282611e0a565b9050919050565b611ea281611e88565b82525050565b5f602082019050611ebb5f830184611e99565b92915050565b5f819050919050565b611ed381611ec1565b82525050565b5f602082019050611eec5f830184611eca565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611f3482611ef2565b611f3e8185611efc565b9350611f4e818560208601611f0c565b611f5781611f1a565b840191505092915050565b5f6020820190508181035f830152611f7a8184611f2a565b905092915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112611fa357611fa2611f82565b5b8235905067ffffffffffffffff811115611fc057611fbf611f86565b5b602083019150836001820283011115611fdc57611fdb611f8a565b5b9250929050565b5f805f8060408587031215611ffb57611ffa611d28565b5b5f85013567ffffffffffffffff81111561201857612017611d2c565b5b61202487828801611f8e565b9450945050602085013567ffffffffffffffff81111561204757612046611d2c565b5b61205387828801611f8e565b925092505092959194509250565b5f61206b82611c69565b9050919050565b61207b81612061565b8114612085575f80fd5b50565b5f8135905061209681612072565b92915050565b5f6120a682611c69565b9050919050565b6120b68161209c565b81146120c0575f80fd5b50565b5f813590506120d1816120ad565b92915050565b5f805f80608085870312156120ef576120ee611d28565b5b5f6120fc87828801612088565b945050602061210d87828801611d46565b935050604061211e878288016120c3565b925050606061212f87828801611d46565b91505092959194509250565b5f61214582611e0a565b9050919050565b6121558161213b565b82525050565b5f60208201905061216e5f83018461214c565b92915050565b5f819050919050565b61218681612174565b8114612190575f80fd5b50565b5f813590506121a18161217d565b92915050565b5f805f604084860312156121be576121bd611d28565b5b5f6121cb86828701612193565b935050602084013567ffffffffffffffff8111156121ec576121eb611d2c565b5b6121f886828701611f8e565b92509250509250925092565b5f61220e82611c69565b9050919050565b61221e81612204565b8114612228575f80fd5b50565b5f8151905061223981612215565b92915050565b5f6020828403121561225457612253611d28565b5b5f6122618482850161222b565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6122a182611ec1565b91506122ac83611ec1565b92508282026122ba81611ec1565b915082820484148315176122d1576122d061226a565b5b5092915050565b5f6040820190506122eb5f830185611db8565b6122f86020830184611eca565b9392505050565b5f8151905061230d81611d30565b92915050565b5f6020828403121561232857612327611d28565b5b5f612335848285016122ff565b91505092915050565b61234781611d85565b8114612351575f80fd5b50565b5f815190506123628161233e565b92915050565b5f6020828403121561237d5761237c611d28565b5b5f61238a84828501612354565b91505092915050565b5f82825260208201905092915050565b828183375f83830152505050565b5f6123bc8385612393565b93506123c98385846123a3565b6123d283611f1a565b840190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6002811061241b5761241a6123dd565b5b50565b5f81905061242b8261240a565b919050565b5f61243a8261241e565b9050919050565b61244a81612430565b82525050565b60028110612461576124606123dd565b5b50565b5f81905061247182612450565b919050565b5f61248082612464565b9050919050565b61249081612476565b82525050565b5f6080820190508181035f8301526124af81888a6123b1565b905081810360208301526124c48186886123b1565b90506124d36040830185612441565b6124e06060830184612487565b979650505050505050565b5f67ffffffffffffffff82169050919050565b612507816124eb565b82525050565b5f6020820190506125205f8301846124fe565b92915050565b61252f81612174565b82525050565b5f6080820190508181035f83015261254e8187896123b1565b905061255d6020830186612526565b61256a6040830185612441565b6125776060830184612487565b9695505050505050565b5f61258b82611ec1565b915061259683611ec1565b92508282039050818111156125ae576125ad61226a565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea164736f6c6343000819000a",
}

// BatchAuthenticatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchAuthenticatorMetaData.ABI instead.
var BatchAuthenticatorABI = BatchAuthenticatorMetaData.ABI

// BatchAuthenticatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchAuthenticatorMetaData.Bin instead.
var BatchAuthenticatorBin = BatchAuthenticatorMetaData.Bin

// DeployBatchAuthenticator deploys a new Ethereum contract, binding an instance of BatchAuthenticator to it.
func DeployBatchAuthenticator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchAuthenticator, error) {
	parsed, err := BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchAuthenticatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchAuthenticator{BatchAuthenticatorCaller: BatchAuthenticatorCaller{contract: contract}, BatchAuthenticatorTransactor: BatchAuthenticatorTransactor{contract: contract}, BatchAuthenticatorFilterer: BatchAuthenticatorFilterer{contract: contract}}, nil
}

// BatchAuthenticator is an auto generated Go binding around an Ethereum contract.
type BatchAuthenticator struct {
	BatchAuthenticatorCaller     // Read-only binding to the contract
	BatchAuthenticatorTransactor // Write-only binding to the contract
	BatchAuthenticatorFilterer   // Log filterer for contract events
}

// BatchAuthenticatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchAuthenticatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchAuthenticatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchAuthenticatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchAuthenticatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchAuthenticatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchAuthenticatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchAuthenticatorSession struct {
	Contract     *BatchAuthenticator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BatchAuthenticatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchAuthenticatorCallerSession struct {
	Contract *BatchAuthenticatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BatchAuthenticatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchAuthenticatorTransactorSession struct {
	Contract     *BatchAuthenticatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BatchAuthenticatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchAuthenticatorRaw struct {
	Contract *BatchAuthenticator // Generic contract binding to access the raw methods on
}

// BatchAuthenticatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchAuthenticatorCallerRaw struct {
	Contract *BatchAuthenticatorCaller // Generic read-only contract binding to access the raw methods on
}

// BatchAuthenticatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchAuthenticatorTransactorRaw struct {
	Contract *BatchAuthenticatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchAuthenticator creates a new instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticator(address common.Address, backend bind.ContractBackend) (*BatchAuthenticator, error) {
	contract, err := bindBatchAuthenticator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticator{BatchAuthenticatorCaller: BatchAuthenticatorCaller{contract: contract}, BatchAuthenticatorTransactor: BatchAuthenticatorTransactor{contract: contract}, BatchAuthenticatorFilterer: BatchAuthenticatorFilterer{contract: contract}}, nil
}

// NewBatchAuthenticatorCaller creates a new read-only instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticatorCaller(address common.Address, caller bind.ContractCaller) (*BatchAuthenticatorCaller, error) {
	contract, err := bindBatchAuthenticator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorCaller{contract: contract}, nil
}

// NewBatchAuthenticatorTransactor creates a new write-only instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchAuthenticatorTransactor, error) {
	contract, err := bindBatchAuthenticator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorTransactor{contract: contract}, nil
}

// NewBatchAuthenticatorFilterer creates a new log filterer instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchAuthenticatorFilterer, error) {
	contract, err := bindBatchAuthenticator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorFilterer{contract: contract}, nil
}

// bindBatchAuthenticator binds a generic wrapper to an already deployed contract.
func bindBatchAuthenticator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchAuthenticator *BatchAuthenticatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchAuthenticator.Contract.BatchAuthenticatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchAuthenticator *BatchAuthenticatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.BatchAuthenticatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchAuthenticator *BatchAuthenticatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.BatchAuthenticatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchAuthenticator *BatchAuthenticatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchAuthenticator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchAuthenticator *BatchAuthenticatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchAuthenticator *BatchAuthenticatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.contract.Transact(opts, method, params...)
}

// ActiveIsEspresso is a free data retrieval call binding the contract method 0xeca919df.
//
// Solidity: function activeIsEspresso() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ActiveIsEspresso(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "activeIsEspresso")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveIsEspresso is a free data retrieval call binding the contract method 0xeca919df.
//
// Solidity: function activeIsEspresso() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) ActiveIsEspresso() (bool, error) {
	return _BatchAuthenticator.Contract.ActiveIsEspresso(&_BatchAuthenticator.CallOpts)
}

// ActiveIsEspresso is a free data retrieval call binding the contract method 0xeca919df.
//
// Solidity: function activeIsEspresso() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ActiveIsEspresso() (bool, error) {
	return _BatchAuthenticator.Contract.ActiveIsEspresso(&_BatchAuthenticator.CallOpts)
}

// EspressoBatcher is a free data retrieval call binding the contract method 0x88da3bb7.
//
// Solidity: function espressoBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) EspressoBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "espressoBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoBatcher is a free data retrieval call binding the contract method 0x88da3bb7.
//
// Solidity: function espressoBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) EspressoBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoBatcher(&_BatchAuthenticator.CallOpts)
}

// EspressoBatcher is a free data retrieval call binding the contract method 0x88da3bb7.
//
// Solidity: function espressoBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) EspressoBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoBatcher(&_BatchAuthenticator.CallOpts)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) EspressoTEEVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "espressoTEEVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) EspressoTEEVerifier() (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoTEEVerifier(&_BatchAuthenticator.CallOpts)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) EspressoTEEVerifier() (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoTEEVerifier(&_BatchAuthenticator.CallOpts)
}

// GetGuardians is a free data retrieval call binding the contract method 0x0665f04b.
//
// Solidity: function getGuardians() view returns(address[])
func (_BatchAuthenticator *BatchAuthenticatorCaller) GetGuardians(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "getGuardians")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGuardians is a free data retrieval call binding the contract method 0x0665f04b.
//
// Solidity: function getGuardians() view returns(address[])
func (_BatchAuthenticator *BatchAuthenticatorSession) GetGuardians() ([]common.Address, error) {
	return _BatchAuthenticator.Contract.GetGuardians(&_BatchAuthenticator.CallOpts)
}

// GetGuardians is a free data retrieval call binding the contract method 0x0665f04b.
//
// Solidity: function getGuardians() view returns(address[])
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) GetGuardians() ([]common.Address, error) {
	return _BatchAuthenticator.Contract.GetGuardians(&_BatchAuthenticator.CallOpts)
}

// GuardianCount is a free data retrieval call binding the contract method 0x54387ad7.
//
// Solidity: function guardianCount() view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorCaller) GuardianCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "guardianCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardianCount is a free data retrieval call binding the contract method 0x54387ad7.
//
// Solidity: function guardianCount() view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorSession) GuardianCount() (*big.Int, error) {
	return _BatchAuthenticator.Contract.GuardianCount(&_BatchAuthenticator.CallOpts)
}

// GuardianCount is a free data retrieval call binding the contract method 0x54387ad7.
//
// Solidity: function guardianCount() view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) GuardianCount() (*big.Int, error) {
	return _BatchAuthenticator.Contract.GuardianCount(&_BatchAuthenticator.CallOpts)
}

// InitVersion is a free data retrieval call binding the contract method 0x38d38c97.
//
// Solidity: function initVersion() view returns(uint8)
func (_BatchAuthenticator *BatchAuthenticatorCaller) InitVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "initVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InitVersion is a free data retrieval call binding the contract method 0x38d38c97.
//
// Solidity: function initVersion() view returns(uint8)
func (_BatchAuthenticator *BatchAuthenticatorSession) InitVersion() (uint8, error) {
	return _BatchAuthenticator.Contract.InitVersion(&_BatchAuthenticator.CallOpts)
}

// InitVersion is a free data retrieval call binding the contract method 0x38d38c97.
//
// Solidity: function initVersion() view returns(uint8)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) InitVersion() (uint8, error) {
	return _BatchAuthenticator.Contract.InitVersion(&_BatchAuthenticator.CallOpts)
}

// IsGuardian is a free data retrieval call binding the contract method 0x0c68ba21.
//
// Solidity: function isGuardian(address account) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) IsGuardian(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "isGuardian", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGuardian is a free data retrieval call binding the contract method 0x0c68ba21.
//
// Solidity: function isGuardian(address account) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) IsGuardian(account common.Address) (bool, error) {
	return _BatchAuthenticator.Contract.IsGuardian(&_BatchAuthenticator.CallOpts, account)
}

// IsGuardian is a free data retrieval call binding the contract method 0x0c68ba21.
//
// Solidity: function isGuardian(address account) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) IsGuardian(account common.Address) (bool, error) {
	return _BatchAuthenticator.Contract.IsGuardian(&_BatchAuthenticator.CallOpts, account)
}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) NitroValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "nitroValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) NitroValidator() (common.Address, error) {
	return _BatchAuthenticator.Contract.NitroValidator(&_BatchAuthenticator.CallOpts)
}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) NitroValidator() (common.Address, error) {
	return _BatchAuthenticator.Contract.NitroValidator(&_BatchAuthenticator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) Owner() (common.Address, error) {
	return _BatchAuthenticator.Contract.Owner(&_BatchAuthenticator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) Owner() (common.Address, error) {
	return _BatchAuthenticator.Contract.Owner(&_BatchAuthenticator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) Paused() (bool, error) {
	return _BatchAuthenticator.Contract.Paused(&_BatchAuthenticator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) Paused() (bool, error) {
	return _BatchAuthenticator.Contract.Paused(&_BatchAuthenticator.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) PendingOwner() (common.Address, error) {
	return _BatchAuthenticator.Contract.PendingOwner(&_BatchAuthenticator.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) PendingOwner() (common.Address, error) {
	return _BatchAuthenticator.Contract.PendingOwner(&_BatchAuthenticator.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) ProxyAdmin() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdmin(&_BatchAuthenticator.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ProxyAdmin() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdmin(&_BatchAuthenticator.CallOpts)
}

// ProxyAdminOwner is a free data retrieval call binding the contract method 0xdad544e0.
//
// Solidity: function proxyAdminOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ProxyAdminOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "proxyAdminOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdminOwner is a free data retrieval call binding the contract method 0xdad544e0.
//
// Solidity: function proxyAdminOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) ProxyAdminOwner() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdminOwner(&_BatchAuthenticator.CallOpts)
}

// ProxyAdminOwner is a free data retrieval call binding the contract method 0xdad544e0.
//
// Solidity: function proxyAdminOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ProxyAdminOwner() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdminOwner(&_BatchAuthenticator.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) SystemConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "systemConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) SystemConfig() (common.Address, error) {
	return _BatchAuthenticator.Contract.SystemConfig(&_BatchAuthenticator.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) SystemConfig() (common.Address, error) {
	return _BatchAuthenticator.Contract.SystemConfig(&_BatchAuthenticator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchAuthenticator *BatchAuthenticatorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchAuthenticator *BatchAuthenticatorSession) Version() (string, error) {
	return _BatchAuthenticator.Contract.Version(&_BatchAuthenticator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) Version() (string, error) {
	return _BatchAuthenticator.Contract.Version(&_BatchAuthenticator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AcceptOwnership(&_BatchAuthenticator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AcceptOwnership(&_BatchAuthenticator.TransactOpts)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AddGuardian(&_BatchAuthenticator.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AddGuardian(&_BatchAuthenticator.TransactOpts, guardian)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) AuthenticateBatchInfo(opts *bind.TransactOpts, commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "authenticateBatchInfo", commitment, _signature)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) AuthenticateBatchInfo(commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AuthenticateBatchInfo(&_BatchAuthenticator.TransactOpts, commitment, _signature)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) AuthenticateBatchInfo(commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AuthenticateBatchInfo(&_BatchAuthenticator.TransactOpts, commitment, _signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _espressoBatcher, address _systemConfig, address _owner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) Initialize(opts *bind.TransactOpts, _espressoTEEVerifier common.Address, _espressoBatcher common.Address, _systemConfig common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "initialize", _espressoTEEVerifier, _espressoBatcher, _systemConfig, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _espressoBatcher, address _systemConfig, address _owner) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) Initialize(_espressoTEEVerifier common.Address, _espressoBatcher common.Address, _systemConfig common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _espressoBatcher, _systemConfig, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _espressoBatcher, address _systemConfig, address _owner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) Initialize(_espressoTEEVerifier common.Address, _espressoBatcher common.Address, _systemConfig common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _espressoBatcher, _systemConfig, _owner)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes verificationData, bytes data) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RegisterSigner(opts *bind.TransactOpts, verificationData []byte, data []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "registerSigner", verificationData, data)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes verificationData, bytes data) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RegisterSigner(verificationData []byte, data []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, verificationData, data)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes verificationData, bytes data) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RegisterSigner(verificationData []byte, data []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, verificationData, data)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RemoveGuardian(&_BatchAuthenticator.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RemoveGuardian(&_BatchAuthenticator.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RenounceOwnership(&_BatchAuthenticator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RenounceOwnership(&_BatchAuthenticator.TransactOpts)
}

// SetEspressoBatcher is a paid mutator transaction binding the contract method 0x2ce53247.
//
// Solidity: function setEspressoBatcher(address _newEspressoBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SetEspressoBatcher(opts *bind.TransactOpts, _newEspressoBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "setEspressoBatcher", _newEspressoBatcher)
}

// SetEspressoBatcher is a paid mutator transaction binding the contract method 0x2ce53247.
//
// Solidity: function setEspressoBatcher(address _newEspressoBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SetEspressoBatcher(_newEspressoBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetEspressoBatcher(&_BatchAuthenticator.TransactOpts, _newEspressoBatcher)
}

// SetEspressoBatcher is a paid mutator transaction binding the contract method 0x2ce53247.
//
// Solidity: function setEspressoBatcher(address _newEspressoBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SetEspressoBatcher(_newEspressoBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetEspressoBatcher(&_BatchAuthenticator.TransactOpts, _newEspressoBatcher)
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SwitchBatcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "switchBatcher")
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SwitchBatcher() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SwitchBatcher(&_BatchAuthenticator.TransactOpts)
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SwitchBatcher() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SwitchBatcher(&_BatchAuthenticator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.TransferOwnership(&_BatchAuthenticator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.TransferOwnership(&_BatchAuthenticator.TransactOpts, newOwner)
}

// BatchAuthenticatorBatchInfoAuthenticatedIterator is returned from FilterBatchInfoAuthenticated and is used to iterate over the raw logs and unpacked data for BatchInfoAuthenticated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorBatchInfoAuthenticatedIterator struct {
	Event *BatchAuthenticatorBatchInfoAuthenticated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorBatchInfoAuthenticatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorBatchInfoAuthenticated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorBatchInfoAuthenticated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorBatchInfoAuthenticatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorBatchInfoAuthenticatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorBatchInfoAuthenticated represents a BatchInfoAuthenticated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorBatchInfoAuthenticated struct {
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchInfoAuthenticated is a free log retrieval operation binding the contract event 0xee0d07d204d979d28885955e59a46f754c4db7378b7df1a95123525aac6e3f80.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterBatchInfoAuthenticated(opts *bind.FilterOpts, commitment [][32]byte) (*BatchAuthenticatorBatchInfoAuthenticatedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "BatchInfoAuthenticated", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorBatchInfoAuthenticatedIterator{contract: _BatchAuthenticator.contract, event: "BatchInfoAuthenticated", logs: logs, sub: sub}, nil
}

// WatchBatchInfoAuthenticated is a free log subscription operation binding the contract event 0xee0d07d204d979d28885955e59a46f754c4db7378b7df1a95123525aac6e3f80.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchBatchInfoAuthenticated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorBatchInfoAuthenticated, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "BatchInfoAuthenticated", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorBatchInfoAuthenticated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "BatchInfoAuthenticated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchInfoAuthenticated is a log parse operation binding the contract event 0xee0d07d204d979d28885955e59a46f754c4db7378b7df1a95123525aac6e3f80.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseBatchInfoAuthenticated(log types.Log) (*BatchAuthenticatorBatchInfoAuthenticated, error) {
	event := new(BatchAuthenticatorBatchInfoAuthenticated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "BatchInfoAuthenticated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorBatcherSwitchedIterator is returned from FilterBatcherSwitched and is used to iterate over the raw logs and unpacked data for BatcherSwitched events raised by the BatchAuthenticator contract.
type BatchAuthenticatorBatcherSwitchedIterator struct {
	Event *BatchAuthenticatorBatcherSwitched // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorBatcherSwitchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorBatcherSwitched)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorBatcherSwitched)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorBatcherSwitchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorBatcherSwitchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorBatcherSwitched represents a BatcherSwitched event raised by the BatchAuthenticator contract.
type BatchAuthenticatorBatcherSwitched struct {
	ActiveIsEspresso bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBatcherSwitched is a free log retrieval operation binding the contract event 0xb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad3.
//
// Solidity: event BatcherSwitched(bool indexed activeIsEspresso)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterBatcherSwitched(opts *bind.FilterOpts, activeIsEspresso []bool) (*BatchAuthenticatorBatcherSwitchedIterator, error) {

	var activeIsEspressoRule []interface{}
	for _, activeIsEspressoItem := range activeIsEspresso {
		activeIsEspressoRule = append(activeIsEspressoRule, activeIsEspressoItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "BatcherSwitched", activeIsEspressoRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorBatcherSwitchedIterator{contract: _BatchAuthenticator.contract, event: "BatcherSwitched", logs: logs, sub: sub}, nil
}

// WatchBatcherSwitched is a free log subscription operation binding the contract event 0xb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad3.
//
// Solidity: event BatcherSwitched(bool indexed activeIsEspresso)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchBatcherSwitched(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorBatcherSwitched, activeIsEspresso []bool) (event.Subscription, error) {

	var activeIsEspressoRule []interface{}
	for _, activeIsEspressoItem := range activeIsEspresso {
		activeIsEspressoRule = append(activeIsEspressoRule, activeIsEspressoItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "BatcherSwitched", activeIsEspressoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorBatcherSwitched)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "BatcherSwitched", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatcherSwitched is a log parse operation binding the contract event 0xb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad3.
//
// Solidity: event BatcherSwitched(bool indexed activeIsEspresso)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseBatcherSwitched(log types.Log) (*BatchAuthenticatorBatcherSwitched, error) {
	event := new(BatchAuthenticatorBatcherSwitched)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "BatcherSwitched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorEspressoBatcherUpdatedIterator is returned from FilterEspressoBatcherUpdated and is used to iterate over the raw logs and unpacked data for EspressoBatcherUpdated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorEspressoBatcherUpdatedIterator struct {
	Event *BatchAuthenticatorEspressoBatcherUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorEspressoBatcherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorEspressoBatcherUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorEspressoBatcherUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorEspressoBatcherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorEspressoBatcherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorEspressoBatcherUpdated represents a EspressoBatcherUpdated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorEspressoBatcherUpdated struct {
	OldEspressoBatcher common.Address
	NewEspressoBatcher common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEspressoBatcherUpdated is a free log retrieval operation binding the contract event 0xc7bef7b97a10ef514a01fa4d5552f5c57e72a37aa901f567bb49bbf9ea449f9c.
//
// Solidity: event EspressoBatcherUpdated(address indexed oldEspressoBatcher, address indexed newEspressoBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterEspressoBatcherUpdated(opts *bind.FilterOpts, oldEspressoBatcher []common.Address, newEspressoBatcher []common.Address) (*BatchAuthenticatorEspressoBatcherUpdatedIterator, error) {

	var oldEspressoBatcherRule []interface{}
	for _, oldEspressoBatcherItem := range oldEspressoBatcher {
		oldEspressoBatcherRule = append(oldEspressoBatcherRule, oldEspressoBatcherItem)
	}
	var newEspressoBatcherRule []interface{}
	for _, newEspressoBatcherItem := range newEspressoBatcher {
		newEspressoBatcherRule = append(newEspressoBatcherRule, newEspressoBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "EspressoBatcherUpdated", oldEspressoBatcherRule, newEspressoBatcherRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorEspressoBatcherUpdatedIterator{contract: _BatchAuthenticator.contract, event: "EspressoBatcherUpdated", logs: logs, sub: sub}, nil
}

// WatchEspressoBatcherUpdated is a free log subscription operation binding the contract event 0xc7bef7b97a10ef514a01fa4d5552f5c57e72a37aa901f567bb49bbf9ea449f9c.
//
// Solidity: event EspressoBatcherUpdated(address indexed oldEspressoBatcher, address indexed newEspressoBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchEspressoBatcherUpdated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorEspressoBatcherUpdated, oldEspressoBatcher []common.Address, newEspressoBatcher []common.Address) (event.Subscription, error) {

	var oldEspressoBatcherRule []interface{}
	for _, oldEspressoBatcherItem := range oldEspressoBatcher {
		oldEspressoBatcherRule = append(oldEspressoBatcherRule, oldEspressoBatcherItem)
	}
	var newEspressoBatcherRule []interface{}
	for _, newEspressoBatcherItem := range newEspressoBatcher {
		newEspressoBatcherRule = append(newEspressoBatcherRule, newEspressoBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "EspressoBatcherUpdated", oldEspressoBatcherRule, newEspressoBatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorEspressoBatcherUpdated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "EspressoBatcherUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEspressoBatcherUpdated is a log parse operation binding the contract event 0xc7bef7b97a10ef514a01fa4d5552f5c57e72a37aa901f567bb49bbf9ea449f9c.
//
// Solidity: event EspressoBatcherUpdated(address indexed oldEspressoBatcher, address indexed newEspressoBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseEspressoBatcherUpdated(log types.Log) (*BatchAuthenticatorEspressoBatcherUpdated, error) {
	event := new(BatchAuthenticatorEspressoBatcherUpdated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "EspressoBatcherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorGuardianAddedIterator is returned from FilterGuardianAdded and is used to iterate over the raw logs and unpacked data for GuardianAdded events raised by the BatchAuthenticator contract.
type BatchAuthenticatorGuardianAddedIterator struct {
	Event *BatchAuthenticatorGuardianAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorGuardianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorGuardianAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorGuardianAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorGuardianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorGuardianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorGuardianAdded represents a GuardianAdded event raised by the BatchAuthenticator contract.
type BatchAuthenticatorGuardianAdded struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdded is a free log retrieval operation binding the contract event 0x038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969.
//
// Solidity: event GuardianAdded(address indexed guardian)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterGuardianAdded(opts *bind.FilterOpts, guardian []common.Address) (*BatchAuthenticatorGuardianAddedIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "GuardianAdded", guardianRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorGuardianAddedIterator{contract: _BatchAuthenticator.contract, event: "GuardianAdded", logs: logs, sub: sub}, nil
}

// WatchGuardianAdded is a free log subscription operation binding the contract event 0x038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969.
//
// Solidity: event GuardianAdded(address indexed guardian)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchGuardianAdded(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorGuardianAdded, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "GuardianAdded", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorGuardianAdded)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGuardianAdded is a log parse operation binding the contract event 0x038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969.
//
// Solidity: event GuardianAdded(address indexed guardian)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseGuardianAdded(log types.Log) (*BatchAuthenticatorGuardianAdded, error) {
	event := new(BatchAuthenticatorGuardianAdded)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorGuardianRemovedIterator is returned from FilterGuardianRemoved and is used to iterate over the raw logs and unpacked data for GuardianRemoved events raised by the BatchAuthenticator contract.
type BatchAuthenticatorGuardianRemovedIterator struct {
	Event *BatchAuthenticatorGuardianRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorGuardianRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorGuardianRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorGuardianRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorGuardianRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorGuardianRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorGuardianRemoved represents a GuardianRemoved event raised by the BatchAuthenticator contract.
type BatchAuthenticatorGuardianRemoved struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianRemoved is a free log retrieval operation binding the contract event 0xb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52.
//
// Solidity: event GuardianRemoved(address indexed guardian)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterGuardianRemoved(opts *bind.FilterOpts, guardian []common.Address) (*BatchAuthenticatorGuardianRemovedIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "GuardianRemoved", guardianRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorGuardianRemovedIterator{contract: _BatchAuthenticator.contract, event: "GuardianRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardianRemoved is a free log subscription operation binding the contract event 0xb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52.
//
// Solidity: event GuardianRemoved(address indexed guardian)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchGuardianRemoved(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorGuardianRemoved, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "GuardianRemoved", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorGuardianRemoved)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGuardianRemoved is a log parse operation binding the contract event 0xb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52.
//
// Solidity: event GuardianRemoved(address indexed guardian)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseGuardianRemoved(log types.Log) (*BatchAuthenticatorGuardianRemoved, error) {
	event := new(BatchAuthenticatorGuardianRemoved)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BatchAuthenticator contract.
type BatchAuthenticatorInitializedIterator struct {
	Event *BatchAuthenticatorInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorInitialized represents a Initialized event raised by the BatchAuthenticator contract.
type BatchAuthenticatorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*BatchAuthenticatorInitializedIterator, error) {

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorInitializedIterator{contract: _BatchAuthenticator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorInitialized) (event.Subscription, error) {

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorInitialized)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseInitialized(log types.Log) (*BatchAuthenticatorInitialized, error) {
	event := new(BatchAuthenticatorInitialized)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the BatchAuthenticator contract.
type BatchAuthenticatorOwnershipTransferStartedIterator struct {
	Event *BatchAuthenticatorOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the BatchAuthenticator contract.
type BatchAuthenticatorOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchAuthenticatorOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorOwnershipTransferStartedIterator{contract: _BatchAuthenticator.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorOwnershipTransferStarted)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseOwnershipTransferStarted(log types.Log) (*BatchAuthenticatorOwnershipTransferStarted, error) {
	event := new(BatchAuthenticatorOwnershipTransferStarted)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BatchAuthenticator contract.
type BatchAuthenticatorOwnershipTransferredIterator struct {
	Event *BatchAuthenticatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorOwnershipTransferred represents a OwnershipTransferred event raised by the BatchAuthenticator contract.
type BatchAuthenticatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchAuthenticatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorOwnershipTransferredIterator{contract: _BatchAuthenticator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorOwnershipTransferred)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseOwnershipTransferred(log types.Log) (*BatchAuthenticatorOwnershipTransferred, error) {
	event := new(BatchAuthenticatorOwnershipTransferred)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorSignerRegistrationInitiatedIterator is returned from FilterSignerRegistrationInitiated and is used to iterate over the raw logs and unpacked data for SignerRegistrationInitiated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorSignerRegistrationInitiatedIterator struct {
	Event *BatchAuthenticatorSignerRegistrationInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchAuthenticatorSignerRegistrationInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorSignerRegistrationInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchAuthenticatorSignerRegistrationInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchAuthenticatorSignerRegistrationInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorSignerRegistrationInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorSignerRegistrationInitiated represents a SignerRegistrationInitiated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorSignerRegistrationInitiated struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerRegistrationInitiated is a free log retrieval operation binding the contract event 0x665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c.
//
// Solidity: event SignerRegistrationInitiated(address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterSignerRegistrationInitiated(opts *bind.FilterOpts, caller []common.Address) (*BatchAuthenticatorSignerRegistrationInitiatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "SignerRegistrationInitiated", callerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorSignerRegistrationInitiatedIterator{contract: _BatchAuthenticator.contract, event: "SignerRegistrationInitiated", logs: logs, sub: sub}, nil
}

// WatchSignerRegistrationInitiated is a free log subscription operation binding the contract event 0x665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c.
//
// Solidity: event SignerRegistrationInitiated(address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchSignerRegistrationInitiated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorSignerRegistrationInitiated, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "SignerRegistrationInitiated", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorSignerRegistrationInitiated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "SignerRegistrationInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerRegistrationInitiated is a log parse operation binding the contract event 0x665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c.
//
// Solidity: event SignerRegistrationInitiated(address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseSignerRegistrationInitiated(log types.Log) (*BatchAuthenticatorSignerRegistrationInitiated, error) {
	event := new(BatchAuthenticatorSignerRegistrationInitiated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "SignerRegistrationInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
