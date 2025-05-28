# Run the tests
tests:
 ./run_all_tests.sh

fast-tests:
 ./run_fast_tests.sh

golint:
 golangci-lint run -E goimports,sqlclosecheck,bodyclose,asciicheck,misspell,errorlint --timeout 5m -e "errors.As" -e "errors.Is" ./...

run-test7: compile-contracts
  go test ./espresso/environment/7_stateless_batcher_test.go -v

run-test9: compile-contracts
  go test ./espresso/environment/9_pipeline_enhancement_test.go -v

run-test12: compile-contracts
  go test ./espresso/environment/12_enforce_majority_rule_test.go -v

compile-contracts:
 (cd packages/contracts-bedrock && just build-dev)

build-batcher-enclave-image:
 (cd kurtosis-devnet && just op-batcher-enclave-image)

run-test4: compile-contracts
 go test ./espresso/environment/4_confirmation_integrity_with_reorgs_test.go -v

espresso-tests: compile-contracts
 go test -timeout=30m -p=1 -count=1 ./espresso/environment

espresso-enclave-tests: compile-contracts build-batcher-enclave-image
 ESPRESSO_RUN_ENCLAVE_TESTS=true go test -timeout=30m -p=1 -count=1 ./espresso/enclave_tests/...

IMAGE_NAME := "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:release-colorful-snake"
remove-espresso-containers:
  docker remove --force $(docker ps -q --filter ancestor={{IMAGE_NAME}})

smoke-tests: compile-contracts
 go test -run ^TestEspressoDockerDevNodeSmokeTest$ ./espresso/environment -v

# Clean up everything before running the tests
nuke:
 make nuke


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
