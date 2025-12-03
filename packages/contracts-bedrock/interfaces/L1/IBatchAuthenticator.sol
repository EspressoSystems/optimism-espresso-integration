// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IBatchAuthenticator {
    event Initialized(uint8 version);
    event OwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner
    );

    function authenticateBatchInfo(
        bytes32 commitment,
        bytes memory _signature
    ) external;

    function decodeAttestationTbs(
        bytes memory attestation
    ) external view returns (bytes memory, bytes memory);

    function espressoTEEVerifier() external view returns (address);

    function nitroValidator() external view returns (address);

    function owner() external view returns (address);

    function teeBatcher() external view returns (address);

    function nonTeeBatcher() external view returns (address);

    function registerSigner(
        bytes memory attestationTbs,
        bytes memory signature
    ) external;

    function renounceOwnership() external;

    function transferOwnership(address newOwner) external;

    function validBatchInfo(bytes32) external view returns (bool);

    function activeIsTee() external view returns (bool);

    function switchBatcher() external;

    function __constructor__(
        address _espressoTEEVerifier,
        address _teeBatcher,
        address _nonTeeBatcher,
        address _owner
    ) external;
}
