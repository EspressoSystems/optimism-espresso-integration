#!/bin/bash
set -euo pipefail
set -x

echo "[*] Setting up Cachix"
cachix authtoken $1
cachix use espresso-systems-private
echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf

echo "[*] Cloning repo and checking out branch $BRANCH_NAME..."
git clone https://github.com/EspressoSystems/optimism-espresso-integration.git
cd optimism-espresso-integration
git checkout "$BRANCH_NAME"
git submodule update --init --recursive
# Poblate cachix cahe
nix flake archive --json | jq -r '.path,(.inputs|to_entries[].value.path)' | cachix push espresso-systems-private

echo "[*] Starting Docker..."
sudo systemctl enable --now docker
sudo usermod -a -G docker ec2-user
sudo chown ec2-user /var/run/docker.sock

echo "[*] Configuring Nitro Enclaves..."
sudo systemctl stop nitro-enclaves-allocator.service || true
echo -e '---\nmemory_mib: 4096\ncpu_count: 2' | sudo tee /etc/nitro_enclaves/allocator.yaml
sudo systemctl start nitro-enclaves-allocator.service


echo "[*] Running tests in nix develop shell..."

nix develop --command bash -c "source ./espresso/.env && just compile-contracts-fast && just build-batcher-enclave-image && just espresso-enclave-tests"
