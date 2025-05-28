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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"},{\"name\":\"_preApprovedBatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroValidator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preApprovedBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60e0604052346100665761001a610014610169565b9061025b565b61002261006b565b611ba56102fc82396080518181816101bc01526112e5015260a05181818161081901528181610c3e01526111ca015260c05181818160f10152610ad90152611ba590f35b610071565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061009d90610075565b810190811060018060401b038211176100b557604052565b61007f565b906100cd6100c661006b565b9283610093565b565b5f80fd5b60018060a01b031690565b6100e7906100d3565b90565b6100f3906100de565b90565b6100ff816100ea565b0361010657565b5f80fd5b90505190610117826100f6565b565b610122816100de565b0361012957565b5f80fd5b9050519061013a82610119565b565b91906040838203126101645780610158610161925f860161010a565b9360200161012d565b90565b6100cf565b610187611ea18038038061017c816100ba565b92833981019061013c565b9091565b61019590516100ea565b90565b90565b6101af6101aa6101b4926100d3565b610198565b6100d3565b90565b6101c09061019b565b90565b6101cc906101b7565b90565b60e01b90565b6101de906100de565b90565b6101ea816101d5565b036101f157565b5f80fd5b90505190610202826101e1565b565b9060208282031261021d5761021a915f016101f5565b90565b6100cf565b5f0190565b61022f61006b565b3d5f823e3d90fd5b610240906101b7565b90565b61024c9061019b565b90565b61025890610243565b90565b60a05260805261028e602061027861027360a061018b565b6101c3565b63d80a4c289061028661006b565b9384926101cf565b8252818061029e60048201610222565b03915afa9081156102f6576102c3916102be915f916102c8575b50610237565b61024f565b60c052565b6102e9915060203d81116102ef575b6102e18183610093565b810190610204565b5f6102b8565b503d6102d7565b61022756fe60806040526004361015610013575b610918565b61001d5f356100cc565b80631b076a4c146100c75780631f568b18146100c257806354fd4d50146100bd578063715018a6146100b85780638da5cb5b146100b3578063a903a277146100ae578063ba58e82a146100a9578063f2fde38b146100a4578063f81f20831461009f578063fa14fe6d1461009a5763fc619e410361000e576108e4565b610869565b6107e2565b6106d9565b610661565b610585565b61041f565b6103ec565b6103b2565b61020c565b610185565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126100ea57565b6100dc565b7f000000000000000000000000000000000000000000000000000000000000000090565b73ffffffffffffffffffffffffffffffffffffffff1690565b90565b61014361013e61014892610113565b61012c565b610113565b90565b6101549061012f565b90565b6101609061014b565b90565b61016c90610157565b9052565b9190610183905f60208501940190610163565b565b346101b5576101953660046100e0565b6101b16101a06100ef565b6101a86100d2565b91829182610170565b0390f35b6100d8565b7f000000000000000000000000000000000000000000000000000000000000000090565b6101e790610113565b90565b6101f3906101de565b9052565b919061020a905f602085019401906101ea565b565b3461023c5761021c3660046100e0565b6102386102276101ba565b61022f6100d2565b918291826101f7565b0390f35b6100d8565b601f801991011690565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b9061028290610241565b810190811067ffffffffffffffff82111761029c57604052565b61024b565b906102b46102ad6100d2565b9283610278565b565b67ffffffffffffffff81116102d4576102d0602091610241565b0190565b61024b565b906102eb6102e6836102b6565b6102a1565b918252565b5f7f312e302e30000000000000000000000000000000000000000000000000000000910152565b61032160056102d9565b9061032e602083016102f0565b565b610338610317565b90565b610343610330565b90565b61034e61033b565b90565b5190565b60209181520190565b90825f9392825e0152565b6103886103916020936103969361037f81610351565b93848093610355565b9586910161035e565b610241565b0190565b6103af9160208201915f818403910152610369565b90565b346103e2576103c23660046100e0565b6103de6103cd610346565b6103d56100d2565b9182918261039a565b0390f35b6100d8565b5f0190565b3461041a576103fc3660046100e0565b61040461096c565b61040c6100d2565b80610416816103e7565b0390f35b6100d8565b3461044f5761042f3660046100e0565b61044b61043a6109b9565b6104426100d2565b918291826101f7565b0390f35b6100d8565b5f80fd5b5f80fd5b5f80fd5b67ffffffffffffffff811161047e5761047a602091610241565b0190565b61024b565b90825f939282370152565b909291926104a361049e82610460565b6102a1565b938185526020850190828401116104bf576104bd92610483565b565b61045c565b9080601f830112156104e2578160206104df9335910161048e565b90565b610458565b90602082820312610517575f82013567ffffffffffffffff81116105125761050f92016104c4565b90565b610454565b6100dc565b5190565b60209181520190565b6105486105516020936105569361053f8161051c565b93848093610520565b9586910161035e565b610241565b0190565b90916105746105829360408401908482035f860152610529565b916020818403910152610529565b90565b346105b65761059d6105983660046104e7565b610abc565b906105b26105a96100d2565b9283928361055a565b0390f35b6100d8565b5f80fd5b5f80fd5b909182601f830112156105fd5781359167ffffffffffffffff83116105f85760200192600183028401116105f357565b6105bf565b6105bb565b610458565b909160408284031261065c575f82013567ffffffffffffffff8111610657578361062d9184016105c3565b929093602082013567ffffffffffffffff81116106525761064e92016105c3565b9091565b610454565b610454565b6100dc565b346106935761067d610674366004610602565b92919091610c36565b6106856100d2565b8061068f816103e7565b0390f35b6100d8565b6106a1816101de565b036106a857565b5f80fd5b905035906106b982610698565b565b906020828203126106d4576106d1915f016106ac565b90565b6100dc565b34610707576106f16106ec3660046106bb565b610dee565b6106f96100d2565b80610703816103e7565b0390f35b6100d8565b90565b6107188161070c565b0361071f57565b5f80fd5b905035906107308261070f565b565b9060208282031261074b57610748915f01610723565b90565b6100dc565b6107599061070c565b90565b9061076690610750565b5f5260205260405f2090565b1c90565b60ff1690565b61078c9060086107919302610772565b610776565b90565b9061079f915461077c565b90565b6107b8906107b36065915f9261075c565b610794565b90565b151590565b6107c9906107bb565b9052565b91906107e0905f602085019401906107c0565b565b346108125761080e6107fd6107f8366004610732565b6107a2565b6108056100d2565b918291826107cd565b0390f35b6100d8565b7f000000000000000000000000000000000000000000000000000000000000000090565b6108449061014b565b90565b6108509061083b565b9052565b9190610867905f60208501940190610847565b565b34610899576108793660046100e0565b610895610884610817565b61088c6100d2565b91829182610854565b0390f35b6100d8565b9190916040818403126108df576108b7835f8301610723565b92602082013567ffffffffffffffff81116108da576108d692016105c3565b9091565b610454565b6100dc565b34610913576108fd6108f736600461089e565b91611147565b6109056100d2565b8061090f816103e7565b0390f35b6100d8565b5f80fd5b6109246114a9565b61092c610959565b565b90565b61094561094061094a9261092e565b61012c565b610113565b90565b61095690610931565b90565b61096a6109655f61094d565b61152d565b565b61097461091c565b565b5f90565b5f1c90565b73ffffffffffffffffffffffffffffffffffffffff1690565b6109a46109a99161097a565b61097f565b90565b6109b69054610998565b90565b6109c1610976565b506109cc60336109ac565b90565b606090565b5f80fd5b60e01b90565b909291926109f36109ee82610460565b6102a1565b93818552602085019082840111610a0f57610a0d9261035e565b565b61045c565b9080601f83011215610a3257816020610a2f935191016109de565b90565b610458565b919091604081840312610a8f575f81015167ffffffffffffffff8111610a8a5783610a63918301610a14565b92602082015167ffffffffffffffff8111610a8557610a829201610a14565b90565b610454565b610454565b6100dc565b610aa99160208201915f818403910152610529565b90565b610ab46100d2565b3d5f823e3d90fd5b905f610b2492610aca6109cf565b50610ad36109cf565b50610afd7f0000000000000000000000000000000000000000000000000000000000000000610157565b610b1963a903a277610b0d6100d2565b968794859384936109d8565b835260048301610a94565b03915afa8015610b64575f80939091610b3d575b509190565b9050610b5c9192503d805f833e610b548183610278565b810190610a37565b91905f610b38565b610aac565b5f910312610b7357565b6100dc565b9190610b9281610b8b81610b9795610520565b8095610483565b610241565b0190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60021115610bd257565b610b9b565b90610be182610bc8565b565b610bec90610bd7565b90565b610bf890610be3565b9052565b959492610c3494610c1e610c2c9360409560608b01918b83035f8d0152610b78565b9188830360208a0152610b78565b940190610bef565b565b929192610c627f000000000000000000000000000000000000000000000000000000000000000061083b565b906335ecb4c190929493600191833b15610ce457610ca1610c96935f97938894610c8a6100d2565b9a8b998a9889976109d8565b875260048701610bfc565b03925af18015610cdf57610cb3575b50565b610cd2905f3d8111610cd8575b610cca8183610278565b810190610b69565b5f610cb0565b503d610cc0565b610aac565b6109d4565b610cfa90610cf56114a9565b610dbe565b565b60207f6464726573730000000000000000000000000000000000000000000000000000917f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201520152565b610d566026604092610355565b610d5f81610cfc565b0190565b610d789060208101905f818303910152610d49565b90565b15610d8257565b610d8a6100d2565b7f08c379a000000000000000000000000000000000000000000000000000000000815280610dba60048201610d63565b0390fd5b610dec90610de781610de0610dda610dd55f61094d565b6101de565b916101de565b1415610d7b565b61152d565b565b610df790610ce9565b565b610e0491369161048e565b90565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b90610e3e8261051c565b811015610e5057600160209102010190565b610e07565b90565b90565b610e6f610e6a610e7492610e55565b61012c565b610e58565b90565b7fff000000000000000000000000000000000000000000000000000000000000001690565b610ea69051610e77565b90565b60f81c90565b60ff1690565b610ec9610ec4610ece92610eaf565b61012c565b610eaf565b90565b610edd610ee291610ea9565b610eb5565b90565b610ef9610ef4610efe9261092e565b61012c565b610eaf565b90565b90565b610f18610f13610f1d92610f01565b61012c565b610eaf565b90565b90565b610f37610f32610f3c92610f20565b61012c565b610eaf565b90565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b610f78610f7e91610eaf565b91610eaf565b019060ff8211610f8a57565b610f3f565b60f81b90565b610fa9610fa4610fae92610eaf565b610f8f565b610e77565b90565b5f7f496e76616c6964207369676e6174757265000000000000000000000000000000910152565b610fe56011602092610355565b610fee81610fb1565b0190565b6110079060208101905f818303910152610fd8565b90565b611013906101de565b90565b61101f8161100a565b0361102657565b5f80fd5b9050519061103782611016565b565b906020828203126110525761104f915f0161102a565b90565b6100dc565b6110609061014b565b90565b61106c816107bb565b0361107357565b5f80fd5b9050519061108482611063565b565b9060208282031261109f5761109c915f01611077565b90565b6100dc565b5f7f496e76616c6964207369676e6572000000000000000000000000000000000000910152565b6110d8600e602092610355565b6110e1816110a4565b0190565b6110fa9060208101905f8183039101526110cb565b90565b5f1b90565b9061110e60ff916110fd565b9181191691161790565b611121906107bb565b90565b90565b9061113c61113761114392611118565b611124565b8254611102565b9055565b91611155906111a092610df9565b61117961117461116f836111696040610e5b565b90610e34565b610e9c565b610ed1565b8061118c6111865f610ee5565b91610eaf565b1480156113f3575b6113b8575b508261158e565b806111bb6111b56111b05f61094d565b6101de565b916101de565b1461137c5761120460206111ee7f000000000000000000000000000000000000000000000000000000000000000061083b565b63d80a4c28906111fc6100d2565b9384926109d8565b82528180611214600482016103e7565b03915afa80156113775761123560209161125f935f9161134a575b50611057565b630123d0c19061125485926112486100d2565b958694859384936109d8565b8352600483016101f7565b03915afa80156113455761127b915f91611317575b50156107bb565b90816112db575b5061129f5761129d90611298600191606561075c565b611127565b565b6112a76100d2565b7f08c379a0000000000000000000000000000000000000000000000000000000008152806112d7600482016110e5565b0390fd5b905061130f6113097f00000000000000000000000000000000000000000000000000000000000000006101de565b916101de565b14155f611282565b611338915060203d811161133e575b6113308183610278565b810190611086565b5f611274565b503d611326565b610aac565b61136a9150833d8111611370575b6113628183610278565b810190611039565b5f61122f565b503d611358565b610aac565b6113846100d2565b7f08c379a0000000000000000000000000000000000000000000000000000000008152806113b460048201610ff2565b0390fd5b6113cf6113d4916113c9601b610f23565b90610f6c565b610f95565b6113ec826113e66040935f1a93610e5b565b90610e34565b535f611199565b50806114086114026001610f04565b91610eaf565b14611194565b5f7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572910152565b61144160208092610355565b61144a8161140e565b0190565b6114639060208101905f818303910152611435565b90565b1561146d57565b6114756100d2565b7f08c379a0000000000000000000000000000000000000000000000000000000008152806114a56004820161144e565b0390fd5b6114d36114b46109b9565b6114cd6114c76114c26115af565b6101de565b916101de565b14611466565b565b906114f473ffffffffffffffffffffffffffffffffffffffff916110fd565b9181191691161790565b6115079061014b565b90565b90565b9061152261151d611529926114fe565b61150a565b82546114d5565b9055565b61153760336109ac565b61154282603361150d565b906115766115707f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936114fe565b916114fe565b9161157f6100d2565b80611589816103e7565b0390a3565b6115ac916115a49161159e610976565b506115e7565b919091611836565b90565b6115b7610976565b503390565b5f90565b90565b6115d76115d26115dc926115c0565b61012c565b610e58565b90565b5f90565b5f90565b6115ef610976565b506115f86115bc565b506116028261051c565b61161561160f60416115c3565b91610e58565b145f1461165a57611654916116286115df565b506116316115df565b5061163a6115e3565b506020810151606060408301519201515f1a909192611a74565b91909190565b50506116655f61094d565b90600290565b6005111561167557565b610b9b565b906116848261166b565b565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202776272076616c5f8201520152565b6116e06022604092610355565b6116e981611686565b0190565b6117029060208101905f8183039101526116d3565b90565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201520152565b61175f6022604092610355565b61176881611705565b0190565b6117819060208101905f818303910152611752565b90565b5f7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800910152565b6117b8601f602092610355565b6117c181611784565b0190565b6117da9060208101905f8183039101526117ab565b90565b5f7f45434453413a20696e76616c6964207369676e61747572650000000000000000910152565b6118116018602092610355565b61181a816117dd565b0190565b6118339060208101905f818303910152611804565b90565b806118496118435f61167a565b9161167a565b145f146118535750565b80611867611861600161167a565b9161167a565b145f146118aa576118766100d2565b7f08c379a0000000000000000000000000000000000000000000000000000000008152806118a66004820161181e565b0390fd5b806118be6118b8600261167a565b9161167a565b145f14611901576118cd6100d2565b7f08c379a0000000000000000000000000000000000000000000000000000000008152806118fd600482016117c5565b0390fd5b8061191561190f600361167a565b9161167a565b145f14611958576119246100d2565b7f08c379a0000000000000000000000000000000000000000000000000000000008152806119546004820161176c565b0390fd5b61196b611965600461167a565b9161167a565b1461197257565b61197a6100d2565b7f08c379a0000000000000000000000000000000000000000000000000000000008152806119aa600482016116ed565b0390fd5b6119c26119bd6119c792610e58565b61012c565b610e58565b90565b6119d66119db9161097a565b6119ae565b90565b90565b6119f56119f06119fa926119de565b61012c565b610e58565b90565b90565b611a14611a0f611a19926119fd565b61012c565b610eaf565b90565b611a259061070c565b9052565b611a3290610eaf565b9052565b611a6b611a7294611a61606094989795611a57608086019a5f870190611a1c565b6020850190611a29565b6040830190611a1c565b0190611a1c565b565b929190611a7f610976565b50611a886115bc565b50611a92836119ca565b611ac4611abe7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a06119e1565b91610e58565b11611b855780611add611ad7601b610f23565b91610eaf565b141580611b69575b611b5657611b045f936020959293611afb6100d2565b94859485611a36565b838052039060015afa15611b5157611b1c5f516110fd565b80611b37611b31611b2c5f61094d565b6101de565b916101de565b14611b4157905f90565b50611b4b5f61094d565b90600190565b610aac565b50505050611b635f61094d565b90600490565b5080611b7e611b78601c611a00565b91610eaf565b1415611ae5565b50505050611b925f61094d565b9060039056fea164736f6c634300081c000a",
}

// BatchAuthenticatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchAuthenticatorMetaData.ABI instead.
var BatchAuthenticatorABI = BatchAuthenticatorMetaData.ABI

// BatchAuthenticatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchAuthenticatorMetaData.Bin instead.
var BatchAuthenticatorBin = BatchAuthenticatorMetaData.Bin

// DeployBatchAuthenticator deploys a new Ethereum contract, binding an instance of BatchAuthenticator to it.
func DeployBatchAuthenticator(auth *bind.TransactOpts, backend bind.ContractBackend, _espressoTEEVerifier common.Address, _preApprovedBatcher common.Address) (common.Address, *types.Transaction, *BatchAuthenticator, error) {
	parsed, err := BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchAuthenticatorBin), backend, _espressoTEEVerifier, _preApprovedBatcher)
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
