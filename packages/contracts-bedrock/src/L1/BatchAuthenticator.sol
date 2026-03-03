// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { OwnableUpgradeable } from
    "lib/espresso-tee-contracts/lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { ServiceType } from "@espresso-tee-contracts/types/Types.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { OwnableWithGuardiansUpgradeable } from "@espresso-tee-contracts/OwnableWithGuardiansUpgradeable.sol";
import { ProxyAdminOwnedBase } from "src/L1/ProxyAdminOwnedBase.sol";
import { ReinitializableBase } from "src/universal/ReinitializableBase.sol";

/// @notice Upgradeable contract that authenticates batch information using the Transparent Proxy
///         pattern.
///         Supports switching between TEE and non-TEE batchers.
contract BatchAuthenticator is
    IBatchAuthenticator,
    ISemver,
    OwnableWithGuardiansUpgradeable,
    ProxyAdminOwnedBase,
    ReinitializableBase
{
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Mapping of batches verified by this contract
    mapping(bytes32 => bool) public validBatchInfo;

    /// @notice Address of the TEE batcher whose signatures may authenticate batches.
    address public teeBatcher;

    /// @notice Address of the non-TEE (fallback) batcher that can post when TEE is inactive.
    address public nonTeeBatcher;

    /// @notice Address of the Espresso TEE Verifier contract.
    IEspressoTEEVerifier public espressoTEEVerifier;

    /// @notice Flag indicating which batcher is currently active.
    /// @dev When true the TEE batcher is active; when false the non-TEE batcher is active.
    bool public activeIsTee;

    /// @notice Constructor disables initializers on implementation
    constructor() ReinitializableBase(1) {
        _disableInitializers();
    }

    function initialize(
        IEspressoTEEVerifier _espressoTEEVerifier,
        address _teeBatcher,
        address _nonTeeBatcher,
        address _owner
    )
        external
        reinitializer(initVersion())
    {
        // Initialization transactions must come from the ProxyAdmin or its owner.
        _assertOnlyProxyAdminOrProxyAdminOwner();

        // Initialize OwnableWithGuardians with the provided owner address
        __OwnableWithGuardians_init(_owner);

        if (_teeBatcher == address(0)) revert InvalidAddress(_teeBatcher);
        if (_nonTeeBatcher == address(0)) revert InvalidAddress(_nonTeeBatcher);
        if (address(_espressoTEEVerifier) == address(0)) {
            revert InvalidAddress(address(_espressoTEEVerifier));
        }

        espressoTEEVerifier = _espressoTEEVerifier;
        teeBatcher = _teeBatcher;
        nonTeeBatcher = _nonTeeBatcher;
        // By default, start with the TEE batcher active.
        activeIsTee = true;
    }

    /// @notice Returns the owner of the contract.
    function owner() public view override(IBatchAuthenticator, OwnableUpgradeable) returns (address) {
        return super.owner();
    }

    /// @notice Toggles the active batcher between the TEE and non-TEE batcher.
    function switchBatcher() external onlyGuardianOrOwner {
        activeIsTee = !activeIsTee;
        emit BatcherSwitched(activeIsTee);
    }

    /// @notice Updates the TEE batcher address.
    function setTeeBatcher(address _newTeeBatcher) external onlyOwner {
        if (_newTeeBatcher == address(0)) revert InvalidAddress(_newTeeBatcher);
        address oldTeeBatcher = teeBatcher;
        teeBatcher = _newTeeBatcher;
        emit TeeBatcherUpdated(oldTeeBatcher, _newTeeBatcher);
    }

    /// @notice Updates the non-TEE batcher address.
    function setNonTeeBatcher(address _newNonTeeBatcher) external onlyOwner {
        if (_newNonTeeBatcher == address(0)) {
            revert InvalidAddress(_newNonTeeBatcher);
        }
        address oldNonTeeBatcher = nonTeeBatcher;
        nonTeeBatcher = _newNonTeeBatcher;
        emit NonTeeBatcherUpdated(oldNonTeeBatcher, _newNonTeeBatcher);
    }

    function authenticateBatchInfo(bytes32 commitment, bytes calldata _signature) external {
        // Setting TEEType as Nitro because OP integration only supports AWS Nitro currently
        espressoTEEVerifier.verify(_signature, commitment, IEspressoTEEVerifier.TeeType.NITRO, ServiceType.BatchPoster);

        validBatchInfo[commitment] = true;
        emit BatchInfoAuthenticated(commitment);
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        espressoTEEVerifier.registerService(
            attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO, ServiceType.BatchPoster
        );
        emit SignerRegistrationInitiated(msg.sender);
    }

    /// @notice Returns the address of the Nitro TEE validator.
    function nitroValidator() external view returns (address) {
        return address(espressoTEEVerifier.espressoNitroTEEVerifier());
    }

    /// @notice Validates a batch submission in TEE mode.
    /// @param sender The address attempting to submit the batch.
    /// @param data The batch data being submitted.
    /// @dev Checks sender is teeBatcher and batch is authenticated.
    ///      Handles both blob and calldata batches.
    function validateTeeBatch(address sender, bytes calldata data) public view {
        // Check sender authorization
        if (sender != teeBatcher) {
            revert(
                string(
                    abi.encodePacked(
                        "BatchInbox: batcher not authorized to post in TEE mode. Expected: ",
                        Strings.toHexString(uint160(teeBatcher), 20),
                        ", Actual: ",
                        Strings.toHexString(uint160(sender), 20)
                    )
                )
            );
        }

        // Check batch authentication
        if (blobhash(0) != 0) {
            // Blob batch: concatenate all blob hashes
            uint256 numBlobs = 0;
            while (blobhash(numBlobs) != 0) {
                numBlobs++;
            }
            bytes memory concatenatedHashes = new bytes(32 * numBlobs);
            for (uint256 i = 0; i < numBlobs; i++) {
                assembly {
                    mstore(add(concatenatedHashes, add(0x20, mul(i, 32))), blobhash(i))
                }
            }
            bytes32 hash = keccak256(concatenatedHashes);
            if (!validBatchInfo[hash]) {
                revert("Invalid blob batch");
            }
        } else {
            // Calldata batch
            bytes32 hash = keccak256(data);
            if (!validBatchInfo[hash]) {
                revert("Invalid calldata batch");
            }
        }
    }

    /// @notice Validates a batch submission in non-TEE (fallback) mode.
    /// @param sender The address attempting to submit the batch.
    /// @dev Only checks sender is nonTeeBatcher. No batch authentication required.
    function validateNonTeeBatch(address sender) public view {
        if (sender != nonTeeBatcher) {
            revert(
                string(
                    abi.encodePacked(
                        "BatchInbox: batcher not authorized to post in fallback mode. Expected: ",
                        Strings.toHexString(uint160(nonTeeBatcher), 20),
                        ", Actual: ",
                        Strings.toHexString(uint160(sender), 20)
                    )
                )
            );
        }
    }

    /// @notice Validates a batch submission based on current batcher mode.
    /// @param sender The address attempting to submit the batch.
    /// @param data The batch data being submitted.
    /// @dev Routes to validateTeeBatch or validateNonTeeBatch based on activeIsTee.
    function validateBatch(address sender, bytes calldata data) external view {
        if (activeIsTee) {
            validateTeeBatch(sender, data);
        } else {
            validateNonTeeBatch(sender);
        }
    }
}
