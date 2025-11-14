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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_batchAuthenticator\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchAuthenticator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610f42380380610f428339818101604052810190610031919061022e565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561009957505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b6100d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100cf906102d8565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505060015f5f6101000a81548160ff0219169083151502179055505050506102f6565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101c282610199565b9050919050565b6101d2816101b8565b81146101dc575f5ffd5b50565b5f815190506101ed816101c9565b92915050565b5f6101fd826101b8565b9050919050565b61020d816101f3565b8114610217575f5ffd5b50565b5f8151905061022881610204565b92915050565b5f5f5f6060848603121561024557610244610195565b5b5f610252868287016101df565b9350506020610263868287016101df565b92505060406102748682870161021a565b9150509250925092565b5f82825260208201905092915050565b7f4261746368496e626f783a207a65726f206261746368657200000000000000005f82015250565b5f6102c260188361027e565b91506102cd8261028e565b602082019050919050565b5f6020820190508181035f8301526102ef816102b6565b9050919050565b60805160a05160c051610bf361034f5f395f81816101d3015281816102cd015281816104cc015261062a01525f8181606f015281816105ba015261068f01525f818160950152818161060601526106630152610bf35ff3fe608060405234801561000f575f5ffd5b5060043610610059575f3560e01c80637877a9ed146103c6578063b1bd4285146103e4578063bc347f4714610402578063d909ba7c1461040c578063e75845731461042a5761005a565b5b5f5f5f9054906101000a900460ff16610093577f00000000000000000000000000000000000000000000000000000000000000006100b5565b7f00000000000000000000000000000000000000000000000000000000000000005b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610125576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011c90610712565b60405180910390fd5b5f5f1b5f49146102b0575f5f67ffffffffffffffff81111561014a57610149610730565b5b6040519080825280601f01601f19166020018201604052801561017c5781602001600182028036833780820191505090505b5090505f5f90505b5f5f1b8149146101c6578181496040516020016101a29291906107d8565b604051602081830303815290604052915080806101be90610835565b915050610184565b5f828051906020012090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083826040518263ffffffff1660e01b815260040161022a919061088b565b602060405180830381865afa158015610245573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061026991906108dd565b6102a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029f90610952565b60405180910390fd5b5050506103a4565b5f5f366040516102c19291906109a2565b604051809103902090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083826040518263ffffffff1660e01b8152600401610324919061088b565b602060405180830381865afa15801561033f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061036391906108dd565b6103a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039990610a04565b60405180910390fd5b505b6103c45f366040516103b79291906109a2565b6040518091039020610448565b005b6103ce6105a7565b6040516103db9190610a31565b60405180910390f35b6103ec6105b8565b6040516103f99190610a89565b60405180910390f35b61040a6105dc565b005b610414610604565b6040516104219190610a89565b60405180910390f35b610432610628565b60405161043f9190610afd565b60405180910390f35b5f5f61045261064c565b915091508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bb90610b60565b60405180910390fd5b80156105a2577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083846040518263ffffffff1660e01b8152600401610523919061088b565b602060405180830381865afa15801561053e573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056291906108dd565b6105a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059890610bc8565b60405180910390fd5b5b505050565b5f5f9054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f9054906101000a900460ff16155f5f6101000a81548160ff021916908315150217905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f5f9054906101000a900460ff161561068d577f00000000000000000000000000000000000000000000000000000000000000006001915091506106b4565b7f00000000000000000000000000000000000000000000000000000000000000005f915091505b9091565b5f82825260208201905092915050565b7f4261746368496e626f783a20756e617574686f72697a656420626174636865725f82015250565b5f6106fc6020836106b8565b9150610707826106c8565b602082019050919050565b5f6020820190508181035f830152610729816106f0565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6107898261075d565b6107938185610767565b93506107a3818560208601610771565b80840191505092915050565b5f819050919050565b5f819050919050565b6107d26107cd826107af565b6107b8565b82525050565b5f6107e3828561077f565b91506107ef82846107c1565b6020820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f819050919050565b5f61083f8261082c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610871576108706107ff565b5b600182019050919050565b610885816107af565b82525050565b5f60208201905061089e5f83018461087c565b92915050565b5f5ffd5b5f8115159050919050565b6108bc816108a8565b81146108c6575f5ffd5b50565b5f815190506108d7816108b3565b92915050565b5f602082840312156108f2576108f16108a4565b5b5f6108ff848285016108c9565b91505092915050565b7f496e76616c696420626c6f6220626174636800000000000000000000000000005f82015250565b5f61093c6012836106b8565b915061094782610908565b602082019050919050565b5f6020820190508181035f83015261096981610930565b9050919050565b828183375f83830152505050565b5f6109898385610767565b9350610996838584610970565b82840190509392505050565b5f6109ae82848661097e565b91508190509392505050565b7f496e76616c69642063616c6c64617461206261746368000000000000000000005f82015250565b5f6109ee6016836106b8565b91506109f9826109ba565b602082019050919050565b5f6020820190508181035f830152610a1b816109e2565b9050919050565b610a2b816108a8565b82525050565b5f602082019050610a445f830184610a22565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a7382610a4a565b9050919050565b610a8381610a69565b82525050565b5f602082019050610a9c5f830184610a7a565b92915050565b5f819050919050565b5f610ac5610ac0610abb84610a4a565b610aa2565b610a4a565b9050919050565b5f610ad682610aab565b9050919050565b5f610ae782610acc565b9050919050565b610af781610add565b82525050565b5f602082019050610b105f830184610aee565b92915050565b7f4261746368496e626f783a20696e6163746976652062617463686572000000005f82015250565b5f610b4a601c836106b8565b9150610b5582610b16565b602082019050919050565b5f6020820190508181035f830152610b7781610b3e565b9050919050565b7f4261746368496e626f783a20696e76616c6964206261746368000000000000005f82015250565b5f610bb26019836106b8565b9150610bbd82610b7e565b602082019050919050565b5f6020820190508181035f830152610bdf81610ba6565b905091905056fea164736f6c634300081c000a",
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
