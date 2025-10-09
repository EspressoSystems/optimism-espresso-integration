#!/bin/bash
# Enclave Batcher Runner Script
# Supports both local (docker-compose) and AWS ECS deployments

set -e

# Required environment variables - will fail if not set
: ${L1_RPC_URL:?Error: L1_RPC_URL is required}
: ${L2_RPC_URL:?Error: L2_RPC_URL is required}
: ${ROLLUP_RPC_URL:?Error: ROLLUP_RPC_URL is required}
: ${ESPRESSO_URL1:?Error: ESPRESSO_URL1 is required}
: ${OPERATOR_PRIVATE_KEY:?Error: OPERATOR_PRIVATE_KEY is required}

# Optional configuration with defaults
TAG="${TAG:-op-batcher-enclavetool}"
ESPRESSO_URL2="${ESPRESSO_URL2:-$ESPRESSO_URL1}"  # Default to same as URL1 if not set
ENCLAVE_DEBUG="${ENCLAVE_DEBUG:-false}"
MONITOR_INTERVAL="${MONITOR_INTERVAL:-30}"
MEMORY_MB="${ENCLAVE_MEMORY_MB:-4096}"
CPU_COUNT="${ENCLAVE_CPU_COUNT:-2}"

# Deployment mode detection
DEPLOYMENT_MODE="${DEPLOYMENT_MODE:-aws}"  # 'local' or 'aws'

echo "=== Enclave Batcher Configuration ==="
echo "Deployment Mode: $DEPLOYMENT_MODE"
echo "L1 RPC URL: $L1_RPC_URL"
echo "L2 RPC URL: $L2_RPC_URL"
echo "Rollup RPC URL: $ROLLUP_RPC_URL"
echo "Espresso URLs: $ESPRESSO_URL1, $ESPRESSO_URL2"
echo "Debug Mode: $ENCLAVE_DEBUG"
echo "Monitor Interval: $MONITOR_INTERVAL seconds"
echo "Memory: ${MEMORY_MB}MB"
echo "CPU Count: $CPU_COUNT"
echo "====================================="

# Batcher arguments
BATCHER_ARGS="--l1-eth-rpc=$L1_RPC_URL"
BATCHER_ARGS="$BATCHER_ARGS,--l2-eth-rpc=$L2_RPC_URL"
BATCHER_ARGS="$BATCHER_ARGS,--rollup-rpc=$ROLLUP_RPC_URL"
BATCHER_ARGS="$BATCHER_ARGS,--espresso-url=$ESPRESSO_URL1"
BATCHER_ARGS="$BATCHER_ARGS,--espresso-url=$ESPRESSO_URL2"
BATCHER_ARGS="$BATCHER_ARGS,--testing-espresso-batcher-private-key=$OPERATOR_PRIVATE_KEY"
BATCHER_ARGS="$BATCHER_ARGS,--mnemonic=test test test test test test test test test test test junk"
BATCHER_ARGS="$BATCHER_ARGS,--hd-path=m/44'/60'/0'/0/0"
BATCHER_ARGS="$BATCHER_ARGS,--throttle-threshold=0"
BATCHER_ARGS="$BATCHER_ARGS,--max-channel-duration=1"
BATCHER_ARGS="$BATCHER_ARGS,--target-num-frames=1"
BATCHER_ARGS="$BATCHER_ARGS,--espresso-light-client-addr=0x703848f4c85f18e3acd8196c8ec91eb0b7bd0797"

# Add debug arguments if enabled
if [ "$ENCLAVE_DEBUG" = "true" ]; then
    BATCHER_ARGS="$BATCHER_ARGS,--log.level=debug"
    echo "Debug logging enabled"
fi

# Build the enclave image
echo "Building enclave image with tag: $TAG"
cd /source

if ! enclave-tools build --op-root /source --tag "$TAG" 2>&1 | tee /tmp/build_output.log; then
    echo "ERROR: Failed to build enclave image"
    echo "Build output was:"
    cat /tmp/build_output.log
    exit 1
fi

echo "Build completed successfully"

# Extract PCR0 from build output
# Works whether the line is `... PCR0: 0xABCD ...` or `... PCR0=abcd123 ...`
PCR0="$(sed -n -E 's/.*PCR0[:=][[:space:]]*(0[xX])?([[:xdigit:]]+).*/\2/p;q' /tmp/build_output.log)"


# Get batch authenticator address from deployment state
BATCH_AUTHENTICATOR_ADDRESS=$(jq -r '.opChainDeployments[0].batchAuthenticatorAddress' /source/espresso/deployment/deployer/state.json 2>/dev/null || echo "")

# Register PCR0 if all required values are present
if [ -n "$PCR0" ] && [ -n "$BATCH_AUTHENTICATOR_ADDRESS" ] && [ -n "$OPERATOR_PRIVATE_KEY" ]; then
    echo "Registering PCR0: $PCR0 with authenticator: $BATCH_AUTHENTICATOR_ADDRESS"
    enclave-tools register \
        --authenticator "$BATCH_AUTHENTICATOR_ADDRESS" \
        --l1-url "$L1_RPC_URL" \
        --private-key "$OPERATOR_PRIVATE_KEY" \
        --pcr0 "$PCR0"

    if [ $? -ne 0 ]; then
        echo "WARNING: Failed to register PCR0, continuing anyway..."
    else
        echo "PCR0 registration successful"
    fi
else
    echo "Skipping PCR0 registration - missing required values:"
    echo "  PCR0: ${PCR0:-[missing]}"
    echo "  BATCH_AUTHENTICATOR_ADDRESS: ${BATCH_AUTHENTICATOR_ADDRESS:-[missing]}"
    echo "  OPERATOR_PRIVATE_KEY: ${OPERATOR_PRIVATE_KEY:+[set]}"
fi

# Setup tracking files for local deployment
if [ "$DEPLOYMENT_MODE" = "local" ]; then
    PID_FILE="/tmp/enclave-tools.pid"
    CONTAINER_TRACKER_FILE="/tmp/enclave-containers.txt"
    STATUS_FILE="/tmp/enclave-status.json"

    # Cleanup function for local deployment
    cleanup() {
        echo "Cleaning up enclave resources..."
        if [ -f "$PID_FILE" ]; then
            STORED_PID=$(cat "$PID_FILE")
            if kill -0 "$STORED_PID" 2>/dev/null; then
                echo "Terminating enclave-tools process (PID: $STORED_PID)"
                kill -TERM "$STORED_PID" 2>/dev/null || true
                sleep 5
                kill -KILL "$STORED_PID" 2>/dev/null || true
            fi
            rm -f "$PID_FILE"
        fi

        # Clean up any remaining enclave containers
        if [ -f "$CONTAINER_TRACKER_FILE" ]; then
            while IFS= read -r container_id; do
                if [ -n "$container_id" ] && docker ps -q --filter id="$container_id" | grep -q "$container_id"; then
                    echo "Stopping tracked enclave container: $container_id"
                    docker stop "$container_id" 2>/dev/null || true
                    docker rm "$container_id" 2>/dev/null || true
                fi
            done < "$CONTAINER_TRACKER_FILE"
            rm -f "$CONTAINER_TRACKER_FILE"
        fi

        rm -f "$STATUS_FILE"
        exit 0
    }

    # Setup signal handlers for local deployment
    trap cleanup SIGTERM SIGINT EXIT

    # Get Docker network for local deployment
    DOCKER_NETWORK=$(docker network ls --filter name=espresso --format "{{.Name}}" | head -1)
    if [ -z "$DOCKER_NETWORK" ]; then
        DOCKER_NETWORK="espresso_default"
    fi
    echo "Using Docker network: $DOCKER_NETWORK"
    export DOCKER_DEFAULT_NETWORK="$DOCKER_NETWORK"
    export ENCLAVE_DOCKER_NETWORK="$DOCKER_NETWORK"
fi

# Run the enclave
echo "Starting enclave with command:"
echo "  enclave-tools run --image \"$TAG\" --args \"$BATCHER_ARGS\""

enclave-tools run --image "$TAG" --args "$BATCHER_ARGS" &
ENCLAVE_TOOLS_PID=$!

if [ "$DEPLOYMENT_MODE" = "local" ]; then
    echo "$ENCLAVE_TOOLS_PID" > "$PID_FILE"
    echo "Enclave-tools started with PID: $ENCLAVE_TOOLS_PID (stored in $PID_FILE)"
else
    echo "Enclave-tools started with PID: $ENCLAVE_TOOLS_PID"
fi

# Wait for enclave-tools to finish starting the enclave container
echo "Waiting for enclave-tools to complete startup..."
wait $ENCLAVE_TOOLS_PID
ENCLAVE_TOOLS_EXIT_CODE=$?
echo "Enclave-tools process completed with exit code: $ENCLAVE_TOOLS_EXIT_CODE"

if [ "$DEPLOYMENT_MODE" = "local" ]; then
    rm -f "$PID_FILE"
fi

# Check if enclave-tools failed
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

# Get container details
CONTAINER_ID=$(docker ps --filter "name=$CONTAINER_NAME" --format "{{.ID}}" | head -1)
CONTAINER_IMAGE=$(docker inspect "$CONTAINER_NAME" --format '{{.Config.Image}}' 2>/dev/null)
STARTED_AT=$(docker inspect "$CONTAINER_NAME" --format '{{.State.StartedAt}}' 2>/dev/null)

echo "Container Details:"
echo "  ID: $CONTAINER_ID"
echo "  Image: $CONTAINER_IMAGE"
echo "  Started: $STARTED_AT"

# Setup status tracking for local deployment
if [ "$DEPLOYMENT_MODE" = "local" ]; then
    echo "$CONTAINER_NAME" >> "$CONTAINER_TRACKER_FILE"

    # Create initial status file
    cat > "$STATUS_FILE" <<EOF
{
  "container_id": "$CONTAINER_ID",
  "container_name": "$CONTAINER_NAME",
  "container_image": "$CONTAINER_IMAGE",
  "started_at": "$STARTED_AT",
  "last_updated": "$(date -Iseconds)",
  "status": "running",
  "enclave_tools_exit_code": $ENCLAVE_TOOLS_EXIT_CODE
}
EOF
fi

# Start capturing container logs in background
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
    # Check if the container is still running
    CONTAINER_STATUS=$(docker inspect "$CONTAINER_NAME" 2>/dev/null | jq -r '.[0].State.Status' 2>/dev/null || echo "")

    if [ -z "$CONTAINER_STATUS" ] || [ "$CONTAINER_STATUS" != "running" ]; then
        echo "$(date): Container $CONTAINER_NAME is no longer running (status: $CONTAINER_STATUS)"

        # Get exit code if available
        EXIT_CODE=$(docker inspect "$CONTAINER_NAME" 2>/dev/null | jq -r '.[0].State.ExitCode' 2>/dev/null || echo "unknown")
        echo "Container exit code: $EXIT_CODE"

        # Update status file for local deployment
        if [ "$DEPLOYMENT_MODE" = "local" ] && [ -n "$STATUS_FILE" ]; then
            cat > "$STATUS_FILE" <<EOF
{
  "container_id": "$CONTAINER_ID",
  "container_name": "$CONTAINER_NAME",
  "container_image": "$CONTAINER_IMAGE",
  "started_at": "$STARTED_AT",
  "last_updated": "$(date -Iseconds)",
  "status": "exited",
  "exit_code": "$EXIT_CODE",
  "enclave_tools_exit_code": $ENCLAVE_TOOLS_EXIT_CODE
}
EOF
        fi
        break
    fi

    # Log current status periodically
    if [ $(($MONITOR_COUNT % 10)) -eq 0 ]; then
        echo "$(date): Container $CONTAINER_NAME status: $CONTAINER_STATUS"

        # Show container resource usage
        docker stats --no-stream "$CONTAINER_NAME" 2>/dev/null || echo "Could not get container stats"

        # Update status file for local deployment
        if [ "$DEPLOYMENT_MODE" = "local" ] && [ -n "$STATUS_FILE" ]; then
            cat > "$STATUS_FILE" <<EOF
{
  "container_id": "$CONTAINER_ID",
  "container_name": "$CONTAINER_NAME",
  "container_image": "$CONTAINER_IMAGE",
  "started_at": "$STARTED_AT",
  "last_updated": "$(date -Iseconds)",
  "status": "$CONTAINER_STATUS",
  "monitor_count": $MONITOR_COUNT,
  "enclave_tools_exit_code": $ENCLAVE_TOOLS_EXIT_CODE
}
EOF
        fi
    fi

    MONITOR_COUNT=$((MONITOR_COUNT + 1))
    sleep "$MONITOR_INTERVAL"
done

echo "Enclave monitoring ended"

# Clean up log capture if still running
if kill -0 $LOG_PID 2>/dev/null; then
    echo "Stopping log capture..."
    kill $LOG_PID 2>/dev/null || true
fi

echo "Script exiting..."
