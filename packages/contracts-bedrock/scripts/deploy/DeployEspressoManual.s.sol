// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { console2 as console } from "forge-std/console2.sol";
import { DeployEspresso, DeployEspressoInput, DeployEspressoOutput } from "scripts/deploy/DeployEspresso.s.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { DeployAllTEEVerifiers } from "@espresso-tee-scripts/DeployAllTEEVerifiers.s.sol";

/// @title DeployEspressoManual
/// @notice Manual deployment script for Espresso contracts.
/// @dev Deploys the complete Espresso stack: TEE Verifiers + Batch Authenticator + Batch Inbox.
///      Uses submodule script `DeployAllTEEVerifiers` for TEE stack deployment.
///
///      REQUIRED Environment Variables:
///      - NON_TEE_BATCHER: Address of the non-TEE batcher.
///      - TEE_BATCHER: Address of the TEE batcher.
///      - BATCH_AUTHENTICATOR_OWNER: Address of the owner for the BatchAuthenticator/BatchInbox.
///      - SGX_QUOTE_VERIFIER_ADDRESS: Address for SGX verifier.
///      - NITRO_ENCLAVE_VERIFIER: Address of the existing AWS Nitro Enclave Verifier.
///      - ENCLAVE_HASH: Enclave hash for Espresso Nitro TEE Verifier (bytes32, cannot be zero).
///
///      OPTIONAL Environment Variables:
///      - BATCH_DEPLOYER_KEY: Private key for deploying Batch contracts (uses TEE_DEPLOYER_KEY or default signer if not
///        set).
///      - TEE_DEPLOYER_KEY: Private key for deploying TEE contracts (uses default signer if not set).
///                          Note: This is used by the submodule script via the --private-key flag,
///                          but you can also set it here to inform the deployment summary.
///      - SALT: Salt for CREATE2 deployments (defaults to hash of timestamp + deployer).
///
///      Usage (same key for both):
///      FOUNDRY_PROFILE=espresso forge script scripts/deploy/DeployEspressoManual.s.sol:DeployEspressoManual \
///        --rpc-url $ETH_RPC_URL --private-key $DEPLOYER_KEY --broadcast --verify
///
///      Usage (different keys):
///      # Set BATCH_DEPLOYER_KEY in .env, then run with TEE key:
///      FOUNDRY_PROFILE=espresso forge script scripts/deploy/DeployEspressoManual.s.sol:DeployEspressoManual \
///        --rpc-url $ETH_RPC_URL --private-key $TEE_DEPLOYER_KEY --broadcast --verify
contract DeployEspressoManual is DeployEspresso {
    function run() public {
        console.log("=== Espresso Manual Deployment ===");

        // 1. Load deployer keys
        (uint256 batchKey, address batchAddr, address teeAddr) = _loadDeployerKeys();

        // 2. Deploy TEE Stack via submodule script (uses the key passed to forge script)
        console.log("[1/3] Deploying TEE Verifiers...");
        console.log("      TEE Deployer: %s", teeAddr);
        vm.setEnv("PROXY_ADMIN_OWNER", vm.toString(teeAddr));
        DeployAllTEEVerifiers teeDeployer = new DeployAllTEEVerifiers();
        teeDeployer.run();
        address teeVerifier = _loadDeployedTEEVerifier();
        console.log("      TEE Verifier: %s", teeVerifier);

        // 3. Deploy Batch Stack (uses BATCH_DEPLOYER_KEY if set, otherwise default signer)
        console.log("[2/3] Deploying Batch Authenticator...");
        console.log("[3/3] Deploying Batch Inbox...");
        console.log("      Batch Deployer: %s", batchAddr);

        DeployEspressoInput input = _prepareBatchInput(batchAddr);
        DeployEspressoOutput output = new DeployEspressoOutput();

        if (batchKey != 0) {
            runDeployBatchStackWithKey(input, output, IEspressoTEEVerifier(teeVerifier), batchKey);
        } else {
            runDeployBatchStack(input, output, IEspressoTEEVerifier(teeVerifier));
        }

        // 4. Summary
        console.log("\n=== Deployment Summary ===");
        console.log("TEE Deployer:        %s", teeAddr);
        console.log("Batch Deployer:      %s", batchAddr);
        console.log("TEE Verifier:        %s", teeVerifier);
        console.log("Batch Authenticator: %s", output.batchAuthenticatorAddress());
        console.log("Batch Inbox:         %s", output.batchInboxAddress());
        console.log("==========================");
    }

    /// @notice Load deployer keys from environment.
    /// @return batchKey The private key for batch deployments (0 if using default signer).
    /// @return batchAddr The address that will deploy batch contracts.
    /// @return teeAddr The address that will deploy TEE contracts (always msg.sender, as it uses forge's
    /// --private-key).
    function _loadDeployerKeys() internal view returns (uint256 batchKey, address batchAddr, address teeAddr) {
        // TEE contracts use the key passed to forge script (msg.sender)
        teeAddr = msg.sender;

        // Batch contracts can use a separate key
        try vm.envUint("BATCH_DEPLOYER_KEY") returns (uint256 k) {
            batchKey = k;
            batchAddr = vm.addr(k);
            if (batchAddr != teeAddr) {
                console.log("Using separate keys: TEE=%s, Batch=%s", teeAddr, batchAddr);
            } else {
                console.log("Using same key for TEE and Batch: %s", teeAddr);
            }
        } catch {
            // No separate batch key, use the same as TEE
            batchAddr = teeAddr;
            console.log("Using single deployer for all contracts: %s", teeAddr);
        }
    }

    /// @notice Prepare the deployment input from environment variables.
    function _prepareBatchInput(address deployerAddr) internal returns (DeployEspressoInput input) {
        input = new DeployEspressoInput();
        input.set(input.nonTeeBatcher.selector, vm.envAddress("NON_TEE_BATCHER"));
        input.set(input.teeBatcher.selector, vm.envAddress("TEE_BATCHER"));
        input.set(input.proxyAdminOwner.selector, vm.envAddress("BATCH_AUTHENTICATOR_OWNER"));

        bytes32 salt;
        try vm.envBytes32("SALT") returns (bytes32 s) {
            salt = s;
        } catch {
            salt = keccak256(abi.encodePacked(block.timestamp, deployerAddr));
        }
        input.set(input.salt.selector, salt);
    }

    /// @notice Load the deployed TEE Verifier proxy address from the deployment artifact.
    function _loadDeployedTEEVerifier() internal view returns (address) {
        string memory path =
            string.concat(vm.projectRoot(), "/deployments/", vm.toString(block.chainid), "-espresso-tee-verifier.json");
        return vm.parseJsonAddress(vm.readFile(path), ".proxy");
    }
}
