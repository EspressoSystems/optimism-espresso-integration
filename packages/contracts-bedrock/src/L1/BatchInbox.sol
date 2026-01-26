// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

/// @title BatchInbox
/// @notice Receives batches from either a TEE batcher or a non-TEE batcher and delegates
///         validation to the BatchAuthenticator contract.
/// @dev This contract intentionally has no public/external functions to ensure ALL calls
///      route to the fallback function. This prevents function selector collisions that
///      could bypass batch authentication. The batchAuthenticator address can be read
///      directly from storage slot 0 if needed.
contract BatchInbox {
    /// @notice Contract responsible for validating batch submissions.
    /// @dev Internal to prevent function selector collision with fallback.
    ///      Can be read externally via storage slot 0.
    IBatchAuthenticator internal immutable _batchAuthenticator;

    /// @notice Initializes the contract with the batch authenticator.
    /// @param batchAuthenticator_ Address of the batch authenticator contract.
    constructor(IBatchAuthenticator batchAuthenticator_) {
        _batchAuthenticator = batchAuthenticator_;
    }

    /// @notice Fallback entry point for batch submissions.
    /// @dev Delegates all validation logic to BatchAuthenticator.validateBatch().
    ///      This allows the validation logic to be upgraded via the BatchAuthenticator proxy.
    fallback() external {
        _batchAuthenticator.validateBatch(msg.sender, msg.data);
    }
}
