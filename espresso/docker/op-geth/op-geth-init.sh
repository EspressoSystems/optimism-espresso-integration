#!/bin/sh
set -euo pipefail

# Set default values.
L1_HTTP_PORT=${L1_HTTP_PORT:-8545}
OP_HTTP_PORT=${OP_HTTP_PORT:-8546}
OP_ENGINE_PORT=${OP_ENGINE_PORT:-8552}
L2_CHAIN_ID=${L2_CHAIN_ID:-22266222}

# Mode can be "genesis", "rollup", or "geth" (default).
MODE=${MODE:-geth}
# This op-geth build in the devnet path is unstable with Jovian-at-genesis
# (extraData/minBaseFee handling mismatch). Keep Jovian disabled by default.
# Keep Cancun/Prague delayed together to preserve valid fork ordering and use
# pre-Cancun Engine API flow in this devnet profile.
L2_CANCUN_TIME=${L2_CANCUN_TIME:-4102444800}
L2_PRAGUE_TIME=${L2_PRAGUE_TIME:-4102444800}
L2_ECOTONE_TIME=${L2_ECOTONE_TIME:-4102444800}
L2_FJORD_TIME=${L2_FJORD_TIME:-4102444800}
L2_GRANITE_TIME=${L2_GRANITE_TIME:-4102444800}
L2_HOLOCENE_TIME=${L2_HOLOCENE_TIME:-4102444800}
L2_ISTHMUS_TIME=${L2_ISTHMUS_TIME:-4102444800}
L2_JOVIAN_TIME=${L2_JOVIAN_TIME:-4102444800}

ensure_genesis_eip1559_extradata() {
  local genesis_path="${1:-/config/genesis.json}"
  local in_place="${2:-false}"
  if [[ ! -f "$genesis_path" ]]; then
    return
  fi

  GENESIS_FILE_FOR_GETH="$genesis_path"
  denom="$(jq -r '.config.optimism.eip1559DenominatorCanyon // .config.optimism.eip1559Denominator // empty' "$genesis_path")"
  elasticity="$(jq -r '.config.optimism.eip1559Elasticity // empty' "$genesis_path")"
  current_extra_data="$(jq -r '.extraData // empty' "$genesis_path")"

  if [[ -z "$denom" || -z "$elasticity" || -z "$current_extra_data" ]]; then
    return
  fi

  # With Jovian active at genesis, op-geth expects 17-byte extraData:
  # version(1 byte) + denominator(4 bytes) + elasticity(4 bytes) + minBaseFee(8 bytes).
  # Otherwise it expects legacy 8-byte encoding.
  jovian_time="$(jq -r '.config.jovianTime // .config.jovian_time // .config.optimism.jovianTime // .config.optimism.jovian_time // empty' "$genesis_path")"
  if [ "$jovian_time" = "0" ]; then
    expected_extra_data="$(printf '0x01%08x%08x%016x' "$denom" "$elasticity" 0)"
  else
    expected_extra_data="$(printf '0x%08x%08x' "$denom" "$elasticity")"
  fi
  if [[ "$current_extra_data" == "$expected_extra_data" ]]; then
    return
  fi

  echo "Normalizing genesis extraData for EIP-1559 compatibility: ${current_extra_data} -> ${expected_extra_data}"
  if [[ "$in_place" == "true" ]]; then
    dasel put -f "$genesis_path" -s .extraData -t string -v "${expected_extra_data}"
    GENESIS_FILE_FOR_GETH="$genesis_path"
  elif [[ "$genesis_path" == "/config/genesis.json" ]]; then
    tmp_genesis="/tmp/geth-genesis.json"
    cp "$genesis_path" "$tmp_genesis"
    dasel put -f "$tmp_genesis" -s .extraData -t string -v "${expected_extra_data}"
    GENESIS_FILE_FOR_GETH="$tmp_genesis"
  else
    dasel put -f "$genesis_path" -s .extraData -t string -v "${expected_extra_data}"
    GENESIS_FILE_FOR_GETH="$genesis_path"
  fi
}

if [ "$MODE" = "genesis" ]; then
  echo "=== Running L2 Genesis Mode ==="

  echo "Generating genesis..."
  op-deployer inspect genesis --workdir /deployer --outfile /config/genesis.json $L2_CHAIN_ID

  echo "Updating genesis timestamp..."
  # Use environment variable or fallback to the current time.
  GENESIS_TIMESTAMP=${GENESIS_TIMESTAMP:-$(printf '0x%x\n' $(date +%s))}
  dasel put -f /config/genesis.json -s .timestamp -v "$GENESIS_TIMESTAMP"
  # Keep Cancun+OP hardforks disabled unless explicitly overridden.
  # op-geth chain config uses .config.<fork>Time, while some generators also
  # emit .config.optimism.<fork>Time. Force both forms to stay consistent.
  dasel put -f /config/genesis.json -s .config.cancunTime -t int -v "${L2_CANCUN_TIME}"
  dasel put -f /config/genesis.json -s .config.pragueTime -t int -v "${L2_PRAGUE_TIME}"
  dasel put -f /config/genesis.json -s .config.ecotoneTime -t int -v "${L2_ECOTONE_TIME}"
  dasel put -f /config/genesis.json -s .config.optimism.ecotoneTime -t int -v "${L2_ECOTONE_TIME}"
  dasel put -f /config/genesis.json -s .config.fjordTime -t int -v "${L2_FJORD_TIME}"
  dasel put -f /config/genesis.json -s .config.optimism.fjordTime -t int -v "${L2_FJORD_TIME}"
  dasel put -f /config/genesis.json -s .config.graniteTime -t int -v "${L2_GRANITE_TIME}"
  dasel put -f /config/genesis.json -s .config.optimism.graniteTime -t int -v "${L2_GRANITE_TIME}"
  dasel put -f /config/genesis.json -s .config.holoceneTime -t int -v "${L2_HOLOCENE_TIME}"
  dasel put -f /config/genesis.json -s .config.optimism.holoceneTime -t int -v "${L2_HOLOCENE_TIME}"
  dasel put -f /config/genesis.json -s .config.isthmusTime -t int -v "${L2_ISTHMUS_TIME}"
  dasel put -f /config/genesis.json -s .config.optimism.isthmusTime -t int -v "${L2_ISTHMUS_TIME}"
  dasel put -f /config/genesis.json -s .config.jovianTime -t int -v "${L2_JOVIAN_TIME}"
  dasel put -f /config/genesis.json -s .config.optimism.jovianTime -t int -v "${L2_JOVIAN_TIME}"
  # Ensure generated genesis uses the eip1559 extraData layout expected by op-geth.
  ensure_genesis_eip1559_extradata /config/genesis.json true

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

  # Some genesis generators emit versioned extraData that this op-geth build
  # rejects. Rewrite to plain eip1559 params expected by CalcBaseFee.
  ensure_genesis_eip1559_extradata

  # Always initialize against the current genesis.
  # If an incompatible/stale chain database is present (e.g. chain ID 1),
  # reset the datadir and re-init so op-node sees the expected L2 chain ID.
  echo "Initializing OP Geth database..."
  if ! geth --gcmode=archive init --state.scheme=hash --datadir=/data/geth "${GENESIS_FILE_FOR_GETH:-/config/genesis.json}"; then
      echo "Initial geth init failed; resetting /data/geth and retrying..."
      rm -rf /data/geth/geth /data/geth/keystore /data/geth/geth.ipc /data/geth/LOCK
      mkdir -p /data/geth
      geth --gcmode=archive init --state.scheme=hash --datadir=/data/geth "${GENESIS_FILE_FOR_GETH:-/config/genesis.json}"
  fi
  echo "OP Geth initialization completed"

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
    # Strip daFootprintGasScalar so succinct-proposer (older schema) can parse the config;
    # op-node treats missing value as 0 and L1BlockInfo uses DAFootprintGasScalarDefault (400).
    if [[ -f "/config/rollup.json" ]]; then
      tmp_rollup="$(mktemp)"
      jq 'del(.caff_node_config) | del(.genesis.system_config.daFootprintGasScalar) | del(.chain_op_config.daFootprintGasScalar) | walk(if type == "object" then del(.UseFetchAPI, .useFetchAPI, .OriginHeight, .originHeight) else . end)' /config/rollup.json > "$tmp_rollup" && mv "$tmp_rollup" /config/rollup.json
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

  force_safe_rollup_system_scalar() {
    if [[ ! -f "/config/rollup.json" ]]; then
      return
    fi
    # Keep the scalar in strict Bedrock encoding (version 0 with zero padding).
    # This avoids malformed scalar padding triggering extreme L1-cost math paths
    # in op-geth txpool during devnet smoke tests.
    safe_scalar="0x00000000000000000000000000000000000000000000000000000000000f4240"
    dasel put -f /config/rollup.json -s .genesis.system_config.scalar -t string -v "${safe_scalar}"
  }

  force_rollup_fork_times_from_env() {
    if [[ ! -f "/config/rollup.json" ]]; then
      return
    fi
    # Keep rollup hardfork schedule deterministic for devnet:
    # disable Cancun+OP forks at genesis unless explicitly overridden.
    dasel put -f /config/rollup.json -s .ecotone_time -t int -v "${L2_ECOTONE_TIME}"
    dasel put -f /config/rollup.json -s .fjord_time -t int -v "${L2_FJORD_TIME}"
    dasel put -f /config/rollup.json -s .granite_time -t int -v "${L2_GRANITE_TIME}"
    dasel put -f /config/rollup.json -s .holocene_time -t int -v "${L2_HOLOCENE_TIME}"
    dasel put -f /config/rollup.json -s .isthmus_time -t int -v "${L2_ISTHMUS_TIME}"
    dasel put -f /config/rollup.json -s .jovian_time -t int -v "${L2_JOVIAN_TIME}"
  }

  sync_rollup_fork_times_from_genesis() {
    if [[ ! -f "/config/rollup.json" || ! -f "/config/genesis.json" ]]; then
      return
    fi

    rollup_ecotone="$(jq -r '.ecotone_time // empty' /config/rollup.json)"
    rollup_fjord="$(jq -r '.fjord_time // empty' /config/rollup.json)"
    rollup_granite="$(jq -r '.granite_time // empty' /config/rollup.json)"
    rollup_holocene="$(jq -r '.holocene_time // empty' /config/rollup.json)"
    rollup_isthmus="$(jq -r '.isthmus_time // empty' /config/rollup.json)"
    rollup_jovian="$(jq -r '.jovian_time // empty' /config/rollup.json)"
    genesis_ecotone="$(jq -r '.config.ecotoneTime // .config.ecotone_time // .config.optimism.ecotoneTime // .config.optimism.ecotone_time // empty' /config/genesis.json)"
    genesis_fjord="$(jq -r '.config.fjordTime // .config.fjord_time // .config.optimism.fjordTime // .config.optimism.fjord_time // empty' /config/genesis.json)"
    genesis_granite="$(jq -r '.config.graniteTime // .config.granite_time // .config.optimism.graniteTime // .config.optimism.granite_time // empty' /config/genesis.json)"
    genesis_holocene="$(jq -r '.config.holoceneTime // .config.holocene_time // .config.optimism.holoceneTime // .config.optimism.holocene_time // empty' /config/genesis.json)"
    genesis_isthmus="$(jq -r '.config.isthmusTime // .config.isthmus_time // .config.optimism.isthmusTime // .config.optimism.isthmus_time // empty' /config/genesis.json)"
    genesis_jovian="$(jq -r '.config.jovianTime // .config.jovian_time // .config.optimism.jovianTime // .config.optimism.jovian_time // empty' /config/genesis.json)"

    # Keep rollup fork timing aligned with genesis config so op-node and op-geth agree.
    if [[ -n "$genesis_ecotone" && "$genesis_ecotone" != "null" ]]; then
      if [[ -z "$rollup_ecotone" || "$rollup_ecotone" == "null" || "$rollup_ecotone" != "$genesis_ecotone" ]]; then
        echo "Patching rollup ecotone_time to ${genesis_ecotone}"
        dasel put -f /config/rollup.json -s .ecotone_time -t int -v "${genesis_ecotone}"
      fi
    fi
    if [[ -n "$genesis_fjord" && "$genesis_fjord" != "null" ]]; then
      if [[ -z "$rollup_fjord" || "$rollup_fjord" == "null" || "$rollup_fjord" != "$genesis_fjord" ]]; then
        echo "Patching rollup fjord_time to ${genesis_fjord}"
        dasel put -f /config/rollup.json -s .fjord_time -t int -v "${genesis_fjord}"
      fi
    fi
    if [[ -n "$genesis_granite" && "$genesis_granite" != "null" ]]; then
      if [[ -z "$rollup_granite" || "$rollup_granite" == "null" || "$rollup_granite" != "$genesis_granite" ]]; then
        echo "Patching rollup granite_time to ${genesis_granite}"
        dasel put -f /config/rollup.json -s .granite_time -t int -v "${genesis_granite}"
      fi
    fi
    if [[ -n "$genesis_holocene" && "$genesis_holocene" != "null" ]]; then
      if [[ -z "$rollup_holocene" || "$rollup_holocene" == "null" || "$rollup_holocene" != "$genesis_holocene" ]]; then
        echo "Patching rollup holocene_time to ${genesis_holocene}"
        dasel put -f /config/rollup.json -s .holocene_time -t int -v "${genesis_holocene}"
      fi
    fi
    if [[ -n "$genesis_isthmus" && "$genesis_isthmus" != "null" ]]; then
      if [[ -z "$rollup_isthmus" || "$rollup_isthmus" == "null" || "$rollup_isthmus" != "$genesis_isthmus" ]]; then
        echo "Patching rollup isthmus_time to ${genesis_isthmus}"
        dasel put -f /config/rollup.json -s .isthmus_time -t int -v "${genesis_isthmus}"
      fi
    fi
    if [[ -n "$genesis_jovian" && "$genesis_jovian" != "null" ]]; then
      if [[ -z "$rollup_jovian" || "$rollup_jovian" == "null" || "$rollup_jovian" != "$genesis_jovian" ]]; then
        echo "Patching rollup jovian_time to ${genesis_jovian}"
        dasel put -f /config/rollup.json -s .jovian_time -t int -v "${genesis_jovian}"
      fi
    fi
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
    force_safe_rollup_system_scalar
    force_rollup_fork_times_from_env
    sync_rollup_fork_times_from_genesis

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
    force_safe_rollup_system_scalar
    force_rollup_fork_times_from_env
    sync_rollup_fork_times_from_genesis

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
