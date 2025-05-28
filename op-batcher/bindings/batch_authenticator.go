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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"},{\"name\":\"_preApprovedBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInformation\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroValidator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preApprovedBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60e0806040523461010b576040816112a0803803809161001f8285610122565b83398101031261010b5780516001600160a01b0381169182820361010b5760200151916001600160a01b038316830361010b5760049260209260a05260805260405192838092631b01498560e31b82525afa908115610117575f916100d1575b506001600160a01b031660c052604051611146908161015a82396080518181816108b10152610b50015260a05181818160e10152818161032f0152610768015260c0518181816104bd0152610bbc0152f35b90506020813d60201161010f575b816100ec60209383610122565b8101031261010b57516001600160a01b038116810361010b575f61007f565b5f80fd5b3d91506100df565b6040513d5f823e3d90fd5b601f909101601f19168101906001600160401b0382119082101761014557604052565b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f905f3560e01c9081631b076a4c14610b74575080631f568b1814610b0657806354fd4d5014610a8957806359cd323314610675578063715018a6146105d95780638da5cb5b14610588578063a903a27714610405578063ba58e82a146102a4578063f2fde38b14610155578063f81f2083146101085763fa14fe6d14610097575f80fd5b3461010557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261010557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b80fd5b50346101055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101055760ff60406020926004358152606584522054166040519015158152f35b50346101055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101055760043573ffffffffffffffffffffffffffffffffffffffff81168091036102a0576101ae610db3565b801561021c5773ffffffffffffffffffffffffffffffffffffffff603354827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617603355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b5080fd5b50346104015760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104015760043567ffffffffffffffff8111610401576102f4903690600401610ccb565b9060243567ffffffffffffffff811161040157610315903690600401610ccb565b92909173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15610401575f61039f936103cf8296604051988997889687957f35ecb4c1000000000000000000000000000000000000000000000000000000008752606060048801526064870191610d75565b917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc858403016024860152610d75565b6001604483015203925af180156103f6576103e8575080f35b6103f491505f90610be0565b005b6040513d5f823e3d90fd5b5f80fd5b346104015760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104015760043567ffffffffffffffff81116104015736602382011215610401575f6104696104a4923690602481600401359101610cf9565b604051809381927fa903a277000000000000000000000000000000000000000000000000000000008352602060048401526024830190610c88565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156103f6575f905f9261051d575b5061051961050b92604051938493604085526040850190610c88565b908382036020850152610c88565b0390f35b9150503d805f833e61052f8183610be0565b81019060408183031261040157805167ffffffffffffffff81116104015782610559918301610d2f565b9060208101519067ffffffffffffffff82116104015761050b93610519926105819201610d2f565b92506104ef565b34610401575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261040157602073ffffffffffffffffffffffffffffffffffffffff60335416604051908152f35b34610401575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104015761060f610db3565b5f73ffffffffffffffffffffffffffffffffffffffff6033547fffffffffffffffffffffffff00000000000000000000000000000000000000008116603355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346104015760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104015760043560243567ffffffffffffffff8111610401576106ca6106d1913690600401610ccb565b3691610cf9565b805160401015610a255760608101805160f81c80158015610a7f575b6109c2575b505061071e61071673ffffffffffffffffffffffffffffffffffffffff9284610e32565b919091610e67565b168015610964576040517fd80a4c2800000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156103f6575f91610914575b50602073ffffffffffffffffffffffffffffffffffffffff916024604051809481937f0123d0c1000000000000000000000000000000000000000000000000000000008352876004840152165afa9081156103f6575f916108d9575b50159081610898575b5061083a575f52606560205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790555f80f35b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c6964207369676e65720000000000000000000000000000000000006044820152fd5b905073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161415826107fd565b90506020813d60201161090c575b816108f460209383610be0565b810103126104015751801515810361040157836107f4565b3d91506108e7565b90506020813d60201161095c575b8161092f60209383610be0565b81010312610401575173ffffffffffffffffffffffffffffffffffffffff81168103610401576020610798565b3d9150610922565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c6964207369676e61747572650000000000000000000000000000006044820152fd5b601b0160ff8111610a5257825160401015610a255773ffffffffffffffffffffffffffffffffffffffff9261071e927fff000000000000000000000000000000000000000000000000000000000000006107169360f81b165f1a905392506106f2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b50600181146106ed565b34610401575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261040157610519604051610ac8604082610be0565b600581527f312e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190610c88565b34610401575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261040157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610401575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104015760209073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610c2157604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b67ffffffffffffffff8111610c2157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b9181601f840112156104015782359167ffffffffffffffff8311610401576020838186019501011161040157565b929192610d0582610c4e565b91610d136040519384610be0565b829481845281830111610401578281602093845f960137010152565b81601f8201121561040157805190610d4682610c4e565b92610d546040519485610be0565b8284526020838301011161040157815f9260208093018386015e8301015290565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe093818652868601375f8582860101520116010190565b73ffffffffffffffffffffffffffffffffffffffff603354163303610dd457565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b9060418151145f14610e5e57610e5a91602082015190606060408401519301515f1a9061108a565b9091565b50505f90600290565b600581101561105d5780610e785750565b60018103610ede5760646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152fd5b60028103610f445760646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152fd5b60038103610fd05760846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152fd5b600414610fd957565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161112e5760ff1690601b82141580611123575b611118576020935f93608093604051938452868401526040830152606082015282805260015afa156103f6575f5173ffffffffffffffffffffffffffffffffffffffff81161561111057905f90565b505f90600190565b505050505f90600490565b50601c8214156110c1565b505050505f9060039056fea164736f6c634300081c000a",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// BindingsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingsMetaData.Bin instead.
var BindingsBin = BindingsMetaData.Bin

// DeployBindings deploys a new Ethereum contract, binding an instance of Bindings to it.
func DeployBindings(auth *bind.TransactOpts, backend bind.ContractBackend, _espressoTEEVerifier common.Address, _preApprovedBatcher common.Address) (common.Address, *types.Transaction, *Bindings, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingsBin), backend, _espressoTEEVerifier, _preApprovedBatcher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) view returns(bytes, bytes)
func (_Bindings *BindingsCaller) DecodeAttestationTbs(opts *bind.CallOpts, attestation []byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "decodeAttestationTbs", attestation)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) view returns(bytes, bytes)
func (_Bindings *BindingsSession) DecodeAttestationTbs(attestation []byte) ([]byte, []byte, error) {
	return _Bindings.Contract.DecodeAttestationTbs(&_Bindings.CallOpts, attestation)
}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) view returns(bytes, bytes)
func (_Bindings *BindingsCallerSession) DecodeAttestationTbs(attestation []byte) ([]byte, []byte, error) {
	return _Bindings.Contract.DecodeAttestationTbs(&_Bindings.CallOpts, attestation)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_Bindings *BindingsCaller) EspressoTEEVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "espressoTEEVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_Bindings *BindingsSession) EspressoTEEVerifier() (common.Address, error) {
	return _Bindings.Contract.EspressoTEEVerifier(&_Bindings.CallOpts)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_Bindings *BindingsCallerSession) EspressoTEEVerifier() (common.Address, error) {
	return _Bindings.Contract.EspressoTEEVerifier(&_Bindings.CallOpts)
}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_Bindings *BindingsCaller) NitroValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "nitroValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_Bindings *BindingsSession) NitroValidator() (common.Address, error) {
	return _Bindings.Contract.NitroValidator(&_Bindings.CallOpts)
}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_Bindings *BindingsCallerSession) NitroValidator() (common.Address, error) {
	return _Bindings.Contract.NitroValidator(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// PreApprovedBatcher is a free data retrieval call binding the contract method 0x1f568b18.
//
// Solidity: function preApprovedBatcher() view returns(address)
func (_Bindings *BindingsCaller) PreApprovedBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "preApprovedBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreApprovedBatcher is a free data retrieval call binding the contract method 0x1f568b18.
//
// Solidity: function preApprovedBatcher() view returns(address)
func (_Bindings *BindingsSession) PreApprovedBatcher() (common.Address, error) {
	return _Bindings.Contract.PreApprovedBatcher(&_Bindings.CallOpts)
}

// PreApprovedBatcher is a free data retrieval call binding the contract method 0x1f568b18.
//
// Solidity: function preApprovedBatcher() view returns(address)
func (_Bindings *BindingsCallerSession) PreApprovedBatcher() (common.Address, error) {
	return _Bindings.Contract.PreApprovedBatcher(&_Bindings.CallOpts)
}

// ValidBatchInfo is a free data retrieval call binding the contract method 0xf81f2083.
//
// Solidity: function validBatchInfo(bytes32 ) view returns(bool)
func (_Bindings *BindingsCaller) ValidBatchInfo(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "validBatchInfo", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidBatchInfo is a free data retrieval call binding the contract method 0xf81f2083.
//
// Solidity: function validBatchInfo(bytes32 ) view returns(bool)
func (_Bindings *BindingsSession) ValidBatchInfo(arg0 [32]byte) (bool, error) {
	return _Bindings.Contract.ValidBatchInfo(&_Bindings.CallOpts, arg0)
}

// ValidBatchInfo is a free data retrieval call binding the contract method 0xf81f2083.
//
// Solidity: function validBatchInfo(bytes32 ) view returns(bool)
func (_Bindings *BindingsCallerSession) ValidBatchInfo(arg0 [32]byte) (bool, error) {
	return _Bindings.Contract.ValidBatchInfo(&_Bindings.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Bindings *BindingsCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Bindings *BindingsSession) Version() (string, error) {
	return _Bindings.Contract.Version(&_Bindings.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Bindings *BindingsCallerSession) Version() (string, error) {
	return _Bindings.Contract.Version(&_Bindings.CallOpts)
}

// AuthenticateBatchInformation is a paid mutator transaction binding the contract method 0x59cd3233.
//
// Solidity: function authenticateBatchInformation(bytes32 commitment, bytes _signature) returns()
func (_Bindings *BindingsTransactor) AuthenticateBatchInformation(opts *bind.TransactOpts, commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "authenticateBatchInformation", commitment, _signature)
}

// AuthenticateBatchInformation is a paid mutator transaction binding the contract method 0x59cd3233.
//
// Solidity: function authenticateBatchInformation(bytes32 commitment, bytes _signature) returns()
func (_Bindings *BindingsSession) AuthenticateBatchInformation(commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Bindings.Contract.AuthenticateBatchInformation(&_Bindings.TransactOpts, commitment, _signature)
}

// AuthenticateBatchInformation is a paid mutator transaction binding the contract method 0x59cd3233.
//
// Solidity: function authenticateBatchInformation(bytes32 commitment, bytes _signature) returns()
func (_Bindings *BindingsTransactorSession) AuthenticateBatchInformation(commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Bindings.Contract.AuthenticateBatchInformation(&_Bindings.TransactOpts, commitment, _signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_Bindings *BindingsTransactor) RegisterSigner(opts *bind.TransactOpts, attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerSigner", attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_Bindings *BindingsSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterSigner(&_Bindings.TransactOpts, attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_Bindings *BindingsTransactorSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterSigner(&_Bindings.TransactOpts, attestationTbs, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

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
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
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
		it.Event = new(BindingsInitialized)
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
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
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
		it.Event = new(BindingsOwnershipTransferred)
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
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
