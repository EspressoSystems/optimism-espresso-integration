// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";

interface IBatchAuthenticator {
    /// @notice Error thrown when an invalid address (zero address) is provided.
    error InvalidAddress(address contract_);

    /// @notice Emitted when a batch info is authenticated.
    event BatchInfoAuthenticated(bytes32 indexed commitment);

    /// @notice Emitted when a signer registration is initiated through this contract.
    event SignerRegistrationInitiated(address indexed caller);

    /// @notice Emitted when the TEE batcher address is updated.
    event TeeBatcherUpdated(address indexed oldTeeBatcher, address indexed newTeeBatcher);

    /// @notice Emitted when the non-TEE batcher address is updated.
    event NonTeeBatcherUpdated(address indexed oldNonTeeBatcher, address indexed newNonTeeBatcher);

    /// @notice Emitted when the active batcher is switched.
    event BatcherSwitched(bool indexed activeIsTee);

    function authenticateBatchInfo(bytes32 commitment, bytes memory _signature) external;

    function espressoTEEVerifier() external view returns (IEspressoTEEVerifier);

    function nitroValidator() external view returns (address);

    function owner() external view returns (address);

    function teeBatcher() external view returns (address);

    function nonTeeBatcher() external view returns (address);

    function registerSigner(bytes memory attestationTbs, bytes memory signature) external;

    function validBatchInfo(bytes32) external view returns (bool);

    function activeIsTee() external view returns (bool);

    function switchBatcher() external;

    function setTeeBatcher(address _newTeeBatcher) external;

    function setNonTeeBatcher(address _newNonTeeBatcher) external;

    function validateBatch(address sender, bytes calldata data) external view;
}
