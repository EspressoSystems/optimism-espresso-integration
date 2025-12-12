// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

/// @title BatchInbox
/// @notice Receives batches from either a TEE batcher or a non-TEE batcher and enforces
///         that TEE batches are authenticated by the configured batch authenticator.
contract BatchInbox is Ownable {
    /// @notice Address of the non-TEE (fallback) batcher.
    address public immutable nonTeeBatcher;

    /// @notice Contract responsible for authenticating TEE batch commitments.
    IBatchAuthenticator public immutable batchAuthenticator;

    /// @notice Initializes the contract with the batch authenticator.
    /// @param _batchAuthenticator Address of the batch authenticator contract.
    constructor(IBatchAuthenticator _batchAuthenticator, address _owner) Ownable() {
        address _nonTeeBatcher = _batchAuthenticator.nonTeeBatcher();
        nonTeeBatcher = _nonTeeBatcher;
        batchAuthenticator = _batchAuthenticator;
        _transferOwnership(_owner);
    }

    /// @notice Fallback entry point for batch submissions.
    /// @dev Enforces that the caller matches the currently active batcher and, when
    ///      the TEE batcher is active, that the batch commitment is approved by
    ///      the batch authenticator. For non-TEE batches, only the caller check
    ///      is enforced.
    fallback() external {
        // TEE batchers require batch authentication
        if (batchAuthenticator.activeIsTee()) {
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
        } else {
            // Non TEE batcher require batcher address authentication
            if (msg.sender != nonTeeBatcher) {
                // For the non active TEE case, the batcher must be authenticated in the Inbox contract
                revert("BatchInbox: unauthorized batcher");
            }
        }
    }
}
