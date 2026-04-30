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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsEspresso\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"_commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"espressoBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGuardians\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guardianCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_espressoBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"contractISystemConfig\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isGuardian\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProxyAdmin\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdminOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"_verificationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEspressoBatcher\",\"inputs\":[{\"name\":\"_newEspressoBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISystemConfig\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BatchInfoAuthenticated\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatcherSwitched\",\"inputs\":[{\"name\":\"activeIsEspresso\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EspressoBatcherUpdated\",\"inputs\":[{\"name\":\"oldEspressoBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEspressoBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianAdded\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianRemoved\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerRegistrationInitiated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BatchAuthenticator_Paused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"contract_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidGuardianAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotGuardian\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotGuardianOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnerCantBeGuardian\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOrProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotResolvedDelegateProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotSharedProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_ProxyAdminNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReinitializableBase_ZeroInitVersion\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50600160805261001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051611c0f6100f35f395f81816102480152610e590152611c0f5ff3fe608060405234801561000f575f80fd5b506004361061019a575f3560e01c806379ba5097116100e8578063dad544e011610093578063f2fde38b1161006e578063f2fde38b1461038e578063f8c8765e146103a1578063fa14fe6d146103b4578063fc619e41146103d4575f80fd5b8063dad544e014610359578063e30c397814610361578063eca919df14610369575f80fd5b8063a526d83b116100c3578063a526d83b1461032b578063ba58e82a1461033e578063bc347f4714610351575f80fd5b806379ba5097146102fc57806388da3bb7146103045780638da5cb5b14610323575f80fd5b80633e47158c116101485780635c975abb116101235780635c975abb146102d957806371404156146102e1578063715018a6146102f4575f80fd5b80633e47158c1461027257806354387ad71461027a57806354fd4d5014610290575f80fd5b80632ce53247116101785780632ce532471461020c57806333d7e2bd1461022157806338d38c9714610241575f80fd5b80630665f04b1461019e5780630c68ba21146101bc5780631b076a4c146101df575b5f80fd5b6101a66103e7565b6040516101b391906117f1565b60405180910390f35b6101cf6101ca36600461186b565b610417565b60405190151581526020016101b3565b6101e7610472565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b3565b61021f61021a36600461186b565b610503565b005b6002546101e79073ffffffffffffffffffffffffffffffffffffffff1681565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101b3565b6101e76105e9565b6102826107ef565b6040519081526020016101b3565b6102cc6040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516101b39190611886565b6101cf610819565b61021f6102ef36600461186b565b6108aa565b61021f61092a565b61021f61093d565b5f546101e79073ffffffffffffffffffffffffffffffffffffffff1681565b6101e76109b5565b61021f61033936600461186b565b6109be565b61021f61034c36600461191e565b610b32565b61021f610c2c565b6101e7610d53565b6101e7610da4565b6001546101cf9074010000000000000000000000000000000000000000900460ff1681565b61021f61039c36600461186b565b610de5565b61021f6103af366004611985565b610e57565b6001546101e79073ffffffffffffffffffffffffffffffffffffffff1681565b61021f6103e23660046119de565b611167565b60606104127f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f00611271565b905090565b5f61046c827f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f005b9073ffffffffffffffffffffffffffffffffffffffff165f9081526001919091016020526040902054151590565b92915050565b600154604080517fd80a4c2800000000000000000000000000000000000000000000000000000000815290515f9273ffffffffffffffffffffffffffffffffffffffff169163d80a4c289160048083019260209291908290030181865afa1580156104df573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104129190611a26565b61050b611284565b73ffffffffffffffffffffffffffffffffffffffff8116610575576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917fc7bef7b97a10ef514a01fa4d5552f5c57e72a37aa901f567bb49bbf9ea449f9c9190a35050565b5f806106137fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b905073ffffffffffffffffffffffffffffffffffffffff81161561063657919050565b6040518060400160405280601a81526020017f4f564d5f4c3143726f7373446f6d61696e4d657373656e6765720000000000008152505160026106799190611a6e565b604080513060208201525f918101919091527f4f564d5f4c3143726f7373446f6d61696e4d657373656e67657200000000000091909117906106d3906060015b604051602081830303815290604052805190602001205490565b1461070a576040517f54e433cd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080513060208201526001918101919091525f9061072b906060016106b9565b905073ffffffffffffffffffffffffffffffffffffffff8116156107bd578073ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610792573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107b69190611a26565b9250505090565b6040517f332144db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6104127f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f006112dc565b600254604080517f5c975abb00000000000000000000000000000000000000000000000000000000815290515f9273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa158015610886573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104129190611a85565b6108b2611284565b7f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f006108dd81836112e5565b6108e5575050565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52905f90a2505b50565b610932611284565b61093b5f611306565b565b3380610947610da4565b73ffffffffffffffffffffffffffffffffffffffff16146109ac576040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260240161056c565b61092781611306565b5f610412611356565b6109c6611284565b73ffffffffffffffffffffffffffffffffffffffff8116610a13576040517f1b08105400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a1b6109b5565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610a865750610a57610da4565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15610abd576040517f3af3c41c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f00610ae8818361137e565b15610b2e5760405173ffffffffffffffffffffffffffffffffffffffff8316907f038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969905f90a25b5050565b610b3a610819565b15610b71576040517fb3a266c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517fdac79fc800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063dac79fc890610bcf9087908790879087905f90600401611b24565b5f604051808303815f87803b158015610be6575f80fd5b505af1158015610bf8573d5f803e3d5ffd5b50506040513392507f665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c91505f90a250505050565b610c56337f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f0061043e565b158015610c965750610c666109b5565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610ccf576040517fd53780c400000000000000000000000000000000000000000000000000000000815233600482015260240161056c565b6001805460ff7401000000000000000000000000000000000000000080830482161581027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9093169290921792839055604051919092049091161515907fb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad3905f90a2565b5f610d5c6105e9565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104df573d5f803e3d5ffd5b5f807f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005b5473ffffffffffffffffffffffffffffffffffffffff1692915050565b610ded611284565b610e17817f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f0061043e565b15610e4e576040517f3af3c41c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109278161139f565b7f000000000000000000000000000000000000000000000000000000000000000060ff165f610e84611456565b805490915068010000000000000000900460ff1680610eb15750805467ffffffffffffffff808416911610155b15610ee8576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001667ffffffffffffffff83161768010000000000000000178155610f2d61147e565b610f36836114ff565b73ffffffffffffffffffffffffffffffffffffffff8516610f9b576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260240161056c565b73ffffffffffffffffffffffffffffffffffffffff8416611000576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161056c565b73ffffffffffffffffffffffffffffffffffffffff8616611065576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8716600482015260240161056c565b600180545f805473ffffffffffffffffffffffffffffffffffffffff8981167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556002805489841692169190911790557fffffffffffffffffffffff000000000000000000000000000000000000000000909116908816177401000000000000000000000000000000000000000017905580547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050565b61116f610819565b156111a6576040517fb3a266c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517fa81d9c5c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063a81d9c5c90611202908590859088905f90600401611b64565b602060405180830381865afa15801561121d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112419190611a85565b5060405183907fee0d07d204d979d28885955e59a46f754c4db7378b7df1a95123525aac6e3f80905f90a2505050565b60605f61127d83611510565b9392505050565b3361128d6109b5565b73ffffffffffffffffffffffffffffffffffffffff161461093b576040517f118cdaa700000000000000000000000000000000000000000000000000000000815233600482015260240161056c565b5f61046c825490565b5f61127d8373ffffffffffffffffffffffffffffffffffffffff8416611569565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155610b2e8261164c565b5f807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300610dc8565b5f61127d8373ffffffffffffffffffffffffffffffffffffffff84166116e1565b6113a7611284565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831690811782556114106109b5565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061046c565b336114876105e9565b73ffffffffffffffffffffffffffffffffffffffff16141580156114c85750336114af610d53565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561093b576040517fc4050a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61150761172d565b6109278161176b565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561155d57602002820191905f5260205f20905b815481526020019060010190808311611549575b50505050509050919050565b5f8181526001830160205260408120548015611643575f61158b600183611b95565b85549091505f9061159e90600190611b95565b90508082146115fd575f865f0182815481106115bc576115bc611ba8565b905f5260205f200154905080875f0184815481106115dc576115dc611ba8565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061160e5761160e611bd5565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061046c565b5f91505061046c565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f81815260018301602052604081205461172657508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561046c565b505f61046c565b61173561177c565b61093b576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61177361172d565b6109278161179a565b5f611785611456565b5468010000000000000000900460ff16919050565b6117a261172d565b73ffffffffffffffffffffffffffffffffffffffff81166109ac576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f600482015260240161056c565b602080825282518282018190525f9190848201906040850190845b8181101561183e57835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161180c565b50909695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610927575f80fd5b5f6020828403121561187b575f80fd5b813561127d8161184a565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f8083601f8401126118e9575f80fd5b50813567ffffffffffffffff811115611900575f80fd5b602083019150836020828501011115611917575f80fd5b9250929050565b5f805f8060408587031215611931575f80fd5b843567ffffffffffffffff80821115611948575f80fd5b611954888389016118d9565b9096509450602087013591508082111561196c575f80fd5b50611979878288016118d9565b95989497509550505050565b5f805f8060808587031215611998575f80fd5b84356119a38161184a565b935060208501356119b38161184a565b925060408501356119c38161184a565b915060608501356119d38161184a565b939692955090935050565b5f805f604084860312156119f0575f80fd5b83359250602084013567ffffffffffffffff811115611a0d575f80fd5b611a19868287016118d9565b9497909650939450505050565b5f60208284031215611a36575f80fd5b815161127d8161184a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808202811582820484141761046c5761046c611a41565b5f60208284031215611a95575f80fd5b8151801515811461127d575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60018110611b20577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b606081525f611b37606083018789611aa4565b8281036020840152611b4a818688611aa4565b915050611b5a6040830184611aeb565b9695505050505050565b606081525f611b77606083018688611aa4565b9050836020830152611b8c6040830184611aeb565b95945050505050565b8181038181111561046c5761046c611a41565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea164736f6c6343000819000a",
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
// Solidity: function authenticateBatchInfo(bytes32 _commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) AuthenticateBatchInfo(opts *bind.TransactOpts, _commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "authenticateBatchInfo", _commitment, _signature)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 _commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) AuthenticateBatchInfo(_commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AuthenticateBatchInfo(&_BatchAuthenticator.TransactOpts, _commitment, _signature)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 _commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) AuthenticateBatchInfo(_commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AuthenticateBatchInfo(&_BatchAuthenticator.TransactOpts, _commitment, _signature)
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
// Solidity: function registerSigner(bytes _verificationData, bytes _data) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RegisterSigner(opts *bind.TransactOpts, _verificationData []byte, _data []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "registerSigner", _verificationData, _data)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes _verificationData, bytes _data) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RegisterSigner(_verificationData []byte, _data []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, _verificationData, _data)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes _verificationData, bytes _data) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RegisterSigner(_verificationData []byte, _data []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, _verificationData, _data)
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
