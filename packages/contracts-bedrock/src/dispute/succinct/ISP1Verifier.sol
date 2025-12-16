// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title SP1 Verifier Interface
/// @author Succinct Labs
/// @notice This contract is the interface for the SP1 Verifier.
interface ISP1Verifier {
    /// @notice Verifies a proof with given public values and vkey.
    /// @dev It is expected that the first 4 bytes of proofBytes must match the first 4 bytes of
    /// target verifier's VERIFIER_HASH.
    /// @param programVKey The verification key for the RISC-V program.
    /// @param publicValues The public values encoded as bytes.
    /// @param proofBytes The proof of the program execution the SP1 zkVM encoded as bytes.
    function verifyProof(bytes32 programVKey, bytes calldata publicValues, bytes calldata proofBytes) external view;
}

interface ISP1VerifierWithHash is ISP1Verifier {
    /// @notice Returns the hash of the verifier.
    function VERIFIER_HASH() external pure returns (bytes32);
}

/// @title SP1 Mock Verifier
/// @author Succinct Labs
/// @notice A mock verifier for local testing that accepts any proof.
contract SP1MockVerifier is ISP1Verifier {
    /// @notice Verifies a mock proof with given public values and vkey.
    /// @dev For testing, accepts empty proofs.
    function verifyProof(bytes32, bytes calldata, bytes calldata proofBytes) external pure {
        assert(proofBytes.length == 0);
    }
}
