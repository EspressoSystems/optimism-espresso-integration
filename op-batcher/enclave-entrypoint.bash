#!/usr/bin/env bash

# Entrypoint for op-batcher running in enclaver image.
# Main goal of the script is to rewrite the URLs passed to the batcher to use the Odyn proxy
# and recover batcher's CLI arguments from ENCLAVE_BATCHER_ARGS env variable (there's no way
# to directly pass commandline arguments when starting EIF images)

# Log environment information for debugging
echo "=== Enclave Environment Debug Info ==="
echo "PATH: $PATH"
echo "Working directory: $(pwd)"
echo "Available commands:"
which op-batcher trurl nc socat 2>&1 || echo "Some commands not found"
echo "======================================"

# We will need to start a proxy for each of those urls
URL_ARG_RE='^(--altda\.da-server|--espresso\.urls|--espresso.\l1-url|--espresso.rollup-l1-url|--l1-eth-rpc|--l2-eth-rpc|--rollup-rpc|--signer\.endpoint)(=|$)'

# Re-populate the arguments passed through the environment
if [ -n "$ENCLAVE_BATCHER_ARGS" ]; then
  eval set -- "$ENCLAVE_BATCHER_ARGS"
fi

if ! ODYN_PROXY_PORT=$(trurl --url "$http_proxy" --get "{port}"); then
        echo "Failed to parse HTTP_PROXY with" >&2
        return 1
 fi

echo "[DEBUG] Testing Odyn proxy connectivity on port $ODYN_PROXY_PORT" >&2
if nc -z 127.0.0.1 $ODYN_PROXY_PORT 2>/dev/null; then
  echo "Odyn proxy functional on port $ODYN_PROXY_PORT"
else
  echo "[ERROR] Odyn proxy unreachable on port $ODYN_PROXY_PORT"
  echo "[DEBUG] Network interfaces:" >&2
  ip addr show 2>&1 || ifconfig 2>&1 || echo "No network info available"
  exit 1
fi

unset http_proxy HTTP_PROXY https_proxy HTTPS_PROXY

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

wait_for_port() {
  local port="$1"

  for ((i=0; i<100; i++)); do
      if nc -z 127.0.0.1 "$port" 2>/dev/null; then
          return 0
      fi
      sleep 0.3
  done

  echo "Error: socat did not open port $port in time" >&2
  return 1
}

launch_socat() {
    local original_url="$1"
    local socat_port="$2"

    echo "[DEBUG] launch_socat called with URL: $original_url, port: $socat_port" >&2

    local host port scheme
    if ! read -r host port scheme < <(trurl --url "$original_url" --default-port --get "{host} {port} {scheme}"); then
        echo "Failed to parse URL" >&2
        return 1
    fi

    echo "[DEBUG] Parsed URL - host: $host, port: $port, scheme: $scheme" >&2

    # If original host was 127.0.0.1, we need to map it to `host` inside the enclave
    if [[ "$host" == "localhost" ]] || [[ "$host" == "127.0.0.1" ]] || [[ "$host" == "::1" ]]; then
      echo "Rewriting '$host' to 'host'" >&2
      host="host"
    fi

    if [[ "$scheme" != "http" ]] && [[ "$scheme" != "https" ]]; then
        echo "Invalid scheme: '$scheme'. Only http and https are supported." >&2
        return 1
    fi

    echo "[DEBUG] Starting socat proxy: ${host}:${port} via Odyn proxy at port ${ODYN_PROXY_PORT}" >&2
    # start socat
    socat -t 10 -d TCP4-LISTEN:"${socat_port}",reuseaddr,fork PROXY:127.0.0.1:"$host":"$port",proxyport="${ODYN_PROXY_PORT}" > /dev/null 2>&1 &
    socat_pid=$!
    disown "$socat_pid"

    echo "[DEBUG] socat started with PID: $socat_pid" >&2

    wait_for_port "${socat_port}" || {
      echo "[ERROR] Failed to start socat proxy on port ${socat_port}" >&2
      kill "$socat_pid" 2>/dev/null
      wait "$socat_pid" 2>/dev/null
      return 1
    }

    echo "[DEBUG] socat proxy ready on port ${socat_port}" >&2

    # return socat-proxied url
    local rewritten_url
    rewritten_url="$(trurl --url "$original_url" --set host="127.0.0.1" --set port="$socat_port")"

    echo "[DEBUG] URL rewrite:" >&2
    echo "[DEBUG]   Original:  $original_url" >&2
    echo "[DEBUG]   Rewritten: $rewritten_url" >&2

    # Verify path is preserved
    local original_path rewritten_path
    original_path="$(trurl --url "$original_url" --get "{path}")"
    rewritten_path="$(trurl --url "$rewritten_url" --get "{path}")"
    if [[ "$original_path" != "$rewritten_path" ]]; then
        echo "[ERROR] URL path was not preserved!" >&2
        echo "[ERROR]   Original path:  $original_path" >&2
        echo "[ERROR]   Rewritten path: $rewritten_path" >&2
    else
        echo "[DEBUG]   Path preserved: $original_path" >&2
    fi

    echo "$rewritten_url"

    return 0
}

# Initialize arrays for filtered arguments and extracted URLs
filtered_args=()
url_args=()

SOCAT_PORT=10001
echo "Arguments: $@"
# Process all arguments
while [ $# -gt 0 ]; do
    echo "Processing argument: $1"
    # Check if the argument matches the URL pattern
    if [[ $1 =~ $URL_ARG_RE ]]; then
      echo "Found URL argument: $1"
      # Extract the flag part and possible value part
      flag=${BASH_REMATCH[1]}

      # extract value from "--flag=value" or "--flag value"
      if [[ "$1" == *=* ]]; then
        value="${1#*=}"
      else
        shift || { echo "$flag missing value"; exit 1; }
        value="$1"
      fi

      # Handle comma-separated values for any flag
      if [[ "$value" == *","* ]]; then
        IFS=',' read -r -a parts <<< "$value"
        for part in "${parts[@]}"; do
          if ! new_url=$(launch_socat "$part" "$SOCAT_PORT"); then
            echo "Failed to launch socat for $flag=$part"; exit 1
          fi
          echo "Rewritten: $new_url"
          url_args+=("${flag}=${new_url}")
          ((SOCAT_PORT++))
        done
      else
        if ! new_url=$(launch_socat "$value" "$SOCAT_PORT"); then
          echo "Failed to launch socat for $flag=$value"; exit 1
        fi
        echo "Rewritten: $new_url"
        url_args+=("$flag" "$new_url")
        ((SOCAT_PORT++))
      fi
    else
      filtered_args+=("$1")
    fi
    shift
  done


# Combine the rewritten URL arguments with the other arguments
all_args=("${filtered_args[@]}" "${url_args[@]}")

echo "=== Final op-batcher arguments ===" >&2
echo "Total arguments: ${#all_args[@]}" >&2
for i in "${!all_args[@]}"; do
  echo "  [$i]: ${all_args[$i]}" >&2
done
echo "===================================" >&2

echo "${all_args[@]}"
echo "[DEBUG] Launching op-batcher..." >&2
op-batcher "${all_args[@]}"
exit_code=$?
echo "Debug: op-batcher exited with code $exit_code"
exit $exit_code
