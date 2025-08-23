#!/bin/bash
set -euo pipefail

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
# NOTE: if loopback doesn't work from inside the enclave, set HOST_IP=host
HOST_IP="${HOST_IP:-127.0.0.1}"
TAG="${TAG:-op-batcher-enclavetool}"

echo "Using HOST_IP: $HOST_IP"
echo "Ports -> L1:$L1_HTTP_PORT  L2:$OP_HTTP_PORT  Rollup:$ROLLUP_PORT  EspressoAPI:$ESPRESSO_SEQUENCER_API_PORT"

# Build enclave-tools if not already built
if [[ ! -f "/app/op-batcher/bin/enclave-tools" ]]; then
    echo "Building enclave-tools..."
    cd /app/op-batcher
    just enclave-tools
    cd -
fi

# Batcher arguments for both build and run
BATCHER_ARGS="--l1-eth-rpc=http://$HOST_IP:$L1_HTTP_PORT,--l2-eth-rpc=http://$HOST_IP:$OP_HTTP_PORT,--rollup-rpc=http://$HOST_IP:$ROLLUP_PORT,--espresso-url=http://$HOST_IP:$ESPRESSO_SEQUENCER_API_PORT,--espresso-url=http://$HOST_IP:$ESPRESSO_SEQUENCER_API_PORT,--testing-espresso-batcher-private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80,--mnemonic=test test test test test test test test test test test junk,--hd-path=m/44'/60'/0'/0/0,--throttle-threshold=0,--max-channel-duration=1,--target-num-frames=1,--espresso-light-client-addr=0x703848f4c85f18e3acd8196c8ec91eb0b7bd0797"

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
    # Use HOST_IP for network communication
    /app/op-batcher/bin/enclave-tools register \
        --authenticator "$BATCH_AUTHENTICATOR_ADDRESS" \
        --l1-url "http://$HOST_IP:$L1_HTTP_PORT" \
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
    --args "$BATCHER_ARGS"
