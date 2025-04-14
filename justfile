# Run the tests
tests:
 ./run_all_tests.sh

fast-tests:
 ./run_fast_tests.sh


compile-contracts:
 (cd packages/contracts-bedrock && just build-dev)

espresso-tests: compile-contracts
 #go test ./espresso/environment
 go test -run ^TestE2eDevNetWithEspressoSimpleTransactions$ ./espresso/environment -v

IMAGE_NAME := "ghcr.io/espressosystems/espresso-sequencer/espresso-dev-node:20250409-dev-node-pos-preview"
remove-espresso-containers:
  docker stop $(docker ps -q --filter ancestor={{IMAGE_NAME}})

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
