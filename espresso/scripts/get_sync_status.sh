#!/bin/bash
# This is a convenience script to fetch data from the optimism node for
# "optimism_syncStatus" RPC method.

echo "NODE $NODE_NAME"
JSON_DATA=$(curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"optimism_syncStatus","params":[],"id":1}' $RPC_ADDRESS 2>/dev/null)

# Make sure the the RPC call was successful
if [ $? -ne 0 ]; then
    echo "Failed to connect to $RPC_ADDRESS"
    exit 1
fi


# Store the results for easier processing
RESULT=$(echo $JSON_DATA | jq .result)

# Extract and print some fields from the JSON response
output_block_details() {
    BLOCK=$(echo $RESULT | jq -r .$1)
    echo "$1: ($(echo $BLOCK | jq -r .number))"
    echo "  hash: $(echo $BLOCK | jq -r .hash)"
    echo "  parentHash: $(echo $BLOCK | jq -r .parentHash)"
    echo "  timestamp: $(echo $BLOCK | jq -r .timestamp)"
}

# Output the block details in a simple format
output_block_details "current_l1"
output_block_details "current_l1_finalized"
output_block_details "head_l1"
output_block_details "safe_l1"
output_block_details "finalized_l1"
echo
output_block_details "unsafe_l2"
output_block_details "safe_l2"
output_block_details "finalized_l2"
