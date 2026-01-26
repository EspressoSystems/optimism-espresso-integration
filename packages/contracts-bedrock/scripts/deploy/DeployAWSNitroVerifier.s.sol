// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { EspressoNitroTEEVerifier } from "@espresso-tee-contracts/EspressoNitroTEEVerifier.sol";
import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";
import { Script } from "forge-std/Script.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { INitroEnclaveVerifier } from "aws-nitro-enclave-attestation/interfaces/INitroEnclaveVerifier.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { ServiceType } from "@espresso-tee-contracts/types/Types.sol";

contract MockEspressoNitroTEEVerifier is IEspressoNitroTEEVerifier {
    constructor() { }

    function isSignerValid(address signer, ServiceType) external pure override returns (bool) {
        // Added this special condition for test TestE2eDevnetWithUnattestedBatcherKey
        if (signer == address(0xe16d5c4080C0faD6D2Ef4eb07C657674a217271C)) {
            return false;
        }
        return true;
    }

    function registeredEnclaveHash(bytes32, ServiceType) external pure override returns (bool) {
        return true;
    }

    function registerService(bytes calldata, bytes calldata, ServiceType) external override { }

    function setEnclaveHash(bytes32, bool, ServiceType) external override { }

    function deleteEnclaveHashes(bytes32[] memory, ServiceType) external override { }

    function setNitroEnclaveVerifier(address) external override { }

    function nitroEnclaveVerifier() external pure override returns (INitroEnclaveVerifier) {
        return INitroEnclaveVerifier(address(0));
    }

    function teeVerifier() external pure override returns (address) {
        return address(0);
    }
}

contract DeployAWSNitroVerifierInput is BaseDeployIO {
    bytes32 internal _enclaveHash;
    address internal _nitroEnclaveVerifier;

    function set(bytes4 _sel, bytes32 _val) public {
        if (_sel == this.enclaveHash.selector) _enclaveHash = _val;
        else revert("DeployAWSNitroVerifierInput: unknown selector");
    }

    function enclaveHash() public view returns (bytes32) {
        return _enclaveHash;
    }

    function set(bytes4 _sel, address _val) public {
        if (_sel == this.nitroEnclaveVerifier.selector) {
            _nitroEnclaveVerifier = _val;
        } else {
            revert("DeployAWSNitroVerifierInput: unknown selector");
        }
    }

    function nitroEnclaveVerifier() public view returns (address) {
        return _nitroEnclaveVerifier;
    }
}

contract DeployAWSNitroVerifierOutput is BaseDeployIO {
    address internal _nitroTEEVerifierAddress;

    function set(bytes4 _sel, address _addr) public {
        require(_addr != address(0), "DeployAWSNitroVerifierOutput: cannot set zero address");
        if (_sel == this.nitroTEEVerifierAddress.selector) {
            _nitroTEEVerifierAddress = _addr;
        } else {
            revert("DeployAWSNitroVerifierOutput: unknown selector");
        }
    }

    function nitroTEEVerifierAddress() public view returns (address) {
        require(_nitroTEEVerifierAddress != address(0), "nitro TEE verifier address not set");
        return _nitroTEEVerifierAddress;
    }
}

contract DeployAWSNitroVerifier is Script {
    function run(DeployAWSNitroVerifierInput input, DeployAWSNitroVerifierOutput output) public {
        deployNitroTEEVerifier(input, output);
        checkOutput(output);
    }

    function deployNitroTEEVerifier(
        DeployAWSNitroVerifierInput input,
        DeployAWSNitroVerifierOutput output
    )
        public
        returns (IEspressoNitroTEEVerifier)
    {
        address nitroEnclaveVerifier = input.nitroEnclaveVerifier();

        IEspressoNitroTEEVerifier proxyAddr;
        if (nitroEnclaveVerifier == address(0)) {
            // Deploy mock without proxy for testing
            vm.broadcast(msg.sender);
            proxyAddr = new MockEspressoNitroTEEVerifier();
            vm.label(address(proxyAddr), "MockNitroTEEVerifier");
        } else {
            // Deploy implementation
            vm.broadcast(msg.sender);
            EspressoNitroTEEVerifier impl = new EspressoNitroTEEVerifier();
            vm.label(address(impl), "NitroTEEVerifierImpl");

            // Prepare initialization data
            bytes memory initData = abi.encodeWithSelector(
                EspressoNitroTEEVerifier.initialize.selector,
                INitroEnclaveVerifier(nitroEnclaveVerifier),
                msg.sender // initial owner
            );

            // Deploy proxy
            vm.broadcast(msg.sender);
            ERC1967Proxy proxy = new ERC1967Proxy(address(impl), initData);
            vm.label(address(proxy), "NitroTEEVerifierProxy");
            proxyAddr = IEspressoNitroTEEVerifier(address(proxy));
        }

        output.set(output.nitroTEEVerifierAddress.selector, address(proxyAddr));
        return proxyAddr;
    }

    function checkOutput(DeployAWSNitroVerifierOutput output) public view {
        address[] memory addresses = Solarray.addresses(address(output.nitroTEEVerifierAddress()));
        DeployUtils.assertValidContractAddresses(addresses);
    }
}
