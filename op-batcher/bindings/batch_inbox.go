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
<<<<<<< Updated upstream
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_batchAuthenticator\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60a0604052348015600e575f5ffd5b506040516103e33803806103e3833981016040819052602b91603b565b6001600160a01b03166080526066565b5f60208284031215604a575f5ffd5b81516001600160a01b0381168114605f575f5ffd5b9392505050565b60805161035f6100845f395f8181609d01526101d0015261035f5ff3fe608060405234801561000f575f5ffd5b505f491561018857604080515f80825260208201909252905b804915610067578181496040516020016100439291906102b4565b6040516020818303038152906040529150808061005f906102ce565b915050610028565b815160208301206040517ff81f2083000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063f81f208390602401602060405180830381865afa1580156100f7573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061011b919061032a565b610186576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e76616c696420626c6f62206261746368000000000000000000000000000060448201526064015b60405180910390fd5b005b5f5f36604051610199929190610350565b6040519081900381207ff81f20830000000000000000000000000000000000000000000000000000000082526004820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063f81f208390602401602060405180830381865afa15801561022a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061024e919061032a565b610186576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c69642063616c6c6461746120626174636800000000000000000000604482015260640161017d565b5f83518060208601845e9190910191825250602001919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610323577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5060010190565b5f6020828403121561033a575f5ffd5b81518015158114610349575f5ffd5b9392505050565b818382375f910190815291905056",
=======
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_batchAuthenticator\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchAuthenticator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBatchAuthenticator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051610f9b380380610f9b8339818101604052810190610031919061029d565b61004d61004261013c60201b60201c565b61014360201b60201c565b5f8273ffffffffffffffffffffffffffffffffffffffff1663b1bd42856040518163ffffffff1660e01b8152600401602060405180830381865afa158015610097573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100bb91906102db565b90508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508273ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250506101348261014360201b60201c565b505050610306565b5f33905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61023182610208565b9050919050565b5f61024282610227565b9050919050565b61025281610238565b811461025c575f5ffd5b50565b5f8151905061026d81610249565b92915050565b61027c81610227565b8114610286575f5ffd5b50565b5f8151905061029781610273565b92915050565b5f5f604083850312156102b3576102b2610204565b5b5f6102c08582860161025f565b92505060206102d185828601610289565b9150509250929050565b5f602082840312156102f0576102ef610204565b5b5f6102fd84828501610289565b91505092915050565b60805160a051610c596103425f395f8181605c0152818161019a0152818161029401526104e101525f818161037201526104bd0152610c595ff3fe608060405234801561000f575f5ffd5b5060043610610059575f3560e01c8063715018a6146104015780638da5cb5b1461040b578063b1bd428514610429578063e758457314610447578063f2fde38b146104655761005a565b5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637877a9ed6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100e79190610704565b15610370575f5f1b5f4914610277575f5f67ffffffffffffffff8111156101115761011061072f565b5b6040519080825280601f01601f1916602001820160405280156101435781602001600182028036833780820191505090505b5090505f5f90505b5f5f1b81491461018d578181496040516020016101699291906107d7565b6040516020818303038152906040529150808061018590610834565b91505061014b565b5f828051906020012090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083826040518263ffffffff1660e01b81526004016101f1919061088a565b602060405180830381865afa15801561020c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102309190610704565b61026f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610266906108fd565b60405180910390fd5b50505061036b565b5f5f3660405161028892919061094d565b604051809103902090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f81f2083826040518263ffffffff1660e01b81526004016102eb919061088a565b602060405180830381865afa158015610306573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032a9190610704565b610369576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610360906109af565b60405180910390fd5b505b6103ff565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f590610a17565b60405180910390fd5b5b005b610409610481565b005b610413610494565b6040516104209190610a74565b60405180910390f35b6104316104bb565b60405161043e9190610a74565b60405180910390f35b61044f6104df565b60405161045c9190610ae8565b60405180910390f35b61047f600480360381019061047a9190610b2b565b610503565b005b610489610585565b6104925f610603565b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b61050b610585565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610579576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057090610bc6565b60405180910390fd5b61058281610603565b50565b61058d6106c4565b73ffffffffffffffffffffffffffffffffffffffff166105ab610494565b73ffffffffffffffffffffffffffffffffffffffff1614610601576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f890610c2e565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f5ffd5b5f8115159050919050565b6106e3816106cf565b81146106ed575f5ffd5b50565b5f815190506106fe816106da565b92915050565b5f60208284031215610719576107186106cb565b5b5f610726848285016106f0565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6107888261075c565b6107928185610766565b93506107a2818560208601610770565b80840191505092915050565b5f819050919050565b5f819050919050565b6107d16107cc826107ae565b6107b7565b82525050565b5f6107e2828561077e565b91506107ee82846107c0565b6020820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f819050919050565b5f61083e8261082b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036108705761086f6107fe565b5b600182019050919050565b610884816107ae565b82525050565b5f60208201905061089d5f83018461087b565b92915050565b5f82825260208201905092915050565b7f496e76616c696420626c6f6220626174636800000000000000000000000000005f82015250565b5f6108e76012836108a3565b91506108f2826108b3565b602082019050919050565b5f6020820190508181035f830152610914816108db565b9050919050565b828183375f83830152505050565b5f6109348385610766565b935061094183858461091b565b82840190509392505050565b5f610959828486610929565b91508190509392505050565b7f496e76616c69642063616c6c64617461206261746368000000000000000000005f82015250565b5f6109996016836108a3565b91506109a482610965565b602082019050919050565b5f6020820190508181035f8301526109c68161098d565b9050919050565b7f4261746368496e626f783a20756e617574686f72697a656420626174636865725f82015250565b5f610a016020836108a3565b9150610a0c826109cd565b602082019050919050565b5f6020820190508181035f830152610a2e816109f5565b9050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a5e82610a35565b9050919050565b610a6e81610a54565b82525050565b5f602082019050610a875f830184610a65565b92915050565b5f819050919050565b5f610ab0610aab610aa684610a35565b610a8d565b610a35565b9050919050565b5f610ac182610a96565b9050919050565b5f610ad282610ab7565b9050919050565b610ae281610ac8565b82525050565b5f602082019050610afb5f830184610ad9565b92915050565b610b0a81610a54565b8114610b14575f5ffd5b50565b5f81359050610b2581610b01565b92915050565b5f60208284031215610b4057610b3f6106cb565b5b5f610b4d84828501610b17565b91505092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f610bb06026836108a3565b9150610bbb82610b56565b604082019050919050565b5f6020820190508181035f830152610bdd81610ba4565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f610c186020836108a3565b9150610c2382610be4565b602082019050919050565b5f6020820190508181035f830152610c4581610c0c565b905091905056fea164736f6c634300081e000a",
>>>>>>> Stashed changes
}

// BatchInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchInboxMetaData.ABI instead.
var BatchInboxABI = BatchInboxMetaData.ABI

// BatchInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchInboxMetaData.Bin instead.
var BatchInboxBin = BatchInboxMetaData.Bin

// DeployBatchInbox deploys a new Ethereum contract, binding an instance of BatchInbox to it.
func DeployBatchInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _batchAuthenticator common.Address) (common.Address, *types.Transaction, *BatchInbox, error) {
	parsed, err := BatchInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchInboxBin), backend, _batchAuthenticator)
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
