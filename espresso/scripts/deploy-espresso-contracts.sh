#!/usr/bin/env bash
set -euo pipefail

# This script deploys Espresso contracts to a kurtosis devnet `celo-isthmus-devnet`

# Function to get L1 RPC port from kurtosis
get_l1_rpc_port() {
    kurtosis enclave inspect celo-isthmus-devnet | grep -A 10 "el-1-geth-teku" | grep "rpc: " | grep -v "engine-rpc" | sed -n 's/.*127.0.0.1:\([0-9]*\).*/\1/p'
}

# Function to get user key from devnet descriptor
get_user_key() {
    # Clean up any existing directory
    rm -rf devnet-descriptor-0

    # Download the devnet descriptor
    kurtosis files download celo-isthmus-devnet devnet-descriptor-0 >/dev/null 2>&1
    if [ $? -ne 0 ]; then
        echo "Error: Failed to download devnet descriptor" >&2
        return 1
    fi

    # Fix line endings and validate JSON
    dos2unix devnet-descriptor-0/env.json >/dev/null 2>&1
    if ! jq empty devnet-descriptor-0/env.json >/dev/null 2>&1; then
        echo "Error: Invalid JSON file" >&2
        rm -rf devnet-descriptor-0
        return 1
    fi

    # Extract the user key from env.json
    local user_key
    user_key=$(jq -r '.l1.wallets."user-key-12".private_key' devnet-descriptor-0/env.json)
    if [ -z "$user_key" ] || [ "$user_key" = "null" ]; then
        echo "Error: Could not find user key in devnet descriptor" >&2
        rm -rf devnet-descriptor-0
        return 1
    fi

    # Clean up the downloaded files
    rm -rf devnet-descriptor-0

    # Return the user key
    echo "$user_key"
    return 0
}

# Default L1 RPC URL for kurtosis devnet
L1_RPC_PORT=$(get_l1_rpc_port)
if [ -z "$L1_RPC_PORT" ]; then
    echo "Error: Could not get L1 RPC port from kurtosis. Please ensure the devnet is running."
    exit 1
fi
echo "Using L1 RPC Port: $L1_RPC_PORT"
L1_RPC_URL=${L1_RPC_URL:-"http://localhost:$L1_RPC_PORT"}

# Get the project root directory
PROJECT_ROOT=$(realpath "$(dirname "${BASH_SOURCE[0]}")/../..")
ENV_FILE="$PROJECT_ROOT/espresso/.env"

# Load environment variables if .env exists
if [ -f "$ENV_FILE" ]; then
    source "$ENV_FILE"
fi

# Function to get balance in ETH
get_balance() {
    local address=$1
    cast balance --rpc-url "$L1_RPC_URL" "$address" | cast --from-wei
}

# Function to transfer ETH
transfer_eth() {
    local from_key=$1
    local to_address=$2
    local amount=$3
    cast send --rpc-url "$L1_RPC_URL" --private-key "$from_key" "$to_address" --value "${amount}eth"
    echo "✓ Transferred $amount ETH to $to_address"
}

# Check if OPERATOR_PRIVATE_KEY is set in .env
if [ -z "${OPERATOR_PRIVATE_KEY:-}" ]; then
    echo "Error: OPERATOR_PRIVATE_KEY not found in .env file"
    exit 1
fi

# Use the operator's private key from .env
PRIVATE_KEY=$OPERATOR_PRIVATE_KEY

# Get operator's address
OPERATOR_ADDRESS=$(cast wallet address --private-key "$PRIVATE_KEY")
echo "Using operator address: $OPERATOR_ADDRESS"

# Check operator balance and transfer funds if needed
OPERATOR_BALANCE=$(get_balance "$OPERATOR_ADDRESS")
echo "Operator balance: $OPERATOR_BALANCE ETH"

# Convert balance to a number that can be compared
OPERATOR_BALANCE_NUM=$(echo "$OPERATOR_BALANCE" | bc)
if (( $(echo "$OPERATOR_BALANCE_NUM < 5" | bc -l) )); then
    echo "Operator balance is less than 5 ETH. Transferring funds..."
    # Get user key from devnet descriptor
    USER_KEY=$(get_user_key)
    if [ $? -ne 0 ]; then
        echo "Error: Failed to get user key from devnet descriptor" >&2
        exit 1
    fi
    # Remove 0x prefix if present
    USER_KEY=${USER_KEY#0x}
    cast send --private-key "$USER_KEY" --rpc-url "$L1_RPC_URL" "$OPERATOR_ADDRESS" --value 2ether
    if [ $? -ne 0 ]; then
        echo "Error: Failed to transfer funds to operator" >&2
        exit 1
    fi
    echo "Successfully transferred 2 ETH to operator"
fi

# Function to validate contract deployment
validate_contract() {
    local address=$1
    local name=$2
    echo "Validating $name at $address..."
    # Check if address exists on chain
    if ! cast code "$address" --rpc-url "$L1_RPC_URL" > /dev/null 2>&1; then
        echo "Error: $name contract not found at $address"
        exit 1
    fi
    echo "✓ $name contract validated"
}

# Function to update .env file
update_env() {
    local name=$1
    local address=$2
    # Remove existing line if it exists
    sed -i "/$name=/d" "$ENV_FILE"
    # Add new line
    echo "$name=$address" >> "$ENV_FILE"
    echo "✓ Updated $name in .env"
}

# Ensure we're in the contracts directory
cd "$PROJECT_ROOT/packages/contracts-bedrock"

# First, deploy the EspressoTEEVerifier
echo "Deploying EspressoTEEVerifier..."
TEE_VERIFIER_RESULT=$(forge create \
    --rpc-url "$L1_RPC_URL" \
    --private-key "$PRIVATE_KEY" \
    --broadcast \
    lib/espresso-tee-contracts/src/EspressoTEEVerifier.sol:EspressoTEEVerifier \
    --constructor-args "$ESPRESSO_SEQUENCER_PLONK_VERIFIER_ADDRESS" "$ESPRESSO_SEQUENCER_PLONK_VERIFIER_V2_ADDRESS")

# Extract EspressoTEEVerifier address
TEE_VERIFIER_ADDRESS=$(echo "$TEE_VERIFIER_RESULT" | grep "Deployed to:" | awk '{print $3}')
echo "EspressoTEEVerifier deployed to: $TEE_VERIFIER_ADDRESS"

# Validate EspressoTEEVerifier deployment
validate_contract "$TEE_VERIFIER_ADDRESS" "EspressoTEEVerifier"

# Deploy BatchAuthenticator with TEE verifier
echo "Deploying BatchAuthenticator..."
AUTHENTICATOR_RESULT=$(forge create \
    --rpc-url "$L1_RPC_URL" \
    --private-key "$PRIVATE_KEY" \
    --broadcast \
    src/L1/BatchAuthenticator.sol:BatchAuthenticator \
    --constructor-args "$TEE_VERIFIER_ADDRESS" "$OPERATOR_ADDRESS")

# Extract BatchAuthenticator address
AUTHENTICATOR_ADDRESS=$(echo "$AUTHENTICATOR_RESULT" | grep "Deployed to:" | awk '{print $3}')
echo "BatchAuthenticator deployed to: $AUTHENTICATOR_ADDRESS"

# Validate BatchAuthenticator deployment
validate_contract "$AUTHENTICATOR_ADDRESS" "BatchAuthenticator"

# Deploy BatchInbox with BatchAuthenticator address
echo "Deploying BatchInbox..."
INBOX_RESULT=$(forge create \
    --rpc-url "$L1_RPC_URL" \
    --private-key "$PRIVATE_KEY" \
    --broadcast \
    src/L1/BatchInbox.sol:BatchInbox \
    --constructor-args "$AUTHENTICATOR_ADDRESS")

# Extract BatchInbox address
INBOX_ADDRESS=$(echo "$INBOX_RESULT" | grep "Deployed to:" | awk '{print $3}')
echo "BatchInbox deployed to: $INBOX_ADDRESS"

# Validate BatchInbox deployment
validate_contract "$INBOX_ADDRESS" "BatchInbox"

# Update .env file with new contract addresses
echo "Updating .env file with contract addresses..."
update_env "ESPRESSO_TEE_VERIFIER_ADDRESS" "$TEE_VERIFIER_ADDRESS"
update_env "BATCH_AUTHENTICATOR_ADDRESS" "$AUTHENTICATOR_ADDRESS"
update_env "BATCH_INBOX_ADDRESS" "$INBOX_ADDRESS"

echo "Deployment complete! Contract addresses have been added to $ENV_FILE"

# Output addresses in a format suitable for environment variables
echo ""
echo "Added these to your .env file:"
echo "ESPRESSO_TEE_VERIFIER_ADDRESS=$TEE_VERIFIER_ADDRESS"
echo "BATCH_AUTHENTICATOR_ADDRESS=$AUTHENTICATOR_ADDRESS"
echo "BATCH_INBOX_ADDRESS=$INBOX_ADDRESS"
