#!/bin/bash
set -euo pipefail
set -x

echo "[*] Setting up Nix"
sh <(curl --proto '=https' --tlsv1.2 -L https://nixos.org/nix/install) --daemon --no-confirm
source /etc/profile.d/nix.sh
nix-env -iA cachix -f https://cachix.org/api/v1/install
mkdir -p ~/.config/nix
echo "trusted-users = root ec2-user" | sudo tee -a /etc/nix/nix.conf && sudo pkill nix-daemon


echo "[*] Installing dependencies..."
sudo yum update -y
sudo yum install -y git docker
sudo amazon-linux-extras enable aws-nitro-enclaves-cli
sudo yum install -y aws-nitro-enclaves-cli-1.4.2


# Workaround due to https://github.com/foundry-rs/foundry/issues/4736
sudo yum install -y gcc
curl https://sh.rustup.rs -sSf | sh -s -- -y
. $HOME/.cargo/env
cargo install svm-rs
svm install 0.8.15
svm install 0.8.19
svm install 0.8.22
svm install 0.8.25
svm install 0.8.28
svm install 0.8.30
