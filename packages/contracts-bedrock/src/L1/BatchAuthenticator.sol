// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { EspressoTEEVerifier } from "@espresso-tee-contracts/EspressoTEEVerifier.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";

interface INitroValidator {
    function decodeAttestationTbs(bytes memory attestation)
        external
        pure
        returns (bytes memory attestationTbs, bytes memory signature);
}

contract BatchAuthenticator is ISemver, OwnableUpgradeable {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Mapping of batches verified by this contract
    mapping(bytes32 => bool) public validBatchInfo;

    address public immutable preApprovedBatcher;

    EspressoTEEVerifier public immutable espressoTEEVerifier;
    INitroValidator public immutable nitroValidator;

    constructor(EspressoTEEVerifier _espressoTEEVerifier, address _preApprovedBatcher) OwnableUpgradeable() {
        espressoTEEVerifier = _espressoTEEVerifier;
        preApprovedBatcher = _preApprovedBatcher;
        nitroValidator = INitroValidator(address(espressoTEEVerifier.espressoNitroTEEVerifier()));
    }

    function decodeAttestationTbs(bytes memory attestation) external view returns (bytes memory, bytes memory) {
        return nitroValidator.decodeAttestationTbs(attestation);
    }

    function authenticateBatchInfo(bytes32 commitment, bytes calldata _signature) external {
        // https://github.com/ethereum/go-ethereum/issues/19751#issuecomment-504900739
        bytes memory signature = _signature;
        uint8 v = uint8(signature[64]);
        if (v == 0 || v == 1) {
            v += 27;
            signature[64] = bytes1(v);
        }
        address signer = ECDSA.recover(commitment, signature);

        if (signer == address(0)) {
            revert("Invalid signature");
        }

        if (!espressoTEEVerifier.espressoNitroTEEVerifier().registeredSigners(signer) && signer != preApprovedBatcher) {
            revert("Invalid signer");
        }

        validBatchInfo[commitment] = true;
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        espressoTEEVerifier.registerSigner(attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO);
    }
}
