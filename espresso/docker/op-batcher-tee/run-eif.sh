#!/bin/bash
# Run a pre-built Enclaver EIF image for the op-batcher.
#
# This script is the production runtime counterpart to run-enclave.sh.
# It expects the EIF image to have already been built and published by the
# dedicated infra repo CI workflow (see espresso/docs/eif-build-workflow.md),
# and the PCR0 to have already been registered on-chain.
#
# run-enclave.sh  — builds EIF from source, registers PCR0, then runs (local/dev)
# run-eif.sh      — pulls pre-built EIF, assembles args, runs (production/infra repo)

set -e

# Required environment variables
: ${EIF_IMAGE:?Error: EIF_IMAGE is required (e.g. ghcr.io/espressosystems/.../op-batcher-eif:TAG)}
: ${L1_RPC_URL:?Error: L1_RPC_URL is required}
: ${L2_RPC_URL:?Error: L2_RPC_URL is required}
: ${ROLLUP_RPC_URL:?Error: ROLLUP_RPC_URL is required}
: ${ESPRESSO_URL1:?Error: ESPRESSO_URL1 is required}
: ${OPERATOR_PRIVATE_KEY:?Error: OPERATOR_PRIVATE_KEY is required}
: ${ESPRESSO_ATTESTATION_SERVICE_URL:?Error: ESPRESSO_ATTESTATION_SERVICE_URL is required}
: ${EIGENDA_PROXY_URL:?Error: EIGENDA_PROXY_URL is required}

# Optional configuration with defaults
ESPRESSO_URL2="${ESPRESSO_URL2:-$ESPRESSO_URL1}"
ESPRESSO_ORIGIN_HEIGHT_ESPRESSO="${ESPRESSO_ORIGIN_HEIGHT_ESPRESSO:-0}"
ESPRESSO_ORIGIN_HEIGHT_L2="${ESPRESSO_ORIGIN_HEIGHT_L2:-0}"
ENCLAVE_DEBUG="${ENCLAVE_DEBUG:-false}"
MONITOR_INTERVAL="${MONITOR_INTERVAL:-30}"
MAX_CHANNEL_DURATION="${MAX_CHANNEL_DURATION:-2}"
TARGET_NUM_FRAMES="${TARGET_NUM_FRAMES:-1}"
MAX_L1_TX_SIZE_BYTES="${MAX_L1_TX_SIZE_BYTES:-120000}"
ALTDA_MAX_CONCURRENT_DA_REQUESTS="${ALTDA_MAX_CONCURRENT_DA_REQUESTS:-1}"

# Get light client address from env var or use default
if [ -n "$ESPRESSO_LIGHT_CLIENT_ADDR" ]; then
    echo "Using ESPRESSO_LIGHT_CLIENT_ADDR from environment variable"
else
    ESPRESSO_LIGHT_CLIENT_ADDR="0x303872bb82a191771321d4828888920100d0b3e4"
    echo "ESPRESSO_LIGHT_CLIENT_ADDR not set, using default"
fi

# Override OP_BATCHER_ESPRESSO_LIGHT_CLIENT_ADDR so the batcher's env var matches,
# preventing any outer deployment env from leaking a stale value into the enclave.
export OP_BATCHER_ESPRESSO_LIGHT_CLIENT_ADDR="$ESPRESSO_LIGHT_CLIENT_ADDR"

echo "=== Enclave Batcher Configuration ==="
echo "EIF Image: $EIF_IMAGE"
echo "L1 RPC URL: $L1_RPC_URL"
echo "L2 RPC URL: $L2_RPC_URL"
echo "Rollup RPC URL: $ROLLUP_RPC_URL"
echo "Espresso URLs: $ESPRESSO_URL1, $ESPRESSO_URL2"
echo "Attestation service url: $ESPRESSO_ATTESTATION_SERVICE_URL"
echo "EigenDA Proxy URL: $EIGENDA_PROXY_URL"
echo "Light Client Address: $ESPRESSO_LIGHT_CLIENT_ADDR"
echo "Espresso Origin Height: $ESPRESSO_ORIGIN_HEIGHT_ESPRESSO"
echo "L2 Origin Height: $ESPRESSO_ORIGIN_HEIGHT_L2"
echo "Debug Mode: $ENCLAVE_DEBUG"
echo "Monitor Interval: $MONITOR_INTERVAL seconds"
echo "Max Channel Duration: $MAX_CHANNEL_DURATION"
echo "Target Num Frames: $TARGET_NUM_FRAMES"
echo "Max L1 Tx Size Bytes: $MAX_L1_TX_SIZE_BYTES"
echo "AltDA Max Concurrent DA Requests: $ALTDA_MAX_CONCURRENT_DA_REQUESTS"
echo "====================================="

# Batcher arguments
BATCHER_ARGS="--l1-eth-rpc=$L1_RPC_URL"
BATCHER_ARGS="$BATCHER_ARGS,--l2-eth-rpc=$L2_RPC_URL"
BATCHER_ARGS="$BATCHER_ARGS,--rollup-rpc=$ROLLUP_RPC_URL"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.enabled=true"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.urls=$ESPRESSO_URL1"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.urls=$ESPRESSO_URL2"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.espresso-attestation-service=$ESPRESSO_ATTESTATION_SERVICE_URL"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.origin-height-espresso=$ESPRESSO_ORIGIN_HEIGHT_ESPRESSO"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.origin-height-l2=$ESPRESSO_ORIGIN_HEIGHT_L2"

# Use private key if provided, otherwise fall back to test mnemonic
if [ -n "$OP_BATCHER_PRIVATE_KEY" ]; then
    echo "Using OP_BATCHER_PRIVATE_KEY for authentication"
    BATCHER_ARGS="$BATCHER_ARGS,--private-key=$OP_BATCHER_PRIVATE_KEY"
else
    echo "Using test mnemonic for authentication (local development mode)"
    BATCHER_ARGS="$BATCHER_ARGS,--mnemonic=test test test test test test test test test test test junk"
    BATCHER_ARGS="$BATCHER_ARGS,--hd-path=m/44'/60'/0'/0/0"
fi

BATCHER_ARGS="$BATCHER_ARGS,--throttle-threshold=0"
BATCHER_ARGS="$BATCHER_ARGS,--max-channel-duration=$MAX_CHANNEL_DURATION"
BATCHER_ARGS="$BATCHER_ARGS,--target-num-frames=$TARGET_NUM_FRAMES"
BATCHER_ARGS="$BATCHER_ARGS,--max-l1-tx-size-bytes=$MAX_L1_TX_SIZE_BYTES"
BATCHER_ARGS="$BATCHER_ARGS,--max-pending-tx=32"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.light-client-addr=$ESPRESSO_LIGHT_CLIENT_ADDR"
BATCHER_ARGS="$BATCHER_ARGS,--espresso.espresso-attestation-service=$ESPRESSO_ATTESTATION_SERVICE_URL"
BATCHER_ARGS="$BATCHER_ARGS,--altda.enabled=true"
BATCHER_ARGS="$BATCHER_ARGS,--altda.da-server=$EIGENDA_PROXY_URL"
BATCHER_ARGS="$BATCHER_ARGS,--altda.da-service=true"
BATCHER_ARGS="$BATCHER_ARGS,--altda.verify-on-read=false"
BATCHER_ARGS="$BATCHER_ARGS,--altda.max-concurrent-da-requests=$ALTDA_MAX_CONCURRENT_DA_REQUESTS"
BATCHER_ARGS="$BATCHER_ARGS,--altda.put-timeout=30s"
BATCHER_ARGS="$BATCHER_ARGS,--altda.get-timeout=30s"
BATCHER_ARGS="$BATCHER_ARGS,--data-availability-type=calldata"

if [ "$ENCLAVE_DEBUG" = "true" ]; then
    BATCHER_ARGS="$BATCHER_ARGS,--log.level=debug"
    echo "Debug logging enabled"
fi

# Run the pre-built EIF image (args contain sensitive data and are not logged)
echo "Starting enclave with image: $EIF_IMAGE"
enclave-tools run --image "$EIF_IMAGE" --args "$BATCHER_ARGS" &
ENCLAVE_TOOLS_PID=$!
echo "Enclave-tools started with PID: $ENCLAVE_TOOLS_PID"

# Wait for enclave-tools to finish starting the enclave container
echo "Waiting for enclave-tools to complete startup..."
wait $ENCLAVE_TOOLS_PID
ENCLAVE_TOOLS_EXIT_CODE=$?
echo "Enclave-tools process completed with exit code: $ENCLAVE_TOOLS_EXIT_CODE"

if [ $ENCLAVE_TOOLS_EXIT_CODE -ne 0 ]; then
    echo "ERROR: enclave-tools failed with exit code $ENCLAVE_TOOLS_EXIT_CODE"
    exit $ENCLAVE_TOOLS_EXIT_CODE
fi

# Wait for container to fully initialize
sleep 5

# Find the enclave container that was started
echo "Looking for running enclave container..."
CONTAINER_NAME=$(docker ps --format "table {{.Names}}" | grep "batcher-enclaver-" | head -1)

if [ -z "$CONTAINER_NAME" ]; then
    echo "ERROR: No enclave container found after waiting."
    echo "Checking all Docker containers:"
    docker ps -a
    exit 1
fi

echo "Found enclave container: $CONTAINER_NAME"

CONTAINER_ID=$(docker ps --filter "name=$CONTAINER_NAME" --format "{{.ID}}" | head -1)
CONTAINER_IMAGE=$(docker inspect "$CONTAINER_NAME" --format '{{.Config.Image}}' 2>/dev/null)
STARTED_AT=$(docker inspect "$CONTAINER_NAME" --format '{{.State.StartedAt}}' 2>/dev/null)

echo "Container Details:"
echo "  ID: $CONTAINER_ID"
echo "  Image: $CONTAINER_IMAGE"
echo "  Started: $STARTED_AT"

# Capture container logs in background
echo "Starting log capture for container $CONTAINER_NAME"
(
    docker logs -f "$CONTAINER_NAME" 2>&1 | while read line; do
        echo "[ENCLAVE] $line"
    done
) &
LOG_PID=$!
echo "Log capture started with PID: $LOG_PID"

# Monitor the container
echo "Monitoring enclave container $CONTAINER_NAME..."
MONITOR_COUNT=0

while true; do
    CONTAINER_STATUS=$(docker inspect "$CONTAINER_NAME" 2>/dev/null | jq -r '.[0].State.Status' 2>/dev/null || echo "")

    if [ -z "$CONTAINER_STATUS" ] || [ "$CONTAINER_STATUS" != "running" ]; then
        echo "$(date): Container $CONTAINER_NAME is no longer running (status: $CONTAINER_STATUS)"
        EXIT_CODE=$(docker inspect "$CONTAINER_NAME" 2>/dev/null | jq -r '.[0].State.ExitCode' 2>/dev/null || echo "unknown")
        echo "Container exit code: $EXIT_CODE"
        break
    fi

    if [ $(($MONITOR_COUNT % 10)) -eq 0 ]; then
        echo "$(date): Container $CONTAINER_NAME status: $CONTAINER_STATUS"
        docker stats --no-stream "$CONTAINER_NAME" 2>/dev/null || echo "Could not get container stats"
    fi

    MONITOR_COUNT=$((MONITOR_COUNT + 1))
    sleep "$MONITOR_INTERVAL"
done

echo "Enclave monitoring ended"

if kill -0 $LOG_PID 2>/dev/null; then
    echo "Stopping log capture..."
    kill $LOG_PID 2>/dev/null || true
fi

echo "Script exiting..."
