#!/bin/bash
# Configure the shell to trigger the exit trap on any non successful error code
set -eu

# Configure a trap handler to run at the end of any unsuccessful script.  This
# will allow us to exist immediately following the failed test, so that it can
# be inspected / debugged.
trap "exit" INT TERM
trap end EXIT
end(){
    if [[ $? -ne 0 ]]; then
        echo "Tests failed :("
        echo "Figure out why"
        exit 1
    fi
}

# We run nuke before we run the tests, in order to make sure we're starting
# from a clean slate.
make nuke

# Some of the following tests depend on the existence of the forge-artifacts
# folder under `packages/contracts-bedrock`.  This folder is created by running
# the cannon tests, so we need to ensure that it is run first.
make -C ./cannon test
(cd packages/contracts-bedrock && just test)

just -f ./op-alt-da/justfile test
just -f ./op-batcher/justfile test
just -f ./op-chain-ops/justfile test
just -f ./op-challenger/justfile test
just -f ./op-conductor/justfile test
just -f ./op-dispute-mon/justfile test
just -f ./op-dripper/justfile test
make -C ./op-program test # TODO Philippe why did we need to put this test before the op-e2e one?
make -C ./op-e2e test
just -f ./op-node/justfile test
just -f ./op-proposer/justfile test
just -f ./op-service/justfile test
just -f ./op-supervisor/justfile test

# Just to be nice we run nuke again, so we don't have any residual state
# left around.
make nuke

echo Ok!
