#!/bin/bash
set -euo pipefail

# Set the default ports if not provided.
L1_HTTP_PORT=${L1_HTTP_PORT:-8545}
L1_ENGINE_PORT=${L1_ENGINE_PORT:-8551}
L1_CHAIN_ID=${L1_CHAIN_ID:-11155111}

# Mode can be "genesis" or "geth" (default).
MODE=${MODE:-geth}

if [[ "$MODE" == "genesis" ]]; then
  echo "Running Genesis Initialization"

  # Create config directory if it doesn't exist.
  mkdir -p /config

  # Use pre-built genesis with deployed contracts instead of empty template
  if [[ ! -f "/config/genesis.json" ]]; then
      echo "Copying pre-built genesis with deployed contracts..."
      if [[ -f "/deployment/l1-config/genesis.json" ]]; then
          echo "Using pre-built genesis from deployment artifacts..."
          cp /deployment/l1-config/genesis.json /config/genesis.json
      else
          echo "Pre-built genesis not found, falling back to template..."
          cp /templates/devnet-genesis-template.json /config/genesis.json
      fi
  fi

  echo "Updating genesis timestamp..."
  dasel put -f /config/genesis.json -s .timestamp -v $(printf '0x%x\n' $(date +%s))

  echo "Generating consensus layer genesis..."
  eth-beacon-genesis devnet \
                    --quiet \
                    --eth1-config "/config/genesis.json" \
                    --config "/templates/beacon-config.yaml" \
                    --mnemonics "/templates/mnemonics.yaml" \
                    --state-output "/config/genesis.ssz"
  cp -r /templates/beacon-config.yaml /config/config.yaml

  echo "Generating validator keys..."
  rm -rf /config/keystore && \
  eth2-val-tools keystores --out-loc /config/keystore \
                            --source-mnemonic "$(yq -r '.[0].mnemonic' "/templates/mnemonics.yaml")" \
                            --source-min 0 \
                            --source-max 1
  mkdir -p /data/lighthouse-validator
  mkdir -p /data/lighthouse-validator/validators
  cp -r /config/keystore/keys/* /data/lighthouse-validator/validators/
  cp -r /config/keystore/secrets/ /data/lighthouse-validator/

  if [[ ! -f "/config/jwt.txt" ]]; then
      echo "Generating JWT secret..."
      openssl rand -hex 32 > "/config/jwt.txt"
  fi

  echo "0" > /config/deposit_contract_block.txt
  echo "0x00000000219ab540356cBB839Cbe05303d7705Fa" > /config/deposit_contract.txt

  echo "Genesis initialization complete"
  exit 0

elif [[ "$MODE" == "geth" ]]; then
  echo "=== Starting L1 Geth ==="

  # Wait for genesis.json to be available (in case genesis container is still running).
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
  if [[ ! -d "/data/geth" ]]; then
      echo "Initializing L1 Geth database..."
      geth --datadir /data --gcmode=archive --state.scheme=hash init /config/genesis.json
      echo "L1 Geth initialization completed"
  else
      echo "Geth database already initialized, skipping..."
  fi

  # Start Geth with the specified configuration.
  echo "Starting Geth..."
  exec geth \
    --datadir /data/geth \
    --http \
    --http.addr=0.0.0.0 \
    --http.api=eth,net,web3,admin,engine,miner \
    --http.port=${L1_HTTP_PORT} \
    --http.vhosts=* \
    --http.corsdomain=* \
    --authrpc.addr=0.0.0.0 \
    --authrpc.port=${L1_ENGINE_PORT} \
    --authrpc.vhosts=* \
    --authrpc.jwtsecret=/config/jwt.txt \
    --nodiscover \
    --maxpeers 0 \
    --networkid ${L1_CHAIN_ID} \
    --syncmode=full \
    --gcmode=archive

else
  echo "Unknown MODE: $MODE. Use 'genesis' or 'geth'"
  exit 1
fi
