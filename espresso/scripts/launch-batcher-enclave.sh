#!/bin/bash
set -euo pipefail

# Default environment variables if not set
export ESPRESSO_RUN_ENCLAVE_TESTS=1

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

# Step 1: Check and build the intermediate Docker image (op-batcher-enclave:tests)
# This is the base image that will be used by enclaver to build the enclave image
if ! docker image inspect op-batcher-enclave:tests >/dev/null 2>&1; then
    echo "Building intermediate batcher image..."
    docker build -t op-batcher-enclave:tests \
        -f docker/op-stack/Dockerfile \
        --target op-batcher-target \
        ../ || { echo "Failed to build batcher image"; exit 1; }
else
    echo "Using existing intermediate batcher image"
fi

# Create enclaver manifest
cat > batcher-manifest.yaml << EOL
version: v1
name: op-batcher-use-enclaver
target: op-batcher-use-enclaver:tests
sources:
  app: op-batcher-enclave:tests
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

# Step 2: Check and build the final enclave image (op-batcher-use-enclaver:tests)
# This is built by enclaver using the intermediate image as input
if ! docker image inspect op-batcher-use-enclaver:tests >/dev/null 2>&1; then
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
    if ! docker image inspect op-batcher-use-enclaver:tests >/dev/null 2>&1; then
        echo "Error: enclaver build succeeded but image not found"
        exit 1
    fi
else
    echo "Using existing enclaver image"
fi

# Check if docker is running and containers exist
if ! docker info >/dev/null 2>&1; then
    echo "Error: Docker is not running or not accessible"
    echo "Please start Docker and ensure you have the necessary permissions"
    exit 1
fi

echo "This script requires the following containers to be running:"
echo "- l1-geth"
echo "- op-geth"
echo "- op-node-sequencer"
echo "- espresso-dev-node"
echo "
To start all required containers, run:"
echo "cd ../.. && docker compose up -d"
echo ""

echo "Checking for required containers..."

# Get container IDs and check they exist
L1_CONTAINER=$(docker ps -q --filter name=espresso-l1-geth)
OP_CONTAINER=$(docker ps -q --filter name=espresso-op-geth)
ROLLUP_CONTAINER=$(docker ps -q --filter name=espresso-op-node-sequencer)
ESPRESSO_CONTAINER=$(docker ps -q --filter name=espresso-espresso-dev-node)

# Show running containers for debugging
echo "\nCurrently running containers:"
docker ps --format "table {{.Names}}\t{{.Status}}"

echo "\nChecking required containers:"

# Check all required containers are running
if [ -z "$L1_CONTAINER" ]; then
    echo "❌ espresso-l1-geth container not found"
    echo "Please start all containers using 'docker compose up -d'"
    exit 1
else
    echo "✓ espresso-l1-geth is running"
fi

if [ -z "$OP_CONTAINER" ]; then
    echo "❌ espresso-op-geth container not found"
    echo "Please start all containers using 'docker compose up -d'"
    exit 1
else
    echo "✓ espresso-op-geth is running"
fi

if [ -z "$ROLLUP_CONTAINER" ]; then
    echo "❌ espresso-op-node-sequencer container not found"
    echo "Please start all containers using 'docker compose up -d'"
    exit 1
else
    echo "✓ espresso-op-node-sequencer is running"
fi

if [ -z "$ESPRESSO_CONTAINER" ]; then
    echo "❌ espresso-espresso-dev-node container not found"
    echo "Please start all containers using 'docker compose up -d'"
    exit 1
else
    echo "✓ espresso-espresso-dev-node is running"
fi

# Set default ports as used in docker-compose.yml
L1_HTTP_PORT=8545
OP_HTTP_PORT=8547
ROLLUP_PORT=8548
ESPRESSO_SEQUENCER_API_PORT=50051

echo "Using ports:"
echo "L1 HTTP: $L1_HTTP_PORT"
echo "OP HTTP: $OP_HTTP_PORT"
echo "Rollup: $ROLLUP_PORT"
echo "Espresso Sequencer: $ESPRESSO_SEQUENCER_API_PORT"

# Run the batcher in enclave
echo "Running batcher in enclave..."
docker run \
    --rm \
    --privileged \
    --net=host \
    --name=batcher-enclaver-${RANDOM} \
    --device=/dev/nitro_enclaves \
    op-batcher-use-enclaver:tests

echo "Batcher started in enclave mode"
