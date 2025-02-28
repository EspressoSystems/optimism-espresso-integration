// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { NitroValidator } from "@nitro-validator/NitroValidator.sol";
import { CborDecode } from "@nitro-validator/CborDecode.sol";
import { CertManager } from "@nitro-validator/CertManager.sol";
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";

contract BatchInbox is OwnableUpgradeable, ISemver, NitroValidator {
    /// @notice Semantic version.
    /// @custom:semver 1.0.1
    string public constant version = "1.0.1";

    using CborDecode for bytes;

    uint256 public constant MAX_AGE = 60 minutes;

    /// @notice Mapping of valid PCR0 values for batchers
    mapping(bytes32 => bool) public validPCR0s;

    // @notice Mapping of attested batchers, effectively used as a hashset, i.e. the value is unused
    mapping(address => bool) public attestedBatchers;

    constructor(
        CertManager certManager,
        address preApprovedBatcherKey
    )
        NitroValidator(certManager)
        OwnableUpgradeable()
    {
        initialize(preApprovedBatcherKey);
    }

    function initialize(address preApprovedBatcherKey) public initializer {
        __Ownable_init();
        if (preApprovedBatcherKey != address(0)) {
            attestedBatchers[preApprovedBatcherKey] = true;
        }
    }

    function registerPCR0(bytes calldata pcr0) external onlyOwner {
        validPCR0s[keccak256(pcr0)] = true;
    }

    function submitBatch(bytes calldata commitment, bytes calldata _signature) external {
        // https://github.com/ethereum/go-ethereum/issues/19751#issuecomment-504900739
        bytes memory signature = _signature;
        uint8 v = uint8(signature[64]);
        if (v == 0 || v == 1) {
            v += 27;
            signature[64] = bytes1(v);
        }

        address signer = ECDSA.recover(keccak256(commitment), signature);
        require(signer != address(0), "could not extract signer");
        require(attestedBatchers[signer], "unauthorized batcher");
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        Ptrs memory ptrs = validateAttestation(attestationTbs, signature);
        bytes32 pcr0 = attestationTbs.keccak(ptrs.pcrs[0]);
        require(validPCR0s[pcr0], "invalid pcr0 in attestation");
        require(ptrs.timestamp + MAX_AGE > block.timestamp, "attestation too old");
        bytes memory publicKey = attestationTbs.slice(ptrs.publicKey);
        require(publicKey.length == 20, "invalid publicKey");
        address signer = address(bytes20(publicKey));
        // Register this batcher as a trusted signer
        attestedBatchers[signer] = true;
    }
}
