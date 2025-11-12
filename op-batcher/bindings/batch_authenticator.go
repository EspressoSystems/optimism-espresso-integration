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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"},{\"name\":\"_preApprovedBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroValidator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preApprovedBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60e0604052346100665761001a610014610176565b91610269565b61002261006b565b611a2861041e82396080518181816101ae0152611214015260a0518181816107f201528181610bef0152611113015260c05181818160f10152610aa30152611a2890f35b610071565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061009d90610075565b810190811060018060401b038211176100b557604052565b61007f565b906100cd6100c661006b565b9283610093565b565b5f80fd5b60018060a01b031690565b6100e7906100d3565b90565b6100f3906100de565b90565b6100ff816100ea565b0361010657565b5f80fd5b90505190610117826100f6565b565b610122816100de565b0361012957565b5f80fd5b9050519061013a82610119565b565b90916060828403126101715761016e610157845f850161010a565b93610165816020860161012d565b9360400161012d565b90565b6100cf565b610194611e4680380380610189816100ba565b92833981019061013c565b909192565b6101a390516100ea565b90565b90565b6101bd6101b86101c2926100d3565b6101a6565b6100d3565b90565b6101ce906101a9565b90565b6101da906101c5565b90565b60e01b90565b6101ec906100de565b90565b6101f8816101e3565b036101ff57565b5f80fd5b90505190610210826101ef565b565b9060208282031261022b57610228915f01610203565b90565b6100cf565b5f0190565b61023d61006b565b3d5f823e3d90fd5b61024e906101c5565b90565b61025a906101a9565b90565b61026690610251565b90565b906102a7929161027761031b565b60a052608052602061029161028c60a0610199565b6101d1565b63d80a4c289061029f61006b565b9485926101dd565b825281806102b760048201610230565b03915afa8015610316576102d96102de916102e6945f916102e8575b50610245565b61025d565b60c0526103ad565b565b610309915060203d811161030f575b6103018183610093565b810190610212565b5f6102d3565b503d6102f7565b610235565b61032b610326610410565b6103ad565b565b5f1c90565b60018060a01b031690565b61034961034e9161032d565b610332565b90565b61035b905461033d565b90565b5f1b90565b9061037460018060a01b039161035e565b9181191691161790565b610387906101c5565b90565b90565b906103a261039d6103a99261037e565b61038a565b8254610363565b9055565b6103b65f610351565b6103c0825f61038d565b906103f46103ee7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361037e565b9161037e565b916103fd61006b565b8061040781610230565b0390a3565b5f90565b61041861040c565b50339056fe60806040526004361015610013575b6108f1565b61001d5f356100cc565b80631b076a4c146100c75780631f568b18146100c257806354fd4d50146100bd578063715018a6146100b85780638da5cb5b146100b3578063a903a277146100ae578063ba58e82a146100a9578063f2fde38b146100a4578063f81f20831461009f578063fa14fe6d1461009a5763fc619e410361000e576108bd565b610842565b6107bb565b6106b2565b61063a565b61055e565b6103f8565b6103c5565b61038b565b6101fe565b610177565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126100ea57565b6100dc565b7f000000000000000000000000000000000000000000000000000000000000000090565b60018060a01b031690565b90565b61013561013061013a92610113565b61011e565b610113565b90565b61014690610121565b90565b6101529061013d565b90565b61015e90610149565b9052565b9190610175905f60208501940190610155565b565b346101a7576101873660046100e0565b6101a36101926100ef565b61019a6100d2565b91829182610162565b0390f35b6100d8565b7f000000000000000000000000000000000000000000000000000000000000000090565b6101d990610113565b90565b6101e5906101d0565b9052565b91906101fc905f602085019401906101dc565b565b3461022e5761020e3660046100e0565b61022a6102196101ac565b6102216100d2565b918291826101e9565b0390f35b6100d8565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061025b90610233565b810190811067ffffffffffffffff82111761027557604052565b61023d565b9061028d6102866100d2565b9283610251565b565b67ffffffffffffffff81116102ad576102a9602091610233565b0190565b61023d565b906102c46102bf8361028f565b61027a565b918252565b5f7f312e302e30000000000000000000000000000000000000000000000000000000910152565b6102fa60056102b2565b90610307602083016102c9565b565b6103116102f0565b90565b61031c610309565b90565b610327610314565b90565b5190565b60209181520190565b90825f9392825e0152565b61036161036a60209361036f936103588161032a565b9384809361032e565b95869101610337565b610233565b0190565b6103889160208201915f818403910152610342565b90565b346103bb5761039b3660046100e0565b6103b76103a661031f565b6103ae6100d2565b91829182610373565b0390f35b6100d8565b5f0190565b346103f3576103d53660046100e0565b6103dd610945565b6103e56100d2565b806103ef816103c0565b0390f35b6100d8565b34610428576104083660046100e0565b610424610413610984565b61041b6100d2565b918291826101e9565b0390f35b6100d8565b5f80fd5b5f80fd5b5f80fd5b67ffffffffffffffff811161045757610453602091610233565b0190565b61023d565b90825f939282370152565b9092919261047c61047782610439565b61027a565b93818552602085019082840111610498576104969261045c565b565b610435565b9080601f830112156104bb578160206104b893359101610467565b90565b610431565b906020828203126104f0575f82013567ffffffffffffffff81116104eb576104e8920161049d565b90565b61042d565b6100dc565b5190565b60209181520190565b61052161052a60209361052f93610518816104f5565b938480936104f9565b95869101610337565b610233565b0190565b909161054d61055b9360408401908482035f860152610502565b916020818403910152610502565b90565b3461058f576105766105713660046104c0565b610a86565b9061058b6105826100d2565b92839283610533565b0390f35b6100d8565b5f80fd5b5f80fd5b909182601f830112156105d65781359167ffffffffffffffff83116105d15760200192600183028401116105cc57565b610598565b610594565b610431565b9091604082840312610635575f82013567ffffffffffffffff8111610630578361060691840161059c565b929093602082013567ffffffffffffffff811161062b57610627920161059c565b9091565b61042d565b61042d565b6100dc565b3461066c5761065661064d3660046105db565b92919091610be7565b61065e6100d2565b80610668816103c0565b0390f35b6100d8565b61067a816101d0565b0361068157565b5f80fd5b9050359061069282610671565b565b906020828203126106ad576106aa915f01610685565b90565b6100dc565b346106e0576106ca6106c5366004610694565b610d85565b6106d26100d2565b806106dc816103c0565b0390f35b6100d8565b90565b6106f1816106e5565b036106f857565b5f80fd5b90503590610709826106e8565b565b9060208282031261072457610721915f016106fc565b90565b6100dc565b610732906106e5565b90565b9061073f90610729565b5f5260205260405f2090565b1c90565b60ff1690565b61076590600861076a930261074b565b61074f565b90565b906107789154610755565b90565b6107919061078c6001915f92610735565b61076d565b90565b151590565b6107a290610794565b9052565b91906107b9905f60208501940190610799565b565b346107eb576107e76107d66107d136600461070b565b61077b565b6107de6100d2565b918291826107a6565b0390f35b6100d8565b7f000000000000000000000000000000000000000000000000000000000000000090565b61081d9061013d565b90565b61082990610814565b9052565b9190610840905f60208501940190610820565b565b34610872576108523660046100e0565b61086e61085d6107f0565b6108656100d2565b9182918261082d565b0390f35b6100d8565b9190916040818403126108b857610890835f83016106fc565b92602082013567ffffffffffffffff81116108b3576108af920161059c565b9091565b61042d565b6100dc565b346108ec576108d66108d0366004610877565b91611090565b6108de6100d2565b806108e8816103c0565b0390f35b6100d8565b5f80fd5b6108fd6113a4565b610905610932565b565b90565b61091e61091961092392610907565b61011e565b610113565b90565b61092f9061090a565b90565b61094361093e5f610926565b61141a565b565b61094d6108f5565b565b5f90565b5f1c90565b60018060a01b031690565b61096f61097491610953565b610958565b90565b6109819054610963565b90565b61098c61094f565b506109965f610977565b90565b606090565b5f80fd5b60e01b90565b909291926109bd6109b882610439565b61027a565b938185526020850190828401116109d9576109d792610337565b565b610435565b9080601f830112156109fc578160206109f9935191016109a8565b90565b610431565b919091604081840312610a59575f81015167ffffffffffffffff8111610a545783610a2d9183016109de565b92602082015167ffffffffffffffff8111610a4f57610a4c92016109de565b90565b61042d565b61042d565b6100dc565b610a739160208201915f818403910152610502565b90565b610a7e6100d2565b3d5f823e3d90fd5b905f610aee92610a94610999565b50610a9d610999565b50610ac77f0000000000000000000000000000000000000000000000000000000000000000610149565b610ae363a903a277610ad76100d2565b968794859384936109a2565b835260048301610a5e565b03915afa8015610b2e575f80939091610b07575b509190565b9050610b269192503d805f833e610b1e8183610251565b810190610a01565b91905f610b02565b610a76565b5f910312610b3d57565b6100dc565b9190610b5c81610b5581610b61956104f9565b809561045c565b610233565b0190565b634e487b7160e01b5f52602160045260245ffd5b60021115610b8357565b610b65565b90610b9282610b79565b565b610b9d90610b88565b90565b610ba990610b94565b9052565b959492610be594610bcf610bdd9360409560608b01918b83035f8d0152610b42565b9188830360208a0152610b42565b940190610ba0565b565b929192610c137f0000000000000000000000000000000000000000000000000000000000000000610814565b906335ecb4c190929493600191833b15610c9557610c52610c47935f97938894610c3b6100d2565b9a8b998a9889976109a2565b875260048701610bad565b03925af18015610c9057610c64575b50565b610c83905f3d8111610c89575b610c7b8183610251565b810190610b33565b5f610c61565b503d610c71565b610a76565b61099e565b610cab90610ca66113a4565b610d55565b565b60207f6464726573730000000000000000000000000000000000000000000000000000917f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201520152565b610d07602660409261032e565b610d1081610cad565b0190565b610d299060208101905f818303910152610cfa565b90565b15610d3357565b610d3b6100d2565b62461bcd60e51b815280610d5160048201610d14565b0390fd5b610d8390610d7e81610d77610d71610d6c5f610926565b6101d0565b916101d0565b1415610d2c565b61141a565b565b610d8e90610c9a565b565b610d9b913691610467565b90565b634e487b7160e01b5f52603260045260245ffd5b90610dbc826104f5565b811015610dce57600160209102010190565b610d9e565b90565b90565b610ded610de8610df292610dd3565b61011e565b610dd6565b90565b60ff60f81b1690565b610e089051610df5565b90565b60f81c90565b60ff1690565b610e2b610e26610e3092610e11565b61011e565b610e11565b90565b610e3f610e4491610e0b565b610e17565b90565b610e5b610e56610e6092610907565b61011e565b610e11565b90565b90565b610e7a610e75610e7f92610e63565b61011e565b610e11565b90565b90565b610e99610e94610e9e92610e82565b61011e565b610e11565b90565b634e487b7160e01b5f52601160045260245ffd5b610ec1610ec791610e11565b91610e11565b019060ff8211610ed357565b610ea1565b60f81b90565b610ef2610eed610ef792610e11565b610ed8565b610df5565b90565b5f7f496e76616c6964207369676e6174757265000000000000000000000000000000910152565b610f2e601160209261032e565b610f3781610efa565b0190565b610f509060208101905f818303910152610f21565b90565b610f5c906101d0565b90565b610f6881610f53565b03610f6f57565b5f80fd5b90505190610f8082610f5f565b565b90602082820312610f9b57610f98915f01610f73565b90565b6100dc565b610fa99061013d565b90565b610fb581610794565b03610fbc57565b5f80fd5b90505190610fcd82610fac565b565b90602082820312610fe857610fe5915f01610fc0565b90565b6100dc565b5f7f496e76616c6964207369676e6572000000000000000000000000000000000000910152565b611021600e60209261032e565b61102a81610fed565b0190565b6110439060208101905f818303910152611014565b90565b5f1b90565b9061105760ff91611046565b9181191691161790565b61106a90610794565b90565b90565b9061108561108061108c92611061565b61106d565b825461104b565b9055565b9161109e906110e992610d90565b6110c26110bd6110b8836110b26040610dd9565b90610db2565b610dfe565b610e33565b806110d56110cf5f610e47565b91610e11565b148015611308575b6112cd575b5082611479565b806111046110fe6110f95f610926565b6101d0565b916101d0565b146112ab5761114d60206111377f0000000000000000000000000000000000000000000000000000000000000000610814565b63d80a4c28906111456100d2565b9384926109a2565b8252818061115d600482016103c0565b03915afa80156112a65761117e6020916111a8935f91611279575b50610fa0565b630123d0c19061119d85926111916100d2565b958694859384936109a2565b8352600483016101e9565b03915afa8015611274576111c4915f91611246575b5015610794565b908161120a575b506111e8576111e6906111e16001916001610735565b611070565b565b6111f06100d2565b62461bcd60e51b8152806112066004820161102e565b0390fd5b905061123e6112387f00000000000000000000000000000000000000000000000000000000000000006101d0565b916101d0565b14155f6111cb565b611267915060203d811161126d575b61125f8183610251565b810190610fcf565b5f6111bd565b503d611255565b610a76565b6112999150833d811161129f575b6112918183610251565b810190610f82565b5f611178565b503d611287565b610a76565b6112b36100d2565b62461bcd60e51b8152806112c960048201610f3b565b0390fd5b6112e46112e9916112de601b610e85565b90610eb5565b610ede565b611301826112fb6040935f1a93610dd9565b90610db2565b535f6110e2565b508061131d6113176001610e66565b91610e11565b146110dd565b5f7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572910152565b6113566020809261032e565b61135f81611323565b0190565b6113789060208101905f81830391015261134a565b90565b1561138257565b61138a6100d2565b62461bcd60e51b8152806113a060048201611363565b0390fd5b6113ce6113af610984565b6113c86113c26113bd61149a565b6101d0565b916101d0565b1461137b565b565b906113e160018060a01b0391611046565b9181191691161790565b6113f49061013d565b90565b90565b9061140f61140a611416926113eb565b6113f7565b82546113d0565b9055565b6114235f610977565b61142d825f6113fa565b9061146161145b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936113eb565b916113eb565b9161146a6100d2565b80611474816103c0565b0390a3565b6114979161148f9161148961094f565b506114d2565b919091611721565b90565b6114a261094f565b503390565b5f90565b90565b6114c26114bd6114c7926114ab565b61011e565b610dd6565b90565b5f90565b5f90565b6114da61094f565b506114e36114a7565b506114ed826104f5565b6115006114fa60416114ae565b91610dd6565b145f146115455761153f916115136114ca565b5061151c6114ca565b506115256114ce565b506020810151606060408301519201515f1a9091926118f7565b91909190565b50506115505f610926565b90600290565b6005111561156057565b610b65565b9061156f82611556565b565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202776272076616c5f8201520152565b6115cb602260409261032e565b6115d481611571565b0190565b6115ed9060208101905f8183039101526115be565b90565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201520152565b61164a602260409261032e565b611653816115f0565b0190565b61166c9060208101905f81830391015261163d565b90565b5f7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800910152565b6116a3601f60209261032e565b6116ac8161166f565b0190565b6116c59060208101905f818303910152611696565b90565b5f7f45434453413a20696e76616c6964207369676e61747572650000000000000000910152565b6116fc601860209261032e565b611705816116c8565b0190565b61171e9060208101905f8183039101526116ef565b90565b8061173461172e5f611565565b91611565565b145f1461173e5750565b8061175261174c6001611565565b91611565565b145f1461177b576117616100d2565b62461bcd60e51b81528061177760048201611709565b0390fd5b8061178f6117896002611565565b91611565565b145f146117b85761179e6100d2565b62461bcd60e51b8152806117b4600482016116b0565b0390fd5b806117cc6117c66003611565565b91611565565b145f146117f5576117db6100d2565b62461bcd60e51b8152806117f160048201611657565b0390fd5b6118086118026004611565565b91611565565b1461180f57565b6118176100d2565b62461bcd60e51b81528061182d600482016115d8565b0390fd5b61184561184061184a92610dd6565b61011e565b610dd6565b90565b61185961185e91610953565b611831565b90565b90565b61187861187361187d92611861565b61011e565b610dd6565b90565b90565b61189761189261189c92611880565b61011e565b610e11565b90565b6118a8906106e5565b9052565b6118b590610e11565b9052565b6118ee6118f5946118e46060949897956118da608086019a5f87019061189f565b60208501906118ac565b604083019061189f565b019061189f565b565b92919061190261094f565b5061190b6114a7565b506119158361184d565b6119476119417f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0611864565b91610dd6565b11611a08578061196061195a601b610e85565b91610e11565b1415806119ec575b6119d9576119875f93602095929361197e6100d2565b948594856118b9565b838052039060015afa156119d45761199f5f51611046565b806119ba6119b46119af5f610926565b6101d0565b916101d0565b146119c457905f90565b506119ce5f610926565b90600190565b610a76565b505050506119e65f610926565b90600490565b5080611a016119fb601c611883565b91610e11565b1415611968565b50505050611a155f610926565b9060039056fea164736f6c634300081d000a",
}

// BatchAuthenticatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchAuthenticatorMetaData.ABI instead.
var BatchAuthenticatorABI = BatchAuthenticatorMetaData.ABI

// BatchAuthenticatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchAuthenticatorMetaData.Bin instead.
var BatchAuthenticatorBin = BatchAuthenticatorMetaData.Bin

// DeployBatchAuthenticator deploys a new Ethereum contract, binding an instance of BatchAuthenticator to it.
func DeployBatchAuthenticator(auth *bind.TransactOpts, backend bind.ContractBackend, _espressoTEEVerifier common.Address, _preApprovedBatcher common.Address, _owner common.Address) (common.Address, *types.Transaction, *BatchAuthenticator, error) {
	parsed, err := BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchAuthenticatorBin), backend, _espressoTEEVerifier, _preApprovedBatcher, _owner)
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

// PreApprovedBatcher is a free data retrieval call binding the contract method 0x1f568b18.
//
// Solidity: function preApprovedBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCaller) PreApprovedBatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchAuthenticator.contract.Call(opts, &out, "preApprovedBatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreApprovedBatcher is a free data retrieval call binding the contract method 0x1f568b18.
//
// Solidity: function preApprovedBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorSession) PreApprovedBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.PreApprovedBatcher(&_BatchAuthenticator.CallOpts)
}

// PreApprovedBatcher is a free data retrieval call binding the contract method 0x1f568b18.
//
// Solidity: function preApprovedBatcher() view returns(address)
func (_BatchAuthenticator *BatchAuthenticatorCallerSession) PreApprovedBatcher() (common.Address, error) {
	return _BatchAuthenticator.Contract.PreApprovedBatcher(&_BatchAuthenticator.CallOpts)
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
