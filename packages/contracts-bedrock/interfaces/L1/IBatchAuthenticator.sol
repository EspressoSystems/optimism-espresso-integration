// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IBatchAuthenticator {
    event Initialized(uint8 version);
    event OwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner
    );

    function decodeAttestationTbs(
        bytes memory attestation
    ) external view returns (bytes memory, bytes memory);

    function espressoTEEVerifier() external view returns (address);

    function owner() external view returns (address);

    function registerSigner(
        bytes memory attestationTbs,
        bytes memory signature
    ) external;

    function renounceOwnership() external;

    function transferOwnership(address newOwner) external;

    function validBatches(bytes32) external view returns (bool);

    function authenticateBatch(
        bytes32 commitment,
        bytes memory signature
    ) external;

    function version() external view returns (string memory);

    function __constructor__(
        address _espressoTEEVerifier,
        address _preApprovedBatcher
    ) external;
}
