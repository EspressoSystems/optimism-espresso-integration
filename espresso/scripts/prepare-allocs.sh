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

cast rpc anvil_setBalance "${OPERATOR_ADDRESS}" 0x100000000000000000000000000000000000 --rpc-url "${ANVIL_URL}"
cast rpc anvil_setBalance "${PROPOSER_ADDRESS}" 0x100000000000000000000000000000000000 --rpc-url "${ANVIL_URL}"

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
                      --proof-maturity-delay-seconds=12 \
                      --dispute-game-finality-delay-seconds=6 \
                      --outfile="${DEPLOYER_DIR}/bootstrap_implementations.json"

op-deployer init --l1-chain-id "${L1_CHAIN_ID}" \
                 --l2-chain-ids "${L2_CHAIN_ID}" \
                 --intent-type standard-overrides \
                 --outdir ${DEPLOYER_DIR}

dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].espressoEnabled -t bool -v true

# Configure Espresso batchers for devnet. We reuse the operator address for the
# TEE batcher, but use a separate address for the non-TEE fallback batcher.
# We use Anvil test account #3 for the fallback batcher (already prefunded by Anvil):
# Private key: 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
# Address: 0x90F79bf6EB2c4f870365E785982E1f101E93b906
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].nonTeeBatcher -v "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
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
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].operatorFeeVaultRecipient -v "${OPERATOR_ADDRESS}"
dasel put -f "${DEPLOYER_DIR}/intent.toml" -s .chains.[0].chainFeesRecipient -v "${OPERATOR_ADDRESS}"
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
MAX_CHALLENGE_DURATION=10  # 10 seconds for devnet testing
MAX_PROVE_DURATION=60     # 60 seconds for devnet testing
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

# Deploy op-succinct contracts using local contracts-bedrock deployment script

echo "Deploying op-succinct contracts from local packages/contracts-bedrock..."

# Export environment variables consumed by DeployOPSuccinctFDG.s.sol
export FACTORY_ADDRESS="${DISPUTE_GAME_FACTORY}"
export ANCHOR_STATE_REGISTRY_ADDRESS="${ANCHOR_STATE_REGISTRY}"
export GAME_TYPE="${GAME_TYPE}"
export INITIAL_BOND_WEI="${INITIAL_BOND_WEI}"
export CHALLENGER_BOND_WEI="${CHALLENGER_BOND_WEI}"
export MAX_CHALLENGE_DURATION="${MAX_CHALLENGE_DURATION}"
export MAX_PROVE_DURATION="${MAX_PROVE_DURATION}"
export USE_SP1_MOCK_VERIFIER="true"

pushd "${OP_ROOT}/packages/contracts-bedrock"

echo "Running local DeployOPSuccinctFDG.s.sol script..."
forge script scripts/deploy/DeployOPSuccinctFDG.s.sol \
	--broadcast \
	--no-storage-caching \
	--slow \
	--rpc-url "${ANVIL_URL}" \
	--private-key "${OPERATOR_PRIVATE_KEY}" || echo "Warning: local op-succinct deployment failed, continuing..."

popd

# Wire OPSuccinctFaultDisputeGame implementation into DisputeGameFactory
echo "Configuring DisputeGameFactory for game type ${GAME_TYPE}..."

# Locate the deployed OPSuccinctFaultDisputeGame implementation address from the forge broadcast
BROADCAST_DIR="${OP_ROOT}/packages/contracts-bedrock/broadcast/DeployOPSuccinctFDG.s.sol/${L1_CHAIN_ID}"
BROADCAST_FILE="${BROADCAST_DIR}/run-latest.json"

if [ -f "${BROADCAST_FILE}" ]; then
	FDG_IMPL=$(jq -r '.transactions[] | select(.contractName=="OPSuccinctFaultDisputeGame") | .contractAddress' "${BROADCAST_FILE}")
	if [ -z "${FDG_IMPL}" ] || [ "${FDG_IMPL}" = "null" ]; then
		echo "Warning: could not find OPSuccinctFaultDisputeGame in ${BROADCAST_FILE}; skipping factory config"
	else
		echo "  OPSuccinctFaultDisputeGame implementation: ${FDG_IMPL}"
		# Set implementation for GAME_TYPE on the factory
		cast send "${DISPUTE_GAME_FACTORY}" \
			"setImplementation(uint32,address)" "${GAME_TYPE}" "${FDG_IMPL}" \
			--rpc-url "${ANVIL_URL}" \
			--private-key "${OPERATOR_PRIVATE_KEY}" || echo "Warning: setImplementation failed, continuing..."

		# Set init bond for GAME_TYPE to match challenger bond
		cast send "${DISPUTE_GAME_FACTORY}" \
			"setInitBond(uint32,uint256)" "${GAME_TYPE}" "${CHALLENGER_BOND_WEI}" \
			--rpc-url "${ANVIL_URL}" \
			--private-key "${OPERATOR_PRIVATE_KEY}" || echo "Warning: setInitBond failed, continuing..."
	fi
else
	echo "Warning: broadcast file ${BROADCAST_FILE} not found; skipping factory config"
fi

# Manually activate game type on AnchorStateRegistry (op-succinct script targets OptimismPortal2)
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

echo "Local op-succinct contracts deployment complete!"

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
