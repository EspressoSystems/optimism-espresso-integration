// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// Testing
import { Test } from "forge-std/Test.sol";

// Contracts
import { BatchInbox } from "src/L1/BatchInbox.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

/// @title MockBatchAuthenticator
/// @notice Mock implementation for testing - only implements validBatchInfo
contract MockBatchAuthenticator {
    mapping(bytes32 => bool) private validHashes;

    function setValidBatchInfo(bytes32 hash, bool valid) external {
        validHashes[hash] = valid;
    }

    function validBatchInfo(bytes32 hash) external view returns (bool) {
        return validHashes[hash];
    }
}

/// @title BatchInbox_Test
/// @notice Base test contract with common setup
contract BatchInbox_Test is Test {
    BatchInbox public inbox;
    MockBatchAuthenticator public authenticator;

    address public teeBatcher = address(0x1234);
    address public nonTeeBatcher = address(0x5678);
    address public deployer = address(0xABCD);
    address public unauthorized = address(0xDEAD);

    function setUp() public virtual {
        authenticator = new MockBatchAuthenticator();
        inbox = new BatchInbox(teeBatcher, nonTeeBatcher, IBatchAuthenticator(address(authenticator)));
    }
}

/// @title BatchInbox_Constructor_Test
/// @notice Tests for the BatchInbox constructor
contract BatchInbox_Constructor_Test is Test {
    address teeBatcher = address(0x1234);
    address nonTeeBatcher = address(0x5678);
    address batchAuthenticator = address(0x9ABC);

    /// @notice Test that constructor reverts when TEE batcher is zero address
    function test_constructor_revertsWhenTeeBatcherIsZero() external {
        vm.expectRevert("BatchInbox: zero batcher");
        new BatchInbox(address(0), nonTeeBatcher, IBatchAuthenticator(batchAuthenticator));
    }

    /// @notice Test that constructor reverts when non-TEE batcher is zero address
    function test_constructor_revertsWhenNonTeeBatcherIsZero() external {
        vm.expectRevert("BatchInbox: zero batcher");
        new BatchInbox(teeBatcher, address(0), IBatchAuthenticator(batchAuthenticator));
    }

    /// @notice Test that constructor reverts when both batchers are zero addresses
    function test_constructor_revertsWhenBothBatchersAreZero() external {
        vm.expectRevert("BatchInbox: zero batcher");
        new BatchInbox(address(0), address(0), IBatchAuthenticator(batchAuthenticator));
    }

    /// @notice Test that constructor succeeds with valid addresses
    function test_constructor_succeedsWithValidAddresses() external {
        BatchInbox testInbox = new BatchInbox(teeBatcher, nonTeeBatcher, IBatchAuthenticator(batchAuthenticator));

        assertEq(testInbox.teeBatcher(), teeBatcher, "TEE batcher should match");
        assertEq(testInbox.nonTeeBatcher(), nonTeeBatcher, "Non-TEE batcher should match");
        assertEq(address(testInbox.batchAuthenticator()), batchAuthenticator, "Batch authenticator should match");
        assertTrue(testInbox.activeIsTee(), "Active batcher should be TEE by default");
    }
}

/// @title BatchInbox_SwitchBatcher_Test
/// @notice Tests for switching between batchers
contract BatchInbox_SwitchBatcher_Test is BatchInbox_Test {
    /// @notice Test that switchBatcher toggles the active batcher
    function test_switchBatcher_togglesActiveBatcher() external {
        // Initially TEE batcher is active
        assertTrue(inbox.activeIsTee(), "Should start with TEE batcher active");

        // Switch to non-TEE batcher
        inbox.switchBatcher();
        assertFalse(inbox.activeIsTee(), "Should switch to non-TEE batcher");

        // Switch back to TEE batcher
        inbox.switchBatcher();
        assertTrue(inbox.activeIsTee(), "Should switch back to TEE batcher");
    }
}

/// @title BatchInbox_Fallback_Test
/// @notice Tests for the fallback function
contract BatchInbox_Fallback_Test is BatchInbox_Test {
    /// @notice Test that non-TEE batcher can post after switching
    function test_fallback_nonTeeBatcherCanPostAfterSwitch() external {
        // Switch to non-TEE batcher
        inbox.switchBatcher();

        // Non-TEE batcher should be able to post
        vm.prank(nonTeeBatcher);
        (bool success,) = address(inbox).call("hello");
        assertTrue(success, "Non-TEE batcher should be able to post");
    }

    /// @notice Test that inactive batcher reverts
    function test_fallback_inactiveBatcherReverts() external {
        // Switch to non-TEE batcher (making TEE batcher inactive)
        inbox.switchBatcher();

        // TEE batcher (now inactive) should revert
        vm.prank(teeBatcher);
        (bool success, bytes memory returnData) = address(inbox).call("unauthorized");
        assertFalse(success, "Should revert");
        // Check the revert reason
        assertEq(
            string(returnData), string(abi.encodeWithSignature("Error(string)", "BatchInbox: unauthorized batcher"))
        );
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
        assertEq(string(returnData), string(abi.encodeWithSignature("Error(string)", "Invalid calldata batch")));
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
        inbox.switchBatcher();

        bytes memory data = "no-auth-needed";
        // Don't set any authentication

        // Non-TEE batcher should succeed without authentication
        vm.prank(nonTeeBatcher);
        (bool success,) = address(inbox).call(data);
        assertTrue(success, "Non-TEE batcher should not require auth");
    }

    /// @notice Test that unauthorized address cannot post
    function test_fallback_unauthorizedAddressReverts() external {
        // Switch to non-TEE batcher. In this case the batch inbox should revert if the batcher is not authorized.
        inbox.switchBatcher();
        vm.prank(unauthorized);
        (bool success,) = address(inbox).call("unauthorized");
        assertFalse(success, "Unauthorized should revert when non-TEE is active");
    }
}
