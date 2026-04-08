// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {IBatchAuthenticator} from "interfaces/L1/IBatchAuthenticator.sol";
import {ISystemConfig} from "interfaces/L1/ISystemConfig.sol";
import {IEspressoTEEVerifier} from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import {ProxyAdmin} from "src/universal/ProxyAdmin.sol";
import {Proxy} from "src/universal/Proxy.sol";
import {BatchAuthenticator} from "src/L1/BatchAuthenticator.sol";

/// @notice Deploys only the BatchAuthenticator (proxy + impl) against an existing TEEVerifier.
///
/// Usage:
///   forge script scripts/deploy/DeployBatchAuthenticator.s.sol:DeployBatchAuthenticator \
///     --rpc-url <L1_RPC_URL> \
///     --broadcast \
///     --private-key <DEPLOYER_KEY> \
///     --verify \
///     --etherscan-api-key <API_KEY> \
///     --sig "run(address,address,address,address)" \
///     <ESPRESSO_BATCHER_ADDRESS> \
///     <SYSTEM_CONFIG_ADDRESS> \
///     <TEE_VERIFIER_ADDRESS> \
///     <PROXY_ADMIN_OWNER>
contract DeployBatchAuthenticator is Script {
    function run(
        address espressoBatcher,
        address systemConfig,
        address teeVerifier,
        address proxyAdminOwner
    ) public {
        require(espressoBatcher != address(0), "espressoBatcher required");
        require(systemConfig != address(0), "systemConfig required");
        require(teeVerifier != address(0), "teeVerifier required");

        if (proxyAdminOwner == address(0)) {
            proxyAdminOwner = msg.sender;
            console.log("WARN: proxyAdminOwner not set, defaulting to msg.sender");
        }

        vm.startBroadcast(msg.sender);

        ProxyAdmin proxyAdmin = new ProxyAdmin(msg.sender);
        vm.label(address(proxyAdmin), "BatchAuthenticatorProxyAdmin");
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.label(address(proxy), "BatchAuthenticatorProxy");
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);
        BatchAuthenticator impl = new BatchAuthenticator();
        vm.label(address(impl), "BatchAuthenticatorImpl");

        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize,
            (
                IEspressoTEEVerifier(teeVerifier),
                espressoBatcher,
                ISystemConfig(systemConfig),
                proxyAdminOwner
            )
        );
        proxyAdmin.upgradeAndCall(
            payable(address(proxy)),
            address(impl),
            initData
        );

        if (proxyAdminOwner != msg.sender) {
            proxyAdmin.transferOwnership(proxyAdminOwner);
        }

        vm.stopBroadcast();

        console.log("BatchAuthenticator (proxy):", address(proxy));
        console.log("BatchAuthenticator (impl): ", address(impl));
        console.log("ProxyAdmin:                ", address(proxyAdmin));
    }
}
