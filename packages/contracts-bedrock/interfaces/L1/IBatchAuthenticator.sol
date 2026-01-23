// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IBatchAuthenticator {
    event Initialized(uint8 version);

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

    function validBatchInfo(bytes32) external view returns (bool);

    function activeIsTee() external view returns (bool);

    function switchBatcher() external;

    function setTeeBatcher(address _newTeeBatcher) external;

    function setNonTeeBatcher(address _newNonTeeBatcher) external;
}
