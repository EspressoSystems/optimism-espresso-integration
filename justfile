# Variable
gid := `id -g`
uid := `id -u`

# Build all Rust binaries (release) for sysgo tests.
build-rust-release:
  cd rust && cargo build --release --bin kona-node --bin kona-supervisor
  cd op-rbuilder && cargo build --release -p op-rbuilder --bin op-rbuilder
  cd rollup-boost && cargo build --release -p rollup-boost --bin rollup-boost

# Checks that locked NUT bundles have not been modified.
check-nut-locks:
  go run ./ops/scripts/check-nut-locks
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
[script('bash')]
update-op-geth ref:
    set -euo pipefail
    ref="{{ref}}"
    if [ -z "$ref" ]; then echo "error: provide a hash/tag/branch"; exit 1; fi
    tmpl=$(printf "\173\173.Version\175\175")
    ver=$(go list -m -f "$tmpl" github.com/ethereum-optimism/op-geth@"$ref")
    if [ -z "$ver" ]; then echo "error: couldn't resolve $ref"; exit 1; fi
    go mod edit -replace=github.com/ethereum/go-ethereum=github.com/ethereum-optimism/op-geth@"$ver"
    go mod tidy
    echo "Updated op-geth to $ver"

# Prints the latest stable semver tag for a component (excludes pre-releases).
latest-tag component:
    @git tag -l '{{ component }}/v*' --sort=-v:refname | grep -E '^[^/]+/v[0-9]+\.[0-9]+\.[0-9]+$' | head -1

# Prints the latest RC tag for a component.
latest-rc-tag component:
    @git tag -l '{{ component }}/v*' --sort=-v:refname | grep -E '^[^/]+/v[0-9]+\.[0-9]+\.[0-9]+-rc\.[0-9]+$' | head -1

# Generates release notes between two tags using git-cliff.
# <from> and <to> can be explicit tags (e.g. v1.16.5), or:
#   'latest'    - resolves to the latest stable tag (vX.Y.Z)
#   'latest-rc' - resolves to the latest RC tag (vX.Y.Z-rc.N)
#   'develop'   - (only for <to>) uses the develop branch tip with --unreleased
#
# Set <mode> to 'offline' to skip GitHub API calls (faster, but no PR metadata).
#
# Examples:
#   just release-notes op-node                          # latest stable -> latest RC (default)
#   just release-notes op-node latest develop           # all unreleased changes since the latest stable release
#   just release-notes op-node latest develop offline   # same, but without GitHub API calls
#   just release-notes op-node v1.16.5 v1.16.6          # explicit tags
#
# Requires GITHUB_TOKEN for git-cliff's GitHub integration (unless mode=offline):
#   GITHUB_TOKEN=$(gh auth token) just release-notes op-node
[script('zsh')]
release-notes component from='latest' to='latest-rc' mode='':
    set -euo pipefail
    if [ "{{ mode }}" != "offline" ] && [ -z "${GITHUB_TOKEN:-}" ]; then
        echo "warning: GITHUB_TOKEN is not set. Set it like: GITHUB_TOKEN=\$(gh auth token) just release-notes ..."
        exit 1
    fi
    resolve_tag() {
        case "$1" in
            latest)    git tag -l "{{ component }}/v*" --sort=-v:refname | grep -E '^[^/]+/v[0-9]+\.[0-9]+\.[0-9]+$' | head -1 ;;
            latest-rc) git tag -l "{{ component }}/v*" --sort=-v:refname | grep -E '^[^/]+/v[0-9]+\.[0-9]+\.[0-9]+-rc\.[0-9]+$' | head -1 ;;
            v[0-9]*) echo "{{ component }}/$1" ;;
            *)       echo "error: invalid tag '$1'; expected 'latest', 'latest-rc', or 'vX.Y.Z...'" >&2; return 1 ;;
        esac
    }
    from_tag=$(resolve_tag "{{ from }}")
    if [ -z "$from_tag" ]; then echo "error: could not resolve from tag '{{ from }}' for {{ component }}"; exit 1; fi
    include_path_args=()
    case "{{ component }}" in
        op-node|op-batcher|op-proposer|op-challenger)
            include_path_args=(
                --include-path "{{ component }}/**/*"
                --include-path "go.*"
                --include-path "op-core/**/*"
                --include-path "op-service/**/*"
            )
            ;;
        op-reth)
            include_path_args=(
                --include-path "rust/{{ component }}/**/*"
                --include-path "rust/Cargo.toml"
                --include-path "rust/op-alloy/**/*"
                --include-path "rust/alloy-op*/**/*"
            )
            ;;
        kona-*)
            include_path_args=(
                --include-path "rust/kona/**/*"
                --include-path "rust/Cargo.toml"
                --include-path "rust/op-alloy/**/*"
                --include-path "rust/alloy-op*/**/*"
            )
            ;;
        *)
            echo "error: component must be one of: op-node, op-batcher, op-proposer, op-challenger, op-reth, kona-*; is {{ component }}"
            exit 1
            ;;
    esac
    tag_args=()
    if [ "{{ to }}" = "develop" ]; then
        tag_args=(--unreleased)
        range_end="develop"
    else
        to_tag=$(resolve_tag "{{ to }}")
        if [ -z "$to_tag" ]; then echo "error: could not resolve to tag '{{ to }}' for {{ component }}"; exit 1; fi
        tag_args=(--tag "$to_tag")
        range_end="$to_tag"
    fi
    echo "Generating release notes for ${from_tag}..${range_end}"
    offline_args=()
    if [ "{{ mode }}" = "offline" ]; then
        offline_args=(--offline)
    fi
    git cliff \
        --config .github/cliff.toml \
        "${include_path_args[@]}" \
        --tag-pattern "${from_tag}" \
        "${tag_args[@]}" \
        "${offline_args[@]}" \
        -- "${from_tag}..${range_end}"
