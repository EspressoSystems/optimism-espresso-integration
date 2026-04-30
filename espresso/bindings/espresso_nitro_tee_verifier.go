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

// EspressoNitroTEEVerifierMetaData contains all meta data concerning the EspressoNitroTEEVerifier contract.
var EspressoNitroTEEVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"teeVerifier_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nitroEnclaveVerifier_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteEnclaveHashes\",\"inputs\":[{\"name\":\"enclaveHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isSignerValid\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroEnclaveVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroEnclaveVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerService\",\"inputs\":[{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNitroEnclaveVerifier\",\"inputs\":[{\"name\":\"nitroEnclaveVerifier_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeletedEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EnclaveHashSet\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"valid\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NitroEnclaveVerifierSet\",\"inputs\":[{\"name\":\"nitroEnclaveVerifierAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ServiceRegistered\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeeVerifierSet\",\"inputs\":[{\"name\":\"teeVerifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNitroEnclaveVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTEEVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedTEEVerifier\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[{\"name\":\"result\",\"type\":\"uint8\",\"internalType\":\"enumVerificationResult\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516114d13803806114d183398101604081905261002e91610164565b8161003881610049565b50610042816100b9565b5050610195565b6001600160a01b03811661007057604051636f987f2f60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383169081179091556040517f99422856f4a571d48b77a2845d7c0e43e0b2c6d46cf7149ece07581dbfd9a302905f90a250565b6001600160a01b03811615806100d757506001600160a01b0381163b155b156100f557604051633ffce40d60e11b815260040160405180910390fd5b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f677ca5a363c501f3c7f7291bec1cd7edc4fe7a33f375571edd1d7d3067031fe69060200160405180910390a150565b80516001600160a01b038116811461015f575f80fd5b919050565b5f8060408385031215610175575f80fd5b61017e83610149565b915061018c60208401610149565b90509250929050565b61132f806101a25f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c8063966989ee11610058578063966989ee146100ec578063a628a19e1461010e578063cda752ff14610121578063ed0fb60314610160575f80fd5b80630b1c4cde14610089578063446079031461009e5780636b8c01a6146100c657806393b5552e146100d9575b5f80fd5b61009c610097366004610bd7565b61017e565b005b6100b16100ac366004610c3e565b610568565b60405190151581526020015b60405180910390f35b61009c6100d4366004610d64565b6105e4565b61009c6100e7366004610df5565b6106e7565b6100b16100fa366004610e27565b5f9081526020819052604090205460ff1690565b61009c61011c366004610c3e565b6107a3565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100bd565b60045473ffffffffffffffffffffffffffffffffffffffff1661013b565b600480546040517f6d114be30000000000000000000000000000000000000000000000000000000081525f9273ffffffffffffffffffffffffffffffffffffffff90921691636d114be3916101de91899189916002918a918a9101610ee7565b5f604051808303815f875af11580156101f9573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261023e9190810190611155565b90505f8151600381111561025457610254610e85565b146102975780516040517f470e419c00000000000000000000000000000000000000000000000000000000815261028e91906004016112a4565b60405180910390fd5b6102a081610802565b5f8160e001515f815181106102b7576102b76112b7565b6020026020010151602001515f01518260e001515f815181106102dc576102dc6112b7565b602002602001015160200151602001516040516020016103289291909182527fffffffffffffffffffffffffffffffff0000000000000000000000000000000016602082015260300190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f8181529283905291205490915060ff166103a8576040517fc2c507ca0000000000000000000000000000000000000000000000000000000081526004810182905260240161028e565b5f60018360c00151516103bb91906112e4565b67ffffffffffffffff8111156103d3576103d3610c78565b6040519080825280601f01601f1916602001820160405280156103fd576020820181803683370190505b50905060015b8360c001515181101561049f578360c001518181518110610426576104266112b7565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016826104596001846112e4565b81518110610469576104696112b7565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101610403565b50805160208083019190912073ffffffffffffffffffffffffffffffffffffffff81165f9081526001909252604090912054819060ff1661055d5773ffffffffffffffffffffffffffffffffffffffff81165f81815260016020818152604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690931790925560039052808220879055518692917f407efaf4a47650e02c37b00e26896263dfe3932b230a4977fc9b95110aea25cb91a35b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001602052604081205460ff1661059b57505f919050565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260036020526040902054806105cd57505f92915050565b5f9081526020819052604090205460ff1692915050565b60025473ffffffffffffffffffffffffffffffffffffffff163314610637576040517f9cb024bc00000000000000000000000000000000000000000000000000000000815233600482015260240161028e565b5f5b81518110156106e3575f828281518110610655576106556112b7565b6020908102919091018101515f8181529182905260409091205490915060ff1661067f57506106db565b5f8181526020819052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555182917f51145aaaa8e2b82c555348e07daa928f1c6e48c375ef759366eda54009d84ca691a2505b600101610639565b5050565b60025473ffffffffffffffffffffffffffffffffffffffff16331461073a576040517f9cb024bc00000000000000000000000000000000000000000000000000000000815233600482015260240161028e565b5f8281526020819052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168415159081179091559051909184917f2282c24f65eac8254df1107716a961b677b872ed0e1d2a9f6bafc154441eb7fd9190a35050565b60025473ffffffffffffffffffffffffffffffffffffffff1633146107f6576040517f9cb024bc00000000000000000000000000000000000000000000000000000000815233600482015260240161028e565b6107ff81610aaa565b50565b5f8160e001515111610870576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f5043522061727261792063616e6e6f7420626520656d70747900000000000000604482015260640161028e565b8060e001515f81518110610886576108866112b7565b60200260200101515f015167ffffffffffffffff165f14610929576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f466972737420504352206d75737420626520504352302028636f6465206d656160448201527f737572656d656e74290000000000000000000000000000000000000000000000606482015260840161028e565b8060c00151516041146109be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f496e76616c6964207075626c6963206b6579206c656e677468202d206d75737460448201527f2062652036352062797465730000000000000000000000000000000000000000606482015260840161028e565b8060c001515f815181106109d4576109d46112b7565b6020910101517fff00000000000000000000000000000000000000000000000000000000000000167f0400000000000000000000000000000000000000000000000000000000000000146107ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f5075626c6963206b6579206d75737420626520756e636f6d707265737365642060448201527f2830783034207072656669782900000000000000000000000000000000000000606482015260840161028e565b73ffffffffffffffffffffffffffffffffffffffff81161580610ae2575073ffffffffffffffffffffffffffffffffffffffff81163b155b15610b19576040517f7ff9c81a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f677ca5a363c501f3c7f7291bec1cd7edc4fe7a33f375571edd1d7d3067031fe69060200160405180910390a150565b5f8083601f840112610ba2575f80fd5b50813567ffffffffffffffff811115610bb9575f80fd5b602083019150836020828501011115610bd0575f80fd5b9250929050565b5f805f8060408587031215610bea575f80fd5b843567ffffffffffffffff80821115610c01575f80fd5b610c0d88838901610b92565b90965094506020870135915080821115610c25575f80fd5b50610c3287828801610b92565b95989497509550505050565b5f60208284031215610c4e575f80fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610c71575f80fd5b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040805190810167ffffffffffffffff81118282101715610cc857610cc8610c78565b60405290565b604051610120810167ffffffffffffffff81118282101715610cc857610cc8610c78565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610d3957610d39610c78565b604052919050565b5f67ffffffffffffffff821115610d5a57610d5a610c78565b5060051b60200190565b5f6020808385031215610d75575f80fd5b823567ffffffffffffffff811115610d8b575f80fd5b8301601f81018513610d9b575f80fd5b8035610dae610da982610d41565b610cf2565b81815260059190911b82018301908381019087831115610dcc575f80fd5b928401925b82841015610dea57833582529284019290840190610dd1565b979650505050505050565b5f8060408385031215610e06575f80fd5b8235915060208301358015158114610e1c575f80fd5b809150509250929050565b5f60208284031215610e37575f80fd5b5035919050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600481106107ff577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b606081525f610efa606083018789610e3e565b610f0386610eb2565b8560208401528281036040840152610f1c818587610e3e565b98975050505050505050565b805160048110610f36575f80fd5b919050565b805160ff81168114610f36575f80fd5b805167ffffffffffffffff81168114610f36575f80fd5b5f82601f830112610f71575f80fd5b81516020610f81610da983610d41565b8083825260208201915060208460051b870101935086841115610fa2575f80fd5b602086015b84811015610fbe5780518352918301918301610fa7565b509695505050505050565b5f82601f830112610fd8575f80fd5b815167ffffffffffffffff811115610ff257610ff2610c78565b61102360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610cf2565b818152846020838601011115611037575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f82601f830112611062575f80fd5b81516020611072610da983610d41565b82815260609283028501820192828201919087851115611090575f80fd5b8387015b8581101561114857808903828112156110ab575f80fd5b6110b3610ca5565b6110bc83610f4b565b81526040807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840112156110ee575f80fd5b6110f6610ca5565b848901518152908401519092507fffffffffffffffffffffffffffffffff0000000000000000000000000000000081168114611130575f80fd5b82880152808701919091528452928401928101611094565b5090979650505050505050565b5f60208284031215611165575f80fd5b815167ffffffffffffffff8082111561117c575f80fd5b908301906101208286031215611190575f80fd5b611198610cce565b6111a183610f28565b81526111af60208401610f3b565b60208201526111c060408401610f4b565b60408201526060830151828111156111d6575f80fd5b6111e287828601610f62565b6060830152506080830151828111156111f9575f80fd5b61120587828601610fc9565b60808301525060a08301518281111561121c575f80fd5b61122887828601610fc9565b60a08301525060c08301518281111561123f575f80fd5b61124b87828601610fc9565b60c08301525060e083015182811115611262575f80fd5b61126e87828601611053565b60e0830152506101008084015183811115611287575f80fd5b61129388828701610fc9565b918301919091525095945050505050565b602081016112b183610eb2565b91905290565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8181038181111561131c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9291505056fea164736f6c6343000819000a",
}

// EspressoNitroTEEVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use EspressoNitroTEEVerifierMetaData.ABI instead.
var EspressoNitroTEEVerifierABI = EspressoNitroTEEVerifierMetaData.ABI

// EspressoNitroTEEVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EspressoNitroTEEVerifierMetaData.Bin instead.
var EspressoNitroTEEVerifierBin = EspressoNitroTEEVerifierMetaData.Bin

// DeployEspressoNitroTEEVerifier deploys a new Ethereum contract, binding an instance of EspressoNitroTEEVerifier to it.
func DeployEspressoNitroTEEVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, teeVerifier_ common.Address, nitroEnclaveVerifier_ common.Address) (common.Address, *types.Transaction, *EspressoNitroTEEVerifier, error) {
	parsed, err := EspressoNitroTEEVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EspressoNitroTEEVerifierBin), backend, teeVerifier_, nitroEnclaveVerifier_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EspressoNitroTEEVerifier{EspressoNitroTEEVerifierCaller: EspressoNitroTEEVerifierCaller{contract: contract}, EspressoNitroTEEVerifierTransactor: EspressoNitroTEEVerifierTransactor{contract: contract}, EspressoNitroTEEVerifierFilterer: EspressoNitroTEEVerifierFilterer{contract: contract}}, nil
}

// EspressoNitroTEEVerifier is an auto generated Go binding around an Ethereum contract.
type EspressoNitroTEEVerifier struct {
	EspressoNitroTEEVerifierCaller     // Read-only binding to the contract
	EspressoNitroTEEVerifierTransactor // Write-only binding to the contract
	EspressoNitroTEEVerifierFilterer   // Log filterer for contract events
}

// EspressoNitroTEEVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type EspressoNitroTEEVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspressoNitroTEEVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EspressoNitroTEEVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspressoNitroTEEVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EspressoNitroTEEVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspressoNitroTEEVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EspressoNitroTEEVerifierSession struct {
	Contract     *EspressoNitroTEEVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EspressoNitroTEEVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EspressoNitroTEEVerifierCallerSession struct {
	Contract *EspressoNitroTEEVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// EspressoNitroTEEVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EspressoNitroTEEVerifierTransactorSession struct {
	Contract     *EspressoNitroTEEVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// EspressoNitroTEEVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type EspressoNitroTEEVerifierRaw struct {
	Contract *EspressoNitroTEEVerifier // Generic contract binding to access the raw methods on
}

// EspressoNitroTEEVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EspressoNitroTEEVerifierCallerRaw struct {
	Contract *EspressoNitroTEEVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// EspressoNitroTEEVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EspressoNitroTEEVerifierTransactorRaw struct {
	Contract *EspressoNitroTEEVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEspressoNitroTEEVerifier creates a new instance of EspressoNitroTEEVerifier, bound to a specific deployed contract.
func NewEspressoNitroTEEVerifier(address common.Address, backend bind.ContractBackend) (*EspressoNitroTEEVerifier, error) {
	contract, err := bindEspressoNitroTEEVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifier{EspressoNitroTEEVerifierCaller: EspressoNitroTEEVerifierCaller{contract: contract}, EspressoNitroTEEVerifierTransactor: EspressoNitroTEEVerifierTransactor{contract: contract}, EspressoNitroTEEVerifierFilterer: EspressoNitroTEEVerifierFilterer{contract: contract}}, nil
}

// NewEspressoNitroTEEVerifierCaller creates a new read-only instance of EspressoNitroTEEVerifier, bound to a specific deployed contract.
func NewEspressoNitroTEEVerifierCaller(address common.Address, caller bind.ContractCaller) (*EspressoNitroTEEVerifierCaller, error) {
	contract, err := bindEspressoNitroTEEVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierCaller{contract: contract}, nil
}

// NewEspressoNitroTEEVerifierTransactor creates a new write-only instance of EspressoNitroTEEVerifier, bound to a specific deployed contract.
func NewEspressoNitroTEEVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*EspressoNitroTEEVerifierTransactor, error) {
	contract, err := bindEspressoNitroTEEVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierTransactor{contract: contract}, nil
}

// NewEspressoNitroTEEVerifierFilterer creates a new log filterer instance of EspressoNitroTEEVerifier, bound to a specific deployed contract.
func NewEspressoNitroTEEVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*EspressoNitroTEEVerifierFilterer, error) {
	contract, err := bindEspressoNitroTEEVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierFilterer{contract: contract}, nil
}

// bindEspressoNitroTEEVerifier binds a generic wrapper to an already deployed contract.
func bindEspressoNitroTEEVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EspressoNitroTEEVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EspressoNitroTEEVerifier.Contract.EspressoNitroTEEVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.EspressoNitroTEEVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.EspressoNitroTEEVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EspressoNitroTEEVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.contract.Transact(opts, method, params...)
}

// IsSignerValid is a free data retrieval call binding the contract method 0x44607903.
//
// Solidity: function isSignerValid(address signer) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) IsSignerValid(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "isSignerValid", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignerValid is a free data retrieval call binding the contract method 0x44607903.
//
// Solidity: function isSignerValid(address signer) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) IsSignerValid(signer common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.IsSignerValid(&_EspressoNitroTEEVerifier.CallOpts, signer)
}

// IsSignerValid is a free data retrieval call binding the contract method 0x44607903.
//
// Solidity: function isSignerValid(address signer) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) IsSignerValid(signer common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.IsSignerValid(&_EspressoNitroTEEVerifier.CallOpts, signer)
}

// NitroEnclaveVerifier is a free data retrieval call binding the contract method 0xed0fb603.
//
// Solidity: function nitroEnclaveVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) NitroEnclaveVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "nitroEnclaveVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NitroEnclaveVerifier is a free data retrieval call binding the contract method 0xed0fb603.
//
// Solidity: function nitroEnclaveVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) NitroEnclaveVerifier() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.NitroEnclaveVerifier(&_EspressoNitroTEEVerifier.CallOpts)
}

// NitroEnclaveVerifier is a free data retrieval call binding the contract method 0xed0fb603.
//
// Solidity: function nitroEnclaveVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) NitroEnclaveVerifier() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.NitroEnclaveVerifier(&_EspressoNitroTEEVerifier.CallOpts)
}

// RegisteredEnclaveHash is a free data retrieval call binding the contract method 0x966989ee.
//
// Solidity: function registeredEnclaveHash(bytes32 enclaveHash) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) RegisteredEnclaveHash(opts *bind.CallOpts, enclaveHash [32]byte) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "registeredEnclaveHash", enclaveHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredEnclaveHash is a free data retrieval call binding the contract method 0x966989ee.
//
// Solidity: function registeredEnclaveHash(bytes32 enclaveHash) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RegisteredEnclaveHash(enclaveHash [32]byte) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredEnclaveHash(&_EspressoNitroTEEVerifier.CallOpts, enclaveHash)
}

// RegisteredEnclaveHash is a free data retrieval call binding the contract method 0x966989ee.
//
// Solidity: function registeredEnclaveHash(bytes32 enclaveHash) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) RegisteredEnclaveHash(enclaveHash [32]byte) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredEnclaveHash(&_EspressoNitroTEEVerifier.CallOpts, enclaveHash)
}

// TeeVerifier is a free data retrieval call binding the contract method 0xcda752ff.
//
// Solidity: function teeVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) TeeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "teeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeVerifier is a free data retrieval call binding the contract method 0xcda752ff.
//
// Solidity: function teeVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) TeeVerifier() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.TeeVerifier(&_EspressoNitroTEEVerifier.CallOpts)
}

// TeeVerifier is a free data retrieval call binding the contract method 0xcda752ff.
//
// Solidity: function teeVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) TeeVerifier() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.TeeVerifier(&_EspressoNitroTEEVerifier.CallOpts)
}

// DeleteEnclaveHashes is a paid mutator transaction binding the contract method 0x6b8c01a6.
//
// Solidity: function deleteEnclaveHashes(bytes32[] enclaveHashes) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) DeleteEnclaveHashes(opts *bind.TransactOpts, enclaveHashes [][32]byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "deleteEnclaveHashes", enclaveHashes)
}

// DeleteEnclaveHashes is a paid mutator transaction binding the contract method 0x6b8c01a6.
//
// Solidity: function deleteEnclaveHashes(bytes32[] enclaveHashes) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) DeleteEnclaveHashes(enclaveHashes [][32]byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.DeleteEnclaveHashes(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHashes)
}

// DeleteEnclaveHashes is a paid mutator transaction binding the contract method 0x6b8c01a6.
//
// Solidity: function deleteEnclaveHashes(bytes32[] enclaveHashes) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) DeleteEnclaveHashes(enclaveHashes [][32]byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.DeleteEnclaveHashes(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHashes)
}

// RegisterService is a paid mutator transaction binding the contract method 0x0b1c4cde.
//
// Solidity: function registerService(bytes output, bytes proofBytes) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) RegisterService(opts *bind.TransactOpts, output []byte, proofBytes []byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "registerService", output, proofBytes)
}

// RegisterService is a paid mutator transaction binding the contract method 0x0b1c4cde.
//
// Solidity: function registerService(bytes output, bytes proofBytes) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RegisterService(output []byte, proofBytes []byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisterService(&_EspressoNitroTEEVerifier.TransactOpts, output, proofBytes)
}

// RegisterService is a paid mutator transaction binding the contract method 0x0b1c4cde.
//
// Solidity: function registerService(bytes output, bytes proofBytes) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) RegisterService(output []byte, proofBytes []byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisterService(&_EspressoNitroTEEVerifier.TransactOpts, output, proofBytes)
}

// SetEnclaveHash is a paid mutator transaction binding the contract method 0x93b5552e.
//
// Solidity: function setEnclaveHash(bytes32 enclaveHash, bool valid) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) SetEnclaveHash(opts *bind.TransactOpts, enclaveHash [32]byte, valid bool) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "setEnclaveHash", enclaveHash, valid)
}

// SetEnclaveHash is a paid mutator transaction binding the contract method 0x93b5552e.
//
// Solidity: function setEnclaveHash(bytes32 enclaveHash, bool valid) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) SetEnclaveHash(enclaveHash [32]byte, valid bool) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetEnclaveHash(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHash, valid)
}

// SetEnclaveHash is a paid mutator transaction binding the contract method 0x93b5552e.
//
// Solidity: function setEnclaveHash(bytes32 enclaveHash, bool valid) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) SetEnclaveHash(enclaveHash [32]byte, valid bool) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetEnclaveHash(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHash, valid)
}

// SetNitroEnclaveVerifier is a paid mutator transaction binding the contract method 0xa628a19e.
//
// Solidity: function setNitroEnclaveVerifier(address nitroEnclaveVerifier_) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) SetNitroEnclaveVerifier(opts *bind.TransactOpts, nitroEnclaveVerifier_ common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "setNitroEnclaveVerifier", nitroEnclaveVerifier_)
}

// SetNitroEnclaveVerifier is a paid mutator transaction binding the contract method 0xa628a19e.
//
// Solidity: function setNitroEnclaveVerifier(address nitroEnclaveVerifier_) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) SetNitroEnclaveVerifier(nitroEnclaveVerifier_ common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetNitroEnclaveVerifier(&_EspressoNitroTEEVerifier.TransactOpts, nitroEnclaveVerifier_)
}

// SetNitroEnclaveVerifier is a paid mutator transaction binding the contract method 0xa628a19e.
//
// Solidity: function setNitroEnclaveVerifier(address nitroEnclaveVerifier_) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) SetNitroEnclaveVerifier(nitroEnclaveVerifier_ common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetNitroEnclaveVerifier(&_EspressoNitroTEEVerifier.TransactOpts, nitroEnclaveVerifier_)
}

// EspressoNitroTEEVerifierDeletedEnclaveHashIterator is returned from FilterDeletedEnclaveHash and is used to iterate over the raw logs and unpacked data for DeletedEnclaveHash events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierDeletedEnclaveHashIterator struct {
	Event *EspressoNitroTEEVerifierDeletedEnclaveHash // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierDeletedEnclaveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierDeletedEnclaveHash)
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
		it.Event = new(EspressoNitroTEEVerifierDeletedEnclaveHash)
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
func (it *EspressoNitroTEEVerifierDeletedEnclaveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierDeletedEnclaveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierDeletedEnclaveHash represents a DeletedEnclaveHash event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierDeletedEnclaveHash struct {
	EnclaveHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeletedEnclaveHash is a free log retrieval operation binding the contract event 0x51145aaaa8e2b82c555348e07daa928f1c6e48c375ef759366eda54009d84ca6.
//
// Solidity: event DeletedEnclaveHash(bytes32 indexed enclaveHash)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterDeletedEnclaveHash(opts *bind.FilterOpts, enclaveHash [][32]byte) (*EspressoNitroTEEVerifierDeletedEnclaveHashIterator, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "DeletedEnclaveHash", enclaveHashRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierDeletedEnclaveHashIterator{contract: _EspressoNitroTEEVerifier.contract, event: "DeletedEnclaveHash", logs: logs, sub: sub}, nil
}

// WatchDeletedEnclaveHash is a free log subscription operation binding the contract event 0x51145aaaa8e2b82c555348e07daa928f1c6e48c375ef759366eda54009d84ca6.
//
// Solidity: event DeletedEnclaveHash(bytes32 indexed enclaveHash)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchDeletedEnclaveHash(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierDeletedEnclaveHash, enclaveHash [][32]byte) (event.Subscription, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "DeletedEnclaveHash", enclaveHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierDeletedEnclaveHash)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "DeletedEnclaveHash", log); err != nil {
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

// ParseDeletedEnclaveHash is a log parse operation binding the contract event 0x51145aaaa8e2b82c555348e07daa928f1c6e48c375ef759366eda54009d84ca6.
//
// Solidity: event DeletedEnclaveHash(bytes32 indexed enclaveHash)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseDeletedEnclaveHash(log types.Log) (*EspressoNitroTEEVerifierDeletedEnclaveHash, error) {
	event := new(EspressoNitroTEEVerifierDeletedEnclaveHash)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "DeletedEnclaveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierEnclaveHashSetIterator is returned from FilterEnclaveHashSet and is used to iterate over the raw logs and unpacked data for EnclaveHashSet events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierEnclaveHashSetIterator struct {
	Event *EspressoNitroTEEVerifierEnclaveHashSet // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierEnclaveHashSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierEnclaveHashSet)
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
		it.Event = new(EspressoNitroTEEVerifierEnclaveHashSet)
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
func (it *EspressoNitroTEEVerifierEnclaveHashSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierEnclaveHashSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierEnclaveHashSet represents a EnclaveHashSet event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierEnclaveHashSet struct {
	EnclaveHash [32]byte
	Valid       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEnclaveHashSet is a free log retrieval operation binding the contract event 0x2282c24f65eac8254df1107716a961b677b872ed0e1d2a9f6bafc154441eb7fd.
//
// Solidity: event EnclaveHashSet(bytes32 indexed enclaveHash, bool indexed valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterEnclaveHashSet(opts *bind.FilterOpts, enclaveHash [][32]byte, valid []bool) (*EspressoNitroTEEVerifierEnclaveHashSetIterator, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var validRule []interface{}
	for _, validItem := range valid {
		validRule = append(validRule, validItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "EnclaveHashSet", enclaveHashRule, validRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierEnclaveHashSetIterator{contract: _EspressoNitroTEEVerifier.contract, event: "EnclaveHashSet", logs: logs, sub: sub}, nil
}

// WatchEnclaveHashSet is a free log subscription operation binding the contract event 0x2282c24f65eac8254df1107716a961b677b872ed0e1d2a9f6bafc154441eb7fd.
//
// Solidity: event EnclaveHashSet(bytes32 indexed enclaveHash, bool indexed valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchEnclaveHashSet(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierEnclaveHashSet, enclaveHash [][32]byte, valid []bool) (event.Subscription, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var validRule []interface{}
	for _, validItem := range valid {
		validRule = append(validRule, validItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "EnclaveHashSet", enclaveHashRule, validRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierEnclaveHashSet)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "EnclaveHashSet", log); err != nil {
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

// ParseEnclaveHashSet is a log parse operation binding the contract event 0x2282c24f65eac8254df1107716a961b677b872ed0e1d2a9f6bafc154441eb7fd.
//
// Solidity: event EnclaveHashSet(bytes32 indexed enclaveHash, bool indexed valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseEnclaveHashSet(log types.Log) (*EspressoNitroTEEVerifierEnclaveHashSet, error) {
	event := new(EspressoNitroTEEVerifierEnclaveHashSet)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "EnclaveHashSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierNitroEnclaveVerifierSetIterator is returned from FilterNitroEnclaveVerifierSet and is used to iterate over the raw logs and unpacked data for NitroEnclaveVerifierSet events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierNitroEnclaveVerifierSetIterator struct {
	Event *EspressoNitroTEEVerifierNitroEnclaveVerifierSet // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierNitroEnclaveVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierNitroEnclaveVerifierSet)
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
		it.Event = new(EspressoNitroTEEVerifierNitroEnclaveVerifierSet)
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
func (it *EspressoNitroTEEVerifierNitroEnclaveVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierNitroEnclaveVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierNitroEnclaveVerifierSet represents a NitroEnclaveVerifierSet event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierNitroEnclaveVerifierSet struct {
	NitroEnclaveVerifierAddress common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNitroEnclaveVerifierSet is a free log retrieval operation binding the contract event 0x677ca5a363c501f3c7f7291bec1cd7edc4fe7a33f375571edd1d7d3067031fe6.
//
// Solidity: event NitroEnclaveVerifierSet(address nitroEnclaveVerifierAddress)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterNitroEnclaveVerifierSet(opts *bind.FilterOpts) (*EspressoNitroTEEVerifierNitroEnclaveVerifierSetIterator, error) {

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "NitroEnclaveVerifierSet")
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierNitroEnclaveVerifierSetIterator{contract: _EspressoNitroTEEVerifier.contract, event: "NitroEnclaveVerifierSet", logs: logs, sub: sub}, nil
}

// WatchNitroEnclaveVerifierSet is a free log subscription operation binding the contract event 0x677ca5a363c501f3c7f7291bec1cd7edc4fe7a33f375571edd1d7d3067031fe6.
//
// Solidity: event NitroEnclaveVerifierSet(address nitroEnclaveVerifierAddress)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchNitroEnclaveVerifierSet(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierNitroEnclaveVerifierSet) (event.Subscription, error) {

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "NitroEnclaveVerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierNitroEnclaveVerifierSet)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "NitroEnclaveVerifierSet", log); err != nil {
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

// ParseNitroEnclaveVerifierSet is a log parse operation binding the contract event 0x677ca5a363c501f3c7f7291bec1cd7edc4fe7a33f375571edd1d7d3067031fe6.
//
// Solidity: event NitroEnclaveVerifierSet(address nitroEnclaveVerifierAddress)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseNitroEnclaveVerifierSet(log types.Log) (*EspressoNitroTEEVerifierNitroEnclaveVerifierSet, error) {
	event := new(EspressoNitroTEEVerifierNitroEnclaveVerifierSet)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "NitroEnclaveVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierServiceRegisteredIterator is returned from FilterServiceRegistered and is used to iterate over the raw logs and unpacked data for ServiceRegistered events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierServiceRegisteredIterator struct {
	Event *EspressoNitroTEEVerifierServiceRegistered // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierServiceRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierServiceRegistered)
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
		it.Event = new(EspressoNitroTEEVerifierServiceRegistered)
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
func (it *EspressoNitroTEEVerifierServiceRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierServiceRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierServiceRegistered represents a ServiceRegistered event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierServiceRegistered struct {
	Signer      common.Address
	EnclaveHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceRegistered is a free log retrieval operation binding the contract event 0x407efaf4a47650e02c37b00e26896263dfe3932b230a4977fc9b95110aea25cb.
//
// Solidity: event ServiceRegistered(address indexed signer, bytes32 indexed enclaveHash)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterServiceRegistered(opts *bind.FilterOpts, signer []common.Address, enclaveHash [][32]byte) (*EspressoNitroTEEVerifierServiceRegisteredIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "ServiceRegistered", signerRule, enclaveHashRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierServiceRegisteredIterator{contract: _EspressoNitroTEEVerifier.contract, event: "ServiceRegistered", logs: logs, sub: sub}, nil
}

// WatchServiceRegistered is a free log subscription operation binding the contract event 0x407efaf4a47650e02c37b00e26896263dfe3932b230a4977fc9b95110aea25cb.
//
// Solidity: event ServiceRegistered(address indexed signer, bytes32 indexed enclaveHash)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchServiceRegistered(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierServiceRegistered, signer []common.Address, enclaveHash [][32]byte) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "ServiceRegistered", signerRule, enclaveHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierServiceRegistered)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "ServiceRegistered", log); err != nil {
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

// ParseServiceRegistered is a log parse operation binding the contract event 0x407efaf4a47650e02c37b00e26896263dfe3932b230a4977fc9b95110aea25cb.
//
// Solidity: event ServiceRegistered(address indexed signer, bytes32 indexed enclaveHash)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseServiceRegistered(log types.Log) (*EspressoNitroTEEVerifierServiceRegistered, error) {
	event := new(EspressoNitroTEEVerifierServiceRegistered)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "ServiceRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierTeeVerifierSetIterator is returned from FilterTeeVerifierSet and is used to iterate over the raw logs and unpacked data for TeeVerifierSet events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierTeeVerifierSetIterator struct {
	Event *EspressoNitroTEEVerifierTeeVerifierSet // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierTeeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierTeeVerifierSet)
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
		it.Event = new(EspressoNitroTEEVerifierTeeVerifierSet)
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
func (it *EspressoNitroTEEVerifierTeeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierTeeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierTeeVerifierSet represents a TeeVerifierSet event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierTeeVerifierSet struct {
	TeeVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTeeVerifierSet is a free log retrieval operation binding the contract event 0x99422856f4a571d48b77a2845d7c0e43e0b2c6d46cf7149ece07581dbfd9a302.
//
// Solidity: event TeeVerifierSet(address indexed teeVerifier)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterTeeVerifierSet(opts *bind.FilterOpts, teeVerifier []common.Address) (*EspressoNitroTEEVerifierTeeVerifierSetIterator, error) {

	var teeVerifierRule []interface{}
	for _, teeVerifierItem := range teeVerifier {
		teeVerifierRule = append(teeVerifierRule, teeVerifierItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "TeeVerifierSet", teeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierTeeVerifierSetIterator{contract: _EspressoNitroTEEVerifier.contract, event: "TeeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchTeeVerifierSet is a free log subscription operation binding the contract event 0x99422856f4a571d48b77a2845d7c0e43e0b2c6d46cf7149ece07581dbfd9a302.
//
// Solidity: event TeeVerifierSet(address indexed teeVerifier)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchTeeVerifierSet(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierTeeVerifierSet, teeVerifier []common.Address) (event.Subscription, error) {

	var teeVerifierRule []interface{}
	for _, teeVerifierItem := range teeVerifier {
		teeVerifierRule = append(teeVerifierRule, teeVerifierItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "TeeVerifierSet", teeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierTeeVerifierSet)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "TeeVerifierSet", log); err != nil {
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

// ParseTeeVerifierSet is a log parse operation binding the contract event 0x99422856f4a571d48b77a2845d7c0e43e0b2c6d46cf7149ece07581dbfd9a302.
//
// Solidity: event TeeVerifierSet(address indexed teeVerifier)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseTeeVerifierSet(log types.Log) (*EspressoNitroTEEVerifierTeeVerifierSet, error) {
	event := new(EspressoNitroTEEVerifierTeeVerifierSet)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "TeeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
