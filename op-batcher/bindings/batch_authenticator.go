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

// BatchAuthenticatorMetaData contains all meta data concerning the BatchAuthenticator contract.
var BatchAuthenticatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051611b25380380611b2583398181016040528101906100319190610358565b61004d6100426101f760201b60201c565b6101fe60201b60201c565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036100bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100b29061043c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610129576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610120906104ca565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050600160025f6101000a81548160ff0219169083151502179055506101ee816101fe60201b60201c565b505050506104e8565b5f33905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102ec826102c3565b9050919050565b5f6102fd826102e2565b9050919050565b61030d816102f3565b8114610317575f5ffd5b50565b5f8151905061032881610304565b92915050565b610337816102e2565b8114610341575f5ffd5b50565b5f815190506103528161032e565b92915050565b5f5f5f5f608085870312156103705761036f6102bf565b5b5f61037d8782880161031a565b945050602061038e87828801610344565b935050604061039f87828801610344565b92505060606103b087828801610344565b91505092959194509250565b5f82825260208201905092915050565b7f426174636841757468656e74696361746f723a207a65726f20746565206261745f8201527f6368657200000000000000000000000000000000000000000000000000000000602082015250565b5f6104266024836103bc565b9150610431826103cc565b604082019050919050565b5f6020820190508181035f8301526104538161041a565b9050919050565b7f426174636841757468656e74696361746f723a207a65726f206e6f6e2d7465655f8201527f2062617463686572000000000000000000000000000000000000000000000000602082015250565b5f6104b46028836103bc565b91506104bf8261045a565b604082019050919050565b5f6020820190508181035f8301526104e1816104a8565b9050919050565b60805160a05160c0516115fe6105275f395f81816102ad0152818161043701526105b101525f61028901525f818161037401526106bd01526115fe5ff3fe608060405234801561000f575f5ffd5b50600436106100b2575f3560e01c8063bc347f471161006f578063bc347f4714610154578063d909ba7c1461015e578063f2fde38b1461017c578063f81f208314610198578063fa14fe6d146101c8578063fc619e41146101e6576100b2565b806354fd4d50146100b6578063715018a6146100d45780637877a9ed146100de5780638da5cb5b146100fc578063b1bd42851461011a578063ba58e82a14610138575b5f5ffd5b6100be610202565b6040516100cb9190610c6f565b60405180910390f35b6100dc61023b565b005b6100e661024e565b6040516100f39190610ca9565b60405180910390f35b610104610260565b6040516101119190610d01565b60405180910390f35b610122610287565b60405161012f9190610d01565b60405180910390f35b610152600480360381019061014d9190610d83565b6102ab565b005b61015c610340565b005b610166610372565b6040516101739190610d01565b60405180910390f35b61019660048036038101906101919190610e2b565b610396565b005b6101b260048036038101906101ad9190610e89565b610418565b6040516101bf9190610ca9565b60405180910390f35b6101d0610435565b6040516101dd9190610f0f565b60405180910390f35b61020060048036038101906101fb9190610f28565b610459565b005b6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b61024361077c565b61024c5f6107fa565b565b60025f9054906101000a900460ff1681565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166335ecb4c18585858560016040518663ffffffff1660e01b815260040161030d959493929190611042565b5f604051808303815f87803b158015610324575f5ffd5b505af1158015610336573d5f5f3e3d5ffd5b5050505050505050565b61034861077c565b60025f9054906101000a900460ff161560025f6101000a81548160ff021916908315150217905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b61039e61077c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361040c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610403906110f9565b60405180910390fd5b610415816107fa565b50565b6001602052805f5260405f205f915054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f82828080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090505f816040815181106104b5576104b4611117565b5b602001015160f81c60f81b60f81c90505f8160ff1614806104d9575060018160ff16145b1561053457601b816104eb919061117d565b90508060f81b8260408151811061050557610504611117565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b5f61053f86846108bb565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a6906111fb565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d80a4c286040518163ffffffff1660e01b8152600401602060405180830381865afa158015610618573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061063c9190611254565b73ffffffffffffffffffffffffffffffffffffffff16630123d0c1826040518263ffffffff1660e01b81526004016106749190610d01565b602060405180830381865afa15801561068f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106b391906112a9565b15801561070c57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b1561074c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107439061131e565b60405180910390fd5b6001805f8881526020019081526020015f205f6101000a81548160ff021916908315150217905550505050505050565b6107846108e0565b73ffffffffffffffffffffffffffffffffffffffff166107a2610260565b73ffffffffffffffffffffffffffffffffffffffff16146107f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ef90611386565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5f5f6108c885856108e7565b915091506108d581610933565b819250505092915050565b5f33905090565b5f5f6041835103610924575f5f5f602086015192506040860151915060608601515f1a905061091887828585610afe565b9450945050505061092c565b5f6002915091505b9250929050565b5f600481111561094657610945610fcf565b5b81600481111561095957610958610fcf565b5b0315610afb576001600481111561097357610972610fcf565b5b81600481111561098657610985610fcf565b5b036109c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109bd906113ee565b60405180910390fd5b600260048111156109da576109d9610fcf565b5b8160048111156109ed576109ec610fcf565b5b03610a2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2490611456565b60405180910390fd5b60036004811115610a4157610a40610fcf565b5b816004811115610a5457610a53610fcf565b5b03610a94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8b906114e4565b60405180910390fd5b600480811115610aa757610aa6610fcf565b5b816004811115610aba57610ab9610fcf565b5b03610afa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af190611572565b60405180910390fd5b5b50565b5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0835f1c1115610b36575f600391509150610bf6565b601b8560ff1614158015610b4e5750601c8560ff1614155b15610b5f575f600491509150610bf6565b5f6001878787876040515f8152602001604052604051610b8294939291906115ae565b6020604051602081039080840390855afa158015610ba2573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bee575f60019250925050610bf6565b805f92509250505b94509492505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610c4182610bff565b610c4b8185610c09565b9350610c5b818560208601610c19565b610c6481610c27565b840191505092915050565b5f6020820190508181035f830152610c878184610c37565b905092915050565b5f8115159050919050565b610ca381610c8f565b82525050565b5f602082019050610cbc5f830184610c9a565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ceb82610cc2565b9050919050565b610cfb81610ce1565b82525050565b5f602082019050610d145f830184610cf2565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610d4357610d42610d22565b5b8235905067ffffffffffffffff811115610d6057610d5f610d26565b5b602083019150836001820283011115610d7c57610d7b610d2a565b5b9250929050565b5f5f5f5f60408587031215610d9b57610d9a610d1a565b5b5f85013567ffffffffffffffff811115610db857610db7610d1e565b5b610dc487828801610d2e565b9450945050602085013567ffffffffffffffff811115610de757610de6610d1e565b5b610df387828801610d2e565b925092505092959194509250565b610e0a81610ce1565b8114610e14575f5ffd5b50565b5f81359050610e2581610e01565b92915050565b5f60208284031215610e4057610e3f610d1a565b5b5f610e4d84828501610e17565b91505092915050565b5f819050919050565b610e6881610e56565b8114610e72575f5ffd5b50565b5f81359050610e8381610e5f565b92915050565b5f60208284031215610e9e57610e9d610d1a565b5b5f610eab84828501610e75565b91505092915050565b5f819050919050565b5f610ed7610ed2610ecd84610cc2565b610eb4565b610cc2565b9050919050565b5f610ee882610ebd565b9050919050565b5f610ef982610ede565b9050919050565b610f0981610eef565b82525050565b5f602082019050610f225f830184610f00565b92915050565b5f5f5f60408486031215610f3f57610f3e610d1a565b5b5f610f4c86828701610e75565b935050602084013567ffffffffffffffff811115610f6d57610f6c610d1e565b5b610f7986828701610d2e565b92509250509250925092565b5f82825260208201905092915050565b828183375f83830152505050565b5f610fae8385610f85565b9350610fbb838584610f95565b610fc483610c27565b840190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6002811061100d5761100c610fcf565b5b50565b5f81905061101d82610ffc565b919050565b5f61102c82611010565b9050919050565b61103c81611022565b82525050565b5f6060820190508181035f83015261105b818789610fa3565b90508181036020830152611070818587610fa3565b905061107f6040830184611033565b9695505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6110e3602683610c09565b91506110ee82611089565b604082019050919050565b5f6020820190508181035f830152611110816110d7565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60ff82169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61118782611144565b915061119283611144565b9250828201905060ff8111156111ab576111aa611150565b5b92915050565b7f496e76616c6964207369676e61747572650000000000000000000000000000005f82015250565b5f6111e5601183610c09565b91506111f0826111b1565b602082019050919050565b5f6020820190508181035f830152611212816111d9565b9050919050565b5f61122382610ce1565b9050919050565b61123381611219565b811461123d575f5ffd5b50565b5f8151905061124e8161122a565b92915050565b5f6020828403121561126957611268610d1a565b5b5f61127684828501611240565b91505092915050565b61128881610c8f565b8114611292575f5ffd5b50565b5f815190506112a38161127f565b92915050565b5f602082840312156112be576112bd610d1a565b5b5f6112cb84828501611295565b91505092915050565b7f496e76616c6964207369676e65720000000000000000000000000000000000005f82015250565b5f611308600e83610c09565b9150611313826112d4565b602082019050919050565b5f6020820190508181035f830152611335816112fc565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611370602083610c09565b915061137b8261133c565b602082019050919050565b5f6020820190508181035f83015261139d81611364565b9050919050565b7f45434453413a20696e76616c6964207369676e617475726500000000000000005f82015250565b5f6113d8601883610c09565b91506113e3826113a4565b602082019050919050565b5f6020820190508181035f830152611405816113cc565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e677468005f82015250565b5f611440601f83610c09565b915061144b8261140c565b602082019050919050565b5f6020820190508181035f83015261146d81611434565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f6114ce602283610c09565b91506114d982611474565b604082019050919050565b5f6020820190508181035f8301526114fb816114c2565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202776272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f61155c602283610c09565b915061156782611502565b604082019050919050565b5f6020820190508181035f83015261158981611550565b9050919050565b61159981610e56565b82525050565b6115a881611144565b82525050565b5f6080820190506115c15f830187611590565b6115ce602083018661159f565b6115db6040830185611590565b6115e86060830184611590565b9594505050505056fea164736f6c634300081e000a",
}

// BatchAuthenticatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchAuthenticatorMetaData.ABI instead.
var BatchAuthenticatorABI = BatchAuthenticatorMetaData.ABI

// BatchAuthenticatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchAuthenticatorMetaData.Bin instead.
var BatchAuthenticatorBin = BatchAuthenticatorMetaData.Bin

// DeployBatchAuthenticator deploys a new Ethereum contract, binding an instance of BatchAuthenticator to it.
func DeployBatchAuthenticator(auth *bind.TransactOpts, backend bind.ContractBackend, _espressoTEEVerifier common.Address, _teeBatcher common.Address, _nonTeeBatcher common.Address, _owner common.Address) (common.Address, *types.Transaction, *BatchAuthenticator, error) {
	parsed, err := BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchAuthenticatorBin), backend, _espressoTEEVerifier, _teeBatcher, _nonTeeBatcher, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchAuthenticator{BatchAuthenticatorCaller: BatchAuthenticatorCaller{contract: contract}, BatchAuthenticatorTransactor: BatchAuthenticatorTransactor{contract: contract}, BatchAuthenticatorFilterer: BatchAuthenticatorFilterer{contract: contract}}, nil
}

// BatchAuthenticator is an auto generated Go binding around an Ethereum contract.
type BatchAuthenticator struct {
	BatchAuthenticatorCaller     // Read-only binding to the contract
	BatchAuthenticatorTransactor // Write-only binding to the contract
	BatchAuthenticatorFilterer   // Log filterer for contract events
}

// BatchAuthenticatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchAuthenticatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchAuthenticatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchAuthenticatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchAuthenticatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchAuthenticatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchAuthenticatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchAuthenticatorSession struct {
	Contract     *BatchAuthenticator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BatchAuthenticatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchAuthenticatorCallerSession struct {
	Contract *BatchAuthenticatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BatchAuthenticatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchAuthenticatorTransactorSession struct {
	Contract     *BatchAuthenticatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BatchAuthenticatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchAuthenticatorRaw struct {
	Contract *BatchAuthenticator // Generic contract binding to access the raw methods on
}

// BatchAuthenticatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchAuthenticatorCallerRaw struct {
	Contract *BatchAuthenticatorCaller // Generic read-only contract binding to access the raw methods on
}

// BatchAuthenticatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchAuthenticatorTransactorRaw struct {
	Contract *BatchAuthenticatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchAuthenticator creates a new instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticator(address common.Address, backend bind.ContractBackend) (*BatchAuthenticator, error) {
	contract, err := bindBatchAuthenticator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticator{BatchAuthenticatorCaller: BatchAuthenticatorCaller{contract: contract}, BatchAuthenticatorTransactor: BatchAuthenticatorTransactor{contract: contract}, BatchAuthenticatorFilterer: BatchAuthenticatorFilterer{contract: contract}}, nil
}

// NewBatchAuthenticatorCaller creates a new read-only instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticatorCaller(address common.Address, caller bind.ContractCaller) (*BatchAuthenticatorCaller, error) {
	contract, err := bindBatchAuthenticator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorCaller{contract: contract}, nil
}

// NewBatchAuthenticatorTransactor creates a new write-only instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchAuthenticatorTransactor, error) {
	contract, err := bindBatchAuthenticator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorTransactor{contract: contract}, nil
}

// NewBatchAuthenticatorFilterer creates a new log filterer instance of BatchAuthenticator, bound to a specific deployed contract.
func NewBatchAuthenticatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchAuthenticatorFilterer, error) {
	contract, err := bindBatchAuthenticator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorFilterer{contract: contract}, nil
}

// bindBatchAuthenticator binds a generic wrapper to an already deployed contract.
func bindBatchAuthenticator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchAuthenticator *BatchAuthenticatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchAuthenticator.Contract.BatchAuthenticatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchAuthenticator *BatchAuthenticatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.BatchAuthenticatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchAuthenticator *BatchAuthenticatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.BatchAuthenticatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchAuthenticator *BatchAuthenticatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchAuthenticator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchAuthenticator *BatchAuthenticatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchAuthenticator *BatchAuthenticatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.contract.Transact(opts, method, params...)
}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ActiveIsTee(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "activeIsTee")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) ActiveIsTee() (bool, error) {
	return _BatchAuthenticator.Contract.ActiveIsTee(&_BatchAuthenticator.CallOpts)
}

// ActiveIsTee is a free data retrieval call binding the contract method 0x7877a9ed.
//
// Solidity: function activeIsTee() view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ActiveIsTee() (bool, error) {
	return _BatchAuthenticator.Contract.ActiveIsTee(&_BatchAuthenticator.CallOpts)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) EspressoTEEVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "espressoTEEVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) EspressoTEEVerifier() (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoTEEVerifier(&_BatchAuthenticator.CallOpts)
}

// EspressoTEEVerifier is a free data retrieval call binding the contract method 0xfa14fe6d.
//
// Solidity: function espressoTEEVerifier() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) EspressoTEEVerifier() (common.Address, error) {
	return _BatchAuthenticator.Contract.EspressoTEEVerifier(&_BatchAuthenticator.CallOpts)
}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) NonTeeBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "nonTeeBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) NonTeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.NonTeeBatcher(&_BatchAuthenticator.CallOpts)
}

// NonTeeBatcher is a free data retrieval call binding the contract method 0xb1bd4285.
//
// Solidity: function nonTeeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) NonTeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.NonTeeBatcher(&_BatchAuthenticator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) Owner() (common.Address, error) {
	return _BatchAuthenticator.Contract.Owner(&_BatchAuthenticator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) Owner() (common.Address, error) {
	return _BatchAuthenticator.Contract.Owner(&_BatchAuthenticator.CallOpts)
}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) TeeBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "teeBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) TeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.TeeBatcher(&_BatchAuthenticator.CallOpts)
}

// TeeBatcher is a free data retrieval call binding the contract method 0xd909ba7c.
//
// Solidity: function teeBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) TeeBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.TeeBatcher(&_BatchAuthenticator.CallOpts)
}

// ValidBatchInfo is a free data retrieval call binding the contract method 0xf81f2083.
//
// Solidity: function validBatchInfo(bytes32 ) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ValidBatchInfo(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "validBatchInfo", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidBatchInfo is a free data retrieval call binding the contract method 0xf81f2083.
//
// Solidity: function validBatchInfo(bytes32 ) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorSession) ValidBatchInfo(arg0 [32]byte) (bool, error) {
	return _BatchAuthenticator.Contract.ValidBatchInfo(&_BatchAuthenticator.CallOpts, arg0)
}

// ValidBatchInfo is a free data retrieval call binding the contract method 0xf81f2083.
//
// Solidity: function validBatchInfo(bytes32 ) view returns(bool)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ValidBatchInfo(arg0 [32]byte) (bool, error) {
	return _BatchAuthenticator.Contract.ValidBatchInfo(&_BatchAuthenticator.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchAuthenticator *BatchAuthenticatorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchAuthenticator *BatchAuthenticatorSession) Version() (string, error) {
	return _BatchAuthenticator.Contract.Version(&_BatchAuthenticator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) Version() (string, error) {
	return _BatchAuthenticator.Contract.Version(&_BatchAuthenticator.CallOpts)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) AuthenticateBatchInfo(opts *bind.TransactOpts, commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "authenticateBatchInfo", commitment, _signature)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) AuthenticateBatchInfo(commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AuthenticateBatchInfo(&_BatchAuthenticator.TransactOpts, commitment, _signature)
}

// AuthenticateBatchInfo is a paid mutator transaction binding the contract method 0xfc619e41.
//
// Solidity: function authenticateBatchInfo(bytes32 commitment, bytes _signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) AuthenticateBatchInfo(commitment [32]byte, _signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.AuthenticateBatchInfo(&_BatchAuthenticator.TransactOpts, commitment, _signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RegisterSigner(opts *bind.TransactOpts, attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "registerSigner", attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, attestationTbs, signature)
}

// RegisterSigner is a paid mutator transaction binding the contract method 0xba58e82a.
//
// Solidity: function registerSigner(bytes attestationTbs, bytes signature) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RegisterSigner(attestationTbs []byte, signature []byte) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSigner(&_BatchAuthenticator.TransactOpts, attestationTbs, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RenounceOwnership(&_BatchAuthenticator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RenounceOwnership(&_BatchAuthenticator.TransactOpts)
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SwitchBatcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "switchBatcher")
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SwitchBatcher() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SwitchBatcher(&_BatchAuthenticator.TransactOpts)
}

// SwitchBatcher is a paid mutator transaction binding the contract method 0xbc347f47.
//
// Solidity: function switchBatcher() returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SwitchBatcher() (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SwitchBatcher(&_BatchAuthenticator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.TransferOwnership(&_BatchAuthenticator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.TransferOwnership(&_BatchAuthenticator.TransactOpts, newOwner)
}

// BatchAuthenticatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BatchAuthenticator contract.
type BatchAuthenticatorOwnershipTransferredIterator struct {
	Event *BatchAuthenticatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorOwnershipTransferred)
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
		it.Event = new(BatchAuthenticatorOwnershipTransferred)
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
func (it *BatchAuthenticatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorOwnershipTransferred represents a OwnershipTransferred event raised by the BatchAuthenticator contract.
type BatchAuthenticatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchAuthenticatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorOwnershipTransferredIterator{contract: _BatchAuthenticator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorOwnershipTransferred)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseOwnershipTransferred(log types.Log) (*BatchAuthenticatorOwnershipTransferred, error) {
	event := new(BatchAuthenticatorOwnershipTransferred)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
