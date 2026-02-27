# OP Geth Dockerfile

# Build arguments.
ARG TARGETOS
ARG TARGETARCH
ARG GIT_COMMIT
ARG GIT_DATE

# CGO builder for components that need Espresso crypto linking (go.mod requires go >= 1.24.0)
FROM golang:1.24-alpine AS op-cgo-builder
# Install dependencies
RUN apk add musl-dev gcc g++ curl tar gzip make linux-headers git jq bash
# Install just (fixed version and arch to avoid mise.toml/yq dependency in build)
ARG TARGETARCH
RUN case "$TARGETARCH" in \
  "amd64") JUST_ARCH="x86_64-unknown-linux-musl" ;; \
  "arm64") JUST_ARCH="aarch64-unknown-linux-musl" ;; \
  *) JUST_ARCH="x86_64-unknown-linux-musl" ;; \
  esac && \
  wget -q "https://github.com/casey/just/releases/download/1.37.0/just-1.37.0-${JUST_ARCH}.tar.gz" -O /tmp/just.tar.gz && \
  tar -xzf /tmp/just.tar.gz -C /usr/local/bin just && rm /tmp/just.tar.gz && just --version

# Go sources
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
# Warm-up the cache
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go mod download
COPY . /app

# Build op-deployer
FROM op-cgo-builder AS op-deployer-builder
ARG OP_DEPLOYER_VERSION=v0.0.0
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH GITCOMMIT=$GIT_COMMIT GITDATE=$GIT_DATE VERSION="$OP_DEPLOYER_VERSION"
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build cd op-deployer && \
  go build -ldflags '-linkmode external -extldflags "-static"' -o /op-deployer ./cmd/op-deployer

FROM us-docker.pkg.dev/oplabs-tools-artifacts/images/op-geth:v1.101503.2-rc.3

# For healtcheck and JSON operations.
RUN apk add curl jq openssl

# Install dasel for JSON manipulation.
RUN case $(uname -m) in \
  "arm64"|"aarch64") ARCH="arm64" ;; \
  *) ARCH="amd64" ;; \
  esac && \
  curl -sSL "https://github.com/TomWright/dasel/releases/download/v2.8.1/dasel_linux_$ARCH" -o /usr/local/bin/dasel && \
  chmod +x /usr/local/bin/dasel

# Copy binary from builder stage.
COPY --from=op-deployer-builder /op-deployer /usr/local/bin/

# Include the deployment artifacts and the initialization script.
COPY espresso/deployment/ /deployment/
COPY espresso/deployment/deployer /deployer
COPY espresso/docker/op-geth/op-geth-init.sh /op-geth-init.sh

# Run the initialization script.
RUN chmod +x /op-geth-init.sh

# Create data directory.
RUN mkdir -p /data

# Expose the RPC ports.
EXPOSE 8546 8552

# Set working directory.
WORKDIR /data

# Use the initialization script as the entrypoint.
ENTRYPOINT ["/op-geth-init.sh"]
