// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import { CertManager } from "@nitro-validator/CertManager.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { EspressoNitroTEEVerifier } from "@espresso-tee-contracts/EspressoNitroTEEVerifier.sol";
import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";
import { Script } from "forge-std/Script.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";

contract DeployAWSNitroVerifierInput is BaseDeployIO {
    bytes32 internal _enclaveHash;

    function set(bytes4 _sel, bytes32 _val) public {
        if (_sel == this.enclaveHash.selector) _enclaveHash = _val;
        else revert("DeployAWSNitroVerifierInput: unknown selector");
    }

    function enclaveHash() public view returns (bytes32) {
        require(_enclaveHash != 0, "DeployAWSNitroVerifierInput: enclaveHash not set");
        return _enclaveHash;
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
        CertManager manager = deployCertManager();
        deployNitroTEEVerifier(input, output, manager);
        checkOutput(output);
    }

    function deployNitroTEEVerifier(
        DeployAWSNitroVerifierInput input,
        DeployAWSNitroVerifierOutput output,
        CertManager certManager
    )
        public
        returns (IEspressoNitroTEEVerifier)
    {
        bytes32 enclaveHash = input.enclaveHash();
        vm.broadcast(msg.sender);
        IEspressoNitroTEEVerifier impl = new EspressoNitroTEEVerifier(enclaveHash, certManager);
        vm.label(address(impl), "NitroTEEVerifierImpl");
        output.set(output.nitroTEEVerifierAddress.selector, address(impl));
        return impl;
    }

    function deployCertManager() public returns (CertManager) {
        vm.broadcast(msg.sender);
        CertManager impl = new CertManager();
        vm.label(address(impl), "CertManagerImpl");
        return impl;
    }

    function checkOutput(DeployAWSNitroVerifierOutput output) public view {
        address[] memory addresses = Solarray.addresses(address(output.nitroTEEVerifierAddress()));
        DeployUtils.assertValidContractAddresses(addresses);
    }
}
