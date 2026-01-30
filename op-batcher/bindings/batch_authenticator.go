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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProxyAdmin\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdminOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNonTeeBatcher\",\"inputs\":[{\"name\":\"_newNonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTeeBatcher\",\"inputs\":[{\"name\":\"_newTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateBatch\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateNonTeeBatch\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateTeeBatch\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BatchInfoAuthenticated\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonTeeBatcherUpdated\",\"inputs\":[{\"name\":\"oldNonTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newNonTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerRegistrationInitiated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeeBatcherUpdated\",\"inputs\":[{\"name\":\"oldTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTeeBatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"contract_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOrProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotResolvedDelegateProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_NotSharedProxyAdminOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyAdminOwnedBase_ProxyAdminNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReinitializableBase_ZeroInitVersion\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50600160805261001d610022565b6100df565b5f54610100900460ff161561008d5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156100dd575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805161209b6100fe5f395f81816101b50152610c15015261209b5ff3fe608060405234801561000f575f5ffd5b5060043610610163575f3560e01c8063995bd81e116100c7578063d909ba7c1161007d578063f81f208311610063578063f81f20831461032f578063fa14fe6d14610351578063fc619e4114610371575f5ffd5b8063d909ba7c14610307578063dad544e014610327575f5ffd5b8063ba58e82a116100ad578063ba58e82a146102d9578063bc347f47146102ec578063c0c53b8b146102f4575f5ffd5b8063995bd81e146102a6578063b1bd4285146102b9575f5ffd5b80636f7eda471161011c5780638c8fc531116101025780638c8fc531146102785780638da5cb5b1461028b57806391a1a35d14610293575f5ffd5b80636f7eda47146102305780637877a9ed14610243575f5ffd5b806338d38c971161014c57806338d38c97146101ae5780633e47158c146101df57806354fd4d50146101e7575f5ffd5b80631b076a4c14610167578063361d2d7b14610199575b5f5ffd5b61016f610384565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101ac6101a7366004611a89565b610418565b005b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610190565b61016f6104fa565b6102236040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101909190611aa4565b6101ac61023e366004611a89565b610700565b6004546102689074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610190565b6101ac610286366004611a89565b6107e3565b61016f6108c6565b6101ac6102a1366004611b35565b6108cf565b6101ac6102b4366004611b35565b61090b565b60035461016f9073ffffffffffffffffffffffffffffffffffffffff1681565b6101ac6102e7366004611b86565b610b01565b6101ac610bbe565b6101ac610302366004611bf2565b610c13565b60025461016f9073ffffffffffffffffffffffffffffffffffffffff1681565b61016f610f3f565b61026861033d366004611c3a565b60016020525f908152604090205460ff1681565b60045461016f9073ffffffffffffffffffffffffffffffffffffffff1681565b6101ac61037f366004611c51565b610f90565b5f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d80a4c286040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103ef573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104139190611c80565b905090565b60035473ffffffffffffffffffffffffffffffffffffffff8281169116146104f75760035461045e9073ffffffffffffffffffffffffffffffffffffffff166014611382565b61047f8273ffffffffffffffffffffffffffffffffffffffff166014611382565b604051602001610490929190611cb2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526104ee91600401611aa4565b60405180910390fd5b50565b5f806105247fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b905073ffffffffffffffffffffffffffffffffffffffff81161561054757919050565b6040518060400160405280601a81526020017f4f564d5f4c3143726f7373446f6d61696e4d657373656e67657200000000000081525051600261058a9190611d95565b604080513060208201525f918101919091527f4f564d5f4c3143726f7373446f6d61696e4d657373656e67657200000000000091909117906105e4906060015b604051602081830303815290604052805190602001205490565b1461061b576040517f54e433cd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080513060208201526001918101919091525f9061063c906060016105ca565b905073ffffffffffffffffffffffffffffffffffffffff8116156106ce578073ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106a3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106c79190611c80565b9250505090565b6040517f332144db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107086115c8565b73ffffffffffffffffffffffffffffffffffffffff811661076d576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016104ee565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b3327905f90a35050565b6107eb6115c8565b73ffffffffffffffffffffffffffffffffffffffff8116610850576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016104ee565b6003805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5905f90a35050565b5f610413610f3f565b60045474010000000000000000000000000000000000000000900460ff1615610902576108fd83838361090b565b505050565b6108fd83610418565b60025473ffffffffffffffffffffffffffffffffffffffff848116911614610983576002546109519073ffffffffffffffffffffffffffffffffffffffff166014611382565b6109728473ffffffffffffffffffffffffffffffffffffffff166014611382565b604051602001610490929190611dac565b5f4915610a6557604080515f80825260208201909252905b8049156109da578181496040516020016109b6929190611e56565b604051602081830303815290604052915080806109d290611e6e565b91505061099b565b81516020808401919091205f818152600190925260409091205460ff16610a5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e76616c696420626c6f62206261746368000000000000000000000000000060448201526064016104ee565b505050505050565b5f8282604051610a76929190611ea5565b60408051918290039091205f8181526001602052919091205490915060ff16610afb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c69642063616c6c646174612062617463680000000000000000000060448201526064016104ee565b50505050565b600480546040517f7f82ea6c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911691637f82ea6c91610b619188918891889188916001915f9101611f5d565b5f604051808303815f87803b158015610b78575f5ffd5b505af1158015610b8a573d5f5f3e3d5ffd5b50506040513392507f665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c91505f90a250505050565b610bc66115c8565b600480547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff8116740100000000000000000000000000000000000000009182900460ff1615909102179055565b7f00000000000000000000000000000000000000000000000000000000000000005f54610100900460ff16158015610c5157505f5460ff8083169116105b610cdd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104ee565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff831617610100179055610d15611620565b73ffffffffffffffffffffffffffffffffffffffff8316610d7a576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016104ee565b73ffffffffffffffffffffffffffffffffffffffff8216610ddf576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016104ee565b73ffffffffffffffffffffffffffffffffffffffff8416610e44576040517f8e4c8aa600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016104ee565b600480546002805473ffffffffffffffffffffffffffffffffffffffff8781167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556003805487841692169190911790557fffffffffffffffffffffff00000000000000000000000000000000000000000090911690861617740100000000000000000000000000000000000000001790555f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b5f610f486104fa565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103ef573d5f5f3e3d5ffd5b5f82828080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505082519293505060419091149050611037576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f496e76616c6964207369676e6174757265206c656e677468000000000000000060448201526064016104ee565b5f8160408151811061104b5761104b611faf565b016020015160f81c905080158061106557508060ff166001145b156110bd57611075601b82611fdc565b90508060f81b8260408151811061108e5761108e611faf565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b5f6110c886846116a1565b905073ffffffffffffffffffffffffffffffffffffffff811661116d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f426174636841757468656e74696361746f723a20696e76616c6964207369676e60448201527f617475726500000000000000000000000000000000000000000000000000000060648201526084016104ee565b60048054604080517fd80a4c28000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9092169263d80a4c289282820192602092908290030181865afa1580156111d8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111fc9190611c80565b73ffffffffffffffffffffffffffffffffffffffff16636d8f5aa9825f6040518363ffffffff1660e01b8152600401611236929190611ff5565b602060405180830381865afa158015611251573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112759190612028565b611301576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f426174636841757468656e74696361746f723a20696e76616c6964207369676e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084016104ee565b5f86815260016020819052604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169092179091555173ffffffffffffffffffffffffffffffffffffffff83169188917f731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c9190a3505050505050565b60605f611390836002611d95565b61139b906002612047565b67ffffffffffffffff8111156113b3576113b3611e29565b6040519080825280601f01601f1916602001820160405280156113dd576020820181803683370190505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f8151811061141357611413611faf565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061147557611475611faf565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f6114af846002611d95565b6114ba906001612047565b90505b6001811115611556577f303132333435363738396162636465660000000000000000000000000000000085600f16601081106114fb576114fb611faf565b1a60f81b82828151811061151157611511611faf565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535060049490941c9361154f8161205a565b90506114bd565b5083156115bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016104ee565b90505b92915050565b336115d1610f3f565b73ffffffffffffffffffffffffffffffffffffffff161461161e576040517f7f12c64b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b336116296104fa565b73ffffffffffffffffffffffffffffffffffffffff161415801561166a575033611651610f3f565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561161e576040517fc4050a2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f6116ae85856116c3565b915091506116bb81611705565b509392505050565b5f5f82516041036116f7576020830151604084015160608501515f1a6116eb87828585611958565b945094505050506116fe565b505f905060025b9250929050565b5f81600481111561171857611718611efb565b036117205750565b600181600481111561173457611734611efb565b0361179b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104ee565b60028160048111156117af576117af611efb565b03611816576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104ee565b600381600481111561182a5761182a611efb565b036118b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104ee565b60048160048111156118cb576118cb611efb565b036104f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104ee565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561198d57505f90506003611a5f565b8460ff16601b141580156119a557508460ff16601c14155b156119b557505f90506004611a5f565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611a06573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611a59575f60019250925050611a5f565b91505f90505b94509492505050565b73ffffffffffffffffffffffffffffffffffffffff811681146104f7575f5ffd5b5f60208284031215611a99575f5ffd5b81356115bf81611a68565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f83601f840112611b07575f5ffd5b50813567ffffffffffffffff811115611b1e575f5ffd5b6020830191508360208285010111156116fe575f5ffd5b5f5f5f60408486031215611b47575f5ffd5b8335611b5281611a68565b9250602084013567ffffffffffffffff811115611b6d575f5ffd5b611b7986828701611af7565b9497909650939450505050565b5f5f5f5f60408587031215611b99575f5ffd5b843567ffffffffffffffff811115611baf575f5ffd5b611bbb87828801611af7565b909550935050602085013567ffffffffffffffff811115611bda575f5ffd5b611be687828801611af7565b95989497509550505050565b5f5f5f60608486031215611c04575f5ffd5b8335611c0f81611a68565b92506020840135611c1f81611a68565b91506040840135611c2f81611a68565b809150509250925092565b5f60208284031215611c4a575f5ffd5b5035919050565b5f5f5f60408486031215611c63575f5ffd5b83359250602084013567ffffffffffffffff811115611b6d575f5ffd5b5f60208284031215611c90575f5ffd5b81516115bf81611a68565b5f81518060208401855e5f93019283525090919050565b7f4261746368496e626f783a2062617463686572206e6f7420617574686f72697a81527f656420746f20706f737420696e2066616c6c6261636b206d6f64652e2045787060208201527f65637465643a200000000000000000000000000000000000000000000000000060408201525f611d2f6047830185611c9b565b7f2c2041637475616c3a20000000000000000000000000000000000000000000008152611d5f600a820185611c9b565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820281158282048414176115c2576115c2611d68565b7f4261746368496e626f783a2062617463686572206e6f7420617574686f72697a81527f656420746f20706f737420696e20544545206d6f64652e20457870656374656460208201527f3a2000000000000000000000000000000000000000000000000000000000000060408201525f611d2f6042830185611c9b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f611e618285611c9b565b9283525050602001919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e9e57611e9e611d68565b5060010190565b818382375f9101908152919050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600281106104f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b608081525f611f7060808301888a611eb4565b8281036020840152611f83818789611eb4565b915050611f8f84611f28565b836040830152611f9e83611f28565b826060830152979650505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b60ff81811683821601908111156115c2576115c2611d68565b73ffffffffffffffffffffffffffffffffffffffff831681526040810161201b83611f28565b8260208301529392505050565b5f60208284031215612038575f5ffd5b815180151581146115bf575f5ffd5b808201808211156115c2576115c2611d68565b5f8161206857612068611d68565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea164736f6c634300081c000a",
}

// BatchAuthenticatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchAuthenticatorMetaData.ABI instead.
var BatchAuthenticatorABI = BatchAuthenticatorMetaData.ABI

// BatchAuthenticatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchAuthenticatorMetaData.Bin instead.
var BatchAuthenticatorBin = BatchAuthenticatorMetaData.Bin

// DeployBatchAuthenticator deploys a new Ethereum contract, binding an instance of BatchAuthenticator to it.
func DeployBatchAuthenticator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchAuthenticator, error) {
	parsed, err := BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchAuthenticatorBin), backend)
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

// InitVersion is a free data retrieval call binding the contract method 0x38d38c97.
//
// Solidity: function initVersion() view returns(uint8)
func (_BatchAuthenticator *BatchAuthenticatorCaller) InitVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "initVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InitVersion is a free data retrieval call binding the contract method 0x38d38c97.
//
// Solidity: function initVersion() view returns(uint8)
func (_BatchAuthenticator *BatchAuthenticatorSession) InitVersion() (uint8, error) {
	return _BatchAuthenticator.Contract.InitVersion(&_BatchAuthenticator.CallOpts)
}

// InitVersion is a free data retrieval call binding the contract method 0x38d38c97.
//
// Solidity: function initVersion() view returns(uint8)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) InitVersion() (uint8, error) {
	return _BatchAuthenticator.Contract.InitVersion(&_BatchAuthenticator.CallOpts)
}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) NitroValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "nitroValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) NitroValidator() (common.Address, error) {
	return _BatchAuthenticator.Contract.NitroValidator(&_BatchAuthenticator.CallOpts)
}

// NitroValidator is a free data retrieval call binding the contract method 0x1b076a4c.
//
// Solidity: function nitroValidator() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) NitroValidator() (common.Address, error) {
	return _BatchAuthenticator.Contract.NitroValidator(&_BatchAuthenticator.CallOpts)
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

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) ProxyAdmin() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdmin(&_BatchAuthenticator.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ProxyAdmin() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdmin(&_BatchAuthenticator.CallOpts)
}

// ProxyAdminOwner is a free data retrieval call binding the contract method 0xdad544e0.
//
// Solidity: function proxyAdminOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) ProxyAdminOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "proxyAdminOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdminOwner is a free data retrieval call binding the contract method 0xdad544e0.
//
// Solidity: function proxyAdminOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) ProxyAdminOwner() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdminOwner(&_BatchAuthenticator.CallOpts)
}

// ProxyAdminOwner is a free data retrieval call binding the contract method 0xdad544e0.
//
// Solidity: function proxyAdminOwner() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ProxyAdminOwner() (common.Address, error) {
	return _BatchAuthenticator.Contract.ProxyAdminOwner(&_BatchAuthenticator.CallOpts)
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

// ValidateBatch is a free data retrieval call binding the contract method 0x91a1a35d.
//
// Solidity: function validateBatch(address sender, bytes data) view returns()
func (_BatchAuthenticator *BatchAuthenticatorCaller) ValidateBatch(opts *bind.CallOpts, sender common.Address, data []byte) error {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "validateBatch", sender, data)

	if err != nil {
		return err
	}

	return err

}

// ValidateBatch is a free data retrieval call binding the contract method 0x91a1a35d.
//
// Solidity: function validateBatch(address sender, bytes data) view returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) ValidateBatch(sender common.Address, data []byte) error {
	return _BatchAuthenticator.Contract.ValidateBatch(&_BatchAuthenticator.CallOpts, sender, data)
}

// ValidateBatch is a free data retrieval call binding the contract method 0x91a1a35d.
//
// Solidity: function validateBatch(address sender, bytes data) view returns()
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ValidateBatch(sender common.Address, data []byte) error {
	return _BatchAuthenticator.Contract.ValidateBatch(&_BatchAuthenticator.CallOpts, sender, data)
}

// ValidateNonTeeBatch is a free data retrieval call binding the contract method 0x361d2d7b.
//
// Solidity: function validateNonTeeBatch(address sender) view returns()
func (_BatchAuthenticator *BatchAuthenticatorCaller) ValidateNonTeeBatch(opts *bind.CallOpts, sender common.Address) error {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "validateNonTeeBatch", sender)

	if err != nil {
		return err
	}

	return err

}

// ValidateNonTeeBatch is a free data retrieval call binding the contract method 0x361d2d7b.
//
// Solidity: function validateNonTeeBatch(address sender) view returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) ValidateNonTeeBatch(sender common.Address) error {
	return _BatchAuthenticator.Contract.ValidateNonTeeBatch(&_BatchAuthenticator.CallOpts, sender)
}

// ValidateNonTeeBatch is a free data retrieval call binding the contract method 0x361d2d7b.
//
// Solidity: function validateNonTeeBatch(address sender) view returns()
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ValidateNonTeeBatch(sender common.Address) error {
	return _BatchAuthenticator.Contract.ValidateNonTeeBatch(&_BatchAuthenticator.CallOpts, sender)
}

// ValidateTeeBatch is a free data retrieval call binding the contract method 0x995bd81e.
//
// Solidity: function validateTeeBatch(address sender, bytes data) view returns()
func (_BatchAuthenticator *BatchAuthenticatorCaller) ValidateTeeBatch(opts *bind.CallOpts, sender common.Address, data []byte) error {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "validateTeeBatch", sender, data)

	if err != nil {
		return err
	}

	return err

}

// ValidateTeeBatch is a free data retrieval call binding the contract method 0x995bd81e.
//
// Solidity: function validateTeeBatch(address sender, bytes data) view returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) ValidateTeeBatch(sender common.Address, data []byte) error {
	return _BatchAuthenticator.Contract.ValidateTeeBatch(&_BatchAuthenticator.CallOpts, sender, data)
}

// ValidateTeeBatch is a free data retrieval call binding the contract method 0x995bd81e.
//
// Solidity: function validateTeeBatch(address sender, bytes data) view returns()
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) ValidateTeeBatch(sender common.Address, data []byte) error {
	return _BatchAuthenticator.Contract.ValidateTeeBatch(&_BatchAuthenticator.CallOpts, sender, data)
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

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _teeBatcher, address _nonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) Initialize(opts *bind.TransactOpts, _espressoTEEVerifier common.Address, _teeBatcher common.Address, _nonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "initialize", _espressoTEEVerifier, _teeBatcher, _nonTeeBatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _teeBatcher, address _nonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) Initialize(_espressoTEEVerifier common.Address, _teeBatcher common.Address, _nonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _teeBatcher, _nonTeeBatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _espressoTEEVerifier, address _teeBatcher, address _nonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) Initialize(_espressoTEEVerifier common.Address, _teeBatcher common.Address, _nonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.Initialize(&_BatchAuthenticator.TransactOpts, _espressoTEEVerifier, _teeBatcher, _nonTeeBatcher)
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

// SetNonTeeBatcher is a paid mutator transaction binding the contract method 0x8c8fc531.
//
// Solidity: function setNonTeeBatcher(address _newNonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SetNonTeeBatcher(opts *bind.TransactOpts, _newNonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "setNonTeeBatcher", _newNonTeeBatcher)
}

// SetNonTeeBatcher is a paid mutator transaction binding the contract method 0x8c8fc531.
//
// Solidity: function setNonTeeBatcher(address _newNonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SetNonTeeBatcher(_newNonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetNonTeeBatcher(&_BatchAuthenticator.TransactOpts, _newNonTeeBatcher)
}

// SetNonTeeBatcher is a paid mutator transaction binding the contract method 0x8c8fc531.
//
// Solidity: function setNonTeeBatcher(address _newNonTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SetNonTeeBatcher(_newNonTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetNonTeeBatcher(&_BatchAuthenticator.TransactOpts, _newNonTeeBatcher)
}

// SetTeeBatcher is a paid mutator transaction binding the contract method 0x6f7eda47.
//
// Solidity: function setTeeBatcher(address _newTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) SetTeeBatcher(opts *bind.TransactOpts, _newTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "setTeeBatcher", _newTeeBatcher)
}

// SetTeeBatcher is a paid mutator transaction binding the contract method 0x6f7eda47.
//
// Solidity: function setTeeBatcher(address _newTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) SetTeeBatcher(_newTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetTeeBatcher(&_BatchAuthenticator.TransactOpts, _newTeeBatcher)
}

// SetTeeBatcher is a paid mutator transaction binding the contract method 0x6f7eda47.
//
// Solidity: function setTeeBatcher(address _newTeeBatcher) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) SetTeeBatcher(_newTeeBatcher common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.SetTeeBatcher(&_BatchAuthenticator.TransactOpts, _newTeeBatcher)
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

// BatchAuthenticatorBatchInfoAuthenticatedIterator is returned from FilterBatchInfoAuthenticated and is used to iterate over the raw logs and unpacked data for BatchInfoAuthenticated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorBatchInfoAuthenticatedIterator struct {
	Event *BatchAuthenticatorBatchInfoAuthenticated // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorBatchInfoAuthenticatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorBatchInfoAuthenticated)
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
		it.Event = new(BatchAuthenticatorBatchInfoAuthenticated)
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
func (it *BatchAuthenticatorBatchInfoAuthenticatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorBatchInfoAuthenticatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorBatchInfoAuthenticated represents a BatchInfoAuthenticated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorBatchInfoAuthenticated struct {
	Commitment [32]byte
	Signer     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchInfoAuthenticated is a free log retrieval operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterBatchInfoAuthenticated(opts *bind.FilterOpts, commitment [][32]byte, signer []common.Address) (*BatchAuthenticatorBatchInfoAuthenticatedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "BatchInfoAuthenticated", commitmentRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorBatchInfoAuthenticatedIterator{contract: _BatchAuthenticator.contract, event: "BatchInfoAuthenticated", logs: logs, sub: sub}, nil
}

// WatchBatchInfoAuthenticated is a free log subscription operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchBatchInfoAuthenticated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorBatchInfoAuthenticated, commitment [][32]byte, signer []common.Address) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "BatchInfoAuthenticated", commitmentRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorBatchInfoAuthenticated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "BatchInfoAuthenticated", log); err != nil {
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

// ParseBatchInfoAuthenticated is a log parse operation binding the contract event 0x731978a77d438b0ea35a9034fb28d9cf9372e1649f18c213110adcfab65c5c5c.
//
// Solidity: event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseBatchInfoAuthenticated(log types.Log) (*BatchAuthenticatorBatchInfoAuthenticated, error) {
	event := new(BatchAuthenticatorBatchInfoAuthenticated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "BatchInfoAuthenticated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BatchAuthenticator contract.
type BatchAuthenticatorInitializedIterator struct {
	Event *BatchAuthenticatorInitialized // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorInitialized)
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
		it.Event = new(BatchAuthenticatorInitialized)
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
func (it *BatchAuthenticatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorInitialized represents a Initialized event raised by the BatchAuthenticator contract.
type BatchAuthenticatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*BatchAuthenticatorInitializedIterator, error) {

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorInitializedIterator{contract: _BatchAuthenticator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorInitialized) (event.Subscription, error) {

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorInitialized)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseInitialized(log types.Log) (*BatchAuthenticatorInitialized, error) {
	event := new(BatchAuthenticatorInitialized)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorNonTeeBatcherUpdatedIterator is returned from FilterNonTeeBatcherUpdated and is used to iterate over the raw logs and unpacked data for NonTeeBatcherUpdated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorNonTeeBatcherUpdatedIterator struct {
	Event *BatchAuthenticatorNonTeeBatcherUpdated // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorNonTeeBatcherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorNonTeeBatcherUpdated)
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
		it.Event = new(BatchAuthenticatorNonTeeBatcherUpdated)
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
func (it *BatchAuthenticatorNonTeeBatcherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorNonTeeBatcherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorNonTeeBatcherUpdated represents a NonTeeBatcherUpdated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorNonTeeBatcherUpdated struct {
	OldNonTeeBatcher common.Address
	NewNonTeeBatcher common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonTeeBatcherUpdated is a free log retrieval operation binding the contract event 0xc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5.
//
// Solidity: event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterNonTeeBatcherUpdated(opts *bind.FilterOpts, oldNonTeeBatcher []common.Address, newNonTeeBatcher []common.Address) (*BatchAuthenticatorNonTeeBatcherUpdatedIterator, error) {

	var oldNonTeeBatcherRule []interface{}
	for _, oldNonTeeBatcherItem := range oldNonTeeBatcher {
		oldNonTeeBatcherRule = append(oldNonTeeBatcherRule, oldNonTeeBatcherItem)
	}
	var newNonTeeBatcherRule []interface{}
	for _, newNonTeeBatcherItem := range newNonTeeBatcher {
		newNonTeeBatcherRule = append(newNonTeeBatcherRule, newNonTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "NonTeeBatcherUpdated", oldNonTeeBatcherRule, newNonTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorNonTeeBatcherUpdatedIterator{contract: _BatchAuthenticator.contract, event: "NonTeeBatcherUpdated", logs: logs, sub: sub}, nil
}

// WatchNonTeeBatcherUpdated is a free log subscription operation binding the contract event 0xc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5.
//
// Solidity: event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchNonTeeBatcherUpdated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorNonTeeBatcherUpdated, oldNonTeeBatcher []common.Address, newNonTeeBatcher []common.Address) (event.Subscription, error) {

	var oldNonTeeBatcherRule []interface{}
	for _, oldNonTeeBatcherItem := range oldNonTeeBatcher {
		oldNonTeeBatcherRule = append(oldNonTeeBatcherRule, oldNonTeeBatcherItem)
	}
	var newNonTeeBatcherRule []interface{}
	for _, newNonTeeBatcherItem := range newNonTeeBatcher {
		newNonTeeBatcherRule = append(newNonTeeBatcherRule, newNonTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "NonTeeBatcherUpdated", oldNonTeeBatcherRule, newNonTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorNonTeeBatcherUpdated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "NonTeeBatcherUpdated", log); err != nil {
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

// ParseNonTeeBatcherUpdated is a log parse operation binding the contract event 0xc85a6d3bc379dcb6b0bfec0a8d348be6dd937e2a34b2e2faaeb5762fc586aee5.
//
// Solidity: event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseNonTeeBatcherUpdated(log types.Log) (*BatchAuthenticatorNonTeeBatcherUpdated, error) {
	event := new(BatchAuthenticatorNonTeeBatcherUpdated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "NonTeeBatcherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorSignerRegistrationInitiatedIterator is returned from FilterSignerRegistrationInitiated and is used to iterate over the raw logs and unpacked data for SignerRegistrationInitiated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorSignerRegistrationInitiatedIterator struct {
	Event *BatchAuthenticatorSignerRegistrationInitiated // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorSignerRegistrationInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorSignerRegistrationInitiated)
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
		it.Event = new(BatchAuthenticatorSignerRegistrationInitiated)
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
func (it *BatchAuthenticatorSignerRegistrationInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorSignerRegistrationInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorSignerRegistrationInitiated represents a SignerRegistrationInitiated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorSignerRegistrationInitiated struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerRegistrationInitiated is a free log retrieval operation binding the contract event 0x665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c.
//
// Solidity: event SignerRegistrationInitiated(address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterSignerRegistrationInitiated(opts *bind.FilterOpts, caller []common.Address) (*BatchAuthenticatorSignerRegistrationInitiatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "SignerRegistrationInitiated", callerRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorSignerRegistrationInitiatedIterator{contract: _BatchAuthenticator.contract, event: "SignerRegistrationInitiated", logs: logs, sub: sub}, nil
}

// WatchSignerRegistrationInitiated is a free log subscription operation binding the contract event 0x665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c.
//
// Solidity: event SignerRegistrationInitiated(address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchSignerRegistrationInitiated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorSignerRegistrationInitiated, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "SignerRegistrationInitiated", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorSignerRegistrationInitiated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "SignerRegistrationInitiated", log); err != nil {
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

// ParseSignerRegistrationInitiated is a log parse operation binding the contract event 0x665b016a0ac50d1280744eaaff1cf21254d0fd30e4c3987d291913c32163416c.
//
// Solidity: event SignerRegistrationInitiated(address indexed caller)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseSignerRegistrationInitiated(log types.Log) (*BatchAuthenticatorSignerRegistrationInitiated, error) {
	event := new(BatchAuthenticatorSignerRegistrationInitiated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "SignerRegistrationInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchAuthenticatorTeeBatcherUpdatedIterator is returned from FilterTeeBatcherUpdated and is used to iterate over the raw logs and unpacked data for TeeBatcherUpdated events raised by the BatchAuthenticator contract.
type BatchAuthenticatorTeeBatcherUpdatedIterator struct {
	Event *BatchAuthenticatorTeeBatcherUpdated // Event containing the contract specifics and raw log

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
func (it *BatchAuthenticatorTeeBatcherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchAuthenticatorTeeBatcherUpdated)
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
		it.Event = new(BatchAuthenticatorTeeBatcherUpdated)
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
func (it *BatchAuthenticatorTeeBatcherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchAuthenticatorTeeBatcherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchAuthenticatorTeeBatcherUpdated represents a TeeBatcherUpdated event raised by the BatchAuthenticator contract.
type BatchAuthenticatorTeeBatcherUpdated struct {
	OldTeeBatcher common.Address
	NewTeeBatcher common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTeeBatcherUpdated is a free log retrieval operation binding the contract event 0x5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b3327.
//
// Solidity: event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) FilterTeeBatcherUpdated(opts *bind.FilterOpts, oldTeeBatcher []common.Address, newTeeBatcher []common.Address) (*BatchAuthenticatorTeeBatcherUpdatedIterator, error) {

	var oldTeeBatcherRule []interface{}
	for _, oldTeeBatcherItem := range oldTeeBatcher {
		oldTeeBatcherRule = append(oldTeeBatcherRule, oldTeeBatcherItem)
	}
	var newTeeBatcherRule []interface{}
	for _, newTeeBatcherItem := range newTeeBatcher {
		newTeeBatcherRule = append(newTeeBatcherRule, newTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.FilterLogs(opts, "TeeBatcherUpdated", oldTeeBatcherRule, newTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return &BatchAuthenticatorTeeBatcherUpdatedIterator{contract: _BatchAuthenticator.contract, event: "TeeBatcherUpdated", logs: logs, sub: sub}, nil
}

// WatchTeeBatcherUpdated is a free log subscription operation binding the contract event 0x5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b3327.
//
// Solidity: event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) WatchTeeBatcherUpdated(opts *bind.WatchOpts, sink chan<- *BatchAuthenticatorTeeBatcherUpdated, oldTeeBatcher []common.Address, newTeeBatcher []common.Address) (event.Subscription, error) {

	var oldTeeBatcherRule []interface{}
	for _, oldTeeBatcherItem := range oldTeeBatcher {
		oldTeeBatcherRule = append(oldTeeBatcherRule, oldTeeBatcherItem)
	}
	var newTeeBatcherRule []interface{}
	for _, newTeeBatcherItem := range newTeeBatcher {
		newTeeBatcherRule = append(newTeeBatcherRule, newTeeBatcherItem)
	}

	logs, sub, err := _BatchAuthenticator.contract.WatchLogs(opts, "TeeBatcherUpdated", oldTeeBatcherRule, newTeeBatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchAuthenticatorTeeBatcherUpdated)
				if err := _BatchAuthenticator.contract.UnpackLog(event, "TeeBatcherUpdated", log); err != nil {
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

// ParseTeeBatcherUpdated is a log parse operation binding the contract event 0x5186a10c46a3a9c7ec5470c24b80c6414eba1320cf76bf72ef5135773c7b3327.
//
// Solidity: event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher)
func (_BatchAuthenticator *BatchAuthenticatorFilterer) ParseTeeBatcherUpdated(log types.Log) (*BatchAuthenticatorTeeBatcherUpdated, error) {
	event := new(BatchAuthenticatorTeeBatcherUpdated)
	if err := _BatchAuthenticator.contract.UnpackLog(event, "TeeBatcherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
