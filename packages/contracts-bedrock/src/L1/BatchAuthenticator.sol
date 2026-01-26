// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { ProxyAdminOwnedBase } from "src/L1/ProxyAdminOwnedBase.sol";
import { ReinitializableBase } from "src/universal/ReinitializableBase.sol";

/// @notice Upgradeable contract that authenticates batch information using the Transparent Proxy
///         pattern.
///         Supports switching between TEE and non-TEE batchers.
contract BatchAuthenticator is IBatchAuthenticator, ISemver, Initializable, ProxyAdminOwnedBase, ReinitializableBase {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Mapping of batches verified by this contract
    mapping(bytes32 => bool) public validBatchInfo;

    /// @notice Address of the TEE batcher whose signatures may authenticate batches.
    address public teeBatcher;

    /// @notice Address of the non-TEE (fallback) batcher that can post when TEE is inactive.
    address public nonTeeBatcher;

    /// @notice Address of the Espresso TEE Verifier contract.
    IEspressoTEEVerifier public espressoTEEVerifier;

    /// @notice Flag indicating which batcher is currently active.
    /// @dev When true the TEE batcher is active; when false the non-TEE batcher is active.
    bool public activeIsTee;

    /// @notice Constructor disables initializers on implementation
    constructor() ReinitializableBase(1) {
        _disableInitializers();
    }

    function initialize(
        IEspressoTEEVerifier _espressoTEEVerifier,
        address _teeBatcher,
        address _nonTeeBatcher
    )
        external
        reinitializer(initVersion())
    {
        // Initialization transactions must come from the ProxyAdmin or its owner.
        _assertOnlyProxyAdminOrProxyAdminOwner();

        if (_teeBatcher == address(0)) revert InvalidAddress(_teeBatcher);
        if (_nonTeeBatcher == address(0)) revert InvalidAddress(_nonTeeBatcher);
        if (address(_espressoTEEVerifier) == address(0)) revert InvalidAddress(address(_espressoTEEVerifier));

        espressoTEEVerifier = _espressoTEEVerifier;
        teeBatcher = _teeBatcher;
        nonTeeBatcher = _nonTeeBatcher;
        // By default, start with the TEE batcher active.
        activeIsTee = true;
    }

    /// @notice Returns the owner of the ProxyAdmin that owns this proxy.
    function owner() external view returns (address) {
        return proxyAdminOwner();
    }

    /// @notice Toggles the active batcher between the TEE and non-TEE batcher.
    function switchBatcher() external {
        _assertOnlyProxyAdminOwner();
        activeIsTee = !activeIsTee;
    }

    /// @notice Updates the TEE batcher address.
    function setTeeBatcher(address _newTeeBatcher) external {
        _assertOnlyProxyAdminOwner();
        if (_newTeeBatcher == address(0)) revert InvalidAddress(_newTeeBatcher);
        address oldTeeBatcher = teeBatcher;
        teeBatcher = _newTeeBatcher;
        emit TeeBatcherUpdated(oldTeeBatcher, _newTeeBatcher);
    }

    /// @notice Updates the non-TEE batcher address.
    function setNonTeeBatcher(address _newNonTeeBatcher) external {
        _assertOnlyProxyAdminOwner();
        if (_newNonTeeBatcher == address(0)) revert InvalidAddress(_newNonTeeBatcher);
        address oldNonTeeBatcher = nonTeeBatcher;
        nonTeeBatcher = _newNonTeeBatcher;
        emit NonTeeBatcherUpdated(oldNonTeeBatcher, _newNonTeeBatcher);
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

        require(
            espressoTEEVerifier.espressoNitroTEEVerifier().registeredSigners(signer),
            "BatchAuthenticator: invalid signer"
        );

        validBatchInfo[commitment] = true;
        emit BatchInfoAuthenticated(commitment, signer);
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        espressoTEEVerifier.registerSigner(attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO);
        emit SignerRegistrationInitiated(msg.sender);
    }

    /// @notice Returns the address of the Nitro TEE validator.
    function nitroValidator() external view returns (address) {
        return address(espressoTEEVerifier.espressoNitroTEEVerifier());
    }
}
