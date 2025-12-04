// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";

interface INitroValidator {
    function decodeAttestationTbs(bytes memory attestation)
        external
        pure
        returns (bytes memory attestationTbs, bytes memory signature);
}

contract BatchAuthenticator is ISemver, Ownable {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Mapping of batches verified by this contract
    mapping(bytes32 => bool) public validBatchInfo;

    /// @notice Address of the TEE batcher whose signatures may authenticate batches.
    address public immutable teeBatcher;

    /// @notice Address of the non-TEE (fallback) batcher that can post when TEE is inactive.
    address public immutable nonTeeBatcher;

    IEspressoTEEVerifier public immutable espressoTEEVerifier;
    INitroValidator public immutable nitroValidator;

    /// @notice Flag indicating which batcher is currently active.
    /// @dev When true the TEE batcher is active; when false the non-TEE batcher is active.
    bool public activeIsTee;

    constructor(
        IEspressoTEEVerifier _espressoTEEVerifier,
        address _teeBatcher,
        address _nonTeeBatcher,
        address _owner
    )
        Ownable()
    {
        require(_teeBatcher != address(0), "BatchAuthenticator: zero tee batcher");
        require(_nonTeeBatcher != address(0), "BatchAuthenticator: zero non-tee batcher");

        espressoTEEVerifier = _espressoTEEVerifier;
        teeBatcher = _teeBatcher;
        nonTeeBatcher = _nonTeeBatcher;
        nitroValidator = INitroValidator(address(espressoTEEVerifier.espressoNitroTEEVerifier()));
        // By default, start with the TEE batcher active.
        activeIsTee = true;
        _transferOwnership(_owner);
    }

    function decodeAttestationTbs(bytes memory attestation) external view returns (bytes memory, bytes memory) {
        return nitroValidator.decodeAttestationTbs(attestation);
    }

    /// @notice Toggles the active batcher between the TEE and non-TEE batcher.
    function switchBatcher() external onlyOwner {
        activeIsTee = !activeIsTee;
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

        if (!espressoTEEVerifier.espressoNitroTEEVerifier().registeredSigners(signer) && signer != teeBatcher) {
            revert("Invalid signer");
        }

        validBatchInfo[commitment] = true;
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        espressoTEEVerifier.registerSigner(attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO);
    }

    function registerSignerWithoutAttestationVerification(
        bytes32 pcr0Hash,
        bytes calldata attestationTbs,
        bytes calldata signature,
        address enclaveAddress
    )
        external
    {
        espressoTEEVerifier.espressoNitroTEEVerifier().registerSignerWithoutAttestationVerification(
            pcr0Hash, attestationTbs, signature, enclaveAddress
        );
    }
}
