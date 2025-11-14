// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

contract BatchInbox {
    address public immutable teeBatcher;
    address public immutable nonTeeBatcher;
    IBatchAuthenticator public immutable batchAuthenticator;

    // true if teeBatcher is active, false if nonTeeBatcher is active
    bool public activeIsTee;

    constructor(address _teeBatcher, address _nonTeeBatcher, IBatchAuthenticator _batchAuthenticator) {
        require(_teeBatcher != address(0) && _nonTeeBatcher != address(0), "BatchInbox: zero batcher");
        teeBatcher = _teeBatcher;
        nonTeeBatcher = _nonTeeBatcher;
        batchAuthenticator = _batchAuthenticator;
        // By default, start with the TEE batcher active
        activeIsTee = true;
    }

    function switchBatcher() external {
        activeIsTee = !activeIsTee;
    }

    fallback() external {
        address expectedBatcher = activeIsTee ? teeBatcher : nonTeeBatcher;
        if (msg.sender != expectedBatcher) {
            revert("BatchInbox: unauthorized batcher");
        }

        // Only TEE batchers require authentication
        if (activeIsTee) {
            if (blobhash(0) != 0) {
                bytes memory concatenatedHashes = new bytes(0);
                uint256 currentBlob = 0;
                while (blobhash(currentBlob) != 0) {
                    concatenatedHashes = bytes.concat(concatenatedHashes, blobhash(currentBlob));
                    currentBlob++;
                }
                bytes32 hash = keccak256(concatenatedHashes);
                if (!batchAuthenticator.validBatchInfo(hash)) {
                    revert("Invalid blob batch");
                }
            } else {
                bytes32 hash = keccak256(msg.data);
                if (!batchAuthenticator.validBatchInfo(hash)) {
                    revert("Invalid calldata batch");
                }
            }
        }
    }

    function _requireAuthorized(bytes32 commitment) internal view {
        (address active, bool isTee) = _activeBatcher();
        require(msg.sender == active, "BatchInbox: inactive batcher");
        if (isTee) {
            require(batchAuthenticator.validBatchInfo(commitment), "BatchInbox: invalid batch");
        }
    }

    function _activeBatcher() internal view returns (address active, bool isTee) {
        if (activeIsTee) {
            return (teeBatcher, true);
        } else {
            return (nonTeeBatcher, false);
        }
    }
}
