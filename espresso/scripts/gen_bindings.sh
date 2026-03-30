#!/usr/bin/env bash
function generate_go_bindings() {
  local json_file="$1"
  local contract_name_full
  local contract_name
  local base_name
  local abi_data
  local bin_data

  if [[ -z "$json_file" ]]; then
    echo "Error: Please provide the path to the .json file." >&2
    return 1
  fi

  # If exact file doesn't exist, try to find a versioned variant.
  # Forge sometimes produces artifacts like Contract.0.8.25.json instead of Contract.json.
  if [[ ! -f "$json_file" ]]; then
    local dir name versioned_file
    dir=$(dirname "$json_file")
    name=$(basename "$json_file" .json)
    versioned_file=$(ls "$dir/$name".*.json 2>/dev/null | sort -V | head -n1)
    if [[ -n "$versioned_file" ]]; then
      echo "Note: $json_file not found, using versioned artifact: $versioned_file" >&2
      json_file="$versioned_file"
    fi
  fi

  if [[ ! -f "$json_file" ]]; then
    echo "Error: File not found: $json_file" >&2
    return 1
  fi

  base_name=$(basename "$json_file")
  contract_name_full="${base_name%.json}"
  contract_name="${contract_name_full#I}"   # Remove leading 'I' if present
  IFS='.' read -r contract_name _ <<< "$contract_name"

  if ! cd "$(dirname "$json_file")"; then
    echo "Error: Could not change directory to $(dirname "$json_file")" >&2
    return 1
  fi

  if ! abi_data=$(cat "$base_name" | jq -r '.abi'); then
    echo "Error extracting ABI from $base_name" >&2
    return 1
  fi

  if ! bin_data=$(cat "$base_name" | jq -r '.bytecode.object'); then
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


if bindings=$(generate_go_bindings "$1"); then
  echo "$bindings"
else
  exit 1 # Propagate the error exit code from the function
fi
