// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Test } from "forge-std/Test.sol";
import { console2 as console } from "forge-std/console2.sol";

import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { IProxy } from "interfaces/universal/IProxy.sol";
import { EIP1967Helper } from "test/mocks/EIP1967Helper.sol";

import { Config } from "scripts/libraries/Config.sol";
import { Chains } from "scripts/libraries/Chains.sol";

/// @notice Mock implementation of IEspressoTEEVerifier and IEspressoNitroTEEVerifier.
///         Supports only the Nitro TEE verifier.
/// @dev Note: BatchAuthenticator only uses registeredSigners() and doesn't check
///      enclave hashes, so registeredEnclaveHash() always returns false in this mock.
contract MockEspressoTEEVerifier is IEspressoTEEVerifier, IEspressoNitroTEEVerifier {
    mapping(address => bool) private _registeredSigners;
    // Note: Enclave hash checks are not used by BatchAuthenticator, so we don't
    // maintain a mapping for them in this mock. If needed for future tests,
    // we could add: mapping(bytes32 => bool) private _registeredEnclaveHashes;

    function espressoNitroTEEVerifier() external view override returns (IEspressoNitroTEEVerifier) {
        return this;
    }

    function espressoSGXTEEVerifier() external pure override returns (IEspressoSGXTEEVerifier) {
        return IEspressoSGXTEEVerifier(address(0));
    }

    function verify(bytes memory, bytes32, TeeType teeType) external pure override returns (bool) {
        if (teeType == TeeType.NITRO) {
            return true;
        }
        // SGX is not supported.
        return false;
    }

    function registerSigner(bytes calldata, bytes calldata, TeeType teeType) external pure override {
        require(teeType == TeeType.NITRO, "MockEspressoTEEVerifier: only NITRO supported");
    }

    function registeredSigners(address signer, TeeType teeType) external view override returns (bool) {
        if (teeType == TeeType.NITRO) {
            return _registeredSigners[signer];
        }
        // SGX is not supported.
        return false;
    }

    /// @notice Always returns false - BatchAuthenticator doesn't use enclave hash checks.
    function registeredEnclaveHashes(bytes32, TeeType) external pure override returns (bool) {
        // BatchAuthenticator only checks registeredSigners, not enclave hashes
        return false;
    }

    function setEspressoSGXTEEVerifier(IEspressoSGXTEEVerifier) external pure override {
        // No-op: SGX is not supported.
    }

    function setEspressoNitroTEEVerifier(IEspressoNitroTEEVerifier) external pure override {
        // No-op: this contract can only be used as the Nitro TEE verifier.
    }

    function registeredSigners(address signer) external view override returns (bool) {
        return _registeredSigners[signer];
    }

    /// @notice Always returns false - BatchAuthenticator doesn't use enclave hash checks.
    function registeredEnclaveHash(bytes32) external pure override returns (bool) {
        // BatchAuthenticator only checks registeredSigners, not enclave hashes
        return false;
    }

    function registerSigner(bytes calldata, bytes calldata) external pure override {
        // No-op for testing.
    }

    function setEnclaveHash(bytes32, bool) external pure override {
        // No-op for testing.
    }

    function deleteRegisteredSigners(address[] memory) external pure override {
        // No-op for testing.
    }

    /// @notice Test helper to set registered signers in the mock.
    function setRegisteredSigner(address signer, bool value) external {
        _registeredSigners[signer] = value;
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
    function test_initialize_revertsWhenAlreadyInitialized() external {
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
    function test_initialize_revertsWhenTeeBatcherIsZero() external {
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
    function test_initialize_revertsWhenNonTeeBatcherIsZero() external {
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
    function test_initialize_revertsWhenVerifierIsZero() external {
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
    function test_initialize_succeedsWithValidAddresses() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        assertEq(address(authenticator.espressoTEEVerifier()), address(teeVerifier));
        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
        assertTrue(authenticator.activeIsTee());
    }

    /// @notice Test that switchBatcher can only be called by ProxyAdmin or owner.
    function test_switchBatcher_onlyProxyAdminOrOwner() external {
        BatchAuthenticator authenticator = _deployAndInitializeProxy();

        // ProxyAdmin owner can switch.
        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();
        assertFalse(authenticator.activeIsTee());

        // Switch back.
        vm.prank(proxyAdminOwner);
        authenticator.switchBatcher();
        assertTrue(authenticator.activeIsTee());

        // ProxyAdmin can switch.
        vm.prank(address(proxyAdmin));
        authenticator.switchBatcher();
        assertFalse(authenticator.activeIsTee());

        // Switch back.
        vm.prank(address(proxyAdmin));
        authenticator.switchBatcher();
        assertTrue(authenticator.activeIsTee());

        // Unauthorized cannot switch.
        vm.prank(unauthorized);
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
        bytes memory signature = "test signature";

        vm.expectEmit(true, false, false, true);
        emit SignerRegistrationInitiated(address(this));

        authenticator.registerSigner(attestationTbs, signature);
    }

    /// @notice Test upgrade to new implementation.
    function test_upgrade_succeeds() external {
        // Create and initialize a proxy.
        BatchAuthenticator authenticator = _deployAndInitializeProxy();
        Proxy proxy = Proxy(payable(address(authenticator)));

        // Deploy new implementation.
        BatchAuthenticator newImpl = new BatchAuthenticator();

        // Upgrade.
        vm.prank(proxyAdminOwner);
        proxyAdmin.upgrade(payable(address(proxy)), address(newImpl));

        // Verify implementation changed.
        address newImplementation = EIP1967Helper.getImplementation(address(proxy));
        assertEq(newImplementation, address(newImpl));

        // Verify state is preserved
        assertEq(address(authenticator.espressoTEEVerifier()), address(teeVerifier));
        assertEq(authenticator.teeBatcher(), teeBatcher);
        assertEq(authenticator.nonTeeBatcher(), nonTeeBatcher);
    }

    // Event declarations for expectEmit.
    event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer);
    event SignerRegistrationInitiated(address indexed caller);
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
        // Create a fork of Sepolia using public Infura endpoint.
        string memory forkUrl = "https://sepolia.infura.io/v3/b9794ad1ddf84dfb8c34d6bb5dca2001";
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
        assertEq(authenticator.version(), "2.0.0");

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
        address signer = address(0x9999);
        bytes32 commitment = keccak256("test commitment on sepolia");

        // Register the signer.
        teeVerifier.setRegisteredSigner(signer, true);

        // Create a signature.
        uint256 privateKey = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, commitment);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Authenticate.
        vm.expectEmit(true, true, false, false);
        emit BatchInfoAuthenticated(commitment, vm.addr(privateKey));
        authenticator.authenticateBatchInfo(commitment, signature);

        assertTrue(authenticator.validBatchInfo(commitment));
    }

    /// @notice Test upgrade on Sepolia fork.
    function testFork_upgrade_preservesState() external {
        // Initialize the authenticator.
        bytes32 commitment = keccak256("test commitment");
        address signer = address(0x9999);
        teeVerifier.setRegisteredSigner(signer, true);
        uint256 privateKey = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
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
        assertEq(authenticator.version(), "2.0.0");
        assertTrue(authenticator.activeIsTee());

        // Verify the fork is working by testing that we can read the block number.
        uint256 blockNum = block.number;
        assertGt(blockNum, 0);
        console.log("Sepolia block number:", blockNum);
    }

    // Event declarations for expectEmit.
    event BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer);
    event SignerRegistrationInitiated(address indexed caller);
}
