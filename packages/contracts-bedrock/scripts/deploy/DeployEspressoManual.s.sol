// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { Script } from "forge-std/Script.sol";
import { console2 as console } from "forge-std/console2.sol";
import { DeployEspresso, DeployEspressoInput, DeployEspressoOutput } from "scripts/deploy/DeployEspresso.s.sol";
import {
    DeployAWSNitroVerifier,
    DeployAWSNitroVerifierInput,
    DeployAWSNitroVerifierOutput
} from "scripts/deploy/DeployAWSNitroVerifier.s.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { IBatchInbox } from "interfaces/L1/IBatchInbox.sol";
import { EspressoNitroTEEVerifier } from "@espresso-tee-contracts/EspressoNitroTEEVerifier.sol";
import { EspressoTEEVerifier } from "@espresso-tee-contracts/EspressoTEEVerifier.sol";
import { INitroEnclaveVerifier } from "aws-nitro-enclave-attestation/interfaces/INitroEnclaveVerifier.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { IProxyAdmin } from "interfaces/universal/IProxyAdmin.sol";
import { IProxy } from "interfaces/universal/IProxy.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { BatchInbox } from "src/L1/BatchInbox.sol";

/// @title DeployEspressoManual
/// @notice Manual deployment script for Espresso contracts.
/// @dev Deploys the complete Espresso stack: TEE Verifiers + Batch Authenticator + Batch Inbox.
///
///      REQUIRED Environment Variables:
///      - BATCH_DEPLOYER_KEY: Private key for deploying Batch contracts.
///      - TEE_DEPLOYER_KEY: Private key for deploying TEE contracts.
///      - NON_TEE_BATCHER: Address of the non-TEE batcher.
///      - TEE_BATCHER: Address of the TEE batcher.
///      - BATCH_AUTHENTICATOR_OWNER: Address of the owner for the BatchAuthenticator/BatchInbox.
///      - UNDERLYING_NITRO_VERIFIER: Address of the existing AWS Nitro Enclave Verifier.
///      - ENCLAVE_HASH: Enclave hash for Espresso Nitro TEE Verifier (bytes32, cannot be zero).
///
///      OPTIONAL Environment Variables:
///      - SALT: Salt for CREATE2 deployments. Defaults to hash of timestamp + deployer.
///
///      Deployment Order:
///      1. EspressoNitroTEEVerifier (wraps UNDERLYING_NITRO_VERIFIER)
///      2. EspressoTEEVerifier (routes to EspressoNitroTEEVerifier)
///      3. ProxyAdmin (for BatchAuthenticator proxy)
///      4. Proxy + BatchAuthenticator implementation
///      5. BatchInbox
///
///      Usage:
///      FOUNDRY_PROFILE=espresso forge script scripts/deploy/DeployEspressoManual.s.sol:DeployEspressoManual \
///        --rpc-url $ETH_RPC_URL --broadcast --verify --etherscan-api-key $ETHERSCAN_API_KEY
contract DeployEspressoManual is DeployEspresso {
    function run() public {
        console.log("=== Espresso Manual Deployment (Multi-Key) ===");

        // 1. Setup Keys
        (uint256 batchKey, uint256 teeKey, address batchAddr) = setupKeys();
        bool hasBatchKey = batchKey != 0;
        bool hasTeeKey = teeKey != 0;

        // 2. Deploy TEE Stack
        address nitroEnclaveVerifier = getNitroEnclaveVerifier();
        address espressoNitroVerifier = deployEspressoNitroVerifier(nitroEnclaveVerifier, teeKey, hasTeeKey);
        address espressoTEEVerifier = deployEspressoTEEVerifier(espressoNitroVerifier, teeKey, hasTeeKey);

        // 3. Deploy Batch Stack
        DeployEspressoOutput output = deployBatchStack(espressoTEEVerifier, batchAddr, batchKey, hasBatchKey);

        console.log("\n=== Final Deployment Summary ===");
        console.log("Underlying Nitro Verifier: %s", nitroEnclaveVerifier);
        console.log("Espresso Nitro Verifier:   %s", espressoNitroVerifier);
        console.log("Espresso TEE Verifier:     %s", espressoTEEVerifier);
        console.log("Batch Authenticator:       %s", output.batchAuthenticatorAddress());
        console.log("Batch Inbox:               %s", output.batchInboxAddress());
        console.log("==================================");
    }

    function setupKeys() internal view returns (uint256 batchKey, uint256 teeKey, address batchAddr) {
        bool hasBatch = false;
        try vm.envUint("BATCH_DEPLOYER_KEY") returns (uint256 k) {
            batchKey = k;
            hasBatch = true;
            console.log("Loaded BATCH_DEPLOYER_KEY");
        } catch {
            console.log("Using Default Key for Batch Contracts");
        }

        try vm.envUint("TEE_DEPLOYER_KEY") returns (uint256 k) {
            teeKey = k;
            console.log("Loaded TEE_DEPLOYER_KEY");
        } catch {
            if (hasBatch) {
                teeKey = batchKey;
                console.log("Using BATCH_DEPLOYER_KEY for TEE Contracts");
            } else {
                console.log("Using Default Key for TEE Contracts");
            }
        }

        batchAddr = hasBatch ? vm.addr(batchKey) : msg.sender;
    }

    function getNitroEnclaveVerifier() internal view returns (address nitroEnclaveVerifier) {
        try vm.envAddress("UNDERLYING_NITRO_VERIFIER") returns (address u) {
            nitroEnclaveVerifier = u;
            console.log("[1/5] Underlying Verifier: Existing at %s", nitroEnclaveVerifier);
        } catch {
            console.log("Error: UNDERLYING_NITRO_VERIFIER env var not set.");
            console.log("Please set this to the deployed NitroEnclaveVerifier address.");
            console.log("Refer to https://github.com/automata-network/aws-nitro-enclave-attestation for deployments.");
            revert("Missing Required Env Var: UNDERLYING_NITRO_VERIFIER");
        }
    }

    function deployEspressoNitroVerifier(address nitroEnclaveVerifier, uint256 key, bool hasKey) internal returns (address espressoNitroTEEVerifier) {
        console.log("[2/5] Espresso Nitro Verifier: Deploying...");
        bytes32 enclaveHash;
        try vm.envBytes32("ENCLAVE_HASH") returns (bytes32 h) {
            enclaveHash = h;
            console.log("      -> Enc Hash: %s", vm.toString(enclaveHash));
        } catch {
            console.log("Error: ENCLAVE_HASH env var not set.");
            console.log("This is required by the EspressoNitroTEEVerifier constructor.");
            revert("Missing Required Env Var: ENCLAVE_HASH");
        }

        if (hasKey) vm.startBroadcast(key);
        else vm.startBroadcast();

        EspressoNitroTEEVerifier instance = new EspressoNitroTEEVerifier(
            enclaveHash,
            INitroEnclaveVerifier(nitroEnclaveVerifier)
        );

        vm.stopBroadcast();
        espressoNitroTEEVerifier = address(instance);
        console.log("      -> Deployed to: %s", espressoNitroTEEVerifier);
    }

    function deployEspressoTEEVerifier(
        address espressoNitroTEEVerifier,
        uint256 key,
        bool hasKey
    )
        internal
        returns (address espressoTEEVerifier)
    {
        console.log("[3/5] Espresso TEE Verifier: Deploying...");

        if (hasKey) vm.startBroadcast(key);
        else vm.startBroadcast();

        // Nitro only for now
        IEspressoTEEVerifier espressoTEEInstance = IEspressoTEEVerifier(address(new EspressoTEEVerifier(
            IEspressoSGXTEEVerifier(address(0)),
            IEspressoNitroTEEVerifier(espressoNitroTEEVerifier)
        )));

        vm.stopBroadcast();
        espressoTEEVerifier = address(espressoTEEInstance);
        console.log("      -> Deployed to: %s", espressoTEEVerifier);
    }

    function deployBatchStack(
        address espressoTEEVerifier,
        address deployerAddr,
        uint256 key,
        bool hasKey
    )
        internal
        returns (DeployEspressoOutput output)
    {
        address nonTeeBatcher = vm.envAddress("NON_TEE_BATCHER");
        address teeBatcher = vm.envAddress("TEE_BATCHER");
        address batchAuthenticatorOwner = vm.envAddress("BATCH_AUTHENTICATOR_OWNER");

        bytes32 salt;
        try vm.envBytes32("SALT") returns (bytes32 s) {
            salt = s;
        } catch {
            salt = keccak256(abi.encodePacked(block.timestamp, deployerAddr));
        }

        output = new DeployEspressoOutput();

        console.log("[4/5] Batch Authenticator: Deploying...");
        if (hasKey) vm.startBroadcast(key);
        else vm.startBroadcast();

        // 1. Deploy ProxyAdmin
        ProxyAdmin proxyAdmin = ProxyAdmin(
            payable(
                DeployUtils.create1({
                    _name: "ProxyAdmin",
                    _args: DeployUtils.encodeConstructor(abi.encodeCall(IProxyAdmin.__constructor__, (deployerAddr)))
                })
            )
        );
        vm.label(address(proxyAdmin), "BatchAuthenticatorProxyAdmin");

        // 2. Deploy Proxy
        Proxy proxy = Proxy(
            payable(
                DeployUtils.create1({
                    _name: "Proxy",
                    _args: DeployUtils.encodeConstructor(abi.encodeCall(IProxy.__constructor__, (address(proxyAdmin))))
                })
            )
        );
        vm.label(address(proxy), "BatchAuthenticatorProxy");

        // 3. Set Proxy Type
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        // 4. Deploy Implementation
        BatchAuthenticator impl = new BatchAuthenticator();
        vm.label(address(impl), "BatchAuthenticatorImpl");

        // 5. Initialize Proxy
        bytes memory initData =
            abi.encodeCall(BatchAuthenticator.initialize, (IEspressoTEEVerifier(espressoTEEVerifier), teeBatcher, nonTeeBatcher));
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(impl), initData);

        // 6. Transfer Ownership (if needed)
        // deployerAddr owns it now. If we want someone else, transfer.
        if (batchAuthenticatorOwner != deployerAddr) {
            proxyAdmin.transferOwnership(batchAuthenticatorOwner);
        }

        IBatchAuthenticator ba = IBatchAuthenticator(address(proxy));
        output.set(output.batchAuthenticatorAddress.selector, address(ba));
        console.log("      -> Deployed to: %s", address(ba));

        console.log("[5/5] Batch Inbox: Deploying...");
        // (Broadcast continues from above)

        // 7. Deploy BatchInbox
        BatchInbox inboxInstance = new BatchInbox(ba, batchAuthenticatorOwner);
        IBatchInbox inboxImpl = IBatchInbox(address(inboxInstance));

        vm.label(address(inboxImpl), "BatchInboxImpl");
        output.set(output.batchInboxAddress.selector, address(inboxImpl));

        vm.stopBroadcast();
        console.log("      -> Deployed to: %s", address(inboxImpl));
    }
}
