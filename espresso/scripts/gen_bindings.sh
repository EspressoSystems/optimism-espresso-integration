#!/usr/bin/env bash

function generate_go_bindings() {
  local json_file="$1"
  local contract_name
  local base_name
  local abi_data
  local bin_data

  if [[ -z "$json_file" ]]; then
    echo "Error: Please provide the path to the .json file." >&2
    return 1
  fi

  if [[ ! -f "$json_file" ]]; then
    echo "Error: File not found: $json_file" >&2
    return 1
  fi

  base_name=$(basename "$json_file")
  contract_name_full="${base_name%.json}"
  contract_name="${contract_name_full#I}" # Remove leading 'I' if present
  IFS='.' read -r contract_name _ <<< "$contract_name" # Remove everything after '.'

  cd "$(dirname "$json_file")" || {
    echo "Error: Could not change directory to $(dirname "$json_file")" >&2
    return 1
  }

  abi_data=$(cat "$base_name" | jq -r '.abi')
  if [[ $? -ne 0 ]]; then
    echo "Error extracting ABI from $base_name" >&2
    return 1
  fi

  bin_data=$(cat "$base_name" | jq -r '.bytecode.object')
  if [[ $? -ne 0 ]]; then
    echo "Error extracting bytecode from $base_name" >&2
    return 1
  fi

  abigen --abi <(echo "$abi_data") --bin <(echo "$bin_data") --pkg bindings --type "$contract_name"
  local abigen_status=$?
  if [[ $abigen_status -ne 0 ]]; then
    echo "Error running abigen for $contract_name (exit code: $abigen_status)" >&2
    return $abigen_status
  fi

  return 0 # Indicate success
}

# Main execution block
if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <path_to_contract_json>" >&2
  exit 1
fi

bindings=$(generate_go_bindings "$1")
if [[ $? -eq 0 ]]; then
  echo "$bindings"
else
  exit 1 # Propagate the error exit code from the function
fi
