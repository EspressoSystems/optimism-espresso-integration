// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

contract BatchInbox {
    IBatchAuthenticator immutable batchAuthenticator;

    constructor(IBatchAuthenticator _batchAuthenticator) {
        batchAuthenticator = _batchAuthenticator;
    }

    fallback() external {
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
