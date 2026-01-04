#!/usr/bin/env bash
set -euxo pipefail

# Deploy op-succinct contracts to Sepolia (or local Anvil if L1_RPC_URL points to localhost)
#
# Usage:
#   export L1_RPC_URL="https://sepolia.infura.io/v3/YOUR_KEY"  # or http://localhost:8545 for Anvil
#   export OPERATOR_PRIVATE_KEY="0x..."
#
#   # After op-deployer apply completes, run:
#   ./espresso/scripts/deploy-op-succinct-sepolia.sh [path-to-state.json]
#
# If state.json path is not provided, defaults to "deployer/state.json"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Try to auto-detect contracts directory
if [ -d "${SCRIPT_DIR}/../../../packages/contracts-bedrock" ]; then
    # We're in the integration repo
    OP_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
    CONTRACTS_DIR="${OP_ROOT}/packages/contracts-bedrock"
elif [ -d "/contracts" ]; then
    # We're in a container with contracts mounted
    CONTRACTS_DIR="/contracts"
elif [ -d "./packages/contracts-bedrock" ]; then
    # We're in repo root
    CONTRACTS_DIR="./packages/contracts-bedrock"
elif [ -n "${CONTRACTS_DIR:-}" ]; then
    # Use provided CONTRACTS_DIR
    :
else
    CONTRACTS_DIR="${SCRIPT_DIR}/../../../packages/contracts-bedrock"
    OP_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd 2>/dev/null || echo "")"
fi

STATE_FILE="${1:-deployer/state.json}"
OUTPUT_FILE="${OUTPUT_FILE:-succinct-contracts.json}"

# Check if deployment already exists (skip redeployment unless FORCE_DEPLOY is set)
DEPLOYMENT_OUTPUT_DIR="$(dirname "${STATE_FILE}")"
DEPLOYMENT_OUTPUT_FILE="${DEPLOYMENT_OUTPUT_DIR}/${OUTPUT_FILE}"
if [ -f "${DEPLOYMENT_OUTPUT_FILE}" ] && [ -z "${FORCE_DEPLOY:-}" ]; then
    echo "=========================================="
    echo "Existing deployment found!"
    echo "=========================================="
    echo "Deployment file: ${DEPLOYMENT_OUTPUT_FILE}"
    echo ""
    echo "Deployed contracts:"
    cat "${DEPLOYMENT_OUTPUT_FILE}" | jq -r '.contracts | to_entries[] | "  \(.key): \(.value)"'
    echo ""
    echo "To redeploy, delete the file or set FORCE_DEPLOY=1"
    echo "To use existing deployment, read addresses from: ${DEPLOYMENT_OUTPUT_FILE}"
    exit 0
fi

# Defaults for Sepolia (can be overridden via environment variables)
SEPOLIA_CHAIN_ID=11155111
GAME_TYPE=42
CHALLENGER_BOND_WEI=1000000000000000  # 0.001 ETH
MAX_CHALLENGE_DURATION="${MAX_CHALLENGE_DURATION:-300}"   # 5 minutes (10s for local Anvil)
MAX_PROVE_DURATION="${MAX_PROVE_DURATION:-1800}"          # 30 minutes (60s for local Anvil)

# Check required environment variables
if [ -z "${L1_RPC_URL:-}" ]; then
    echo "ERROR: L1_RPC_URL environment variable must be set"
    echo "  For Sepolia: export L1_RPC_URL=https://sepolia.infura.io/v3/YOUR_KEY"
    echo "  For local:   export L1_RPC_URL=http://localhost:8545"
    exit 1
fi

if [ -z "${OPERATOR_PRIVATE_KEY:-}" ]; then
    echo "ERROR: OPERATOR_PRIVATE_KEY environment variable must be set"
    exit 1
fi

# Auto-detect chain ID if not set (default to Sepolia)
if [ -z "${L1_CHAIN_ID:-}" ]; then
    if echo "${L1_RPC_URL}" | grep -q "localhost\|127.0.0.1"; then
        # Local Anvil - try to detect or use common devnet chain ID
        L1_CHAIN_ID="${L1_CHAIN_ID:-31337}"
        echo "Detected local RPC, using chain ID: ${L1_CHAIN_ID}"
    else
        # Assume Sepolia for real RPC
        L1_CHAIN_ID="${SEPOLIA_CHAIN_ID}"
        echo "Using Sepolia chain ID: ${L1_CHAIN_ID}"
    fi
fi

# Read contract addresses from state.json
if [ ! -f "${STATE_FILE}" ]; then
    echo "ERROR: State file not found: ${STATE_FILE}"
    echo "Make sure op-deployer apply has completed and state.json exists"
    exit 1
fi

DISPUTE_GAME_FACTORY=$(jq -r '.opChainDeployments[0].DisputeGameFactoryProxy' "${STATE_FILE}")
ANCHOR_STATE_REGISTRY=$(jq -r '.opChainDeployments[0].AnchorStateRegistryProxy' "${STATE_FILE}")

if [ -z "${DISPUTE_GAME_FACTORY}" ] || [ "${DISPUTE_GAME_FACTORY}" = "null" ]; then
    echo "ERROR: Could not read DisputeGameFactoryProxy from ${STATE_FILE}"
    exit 1
fi

if [ -z "${ANCHOR_STATE_REGISTRY}" ] || [ "${ANCHOR_STATE_REGISTRY}" = "null" ]; then
    echo "ERROR: Could not read AnchorStateRegistryProxy from ${STATE_FILE}"
    exit 1
fi

echo "=========================================="
echo "Deploying op-succinct contracts"
echo "=========================================="
echo "RPC URL: ${L1_RPC_URL}"
echo "Chain ID: ${L1_CHAIN_ID}"
echo "State file: ${STATE_FILE}"
echo "DisputeGameFactory: ${DISPUTE_GAME_FACTORY}"
echo "AnchorStateRegistry: ${ANCHOR_STATE_REGISTRY}"
echo "=========================================="
echo ""

# Check contracts directory
if [ ! -d "${CONTRACTS_DIR}" ]; then
    echo "ERROR: Could not find contracts-bedrock directory"
    echo "Tried: ${CONTRACTS_DIR}"
    echo "Please set CONTRACTS_DIR environment variable pointing to contracts-bedrock directory"
    exit 1
fi

# Export environment variables for forge script
export FACTORY_ADDRESS="${DISPUTE_GAME_FACTORY}"
export ANCHOR_STATE_REGISTRY_ADDRESS="${ANCHOR_STATE_REGISTRY}"
export GAME_TYPE="${GAME_TYPE}"
export INITIAL_BOND_WEI="${CHALLENGER_BOND_WEI}"
export CHALLENGER_BOND_WEI="${CHALLENGER_BOND_WEI}"
export MAX_CHALLENGE_DURATION="${MAX_CHALLENGE_DURATION}"
export MAX_PROVE_DURATION="${MAX_PROVE_DURATION}"
export USE_SP1_MOCK_VERIFIER="true"

# Deploy contracts
echo "Deploying contracts with forge..."
pushd "${CONTRACTS_DIR}" > /dev/null

if ! forge script scripts/deploy/DeployOPSuccinctFDG.s.sol \
    --broadcast \
    --rpc-url "${L1_RPC_URL}" \
    --private-key "${OPERATOR_PRIVATE_KEY}" \
    --slow; then
    echo "ERROR: Contract deployment failed"
    popd > /dev/null
    exit 1
fi

popd > /dev/null

# Find deployed contract addresses
BROADCAST_DIR="${CONTRACTS_DIR}/broadcast/DeployOPSuccinctFDG.s.sol/${L1_CHAIN_ID}"
BROADCAST_FILE=$(find "${BROADCAST_DIR}" -name "run-*.json" | sort -V | tail -1)

if [ -z "${BROADCAST_FILE}" ] || [ ! -f "${BROADCAST_FILE}" ]; then
    echo "ERROR: Could not find forge broadcast file in ${BROADCAST_DIR}"
    exit 1
fi

FDG_IMPL=$(jq -r '.transactions[] | select(.contractName=="OPSuccinctFaultDisputeGame") | .contractAddress' "${BROADCAST_FILE}" | head -1)

if [ -z "${FDG_IMPL}" ] || [ "${FDG_IMPL}" = "null" ]; then
    echo "ERROR: Could not find OPSuccinctFaultDisputeGame deployment"
    exit 1
fi

echo ""
echo "Deployed OPSuccinctFaultDisputeGame: ${FDG_IMPL}"
echo ""

# Configure DisputeGameFactory
echo "Configuring DisputeGameFactory..."
if ! cast send "${DISPUTE_GAME_FACTORY}" \
    "setImplementation(uint32,address)" "${GAME_TYPE}" "${FDG_IMPL}" \
    --rpc-url "${L1_RPC_URL}" \
    --private-key "${OPERATOR_PRIVATE_KEY}"; then
    echo "ERROR: setImplementation failed"
    exit 1
fi

if ! cast send "${DISPUTE_GAME_FACTORY}" \
    "setInitBond(uint32,uint256)" "${GAME_TYPE}" "${CHALLENGER_BOND_WEI}" \
    --rpc-url "${L1_RPC_URL}" \
    --private-key "${OPERATOR_PRIVATE_KEY}"; then
    echo "ERROR: setInitBond failed"
    exit 1
fi

# Activate game type on AnchorStateRegistry
echo "Activating game type ${GAME_TYPE} on AnchorStateRegistry..."
if ! cast send "${ANCHOR_STATE_REGISTRY}" \
    "setRespectedGameType(uint32)" "${GAME_TYPE}" \
    --rpc-url "${L1_RPC_URL}" \
    --private-key "${OPERATOR_PRIVATE_KEY}"; then
    echo "ERROR: setRespectedGameType failed"
    exit 1
fi

# Save deployment addresses to a JSON file for easy retrieval
OUTPUT_FILE="${OUTPUT_FILE:-succinct-contracts.json}"
DEPLOYMENT_OUTPUT_DIR="$(dirname "${STATE_FILE}")"
DEPLOYMENT_OUTPUT_FILE="${DEPLOYMENT_OUTPUT_DIR}/${OUTPUT_FILE}"

# Ensure output directory exists
mkdir -p "${DEPLOYMENT_OUTPUT_DIR}"

# Extract AccessManager and SP1Verifier addresses if available
ACCESS_MANAGER=$(jq -r '.transactions[] | select(.contractName=="AccessManager") | .contractAddress' "${BROADCAST_FILE}" | head -1 || echo "")
SP1_VERIFIER=$(jq -r '.transactions[] | select(.contractName=="SP1MockVerifier") | .contractAddress' "${BROADCAST_FILE}" | head -1 || echo "")

# Create output JSON
cat > "${DEPLOYMENT_OUTPUT_FILE}" << EOF
{
  "chainId": ${L1_CHAIN_ID},
  "gameType": ${GAME_TYPE},
  "contracts": {
    "OPSuccinctFaultDisputeGame": "${FDG_IMPL}",
    "AccessManager": "${ACCESS_MANAGER:-}",
    "SP1MockVerifier": "${SP1_VERIFIER:-}"
  },
  "factory": "${DISPUTE_GAME_FACTORY}",
  "anchorStateRegistry": "${ANCHOR_STATE_REGISTRY}",
  "config": {
    "challengerBondWei": "${CHALLENGER_BOND_WEI}",
    "maxChallengeDuration": ${MAX_CHALLENGE_DURATION},
    "maxProveDuration": ${MAX_PROVE_DURATION}
  },
  "deployedAt": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
}
EOF

echo ""
echo "=========================================="
echo "✓ Deployment complete!"
echo "=========================================="
echo "OPSuccinctFaultDisputeGame: ${FDG_IMPL}"
if [ -n "${ACCESS_MANAGER}" ] && [ "${ACCESS_MANAGER}" != "null" ]; then
    echo "AccessManager: ${ACCESS_MANAGER}"
fi
if [ -n "${SP1_VERIFIER}" ] && [ "${SP1_VERIFIER}" != "null" ]; then
    echo "SP1MockVerifier: ${SP1_VERIFIER}"
fi
echo "Game Type: ${GAME_TYPE}"
echo ""
echo "Deployment addresses saved to: ${DEPLOYMENT_OUTPUT_FILE}"
echo ""
echo "To use these addresses, read from the JSON file:"
echo "  jq -r '.contracts.OPSuccinctFaultDisputeGame' ${DEPLOYMENT_OUTPUT_FILE}"
echo ""

