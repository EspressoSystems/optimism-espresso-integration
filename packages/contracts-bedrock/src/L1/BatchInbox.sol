// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { IBatchVerifier } from "interfaces/L1/IBatchVerifier.sol";

contract BatchInbox {
    IBatchVerifier immutable batchVerifier;

    constructor(IBatchVerifier _batchVerifier) {
        batchVerifier = _batchVerifier;
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
            if (!batchVerifier.validBatches(hash)) {
                revert("Invalid blob batch");
            }
        } else {
            bytes32 hash = keccak256(msg.data);
            if (!batchVerifier.validBatches(hash)) {
                revert("Invalid calldata batch");
            }
        }
    }
}
