#!/bin/bash
set -e

# Set the default port if not provided.
ESPRESSO_L1_PORT=${ESPRESSO_L1_PORT:-8545}

# Initialize database.
echo "Initializing L1 Geth database..."∂
rm -rf /data/geth || true
geth --datadir /data init /l1-genesis-devnet.json
echo "L1 Geth initialization completed"

# Start Geth with the specified configuration.
exec geth --datadir /data \
  --http \
  --http.addr=0.0.0.0 \
  --http.api=eth,net,web3,admin \
  --http.port=${ESPRESSO_L1_PORT} \
  --http.vhosts=* \
  --http.corsdomain=* \
  --nodiscover \
  --dev \
  --dev.period=12 \
  --miner.etherbase=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC \
  --mine \
  --allow-insecure-unlock \
  --rpc.allow-unprotected-txs
