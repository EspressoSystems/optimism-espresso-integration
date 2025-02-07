#!/bin/sh
set -e

# Build the docker image, only use "--no-cache" when needed
docker build -t op-batcher:app -f kurtosis-devnet/enclaver/Dockerfile .

# Build the enclave
sudo enclaver build --file kurtosis-devnet/enclaver/enclaver.yaml

# Run the enclave
sudo enclaver run enclave-batcher:latest
