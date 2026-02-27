#!/usr/bin/env bash
# Run a single devnet test (TestSmokeWithoutTEE) after full setup.
# Usage from repo root (with Docker running):
#   nix develop . --command ./espresso/scripts/run-one-devnet-test.sh
set -euo pipefail

export OP_E2E_SKIP_ALLOC_GEN=1
export ESPRESSO_DEVNET_TESTS_LIVENESS_PERIOD=30s
export ESPRESSO_DEVNET_TESTS_OUTAGE_PERIOD=30s

OP_ROOT="${1:-$(cd "$(dirname "$0")/../.." && pwd)}"
OP_ROOT="$(realpath "${OP_ROOT}")"
cd "${OP_ROOT}"

# Load env from espresso/.env
if [[ -f espresso/.env ]]; then
  set -a
  # shellcheck source=/dev/null
  source <(grep -v '^#' espresso/.env | grep -v '^$' | sed 's/^export //')
  set +a
fi

echo "== Compile contracts =="
cd packages/contracts-bedrock && just build && just fix-proxy-artifact && cd "${OP_ROOT}"

echo "== Build op-deployer and prepare allocs =="
cd op-deployer && just && export PATH="${PWD}/bin:${PATH}" && cd "${OP_ROOT}"
cd espresso && ./scripts/prepare-allocs.sh && cd "${OP_ROOT}"

echo "== Build devnet images =="
cd espresso && COMPOSE_PROFILES=default docker compose build && cd "${OP_ROOT}"

echo "== Pre-generate L1 beacon genesis =="
cd "${OP_ROOT}/espresso"
dasel put -f deployment/l1-config/genesis.json -s .timestamp -v "$(printf '0x%x' $(date +%s))"
eth-beacon-genesis devnet \
  --quiet \
  --eth1-config deployment/l1-config/genesis.json \
  --config docker/l1-geth/beacon-config.yaml \
  --mnemonics docker/l1-geth/mnemonics.yaml \
  --state-output deployment/l1-config/genesis.ssz
cp docker/l1-geth/beacon-config.yaml deployment/l1-config/config.yaml
openssl rand -hex 32 > deployment/l1-config/jwt.txt
echo 0 > deployment/l1-config/deposit_contract_block.txt
echo 0x00000000219ab540356cBB839Cbe05303d7705Fa > deployment/l1-config/deposit_contract.txt
cd "${OP_ROOT}"

echo "== Verify pre-generation =="
test -f espresso/deployment/l1-config/genesis.ssz || (echo "Missing genesis.ssz" && exit 1)
test -s espresso/deployment/l1-config/genesis.ssz || (echo "genesis.ssz empty" && exit 1)
echo "Pre-generated L1 beacon files OK."

echo "== Run TestSmokeWithoutTEE =="
go test -timeout 25m -p 1 -count 1 -run 'TestSmokeWithoutTEE' -v ./espresso/devnet-tests/...

echo "== Test passed =="
