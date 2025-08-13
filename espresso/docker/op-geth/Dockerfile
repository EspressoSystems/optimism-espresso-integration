# OP Geth Dockerfile

# Build arguments.
ARG TARGETOS
ARG TARGETARCH
ARG GIT_COMMIT
ARG GIT_DATE

# Rust builder for Espresso crypto libraries
FROM rust:1.88.0-alpine3.22 AS rust-builder
# TODO: Check the hash of the Espresso GO library when switch to the new one.
# <https://app.asana.com/1/1208976916964769/project/1209392461754458/task/1210755885228839?focus=true>
ARG ESPRESSO_NETWORK_GO_VER=0.0.34
RUN apk add perl make openssl-dev musl-dev gcc
ADD https://github.com/EspressoSystems/espresso-network-go/archive/refs/tags/v$ESPRESSO_NETWORK_GO_VER.tar.gz /source.tgz
RUN tar -oxzf /source.tgz
WORKDIR /espresso-network-go-$ESPRESSO_NETWORK_GO_VER
RUN --mount=type=cache,target=/usr/local/cargo/registry \
  --mount=type=cache,target=/usr/local/cargo/git/db \
  --mount=type=cache,target=/espresso-network-go/verification/rust/target \
  cargo build --release  --locked  --manifest-path ./verification/rust/Cargo.toml
RUN mkdir -p /libespresso
RUN cp ./verification/rust/target/release/libespresso_crypto_helper.a \
  /libespresso/libespresso_crypto_helper-aarch64-unknown-linux-gnu.a
RUN cp ./verification/rust/target/release/libespresso_crypto_helper.a \
  /libespresso/libespresso_crypto_helper-x86_64-unknown-linux-gnu.a

# CGO builder for components that need Espresso crypto linking
FROM alpine:3.22 AS op-cgo-builder
# Install dependencies
RUN apk add musl-dev gcc go g++ curl tar gzip make gcc linux-headers git jq bash yq
# Install just from mise
COPY ./mise.toml .
RUN curl -L https://github.com/casey/just/releases/download/$(yq '.tools.just' mise.toml)/just-$(yq '.tools.just' mise.toml)-x86_64-unknown-linux-musl.tar.gz | \
  tar xz -C /usr/local/bin just
# Go sources
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
# Copy rust libs for dynamic linking
COPY --from=rust-builder /libespresso/* /lib
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

# Add a cache-busting layer
RUN echo "Cache bust: $(date)" > /cache-bust.txt

# For healtcheck and JSON operations.
RUN apk add curl jq openssl

# Install dasel for JSON manipulation.
RUN curl -sSL "https://github.com/TomWright/dasel/releases/latest/download/dasel_linux_amd64" -o /usr/local/bin/dasel && \
    chmod +x /usr/local/bin/dasel

# Copy binary from builder stage.
COPY --from=op-deployer-builder /op-deployer /usr/local/bin/

# Include the deployer and the initialization script.
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
