// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {CertManager} from "@nitro-validator/CertManager.sol";

type CborElement is uint256;

library NitroValidator {
    struct Ptrs {
        CborElement moduleID;
        uint64 timestamp;
        CborElement digest;
        CborElement[] pcrs;
        CborElement cert;
        CborElement[] cabundle;
        CborElement publicKey;
        CborElement userData;
        CborElement nonce;
    }
}

interface IBatchInbox {
    event Initialized(uint8 version);
    event OwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner
    );

    function ATTESTATION_DIGEST() external view returns (bytes32);

    function ATTESTATION_TBS_PREFIX() external view returns (bytes32);

    function CABUNDLE_KEY() external view returns (bytes32);

    function CERTIFICATE_KEY() external view returns (bytes32);

    function DIGEST_KEY() external view returns (bytes32);

    function MAX_AGE() external view returns (uint256);

    function MODULE_ID_KEY() external view returns (bytes32);

    function NONCE_KEY() external view returns (bytes32);

    function PCRS_KEY() external view returns (bytes32);

    function PUBLIC_KEY_KEY() external view returns (bytes32);

    function TIMESTAMP_KEY() external view returns (bytes32);

    function USER_DATA_KEY() external view returns (bytes32);

    function attestedBatchers(address) external view returns (bool);

    function certManager() external view returns (address);

    function decodeAttestationTbs(
        bytes memory attestation
    )
        external
        pure
        returns (bytes memory attestationTbs, bytes memory signature);

    function initialize(address preApprovedBatcherKey) external;

    function owner() external view returns (address);

    function registerPCR0(bytes memory pcr0) external;

    function registerSigner(
        bytes memory attestationTbs,
        bytes memory signature
    ) external;

    function renounceOwnership() external;

    function submitBatch(
        bytes memory commitment,
        bytes memory _signature
    ) external;

    function transferOwnership(address newOwner) external;

    function validPCR0s(bytes32) external view returns (bool);

    function validateAttestation(
        bytes memory attestationTbs,
        bytes memory signature
    ) external returns (NitroValidator.Ptrs memory);

    function version() external view returns (string memory);

    function __constructor__(
        CertManager certManager,
        address preApprovedBatcherKey
    ) external;
}
