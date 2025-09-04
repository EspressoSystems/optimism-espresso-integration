#!/bin/bash
# Script to build enclave image outside Docker and save it for consistent PCR0

set -e

# Load environment from .env
SCRIPT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
ENV_FILE="${SCRIPT_DIR}/../.env"
if [[ ! -f "$ENV_FILE" ]]; then
  echo "Error: $ENV_FILE not found"; exit 1
fi
set -a
source "$ENV_FILE"
set +a

# Configuration
TAG="${TAG:-op-batcher-enclavetool}"
SHARED_DIR="${SCRIPT_DIR}/../shared"

echo "=== Enclave Image Builder ==="
echo "Tag: $TAG"
echo "Shared directory: $SHARED_DIR"

# Create shared directory
mkdir -p "$SHARED_DIR"

# Ensure enclave-tools is built
if [[ ! -f "../op-batcher/bin/enclave-tools" ]]; then
    echo "Building enclave-tools..."
    cd ../op-batcher
    just enclave-tools
    cd -
fi

echo "Building enclave image WITHOUT args for consistent PCR0..."
BUILD_OUTPUT=$(../op-batcher/bin/enclave-tools build \
    --op-root ../ \
    --tag "$TAG" 2>&1)

if [ $? -ne 0 ]; then
    echo "Failed to build enclave image"
    echo "$BUILD_OUTPUT"
    exit 1
fi

echo "$BUILD_OUTPUT"

# Extract PCR0
PCR0=$(echo "$BUILD_OUTPUT" | grep "PCR0:" | sed 's/.*PCR0: //')
echo "Extracted PCR0: $PCR0"

# Save PCR0 to shared file
echo "$PCR0" > "$SHARED_DIR/pcr0-${TAG}.txt"
echo "Saved PCR0 to $SHARED_DIR/pcr0-${TAG}.txt"

# Save the enclave image
echo "Saving enclave image..."
docker save "$TAG" -o "$SHARED_DIR/enclave-image-${TAG}.tar"
echo "Saved image to $SHARED_DIR/enclave-image-${TAG}.tar"

# Now register the PCR0
HOST_IP="${HOST_IP:-127.0.0.1}"
BATCH_AUTHENTICATOR_ADDRESS=$(jq -r '.opChainDeployments[0].batchAuthenticatorAddress' deployment/deployer/state.json)

if [[ -n "$PCR0" && -n "$BATCH_AUTHENTICATOR_ADDRESS" && -n "$OPERATOR_PRIVATE_KEY" ]]; then
    echo "=== Registering PCR0 ==="
    echo "PCR0: $PCR0"
    echo "Authenticator: $BATCH_AUTHENTICATOR_ADDRESS"
    echo "L1 URL: http://$HOST_IP:$L1_HTTP_PORT"
    
    ../op-batcher/bin/enclave-tools register \
        --authenticator "$BATCH_AUTHENTICATOR_ADDRESS" \
        --l1-url "http://$HOST_IP:$L1_HTTP_PORT" \
        --private-key "$OPERATOR_PRIVATE_KEY" \
        --pcr0 "$PCR0"
    
    if [ $? -ne 0 ]; then
        echo "ERROR: Failed to register PCR0"
        exit 1
    fi
    echo "PCR0 registered successfully!"
else
    echo "ERROR: Missing required values for registration"
    echo "  PCR0: ${PCR0:-[missing]}"
    echo "  BATCH_AUTHENTICATOR_ADDRESS: ${BATCH_AUTHENTICATOR_ADDRESS:-[missing]}"
    echo "  OPERATOR_PRIVATE_KEY: ${OPERATOR_PRIVATE_KEY:+[set]}"
    exit 1
fi

echo "=== Build Complete ==="
echo "Files created:"
echo "  - $SHARED_DIR/pcr0-${TAG}.txt"
echo "  - $SHARED_DIR/enclave-image-${TAG}.tar"
echo ""
echo "Now you can run: docker compose up -d op-batcher-tee"