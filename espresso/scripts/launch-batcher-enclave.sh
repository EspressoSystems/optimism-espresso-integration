#!/bin/bash
set -euo pipefail

# Default environment variables if not set
export ENCLAVE_INTERMEDIATE_IMAGE_TAG="op-batcher-enclave:tests"
export ENCLAVE_IMAGE_TAG="op-batcher-enclaver:tests"

# Required for enclave operations
if [[ ! -e /dev/nitro_enclaves ]]; then
    echo "Error: /dev/nitro_enclaves device not found. Are you running on a Nitro-enabled instance?"
    exit 1
fi

# Check if docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Error: Docker is not running or not accessible"
    exit 1
fi

# Step 1: Check and build the intermediate Docker image
# This is the base image that will be used by enclaver to build the enclave image
if ! docker image inspect $ENCLAVE_INTERMEDIATE_IMAGE_TAG >/dev/null 2>&1; then
    echo "Building enclave image..."
    docker build -t $ENCLAVE_INTERMEDIATE_IMAGE_TAG \
        -f ../ops/docker/op-stack-go/Dockerfile \
        --target op-batcher-enclave-target \
        --build-arg ENCLAVE_BATCHER_ARGS="--l1-eth-rpc=http://l1-geth:8545 \
          --l2-eth-rpc=http://l2-geth:8547 \
          --rollup-rpc=http://rollup-node:8548 \
          --espresso-url=http://op-proposer:50051 \
          --testing-espresso-batcher-private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
          --mnemonic=test\ test\ test\ test\ test\ test\ test\ test\ test\ test\ test\ junk \
          --hd-path=m/44\'/60\'/0\'/0/0 \
          --throttle-threshold=0 --max-channel-duration=1 --target-num-frames=1 \
          --espresso-light-client-addr=0x703848f4c85f18e3acd8196c8ec91eb0b7bd0797" \
        ../
    if [ $? -ne 0 ]; then
        echo "Failed to build batcher image"; exit 1
    fi
else
    echo "Using existing intermediate batcher image"
fi

# Create enclaver manifest
cat > batcher-manifest.yaml << EOL
version: v1
name: op-batcher-enclaver
target: $ENCLAVE_IMAGE_TAG
sources:
  app: $ENCLAVE_INTERMEDIATE_IMAGE_TAG
defaults:
  cpu_count: 2
  memory_mb: 4096
egress:
  proxy_port: 10000
  allow:
    - "0.0.0.0/0"
    - "**"
    - "::/0"
EOL

# Step 2: Check and build the final enclave image (op-batcher-enclaver:tests)
# This is built by enclaver using the intermediate image as input
if ! docker image inspect $ENCLAVE_IMAGE_TAG >/dev/null 2>&1; then
    echo "Building enclaver image..."
    echo "Using manifest:"
    cat batcher-manifest.yaml

    echo "\nRunning enclaver build..."
    ENCLAVER_OUTPUT=$(enclaver build --file batcher-manifest.yaml 2>&1)
    if [ $? -ne 0 ]; then
        echo "Failed to build enclaver image"
        echo "Build output:"
        echo "$ENCLAVER_OUTPUT"
        exit 1
    fi
    echo "Build output:"
    echo "$ENCLAVER_OUTPUT"
    # Verify the image was built
    if ! docker image inspect $ENCLAVE_IMAGE_TAG >/dev/null 2>&1; then
        echo "Error: enclaver build succeeded but image not found"
        exit 1
    fi
else
    echo "Using existing enclaver image"
fi

# Run the batcher in enclave
echo "Running batcher in enclave..."
docker run \
    --rm \
    --privileged \
    --net=host \
    --name=batcher-enclaver-${RANDOM} \
    --device=/dev/nitro_enclaves \
    $ENCLAVE_IMAGE_TAG

echo "Batcher started in enclave mode"
