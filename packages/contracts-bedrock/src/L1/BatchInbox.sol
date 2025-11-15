// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

/// @title BatchInbox
/// @notice Receives batches from either a TEE batcher or a non-TEE batcher and enforces
///         that TEE batches are authenticated by the configured batch authenticator.
contract BatchInbox {
    /// @notice Address of the TEE-based batcher.
    address public immutable teeBatcher;

    /// @notice Address of the non-TEE (fallback) batcher.
    address public immutable nonTeeBatcher;

    /// @notice Contract responsible for authenticating TEE batch commitments.
    IBatchAuthenticator public immutable batchAuthenticator;

    /// @notice Flag indicating which batcher is currently active.
    /// @dev When true the TEE batcher is active; when false the non-TEE batcher is active.
    bool public activeIsTee;

    /// @notice Initializes the contract with the TEE and non-TEE batcher addresses
    ///         and the batch authenticator.
    /// @param _teeBatcher Address of the TEE batcher.
    /// @param _nonTeeBatcher Address of the non-TEE batcher.
    /// @param _batchAuthenticator Address of the batch authenticator contract.
    constructor(address _teeBatcher, address _nonTeeBatcher, IBatchAuthenticator _batchAuthenticator) {
        require(_teeBatcher != address(0) && _nonTeeBatcher != address(0), "BatchInbox: zero batcher");
        teeBatcher = _teeBatcher;
        nonTeeBatcher = _nonTeeBatcher;
        batchAuthenticator = _batchAuthenticator;
        // By default, start with the TEE batcher active
        activeIsTee = true;
    }

    /// @notice Toggles the active batcher between the TEE and non-TEE batcher.
    function switchBatcher() external {
        activeIsTee = !activeIsTee;
    }

    /// @notice Fallback entry point for batch submissions.
    /// @dev Enforces that the caller matches the currently active batcher and, when
    ///      the TEE batcher is active, that the batch commitment is approved by
    ///      the batch authenticator. For non-TEE batches, only the caller check
    ///      is enforced.
    fallback() external {
        address expectedBatcher = activeIsTee ? teeBatcher : nonTeeBatcher;
        if (msg.sender != expectedBatcher) {
            revert("BatchInbox: unauthorized batcher");
        }

        // Only TEE batchers require authentication
        if (activeIsTee) {
            if (blobhash(0) != 0) {
                bytes memory concatenatedHashes = new bytes(0);
                uint256 currentBlob = 0;
                while (blobhash(currentBlob) != 0) {
                    concatenatedHashes = bytes.concat(concatenatedHashes, blobhash(currentBlob));
                    currentBlob++;
                }
                bytes32 hash = keccak256(concatenatedHashes);
                if (!batchAuthenticator.validBatchInfo(hash)) {
                    revert("Invalid blob batch");
                }
            } else {
                bytes32 hash = keccak256(msg.data);
                if (!batchAuthenticator.validBatchInfo(hash)) {
                    revert("Invalid calldata batch");
                }
            }
        }
    }

    /// @notice Returns the currently active batcher and whether it is the TEE batcher.
    /// @return active Address of the currently active batcher.
    /// @return isTee True if the active batcher is the TEE batcher, false if it is the non-TEE batcher.
    function _activeBatcher() internal view returns (address active, bool isTee) {
        if (activeIsTee) {
            return (teeBatcher, true);
        } else {
            return (nonTeeBatcher, false);
        }
    }
}
