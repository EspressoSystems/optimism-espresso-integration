#!/bin/bash
set -euo pipefail
set -x

echo "[*] Setting up Cachix"
cachix authtoken $1
# Retry cachix use (cachix.org can return 502 Bad Gateway transiently)
for attempt in 1 2 3 4 5; do
  if cachix use espresso-systems-private; then
    break
  fi
  if [[ $attempt -eq 5 ]]; then
    echo "[!] Cachix still failing after 5 attempts (e.g. cachix.org 502). Continuing without cache."
  else
    echo "[*] Cachix failed (attempt $attempt/5), retrying in 60s..."
    sleep 60
  fi
done
echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf

echo "[*] Cloning repo and checking out branch $BRANCH_NAME..."
git clone https://github.com/EspressoSystems/optimism-espresso-integration.git
cd optimism-espresso-integration
git checkout "$BRANCH_NAME"
git submodule update --init --recursive
# Populate Cachix cache (best-effort; do not fail if Cachix is down)
nix flake archive --json | jq -r '.path,(.inputs|to_entries[].value.path)' | cachix push espresso-systems-private || true

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
