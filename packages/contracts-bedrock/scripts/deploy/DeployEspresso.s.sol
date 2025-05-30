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
    function run(DeployEspressoInput input, DeployEspressoOutput output) public {
        IEspressoTEEVerifier teeVerifier = deployTEEVerifier(input);
        IBatchAuthenticator batchAuthenticator = deployBatchAuthenticator(input, output, teeVerifier);
        deployBatchInbox(input, output, batchAuthenticator);
        checkOutput(output);
    }

    function deployBatchAuthenticator(
        DeployEspressoInput input,
        DeployEspressoOutput output,
        IEspressoTEEVerifier teeVerifier
    )
        public
        returns (IBatchAuthenticator)
    {
        bytes32 salt = input.salt();
        address preApprovedBatcherKey = input.preApprovedBatcherKey();
        vm.broadcast(msg.sender);
        IBatchAuthenticator impl = IBatchAuthenticator(
            DeployUtils.create2({
                _name: "BatchAuthenticator",
                _salt: salt,
                _args: DeployUtils.encodeConstructor(
                    abi.encodeCall(IBatchAuthenticator.__constructor__, (address(teeVerifier), preApprovedBatcherKey))
                )
            })
        );
        vm.label(address(impl), "BatchAuthenticatorImpl");
        output.set(output.batchAuthenticatorAddress.selector, address(impl));
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
        IBatchAuthenticator batchAuthenticator
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
                    abi.encodeCall(IBatchInbox.__constructor__, (address(batchAuthenticator)))
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
