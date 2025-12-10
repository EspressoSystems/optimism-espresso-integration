// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { EspressoNitroTEEVerifier } from "@espresso-tee-contracts/EspressoNitroTEEVerifier.sol";
import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";
import { Script } from "forge-std/Script.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { INitroEnclaveVerifier } from "aws-nitro-enclave-attestation/interfaces/INitroEnclaveVerifier.sol";

contract MockEspressoNitroTEEVerifier is IEspressoNitroTEEVerifier {
    constructor() { }

    function registeredSigners(address signer) external pure override returns (bool) {
        // Added this special condition for test TestE2eDevnetWithUnattestedBatcherKey
        if (signer == address(0xe16d5c4080C0faD6D2Ef4eb07C657674a217271C)) {
            return false;
        }
        return true;
    }

    function registeredEnclaveHash(bytes32) external pure override returns (bool) {
        return true;
    }

    function registerSigner(bytes calldata, bytes calldata) external override { }

    function setEnclaveHash(bytes32, bool) external override { }

    function deleteRegisteredSigners(address[] memory) external override { }
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
        vm.broadcast(msg.sender);
        bytes32 enclaveHash = input.enclaveHash();
        address nitroEnclaveVerifier = input.nitroEnclaveVerifier();

        IEspressoNitroTEEVerifier impl;
        if (nitroEnclaveVerifier == address(0)) {
            impl = new MockEspressoNitroTEEVerifier();
        } else {
            impl = new EspressoNitroTEEVerifier(enclaveHash, INitroEnclaveVerifier(nitroEnclaveVerifier));
        }
        vm.label(address(impl), "NitroTEEVerifierImpl");
        output.set(output.nitroTEEVerifierAddress.selector, address(impl));
        return impl;
    }

    function checkOutput(DeployAWSNitroVerifierOutput output) public view {
        address[] memory addresses = Solarray.addresses(address(output.nitroTEEVerifierAddress()));
        DeployUtils.assertValidContractAddresses(addresses);
    }
}
