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
    --tag "$TAG" )

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

echo "=== Build Complete ==="
echo "Files created:"
echo "  - $SHARED_DIR/pcr0-${TAG}.txt"
echo "  - $SHARED_DIR/enclave-image-${TAG}.tar"
echo ""
echo "Now you can run: docker compose up -d op-batcher-tee"
