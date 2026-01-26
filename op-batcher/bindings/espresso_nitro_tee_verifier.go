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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GUARDIAN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_nitroEnclaveVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINitroEnclaveVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteEnclaveHashes\",\"inputs\":[{\"name\":\"enclaveHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"service\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enclaveHashSigners\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"service\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGuardians\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"guardianCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"nitroEnclaveVerifier\",\"type\":\"address\",\"internalType\":\"contractINitroEnclaveVerifier\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isGuardian\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerService\",\"inputs\":[{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"service\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"service\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registeredEnclaveHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"},{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registeredService\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"service\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registeredServices\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"service\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNitroEnclaveVerifier\",\"inputs\":[{\"name\":\"nitroEnclaveVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DeletedEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"service\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumServiceType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeletedRegisteredService\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"service\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumServiceType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EnclaveHashSet\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"valid\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"service\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumServiceType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianAdded\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianRemoved\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NitroEnclaveVerifierSet\",\"inputs\":[{\"name\":\"nitroEnclaveVerifierAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ServiceRegistered\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"service\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumServiceType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEnclaveHash\",\"inputs\":[{\"name\":\"enclaveHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"service\",\"type\":\"uint8\",\"internalType\":\"enumServiceType\"}]},{\"type\":\"error\",\"name\":\"InvalidGuardianAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNitroEnclaveVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotGuardian\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotGuardianOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[{\"name\":\"result\",\"type\":\"uint8\",\"internalType\":\"enumVerificationResult\"}]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610042575f5ffd5b5061005161006460201b60201c565b61005f61006460201b60201c565b6101df565b5f61007361016260201b60201c565b9050805f0160089054906101000a900460ff16156100bd576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461015f5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161015691906101c6565b60405180910390a15b50565b5f5f61017261017b60201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b6101c0816101a4565b82525050565b5f6020820190506101d95f8301846101b7565b92915050565b6080516143a46102055f395f818161200401528181612059015261221301526143a45ff3fe608060405260043610610203575f3560e01c80638da5cb5b11610117578063a526d83b1161009f578063cd8f69971161006e578063cd8f699714610787578063d547741f146107af578063dac79fc8146107d7578063e30c3978146107ff578063f2fde38b1461082957610203565b8063a526d83b146106d1578063a628a19e146106f9578063ad3cb1cc14610721578063ca15c8731461074b57610203565b80639546922e116100e65780639546922e146105c95780639ca6e7c2146106055780639f3eb6721461062f578063a217fddf1461066b578063a3246ad31461069557610203565b80638da5cb5b146104eb5780638fdeb2c1146105155780639010d07c1461055157806391d148541461058d57610203565b806336568abe1161019a57806354387ad71161016957806354387ad71461043157806361ff41801461045b5780637140415614610497578063715018a6146104bf57806379ba5097146104d557610203565b806336568abe1461039b578063485cc955146103c35780634f1ef286146103eb57806352d1902d1461040757610203565b80630f1f0f86116101d65780630f1f0f86146102e5578063248a9ca31461030d57806324ea54f4146103495780632f2ff15d1461037357610203565b806301ffc9a7146102075780630665f04b14610243578063094d5de21461026d5780630c68ba21146102a9575b5f5ffd5b348015610212575f5ffd5b5061022d60048036038101906102289190612fd1565b610851565b60405161023a9190613016565b60405180910390f35b34801561024e575f5ffd5b506102576108ca565b6040516102649190613116565b60405180910390f35b348015610278575f5ffd5b50610293600480360381019061028e919061318c565b6109de565b6040516102a09190613116565b60405180910390f35b3480156102b4575f5ffd5b506102cf60048036038101906102ca91906131f4565b610a39565b6040516102dc9190613016565b60405180910390f35b3480156102f0575f5ffd5b5061030b60048036038101906103069190613249565b610a6b565b005b348015610318575f5ffd5b50610333600480360381019061032e9190613299565b610bb9565b60405161034091906132d3565b60405180910390f35b348015610354575f5ffd5b5061035d610be3565b60405161036a91906132d3565b60405180910390f35b34801561037e575f5ffd5b50610399600480360381019061039491906132ec565b610c07565b005b3480156103a6575f5ffd5b506103c160048036038101906103bc91906132ec565b610c29565b005b3480156103ce575f5ffd5b506103e960048036038101906103e49190613365565b610ca4565b005b610405600480360381019061040091906134df565b610ed4565b005b348015610412575f5ffd5b5061041b610ef3565b60405161042891906132d3565b60405180910390f35b34801561043c575f5ffd5b50610445610f24565b6040516104529190613551565b60405180910390f35b348015610466575f5ffd5b50610481600480360381019061047c919061356a565b610f53565b60405161048e9190613016565b60405180910390f35b3480156104a2575f5ffd5b506104bd60048036038101906104b891906131f4565b610f7c565b005b3480156104ca575f5ffd5b506104d3611024565b005b3480156104e0575f5ffd5b506104e9611037565b005b3480156104f6575f5ffd5b506104ff6110c5565b60405161050c91906135b7565b60405180910390f35b348015610520575f5ffd5b5061053b600480360381019061053691906135d0565b6110fa565b6040516105489190613016565b60405180910390f35b34801561055c575f5ffd5b5061057760048036038101906105729190613638565b611124565b60405161058491906135b7565b60405180910390f35b348015610598575f5ffd5b506105b360048036038101906105ae91906132ec565b61115d565b6040516105c09190613016565b60405180910390f35b3480156105d4575f5ffd5b506105ef60048036038101906105ea9190613676565b6111ce565b6040516105fc9190613016565b60405180910390f35b348015610610575f5ffd5b50610619611254565b604051610626919061370f565b60405180910390f35b34801561063a575f5ffd5b506106556004803603810190610650919061318c565b611279565b6040516106629190613016565b60405180910390f35b348015610676575f5ffd5b5061067f6112d2565b60405161068c91906132d3565b60405180910390f35b3480156106a0575f5ffd5b506106bb60048036038101906106b69190613299565b6112d8565b6040516106c89190613116565b60405180910390f35b3480156106dc575f5ffd5b506106f760048036038101906106f291906131f4565b611307565b005b348015610704575f5ffd5b5061071f600480360381019061071a91906131f4565b611413565b005b34801561072c575f5ffd5b506107356114fa565b6040516107429190613788565b60405180910390f35b348015610756575f5ffd5b50610771600480360381019061076c9190613299565b611533565b60405161077e9190613551565b60405180910390f35b348015610792575f5ffd5b506107ad60048036038101906107a8919061386c565b611561565b005b3480156107ba575f5ffd5b506107d560048036038101906107d091906132ec565b61186b565b005b3480156107e2575f5ffd5b506107fd60048036038101906107f8919061391f565b61188d565b005b34801561080a575f5ffd5b50610813611d73565b60405161082091906135b7565b60405180910390f35b348015610834575f5ffd5b5061084f600480360381019061084a91906131f4565b611da8565b005b5f7f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806108c357506108c282611e61565b5b9050919050565b60605f6108f67f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a5041611533565b90505f8167ffffffffffffffff811115610913576109126133bb565b5b6040519080825280602002602001820160405280156109415781602001602082028036833780820191505090505b5090505f5f90505b828110156109d55761097b7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a504182611124565b82828151811061098e5761098d6139b0565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610949565b50809250505090565b60605f60025f8460018111156109f7576109f66139dd565b5b6001811115610a0957610a086139dd565b5b81526020019081526020015f205f8581526020019081526020015f209050610a3081611eda565b91505092915050565b5f610a647f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50418361115d565b9050919050565b610a957f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50413361115d565b158015610ad55750610aa56110c5565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610b1757336040517fd53780c4000000000000000000000000000000000000000000000000000000008152600401610b0e91906135b7565b60405180910390fd5b815f5f836001811115610b2d57610b2c6139dd565b5b6001811115610b3f57610b3e6139dd565b5b81526020019081526020015f205f8581526020019081526020015f205f6101000a81548160ff021916908315150217905550806001811115610b8457610b836139dd565b5b821515847f09821cb8037e04ceb9fc83e9a7c52c75b73a08d32a0f2f0f30cc83f2e442609060405160405180910390a4505050565b5f5f610bc3611ef9565b9050805f015f8481526020019081526020015f2060010154915050919050565b7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a504181565b610c1082610bb9565b610c1981611f20565b610c238383611f34565b50505050565b610c31611f84565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c95576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c9f8282611f8b565b505050565b5f610cad611fdb565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015610cf55750825b90505f60018367ffffffffffffffff16148015610d2857505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610d36575080155b15610d6d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610dba576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1603610e28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1f90613a7a565b60405180910390fd5b610e3186611fee565b8660035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508315610ecb575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610ec29190613ae4565b60405180910390a15b50505050505050565b610edc612002565b610ee5826120e8565b610eef82826120f3565b5050565b5f610efc612211565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b5f610f4e7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a5041611533565b905090565b5f602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b610f84612298565b610fae7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50418261115d565b1561102157610fdd7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50418261186b565b8073ffffffffffffffffffffffffffffffffffffffff167fb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c5260405160405180910390a25b50565b61102c612298565b6110355f61231f565b565b5f611040611f84565b90508073ffffffffffffffffffffffffffffffffffffffff16611061611d73565b73ffffffffffffffffffffffffffffffffffffffff16146110b957806040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016110b091906135b7565b60405180910390fd5b6110c28161231f565b50565b5f5f6110cf612385565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6001602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b5f5f61112e6123ac565b905061115483825f015f8781526020019081526020015f206123d390919063ffffffff16565b91505092915050565b5f5f611167611ef9565b9050805f015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1691505092915050565b5f60015f8360018111156111e5576111e46139dd565b5b60018111156111f7576111f66139dd565b5b81526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f5f83600181111561128f5761128e6139dd565b5b60018111156112a1576112a06139dd565b5b81526020019081526020015f205f8481526020019081526020015f205f9054906101000a900460ff16905092915050565b5f5f1b81565b60605f6112e36123ac565b90506112ff815f015f8581526020019081526020015f20611eda565b915050919050565b61130f612298565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611374576040517f1b08105400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61139e7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50418261115d565b611410576113cc7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a504182610c07565b8073ffffffffffffffffffffffffffffffffffffffff167f038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f96960405160405180910390a25b50565b61141b612298565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611480576040517f7ff9c81a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f677ca5a363c501f3c7f7291bec1cd7edc4fe7a33f375571edd1d7d3067031fe6816040516114ef91906135b7565b60405180910390a150565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f5f61153d6123ac565b9050611559815f015f8581526020019081526020015f206123ea565b915050919050565b61158b7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50413361115d565b1580156115cb575061159b6110c5565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b1561160d57336040517fd53780c400000000000000000000000000000000000000000000000000000000815260040161160491906135b7565b60405180910390fd5b5f5f90505b8251811015611866575f60025f846001811115611632576116316139dd565b5b6001811115611644576116436139dd565b5b81526020019081526020015f205f858481518110611665576116646139b0565b5b602002602001015181526020019081526020015f2090505b5f611687826123ea565b1115611792575f6116a15f836123d390919063ffffffff16565b905060015f8560018111156116b9576116b86139dd565b5b60018111156116cb576116ca6139dd565b5b81526020019081526020015f205f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81549060ff021916905561173581836123fd90919063ffffffff16565b50836001811115611749576117486139dd565b5b8173ffffffffffffffffffffffffffffffffffffffff167fbc364e6a17bd1d2abf3aff8b02c8660f8b601f3cfe343ea66df756115643f3a460405160405180910390a35061167d565b5f5f8460018111156117a7576117a66139dd565b5b60018111156117b9576117b86139dd565b5b81526020019081526020015f205f8584815181106117da576117d96139b0565b5b602002602001015181526020019081526020015f205f6101000a81549060ff0219169055826001811115611811576118106139dd565b5b848381518110611824576118236139b0565b5b60200260200101517f4f4ccf0f17d7016865671af778a980fe75a07abbbd28dcf80f9311ada39fa4e860405160405180910390a3508080600101915050611612565b505050565b61187482610bb9565b61187d81611f20565b6118878383611f8b565b50505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d114be38787600288886040518663ffffffff1660e01b81526004016118f1959493929190613b7f565b5f604051808303815f875af115801561190c573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190611934919061413a565b90505f6003811115611949576119486139dd565b5b815f0151600381111561195f5761195e6139dd565b5b146119a457805f01516040517f470e419c00000000000000000000000000000000000000000000000000000000815260040161199b91906141c7565b60405180910390fd5b5f8160e001515f815181106119bc576119bb6139b0565b5b6020026020010151602001515f01518260e001515f815181106119e2576119e16139b0565b5b60200260200101516020015160200151604051602001611a03929190614220565b6040516020818303038152906040528051906020012090505f5f846001811115611a3057611a2f6139dd565b5b6001811115611a4257611a416139dd565b5b81526020019081526020015f205f8281526020019081526020015f205f9054906101000a900460ff16611aae5780836040517f9f2e2b21000000000000000000000000000000000000000000000000000000008152600401611aa5929190614291565b60405180910390fd5b5f60018360c0015151611ac191906142e5565b67ffffffffffffffff811115611ada57611ad96133bb565b5b6040519080825280601f01601f191660200182016040528015611b0c5781602001600182028036833780820191505090505b5090505f600190505b8360c0015151811015611b9e578360c001518181518110611b3957611b386139b0565b5b602001015160f81c60f81b82600183611b5291906142e5565b81518110611b6357611b626139b0565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053508080600101915050611b15565b505f818051906020012090505f815f1c905060015f876001811115611bc657611bc56139dd565b5b6001811115611bd857611bd76139dd565b5b81526020019081526020015f205f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16611d67576001805f886001811115611c4857611c476139dd565b5b6001811115611c5a57611c596139dd565b5b81526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550611d0e8160025f896001811115611cd257611cd16139dd565b5b6001811115611ce457611ce36139dd565b5b81526020019081526020015f205f8781526020019081526020015f2061242a90919063ffffffff16565b50856001811115611d2257611d216139dd565b5b848273ffffffffffffffffffffffffffffffffffffffff167f0fa700ad17f1b256f57e62054a779bd9fe08e585f8cacb3def28cca47b25cdc160405160405180910390a45b50505050505050505050565b5f5f611d7d612457565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b611db0612298565b5f611db9612457565b905081815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff16611e1b6110c5565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480611ed35750611ed28261247e565b5b9050919050565b60605f611ee8835f016124e7565b905060608190508092505050919050565b5f7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800905090565b611f3181611f2c611f84565b612540565b50565b5f5f611f3e6123ac565b90505f611f4b8585612591565b90508015611f7957611f7784835f015f8881526020019081526020015f2061242a90919063ffffffff16565b505b809250505092915050565b5f33905090565b5f5f611f956123ac565b90505f611fa28585612689565b90508015611fd057611fce84835f015f8881526020019081526020015f206123fd90919063ffffffff16565b505b809250505092915050565b5f5f611fe5612781565b90508091505090565b611ff66127aa565b611fff816127ea565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614806120af57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661209661281f565b73ffffffffffffffffffffffffffffffffffffffff1614155b156120e6576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6120f0612298565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561215b57506040513d601f19601f820116820180604052508101906121589190614318565b60015b61219c57816040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815260040161219391906135b7565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b811461220257806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016121f991906132d3565b60405180910390fd5b61220c8383612872565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614612296576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6122a0611f84565b73ffffffffffffffffffffffffffffffffffffffff166122be6110c5565b73ffffffffffffffffffffffffffffffffffffffff161461231d576122e1611f84565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161231491906135b7565b60405180910390fd5b565b5f6123286110c5565b9050612333826128e4565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612374576123725f5f1b82611f8b565b505b6123805f5f1b83611f34565b505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000905090565b5f6123e0835f0183612921565b5f1c905092915050565b5f6123f6825f01612948565b9050919050565b5f612422835f018373ffffffffffffffffffffffffffffffffffffffff165f1b612957565b905092915050565b5f61244f835f018373ffffffffffffffffffffffffffffffffffffffff165f1b612a53565b905092915050565b5f7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00905090565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561253457602002820191905f5260205f20905b815481526020019060010190808311612520575b50505050509050919050565b61254a828261115d565b61258d5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401612584929190614343565b60405180910390fd5b5050565b5f5f61259b611ef9565b90506125a7848461115d565b61267e576001815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061261a611f84565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050612683565b5f9150505b92915050565b5f5f612693611ef9565b905061269f848461115d565b15612776575f815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612712611f84565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a4600191505061277b565b5f9150505b92915050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b6127b2612aba565b6127e8576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6127f26127aa565b6127fb81612ad8565b612803612aec565b61280b612af6565b612813612b00565b61281c81612b0a565b50565b5f61284b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b612b4e565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61287b82612b57565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f815111156128d7576128d18282612c20565b506128e0565b6128df612d11565b5b5050565b5f6128ed612457565b9050805f015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905561291d82612d4d565b5050565b5f825f018281548110612937576129366139b0565b5b905f5260205f200154905092915050565b5f815f01805490509050919050565b5f5f836001015f8481526020019081526020015f205490505f8114612a48575f60018261298491906142e5565b90505f6001865f018054905061299a91906142e5565b9050808214612a00575f865f0182815481106129b9576129b86139b0565b5b905f5260205f200154905080875f0184815481106129da576129d96139b0565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f01805480612a1357612a1261436a565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050612a4d565b5f9150505b92915050565b5f612a5e8383612e1e565b612ab057825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f208190555060019050612ab4565b5f90505b92915050565b5f612ac3611fdb565b5f0160089054906101000a900460ff16905090565b612ae06127aa565b612ae981612e3e565b50565b612af46127aa565b565b612afe6127aa565b565b612b086127aa565b565b612b126127aa565b612b1e5f5f1b82611f34565b50612b4b7f55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a50415f5f1b612ec2565b50565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03612bb257806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401612ba991906135b7565b60405180910390fd5b80612bde7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b612b4e565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f612c2d8484612f28565b9050808015612c6357505f612c40612f3c565b1180612c6257505f8473ffffffffffffffffffffffffffffffffffffffff163b115b5b15612c7857612c70612f43565b915050612d0b565b8015612cbb57836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401612cb291906135b7565b60405180910390fd5b5f612cc4612f3c565b1115612cd757612cd2612f60565b612d09565b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b92915050565b5f341115612d4b576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f612d56612385565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f5f836001015f8481526020019081526020015f20541415905092915050565b612e466127aa565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612eb6575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401612ead91906135b7565b60405180910390fd5b612ebf8161231f565b50565b5f612ecb611ef9565b90505f612ed784610bb9565b905082825f015f8681526020019081526020015f20600101819055508281857fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a450505050565b5f5f5f835160208501865af4905092915050565b5f3d905090565b606060405190503d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b612fb081612f7c565b8114612fba575f5ffd5b50565b5f81359050612fcb81612fa7565b92915050565b5f60208284031215612fe657612fe5612f74565b5b5f612ff384828501612fbd565b91505092915050565b5f8115159050919050565b61301081612ffc565b82525050565b5f6020820190506130295f830184613007565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61308182613058565b9050919050565b61309181613077565b82525050565b5f6130a28383613088565b60208301905092915050565b5f602082019050919050565b5f6130c48261302f565b6130ce8185613039565b93506130d983613049565b805f5b838110156131095781516130f08882613097565b97506130fb836130ae565b9250506001810190506130dc565b5085935050505092915050565b5f6020820190508181035f83015261312e81846130ba565b905092915050565b5f819050919050565b61314881613136565b8114613152575f5ffd5b50565b5f813590506131638161313f565b92915050565b60028110613175575f5ffd5b50565b5f8135905061318681613169565b92915050565b5f5f604083850312156131a2576131a1612f74565b5b5f6131af85828601613155565b92505060206131c085828601613178565b9150509250929050565b6131d381613077565b81146131dd575f5ffd5b50565b5f813590506131ee816131ca565b92915050565b5f6020828403121561320957613208612f74565b5b5f613216848285016131e0565b91505092915050565b61322881612ffc565b8114613232575f5ffd5b50565b5f813590506132438161321f565b92915050565b5f5f5f606084860312156132605761325f612f74565b5b5f61326d86828701613155565b935050602061327e86828701613235565b925050604061328f86828701613178565b9150509250925092565b5f602082840312156132ae576132ad612f74565b5b5f6132bb84828501613155565b91505092915050565b6132cd81613136565b82525050565b5f6020820190506132e65f8301846132c4565b92915050565b5f5f6040838503121561330257613301612f74565b5b5f61330f85828601613155565b9250506020613320858286016131e0565b9150509250929050565b5f61333482613077565b9050919050565b6133448161332a565b811461334e575f5ffd5b50565b5f8135905061335f8161333b565b92915050565b5f5f6040838503121561337b5761337a612f74565b5b5f61338885828601613351565b9250506020613399858286016131e0565b9150509250929050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6133f1826133ab565b810181811067ffffffffffffffff821117156134105761340f6133bb565b5b80604052505050565b5f613422612f6b565b905061342e82826133e8565b919050565b5f67ffffffffffffffff82111561344d5761344c6133bb565b5b613456826133ab565b9050602081019050919050565b828183375f83830152505050565b5f61348361347e84613433565b613419565b90508281526020810184848401111561349f5761349e6133a7565b5b6134aa848285613463565b509392505050565b5f82601f8301126134c6576134c56133a3565b5b81356134d6848260208601613471565b91505092915050565b5f5f604083850312156134f5576134f4612f74565b5b5f613502858286016131e0565b925050602083013567ffffffffffffffff81111561352357613522612f78565b5b61352f858286016134b2565b9150509250929050565b5f819050919050565b61354b81613539565b82525050565b5f6020820190506135645f830184613542565b92915050565b5f5f604083850312156135805761357f612f74565b5b5f61358d85828601613178565b925050602061359e85828601613155565b9150509250929050565b6135b181613077565b82525050565b5f6020820190506135ca5f8301846135a8565b92915050565b5f5f604083850312156135e6576135e5612f74565b5b5f6135f385828601613178565b9250506020613604858286016131e0565b9150509250929050565b61361781613539565b8114613621575f5ffd5b50565b5f813590506136328161360e565b92915050565b5f5f6040838503121561364e5761364d612f74565b5b5f61365b85828601613155565b925050602061366c85828601613624565b9150509250929050565b5f5f6040838503121561368c5761368b612f74565b5b5f613699858286016131e0565b92505060206136aa85828601613178565b9150509250929050565b5f819050919050565b5f6136d76136d26136cd84613058565b6136b4565b613058565b9050919050565b5f6136e8826136bd565b9050919050565b5f6136f9826136de565b9050919050565b613709816136ef565b82525050565b5f6020820190506137225f830184613700565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61375a82613728565b6137648185613732565b9350613774818560208601613742565b61377d816133ab565b840191505092915050565b5f6020820190508181035f8301526137a08184613750565b905092915050565b5f67ffffffffffffffff8211156137c2576137c16133bb565b5b602082029050602081019050919050565b5f5ffd5b5f6137e96137e4846137a8565b613419565b9050808382526020820190506020840283018581111561380c5761380b6137d3565b5b835b8181101561383557806138218882613155565b84526020840193505060208101905061380e565b5050509392505050565b5f82601f830112613853576138526133a3565b5b81356138638482602086016137d7565b91505092915050565b5f5f6040838503121561388257613881612f74565b5b5f83013567ffffffffffffffff81111561389f5761389e612f78565b5b6138ab8582860161383f565b92505060206138bc85828601613178565b9150509250929050565b5f5ffd5b5f5f83601f8401126138df576138de6133a3565b5b8235905067ffffffffffffffff8111156138fc576138fb6138c6565b5b602083019150836001820283011115613918576139176137d3565b5b9250929050565b5f5f5f5f5f6060868803121561393857613937612f74565b5b5f86013567ffffffffffffffff81111561395557613954612f78565b5b613961888289016138ca565b9550955050602086013567ffffffffffffffff81111561398457613983612f78565b5b613990888289016138ca565b935093505060406139a388828901613178565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e6974726f456e636c61766556657269666965722063616e6e6f74206265207a5f8201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b5f613a64602383613732565b9150613a6f82613a0a565b604082019050919050565b5f6020820190508181035f830152613a9181613a58565b9050919050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f613ace613ac9613ac484613a98565b6136b4565b613aa1565b9050919050565b613ade81613ab4565b82525050565b5f602082019050613af75f830184613ad5565b92915050565b5f82825260208201905092915050565b5f613b188385613afd565b9350613b25838584613463565b613b2e836133ab565b840190509392505050565b60038110613b4a57613b496139dd565b5b50565b5f819050613b5a82613b39565b919050565b5f613b6982613b4d565b9050919050565b613b7981613b5f565b82525050565b5f6060820190508181035f830152613b98818789613b0d565b9050613ba76020830186613b70565b8181036040830152613bba818486613b0d565b90509695505050505050565b5f5ffd5b5f5ffd5b60048110613bda575f5ffd5b50565b5f81519050613beb81613bce565b92915050565b5f60ff82169050919050565b613c0681613bf1565b8114613c10575f5ffd5b50565b5f81519050613c2181613bfd565b92915050565b613c3081613aa1565b8114613c3a575f5ffd5b50565b5f81519050613c4b81613c27565b92915050565b5f81519050613c5f8161313f565b92915050565b5f613c77613c72846137a8565b613419565b90508083825260208201905060208402830185811115613c9a57613c996137d3565b5b835b81811015613cc35780613caf8882613c51565b845260208401935050602081019050613c9c565b5050509392505050565b5f82601f830112613ce157613ce06133a3565b5b8151613cf1848260208601613c65565b91505092915050565b5f613d0c613d0784613433565b613419565b905082815260208101848484011115613d2857613d276133a7565b5b613d33848285613742565b509392505050565b5f82601f830112613d4f57613d4e6133a3565b5b8151613d5f848260208601613cfa565b91505092915050565b5f67ffffffffffffffff821115613d8257613d816133bb565b5b602082029050602081019050919050565b5f7fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b613dc781613d93565b8114613dd1575f5ffd5b50565b5f81519050613de281613dbe565b92915050565b5f60408284031215613dfd57613dfc613bc6565b5b613e076040613419565b90505f613e1684828501613c51565b5f830152506020613e2984828501613dd4565b60208301525092915050565b5f60608284031215613e4a57613e49613bc6565b5b613e546040613419565b90505f613e6384828501613c3d565b5f830152506020613e7684828501613de8565b60208301525092915050565b5f613e94613e8f84613d68565b613419565b90508083825260208201905060608402830185811115613eb757613eb66137d3565b5b835b81811015613ee05780613ecc8882613e35565b845260208401935050606081019050613eb9565b5050509392505050565b5f82601f830112613efe57613efd6133a3565b5b8151613f0e848260208601613e82565b91505092915050565b5f67ffffffffffffffff821115613f3157613f306133bb565b5b613f3a826133ab565b9050602081019050919050565b5f613f59613f5484613f17565b613419565b905082815260208101848484011115613f7557613f746133a7565b5b613f80848285613742565b509392505050565b5f82601f830112613f9c57613f9b6133a3565b5b8151613fac848260208601613f47565b91505092915050565b5f6101208284031215613fcb57613fca613bc6565b5b613fd6610120613419565b90505f613fe584828501613bdd565b5f830152506020613ff884828501613c13565b602083015250604061400c84828501613c3d565b604083015250606082015167ffffffffffffffff8111156140305761402f613bca565b5b61403c84828501613ccd565b606083015250608082015167ffffffffffffffff8111156140605761405f613bca565b5b61406c84828501613d3b565b60808301525060a082015167ffffffffffffffff8111156140905761408f613bca565b5b61409c84828501613d3b565b60a08301525060c082015167ffffffffffffffff8111156140c0576140bf613bca565b5b6140cc84828501613d3b565b60c08301525060e082015167ffffffffffffffff8111156140f0576140ef613bca565b5b6140fc84828501613eea565b60e08301525061010082015167ffffffffffffffff81111561412157614120613bca565b5b61412d84828501613f88565b6101008301525092915050565b5f6020828403121561414f5761414e612f74565b5b5f82015167ffffffffffffffff81111561416c5761416b612f78565b5b61417884828501613fb5565b91505092915050565b60048110614192576141916139dd565b5b50565b5f8190506141a282614181565b919050565b5f6141b182614195565b9050919050565b6141c1816141a7565b82525050565b5f6020820190506141da5f8301846141b8565b92915050565b5f819050919050565b6141fa6141f582613136565b6141e0565b82525050565b5f819050919050565b61421a61421582613d93565b614200565b82525050565b5f61422b82856141e9565b60208201915061423b8284614209565b6010820191508190509392505050565b6002811061425c5761425b6139dd565b5b50565b5f81905061426c8261424b565b919050565b5f61427b8261425f565b9050919050565b61428b81614271565b82525050565b5f6040820190506142a45f8301856132c4565b6142b16020830184614282565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6142ef82613539565b91506142fa83613539565b9250828203905081811115614312576143116142b8565b5b92915050565b5f6020828403121561432d5761432c612f74565b5b5f61433a84828501613c51565b91505092915050565b5f6040820190506143565f8301856135a8565b61436360208301846132c4565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea164736f6c634300081c000a",
}

// EspressoNitroTEEVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use EspressoNitroTEEVerifierMetaData.ABI instead.
var EspressoNitroTEEVerifierABI = EspressoNitroTEEVerifierMetaData.ABI

// EspressoNitroTEEVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EspressoNitroTEEVerifierMetaData.Bin instead.
var EspressoNitroTEEVerifierBin = EspressoNitroTEEVerifierMetaData.Bin

// DeployEspressoNitroTEEVerifier deploys a new Ethereum contract, binding an instance of EspressoNitroTEEVerifier to it.
func DeployEspressoNitroTEEVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EspressoNitroTEEVerifier, error) {
	parsed, err := EspressoNitroTEEVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EspressoNitroTEEVerifierBin), backend)
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.DEFAULTADMINROLE(&_EspressoNitroTEEVerifier.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.DEFAULTADMINROLE(&_EspressoNitroTEEVerifier.CallOpts)
}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) GUARDIANROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "GUARDIAN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GUARDIANROLE() ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.GUARDIANROLE(&_EspressoNitroTEEVerifier.CallOpts)
}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) GUARDIANROLE() ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.GUARDIANROLE(&_EspressoNitroTEEVerifier.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EspressoNitroTEEVerifier.Contract.UPGRADEINTERFACEVERSION(&_EspressoNitroTEEVerifier.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EspressoNitroTEEVerifier.Contract.UPGRADEINTERFACEVERSION(&_EspressoNitroTEEVerifier.CallOpts)
}

// NitroEnclaveVerifier is a free data retrieval call binding the contract method 0x9ca6e7c2.
//
// Solidity: function _nitroEnclaveVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) NitroEnclaveVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "_nitroEnclaveVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NitroEnclaveVerifier is a free data retrieval call binding the contract method 0x9ca6e7c2.
//
// Solidity: function _nitroEnclaveVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) NitroEnclaveVerifier() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.NitroEnclaveVerifier(&_EspressoNitroTEEVerifier.CallOpts)
}

// NitroEnclaveVerifier is a free data retrieval call binding the contract method 0x9ca6e7c2.
//
// Solidity: function _nitroEnclaveVerifier() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) NitroEnclaveVerifier() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.NitroEnclaveVerifier(&_EspressoNitroTEEVerifier.CallOpts)
}

// EnclaveHashSigners is a free data retrieval call binding the contract method 0x094d5de2.
//
// Solidity: function enclaveHashSigners(bytes32 enclaveHash, uint8 service) view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) EnclaveHashSigners(opts *bind.CallOpts, enclaveHash [32]byte, service uint8) ([]common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "enclaveHashSigners", enclaveHash, service)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnclaveHashSigners is a free data retrieval call binding the contract method 0x094d5de2.
//
// Solidity: function enclaveHashSigners(bytes32 enclaveHash, uint8 service) view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) EnclaveHashSigners(enclaveHash [32]byte, service uint8) ([]common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.EnclaveHashSigners(&_EspressoNitroTEEVerifier.CallOpts, enclaveHash, service)
}

// EnclaveHashSigners is a free data retrieval call binding the contract method 0x094d5de2.
//
// Solidity: function enclaveHashSigners(bytes32 enclaveHash, uint8 service) view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) EnclaveHashSigners(enclaveHash [32]byte, service uint8) ([]common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.EnclaveHashSigners(&_EspressoNitroTEEVerifier.CallOpts, enclaveHash, service)
}

// GetGuardians is a free data retrieval call binding the contract method 0x0665f04b.
//
// Solidity: function getGuardians() view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) GetGuardians(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "getGuardians")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGuardians is a free data retrieval call binding the contract method 0x0665f04b.
//
// Solidity: function getGuardians() view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GetGuardians() ([]common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.GetGuardians(&_EspressoNitroTEEVerifier.CallOpts)
}

// GetGuardians is a free data retrieval call binding the contract method 0x0665f04b.
//
// Solidity: function getGuardians() view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) GetGuardians() ([]common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.GetGuardians(&_EspressoNitroTEEVerifier.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleAdmin(&_EspressoNitroTEEVerifier.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleAdmin(&_EspressoNitroTEEVerifier.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleMember(&_EspressoNitroTEEVerifier.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleMember(&_EspressoNitroTEEVerifier.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleMemberCount(&_EspressoNitroTEEVerifier.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleMemberCount(&_EspressoNitroTEEVerifier.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleMembers(&_EspressoNitroTEEVerifier.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.GetRoleMembers(&_EspressoNitroTEEVerifier.CallOpts, role)
}

// GuardianCount is a free data retrieval call binding the contract method 0x54387ad7.
//
// Solidity: function guardianCount() view returns(uint256)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) GuardianCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "guardianCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardianCount is a free data retrieval call binding the contract method 0x54387ad7.
//
// Solidity: function guardianCount() view returns(uint256)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GuardianCount() (*big.Int, error) {
	return _EspressoNitroTEEVerifier.Contract.GuardianCount(&_EspressoNitroTEEVerifier.CallOpts)
}

// GuardianCount is a free data retrieval call binding the contract method 0x54387ad7.
//
// Solidity: function guardianCount() view returns(uint256)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) GuardianCount() (*big.Int, error) {
	return _EspressoNitroTEEVerifier.Contract.GuardianCount(&_EspressoNitroTEEVerifier.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.HasRole(&_EspressoNitroTEEVerifier.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.HasRole(&_EspressoNitroTEEVerifier.CallOpts, role, account)
}

// IsGuardian is a free data retrieval call binding the contract method 0x0c68ba21.
//
// Solidity: function isGuardian(address account) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) IsGuardian(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "isGuardian", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGuardian is a free data retrieval call binding the contract method 0x0c68ba21.
//
// Solidity: function isGuardian(address account) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) IsGuardian(account common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.IsGuardian(&_EspressoNitroTEEVerifier.CallOpts, account)
}

// IsGuardian is a free data retrieval call binding the contract method 0x0c68ba21.
//
// Solidity: function isGuardian(address account) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) IsGuardian(account common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.IsGuardian(&_EspressoNitroTEEVerifier.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) Owner() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.Owner(&_EspressoNitroTEEVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) Owner() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.Owner(&_EspressoNitroTEEVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) PendingOwner() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.PendingOwner(&_EspressoNitroTEEVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) PendingOwner() (common.Address, error) {
	return _EspressoNitroTEEVerifier.Contract.PendingOwner(&_EspressoNitroTEEVerifier.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) ProxiableUUID() ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.ProxiableUUID(&_EspressoNitroTEEVerifier.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EspressoNitroTEEVerifier.Contract.ProxiableUUID(&_EspressoNitroTEEVerifier.CallOpts)
}

// RegisteredEnclaveHash is a free data retrieval call binding the contract method 0x9f3eb672.
//
// Solidity: function registeredEnclaveHash(bytes32 enclaveHash, uint8 service) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) RegisteredEnclaveHash(opts *bind.CallOpts, enclaveHash [32]byte, service uint8) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "registeredEnclaveHash", enclaveHash, service)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredEnclaveHash is a free data retrieval call binding the contract method 0x9f3eb672.
//
// Solidity: function registeredEnclaveHash(bytes32 enclaveHash, uint8 service) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RegisteredEnclaveHash(enclaveHash [32]byte, service uint8) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredEnclaveHash(&_EspressoNitroTEEVerifier.CallOpts, enclaveHash, service)
}

// RegisteredEnclaveHash is a free data retrieval call binding the contract method 0x9f3eb672.
//
// Solidity: function registeredEnclaveHash(bytes32 enclaveHash, uint8 service) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) RegisteredEnclaveHash(enclaveHash [32]byte, service uint8) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredEnclaveHash(&_EspressoNitroTEEVerifier.CallOpts, enclaveHash, service)
}

// RegisteredEnclaveHashes is a free data retrieval call binding the contract method 0x61ff4180.
//
// Solidity: function registeredEnclaveHashes(uint8 , bytes32 enclaveHash) view returns(bool valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) RegisteredEnclaveHashes(opts *bind.CallOpts, arg0 uint8, enclaveHash [32]byte) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "registeredEnclaveHashes", arg0, enclaveHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredEnclaveHashes is a free data retrieval call binding the contract method 0x61ff4180.
//
// Solidity: function registeredEnclaveHashes(uint8 , bytes32 enclaveHash) view returns(bool valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RegisteredEnclaveHashes(arg0 uint8, enclaveHash [32]byte) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredEnclaveHashes(&_EspressoNitroTEEVerifier.CallOpts, arg0, enclaveHash)
}

// RegisteredEnclaveHashes is a free data retrieval call binding the contract method 0x61ff4180.
//
// Solidity: function registeredEnclaveHashes(uint8 , bytes32 enclaveHash) view returns(bool valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) RegisteredEnclaveHashes(arg0 uint8, enclaveHash [32]byte) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredEnclaveHashes(&_EspressoNitroTEEVerifier.CallOpts, arg0, enclaveHash)
}

// RegisteredService is a free data retrieval call binding the contract method 0x9546922e.
//
// Solidity: function registeredService(address signer, uint8 service) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) RegisteredService(opts *bind.CallOpts, signer common.Address, service uint8) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "registeredService", signer, service)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredService is a free data retrieval call binding the contract method 0x9546922e.
//
// Solidity: function registeredService(address signer, uint8 service) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RegisteredService(signer common.Address, service uint8) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredService(&_EspressoNitroTEEVerifier.CallOpts, signer, service)
}

// RegisteredService is a free data retrieval call binding the contract method 0x9546922e.
//
// Solidity: function registeredService(address signer, uint8 service) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) RegisteredService(signer common.Address, service uint8) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredService(&_EspressoNitroTEEVerifier.CallOpts, signer, service)
}

// RegisteredServices is a free data retrieval call binding the contract method 0x8fdeb2c1.
//
// Solidity: function registeredServices(uint8 , address signer) view returns(bool valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) RegisteredServices(opts *bind.CallOpts, arg0 uint8, signer common.Address) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "registeredServices", arg0, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredServices is a free data retrieval call binding the contract method 0x8fdeb2c1.
//
// Solidity: function registeredServices(uint8 , address signer) view returns(bool valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RegisteredServices(arg0 uint8, signer common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredServices(&_EspressoNitroTEEVerifier.CallOpts, arg0, signer)
}

// RegisteredServices is a free data retrieval call binding the contract method 0x8fdeb2c1.
//
// Solidity: function registeredServices(uint8 , address signer) view returns(bool valid)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) RegisteredServices(arg0 uint8, signer common.Address) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisteredServices(&_EspressoNitroTEEVerifier.CallOpts, arg0, signer)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EspressoNitroTEEVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.SupportsInterface(&_EspressoNitroTEEVerifier.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EspressoNitroTEEVerifier.Contract.SupportsInterface(&_EspressoNitroTEEVerifier.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.AcceptOwnership(&_EspressoNitroTEEVerifier.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.AcceptOwnership(&_EspressoNitroTEEVerifier.TransactOpts)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.AddGuardian(&_EspressoNitroTEEVerifier.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.AddGuardian(&_EspressoNitroTEEVerifier.TransactOpts, guardian)
}

// DeleteEnclaveHashes is a paid mutator transaction binding the contract method 0xcd8f6997.
//
// Solidity: function deleteEnclaveHashes(bytes32[] enclaveHashes, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) DeleteEnclaveHashes(opts *bind.TransactOpts, enclaveHashes [][32]byte, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "deleteEnclaveHashes", enclaveHashes, service)
}

// DeleteEnclaveHashes is a paid mutator transaction binding the contract method 0xcd8f6997.
//
// Solidity: function deleteEnclaveHashes(bytes32[] enclaveHashes, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) DeleteEnclaveHashes(enclaveHashes [][32]byte, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.DeleteEnclaveHashes(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHashes, service)
}

// DeleteEnclaveHashes is a paid mutator transaction binding the contract method 0xcd8f6997.
//
// Solidity: function deleteEnclaveHashes(bytes32[] enclaveHashes, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) DeleteEnclaveHashes(enclaveHashes [][32]byte, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.DeleteEnclaveHashes(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHashes, service)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.GrantRole(&_EspressoNitroTEEVerifier.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.GrantRole(&_EspressoNitroTEEVerifier.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address nitroEnclaveVerifier, address initialOwner) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) Initialize(opts *bind.TransactOpts, nitroEnclaveVerifier common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "initialize", nitroEnclaveVerifier, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address nitroEnclaveVerifier, address initialOwner) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) Initialize(nitroEnclaveVerifier common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.Initialize(&_EspressoNitroTEEVerifier.TransactOpts, nitroEnclaveVerifier, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address nitroEnclaveVerifier, address initialOwner) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) Initialize(nitroEnclaveVerifier common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.Initialize(&_EspressoNitroTEEVerifier.TransactOpts, nitroEnclaveVerifier, initialOwner)
}

// RegisterService is a paid mutator transaction binding the contract method 0xdac79fc8.
//
// Solidity: function registerService(bytes output, bytes proofBytes, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) RegisterService(opts *bind.TransactOpts, output []byte, proofBytes []byte, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "registerService", output, proofBytes, service)
}

// RegisterService is a paid mutator transaction binding the contract method 0xdac79fc8.
//
// Solidity: function registerService(bytes output, bytes proofBytes, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RegisterService(output []byte, proofBytes []byte, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisterService(&_EspressoNitroTEEVerifier.TransactOpts, output, proofBytes, service)
}

// RegisterService is a paid mutator transaction binding the contract method 0xdac79fc8.
//
// Solidity: function registerService(bytes output, bytes proofBytes, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) RegisterService(output []byte, proofBytes []byte, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RegisterService(&_EspressoNitroTEEVerifier.TransactOpts, output, proofBytes, service)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RemoveGuardian(&_EspressoNitroTEEVerifier.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RemoveGuardian(&_EspressoNitroTEEVerifier.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RenounceOwnership(&_EspressoNitroTEEVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RenounceOwnership(&_EspressoNitroTEEVerifier.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RenounceRole(&_EspressoNitroTEEVerifier.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RenounceRole(&_EspressoNitroTEEVerifier.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RevokeRole(&_EspressoNitroTEEVerifier.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.RevokeRole(&_EspressoNitroTEEVerifier.TransactOpts, role, account)
}

// SetEnclaveHash is a paid mutator transaction binding the contract method 0x0f1f0f86.
//
// Solidity: function setEnclaveHash(bytes32 enclaveHash, bool valid, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) SetEnclaveHash(opts *bind.TransactOpts, enclaveHash [32]byte, valid bool, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "setEnclaveHash", enclaveHash, valid, service)
}

// SetEnclaveHash is a paid mutator transaction binding the contract method 0x0f1f0f86.
//
// Solidity: function setEnclaveHash(bytes32 enclaveHash, bool valid, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) SetEnclaveHash(enclaveHash [32]byte, valid bool, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetEnclaveHash(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHash, valid, service)
}

// SetEnclaveHash is a paid mutator transaction binding the contract method 0x0f1f0f86.
//
// Solidity: function setEnclaveHash(bytes32 enclaveHash, bool valid, uint8 service) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) SetEnclaveHash(enclaveHash [32]byte, valid bool, service uint8) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetEnclaveHash(&_EspressoNitroTEEVerifier.TransactOpts, enclaveHash, valid, service)
}

// SetNitroEnclaveVerifier is a paid mutator transaction binding the contract method 0xa628a19e.
//
// Solidity: function setNitroEnclaveVerifier(address nitroEnclaveVerifier) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) SetNitroEnclaveVerifier(opts *bind.TransactOpts, nitroEnclaveVerifier common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "setNitroEnclaveVerifier", nitroEnclaveVerifier)
}

// SetNitroEnclaveVerifier is a paid mutator transaction binding the contract method 0xa628a19e.
//
// Solidity: function setNitroEnclaveVerifier(address nitroEnclaveVerifier) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) SetNitroEnclaveVerifier(nitroEnclaveVerifier common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetNitroEnclaveVerifier(&_EspressoNitroTEEVerifier.TransactOpts, nitroEnclaveVerifier)
}

// SetNitroEnclaveVerifier is a paid mutator transaction binding the contract method 0xa628a19e.
//
// Solidity: function setNitroEnclaveVerifier(address nitroEnclaveVerifier) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) SetNitroEnclaveVerifier(nitroEnclaveVerifier common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.SetNitroEnclaveVerifier(&_EspressoNitroTEEVerifier.TransactOpts, nitroEnclaveVerifier)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.TransferOwnership(&_EspressoNitroTEEVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.TransferOwnership(&_EspressoNitroTEEVerifier.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.UpgradeToAndCall(&_EspressoNitroTEEVerifier.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EspressoNitroTEEVerifier.Contract.UpgradeToAndCall(&_EspressoNitroTEEVerifier.TransactOpts, newImplementation, data)
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
	Service     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeletedEnclaveHash is a free log retrieval operation binding the contract event 0x4f4ccf0f17d7016865671af778a980fe75a07abbbd28dcf80f9311ada39fa4e8.
//
// Solidity: event DeletedEnclaveHash(bytes32 indexed enclaveHash, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterDeletedEnclaveHash(opts *bind.FilterOpts, enclaveHash [][32]byte, service []uint8) (*EspressoNitroTEEVerifierDeletedEnclaveHashIterator, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "DeletedEnclaveHash", enclaveHashRule, serviceRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierDeletedEnclaveHashIterator{contract: _EspressoNitroTEEVerifier.contract, event: "DeletedEnclaveHash", logs: logs, sub: sub}, nil
}

// WatchDeletedEnclaveHash is a free log subscription operation binding the contract event 0x4f4ccf0f17d7016865671af778a980fe75a07abbbd28dcf80f9311ada39fa4e8.
//
// Solidity: event DeletedEnclaveHash(bytes32 indexed enclaveHash, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchDeletedEnclaveHash(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierDeletedEnclaveHash, enclaveHash [][32]byte, service []uint8) (event.Subscription, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "DeletedEnclaveHash", enclaveHashRule, serviceRule)
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

// ParseDeletedEnclaveHash is a log parse operation binding the contract event 0x4f4ccf0f17d7016865671af778a980fe75a07abbbd28dcf80f9311ada39fa4e8.
//
// Solidity: event DeletedEnclaveHash(bytes32 indexed enclaveHash, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseDeletedEnclaveHash(log types.Log) (*EspressoNitroTEEVerifierDeletedEnclaveHash, error) {
	event := new(EspressoNitroTEEVerifierDeletedEnclaveHash)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "DeletedEnclaveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierDeletedRegisteredServiceIterator is returned from FilterDeletedRegisteredService and is used to iterate over the raw logs and unpacked data for DeletedRegisteredService events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierDeletedRegisteredServiceIterator struct {
	Event *EspressoNitroTEEVerifierDeletedRegisteredService // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierDeletedRegisteredServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierDeletedRegisteredService)
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
		it.Event = new(EspressoNitroTEEVerifierDeletedRegisteredService)
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
func (it *EspressoNitroTEEVerifierDeletedRegisteredServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierDeletedRegisteredServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierDeletedRegisteredService represents a DeletedRegisteredService event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierDeletedRegisteredService struct {
	Signer  common.Address
	Service uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeletedRegisteredService is a free log retrieval operation binding the contract event 0xbc364e6a17bd1d2abf3aff8b02c8660f8b601f3cfe343ea66df756115643f3a4.
//
// Solidity: event DeletedRegisteredService(address indexed signer, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterDeletedRegisteredService(opts *bind.FilterOpts, signer []common.Address, service []uint8) (*EspressoNitroTEEVerifierDeletedRegisteredServiceIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "DeletedRegisteredService", signerRule, serviceRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierDeletedRegisteredServiceIterator{contract: _EspressoNitroTEEVerifier.contract, event: "DeletedRegisteredService", logs: logs, sub: sub}, nil
}

// WatchDeletedRegisteredService is a free log subscription operation binding the contract event 0xbc364e6a17bd1d2abf3aff8b02c8660f8b601f3cfe343ea66df756115643f3a4.
//
// Solidity: event DeletedRegisteredService(address indexed signer, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchDeletedRegisteredService(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierDeletedRegisteredService, signer []common.Address, service []uint8) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "DeletedRegisteredService", signerRule, serviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierDeletedRegisteredService)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "DeletedRegisteredService", log); err != nil {
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

// ParseDeletedRegisteredService is a log parse operation binding the contract event 0xbc364e6a17bd1d2abf3aff8b02c8660f8b601f3cfe343ea66df756115643f3a4.
//
// Solidity: event DeletedRegisteredService(address indexed signer, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseDeletedRegisteredService(log types.Log) (*EspressoNitroTEEVerifierDeletedRegisteredService, error) {
	event := new(EspressoNitroTEEVerifierDeletedRegisteredService)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "DeletedRegisteredService", log); err != nil {
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
	Service     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEnclaveHashSet is a free log retrieval operation binding the contract event 0x09821cb8037e04ceb9fc83e9a7c52c75b73a08d32a0f2f0f30cc83f2e4426090.
//
// Solidity: event EnclaveHashSet(bytes32 indexed enclaveHash, bool indexed valid, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterEnclaveHashSet(opts *bind.FilterOpts, enclaveHash [][32]byte, valid []bool, service []uint8) (*EspressoNitroTEEVerifierEnclaveHashSetIterator, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var validRule []interface{}
	for _, validItem := range valid {
		validRule = append(validRule, validItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "EnclaveHashSet", enclaveHashRule, validRule, serviceRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierEnclaveHashSetIterator{contract: _EspressoNitroTEEVerifier.contract, event: "EnclaveHashSet", logs: logs, sub: sub}, nil
}

// WatchEnclaveHashSet is a free log subscription operation binding the contract event 0x09821cb8037e04ceb9fc83e9a7c52c75b73a08d32a0f2f0f30cc83f2e4426090.
//
// Solidity: event EnclaveHashSet(bytes32 indexed enclaveHash, bool indexed valid, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchEnclaveHashSet(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierEnclaveHashSet, enclaveHash [][32]byte, valid []bool, service []uint8) (event.Subscription, error) {

	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var validRule []interface{}
	for _, validItem := range valid {
		validRule = append(validRule, validItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "EnclaveHashSet", enclaveHashRule, validRule, serviceRule)
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

// ParseEnclaveHashSet is a log parse operation binding the contract event 0x09821cb8037e04ceb9fc83e9a7c52c75b73a08d32a0f2f0f30cc83f2e4426090.
//
// Solidity: event EnclaveHashSet(bytes32 indexed enclaveHash, bool indexed valid, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseEnclaveHashSet(log types.Log) (*EspressoNitroTEEVerifierEnclaveHashSet, error) {
	event := new(EspressoNitroTEEVerifierEnclaveHashSet)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "EnclaveHashSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierGuardianAddedIterator is returned from FilterGuardianAdded and is used to iterate over the raw logs and unpacked data for GuardianAdded events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierGuardianAddedIterator struct {
	Event *EspressoNitroTEEVerifierGuardianAdded // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierGuardianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierGuardianAdded)
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
		it.Event = new(EspressoNitroTEEVerifierGuardianAdded)
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
func (it *EspressoNitroTEEVerifierGuardianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierGuardianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierGuardianAdded represents a GuardianAdded event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierGuardianAdded struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdded is a free log retrieval operation binding the contract event 0x038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969.
//
// Solidity: event GuardianAdded(address indexed guardian)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterGuardianAdded(opts *bind.FilterOpts, guardian []common.Address) (*EspressoNitroTEEVerifierGuardianAddedIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "GuardianAdded", guardianRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierGuardianAddedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "GuardianAdded", logs: logs, sub: sub}, nil
}

// WatchGuardianAdded is a free log subscription operation binding the contract event 0x038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969.
//
// Solidity: event GuardianAdded(address indexed guardian)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchGuardianAdded(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierGuardianAdded, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "GuardianAdded", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierGuardianAdded)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
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

// ParseGuardianAdded is a log parse operation binding the contract event 0x038596bb31e2e7d3d9f184d4c98b310103f6d7f5830e5eec32bffe6f1728f969.
//
// Solidity: event GuardianAdded(address indexed guardian)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseGuardianAdded(log types.Log) (*EspressoNitroTEEVerifierGuardianAdded, error) {
	event := new(EspressoNitroTEEVerifierGuardianAdded)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierGuardianRemovedIterator is returned from FilterGuardianRemoved and is used to iterate over the raw logs and unpacked data for GuardianRemoved events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierGuardianRemovedIterator struct {
	Event *EspressoNitroTEEVerifierGuardianRemoved // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierGuardianRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierGuardianRemoved)
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
		it.Event = new(EspressoNitroTEEVerifierGuardianRemoved)
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
func (it *EspressoNitroTEEVerifierGuardianRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierGuardianRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierGuardianRemoved represents a GuardianRemoved event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierGuardianRemoved struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianRemoved is a free log retrieval operation binding the contract event 0xb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52.
//
// Solidity: event GuardianRemoved(address indexed guardian)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterGuardianRemoved(opts *bind.FilterOpts, guardian []common.Address) (*EspressoNitroTEEVerifierGuardianRemovedIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "GuardianRemoved", guardianRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierGuardianRemovedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "GuardianRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardianRemoved is a free log subscription operation binding the contract event 0xb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52.
//
// Solidity: event GuardianRemoved(address indexed guardian)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchGuardianRemoved(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierGuardianRemoved, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "GuardianRemoved", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierGuardianRemoved)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
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

// ParseGuardianRemoved is a log parse operation binding the contract event 0xb8107d0c6b40be480ce3172ee66ba6d64b71f6b1685a851340036e6e2e3e3c52.
//
// Solidity: event GuardianRemoved(address indexed guardian)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseGuardianRemoved(log types.Log) (*EspressoNitroTEEVerifierGuardianRemoved, error) {
	event := new(EspressoNitroTEEVerifierGuardianRemoved)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierInitializedIterator struct {
	Event *EspressoNitroTEEVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierInitialized)
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
		it.Event = new(EspressoNitroTEEVerifierInitialized)
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
func (it *EspressoNitroTEEVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierInitialized represents a Initialized event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*EspressoNitroTEEVerifierInitializedIterator, error) {

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierInitializedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierInitialized)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseInitialized(log types.Log) (*EspressoNitroTEEVerifierInitialized, error) {
	event := new(EspressoNitroTEEVerifierInitialized)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// EspressoNitroTEEVerifierOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierOwnershipTransferStartedIterator struct {
	Event *EspressoNitroTEEVerifierOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierOwnershipTransferStarted)
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
		it.Event = new(EspressoNitroTEEVerifierOwnershipTransferStarted)
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
func (it *EspressoNitroTEEVerifierOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EspressoNitroTEEVerifierOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierOwnershipTransferStartedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierOwnershipTransferStarted)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseOwnershipTransferStarted(log types.Log) (*EspressoNitroTEEVerifierOwnershipTransferStarted, error) {
	event := new(EspressoNitroTEEVerifierOwnershipTransferStarted)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierOwnershipTransferredIterator struct {
	Event *EspressoNitroTEEVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierOwnershipTransferred)
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
		it.Event = new(EspressoNitroTEEVerifierOwnershipTransferred)
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
func (it *EspressoNitroTEEVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EspressoNitroTEEVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierOwnershipTransferredIterator{contract: _EspressoNitroTEEVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierOwnershipTransferred)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*EspressoNitroTEEVerifierOwnershipTransferred, error) {
	event := new(EspressoNitroTEEVerifierOwnershipTransferred)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierRoleAdminChangedIterator struct {
	Event *EspressoNitroTEEVerifierRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierRoleAdminChanged)
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
		it.Event = new(EspressoNitroTEEVerifierRoleAdminChanged)
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
func (it *EspressoNitroTEEVerifierRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierRoleAdminChanged represents a RoleAdminChanged event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EspressoNitroTEEVerifierRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierRoleAdminChangedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierRoleAdminChanged)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseRoleAdminChanged(log types.Log) (*EspressoNitroTEEVerifierRoleAdminChanged, error) {
	event := new(EspressoNitroTEEVerifierRoleAdminChanged)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierRoleGrantedIterator struct {
	Event *EspressoNitroTEEVerifierRoleGranted // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierRoleGranted)
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
		it.Event = new(EspressoNitroTEEVerifierRoleGranted)
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
func (it *EspressoNitroTEEVerifierRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierRoleGranted represents a RoleGranted event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EspressoNitroTEEVerifierRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierRoleGrantedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierRoleGranted)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseRoleGranted(log types.Log) (*EspressoNitroTEEVerifierRoleGranted, error) {
	event := new(EspressoNitroTEEVerifierRoleGranted)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierRoleRevokedIterator struct {
	Event *EspressoNitroTEEVerifierRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierRoleRevoked)
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
		it.Event = new(EspressoNitroTEEVerifierRoleRevoked)
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
func (it *EspressoNitroTEEVerifierRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierRoleRevoked represents a RoleRevoked event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EspressoNitroTEEVerifierRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierRoleRevokedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierRoleRevoked)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseRoleRevoked(log types.Log) (*EspressoNitroTEEVerifierRoleRevoked, error) {
	event := new(EspressoNitroTEEVerifierRoleRevoked)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
	Service     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceRegistered is a free log retrieval operation binding the contract event 0x0fa700ad17f1b256f57e62054a779bd9fe08e585f8cacb3def28cca47b25cdc1.
//
// Solidity: event ServiceRegistered(address indexed signer, bytes32 indexed enclaveHash, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterServiceRegistered(opts *bind.FilterOpts, signer []common.Address, enclaveHash [][32]byte, service []uint8) (*EspressoNitroTEEVerifierServiceRegisteredIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "ServiceRegistered", signerRule, enclaveHashRule, serviceRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierServiceRegisteredIterator{contract: _EspressoNitroTEEVerifier.contract, event: "ServiceRegistered", logs: logs, sub: sub}, nil
}

// WatchServiceRegistered is a free log subscription operation binding the contract event 0x0fa700ad17f1b256f57e62054a779bd9fe08e585f8cacb3def28cca47b25cdc1.
//
// Solidity: event ServiceRegistered(address indexed signer, bytes32 indexed enclaveHash, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchServiceRegistered(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierServiceRegistered, signer []common.Address, enclaveHash [][32]byte, service []uint8) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var enclaveHashRule []interface{}
	for _, enclaveHashItem := range enclaveHash {
		enclaveHashRule = append(enclaveHashRule, enclaveHashItem)
	}
	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "ServiceRegistered", signerRule, enclaveHashRule, serviceRule)
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

// ParseServiceRegistered is a log parse operation binding the contract event 0x0fa700ad17f1b256f57e62054a779bd9fe08e585f8cacb3def28cca47b25cdc1.
//
// Solidity: event ServiceRegistered(address indexed signer, bytes32 indexed enclaveHash, uint8 indexed service)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseServiceRegistered(log types.Log) (*EspressoNitroTEEVerifierServiceRegistered, error) {
	event := new(EspressoNitroTEEVerifierServiceRegistered)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "ServiceRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EspressoNitroTEEVerifierUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierUpgradedIterator struct {
	Event *EspressoNitroTEEVerifierUpgraded // Event containing the contract specifics and raw log

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
func (it *EspressoNitroTEEVerifierUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EspressoNitroTEEVerifierUpgraded)
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
		it.Event = new(EspressoNitroTEEVerifierUpgraded)
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
func (it *EspressoNitroTEEVerifierUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EspressoNitroTEEVerifierUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EspressoNitroTEEVerifierUpgraded represents a Upgraded event raised by the EspressoNitroTEEVerifier contract.
type EspressoNitroTEEVerifierUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EspressoNitroTEEVerifierUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EspressoNitroTEEVerifierUpgradedIterator{contract: _EspressoNitroTEEVerifier.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EspressoNitroTEEVerifierUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EspressoNitroTEEVerifier.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EspressoNitroTEEVerifierUpgraded)
				if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EspressoNitroTEEVerifier *EspressoNitroTEEVerifierFilterer) ParseUpgraded(log types.Log) (*EspressoNitroTEEVerifierUpgraded, error) {
	event := new(EspressoNitroTEEVerifierUpgraded)
	if err := _EspressoNitroTEEVerifier.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
