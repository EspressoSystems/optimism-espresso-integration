// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { EspressoNitroTEEVerifier } from "@espresso-tee-contracts/EspressoNitroTEEVerifier.sol";
import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";
import { Script } from "forge-std/Script.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { INitroEnclaveVerifier } from "aws-nitro-enclave-attestation/interfaces/INitroEnclaveVerifier.sol";
import { IProxyAdmin } from "interfaces/universal/IProxyAdmin.sol";
import { IProxy } from "interfaces/universal/IProxy.sol";
import { MockEspressoNitroTEEVerifier } from "test/mocks/MockEspressoTEEVerifiers.sol";

contract DeployAWSNitroVerifierInput is BaseDeployIO {
    address internal _nitroEnclaveVerifier;
    address internal _teeVerifierAddress;
    address internal _proxyAdminOwner;

    function set(bytes4 _sel, address _val) public {
        if (_sel == this.nitroEnclaveVerifier.selector) {
            _nitroEnclaveVerifier = _val;
        } else if (_sel == this.teeVerifierAddress.selector) {
            _teeVerifierAddress = _val;
        } else if (_sel == this.proxyAdminOwner.selector) {
            _proxyAdminOwner = _val;
        } else {
            revert("DeployAWSNitroVerifierInput: unknown selector");
        }
    }

    function nitroEnclaveVerifier() public view returns (address) {
        return _nitroEnclaveVerifier;
    }

    /// @notice The address of the main EspressoTEEVerifier contract that controls admin functions.
    ///         Can be address(0) during initial deployment if TEEVerifier doesn't exist yet.
    function teeVerifierAddress() public view returns (address) {
        return _teeVerifierAddress;
    }

    /// @notice The address that will own the ProxyAdmin. Defaults to msg.sender if not set.
    function proxyAdminOwner() public view returns (address) {
        return _proxyAdminOwner;
    }
}

contract DeployAWSNitroVerifierOutput is BaseDeployIO {
    address internal _nitroTEEVerifierProxy;
    address internal _nitroTEEVerifierImpl;
    address internal _proxyAdmin;

    function set(bytes4 _sel, address _addr) public {
        require(_addr != address(0), "DeployAWSNitroVerifierOutput: cannot set zero address");
        if (_sel == this.nitroTEEVerifierProxy.selector) {
            _nitroTEEVerifierProxy = _addr;
        } else if (_sel == this.nitroTEEVerifierImpl.selector) {
            _nitroTEEVerifierImpl = _addr;
        } else if (_sel == this.proxyAdmin.selector) {
            _proxyAdmin = _addr;
        } else {
            revert("DeployAWSNitroVerifierOutput: unknown selector");
        }
    }

    function nitroTEEVerifierProxy() public view returns (address) {
        require(_nitroTEEVerifierProxy != address(0), "nitro TEE verifier proxy not set");
        return _nitroTEEVerifierProxy;
    }

    function nitroTEEVerifierImpl() public view returns (address) {
        require(_nitroTEEVerifierImpl != address(0), "nitro TEE verifier impl not set");
        return _nitroTEEVerifierImpl;
    }

    function proxyAdmin() public view returns (address) {
        require(_proxyAdmin != address(0), "proxy admin not set");
        return _proxyAdmin;
    }

    /// @notice Alias for nitroTEEVerifierProxy for backward compatibility
    function nitroTEEVerifierAddress() public view returns (address) {
        return nitroTEEVerifierProxy();
    }
}

contract DeployAWSNitroVerifier is Script {
    struct ProxyDeployment {
        IProxyAdmin proxyAdmin;
        IProxy proxy;
    }

    function run(DeployAWSNitroVerifierInput input, DeployAWSNitroVerifierOutput output) public {
        deployNitroTEEVerifier(input, output);
        checkOutput(output);
    }

    /// @notice Deploys a contract by name using vm.getCode and CREATE. Avoids importing DeployUtils
    ///         (which transitively imports Blueprint/Bytes) to prevent compilation group merging with
    ///         the OZ v5 chain (from EspressoNitroTEEVerifier), which would create duplicate versioned
    ///         artifacts that break vm.getCode lookups.
    function _create1(string memory _name, bytes memory _args) internal returns (address payable addr_) {
        bytes memory bytecode = abi.encodePacked(vm.getCode(_name), _args);
        assembly {
            addr_ := create(0, add(bytecode, 0x20), mload(bytecode))
        }
        require(addr_ != address(0), "DeployAWSNitroVerifier: deployment failed");
    }

    /// @notice Deploys ProxyAdmin and Proxy contracts
    /// @param labelPrefix Prefix for vm.label (e.g., "Mock" or "")
    /// @return deployment Struct containing the deployed ProxyAdmin and Proxy
    function deployProxyInfrastructure(string memory labelPrefix)
        internal
        returns (ProxyDeployment memory deployment)
    {
        vm.broadcast(msg.sender);
        deployment.proxyAdmin = IProxyAdmin(_create1("ProxyAdmin", abi.encode(msg.sender)));
        vm.label(address(deployment.proxyAdmin), string.concat(labelPrefix, "NitroTEEVerifierProxyAdmin"));

        vm.broadcast(msg.sender);
        deployment.proxy = IProxy(payable(_create1("universal/Proxy.sol:Proxy", abi.encode(address(deployment.proxyAdmin)))));
        vm.label(address(deployment.proxy), string.concat(labelPrefix, "NitroTEEVerifierProxy"));

        vm.broadcast(msg.sender);
        deployment.proxyAdmin.setProxyType(address(deployment.proxy), IProxyAdmin.ProxyType.ERC1967);
    }

    function deployNitroTEEVerifier(
        DeployAWSNitroVerifierInput input,
        DeployAWSNitroVerifierOutput output
    )
        public
        returns (IEspressoNitroTEEVerifier)
    {
        address nitroEnclaveVerifier = input.nitroEnclaveVerifier();
        address proxyAdminOwner = input.proxyAdminOwner();
        if (proxyAdminOwner == address(0)) {
            proxyAdminOwner = msg.sender;
        }

        // Deploy proxy infrastructure
        ProxyDeployment memory deployment = deployProxyInfrastructure(nitroEnclaveVerifier == address(0) ? "Mock" : "");

        address implAddress;

        if (nitroEnclaveVerifier == address(0)) {
            // Deploy mock implementation
            vm.broadcast(msg.sender);
            MockEspressoNitroTEEVerifier mockImpl = new MockEspressoNitroTEEVerifier();
            vm.label(address(mockImpl), "MockNitroTEEVerifierImpl");
            implAddress = address(mockImpl);

            // Upgrade proxy to point to mock implementation
            vm.broadcast(msg.sender);
            deployment.proxyAdmin.upgrade(payable(address(deployment.proxy)), implAddress);
        } else {
            // Deploy production implementation
            address teeVerifierAddress = input.teeVerifierAddress();

            vm.broadcast(msg.sender);
            EspressoNitroTEEVerifier impl = new EspressoNitroTEEVerifier();
            vm.label(address(impl), "NitroTEEVerifierImpl");
            implAddress = address(impl);

            // Initialize the proxy
            bytes memory initData = abi.encodeCall(
                EspressoNitroTEEVerifier.initialize, (teeVerifierAddress, INitroEnclaveVerifier(nitroEnclaveVerifier))
            );
            vm.broadcast(msg.sender);
            deployment.proxyAdmin.upgradeAndCall(payable(address(deployment.proxy)), implAddress, initData);
        }

        // Transfer ownership if needed
        if (proxyAdminOwner != msg.sender) {
            vm.broadcast(msg.sender);
            deployment.proxyAdmin.transferOwnership(proxyAdminOwner);
        }

        // Set outputs
        output.set(output.nitroTEEVerifierProxy.selector, address(deployment.proxy));
        output.set(output.nitroTEEVerifierImpl.selector, implAddress);
        output.set(output.proxyAdmin.selector, address(deployment.proxyAdmin));

        return IEspressoNitroTEEVerifier(address(deployment.proxy));
    }

    function checkOutput(DeployAWSNitroVerifierOutput output) public view {
        address[] memory addresses =
            Solarray.addresses(output.nitroTEEVerifierProxy(), output.nitroTEEVerifierImpl(), output.proxyAdmin());
        for (uint256 i = 0; i < addresses.length; i++) {
            require(addresses[i] != address(0) && addresses[i].code.length > 0, "DeployAWSNitroVerifier: invalid address");
        }
    }
}
