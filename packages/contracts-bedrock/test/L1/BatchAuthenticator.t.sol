// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Testing
import { Test } from "forge-std/Test.sol";

// Contracts
import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";

contract MockNitroTEEVerifier is IEspressoNitroTEEVerifier {
    mapping(address => bool) private _registeredSigners;

    function registeredSigners(address signer) external view override returns (bool) {
        return _registeredSigners[signer];
    }

    function registeredEnclaveHash(bytes32) external pure override returns (bool) {
        return false;
    }

    function registerSigner(bytes calldata, bytes calldata) external pure override { }

    function setEnclaveHash(bytes32, bool) external pure override { }

    function deleteRegisteredSigners(address[] memory) external pure override { }

    // Test helper
    function setRegisteredSigner(address signer, bool value) external {
        _registeredSigners[signer] = value;
    }
}

contract MockEspressoTEEVerifier is IEspressoTEEVerifier {
    IEspressoNitroTEEVerifier public nitro;
    IEspressoSGXTEEVerifier public sgx;

    constructor(IEspressoNitroTEEVerifier _nitro) {
        nitro = _nitro;
        sgx = IEspressoSGXTEEVerifier(address(0));
    }

    function espressoNitroTEEVerifier() external view override returns (IEspressoNitroTEEVerifier) {
        return nitro;
    }

    function espressoSGXTEEVerifier() external view override returns (IEspressoSGXTEEVerifier) {
        return sgx;
    }

    function verify(bytes memory, bytes32, TeeType) external pure override returns (bool) {
        return true;
    }

    function registerSigner(bytes calldata, bytes calldata, TeeType) external pure override { }

    function registeredSigners(address, TeeType) external pure override returns (bool) {
        return false;
    }

    function registeredEnclaveHashes(bytes32, TeeType) external pure override returns (bool) {
        return false;
    }

    function setEspressoSGXTEEVerifier(IEspressoSGXTEEVerifier _sgx) external override {
        sgx = _sgx;
    }

    function setEspressoNitroTEEVerifier(IEspressoNitroTEEVerifier _nitro) external override {
        nitro = _nitro;
    }
}

/// @title BatchAuthenticator_SwitchBatcher_Test
/// @notice Tests ownership restrictions on BatchAuthenticator switchBatcher behavior
contract BatchAuthenticator_SwitchBatcher_Test is Test {
    address public deployer = address(0xABCD);
    address public unauthorized = address(0xDEAD);

    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);

    MockNitroTEEVerifier public nitroVerifier;
    MockEspressoTEEVerifier public teeVerifier;
    BatchAuthenticator public authenticator;

    function setUp() public {
        nitroVerifier = new MockNitroTEEVerifier();
        teeVerifier = new MockEspressoTEEVerifier(nitroVerifier);

        vm.prank(deployer);
        authenticator =
            new BatchAuthenticator(IEspressoTEEVerifier(address(teeVerifier)), teeBatcher, nonTeeBatcher, deployer);
    }

    /// @notice Test that only the owner can switch the active batcher
    function test_switchBatcher_revertsForNonOwner() external {
        // Owner can switch batcher successfully.
        vm.startPrank(deployer);
        bool initialIsTee = authenticator.activeIsTee();
        authenticator.switchBatcher();
        assertEq(authenticator.activeIsTee(), !initialIsTee, "owner should be able to switch batcher");
        vm.stopPrank();

        // Non-owner cannot switch batcher.
        vm.startPrank(unauthorized);
        vm.expectRevert("Ownable: caller is not the owner");
        authenticator.switchBatcher();
        vm.stopPrank();
    }
}

contract BatchAuthenticator_Constructor_Test is Test {
    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);

    address public owner = address(0xBEEF);

    MockNitroTEEVerifier public nitroVerifier;
    MockEspressoTEEVerifier public teeVerifier;

    function setUp() public {
        nitroVerifier = new MockNitroTEEVerifier();
        teeVerifier = new MockEspressoTEEVerifier(nitroVerifier);
    }

    function test_constructor_revertsWhenTeeBatcherIsZero() external {
        vm.expectRevert("BatchAuthenticator: zero tee batcher");
        new BatchAuthenticator(IEspressoTEEVerifier(address(teeVerifier)), address(0), nonTeeBatcher, owner);
    }

    function test_constructor_revertsWhenNonTeeBatcherIsZero() external {
        vm.expectRevert("BatchAuthenticator: zero non-tee batcher");
        new BatchAuthenticator(IEspressoTEEVerifier(address(teeVerifier)), teeBatcher, address(0), owner);
    }

    function test_constructor_succeedsWithValidAddresses() external {
        BatchAuthenticator authenticator =
            new BatchAuthenticator(IEspressoTEEVerifier(address(teeVerifier)), teeBatcher, nonTeeBatcher, owner);
        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
    }
}
