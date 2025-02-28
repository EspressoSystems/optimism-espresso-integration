// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";
import { IBatchVerifier } from "interfaces/L1/IBatchVerifier.sol";

contract BatchInbox {
    IBatchVerifier public immutable batchVerifier;

    constructor(IBatchVerifier _batchVerifier) {
        batchVerifier = _batchVerifier;
    }

    fallback(bytes calldata data) external returns (bytes memory) {
        if (!batchVerifier.validBatches(keccak256(data))) {
            revert("Invalid batch");
        }
        return "";
    }
}
