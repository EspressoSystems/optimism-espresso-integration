// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";

contract BatchInbox {
    string public constant version = "1.1.0";

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

    function postCalldata(bytes calldata data) external {
        _requireAuthorized(keccak256(data));
    }

    function postBlobs() external {
        _requireAuthorized(_commitmentFromBlobs());
    }

    function _activeBatcher() internal view returns (address active, bool isTee) {
        if (activeIsTee) {
            return (teeBatcher, true);
        } else {
            return (nonTeeBatcher, false);
        }
    }

    function _requireAuthorized(bytes32 commitment) internal view {
        (address active, bool isTee) = _activeBatcher();
        require(msg.sender == active, "BatchInbox: inactive batcher");
        if (isTee) {
            require(batchAuthenticator.validBatchInfo(commitment), "BatchInbox: invalid batch");
        }
    }

    function _commitmentFromBlobs() internal view returns (bytes32) {
        bytes memory concatenatedHashes;
        uint256 i;
        while (blobhash(i) != 0) {
            concatenatedHashes = bytes.concat(concatenatedHashes, blobhash(i));
            unchecked {
                i++;
            }
        }
        return keccak256(concatenatedHashes);
    }
}
