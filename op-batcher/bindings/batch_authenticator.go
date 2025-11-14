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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"},{\"name\":\"_preApprovedBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroValidator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preApprovedBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSignerWithoutAttestationVerification\",\"inputs\":[{\"name\":\"pcr0Hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"enclaveAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60e06040523461006e5761001a61001461017e565b91610271565b610022610073565b611c65610426823960805181818161031d015261145e015260a0518181816108b801528181610ac501528181610e92015261135d015260c05181818161026b0152610d780152611c6590f35b610079565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100a59061007d565b810190811060018060401b038211176100bd57604052565b610087565b906100d56100ce610073565b928361009b565b565b5f80fd5b60018060a01b031690565b6100ef906100db565b90565b6100fb906100e6565b90565b610107816100f2565b0361010e57565b5f80fd5b9050519061011f826100fe565b565b61012a816100e6565b0361013157565b5f80fd5b9050519061014282610121565b565b90916060828403126101795761017661015f845f8501610112565b9361016d8160208601610135565b93604001610135565b90565b6100d7565b61019c61208b80380380610191816100c2565b928339810190610144565b909192565b6101ab90516100f2565b90565b90565b6101c56101c06101ca926100db565b6101ae565b6100db565b90565b6101d6906101b1565b90565b6101e2906101cd565b90565b60e01b90565b6101f4906100e6565b90565b610200816101eb565b0361020757565b5f80fd5b90505190610218826101f7565b565b9060208282031261023357610230915f0161020b565b90565b6100d7565b5f0190565b610245610073565b3d5f823e3d90fd5b610256906101cd565b90565b610262906101b1565b90565b61026e90610259565b90565b906102af929161027f610323565b60a052608052602061029961029460a06101a1565b6101d9565b63d80a4c28906102a7610073565b9485926101e5565b825281806102bf60048201610238565b03915afa801561031e576102e16102e6916102ee945f916102f0575b5061024d565b610265565b60c0526103b5565b565b610311915060203d8111610317575b610309818361009b565b81019061021a565b5f6102db565b503d6102ff565b61023d565b61033361032e610418565b6103b5565b565b5f1c90565b60018060a01b031690565b61035161035691610335565b61033a565b90565b6103639054610345565b90565b5f1b90565b9061037c60018060a01b0391610366565b9181191691161790565b61038f906101cd565b90565b90565b906103aa6103a56103b192610386565b610392565b825461036b565b9055565b6103be5f610359565b6103c8825f610395565b906103fc6103f67f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610386565b91610386565b91610405610073565b8061040f81610238565b0390a3565b5f90565b610420610414565b50339056fe60806040526004361015610013575b6109b7565b61001d5f356100dc565b806302afd6e3146100d75780631b076a4c146100d25780631f568b18146100cd57806354fd4d50146100c8578063715018a6146100c35780638da5cb5b146100be578063a903a277146100b9578063ba58e82a146100b4578063f2fde38b146100af578063f81f2083146100aa578063fa14fe6d146100a55763fc619e410361000e57610983565b610908565b610881565b61079e565b610749565b6106b4565b610556565b610523565b6104ee565b610361565b6102e6565b610220565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b610100816100f4565b0361010757565b5f80fd5b90503590610118826100f7565b565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156101605781359167ffffffffffffffff831161015b57602001926001830284011161015657565b610122565b61011e565b61011a565b60018060a01b031690565b61017990610165565b90565b61018581610170565b0361018c57565b5f80fd5b9050359061019d8261017c565b565b9190608083820312610216576101b7815f850161010b565b92602081013567ffffffffffffffff811161021157826101d8918301610126565b929093604083013567ffffffffffffffff811161020c576101fe83610209928601610126565b939094606001610190565b90565b6100f0565b6100f0565b6100ec565b5f0190565b346102555761023f61023336600461019f565b94939093929192610ab6565b6102476100e2565b806102518161021b565b0390f35b6100e8565b5f91031261026457565b6100ec565b7f000000000000000000000000000000000000000000000000000000000000000090565b90565b6102a461029f6102a992610165565b61028d565b610165565b90565b6102b590610290565b90565b6102c1906102ac565b90565b6102cd906102b8565b9052565b91906102e4905f602085019401906102c4565b565b34610316576102f636600461025a565b610312610301610269565b6103096100e2565b918291826102d1565b0390f35b6100e8565b7f000000000000000000000000000000000000000000000000000000000000000090565b61034890610170565b9052565b919061035f905f6020850194019061033f565b565b346103915761037136600461025a565b61038d61037c61031b565b6103846100e2565b9182918261034c565b0390f35b6100e8565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906103be90610396565b810190811067ffffffffffffffff8211176103d857604052565b6103a0565b906103f06103e96100e2565b92836103b4565b565b67ffffffffffffffff81116104105761040c602091610396565b0190565b6103a0565b90610427610422836103f2565b6103dd565b918252565b5f7f312e302e30000000000000000000000000000000000000000000000000000000910152565b61045d6005610415565b9061046a6020830161042c565b565b610474610453565b90565b61047f61046c565b90565b61048a610477565b90565b5190565b60209181520190565b90825f9392825e0152565b6104c46104cd6020936104d2936104bb8161048d565b93848093610491565b9586910161049a565b610396565b0190565b6104eb9160208201915f8184039101526104a5565b90565b3461051e576104fe36600461025a565b61051a610509610482565b6105116100e2565b918291826104d6565b0390f35b6100e8565b346105515761053336600461025a565b61053b610c34565b6105436100e2565b8061054d8161021b565b0390f35b6100e8565b346105865761056636600461025a565b610582610571610c73565b6105796100e2565b9182918261034c565b0390f35b6100e8565b5f80fd5b67ffffffffffffffff81116105ad576105a9602091610396565b0190565b6103a0565b90825f939282370152565b909291926105d26105cd8261058f565b6103dd565b938185526020850190828401116105ee576105ec926105b2565b565b61058b565b9080601f830112156106115781602061060e933591016105bd565b90565b61011a565b90602082820312610646575f82013567ffffffffffffffff81116106415761063e92016105f3565b90565b6100f0565b6100ec565b5190565b60209181520190565b6106776106806020936106859361066e8161064b565b9384809361064f565b9586910161049a565b610396565b0190565b90916106a36106b19360408401908482035f860152610658565b916020818403910152610658565b90565b346106e5576106cc6106c7366004610616565b610d5b565b906106e16106d86100e2565b92839283610689565b0390f35b6100e8565b9091604082840312610744575f82013567ffffffffffffffff811161073f5783610715918401610126565b929093602082013567ffffffffffffffff811161073a576107369201610126565b9091565b6100f0565b6100f0565b6100ec565b3461077b5761076561075c3660046106ea565b92919091610e8a565b61076d6100e2565b806107778161021b565b0390f35b6100e8565b9060208282031261079957610796915f01610190565b90565b6100ec565b346107cc576107b66107b1366004610780565b611028565b6107be6100e2565b806107c88161021b565b0390f35b6100e8565b906020828203126107ea576107e7915f0161010b565b90565b6100ec565b6107f8906100f4565b90565b90610805906107ef565b5f5260205260405f2090565b1c90565b60ff1690565b61082b9060086108309302610811565b610815565b90565b9061083e915461081b565b90565b610857906108526001915f926107fb565b610833565b90565b151590565b6108689061085a565b9052565b919061087f905f6020850194019061085f565b565b346108b1576108ad61089c6108973660046107d1565b610841565b6108a46100e2565b9182918261086c565b0390f35b6100e8565b7f000000000000000000000000000000000000000000000000000000000000000090565b6108e3906102ac565b90565b6108ef906108da565b9052565b9190610906905f602085019401906108e6565b565b346109385761091836600461025a565b6109346109236108b6565b61092b6100e2565b918291826108f3565b0390f35b6100e8565b91909160408184031261097e57610956835f830161010b565b92602082013567ffffffffffffffff8111610979576109759201610126565b9091565b6100f0565b6100ec565b346109b25761099c61099636600461093d565b916112da565b6109a46100e2565b806109ae8161021b565b0390f35b6100e8565b5f80fd5b5f80fd5b60e01b90565b6109ce90610170565b90565b6109da816109c5565b036109e157565b5f80fd5b905051906109f2826109d1565b565b90602082820312610a0d57610a0a915f016109e5565b90565b6100ec565b610a1a6100e2565b3d5f823e3d90fd5b610a2b906102ac565b90565b5f910312610a3857565b6100ec565b610a46906100f4565b9052565b9190610a6481610a5d81610a699561064f565b80956105b2565b610396565b0190565b9695939094610a9e88606095610aac95610a91610ab49a5f60808601950190610a3d565b8b830360208d0152610a4a565b9188830360408a0152610a4a565b94019061033f565b565b9194909293610aff6020610ae97f00000000000000000000000000000000000000000000000000000000000000006108da565b63d80a4c2890610af76100e2565b9384926109bf565b82528180610b0f6004820161021b565b03915afa8015610bdf57610b2a915f91610bb1575b50610a22565b926302afd6e390949695919295843b15610bac575f96610b5e948894610b6993610b526100e2565b9b8c9a8b998a986109bf565b885260048801610a6d565b03925af18015610ba757610b7b575b50565b610b9a905f3d8111610ba0575b610b9281836103b4565b810190610a2e565b5f610b78565b503d610b88565b610a12565b6109bb565b610bd2915060203d8111610bd8575b610bca81836103b4565b8101906109f4565b5f610b24565b503d610bc0565b610a12565b610bec6115ee565b610bf4610c21565b565b90565b610c0d610c08610c1292610bf6565b61028d565b610165565b90565b610c1e90610bf9565b90565b610c32610c2d5f610c15565b611664565b565b610c3c610be4565b565b5f90565b5f1c90565b60018060a01b031690565b610c5e610c6391610c42565b610c47565b90565b610c709054610c52565b90565b610c7b610c3e565b50610c855f610c66565b90565b606090565b90929192610ca2610c9d8261058f565b6103dd565b93818552602085019082840111610cbe57610cbc9261049a565b565b61058b565b9080601f83011215610ce157816020610cde93519101610c8d565b90565b61011a565b919091604081840312610d3e575f81015167ffffffffffffffff8111610d395783610d12918301610cc3565b92602082015167ffffffffffffffff8111610d3457610d319201610cc3565b90565b6100f0565b6100f0565b6100ec565b610d589160208201915f818403910152610658565b90565b905f610dc392610d69610c88565b50610d72610c88565b50610d9c7f00000000000000000000000000000000000000000000000000000000000000006102b8565b610db863a903a277610dac6100e2565b968794859384936109bf565b835260048301610d43565b03915afa8015610e03575f80939091610ddc575b509190565b9050610dfb9192503d805f833e610df381836103b4565b810190610ce6565b91905f610dd7565b610a12565b634e487b7160e01b5f52602160045260245ffd5b60021115610e2657565b610e08565b90610e3582610e1c565b565b610e4090610e2b565b90565b610e4c90610e37565b9052565b959492610e8894610e72610e809360409560608b01918b83035f8d0152610a4a565b9188830360208a0152610a4a565b940190610e43565b565b929192610eb67f00000000000000000000000000000000000000000000000000000000000000006108da565b906335ecb4c190929493600191833b15610f3857610ef5610eea935f97938894610ede6100e2565b9a8b998a9889976109bf565b875260048701610e50565b03925af18015610f3357610f07575b50565b610f26905f3d8111610f2c575b610f1e81836103b4565b810190610a2e565b5f610f04565b503d610f14565b610a12565b6109bb565b610f4e90610f496115ee565b610ff8565b565b60207f6464726573730000000000000000000000000000000000000000000000000000917f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201520152565b610faa6026604092610491565b610fb381610f50565b0190565b610fcc9060208101905f818303910152610f9d565b90565b15610fd657565b610fde6100e2565b62461bcd60e51b815280610ff460048201610fb7565b0390fd5b611026906110218161101a61101461100f5f610c15565b610170565b91610170565b1415610fcf565b611664565b565b61103190610f3d565b565b61103e9136916105bd565b90565b634e487b7160e01b5f52603260045260245ffd5b9061105f8261064b565b81101561107157600160209102010190565b611041565b90565b90565b61109061108b61109592611076565b61028d565b611079565b90565b60ff60f81b1690565b6110ab9051611098565b90565b60f81c90565b60ff1690565b6110ce6110c96110d3926110b4565b61028d565b6110b4565b90565b6110e26110e7916110ae565b6110ba565b90565b6110fe6110f961110392610bf6565b61028d565b6110b4565b90565b90565b61111d61111861112292611106565b61028d565b6110b4565b90565b90565b61113c61113761114192611125565b61028d565b6110b4565b90565b634e487b7160e01b5f52601160045260245ffd5b61116461116a916110b4565b916110b4565b019060ff821161117657565b611144565b60f81b90565b61119561119061119a926110b4565b61117b565b611098565b90565b5f7f496e76616c6964207369676e6174757265000000000000000000000000000000910152565b6111d16011602092610491565b6111da8161119d565b0190565b6111f39060208101905f8183039101526111c4565b90565b6111ff8161085a565b0361120657565b5f80fd5b90505190611217826111f6565b565b906020828203126112325761122f915f0161120a565b90565b6100ec565b5f7f496e76616c6964207369676e6572000000000000000000000000000000000000910152565b61126b600e602092610491565b61127481611237565b0190565b61128d9060208101905f81830391015261125e565b90565b5f1b90565b906112a160ff91611290565b9181191691161790565b6112b49061085a565b90565b90565b906112cf6112ca6112d6926112ab565b6112b7565b8254611295565b9055565b916112e89061133392611033565b61130c611307611302836112fc604061107c565b90611055565b6110a1565b6110d6565b8061131f6113195f6110ea565b916110b4565b148015611552575b611517575b50826116c3565b8061134e6113486113435f610c15565b610170565b91610170565b146114f55761139760206113817f00000000000000000000000000000000000000000000000000000000000000006108da565b63d80a4c289061138f6100e2565b9384926109bf565b825281806113a76004820161021b565b03915afa80156114f0576113c86020916113f2935f916114c3575b50610a22565b630123d0c1906113e785926113db6100e2565b958694859384936109bf565b83526004830161034c565b03915afa80156114be5761140e915f91611490575b501561085a565b9081611454575b50611432576114309061142b60019160016107fb565b6112ba565b565b61143a6100e2565b62461bcd60e51b81528061145060048201611278565b0390fd5b90506114886114827f0000000000000000000000000000000000000000000000000000000000000000610170565b91610170565b14155f611415565b6114b1915060203d81116114b7575b6114a981836103b4565b810190611219565b5f611407565b503d61149f565b610a12565b6114e39150833d81116114e9575b6114db81836103b4565b8101906109f4565b5f6113c2565b503d6114d1565b610a12565b6114fd6100e2565b62461bcd60e51b815280611513600482016111de565b0390fd5b61152e61153391611528601b611128565b90611158565b611181565b61154b826115456040935f1a9361107c565b90611055565b535f61132c565b50806115676115616001611109565b916110b4565b14611327565b5f7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572910152565b6115a060208092610491565b6115a98161156d565b0190565b6115c29060208101905f818303910152611594565b90565b156115cc57565b6115d46100e2565b62461bcd60e51b8152806115ea600482016115ad565b0390fd5b6116186115f9610c73565b61161261160c6116076116e4565b610170565b91610170565b146115c5565b565b9061162b60018060a01b0391611290565b9181191691161790565b61163e906102ac565b90565b90565b9061165961165461166092611635565b611641565b825461161a565b9055565b61166d5f610c66565b611677825f611644565b906116ab6116a57f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611635565b91611635565b916116b46100e2565b806116be8161021b565b0390a3565b6116e1916116d9916116d3610c3e565b5061171c565b91909161196b565b90565b6116ec610c3e565b503390565b5f90565b90565b61170c611707611711926116f5565b61028d565b611079565b90565b5f90565b5f90565b611724610c3e565b5061172d6116f1565b506117378261064b565b61174a61174460416116f8565b91611079565b145f1461178f576117899161175d611714565b50611766611714565b5061176f611718565b506020810151606060408301519201515f1a909192611b34565b91909190565b505061179a5f610c15565b90600290565b600511156117aa57565b610e08565b906117b9826117a0565b565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202776272076616c5f8201520152565b6118156022604092610491565b61181e816117bb565b0190565b6118379060208101905f818303910152611808565b90565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201520152565b6118946022604092610491565b61189d8161183a565b0190565b6118b69060208101905f818303910152611887565b90565b5f7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800910152565b6118ed601f602092610491565b6118f6816118b9565b0190565b61190f9060208101905f8183039101526118e0565b90565b5f7f45434453413a20696e76616c6964207369676e61747572650000000000000000910152565b6119466018602092610491565b61194f81611912565b0190565b6119689060208101905f818303910152611939565b90565b8061197e6119785f6117af565b916117af565b145f146119885750565b8061199c61199660016117af565b916117af565b145f146119c5576119ab6100e2565b62461bcd60e51b8152806119c160048201611953565b0390fd5b806119d96119d360026117af565b916117af565b145f14611a02576119e86100e2565b62461bcd60e51b8152806119fe600482016118fa565b0390fd5b80611a16611a1060036117af565b916117af565b145f14611a3f57611a256100e2565b62461bcd60e51b815280611a3b600482016118a1565b0390fd5b611a52611a4c60046117af565b916117af565b14611a5957565b611a616100e2565b62461bcd60e51b815280611a7760048201611822565b0390fd5b611a8f611a8a611a9492611079565b61028d565b611079565b90565b611aa3611aa891610c42565b611a7b565b90565b90565b611ac2611abd611ac792611aab565b61028d565b611079565b90565b90565b611ae1611adc611ae692611aca565b61028d565b6110b4565b90565b611af2906110b4565b9052565b611b2b611b3294611b21606094989795611b17608086019a5f870190610a3d565b6020850190611ae9565b6040830190610a3d565b0190610a3d565b565b929190611b3f610c3e565b50611b486116f1565b50611b5283611a97565b611b84611b7e7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0611aae565b91611079565b11611c455780611b9d611b97601b611128565b916110b4565b141580611c29575b611c1657611bc45f936020959293611bbb6100e2565b94859485611af6565b838052039060015afa15611c1157611bdc5f51611290565b80611bf7611bf1611bec5f610c15565b610170565b91610170565b14611c0157905f90565b50611c0b5f610c15565b90600190565b610a12565b50505050611c235f610c15565b90600490565b5080611c3e611c38601c611acd565b916110b4565b1415611ba5565b50505050611c525f610c15565b9060039056fea164736f6c634300081d000a",
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

// RegisterSignerWithoutAttestationVerification is a paid mutator transaction binding the contract method 0x02afd6e3.
//
// Solidity: function registerSignerWithoutAttestationVerification(bytes32 pcr0Hash, bytes attestationTbs, bytes signature, address enclaveAddress) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactor) RegisterSignerWithoutAttestationVerification(opts *bind.TransactOpts, pcr0Hash [32]byte, attestationTbs []byte, signature []byte, enclaveAddress common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.contract.Transact(opts, "registerSignerWithoutAttestationVerification", pcr0Hash, attestationTbs, signature, enclaveAddress)
}

// RegisterSignerWithoutAttestationVerification is a paid mutator transaction binding the contract method 0x02afd6e3.
//
// Solidity: function registerSignerWithoutAttestationVerification(bytes32 pcr0Hash, bytes attestationTbs, bytes signature, address enclaveAddress) returns()
func (_BatchAuthenticator *BatchAuthenticatorSession) RegisterSignerWithoutAttestationVerification(pcr0Hash [32]byte, attestationTbs []byte, signature []byte, enclaveAddress common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSignerWithoutAttestationVerification(&_BatchAuthenticator.TransactOpts, pcr0Hash, attestationTbs, signature, enclaveAddress)
}

// RegisterSignerWithoutAttestationVerification is a paid mutator transaction binding the contract method 0x02afd6e3.
//
// Solidity: function registerSignerWithoutAttestationVerification(bytes32 pcr0Hash, bytes attestationTbs, bytes signature, address enclaveAddress) returns()
func (_BatchAuthenticator *BatchAuthenticatorTransactorSession) RegisterSignerWithoutAttestationVerification(pcr0Hash [32]byte, attestationTbs []byte, signature []byte, enclaveAddress common.Address) (*types.Transaction, error) {
	return _BatchAuthenticator.Contract.RegisterSignerWithoutAttestationVerification(&_BatchAuthenticator.TransactOpts, pcr0Hash, attestationTbs, signature, enclaveAddress)
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
