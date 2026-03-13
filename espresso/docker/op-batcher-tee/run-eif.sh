#!/bin/sh
# REFERENCE TEMPLATE — canonical copy lives in the infra repo
#
# Copy this file to the root of the infra repo alongside build-eif.yml.
# Once copied, edit it there — changes here require a new op-batcher-tee CI
# image, which is unnecessary since run-eif.sh is an outer-layer script that
# does NOT affect PCR0.
#
# Run the pre-built EIF for the op-batcher.
#
# This script is the ENTRYPOINT of the op-batcher-eif image produced by the
# build-eif workflow.  The EIF is baked in at /enclave/application.eif by
# enclaver; this script assembles the batcher arguments from environment
# variables and delivers them to the enclave via enclaver-run.
#
# The outer image is scratch-based (enclaver v0.5.0) so this script must be
# POSIX sh and rely only on busybox (installed alongside it in build-eif.yml).
#
# run-enclave.sh — builds EIF from source, registers PCR0, then runs (local/dev)
# run-eif.sh     — starts enclaver-run against baked-in EIF (production/infra repo)

set -e

# Required environment variables
: ${L1_RPC_URL:?Error: L1_RPC_URL is required}
: ${L2_RPC_URL:?Error: L2_RPC_URL is required}
: ${ROLLUP_RPC_URL:?Error: ROLLUP_RPC_URL is required}
: ${ESPRESSO_URL1:?Error: ESPRESSO_URL1 is required}
: ${OP_BATCHER_PRIVATE_KEY:?Error: OP_BATCHER_PRIVATE_KEY is required}
: ${ESPRESSO_ATTESTATION_SERVICE_URL:?Error: ESPRESSO_ATTESTATION_SERVICE_URL is required}
: ${EIGENDA_PROXY_URL:?Error: EIGENDA_PROXY_URL is required}

# Optional configuration with defaults
ESPRESSO_URL2="${ESPRESSO_URL2:-$ESPRESSO_URL1}"
ESPRESSO_ORIGIN_HEIGHT_ESPRESSO="${ESPRESSO_ORIGIN_HEIGHT_ESPRESSO:-0}"
ESPRESSO_ORIGIN_HEIGHT_L2="${ESPRESSO_ORIGIN_HEIGHT_L2:-0}"
ENCLAVE_DEBUG="${ENCLAVE_DEBUG:-false}"
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
echo "Max Channel Duration: $MAX_CHANNEL_DURATION"
echo "Target Num Frames: $TARGET_NUM_FRAMES"
echo "Max L1 Tx Size Bytes: $MAX_L1_TX_SIZE_BYTES"
echo "AltDA Max Concurrent DA Requests: $ALTDA_MAX_CONCURRENT_DA_REQUESTS"
echo "====================================="

# Send batcher args as a NUL-separated stream.
# Protocol matches enclave-entrypoint.bash: each arg is NUL-terminated;
# a second consecutive NUL (empty string) signals end-of-args.
# NOTE: private key is not logged here — enclave-entrypoint.bash redacts it.
send_batcher_args() {
    printf '%s\0' \
        "--l1-eth-rpc=$L1_RPC_URL" \
        "--l2-eth-rpc=$L2_RPC_URL" \
        "--rollup-rpc=$ROLLUP_RPC_URL" \
        "--espresso.enabled=true" \
        "--espresso.urls=$ESPRESSO_URL1" \
        "--espresso.urls=$ESPRESSO_URL2" \
        "--espresso.espresso-attestation-service=$ESPRESSO_ATTESTATION_SERVICE_URL" \
        "--espresso.origin-height-espresso=$ESPRESSO_ORIGIN_HEIGHT_ESPRESSO" \
        "--espresso.origin-height-l2=$ESPRESSO_ORIGIN_HEIGHT_L2" \
        "--private-key=$OP_BATCHER_PRIVATE_KEY" \
        "--throttle-threshold=0" \
        "--max-channel-duration=$MAX_CHANNEL_DURATION" \
        "--target-num-frames=$TARGET_NUM_FRAMES" \
        "--max-l1-tx-size-bytes=$MAX_L1_TX_SIZE_BYTES" \
        "--max-pending-tx=32" \
        "--espresso.light-client-addr=$ESPRESSO_LIGHT_CLIENT_ADDR" \
        "--espresso.espresso-attestation-service=$ESPRESSO_ATTESTATION_SERVICE_URL" \
        "--altda.enabled=true" \
        "--altda.da-server=$EIGENDA_PROXY_URL" \
        "--altda.da-service=true" \
        "--altda.verify-on-read=false" \
        "--altda.max-concurrent-da-requests=$ALTDA_MAX_CONCURRENT_DA_REQUESTS" \
        "--altda.put-timeout=30s" \
        "--altda.get-timeout=30s" \
        "--data-availability-type=calldata"
    if [ "$ENCLAVE_DEBUG" = "true" ]; then
        printf '%s\0' "--log.level=debug"
        echo "Debug logging enabled" >&2
    fi
    printf '\0'  # double-NUL terminator
}

# ---------------------------------------------------------------------------
# Enclave lifecycle helpers
# ---------------------------------------------------------------------------

# List IDs of all running Nitro enclaves (one per line).
enclave_list_ids() {
    /bin/nitro-cli describe-enclaves 2>&1 | awk -F'"' '/"EnclaveID"/{print $4}'
}

# Terminate all running enclaves by their specific ID.
enclave_terminate_all() {
    for id in $(enclave_list_ids); do
        echo "Terminating enclave: $id"
        /bin/nitro-cli terminate-enclave --enclave-id "$id" 2>/dev/null || true
    done
}

# Terminate our specific enclave by ID (set after startup).
ENCLAVE_ID=""
enclave_shutdown() {
    echo "Received shutdown signal"
    if [ -n "$ENCLAVE_ID" ]; then
        echo "Terminating enclave: $ENCLAVE_ID"
        /bin/nitro-cli terminate-enclave --enclave-id "$ENCLAVE_ID" 2>/dev/null || true
    fi
    kill "$ENCLAVER_PID" 2>/dev/null
    wait "$ENCLAVER_PID" 2>/dev/null
    exit 0
}

trap 'enclave_shutdown' TERM INT

# ---------------------------------------------------------------------------
# Startup: ensure a clean slate before launching our enclave
# ---------------------------------------------------------------------------

# Terminate any stale enclaves left by a previous task.
echo "describe-enclaves output: $(/bin/nitro-cli describe-enclaves 2>&1)"
enclave_terminate_all

# Assert no enclaves are running — guarantees the ID we capture later is ours.
LEFTOVER=$(enclave_list_ids)
if [ -n "$LEFTOVER" ]; then
    echo "ERROR: enclave still running after cleanup: $LEFTOVER"
    exit 1
fi

# Start enclaver-run — reads /enclave/enclaver.yaml, starts the Nitro enclave
# from /enclave/application.eif, and proxies TCP:8337 → vsock:8337.
echo "Starting enclaver-run..."
/usr/local/bin/enclaver-run &
ENCLAVER_PID=$!
echo "enclaver-run started with PID: $ENCLAVER_PID"

# Wait for the ingress port (enclaver-run's vsock bridge) to be ready.
echo "Waiting for enclave ingress port 8337..."
i=0
while [ $i -lt 120 ]; do
    if nc -z 127.0.0.1 8337 2>/dev/null; then
        echo "Enclave ingress port 8337 ready"
        break
    fi
    if ! kill -0 "$ENCLAVER_PID" 2>/dev/null; then
        echo "ERROR: enclaver-run exited prematurely"
        exit 1
    fi
    sleep 1
    i=$((i + 1))
done

if ! nc -z 127.0.0.1 8337 2>/dev/null; then
    echo "ERROR: Enclave ingress port 8337 did not open within 120 seconds"
    exit 1
fi

# Capture the ID of the enclave we just started.
# The pre-start assertion above guarantees this is exactly our enclave.
ENCLAVE_ID=$(enclave_list_ids)
if [ -z "$ENCLAVE_ID" ]; then
    echo "ERROR: enclave not found after startup"
    exit 1
fi
echo "Enclave started with ID: $ENCLAVE_ID"

# Deliver batcher arguments to the enclave's nc listener (args not logged here).
echo "Sending batcher arguments to enclave..."
send_batcher_args | timeout 30 nc 127.0.0.1 8337
echo "Arguments sent to enclave"

# Wait for enclaver-run — it stays alive as long as the enclave is running.
echo "Monitoring enclaver process $ENCLAVER_PID..."
wait "$ENCLAVER_PID"
EXIT_CODE=$?
echo "enclaver-run exited with code: $EXIT_CODE"
exit "$EXIT_CODE"
