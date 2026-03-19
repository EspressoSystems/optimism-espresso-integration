// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { ServiceType } from "@espresso-tee-contracts/types/Types.sol";
import { INitroEnclaveVerifier } from "aws-nitro-enclave-attestation/interfaces/INitroEnclaveVerifier.sol";
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/// @notice Mock implementation of IEspressoNitroTEEVerifier for testing without real attestation verification.
///         Used by deployment scripts and tests.
contract MockEspressoNitroTEEVerifier is IEspressoNitroTEEVerifier {
    address internal _teeVerifier;
    mapping(ServiceType => mapping(address => bool)) private _registeredServices;

    constructor() { }

    function isSignerValid(address signer, ServiceType service) external view override returns (bool) {
        // Special condition for test TestE2eDevnetWithUnattestedBatcherKey
        if (signer == address(0xe16d5c4080C0faD6D2Ef4eb07C657674a217271C)) {
            return false;
        }
        // If signer was explicitly registered, return true
        if (_registeredServices[service][signer]) {
            return true;
        }
        // Default permissive behavior for deployment scripts (when no signers registered)
        // This allows the mock to work in both test (explicit registration) and deploy (permissive) modes
        return true;
    }

    function registeredEnclaveHash(bytes32, ServiceType) external pure override returns (bool) {
        return true;
    }

    function registerService(bytes calldata attestation, bytes calldata, ServiceType service) external override {
        if (attestation.length >= 20) {
            address signer = address(uint160(bytes20(attestation[:20])));
            _registeredServices[service][signer] = true;
        }
    }

    function setEnclaveHash(bytes32, bool, ServiceType) external override { }

    function deleteEnclaveHashes(bytes32[] memory, ServiceType) external override { }

    function setNitroEnclaveVerifier(address) external override { }

    function nitroEnclaveVerifier() external pure override returns (INitroEnclaveVerifier) {
        return INitroEnclaveVerifier(address(0));
    }

    function teeVerifier() external view override returns (address) {
        return _teeVerifier;
    }

    /// @notice Test helper to directly set registered signer status for a service type.
    function setRegisteredSigner(address signer, bool value, ServiceType service) external {
        _registeredServices[service][signer] = value;
    }

    /// @notice Test helper to directly set registered signer status (defaults to BatchPoster).
    function setRegisteredSigner(address signer, bool value) external {
        _registeredServices[ServiceType.BatchPoster][signer] = value;
    }
}

/// @notice Mock implementation of IEspressoTEEVerifier for testing.
///         Can optionally wrap a MockEspressoNitroTEEVerifier or act as its own Nitro verifier.
contract MockEspressoTEEVerifier is IEspressoTEEVerifier, IEspressoNitroTEEVerifier {
    IEspressoNitroTEEVerifier private _nitroVerifier;
    mapping(ServiceType => mapping(address => bool)) private _registeredServices;
    bool private _useExternalNitroVerifier;

    /// @notice Constructor that optionally takes an external Nitro verifier.
    /// @param nitroVerifier_ The external Nitro verifier to use. If address(0), acts as standalone.
    constructor(IEspressoNitroTEEVerifier nitroVerifier_) {
        if (address(nitroVerifier_) != address(0)) {
            _nitroVerifier = nitroVerifier_;
            _useExternalNitroVerifier = true;
        } else {
            _useExternalNitroVerifier = false;
        }
    }

    // ============ IEspressoTEEVerifier Implementation ============

    function espressoNitroTEEVerifier() external view override returns (IEspressoNitroTEEVerifier) {
        if (_useExternalNitroVerifier) {
            return _nitroVerifier;
        }
        return this;
    }

    function espressoSGXTEEVerifier() external pure override returns (IEspressoSGXTEEVerifier) {
        return IEspressoSGXTEEVerifier(address(0));
    }

    function verify(
        bytes memory signature,
        bytes32 userDataHash,
        TeeType teeType,
        ServiceType service
    )
        external
        view
        override
        returns (bool)
    {
        if (teeType != TeeType.NITRO) {
            revert InvalidSignature();
        }
        address signer = ECDSA.recover(userDataHash, signature);
        IEspressoNitroTEEVerifier nitroVerifier =
            _useExternalNitroVerifier ? _nitroVerifier : IEspressoNitroTEEVerifier(address(this));
        if (!nitroVerifier.isSignerValid(signer, service)) {
            revert InvalidSignature();
        }
        return true;
    }

    function registerService(
        bytes calldata attestation,
        bytes calldata,
        TeeType teeType,
        ServiceType serviceType
    )
        external
        override
    {
        require(teeType == TeeType.NITRO, "MockEspressoTEEVerifier: only NITRO supported");
        if (attestation.length >= 20) {
            address signer = address(uint160(bytes20(attestation[:20])));
            _registeredServices[serviceType][signer] = true;
        }
    }

    function registeredEnclaveHashes(bytes32, TeeType, ServiceType) external pure override returns (bool) {
        return false;
    }

    function setEspressoSGXTEEVerifier(IEspressoSGXTEEVerifier) external override { }

    function setEspressoNitroTEEVerifier(IEspressoNitroTEEVerifier verifier) external override {
        _nitroVerifier = verifier;
        _useExternalNitroVerifier = address(verifier) != address(0);
    }

    function setEnclaveHash(bytes32, bool, TeeType, ServiceType) external override { }

    function deleteEnclaveHashes(bytes32[] memory, TeeType, ServiceType) external override { }

    function setQuoteVerifier(address) external override { }

    function setNitroEnclaveVerifier(address) external override(IEspressoNitroTEEVerifier, IEspressoTEEVerifier) { }

    // ============ IEspressoNitroTEEVerifier Implementation (for standalone mode) ============

    function isSignerValid(address signer, ServiceType service) external view override returns (bool) {
        return _registeredServices[service][signer];
    }

    function registeredEnclaveHash(bytes32, ServiceType) external pure override returns (bool) {
        return false;
    }

    function registerService(bytes calldata attestation, bytes calldata, ServiceType service) external override {
        if (attestation.length >= 20) {
            address signer = address(uint160(bytes20(attestation[:20])));
            _registeredServices[service][signer] = true;
        }
    }

    function setEnclaveHash(bytes32, bool, ServiceType) external pure override { }

    function deleteEnclaveHashes(bytes32[] memory, ServiceType) external pure override { }

    function nitroEnclaveVerifier() external pure override returns (INitroEnclaveVerifier) {
        return INitroEnclaveVerifier(address(0));
    }

    function teeVerifier() external view override returns (address) {
        return address(this);
    }

    // ============ Test Helpers ============

    /// @notice Test helper to directly set registered signer status.
    function setRegisteredSigner(address signer, bool value) external {
        if (value) {
            _registeredServices[ServiceType.BatchPoster][signer] = true;
        } else {
            revert("MockEspressoTEEVerifier: unregistering not supported");
        }
    }
}
