// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { OwnableUpgradeable } from
    "lib/espresso-tee-contracts/lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { ServiceType } from "@espresso-tee-contracts/types/Types.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { ISystemConfig } from "interfaces/L1/ISystemConfig.sol";
import { OwnableWithGuardiansUpgradeable } from "@espresso-tee-contracts/OwnableWithGuardiansUpgradeable.sol";
import { ProxyAdminOwnedBase } from "src/L1/ProxyAdminOwnedBase.sol";
import { ReinitializableBase } from "src/universal/ReinitializableBase.sol";

/// @notice Upgradeable contract that authenticates batch information using the Transparent Proxy
///         pattern.
///         Supports switching between TEE and non-TEE batchers.
contract BatchAuthenticator is
    IBatchAuthenticator,
    ISemver,
    OwnableWithGuardiansUpgradeable,
    ProxyAdminOwnedBase,
    ReinitializableBase
{
    /// @notice Semantic version.
    /// @custom:semver 1.1.0
    string public constant version = "1.1.0";

    /// @notice Address of the TEE batcher whose signatures may authenticate batches.
    address public teeBatcher;

    /// @notice Address of the Espresso TEE Verifier contract.
    IEspressoTEEVerifier public espressoTEEVerifier;

    /// @notice Flag indicating which batcher is currently active.
    /// @dev When true the TEE batcher is active; when false the non-TEE batcher is active.
    bool public activeIsTee;

    /// @notice The SystemConfig contract, used to check the paused status.
    ISystemConfig public systemConfig;

    /// @notice Constructor disables initializers on implementation
    constructor() ReinitializableBase(1) {
        _disableInitializers();
    }

    function initialize(
        IEspressoTEEVerifier _espressoTEEVerifier,
        address _teeBatcher,
        ISystemConfig _systemConfig,
        address _owner
    )
        external
        reinitializer(initVersion())
    {
        // Initialization transactions must come from the ProxyAdmin or its owner.
        _assertOnlyProxyAdminOrProxyAdminOwner();

        // Initialize OwnableWithGuardians with the provided owner address
        __OwnableWithGuardians_init(_owner);

        if (_teeBatcher == address(0)) revert InvalidAddress(_teeBatcher);
        if (address(_systemConfig) == address(0)) revert InvalidAddress(address(_systemConfig));
        if (address(_espressoTEEVerifier) == address(0)) {
            revert InvalidAddress(address(_espressoTEEVerifier));
        }

        espressoTEEVerifier = _espressoTEEVerifier;
        teeBatcher = _teeBatcher;
        systemConfig = _systemConfig;
        // By default, start with the TEE batcher active.
        activeIsTee = true;
    }

    /// @notice Returns the owner of the contract.
    function owner() public view override(IBatchAuthenticator, OwnableUpgradeable) returns (address) {
        return super.owner();
    }

    /// @notice Getter for the current paused status.
    function paused() public view returns (bool) {
        return systemConfig.paused();
    }

    /// @notice Toggles the active batcher between the TEE and non-TEE batcher.
    function switchBatcher() external onlyGuardianOrOwner {
        activeIsTee = !activeIsTee;
        emit BatcherSwitched(activeIsTee);
    }

    /// @notice Updates the TEE batcher address.
    function setTeeBatcher(address _newTeeBatcher) external onlyOwner {
        if (_newTeeBatcher == address(0)) revert InvalidAddress(_newTeeBatcher);
        address oldTeeBatcher = teeBatcher;
        teeBatcher = _newTeeBatcher;
        emit TeeBatcherUpdated(oldTeeBatcher, _newTeeBatcher);
    }

    function authenticateBatchInfo(bytes32 commitment, bytes calldata _signature) external {
        if (paused()) revert BatchAuthenticator_Paused();

        // Setting TEEType as Nitro because OP integration only supports AWS Nitro currently
        espressoTEEVerifier.verify(_signature, commitment, IEspressoTEEVerifier.TeeType.NITRO, ServiceType.BatchPoster);

        emit BatchInfoAuthenticated(commitment);
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external {
        if (paused()) revert BatchAuthenticator_Paused();

        espressoTEEVerifier.registerService(
            attestationTbs, signature, IEspressoTEEVerifier.TeeType.NITRO, ServiceType.BatchPoster
        );
        emit SignerRegistrationInitiated(msg.sender);
    }

    /// @notice Returns the address of the Nitro TEE validator.
    function nitroValidator() external view returns (address) {
        return address(espressoTEEVerifier.espressoNitroTEEVerifier());
    }

    // NOTE: This contract only provides authenticateBatchInfo (which emits BatchInfoAuthenticated events)
    // and signer management. Batch authentication is performed off-chain by the derivation pipeline,
    // which scans L1 receipts for BatchInfoAuthenticated events in a lookback window.
    // Batch data is sent as plain transactions to the BatchInbox EOA address.
}
