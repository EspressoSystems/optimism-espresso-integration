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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GUARDIAN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGuardians\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"guardianCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isGuardian\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProxyAdmin\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdminOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNonTeeBatcher\",\"inputs\":[{\"name\":\"_newNonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTeeBatcher\",\"inputs\":[{\"name\":\"_newTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BatchInfoAuthenticated\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatcherSwitched\",\"inputs\":[{\"name\":\"activeIsTee\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianAdded\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianRemoved\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonTeeBatcherUpdated\",\"inputs\":[{\"name\":\"oldNonTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newNonTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerRegistrationInitiated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeeBatcherUpdated\",\"inputs\":[{\"name\":\"oldTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"contract_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidGuardianAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotGuardian\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotGuardianOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOrProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotResolvedDelegateProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotSharedProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_ProxyAdminNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReinitializableBase_ZeroInitVersion\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50600160805261001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051612d426100f35f395f8181610375015261135d0152612d425ff3fe608060405234801561000f575f5ffd5b506004361061024f575f3560e01c80638da5cb5b1161013d578063ca15c873116100b8578063e30c397811610088578063f8c8765e1161006e578063f8c8765e146105b5578063fa14fe6d146105c8578063fc619e41146105e8575f5ffd5b8063e30c39781461059a578063f2fde38b146105a2575f5ffd5b8063ca15c8731461054d578063d547741f14610560578063d909ba7c14610573578063dad544e014610592575f5ffd5b8063a3246ad31161010d578063b1bd4285116100f3578063b1bd428514610512578063ba58e82a14610532578063bc347f4714610545575f5ffd5b8063a3246ad3146104ec578063a526d83b146104ff575f5ffd5b80638da5cb5b146104665780639010d07c1461046e57806391d1485414610481578063a217fddf146104e5575f5ffd5b80633e47158c116101cd578063714041561161019d5780637877a9ed116101835780637877a9ed1461042657806379ba50971461044b5780638c8fc53114610453575f5ffd5b8063714041561461040b578063715018a61461041e575f5ffd5b80633e47158c1461039f57806354387ad7146103a757806354fd4d50146103af5780636f7eda47146103f8575f5ffd5b8063248a9ca3116102225780632f2ff15d116102085780632f2ff15d1461034657806336568abe1461035b57806338d38c971461036e575f5ffd5b8063248a9ca3146102d057806324ea54f41461031f575f5ffd5b806301ffc9a7146102535780630665f04b1461027b5780630c68ba21146102905780631b076a4c146102a3575b5f5ffd5b610266610261366004612800565b6105fb565b60405190151581526020015b60405180910390f35b610283610656565b604051610272919061283f565b61026661029e3660046128b8565b610744565b6102ab610790565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610272565b6103116102de3660046128d3565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610272565b6103117f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a504181565b6103596103543660046128ea565b610826565b005b6103596103693660046128ea565b61086f565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610272565b6102ab6108cd565b610311610ad3565b6103eb6040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516102729190612918565b6103596104063660046128b8565b610afd565b6103596104193660046128b8565b610be3565b610359610ca5565b6002546102669074010000000000000000000000000000000000000000900460ff1681565b610359610cb8565b6103596104613660046128b8565b610d30565b6102ab610e13565b6102ab61047c36600461296b565b610e1c565b61026661048f3660046128ea565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6103115f81565b6102836104fa3660046128d3565b610e5c565b61035961050d3660046128b8565b610e9f565b6001546102ab9073ffffffffffffffffffffffffffffffffffffffff1681565b6103596105403660046129c9565b610fac565b61035961106a565b61031161055b3660046128d3565b611198565b61035961056e3660046128ea565b6111cf565b5f546102ab9073ffffffffffffffffffffffffffffffffffffffff1681565b6102ab611212565b6102ab611263565b6103596105b03660046128b8565b6112a4565b6103596105c3366004612a35565b61135b565b6002546102ab9073ffffffffffffffffffffffffffffffffffffffff1681565b6103596105f6366004612a8e565b61166b565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f000000000000000000000000000000000000000000000000000000001480610650575061065082611a25565b92915050565b60605f6106827f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a5041611198565b90505f8167ffffffffffffffff81111561069e5761069e612ad6565b6040519080825280602002602001820160405280156106c7578160200160208202803683370190505b5090505f5b8281101561073d576106fe7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a504182610e1c565b82828151811061071057610710612b03565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526001016106cc565b5092915050565b73ffffffffffffffffffffffffffffffffffffffff81165f9081527f18476f5b3d6d00091ddd56161ac5e9ba807d29b59f48f8df98938ee352a7cf23602052604081205460ff16610650565b600254604080517fd80a4c2800000000000000000000000000000000000000000000000000000000815290515f9273ffffffffffffffffffffffffffffffffffffffff169163d80a4c289160048083019260209291908290030181865afa1580156107fd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108219190612b30565b905090565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461085f81611abb565b6108698383611ac5565b50505050565b73ffffffffffffffffffffffffffffffffffffffff811633146108be576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108c88282611b1a565b505050565b5f806108f77fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b905073ffffffffffffffffffffffffffffffffffffffff81161561091a57919050565b6040518060400160405280601a81526020017f4f564d5f4c3143726f7373446f6d61696e4d657373656e67657200000000000081525051600261095d9190612b78565b604080513060208201525f918101919091527f4f564d5f4c3143726f7373446f6d61696e4d657373656e67657200000000000091909117906109b7906060015b604051602081830303815290604052805190602001205490565b146109ee576040517f54e433cd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080513060208201526001918101919091525f90610a0f9060600161099d565b905073ffffffffffffffffffffffffffffffffffffffff811615610aa1578073ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a76573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a9a9190612b30565b9250505090565b6040517f332144db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6108217f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a5041611198565b610b05611b66565b73ffffffffffffffffffffffffffffffffffffffff8116610b6f576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b33279190a35050565b610beb611b66565b73ffffffffffffffffffffffffffffffffffffffff81165f9081527f18476f5b3d6d00091ddd56161ac5e9ba807d29b59f48f8df98938ee352a7cf23602052604090205460ff1615610ca257610c617f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a5041826111cf565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52905f90a25b50565b610cad611b66565b610cb65f611bbe565b565b3380610cc2611263565b73ffffffffffffffffffffffffffffffffffffffff1614610d27576040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610b66565b610ca281611bbe565b610d38611b66565b73ffffffffffffffffffffffffffffffffffffffff8116610d9d576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610b66565b6001805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5905f90a35050565b5f610821611c04565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220610e549084611c2c565b949350505050565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604090912060609190610e9890611c37565b9392505050565b610ea7611b66565b73ffffffffffffffffffffffffffffffffffffffff8116610ef4576040517f1b08105400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f9081527f18476f5b3d6d00091ddd56161ac5e9ba807d29b59f48f8df98938ee352a7cf23602052604090205460ff16610ca257610f697f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a504182610826565b60405173ffffffffffffffffffffffffffffffffffffffff8216907f038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969905f90a250565b6002546040517f7f82ea6c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690637f82ea6c9061100d9087908790879087906001905f90600401612c38565b5f604051808303815f87803b158015611024575f5ffd5b505af1158015611036573d5f5f3e3d5ffd5b50506040513392507f665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c91505f90a250505050565b335f9081527f18476f5b3d6d00091ddd56161ac5e9ba807d29b59f48f8df98938ee352a7cf23602052604090205460ff161580156110db57506110ab610e13565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15611114576040517fd53780c4000000000000000000000000000000000000000000000000000000008152336004820152602401610b66565b6002805460ff7401000000000000000000000000000000000000000080830482161581027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9093169290921792839055604051919092049091161515907fb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad3905f90a2565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220610e9890611c43565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461120881611abb565b6108698383611b1a565b5f61121b6108cd565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107fd573d5f5f3e3d5ffd5b5f807f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005b5473ffffffffffffffffffffffffffffffffffffffff1692915050565b6112ac611b66565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081178255611315610e13565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b7f000000000000000000000000000000000000000000000000000000000000000060ff165f611388611c4c565b805490915068010000000000000000900460ff16806113b55750805467ffffffffffffffff808416911610155b156113ec576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001667ffffffffffffffff83161768010000000000000000178155611431611c74565b61143a83611cf5565b73ffffffffffffffffffffffffffffffffffffffff851661149f576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610b66565b73ffffffffffffffffffffffffffffffffffffffff8416611504576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610b66565b73ffffffffffffffffffffffffffffffffffffffff8616611569576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87166004820152602401610b66565b600280545f805473ffffffffffffffffffffffffffffffffffffffff8981167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556001805489841692169190911790557fffffffffffffffffffffff000000000000000000000000000000000000000000909116908816177401000000000000000000000000000000000000000017905580547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050565b5f82828080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505082519293505060419091149050611712576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f496e76616c6964207369676e6174757265206c656e67746800000000000000006044820152606401610b66565b5f8160408151811061172657611726612b03565b016020015160f81c905080158061174057508060ff166001145b1561179857611750601b82612c8a565b90508060f81b8260408151811061176957611769612b03565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b5f6117a38684611d1f565b905073ffffffffffffffffffffffffffffffffffffffff8116611848576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f426174636841757468656e74696361746f723a20696e76616c6964207369676e60448201527f61747572650000000000000000000000000000000000000000000000000000006064820152608401610b66565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d80a4c286040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118b2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118d69190612b30565b73ffffffffffffffffffffffffffffffffffffffff16636d8f5aa9825f6040518363ffffffff1660e01b8152600401611910929190612ca3565b602060405180830381865afa15801561192b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061194f9190612cd6565b6119db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f426174636841757468656e74696361746f723a20696e76616c6964207369676e60448201527f65720000000000000000000000000000000000000000000000000000000000006064820152608401610b66565b60405173ffffffffffffffffffffffffffffffffffffffff82169087907f731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c905f90a3505050505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061065057507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610650565b610ca28133611d41565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081611af28585611deb565b90508015610e54575f858152602083905260409020611b119085611f09565b50949350505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081611b478585611f2a565b90508015610e54575f858152602083905260409020611b119085612006565b33611b6f610e13565b73ffffffffffffffffffffffffffffffffffffffff1614610cb6576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610b66565b5f611bc7610e13565b9050611bd282612027565b73ffffffffffffffffffffffffffffffffffffffff811615611bfa57611bf85f82611b1a565b505b6108c85f83611ac5565b5f807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300611287565b5f610e988383612077565b60605f610e988361209d565b5f610650825490565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610650565b33611c7d6108cd565b73ffffffffffffffffffffffffffffffffffffffff1614158015611cbe575033611ca5611212565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610cb6576040517fc4050a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611cfd6120f6565b611d0681612134565b611d0e612145565b611d16612145565b610ca28161214d565b5f5f5f611d2c858561218a565b91509150611d39816121cc565b509392505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611de7576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610b66565b5050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611f00575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611e9c3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610650565b5f915050610650565b5f610e988373ffffffffffffffffffffffffffffffffffffffff841661241f565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611f00575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610650565b5f610e988373ffffffffffffffffffffffffffffffffffffffff841661246b565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155611de782612545565b5f825f01828154811061208c5761208c612b03565b905f5260205f200154905092915050565b6060815f018054806020026020016040519081016040528092919081815260200182805480156120ea57602002820191905f5260205f20905b8154815260200190600101908083116120d6575b50505050509050919050565b6120fe6125da565b610cb6576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61213c6120f6565b610ca2816125f8565b610cb66120f6565b6121556120f6565b61215f5f82611ac5565b50610ca27f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50415f61264f565b5f5f82516041036121be576020830151604084015160608501515f1a6121b2878285856126f0565b945094505050506121c5565b505f905060025b9250929050565b5f8160048111156121df576121df612bd6565b036121e75750565b60018160048111156121fb576121fb612bd6565b03612262576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610b66565b600281600481111561227657612276612bd6565b036122dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610b66565b60038160048111156122f1576122f1612bd6565b0361237e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610b66565b600481600481111561239257612392612bd6565b03610ca2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610b66565b5f81815260018301602052604081205461246457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610650565b505f610650565b5f8181526001830160205260408120548015611f00575f61248d600183612cf5565b85549091505f906124a090600190612cf5565b90508082146124ff575f865f0182815481106124be576124be612b03565b905f5260205f200154905080875f0184815481106124de576124de612b03565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061251057612510612d08565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610650565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f6125e3611c4c565b5468010000000000000000900460ff16919050565b6126006120f6565b73ffffffffffffffffffffffffffffffffffffffff8116610d27576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610b66565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f6126a8845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561272557505f905060036127f7565b8460ff16601b1415801561273d57508460ff16601c14155b1561274d57505f905060046127f7565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561279e573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166127f1575f600192509250506127f7565b91505f90505b94509492505050565b5f60208284031215612810575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610e98575f5ffd5b602080825282518282018190525f918401906040840190835b8181101561288c57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612858565b509095945050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610ca2575f5ffd5b5f602082840312156128c8575f5ffd5b8135610e9881612897565b5f602082840312156128e3575f5ffd5b5035919050565b5f5f604083850312156128fb575f5ffd5b82359150602083013561290d81612897565b809150509250929050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f6040838503121561297c575f5ffd5b50508035926020909101359150565b5f5f83601f84011261299b575f5ffd5b50813567ffffffffffffffff8111156129b2575f5ffd5b6020830191508360208285010111156121c5575f5ffd5b5f5f5f5f604085870312156129dc575f5ffd5b843567ffffffffffffffff8111156129f2575f5ffd5b6129fe8782880161298b565b909550935050602085013567ffffffffffffffff811115612a1d575f5ffd5b612a298782880161298b565b95989497509550505050565b5f5f5f5f60808587031215612a48575f5ffd5b8435612a5381612897565b93506020850135612a6381612897565b92506040850135612a7381612897565b91506060850135612a8381612897565b939692955090935050565b5f5f5f60408486031215612aa0575f5ffd5b83359250602084013567ffffffffffffffff811115612abd575f5ffd5b612ac98682870161298b565b9497909650939450505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215612b40575f5ffd5b8151610e9881612897565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808202811582820484141761065057610650612b4b565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60028110610ca2577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b608081525f612c4b60808301888a612b8f565b8281036020840152612c5e818789612b8f565b915050612c6a84612c03565b836040830152612c7983612c03565b826060830152979650505050505050565b60ff818116838216019081111561065057610650612b4b565b73ffffffffffffffffffffffffffffffffffffffff8316815260408101612cc983612c03565b8260208301529392505050565b5f60208284031215612ce6575f5ffd5b81518015158114610e98575f5ffd5b8181038181111561065057610650612b4b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea164736f6c634300081c000a",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BatchAuthenticator.Contract.DEFAULTADMINROLE(&_BatchAuthenticator.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BatchAuthenticator.Contract.DEFAULTADMINROLE(&_BatchAuthenticator.CallOpts)
}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorCaller) GUARDIANROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "GUARDIAN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorSession) GUARDIANROLE() ([32]byte, error) {
	return _BatchAuthenticator.Contract.GUARDIANROLE(&_BatchAuthenticator.CallOpts)
}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) GUARDIANROLE() ([32]byte, error) {
	return _BatchAuthenticator.Contract.GUARDIANROLE(&_BatchAuthenticator.CallOpts)
}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ActiveIsTee(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "activeIsTee")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) ActiveIsTee() (bool, error) {
	return _BatchAuthenticator.Contract.ActiveIsTee(&_BatchAuthenticator.CallOpts)
}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ActiveIsTee() (bool, error) {
	return _BatchAuthenticator.Contract.ActiveIsTee(&_BatchAuthenticator.CallOpts)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BatchAuthenticator.Contract.GetRoleAdmin(&_BatchAuthenticator.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BatchAuthenticator.Contract.GetRoleAdmin(&_BatchAuthenticator.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BatchAuthenticator.Contract.GetRoleMember(&_BatchAuthenticator.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BatchAuthenticator.Contract.GetRoleMember(&_BatchAuthenticator.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BatchAuthenticator.Contract.GetRoleMemberCount(&_BatchAuthenticator.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BatchAuthenticator.Contract.GetRoleMemberCount(&_BatchAuthenticator.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BatchAuthenticator *BatchAuthenticatorCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BatchAuthenticator *BatchAuthenticatorSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BatchAuthenticator.Contract.GetRoleMembers(&_BatchAuthenticator.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BatchAuthenticator.Contract.GetRoleMembers(&_BatchAuthenticator.CallOpts, role)
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

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BatchAuthenticator.Contract.HasRole(&_BatchAuthenticator.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BatchAuthenticator.Contract.HasRole(&_BatchAuthenticator.CallOpts, role, account)
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

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) NonTeeBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "nonTeeBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) NonTeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.NonTeeBatcher(&_BatchAuthenticator.CallOpts)
}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) NonTeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.NonTeeBatcher(&_BatchAuthenticator.CallOpts)
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BatchAuthenticator.Contract.SupportsInterface(&_BatchAuthenticator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BatchAuthenticator.Contract.SupportsInterface(&_BatchAuthenticator.CallOpts, interfaceId)
}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) TeeBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "teeBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) TeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.TeeBatcher(&_BatchAuthenticator.CallOpts)
}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) TeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.TeeBatcher(&_BatchAuthenticator.CallOpts)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.GrantRole(&_BatchAuthenticator.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.GrantRole(&_BatchAuthenticator.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _teeBatcher, address _nonTeeBatcher, address _owner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) Initialize(opts *bind.TransactOpts, _espressoTEEVerifier common.Address, _teeBatcher common.Address, _nonTeeBatcher common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "initialize", _espressoTEEVerifier, _teeBatcher, _nonTeeBatcher, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _teeBatcher, address _nonTeeBatcher, address _owner) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) Initialize(_espressoTEEVerifier common.Address, _teeBatcher common.Address, _nonTeeBatcher common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _teeBatcher, _nonTeeBatcher, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _teeBatcher, address _nonTeeBatcher, address _owner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) Initialize(_espressoTEEVerifier common.Address, _teeBatcher common.Address, _nonTeeBatcher common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _teeBatcher, _nonTeeBatcher, _owner)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RegisterSigner(opts *bind.TransactOpts, attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "registerSigner", attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, attestationTbs, signature)
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

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RenounceRole(&_BatchAuthenticator.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RenounceRole(&_BatchAuthenticator.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RevokeRole(&_BatchAuthenticator.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RevokeRole(&_BatchAuthenticator.TransactOpts, role, account)
}

// SetNonTeeBatcher is a paid mutator transaction binding the contract method 0x8c8fc531.
//
// Solidity: function setNonTeeBatcher(address _newNonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SetNonTeeBatcher(opts *bind.TransactOpts, _newNonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "setNonTeeBatcher", _newNonTeeBatcher)
}

// SetNonTeeBatcher is a paid mutator transaction binding the contract method 0x8c8fc531.
//
// Solidity: function setNonTeeBatcher(address _newNonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SetNonTeeBatcher(_newNonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetNonTeeBatcher(&_BatchAuthenticator.TransactOpts, _newNonTeeBatcher)
}

// SetNonTeeBatcher is a paid mutator transaction binding the contract method 0x8c8fc531.
//
// Solidity: function setNonTeeBatcher(address _newNonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SetNonTeeBatcher(_newNonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetNonTeeBatcher(&_BatchAuthenticator.TransactOpts, _newNonTeeBatcher)
}

// SetTeeBatcher is a paid mutator transaction binding the contract method 0x6f7eda47.
//
// Solidity: function setTeeBatcher(address _newTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SetTeeBatcher(opts *bind.TransactOpts, _newTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "setTeeBatcher", _newTeeBatcher)
}

// SetTeeBatcher is a paid mutator transaction binding the contract method 0x6f7eda47.
//
// Solidity: function setTeeBatcher(address _newTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SetTeeBatcher(_newTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetTeeBatcher(&_BatchAuthenticator.TransactOpts, _newTeeBatcher)
}

// SetTeeBatcher is a paid mutator transaction binding the contract method 0x6f7eda47.
//
// Solidity: function setTeeBatcher(address _newTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SetTeeBatcher(_newTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetTeeBatcher(&_BatchAuthenticator.TransactOpts, _newTeeBatcher)
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
	Signer     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchInfoAuthenticated is a free log retrieval operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterBatchInfoAuthenticated(opts *bind.FilterOpts, commitment [][32]byte, signer []common.Address) (*BatchAuthenticatorBatchInfoAuthenticatedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "BatchInfoAuthenticated", commitmentRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorBatchInfoAuthenticatedIterator{contract: _BatchAuthenticator.contract, event: "BatchInfoAuthenticated", logs: logs, sub: sub}, nil
}

// WatchBatchInfoAuthenticated is a free log subscription operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchBatchInfoAuthenticated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorBatchInfoAuthenticated, commitment [][32]byte, signer []common.Address) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "BatchInfoAuthenticated", commitmentRule, signerRule)
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

// ParseBatchInfoAuthenticated is a log parse operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer)
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
	ActiveIsTee bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatcherSwitched is a free log retrieval operation binding the contract event 0xb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad3.
//
// Solidity: event BatcherSwitched(bool indexed activeIsTee)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterBatcherSwitched(opts *bind.FilterOpts, activeIsTee []bool) (*BatchAuthenticatorBatcherSwitchedIterator, error) {

	var activeIsTeeRule []interface{}
	for _, activeIsTeeItem := range activeIsTee {
		activeIsTeeRule = append(activeIsTeeRule, activeIsTeeItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "BatcherSwitched", activeIsTeeRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorBatcherSwitchedIterator{contract: _BatchAuthenticator.contract, event: "BatcherSwitched", logs: logs, sub: sub}, nil
}

// WatchBatcherSwitched is a free log subscription operation binding the contract event 0xb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad3.
//
// Solidity: event BatcherSwitched(bool indexed activeIsTee)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchBatcherSwitched(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorBatcherSwitched, activeIsTee []bool) (event.Subscription, error) {

	var activeIsTeeRule []interface{}
	for _, activeIsTeeItem := range activeIsTee {
		activeIsTeeRule = append(activeIsTeeRule, activeIsTeeItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "BatcherSwitched", activeIsTeeRule)
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
// Solidity: event BatcherSwitched(bool indexed activeIsTee)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseBatcherSwitched(log types.Log) (*BatchAuthenticatorBatcherSwitched, error) {
	event := new(BatchAuthenticatorBatcherSwitched)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "BatcherSwitched", log); err != nil {
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

// BatchAuthenticatorNonTeeBatcherUpdatedIterator is returned from FilterNonTeeBatcherUpdated and is used to iterate over the raw logs and unpacked data for NonTeeBatcherUpdated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorNonTeeBatcherUpdatedIterator struct {
	Event *BatchAuthenticatorNonTeeBatcherUpdated // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorNonTeeBatcherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorNonTeeBatcherUpdated)
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
		it.Event = new(BatchAuthenticatorNonTeeBatcherUpdated)
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
func (it *BatchAuthenticatorNonTeeBatcherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorNonTeeBatcherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorNonTeeBatcherUpdated represents a NonTeeBatcherUpdated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorNonTeeBatcherUpdated struct {
	OldNonTeeBatcher common.Address
	NewNonTeeBatcher common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonTeeBatcherUpdated is a free log retrieval operation binding the contract event 0xc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5.
//
// Solidity: event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterNonTeeBatcherUpdated(opts *bind.FilterOpts, oldNonTeeBatcher []common.Address, newNonTeeBatcher []common.Address) (*BatchAuthenticatorNonTeeBatcherUpdatedIterator, error) {

	var oldNonTeeBatcherRule []interface{}
	for _, oldNonTeeBatcherItem := range oldNonTeeBatcher {
		oldNonTeeBatcherRule = append(oldNonTeeBatcherRule, oldNonTeeBatcherItem)
	}
	var newNonTeeBatcherRule []interface{}
	for _, newNonTeeBatcherItem := range newNonTeeBatcher {
		newNonTeeBatcherRule = append(newNonTeeBatcherRule, newNonTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "NonTeeBatcherUpdated", oldNonTeeBatcherRule, newNonTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorNonTeeBatcherUpdatedIterator{contract: _BatchAuthenticator.contract, event: "NonTeeBatcherUpdated", logs: logs, sub: sub}, nil
}

// WatchNonTeeBatcherUpdated is a free log subscription operation binding the contract event 0xc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5.
//
// Solidity: event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchNonTeeBatcherUpdated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorNonTeeBatcherUpdated, oldNonTeeBatcher []common.Address, newNonTeeBatcher []common.Address) (event.Subscription, error) {

	var oldNonTeeBatcherRule []interface{}
	for _, oldNonTeeBatcherItem := range oldNonTeeBatcher {
		oldNonTeeBatcherRule = append(oldNonTeeBatcherRule, oldNonTeeBatcherItem)
	}
	var newNonTeeBatcherRule []interface{}
	for _, newNonTeeBatcherItem := range newNonTeeBatcher {
		newNonTeeBatcherRule = append(newNonTeeBatcherRule, newNonTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "NonTeeBatcherUpdated", oldNonTeeBatcherRule, newNonTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorNonTeeBatcherUpdated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "NonTeeBatcherUpdated", log); err != nil {
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

// ParseNonTeeBatcherUpdated is a log parse operation binding the contract event 0xc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5.
//
// Solidity: event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseNonTeeBatcherUpdated(log types.Log) (*BatchAuthenticatorNonTeeBatcherUpdated, error) {
	event := new(BatchAuthenticatorNonTeeBatcherUpdated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "NonTeeBatcherUpdated", log); err != nil {
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

// BatchAuthenticatorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BatchAuthenticator contract.
type BatchAuthenticatorRoleAdminChangedIterator struct {
	Event *BatchAuthenticatorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorRoleAdminChanged)
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
		it.Event = new(BatchAuthenticatorRoleAdminChanged)
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
func (it *BatchAuthenticatorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorRoleAdminChanged represents a RoleAdminChanged event raised by the BatchAuthenticator contract.
type BatchAuthenticatorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BatchAuthenticatorRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorRoleAdminChangedIterator{contract: _BatchAuthenticator.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorRoleAdminChanged)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseRoleAdminChanged(log types.Log) (*BatchAuthenticatorRoleAdminChanged, error) {
	event := new(BatchAuthenticatorRoleAdminChanged)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BatchAuthenticator contract.
type BatchAuthenticatorRoleGrantedIterator struct {
	Event *BatchAuthenticatorRoleGranted // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorRoleGranted)
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
		it.Event = new(BatchAuthenticatorRoleGranted)
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
func (it *BatchAuthenticatorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorRoleGranted represents a RoleGranted event raised by the BatchAuthenticator contract.
type BatchAuthenticatorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BatchAuthenticatorRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorRoleGrantedIterator{contract: _BatchAuthenticator.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorRoleGranted)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseRoleGranted(log types.Log) (*BatchAuthenticatorRoleGranted, error) {
	event := new(BatchAuthenticatorRoleGranted)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BatchAuthenticator contract.
type BatchAuthenticatorRoleRevokedIterator struct {
	Event *BatchAuthenticatorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorRoleRevoked)
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
		it.Event = new(BatchAuthenticatorRoleRevoked)
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
func (it *BatchAuthenticatorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorRoleRevoked represents a RoleRevoked event raised by the BatchAuthenticator contract.
type BatchAuthenticatorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BatchAuthenticatorRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorRoleRevokedIterator{contract: _BatchAuthenticator.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorRoleRevoked)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseRoleRevoked(log types.Log) (*BatchAuthenticatorRoleRevoked, error) {
	event := new(BatchAuthenticatorRoleRevoked)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// BatchAuthenticatorTeeBatcherUpdatedIterator is returned from FilterTeeBatcherUpdated and is used to iterate over the raw logs and unpacked data for TeeBatcherUpdated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorTeeBatcherUpdatedIterator struct {
	Event *BatchAuthenticatorTeeBatcherUpdated // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorTeeBatcherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorTeeBatcherUpdated)
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
		it.Event = new(BatchAuthenticatorTeeBatcherUpdated)
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
func (it *BatchAuthenticatorTeeBatcherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorTeeBatcherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorTeeBatcherUpdated represents a TeeBatcherUpdated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorTeeBatcherUpdated struct {
	OldTeeBatcher common.Address
	NewTeeBatcher common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTeeBatcherUpdated is a free log retrieval operation binding the contract event 0x5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b3327.
//
// Solidity: event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterTeeBatcherUpdated(opts *bind.FilterOpts, oldTeeBatcher []common.Address, newTeeBatcher []common.Address) (*BatchAuthenticatorTeeBatcherUpdatedIterator, error) {

	var oldTeeBatcherRule []interface{}
	for _, oldTeeBatcherItem := range oldTeeBatcher {
		oldTeeBatcherRule = append(oldTeeBatcherRule, oldTeeBatcherItem)
	}
	var newTeeBatcherRule []interface{}
	for _, newTeeBatcherItem := range newTeeBatcher {
		newTeeBatcherRule = append(newTeeBatcherRule, newTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "TeeBatcherUpdated", oldTeeBatcherRule, newTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorTeeBatcherUpdatedIterator{contract: _BatchAuthenticator.contract, event: "TeeBatcherUpdated", logs: logs, sub: sub}, nil
}

// WatchTeeBatcherUpdated is a free log subscription operation binding the contract event 0x5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b3327.
//
// Solidity: event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchTeeBatcherUpdated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorTeeBatcherUpdated, oldTeeBatcher []common.Address, newTeeBatcher []common.Address) (event.Subscription, error) {

	var oldTeeBatcherRule []interface{}
	for _, oldTeeBatcherItem := range oldTeeBatcher {
		oldTeeBatcherRule = append(oldTeeBatcherRule, oldTeeBatcherItem)
	}
	var newTeeBatcherRule []interface{}
	for _, newTeeBatcherItem := range newTeeBatcher {
		newTeeBatcherRule = append(newTeeBatcherRule, newTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "TeeBatcherUpdated", oldTeeBatcherRule, newTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorTeeBatcherUpdated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "TeeBatcherUpdated", log); err != nil {
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

// ParseTeeBatcherUpdated is a log parse operation binding the contract event 0x5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b3327.
//
// Solidity: event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseTeeBatcherUpdated(log types.Log) (*BatchAuthenticatorTeeBatcherUpdated, error) {
	event := new(BatchAuthenticatorTeeBatcherUpdated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "TeeBatcherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
