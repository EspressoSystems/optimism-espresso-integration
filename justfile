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

devnet-tests: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v ./espresso/devnet-tests/...

devnet-smoke-test-without-tee: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -run 'TestSmokeWithoutTEE' -v ./espresso/devnet-tests/...

devnet-challenge-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestChallengeGame ./espresso/devnet-tests/...


devnet-withdraw-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestWithdrawal ./espresso/devnet-tests/...

build-devnet: compile-contracts
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
 (cd packages/contracts-bedrock && forge build --offline --skip "/**/test/**")

build-batcher-enclave-image:
 (cd kurtosis-devnet && just op-batcher-enclave-image)

espresso_tests_timeout := "35m"
espresso-tests timeout=espresso_tests_timeout: compile-contracts
 go test -timeout={{timeout}} -p=1 -count=1 ./espresso/environment

espresso-enclave-tests:
  ESPRESSO_RUN_ENCLAVE_TESTS=true go test -timeout={{espresso_tests_timeout}} -p=1 -count=1 ./espresso/enclave-tests/...


IMAGE_NAME := "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-fix-cors"
remove-espresso-containers:
  docker remove --force $(docker ps -q --filter ancestor={{IMAGE_NAME}})

forge_artifacts_dir:="packages/contracts-bedrock/forge-artifacts"
bindings_dir:="op-batcher/bindings"
gen_bindings_cmd:="./espresso/scripts/gen_bindings.sh"
gen-bindings:
  {{gen_bindings_cmd}} {{forge_artifacts_dir}}/BatchInbox.sol/BatchInbox.json > ./{{bindings_dir}}/batch_inbox.go
  {{gen_bindings_cmd}} {{forge_artifacts_dir}}/BatchAuthenticator.sol/BatchAuthenticator.json > ./{{bindings_dir}}/batch_authenticator.go

smoke-tests: compile-contracts
 go test -run ^TestEspressoDockerDevNodeSmokeTest$ ./espresso/environment -v

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
