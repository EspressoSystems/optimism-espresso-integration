# L1 Geth Dockerfile, simplified from ops/docker/deployment-utils/Dockerfile
FROM debian:12.7-slim

ENV DEBIAN_FRONTEND=noninteractive

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    curl \
    jq \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install Geth
# TODO (Keyao) Add SHA valication back
RUN ARCH=$(dpkg --print-architecture) && \
    echo "Detected architecture: $ARCH" && \
    case "$ARCH" in \
        "amd64") \
            GETH_URL="https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.15.11-36b2371c.tar.gz" && \
            GETH_DIR="geth-linux-amd64-1.15.11-36b2371c" \
            ;; \
        "arm64") \
            GETH_URL="https://gethstore.blob.core.windows.net/builds/geth-linux-arm64-1.15.11-36b2371c.tar.gz" && \
            GETH_DIR="geth-linux-arm64-1.15.11-36b2371c" \
            ;; \
        *) \
            echo "Unsupported architecture: $ARCH" && exit 1 \
            ;; \
    esac && \
    echo "Downloading: $GETH_URL" && \
    curl -L "$GETH_URL" -o geth.tar.gz && \
    tar -xvf geth.tar.gz && \
    mv "$GETH_DIR/geth" /usr/local/bin/geth && \
    rm -rf geth.tar.gz "$GETH_DIR" && \
    chmod +x /usr/local/bin/geth

# Create data directory
RUN mkdir -p /data

# Expose the RPC port
EXPOSE 8545

# Set working directory
WORKDIR /data

# Healthcheck
HEALTHCHECK --interval=3s --timeout=2s --retries=40 \
    CMD curl -f http://localhost:8545 || exit 1

# Default command
CMD ["geth", "--help"]
