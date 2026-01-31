// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IBatchInbox {
    fallback() external;

    function __constructor__(address _batchAuthenticator) external;
}
