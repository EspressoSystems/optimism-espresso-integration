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

// NitroValidatorPtrs is an auto generated low-level Go binding around an user-defined struct.
type NitroValidatorPtrs struct {
	ModuleID  *big.Int
	Timestamp uint64
	Digest    *big.Int
	Pcrs      []*big.Int
	Cert      *big.Int
	Cabundle  []*big.Int
	PublicKey *big.Int
	UserData  *big.Int
	Nonce     *big.Int
}

// BatchInboxMetaData contains all meta data concerning the BatchInbox contract.
var BatchInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ATTESTATION_DIGEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ATTESTATION_TBS_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CABUNDLE_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CERTIFICATE_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DIGEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_AGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE_ID_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NONCE_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PCRS_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBLIC_KEY_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMESTAMP_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USER_DATA_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"__constructor__\",\"inputs\":[{\"name\":\"certManager\",\"type\":\"address\",\"internalType\":\"contractCertManager\"},{\"name\":\"preApprovedBatcherKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attestedBatchers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"certManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"preApprovedBatcherKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerPCR0\",\"inputs\":[{\"name\":\"pcr0\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitBatch\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validPCR0s\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateAttestation\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNitroValidator.Ptrs\",\"components\":[{\"name\":\"moduleID\",\"type\":\"uint256\",\"internalType\":\"CborElement\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"digest\",\"type\":\"uint256\",\"internalType\":\"CborElement\"},{\"name\":\"pcrs\",\"type\":\"uint256[]\",\"internalType\":\"CborElement[]\"},{\"name\":\"cert\",\"type\":\"uint256\",\"internalType\":\"CborElement\"},{\"name\":\"cabundle\",\"type\":\"uint256[]\",\"internalType\":\"CborElement[]\"},{\"name\":\"publicKey\",\"type\":\"uint256\",\"internalType\":\"CborElement\"},{\"name\":\"userData\",\"type\":\"uint256\",\"internalType\":\"CborElement\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"CborElement\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// BatchInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchInboxMetaData.ABI instead.
var BatchInboxABI = BatchInboxMetaData.ABI

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

// ATTESTATIONDIGEST is a free data retrieval call binding the contract method 0x3893af6d.
//
// Solidity: function ATTESTATION_DIGEST() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) ATTESTATIONDIGEST(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "ATTESTATION_DIGEST")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ATTESTATIONDIGEST is a free data retrieval call binding the contract method 0x3893af6d.
//
// Solidity: function ATTESTATION_DIGEST() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) ATTESTATIONDIGEST() ([32]byte, error) {
	return _BatchInbox.Contract.ATTESTATIONDIGEST(&_BatchInbox.CallOpts)
}

// ATTESTATIONDIGEST is a free data retrieval call binding the contract method 0x3893af6d.
//
// Solidity: function ATTESTATION_DIGEST() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) ATTESTATIONDIGEST() ([32]byte, error) {
	return _BatchInbox.Contract.ATTESTATIONDIGEST(&_BatchInbox.CallOpts)
}

// ATTESTATIONTBSPREFIX is a free data retrieval call binding the contract method 0x2d4bad8a.
//
// Solidity: function ATTESTATION_TBS_PREFIX() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) ATTESTATIONTBSPREFIX(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "ATTESTATION_TBS_PREFIX")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ATTESTATIONTBSPREFIX is a free data retrieval call binding the contract method 0x2d4bad8a.
//
// Solidity: function ATTESTATION_TBS_PREFIX() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) ATTESTATIONTBSPREFIX() ([32]byte, error) {
	return _BatchInbox.Contract.ATTESTATIONTBSPREFIX(&_BatchInbox.CallOpts)
}

// ATTESTATIONTBSPREFIX is a free data retrieval call binding the contract method 0x2d4bad8a.
//
// Solidity: function ATTESTATION_TBS_PREFIX() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) ATTESTATIONTBSPREFIX() ([32]byte, error) {
	return _BatchInbox.Contract.ATTESTATIONTBSPREFIX(&_BatchInbox.CallOpts)
}

// CABUNDLEKEY is a free data retrieval call binding the contract method 0x9cc3eb48.
//
// Solidity: function CABUNDLE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) CABUNDLEKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "CABUNDLE_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CABUNDLEKEY is a free data retrieval call binding the contract method 0x9cc3eb48.
//
// Solidity: function CABUNDLE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) CABUNDLEKEY() ([32]byte, error) {
	return _BatchInbox.Contract.CABUNDLEKEY(&_BatchInbox.CallOpts)
}

// CABUNDLEKEY is a free data retrieval call binding the contract method 0x9cc3eb48.
//
// Solidity: function CABUNDLE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) CABUNDLEKEY() ([32]byte, error) {
	return _BatchInbox.Contract.CABUNDLEKEY(&_BatchInbox.CallOpts)
}

// CERTIFICATEKEY is a free data retrieval call binding the contract method 0xae951149.
//
// Solidity: function CERTIFICATE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) CERTIFICATEKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "CERTIFICATE_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CERTIFICATEKEY is a free data retrieval call binding the contract method 0xae951149.
//
// Solidity: function CERTIFICATE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) CERTIFICATEKEY() ([32]byte, error) {
	return _BatchInbox.Contract.CERTIFICATEKEY(&_BatchInbox.CallOpts)
}

// CERTIFICATEKEY is a free data retrieval call binding the contract method 0xae951149.
//
// Solidity: function CERTIFICATE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) CERTIFICATEKEY() ([32]byte, error) {
	return _BatchInbox.Contract.CERTIFICATEKEY(&_BatchInbox.CallOpts)
}

// DIGESTKEY is a free data retrieval call binding the contract method 0x6be1e68b.
//
// Solidity: function DIGEST_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) DIGESTKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "DIGEST_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIGESTKEY is a free data retrieval call binding the contract method 0x6be1e68b.
//
// Solidity: function DIGEST_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) DIGESTKEY() ([32]byte, error) {
	return _BatchInbox.Contract.DIGESTKEY(&_BatchInbox.CallOpts)
}

// DIGESTKEY is a free data retrieval call binding the contract method 0x6be1e68b.
//
// Solidity: function DIGEST_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) DIGESTKEY() ([32]byte, error) {
	return _BatchInbox.Contract.DIGESTKEY(&_BatchInbox.CallOpts)
}

// MAXAGE is a free data retrieval call binding the contract method 0x0dcaeaf2.
//
// Solidity: function MAX_AGE() view returns(uint256)
func (_BatchInbox *BatchInboxCaller) MAXAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "MAX_AGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAGE is a free data retrieval call binding the contract method 0x0dcaeaf2.
//
// Solidity: function MAX_AGE() view returns(uint256)
func (_BatchInbox *BatchInboxSession) MAXAGE() (*big.Int, error) {
	return _BatchInbox.Contract.MAXAGE(&_BatchInbox.CallOpts)
}

// MAXAGE is a free data retrieval call binding the contract method 0x0dcaeaf2.
//
// Solidity: function MAX_AGE() view returns(uint256)
func (_BatchInbox *BatchInboxCallerSession) MAXAGE() (*big.Int, error) {
	return _BatchInbox.Contract.MAXAGE(&_BatchInbox.CallOpts)
}

// MODULEIDKEY is a free data retrieval call binding the contract method 0x9adb2d68.
//
// Solidity: function MODULE_ID_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) MODULEIDKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "MODULE_ID_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODULEIDKEY is a free data retrieval call binding the contract method 0x9adb2d68.
//
// Solidity: function MODULE_ID_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) MODULEIDKEY() ([32]byte, error) {
	return _BatchInbox.Contract.MODULEIDKEY(&_BatchInbox.CallOpts)
}

// MODULEIDKEY is a free data retrieval call binding the contract method 0x9adb2d68.
//
// Solidity: function MODULE_ID_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) MODULEIDKEY() ([32]byte, error) {
	return _BatchInbox.Contract.MODULEIDKEY(&_BatchInbox.CallOpts)
}

// NONCEKEY is a free data retrieval call binding the contract method 0x6378aad5.
//
// Solidity: function NONCE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) NONCEKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "NONCE_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NONCEKEY is a free data retrieval call binding the contract method 0x6378aad5.
//
// Solidity: function NONCE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) NONCEKEY() ([32]byte, error) {
	return _BatchInbox.Contract.NONCEKEY(&_BatchInbox.CallOpts)
}

// NONCEKEY is a free data retrieval call binding the contract method 0x6378aad5.
//
// Solidity: function NONCE_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) NONCEKEY() ([32]byte, error) {
	return _BatchInbox.Contract.NONCEKEY(&_BatchInbox.CallOpts)
}

// PCRSKEY is a free data retrieval call binding the contract method 0xb22bed7e.
//
// Solidity: function PCRS_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) PCRSKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "PCRS_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCRSKEY is a free data retrieval call binding the contract method 0xb22bed7e.
//
// Solidity: function PCRS_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) PCRSKEY() ([32]byte, error) {
	return _BatchInbox.Contract.PCRSKEY(&_BatchInbox.CallOpts)
}

// PCRSKEY is a free data retrieval call binding the contract method 0xb22bed7e.
//
// Solidity: function PCRS_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) PCRSKEY() ([32]byte, error) {
	return _BatchInbox.Contract.PCRSKEY(&_BatchInbox.CallOpts)
}

// PUBLICKEYKEY is a free data retrieval call binding the contract method 0xe8b6d3fe.
//
// Solidity: function PUBLIC_KEY_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) PUBLICKEYKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "PUBLIC_KEY_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBLICKEYKEY is a free data retrieval call binding the contract method 0xe8b6d3fe.
//
// Solidity: function PUBLIC_KEY_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) PUBLICKEYKEY() ([32]byte, error) {
	return _BatchInbox.Contract.PUBLICKEYKEY(&_BatchInbox.CallOpts)
}

// PUBLICKEYKEY is a free data retrieval call binding the contract method 0xe8b6d3fe.
//
// Solidity: function PUBLIC_KEY_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) PUBLICKEYKEY() ([32]byte, error) {
	return _BatchInbox.Contract.PUBLICKEYKEY(&_BatchInbox.CallOpts)
}

// TIMESTAMPKEY is a free data retrieval call binding the contract method 0xe0a655ff.
//
// Solidity: function TIMESTAMP_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) TIMESTAMPKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "TIMESTAMP_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TIMESTAMPKEY is a free data retrieval call binding the contract method 0xe0a655ff.
//
// Solidity: function TIMESTAMP_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) TIMESTAMPKEY() ([32]byte, error) {
	return _BatchInbox.Contract.TIMESTAMPKEY(&_BatchInbox.CallOpts)
}

// TIMESTAMPKEY is a free data retrieval call binding the contract method 0xe0a655ff.
//
// Solidity: function TIMESTAMP_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) TIMESTAMPKEY() ([32]byte, error) {
	return _BatchInbox.Contract.TIMESTAMPKEY(&_BatchInbox.CallOpts)
}

// USERDATAKEY is a free data retrieval call binding the contract method 0xcebf08d7.
//
// Solidity: function USER_DATA_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCaller) USERDATAKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "USER_DATA_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// USERDATAKEY is a free data retrieval call binding the contract method 0xcebf08d7.
//
// Solidity: function USER_DATA_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxSession) USERDATAKEY() ([32]byte, error) {
	return _BatchInbox.Contract.USERDATAKEY(&_BatchInbox.CallOpts)
}

// USERDATAKEY is a free data retrieval call binding the contract method 0xcebf08d7.
//
// Solidity: function USER_DATA_KEY() view returns(bytes32)
func (_BatchInbox *BatchInboxCallerSession) USERDATAKEY() ([32]byte, error) {
	return _BatchInbox.Contract.USERDATAKEY(&_BatchInbox.CallOpts)
}

// AttestedBatchers is a free data retrieval call binding the contract method 0xf2d8ed17.
//
// Solidity: function attestedBatchers(address ) view returns(bool)
func (_BatchInbox *BatchInboxCaller) AttestedBatchers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "attestedBatchers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestedBatchers is a free data retrieval call binding the contract method 0xf2d8ed17.
//
// Solidity: function attestedBatchers(address ) view returns(bool)
func (_BatchInbox *BatchInboxSession) AttestedBatchers(arg0 common.Address) (bool, error) {
	return _BatchInbox.Contract.AttestedBatchers(&_BatchInbox.CallOpts, arg0)
}

// AttestedBatchers is a free data retrieval call binding the contract method 0xf2d8ed17.
//
// Solidity: function attestedBatchers(address ) view returns(bool)
func (_BatchInbox *BatchInboxCallerSession) AttestedBatchers(arg0 common.Address) (bool, error) {
	return _BatchInbox.Contract.AttestedBatchers(&_BatchInbox.CallOpts, arg0)
}

// CertManager is a free data retrieval call binding the contract method 0x739e8484.
//
// Solidity: function certManager() view returns(address)
func (_BatchInbox *BatchInboxCaller) CertManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "certManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CertManager is a free data retrieval call binding the contract method 0x739e8484.
//
// Solidity: function certManager() view returns(address)
func (_BatchInbox *BatchInboxSession) CertManager() (common.Address, error) {
	return _BatchInbox.Contract.CertManager(&_BatchInbox.CallOpts)
}

// CertManager is a free data retrieval call binding the contract method 0x739e8484.
//
// Solidity: function certManager() view returns(address)
func (_BatchInbox *BatchInboxCallerSession) CertManager() (common.Address, error) {
	return _BatchInbox.Contract.CertManager(&_BatchInbox.CallOpts)
}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) pure returns(bytes attestationTbs, bytes signature)
func (_BatchInbox *BatchInboxCaller) DecodeAttestationTbs(opts *bind.CallOpts, attestation []byte) (struct {
	AttestationTbs []byte
	Signature      []byte
}, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "decodeAttestationTbs", attestation)

	outstruct := new(struct {
		AttestationTbs []byte
		Signature      []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttestationTbs = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Signature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) pure returns(bytes attestationTbs, bytes signature)
func (_BatchInbox *BatchInboxSession) DecodeAttestationTbs(attestation []byte) (struct {
	AttestationTbs []byte
	Signature      []byte
}, error) {
	return _BatchInbox.Contract.DecodeAttestationTbs(&_BatchInbox.CallOpts, attestation)
}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) pure returns(bytes attestationTbs, bytes signature)
func (_BatchInbox *BatchInboxCallerSession) DecodeAttestationTbs(attestation []byte) (struct {
	AttestationTbs []byte
	Signature      []byte
}, error) {
	return _BatchInbox.Contract.DecodeAttestationTbs(&_BatchInbox.CallOpts, attestation)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchInbox *BatchInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchInbox *BatchInboxSession) Owner() (common.Address, error) {
	return _BatchInbox.Contract.Owner(&_BatchInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchInbox *BatchInboxCallerSession) Owner() (common.Address, error) {
	return _BatchInbox.Contract.Owner(&_BatchInbox.CallOpts)
}

// ValidPCR0s is a free data retrieval call binding the contract method 0x295840d9.
//
// Solidity: function validPCR0s(bytes32 ) view returns(bool)
func (_BatchInbox *BatchInboxCaller) ValidPCR0s(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "validPCR0s", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidPCR0s is a free data retrieval call binding the contract method 0x295840d9.
//
// Solidity: function validPCR0s(bytes32 ) view returns(bool)
func (_BatchInbox *BatchInboxSession) ValidPCR0s(arg0 [32]byte) (bool, error) {
	return _BatchInbox.Contract.ValidPCR0s(&_BatchInbox.CallOpts, arg0)
}

// ValidPCR0s is a free data retrieval call binding the contract method 0x295840d9.
//
// Solidity: function validPCR0s(bytes32 ) view returns(bool)
func (_BatchInbox *BatchInboxCallerSession) ValidPCR0s(arg0 [32]byte) (bool, error) {
	return _BatchInbox.Contract.ValidPCR0s(&_BatchInbox.CallOpts, arg0)
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

// Constructor is a paid mutator transaction binding the contract method 0x63f9288d.
//
// Solidity: function __constructor__(address certManager, address preApprovedBatcherKey) returns()
func (_BatchInbox *BatchInboxTransactor) Constructor(opts *bind.TransactOpts, certManager common.Address, preApprovedBatcherKey common.Address) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "__constructor__", certManager, preApprovedBatcherKey)
}

// Constructor is a paid mutator transaction binding the contract method 0x63f9288d.
//
// Solidity: function __constructor__(address certManager, address preApprovedBatcherKey) returns()
func (_BatchInbox *BatchInboxSession) Constructor(certManager common.Address, preApprovedBatcherKey common.Address) (*types.Transaction, error) {
	return _BatchInbox.Contract.Constructor(&_BatchInbox.TransactOpts, certManager, preApprovedBatcherKey)
}

// Constructor is a paid mutator transaction binding the contract method 0x63f9288d.
//
// Solidity: function __constructor__(address certManager, address preApprovedBatcherKey) returns()
func (_BatchInbox *BatchInboxTransactorSession) Constructor(certManager common.Address, preApprovedBatcherKey common.Address) (*types.Transaction, error) {
	return _BatchInbox.Contract.Constructor(&_BatchInbox.TransactOpts, certManager, preApprovedBatcherKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address preApprovedBatcherKey) returns()
func (_BatchInbox *BatchInboxTransactor) Initialize(opts *bind.TransactOpts, preApprovedBatcherKey common.Address) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "initialize", preApprovedBatcherKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address preApprovedBatcherKey) returns()
func (_BatchInbox *BatchInboxSession) Initialize(preApprovedBatcherKey common.Address) (*types.Transaction, error) {
	return _BatchInbox.Contract.Initialize(&_BatchInbox.TransactOpts, preApprovedBatcherKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address preApprovedBatcherKey) returns()
func (_BatchInbox *BatchInboxTransactorSession) Initialize(preApprovedBatcherKey common.Address) (*types.Transaction, error) {
	return _BatchInbox.Contract.Initialize(&_BatchInbox.TransactOpts, preApprovedBatcherKey)
}

// RegisterPCR0 is a paid mutator transaction binding the contract method 0x2c68fa02.
//
// Solidity: function registerPCR0(bytes pcr0) returns()
func (_BatchInbox *BatchInboxTransactor) RegisterPCR0(opts *bind.TransactOpts, pcr0 []byte) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "registerPCR0", pcr0)
}

// RegisterPCR0 is a paid mutator transaction binding the contract method 0x2c68fa02.
//
// Solidity: function registerPCR0(bytes pcr0) returns()
func (_BatchInbox *BatchInboxSession) RegisterPCR0(pcr0 []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.RegisterPCR0(&_BatchInbox.TransactOpts, pcr0)
}

// RegisterPCR0 is a paid mutator transaction binding the contract method 0x2c68fa02.
//
// Solidity: function registerPCR0(bytes pcr0) returns()
func (_BatchInbox *BatchInboxTransactorSession) RegisterPCR0(pcr0 []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.RegisterPCR0(&_BatchInbox.TransactOpts, pcr0)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchInbox *BatchInboxTransactor) RegisterSigner(opts *bind.TransactOpts, attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "registerSigner", attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchInbox *BatchInboxSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.RegisterSigner(&_BatchInbox.TransactOpts, attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchInbox *BatchInboxTransactorSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.RegisterSigner(&_BatchInbox.TransactOpts, attestationTbs, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchInbox *BatchInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchInbox *BatchInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchInbox.Contract.RenounceOwnership(&_BatchInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchInbox *BatchInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchInbox.Contract.RenounceOwnership(&_BatchInbox.TransactOpts)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xc9486c8b.
//
// Solidity: function submitBatch(bytes commitment, bytes _signature) returns()
func (_BatchInbox *BatchInboxTransactor) SubmitBatch(opts *bind.TransactOpts, commitment []byte, _signature []byte) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "submitBatch", commitment, _signature)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xc9486c8b.
//
// Solidity: function submitBatch(bytes commitment, bytes _signature) returns()
func (_BatchInbox *BatchInboxSession) SubmitBatch(commitment []byte, _signature []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.SubmitBatch(&_BatchInbox.TransactOpts, commitment, _signature)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xc9486c8b.
//
// Solidity: function submitBatch(bytes commitment, bytes _signature) returns()
func (_BatchInbox *BatchInboxTransactorSession) SubmitBatch(commitment []byte, _signature []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.SubmitBatch(&_BatchInbox.TransactOpts, commitment, _signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchInbox *BatchInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchInbox *BatchInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchInbox.Contract.TransferOwnership(&_BatchInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchInbox *BatchInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchInbox.Contract.TransferOwnership(&_BatchInbox.TransactOpts, newOwner)
}

// ValidateAttestation is a paid mutator transaction binding the contract method 0x05f7aead.
//
// Solidity: function validateAttestation(bytes attestationTbs, bytes signature) returns((uint256,uint64,uint256,uint256[],uint256,uint256[],uint256,uint256,uint256))
func (_BatchInbox *BatchInboxTransactor) ValidateAttestation(opts *bind.TransactOpts, attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "validateAttestation", attestationTbs, signature)
}

// ValidateAttestation is a paid mutator transaction binding the contract method 0x05f7aead.
//
// Solidity: function validateAttestation(bytes attestationTbs, bytes signature) returns((uint256,uint64,uint256,uint256[],uint256,uint256[],uint256,uint256,uint256))
func (_BatchInbox *BatchInboxSession) ValidateAttestation(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.ValidateAttestation(&_BatchInbox.TransactOpts, attestationTbs, signature)
}

// ValidateAttestation is a paid mutator transaction binding the contract method 0x05f7aead.
//
// Solidity: function validateAttestation(bytes attestationTbs, bytes signature) returns((uint256,uint64,uint256,uint256[],uint256,uint256[],uint256,uint256,uint256))
func (_BatchInbox *BatchInboxTransactorSession) ValidateAttestation(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.ValidateAttestation(&_BatchInbox.TransactOpts, attestationTbs, signature)
}

// BatchInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BatchInbox contract.
type BatchInboxInitializedIterator struct {
	Event *BatchInboxInitialized // Event containing the contract specifics and raw log

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
func (it *BatchInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchInboxInitialized)
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
		it.Event = new(BatchInboxInitialized)
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
func (it *BatchInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchInboxInitialized represents a Initialized event raised by the BatchInbox contract.
type BatchInboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BatchInbox *BatchInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*BatchInboxInitializedIterator, error) {

	logs, sub, err := _BatchInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BatchInboxInitializedIterator{contract: _BatchInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BatchInbox *BatchInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BatchInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _BatchInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchInboxInitialized)
				if err := _BatchInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BatchInbox *BatchInboxFilterer) ParseInitialized(log types.Log) (*BatchInboxInitialized, error) {
	event := new(BatchInboxInitialized)
	if err := _BatchInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BatchInbox contract.
type BatchInboxOwnershipTransferredIterator struct {
	Event *BatchInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BatchInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchInboxOwnershipTransferred)
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
		it.Event = new(BatchInboxOwnershipTransferred)
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
func (it *BatchInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchInboxOwnershipTransferred represents a OwnershipTransferred event raised by the BatchInbox contract.
type BatchInboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchInbox *BatchInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchInboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchInbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchInboxOwnershipTransferredIterator{contract: _BatchInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchInbox *BatchInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatchInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchInbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchInboxOwnershipTransferred)
				if err := _BatchInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BatchInbox *BatchInboxFilterer) ParseOwnershipTransferred(log types.Log) (*BatchInboxOwnershipTransferred, error) {
	event := new(BatchInboxOwnershipTransferred)
	if err := _BatchInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
