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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsEspresso\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"_commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"espressoBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoBatcherAt\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"batcher_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromBlock_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoBatcherAtBlock\",\"inputs\":[{\"name\":\"_l1Block\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoBatcherHistoryLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGuardians\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guardianCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_espressoBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"contractISystemConfig\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activeIsEspresso\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isGuardian\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProxyAdmin\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdminOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"_verificationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActiveIsEspresso\",\"inputs\":[{\"name\":\"_desired\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEspressoBatcher\",\"inputs\":[{\"name\":\"_newEspressoBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISystemConfig\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BatchInfoAuthenticated\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatcherSwitched\",\"inputs\":[{\"name\":\"activeIsEspresso\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EspressoBatcherUpdated\",\"inputs\":[{\"name\":\"oldEspressoBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEspressoBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromBlock\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianAdded\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianRemoved\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerRegistrationInitiated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BatcherChangedThisBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"contract_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidGuardianAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoChange\",\"inputs\":[{\"name\":\"batcher\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotGuardian\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotGuardianOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnerCantBeGuardian\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOrProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotResolvedDelegateProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotSharedProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_ProxyAdminNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReinitializableBase_ZeroInitVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedEspressoBatcher\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedFallbackBatcher\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50600160805261001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516126106100f35f395f818161025e0152610ec201526126105ff3fe608060405234801561000f575f5ffd5b50600436106101b0575f3560e01c806379ba5097116100f3578063e30c397811610093578063fa14fe6d1161006e578063fa14fe6d146103bd578063fc5b5fda146103dc578063fc619e41146103ef578063fd402af714610402575f5ffd5b8063e30c39781461037e578063eca919df14610386578063f2fde38b146103aa575f5ffd5b80638da5cb5b116100ce5780638da5cb5b14610348578063a526d83b14610350578063ba58e82a14610363578063dad544e014610376575f5ffd5b806379ba5097146103255780637d531a781461032d57806388da3bb714610340575f5ffd5b80633e47158c1161015e57806354fd4d501161013957806354fd4d50146102ae5780636c076871146102f7578063714041561461030a578063715018a61461031d575f5ffd5b80633e47158c146102885780634268ecaa1461029057806354387ad7146102a6575f5ffd5b80632ce532471161018e5780632ce532471461022257806333d7e2bd1461023757806338d38c9714610257575f5ffd5b80630665f04b146101b45780630c68ba21146101d25780631b076a4c146101f5575b5f5ffd5b6101bc61044a565b6040516101c991906120dd565b60405180910390f35b6101e56101e0366004612156565b61047a565b60405190151581526020016101c9565b6101fd6104d5565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c9565b610235610230366004612156565b610563565b005b6001546101fd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101c9565b6101fd6106e2565b6102986108e8565b6040519081526020016101c9565b6102986108f2565b6102ea6040518060400160405280600581526020017f312e322e3000000000000000000000000000000000000000000000000000000081525081565b6040516101c99190612171565b6102356103053660046121d1565b61091c565b610235610318366004612156565b610a5a565b610235610ad9565b610235610aec565b6101fd61033b3660046121ec565b610b64565b6101fd610b7a565b6101fd610b85565b61023561035e366004612156565b610b8e565b610235610371366004612258565b610d02565b6101fd610dbc565b6101fd610e0d565b5f546101e59074010000000000000000000000000000000000000000900460ff1681565b6102356103b8366004612156565b610e4e565b5f546101fd9073ffffffffffffffffffffffffffffffffffffffff1681565b6102356103ea3660046122c4565b610ec0565b6102356103fd366004612331565b61122e565b610415610410366004612379565b6114d9565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835267ffffffffffffffff9091166020830152016101c9565b60606104757f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f006114f9565b905090565b5f6104cf827f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f005b9073ffffffffffffffffffffffffffffffffffffffff165f9081526001919091016020526040902054151590565b92915050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d80a4c286040518163ffffffff1660e01b8152600401602060405180830381865afa15801561053f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610475919061239c565b61056b61150c565b5f610574610b7a565b90508073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105f8576040517f81efeac100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024015b60405180910390fd5b435f6106046002611564565b50915050816bffffffffffffffffffffffff16816bffffffffffffffffffffffff1603610669576040517fdd5b488f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526024016105ef565b61067560028386611603565b50508167ffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fcf8f6f62babb05dd1d159c090ad8429ee0df72e16c82701004a5405f908cb0f560405160405180910390a450505050565b5f8061070c7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b905073ffffffffffffffffffffffffffffffffffffffff81161561072f57919050565b6040518060400160405280601a81526020017f4f564d5f4c3143726f7373446f6d61696e4d657373656e67657200000000000081525051600261077291906123e4565b604080513060208201525f918101919091527f4f564d5f4c3143726f7373446f6d61696e4d657373656e67657200000000000091909117906107cc906060015b604051602081830303815290604052805190602001205490565b14610803576040517f54e433cd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080513060208201526001918101919091525f90610824906060016107b2565b905073ffffffffffffffffffffffffffffffffffffffff8116156108b6578073ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561088b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108af919061239c565b9250505090565b6040517f332144db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61047560025490565b5f6104757f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f0061161d565b610946337f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f006104a1565b1580156109865750610956610b85565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b156109bf576040517fd53780c40000000000000000000000000000000000000000000000000000000081523360048201526024016105ef565b5f5460ff7401000000000000000000000000000000000000000090910416151581151514610a57575f80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000083151590810291909117825560405190917fb957d7fc29e5974594db2f2e132076d52f42c0734eae05fd5ea080d1ba175ad391a25b50565b610a6261150c565b7f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f00610a8d8183611626565b610a95575050565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52905f90a25050565b610ae161150c565b610aea5f611647565b565b3380610af6610e0d565b73ffffffffffffffffffffffffffffffffffffffff1614610b5b576040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016105ef565b610a5781611647565b5f6104cf600267ffffffffffffffff8416611697565b5f6104756002611756565b5f6104756117a3565b610b9661150c565b73ffffffffffffffffffffffffffffffffffffffff8116610be3576040517f1b08105400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610beb610b85565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610c565750610c27610e0d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15610c8d576040517f3af3c41c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f00610cb881836117cb565b15610cfe5760405173ffffffffffffffffffffffffffffffffffffffff8316907f038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969905f90a25b5050565b5f80546040517fdac79fc800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169163dac79fc891610d5f9188918891889188919060040161247b565b5f604051808303815f87803b158015610d76575f5ffd5b505af1158015610d88573d5f5f3e3d5ffd5b50506040513392507f665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c91505f90a250505050565b5f610dc56106e2565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561053f573d5f5f3e3d5ffd5b5f807f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005b5473ffffffffffffffffffffffffffffffffffffffff1692915050565b610e5661150c565b610e80817f0f4ac8aae5a4fa6a3612928fcd8255b475ff86b500ae30bb272e61542cfc6f006104a1565b15610eb7576040517f3af3c41c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a57816117ec565b7f000000000000000000000000000000000000000000000000000000000000000060ff165f610eed6118a3565b805490915068010000000000000000900460ff1680610f1a5750805467ffffffffffffffff808416911610155b15610f51576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001667ffffffffffffffff83161768010000000000000000178155610f966118cb565b610f9f8461194c565b73ffffffffffffffffffffffffffffffffffffffff8616611004576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871660048201526024016105ef565b73ffffffffffffffffffffffffffffffffffffffff8516611069576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024016105ef565b73ffffffffffffffffffffffffffffffffffffffff87166110ce576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff881660048201526024016105ef565b5f8054600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8981169190911790915589167fffffffffffffffffffffff0000000000000000000000000000000000000000009091161774010000000000000000000000000000000000000000851515021790556002545f036111c2574361117060028289611603565b505060405167ffffffffffffffff82169073ffffffffffffffffffffffffffffffffffffffff8916905f907fcf8f6f62babb05dd1d159c090ad8429ee0df72e16c82701004a5405f908cb0f5908290a4505b80547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050505050565b5f5474010000000000000000000000000000000000000000900460ff161561139d575f611259610b7a565b90503373ffffffffffffffffffffffffffffffffffffffff8216146112c8576040517f8d1db98a00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044016105ef565b5f80546040517fa81d9c5c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169163a81d9c5c9161132291879187918a916004016124bb565b602060405180830381865afa15801561133d573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061136191906124ec565b611397576040517f8baa579f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061149f565b600154604080517fe81b2c6d00000000000000000000000000000000000000000000000000000000815290515f9273ffffffffffffffffffffffffffffffffffffffff169163e81b2c6d9160048083019260209291908290030181865afa15801561140a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061142e9190612507565b90503373ffffffffffffffffffffffffffffffffffffffff82161461149d576040517f51f905ea00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044016105ef565b505b60405183815233907f731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c9060200160405180910390a2505050565b5f80806114e760028561195d565b60208101519051909590945092505050565b60605f611505836119e6565b9392505050565b33611515610b85565b73ffffffffffffffffffffffffffffffffffffffff1614610aea576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016105ef565b80545f9081908190808203611582575f5f5f935093509350506115fc565b5f61159f8661159260018561251e565b5f91825260209091200190565b6040805180820190915290546bffffffffffffffffffffffff81168083526c0100000000000000000000000090910473ffffffffffffffffffffffffffffffffffffffff16602090920182905260019650945092506115fc915050565b9193909250565b5f80611610858585611a3f565b915091505b935093915050565b5f6104cf825490565b5f6115058373ffffffffffffffffffffffffffffffffffffffff8416611c44565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155610cfe82611d27565b81545f90818160058111156116f9575f6116b084611dbc565b6116ba908561251e565b5f888152602090209091508101546bffffffffffffffffffffffff90811690871610156116e9578091506116f7565b6116f4816001612531565b92505b505b5f61170687878585611ea0565b905080156117495761171d8761159260018461251e565b546c01000000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1661174b565b5f5b979650505050505050565b80545f90801561179b5761176f8361159260018461251e565b546c01000000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff16611505565b5f9392505050565b5f807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300610e31565b5f6115058373ffffffffffffffffffffffffffffffffffffffff8416611f0b565b6117f461150c565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117825561185d610b85565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006104cf565b336118d46106e2565b73ffffffffffffffffffffffffffffffffffffffff16141580156119155750336118fc610dbc565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610aea576040517fc4050a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611954611f57565b610a5781611f95565b604080518082019091525f8082526020820152825f018263ffffffff168154811061198a5761198a612544565b5f918252602091829020604080518082019091529101546bffffffffffffffffffffffff811682526c01000000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff16918101919091529392505050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611a3357602002820191905f5260205f20905b815481526020019060010190808311611a1f575b50505050509050919050565b82545f9081908015611bce575f611a5b8761159260018561251e565b6040805180820190915290546bffffffffffffffffffffffff8082168084526c0100000000000000000000000090920473ffffffffffffffffffffffffffffffffffffffff1660208401529192509087161015611ae4576040517f2520601d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516bffffffffffffffffffffffff808816911603611b525784611b0d8861159260018661251e565b805473ffffffffffffffffffffffffffffffffffffffff929092166c01000000000000000000000000026bffffffffffffffffffffffff909216919091179055611bbe565b604080518082019091526bffffffffffffffffffffffff808816825273ffffffffffffffffffffffffffffffffffffffff80881660208085019182528b54600181018d555f8d815291909120945191519092166c01000000000000000000000000029216919091179101555b6020015192508391506116159050565b5050604080518082019091526bffffffffffffffffffffffff808516825273ffffffffffffffffffffffffffffffffffffffff80851660208085019182528854600181018a555f8a8152918220955192519093166c01000000000000000000000000029190931617920191909155905081611615565b5f8181526001830160205260408120548015611d1e575f611c6660018361251e565b85549091505f90611c799060019061251e565b9050808214611cd8575f865f018281548110611c9757611c97612544565b905f5260205f200154905080875f018481548110611cb757611cb7612544565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611ce957611ce9612571565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506104cf565b5f9150506104cf565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f815f03611dcb57505f919050565b5f6001611dd784611fa6565b901c6001901b90506001818481611df057611df061259e565b048201901c90506001818481611e0857611e0861259e565b048201901c90506001818481611e2057611e2061259e565b048201901c90506001818481611e3857611e3861259e565b048201901c90506001818481611e5057611e5061259e565b048201901c90506001818481611e6857611e6861259e565b048201901c90506001818481611e8057611e8061259e565b048201901c905061150581828581611e9a57611e9a61259e565b04612039565b5f5b81831015611f03575f611eb5848461204e565b5f878152602090209091506bffffffffffffffffffffffff8616908201546bffffffffffffffffffffffff161115611eef57809250611efd565b611efa816001612531565b93505b50611ea2565b509392505050565b5f818152600183016020526040812054611f5057508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556104cf565b505f6104cf565b611f5f612068565b610aea576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f9d611f57565b610a5781612086565b5f80608083901c15611fba57608092831c92015b604083901c15611fcc57604092831c92015b602083901c15611fde57602092831c92015b601083901c15611ff057601092831c92015b600883901c1561200257600892831c92015b600483901c1561201457600492831c92015b600283901c1561202657600292831c92015b600183901c156104cf5760010192915050565b5f8183106120475781611505565b5090919050565b5f61205c60028484186125cb565b61150590848416612531565b5f6120716118a3565b5468010000000000000000900460ff16919050565b61208e611f57565b73ffffffffffffffffffffffffffffffffffffffff8116610b5b576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f60048201526024016105ef565b602080825282518282018190525f918401906040840190835b8181101561212a57835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016120f6565b509095945050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a57575f5ffd5b5f60208284031215612166575f5ffd5b813561150581612135565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b8015158114610a57575f5ffd5b5f602082840312156121e1575f5ffd5b8135611505816121c4565b5f602082840312156121fc575f5ffd5b813567ffffffffffffffff81168114611505575f5ffd5b5f5f83601f840112612223575f5ffd5b50813567ffffffffffffffff81111561223a575f5ffd5b602083019150836020828501011115612251575f5ffd5b9250929050565b5f5f5f5f6040858703121561226b575f5ffd5b843567ffffffffffffffff811115612281575f5ffd5b61228d87828801612213565b909550935050602085013567ffffffffffffffff8111156122ac575f5ffd5b6122b887828801612213565b95989497509550505050565b5f5f5f5f5f60a086880312156122d8575f5ffd5b85356122e381612135565b945060208601356122f381612135565b9350604086013561230381612135565b9250606086013561231381612135565b91506080860135612323816121c4565b809150509295509295909350565b5f5f5f60408486031215612343575f5ffd5b83359250602084013567ffffffffffffffff811115612360575f5ffd5b61236c86828701612213565b9497909650939450505050565b5f60208284031215612389575f5ffd5b813563ffffffff81168114611505575f5ffd5b5f602082840312156123ac575f5ffd5b815161150581612135565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820281158282048414176104cf576104cf6123b7565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60018110612477577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b606081525f61248e6060830187896123fb565b82810360208401526124a18186886123fb565b9150506124b16040830184612442565b9695505050505050565b606081525f6124ce6060830186886123fb565b90508360208301526124e36040830184612442565b95945050505050565b5f602082840312156124fc575f5ffd5b8151611505816121c4565b5f60208284031215612517575f5ffd5b5051919050565b818103818111156104cf576104cf6123b7565b808201808211156104cf576104cf6123b7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f826125fe577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c634300081e000a",
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

// EspressoBatcherAt is a free data retrieval call binding the contract method 0xfd402af7.
//
// Solidity: function espressoBatcherAt(uint32 _index) view returns(address batcher_, uint64 fromBlock_)
func (_BatchAuthenticator *BatchAuthenticatorCaller) EspressoBatcherAt(opts *bind.CallOpts, _index uint32) (struct {
	Batcher   common.Address
	FromBlock uint64
}, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "espressoBatcherAt", _index)

	outstruct := new(struct {
		Batcher   common.Address
		FromBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Batcher = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FromBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// EspressoBatcherAt is a free data retrieval call binding the contract method 0xfd402af7.
//
// Solidity: function espressoBatcherAt(uint32 _index) view returns(address batcher_, uint64 fromBlock_)
func (_BatchAuthenticator *BatchAuthenticatorSession) EspressoBatcherAt(_index uint32) (struct {
	Batcher   common.Address
	FromBlock uint64
}, error) {
	return _BatchAuthenticator.Contract.EspressoBatcherAt(&_BatchAuthenticator.CallOpts, _index)
}

// EspressoBatcherAt is a free data retrieval call binding the contract method 0xfd402af7.
//
// Solidity: function espressoBatcherAt(uint32 _index) view returns(address batcher_, uint64 fromBlock_)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) EspressoBatcherAt(_index uint32) (struct {
	Batcher   common.Address
	FromBlock uint64
}, error) {
	return _BatchAuthenticator.Contract.EspressoBatcherAt(&_BatchAuthenticator.CallOpts, _index)
}

// EspressoBatcherAtBlock is a free data retrieval call binding the contract method 0x7d531a78.
//
// Solidity: function espressoBatcherAtBlock(uint64 _l1Block) view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) EspressoBatcherAtBlock(opts *bind.CallOpts, _l1Block uint64) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "espressoBatcherAtBlock", _l1Block)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoBatcherAtBlock is a free data retrieval call binding the contract method 0x7d531a78.
//
// Solidity: function espressoBatcherAtBlock(uint64 _l1Block) view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) EspressoBatcherAtBlock(_l1Block uint64) (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoBatcherAtBlock(&_BatchAuthenticator.CallOpts, _l1Block)
}

// EspressoBatcherAtBlock is a free data retrieval call binding the contract method 0x7d531a78.
//
// Solidity: function espressoBatcherAtBlock(uint64 _l1Block) view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) EspressoBatcherAtBlock(_l1Block uint64) (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoBatcherAtBlock(&_BatchAuthenticator.CallOpts, _l1Block)
}

// EspressoBatcherHistoryLength is a free data retrieval call binding the contract method 0x4268ecaa.
//
// Solidity: function espressoBatcherHistoryLength() view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorCaller) EspressoBatcherHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "espressoBatcherHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EspressoBatcherHistoryLength is a free data retrieval call binding the contract method 0x4268ecaa.
//
// Solidity: function espressoBatcherHistoryLength() view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorSession) EspressoBatcherHistoryLength() (*big.Int, error) {
	return _BatchAuthenticator.Contract.EspressoBatcherHistoryLength(&_BatchAuthenticator.CallOpts)
}

// EspressoBatcherHistoryLength is a free data retrieval call binding the contract method 0x4268ecaa.
//
// Solidity: function espressoBatcherHistoryLength() view returns(uint256)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) EspressoBatcherHistoryLength() (*big.Int, error) {
	return _BatchAuthenticator.Contract.EspressoBatcherHistoryLength(&_BatchAuthenticator.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xfc5b5fda.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _espressoBatcher, address _systemConfig, address _owner, bool _activeIsEspresso) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) Initialize(opts *bind.TransactOpts, _espressoTEEVerifier common.Address, _espressoBatcher common.Address, _systemConfig common.Address, _owner common.Address, _activeIsEspresso bool) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "initialize", _espressoTEEVerifier, _espressoBatcher, _systemConfig, _owner, _activeIsEspresso)
}

// Initialize is a paid mutator transaction binding the contract method 0xfc5b5fda.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _espressoBatcher, address _systemConfig, address _owner, bool _activeIsEspresso) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) Initialize(_espressoTEEVerifier common.Address, _espressoBatcher common.Address, _systemConfig common.Address, _owner common.Address, _activeIsEspresso bool) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _espressoBatcher, _systemConfig, _owner, _activeIsEspresso)
}

// Initialize is a paid mutator transaction binding the contract method 0xfc5b5fda.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _espressoBatcher, address _systemConfig, address _owner, bool _activeIsEspresso) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) Initialize(_espressoTEEVerifier common.Address, _espressoBatcher common.Address, _systemConfig common.Address, _owner common.Address, _activeIsEspresso bool) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _espressoBatcher, _systemConfig, _owner, _activeIsEspresso)
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

// SetActiveIsEspresso is a paid mutator transaction binding the contract method 0x6c076871.
//
// Solidity: function setActiveIsEspresso(bool _desired) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SetActiveIsEspresso(opts *bind.TransactOpts, _desired bool) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "setActiveIsEspresso", _desired)
}

// SetActiveIsEspresso is a paid mutator transaction binding the contract method 0x6c076871.
//
// Solidity: function setActiveIsEspresso(bool _desired) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SetActiveIsEspresso(_desired bool) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetActiveIsEspresso(&_BatchAuthenticator.TransactOpts, _desired)
}

// SetActiveIsEspresso is a paid mutator transaction binding the contract method 0x6c076871.
//
// Solidity: function setActiveIsEspresso(bool _desired) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SetActiveIsEspresso(_desired bool) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetActiveIsEspresso(&_BatchAuthenticator.TransactOpts, _desired)
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
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchInfoAuthenticated is a free log retrieval operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 commitment, address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterBatchInfoAuthenticated(opts *bind.FilterOpts, caller []common.Address) (*BatchAuthenticatorBatchInfoAuthenticatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "BatchInfoAuthenticated", callerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorBatchInfoAuthenticatedIterator{contract: _BatchAuthenticator.contract, event: "BatchInfoAuthenticated", logs: logs, sub: sub}, nil
}

// WatchBatchInfoAuthenticated is a free log subscription operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 commitment, address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchBatchInfoAuthenticated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorBatchInfoAuthenticated, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "BatchInfoAuthenticated", callerRule)
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
// Solidity: event BatchInfoAuthenticated(bytes32 commitment, address indexed caller)
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
	FromBlock          uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEspressoBatcherUpdated is a free log retrieval operation binding the contract event 0xcf8f6f62babb05dd1d159c090ad8429ee0df72e16c82701004a5405f908cb0f5.
//
// Solidity: event EspressoBatcherUpdated(address indexed oldEspressoBatcher, address indexed newEspressoBatcher, uint64 indexed fromBlock)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterEspressoBatcherUpdated(opts *bind.FilterOpts, oldEspressoBatcher []common.Address, newEspressoBatcher []common.Address, fromBlock []uint64) (*BatchAuthenticatorEspressoBatcherUpdatedIterator, error) {

	var oldEspressoBatcherRule []interface{}
	for _, oldEspressoBatcherItem := range oldEspressoBatcher {
		oldEspressoBatcherRule = append(oldEspressoBatcherRule, oldEspressoBatcherItem)
	}
	var newEspressoBatcherRule []interface{}
	for _, newEspressoBatcherItem := range newEspressoBatcher {
		newEspressoBatcherRule = append(newEspressoBatcherRule, newEspressoBatcherItem)
	}
	var fromBlockRule []interface{}
	for _, fromBlockItem := range fromBlock {
		fromBlockRule = append(fromBlockRule, fromBlockItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "EspressoBatcherUpdated", oldEspressoBatcherRule, newEspressoBatcherRule, fromBlockRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorEspressoBatcherUpdatedIterator{contract: _BatchAuthenticator.contract, event: "EspressoBatcherUpdated", logs: logs, sub: sub}, nil
}

// WatchEspressoBatcherUpdated is a free log subscription operation binding the contract event 0xcf8f6f62babb05dd1d159c090ad8429ee0df72e16c82701004a5405f908cb0f5.
//
// Solidity: event EspressoBatcherUpdated(address indexed oldEspressoBatcher, address indexed newEspressoBatcher, uint64 indexed fromBlock)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchEspressoBatcherUpdated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorEspressoBatcherUpdated, oldEspressoBatcher []common.Address, newEspressoBatcher []common.Address, fromBlock []uint64) (event.Subscription, error) {

	var oldEspressoBatcherRule []interface{}
	for _, oldEspressoBatcherItem := range oldEspressoBatcher {
		oldEspressoBatcherRule = append(oldEspressoBatcherRule, oldEspressoBatcherItem)
	}
	var newEspressoBatcherRule []interface{}
	for _, newEspressoBatcherItem := range newEspressoBatcher {
		newEspressoBatcherRule = append(newEspressoBatcherRule, newEspressoBatcherItem)
	}
	var fromBlockRule []interface{}
	for _, fromBlockItem := range fromBlock {
		fromBlockRule = append(fromBlockRule, fromBlockItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "EspressoBatcherUpdated", oldEspressoBatcherRule, newEspressoBatcherRule, fromBlockRule)
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

// ParseEspressoBatcherUpdated is a log parse operation binding the contract event 0xcf8f6f62babb05dd1d159c090ad8429ee0df72e16c82701004a5405f908cb0f5.
//
// Solidity: event EspressoBatcherUpdated(address indexed oldEspressoBatcher, address indexed newEspressoBatcher, uint64 indexed fromBlock)
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
