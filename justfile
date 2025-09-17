# Variable
gid := `id -g`
uid := `id -u`

# Run the tests
tests:
 ./run_all_tests.sh

fast-tests:
 ./run_fast_tests.sh

devnet-tests: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v ./espresso/devnet-tests/...

devnet-smoke-test: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -run 'TestSmoke' -v ./espresso/devnet-tests/...

devnet-test-withdraw: build-devnet
  U_ID={{uid}} GID={{gid}} go test -timeout 30m -p 1 -count 1 -v -run TestWithdraw ./espresso/devnet-tests/...

build-devnet: compile-contracts
  rm -Rf espresso/deployment
  (cd op-deployer && just)
  (cd espresso && ./scripts/prepare-allocs.sh && docker compose build)

golint:
 golangci-lint run -E goimports,sqlclosecheck,bodyclose,asciicheck,misspell,errorlint --timeout 5m -e "errors.As" -e "errors.Is" ./...


compile-contracts:
 (cd packages/contracts-bedrock && just build-dev)

compile-contracts-fast:
 (cd packages/contracts-bedrock && forge build --offline --skip "/**/test/**")

build-batcher-enclave-image:
 (cd kurtosis-devnet && just op-batcher-enclave-image)

espresso_tests_timeout := "35m"
espresso-tests timeout=espresso_tests_timeout: compile-contracts
 go test -timeout={{timeout}} -p=1 -count=1 ./espresso/environment

espresso-enclave-tests:
  ESPRESSO_RUN_ENCLAVE_TESTS=true go test -timeout={{espresso_tests_timeout}} -p=1 -count=1 ./espresso/enclave-tests/...


IMAGE_NAME := "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-colorful-snake"
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

# Generates a table of contents for the README.md file.
toc:
  md_toc -p github README.md
