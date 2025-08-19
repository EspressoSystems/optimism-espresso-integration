#!/bin/bash

OP_RPC_SEQUENCER=${OP_RPC_SEQUENCER:-http://localhost:9545}
OP_RPC_VERIFIER=${OP_RPC_VERIFIER:-http://localhost:9546}
OP_RPC_CAFF=${OP_RPC_CAFF:-http://localhost:9547}

set -euC pipefail

# Change the current directory to the script's directory
cd "$(dirname "$0")"

# If the tmux session already exists, we will attach to it.
if tmux has-session -t '=get_sync_status' 2>/dev/null; then
  echo "Tmux session 'get_sync_status' already exists. Exiting."
  tmux kill-session -t get_sync_status || true
fi

# Create a new tmux session, detached, named "get_sync_status"
tmux new-session -d -s get_sync_status \; \
    send-keys "NODE_NAME=sequencer RPC_ADDRESS=$OP_RPC_SEQUENCER watch -p -n 1 -c -d ./get_sync_status.sh" ENTER \; \
    split-window -h "NODE_NAME=verifier RPC_ADDRESS=$OP_RPC_VERIFIER watch -p -n 1 -c -d ./get_sync_status.sh" \; \
    split-window -h "NODE_NAME=caff-node RPC_ADDRESS=$OP_RPC_CAFF watch -p -n 1 -c -d ./get_sync_status.sh" \; \
    select-layout even-horizontal \; \
    attach
