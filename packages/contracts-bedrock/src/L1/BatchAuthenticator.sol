// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";

contract BatchAuthenticator is ISemver, Ownable {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Emitted when a batch info is authenticated.
    event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer);

    /// @notice Emitted when a signer registration is initiated through this contract.
    event SignerRegistrationInitiated(address indexed caller);

    /// @notice Mapping of batches verified by this contract
    mapping(bytes32 => bool) public validBatchInfo;

    /// @notice Address of the TEE batcher whose signatures may authenticate batches.
    address public immutable teeBatcher;

    /// @notice Address of the non-TEE (fallback) batcher that can post when TEE is inactive.
    address public immutable nonTeeBatcher;

    /// @notice Key of pre-registered TEE batcher that can authenticate batches without
    /// calling registerSigner first. For testing only.
    address public immutable preRegisteredBatcher;

    IEspressoTEEVerifier public immutable espressoTEEVerifier;

    /// @notice Flag indicating which batcher is currently active.
    /// @dev When true the TEE batcher is active; when false the non-TEE batcher is active.
    bool public activeIsTee;

    constructor(
        IEspressoTEEVerifier _espressoTEEVerifier,
        address _teeBatcher,
        address _nonTeeBatcher,
        address _preRegisteredBatcher,
        address _owner
    )
        Ownable()
    {
        require(_teeBatcher != address(0), "BatchAuthenticator: zero tee batcher");
        require(_nonTeeBatcher != address(0), "BatchAuthenticator: zero non-tee batcher");

        espressoTEEVerifier = _espressoTEEVerifier;
        teeBatcher = _teeBatcher;
        nonTeeBatcher = _nonTeeBatcher;
        // By default, start with the TEE batcher active.
        activeIsTee = true;
        preRegisteredBatcher = _preRegisteredBatcher;
        _transferOwnership(_owner);
    }

    /// @notice Toggles the active batcher between the TEE and non-TEE batcher.
    function switchBatcher() external onlyOwner {
        activeIsTee = !activeIsTee;
    }

    function authenticateBatchInfo(bytes32 commitment, bytes calldata _signature) external {
        // https://github.com/ethereum/go-ethereum/issues/19751#issuecomment-504900739
        bytes memory signature = _signature;
        require(signature.length == 65, "Invalid signature length");
        uint8 v = uint8(signature[64]);
        if (v == 0 || v == 1) {
            v += 27;
            signature[64] = bytes1(v);
        }
        address signer = ECDSA.recover(commitment, signature);

        require(signer != address(0), "BatchAuthenticator: invalid signature");

        if (signer != preRegisteredBatcher) {
            require(
                espressoTEEVerifier.espressoNitroTEEVerifier().registeredSigners(signer),
                "BatchAuthenticator: invalid signer"
            );
        }

        validBatchInfo[commitment] = true;
        emit BatchInfoAuthenticated(commitment, signer);
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        espressoTEEVerifier.registerSigner(attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO);
        emit SignerRegistrationInitiated(msg.sender);
    }
}
