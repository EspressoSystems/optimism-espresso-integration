// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Test } from "forge-std/Test.sol";
import { console2 as console } from "forge-std/console2.sol";

import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { EspressoTEEVerifierMock } from "@espresso-tee-contracts/mocks/EspressoTEEVerifier.sol";
import { IProxy } from "interfaces/universal/IProxy.sol";
import { EIP1967Helper } from "test/mocks/EIP1967Helper.sol";

import { Config } from "scripts/libraries/Config.sol";
import { Chains } from "scripts/libraries/Chains.sol";

/// @notice Mock that implements IEspressoTEEVerifier and IEspressoNitroTEEVerifier by using
///         composition with EspressoTEEVerifierMock to reuse its logic.
///         Supports only the Nitro TEE verifier.
contract MockEspressoTEEVerifier is IEspressoTEEVerifier, IEspressoNitroTEEVerifier {
    EspressoTEEVerifierMock private _mock;

    constructor() {
        _mock = new EspressoTEEVerifierMock();
    }

    function espressoNitroTEEVerifier() external view override returns (IEspressoNitroTEEVerifier) {
        return this;
    }

    function espressoSGXTEEVerifier() external pure override returns (IEspressoSGXTEEVerifier) {
        return IEspressoSGXTEEVerifier(address(0));
    }

    function verify(
        bytes memory signature,
        bytes32 userDataHash,
        TeeType teeType
    )
        external
        view
        override
        returns (bool)
    {
        if (teeType != TeeType.NITRO) {
            return false;
        }
        EspressoTEEVerifierMock.TeeType mockTeeType = EspressoTEEVerifierMock.TeeType(uint8(teeType));
        return _mock.verify(signature, userDataHash, mockTeeType);
    }

    function registerSigner(bytes calldata attestation, bytes calldata data, TeeType teeType) external override {
        require(teeType == TeeType.NITRO, "MockEspressoTEEVerifier: only NITRO supported");
        EspressoTEEVerifierMock.TeeType mockTeeType = EspressoTEEVerifierMock.TeeType(uint8(teeType));
        _mock.registerSigner(attestation, data, mockTeeType);
    }

    function registeredSigners(address signer, TeeType teeType) external view override returns (bool) {
        if (teeType != TeeType.NITRO) {
            return false;
        }
        EspressoTEEVerifierMock.TeeType mockTeeType = EspressoTEEVerifierMock.TeeType(uint8(teeType));
        return _mock.registeredSigners(signer, mockTeeType);
    }

    function registeredEnclaveHashes(bytes32, TeeType) external pure override returns (bool) {
        return false;
    }

    function setEspressoSGXTEEVerifier(IEspressoSGXTEEVerifier) external pure override {
        // No-op: SGX is not supported.
    }

    function setEspressoNitroTEEVerifier(IEspressoNitroTEEVerifier) external pure override {
        // No-op: this contract can only be used as the Nitro TEE verifier.
    }

    function registeredSigners(address signer) external view override returns (bool) {
        return _mock.registeredSigner(signer);
    }

    function registeredEnclaveHash(bytes32) external pure override returns (bool) {
        return false;
    }

    function registerSigner(bytes calldata, bytes calldata) external pure override {
        // No-op: registration should go through registerSigner(bytes, bytes, TeeType)
    }

    function setEnclaveHash(bytes32, bool) external pure override { }

    function deleteRegisteredSigners(address[] memory) external pure override { }

    /// @notice Test helper to directly set registered signer status.
    function setRegisteredSigner(address signer, bool value) external {
        if (value) {
            bytes memory data = abi.encodePacked(signer);
            this.registerSigner("", data, TeeType.NITRO);
        } else {
            // For false, we can't unregister through the mock's interface,
            // but tests only set to true, so this is fine.
            revert("MockEspressoTEEVerifier: unregistering not supported");
        }
    }
}

/// @notice Tests for the upgradeable BatchAuthenticator contract using the Transparent Proxy pattern.
contract BatchAuthenticator_Test is Test {
    address public deployer = address(0xABCD);
    address public proxyAdminOwner = address(0xBEEF);
    address public unauthorized = address(0xDEAD);

    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);

    MockEspressoTEEVerifier public teeVerifier;
    BatchAuthenticator public implementation;
    ProxyAdmin public proxyAdmin;

    function setUp() public {
        // Deploy the mock TEE verifier and the authenticator implementation.
        teeVerifier = new MockEspressoTEEVerifier();
        implementation = new BatchAuthenticator();

        // Deploy the proxy admin.
        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin(proxyAdminOwner);
    }

    /// @notice Create and initialize a proxy.
    function _deployAndInitializeProxy() internal returns (BatchAuthenticator) {
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.prank(proxyAdminOwner);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize, (IEspressoTEEVerifier(address(teeVerifier)), teeBatcher, nonTeeBatcher)
        );
        vm.prank(proxyAdminOwner);
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(implementation), initData);

        return BatchAuthenticator(address(proxy));
    }

    /// @notice Test that the initialization can only be called once.
    function test_constructor_revertsWhenAlreadyInitialized() external {
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.prank(proxyAdminOwner);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize, (IEspressoTEEVerifier(address(teeVerifier)), teeBatcher, nonTeeBatcher)
        );

        // First initialization succeeds.
        vm.prank(proxyAdminOwner);
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(implementation), initData);

        // Second initialization should revert.
        vm.prank(proxyAdminOwner);
        vm.expectRevert();
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(implementation), initData);
    }

    /// @notice Test that initialize reverts when teeBatcher is zero.
    function test_constructor_revertsWhenTeeBatcherIsZero() external {
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.prank(proxyAdminOwner);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize, (IEspressoTEEVerifier(address(teeVerifier)), address(0), nonTeeBatcher)
        );

        vm.prank(proxyAdminOwner);
        vm.expectRevert("Proxy: delegatecall to new implementation contract failed");
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(implementation), initData);
    }

    /// @notice Test that initialize reverts when nonTeeBatcher is zero.
    function test_constructor_revertsWhenNonTeeBatcherIsZero() external {
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.prank(proxyAdminOwner);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize, (IEspressoTEEVerifier(address(teeVerifier)), teeBatcher, address(0))
        );

        vm.prank(proxyAdminOwner);
        vm.expectRevert("Proxy: delegatecall to new implementation contract failed");
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(implementation), initData);
    }

    /// @notice Test that initialize reverts when verifier is zero.
    function test_constructor_revertsWhenVerifierIsZero() external {
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.prank(proxyAdminOwner);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        bytes memory initData =
            abi.encodeCall(BatchAuthenticator.initialize, (IEspressoTEEVerifier(address(0)), teeBatcher, nonTeeBatcher));

        vm.prank(proxyAdminOwner);
        vm.expectRevert("Proxy: delegatecall to new implementation contract failed");
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(implementation), initData);
    }

    /// @notice Test that initialize succeeds with valid addresses.
    function test_constructor_succeedsWithValidAddresses() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        assertEq(address(authenticator.espressoTEEVerifier()), address(teeVerifier));
        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
        assertTrue(authenticator.activeIsTee());
    }

    /// @notice Test that switchBatcher can only be called by ProxyAdmin owner.
    function test_switchBatcher_onlyProxyAdminOwner() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        // ProxyAdmin owner can switch.
        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();
        assertFalse(authenticator.activeIsTee());

        // Switch back.
        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();
        assertTrue(authenticator.activeIsTee());

        // Unauthorized cannot switch.
        vm.prank(unauthorized);
        vm.expectRevert();
        authenticator.switchBatcher();

        // ProxyAdmin cannot switch.
        vm.prank(address(proxyAdmin));
        vm.expectRevert();
        authenticator.switchBatcher();
    }

    /// @notice Test that authenticateBatchInfo works correctly.
    function test_authenticateBatchInfo_succeeds() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        uint256 privateKey = 1;
        address signer = vm.addr(privateKey);
        bytes32 commitment = keccak256("test commitment");

        // Register signer.
        teeVerifier.setRegisteredSigner(signer, true);

        // Create signature.
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, commitment);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Authenticate.
        vm.expectEmit(true, true, false, false);
        emit BatchInfoAuthenticated(commitment, signer);

        authenticator.authenticateBatchInfo(commitment, signature);

        assertTrue(authenticator.validBatchInfo(commitment));
    }

    /// @notice Test that registerSigner works correctly.
    function test_registerSigner_succeeds() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        bytes memory attestationTbs = "test attestation";
        address signer = address(0x1234);
        bytes memory signature = abi.encodePacked(signer);

        vm.expectEmit(true, false, false, true);
        emit SignerRegistrationInitiated(address(this));

        authenticator.registerSigner(attestationTbs, signature);
    }

    /// @notice Test that setTeeBatcher can only be called by ProxyAdmin owner.
    function test_setTeeBatcher_onlyProxyAdminOwner() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();
        address newTeeBatcher = address(0x9999);

        // ProxyAdmin owner can set.
        vm.expectEmit(true, true, false, false);
        emit TeeBatcherUpdated(teeBatcher, newTeeBatcher);
        vm.prank(proxyAdminOwner);
        authenticator.setTeeBatcher(newTeeBatcher);
        assertEq(authenticator.teeBatcher(), newTeeBatcher);

        // Unauthorized cannot set.
        vm.prank(unauthorized);
        vm.expectRevert();
        authenticator.setTeeBatcher(address(0x7777));

        // ProxyAdmin cannot set.
        vm.prank(address(proxyAdmin));
        vm.expectRevert();
        authenticator.setTeeBatcher(address(0x8888));
    }

    /// @notice Test that setTeeBatcher reverts when zero address is provided.
    function test_setTeeBatcher_revertsWhenZeroAddress() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        vm.prank(proxyAdminOwner);
        vm.expectRevert(abi.encodeWithSelector(IBatchAuthenticator.InvalidAddress.selector, address(0)));
        authenticator.setTeeBatcher(address(0));
    }

    /// @notice Test that setNonTeeBatcher can only be called by ProxyAdmin owner.
    function test_setNonTeeBatcher_onlyProxyAdminOwner() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();
        address newNonTeeBatcher = address(0xAAAA);

        // ProxyAdmin owner can set.
        vm.expectEmit(true, true, false, false);
        emit NonTeeBatcherUpdated(nonTeeBatcher, newNonTeeBatcher);
        vm.prank(proxyAdminOwner);
        authenticator.setNonTeeBatcher(newNonTeeBatcher);
        assertEq(authenticator.nonTeeBatcher(), newNonTeeBatcher);

        // Unauthorized cannot set.
        vm.prank(unauthorized);
        vm.expectRevert();
        authenticator.setNonTeeBatcher(address(0xCCCC));

        // ProxyAdmin cannot set.
        vm.prank(address(proxyAdmin));
        vm.expectRevert();
        authenticator.setNonTeeBatcher(address(0xBBBB));
    }

    /// @notice Test that setNonTeeBatcher reverts when zero address is provided.
    function test_setNonTeeBatcher_revertsWhenZeroAddress() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        vm.prank(proxyAdminOwner);
        vm.expectRevert(abi.encodeWithSelector(IBatchAuthenticator.InvalidAddress.selector, address(0)));
        authenticator.setNonTeeBatcher(address(0));
    }

    /// @notice Test upgrade to new implementation with comprehensive state preservation.
    function test_upgrade_preservesState() external {
        // Create and initialize a proxy.
        BatchAuthenticator authenticator = _deployAndInitializeProxy();
        Proxy proxy = Proxy(payable(address(authenticator)));

        // Set up initial state.
        bytes32 commitment = keccak256("test commitment");
        uint256 privateKey = 1;
        address signer = vm.addr(privateKey);
        teeVerifier.setRegisteredSigner(signer, true);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, commitment);
        bytes memory signature = abi.encodePacked(r, s, v);
        authenticator.authenticateBatchInfo(commitment, signature);
        assertTrue(authenticator.validBatchInfo(commitment));

        // Switch batcher to test boolean flag preservation.
        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();
        assertFalse(authenticator.activeIsTee());

        // Deploy new implementation and upgrade.
        BatchAuthenticator newImpl = new BatchAuthenticator();
        vm.prank(proxyAdminOwner);
        proxyAdmin.upgrade(payable(address(proxy)), address(newImpl));

        // Verify implementation changed.
        address newImplementation = EIP1967Helper.getImplementation(address(proxy));
        assertEq(newImplementation, address(newImpl));

        // Verify state is preserved.
        assertEq(address(authenticator.espressoTEEVerifier()), address(teeVerifier));
        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
        assertTrue(authenticator.validBatchInfo(commitment));
        assertFalse(authenticator.activeIsTee());
    }

    // Event declarations for expectEmit.
    event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer);
    event SignerRegistrationInitiated(address indexed caller);
    event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher);
    event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher);
}

/// @notice Fork tests for BatchAuthenticator on Sepolia.
contract BatchAuthenticator_Fork_Test is Test {
    address public proxyAdminOwner = address(0xBEEF);
    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);

    MockEspressoTEEVerifier public teeVerifier;
    BatchAuthenticator public implementation;
    Proxy public proxy;
    ProxyAdmin public proxyAdmin;
    BatchAuthenticator public authenticator;

    function setUp() public {
        // Create a fork of Sepolia using the execution layer RPC endpoint.
        string memory forkUrl = "https://theserversroom.com/sepolia/54cmzzhcj1o/";
        vm.createSelectFork(forkUrl);

        // Verify we're on Sepolia.
        require(block.chainid == Chains.Sepolia, "Fork test must run on Sepolia");
        console.log("Forked Sepolia at block:", block.number);

        // Deploy mock TEE verifier and authenticator implementation.
        teeVerifier = new MockEspressoTEEVerifier();
        implementation = new BatchAuthenticator();

        // Deploy proxy admin and proxy.
        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin(proxyAdminOwner);
        proxy = new Proxy(address(proxyAdmin));
        vm.prank(proxyAdminOwner);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        // Initialize the proxy.
        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize, (IEspressoTEEVerifier(address(teeVerifier)), teeBatcher, nonTeeBatcher)
        );
        vm.prank(proxyAdminOwner);
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(implementation), initData);

        // Get the proxied contract instance.
        authenticator = BatchAuthenticator(address(proxy));
    }

    /// @notice Test deployment and initialization on Sepolia fork.
    function testFork_deployment_succeeds() external view {
        assertEq(address(authenticator.espressoTEEVerifier()), address(teeVerifier));
        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
        assertTrue(authenticator.activeIsTee());
        assertEq(authenticator.version(), "1.0.0");

        // Verify proxy admin.
        address admin = EIP1967Helper.getAdmin(address(proxy));
        assertEq(admin, address(proxyAdmin));
    }

    /// @notice Test switchBatcher on Sepolia fork.
    function testFork_switchBatcher_succeeds() external {
        assertTrue(authenticator.activeIsTee());

        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();

        assertFalse(authenticator.activeIsTee());

        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();

        assertTrue(authenticator.activeIsTee());
    }

    /// @notice Test authenticateBatchInfo on Sepolia fork.
    function testFork_authenticateBatchInfo_succeeds() external {
        bytes32 commitment = keccak256("test commitment on sepolia");

        // Create a signature.
        uint256 privateKey = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
        address signer = vm.addr(privateKey);

        // Register the signer.
        teeVerifier.setRegisteredSigner(signer, true);

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, commitment);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Authenticate.
        vm.expectEmit(true, true, false, false);
        emit BatchInfoAuthenticated(commitment, signer);
        authenticator.authenticateBatchInfo(commitment, signature);

        assertTrue(authenticator.validBatchInfo(commitment));
    }

    /// @notice Test upgrade on Sepolia fork.
    function testFork_upgrade_preservesState() external {
        // Initialize the authenticator.
        bytes32 commitment = keccak256("test commitment");
        uint256 privateKey = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
        address signer = vm.addr(privateKey);

        // Register the signer.
        teeVerifier.setRegisteredSigner(signer, true);

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, commitment);
        bytes memory signature = abi.encodePacked(r, s, v);
        authenticator.authenticateBatchInfo(commitment, signature);
        assertTrue(authenticator.validBatchInfo(commitment));

        // Switch batcher
        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();
        assertFalse(authenticator.activeIsTee());

        // Deploy new implementation and upgrade.
        BatchAuthenticator newImpl = new BatchAuthenticator();
        vm.prank(proxyAdminOwner);
        proxyAdmin.upgrade(payable(address(proxy)), address(newImpl));

        // Verify state is preserved.
        assertTrue(authenticator.validBatchInfo(commitment));
        assertFalse(authenticator.activeIsTee());
        assertEq(address(authenticator.espressoTEEVerifier()), address(teeVerifier));
        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
    }

    /// @notice Test that contract works with real Sepolia state
    function testFork_integrationWithSepolia() external view {
        // Verify we're on Sepolia.
        assertEq(block.chainid, Chains.Sepolia);

        // Verify contract is functional.
        assertEq(authenticator.version(), "1.0.0");
        assertTrue(authenticator.activeIsTee());

        // Verify the fork is working by testing that we can read the block number.
        uint256 blockNum = block.number;
        assertGt(blockNum, 0);
        console.log("Sepolia block number:", blockNum);
    }

    // Event declarations for expectEmit.
    event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer);
    event SignerRegistrationInitiated(address indexed caller);
    event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher);
    event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher);
}
