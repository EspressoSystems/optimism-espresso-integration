# Variable
gid := `id -g`
uid := `id -u`

# Build all Rust binaries (release) for sysgo tests.
build-rust-release:
  cd kona && cargo build --release --bin kona-node --bin kona-supervisor
  cd op-rbuilder && cargo build --release -p op-rbuilder --bin op-rbuilder
  cd rollup-boost && cargo build --release -p rollup-boost --bin rollup-boost

# Run the tests

## OP stack tests
op-tests:
  ./run_all_tests.sh

fast-op-tests:
 ./run_fast_tests.sh

## Unit tests (no Docker, fast)
# Espresso core (streamer, batch buffer)
espresso-streamer-unit-tests:
  go test -count=1 -v ./espresso/


## Integration tests

espresso_tests_timeout := "35m"
# Max parallel tests. Each test spins up L1+L2+Espresso Docker containers.
# 4 is a good default for 32GB RAM; reduce to 2 on constrained machines.
espresso_tests_parallel := "4"

espresso-tests timeout=espresso_tests_timeout parallel=espresso_tests_parallel: compile-contracts-fast
 go test -timeout={{timeout}} -count=1 -parallel={{parallel}} -v ./espresso/environment

# Run espresso-tests without recompiling contracts (when artifacts already exist).
espresso-tests-no-compile timeout=espresso_tests_timeout parallel=espresso_tests_parallel:
 go test -timeout={{timeout}} -count=1 -parallel={{parallel}} -v ./espresso/environment

# Run a single espresso integration test by name regex.
espresso-test name timeout=espresso_tests_timeout: compile-contracts-fast
 go test -timeout={{timeout}} -count=1 -run '{{name}}' -v ./espresso/environment

# Run a single espresso integration test without recompiling contracts.
espresso-test-no-compile name timeout=espresso_tests_timeout:
 go test -timeout={{timeout}} -count=1 -run '{{name}}' -v ./espresso/environment

## Integration test groups (for targeted development)
# Caff node tests
espresso-tests-caff timeout=espresso_tests_timeout: compile-contracts-fast
 go test -timeout={{timeout}} -count=1 -parallel=3 -run 'TestE2eDevnetWithEspressoWithCaffNodeDeterministicDerivation|TestDeterministicDerivationExecutionState|TestFastDerivationAndCaffNode' -v ./espresso/environment

# Batcher tests (auth, inbox, stateless, fallback)
espresso-tests-batcher timeout=espresso_tests_timeout: compile-contracts-fast
 go test -timeout={{timeout}} -count=1 -parallel=4 -run 'TestE2eDevnetWithInvalidAttestation|TestE2eDevnetWithUnattestedBatcherKey|TestE2eDevnetWithoutAuthenticatingBatches|TestStatelessBatcher|TestBatcherSwitching|TestFallbackMechanism' -v ./espresso/environment

# Derivation pipeline and soft confirmation tests
espresso-tests-derivation timeout=espresso_tests_timeout: compile-contracts-fast
 go test -timeout={{timeout}} -count=1 -parallel=3 -run 'TestPipelineEnhancement|TestSequencerFeedConsistency|TestSoftConfirmation' -v ./espresso/environment

# Reorg and finality tests
espresso-tests-reorg timeout=espresso_tests_timeout: compile-contracts-fast
 go test -timeout={{timeout}} -count=1 -parallel=4 -run 'TestBatcherWaitForFinality|TestCaffNodeWaitForFinality|TestE2eDevnetWithL1Reorg|TestConfirmationIntegrityWithReorgs' -v ./espresso/environment

# Liveness and degradation tests
espresso-tests-liveness timeout=espresso_tests_timeout: compile-contracts-fast
 go test -timeout={{timeout}} -count=1 -parallel=2 -run 'TestE2eDevnetWithEspressoDegradedLiveness' -v ./espresso/environment

espresso-enclave-tests:
  ESPRESSO_RUN_ENCLAVE_TESTS=true go test -timeout={{espresso_tests_timeout}} -count=1 ./espresso/enclave-tests/...

smoke-tests: compile-contracts-fast
 go test -run ^TestEspressoDockerDevNodeSmokeTest$ ./espresso/environment -v


## Devnet tests
devnet-tests: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v ./espresso/devnet-tests/...

devnet-smoke-test-without-tee: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -run 'TestSmokeWithoutTEE' -v ./espresso/devnet-tests/...

devnet-challenge-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestChallengeGame ./espresso/devnet-tests/...


devnet-forced-transaction-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestForcedTransaction ./espresso/devnet-tests/...


devnet-withdraw-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestWithdrawal ./espresso/devnet-tests/...

devnet-batcher-switching-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestBatcherSwitching ./espresso/devnet-tests/...

devnet-batcher-active-publish-only-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestBatcherActivePublishOnly ./espresso/devnet-tests/...

build-devnet: stop-containers compile-contracts-fast
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

# Build the batcher enclave image for devnet tests.
build-batcher-enclave-image:
 (cd kurtosis-devnet && just op-batcher-enclave-image)


IMAGE_NAME := "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-20251120-lip2p-tcp-3855"
remove-espresso-containers:
  docker remove --force $(docker ps -q --filter ancestor={{IMAGE_NAME}})


# Contracts
forge_artifacts_dir:="packages/contracts-bedrock/forge-artifacts"
bindings_dir:="op-batcher/bindings"
gen_bindings_cmd:="./espresso/scripts/gen_bindings.sh"
gen-bindings:
  {{gen_bindings_cmd}} {{forge_artifacts_dir}}/BatchAuthenticator.sol/BatchAuthenticator.json > ./{{bindings_dir}}/batch_authenticator.go
  {{gen_bindings_cmd}} {{forge_artifacts_dir}}/OPSuccinctFaultDisputeGame.sol/OPSuccinctFaultDisputeGame.json > ./{{bindings_dir}}/opsuccinct_fault_dispute_game.go

# Clean up everything before running the tests
nuke:
  make nuke

# Stop the containers
stop-containers:
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
