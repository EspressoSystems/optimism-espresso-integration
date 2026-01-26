// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Testing
import { Test } from "forge-std/Test.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";

// Contracts
import { BatchInbox } from "src/L1/BatchInbox.sol";
import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { MockNitroTEEVerifier, MockEspressoTEEVerifier } from "./BatchAuthenticator.t.sol";

contract TestBatchAuthenticator is BatchAuthenticator {
    // Test helper to bypass signature verification in authenticateBatchInfo.
    function setValidBatchInfo(bytes32 hash, bool valid) external {
        validBatchInfo[hash] = valid;
    }
}

/// @title BatchInbox_Test
/// @notice Base test contract with common setup
contract BatchInbox_Test is Test {
    BatchInbox public inbox;
    TestBatchAuthenticator public authenticator;

    MockNitroTEEVerifier public nitroVerifier;
    MockEspressoTEEVerifier public teeVerifier;

    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);
    address public deployer = address(0xDEF0);
    address public unauthorized = address(0xDEAD);

    function setUp() public virtual {
        nitroVerifier = new MockNitroTEEVerifier();
        teeVerifier = new MockEspressoTEEVerifier(nitroVerifier);

        // Deploy implementation
        TestBatchAuthenticator impl = new TestBatchAuthenticator();

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
        authenticator = TestBatchAuthenticator(address(proxy));

        inbox = new BatchInbox(IBatchAuthenticator(address(authenticator)));
    }
}

/// @notice Minimal authenticator mock that returns a zero non-TEE batcher.
contract ConstructorMockBatchAuthenticatorZeroNonTee {
    function nonTeeBatcher() external pure returns (address) {
        return address(0);
    }
}

/// @notice Minimal authenticator mock that returns a configured non-TEE batcher.
contract ConstructorMockBatchAuthenticatorNonZero {
    address public nonTeeBatcherValue;

    constructor(address _nonTeeBatcherValue) {
        nonTeeBatcherValue = _nonTeeBatcherValue;
    }

    function nonTeeBatcher() external view returns (address) {
        return nonTeeBatcherValue;
    }
}

/// @title BatchInbox_Fallback_Test
/// @notice Tests for the fallback function
/// @dev Behavior matrix:
///      - When the TEE batcher is active (`activeIsTee == true`), any caller must provide
///        a previously authenticated batch commitment via `batchAuthenticator.validBatchInfo`.
///      - When the non-TEE batcher is active (`activeIsTee == false`), only `nonTeeBatcher`
///        may send batches and no additional authentication is required.
contract BatchInbox_Fallback_Test is BatchInbox_Test {
    /// @notice Test that non-TEE batcher can post after switching
    function test_fallback_nonTeeBatcherCanPostAfterSwitch() external {
        // Switch to non-TEE batcher
        vm.prank(deployer);
        authenticator.switchBatcher();

        // Non-TEE batcher should be able to post
        vm.prank(nonTeeBatcher);
        (bool success,) = address(inbox).call("hello");
        assertTrue(success, "Non-TEE batcher should be able to post");
    }

    /// @notice Test that inactive batcher reverts
    function test_fallback_inactiveBatcherReverts() external {
        // Switch to non-TEE batcher (making TEE batcher inactive)
        vm.prank(deployer);
        authenticator.switchBatcher();

        // TEE batcher (now inactive) should revert
        vm.prank(teeBatcher);
        (bool success, bytes memory returnData) = address(inbox).call("unauthorized");
        assertFalse(success, "Should revert");
        // Check the revert reason - contract returns detailed error with addresses
        string memory expectedError = string(
            abi.encodePacked(
                "BatchAuthenticator: batcher not authorized to post in fallback mode. Expected: ",
                Strings.toHexString(uint160(nonTeeBatcher), 20),
                ", Actual: ",
                Strings.toHexString(uint160(teeBatcher), 20)
            )
        );
        assertEq(string(returnData), string(abi.encodeWithSignature("Error(string)", expectedError)));
    }

    /// @notice Test that TEE batcher requires authentication
    function test_fallback_teeBatcherRequiresAuthentication() external {
        // TEE batcher is active by default
        bytes memory data = "needs-auth";
        bytes32 hash = keccak256(data);

        // Don't set the hash as valid in authenticator
        authenticator.setValidBatchInfo(hash, false);

        // TEE batcher should revert due to invalid authentication
        vm.prank(teeBatcher);
        (bool success, bytes memory returnData) = address(inbox).call(data);
        assertFalse(success, "Should revert");
        // Check the revert reason
        assertEq(
            string(returnData),
            string(abi.encodeWithSignature("Error(string)", "BatchAuthenticator: invalid calldata batch"))
        );
    }

    /// @notice Test that TEE batcher succeeds with valid authentication
    function test_fallback_teeBatcherSucceedsWithValidAuth() external {
        // TEE batcher is active by default
        bytes memory data = "valid-batch";
        bytes32 hash = keccak256(data);

        // Set the hash as valid in authenticator
        authenticator.setValidBatchInfo(hash, true);

        // TEE batcher should succeed
        vm.prank(teeBatcher);
        (bool success,) = address(inbox).call(data);
        assertTrue(success, "TEE batcher should succeed with valid auth");
    }

    /// @notice Test that non-TEE batcher doesn't require authentication
    function test_fallback_nonTeeBatcherDoesNotRequireAuth() external {
        // Switch to non-TEE batcher
        vm.prank(deployer);
        authenticator.switchBatcher();

        bytes memory data = "no-auth-needed";
        bytes32 hash = keccak256(data);
        authenticator.setValidBatchInfo(hash, false);

        // Non-TEE batcher should succeed without authentication
        vm.prank(nonTeeBatcher);
        (bool success,) = address(inbox).call(data);
        assertTrue(success, "Non-TEE batcher should not require auth");
    }

    /// @notice Test that unauthorized address cannot post
    function test_fallback_unauthorizedAddressReverts() external {
        // Switch to non-TEE batcher. In this case the batch inbox should revert if the batcher is not authorized.
        vm.prank(deployer);
        authenticator.switchBatcher();
        vm.prank(unauthorized);
        (bool success,) = address(inbox).call("unauthorized");
        assertFalse(success, "Unauthorized should revert when non-TEE is active");
    }

    /// @notice Test that non-TEE batcher is rejected while TEE batcher is active
    function test_fallback_nonTeeBatcherRevertsWhenTeeActiveAndUnauthenticated() external {
        // By default, the TEE batcher is active (activeIsTee == true).
        bytes memory data = "nontee-unauth";
        bytes32 hash = keccak256(data);

        // Even if the batch is authenticated, the non-TEE batcher should revert because it is not authorized to post
        // when TEE is active.
        authenticator.setValidBatchInfo(hash, true);

        vm.prank(nonTeeBatcher);
        (bool success, bytes memory returnData) = address(inbox).call(data);
        assertFalse(success, "Should revert when TEE is active and batch is not authenticated");
        // Check the revert reason - contract checks sender first, so it returns TEE mode error
        string memory expectedError = string(
            abi.encodePacked(
                "BatchAuthenticator: batcher not authorized to post in TEE mode. Expected: ",
                Strings.toHexString(uint160(teeBatcher), 20),
                ", Actual: ",
                Strings.toHexString(uint160(nonTeeBatcher), 20)
            )
        );
        assertEq(string(returnData), string(abi.encodeWithSignature("Error(string)", expectedError)));
    }
}

/// @title BatchInbox_Blob_Test
/// @notice Tests for blob batch submissions
contract BatchInbox_Blob_Test is BatchInbox_Test {
    /// @notice Test that TEE batcher succeeds with valid blob authentication
    function test_fallback_blobBatchSucceedsWithValidAuth() external {
        // TEE batcher is active by default
        bytes32[] memory blobHashes = new bytes32[](2);
        blobHashes[0] = keccak256("blob1");
        blobHashes[1] = keccak256("blob2");

        // Set blob hashes for this transaction
        vm.blobhashes(blobHashes);

        // Calculate expected hash (concatenation of blob hashes)
        bytes memory concatenatedHashes = bytes.concat(blobHashes[0], blobHashes[1]);
        bytes32 expectedHash = keccak256(concatenatedHashes);

        // Set the hash as valid in authenticator
        authenticator.setValidBatchInfo(expectedHash, true);

        // TEE batcher should succeed
        vm.prank(teeBatcher);
        (bool success,) = address(inbox).call("any-calldata");
        assertTrue(success, "TEE batcher should succeed with valid blob auth");
    }

    /// @notice Test that TEE batcher fails with invalid blob authentication
    function test_fallback_blobBatchFailsWithInvalidAuth() external {
        // TEE batcher is active by default
        bytes32[] memory blobHashes = new bytes32[](1);
        blobHashes[0] = keccak256("blob1");

        // Set blob hashes for this transaction
        vm.blobhashes(blobHashes);

        // Don't set the hash as valid - should fail

        // TEE batcher should fail
        vm.prank(teeBatcher);
        (bool success, bytes memory returnData) = address(inbox).call("any-calldata");
        assertFalse(success, "TEE batcher should fail with invalid blob auth");
        assertEq(
            string(returnData),
            string(abi.encodeWithSignature("Error(string)", "BatchAuthenticator: invalid blob batch"))
        );
    }

    /// @notice Test that blob hashes take precedence over calldata
    function test_fallback_blobHashesTakePrecedenceOverCalldata() external {
        // TEE batcher is active by default
        bytes memory data = "valid-calldata";
        bytes32 calldataHash = keccak256(data);

        bytes32[] memory blobHashes = new bytes32[](1);
        blobHashes[0] = keccak256("blob1");

        // Set blob hashes for this transaction
        vm.blobhashes(blobHashes);

        // Only authenticate the calldata hash, not the blob hash
        authenticator.setValidBatchInfo(calldataHash, true);

        // Should fail because blob hashes take precedence and blob hash is not authenticated
        vm.prank(teeBatcher);
        (bool success,) = address(inbox).call(data);
        assertFalse(success, "Should fail because blob hash is checked, not calldata");
    }

    /// @notice Test that non-TEE batcher ignores blobs (no auth required)
    function test_fallback_nonTeeBatcherIgnoresBlobs() external {
        // Switch to non-TEE batcher
        vm.prank(deployer);
        authenticator.switchBatcher();

        bytes32[] memory blobHashes = new bytes32[](1);
        blobHashes[0] = keccak256("blob1");

        // Set blob hashes for this transaction
        vm.blobhashes(blobHashes);

        // Don't authenticate the blob hash

        // Non-TEE batcher should succeed without authentication
        vm.prank(nonTeeBatcher);
        (bool success,) = address(inbox).call("any-calldata");
        assertTrue(success, "Non-TEE batcher should not require blob auth");
    }
}
