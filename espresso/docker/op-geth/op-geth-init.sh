#!/bin/sh
set -euo pipefail

# Set default values.
L1_HTTP_PORT=${L1_HTTP_PORT:-8545}
OP_HTTP_PORT=${OP_HTTP_PORT:-8546}
OP_ENGINE_PORT=${OP_ENGINE_PORT:-8552}
L2_CHAIN_ID=${L2_CHAIN_ID:-22266222}

# Mode can be "genesis", "rollup", or "geth" (default).
MODE=${MODE:-geth}

if [ "$MODE" = "genesis" ]; then
  echo "=== Running L2 Genesis Mode ==="

  echo "Generating genesis..."
  # Use writable cache dir (container runs as U_ID:GID and cannot write to /.op-deployer). Must be before subcommand.
  op-deployer --cache-dir /tmp/op-deployer/cache inspect genesis --workdir /deployer --outfile /config/genesis.json $L2_CHAIN_ID

  echo "Updating genesis timestamp..."
  # Use environment variable or fallback to the current time.
  GENESIS_TIMESTAMP=${GENESIS_TIMESTAMP:-$(printf '0x%x\n' $(date +%s))}
  dasel put -f /config/genesis.json -s .timestamp -v "$GENESIS_TIMESTAMP"

  # Ensure L2 genesis has valid Jovian EIP-1559 extraData so op-geth does not panic in CalcBaseFee.
  # Format: version=1, denominator=250, elasticity=6, minBaseFee=0 (OP Stack Jovian spec).
  JOVIAN_EXTRA_DATA="${JOVIAN_EXTRA_DATA:-0x01000000fa000000060000000000000000}"
  dasel put -f /config/genesis.json -s .extraData -v "$JOVIAN_EXTRA_DATA"

  if [[ ! -f /config/jwt.txt ]]; then
      echo "Generating JWT token..."
      # TODO (Keyao) Use a random value?
      printf "2692310708e4207ecd73bf5597a59ab9cd085380108a7787b3d6be22840e37f0" > /config/jwt.txt
  fi

  # L1 is beacon-based (l1-geth + l1-beacon); "finalized" appears once the beacon has finalized.
  # Fallback to "latest" in case we run before first finality (e.g. startup order change, Jovian timing).
  echo "Waiting for L1 to have a block (finalized, or latest if not yet finalized)..."
  while true; do
    # Try "finalized" first; fall back to "latest" if not available yet. Ignore curl failures so we retry.
    block_num=$(curl -sS -X POST -H "Content-Type: application/json" \
      --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["finalized", false],"id":1}' \
      "$L1_RPC" 2>/dev/null | jq -r '.result.number // empty') || true
    if [[ -z "$block_num" ]]; then
      block_num=$(curl -sS -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", false],"id":1}' \
        "$L1_RPC" 2>/dev/null | jq -r '.result.number // empty') || true
    fi
    if [[ -n "$block_num" && "$block_num" != "null" ]]; then
      echo "Found L1 block, number=$block_num"
      break
    fi
    sleep 3
  done
  echo "L2 genesis setup complete"
  exit 0

elif [ "$MODE" = "geth" ]; then
  echo "=== Starting OP Geth Mode ==="

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
  # rollup.halt=none so geth never exits on protocol-version signal (devnet CI stability).
  echo "Starting OP Geth..."
  exec geth \
    --datadir=/data/geth \
    --gcmode=archive \
    --state.scheme=hash \
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
    --rollup.halt=none \
    --nodiscover

elif [ "$MODE" = "rollup" ]; then
  echo "=== Running L2 Rollup Config Mode ==="

  if [[ -f "/deployment/l2-config/rollup.json" ]]; then
    echo "Using pre-built rollup config..."
    cp /deployment/l2-config/rollup.json /config/rollup.json

    # Still need to update with current L1/L2 state
    echo "Updating L1 genesis info..."
    L1_HASH=$(curl -X POST \
            "${L1_RPC}" \
            -H 'Content-Type: application/json' \
            -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", false],"id":1}' \
            | jq -r ".result.hash")
    dasel put -f /config/rollup.json -s .genesis.l1.hash -t string -v $L1_HASH
    dasel put -f /config/rollup.json -s .genesis.l1.number -t int -v 0

    echo "Updating L2 genesis info..."
    L2_HASH=$(curl -X POST \
            "${OP_RPC}" \
            -H 'Content-Type: application/json' \
            -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", false],"id":1}' \
            | jq -r ".result.hash")
    dasel put -f /config/rollup.json -s .genesis.l2.hash -t string -v $L2_HASH
    dasel put -f /config/rollup.json -s .genesis.l2.number -t int -v 0

    echo "Updating rollup l2_time..."
    dasel put -f /config/rollup.json -s .genesis.l2_time -t int -v $(date +%s)

    # Strip daFootprintGasScalar so succinct-proposer (older schema) can parse the config
    if command -v jq >/dev/null 2>&1; then
      tmp_rollup=$(mktemp)
      jq 'del(.genesis.system_config.daFootprintGasScalar) | del(.chain_op_config.daFootprintGasScalar)' /config/rollup.json > "$tmp_rollup" && mv "$tmp_rollup" /config/rollup.json
    fi
  else
    echo "Pre-built rollup config not found, generating new one..."
    op-deployer --cache-dir /tmp/op-deployer/cache inspect rollup --workdir /deployer --outfile /config/rollup.json $L2_CHAIN_ID

    echo "Updating L1 genesis info..."
    L1_HASH=$(curl -X POST \
            "${L1_RPC}" \
            -H 'Content-Type: application/json' \
            -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", false],"id":1}' \
            | jq -r ".result.hash")
    dasel put -f /config/rollup.json -s .genesis.l1.hash -t string -v $L1_HASH
    dasel put -f /config/rollup.json -s .genesis.l1.number -t int -v 0

    echo "Updating L2 genesis info..."
    L2_HASH=$(curl -X POST \
            "${OP_RPC}" \
            -H 'Content-Type: application/json' \
            -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", false],"id":1}' \
            | jq -r ".result.hash")
    dasel put -f /config/rollup.json -s .genesis.l2.hash -t string -v $L2_HASH
    dasel put -f /config/rollup.json -s .genesis.l2.number -t int -v 0

    echo "Updating rollup l2_time..."
    dasel put -f /config/rollup.json -s .genesis.l2_time -t int -v $(date +%s)

    # Strip daFootprintGasScalar so succinct-proposer (older schema) can parse the config
    if command -v jq >/dev/null 2>&1; then
      tmp_rollup=$(mktemp)
      jq 'del(.genesis.system_config.daFootprintGasScalar) | del(.chain_op_config.daFootprintGasScalar)' /config/rollup.json > "$tmp_rollup" && mv "$tmp_rollup" /config/rollup.json
    fi
  fi

  echo "L2 rollup config complete"
  exit 0

else
    echo "Unknown MODE: $MODE. Use 'genesis', 'rollup', or 'geth'"
    exit 1
fi
