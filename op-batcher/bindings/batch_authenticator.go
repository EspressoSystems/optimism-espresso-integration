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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_espressoTEEVerifier\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"},{\"name\":\"_teeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonTeeBatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIsTee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authenticateBatchInfo\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAttestationTbs\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"espressoTEEVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEspressoTEEVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nitroValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroValidator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonTeeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerSigner\",\"inputs\":[{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSignerWithoutAttestationVerification\",\"inputs\":[{\"name\":\"pcr0Hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"attestationTbs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"enclaveAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchBatcher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teeBatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validBatchInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523461007b5761001e610015610197565b9291909161045e565b610026610080565b611dc1610668823960805181818161088001526115ba015260a0518161075e015260c0518181816109b801528181610bc501528181610f9201526114b9015260e05181818161029b0152610e780152611dc190f35b610086565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100b29061008a565b810190811060018060401b038211176100ca57604052565b610094565b906100e26100db610080565b92836100a8565b565b5f80fd5b60018060a01b031690565b6100fc906100e8565b90565b610108906100f3565b90565b610114816100ff565b0361011b57565b5f80fd5b9050519061012c8261010b565b565b610137816100f3565b0361013e57565b5f80fd5b9050519061014f8261012e565b565b60808183031261019257610167825f830161011f565b9261018f6101788460208501610142565b936101868160408601610142565b93606001610142565b90565b6100e4565b6101b5612429803803806101aa816100cf565b928339810190610151565b90919293565b90565b90565b6101d56101d06101da926101bb565b6101be565b6100e8565b90565b6101e6906101c1565b90565b60209181520190565b60207f6368657200000000000000000000000000000000000000000000000000000000917f426174636841757468656e74696361746f723a207a65726f20746565206261745f8201520152565b61024c60246040926101e9565b610255816101f2565b0190565b61026e9060208101905f81830391015261023f565b90565b1561027857565b610280610080565b62461bcd60e51b81528061029660048201610259565b0390fd5b60207f2062617463686572000000000000000000000000000000000000000000000000917f426174636841757468656e74696361746f723a207a65726f206e6f6e2d7465655f8201520152565b6102f460286040926101e9565b6102fd8161029a565b0190565b6103169060208101905f8183039101526102e7565b90565b1561032057565b610328610080565b62461bcd60e51b81528061033e60048201610301565b0390fd5b61034c90516100ff565b90565b61036361035e610368926100e8565b6101be565b6100e8565b90565b6103749061034f565b90565b6103809061036b565b90565b60e01b90565b610392906100f3565b90565b61039e81610389565b036103a557565b5f80fd5b905051906103b682610395565b565b906020828203126103d1576103ce915f016103a9565b90565b6100e4565b5f0190565b6103e3610080565b3d5f823e3d90fd5b6103f49061036b565b90565b6104009061034f565b90565b61040c906103f7565b90565b5f1b90565b9061042060ff9161040f565b9181191691161790565b151590565b6104389061042a565b90565b90565b9061045361044e61045a9261042f565b61043b565b8254610414565b9055565b906104ea93929161046d61056a565b6104928261048b6104856104805f6101dd565b6100f3565b916100f3565b1415610271565b6104b7836104b06104aa6104a55f6101dd565b6100f3565b916100f3565b1415610319565b60c05260805260a05260206104d46104cf60c0610342565b610377565b63d80a4c28906104e2610080565b948592610383565b825281806104fa600482016103d6565b03915afa80156105655761051c61052191610535945f91610537575b506103eb565b610403565b60e0526105306001600261043e565b6105f7565b565b610558915060203d811161055e575b61055081836100a8565b8101906103b8565b5f610516565b503d610546565b6103db565b61057a61057561065a565b6105f7565b565b5f1c90565b60018060a01b031690565b61059861059d9161057c565b610581565b90565b6105aa905461058c565b90565b906105be60018060a01b039161040f565b9181191691161790565b6105d19061036b565b90565b90565b906105ec6105e76105f3926105c8565b6105d4565b82546105ad565b9055565b6106005f6105a0565b61060a825f6105d7565b9061063e6106387f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936105c8565b916105c8565b91610647610080565b80610651816103d6565b0390a3565b5f90565b610662610656565b50339056fe60806040526004361015610013575b610ab7565b61001d5f3561010c565b806302afd6e3146101075780631b076a4c1461010257806354fd4d50146100fd578063715018a6146100f85780637877a9ed146100f35780638da5cb5b146100ee578063a903a277146100e9578063b1bd4285146100e4578063ba58e82a146100df578063bc347f47146100da578063d909ba7c146100d5578063f2fde38b146100d0578063f81f2083146100cb578063fa14fe6d146100c65763fc619e410361000e57610a83565b610a08565b610981565b6108f5565b6108a2565b61084b565b610814565b610780565b610726565b6105c8565b610571565b6104d8565b6104a3565b610316565b610250565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b61013081610124565b0361013757565b5f80fd5b9050359061014882610127565b565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156101905781359167ffffffffffffffff831161018b57602001926001830284011161018657565b610152565b61014e565b61014a565b60018060a01b031690565b6101a990610195565b90565b6101b5816101a0565b036101bc57565b5f80fd5b905035906101cd826101ac565b565b9190608083820312610246576101e7815f850161013b565b92602081013567ffffffffffffffff81116102415782610208918301610156565b929093604083013567ffffffffffffffff811161023c5761022e83610239928601610156565b9390946060016101c0565b90565b610120565b610120565b61011c565b5f0190565b346102855761026f6102633660046101cf565b94939093929192610bb6565b610277610112565b806102818161024b565b0390f35b610118565b5f91031261029457565b61011c565b7f000000000000000000000000000000000000000000000000000000000000000090565b90565b6102d46102cf6102d992610195565b6102bd565b610195565b90565b6102e5906102c0565b90565b6102f1906102dc565b90565b6102fd906102e8565b9052565b9190610314905f602085019401906102f4565b565b346103465761032636600461028a565b610342610331610299565b610339610112565b91829182610301565b0390f35b610118565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906103739061034b565b810190811067ffffffffffffffff82111761038d57604052565b610355565b906103a561039e610112565b9283610369565b565b67ffffffffffffffff81116103c5576103c160209161034b565b0190565b610355565b906103dc6103d7836103a7565b610392565b918252565b5f7f312e302e30000000000000000000000000000000000000000000000000000000910152565b61041260056103ca565b9061041f602083016103e1565b565b610429610408565b90565b610434610421565b90565b61043f61042c565b90565b5190565b60209181520190565b90825f9392825e0152565b6104796104826020936104879361047081610442565b93848093610446565b9586910161044f565b61034b565b0190565b6104a09160208201915f81840391015261045a565b90565b346104d3576104b336600461028a565b6104cf6104be610437565b6104c6610112565b9182918261048b565b0390f35b610118565b34610506576104e836600461028a565b6104f0610d34565b6104f8610112565b806105028161024b565b0390f35b610118565b1c90565b60ff1690565b61052590600861052a930261050b565b61050f565b90565b906105389154610515565b90565b61054760025f9061052d565b90565b151590565b6105589061054a565b9052565b919061056f905f6020850194019061054f565b565b346105a15761058136600461028a565b61059d61058c61053b565b610594610112565b9182918261055c565b0390f35b610118565b6105af906101a0565b9052565b91906105c6905f602085019401906105a6565b565b346105f8576105d836600461028a565b6105f46105e3610d73565b6105eb610112565b918291826105b3565b0390f35b610118565b5f80fd5b67ffffffffffffffff811161061f5761061b60209161034b565b0190565b610355565b90825f939282370152565b9092919261064461063f82610601565b610392565b938185526020850190828401116106605761065e92610624565b565b6105fd565b9080601f83011215610683578160206106809335910161062f565b90565b61014a565b906020828203126106b8575f82013567ffffffffffffffff81116106b3576106b09201610665565b90565b610120565b61011c565b5190565b60209181520190565b6106e96106f26020936106f7936106e0816106bd565b938480936106c1565b9586910161044f565b61034b565b0190565b90916107156107239360408401908482035f8601526106ca565b9160208184039101526106ca565b90565b346107575761073e610739366004610688565b610e5b565b9061075361074a610112565b928392836106fb565b0390f35b610118565b7f000000000000000000000000000000000000000000000000000000000000000090565b346107b05761079036600461028a565b6107ac61079b61075c565b6107a3610112565b918291826105b3565b0390f35b610118565b909160408284031261080f575f82013567ffffffffffffffff811161080a57836107e0918401610156565b929093602082013567ffffffffffffffff8111610805576108019201610156565b9091565b610120565b610120565b61011c565b34610846576108306108273660046107b5565b92919091610f8a565b610838610112565b806108428161024b565b0390f35b610118565b346108795761085b36600461028a565b6108636110d9565b61086b610112565b806108758161024b565b0390f35b610118565b7f000000000000000000000000000000000000000000000000000000000000000090565b346108d2576108b236600461028a565b6108ce6108bd61087e565b6108c5610112565b918291826105b3565b0390f35b610118565b906020828203126108f0576108ed915f016101c0565b90565b61011c565b346109235761090d6109083660046108d7565b6111ce565b610915610112565b8061091f8161024b565b0390f35b610118565b906020828203126109415761093e915f0161013b565b90565b61011c565b61094f90610124565b90565b9061095c90610946565b5f5260205260405f2090565b61097e906109796001915f92610952565b61052d565b90565b346109b1576109ad61099c610997366004610928565b610968565b6109a4610112565b9182918261055c565b0390f35b610118565b7f000000000000000000000000000000000000000000000000000000000000000090565b6109e3906102dc565b90565b6109ef906109da565b9052565b9190610a06905f602085019401906109e6565b565b34610a3857610a1836600461028a565b610a34610a236109b6565b610a2b610112565b918291826109f3565b0390f35b610118565b919091604081840312610a7e57610a56835f830161013b565b92602082013567ffffffffffffffff8111610a7957610a759201610156565b9091565b610120565b61011c565b34610ab257610a9c610a96366004610a3d565b91611436565b610aa4610112565b80610aae8161024b565b0390f35b610118565b5f80fd5b5f80fd5b60e01b90565b610ace906101a0565b90565b610ada81610ac5565b03610ae157565b5f80fd5b90505190610af282610ad1565b565b90602082820312610b0d57610b0a915f01610ae5565b90565b61011c565b610b1a610112565b3d5f823e3d90fd5b610b2b906102dc565b90565b5f910312610b3857565b61011c565b610b4690610124565b9052565b9190610b6481610b5d81610b69956106c1565b8095610624565b61034b565b0190565b9695939094610b9e88606095610bac95610b91610bb49a5f60808601950190610b3d565b8b830360208d0152610b4a565b9188830360408a0152610b4a565b9401906105a6565b565b9194909293610bff6020610be97f00000000000000000000000000000000000000000000000000000000000000006109da565b63d80a4c2890610bf7610112565b938492610abf565b82528180610c0f6004820161024b565b03915afa8015610cdf57610c2a915f91610cb1575b50610b22565b926302afd6e390949695919295843b15610cac575f96610c5e948894610c6993610c52610112565b9b8c9a8b998a98610abf565b885260048801610b6d565b03925af18015610ca757610c7b575b50565b610c9a905f3d8111610ca0575b610c928183610369565b810190610b2e565b5f610c78565b503d610c88565b610b12565b610abb565b610cd2915060203d8111610cd8575b610cca8183610369565b810190610af4565b5f610c24565b503d610cc0565b610b12565b610cec61174a565b610cf4610d21565b565b90565b610d0d610d08610d1292610cf6565b6102bd565b610195565b90565b610d1e90610cf9565b90565b610d32610d2d5f610d15565b6117c0565b565b610d3c610ce4565b565b5f90565b5f1c90565b60018060a01b031690565b610d5e610d6391610d42565b610d47565b90565b610d709054610d52565b90565b610d7b610d3e565b50610d855f610d66565b90565b606090565b90929192610da2610d9d82610601565b610392565b93818552602085019082840111610dbe57610dbc9261044f565b565b6105fd565b9080601f83011215610de157816020610dde93519101610d8d565b90565b61014a565b919091604081840312610e3e575f81015167ffffffffffffffff8111610e395783610e12918301610dc3565b92602082015167ffffffffffffffff8111610e3457610e319201610dc3565b90565b610120565b610120565b61011c565b610e589160208201915f8184039101526106ca565b90565b905f610ec392610e69610d88565b50610e72610d88565b50610e9c7f00000000000000000000000000000000000000000000000000000000000000006102e8565b610eb863a903a277610eac610112565b96879485938493610abf565b835260048301610e43565b03915afa8015610f03575f80939091610edc575b509190565b9050610efb9192503d805f833e610ef38183610369565b810190610de6565b91905f610ed7565b610b12565b634e487b7160e01b5f52602160045260245ffd5b60021115610f2657565b610f08565b90610f3582610f1c565b565b610f4090610f2b565b90565b610f4c90610f37565b9052565b959492610f8894610f72610f809360409560608b01918b83035f8d0152610b4a565b9188830360208a0152610b4a565b940190610f43565b565b929192610fb67f00000000000000000000000000000000000000000000000000000000000000006109da565b906335ecb4c190929493600191833b1561103857610ff5610fea935f97938894610fde610112565b9a8b998a988997610abf565b875260048701610f50565b03925af1801561103357611007575b50565b611026905f3d811161102c575b61101e8183610369565b810190610b2e565b5f611004565b503d611014565b610b12565b610abb565b61104561174a565b61104d6110ba565b565b61105b61106091610d42565b61050f565b90565b61106d905461104f565b90565b5f1b90565b9061108160ff91611070565b9181191691161790565b6110949061054a565b90565b90565b906110af6110aa6110b69261108b565b611097565b8254611075565b9055565b6110d76110d06110ca6002611063565b1561054a565b600261109a565b565b6110e161103d565b565b6110f4906110ef61174a565b61119e565b565b60207f6464726573730000000000000000000000000000000000000000000000000000917f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201520152565b6111506026604092610446565b611159816110f6565b0190565b6111729060208101905f818303910152611143565b90565b1561117c57565b611184610112565b62461bcd60e51b81528061119a6004820161115d565b0390fd5b6111cc906111c7816111c06111ba6111b55f610d15565b6101a0565b916101a0565b1415611175565b6117c0565b565b6111d7906110e3565b565b6111e491369161062f565b90565b634e487b7160e01b5f52603260045260245ffd5b90611205826106bd565b81101561121757600160209102010190565b6111e7565b90565b90565b61123661123161123b9261121c565b6102bd565b61121f565b90565b60ff60f81b1690565b611251905161123e565b90565b60f81c90565b60ff1690565b61127461126f6112799261125a565b6102bd565b61125a565b90565b61128861128d91611254565b611260565b90565b6112a461129f6112a992610cf6565b6102bd565b61125a565b90565b90565b6112c36112be6112c8926112ac565b6102bd565b61125a565b90565b90565b6112e26112dd6112e7926112cb565b6102bd565b61125a565b90565b634e487b7160e01b5f52601160045260245ffd5b61130a6113109161125a565b9161125a565b019060ff821161131c57565b6112ea565b60f81b90565b61133b6113366113409261125a565b611321565b61123e565b90565b5f7f496e76616c6964207369676e6174757265000000000000000000000000000000910152565b6113776011602092610446565b61138081611343565b0190565b6113999060208101905f81830391015261136a565b90565b6113a58161054a565b036113ac57565b5f80fd5b905051906113bd8261139c565b565b906020828203126113d8576113d5915f016113b0565b90565b61011c565b5f7f496e76616c6964207369676e6572000000000000000000000000000000000000910152565b611411600e602092610446565b61141a816113dd565b0190565b6114339060208101905f818303910152611404565b90565b916114449061148f926111d9565b61146861146361145e836114586040611222565b906111fb565b611247565b61127c565b8061147b6114755f611290565b9161125a565b1480156116ae575b611673575b508261181f565b806114aa6114a461149f5f610d15565b6101a0565b916101a0565b14611651576114f360206114dd7f00000000000000000000000000000000000000000000000000000000000000006109da565b63d80a4c28906114eb610112565b938492610abf565b825281806115036004820161024b565b03915afa801561164c5761152460209161154e935f9161161f575b50610b22565b630123d0c1906115438592611537610112565b95869485938493610abf565b8352600483016105b3565b03915afa801561161a5761156a915f916115ec575b501561054a565b90816115b0575b5061158e5761158c906115876001916001610952565b61109a565b565b611596610112565b62461bcd60e51b8152806115ac6004820161141e565b0390fd5b90506115e46115de7f00000000000000000000000000000000000000000000000000000000000000006101a0565b916101a0565b14155f611571565b61160d915060203d8111611613575b6116058183610369565b8101906113bf565b5f611563565b503d6115fb565b610b12565b61163f9150833d8111611645575b6116378183610369565b810190610af4565b5f61151e565b503d61162d565b610b12565b611659610112565b62461bcd60e51b81528061166f60048201611384565b0390fd5b61168a61168f91611684601b6112ce565b906112fe565b611327565b6116a7826116a16040935f1a93611222565b906111fb565b535f611488565b50806116c36116bd60016112af565b9161125a565b14611483565b5f7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572910152565b6116fc60208092610446565b611705816116c9565b0190565b61171e9060208101905f8183039101526116f0565b90565b1561172857565b611730610112565b62461bcd60e51b81528061174660048201611709565b0390fd5b611774611755610d73565b61176e611768611763611840565b6101a0565b916101a0565b14611721565b565b9061178760018060a01b0391611070565b9181191691161790565b61179a906102dc565b90565b90565b906117b56117b06117bc92611791565b61179d565b8254611776565b9055565b6117c95f610d66565b6117d3825f6117a0565b906118076118017f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611791565b91611791565b91611810610112565b8061181a8161024b565b0390a3565b61183d916118359161182f610d3e565b50611878565b919091611ac7565b90565b611848610d3e565b503390565b5f90565b90565b61186861186361186d92611851565b6102bd565b61121f565b90565b5f90565b5f90565b611880610d3e565b5061188961184d565b50611893826106bd565b6118a66118a06041611854565b9161121f565b145f146118eb576118e5916118b9611870565b506118c2611870565b506118cb611874565b506020810151606060408301519201515f1a909192611c90565b91909190565b50506118f65f610d15565b90600290565b6005111561190657565b610f08565b90611915826118fc565b565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202776272076616c5f8201520152565b6119716022604092610446565b61197a81611917565b0190565b6119939060208101905f818303910152611964565b90565b60207f7565000000000000000000000000000000000000000000000000000000000000917f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201520152565b6119f06022604092610446565b6119f981611996565b0190565b611a129060208101905f8183039101526119e3565b90565b5f7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800910152565b611a49601f602092610446565b611a5281611a15565b0190565b611a6b9060208101905f818303910152611a3c565b90565b5f7f45434453413a20696e76616c6964207369676e61747572650000000000000000910152565b611aa26018602092610446565b611aab81611a6e565b0190565b611ac49060208101905f818303910152611a95565b90565b80611ada611ad45f61190b565b9161190b565b145f14611ae45750565b80611af8611af2600161190b565b9161190b565b145f14611b2157611b07610112565b62461bcd60e51b815280611b1d60048201611aaf565b0390fd5b80611b35611b2f600261190b565b9161190b565b145f14611b5e57611b44610112565b62461bcd60e51b815280611b5a60048201611a56565b0390fd5b80611b72611b6c600361190b565b9161190b565b145f14611b9b57611b81610112565b62461bcd60e51b815280611b97600482016119fd565b0390fd5b611bae611ba8600461190b565b9161190b565b14611bb557565b611bbd610112565b62461bcd60e51b815280611bd36004820161197e565b0390fd5b611beb611be6611bf09261121f565b6102bd565b61121f565b90565b611bff611c0491610d42565b611bd7565b90565b90565b611c1e611c19611c2392611c07565b6102bd565b61121f565b90565b90565b611c3d611c38611c4292611c26565b6102bd565b61125a565b90565b611c4e9061125a565b9052565b611c87611c8e94611c7d606094989795611c73608086019a5f870190610b3d565b6020850190611c45565b6040830190610b3d565b0190610b3d565b565b929190611c9b610d3e565b50611ca461184d565b50611cae83611bf3565b611ce0611cda7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0611c0a565b9161121f565b11611da15780611cf9611cf3601b6112ce565b9161125a565b141580611d85575b611d7257611d205f936020959293611d17610112565b94859485611c52565b838052039060015afa15611d6d57611d385f51611070565b80611d53611d4d611d485f610d15565b6101a0565b916101a0565b14611d5d57905f90565b50611d675f610d15565b90600190565b610b12565b50505050611d7f5f610d15565b90600490565b5080611d9a611d94601c611c29565b9161125a565b1415611d01565b50505050611dae5f610d15565b9060039056fea164736f6c634300081c000a",
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
