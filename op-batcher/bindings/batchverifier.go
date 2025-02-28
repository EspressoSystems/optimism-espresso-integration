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

// BatchVerifierMetaData contains all meta data concerning the BatchVerifier contract.
var BatchVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"__constructor__\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatches\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyBatch\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// BatchVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchVerifierMetaData.ABI instead.
var BatchVerifierABI = BatchVerifierMetaData.ABI

// BatchVerifier is an auto generated Go binding around an Ethereum contract.
type BatchVerifier struct {
	BatchVerifierCaller     // Read-only binding to the contract
	BatchVerifierTransactor // Write-only binding to the contract
	BatchVerifierFilterer   // Log filterer for contract events
}

// BatchVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchVerifierSession struct {
	Contract     *BatchVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchVerifierCallerSession struct {
	Contract *BatchVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BatchVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchVerifierTransactorSession struct {
	Contract     *BatchVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BatchVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchVerifierRaw struct {
	Contract *BatchVerifier // Generic contract binding to access the raw methods on
}

// BatchVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchVerifierCallerRaw struct {
	Contract *BatchVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BatchVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchVerifierTransactorRaw struct {
	Contract *BatchVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchVerifier creates a new instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifier(address common.Address, backend bind.ContractBackend) (*BatchVerifier, error) {
	contract, err := bindBatchVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchVerifier{BatchVerifierCaller: BatchVerifierCaller{contract: contract}, BatchVerifierTransactor: BatchVerifierTransactor{contract: contract}, BatchVerifierFilterer: BatchVerifierFilterer{contract: contract}}, nil
}

// NewBatchVerifierCaller creates a new read-only instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifierCaller(address common.Address, caller bind.ContractCaller) (*BatchVerifierCaller, error) {
	contract, err := bindBatchVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchVerifierCaller{contract: contract}, nil
}

// NewBatchVerifierTransactor creates a new write-only instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchVerifierTransactor, error) {
	contract, err := bindBatchVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchVerifierTransactor{contract: contract}, nil
}

// NewBatchVerifierFilterer creates a new log filterer instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchVerifierFilterer, error) {
	contract, err := bindBatchVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchVerifierFilterer{contract: contract}, nil
}

// bindBatchVerifier binds a generic wrapper to an already deployed contract.
func bindBatchVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchVerifier *BatchVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchVerifier.Contract.BatchVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchVerifier *BatchVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.Contract.BatchVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchVerifier *BatchVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchVerifier.Contract.BatchVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchVerifier *BatchVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchVerifier *BatchVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchVerifier *BatchVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchVerifier.Contract.contract.Transact(opts, method, params...)
}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) view returns(bytes, bytes)
func (_BatchVerifier *BatchVerifierCaller) DecodeAttestationTbs(opts *bind.CallOpts, attestation []byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _BatchVerifier.contract.Call(opts, &out, "decodeAttestationTbs", attestation)

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
func (_BatchVerifier *BatchVerifierSession) DecodeAttestationTbs(attestation []byte) ([]byte, []byte, error) {
	return _BatchVerifier.Contract.DecodeAttestationTbs(&_BatchVerifier.CallOpts, attestation)
}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) view returns(bytes, bytes)
func (_BatchVerifier *BatchVerifierCallerSession) DecodeAttestationTbs(attestation []byte) ([]byte, []byte, error) {
	return _BatchVerifier.Contract.DecodeAttestationTbs(&_BatchVerifier.CallOpts, attestation)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchVerifier *BatchVerifierCaller) EspressoTEEVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchVerifier.contract.Call(opts, &out, "espressoTEEVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchVerifier *BatchVerifierSession) EspressoTEEVerifier() (common.Address, error) {
	return _BatchVerifier.Contract.EspressoTEEVerifier(&_BatchVerifier.CallOpts)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchVerifier *BatchVerifierCallerSession) EspressoTEEVerifier() (common.Address, error) {
	return _BatchVerifier.Contract.EspressoTEEVerifier(&_BatchVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchVerifier *BatchVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchVerifier *BatchVerifierSession) Owner() (common.Address, error) {
	return _BatchVerifier.Contract.Owner(&_BatchVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchVerifier *BatchVerifierCallerSession) Owner() (common.Address, error) {
	return _BatchVerifier.Contract.Owner(&_BatchVerifier.CallOpts)
}

// ValidBatches is a free data retrieval call binding the contract method 0x177db6ae.
//
// Solidity: function validBatches(bytes32 ) view returns(bool)
func (_BatchVerifier *BatchVerifierCaller) ValidBatches(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BatchVerifier.contract.Call(opts, &out, "validBatches", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidBatches is a free data retrieval call binding the contract method 0x177db6ae.
//
// Solidity: function validBatches(bytes32 ) view returns(bool)
func (_BatchVerifier *BatchVerifierSession) ValidBatches(arg0 [32]byte) (bool, error) {
	return _BatchVerifier.Contract.ValidBatches(&_BatchVerifier.CallOpts, arg0)
}

// ValidBatches is a free data retrieval call binding the contract method 0x177db6ae.
//
// Solidity: function validBatches(bytes32 ) view returns(bool)
func (_BatchVerifier *BatchVerifierCallerSession) ValidBatches(arg0 [32]byte) (bool, error) {
	return _BatchVerifier.Contract.ValidBatches(&_BatchVerifier.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchVerifier *BatchVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BatchVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchVerifier *BatchVerifierSession) Version() (string, error) {
	return _BatchVerifier.Contract.Version(&_BatchVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchVerifier *BatchVerifierCallerSession) Version() (string, error) {
	return _BatchVerifier.Contract.Version(&_BatchVerifier.CallOpts)
}

// Constructor is a paid mutator transaction binding the contract method 0x038a609c.
//
// Solidity: function __constructor__(address _espressoTEEVerifier) returns()
func (_BatchVerifier *BatchVerifierTransactor) Constructor(opts *bind.TransactOpts, _espressoTEEVerifier common.Address) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "__constructor__", _espressoTEEVerifier)
}

// Constructor is a paid mutator transaction binding the contract method 0x038a609c.
//
// Solidity: function __constructor__(address _espressoTEEVerifier) returns()
func (_BatchVerifier *BatchVerifierSession) Constructor(_espressoTEEVerifier common.Address) (*types.Transaction, error) {
	return _BatchVerifier.Contract.Constructor(&_BatchVerifier.TransactOpts, _espressoTEEVerifier)
}

// Constructor is a paid mutator transaction binding the contract method 0x038a609c.
//
// Solidity: function __constructor__(address _espressoTEEVerifier) returns()
func (_BatchVerifier *BatchVerifierTransactorSession) Constructor(_espressoTEEVerifier common.Address) (*types.Transaction, error) {
	return _BatchVerifier.Contract.Constructor(&_BatchVerifier.TransactOpts, _espressoTEEVerifier)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchVerifier *BatchVerifierTransactor) RegisterSigner(opts *bind.TransactOpts, attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "registerSigner", attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchVerifier *BatchVerifierSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchVerifier.Contract.RegisterSigner(&_BatchVerifier.TransactOpts, attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchVerifier *BatchVerifierTransactorSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchVerifier.Contract.RegisterSigner(&_BatchVerifier.TransactOpts, attestationTbs, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchVerifier *BatchVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchVerifier *BatchVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchVerifier.Contract.RenounceOwnership(&_BatchVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchVerifier *BatchVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchVerifier.Contract.RenounceOwnership(&_BatchVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchVerifier *BatchVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchVerifier *BatchVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchVerifier.Contract.TransferOwnership(&_BatchVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchVerifier *BatchVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchVerifier.Contract.TransferOwnership(&_BatchVerifier.TransactOpts, newOwner)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0x75d0a7ce.
//
// Solidity: function verifyBatch(bytes32 commitment, bytes signature) returns()
func (_BatchVerifier *BatchVerifierTransactor) VerifyBatch(opts *bind.TransactOpts, commitment [32]byte, signature []byte) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "verifyBatch", commitment, signature)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0x75d0a7ce.
//
// Solidity: function verifyBatch(bytes32 commitment, bytes signature) returns()
func (_BatchVerifier *BatchVerifierSession) VerifyBatch(commitment [32]byte, signature []byte) (*types.Transaction, error) {
	return _BatchVerifier.Contract.VerifyBatch(&_BatchVerifier.TransactOpts, commitment, signature)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0x75d0a7ce.
//
// Solidity: function verifyBatch(bytes32 commitment, bytes signature) returns()
func (_BatchVerifier *BatchVerifierTransactorSession) VerifyBatch(commitment [32]byte, signature []byte) (*types.Transaction, error) {
	return _BatchVerifier.Contract.VerifyBatch(&_BatchVerifier.TransactOpts, commitment, signature)
}

// BatchVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BatchVerifier contract.
type BatchVerifierInitializedIterator struct {
	Event *BatchVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *BatchVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchVerifierInitialized)
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
		it.Event = new(BatchVerifierInitialized)
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
func (it *BatchVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchVerifierInitialized represents a Initialized event raised by the BatchVerifier contract.
type BatchVerifierInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BatchVerifier *BatchVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*BatchVerifierInitializedIterator, error) {

	logs, sub, err := _BatchVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BatchVerifierInitializedIterator{contract: _BatchVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BatchVerifier *BatchVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BatchVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _BatchVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchVerifierInitialized)
				if err := _BatchVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BatchVerifier *BatchVerifierFilterer) ParseInitialized(log types.Log) (*BatchVerifierInitialized, error) {
	event := new(BatchVerifierInitialized)
	if err := _BatchVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BatchVerifier contract.
type BatchVerifierOwnershipTransferredIterator struct {
	Event *BatchVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BatchVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchVerifierOwnershipTransferred)
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
		it.Event = new(BatchVerifierOwnershipTransferred)
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
func (it *BatchVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the BatchVerifier contract.
type BatchVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchVerifier *BatchVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchVerifierOwnershipTransferredIterator{contract: _BatchVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchVerifier *BatchVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatchVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchVerifierOwnershipTransferred)
				if err := _BatchVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BatchVerifier *BatchVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*BatchVerifierOwnershipTransferred, error) {
	event := new(BatchVerifierOwnershipTransferred)
	if err := _BatchVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
