#!/bin/sh
set -e

docker build --no-cache -t op-batcher:app -f kurtosis-devnet/enclaver/Dockerfile .
sudo enclaver build --file kurtosis-devnet/enclaver/enclaver.yaml
sudo enclaver run enclave-batcher:latest
