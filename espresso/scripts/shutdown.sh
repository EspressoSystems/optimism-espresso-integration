#!/bin/bash

# Celo-Espresso Integration shutdown script
# This script shuts down devnet services.

docker compose down -v

# Stop and remove containers built from op-batcher-tee:espresso image
echo "Stopping containers built from op-batcher-tee:espresso image..."
CONTAINERS=$(docker ps -q --filter "ancestor=op-batcher-tee:espresso")
if [ ! -z "$CONTAINERS" ]; then
    echo "Stopping containers: $CONTAINERS"
    docker stop $CONTAINERS
    docker rm $CONTAINERS
    echo "Containers stopped and removed"
else
    echo "No running containers found with op-batcher-tee:espresso image"
fi
