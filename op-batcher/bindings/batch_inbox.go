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

// BatchInboxMetaData contains all meta data concerning the BatchInbox contract.
var BatchInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_batchAuthenticator\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchAuthenticator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postBlobs\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"postCalldata\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610c52380380610c528339818101604052810190610031919061022e565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561009957505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b6100d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100cf906102d8565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505060015f5f6101000a81548160ff0219169083151502179055505050506102f6565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101c282610199565b9050919050565b6101d2816101b8565b81146101dc575f5ffd5b50565b5f815190506101ed816101c9565b92915050565b5f6101fd826101b8565b9050919050565b61020d816101f3565b8114610217575f5ffd5b50565b5f8151905061022881610204565b92915050565b5f5f5f6060848603121561024557610244610195565b5b5f610252868287016101df565b9350506020610263868287016101df565b92505060406102748682870161021a565b9150509250925092565b5f82825260208201905092915050565b7f4261746368496e626f783a207a65726f206261746368657200000000000000005f82015250565b5f6102c260188361027e565b91506102cd8261028e565b602082019050919050565b5f6020820190508181035f8301526102ef816102b6565b9050919050565b60805160a05160c05161091d6103355f395f81816101f701526103c401525f8181610354015261042901525f81816103a001526103fd015261091d5ff3fe608060405234801561000f575f5ffd5b506004361061008a575f3560e01c8063b1bd428511610059578063b1bd42851461010f578063bc347f471461012d578063d909ba7c14610137578063e7584573146101555761008b565b80631ad40238146100ad57806354fd4d50146100c95780637098ae43146100e75780637877a9ed146100f15761008b565b5b6100ab5f3660405161009e9291906104de565b6040518091039020610173565b005b6100c760048036038101906100c2919061055f565b6102d2565b005b6100d16102f6565b6040516100de919061061a565b60405180910390f35b6100ef61032f565b005b6100f9610341565b6040516101069190610654565b60405180910390f35b610117610352565b60405161012491906106ac565b60405180910390f35b610135610376565b005b61013f61039e565b60405161014c91906106ac565b60405180910390f35b61015d6103c2565b60405161016a9190610720565b60405180910390f35b5f5f61017d6103e6565b915091508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e690610783565b60405180910390fd5b80156102cd577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083846040518263ffffffff1660e01b815260040161024e91906107b9565b602060405180830381865afa158015610269573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061028d91906107fc565b6102cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c390610871565b60405180910390fd5b5b505050565b6102f282826040516102e59291906104de565b6040518091039020610173565b5050565b6040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b61033f61033a610452565b610173565b565b5f5f9054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f9054906101000a900460ff16155f5f6101000a81548160ff021916908315150217905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f5f9054906101000a900460ff1615610427577f000000000000000000000000000000000000000000000000000000000000000060019150915061044e565b7f00000000000000000000000000000000000000000000000000000000000000005f915091505b9091565b5f60605f5b5f5f1b814914610493578181496040516020016104759291906108e9565b60405160208183030381529060405291508080600101915050610457565b81805190602001209250505090565b5f81905092915050565b828183375f83830152505050565b5f6104c583856104a2565b93506104d28385846104ac565b82840190509392505050565b5f6104ea8284866104ba565b91508190509392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261051f5761051e6104fe565b5b8235905067ffffffffffffffff81111561053c5761053b610502565b5b60208301915083600182028301111561055857610557610506565b5b9250929050565b5f5f60208385031215610575576105746104f6565b5b5f83013567ffffffffffffffff811115610592576105916104fa565b5b61059e8582860161050a565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6105ec826105aa565b6105f681856105b4565b93506106068185602086016105c4565b61060f816105d2565b840191505092915050565b5f6020820190508181035f83015261063281846105e2565b905092915050565b5f8115159050919050565b61064e8161063a565b82525050565b5f6020820190506106675f830184610645565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106968261066d565b9050919050565b6106a68161068c565b82525050565b5f6020820190506106bf5f83018461069d565b92915050565b5f819050919050565b5f6106e86106e36106de8461066d565b6106c5565b61066d565b9050919050565b5f6106f9826106ce565b9050919050565b5f61070a826106ef565b9050919050565b61071a81610700565b82525050565b5f6020820190506107335f830184610711565b92915050565b7f4261746368496e626f783a20696e6163746976652062617463686572000000005f82015250565b5f61076d601c836105b4565b915061077882610739565b602082019050919050565b5f6020820190508181035f83015261079a81610761565b9050919050565b5f819050919050565b6107b3816107a1565b82525050565b5f6020820190506107cc5f8301846107aa565b92915050565b6107db8161063a565b81146107e5575f5ffd5b50565b5f815190506107f6816107d2565b92915050565b5f60208284031215610811576108106104f6565b5b5f61081e848285016107e8565b91505092915050565b7f4261746368496e626f783a20696e76616c6964206261746368000000000000005f82015250565b5f61085b6019836105b4565b915061086682610827565b602082019050919050565b5f6020820190508181035f8301526108888161084f565b9050919050565b5f81519050919050565b5f6108a38261088f565b6108ad81856104a2565b93506108bd8185602086016105c4565b80840191505092915050565b5f819050919050565b6108e36108de826107a1565b6108c9565b82525050565b5f6108f48285610899565b915061090082846108d2565b602082019150819050939250505056fea164736f6c634300081c000a",
}

// BatchInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchInboxMetaData.ABI instead.
var BatchInboxABI = BatchInboxMetaData.ABI

// BatchInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchInboxMetaData.Bin instead.
var BatchInboxBin = BatchInboxMetaData.Bin

// DeployBatchInbox deploys a new Ethereum contract, binding an instance of BatchInbox to it.
func DeployBatchInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _teeBatcher common.Address, _nonTeeBatcher common.Address, _batchAuthenticator common.Address) (common.Address, *types.Transaction, *BatchInbox, error) {
	parsed, err := BatchInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchInboxBin), backend, _teeBatcher, _nonTeeBatcher, _batchAuthenticator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchInbox{BatchInboxCaller: BatchInboxCaller{contract: contract}, BatchInboxTransactor: BatchInboxTransactor{contract: contract}, BatchInboxFilterer: BatchInboxFilterer{contract: contract}}, nil
}

// BatchInbox is an auto generated Go binding around an Ethereum contract.
type BatchInbox struct {
	BatchInboxCaller     // Read-only binding to the contract
	BatchInboxTransactor // Write-only binding to the contract
	BatchInboxFilterer   // Log filterer for contract events
}

// BatchInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchInboxSession struct {
	Contract     *BatchInbox       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchInboxCallerSession struct {
	Contract *BatchInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BatchInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchInboxTransactorSession struct {
	Contract     *BatchInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BatchInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchInboxRaw struct {
	Contract *BatchInbox // Generic contract binding to access the raw methods on
}

// BatchInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchInboxCallerRaw struct {
	Contract *BatchInboxCaller // Generic read-only contract binding to access the raw methods on
}

// BatchInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchInboxTransactorRaw struct {
	Contract *BatchInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchInbox creates a new instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInbox(address common.Address, backend bind.ContractBackend) (*BatchInbox, error) {
	contract, err := bindBatchInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchInbox{BatchInboxCaller: BatchInboxCaller{contract: contract}, BatchInboxTransactor: BatchInboxTransactor{contract: contract}, BatchInboxFilterer: BatchInboxFilterer{contract: contract}}, nil
}

// NewBatchInboxCaller creates a new read-only instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInboxCaller(address common.Address, caller bind.ContractCaller) (*BatchInboxCaller, error) {
	contract, err := bindBatchInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchInboxCaller{contract: contract}, nil
}

// NewBatchInboxTransactor creates a new write-only instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchInboxTransactor, error) {
	contract, err := bindBatchInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchInboxTransactor{contract: contract}, nil
}

// NewBatchInboxFilterer creates a new log filterer instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchInboxFilterer, error) {
	contract, err := bindBatchInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchInboxFilterer{contract: contract}, nil
}

// bindBatchInbox binds a generic wrapper to an already deployed contract.
func bindBatchInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchInbox *BatchInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchInbox.Contract.BatchInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchInbox *BatchInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchInbox.Contract.BatchInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchInbox *BatchInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchInbox.Contract.BatchInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchInbox *BatchInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchInbox *BatchInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchInbox *BatchInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchInbox.Contract.contract.Transact(opts, method, params...)
}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchInbox *BatchInboxCaller) ActiveIsTee(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "activeIsTee")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchInbox *BatchInboxSession) ActiveIsTee() (bool, error) {
	return _BatchInbox.Contract.ActiveIsTee(&_BatchInbox.CallOpts)
}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchInbox *BatchInboxCallerSession) ActiveIsTee() (bool, error) {
	return _BatchInbox.Contract.ActiveIsTee(&_BatchInbox.CallOpts)
}

// BatchAuthenticator is a free data retrieval call binding the contract method 0xe7584573.
//
// Solidity: function batchAuthenticator() view returns(address)
func (_BatchInbox *BatchInboxCaller) BatchAuthenticator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "batchAuthenticator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchAuthenticator is a free data retrieval call binding the contract method 0xe7584573.
//
// Solidity: function batchAuthenticator() view returns(address)
func (_BatchInbox *BatchInboxSession) BatchAuthenticator() (common.Address, error) {
	return _BatchInbox.Contract.BatchAuthenticator(&_BatchInbox.CallOpts)
}

// BatchAuthenticator is a free data retrieval call binding the contract method 0xe7584573.
//
// Solidity: function batchAuthenticator() view returns(address)
func (_BatchInbox *BatchInboxCallerSession) BatchAuthenticator() (common.Address, error) {
	return _BatchInbox.Contract.BatchAuthenticator(&_BatchInbox.CallOpts)
}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchInbox *BatchInboxCaller) NonTeeBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "nonTeeBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchInbox *BatchInboxSession) NonTeeBatcher() (common.Address, error) {
	return _BatchInbox.Contract.NonTeeBatcher(&_BatchInbox.CallOpts)
}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchInbox *BatchInboxCallerSession) NonTeeBatcher() (common.Address, error) {
	return _BatchInbox.Contract.NonTeeBatcher(&_BatchInbox.CallOpts)
}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchInbox *BatchInboxCaller) TeeBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "teeBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchInbox *BatchInboxSession) TeeBatcher() (common.Address, error) {
	return _BatchInbox.Contract.TeeBatcher(&_BatchInbox.CallOpts)
}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchInbox *BatchInboxCallerSession) TeeBatcher() (common.Address, error) {
	return _BatchInbox.Contract.TeeBatcher(&_BatchInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchInbox *BatchInboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchInbox *BatchInboxSession) Version() (string, error) {
	return _BatchInbox.Contract.Version(&_BatchInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchInbox *BatchInboxCallerSession) Version() (string, error) {
	return _BatchInbox.Contract.Version(&_BatchInbox.CallOpts)
}

// PostBlobs is a paid mutator transaction binding the contract method 0x7098ae43.
//
// Solidity: function postBlobs() returns()
func (_BatchInbox *BatchInboxTransactor) PostBlobs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "postBlobs")
}

// PostBlobs is a paid mutator transaction binding the contract method 0x7098ae43.
//
// Solidity: function postBlobs() returns()
func (_BatchInbox *BatchInboxSession) PostBlobs() (*types.Transaction, error) {
	return _BatchInbox.Contract.PostBlobs(&_BatchInbox.TransactOpts)
}

// PostBlobs is a paid mutator transaction binding the contract method 0x7098ae43.
//
// Solidity: function postBlobs() returns()
func (_BatchInbox *BatchInboxTransactorSession) PostBlobs() (*types.Transaction, error) {
	return _BatchInbox.Contract.PostBlobs(&_BatchInbox.TransactOpts)
}

// PostCalldata is a paid mutator transaction binding the contract method 0x1ad40238.
//
// Solidity: function postCalldata(bytes data) returns()
func (_BatchInbox *BatchInboxTransactor) PostCalldata(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "postCalldata", data)
}

// PostCalldata is a paid mutator transaction binding the contract method 0x1ad40238.
//
// Solidity: function postCalldata(bytes data) returns()
func (_BatchInbox *BatchInboxSession) PostCalldata(data []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.PostCalldata(&_BatchInbox.TransactOpts, data)
}

// PostCalldata is a paid mutator transaction binding the contract method 0x1ad40238.
//
// Solidity: function postCalldata(bytes data) returns()
func (_BatchInbox *BatchInboxTransactorSession) PostCalldata(data []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.PostCalldata(&_BatchInbox.TransactOpts, data)
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchInbox *BatchInboxTransactor) SwitchBatcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "switchBatcher")
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchInbox *BatchInboxSession) SwitchBatcher() (*types.Transaction, error) {
	return _BatchInbox.Contract.SwitchBatcher(&_BatchInbox.TransactOpts)
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchInbox *BatchInboxTransactorSession) SwitchBatcher() (*types.Transaction, error) {
	return _BatchInbox.Contract.SwitchBatcher(&_BatchInbox.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BatchInbox *BatchInboxTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BatchInbox.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BatchInbox *BatchInboxSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.Fallback(&_BatchInbox.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BatchInbox *BatchInboxTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.Fallback(&_BatchInbox.TransactOpts, calldata)
}
