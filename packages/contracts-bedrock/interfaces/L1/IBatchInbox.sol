// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title IBatchInbox
/// @notice Interface for the BatchInbox contract.
/// @dev Note: This contract intentionally has no public/external view functions to ensure
///      ALL calls route to the fallback function. This prevents function selector collisions
///      that could bypass batch authentication.
interface IBatchInbox {
    fallback() external;

    function __constructor__(address batchAuthenticator_) external;
}
