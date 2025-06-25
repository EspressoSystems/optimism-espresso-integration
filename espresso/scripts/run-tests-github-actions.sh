#!/bin/bash
set -euo pipefail
set -x

echo "[*] Setting up Nix and Cachix"
sh <(curl --proto '=https' --tlsv1.2 -L https://nixos.org/nix/install) --daemon
source /etc/profile.d/nix.sh
nix-env -iA cachix -f https://cachix.org/api/v1/install
mkdir -p ~/.config/nix
echo "trusted-users = root ec2-user" | sudo tee -a /etc/nix/nix.conf && sudo pkill nix-daemon
cachix authtoken $1
cachix use espresso-systems-private
echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf

echo "[*] Installing dependencies..."
sudo yum update -y
sudo yum install -y git docker
sudo amazon-linux-extras enable aws-nitro-enclaves-cli
sudo yum install -y aws-nitro-enclaves-cli-1.4.2

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

# Workaround due to https://github.com/foundry-rs/foundry/issues/4736
sudo yum install -y gcc
curl https://sh.rustup.rs -sSf | sh -s -- -y
. $HOME/.cargo/env
cargo install svm-rs
svm install 0.8.30

nix develop --command bash -c "just compile-contracts-fast && just build-batcher-enclave-image && just espresso-enclave-tests"
