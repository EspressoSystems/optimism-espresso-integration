// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

/// @title BatchInbox
/// @notice Receives batches from either a TEE batcher or a non-TEE batcher and enforces
///         that TEE batches are authenticated by the configured batch authenticator.
/// @dev This contract has NO public function selectors - all calls route to the fallback.
contract BatchInbox {
    /// @notice Contract responsible for authenticating TEE batch commitments.
    /// @dev Private to prevent creating a function selector.
    IBatchAuthenticator private immutable batchAuthenticator;

    /// @notice Initializes the contract with the batch authenticator.
    /// @param _batchAuthenticator Address of the batch authenticator contract.
    constructor(IBatchAuthenticator _batchAuthenticator) {
        batchAuthenticator = _batchAuthenticator;
    }

    /// @notice Fallback entry point for batch submissions.
    /// @dev Delegates all validation to the batch authenticator.
    fallback() external {
        batchAuthenticator.validateBatch(msg.sender, msg.data);
    }
}
