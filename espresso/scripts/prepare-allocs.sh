#!/usr/bin/env bash
set -euxo pipefail

source .env

ANVIL_PORT=8545
ANVIL_URL=http://localhost:$ANVIL_PORT

# All variables must be set

OP_ROOT="${1:-$(pwd)/..}"
OP_ROOT=$(realpath "${OP_ROOT}")

DEPLOYMENT_DIR="${OP_ROOT}/espresso/deployment"
DEPLOYER_DIR="${DEPLOYMENT_DIR}/deployer"
L1_CONFIG_DIR="${DEPLOYMENT_DIR}/l1-config"
mkdir -p "${DEPLOYER_DIR}"
mkdir -p "${L1_CONFIG_DIR}"

ANVIL_STATE_FILE="${DEPLOYMENT_DIR}/anvil_state.json"
ARTIFACTS_DIR="file:///${OP_ROOT}/packages/contracts-bedrock/forge-artifacts"

# Start anvil in dev mode and save PID to kill later
anvil --port $ANVIL_PORT --chain-id "${L1_CHAIN_ID}" --disable-gas-limit --disable-code-size-limit --dump-state "${ANVIL_STATE_FILE}" &
ANVIL_PID=$!
echo "Started anvil in dev mode with PID: $ANVIL_PID"

# Function to cleanup anvil process
cleanup() {
    if kill -0 $ANVIL_PID > /dev/null 2>&1; then
        echo "Stopping anvil (PID: $ANVIL_PID)"
        kill $ANVIL_PID
    fi
}
trap cleanup EXIT

# Give anvil a moment to start up
sleep 1

cast rpc anvil_setBalance "${OPERATOR_ADDRESS}" 0x100000000000000000000000000000000000
cast rpc anvil_setBalance "${PROPOSER_ADDRESS}" 0x100000000000000000000000000000000000

op-deployer bootstrap proxy \
                      --l1-rpc-url="${ANVIL_URL}" \
                      --private-key="${OPERATOR_PRIVATE_KEY}" \
                      --artifacts-locator="${ARTIFACTS_DIR}" \
                      --proxy-owner="${OPERATOR_ADDRESS}"

export LOG_LEVEL=debug

op-deployer bootstrap superchain \
                      --l1-rpc-url="${ANVIL_URL}" \
                      --private-key="${OPERATOR_PRIVATE_KEY}" \
                      --artifacts-locator="${ARTIFACTS_DIR}" \
                      --outfile="${DEPLOYER_DIR}/bootstrap_superchain.json" \
                      --superchain-proxy-admin-owner="${OPERATOR_ADDRESS}" \
                      --protocol-versions-owner="${OPERATOR_ADDRESS}" \
                      --guardian="${OPERATOR_ADDRESS}"

op-deployer bootstrap implementations \
                      --l1-rpc-url="${ANVIL_URL}" \
                      --private-key="${OPERATOR_PRIVATE_KEY}" \
                      --artifacts-locator="${ARTIFACTS_DIR}" \
                      --protocol-versions-proxy=`jq -r .protocolVersionsProxyAddress < ${DEPLOYER_DIR}/bootstrap_superchain.json` \
                      --superchain-config-proxy=`jq -r .superchainConfigProxyAddress < ${DEPLOYER_DIR}/bootstrap_superchain.json` \
                      --superchain-proxy-admin=`jq -r .proxyAdminAddress < ${DEPLOYER_DIR}/bootstrap_superchain.json` \
                      --upgrade-controller="${OPERATOR_ADDRESS}" \
                      --challenger="${OPERATOR_ADDRESS}" \
                      --outfile="${DEPLOYER_DIR}/bootstrap_implementations.json"

op-deployer init --l1-chain-id "${L1_CHAIN_ID}" \
                 --l2-chain-ids "${L2_CHAIN_ID}" \
                 --intent-type standard-overrides \
                 --outdir ${DEPLOYER_DIR}

dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].espressoEnabled -t bool -v true

# Configure Espresso batchers for devnet. We reuse the operator address for both
# the non-TEE and TEE batchers to ensure they are non-zero and consistent.
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].nonTeeBatcher -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].teeBatcher -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .l1ContractsLocator -v "${ARTIFACTS_DIR}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .l2ContractsLocator -v "${ARTIFACTS_DIR}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .opcmAddress -v `jq -r .opcmAddress < ${DEPLOYER_DIR}/bootstrap_implementations.json`
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .fundDevAccounts -t bool -v true
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .globalDeployOverrides.faultGameMaxClockDuration -t int -v 302400
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .globalDeployOverrides.faultGameClockExtension -t int -v 10800
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .globalDeployOverrides.preimageOracleChallengePeriod -t int -v 0
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .globalDeployOverrides.dangerouslyAllowCustomDisputeParameters -t bool -v true
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .globalDeployOverrides.proofMaturityDelaySeconds -t int -v 12
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .globalDeployOverrides.disputeGameFinalityDelaySeconds -t int -v 6
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].baseFeeVaultRecipient -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].l1FeeVaultRecipient -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].sequencerFeeVaultRecipient -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].roles.systemConfigOwner -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].roles.unsafeBlockSigner -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].roles.batcher -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].roles.proposer -v "${PROPOSER_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].roles.l1ProxyAdminOwner -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].roles.challenger -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].dangerousAltDAConfig.useAltDA -t bool -v true
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].dangerousAltDAConfig.daCommitmentType -v "GenericCommitment"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].dangerousAltDAConfig.daChallengeWindow -t int -v 6
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].dangerousAltDAConfig.daResolveWindow -t int -v 1

# Fill in a specified create2Salt for the deployer, in order to ensure that the
# contract addresses are deterministic.
dasel put -f "${DEPLOYER_DIR}/state.json" -s create2Salt -v "0xaecea4f57fadb2097ccd56594f2f22715ac52f92971c5913b70a7f1134b68feb"

BATCH_AUTHENTICATOR_OWNER_ADDRESS="${BATCH_AUTHENTICATOR_OWNER_ADDRESS}" op-deployer apply --l1-rpc-url "${ANVIL_URL}" \
                  --workdir "${DEPLOYER_DIR}" \
                  --private-key="${OPERATOR_PRIVATE_KEY}"

# =====================================================================
# Deploy op-succinct contracts (OPSuccinctFaultDisputeGame, AccessManager, SP1MockVerifier)
# =====================================================================
echo "Deploying op-succinct contracts..."

# Configuration for op-succinct
GAME_TYPE=42  # Succinct game type
MAX_CHALLENGE_DURATION=300  # 5 minutes for devnet
MAX_PROVE_DURATION=1800     # 30 minutes for devnet
FALLBACK_TIMEOUT_FP_SECS=3600  # 1 hour for devnet
INITIAL_BOND_WEI=1000000000000000  # 0.001 ETH
CHALLENGER_BOND_WEI=1000000000000000  # 0.001 ETH
DISPUTE_GAME_FINALITY_DELAY_SECONDS=6  # Fast for devnet

# Read existing contract addresses from state.json
DISPUTE_GAME_FACTORY=$(jq -r '.opChainDeployments[0].DisputeGameFactoryProxy' "${DEPLOYER_DIR}/state.json")
OPTIMISM_PORTAL=$(jq -r '.opChainDeployments[0].OptimismPortalProxy' "${DEPLOYER_DIR}/state.json")
ANCHOR_STATE_REGISTRY=$(jq -r '.opChainDeployments[0].AnchorStateRegistryProxy' "${DEPLOYER_DIR}/state.json")

echo "Existing contract addresses:"
echo "  DisputeGameFactory: ${DISPUTE_GAME_FACTORY}"
echo "  OptimismPortal: ${OPTIMISM_PORTAL}"
echo "  AnchorStateRegistry: ${ANCHOR_STATE_REGISTRY}"

# Get the starting output root from the anchor state registry
# The AnchorStateRegistry was initialized by op-deployer with a starting anchor root
STARTING_ROOT=$(cast call "${ANCHOR_STATE_REGISTRY}" "getAnchorRoot()(bytes32,uint256)" --rpc-url "${ANVIL_URL}" | head -1)
STARTING_L2_BLOCK=$(cast call "${ANCHOR_STATE_REGISTRY}" "getAnchorRoot()(bytes32,uint256)" --rpc-url "${ANVIL_URL}" | tail -1)

echo "Starting anchor state:"
echo "  Root: ${STARTING_ROOT}"
echo "  L2 Block: ${STARTING_L2_BLOCK}"

# Always create succinct.env with contract addresses (even if op-succinct repo not found)
# Note: op-succinct uses FACTORY_ADDRESS not DISPUTE_GAME_FACTORY_ADDRESS
SUCCINCT_ENV_FILE="${DEPLOYER_DIR}/succinct.env"
cat > "${SUCCINCT_ENV_FILE}" << EOF
FACTORY_ADDRESS=${DISPUTE_GAME_FACTORY}
DISPUTE_GAME_FACTORY_ADDRESS=${DISPUTE_GAME_FACTORY}
OPTIMISM_PORTAL2_ADDRESS=${OPTIMISM_PORTAL}
ANCHOR_STATE_REGISTRY_ADDRESS=${ANCHOR_STATE_REGISTRY}
EOF
echo "Created succinct env file at ${SUCCINCT_ENV_FILE}"

# Check if op-succinct-espresso repo exists
OP_SUCCINCT_DIR="${OP_ROOT}/../op-succinct-espresso"
if [ ! -d "${OP_SUCCINCT_DIR}" ]; then
    echo "Warning: op-succinct-espresso repo not found at ${OP_SUCCINCT_DIR}"
    echo "Skipping op-succinct contract deployment."
    echo "To enable succinct: git clone https://github.com/EspressoSystems/op-succinct.git ${OP_SUCCINCT_DIR}"
else
    # Get the L1 proxy admin owner from state.json (this is the factory owner)
    L1_PROXY_ADMIN_OWNER=$(jq -r '.appliedIntent.chains[0].roles.l1ProxyAdminOwner' "${DEPLOYER_DIR}/state.json")
    echo "L1 Proxy Admin Owner: ${L1_PROXY_ADMIN_OWNER}"

    # Create the op-succinct FDG config JSON
    # Note: activateContracts is false because op-succinct calls setRespectedGameType
    # on OptimismPortal2, but in our fork this function is on AnchorStateRegistry.
    # We'll activate manually after deployment.
    OP_SUCCINCT_CONFIG="${DEPLOYER_DIR}/opsuccinctfdgconfig.json"
    cat > "${OP_SUCCINCT_CONFIG}" << EOF
{
  "activateContracts": false,
  "aggregationVkey": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "anchorStateRegistryAddress": "${ANCHOR_STATE_REGISTRY}",
  "celoSuperchainConfigAddress": "0x0000000000000000000000000000000000000000",
  "challengerAddresses": ["${OPERATOR_ADDRESS}"],
  "challengerBondWei": ${CHALLENGER_BOND_WEI},
  "configureContracts": true,
  "disputeGameFactoryAddress": "${DISPUTE_GAME_FACTORY}",
  "disputeGameFinalityDelaySeconds": ${DISPUTE_GAME_FINALITY_DELAY_SECONDS},
  "fallbackTimeoutFpSecs": ${FALLBACK_TIMEOUT_FP_SECS},
  "gameType": ${GAME_TYPE},
  "initialBondWei": ${INITIAL_BOND_WEI},
  "maxChallengeDuration": ${MAX_CHALLENGE_DURATION},
  "maxProveDuration": ${MAX_PROVE_DURATION},
  "optimismPortal2Address": "${OPTIMISM_PORTAL}",
  "permissionlessMode": false,
  "proposerAddresses": ["${PROPOSER_ADDRESS}"],
  "rangeVkeyCommitment": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "rollupConfigHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "startingL2BlockNumber": ${STARTING_L2_BLOCK},
  "startingRoot": "${STARTING_ROOT}",
  "useSp1MockVerifier": true,
  "verifierAddress": "0x0000000000000000000000000000000000000000"
}
EOF

    echo "Created op-succinct config at ${OP_SUCCINCT_CONFIG}"

    # Copy config to op-succinct contracts directory
    cp "${OP_SUCCINCT_CONFIG}" "${OP_SUCCINCT_DIR}/contracts/opsuccinctfdgconfig.json"

    # Deploy op-succinct contracts using forge
    pushd "${OP_SUCCINCT_DIR}/contracts"

    # Install dependencies if not already installed
    if [ ! -d "lib/forge-std" ]; then
        echo "Installing forge dependencies..."
        forge install --no-commit
    fi

    # Build contracts
    echo "Building op-succinct contracts..."
    forge build

    # Run the deployment script
    echo "Running op-succinct deployment script..."
    forge script script/fp/DeployOPSuccinctFDG.s.sol \
        --broadcast \
        --no-storage-caching \
        --slow \
        --rpc-url "${ANVIL_URL}" \
        --private-key "${OPERATOR_PRIVATE_KEY}" || echo "Warning: op-succinct deployment failed, continuing..."

    popd

    # Manually activate game type 42 on AnchorStateRegistry
    # (op-succinct script calls wrong contract for setRespectedGameType)
    echo "Activating game type ${GAME_TYPE} on AnchorStateRegistry..."
    cast send "${ANCHOR_STATE_REGISTRY}" \
        "setRespectedGameType(uint32)" "${GAME_TYPE}" \
        --rpc-url "${ANVIL_URL}" \
        --private-key "${OPERATOR_PRIVATE_KEY}" || echo "Warning: setRespectedGameType failed (may need Guardian role)"

    # Verify the starting anchor root is set (should be set by op-deployer initialization)
    echo "Verifying starting anchor root..."
    ANCHOR_ROOT=$(cast call "${ANCHOR_STATE_REGISTRY}" "getAnchorRoot()(bytes32,uint256)" --rpc-url "${ANVIL_URL}" | head -1)
    ANCHOR_BLOCK=$(cast call "${ANCHOR_STATE_REGISTRY}" "getAnchorRoot()(bytes32,uint256)" --rpc-url "${ANVIL_URL}" | tail -1)
    echo "  Anchor root: ${ANCHOR_ROOT}"
    echo "  Anchor L2 block: ${ANCHOR_BLOCK}"

    if [ "${ANCHOR_ROOT}" = "0x0000000000000000000000000000000000000000000000000000000000000000" ]; then
        echo "ERROR: Anchor root is zero! AnchorStateRegistry was not initialized correctly."
        echo "This will cause 'AnchorRootNotFound' errors when creating games."
    fi

    echo "Op-succinct contracts deployment complete!"
fi

# =====================================================================
# End of op-succinct deployment
# =====================================================================

# Dump anvil state via RPC before killing it
cast rpc anvil_dumpState > "${ANVIL_STATE_FILE}"

# Gracefully shutdown anvil
kill -SIGTERM $ANVIL_PID 2>/dev/null || true

# Wait for clean shutdown
sleep 1

# Force kill if still running
kill -9 $ANVIL_PID 2>/dev/null || true

"${OP_ROOT}/espresso/scripts/reshape-allocs.jq" \
                  <(jq .accounts "${ANVIL_STATE_FILE}") \
                  | jq '{ "alloc": map_values(.state) }' \
                  > "${DEPLOYMENT_DIR}/deployer_allocs.json"

jq -s 'reduce .[] as $item ({}; . * $item)'        \
        <(jq '{ "alloc": map_values(.state) }' "${OP_ROOT}/espresso/environment/allocs.json") \
        "${DEPLOYMENT_DIR}/deployer_allocs.json"            \
        "${OP_ROOT}/espresso/docker/l1-geth/devnet-genesis-template.json" \
        > "${L1_CONFIG_DIR}/genesis.json"
