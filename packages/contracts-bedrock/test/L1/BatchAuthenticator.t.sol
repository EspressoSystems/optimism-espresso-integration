// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

// Testing
import { Test } from "forge-std/Test.sol";

// Contracts
import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { INitroEnclaveVerifier } from "aws-nitro-enclave-attestation/interfaces/INitroEnclaveVerifier.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { ServiceType } from "@espresso-tee-contracts/types/Types.sol";

contract MockNitroTEEVerifier is IEspressoNitroTEEVerifier {
    mapping(address => mapping(ServiceType => bool)) private _registeredServices;

    function isSignerValid(address signer, ServiceType serviceType) external view override returns (bool) {
        return _registeredServices[signer][serviceType];
    }

    function registeredEnclaveHash(bytes32, ServiceType) external pure override returns (bool) {
        return false;
    }

    function registerService(bytes calldata, bytes calldata, ServiceType) external pure override { }

    function setEnclaveHash(bytes32, bool, ServiceType) external pure override { }

    function deleteEnclaveHashes(bytes32[] memory, ServiceType) external pure override { }

    function setNitroEnclaveVerifier(address) external pure override { }

    function nitroEnclaveVerifier() external pure override returns (INitroEnclaveVerifier) {
        return INitroEnclaveVerifier(address(0));
    }

    function teeVerifier() external pure override returns (address) {
        return address(0);
    }

    // Test helper
    function setRegisteredService(address signer, ServiceType serviceType, bool value) external {
        _registeredServices[signer][serviceType] = value;
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

    function verify(bytes memory, bytes32, TeeType, ServiceType) external pure override returns (bool) {
        return true;
    }

    function registerService(bytes calldata, bytes calldata, TeeType, ServiceType) external pure override { }

    function registeredEnclaveHashes(bytes32, TeeType, ServiceType) external pure override returns (bool) {
        return false;
    }

    function setEspressoSGXTEEVerifier(IEspressoSGXTEEVerifier _sgx) external override {
        sgx = _sgx;
    }

    function setEspressoNitroTEEVerifier(IEspressoNitroTEEVerifier _nitro) external override {
        nitro = _nitro;
    }

    function setEnclaveHash(bytes32, bool, TeeType, ServiceType) external pure override { }

    function deleteEnclaveHashes(bytes32[] memory, TeeType, ServiceType) external pure override { }

    function setQuoteVerifier(address) external pure override { }

    function setNitroEnclaveVerifier(address) external pure override { }
}

/// @title BatchAuthenticator_SwitchBatcher_Test
/// @notice Tests ownership and guardian restrictions on BatchAuthenticator switchBatcher behavior
contract BatchAuthenticator_SwitchBatcher_Test is Test {
    address public deployer = address(0xABCD);
    address public unauthorized = address(0xDEAD);
    address public guardian = address(0xFEED);

    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);

    MockNitroTEEVerifier public nitroVerifier;
    MockEspressoTEEVerifier public teeVerifier;
    BatchAuthenticator public authenticator;

    function setUp() public {
        nitroVerifier = new MockNitroTEEVerifier();
        teeVerifier = new MockEspressoTEEVerifier(nitroVerifier);

        // Deploy implementation
        BatchAuthenticator impl = new BatchAuthenticator();

        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            BatchAuthenticator.initialize.selector,
            IEspressoTEEVerifier(address(teeVerifier)),
            teeBatcher,
            nonTeeBatcher,
            deployer
        );

        // Deploy proxy
        ERC1967Proxy proxy = new ERC1967Proxy(address(impl), initData);
        authenticator = BatchAuthenticator(address(proxy));
    }

    /// @notice Test that only the owner or guardian can switch the active batcher
    function test_switchBatcher_revertsForNonOwner() external {
        // Owner can switch batcher successfully.
        vm.startPrank(deployer);
        bool initialIsTee = authenticator.activeIsTee();
        authenticator.switchBatcher();
        assertEq(authenticator.activeIsTee(), !initialIsTee, "owner should be able to switch batcher");
        vm.stopPrank();

        // Non-owner cannot switch batcher.
        vm.startPrank(unauthorized);
        vm.expectRevert();
        authenticator.switchBatcher();
        vm.stopPrank();
    }

    /// @notice Test that guardians can switch the active batcher
    function test_switchBatcher_guardianCanSwitch() external {
        // Add guardian
        vm.prank(deployer);
        authenticator.addGuardian(guardian);

        // Guardian can switch batcher
        vm.startPrank(guardian);
        bool initialIsTee = authenticator.activeIsTee();
        authenticator.switchBatcher();
        assertEq(authenticator.activeIsTee(), !initialIsTee, "guardian should be able to switch batcher");
        vm.stopPrank();
    }

    /// @notice Test guardian management functions
    function test_guardianManagement() external {
        vm.startPrank(deployer);

        // Add guardian
        authenticator.addGuardian(guardian);
        assertTrue(authenticator.isGuardian(guardian), "should be guardian");
        assertEq(authenticator.guardianCount(), 1, "should have 1 guardian");

        // Get guardians list
        address[] memory guardians = authenticator.getGuardians();
        assertEq(guardians.length, 1, "should have 1 guardian in list");
        assertEq(guardians[0], guardian, "guardian address should match");

        // Remove guardian
        authenticator.removeGuardian(guardian);
        assertFalse(authenticator.isGuardian(guardian), "should not be guardian");
        assertEq(authenticator.guardianCount(), 0, "should have 0 guardians");

        vm.stopPrank();
    }
}

contract BatchAuthenticator_Initialize_Test is Test {
    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);

    address public owner = address(0xBEEF);

    MockNitroTEEVerifier public nitroVerifier;
    MockEspressoTEEVerifier public teeVerifier;

    function setUp() public {
        nitroVerifier = new MockNitroTEEVerifier();
        teeVerifier = new MockEspressoTEEVerifier(nitroVerifier);
    }

    function test_initialize_revertsWhenTeeBatcherIsZero() external {
        BatchAuthenticator impl = new BatchAuthenticator();

        bytes memory initData = abi.encodeWithSelector(
            BatchAuthenticator.initialize.selector,
            IEspressoTEEVerifier(address(teeVerifier)),
            address(0),
            nonTeeBatcher,
            owner
        );

        vm.expectRevert("BatchAuthenticator: zero tee batcher");
        new ERC1967Proxy(address(impl), initData);
    }

    function test_initialize_revertsWhenNonTeeBatcherIsZero() external {
        BatchAuthenticator impl = new BatchAuthenticator();

        bytes memory initData = abi.encodeWithSelector(
            BatchAuthenticator.initialize.selector,
            IEspressoTEEVerifier(address(teeVerifier)),
            teeBatcher,
            address(0),
            owner
        );

        vm.expectRevert("BatchAuthenticator: zero non-tee batcher");
        new ERC1967Proxy(address(impl), initData);
    }

    function test_initialize_succeedsWithValidAddresses() external {
        BatchAuthenticator impl = new BatchAuthenticator();

        bytes memory initData = abi.encodeWithSelector(
            BatchAuthenticator.initialize.selector,
            IEspressoTEEVerifier(address(teeVerifier)),
            teeBatcher,
            nonTeeBatcher,
            owner
        );

        ERC1967Proxy proxy = new ERC1967Proxy(address(impl), initData);
        BatchAuthenticator authenticator = BatchAuthenticator(address(proxy));

        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
        assertEq(authenticator.owner(), owner);
    }
}
