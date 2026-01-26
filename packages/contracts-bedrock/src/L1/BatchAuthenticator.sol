// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { OwnableWithGuardiansUpgradeable } from "@espresso-tee-contracts/OwnableWithGuardiansUpgradeable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { ServiceType } from "@espresso-tee-contracts/types/Types.sol";

contract BatchAuthenticator is ISemver, OwnableWithGuardiansUpgradeable {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Emitted when a batch info is authenticated.
    event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer);

    /// @notice Emitted when a signer registration is initiated through this contract.
    event SignerRegistrationInitiated(address indexed caller);

    /// @notice Mapping of batches verified by this contract
    mapping(bytes32 => bool) public validBatchInfo;

    /// @notice Address of the TEE batcher whose signatures may authenticate batches.
    address public teeBatcher;

    /// @notice Address of the non-TEE (fallback) batcher that can post when TEE is inactive.
    address public nonTeeBatcher;

    IEspressoTEEVerifier public espressoTEEVerifier;

    /// @notice Flag indicating which batcher is currently active.
    /// @dev When true the TEE batcher is active; when false the non-TEE batcher is active.
    bool public activeIsTee;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initializes the contract with verifier addresses and initial owner
    /// @param _espressoTEEVerifier The Espresso TEE verifier contract address
    /// @param _teeBatcher The TEE batcher address
    /// @param _nonTeeBatcher The non-TEE fallback batcher address
    /// @param _owner The initial owner address
    function initialize(
        IEspressoTEEVerifier _espressoTEEVerifier,
        address _teeBatcher,
        address _nonTeeBatcher,
        address _owner
    )
        public
        initializer
    {
        require(_teeBatcher != address(0), "BatchAuthenticator: zero tee batcher");
        require(_nonTeeBatcher != address(0), "BatchAuthenticator: zero non-tee batcher");

        __OwnableWithGuardians_init(_owner);

        espressoTEEVerifier = _espressoTEEVerifier;
        teeBatcher = _teeBatcher;
        nonTeeBatcher = _nonTeeBatcher;
        // By default, start with the TEE batcher active.
        activeIsTee = true;
    }

    /// @notice Toggles the active batcher between the TEE and non-TEE batcher.
    /// @dev Can be called by either the owner or a guardian for emergency response.
    function switchBatcher() external onlyGuardianOrOwner {
        activeIsTee = !activeIsTee;
    }

    /// @notice Validates a batch submission from the BatchInbox.
    /// @dev Routes to the appropriate validation function based on activeIsTee.
    /// @param sender The address that sent the batch to BatchInbox.
    /// @param data The calldata of the original batch submission.
    function validateBatch(address sender, bytes calldata data) external view {
        if (activeIsTee) {
            _validateTeeBatch(sender, data);
        } else {
            _validateNonTeeBatch(sender);
        }
    }

    /// @notice Validates a TEE batch submission.
    /// @dev Checks that the sender is the TEE batcher and that the batch commitment
    ///      has been pre-authenticated via authenticateBatchInfo().
    ///      Note: blobhash() works here because it returns blob hashes for the current
    ///      transaction, regardless of call depth.
    /// @param sender The address that sent the batch to BatchInbox.
    /// @param data The calldata of the original batch submission.
    function _validateTeeBatch(address sender, bytes calldata data) internal view {
        require(
            sender == teeBatcher,
            string(
                abi.encodePacked(
                    "BatchAuthenticator: batcher not authorized to post in TEE mode. Expected: ",
                    Strings.toHexString(uint160(teeBatcher), 20),
                    ", Actual: ",
                    Strings.toHexString(uint160(sender), 20)
                )
            )
        );

        bytes32 hash;
        if (blobhash(0) != 0) {
            // Blob batch: hash all blob hashes together
            bytes memory concatenatedHashes = new bytes(0);
            uint256 i = 0;
            while (blobhash(i) != 0) {
                concatenatedHashes = bytes.concat(concatenatedHashes, blobhash(i));
                i++;
            }
            hash = keccak256(concatenatedHashes);
            require(validBatchInfo[hash], "BatchAuthenticator: invalid blob batch");
        } else {
            // Calldata batch: hash the calldata
            hash = keccak256(data);
            require(validBatchInfo[hash], "BatchAuthenticator: invalid calldata batch");
        }
    }

    /// @notice Validates a non-TEE (fallback) batch submission.
    /// @dev Only checks that the sender is the non-TEE batcher. No batch content
    ///      validation is performed in fallback mode.
    /// @param sender The address that sent the batch to BatchInbox.
    function _validateNonTeeBatch(address sender) internal view {
        require(
            sender == nonTeeBatcher,
            string(
                abi.encodePacked(
                    "BatchAuthenticator: batcher not authorized to post in fallback mode. Expected: ",
                    Strings.toHexString(uint160(nonTeeBatcher), 20),
                    ", Actual: ",
                    Strings.toHexString(uint160(sender), 20)
                )
            )
        );
    }

    function authenticateBatchInfo(bytes32 commitment, bytes calldata _signature) external {
        // https://github.com/ethereum/go-ethereum/issues/19751#issuecomment-504900739
        bytes memory signature = _signature;
        require(signature.length == 65, "Invalid signature length");
        uint8 v = uint8(signature[64]);
        if (v == 0 || v == 1) {
            v += 27;
            signature[64] = bytes1(v);
        }
        address signer = ECDSA.recover(commitment, signature);

        require(signer != address(0), "BatchAuthenticator: invalid signature");

        require(
            espressoTEEVerifier.registeredService(signer, IEspressoTEEVerifier.TeeType.NITRO, ServiceType.BatchPoster),
            "BatchAuthenticator: invalid signer"
        );

        validBatchInfo[commitment] = true;
        emit BatchInfoAuthenticated(commitment, signer);
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        espressoTEEVerifier.registerService(
            attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO, ServiceType.BatchPoster
        );
        emit SignerRegistrationInitiated(msg.sender);
    }
}
