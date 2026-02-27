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
  op-deployer inspect genesis --workdir /deployer --outfile /config/genesis.json $L2_CHAIN_ID

  echo "Updating genesis timestamp..."
  # Use environment variable or fallback to the current time.
  GENESIS_TIMESTAMP=${GENESIS_TIMESTAMP:-$(printf '0x%x\n' $(date +%s))}
  dasel put -f /config/genesis.json -s .timestamp -v "$GENESIS_TIMESTAMP"

  if [[ ! -f /config/jwt.txt ]]; then
      echo "Generating JWT token..."
      # TODO (Keyao) Use a random value?
      printf "2692310708e4207ecd73bf5597a59ab9cd085380108a7787b3d6be22840e37f0" > /config/jwt.txt
  fi

  # On fresh/small devnets "finalized" can stay unavailable for a long time.
  # Wait a bounded amount, then fall back to "latest" so genesis does not hang forever.
  max_finalized_wait_seconds=${L1_FINALIZED_WAIT_TIMEOUT_SECONDS:-180}
  finalized_attempts=$((max_finalized_wait_seconds / 3))
  if [ "$finalized_attempts" -lt 1 ]; then
    finalized_attempts=1
  fi

  echo "Waiting for L1 finalized block (up to ${max_finalized_wait_seconds}s)..."
  finalized_block=""
  i=0
  while [ "$i" -lt "$finalized_attempts" ]; do
    finalized_block=$(curl -s -X POST -H "Content-Type: application/json" \
      --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["finalized", false],"id":1}' \
      "$L1_RPC" | jq -r '.result.number')
    if [[ -n "$finalized_block" && "$finalized_block" != "null" ]]; then
      echo "Found L1 finalized block: $finalized_block"
      break
    fi
    i=$((i + 1))
    sleep 3
  done

  if [[ -z "$finalized_block" || "$finalized_block" == "null" ]]; then
    echo "No finalized block found within ${max_finalized_wait_seconds}s; falling back to latest block..."
    latest_block=""
    j=0
    while [ "$j" -lt 60 ]; do
      latest_block=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", false],"id":1}' \
        "$L1_RPC" | jq -r '.result.number')
      if [[ -n "$latest_block" && "$latest_block" != "null" ]]; then
        echo "Found L1 latest block: $latest_block"
        break
      fi
      j=$((j + 1))
      sleep 2
    done

    if [[ -z "$latest_block" || "$latest_block" == "null" ]]; then
      echo "Failed to get any L1 block (finalized/latest) after retries" >&2
      exit 1
    fi
  fi

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
    --rollup.halt=major \
    --nodiscover

elif [ "$MODE" = "rollup" ]; then
  echo "=== Running L2 Rollup Config Mode ==="

  sanitize_rollup_config() {
    # Some op-deployer outputs include fields newer than the op-node parser in this repo.
    # Drop known incompatible keys/sections so op-node can decode rollup.json reliably.
    if [[ -f "/config/rollup.json" ]]; then
      tmp_rollup="$(mktemp)"
      jq 'del(.caff_node_config) | walk(if type == "object" then del(.UseFetchAPI, .useFetchAPI, .OriginHeight, .originHeight) else . end)' /config/rollup.json > "$tmp_rollup" && mv "$tmp_rollup" /config/rollup.json
    fi
  }

  ensure_rollup_eip1559_params() {
    if [[ ! -f "/config/rollup.json" ]]; then
      return
    fi

    current_params="$(jq -r '.genesis.system_config.eip1559Params // empty' /config/rollup.json)"
    if [[ "$current_params" != "0x0000000000000000" ]]; then
      return
    fi

    denom="$(jq -r '.chain_op_config.eip1559DenominatorCanyon // .chain_op_config.eip1559Denominator // empty' /config/rollup.json)"
    elasticity="$(jq -r '.chain_op_config.eip1559Elasticity // empty' /config/rollup.json)"

    if [[ -z "$denom" || -z "$elasticity" ]]; then
      echo "rollup.json has zero eip1559Params and no chain_op_config fallback; leaving unchanged"
      return
    fi

    patched_params="$(printf '0x%08x%08x' "$denom" "$elasticity")"
    echo "Patching rollup eip1559Params from 0x0000000000000000 to ${patched_params}"
    dasel put -f /config/rollup.json -s .genesis.system_config.eip1559Params -t string -v "${patched_params}"
  }

  # Retry RPC until ready (L1 and op-geth-sequencer can be slow to accept connections after healthy).
  retry_rpc() {
    local url="$1"
    local name="$2"
    local max=30
    local out
    while [ "$max" -gt 0 ]; do
      out=$(curl -sf -X POST "${url}" \
        -H 'Content-Type: application/json' \
        -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", false],"id":1}' 2>/dev/null | jq -r ".result.hash" 2>/dev/null) || true
      if [ -n "$out" ] && [ "$out" != "null" ]; then
        echo "$out"
        return 0
      fi
      echo "Waiting for ${name} RPC..."
      sleep 2
      max=$((max - 1))
    done
    echo "Timed out waiting for ${name} RPC at ${url}" >&2
    return 1
  }

  if [[ -f "/deployment/l2-config/rollup.json" ]]; then
    echo "Using pre-built rollup config..."
    cp /deployment/l2-config/rollup.json /config/rollup.json
    sanitize_rollup_config
    ensure_rollup_eip1559_params

    echo "Updating L1 genesis info..."
    L1_HASH=$(retry_rpc "${L1_RPC}" "L1")
    dasel put -f /config/rollup.json -s .genesis.l1.hash -t string -v $L1_HASH
    dasel put -f /config/rollup.json -s .genesis.l1.number -t int -v 0

    echo "Updating L2 genesis info..."
    L2_HASH=$(retry_rpc "${OP_RPC}" "OP sequencer")
    dasel put -f /config/rollup.json -s .genesis.l2.hash -t string -v $L2_HASH
    dasel put -f /config/rollup.json -s .genesis.l2.number -t int -v 0

    echo "Updating rollup l2_time..."
    dasel put -f /config/rollup.json -s .genesis.l2_time -t int -v $(date +%s)
  else
    echo "Pre-built rollup config not found, generating new one..."
    op-deployer inspect rollup --workdir /deployer --outfile /config/rollup.json $L2_CHAIN_ID
    sanitize_rollup_config
    ensure_rollup_eip1559_params

    echo "Updating L1 genesis info..."
    L1_HASH=$(retry_rpc "${L1_RPC}" "L1")
    dasel put -f /config/rollup.json -s .genesis.l1.hash -t string -v $L1_HASH
    dasel put -f /config/rollup.json -s .genesis.l1.number -t int -v 0

    echo "Updating L2 genesis info..."
    L2_HASH=$(retry_rpc "${OP_RPC}" "OP sequencer")
    dasel put -f /config/rollup.json -s .genesis.l2.hash -t string -v $L2_HASH
    dasel put -f /config/rollup.json -s .genesis.l2.number -t int -v 0

    echo "Updating rollup l2_time..."
    dasel put -f /config/rollup.json -s .genesis.l2_time -t int -v $(date +%s)
  fi

  echo "L2 rollup config complete"
  exit 0

else
    echo "Unknown MODE: $MODE. Use 'genesis', 'rollup', or 'geth'"
    exit 1
fi
