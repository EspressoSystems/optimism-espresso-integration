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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_batchAuthenticator\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchAuthenticator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postBlobs\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"postCalldata\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610c2f380380610c2f8339818101604052810190610031919061022e565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561009957505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b6100d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100cf906102d8565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505060015f5f6101000a81548160ff0219169083151502179055505050506102f6565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101c282610199565b9050919050565b6101d2816101b8565b81146101dc575f5ffd5b50565b5f815190506101ed816101c9565b92915050565b5f6101fd826101b8565b9050919050565b61020d816101f3565b8114610217575f5ffd5b50565b5f8151905061022881610204565b92915050565b5f5f5f6060848603121561024557610244610195565b5b5f610252868287016101df565b9350506020610263868287016101df565b92505060406102748682870161021a565b9150509250925092565b5f82825260208201905092915050565b7f4261746368496e626f783a207a65726f206261746368657200000000000000005f82015250565b5f6102c260188361027e565b91506102cd8261028e565b602082019050919050565b5f6020820190508181035f8301526102ef816102b6565b9050919050565b60805160a05160c0516108fa6103355f395f818161024201526102e801525f81816101d2015261045601525f818161021e015261042a01526108fa5ff3fe608060405234801561000f575f5ffd5b5060043610610086575f3560e01c8063b1bd428511610059578063b1bd4285146100ec578063bc347f471461010a578063d909ba7c14610114578063e75845731461013257610086565b80631ad402381461008a57806354fd4d50146100a65780637098ae43146100c45780637877a9ed146100ce575b5f5ffd5b6100a4600480360381019061009f91906104e8565b610150565b005b6100ae610174565b6040516100bb91906105a3565b60405180910390f35b6100cc6101ad565b005b6100d66101bf565b6040516100e391906105dd565b60405180910390f35b6100f46101d0565b6040516101019190610635565b60405180910390f35b6101126101f4565b005b61011c61021c565b6040516101299190610635565b60405180910390f35b61013a610240565b60405161014791906106a9565b60405180910390f35b61017082826040516101639291906106fe565b6040518091039020610264565b5050565b6040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6101bd6101b86103c3565b610264565b565b5f5f9054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f9054906101000a900460ff16155f5f6101000a81548160ff021916908315150217905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f61026e610413565b915091508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d790610760565b60405180910390fd5b80156103be577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083846040518263ffffffff1660e01b815260040161033f9190610796565b602060405180830381865afa15801561035a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061037e91906107d9565b6103bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b49061084e565b60405180910390fd5b5b505050565b5f60605f5b5f5f1b814914610404578181496040516020016103e69291906108c6565b604051602081830303815290604052915080806001019150506103c8565b81805190602001209250505090565b5f5f5f5f9054906101000a900460ff1615610454577f000000000000000000000000000000000000000000000000000000000000000060019150915061047b565b7f00000000000000000000000000000000000000000000000000000000000000005f915091505b9091565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126104a8576104a7610487565b5b8235905067ffffffffffffffff8111156104c5576104c461048b565b5b6020830191508360018202830111156104e1576104e061048f565b5b9250929050565b5f5f602083850312156104fe576104fd61047f565b5b5f83013567ffffffffffffffff81111561051b5761051a610483565b5b61052785828601610493565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61057582610533565b61057f818561053d565b935061058f81856020860161054d565b6105988161055b565b840191505092915050565b5f6020820190508181035f8301526105bb818461056b565b905092915050565b5f8115159050919050565b6105d7816105c3565b82525050565b5f6020820190506105f05f8301846105ce565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61061f826105f6565b9050919050565b61062f81610615565b82525050565b5f6020820190506106485f830184610626565b92915050565b5f819050919050565b5f61067161066c610667846105f6565b61064e565b6105f6565b9050919050565b5f61068282610657565b9050919050565b5f61069382610678565b9050919050565b6106a381610689565b82525050565b5f6020820190506106bc5f83018461069a565b92915050565b5f81905092915050565b828183375f83830152505050565b5f6106e583856106c2565b93506106f28385846106cc565b82840190509392505050565b5f61070a8284866106da565b91508190509392505050565b7f4261746368496e626f783a20696e6163746976652062617463686572000000005f82015250565b5f61074a601c8361053d565b915061075582610716565b602082019050919050565b5f6020820190508181035f8301526107778161073e565b9050919050565b5f819050919050565b6107908161077e565b82525050565b5f6020820190506107a95f830184610787565b92915050565b6107b8816105c3565b81146107c2575f5ffd5b50565b5f815190506107d3816107af565b92915050565b5f602082840312156107ee576107ed61047f565b5b5f6107fb848285016107c5565b91505092915050565b7f4261746368496e626f783a20696e76616c6964206261746368000000000000005f82015250565b5f61083860198361053d565b915061084382610804565b602082019050919050565b5f6020820190508181035f8301526108658161082c565b9050919050565b5f81519050919050565b5f6108808261086c565b61088a81856106c2565b935061089a81856020860161054d565b80840191505092915050565b5f819050919050565b6108c06108bb8261077e565b6108a6565b82525050565b5f6108d18285610876565b91506108dd82846108af565b602082019150819050939250505056fea164736f6c634300081c000a",
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
