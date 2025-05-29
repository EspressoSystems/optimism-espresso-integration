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

// EspressoTEEVerifierMetaData contains all meta data concerning the EspressoTEEVerifier contract.
var EspressoTEEVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoSGXTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoSGXTEEVerifier\"},{\"name\":\"_espressoNitroTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoNitroTEEVerifier\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"espressoNitroTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoNitroTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoSGXTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoSGXTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"teeType\",\"type\":\"uint8\",\"internalType\":\"enumIEspressoTEEVerifier.TeeType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredEnclaveHashes\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"teeType\",\"type\":\"uint8\",\"internalType\":\"enumIEspressoTEEVerifier.TeeType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registeredSigners\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"teeType\",\"type\":\"uint8\",\"internalType\":\"enumIEspressoTEEVerifier.TeeType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEspressoNitroTEEVerifier\",\"inputs\":[{\"name\":\"_espressoNitroTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoNitroTEEVerifier\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEspressoSGXTEEVerifier\",\"inputs\":[{\"name\":\"_espressoSGXTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoSGXTEEVerifier\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userDataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedTeeType\",\"inputs\":[]}]",
	Bin: "0x6080346100aa57601f61115d38819003918201601f19168301916001600160401b038311848410176100ae5780849260409485528339810103126100aa5780516001600160a01b03811691908290036100aa57602001516001600160a01b03811691908290036100aa57610072336100c2565b60018060a01b0319600254161760025560018060a01b0319600354161760035561009b336100c2565b60405161104690816101178239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b600180546001600160a01b03199081169091555f80546001600160a01b03938416928116831782559192909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a356fe60806040526004361015610011575f80fd5b5f3560e01c8063330282f5146108c457806335ecb4c11461083c5780633cbe6803146107f35780636b406341146105ad578063715018a6146104eb57806379ba50971461038b57806380710c801461033a5780638da5cb5b146102ea578063bc3a091114610265578063d80a4c2814610214578063e30c3978146101c3578063e9b1a7be146101695763f2fde38b146100a8575f80fd5b346101655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101655773ffffffffffffffffffffffffffffffffffffffff6100f46109b8565b6100fc610d94565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff5f54167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227005f80a3005b5f80fd5b346101655760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610165576101a06109b8565b602435906002821015610165576020916101b991610c8e565b6040519015158152f35b34610165575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b34610165575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016557602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b346101655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101655760043573ffffffffffffffffffffffffffffffffffffffff8116809103610165576102bd610d94565b7fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555f80f35b34610165575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016557602073ffffffffffffffffffffffffffffffffffffffff5f5416604051908152f35b34610165575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016557602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b34610165575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610165573373ffffffffffffffffffffffffffffffffffffffff6001541603610467577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154166001555f54337fffffffffffffffffffffffff00000000000000000000000000000000000000008216175f5573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4f776e61626c6532537465703a2063616c6c6572206973206e6f74207468652060448201527f6e6577206f776e657200000000000000000000000000000000000000000000006064820152fd5b34610165575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016557610521610d94565b7fffffffffffffffffffffffff0000000000000000000000000000000000000000600154166001555f73ffffffffffffffffffffffffffffffffffffffff81547fffffffffffffffffffffffff000000000000000000000000000000000000000081168355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346101655760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101655760043567ffffffffffffffff811161016557366023820112156101655780600401359067ffffffffffffffff82116107c65760405161064360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8601160182610977565b8281523660248484010111610165575f60208461067995602461067196018386013783010152602435610e12565b919091610e47565b73ffffffffffffffffffffffffffffffffffffffff60208160025416926024604051809481937f0123d0c100000000000000000000000000000000000000000000000000000000835216958660048301525afa90811561079c575f916107a7575b501561074557602073ffffffffffffffffffffffffffffffffffffffff60035416916024604051809481937f0123d0c100000000000000000000000000000000000000000000000000000000835260048301525afa90811561079c575f9161076d575b501561074557005b7f8baa579f000000000000000000000000000000000000000000000000000000005f5260045ffd5b61078f915060203d602011610795575b6107878183610977565b810190610bc6565b8161073d565b503d61077d565b6040513d5f823e3d90fd5b6107c0915060203d602011610795576107878183610977565b826106da565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b346101655760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610165576024356002811015610165576101b9602091600435610bde565b346101655760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101655760043567ffffffffffffffff81116101655761088b903690600401610949565b60243567ffffffffffffffff8111610165576108ab903690600401610949565b90604435926002841015610165576108c294610a43565b005b346101655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101655760043573ffffffffffffffffffffffffffffffffffffffff81168091036101655761091c610d94565b7fffffffffffffffffffffffff000000000000000000000000000000000000000060035416176003555f80f35b9181601f840112156101655782359167ffffffffffffffff8311610165576020838186019501011161016557565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176107c657604052565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361016557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe093818652868601375f8582860101520116010190565b9290610a3290610a4095936040865260408601916109db565b9260208185039101526109db565b90565b905f946002811015610b99578015610b1a57600114610a84576004857fd0cb35a1000000000000000000000000000000000000000000000000000000008152fd5b73ffffffffffffffffffffffffffffffffffffffff6003541691823b15610b1657908580949392610ae4604051978896879586947fba58e82a00000000000000000000000000000000000000000000000000000000865260048601610a19565b03925af18015610b0b57610af6575050565b610b01828092610977565b610b085750565b80fd5b6040513d84823e3d90fd5b8580fd5b509092935073ffffffffffffffffffffffffffffffffffffffff6002541690813b15610165575f8094610b7c604051978896879586947fba58e82a00000000000000000000000000000000000000000000000000000000865260048601610a19565b03925af1801561079c57610b8d5750565b5f610b9791610977565b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b90816020910312610165575180151581036101655790565b906002811015610b995715610c15577fd0cb35a1000000000000000000000000000000000000000000000000000000005f5260045ffd5b602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937f966989ee00000000000000000000000000000000000000000000000000000000835260048301525afa90811561079c575f91610c75575090565b610a40915060203d602011610795576107878183610977565b906002811015610b99578015610d3057600114610ccd577fd0cb35a1000000000000000000000000000000000000000000000000000000005f5260045ffd5b602073ffffffffffffffffffffffffffffffffffffffff602481600354169360405194859384927f0123d0c10000000000000000000000000000000000000000000000000000000084521660048301525afa90811561079c575f91610c75575090565b50602073ffffffffffffffffffffffffffffffffffffffff602481600254169360405194859384927f0123d0c10000000000000000000000000000000000000000000000000000000084521660048301525afa90811561079c575f91610c75575090565b73ffffffffffffffffffffffffffffffffffffffff5f54163303610db457565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b9060418151145f14610e3e57610e3a91602082015190606060408401519301515f1a90610fb1565b9091565b50505f90600290565b6005811015610b995780610e585750565b60018103610ebe5760646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152fd5b60028103610f245760646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152fd5b600314610f2d57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152fd5b7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161102e576020935f9360ff60809460405194855216868401526040830152606082015282805260015afa1561079c575f5173ffffffffffffffffffffffffffffffffffffffff81161561102657905f90565b505f90600190565b505050505f9060039056fea164736f6c634300081c000a",
}

// EspressoTEEVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use EspressoTEEVerifierMetaData.ABI instead.
var EspressoTEEVerifierABI = EspressoTEEVerifierMetaData.ABI

// EspressoTEEVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EspressoTEEVerifierMetaData.Bin instead.
var EspressoTEEVerifierBin = EspressoTEEVerifierMetaData.Bin

// DeployEspressoTEEVerifier deploys a new Ethereum contract, binding an instance of EspressoTEEVerifier to it.
func DeployEspressoTEEVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _espressoSGXTEEVerifier common.Address, _espressoNitroTEEVerifier common.Address) (common.Address, *types.Transaction, *EspressoTEEVerifier, error) {
	parsed, err := EspressoTEEVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EspressoTEEVerifierBin), backend, _espressoSGXTEEVerifier, _espressoNitroTEEVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EspressoTEEVerifier{EspressoTEEVerifierCaller: EspressoTEEVerifierCaller{contract: contract}, EspressoTEEVerifierTransactor: EspressoTEEVerifierTransactor{contract: contract}, EspressoTEEVerifierFilterer: EspressoTEEVerifierFilterer{contract: contract}}, nil
}

// EspressoTEEVerifier is an auto generated Go binding around an Ethereum contract.
type EspressoTEEVerifier struct {
	EspressoTEEVerifierCaller     // Read-only binding to the contract
	EspressoTEEVerifierTransactor // Write-only binding to the contract
	EspressoTEEVerifierFilterer   // Log filterer for contract events
}

// EspressoTEEVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type EspressoTEEVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspressoTEEVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EspressoTEEVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspressoTEEVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EspressoTEEVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspressoTEEVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EspressoTEEVerifierSession struct {
	Contract     *EspressoTEEVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EspressoTEEVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EspressoTEEVerifierCallerSession struct {
	Contract *EspressoTEEVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EspressoTEEVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EspressoTEEVerifierTransactorSession struct {
	Contract     *EspressoTEEVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EspressoTEEVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type EspressoTEEVerifierRaw struct {
	Contract *EspressoTEEVerifier // Generic contract binding to access the raw methods on
}

// EspressoTEEVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EspressoTEEVerifierCallerRaw struct {
	Contract *EspressoTEEVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// EspressoTEEVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EspressoTEEVerifierTransactorRaw struct {
	Contract *EspressoTEEVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEspressoTEEVerifier creates a new instance of EspressoTEEVerifier, bound to a specific deployed contract.
func NewEspressoTEEVerifier(address common.Address, backend bind.ContractBackend) (*EspressoTEEVerifier, error) {
	contract, err := bindEspressoTEEVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EspressoTEEVerifier{EspressoTEEVerifierCaller: EspressoTEEVerifierCaller{contract: contract}, EspressoTEEVerifierTransactor: EspressoTEEVerifierTransactor{contract: contract}, EspressoTEEVerifierFilterer: EspressoTEEVerifierFilterer{contract: contract}}, nil
}

// NewEspressoTEEVerifierCaller creates a new read-only instance of EspressoTEEVerifier, bound to a specific deployed contract.
func NewEspressoTEEVerifierCaller(address common.Address, caller bind.ContractCaller) (*EspressoTEEVerifierCaller, error) {
	contract, err := bindEspressoTEEVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EspressoTEEVerifierCaller{contract: contract}, nil
}

// NewEspressoTEEVerifierTransactor creates a new write-only instance of EspressoTEEVerifier, bound to a specific deployed contract.
func NewEspressoTEEVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*EspressoTEEVerifierTransactor, error) {
	contract, err := bindEspressoTEEVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EspressoTEEVerifierTransactor{contract: contract}, nil
}

// NewEspressoTEEVerifierFilterer creates a new log filterer instance of EspressoTEEVerifier, bound to a specific deployed contract.
func NewEspressoTEEVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*EspressoTEEVerifierFilterer, error) {
	contract, err := bindEspressoTEEVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EspressoTEEVerifierFilterer{contract: contract}, nil
}

// bindEspressoTEEVerifier binds a generic wrapper to an already deployed contract.
func bindEspressoTEEVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EspressoTEEVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EspressoTEEVerifier *EspressoTEEVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EspressoTEEVerifier.Contract.EspressoTEEVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EspressoTEEVerifier *EspressoTEEVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.EspressoTEEVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EspressoTEEVerifier *EspressoTEEVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.EspressoTEEVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EspressoTEEVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.contract.Transact(opts, method, params...)
}

// EspressoNitroTEEVerifier is a free data retrieval call binding the contract method 0xd80a4c28.
//
// Solidity: function espressoNitroTEEVerifier() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCaller) EspressoNitroTEEVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoTEEVerifier.contract.Call(opts, &out, "espressoNitroTEEVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoNitroTEEVerifier is a free data retrieval call binding the contract method 0xd80a4c28.
//
// Solidity: function espressoNitroTEEVerifier() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) EspressoNitroTEEVerifier() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.EspressoNitroTEEVerifier(&_EspressoTEEVerifier.CallOpts)
}

// EspressoNitroTEEVerifier is a free data retrieval call binding the contract method 0xd80a4c28.
//
// Solidity: function espressoNitroTEEVerifier() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerSession) EspressoNitroTEEVerifier() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.EspressoNitroTEEVerifier(&_EspressoTEEVerifier.CallOpts)
}

// EspressoSGXTEEVerifier is a free data retrieval call binding the contract method 0x80710c80.
//
// Solidity: function espressoSGXTEEVerifier() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCaller) EspressoSGXTEEVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoTEEVerifier.contract.Call(opts, &out, "espressoSGXTEEVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoSGXTEEVerifier is a free data retrieval call binding the contract method 0x80710c80.
//
// Solidity: function espressoSGXTEEVerifier() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) EspressoSGXTEEVerifier() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.EspressoSGXTEEVerifier(&_EspressoTEEVerifier.CallOpts)
}

// EspressoSGXTEEVerifier is a free data retrieval call binding the contract method 0x80710c80.
//
// Solidity: function espressoSGXTEEVerifier() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerSession) EspressoSGXTEEVerifier() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.EspressoSGXTEEVerifier(&_EspressoTEEVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoTEEVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) Owner() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.Owner(&_EspressoTEEVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerSession) Owner() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.Owner(&_EspressoTEEVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoTEEVerifier.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) PendingOwner() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.PendingOwner(&_EspressoTEEVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerSession) PendingOwner() (common.Address, error) {
	return _EspressoTEEVerifier.Contract.PendingOwner(&_EspressoTEEVerifier.CallOpts)
}

// RegisteredEnclaveHashes is a free data retrieval call binding the contract method 0x3cbe6803.
//
// Solidity: function registeredEnclaveHashes(bytes32 enclaveHash, uint8 teeType) view returns(bool)
func (_EspressoTEEVerifier *EspressoTEEVerifierCaller) RegisteredEnclaveHashes(opts *bind.CallOpts, enclaveHash [32]byte, teeType uint8) (bool, error) {
	var out []interface{}
	err := _EspressoTEEVerifier.contract.Call(opts, &out, "registeredEnclaveHashes", enclaveHash, teeType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredEnclaveHashes is a free data retrieval call binding the contract method 0x3cbe6803.
//
// Solidity: function registeredEnclaveHashes(bytes32 enclaveHash, uint8 teeType) view returns(bool)
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) RegisteredEnclaveHashes(enclaveHash [32]byte, teeType uint8) (bool, error) {
	return _EspressoTEEVerifier.Contract.RegisteredEnclaveHashes(&_EspressoTEEVerifier.CallOpts, enclaveHash, teeType)
}

// RegisteredEnclaveHashes is a free data retrieval call binding the contract method 0x3cbe6803.
//
// Solidity: function registeredEnclaveHashes(bytes32 enclaveHash, uint8 teeType) view returns(bool)
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerSession) RegisteredEnclaveHashes(enclaveHash [32]byte, teeType uint8) (bool, error) {
	return _EspressoTEEVerifier.Contract.RegisteredEnclaveHashes(&_EspressoTEEVerifier.CallOpts, enclaveHash, teeType)
}

// RegisteredSigners is a free data retrieval call binding the contract method 0xe9b1a7be.
//
// Solidity: function registeredSigners(address signer, uint8 teeType) view returns(bool)
func (_EspressoTEEVerifier *EspressoTEEVerifierCaller) RegisteredSigners(opts *bind.CallOpts, signer common.Address, teeType uint8) (bool, error) {
	var out []interface{}
	err := _EspressoTEEVerifier.contract.Call(opts, &out, "registeredSigners", signer, teeType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredSigners is a free data retrieval call binding the contract method 0xe9b1a7be.
//
// Solidity: function registeredSigners(address signer, uint8 teeType) view returns(bool)
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) RegisteredSigners(signer common.Address, teeType uint8) (bool, error) {
	return _EspressoTEEVerifier.Contract.RegisteredSigners(&_EspressoTEEVerifier.CallOpts, signer, teeType)
}

// RegisteredSigners is a free data retrieval call binding the contract method 0xe9b1a7be.
//
// Solidity: function registeredSigners(address signer, uint8 teeType) view returns(bool)
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerSession) RegisteredSigners(signer common.Address, teeType uint8) (bool, error) {
	return _EspressoTEEVerifier.Contract.RegisteredSigners(&_EspressoTEEVerifier.CallOpts, signer, teeType)
}

// Verify is a free data retrieval call binding the contract method 0x6b406341.
//
// Solidity: function verify(bytes signature, bytes32 userDataHash) view returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierCaller) Verify(opts *bind.CallOpts, signature []byte, userDataHash [32]byte) error {
	var out []interface{}
	err := _EspressoTEEVerifier.contract.Call(opts, &out, "verify", signature, userDataHash)

	if err != nil {
		return err
	}

	return err

}

// Verify is a free data retrieval call binding the contract method 0x6b406341.
//
// Solidity: function verify(bytes signature, bytes32 userDataHash) view returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) Verify(signature []byte, userDataHash [32]byte) error {
	return _EspressoTEEVerifier.Contract.Verify(&_EspressoTEEVerifier.CallOpts, signature, userDataHash)
}

// Verify is a free data retrieval call binding the contract method 0x6b406341.
//
// Solidity: function verify(bytes signature, bytes32 userDataHash) view returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierCallerSession) Verify(signature []byte, userDataHash [32]byte) error {
	return _EspressoTEEVerifier.Contract.Verify(&_EspressoTEEVerifier.CallOpts, signature, userDataHash)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoTEEVerifier.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.AcceptOwnership(&_EspressoTEEVerifier.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.AcceptOwnership(&_EspressoTEEVerifier.TransactOpts)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0x35ecb4c1.
//
// Solidity: function registerSigner(bytes attestation, bytes data, uint8 teeType) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactor) RegisterSigner(opts *bind.TransactOpts, attestation []byte, data []byte, teeType uint8) (*types.Transaction, error) {
	return _EspressoTEEVerifier.contract.Transact(opts, "registerSigner", attestation, data, teeType)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0x35ecb4c1.
//
// Solidity: function registerSigner(bytes attestation, bytes data, uint8 teeType) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) RegisterSigner(attestation []byte, data []byte, teeType uint8) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.RegisterSigner(&_EspressoTEEVerifier.TransactOpts, attestation, data, teeType)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0x35ecb4c1.
//
// Solidity: function registerSigner(bytes attestation, bytes data, uint8 teeType) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorSession) RegisterSigner(attestation []byte, data []byte, teeType uint8) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.RegisterSigner(&_EspressoTEEVerifier.TransactOpts, attestation, data, teeType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoTEEVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.RenounceOwnership(&_EspressoTEEVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.RenounceOwnership(&_EspressoTEEVerifier.TransactOpts)
}

// SetEspressoNitroTEEVerifier is a paid mutator transaction binding the contract method 0x330282f5.
//
// Solidity: function setEspressoNitroTEEVerifier(address _espressoNitroTEEVerifier) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactor) SetEspressoNitroTEEVerifier(opts *bind.TransactOpts, _espressoNitroTEEVerifier common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.contract.Transact(opts, "setEspressoNitroTEEVerifier", _espressoNitroTEEVerifier)
}

// SetEspressoNitroTEEVerifier is a paid mutator transaction binding the contract method 0x330282f5.
//
// Solidity: function setEspressoNitroTEEVerifier(address _espressoNitroTEEVerifier) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) SetEspressoNitroTEEVerifier(_espressoNitroTEEVerifier common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.SetEspressoNitroTEEVerifier(&_EspressoTEEVerifier.TransactOpts, _espressoNitroTEEVerifier)
}

// SetEspressoNitroTEEVerifier is a paid mutator transaction binding the contract method 0x330282f5.
//
// Solidity: function setEspressoNitroTEEVerifier(address _espressoNitroTEEVerifier) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorSession) SetEspressoNitroTEEVerifier(_espressoNitroTEEVerifier common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.SetEspressoNitroTEEVerifier(&_EspressoTEEVerifier.TransactOpts, _espressoNitroTEEVerifier)
}

// SetEspressoSGXTEEVerifier is a paid mutator transaction binding the contract method 0xbc3a0911.
//
// Solidity: function setEspressoSGXTEEVerifier(address _espressoSGXTEEVerifier) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactor) SetEspressoSGXTEEVerifier(opts *bind.TransactOpts, _espressoSGXTEEVerifier common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.contract.Transact(opts, "setEspressoSGXTEEVerifier", _espressoSGXTEEVerifier)
}

// SetEspressoSGXTEEVerifier is a paid mutator transaction binding the contract method 0xbc3a0911.
//
// Solidity: function setEspressoSGXTEEVerifier(address _espressoSGXTEEVerifier) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) SetEspressoSGXTEEVerifier(_espressoSGXTEEVerifier common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.SetEspressoSGXTEEVerifier(&_EspressoTEEVerifier.TransactOpts, _espressoSGXTEEVerifier)
}

// SetEspressoSGXTEEVerifier is a paid mutator transaction binding the contract method 0xbc3a0911.
//
// Solidity: function setEspressoSGXTEEVerifier(address _espressoSGXTEEVerifier) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorSession) SetEspressoSGXTEEVerifier(_espressoSGXTEEVerifier common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.SetEspressoSGXTEEVerifier(&_EspressoTEEVerifier.TransactOpts, _espressoSGXTEEVerifier)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.TransferOwnership(&_EspressoTEEVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EspressoTEEVerifier *EspressoTEEVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EspressoTEEVerifier.Contract.TransferOwnership(&_EspressoTEEVerifier.TransactOpts, newOwner)
}

// EspressoTEEVerifierOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the EspressoTEEVerifier contract.
type EspressoTEEVerifierOwnershipTransferStartedIterator struct {
	Event *EspressoTEEVerifierOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *EspressoTEEVerifierOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoTEEVerifierOwnershipTransferStarted)
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
		it.Event = new(EspressoTEEVerifierOwnershipTransferStarted)
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
func (it *EspressoTEEVerifierOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoTEEVerifierOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoTEEVerifierOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the EspressoTEEVerifier contract.
type EspressoTEEVerifierOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_EspressoTEEVerifier *EspressoTEEVerifierFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EspressoTEEVerifierOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoTEEVerifier.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EspressoTEEVerifierOwnershipTransferStartedIterator{contract: _EspressoTEEVerifier.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_EspressoTEEVerifier *EspressoTEEVerifierFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *EspressoTEEVerifierOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoTEEVerifier.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoTEEVerifierOwnershipTransferStarted)
				if err := _EspressoTEEVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_EspressoTEEVerifier *EspressoTEEVerifierFilterer) ParseOwnershipTransferStarted(log types.Log) (*EspressoTEEVerifierOwnershipTransferStarted, error) {
	event := new(EspressoTEEVerifierOwnershipTransferStarted)
	if err := _EspressoTEEVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoTEEVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EspressoTEEVerifier contract.
type EspressoTEEVerifierOwnershipTransferredIterator struct {
	Event *EspressoTEEVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EspressoTEEVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoTEEVerifierOwnershipTransferred)
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
		it.Event = new(EspressoTEEVerifierOwnershipTransferred)
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
func (it *EspressoTEEVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoTEEVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoTEEVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the EspressoTEEVerifier contract.
type EspressoTEEVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EspressoTEEVerifier *EspressoTEEVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EspressoTEEVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoTEEVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EspressoTEEVerifierOwnershipTransferredIterator{contract: _EspressoTEEVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EspressoTEEVerifier *EspressoTEEVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EspressoTEEVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoTEEVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoTEEVerifierOwnershipTransferred)
				if err := _EspressoTEEVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EspressoTEEVerifier *EspressoTEEVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*EspressoTEEVerifierOwnershipTransferred, error) {
	event := new(EspressoTEEVerifierOwnershipTransferred)
	if err := _EspressoTEEVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
