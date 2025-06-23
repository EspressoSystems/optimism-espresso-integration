#!/bin/bash
set -euo pipefail

echo "[*] Setting up Nix..."
sh <(curl --proto '=https' --tlsv1.2 -L https://nixos.org/nix/install) --daemon
source /etc/profile.d/nix.sh

mkdir -p ~/.config/nix
echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf

echo "[*] Installing dependencies..."
sudo yum update -y
sudo yum install -y git docker
sudo amazon-linux-extras enable aws-nitro-enclaves-cli
sudo yum install -y aws-nitro-enclaves-cli-1.4.2

echo "[*] Starting Docker..."
sudo systemctl enable --now docker
sudo usermod -a -G docker ec2-user
sudo chown ec2-user /var/run/docker.sock

echo "[*] Configuring Nitro Enclaves..."
sudo systemctl stop nitro-enclaves-allocator.service || true
echo -e '---\nmemory_mib: 4096\ncpu_count: 2' | sudo tee /etc/nitro_enclaves/allocator.yaml
sudo systemctl start nitro-enclaves-allocator.service

echo "[*] Cloning repo and checking out branch $BRANCH_NAME..."
git clone https://github.com/EspressoSystems/optimism-espresso-integration.git
cd optimism-espresso-integration
git checkout "$BRANCH_NAME"
git submodule update --init --recursive

echo "[*] Running tests in nix develop shell..."
nix develop --command just espresso-enclave-tests
