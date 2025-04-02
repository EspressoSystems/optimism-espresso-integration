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
# make -C ./cannon test
(cd packages/contracts-bedrock && just build-dev)

just -f ./op-batcher/justfile test
just -f ./op-node/justfile test

# Just to be nice we run nuke again, so we don't have any residual state
# left around.
make nuke

echo Ok!
