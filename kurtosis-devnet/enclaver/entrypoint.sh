#!/bin/sh
set -e

# -----------------------------------------------------------------------------
# Configuration
# -----------------------------------------------------------------------------
export PROXY_HOST="127.0.0.1"
export PROXY_PORT="10000"
export REMOTE_HOST="172.31.37.114" # get your ip address with $(hostname -I | awk '{print $1}')

# Define the ports for each service.
export REMOTE_PORT_L2_RPC="32786"   # Used for l2-eth-rpc
export REMOTE_PORT_ROLLUP="32789"   # For rollup-rpc (if needed)
export REMOTE_PORT_L1_RPC="32774"   # Used for l1-eth-rpc
export REMOTE_PORT_ESPRESSO="32780" # For espresso-url
export REMOTE_PORT_ALTDA="32781"    # For altda.da-server

# -----------------------------------------------------------------------------
# Start socat proxies in the background
# -----------------------------------------------------------------------------
echo "Starting socat for L1 RPC on port ${REMOTE_PORT_L1_RPC}..."
socat -d TCP4-LISTEN:"${REMOTE_PORT_L1_RPC}",reuseaddr,fork PROXY:"${PROXY_HOST}":"${REMOTE_HOST}":"${REMOTE_PORT_L1_RPC}",proxyport="${PROXY_PORT}" &

echo "Starting socat for L2 RPC on port ${REMOTE_PORT_L2_RPC}..."
socat -d TCP4-LISTEN:"${REMOTE_PORT_L2_RPC}",reuseaddr,fork PROXY:"${PROXY_HOST}":"${REMOTE_HOST}":"${REMOTE_PORT_L2_RPC}",proxyport="${PROXY_PORT}" &

echo "Starting socat for Rollup on port ${REMOTE_PORT_ROLLUP}..."
socat -d TCP4-LISTEN:"${REMOTE_PORT_ROLLUP}",reuseaddr,fork PROXY:"${PROXY_HOST}":"${REMOTE_HOST}":"${REMOTE_PORT_ROLLUP}",proxyport="${PROXY_PORT}" &

echo "Starting socat for Espresso on port ${REMOTE_PORT_ESPRESSO}..."
socat -d TCP4-LISTEN:"${REMOTE_PORT_ESPRESSO}",reuseaddr,fork PROXY:"${PROXY_HOST}":"${REMOTE_HOST}":"${REMOTE_PORT_ESPRESSO}",proxyport="${PROXY_PORT}" &

echo "Starting socat for ALTDA on port ${REMOTE_PORT_ALTDA}..."
socat -d TCP4-LISTEN:"${REMOTE_PORT_ALTDA}",reuseaddr,fork PROXY:"${PROXY_HOST}":"${REMOTE_HOST}":"${REMOTE_PORT_ALTDA}",proxyport="${PROXY_PORT}" &

# Give socat a moment to initialize.
sleep 10

# -----------------------------------------------------------------------------
# Start op-batcher
# -----------------------------------------------------------------------------
echo "Starting op-batcher..."
exec /usr/local/bin/op-batcher \
     --l2-eth-rpc="http://127.0.0.1:${REMOTE_PORT_L2_RPC}" \
     --rollup-rpc="http://127.0.0.1:${REMOTE_PORT_ROLLUP}" \
     --l1-eth-rpc="http://127.0.0.1:${REMOTE_PORT_L1_RPC}" \
     --espresso-url="http://127.0.0.1:${REMOTE_PORT_ESPRESSO}" \
     --altda.da-server="http://127.0.0.1:${REMOTE_PORT_ALTDA}" \
     --poll-interval=1s \
     --sub-safety-margin=6 \
     --num-confirmations=1 \
     --safe-abort-nonce-too-low-count=3 \
     --resubmission-timeout=30s \
     --rpc.addr=0.0.0.0 \
     --rpc.port=8548 \
     --rpc.enable-admin \
     --max-channel-duration=1 \
     --private-key=0xb3d2d558e3491a3709b7c451100a0366b5872520c7aa020c17a0e7fa35b6a8df \
     --data-availability-type=calldata \
     --metrics.enabled \
     --metrics.addr=0.0.0.0 \
     --metrics.port=9001 \
     --altda.enabled=true
