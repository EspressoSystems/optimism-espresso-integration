// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IBatchInbox {
    function version() external view returns (string memory);

    function __constructor__(address _teeBatcher, address _nonTeeBatcher, address _batchAuthenticator) external;

    function switchBatcher() external;

    function postCalldata(bytes calldata data) external;

    function postBlobs() external;

    fallback() external;
}
