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

// OPSuccinctFaultDisputeGameMetaData contains all meta data concerning the OPSuccinctFaultDisputeGame contract.
var OPSuccinctFaultDisputeGameMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_maxChallengeDuration\",\"type\":\"uint64\",\"internalType\":\"Duration\"},{\"name\":\"_maxProveDuration\",\"type\":\"uint64\",\"internalType\":\"Duration\"},{\"name\":\"_disputeGameFactory\",\"type\":\"address\",\"internalType\":\"contractIDisputeGameFactory\"},{\"name\":\"_sp1Verifier\",\"type\":\"address\",\"internalType\":\"contractISP1Verifier\"},{\"name\":\"_rollupConfigHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_aggregationVkey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_rangeVkeyCommitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_challengerBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_anchorStateRegistry\",\"type\":\"address\",\"internalType\":\"contractIAnchorStateRegistry\"},{\"name\":\"_accessManager\",\"type\":\"address\",\"internalType\":\"contractAccessManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accessManager\",\"inputs\":[],\"outputs\":[{\"name\":\"accessManager_\",\"type\":\"address\",\"internalType\":\"contractAccessManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"anchorStateRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"registry_\",\"type\":\"address\",\"internalType\":\"contractIAnchorStateRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondDistributionMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBondDistributionMode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumOPSuccinctFaultDisputeGame.ProposalStatus\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"challengerBond\",\"inputs\":[],\"outputs\":[{\"name\":\"challengerBond_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimCredit\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimData\",\"inputs\":[],\"outputs\":[{\"name\":\"parentIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"counteredBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claim\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumOPSuccinctFaultDisputeGame.ProposalStatus\"},{\"name\":\"deadline\",\"type\":\"uint64\",\"internalType\":\"Timestamp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"closeGame\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createdAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"Timestamp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"credit\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"credit_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"disputeGameFactory_\",\"type\":\"address\",\"internalType\":\"contractIDisputeGameFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extraData\",\"inputs\":[],\"outputs\":[{\"name\":\"extraData_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"gameCreator\",\"inputs\":[],\"outputs\":[{\"name\":\"creator_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"gameData\",\"inputs\":[],\"outputs\":[{\"name\":\"gameType_\",\"type\":\"uint32\",\"internalType\":\"GameType\"},{\"name\":\"rootClaim_\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"extraData_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gameOver\",\"inputs\":[],\"outputs\":[{\"name\":\"gameOver_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gameType\",\"inputs\":[],\"outputs\":[{\"name\":\"gameType_\",\"type\":\"uint32\",\"internalType\":\"GameType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"l1Head\",\"inputs\":[],\"outputs\":[{\"name\":\"l1Head_\",\"type\":\"bytes32\",\"internalType\":\"Hash\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"l2BlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"l2BlockNumber_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"l2SequenceNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"l2SequenceNumber_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"maxChallengeDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"maxChallengeDuration_\",\"type\":\"uint64\",\"internalType\":\"Duration\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxProveDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"maxProveDuration_\",\"type\":\"uint64\",\"internalType\":\"Duration\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"normalModeCredit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"parentIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"parentIndex_\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"prove\",\"inputs\":[{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumOPSuccinctFaultDisputeGame.ProposalStatus\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"refundModeCredit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumGameStatus\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolvedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"Timestamp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rootClaim\",\"inputs\":[],\"outputs\":[{\"name\":\"rootClaim_\",\"type\":\"bytes32\",\"internalType\":\"Claim\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"startingBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"startingBlockNumber_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startingOutputRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"Hash\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startingRootHash\",\"inputs\":[],\"outputs\":[{\"name\":\"startingRootHash_\",\"type\":\"bytes32\",\"internalType\":\"Hash\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"status\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumGameStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wasRespectedGameTypeWhenCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Challenged\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GameClosed\",\"inputs\":[{\"name\":\"bondDistributionMode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBondDistributionMode\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Proved\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resolved\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumGameStatus\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadAuth\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BondTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimAlreadyChallenged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimAlreadyResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GameNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GameNotOver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GameOver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectBondAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectDisputeGameFactory\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondDistributionMode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidParentGame\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProposalStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoCreditToClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ParentGameNotResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedRootClaim\",\"inputs\":[{\"name\":\"rootClaim\",\"type\":\"bytes32\",\"internalType\":\"Claim\"}]}]",
	Bin: "0x6101e0604052348015610010575f5ffd5b50604051612f95380380612f9583398101604081905261002f916100b4565b602a60c0526001600160401b03998a166080529790981660a0526001600160a01b0395861660e05293851661010052610120929092526101405261016052610180529182166101a052166101c052610169565b80516001600160401b0381168114610098575f5ffd5b919050565b6001600160a01b03811681146100b1575f5ffd5b50565b5f5f5f5f5f5f5f5f5f5f6101408b8d0312156100ce575f5ffd5b6100d78b610082565b99506100e560208c01610082565b985060408b01516100f58161009d565b60608c01519098506101068161009d565b809750505f60808c01519050809650505f60a08c01519050809550505f60c08c015190508094505060e08b015192506101008b01516101448161009d565b6101208c01519092506101568161009d565b809150509295989b9194979a5092959850565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c051612d2761026e5f395f818161082c0152818161177a01526123f901525f81816104e3015281816113f8015281816114dd0152818161157501528181611a1001528181611ac901528181611b7d01528181611e29015261228201525f818161058901528181610c5701526124ee01525f610ee901525f610f6701525f610ec301525f610f2b01525f81816107d7015281816116f6015281816118f601526127b801525f818161067d01528181611e020152818161225a015261270601525f81816106af01526125af01525f818161077e0152611fe10152612d275ff3fe608060405260043610610229575f3560e01c806370872aa511610131578063bdb337d1116100ac578063d2ef73981161007c578063f2b4e61711610062578063f2b4e617146107c9578063fa24f743146107fb578063fdcb60681461081e575f5ffd5b8063d2ef7398146107a2578063d5d44d80146107aa575f5ffd5b8063bdb337d114610712578063c0d8bb7414610726578063cf09e0d014610751578063d2177bdd14610770575f5ffd5b80638b85902b11610101578063bbdc02db116100e7578063bbdc02db1461066f578063bcbe5094146106a1578063bcef3b55146106d3575f5ffd5b80638b85902b1461063057806399735e3214610630575f5ffd5b806370872aa5146105ad578063786b844b146105c15780637948690a146105d55780638129fc1c14610628575f5ffd5b80633ec4d4d6116101c15780635c0cba331161019157806360e274641161017757806360e274641461051b5780636361506d1461053c57806368ccdc861461057b575f5ffd5b80635c0cba33146104d5578063609d333414610507575f5ffd5b80633ec4d4d6146103b4578063529d6a8c1461042657806354fd4d501461045157806357da950e146104a6575f5ffd5b80632810e1d6116101fc5780632810e1d6146102f6578063375bfa5d1461030a578063378dd48c1461033657806337b1b22914610354575f5ffd5b806319effeb41461022d578063200d2ed214610276578063250e69bd146102af57806325fc2ace146102d8575b5f5ffd5b348015610238575f5ffd5b505f546102589068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b348015610281575f5ffd5b505f546102a290700100000000000000000000000000000000900460ff1681565b60405161026d9190612998565b3480156102ba575f5ffd5b506009546102c89060ff1681565b604051901515815260200161026d565b3480156102e3575f5ffd5b506007545b60405190815260200161026d565b348015610301575f5ffd5b506102a2610850565b348015610315575f5ffd5b506103296103243660046129ab565b610dc4565b60405161026d9190612a2d565b348015610341575f5ffd5b506009546102a290610100900460ff1681565b34801561035f575f5ffd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560601c5b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161026d565b3480156103bf575f5ffd5b506001546002546003546004546104149363ffffffff81169373ffffffffffffffffffffffffffffffffffffffff64010000000090920482169391169160ff81169067ffffffffffffffff6101009091041686565b60405161026d96959493929190612a3b565b348015610431575f5ffd5b506102e8610440366004612abc565b60056020525f908152604090205481565b34801561045c575f5ffd5b506104996040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161026d9190612b2a565b3480156104b1575f5ffd5b506007546008546104c0919082565b6040805192835260208301919091520161026d565b3480156104e0575f5ffd5b507f000000000000000000000000000000000000000000000000000000000000000061038f565b348015610512575f5ffd5b50610499611154565b348015610526575f5ffd5b5061053a610535366004612abc565b611162565b005b348015610547575f5ffd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003603401356102e8565b348015610586575f5ffd5b507f00000000000000000000000000000000000000000000000000000000000000006102e8565b3480156105b8575f5ffd5b506008546102e8565b3480156105cc575f5ffd5b5061053a611328565b3480156105e0575f5ffd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036074013560e01c5b60405163ffffffff909116815260200161026d565b61053a6116a3565b34801561063b575f5ffd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003605401356102e8565b34801561067a575f5ffd5b507f0000000000000000000000000000000000000000000000000000000000000000610613565b3480156106ac575f5ffd5b507f0000000000000000000000000000000000000000000000000000000000000000610258565b3480156106de575f5ffd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003601401356102e8565b34801561071d575f5ffd5b506102c861233d565b348015610731575f5ffd5b506102e8610740366004612abc565b60066020525f908152604090205481565b34801561075c575f5ffd5b505f546102589067ffffffffffffffff1681565b34801561077b575f5ffd5b507f0000000000000000000000000000000000000000000000000000000000000000610258565b61032961237b565b3480156107b5575f5ffd5b506102e86107c4366004612abc565b61268c565b3480156107d4575f5ffd5b507f000000000000000000000000000000000000000000000000000000000000000061038f565b348015610806575f5ffd5b5061080f612704565b60405161026d93929190612b3c565b348015610829575f5ffd5b507f000000000000000000000000000000000000000000000000000000000000000061038f565b5f805f54700100000000000000000000000000000000900460ff16600281111561087c5761087c612958565b146108b3576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6108bc612764565b90505f8160028111156108d1576108d1612958565b03610908576040517f92c506ae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600181600281111561091c5761091c612958565b03610999575f8054600191907fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff16700100000000000000000000000000000000835b0217905550600154640100000000900473ffffffffffffffffffffffffffffffffffffffff165f908152600560205260409020479055610cec565b6109a161233d565b6109d7576040517f04643c3900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6004805460ff16908111156109ef576109ef612958565b03610a94575f8054600291907fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff16700100000000000000000000000000000000835b02179055504760055f367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560601c5b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2055610cec565b60016004805460ff1690811115610aad57610aad612958565b03610af3575f8054600191907fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff167001000000000000000000000000000000008361095e565b60026004805460ff1690811115610b0c57610b0c612958565b03610b52575f8054600291907fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000083610a31565b60036004805460ff1690811115610b6b57610b6b612958565b03610cba575f80547fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff16700200000000000000000000000000000000179055610bdf7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033560601c90565b60025473ffffffffffffffffffffffffffffffffffffffff918216911603610c2f5760025473ffffffffffffffffffffffffffffffffffffffff165f908152600560205260409020479055610cec565b60025473ffffffffffffffffffffffffffffffffffffffff165f9081526005602052604090207f000000000000000000000000000000000000000000000000000000000000000090819055610c849047612b96565b60055f367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560601c610a69565b6040517f7492a26900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790555f80547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff16680100000000000000004267ffffffffffffffff16021790819055700100000000000000000000000000000000900460ff166002811115610d7e57610d7e612958565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60905f90a250505f54700100000000000000000000000000000000900460ff1690565b5f610dcd61233d565b15610e04576040517fdf469ccb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6040518060e00160405280610e4560347ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c9003013590565b81526007546020820152604001610e89367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036014013590565b90565b8152602001367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036054013581526020017f000000000000000000000000000000000000000000000000000000000000000081526020017f000000000000000000000000000000000000000000000000000000000000000081526020013373ffffffffffffffffffffffffffffffffffffffff1681525090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166341493c607f000000000000000000000000000000000000000000000000000000000000000083604051602001610ff591905f60e082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015273ffffffffffffffffffffffffffffffffffffffff60c08401511660c083015292915050565b60405160208183030381529060405287876040518563ffffffff1660e01b81526004016110259493929190612ba9565b5f6040518083038186803b15801561103b575f5ffd5b505afa15801561104d573d5f5f3e3d5ffd5b5050600280547fffffffffffffffffffffffff000000000000000000000000000000000000000016331790555050600154640100000000900473ffffffffffffffffffffffffffffffffffffffff166110d057600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660021790556110fc565b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660031790555b60025460405173ffffffffffffffffffffffffffffffffffffffff909116907f5e6565d9ca2f5c8501d6418bf563322a7243ba7ace266d75eac99f4adbb30ba7905f90a2505060045460ff165b92915050565b905090565b606061114f60546024612907565b61116a611328565b5f6002600954610100900460ff16600281111561118957611189612958565b036111b9575073ffffffffffffffffffffffffffffffffffffffff81165f90815260066020526040902054611239565b6001600954610100900460ff1660028111156111d7576111d7612958565b03611207575073ffffffffffffffffffffffffffffffffffffffff81165f90815260056020526040902054611239565b6040517f078a3df400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f03611272576040517f17bfe5f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f81815260066020908152604080832083905560059091528082208290555190919083908381818185875af1925050503d805f81146112e3576040519150601f19603f3d011682016040523d82523d5f602084013e6112e8565b606091505b5050905080611323576040517f83e6cc6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b6002600954610100900460ff16600281111561134657611346612958565b148061136d57506001600954610100900460ff16600281111561136b5761136b612958565b145b1561137457565b5f600954610100900460ff16600281111561139157611391612958565b146113c8576040517f078a3df400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f0314d2b30000000000000000000000000000000000000000000000000000000081523060048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690630314d2b390602401602060405180830381865afa158015611452573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114769190612c12565b9050806114af576040517f4851bd9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f17cf21a90000000000000000000000000000000000000000000000000000000081523060048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906317cf21a9906024015f604051808303815f87803b158015611533575f5ffd5b505af1925050508015611544575060015b506040517f496b9c160000000000000000000000000000000000000000000000000000000081523060048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063496b9c1690602401602060405180830381865afa1580156115cf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115f39190612c12565b9050801561162c57600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055611659565b600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166102001790555b7f9908eaac0645df9d0704d06adc9e07337c951de2f06b5f2836151d48d5e4722f600960019054906101000a900460ff166040516116979190612998565b60405180910390a15050565b5f5471010000000000000000000000000000000000900460ff16156116f4576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611763576040517f940d38c700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016631d3225e3367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560601c6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015611834573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118589190612c12565b61188e576040517fd386ef3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607e36146118a357639824bdab5f526004601cfd5b63ffffffff367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036074013560e01c14611dd5575f73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663bb8aa1fc367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036074013560e01c6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815263ffffffff919091166004820152602401606060405180830381865afa1580156119a4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119c89190612c44565b6040517f04e50fed00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301529194507f000000000000000000000000000000000000000000000000000000000000000090911692506304e50fed9150602401602060405180830381865afa158015611a59573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a7d9190612c12565b1580611b3257506040517f34a346ea00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f000000000000000000000000000000000000000000000000000000000000000016906334a346ea90602401602060405180830381865afa158015611b0e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b329190612c12565b80611be657506040517f5958a19300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f00000000000000000000000000000000000000000000000000000000000000001690635958a19390602401602060405180830381865afa158015611bc2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611be69190612c12565b15611c1d576040517f346119f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060400160405280611c988373ffffffffffffffffffffffffffffffffffffffff1663bcef3b556040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c74573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e869190612c97565b81526020018273ffffffffffffffffffffffffffffffffffffffff16638b85902b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ce6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d0a9190612c97565b905280516007556020015160085560018173ffffffffffffffffffffffffffffffffffffffff1663200d2ed26040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d63573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d879190612cae565b6002811115611d9857611d98612958565b03611dcf576040517f346119f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50611ead565b6040517f7258a80700000000000000000000000000000000000000000000000000000000815263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690637258a807906024016040805180830381865afa158015611e82573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ea69190612ccc565b6008556007555b600854367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036054013511611f48576040517ff40239db000000000000000000000000000000000000000000000000000000008152367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360140135600482015260240160405180910390fd5b6040518060c00160405280611f8b60747ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c9003013560e01c90565b63ffffffff1681525f602082018190526040820152606001367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036014013581525f60208201526040016120107f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1642612cee565b67ffffffffffffffff169052805160018054602084015163ffffffff9093167fffffffffffffffff0000000000000000000000000000000000000000000000009091161764010000000073ffffffffffffffffffffffffffffffffffffffff938416021781556040830151600280547fffffffffffffffffffffffff000000000000000000000000000000000000000016919093161790915560608201516003556080820151600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168383838111156120ee576120ee612958565b021790555060a091909101516003909101805467ffffffffffffffff909216610100027fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff9092169190911790555f80547fffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffff167101000000000000000000000000000000000017815534906006906121b07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033560601c90565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546121f79190612cee565b90915550505f80547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff16179055604080517f3c9f397c00000000000000000000000000000000000000000000000000000000815290517f000000000000000000000000000000000000000000000000000000000000000063ffffffff16917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1691633c9f397c916004808201926020929091908290030181865afa1580156122e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123059190612d01565b600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001663ffffffff9290921692909214179055565b6004545f9067ffffffffffffffff42811661010090920416108061114f57505060025473ffffffffffffffffffffffffffffffffffffffff16151590565b5f806004805460ff169081111561239457612394612958565b146123cb576040517f85c345b000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fff59ae7d0000000000000000000000000000000000000000000000000000000081523360048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063ff59ae7d90602401602060405180830381865afa158015612453573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124779190612c12565b6124ad576040517fd386ef3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124b561233d565b156124ec576040517fdf469ccb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000003414612545576040517f8620aa1900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffff0000000000000000000000000000000000000000ffffffff163364010000000002178155600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790556125d567ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001642612cee565b6004805467ffffffffffffffff92909216610100027fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff909216919091179055335f9081526006602052604081208054349290612632908490612cee565b909155505060015460405164010000000090910473ffffffffffffffffffffffffffffffffffffffff16907f98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c905f90a25060045460ff1690565b5f6002600954610100900460ff1660028111156126ab576126ab612958565b036126d8575073ffffffffffffffffffffffffffffffffffffffff165f9081526006602052604090205490565b5073ffffffffffffffffffffffffffffffffffffffff81165f908152600560205260409020545b919050565b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360140135606061275d611154565b9050909192565b5f63ffffffff367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036074013560e01c14612901575f73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663bb8aa1fc367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036074013560e01c6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815263ffffffff919091166004820152602401606060405180830381865afa158015612866573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061288a9190612c44565b925050508073ffffffffffffffffffffffffffffffffffffffff1663200d2ed26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128d7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128fb9190612cae565b91505090565b50600290565b604051818152367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90038284820160208401378260208301015f815260208101604052505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6003811061299557612995612958565b50565b602081016129a583612985565b91905290565b5f5f602083850312156129bc575f5ffd5b823567ffffffffffffffff8111156129d2575f5ffd5b8301601f810185136129e2575f5ffd5b803567ffffffffffffffff8111156129f8575f5ffd5b856020828401011115612a09575f5ffd5b6020919091019590945092505050565b60058110612a2957612a29612958565b9052565b602081016111498284612a19565b63ffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8681166020830152851660408201526060810184905260c08101612a806080830185612a19565b67ffffffffffffffff831660a0830152979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114612995575f5ffd5b5f60208284031215612acc575f5ffd5b8135612ad781612a9b565b9392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f612ad76020830184612ade565b63ffffffff84168152826020820152606060408201525f612b606060830184612ade565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561114957611149612b69565b848152606060208201525f612bc16060830186612ade565b8281036040840152838152838560208301375f6020858301015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f86011682010191505095945050505050565b5f60208284031215612c22575f5ffd5b81518015158114612ad7575f5ffd5b805163ffffffff811681146126ff575f5ffd5b5f5f5f60608486031215612c56575f5ffd5b612c5f84612c31565b9250602084015167ffffffffffffffff81168114612c7b575f5ffd5b6040850151909250612c8c81612a9b565b809150509250925092565b5f60208284031215612ca7575f5ffd5b5051919050565b5f60208284031215612cbe575f5ffd5b815160038110612ad7575f5ffd5b5f5f60408385031215612cdd575f5ffd5b505080516020909101519092909150565b8082018082111561114957611149612b69565b5f60208284031215612d11575f5ffd5b612ad782612c3156fea164736f6c634300081d000a",
}

// OPSuccinctFaultDisputeGameABI is the input ABI used to generate the binding from.
// Deprecated: Use OPSuccinctFaultDisputeGameMetaData.ABI instead.
var OPSuccinctFaultDisputeGameABI = OPSuccinctFaultDisputeGameMetaData.ABI

// OPSuccinctFaultDisputeGameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OPSuccinctFaultDisputeGameMetaData.Bin instead.
var OPSuccinctFaultDisputeGameBin = OPSuccinctFaultDisputeGameMetaData.Bin

// DeployOPSuccinctFaultDisputeGame deploys a new Ethereum contract, binding an instance of OPSuccinctFaultDisputeGame to it.
func DeployOPSuccinctFaultDisputeGame(auth *bind.TransactOpts, backend bind.ContractBackend, _maxChallengeDuration uint64, _maxProveDuration uint64, _disputeGameFactory common.Address, _sp1Verifier common.Address, _rollupConfigHash [32]byte, _aggregationVkey [32]byte, _rangeVkeyCommitment [32]byte, _challengerBond *big.Int, _anchorStateRegistry common.Address, _accessManager common.Address) (common.Address, *types.Transaction, *OPSuccinctFaultDisputeGame, error) {
	parsed, err := OPSuccinctFaultDisputeGameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OPSuccinctFaultDisputeGameBin), backend, _maxChallengeDuration, _maxProveDuration, _disputeGameFactory, _sp1Verifier, _rollupConfigHash, _aggregationVkey, _rangeVkeyCommitment, _challengerBond, _anchorStateRegistry, _accessManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OPSuccinctFaultDisputeGame{OPSuccinctFaultDisputeGameCaller: OPSuccinctFaultDisputeGameCaller{contract: contract}, OPSuccinctFaultDisputeGameTransactor: OPSuccinctFaultDisputeGameTransactor{contract: contract}, OPSuccinctFaultDisputeGameFilterer: OPSuccinctFaultDisputeGameFilterer{contract: contract}}, nil
}

// OPSuccinctFaultDisputeGame is an auto generated Go binding around an Ethereum contract.
type OPSuccinctFaultDisputeGame struct {
	OPSuccinctFaultDisputeGameCaller     // Read-only binding to the contract
	OPSuccinctFaultDisputeGameTransactor // Write-only binding to the contract
	OPSuccinctFaultDisputeGameFilterer   // Log filterer for contract events
}

// OPSuccinctFaultDisputeGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type OPSuccinctFaultDisputeGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPSuccinctFaultDisputeGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPSuccinctFaultDisputeGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPSuccinctFaultDisputeGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPSuccinctFaultDisputeGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPSuccinctFaultDisputeGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPSuccinctFaultDisputeGameSession struct {
	Contract     *OPSuccinctFaultDisputeGame // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OPSuccinctFaultDisputeGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPSuccinctFaultDisputeGameCallerSession struct {
	Contract *OPSuccinctFaultDisputeGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// OPSuccinctFaultDisputeGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPSuccinctFaultDisputeGameTransactorSession struct {
	Contract     *OPSuccinctFaultDisputeGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// OPSuccinctFaultDisputeGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type OPSuccinctFaultDisputeGameRaw struct {
	Contract *OPSuccinctFaultDisputeGame // Generic contract binding to access the raw methods on
}

// OPSuccinctFaultDisputeGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPSuccinctFaultDisputeGameCallerRaw struct {
	Contract *OPSuccinctFaultDisputeGameCaller // Generic read-only contract binding to access the raw methods on
}

// OPSuccinctFaultDisputeGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPSuccinctFaultDisputeGameTransactorRaw struct {
	Contract *OPSuccinctFaultDisputeGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPSuccinctFaultDisputeGame creates a new instance of OPSuccinctFaultDisputeGame, bound to a specific deployed contract.
func NewOPSuccinctFaultDisputeGame(address common.Address, backend bind.ContractBackend) (*OPSuccinctFaultDisputeGame, error) {
	contract, err := bindOPSuccinctFaultDisputeGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGame{OPSuccinctFaultDisputeGameCaller: OPSuccinctFaultDisputeGameCaller{contract: contract}, OPSuccinctFaultDisputeGameTransactor: OPSuccinctFaultDisputeGameTransactor{contract: contract}, OPSuccinctFaultDisputeGameFilterer: OPSuccinctFaultDisputeGameFilterer{contract: contract}}, nil
}

// NewOPSuccinctFaultDisputeGameCaller creates a new read-only instance of OPSuccinctFaultDisputeGame, bound to a specific deployed contract.
func NewOPSuccinctFaultDisputeGameCaller(address common.Address, caller bind.ContractCaller) (*OPSuccinctFaultDisputeGameCaller, error) {
	contract, err := bindOPSuccinctFaultDisputeGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGameCaller{contract: contract}, nil
}

// NewOPSuccinctFaultDisputeGameTransactor creates a new write-only instance of OPSuccinctFaultDisputeGame, bound to a specific deployed contract.
func NewOPSuccinctFaultDisputeGameTransactor(address common.Address, transactor bind.ContractTransactor) (*OPSuccinctFaultDisputeGameTransactor, error) {
	contract, err := bindOPSuccinctFaultDisputeGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGameTransactor{contract: contract}, nil
}

// NewOPSuccinctFaultDisputeGameFilterer creates a new log filterer instance of OPSuccinctFaultDisputeGame, bound to a specific deployed contract.
func NewOPSuccinctFaultDisputeGameFilterer(address common.Address, filterer bind.ContractFilterer) (*OPSuccinctFaultDisputeGameFilterer, error) {
	contract, err := bindOPSuccinctFaultDisputeGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGameFilterer{contract: contract}, nil
}

// bindOPSuccinctFaultDisputeGame binds a generic wrapper to an already deployed contract.
func bindOPSuccinctFaultDisputeGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OPSuccinctFaultDisputeGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPSuccinctFaultDisputeGame.Contract.OPSuccinctFaultDisputeGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.OPSuccinctFaultDisputeGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.OPSuccinctFaultDisputeGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPSuccinctFaultDisputeGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.contract.Transact(opts, method, params...)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address accessManager_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) AccessManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "accessManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address accessManager_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) AccessManager() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.AccessManager(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address accessManager_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) AccessManager() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.AccessManager(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// AnchorStateRegistry is a free data retrieval call binding the contract method 0x5c0cba33.
//
// Solidity: function anchorStateRegistry() view returns(address registry_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) AnchorStateRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "anchorStateRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnchorStateRegistry is a free data retrieval call binding the contract method 0x5c0cba33.
//
// Solidity: function anchorStateRegistry() view returns(address registry_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) AnchorStateRegistry() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.AnchorStateRegistry(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// AnchorStateRegistry is a free data retrieval call binding the contract method 0x5c0cba33.
//
// Solidity: function anchorStateRegistry() view returns(address registry_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) AnchorStateRegistry() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.AnchorStateRegistry(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// BondDistributionMode is a free data retrieval call binding the contract method 0x378dd48c.
//
// Solidity: function bondDistributionMode() view returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) BondDistributionMode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "bondDistributionMode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BondDistributionMode is a free data retrieval call binding the contract method 0x378dd48c.
//
// Solidity: function bondDistributionMode() view returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) BondDistributionMode() (uint8, error) {
	return _OPSuccinctFaultDisputeGame.Contract.BondDistributionMode(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// BondDistributionMode is a free data retrieval call binding the contract method 0x378dd48c.
//
// Solidity: function bondDistributionMode() view returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) BondDistributionMode() (uint8, error) {
	return _OPSuccinctFaultDisputeGame.Contract.BondDistributionMode(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ChallengerBond is a free data retrieval call binding the contract method 0x68ccdc86.
//
// Solidity: function challengerBond() view returns(uint256 challengerBond_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) ChallengerBond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "challengerBond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengerBond is a free data retrieval call binding the contract method 0x68ccdc86.
//
// Solidity: function challengerBond() view returns(uint256 challengerBond_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) ChallengerBond() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ChallengerBond(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ChallengerBond is a free data retrieval call binding the contract method 0x68ccdc86.
//
// Solidity: function challengerBond() view returns(uint256 challengerBond_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) ChallengerBond() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ChallengerBond(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ClaimData is a free data retrieval call binding the contract method 0x3ec4d4d6.
//
// Solidity: function claimData() view returns(uint32 parentIndex, address counteredBy, address prover, bytes32 claim, uint8 status, uint64 deadline)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) ClaimData(opts *bind.CallOpts) (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Prover      common.Address
	Claim       [32]byte
	Status      uint8
	Deadline    uint64
}, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "claimData")

	outstruct := new(struct {
		ParentIndex uint32
		CounteredBy common.Address
		Prover      common.Address
		Claim       [32]byte
		Status      uint8
		Deadline    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParentIndex = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.CounteredBy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Prover = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Claim = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Deadline = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// ClaimData is a free data retrieval call binding the contract method 0x3ec4d4d6.
//
// Solidity: function claimData() view returns(uint32 parentIndex, address counteredBy, address prover, bytes32 claim, uint8 status, uint64 deadline)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) ClaimData() (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Prover      common.Address
	Claim       [32]byte
	Status      uint8
	Deadline    uint64
}, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ClaimData(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ClaimData is a free data retrieval call binding the contract method 0x3ec4d4d6.
//
// Solidity: function claimData() view returns(uint32 parentIndex, address counteredBy, address prover, bytes32 claim, uint8 status, uint64 deadline)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) ClaimData() (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Prover      common.Address
	Claim       [32]byte
	Status      uint8
	Deadline    uint64
}, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ClaimData(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) CreatedAt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) CreatedAt() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.CreatedAt(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) CreatedAt() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.CreatedAt(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// Credit is a free data retrieval call binding the contract method 0xd5d44d80.
//
// Solidity: function credit(address _recipient) view returns(uint256 credit_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) Credit(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "credit", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Credit is a free data retrieval call binding the contract method 0xd5d44d80.
//
// Solidity: function credit(address _recipient) view returns(uint256 credit_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) Credit(_recipient common.Address) (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Credit(&_OPSuccinctFaultDisputeGame.CallOpts, _recipient)
}

// Credit is a free data retrieval call binding the contract method 0xd5d44d80.
//
// Solidity: function credit(address _recipient) view returns(uint256 credit_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) Credit(_recipient common.Address) (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Credit(&_OPSuccinctFaultDisputeGame.CallOpts, _recipient)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address disputeGameFactory_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) DisputeGameFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "disputeGameFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address disputeGameFactory_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) DisputeGameFactory() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.DisputeGameFactory(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address disputeGameFactory_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) DisputeGameFactory() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.DisputeGameFactory(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) ExtraData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "extraData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) ExtraData() ([]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ExtraData(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) ExtraData() ([]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ExtraData(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameCreator is a free data retrieval call binding the contract method 0x37b1b229.
//
// Solidity: function gameCreator() pure returns(address creator_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) GameCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "gameCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameCreator is a free data retrieval call binding the contract method 0x37b1b229.
//
// Solidity: function gameCreator() pure returns(address creator_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) GameCreator() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameCreator(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameCreator is a free data retrieval call binding the contract method 0x37b1b229.
//
// Solidity: function gameCreator() pure returns(address creator_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) GameCreator() (common.Address, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameCreator(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() view returns(uint32 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) GameData(opts *bind.CallOpts) (struct {
	GameType  uint32
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "gameData")

	outstruct := new(struct {
		GameType  uint32
		RootClaim [32]byte
		ExtraData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameType = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.RootClaim = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ExtraData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() view returns(uint32 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) GameData() (struct {
	GameType  uint32
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameData(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() view returns(uint32 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) GameData() (struct {
	GameType  uint32
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameData(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameOver is a free data retrieval call binding the contract method 0xbdb337d1.
//
// Solidity: function gameOver() view returns(bool gameOver_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) GameOver(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "gameOver")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GameOver is a free data retrieval call binding the contract method 0xbdb337d1.
//
// Solidity: function gameOver() view returns(bool gameOver_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) GameOver() (bool, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameOver(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameOver is a free data retrieval call binding the contract method 0xbdb337d1.
//
// Solidity: function gameOver() view returns(bool gameOver_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) GameOver() (bool, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameOver(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() view returns(uint32 gameType_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) GameType(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "gameType")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() view returns(uint32 gameType_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) GameType() (uint32, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameType(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() view returns(uint32 gameType_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) GameType() (uint32, error) {
	return _OPSuccinctFaultDisputeGame.Contract.GameType(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// L1Head is a free data retrieval call binding the contract method 0x6361506d.
//
// Solidity: function l1Head() pure returns(bytes32 l1Head_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) L1Head(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "l1Head")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1Head is a free data retrieval call binding the contract method 0x6361506d.
//
// Solidity: function l1Head() pure returns(bytes32 l1Head_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) L1Head() ([32]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.L1Head(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// L1Head is a free data retrieval call binding the contract method 0x6361506d.
//
// Solidity: function l1Head() pure returns(bytes32 l1Head_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) L1Head() ([32]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.L1Head(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) L2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "l2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) L2BlockNumber() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.L2BlockNumber(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) L2BlockNumber() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.L2BlockNumber(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// L2SequenceNumber is a free data retrieval call binding the contract method 0x99735e32.
//
// Solidity: function l2SequenceNumber() pure returns(uint256 l2SequenceNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) L2SequenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "l2SequenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2SequenceNumber is a free data retrieval call binding the contract method 0x99735e32.
//
// Solidity: function l2SequenceNumber() pure returns(uint256 l2SequenceNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) L2SequenceNumber() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.L2SequenceNumber(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// L2SequenceNumber is a free data retrieval call binding the contract method 0x99735e32.
//
// Solidity: function l2SequenceNumber() pure returns(uint256 l2SequenceNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) L2SequenceNumber() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.L2SequenceNumber(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// MaxChallengeDuration is a free data retrieval call binding the contract method 0xd2177bdd.
//
// Solidity: function maxChallengeDuration() view returns(uint64 maxChallengeDuration_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) MaxChallengeDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "maxChallengeDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxChallengeDuration is a free data retrieval call binding the contract method 0xd2177bdd.
//
// Solidity: function maxChallengeDuration() view returns(uint64 maxChallengeDuration_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) MaxChallengeDuration() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.MaxChallengeDuration(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// MaxChallengeDuration is a free data retrieval call binding the contract method 0xd2177bdd.
//
// Solidity: function maxChallengeDuration() view returns(uint64 maxChallengeDuration_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) MaxChallengeDuration() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.MaxChallengeDuration(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// MaxProveDuration is a free data retrieval call binding the contract method 0xbcbe5094.
//
// Solidity: function maxProveDuration() view returns(uint64 maxProveDuration_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) MaxProveDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "maxProveDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxProveDuration is a free data retrieval call binding the contract method 0xbcbe5094.
//
// Solidity: function maxProveDuration() view returns(uint64 maxProveDuration_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) MaxProveDuration() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.MaxProveDuration(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// MaxProveDuration is a free data retrieval call binding the contract method 0xbcbe5094.
//
// Solidity: function maxProveDuration() view returns(uint64 maxProveDuration_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) MaxProveDuration() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.MaxProveDuration(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// NormalModeCredit is a free data retrieval call binding the contract method 0x529d6a8c.
//
// Solidity: function normalModeCredit(address ) view returns(uint256)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) NormalModeCredit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "normalModeCredit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalModeCredit is a free data retrieval call binding the contract method 0x529d6a8c.
//
// Solidity: function normalModeCredit(address ) view returns(uint256)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) NormalModeCredit(arg0 common.Address) (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.NormalModeCredit(&_OPSuccinctFaultDisputeGame.CallOpts, arg0)
}

// NormalModeCredit is a free data retrieval call binding the contract method 0x529d6a8c.
//
// Solidity: function normalModeCredit(address ) view returns(uint256)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) NormalModeCredit(arg0 common.Address) (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.NormalModeCredit(&_OPSuccinctFaultDisputeGame.CallOpts, arg0)
}

// ParentIndex is a free data retrieval call binding the contract method 0x7948690a.
//
// Solidity: function parentIndex() pure returns(uint32 parentIndex_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) ParentIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "parentIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ParentIndex is a free data retrieval call binding the contract method 0x7948690a.
//
// Solidity: function parentIndex() pure returns(uint32 parentIndex_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) ParentIndex() (uint32, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ParentIndex(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ParentIndex is a free data retrieval call binding the contract method 0x7948690a.
//
// Solidity: function parentIndex() pure returns(uint32 parentIndex_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) ParentIndex() (uint32, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ParentIndex(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// RefundModeCredit is a free data retrieval call binding the contract method 0xc0d8bb74.
//
// Solidity: function refundModeCredit(address ) view returns(uint256)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) RefundModeCredit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "refundModeCredit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundModeCredit is a free data retrieval call binding the contract method 0xc0d8bb74.
//
// Solidity: function refundModeCredit(address ) view returns(uint256)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) RefundModeCredit(arg0 common.Address) (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.RefundModeCredit(&_OPSuccinctFaultDisputeGame.CallOpts, arg0)
}

// RefundModeCredit is a free data retrieval call binding the contract method 0xc0d8bb74.
//
// Solidity: function refundModeCredit(address ) view returns(uint256)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) RefundModeCredit(arg0 common.Address) (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.RefundModeCredit(&_OPSuccinctFaultDisputeGame.CallOpts, arg0)
}

// ResolvedAt is a free data retrieval call binding the contract method 0x19effeb4.
//
// Solidity: function resolvedAt() view returns(uint64)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) ResolvedAt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "resolvedAt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ResolvedAt is a free data retrieval call binding the contract method 0x19effeb4.
//
// Solidity: function resolvedAt() view returns(uint64)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) ResolvedAt() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ResolvedAt(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// ResolvedAt is a free data retrieval call binding the contract method 0x19effeb4.
//
// Solidity: function resolvedAt() view returns(uint64)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) ResolvedAt() (uint64, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ResolvedAt(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) RootClaim(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "rootClaim")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) RootClaim() ([32]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.RootClaim(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) RootClaim() ([32]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.RootClaim(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256 startingBlockNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256 startingBlockNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) StartingBlockNumber() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.StartingBlockNumber(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256 startingBlockNumber_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _OPSuccinctFaultDisputeGame.Contract.StartingBlockNumber(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// StartingOutputRoot is a free data retrieval call binding the contract method 0x57da950e.
//
// Solidity: function startingOutputRoot() view returns(bytes32 root, uint256 l2BlockNumber)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) StartingOutputRoot(opts *bind.CallOpts) (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "startingOutputRoot")

	outstruct := new(struct {
		Root          [32]byte
		L2BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.L2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StartingOutputRoot is a free data retrieval call binding the contract method 0x57da950e.
//
// Solidity: function startingOutputRoot() view returns(bytes32 root, uint256 l2BlockNumber)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) StartingOutputRoot() (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _OPSuccinctFaultDisputeGame.Contract.StartingOutputRoot(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// StartingOutputRoot is a free data retrieval call binding the contract method 0x57da950e.
//
// Solidity: function startingOutputRoot() view returns(bytes32 root, uint256 l2BlockNumber)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) StartingOutputRoot() (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _OPSuccinctFaultDisputeGame.Contract.StartingOutputRoot(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// StartingRootHash is a free data retrieval call binding the contract method 0x25fc2ace.
//
// Solidity: function startingRootHash() view returns(bytes32 startingRootHash_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) StartingRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "startingRootHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StartingRootHash is a free data retrieval call binding the contract method 0x25fc2ace.
//
// Solidity: function startingRootHash() view returns(bytes32 startingRootHash_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) StartingRootHash() ([32]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.StartingRootHash(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// StartingRootHash is a free data retrieval call binding the contract method 0x25fc2ace.
//
// Solidity: function startingRootHash() view returns(bytes32 startingRootHash_)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) StartingRootHash() ([32]byte, error) {
	return _OPSuccinctFaultDisputeGame.Contract.StartingRootHash(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) Status() (uint8, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Status(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) Status() (uint8, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Status(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) Version() (string, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Version(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) Version() (string, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Version(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// WasRespectedGameTypeWhenCreated is a free data retrieval call binding the contract method 0x250e69bd.
//
// Solidity: function wasRespectedGameTypeWhenCreated() view returns(bool)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCaller) WasRespectedGameTypeWhenCreated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OPSuccinctFaultDisputeGame.contract.Call(opts, &out, "wasRespectedGameTypeWhenCreated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasRespectedGameTypeWhenCreated is a free data retrieval call binding the contract method 0x250e69bd.
//
// Solidity: function wasRespectedGameTypeWhenCreated() view returns(bool)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) WasRespectedGameTypeWhenCreated() (bool, error) {
	return _OPSuccinctFaultDisputeGame.Contract.WasRespectedGameTypeWhenCreated(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// WasRespectedGameTypeWhenCreated is a free data retrieval call binding the contract method 0x250e69bd.
//
// Solidity: function wasRespectedGameTypeWhenCreated() view returns(bool)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameCallerSession) WasRespectedGameTypeWhenCreated() (bool, error) {
	return _OPSuccinctFaultDisputeGame.Contract.WasRespectedGameTypeWhenCreated(&_OPSuccinctFaultDisputeGame.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() payable returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactor) Challenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.contract.Transact(opts, "challenge")
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() payable returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) Challenge() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Challenge(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() payable returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorSession) Challenge() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Challenge(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// ClaimCredit is a paid mutator transaction binding the contract method 0x60e27464.
//
// Solidity: function claimCredit(address _recipient) returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactor) ClaimCredit(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.contract.Transact(opts, "claimCredit", _recipient)
}

// ClaimCredit is a paid mutator transaction binding the contract method 0x60e27464.
//
// Solidity: function claimCredit(address _recipient) returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) ClaimCredit(_recipient common.Address) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ClaimCredit(&_OPSuccinctFaultDisputeGame.TransactOpts, _recipient)
}

// ClaimCredit is a paid mutator transaction binding the contract method 0x60e27464.
//
// Solidity: function claimCredit(address _recipient) returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorSession) ClaimCredit(_recipient common.Address) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.ClaimCredit(&_OPSuccinctFaultDisputeGame.TransactOpts, _recipient)
}

// CloseGame is a paid mutator transaction binding the contract method 0x786b844b.
//
// Solidity: function closeGame() returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactor) CloseGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.contract.Transact(opts, "closeGame")
}

// CloseGame is a paid mutator transaction binding the contract method 0x786b844b.
//
// Solidity: function closeGame() returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) CloseGame() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.CloseGame(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// CloseGame is a paid mutator transaction binding the contract method 0x786b844b.
//
// Solidity: function closeGame() returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorSession) CloseGame() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.CloseGame(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) Initialize() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Initialize(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorSession) Initialize() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Initialize(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// Prove is a paid mutator transaction binding the contract method 0x375bfa5d.
//
// Solidity: function prove(bytes proofBytes) returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactor) Prove(opts *bind.TransactOpts, proofBytes []byte) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.contract.Transact(opts, "prove", proofBytes)
}

// Prove is a paid mutator transaction binding the contract method 0x375bfa5d.
//
// Solidity: function prove(bytes proofBytes) returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) Prove(proofBytes []byte) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Prove(&_OPSuccinctFaultDisputeGame.TransactOpts, proofBytes)
}

// Prove is a paid mutator transaction binding the contract method 0x375bfa5d.
//
// Solidity: function prove(bytes proofBytes) returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorSession) Prove(proofBytes []byte) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Prove(&_OPSuccinctFaultDisputeGame.TransactOpts, proofBytes)
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactor) Resolve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.contract.Transact(opts, "resolve")
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameSession) Resolve() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Resolve(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameTransactorSession) Resolve() (*types.Transaction, error) {
	return _OPSuccinctFaultDisputeGame.Contract.Resolve(&_OPSuccinctFaultDisputeGame.TransactOpts)
}

// OPSuccinctFaultDisputeGameChallengedIterator is returned from FilterChallenged and is used to iterate over the raw logs and unpacked data for Challenged events raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameChallengedIterator struct {
	Event *OPSuccinctFaultDisputeGameChallenged // Event containing the contract specifics and raw log

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
func (it *OPSuccinctFaultDisputeGameChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPSuccinctFaultDisputeGameChallenged)
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
		it.Event = new(OPSuccinctFaultDisputeGameChallenged)
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
func (it *OPSuccinctFaultDisputeGameChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPSuccinctFaultDisputeGameChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPSuccinctFaultDisputeGameChallenged represents a Challenged event raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameChallenged struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallenged is a free log retrieval operation binding the contract event 0x98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c.
//
// Solidity: event Challenged(address indexed challenger)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) FilterChallenged(opts *bind.FilterOpts, challenger []common.Address) (*OPSuccinctFaultDisputeGameChallengedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.FilterLogs(opts, "Challenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGameChallengedIterator{contract: _OPSuccinctFaultDisputeGame.contract, event: "Challenged", logs: logs, sub: sub}, nil
}

// WatchChallenged is a free log subscription operation binding the contract event 0x98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c.
//
// Solidity: event Challenged(address indexed challenger)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) WatchChallenged(opts *bind.WatchOpts, sink chan<- *OPSuccinctFaultDisputeGameChallenged, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.WatchLogs(opts, "Challenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPSuccinctFaultDisputeGameChallenged)
				if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "Challenged", log); err != nil {
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

// ParseChallenged is a log parse operation binding the contract event 0x98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c.
//
// Solidity: event Challenged(address indexed challenger)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) ParseChallenged(log types.Log) (*OPSuccinctFaultDisputeGameChallenged, error) {
	event := new(OPSuccinctFaultDisputeGameChallenged)
	if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "Challenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPSuccinctFaultDisputeGameGameClosedIterator is returned from FilterGameClosed and is used to iterate over the raw logs and unpacked data for GameClosed events raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameGameClosedIterator struct {
	Event *OPSuccinctFaultDisputeGameGameClosed // Event containing the contract specifics and raw log

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
func (it *OPSuccinctFaultDisputeGameGameClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPSuccinctFaultDisputeGameGameClosed)
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
		it.Event = new(OPSuccinctFaultDisputeGameGameClosed)
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
func (it *OPSuccinctFaultDisputeGameGameClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPSuccinctFaultDisputeGameGameClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPSuccinctFaultDisputeGameGameClosed represents a GameClosed event raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameGameClosed struct {
	BondDistributionMode uint8
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGameClosed is a free log retrieval operation binding the contract event 0x9908eaac0645df9d0704d06adc9e07337c951de2f06b5f2836151d48d5e4722f.
//
// Solidity: event GameClosed(uint8 bondDistributionMode)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) FilterGameClosed(opts *bind.FilterOpts) (*OPSuccinctFaultDisputeGameGameClosedIterator, error) {

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.FilterLogs(opts, "GameClosed")
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGameGameClosedIterator{contract: _OPSuccinctFaultDisputeGame.contract, event: "GameClosed", logs: logs, sub: sub}, nil
}

// WatchGameClosed is a free log subscription operation binding the contract event 0x9908eaac0645df9d0704d06adc9e07337c951de2f06b5f2836151d48d5e4722f.
//
// Solidity: event GameClosed(uint8 bondDistributionMode)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) WatchGameClosed(opts *bind.WatchOpts, sink chan<- *OPSuccinctFaultDisputeGameGameClosed) (event.Subscription, error) {

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.WatchLogs(opts, "GameClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPSuccinctFaultDisputeGameGameClosed)
				if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "GameClosed", log); err != nil {
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

// ParseGameClosed is a log parse operation binding the contract event 0x9908eaac0645df9d0704d06adc9e07337c951de2f06b5f2836151d48d5e4722f.
//
// Solidity: event GameClosed(uint8 bondDistributionMode)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) ParseGameClosed(log types.Log) (*OPSuccinctFaultDisputeGameGameClosed, error) {
	event := new(OPSuccinctFaultDisputeGameGameClosed)
	if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "GameClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPSuccinctFaultDisputeGameProvedIterator is returned from FilterProved and is used to iterate over the raw logs and unpacked data for Proved events raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameProvedIterator struct {
	Event *OPSuccinctFaultDisputeGameProved // Event containing the contract specifics and raw log

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
func (it *OPSuccinctFaultDisputeGameProvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPSuccinctFaultDisputeGameProved)
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
		it.Event = new(OPSuccinctFaultDisputeGameProved)
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
func (it *OPSuccinctFaultDisputeGameProvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPSuccinctFaultDisputeGameProvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPSuccinctFaultDisputeGameProved represents a Proved event raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameProved struct {
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProved is a free log retrieval operation binding the contract event 0x5e6565d9ca2f5c8501d6418bf563322a7243ba7ace266d75eac99f4adbb30ba7.
//
// Solidity: event Proved(address indexed prover)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) FilterProved(opts *bind.FilterOpts, prover []common.Address) (*OPSuccinctFaultDisputeGameProvedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.FilterLogs(opts, "Proved", proverRule)
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGameProvedIterator{contract: _OPSuccinctFaultDisputeGame.contract, event: "Proved", logs: logs, sub: sub}, nil
}

// WatchProved is a free log subscription operation binding the contract event 0x5e6565d9ca2f5c8501d6418bf563322a7243ba7ace266d75eac99f4adbb30ba7.
//
// Solidity: event Proved(address indexed prover)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) WatchProved(opts *bind.WatchOpts, sink chan<- *OPSuccinctFaultDisputeGameProved, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.WatchLogs(opts, "Proved", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPSuccinctFaultDisputeGameProved)
				if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "Proved", log); err != nil {
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

// ParseProved is a log parse operation binding the contract event 0x5e6565d9ca2f5c8501d6418bf563322a7243ba7ace266d75eac99f4adbb30ba7.
//
// Solidity: event Proved(address indexed prover)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) ParseProved(log types.Log) (*OPSuccinctFaultDisputeGameProved, error) {
	event := new(OPSuccinctFaultDisputeGameProved)
	if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "Proved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPSuccinctFaultDisputeGameResolvedIterator is returned from FilterResolved and is used to iterate over the raw logs and unpacked data for Resolved events raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameResolvedIterator struct {
	Event *OPSuccinctFaultDisputeGameResolved // Event containing the contract specifics and raw log

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
func (it *OPSuccinctFaultDisputeGameResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPSuccinctFaultDisputeGameResolved)
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
		it.Event = new(OPSuccinctFaultDisputeGameResolved)
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
func (it *OPSuccinctFaultDisputeGameResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPSuccinctFaultDisputeGameResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPSuccinctFaultDisputeGameResolved represents a Resolved event raised by the OPSuccinctFaultDisputeGame contract.
type OPSuccinctFaultDisputeGameResolved struct {
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResolved is a free log retrieval operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) FilterResolved(opts *bind.FilterOpts, status []uint8) (*OPSuccinctFaultDisputeGameResolvedIterator, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.FilterLogs(opts, "Resolved", statusRule)
	if err != nil {
		return nil, err
	}
	return &OPSuccinctFaultDisputeGameResolvedIterator{contract: _OPSuccinctFaultDisputeGame.contract, event: "Resolved", logs: logs, sub: sub}, nil
}

// WatchResolved is a free log subscription operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) WatchResolved(opts *bind.WatchOpts, sink chan<- *OPSuccinctFaultDisputeGameResolved, status []uint8) (event.Subscription, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _OPSuccinctFaultDisputeGame.contract.WatchLogs(opts, "Resolved", statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPSuccinctFaultDisputeGameResolved)
				if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "Resolved", log); err != nil {
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

// ParseResolved is a log parse operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_OPSuccinctFaultDisputeGame *OPSuccinctFaultDisputeGameFilterer) ParseResolved(log types.Log) (*OPSuccinctFaultDisputeGameResolved, error) {
	event := new(OPSuccinctFaultDisputeGameResolved)
	if err := _OPSuccinctFaultDisputeGame.contract.UnpackLog(event, "Resolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
