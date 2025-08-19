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
    "op-geth-sequencer"
    "op-geth-verifier"
    "op-geth-caff-node"
    "l2-rollup"
    "op-node-sequencer"
    "op-node-verifier"
    "caff-node"
    "op-batcher"
    "op-proposer"
)

# Function to display usage
show_usage() {
    echo "Usage: $0 <service-name>"
    echo ""
    echo "Available services:"
    for service in "${VALID_SERVICES[@]}"; do
        echo "  • $service"
    done
    echo ""
    echo "Available aliases:"
    echo "  • dev-node → espresso-dev-node"
    echo "  • sequencer → op-node-sequencer"
    echo "  • verifier → op-node-verifier"
    echo "  • batcher → op-batcher"
    echo ""
    echo "Examples:"
    echo "  $0 op-node-sequencer"
    echo "  $0 sequencer"
    echo "  $0 dev-node"
}

# Function to check if service is valid
is_valid_service() {
    local service="$1"
    for valid_service in "${VALID_SERVICES[@]}"; do
        if [[ "$service" == "$valid_service" ]]; then
            return 0
        fi
    done
    return 1
}

# Function to resolve service name
resolve_service_name() {
    local input="$1"

    # Check if it's an alias
    case "$input" in
        "dev-node")
            echo "espresso-dev-node"
            return 0
            ;;
        "sequencer")
            echo "op-node-sequencer"
            return 0
            ;;
        "verifier")
            echo "op-node-verifier"
            return 0
            ;;
        "batcher")
            echo "op-batcher"
            return 0
            ;;
    esac

    # Check if it's a valid full service name
    if is_valid_service "$input"; then
        echo "$input"
        return 0
    fi

    return 1
}

# Check if argument is provided
if [[ $# -eq 0 ]]; then
    echo "❌ Error: No service name provided"
    echo ""
    show_usage
    exit 1
fi

# Get the service name
INPUT_NAME="$1"

# Resolve the service name
SERVICE_NAME=$(resolve_service_name "$INPUT_NAME")

if [[ $? -ne 0 ]]; then
    echo "❌ Error: Invalid service name or alias '$INPUT_NAME'"
    echo ""
    show_usage
    exit 1
fi

# Show what service we're actually viewing (helpful when using aliases)
if [[ "$INPUT_NAME" != "$SERVICE_NAME" ]]; then
    echo "📋 Alias '$INPUT_NAME' resolved to '$SERVICE_NAME'"
fi

# Run docker compose logs
echo "📋 Showing logs for $SERVICE_NAME (Press Ctrl+C to exit)"
echo "----------------------------------------"
docker compose logs -f $SERVICE_NAME
