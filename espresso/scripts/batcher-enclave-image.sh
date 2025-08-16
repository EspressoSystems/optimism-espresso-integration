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
export ENCLAVE_APP_IMAGE="op-batcher-enclave:app"
export ENCLAVE_TARGET_IMAGE="op-batcher-enclaver:tests"
export MANIFEST_FILE="batcher-enclave.yaml"

# Required for enclave operations
# if [[ ! -e /dev/nitro_enclaves ]]; then
#     echo "Error: /dev/nitro_enclaves device not found. Are you running on a Nitro-enabled instance?"
#     exit 1
# fi

# Check if docker is running
# if ! docker info > /dev/null 2>&1; then
#     echo "Error: Docker is not running or not accessible"
#     exit 1
# fi

echo "Using HOST_IP: $HOST_IP"
echo "Ports -> L1:$L1_HTTP_PORT  L2:$OP_HTTP_PORT  Rollup:$ROLLUP_PORT  EspressoAPI:$ESPRESSO_SEQUENCER_API_PORT"

# Step 1: Build the Docker image using your existing Dockerfile
echo "Building Docker image..."
docker build -t $ENCLAVE_APP_IMAGE \
    -f ../ops/docker/op-stack-go/Dockerfile \
    --target op-batcher-enclave-target \
    --build-arg ENCLAVE_BATCHER_ARGS="--l1-eth-rpc=http://$HOST_IP:$L1_HTTP_PORT \
      --l2-eth-rpc=http://$HOST_IP:$OP_HTTP_PORT \
      --rollup-rpc=http://$HOST_IP:$ROLLUP_PORT \
      --espresso-url=http://$HOST_IP:$ESPRESSO_SEQUENCER_API_PORT,http://$HOST_IP:$ESPRESSO_SEQUENCER_API_PORT \
      --testing-espresso-batcher-private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
      --mnemonic=test\ test\ test\ test\ test\ test\ test\ test\ test\ test\ test\ junk \
      --hd-path=m/44\'/60\'/0\'/0/0 \
      --throttle-threshold=0 --max-channel-duration=1 --target-num-frames=1 \
      --espresso-light-client-addr=0x703848f4c85f18e3acd8196c8ec91eb0b7bd0797" \
    ../

if [ $? -ne 0 ]; then
    echo "Failed to build Docker image"
    exit 1
fi

# Step 2: Create enclaver manifest
echo "Creating enclaver manifest..."
cat > $MANIFEST_FILE << EOL
version: v1
name: "op-batcher-enclave"
target: "$ENCLAVE_TARGET_IMAGE"
sources:
  app: "$ENCLAVE_APP_IMAGE"
defaults:
  memory_mb: 4096
  cpu_count: 2
egress:
  proxy_port: 10000
  allow:
    - "host"
    - "0.0.0.0/0"
    - "**"
    - "::/0"
EOL

echo "Manifest created:"
cat $MANIFEST_FILE

# Step 3: Build the enclave
echo "Building enclave..."
sudo enclaver build --file $MANIFEST_FILE

if [ $? -ne 0 ]; then
    echo "Failed to build enclave"
    exit 1
fi

# Step 4: Run the enclave
# echo "Running enclave..."
# docker run --rm --privileged --net=host \
#   --name batcher-enclaver-$RANDOM \
#   --device=/dev/nitro_enclaves \
#   $ENCLAVE_TARGET_IMAGE

# echo "Enclave execution completed"
