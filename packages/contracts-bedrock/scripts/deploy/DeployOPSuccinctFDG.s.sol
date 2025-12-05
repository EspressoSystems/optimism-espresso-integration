// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Libraries
import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";
import { GameType, Duration } from "src/dispute/lib/Types.sol";

// Interfaces
import { IDisputeGame } from "interfaces/dispute/IDisputeGame.sol";
import { IDisputeGameFactory } from "interfaces/dispute/IDisputeGameFactory.sol";
import { ISP1Verifier } from "src/dispute/succinct/ISP1Verifier.sol";
import { IAnchorStateRegistry } from "interfaces/dispute/IAnchorStateRegistry.sol";

// Contracts
import { AccessManager } from "src/dispute/succinct/AccessManager.sol";
import { OPSuccinctFaultDisputeGame } from "src/dispute/succinct/OPSuccinctFaultDisputeGame.sol";
import { SP1MockVerifier } from "src/dispute/succinct/ISP1Verifier.sol";

/// @title DeployOPSuccinctFDG
/// @notice Deployment script for OPSuccinctFaultDisputeGame and related contracts.
contract DeployOPSuccinctFDG is Script {
    // Storage variables to reduce stack usage
    address public factoryAddr;
    address public registryAddr;
    address public sp1VerifierAddr;
    address public accessManagerAddr;
    address public gameImplAddr;
    uint32 public gameTypeId;
    uint256 public initialBond;

    function run() public {
        _loadConfig();

        vm.startBroadcast();

        _deployAccessManager();
        _deployVerifier();
        _deployGame();

        // Factory configuration is done separately via cast commands
        // since it requires proxy admin privileges
        console.log("Factory configuration should be done via cast with impersonation");

        vm.stopBroadcast();

        console.log("=== Deployment Complete ===");
        console.log("AccessManager:", accessManagerAddr);
        console.log("SP1 Verifier:", sp1VerifierAddr);
        console.log("Game Implementation:", gameImplAddr);
        console.log("Game Type:", gameTypeId);
    }

    function _loadConfig() internal {
        factoryAddr = vm.envAddress("FACTORY_ADDRESS");
        registryAddr = vm.envAddress("ANCHOR_STATE_REGISTRY_ADDRESS");
        gameTypeId = uint32(vm.envOr("GAME_TYPE", uint256(42)));
        initialBond = vm.envOr("INITIAL_BOND_WEI", uint256(0.001 ether));
    }

    function _deployAccessManager() internal {
        AccessManager am = new AccessManager();
        accessManagerAddr = address(am);

        // Configure permissionless mode by default
        if (vm.envOr("PERMISSIONLESS_MODE", true)) {
            am.setProposer(address(0), true);
            am.setChallenger(address(0), true);
        }
    }

    function _deployVerifier() internal {
        if (vm.envOr("USE_SP1_MOCK_VERIFIER", true)) {
            SP1MockVerifier verifier = new SP1MockVerifier();
            sp1VerifierAddr = address(verifier);
        } else {
            sp1VerifierAddr = vm.envAddress("VERIFIER_ADDRESS");
        }
    }

    function _deployGame() internal {
        uint64 maxChallenge = uint64(vm.envOr("MAX_CHALLENGE_DURATION", uint256(300)));
        uint64 maxProve = uint64(vm.envOr("MAX_PROVE_DURATION", uint256(1800)));
        uint256 challengerBond = vm.envOr("CHALLENGER_BOND_WEI", uint256(0.001 ether));

        bytes32 configHash = bytes32(0);
        bytes32 aggVkey = bytes32(0);
        bytes32 rangeVkey = bytes32(0);

        if (!vm.envOr("USE_SP1_MOCK_VERIFIER", true)) {
            configHash = vm.envBytes32("ROLLUP_CONFIG_HASH");
            aggVkey = vm.envBytes32("AGGREGATION_VKEY");
            rangeVkey = vm.envBytes32("RANGE_VKEY_COMMITMENT");
        }

        OPSuccinctFaultDisputeGame game = new OPSuccinctFaultDisputeGame(
            Duration.wrap(maxChallenge),
            Duration.wrap(maxProve),
            IDisputeGameFactory(factoryAddr),
            ISP1Verifier(sp1VerifierAddr),
            configHash,
            aggVkey,
            rangeVkey,
            challengerBond,
            IAnchorStateRegistry(registryAddr),
            AccessManager(accessManagerAddr)
        );
        gameImplAddr = address(game);
    }
}
