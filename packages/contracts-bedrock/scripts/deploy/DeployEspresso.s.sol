// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";
import { IBatchInbox } from "interfaces/L1/IBatchInbox.sol";
import { Script } from "forge-std/Script.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { EspressoTEEVerifier } from "@espresso-tee-contracts/EspressoTEEVerifier.sol";
import { IProxyAdmin } from "interfaces/universal/IProxyAdmin.sol";
import { IProxy } from "interfaces/universal/IProxy.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { console2 as console } from "forge-std/console2.sol";

contract DeployEspressoInput is BaseDeployIO {
    bytes32 internal _salt;
    address internal _nitroTEEVerifier;
    address internal _nonTeeBatcher;
    address internal _teeBatcher;

    function set(bytes4 _sel, bytes32 _val) public {
        if (_sel == this.salt.selector) _salt = _val;
        else revert("DeployEspressoInput: unknown selector");
    }

    function set(bytes4 _sel, address _val) public {
        if (_sel == this.nitroTEEVerifier.selector) {
            _nitroTEEVerifier = _val;
        } else if (_sel == this.nonTeeBatcher.selector) {
            _nonTeeBatcher = _val;
        } else if (_sel == this.teeBatcher.selector) {
            _teeBatcher = _val;
        } else {
            revert("DeployEspressoInput: unknown selector");
        }
    }

    function salt() public view returns (bytes32) {
        require(_salt != 0, "DeployEspressoInput: salt not set");
        return _salt;
    }

    function nitroTEEVerifier() public view returns (address) {
        return _nitroTEEVerifier;
    }

    function nonTeeBatcher() public view returns (address) {
        return _nonTeeBatcher;
    }

    function teeBatcher() public view returns (address) {
        return _teeBatcher;
    }
}

contract DeployEspressoOutput is BaseDeployIO {
    address internal _batchInboxAddress;
    address internal _batchAuthenticatorAddress;

    function set(bytes4 _sel, address _addr) public {
        require(_addr != address(0), "DeployEspressoOutput: cannot set zero address");
        if (_sel == this.batchInboxAddress.selector) {
            _batchInboxAddress = _addr;
        } else if (_sel == this.batchAuthenticatorAddress.selector) {
            _batchAuthenticatorAddress = _addr;
        } else {
            revert("DeployEspressoOutput: unknown selector");
        }
    }

    function batchAuthenticatorAddress() public view returns (address) {
        require(_batchAuthenticatorAddress != address(0), "DeployEspressoOutput: batch authenticator address not set");
        return _batchAuthenticatorAddress;
    }

    function batchInboxAddress() public view returns (address) {
        require(_batchInboxAddress != address(0), "DeployEspressoOutput: batcher inbox address not set");
        return _batchInboxAddress;
    }
}

contract DeployEspresso is Script {
    function run(DeployEspressoInput input, DeployEspressoOutput output, address deployerAddress) public {
        IEspressoTEEVerifier teeVerifier = deployTEEVerifier(input, deployerAddress);
        IBatchAuthenticator batchAuthenticator = deployBatchAuthenticator(input, output, teeVerifier, deployerAddress);
        deployBatchInbox(input, output, batchAuthenticator, deployerAddress);
        checkOutput(output);
    }

    function deployBatchAuthenticator(
        DeployEspressoInput input,
        DeployEspressoOutput output,
        IEspressoTEEVerifier teeVerifier,
        address deployerAddress
    )
        public
        returns (IBatchAuthenticator)
    {
        // Deploy the proxy admin, the proxy, and the batch authenticator implementation.
        // We create ProxyAdmin with msg.sender as the owner to ensure broadcasts come from
        // the expected address, then transfer ownership to deployerAddress afterward.
        // Use DeployUtils.create1 to ensure artifacts are available for vm.getCode calls.
        vm.broadcast(msg.sender);
        ProxyAdmin proxyAdmin = ProxyAdmin(
            payable(
                DeployUtils.create1({
                    _name: "ProxyAdmin",
                    _args: DeployUtils.encodeConstructor(abi.encodeCall(IProxyAdmin.__constructor__, (msg.sender)))
                })
            )
        );
        vm.label(address(proxyAdmin), "BatchAuthenticatorProxyAdmin");
        vm.broadcast(msg.sender);
        Proxy proxy = Proxy(
            payable(
                DeployUtils.create1({
                    _name: "Proxy",
                    _args: DeployUtils.encodeConstructor(abi.encodeCall(IProxy.__constructor__, (address(proxyAdmin))))
                })
            )
        );
        vm.label(address(proxy), "BatchAuthenticatorProxy");
        vm.broadcast(msg.sender);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);
        vm.broadcast(msg.sender);
        BatchAuthenticator impl = new BatchAuthenticator();
        vm.label(address(impl), "BatchAuthenticatorImpl");

        // Initialize the proxy.
        bytes memory initData =
            abi.encodeCall(BatchAuthenticator.initialize, (teeVerifier, input.teeBatcher(), input.nonTeeBatcher()));
        vm.broadcast(msg.sender);
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(impl), initData);

        // Transfer ownership to the desired deployerAddress if it differs from msg.sender.
        if (deployerAddress != msg.sender) {
            vm.broadcast(msg.sender);
            proxyAdmin.transferOwnership(deployerAddress);
        }

        // Return the proxied contract.
        IBatchAuthenticator batchAuthenticator = IBatchAuthenticator(address(proxy));
        output.set(output.batchAuthenticatorAddress.selector, address(batchAuthenticator));
        return batchAuthenticator;
    }

    function deployTEEVerifier(
        DeployEspressoInput input,
        address /* deployerAddress */
    )
        public
        returns (IEspressoTEEVerifier)
    {
        IEspressoNitroTEEVerifier nitroTEEVerifier = IEspressoNitroTEEVerifier(input.nitroTEEVerifier());
        vm.broadcast(msg.sender);
        IEspressoTEEVerifier impl = new EspressoTEEVerifier(
            // SGX TEE verifier is not yet implemented
            IEspressoSGXTEEVerifier(address(0)),
            nitroTEEVerifier
        );
        vm.label(address(impl), "EspressoTEEVerifierImpl");
        return impl;
    }

    function deployBatchInbox(
        DeployEspressoInput input,
        DeployEspressoOutput output,
        IBatchAuthenticator batchAuthenticator,
        address deployerAddress
    )
        public
    {
        bytes32 salt = input.salt();
        vm.broadcast(msg.sender);
        IBatchInbox impl = IBatchInbox(
            DeployUtils.create2({
                _name: "BatchInbox",
                _salt: salt,
                _args: DeployUtils.encodeConstructor(
                    abi.encodeCall(IBatchInbox.__constructor__, (address(batchAuthenticator), deployerAddress))
                )
            })
        );
        vm.label(address(impl), "BatchInboxImpl");
        output.set(output.batchInboxAddress.selector, address(impl));
    }

    function checkOutput(DeployEspressoOutput output) public view {
        address[] memory addresses =
            Solarray.addresses(address(output.batchAuthenticatorAddress()), address(output.batchInboxAddress()));
        DeployUtils.assertValidContractAddresses(addresses);
    }
}
