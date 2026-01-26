// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

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

    function initialize(
        address _espressoTEEVerifier,
        address _teeBatcher,
        address _nonTeeBatcher,
        address _owner
    ) external;

    function validateBatch(address sender, bytes calldata data) external view;

    // Guardian functions
    function addGuardian(address guardian) external;

    function removeGuardian(address guardian) external;

    function isGuardian(address account) external view returns (bool);

    function getGuardians() external view returns (address[] memory);

    function guardianCount() external view returns (uint256);
}
