#!/bin/bash

# Celo-Espresso Integration startup script
# This script builds the deployer, prepares the contracts, and start devnet services.

set -e  # Exit on any error

# NOTE: Start from the espresso/scripts directory
echo "Setting up Celo-Espresso Integration..."

# Step 1: Build the op-deployer
# NOTE: This step needs to be re-run if the op-deployer is modified
echo "👉 Step 1: Building op-deployer..."
cd ../../op-deployer
just
echo "✅ op-deployer build complete"

# Step 2: Compile the contracts
# NOTE: This step needs to be re-run if the contracts are modified
echo "👉 Step 2: Compiling contracts..."
cd ../
just compile-contracts
echo "✅ Contracts compilation complete"

# Step 3: Shut down all containers
echo "👉 Step 3: Shutting down all containers..."
cd espresso
docker compose down -v --remove-orphans
echo "✅ All containers shut down"

# Step 4: Prepare contract allocations
# NOTE: This step needs to be re-run if the contracts are modified
echo "👉 Step 4: Preparing contract allocations..."
./scripts/prepare-allocs.sh
echo "✅ Contract allocations prepared"

# Check for AWS Nitro Enclave support
echo "👉 Checking for AWS Nitro Enclave support..."
if command -v nitro-cli &>/dev/null && \
   (nitro-cli describe-enclaves 2>/dev/null | grep -qE "EnclaveID|\[\]" || [ -e /dev/nitro_enclaves ]); then
    echo "✅ AWS Nitro Enclave support detected"
else
    echo "⚠️  WARNING: AWS Nitro Enclave support not detected! TEE components will fail."
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo ""
    [[ ! $REPLY =~ ^[Yy]$ ]] && { echo "❌ Startup cancelled."; exit 1; }
fi

# Step 5: Prepare op-batcher-enclave image
echo "👉 Step 5: Preparing op-batcher-enclave image..."
docker system prune -f
cd .. && rm -f espresso/shared/*
# Ensure required tools are installed
if ! command -v nitro-cli &>/dev/null; then
    echo "Installing nitro-cli..."
    sudo amazon-linux-extras install aws-nitro-enclaves-cli -y
    sudo yum install aws-nitro-enclaves-cli-devel -y
fi
cd op-batcher && just op-batcher && just enclave-tools && cd ../espresso
docker compose stop op-batcher-tee
docker compose rm -f op-batcher-tee
./scripts/build-enclave-image.sh
echo "✅ op-batcher-enclave image prepared"

# Step 6: Build docker compose
echo "👉 Step 6: Building docker compose..."
COMPOSE_PROFILES=tee docker compose build
echo "✅ Docker compose build complete"

# Step 7: Start services
echo "👉 Step 7: Starting services..."
COMPOSE_PROFILES=tee docker compose up -d
echo "✅ Services started in detached mode"

echo "🎉 Startup complete! All services should now be running."
