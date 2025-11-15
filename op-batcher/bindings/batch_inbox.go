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
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610c86380380610c868339818101604052810190610031919061022e565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561009957505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b6100d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100cf906102d8565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505060015f5f6101000a81548160ff0219169083151502179055505050506102f6565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101c282610199565b9050919050565b6101d2816101b8565b81146101dc575f5ffd5b50565b5f815190506101ed816101c9565b92915050565b5f6101fd826101b8565b9050919050565b61020d816101f3565b8114610217575f5ffd5b50565b5f8151905061022881610204565b92915050565b5f5f5f6060848603121561024557610244610195565b5b5f610252868287016101df565b9350506020610263868287016101df565b92505060406102748682870161021a565b9150509250925092565b5f82825260208201905092915050565b7f4261746368496e626f783a207a65726f206261746368657200000000000000005f82015250565b5f6102c260188361027e565b91506102cd8261028e565b602082019050919050565b5f6020820190508181035f8301526102ef816102b6565b9050919050565b60805160a05160c05161094c61033a5f395f81816101e6015281816102e001526104bf01525f8181606f015261044f01525f81816095015261049b015261094c5ff3fe608060405234801561000f575f5ffd5b5060043610610059575f3560e01c80637877a9ed146103ba578063b1bd4285146103d8578063bc347f47146103f6578063d909ba7c14610400578063e75845731461041e5761005a565b5b5f5f5f9054906101000a900460ff16610093577f00000000000000000000000000000000000000000000000000000000000000006100b5565b7f00000000000000000000000000000000000000000000000000000000000000005b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610125576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011c9061053b565b60405180910390fd5b5f5f9054906101000a900460ff16156103b8575f5f1b5f49146102c3575f5f67ffffffffffffffff81111561015d5761015c610559565b5b6040519080825280601f01601f19166020018201604052801561018f5781602001600182028036833780820191505090505b5090505f5f90505b5f5f1b8149146101d9578181496040516020016101b5929190610601565b604051602081830303815290604052915080806101d19061065e565b915050610197565b5f828051906020012090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083826040518263ffffffff1660e01b815260040161023d91906106b4565b602060405180830381865afa158015610258573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061027c9190610706565b6102bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b29061077b565b60405180910390fd5b5050506103b7565b5f5f366040516102d49291906107cb565b604051809103902090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083826040518263ffffffff1660e01b815260040161033791906106b4565b602060405180830381865afa158015610352573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103769190610706565b6103b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ac9061082d565b60405180910390fd5b505b5b005b6103c261043c565b6040516103cf919061085a565b60405180910390f35b6103e061044d565b6040516103ed91906108b2565b60405180910390f35b6103fe610471565b005b610408610499565b60405161041591906108b2565b60405180910390f35b6104266104bd565b6040516104339190610926565b60405180910390f35b5f5f9054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f9054906101000a900460ff16155f5f6101000a81548160ff021916908315150217905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f82825260208201905092915050565b7f4261746368496e626f783a20756e617574686f72697a656420626174636865725f82015250565b5f6105256020836104e1565b9150610530826104f1565b602082019050919050565b5f6020820190508181035f83015261055281610519565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6105b282610586565b6105bc8185610590565b93506105cc81856020860161059a565b80840191505092915050565b5f819050919050565b5f819050919050565b6105fb6105f6826105d8565b6105e1565b82525050565b5f61060c82856105a8565b915061061882846105ea565b6020820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f819050919050565b5f61066882610655565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361069a57610699610628565b5b600182019050919050565b6106ae816105d8565b82525050565b5f6020820190506106c75f8301846106a5565b92915050565b5f5ffd5b5f8115159050919050565b6106e5816106d1565b81146106ef575f5ffd5b50565b5f81519050610700816106dc565b92915050565b5f6020828403121561071b5761071a6106cd565b5b5f610728848285016106f2565b91505092915050565b7f496e76616c696420626c6f6220626174636800000000000000000000000000005f82015250565b5f6107656012836104e1565b915061077082610731565b602082019050919050565b5f6020820190508181035f83015261079281610759565b9050919050565b828183375f83830152505050565b5f6107b28385610590565b93506107bf838584610799565b82840190509392505050565b5f6107d78284866107a7565b91508190509392505050565b7f496e76616c69642063616c6c64617461206261746368000000000000000000005f82015250565b5f6108176016836104e1565b9150610822826107e3565b602082019050919050565b5f6020820190508181035f8301526108448161080b565b9050919050565b610854816106d1565b82525050565b5f60208201905061086d5f83018461084b565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61089c82610873565b9050919050565b6108ac81610892565b82525050565b5f6020820190506108c55f8301846108a3565b92915050565b5f819050919050565b5f6108ee6108e96108e484610873565b6108cb565b610873565b9050919050565b5f6108ff826108d4565b9050919050565b5f610910826108f5565b9050919050565b61092081610906565b82525050565b5f6020820190506109395f830184610917565b9291505056fea164736f6c634300081c000a",
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
