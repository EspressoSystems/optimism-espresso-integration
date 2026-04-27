#!/bin/bash
set -euo pipefail
set -x

echo "[*] Setting up Nix"
sh <(curl --proto '=https' --tlsv1.2 -sSfL https://nixos.org/nix/install) --daemon --yes
# shellcheck source=/dev/null
source /etc/profile.d/nix.sh
nix-env -iA cachix -f https://cachix.org/api/v1/install
mkdir -p ~/.config/nix
echo "trusted-users = root ec2-user" | sudo tee -a /etc/nix/nix.conf && sudo pkill nix-daemon

echo "[*] Installing dependencies..."
sudo dnf update -y
sudo dnf install -y git docker gcc

# Nitro Enclaves CLI for Amazon Linux 2023
sudo dnf install -y aws-nitro-enclaves-cli aws-nitro-enclaves-cli-devel
sudo systemctl enable docker
sudo systemctl start docker
sudo usermod -aG ne ec2-user || true
sudo usermod -aG docker ec2-user || true

# Rust + svm workaround
curl https://sh.rustup.rs -sSf | sh -s -- -y
# shellcheck source=/dev/null
. "$HOME/.cargo/env"
cargo install svm-rs
svm install 0.8.15
svm install 0.8.19
svm install 0.8.22
svm install 0.8.25
svm install 0.8.28
svm install 0.8.30
