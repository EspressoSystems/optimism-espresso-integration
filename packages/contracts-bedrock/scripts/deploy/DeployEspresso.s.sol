// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";
import { Script } from "forge-std/Script.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { IBatchAuthenticator } from "interfaces/L1/IBatchAuthenticator.sol";
import { IEspressoNitroTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoNitroTEEVerifier.sol";
import { IEspressoSGXTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoSGXTEEVerifier.sol";
import { IEspressoTEEVerifier } from "@espresso-tee-contracts/interface/IEspressoTEEVerifier.sol";
import { EspressoTEEVerifier } from "@espresso-tee-contracts/EspressoTEEVerifier.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { BatchAuthenticator } from "src/L1/BatchAuthenticator.sol";
import { MockEspressoTEEVerifier } from "test/mocks/MockEspressoTEEVerifiers.sol";

contract DeployEspressoInput is BaseDeployIO {
    bytes32 internal _salt;
    address internal _nitroTEEVerifier;
    address internal _teeBatcher;
    address internal _systemConfig;
    address internal _proxyAdminOwner;
    bool internal _useMockTEEVerifier;

    function set(bytes4 _sel, bytes32 _val) public {
        if (_sel == this.salt.selector) _salt = _val;
        else revert("DeployEspressoInput: unknown selector");
    }

    function set(bytes4 _sel, bool _val) public {
        if (_sel == this.useMockTEEVerifier.selector) _useMockTEEVerifier = _val;
        else revert("DeployEspressoInput: unknown selector");
    }

    function set(bytes4 _sel, address _val) public {
        if (_sel == this.nitroTEEVerifier.selector) {
            _nitroTEEVerifier = _val;
        } else if (_sel == this.teeBatcher.selector) {
            _teeBatcher = _val;
        } else if (_sel == this.systemConfig.selector) {
            _systemConfig = _val;
        } else if (_sel == this.proxyAdminOwner.selector) {
            _proxyAdminOwner = _val;
        } else {
            revert("DeployEspressoInput: unknown selector");
        }
    }

    function salt() public view returns (bytes32) {
        require(_salt != 0, "DeployEspressoInput: salt not set");
        return _salt;
    }

    /// @notice Address of the EspressoNitroTEEVerifier proxy (deployed via DeployAWSNitroVerifier)
    function nitroTEEVerifier() public view returns (address) {
        return _nitroTEEVerifier;
    }

    function teeBatcher() public view returns (address) {
        return _teeBatcher;
    }

    function systemConfig() public view returns (address) {
        return _systemConfig;
    }

    /// @notice If true, deploy MockEspressoTEEVerifier instead of production EspressoTEEVerifier.
    ///         Defaults to false. Also uses mock if nitroTEEVerifier is address(0).
    function useMockTEEVerifier() public view returns (bool) {
        return _useMockTEEVerifier;
    }

    /// @notice The address that will own the ProxyAdmin contracts. Defaults to msg.sender if not set.
    function proxyAdminOwner() public view returns (address) {
        return _proxyAdminOwner;
    }
}

contract DeployEspressoOutput is BaseDeployIO {
    address internal _batchAuthenticatorAddress;
    address internal _teeVerifierProxy;
    address internal _teeVerifierImpl;
    address internal _teeVerifierProxyAdmin;

    function set(bytes4 _sel, address _addr) public {
        require(_addr != address(0), "DeployEspressoOutput: cannot set zero address");
        if (_sel == this.batchAuthenticatorAddress.selector) {
            _batchAuthenticatorAddress = _addr;
        } else if (_sel == this.teeVerifierProxy.selector) {
            _teeVerifierProxy = _addr;
        } else if (_sel == this.teeVerifierImpl.selector) {
            _teeVerifierImpl = _addr;
        } else if (_sel == this.teeVerifierProxyAdmin.selector) {
            _teeVerifierProxyAdmin = _addr;
        } else {
            revert("DeployEspressoOutput: unknown selector");
        }
    }

    function batchAuthenticatorAddress() public view returns (address) {
        require(_batchAuthenticatorAddress != address(0), "DeployEspressoOutput: batch authenticator address not set");
        return _batchAuthenticatorAddress;
    }

    function teeVerifierProxy() public view returns (address) {
        require(_teeVerifierProxy != address(0), "DeployEspressoOutput: tee verifier proxy not set");
        return _teeVerifierProxy;
    }

    function teeVerifierImpl() public view returns (address) {
        require(_teeVerifierImpl != address(0), "DeployEspressoOutput: tee verifier impl not set");
        return _teeVerifierImpl;
    }

    function teeVerifierProxyAdmin() public view returns (address) {
        require(_teeVerifierProxyAdmin != address(0), "DeployEspressoOutput: tee verifier proxy admin not set");
        return _teeVerifierProxyAdmin;
    }

    /// @notice Alias for teeVerifierProxy for convenience
    function teeVerifierAddress() public view returns (address) {
        return teeVerifierProxy();
    }
}

contract DeployEspresso is Script {
    function run(DeployEspressoInput input, DeployEspressoOutput output, address deployerAddress) public {
        IEspressoTEEVerifier teeVerifier = deployTEEVerifier(input, output, deployerAddress);
        deployBatchAuthenticator(input, output, teeVerifier);
        checkOutput(input, output);
    }

    function deployBatchAuthenticator(
        DeployEspressoInput input,
        DeployEspressoOutput output,
        IEspressoTEEVerifier teeVerifier
    )
        public
        returns (IBatchAuthenticator)
    {
        // Deploy the proxy admin, the proxy, and the batch authenticator implementation.
        // We create ProxyAdmin with msg.sender as the owner to ensure broadcasts come from
        // the expected address, then transfer ownership to proxyAdminOwner afterward.
        vm.broadcast(msg.sender);
        ProxyAdmin proxyAdmin = new ProxyAdmin(msg.sender);
        vm.label(address(proxyAdmin), "BatchAuthenticatorProxyAdmin");
        vm.broadcast(msg.sender);
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.label(address(proxy), "BatchAuthenticatorProxy");
        vm.broadcast(msg.sender);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);
        vm.broadcast(msg.sender);
        BatchAuthenticator impl = new BatchAuthenticator();
        vm.label(address(impl), "BatchAuthenticatorImpl");

        // Determine the desired BatchAuthenticator owner
        address batchAuthenticatorOwner = input.proxyAdminOwner();
        if (batchAuthenticatorOwner == address(0)) {
            batchAuthenticatorOwner = msg.sender;
        }

        // Initialize the proxy with explicit owner parameter
        bytes memory initData = abi.encodeCall(
            BatchAuthenticator.initialize,
            (teeVerifier, input.teeBatcher(), input.systemConfig(), batchAuthenticatorOwner)
        );
        vm.broadcast(msg.sender);
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(impl), initData);

        // Transfer ProxyAdmin ownership to the desired proxyAdminOwner if different from msg.sender.
        address proxyAdminOwner = input.proxyAdminOwner();
        if (proxyAdminOwner == address(0)) {
            proxyAdminOwner = msg.sender;
        }
        if (proxyAdminOwner != msg.sender) {
            vm.broadcast(msg.sender);
            proxyAdmin.transferOwnership(proxyAdminOwner);
        }

        // Return the proxied contract.
        IBatchAuthenticator batchAuthenticator = IBatchAuthenticator(address(proxy));
        output.set(output.batchAuthenticatorAddress.selector, address(batchAuthenticator));
        return batchAuthenticator;
    }

    function deployTEEVerifier(
        DeployEspressoInput input,
        DeployEspressoOutput output,
        address deployerAddress
    )
        public
        returns (IEspressoTEEVerifier)
    {
        IEspressoNitroTEEVerifier nitroTEEVerifier = IEspressoNitroTEEVerifier(input.nitroTEEVerifier());
        // OP only uses Nitro TEE, not SGX
        IEspressoSGXTEEVerifier sgxTEEVerifier = IEspressoSGXTEEVerifier(address(0));

        // Use mock if explicitly requested or if nitroTEEVerifier is not set
        if (input.useMockTEEVerifier() || address(nitroTEEVerifier) == address(0)) {
            vm.broadcast(msg.sender);
            MockEspressoTEEVerifier mockImpl = new MockEspressoTEEVerifier(nitroTEEVerifier);
            vm.label(address(mockImpl), "MockEspressoTEEVerifier");

            // For mock deployments, we still need valid distinct addresses for the output.
            // Deploy a minimal ProxyAdmin to satisfy the output requirements, even though
            // the mock doesn't use it. This ensures checkOutput validation passes.
            address mockProxyAdminOwner = input.proxyAdminOwner();
            if (mockProxyAdminOwner == address(0)) {
                mockProxyAdminOwner = msg.sender;
            }
            vm.broadcast(msg.sender);
            ProxyAdmin mockProxyAdmin = new ProxyAdmin(mockProxyAdminOwner);
            vm.label(address(mockProxyAdmin), "MockTEEVerifierProxyAdmin");

            output.set(output.teeVerifierProxy.selector, address(mockImpl));
            output.set(output.teeVerifierImpl.selector, address(mockImpl));
            output.set(output.teeVerifierProxyAdmin.selector, address(mockProxyAdmin));
            return IEspressoTEEVerifier(address(mockImpl));
        }

        // Production deployment: Proxy + ProxyAdmin pattern

        // 1. Deploy the ProxyAdmin
        vm.broadcast(msg.sender);
        ProxyAdmin proxyAdmin = new ProxyAdmin(msg.sender);
        vm.label(address(proxyAdmin), "TEEVerifierProxyAdmin");

        // 2. Deploy the Proxy
        vm.broadcast(msg.sender);
        Proxy proxy = new Proxy(address(proxyAdmin));
        vm.label(address(proxy), "TEEVerifierProxy");

        // 3. Set proxy type
        vm.broadcast(msg.sender);
        proxyAdmin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);

        // 4. Deploy the EspressoTEEVerifier implementation
        vm.broadcast(msg.sender);
        EspressoTEEVerifier impl = new EspressoTEEVerifier();
        vm.label(address(impl), "TEEVerifierImpl");

        // 5. Initialize the proxy
        // Note: EspressoTEEVerifier.initialize takes (owner, sgxVerifier, nitroVerifier)
        bytes memory initData =
            abi.encodeCall(EspressoTEEVerifier.initialize, (deployerAddress, sgxTEEVerifier, nitroTEEVerifier));
        vm.broadcast(msg.sender);
        proxyAdmin.upgradeAndCall(payable(address(proxy)), address(impl), initData);

        // 6. Transfer ownership to the desired proxyAdminOwner if different from msg.sender
        address proxyAdminOwner = input.proxyAdminOwner();
        if (proxyAdminOwner == address(0)) {
            proxyAdminOwner = msg.sender;
        }
        if (proxyAdminOwner != msg.sender) {
            vm.broadcast(msg.sender);
            proxyAdmin.transferOwnership(proxyAdminOwner);
        }

        // 7. Set outputs
        output.set(output.teeVerifierProxy.selector, address(proxy));
        output.set(output.teeVerifierImpl.selector, address(impl));
        output.set(output.teeVerifierProxyAdmin.selector, address(proxyAdmin));

        return IEspressoTEEVerifier(address(proxy));
    }

    function checkOutput(DeployEspressoInput input, DeployEspressoOutput output) public view {
        // Check core addresses
        address[] memory coreAddresses =
            Solarray.addresses(output.batchAuthenticatorAddress(), output.teeVerifierProxy());
        DeployUtils.assertValidContractAddresses(coreAddresses);

        // Check that proxy admin is a valid, distinct address (applies to both mock and production)
        address[] memory adminAddresses = Solarray.addresses(output.teeVerifierProxyAdmin());
        DeployUtils.assertValidContractAddresses(adminAddresses);
        require(
            output.teeVerifierProxy() != output.teeVerifierProxyAdmin(),
            "DeployEspresso: proxy and proxy admin should be different"
        );

        // For production deployment, also check impl is valid and distinct from proxy
        if (!input.useMockTEEVerifier() && input.nitroTEEVerifier() != address(0)) {
            address[] memory teeAddresses =
                Solarray.addresses(output.teeVerifierProxy(), output.teeVerifierImpl(), output.teeVerifierProxyAdmin());
            DeployUtils.assertValidContractAddresses(teeAddresses);
            require(
                output.teeVerifierProxy() != output.teeVerifierImpl(),
                "DeployEspresso: proxy and impl should be different"
            );
        }
    }
}
