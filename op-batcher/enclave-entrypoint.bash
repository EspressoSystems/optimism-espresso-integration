#!/usr/bin/env bash

# Entrypoint for op-batcher running in enclaver image.
# Main goal of the script is to use the Odyn HTTP proxy for external connections
# and recover batcher's CLI arguments from ENCLAVE_BATCHER_ARGS env variable (there's no way
# to directly pass commandline arguments when starting EIF images)

# Re-populate the arguments passed through the environment
if [ -n "$ENCLAVE_BATCHER_ARGS" ]; then
  eval set -- "$ENCLAVE_BATCHER_ARGS"
fi

if ! ODYN_PROXY_PORT=$(trurl --url "$http_proxy" --get "{port}"); then
        echo "Failed to parse HTTP_PROXY with" >&2
        return 1
 fi

if nc -z 127.0.0.1 $ODYN_PROXY_PORT 2>/dev/null; then
  echo "Odyn proxy functional"
else
  echo "Odyn proxy unreachable"
  exit 1
fi

# Keep HTTP_PROXY set so Go's net/http client uses it for all requests
export HTTP_PROXY="$http_proxy"
export HTTPS_PROXY="$http_proxy"

# Store the original arguments from ENCLAVE_BATCHER_ARGS
original_args=("$@")

# Launch nc listener to receive null-separated arguments
NC_PORT=8337
received_args=()

echo "Starting nc listener on port $NC_PORT (60 second timeout)"
{
    # Read null-separated arguments until we get \0\0
    while IFS= read -r -d '' arg; do
        if [[ -z "$arg" ]]; then
            # Empty argument signals end (\0\0)
            break
        fi
        received_args+=("$arg")
    done
} < <(nc -l -p "$NC_PORT" -w 60)

if [ ${#received_args[@]} -eq 0 ]; then
    echo "Warning: No arguments received via nc listener within 60 seconds, using original arguments"
    # Use original arguments from ENCLAVE_BATCHER_ARGS
    set -- "${original_args[@]}"
else
    echo "Received ${#received_args[@]} arguments via nc, merging with original arguments"
    # Merge: original args + received args
    set -- "${original_args[@]}" "${received_args[@]}"
fi

# Rewrite localhost URLs to use 'host' inside the enclave
# The enclave's /etc/hosts maps 'host' to the parent instance
declare -a final_args=()
URL_ARG_RE='^(--altda\.da-server|--espresso\.urls|--espresso\.l1-url|--espresso.rollup-l1-url|--l1-eth-rpc|--l2-eth-rpc|--rollup-rpc|--signer\.endpoint)(=|$)'

while [ $# -gt 0 ]; do
    echo "Processing argument: $1"
    # Check if the argument matches the URL pattern
    if [[ $1 =~ $URL_ARG_RE ]]; then
      echo "Found URL argument: $1"
      flag=${BASH_REMATCH[1]}

      # extract value from "--flag=value" or "--flag value"
      if [[ "$1" == *=* ]]; then
        value="${1#*=}"
      else
        shift || { echo "$flag missing value"; exit 1; }
        value="$1"
      fi

      # Rewrite localhost to 'host' for internal services
      if [[ "$value" == http://localhost:* ]] || [[ "$value" == http://127.0.0.1:* ]]; then
        value=$(echo "$value" | sed 's|localhost|host|; s|127\.0\.0\.1|host|')
        echo "Rewrote localhost to host: $value"
      fi

      final_args+=("$flag" "$value")
    else
      final_args+=("$1")
    fi
    shift
done

echo "Final arguments: ${final_args[@]}"
op-batcher "${final_args[@]}"
exit_code=$?
echo "Debug: op-batcher exited with code $exit_code"
exit $exit_code
