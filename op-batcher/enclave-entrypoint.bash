#!/usr/bin/env bash

# Entrypoint for op-batcher running in enclaver image.
# Main goal of the script is to rewrite the URLs passed to the batcher to use the Odyn proxy
# and recover batcher's CLI arguments from ENCLAVE_BATCHER_ARGS env variable (there's no way
# to directly pass commandline arguments when starting EIF images)

# We will need to start a proxy for each of those urls
URL_ARG="^(--altda\.da-server|--espresso-url|--l1-eth-rpc|--l2-eth-rpc|--rollup-rpc|--signer\.endpoint)$"

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

unset http_proxy HTTP_PROXY https_proxy HTTPS_PROXY

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

    local host port scheme
    if ! read -r host port scheme < <(trurl --url "$original_url" --default-port --get "{host} {port} {scheme}"); then
        echo "Failed to parse URL" >&2
        return 1
    fi

    # If original host was 127.0.0.1, we need to map it to `host` inside the enclave
    if [[ "$host" == "localhost" ]] || [[ "$host" == "127.0.0.1" ]] || [[ "$host" == "::1" ]]; then
      echo "Rewriting '$host' to 'host'" >&2
      host="host"
    fi

    if [[ "$scheme" != "http" ]] && [[ "$scheme" != "https" ]]; then
        echo "Invalid scheme: '$scheme'. Only http and https are supported." >&2
        return 1
    fi

    # start socat
    socat -t 10 -d TCP4-LISTEN:"${socat_port}",reuseaddr,fork PROXY:127.0.0.1:"$host":"$port",proxyport="${ODYN_PROXY_PORT}" > /dev/null 2>&1 &
    socat_pid=$!
    disown "$socat_pid"

    wait_for_port "${socat_port}" || {
      kill "$socat_pid" 2>/dev/null
      wait "$socat_pid" 2>/dev/null
      return 1
    }

    # return socat-proxied url
    echo "$(trurl --url "$original_url" --set host="127.0.0.1" --set port="$socat_port")"

    return 0
}

# Initialize arrays for filtered arguments and extracted URLs
filtered_args=()
url_args=()

SOCAT_PORT=10001
# Process all arguments
while [ $# -gt 0 ]; do
    # Check if the argument matches the URL pattern
    if [[ $1 =~ $URL_ARG ]]; then
        # Extract the flag part and possible value part
        flag=${BASH_REMATCH[1]}

        if [ $# -gt 1 ]; then
            shift
            value="$1"
        else
          echo "$flag doesn't have a value"
          exit 1
        fi

        echo "Rewriting $flag=$value"
        if ! new_url=$(launch_socat "$value" "$SOCAT_PORT"); then
            echo "Failed to launch socat for $flag=$value"
            exit 1
        fi
        echo "Rewritten: $new_url"
        url_args+=("$flag" "$new_url")

        ((SOCAT_PORT++))
    else
        # This is not a URL argument, add it to filtered args
        filtered_args+=("$1")
    fi
    shift
done

# Combine the rewritten URL arguments with the other arguments
all_args=("${filtered_args[@]}" "${url_args[@]}")

echo "${all_args[@]}"
op-batcher "${all_args[@]}"
