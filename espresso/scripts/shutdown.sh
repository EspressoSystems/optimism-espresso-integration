#!/bin/bash

# Celo-Espresso Integration shutdown script
# This script shuts down devnet services.

docker compose down -v

# Stop and remove containers built from op-batcher-tee:espresso image
echo "Stopping containers built from op-batcher-tee:espresso image..."
CONTAINERS=$(docker ps -aq --filter "ancestor=op-batcher-tee:espresso")
if [ ! -z "$CONTAINERS" ]; then
    echo "Stopping containers: $CONTAINERS"
    docker stop $CONTAINERS
    docker rm $CONTAINERS
    echo "Containers stopped and removed"
else
    echo "No running containers found with op-batcher-tee:espresso image"
fi

# Stop and remove containers built from op-proposer-tee:espresso image
echo "Stopping containers built from op-proposer-tee:espresso image..."
PROPOSER_CONTAINERS=$(docker ps -aq --filter "ancestor=op-proposer-tee:espresso")
if [ ! -z "$PROPOSER_CONTAINERS" ]; then
    echo "Stopping containers: $PROPOSER_CONTAINERS"
    docker stop $PROPOSER_CONTAINERS 2>/dev/null || true
    docker rm $PROPOSER_CONTAINERS 2>/dev/null || true
    echo "Containers stopped and removed"
else
    echo "No containers found with op-proposer-tee:espresso image"
fi

# Stop and remove batcher-enclaver containers that run the eif
echo "Stopping batcher-enclaver containers..."
ENCLAVE_CONTAINERS=$(docker ps -aq --filter "name=batcher-enclaver-")
if [ ! -z "$ENCLAVE_CONTAINERS" ]; then
    echo "Stopping enclave containers: $ENCLAVE_CONTAINERS"
    docker stop $ENCLAVE_CONTAINERS 2>/dev/null || true
    docker rm $ENCLAVE_CONTAINERS 2>/dev/null || true
    echo "Enclave containers stopped and removed"
else
    echo "No batcher-enclaver containers found"
fi
