#!/usr/bin/env bash

# Entrypoint for op-batcher running in enclaver image.
# Main goal: rewrite URLs to use socat proxies through Odyn
# This ensures both internal and external URLs work correctly through the enclave proxy

set -e

echo "=== Enclave Environment Debug Info ==="
echo "PATH: $PATH"
echo "Working directory: $(pwd)"
echo "Proxy: ${http_proxy:-not set}"
echo "======================================"

# Re-populate the arguments passed through the environment
if [ -n "$ENCLAVE_BATCHER_ARGS" ]; then
  eval set -- "$ENCLAVE_BATCHER_ARGS"
fi

# Verify Odyn proxy is available
if [ -z "$http_proxy" ]; then
    echo "[ERROR] http_proxy not set" >&2
    exit 1
fi

if ! ODYN_PROXY_PORT=$(trurl --url "$http_proxy" --get "{port}"); then
    echo "[ERROR] Failed to parse http_proxy" >&2
    exit 1
fi

echo "[DEBUG] Testing Odyn proxy on port $ODYN_PROXY_PORT..." >&2
if nc -z 127.0.0.1 $ODYN_PROXY_PORT 2>/dev/null; then
  echo "✓ Odyn proxy functional on port $ODYN_PROXY_PORT"
else
  echo "[ERROR] Odyn proxy unreachable on port $ODYN_PROXY_PORT" >&2
  exit 1
fi

# IMPORTANT: Unset proxy environment variables
# We'll use socat to handle all proxying instead of relying on Go's HTTP client
# This avoids DNS resolution issues with the "host" hostname
unset http_proxy HTTP_PROXY https_proxy HTTPS_PROXY no_proxy NO_PROXY

echo "[DEBUG] Using socat for all URL proxying (avoids DNS resolution issues)"
echo ""

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
    set -- "${original_args[@]}"
else
    echo "Received ${#received_args[@]} arguments via nc, merging with original arguments"
    set -- "${original_args[@]}" "${received_args[@]}"
fi

# Helper function to wait for socat to open a port
wait_for_port() {
    local port="$1"

    for ((i=0; i<100; i++)); do
        if nc -z 127.0.0.1 "$port" 2>/dev/null; then
            return 0
        fi
        sleep 0.3
    done

    echo "[ERROR] socat did not open port $port in time" >&2
    return 1
}

# Helper function to launch socat proxy for a URL
launch_socat() {
    local original_url="$1"
    local socat_port="$2"

    local host port scheme
    if ! read -r host port scheme < <(trurl --url "$original_url" --default-port --get "{host} {port} {scheme}"); then
        echo "[ERROR] Failed to parse URL: $original_url" >&2
        return 1
    fi

    # If original host was localhost/127.0.0.1, map it to 'host' inside the enclave
    if [[ "$host" == "localhost" ]] || [[ "$host" == "127.0.0.1" ]] || [[ "$host" == "::1" ]]; then
        echo "[DEBUG] Rewriting '$host' to 'host'" >&2
        host="host"
    fi

    if [[ "$scheme" != "http" ]] && [[ "$scheme" != "https" ]]; then
        echo "[ERROR] Invalid scheme: '$scheme'. Only http and https are supported." >&2
        return 1
    fi

    # Start socat to proxy through Odyn
    echo "[DEBUG] Starting socat: 127.0.0.1:${socat_port} -> PROXY:${host}:${port} via Odyn:${ODYN_PROXY_PORT}" >&2
    socat -t 10 -d TCP4-LISTEN:"${socat_port}",reuseaddr,fork PROXY:127.0.0.1:"$host":"$port",proxyport="${ODYN_PROXY_PORT}" > /dev/null 2>&1 &
    socat_pid=$!
    disown "$socat_pid"

    wait_for_port "${socat_port}" || {
        kill "$socat_pid" 2>/dev/null
        wait "$socat_pid" 2>/dev/null
        return 1
    }

    # Return socat-proxied URL
    local new_url
    new_url="$(trurl --url "$original_url" --set host="127.0.0.1" --set port="$socat_port")"
    echo "$new_url"

    return 0
}

# URL argument regex pattern
URL_ARG_RE='^(--altda\.da-server|--espresso\.urls|--espresso\.l1-url|--espresso\.rollup-l1-url|--l1-eth-rpc|--l2-eth-rpc|--rollup-rpc|--signer\.endpoint)(=|$)'

# Process all arguments
filtered_args=()
url_args=()
SOCAT_PORT=10001

echo "Processing arguments..."
while [ $# -gt 0 ]; do
    echo "[DEBUG] Processing argument: $1" >&2
    # Check if the argument matches the URL pattern
    if [[ $1 =~ $URL_ARG_RE ]]; then
        echo "[DEBUG] Found URL argument: $1" >&2
        flag=${BASH_REMATCH[1]}

        # Extract value from "--flag=value" or "--flag value"
        if [[ "$1" == *=* ]]; then
            value="${1#*=}"
        else
            shift || { echo "[ERROR] $flag missing value"; exit 1; }
            value="$1"
        fi

        # Handle comma-separated values for any flag
        if [[ "$value" == *","* ]]; then
            IFS=',' read -r -a parts <<< "$value"
            rewritten_parts=()
            for part in "${parts[@]}"; do
                if ! new_url=$(launch_socat "$part" "$SOCAT_PORT"); then
                    echo "[ERROR] Failed to launch socat for $flag=$part" >&2
                    exit 1
                fi
                echo "[DEBUG] Rewritten: $part -> $new_url" >&2
                rewritten_parts+=("$new_url")
                ((SOCAT_PORT++))
            done
            # Join with commas
            joined=$(IFS=,; echo "${rewritten_parts[*]}")
            url_args+=("${flag}=${joined}")
        else
            if ! new_url=$(launch_socat "$value" "$SOCAT_PORT"); then
                echo "[ERROR] Failed to launch socat for $flag=$value" >&2
                exit 1
            fi
            echo "[DEBUG] Rewritten: $value -> $new_url" >&2
            url_args+=("$flag" "$new_url")
            ((SOCAT_PORT++))
        fi
    else
        filtered_args+=("$1")
    fi
    shift
done

# Combine all arguments
all_args=("${filtered_args[@]}" "${url_args[@]}")

echo ""
echo "=== Final op-batcher arguments ==="
echo "Total arguments: ${#all_args[@]}"
for i in "${!all_args[@]}"; do
    echo "  [$i]: ${all_args[$i]}" >&2
done
echo "===================================" >&2
echo ""

echo "[DEBUG] Launching op-batcher..." >&2
op-batcher "${all_args[@]}"
exit_code=$?
echo "[DEBUG] op-batcher exited with code $exit_code" >&2
exit $exit_code
