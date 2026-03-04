# Variable
gid := `id -g`
uid := `id -u`

# Build all Rust binaries (release) for sysgo tests.
build-rust-release:
  cd kona && cargo build --release --bin kona-node --bin kona-supervisor
  cd op-rbuilder && cargo build --release -p op-rbuilder --bin op-rbuilder
  cd rollup-boost && cargo build --release -p rollup-boost --bin rollup-boost

# Run the tests
tests:
  ./run_all_tests.sh

fast-tests:
 ./run_fast_tests.sh

# Devnet tests use Docker devnet and do not need op-e2e alloc gen. Skip it to avoid init panic.
_devnet_test_env := "OP_E2E_SKIP_ALLOC_GEN=1"
# Run devnet tests from espresso/ so docker compose finds docker-compose.yml (cwd = espresso when test runs).
_devnet_test_cmd := "cd espresso && U_ID={{uid}} GID={{gid}} {{_devnet_test_env}} go test"

devnet-tests: build-devnet
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -v ./devnet-tests/...

devnet-smoke-test-without-tee: build-devnet
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -run 'TestSmokeWithoutTEE' -v ./devnet-tests/...

# Quick: run test only (no build-devnet). Use after `just build-devnet` for fast re-runs.
devnet-smoke-test-without-tee-quick:
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -run 'TestSmokeWithoutTEE' -v ./devnet-tests/...

# Challenge test: same 30m timeout as other devnet tests. "Quick" = skip build-devnet (faster re-runs), not shorter test.
devnet-challenge-test: build-devnet
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -v -run TestChallengeGame ./devnet-tests/...

# Quick: run challenge test only (no build-devnet). Use after `just build-devnet` for fast re-runs.
devnet-challenge-test-quick:
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -v -run TestChallengeGame ./devnet-tests/...


devnet-forced-transaction-test: build-devnet
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -v -run TestForcedTransaction ./devnet-tests/...


devnet-withdraw-test: build-devnet
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -v -run TestWithdrawal ./devnet-tests/...

devnet-batcher-switching-test: build-devnet
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -v -run TestBatcherSwitching ./devnet-tests/...

devnet-batcher-active-publish-only-test: build-devnet
  {{_devnet_test_cmd}} -timeout 30m -p 1 -count 1 -v -run TestBatcherActivePublishOnly ./devnet-tests/...

build-devnet: stop-containers compile-contracts
  rm -Rf espresso/deployment
  (cd op-deployer && just)
  (cd espresso && ./scripts/prepare-allocs.sh && docker compose build)

# Same as build-devnet but skip stop-containers. Use when Docker returns 500 (e.g. API version mismatch);
# restart Docker Desktop or set DOCKER_API_VERSION=1.41 if needed.
build-devnet-skip-stop: compile-contracts
  rm -Rf espresso/deployment
  (cd op-deployer && just)
  (cd espresso && ./scripts/prepare-allocs.sh && docker compose build)

golint:
 golangci-lint run -E goimports,sqlclosecheck,bodyclose,asciicheck,misspell,errorlint --timeout 5m -e "errors.As" -e "errors.Is" ./...


compile-contracts:
 (cd packages/contracts-bedrock && just build-dev)

run-l1-espresso-contracts-tests: compile-contracts
 (cd packages/contracts-bedrock && forge test --match-path "/**/test/L1/Batch*.t.sol")

compile-contracts-fast:
 (cd packages/contracts-bedrock && forge build --offline --skip "/**/test/**" && just fix-proxy-artifact)

build-batcher-enclave-image:
 (cd kurtosis-devnet && just op-batcher-enclave-image)

espresso_tests_timeout := "35m"
espresso-tests timeout=espresso_tests_timeout: compile-contracts
 go test -timeout={{timeout}} -p=1 -count=1 ./espresso/environment

espresso-enclave-tests:
  ESPRESSO_RUN_ENCLAVE_TESTS=true go test -timeout={{espresso_tests_timeout}} -p=1 -count=1 ./espresso/enclave-tests/...


IMAGE_NAME := "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-20251120-lip2p-tcp-3855"
remove-espresso-containers:
  docker remove --force $(docker ps -q --filter ancestor={{IMAGE_NAME}})

forge_artifacts_dir:="packages/contracts-bedrock/forge-artifacts"
bindings_dir:="op-batcher/bindings"
gen_bindings_cmd:="./espresso/scripts/gen_bindings.sh"
gen-bindings:
  {{gen_bindings_cmd}} {{forge_artifacts_dir}}/BatchInbox.sol/BatchInbox.json > ./{{bindings_dir}}/batch_inbox.go
  {{gen_bindings_cmd}} {{forge_artifacts_dir}}/BatchAuthenticator.sol/BatchAuthenticator.json > ./{{bindings_dir}}/batch_authenticator.go
  {{gen_bindings_cmd}} {{forge_artifacts_dir}}/OPSuccinctFaultDisputeGame.sol/OPSuccinctFaultDisputeGame.json > ./{{bindings_dir}}/opsuccinct_fault_dispute_game.go

smoke-tests: compile-contracts
 go test -run ^TestEspressoDockerDevNodeSmokeTest$ ./espresso/environment -v

# Clean up everything before running the tests
nuke:
  make nuke

# Stop the containers. If you get "500 Internal Server Error" / API version, restart Docker Desktop
# or run: DOCKER_API_VERSION=1.41 just build-devnet-skip-stop (then build-devnet-skip-stop).
# Ensure deployment/deployer/succinct.env exists so compose can load when deployment dir is missing (e.g. first run).
stop-containers:
  mkdir -p espresso/deployment/deployer && touch espresso/deployment/deployer/succinct.env
  (cd espresso && U_ID={{uid}} GID={{gid}} docker compose down -v)

# Checks that TODO comments have corresponding issues.
todo-checker:
  ./ops/scripts/todo-checker.sh

# Runs semgrep on the entire monorepo.
semgrep:
  semgrep scan --config .semgrep/rules/ --error .

# Runs semgrep tests.
semgrep-test:
  semgrep scan --test --config .semgrep/rules/ .semgrep/tests/

# Runs shellcheck.
shellcheck:
  find . -type f -name '*.sh' -not -path '*/node_modules/*' -not -path './packages/contracts-bedrock/lib/*' -not -path './packages/contracts-bedrock/kout*/*' -exec sh -c 'echo "Checking $1"; shellcheck "$1"' _ {} \;
  find . -type f -name '*.sh' -not -path '*/node_modules/*' -not -path './packages/contracts-bedrock/lib/*' -not -path './packages/contracts-bedrock/kout*/*' -exec shfmt --diff {} \;

# Format shell scripts with shfmt.
shfmt-fix:
  find . -type f -name '*.sh' -not -path '*/node_modules/*' -not -path './packages/contracts-bedrock/lib/*' -not -path './packages/contracts-bedrock/kout*/*' -exec shfmt --write {} \;

# Generates a table of contents for the README.md file.
toc:
  md_toc -p github README.md

latest-versions:
  ./ops/scripts/latest-versions.sh

# Usage:
#   just update-op-geth 2f0528b
#   just update-op-geth v1.101602.4
#   just update-op-geth optimism
update-op-geth ref:
	@ref="{{ref}}"; \
	if [ -z "$ref" ]; then echo "error: provide a hash/tag/branch"; exit 1; fi; \
	tmpl=$(printf "\173\173.Version\175\175"); \
	ver=$(go list -m -f "$tmpl" github.com/ethereum-optimism/op-geth@"$ref"); \
	if [ -z "$ver" ]; then echo "error: couldn't resolve $ref"; exit 1; fi; \
	go mod edit -replace=github.com/ethereum/go-ethereum=github.com/ethereum-optimism/op-geth@"$ver"; \
	go mod tidy; \
	echo "Updated op-geth to $ver"
