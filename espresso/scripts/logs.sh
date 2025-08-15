#!/bin/bash

# Celo-Espresso Integration logging script
# This script outputs the logs of a specific service.

# Usage: ./logs.sh <service-name>

# Valid service names
VALID_SERVICES=(
    "l1-genesis"
    "l1-geth"
    "l1-beacon"
    "l1-validator"
    "espresso-dev-node"
    "l2-genesis"
    "op-geth"
    "l2-rollup"
    "op-node-sequencer"
    "op-node-verifier"
    "caff-node"
    "op-batcher"
    "op-proposer"
)

# Get the service name
SERVICE_NAME="$1"

# Run docker compose logs
echo "📋 Showing logs for $SERVICE_NAME (Press Ctrl+C to exit)"
echo "----------------------------------------"
docker compose logs -f $SERVICE_NAME
