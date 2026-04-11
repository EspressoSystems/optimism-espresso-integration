// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {IBatchAuthenticator} from "interfaces/L1/IBatchAuthenticator.sol";
import {ISystemConfig} from "interfaces/L1/ISystemConfig.sol";
import {IEspressoTEEVerifier} from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import {IProxyAdmin} from "interfaces/universal/IProxyAdmin.sol";
import {DeployUtils} from "scripts/libraries/DeployUtils.sol";
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
        address _espressoBatcher,
        address _systemConfig,
        address _teeVerifier,
        address _proxyAdminOwner
    ) public {
        require(_espressoBatcher != address(0), "DeployBatchAuthenticator: espressoBatcher required");
        require(_systemConfig != address(0), "DeployBatchAuthenticator: systemConfig required");
        require(_teeVerifier != address(0), "DeployBatchAuthenticator: teeVerifier required");

        if (_proxyAdminOwner == address(0)) {
            _proxyAdminOwner = msg.sender;
            console.log("WARN: proxyAdminOwner not set, defaulting to msg.sender");
        }

        vm.startBroadcast(msg.sender);

        IProxyAdmin proxyAdmin = IProxyAdmin(DeployUtils.create1({ _name: "ProxyAdmin", _args: abi.encode(msg.sender) }));
        vm.label(address(proxyAdmin), "BatchAuthenticatorProxyAdmin");
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.label(address(proxy), "BatchAuthenticatorProxy");
        proxyAdmin.setProxyType(address(proxy), IProxyAdmin.ProxyType.ERC1967);
        BatchAuthenticator impl = new BatchAuthenticator();
        vm.label(address(impl), "BatchAuthenticatorImpl");

        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize,
            (
                IEspressoTEEVerifier(_teeVerifier),
                _espressoBatcher,
                ISystemConfig(_systemConfig),
                _proxyAdminOwner
            )
        );
        proxyAdmin.upgradeAndCall(
            payable(address(proxy)),
            address(impl),
            initData
        );

        if (_proxyAdminOwner != msg.sender) {
            proxyAdmin.transferOwnership(_proxyAdminOwner);
        }

        vm.stopBroadcast();

        console.log("BatchAuthenticator (proxy):", address(proxy));
        console.log("BatchAuthenticator (impl): ", address(impl));
        console.log("ProxyAdmin:                ", address(proxyAdmin));
    }
}
