#!/bin/bash
set -euo pipefail

# Parse command line arguments
if [[ $# -ne 4 ]]; then
    echo "Usage: $0 <L1_RPC_URL> <L2_RPC_URL> <ROLLUP_RPC_URL> <ESPRESSO_URL>"
    echo "Example: $0 http://127.0.0.1:8545 http://127.0.0.1:8546 http://127.0.0.1:9545 http://127.0.0.1:24000"
    exit 1
fi

L1_RPC_URL="$1"
L2_RPC_URL="$2"
ROLLUP_RPC_URL="$3"
ESPRESSO_URL="$4"

# Extract HOST_IP from L1_RPC_URL for registration purposes
HOST_IP=$(echo "$L1_RPC_URL" | sed -n 's|^https\?://\([^:]*\).*|\1|p')

# --- load .env ---
SCRIPT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
ENV_FILE="${SCRIPT_DIR}/../.env"
if [[ ! -f "$ENV_FILE" ]]; then
  echo "Error: $ENV_FILE not found"; exit 1
fi
# export everything we source
set -a
# shellcheck disable=SC1090
source "$ENV_FILE"
set +a

# Configuration
TAG="${TAG:-op-batcher-enclavetool}"

echo "Service URLs:"
echo "  L1 RPC: $L1_RPC_URL"
echo "  L2 RPC: $L2_RPC_URL"
echo "  Rollup RPC: $ROLLUP_RPC_URL"
echo "  Espresso API: $ESPRESSO_URL"
echo "  Host IP (for registration): $HOST_IP"

# Build enclave-tools if not already built
if [[ ! -f "/app/op-batcher/bin/enclave-tools" ]]; then
    echo "Building enclave-tools..."
    cd /app/op-batcher
    just enclave-tools
    cd -
fi

# Batcher arguments for both build and run
BATCHER_ARGS="--l1-eth-rpc=$L1_RPC_URL,--l2-eth-rpc=$L2_RPC_URL,--rollup-rpc=$ROLLUP_RPC_URL,--espresso-url=$ESPRESSO_URL,--espresso-url=$ESPRESSO_URL,--testing-espresso-batcher-private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80,--mnemonic=test test test test test test test test test test test junk,--hd-path=m/44'/60'/0'/0/0,--throttle-threshold=0,--max-channel-duration=1,--target-num-frames=1,--espresso-light-client-addr=0x703848f4c85f18e3acd8196c8ec91eb0b7bd0797"

# Use enclave-tools to build the image
echo "Building enclave image using enclave-tools..."
echo "Command: /app/op-batcher/bin/enclave-tools build --op-root /source --tag \"$TAG\" --args \"$BATCHER_ARGS\""
echo "Checking if enclaver is available..."
which enclaver || echo "enclaver not found in PATH"
echo "Checking Docker availability..."
docker version || echo "Docker not accessible"
echo "Starting enclave build..."

# Change to source directory for build context
cd /source

# Run the command and capture output while also showing it in real-time
/app/op-batcher/bin/enclave-tools build \
    --op-root /source \
    --tag "$TAG" \
    --args "$BATCHER_ARGS" 2>&1 | tee /tmp/build_output.log

BUILD_EXIT_CODE=${PIPESTATUS[0]}
BUILD_OUTPUT=$(cat /tmp/build_output.log)

if [ $BUILD_EXIT_CODE -ne 0 ]; then
    echo "Failed to build enclave image (exit code: $BUILD_EXIT_CODE)"
    echo "Build output was:"
    echo "$BUILD_OUTPUT"
    exit 1
fi

echo "Build completed successfully"

# Extract PCR0 from build output
PCR0=$(echo "$BUILD_OUTPUT" | grep "PCR0:" | sed 's/.*PCR0: //')

# Get batch authenticator address from deployment state
BATCH_AUTHENTICATOR_ADDRESS=$(jq -r '.opChainDeployments[0].batchAuthenticatorAddress' /source/espresso/deployment/deployer/state.json)

if [[ -n "$PCR0" && -n "$BATCH_AUTHENTICATOR_ADDRESS" && -n "$OPERATOR_PRIVATE_KEY" ]]; then
    echo "Registering PCR0: $PCR0 with authenticator: $BATCH_AUTHENTICATOR_ADDRESS"
    # Use L1_RPC_URL for registration
    /app/op-batcher/bin/enclave-tools register \
        --authenticator "$BATCH_AUTHENTICATOR_ADDRESS" \
        --l1-url "$L1_RPC_URL" \
        --private-key "$OPERATOR_PRIVATE_KEY" \
        --pcr0 "$PCR0"

    if [ $? -ne 0 ]; then
        echo "Failed to register PCR0, continuing anyway..."
    fi
else
    echo "Skipping registration - missing PCR0 ($PCR0), BATCH_AUTHENTICATOR_ADDRESS ($BATCH_AUTHENTICATOR_ADDRESS), or OPERATOR_PRIVATE_KEY"
fi

# Run the enclave
echo "Running enclave..."
echo "Command: /app/op-batcher/bin/enclave-tools run --image \"$TAG\" --args \"$BATCHER_ARGS\""
/app/op-batcher/bin/enclave-tools run \
    --image "$TAG" \
    --args "$BATCHER_ARGS" &

# Get the enclave-tools PID
ENCLAVE_TOOLS_PID=$!
echo "Enclave-tools started with PID: $ENCLAVE_TOOLS_PID"

# Keep the script running and monitor the enclave
echo "Monitoring enclave..."
while true; do
    # Check if enclave-tools process is still running
    if ! kill -0 $ENCLAVE_TOOLS_PID 2>/dev/null; then
        echo "Enclave-tools process has exited"
        break
    fi

    # Check if any enclave is running
    RUNNING_ENCLAVES=$(sudo nitro-cli describe-enclaves 2>/dev/null | jq length 2>/dev/null || echo "0")
    echo "$(date): Running enclaves: $RUNNING_ENCLAVES"

    sleep 10
done

echo "Script exiting..."
