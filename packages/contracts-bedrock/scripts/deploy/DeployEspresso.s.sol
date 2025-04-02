// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";
import { IDataAvailabilityChallenge } from "interfaces/L1/IDataAvailabilityChallenge.sol";
import { IProxy } from "interfaces/universal/IProxy.sol";
import { IBatchInbox } from "interfaces/L1/IBatchInbox.sol";
import { Script } from "forge-std/Script.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { IProxyAdmin } from "interfaces/universal/IProxyAdmin.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { IBatchVerifier } from "interfaces/L1/IBatchVerifier.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { EspressoTEEVerifier } from "@espresso-tee-contracts/EspressoTEEVerifier.sol";

contract DeployEspressoInput is BaseDeployIO {
    bytes32 internal _salt;
    address internal _preApprovedBatcherKey;
    address internal _nitroTEEVerifier;

    function set(bytes4 _sel, bytes32 _val) public {
        if (_sel == this.salt.selector) _salt = _val;
        else revert("DeployEspressoInput: unknown selector");
    }

    function set(bytes4 _sel, address _val) public {
        if (_sel == this.preApprovedBatcherKey.selector) {
            _preApprovedBatcherKey = _val;
        } else if (_sel == this.nitroTEEVerifier.selector) {
            _nitroTEEVerifier = _val;
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

    function preApprovedBatcherKey() public view returns (address) {
        return _preApprovedBatcherKey;
    }
}

contract DeployEspressoOutput is BaseDeployIO {
    address internal _batchInboxAddress;
    address internal _batchVerifierAddress;

    function set(bytes4 _sel, address _addr) public {
        require(_addr != address(0), "DeployEspressoOutput: cannot set zero address");
        if (_sel == this.batchInboxAddress.selector) {
            _batchInboxAddress = _addr;
        } else if (_sel == this.batchVerifierAddress.selector) {
            _batchVerifierAddress = _addr;
        } else {
            revert("DeployEspressoOutput: unknown selector");
        }
    }

    function batchVerifierAddress() public view returns (address) {
        require(_batchVerifierAddress != address(0), "DeployEspressoOutput: batch verifier address not set");
        return _batchVerifierAddress;
    }

    function batchInboxAddress() public view returns (address) {
        require(_batchInboxAddress != address(0), "DeployEspressoOutput: batcher inbox address not set");
        return _batchInboxAddress;
    }
}

contract DeployEspresso is Script {
    function run(DeployEspressoInput input, DeployEspressoOutput output) public {
        IEspressoTEEVerifier teeVerifier = deployTEEVerifier(input);
        IBatchVerifier batchVerifier = deployBatchVerifier(input, output, teeVerifier);
        deployBatchInbox(input, output, batchVerifier);
        checkOutput(output);
    }

    function deployBatchVerifier(
        DeployEspressoInput input,
        DeployEspressoOutput output,
        IEspressoTEEVerifier teeVerifier
    )
        public
        returns (IBatchVerifier)
    {
        bytes32 salt = input.salt();
        address preApprovedBatcherKey = input.preApprovedBatcherKey();
        vm.broadcast(msg.sender);
        IBatchVerifier impl = IBatchVerifier(
            DeployUtils.create2({
                _name: "BatchVerifier",
                _salt: salt,
                _args: DeployUtils.encodeConstructor(
                    abi.encodeCall(IBatchVerifier.__constructor__, (address(teeVerifier), preApprovedBatcherKey))
                )
            })
        );
        vm.label(address(impl), "BatchVerifierImpl");
        output.set(output.batchVerifierAddress.selector, address(impl));
        return impl;
    }

    function deployTEEVerifier(DeployEspressoInput input) public returns (IEspressoTEEVerifier) {
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
        IBatchVerifier batchVerifier
    )
        public
    {
        bytes32 salt = input.salt();
        vm.broadcast(msg.sender);
        IBatchInbox impl = IBatchInbox(
            DeployUtils.create2({
                _name: "BatchInbox",
                _salt: salt,
                _args: DeployUtils.encodeConstructor(abi.encodeCall(IBatchInbox.__constructor__, (address(batchVerifier))))
            })
        );
        vm.label(address(impl), "BatchInboxImpl");
        output.set(output.batchInboxAddress.selector, address(impl));
    }

    function checkOutput(DeployEspressoOutput output) public view {
        address[] memory addresses =
            Solarray.addresses(address(output.batchVerifierAddress()), address(output.batchInboxAddress()));
        DeployUtils.assertValidContractAddresses(addresses);
    }
}
