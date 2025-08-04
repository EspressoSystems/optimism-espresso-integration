#!/bin/sh
set -euo pipefail

# Set default values.
OP_HTTP_PORT=${OP_HTTP_PORT:-8546}
OP_ENGINE_PORT=${OP_ENGINE_PORT:-8552}
L2_CHAIN_ID=${L2_CHAIN_ID:-22266222}

# Wait for genesis.json to be available.
while [[ ! -f "/config/genesis.json" ]]; do
    echo "Waiting for genesis.json to be generated..."
    sleep 2
done

# Wait for JWT secret to be available.
while [[ ! -f "/config/jwt.txt" ]]; do
    echo "Waiting for JWT secret to be generated..."
    sleep 2
done

# Initialize database if not already done.
if [ ! -d "/data/geth" ]; then
    echo "Initializing OP Geth database..."
    geth --gcmode=archive init --state.scheme=hash --datadir=/data/geth /config/genesis.json
    echo "OP Geth initialization completed"
else
    echo "OP Geth database already initialized, skipping..."
fi

# Start OP Geth with the specified configuration.
echo "Starting OP Geth..."
exec geth \
  --datadir=/data/geth \
  --networkid=${L2_CHAIN_ID} \
  --http \
  --http.addr=0.0.0.0 \
  --http.port=${OP_HTTP_PORT} \
  --http.api=eth,net,web3,debug,admin,txpool \
  --http.vhosts=* \
  --http.corsdomain=* \
  --authrpc.addr=0.0.0.0 \
  --authrpc.port=${OP_ENGINE_PORT} \
  --authrpc.vhosts=* \
  --authrpc.jwtsecret=/config/jwt.txt \
  --rollup.disabletxpoolgossip=true \
  --rollup.halt=major \
  --nodiscover \
  --networkid ${L2_CHAIN_ID}
