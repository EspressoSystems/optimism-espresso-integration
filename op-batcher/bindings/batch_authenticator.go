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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroValidator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516120ac3803806120ac833981810160405281019061003291906103fb565b61004e61004361029a60201b60201c565b6102a160201b60201c565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036100bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100b3906104df565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361012a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101219061056d565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff168152505060c05173ffffffffffffffffffffffffffffffffffffffff1663d80a4c286040518163ffffffff1660e01b8152600401602060405180830381865afa158015610211573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061023591906105c6565b73ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1681525050600160025f6101000a81548160ff021916908315150217905550610291816102a160201b60201c565b505050506105f1565b5f33905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61038f82610366565b9050919050565b5f6103a082610385565b9050919050565b6103b081610396565b81146103ba575f5ffd5b50565b5f815190506103cb816103a7565b92915050565b6103da81610385565b81146103e4575f5ffd5b50565b5f815190506103f5816103d1565b92915050565b5f5f5f5f6080858703121561041357610412610362565b5b5f610420878288016103bd565b9450506020610431878288016103e7565b9350506040610442878288016103e7565b9250506060610453878288016103e7565b91505092959194509250565b5f82825260208201905092915050565b7f426174636841757468656e74696361746f723a207a65726f20746565206261745f8201527f6368657200000000000000000000000000000000000000000000000000000000602082015250565b5f6104c960248361045f565b91506104d48261046f565b604082019050919050565b5f6020820190508181035f8301526104f6816104bd565b9050919050565b7f426174636841757468656e74696361746f723a207a65726f206e6f6e2d7465655f8201527f2062617463686572000000000000000000000000000000000000000000000000602082015250565b5f61055760288361045f565b9150610562826104fd565b604082019050919050565b5f6020820190508181035f8301526105848161054b565b9050919050565b5f61059582610385565b9050919050565b6105a58161058b565b81146105af575f5ffd5b50565b5f815190506105c08161059c565b92915050565b5f602082840312156105db576105da610362565b5b5f6105e8848285016105b2565b91505092915050565b60805160a05160c05160e051611a6c6106405f395f8181610289015261033501525f81816103fd01528181610587015261070101525f6103d901525f81816104c4015261080d0152611a6c5ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c8063ba58e82a1161008a578063f2fde38b11610064578063f2fde38b14610201578063f81f20831461021d578063fa14fe6d1461024d578063fc619e411461026b576100e8565b8063ba58e82a146101bd578063bc347f47146101d9578063d909ba7c146101e3576100e8565b80637877a9ed116100c65780637877a9ed146101325780638da5cb5b14610150578063a903a2771461016e578063b1bd42851461019f576100e8565b80631b076a4c146100ec57806354fd4d501461010a578063715018a614610128575b5f5ffd5b6100f4610287565b6040516101019190610dc9565b60405180910390f35b6101126102ab565b60405161011f9190610e52565b60405180910390f35b6101306102e4565b005b61013a6102f7565b6040516101479190610e8c565b60405180910390f35b610158610309565b6040516101659190610ec5565b60405180910390f35b6101886004803603810190610183919061101b565b610330565b6040516101969291906110b4565b60405180910390f35b6101a76103d7565b6040516101b49190610ec5565b60405180910390f35b6101d760048036038101906101d29190611146565b6103fb565b005b6101e1610490565b005b6101eb6104c2565b6040516101f89190610ec5565b60405180910390f35b61021b600480360381019061021691906111ee565b6104e6565b005b6102376004803603810190610232919061124c565b610568565b6040516102449190610e8c565b60405180910390f35b610255610585565b6040516102629190611297565b60405180910390f35b610285600480360381019061028091906112b0565b6105a9565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6102ec6108cc565b6102f55f61094a565b565b60025f9054906101000a900460ff1681565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a903a277846040518263ffffffff1660e01b815260040161038c919061130d565b5f60405180830381865afa1580156103a6573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906103ce919061139b565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166335ecb4c18585858560016040518663ffffffff1660e01b815260040161045d9594939291906114b0565b5f604051808303815f87803b158015610474575f5ffd5b505af1158015610486573d5f5f3e3d5ffd5b5050505050505050565b6104986108cc565b60025f9054906101000a900460ff161560025f6101000a81548160ff021916908315150217905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b6104ee6108cc565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361055c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055390611567565b60405180910390fd5b6105658161094a565b50565b6001602052805f5260405f205f915054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f82828080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090505f8160408151811061060557610604611585565b5b602001015160f81c60f81b60f81c90505f8160ff161480610629575060018160ff16145b1561068457601b8161063b91906115eb565b90508060f81b8260408151811061065557610654611585565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b5f61068f8684610a0b565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f690611669565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d80a4c286040518163ffffffff1660e01b8152600401602060405180830381865afa158015610768573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078c91906116c2565b73ffffffffffffffffffffffffffffffffffffffff16630123d0c1826040518263ffffffff1660e01b81526004016107c49190610ec5565b602060405180830381865afa1580156107df573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108039190611717565b15801561085c57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b1561089c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108939061178c565b60405180910390fd5b6001805f8881526020019081526020015f205f6101000a81548160ff021916908315150217905550505050505050565b6108d4610a30565b73ffffffffffffffffffffffffffffffffffffffff166108f2610309565b73ffffffffffffffffffffffffffffffffffffffff1614610948576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093f906117f4565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5f5f610a188585610a37565b91509150610a2581610a83565b819250505092915050565b5f33905090565b5f5f6041835103610a74575f5f5f602086015192506040860151915060608601515f1a9050610a6887828585610c4e565b94509450505050610a7c565b5f6002915091505b9250929050565b5f6004811115610a9657610a9561143d565b5b816004811115610aa957610aa861143d565b5b0315610c4b5760016004811115610ac357610ac261143d565b5b816004811115610ad657610ad561143d565b5b03610b16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0d9061185c565b60405180910390fd5b60026004811115610b2a57610b2961143d565b5b816004811115610b3d57610b3c61143d565b5b03610b7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b74906118c4565b60405180910390fd5b60036004811115610b9157610b9061143d565b5b816004811115610ba457610ba361143d565b5b03610be4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bdb90611952565b60405180910390fd5b600480811115610bf757610bf661143d565b5b816004811115610c0a57610c0961143d565b5b03610c4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c41906119e0565b60405180910390fd5b5b50565b5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0835f1c1115610c86575f600391509150610d46565b601b8560ff1614158015610c9e5750601c8560ff1614155b15610caf575f600491509150610d46565b5f6001878787876040515f8152602001604052604051610cd29493929190611a1c565b6020604051602081039080840390855afa158015610cf2573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d3e575f60019250925050610d46565b805f92509250505b94509492505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610d91610d8c610d8784610d4f565b610d6e565b610d4f565b9050919050565b5f610da282610d77565b9050919050565b5f610db382610d98565b9050919050565b610dc381610da9565b82525050565b5f602082019050610ddc5f830184610dba565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e2482610de2565b610e2e8185610dec565b9350610e3e818560208601610dfc565b610e4781610e0a565b840191505092915050565b5f6020820190508181035f830152610e6a8184610e1a565b905092915050565b5f8115159050919050565b610e8681610e72565b82525050565b5f602082019050610e9f5f830184610e7d565b92915050565b5f610eaf82610d4f565b9050919050565b610ebf81610ea5565b82525050565b5f602082019050610ed85f830184610eb6565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f2d82610e0a565b810181811067ffffffffffffffff82111715610f4c57610f4b610ef7565b5b80604052505050565b5f610f5e610ede565b9050610f6a8282610f24565b919050565b5f67ffffffffffffffff821115610f8957610f88610ef7565b5b610f9282610e0a565b9050602081019050919050565b828183375f83830152505050565b5f610fbf610fba84610f6f565b610f55565b905082815260208101848484011115610fdb57610fda610ef3565b5b610fe6848285610f9f565b509392505050565b5f82601f83011261100257611001610eef565b5b8135611012848260208601610fad565b91505092915050565b5f602082840312156110305761102f610ee7565b5b5f82013567ffffffffffffffff81111561104d5761104c610eeb565b5b61105984828501610fee565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f61108682611062565b611090818561106c565b93506110a0818560208601610dfc565b6110a981610e0a565b840191505092915050565b5f6040820190508181035f8301526110cc818561107c565b905081810360208301526110e0818461107c565b90509392505050565b5f5ffd5b5f5ffd5b5f5f83601f84011261110657611105610eef565b5b8235905067ffffffffffffffff811115611123576111226110e9565b5b60208301915083600182028301111561113f5761113e6110ed565b5b9250929050565b5f5f5f5f6040858703121561115e5761115d610ee7565b5b5f85013567ffffffffffffffff81111561117b5761117a610eeb565b5b611187878288016110f1565b9450945050602085013567ffffffffffffffff8111156111aa576111a9610eeb565b5b6111b6878288016110f1565b925092505092959194509250565b6111cd81610ea5565b81146111d7575f5ffd5b50565b5f813590506111e8816111c4565b92915050565b5f6020828403121561120357611202610ee7565b5b5f611210848285016111da565b91505092915050565b5f819050919050565b61122b81611219565b8114611235575f5ffd5b50565b5f8135905061124681611222565b92915050565b5f6020828403121561126157611260610ee7565b5b5f61126e84828501611238565b91505092915050565b5f61128182610d98565b9050919050565b61129181611277565b82525050565b5f6020820190506112aa5f830184611288565b92915050565b5f5f5f604084860312156112c7576112c6610ee7565b5b5f6112d486828701611238565b935050602084013567ffffffffffffffff8111156112f5576112f4610eeb565b5b611301868287016110f1565b92509250509250925092565b5f6020820190508181035f830152611325818461107c565b905092915050565b5f61133f61133a84610f6f565b610f55565b90508281526020810184848401111561135b5761135a610ef3565b5b611366848285610dfc565b509392505050565b5f82601f83011261138257611381610eef565b5b815161139284826020860161132d565b91505092915050565b5f5f604083850312156113b1576113b0610ee7565b5b5f83015167ffffffffffffffff8111156113ce576113cd610eeb565b5b6113da8582860161136e565b925050602083015167ffffffffffffffff8111156113fb576113fa610eeb565b5b6114078582860161136e565b9150509250929050565b5f61141c838561106c565b9350611429838584610f9f565b61143283610e0a565b840190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6002811061147b5761147a61143d565b5b50565b5f81905061148b8261146a565b919050565b5f61149a8261147e565b9050919050565b6114aa81611490565b82525050565b5f6060820190508181035f8301526114c9818789611411565b905081810360208301526114de818587611411565b90506114ed60408301846114a1565b9695505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611551602683610dec565b915061155c826114f7565b604082019050919050565b5f6020820190508181035f83015261157e81611545565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60ff82169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6115f5826115b2565b9150611600836115b2565b9250828201905060ff811115611619576116186115be565b5b92915050565b7f496e76616c6964207369676e61747572650000000000000000000000000000005f82015250565b5f611653601183610dec565b915061165e8261161f565b602082019050919050565b5f6020820190508181035f83015261168081611647565b9050919050565b5f61169182610ea5565b9050919050565b6116a181611687565b81146116ab575f5ffd5b50565b5f815190506116bc81611698565b92915050565b5f602082840312156116d7576116d6610ee7565b5b5f6116e4848285016116ae565b91505092915050565b6116f681610e72565b8114611700575f5ffd5b50565b5f81519050611711816116ed565b92915050565b5f6020828403121561172c5761172b610ee7565b5b5f61173984828501611703565b91505092915050565b7f496e76616c6964207369676e65720000000000000000000000000000000000005f82015250565b5f611776600e83610dec565b915061178182611742565b602082019050919050565b5f6020820190508181035f8301526117a38161176a565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6117de602083610dec565b91506117e9826117aa565b602082019050919050565b5f6020820190508181035f83015261180b816117d2565b9050919050565b7f45434453413a20696e76616c6964207369676e617475726500000000000000005f82015250565b5f611846601883610dec565b915061185182611812565b602082019050919050565b5f6020820190508181035f8301526118738161183a565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e677468005f82015250565b5f6118ae601f83610dec565b91506118b98261187a565b602082019050919050565b5f6020820190508181035f8301526118db816118a2565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f61193c602283610dec565b9150611947826118e2565b604082019050919050565b5f6020820190508181035f83015261196981611930565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202776272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f6119ca602283610dec565b91506119d582611970565b604082019050919050565b5f6020820190508181035f8301526119f7816119be565b9050919050565b611a0781611219565b82525050565b611a16816115b2565b82525050565b5f608082019050611a2f5f8301876119fe565b611a3c6020830186611a0d565b611a4960408301856119fe565b611a5660608301846119fe565b9594505050505056fea164736f6c634300081e000a",
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

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) view returns(bytes, bytes)
func (_BatchAuthenticator *BatchAuthenticatorCaller) DecodeAttestationTbs(opts *bind.CallOpts, attestation []byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "decodeAttestationTbs", attestation)

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
func (_BatchAuthenticator *BatchAuthenticatorSession) DecodeAttestationTbs(attestation []byte) ([]byte, []byte, error) {
	return _BatchAuthenticator.Contract.DecodeAttestationTbs(&_BatchAuthenticator.CallOpts, attestation)
}

// DecodeAttestationTbs is a free data retrieval call binding the contract method 0xa903a277.
//
// Solidity: function decodeAttestationTbs(bytes attestation) view returns(bytes, bytes)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) DecodeAttestationTbs(attestation []byte) ([]byte, []byte, error) {
	return _BatchAuthenticator.Contract.DecodeAttestationTbs(&_BatchAuthenticator.CallOpts, attestation)
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
