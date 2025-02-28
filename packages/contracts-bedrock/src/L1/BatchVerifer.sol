// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { EspressoTEEVerifier } from "@espresso-tee-contracts/EspressoTEEVerifier.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";

interface INitroValidator {
    function decodeAttestationTbs(bytes memory attestation)
        external
        pure
        returns (bytes memory attestationTbs, bytes memory signature);
}

contract BatchVerifier is ISemver, OwnableUpgradeable {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Mapping of batches verified by this contract
    mapping(bytes32 => bool) public validBatches;

    EspressoTEEVerifier public immutable espressoTEEVerifier;
    INitroValidator public immutable nitroValidator;

    constructor(EspressoTEEVerifier _espressoTEEVerifier) OwnableUpgradeable() {
        espressoTEEVerifier = _espressoTEEVerifier;
        nitroValidator = INitroValidator(address(espressoTEEVerifier.espressoNitroTEEVerifier()));
    }

    function decodeAttestationTbs(bytes memory attestation) external view returns (bytes memory, bytes memory) {
        return nitroValidator.decodeAttestationTbs(attestation);
    }

    function verifyBatch(bytes32 commitment, bytes calldata signature) external {
        address signer = ECDSA.recover(commitment, signature);

        if (!espressoTEEVerifier.espressoNitroTEEVerifier().registeredSigners(signer) || signer == address(0)) {
            revert("Invalid signature");
        }

        validBatches[commitment] = true;
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        espressoTEEVerifier.registerSigner(attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO);
    }
}
