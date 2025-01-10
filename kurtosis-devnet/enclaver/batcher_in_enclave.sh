#!/bin/sh
set -e

# Build the docker image
docker build --no-cache -t op-batcher:app -f kurtosis-devnet/enclaver/Dockerfile .

# Build the enclave
sudo enclaver build --file kurtosis-devnet/enclaver/enclaver.yaml

# Run the enclave
sudo enclaver run enclave-batcher:latest
