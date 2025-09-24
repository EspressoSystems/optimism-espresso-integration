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
./scripts/shutdown.sh
echo "✅ All containers shut down"

# Step 4: Prepare contract allocations
# NOTE: This step needs to be re-run if the contracts are modified
echo "👉 Step 4: Preparing contract allocations..."
./scripts/prepare-allocs.sh
echo "✅ Contract allocations prepared"

# Step 5: Build docker compose
echo "👉 Step 5: Building docker compose..."
if [ "$USE_TEE" = "True" ] || [ "$USE_TEE" = "true" ]; then
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
    echo "Building with TEE profile..."
    COMPOSE_PROFILES=tee docker compose build
else
    docker compose build
fi
echo "✅ Docker compose build complete"

# Step 6: Start services
echo "👉 Step 6: Starting services..."
if [ "$USE_TEE" = "True" ] || [ "$USE_TEE" = "true" ]; then
    COMPOSE_PROFILES=tee docker compose up -d
else
    docker compose up -d
fi
echo "✅ Services started in detached mode"

echo "🎉 Startup complete! All services should now be running."
