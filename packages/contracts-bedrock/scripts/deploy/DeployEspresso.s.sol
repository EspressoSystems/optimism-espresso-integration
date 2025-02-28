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
import { CertManager, ICertManager } from "@nitro-validator/CertManager.sol";

contract DeployEspressoInput is BaseDeployIO {
    bytes32 internal _salt;
    address internal _preApprovedBatcherKey;

    function set(bytes4 _sel, bytes32 _val) public {
        if (_sel == this.salt.selector) _salt = _val;
        else revert("DeployEspressoInput: unknown selector");
    }

    function set(bytes4 _sel, address _val) public {
        if (_sel == this.preApprovedBatcherKey.selector) {
            _preApprovedBatcherKey = _val;
        } else {
            revert("DeployEspressoInput: unknown selector");
        }
    }

    function salt() public view returns (bytes32) {
        require(_salt != 0, "DeployEspressoInput: salt not set");
        return _salt;
    }

    function preApprovedBatcherKey() public view returns (address) {
        return _preApprovedBatcherKey;
    }
}

contract DeployEspressoOutput is BaseDeployIO {
    address internal _batchInboxAddress;

    function set(bytes4 _sel, address _addr) public {
        require(_addr != address(0), "DeployEspressoOutput: cannot set zero address");
        if (_sel == this.batchInboxAddress.selector) {
            _batchInboxAddress = _addr;
        } else {
            revert("DeployEspressoOutput: unknown selector");
        }
    }

    function batchInboxAddress() public view returns (address) {
        require(_batchInboxAddress != address(0), "DeployEspressoOutput: batcher inbox address not set");
        return _batchInboxAddress;
    }
}

contract DeployEspresso is Script {
    function run(DeployEspressoInput _ei, DeployEspressoOutput _eo) public {
        CertManager certManager = deployCertManager();
        deployBatchInbox(_ei, _eo, certManager);
        checkOutput(_eo);
    }

    function deployCertManager() public returns (CertManager) {
        vm.broadcast(msg.sender);
        CertManager impl = new CertManager();
        vm.label(address(impl), "CertManagerImpl");
        return impl;
    }

    function deployBatchInbox(DeployEspressoInput _ei, DeployEspressoOutput _eo, CertManager _certManager) public {
        bytes32 salt = _ei.salt();
        address preApprovedBatcher = _ei.preApprovedBatcherKey();
        vm.broadcast(msg.sender);
        IBatchInbox impl = IBatchInbox(
            DeployUtils.create2({
                _name: "BatchInbox",
                _salt: salt,
                _args: DeployUtils.encodeConstructor(
                    abi.encodeCall(IBatchInbox.__constructor__, (_certManager, preApprovedBatcher))
                )
            })
        );
        vm.label(address(impl), "BatchInboxImpl");
        _eo.set(_eo.batchInboxAddress.selector, address(impl));
    }

    function checkOutput(DeployEspressoOutput _eo) public view {
        address[] memory addresses = Solarray.addresses(address(_eo.batchInboxAddress()));
        DeployUtils.assertValidContractAddresses(addresses);
    }
}
