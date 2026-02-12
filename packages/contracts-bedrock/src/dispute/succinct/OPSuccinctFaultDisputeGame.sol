// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Libraries
import { Clone } from "@solady/utils/Clone.sol";
import {
    BondDistributionMode,
    Claim,
    Clock,
    Duration,
    GameStatus,
    GameType,
    Hash,
    LibClock,
    Timestamp
} from "src/dispute/lib/Types.sol";
import {
    AlreadyInitialized,
    AnchorRootNotFound,
    BadAuth,
    BondTransferFailed,
    ClaimAlreadyResolved,
    ClockTimeExceeded,
    GameNotFinalized,
    GameNotInProgress,
    IncorrectBondAmount,
    InvalidBondDistributionMode,
    NoCreditToClaim,
    UnexpectedRootClaim
} from "src/dispute/lib/Errors.sol";
import {
    ClaimAlreadyChallenged,
    GameNotOver,
    GameOver,
    IncorrectDisputeGameFactory,
    InvalidParentGame,
    InvalidProposalStatus,
    ParentGameNotResolved
} from "src/dispute/succinct/lib/Errors.sol";
import { AggregationOutputs } from "src/dispute/succinct/lib/Types.sol";

// Interfaces
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IDisputeGameFactory } from "interfaces/dispute/IDisputeGameFactory.sol";
import { IDisputeGame } from "interfaces/dispute/IDisputeGame.sol";
import { ISP1Verifier } from "src/dispute/succinct/ISP1Verifier.sol";
import { IAnchorStateRegistry } from "interfaces/dispute/IAnchorStateRegistry.sol";

// Contracts
import { AccessManager } from "src/dispute/succinct/AccessManager.sol";

/// @notice Represents an output root and the L2 block number at which it was generated.
/// @custom:field root The output root hash.
/// @custom:field l2BlockNumber The L2 block number at which the output root was generated.
struct OutputRoot {
    Hash root;
    uint256 l2BlockNumber;
}

/// @title OPSuccinctFaultDisputeGame
/// @notice An implementation of the `IFaultDisputeGame` interface using ZK proofs.
contract OPSuccinctFaultDisputeGame is Clone, ISemver, IDisputeGame {
    ////////////////////////////////////////////////////////////////
    //                         Enums                              //
    ////////////////////////////////////////////////////////////////

    enum ProposalStatus {
        // The initial state of a new proposal.
        Unchallenged,
        // A proposal that has been challenged but not yet proven.
        Challenged,
        // An unchallenged proposal that has been proven valid with a verified proof.
        UnchallengedAndValidProofProvided,
        // A challenged proposal that has been proven valid with a verified proof.
        ChallengedAndValidProofProvided,
        // The final state after resolution, either GameStatus.CHALLENGER_WINS or GameStatus.DEFENDER_WINS.
        Resolved
    }

    ////////////////////////////////////////////////////////////////
    //                         Structs                            //
    ////////////////////////////////////////////////////////////////

    /// @notice The `ClaimData` struct represents the data associated with a Claim.
    struct ClaimData {
        uint32 parentIndex;
        address counteredBy;
        address prover;
        Claim claim;
        ProposalStatus status;
        Timestamp deadline;
    }

    ////////////////////////////////////////////////////////////////
    //                         Events                             //
    ////////////////////////////////////////////////////////////////

    /// @notice Emitted when the game is challenged.
    /// @param challenger The address of the challenger.
    event Challenged(address indexed challenger);

    /// @notice Emitted when the game is proved.
    /// @param prover The address of the prover.
    event Proved(address indexed prover);

    /// @notice Emitted when the game is closed.
    event GameClosed(BondDistributionMode bondDistributionMode);

    ////////////////////////////////////////////////////////////////
    //                         State Vars                         //
    ////////////////////////////////////////////////////////////////

    /// @notice The maximum duration allowed for a challenger to challenge a game.
    Duration internal immutable MAX_CHALLENGE_DURATION;

    /// @notice The maximum duration allowed for a proposer to prove against a challenge.
    Duration internal immutable MAX_PROVE_DURATION;

    /// @notice The game type ID.
    GameType internal immutable GAME_TYPE;

    /// @notice The dispute game factory.
    IDisputeGameFactory internal immutable DISPUTE_GAME_FACTORY;

    /// @notice The SP1 verifier.
    ISP1Verifier internal immutable SP1_VERIFIER;

    /// @notice The rollup config hash.
    bytes32 internal immutable ROLLUP_CONFIG_HASH;

    /// @notice The vkey for the aggregation program.
    bytes32 internal immutable AGGREGATION_VKEY;

    /// @notice The 32 byte commitment to the BabyBear representation of the verification key of the range SP1 program.
    bytes32 internal immutable RANGE_VKEY_COMMITMENT;

    /// @notice The challenger bond for the game.
    uint256 internal immutable CHALLENGER_BOND;

    /// @notice The anchor state registry.
    IAnchorStateRegistry internal immutable ANCHOR_STATE_REGISTRY;

    /// @notice The access manager.
    AccessManager internal immutable ACCESS_MANAGER;

    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice The starting timestamp of the game.
    Timestamp public createdAt;

    /// @notice The timestamp of the game's global resolution.
    Timestamp public resolvedAt;

    /// @notice The current status of the game.
    GameStatus public status;

    /// @notice Flag for the `initialize` function to prevent re-initialization.
    bool internal initialized;

    /// @notice The claim made by the proposer.
    ClaimData public claimData;

    /// @notice Credited balances for winning participants.
    mapping(address => uint256) public normalModeCredit;

    /// @notice A mapping of each claimant's refund mode credit.
    mapping(address => uint256) public refundModeCredit;

    /// @notice The starting output root of the game that is proven from in case of a challenge.
    OutputRoot public startingOutputRoot;

    /// @notice A boolean for whether or not the game type was respected when the game was created.
    bool public wasRespectedGameTypeWhenCreated;

    /// @notice The bond distribution mode of the game.
    BondDistributionMode public bondDistributionMode;

    /// @param _maxChallengeDuration The maximum duration allowed for a challenger to challenge a game.
    /// @param _maxProveDuration The maximum duration allowed for a proposer to prove against a challenge.
    /// @param _disputeGameFactory The factory that creates the dispute games.
    /// @param _sp1Verifier The address of the SP1 verifier.
    /// @param _rollupConfigHash The rollup config hash for the L2 network.
    /// @param _aggregationVkey The vkey for the aggregation program.
    /// @param _rangeVkeyCommitment The commitment to the range vkey.
    /// @param _challengerBond The bond amount that must be submitted by the challenger.
    /// @param _anchorStateRegistry The anchor state registry for the L2 network.
    /// @param _accessManager The access manager for proposer/challenger permissions.
    constructor(
        Duration _maxChallengeDuration,
        Duration _maxProveDuration,
        IDisputeGameFactory _disputeGameFactory,
        ISP1Verifier _sp1Verifier,
        bytes32 _rollupConfigHash,
        bytes32 _aggregationVkey,
        bytes32 _rangeVkeyCommitment,
        uint256 _challengerBond,
        IAnchorStateRegistry _anchorStateRegistry,
        AccessManager _accessManager
    ) {
        // Set up initial game state.
        GAME_TYPE = GameType.wrap(42);
        MAX_CHALLENGE_DURATION = _maxChallengeDuration;
        MAX_PROVE_DURATION = _maxProveDuration;
        DISPUTE_GAME_FACTORY = _disputeGameFactory;
        SP1_VERIFIER = _sp1Verifier;
        ROLLUP_CONFIG_HASH = _rollupConfigHash;
        AGGREGATION_VKEY = _aggregationVkey;
        RANGE_VKEY_COMMITMENT = _rangeVkeyCommitment;
        CHALLENGER_BOND = _challengerBond;
        ANCHOR_STATE_REGISTRY = _anchorStateRegistry;
        ACCESS_MANAGER = _accessManager;
    }

    /// @notice Initializes the contract.
    /// @dev This function may only be called once.
    function initialize() external payable virtual {
        // INVARIANT: The game must not have already been initialized.
        if (initialized) revert AlreadyInitialized();

        // INVARIANT: The game can only be initialized by the dispute game factory.
        if (address(DISPUTE_GAME_FACTORY) != msg.sender) {
            revert IncorrectDisputeGameFactory();
        }

        // INVARIANT: The proposer must be whitelisted.
        if (!ACCESS_MANAGER.isAllowedProposer(gameCreator())) revert BadAuth();

        // Revert if the calldata size is not the expected length.
        assembly {
            if iszero(eq(calldatasize(), 0x7E)) {
                // Store the selector for `BadExtraData()` & revert
                mstore(0x00, 0x9824bdab)
                revert(0x1C, 0x04)
            }
        }

        // The first game is initialized with a parent index of uint32.max
        if (parentIndex() != type(uint32).max) {
            // For subsequent games, get the parent game's information
            (,, IDisputeGame proxy) = DISPUTE_GAME_FACTORY.gameAtIndex(parentIndex());

            if (
                !ANCHOR_STATE_REGISTRY.isGameRespected(proxy) || ANCHOR_STATE_REGISTRY.isGameBlacklisted(proxy)
                    || ANCHOR_STATE_REGISTRY.isGameRetired(proxy)
            ) {
                revert InvalidParentGame();
            }

            startingOutputRoot = OutputRoot({
                l2BlockNumber: OPSuccinctFaultDisputeGame(address(proxy)).l2BlockNumber(),
                root: Hash.wrap(OPSuccinctFaultDisputeGame(address(proxy)).rootClaim().raw())
            });

            // INVARIANT: The parent game must be a valid game.
            if (proxy.status() == GameStatus.CHALLENGER_WINS) {
                revert InvalidParentGame();
            }
        } else {
            // When there is no parent game, the starting output root is the anchor state for the game type.
            (startingOutputRoot.root, startingOutputRoot.l2BlockNumber) =
                IAnchorStateRegistry(ANCHOR_STATE_REGISTRY).anchors(GAME_TYPE);
        }

        // Do not allow the game to be initialized if the root claim corresponds to a block at or before the
        // configured starting block number.
        if (l2BlockNumber() <= startingOutputRoot.l2BlockNumber) {
            revert UnexpectedRootClaim(rootClaim());
        }

        // Set the root claim
        claimData = ClaimData({
            parentIndex: parentIndex(),
            counteredBy: address(0),
            prover: address(0),
            claim: rootClaim(),
            status: ProposalStatus.Unchallenged,
            deadline: Timestamp.wrap(uint64(block.timestamp + MAX_CHALLENGE_DURATION.raw()))
        });

        // Set the game as initialized.
        initialized = true;

        // Deposit the bond.
        refundModeCredit[gameCreator()] += msg.value;

        // Set the game's starting timestamp
        createdAt = Timestamp.wrap(uint64(block.timestamp));

        // Set whether the game type was respected when the game was created.
        wasRespectedGameTypeWhenCreated =
            GameType.unwrap(ANCHOR_STATE_REGISTRY.respectedGameType()) == GameType.unwrap(GAME_TYPE);
    }

    /// @notice The L2 block number for which this game is proposing an output root.
    function l2BlockNumber() public pure returns (uint256 l2BlockNumber_) {
        l2BlockNumber_ = _getArgUint256(0x54);
    }

    /// @notice The L2 sequence number (block number) for which this game is proposing an output root.
    /// @dev Required by IDisputeGame interface. Returns the same value as l2BlockNumber().
    function l2SequenceNumber() public pure returns (uint256 l2SequenceNumber_) {
        l2SequenceNumber_ = _getArgUint256(0x54);
    }

    /// @notice The parent index of the game.
    function parentIndex() public pure returns (uint32 parentIndex_) {
        parentIndex_ = _getArgUint32(0x74);
    }

    /// @notice Only the starting block number of the game.
    function startingBlockNumber() external view returns (uint256 startingBlockNumber_) {
        startingBlockNumber_ = startingOutputRoot.l2BlockNumber;
    }

    /// @notice Starting output root of the game.
    function startingRootHash() external view returns (Hash startingRootHash_) {
        startingRootHash_ = startingOutputRoot.root;
    }

    ////////////////////////////////////////////////////////////////
    //                    `IDisputeGame` impl                     //
    ////////////////////////////////////////////////////////////////

    /// @notice Challenges the game.
    function challenge() external payable returns (ProposalStatus) {
        // INVARIANT: Can only challenge a game that has not been challenged yet.
        if (claimData.status != ProposalStatus.Unchallenged) {
            revert ClaimAlreadyChallenged();
        }

        // INVARIANT: The challenger must be whitelisted.
        if (!ACCESS_MANAGER.isAllowedChallenger(msg.sender)) revert BadAuth();

        // INVARIANT: Cannot challenge if the game is over.
        if (gameOver()) revert GameOver();

        // If the required bond is not met, revert.
        if (msg.value != CHALLENGER_BOND) revert IncorrectBondAmount();

        // Update the counteredBy address
        claimData.counteredBy = msg.sender;

        // Update the status of the proposal
        claimData.status = ProposalStatus.Challenged;

        // Update the clock to the current block timestamp, which marks the start of the challenge.
        claimData.deadline = Timestamp.wrap(uint64(block.timestamp + MAX_PROVE_DURATION.raw()));

        // Deposit the bond.
        refundModeCredit[msg.sender] += msg.value;

        emit Challenged(claimData.counteredBy);

        return claimData.status;
    }

    /// @notice Proves the game.
    /// @param proofBytes The proof bytes to validate the claim.
    function prove(bytes calldata proofBytes) external returns (ProposalStatus) {
        // INVARIANT: Cannot prove if the game is over.
        if (gameOver()) revert GameOver();

        // Decode the public values to check the claim root
        AggregationOutputs memory publicValues = AggregationOutputs({
            l1Head: Hash.unwrap(l1Head()),
            l2PreRoot: Hash.unwrap(startingOutputRoot.root),
            claimRoot: rootClaim().raw(),
            claimBlockNum: l2BlockNumber(),
            rollupConfigHash: ROLLUP_CONFIG_HASH,
            rangeVkeyCommitment: RANGE_VKEY_COMMITMENT,
            proverAddress: msg.sender
        });

        // Verify the proof. Reverts if the proof is invalid.
        SP1_VERIFIER.verifyProof(AGGREGATION_VKEY, abi.encode(publicValues), proofBytes);

        // Update the prover address
        claimData.prover = msg.sender;

        // Update the status of the proposal
        if (claimData.counteredBy == address(0)) {
            claimData.status = ProposalStatus.UnchallengedAndValidProofProvided;
        } else {
            claimData.status = ProposalStatus.ChallengedAndValidProofProvided;
        }

        emit Proved(claimData.prover);

        return claimData.status;
    }

    /// @notice Returns the status of the parent game.
    function getParentGameStatus() private view returns (GameStatus) {
        if (parentIndex() != type(uint32).max) {
            (,, IDisputeGame parentGame) = DISPUTE_GAME_FACTORY.gameAtIndex(parentIndex());
            return parentGame.status();
        } else {
            return GameStatus.DEFENDER_WINS;
        }
    }

    /// @notice Resolves the game after the clock expires.
    function resolve() external returns (GameStatus) {
        // INVARIANT: Resolution cannot occur unless the game has already been resolved.
        if (status != GameStatus.IN_PROGRESS) revert ClaimAlreadyResolved();

        // INVARIANT: Cannot resolve a game if the parent game has not been resolved.
        GameStatus parentGameStatus = getParentGameStatus();
        if (parentGameStatus == GameStatus.IN_PROGRESS) {
            revert ParentGameNotResolved();
        }

        // INVARIANT: If the parent game's claim is invalid, then the current game's claim is invalid.
        if (parentGameStatus == GameStatus.CHALLENGER_WINS) {
            status = GameStatus.CHALLENGER_WINS;
            normalModeCredit[claimData.counteredBy] = address(this).balance;
        } else {
            // INVARIANT: Game must be completed either by clock expiration or valid proof.
            if (!gameOver()) revert GameNotOver();

            // Determine status based on claim status.
            if (claimData.status == ProposalStatus.Unchallenged) {
                status = GameStatus.DEFENDER_WINS;
                normalModeCredit[gameCreator()] = address(this).balance;
            } else if (claimData.status == ProposalStatus.Challenged) {
                status = GameStatus.CHALLENGER_WINS;
                normalModeCredit[claimData.counteredBy] = address(this).balance;
            } else if (claimData.status == ProposalStatus.UnchallengedAndValidProofProvided) {
                status = GameStatus.DEFENDER_WINS;
                normalModeCredit[gameCreator()] = address(this).balance;
            } else if (claimData.status == ProposalStatus.ChallengedAndValidProofProvided) {
                status = GameStatus.DEFENDER_WINS;

                if (claimData.prover == gameCreator()) {
                    normalModeCredit[claimData.prover] = address(this).balance;
                } else {
                    normalModeCredit[claimData.prover] = CHALLENGER_BOND;
                    normalModeCredit[gameCreator()] = address(this).balance - CHALLENGER_BOND;
                }
            } else {
                revert InvalidProposalStatus();
            }
        }

        // Mark the game as resolved.
        claimData.status = ProposalStatus.Resolved;
        resolvedAt = Timestamp.wrap(uint64(block.timestamp));
        emit Resolved(status);

        return status;
    }

    /// @notice Claim the credit belonging to the recipient address.
    /// @param _recipient The owner and recipient of the credit.
    function claimCredit(address _recipient) external {
        closeGame();

        uint256 recipientCredit;
        if (bondDistributionMode == BondDistributionMode.REFUND) {
            recipientCredit = refundModeCredit[_recipient];
        } else if (bondDistributionMode == BondDistributionMode.NORMAL) {
            recipientCredit = normalModeCredit[_recipient];
        } else {
            revert InvalidBondDistributionMode();
        }

        if (recipientCredit == 0) revert NoCreditToClaim();

        refundModeCredit[_recipient] = 0;
        normalModeCredit[_recipient] = 0;

        (bool success,) = _recipient.call{ value: recipientCredit }(hex"");
        if (!success) revert BondTransferFailed();
    }

    /// @notice Closes out the game and determines the bond distribution mode.
    function closeGame() public {
        if (bondDistributionMode == BondDistributionMode.REFUND || bondDistributionMode == BondDistributionMode.NORMAL)
        {
            return;
        } else if (bondDistributionMode != BondDistributionMode.UNDECIDED) {
            revert InvalidBondDistributionMode();
        }

        bool finalized = ANCHOR_STATE_REGISTRY.isGameFinalized(IDisputeGame(address(this)));
        if (!finalized) {
            revert GameNotFinalized();
        }

        try ANCHOR_STATE_REGISTRY.setAnchorState(IDisputeGame(address(this))) { } catch { }

        bool properGame = ANCHOR_STATE_REGISTRY.isGameProper(IDisputeGame(address(this)));

        if (properGame) {
            bondDistributionMode = BondDistributionMode.NORMAL;
        } else {
            bondDistributionMode = BondDistributionMode.REFUND;
        }

        emit GameClosed(bondDistributionMode);
    }

    /// @notice Determines if the game is finished.
    function gameOver() public view returns (bool gameOver_) {
        gameOver_ = claimData.deadline.raw() < uint64(block.timestamp) || claimData.prover != address(0);
    }

    /// @notice Getter for the game type.
    function gameType() public view returns (GameType gameType_) {
        gameType_ = GAME_TYPE;
    }

    /// @notice Getter for the creator of the dispute game.
    function gameCreator() public pure returns (address creator_) {
        creator_ = _getArgAddress(0x00);
    }

    /// @notice Getter for the root claim.
    function rootClaim() public pure returns (Claim rootClaim_) {
        rootClaim_ = Claim.wrap(_getArgBytes32(0x14));
    }

    /// @notice Getter for the parent hash of the L1 block when the dispute game was created.
    function l1Head() public pure returns (Hash l1Head_) {
        l1Head_ = Hash.wrap(_getArgBytes32(0x34));
    }

    /// @notice Getter for the extra data.
    function extraData() public pure returns (bytes memory extraData_) {
        extraData_ = _getArgBytes(0x54, 0x24);
    }

    /// @notice Returns the game data.
    function gameData() external view returns (GameType gameType_, Claim rootClaim_, bytes memory extraData_) {
        gameType_ = gameType();
        rootClaim_ = rootClaim();
        extraData_ = extraData();
    }

    ////////////////////////////////////////////////////////////////
    //                       MISC EXTERNAL                        //
    ////////////////////////////////////////////////////////////////

    /// @notice Returns the credit balance of a given recipient.
    function credit(address _recipient) external view returns (uint256 credit_) {
        if (bondDistributionMode == BondDistributionMode.REFUND) {
            credit_ = refundModeCredit[_recipient];
        } else {
            credit_ = normalModeCredit[_recipient];
        }
    }

    ////////////////////////////////////////////////////////////////
    //                     IMMUTABLE GETTERS                      //
    ////////////////////////////////////////////////////////////////

    /// @notice Returns the max challenge duration.
    function maxChallengeDuration() external view returns (Duration maxChallengeDuration_) {
        maxChallengeDuration_ = MAX_CHALLENGE_DURATION;
    }

    /// @notice Returns the max prove duration.
    function maxProveDuration() external view returns (Duration maxProveDuration_) {
        maxProveDuration_ = MAX_PROVE_DURATION;
    }

    /// @notice Returns the dispute game factory.
    function disputeGameFactory() external view returns (IDisputeGameFactory disputeGameFactory_) {
        disputeGameFactory_ = DISPUTE_GAME_FACTORY;
    }

    /// @notice Returns the challenger bond amount.
    function challengerBond() external view returns (uint256 challengerBond_) {
        challengerBond_ = CHALLENGER_BOND;
    }

    /// @notice Returns the anchor state registry contract.
    function anchorStateRegistry() external view returns (IAnchorStateRegistry registry_) {
        registry_ = ANCHOR_STATE_REGISTRY;
    }

    /// @notice Returns the access manager contract.
    function accessManager() external view returns (AccessManager accessManager_) {
        accessManager_ = ACCESS_MANAGER;
    }
}
