#!/usr/bin/env bash

# Entrypoint for op-batcher running in enclaver image.
# Uses HTTPS_PROXY for external URLs (preserving SNI/Host headers)
# Only rewrites internal localhost URLs to map to enclave's "host"

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

# CRITICAL: Preserve proxy environment variables for Go's HTTP client
# This allows external HTTPS URLs to work with correct SNI and Host headers
export HTTPS_PROXY="$http_proxy"
export HTTP_PROXY="$http_proxy"
export https_proxy="$http_proxy"
export NO_PROXY="localhost,127.0.0.1,::1,host"
export no_proxy="$NO_PROXY"

echo "[DEBUG] Proxy environment configured:"
echo "  HTTPS_PROXY=$HTTPS_PROXY"
echo "  NO_PROXY=$NO_PROXY"
echo "[DEBUG] External URLs will use proxy with correct SNI/Host headers"
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

# Helper function to check if URL needs rewriting
# Only localhost/127.0.0.1 URLs need to be mapped to "host" inside enclave
is_local_url() {
    local url="$1"
    local host
    host="$(trurl --url "$url" --get "{host}" 2>/dev/null)" || return 1

    if [[ "$host" == "localhost" ]] || [[ "$host" == "127.0.0.1" ]] || [[ "$host" == "::1" ]]; then
        return 0  # is local
    fi
    return 1  # is external
}

# Helper function to rewrite localhost to "host" for enclave internal services
rewrite_local_url() {
    local url="$1"
    local host port scheme path

    host="$(trurl --url "$url" --get "{host}")" || return 1
    port="$(trurl --url "$url" --get "{port}")" || return 1
    scheme="$(trurl --url "$url" --get "{scheme}")" || return 1
    path="$(trurl --url "$url" --get "{path}")" || path=""

    # Map localhost to "host" (enclave's parent)
    if [[ "$host" == "localhost" ]] || [[ "$host" == "127.0.0.1" ]] || [[ "$host" == "::1" ]]; then
        host="host"
    fi

    # Reconstruct URL
    local new_url="$scheme://$host"
    if [ -n "$port" ]; then
        new_url="$new_url:$port"
    fi
    new_url="$new_url$path"

    echo "$new_url"
}

# URL argument regex pattern
URL_ARG_RE='^(--altda\.da-server|--espresso\.urls|--espresso\.l1-url|--espresso\.rollup-l1-url|--l1-eth-rpc|--l2-eth-rpc|--rollup-rpc|--signer\.endpoint)(=|$)'

# Process all arguments
filtered_args=()
url_args=()

echo "Processing arguments..."
while [ $# -gt 0 ]; do
    # Check if the argument matches the URL pattern
    if [[ $1 =~ $URL_ARG_RE ]]; then
        flag=${BASH_REMATCH[1]}

        # Extract value from "--flag=value" or "--flag value"
        if [[ "$1" == *=* ]]; then
            value="${1#*=}"
        else
            shift || { echo "$flag missing value"; exit 1; }
            value="$1"
        fi

        # Handle comma-separated values for any flag
        if [[ "$value" == *","* ]]; then
            IFS=',' read -r -a parts <<< "$value"
            rewritten_parts=()
            for part in "${parts[@]}"; do
                if is_local_url "$part"; then
                    new_url=$(rewrite_local_url "$part")
                    echo "[DEBUG] Rewriting local URL: $part -> $new_url" >&2
                    rewritten_parts+=("$new_url")
                else
                    echo "[DEBUG] Keeping external URL unchanged: $part" >&2
                    rewritten_parts+=("$part")
                fi
            done
            # Join with commas
            joined=$(IFS=,; echo "${rewritten_parts[*]}")
            url_args+=("${flag}=${joined}")
        else
            if is_local_url "$value"; then
                new_url=$(rewrite_local_url "$value")
                echo "[DEBUG] Rewriting local URL: $value -> $new_url" >&2
                url_args+=("$flag" "$new_url")
            else
                echo "[DEBUG] Keeping external URL unchanged: $value" >&2
                url_args+=("$flag" "$value")
            fi
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
exec op-batcher "${all_args[@]}"
