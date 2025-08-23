# OP Stack Dockerfile, simplified from ops/docker/op-stack-go/Dockerfile

# Build arguments
ARG TARGET_BASE_IMAGE=alpine:3.22
ARG TARGETOS
ARG TARGETARCH

# Base builder image
FROM golang:1.24.5-alpine3.22 AS builder

RUN apk add --no-cache curl tar gzip make gcc musl-dev linux-headers git jq bash

# Install mise for toolchain management
RUN curl https://mise.run | MISE_INSTALL_PATH=/usr/local/bin/mise sh

# Install yq
RUN case "$TARGETARCH" in \
  "amd64") YQ_ARCH="amd64" ;; \
  "arm64") YQ_ARCH="arm64" ;; \
  *) YQ_ARCH="amd64" ;; \
  esac && \
  wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_$YQ_ARCH -O /usr/local/bin/yq && \
  chmod +x /usr/local/bin/yq

# Install versioned toolchain
COPY ./mise.toml .
RUN mise trust && mise install -v -y just && cp $(mise which just) /usr/local/bin/just && just --version

# Copy and download Go dependencies
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go mod download

# Copy source code
COPY . /app

# Build arguments for git metadata
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

# Build op-node
FROM op-cgo-builder AS op-node-builder
ARG OP_NODE_VERSION=v0.0.0
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build cd op-node && \
  CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
  go build -a -ldflags '-extldflags "-static"' \
  -o bin/op-node ./cmd/main.go

# Build op-batcher
FROM op-cgo-builder AS op-batcher-builder
ARG OP_BATCHER_VERSION=v0.0.0
WORKDIR /app/op-batcher
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH GITCOMMIT=$GIT_COMMIT GITDATE=$GIT_DATE VERSION="$OP_BATCHER_VERSION"
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build just op-batcher

# Build enclave-tools
FROM op-cgo-builder AS enclave-tools-builder
ARG ENCLAVE_TOOLS_VERSION=v0.0.0
WORKDIR /app/op-batcher
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH GITCOMMIT=$GIT_COMMIT GITDATE=$GIT_DATE VERSION="$ENCLAVE_TOOLS_VERSION"
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build just enclave-tools

# Build op-proposer
FROM builder AS op-proposer-builder
ARG OP_PROPOSER_VERSION=v0.0.0
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build cd op-proposer && make op-proposer  \
  GOOS=$TARGETOS GOARCH=$TARGETARCH GITCOMMIT=$GIT_COMMIT GITDATE=$GIT_DATE VERSION="$OP_PROPOSER_VERSION"

FROM golang:1.24-alpine AS deployment-utils-builder
RUN apk add gcc lld musl-dev # For CGO
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go install -ldflags '-linkmode external -extldflags "-static"' github.com/tomwright/dasel/v2/cmd/dasel@v2.8.1
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go install -ldflags '-linkmode external -extldflags "-static"' github.com/mikefarah/yq/v4@v4.47.1


# Final runtime images
FROM $TARGET_BASE_IMAGE AS op-node-target
RUN apk add gcc
ENV AZTEC_SRS_PATH /aztec/kzg10-aztec20-srs-1048584.bin
ADD "https://github.com/EspressoSystems/ark-srs/releases/download/v0.2.0/kzg10-aztec20-srs-1048584.bin" /aztec/kzg10-aztec20-srs-1048584.bin
COPY --from=op-node-builder /app/op-node/bin/op-node /usr/local/bin/

# Create config directory
RUN mkdir -p /config

# Include the deployment and contarcts.
COPY espresso/deployment/ /deployment/

CMD ["op-node"]

FROM $TARGET_BASE_IMAGE AS op-batcher-target
RUN apk add gcc
ENV AZTEC_SRS_PATH /aztec/kzg10-aztec20-srs-1048584.bin
ADD "https://github.com/EspressoSystems/ark-srs/releases/download/v0.2.0/kzg10-aztec20-srs-1048584.bin" /aztec/kzg10-aztec20-srs-1048584.bin
COPY --from=op-batcher-builder /app/op-batcher/bin/op-batcher /usr/local/bin/
CMD ["op-batcher"]

FROM $TARGET_BASE_IMAGE AS op-batcher-enclave-target
RUN apk add gcc docker bash jq curl wget
# Install enclaver for EIF creation
RUN curl -L https://github.com/enclaver-io/enclaver/releases/download/v0.5.0/enclaver-linux-x86_64-v0.5.0.tar.gz | tar xz --strip-components=1 -C /usr/local/bin enclaver-linux-x86_64-v0.5.0/enclaver
ENV AZTEC_SRS_PATH /aztec/kzg10-aztec20-srs-1048584.bin
ADD "https://github.com/EspressoSystems/ark-srs/releases/download/v0.2.0/kzg10-aztec20-srs-1048584.bin" /aztec/kzg10-aztec20-srs-1048584.bin
COPY --from=enclave-tools-builder /app/op-batcher/bin/enclave-tools /usr/local/bin/
# Copy go.mod to establish monorepo root for enclave-tools
COPY --from=enclave-tools-builder /app/go.mod /app/go.mod
# Copy entire packages and op-e2e directories to satisfy op-e2e config dependencies
COPY --from=enclave-tools-builder /app/packages /app/packages
COPY --from=enclave-tools-builder /app/op-e2e /app/op-e2e
# Copy ops directory required by enclave-tools
COPY --from=enclave-tools-builder /app/ops /app/ops
# Copy entire espresso directory and op-batcher to get all dependencies
COPY --from=enclave-tools-builder /app/espresso /app/espresso
COPY --from=enclave-tools-builder /app/op-batcher /app/op-batcher
# Create .env file with necessary variables (since it's excluded by .gitignore)
RUN echo "# Environment variables for enclave" > /app/espresso/.env && \
    echo "ESPRESSO_DEV_NODE_L1_DEPLOYMENT=skip" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_ESP_TOKEN_ADDRESS=0x00c042c4d5d913277ce16611a2ce6e9003554ad5" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_PLONK_VERIFIER_V2_ADDRESS=0x422a3492e218383753d8006c7bfa97815b44373f" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_STAKE_TABLE_ADDRESS=0x63e6dde6763c3466c7b45be880f7ee5dc2ca3e25" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_LIGHT_CLIENT_PROXY_ADDRESS=0x703848f4c85f18e3acd8196c8ec91eb0b7bd0797" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_FEE_CONTRACT_PROXY_ADDRESS=0x72ae2643518179cf01bca3278a37cead408de8b2" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_FEE_CONTRACT_ADDRESS=0x8f0342a7060e76dfc7f6e9debfad9b9ec919952c" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_STAKE_TABLE_PROXY_ADDRESS=0x9f5eac3d8e082f47631f1551f1343f23cd427162" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_ESP_TOKEN_PROXY_ADDRESS=0x9fcf7d13d10dedf17d0f24c62f0cf4ed462f65b7" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_PLONK_VERIFIER_ADDRESS=0xb4b46bdaa835f8e4b4d8e208b6559cd267851051" >> /app/espresso/.env && \
    echo "ESPRESSO_SEQUENCER_API_PORT=24000" >> /app/espresso/.env && \
    echo "ESPRESSO_DEV_NODE_PORT=24002" >> /app/espresso/.env && \
    echo "ESPRESSO_BUILDER_PORT=31003" >> /app/espresso/.env && \
    echo "L1_ENGINE_PORT=8551" >> /app/espresso/.env && \
    echo "L1_HTTP_PORT=8545" >> /app/espresso/.env && \
    echo "L1_BEACON_PORT=5052" >> /app/espresso/.env && \
    echo "ROLLUP_PORT=9545" >> /app/espresso/.env && \
    echo "VERIFIER_PORT=9546" >> /app/espresso/.env && \
    echo "CAFF_PORT=9547" >> /app/espresso/.env && \
    echo "OP_ENGINE_PORT=8552" >> /app/espresso/.env && \
    echo "OP_HTTP_PORT=8546" >> /app/espresso/.env && \
    echo "OPERATOR_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" >> /app/espresso/.env && \
    echo "OPERATOR_ADDRESS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" >> /app/espresso/.env && \
    echo "L1_CHAIN_ID=11155111" >> /app/espresso/.env && \
    echo "L2_CHAIN_ID=22266222" >> /app/espresso/.env && \
    echo "COMPOSE_PROFILES=default" >> /app/espresso/.env
RUN chmod +x /app/espresso/scripts/batcher-enclave-tool-image.sh /app/espresso/scripts/prepare-allocs.sh
WORKDIR /app
CMD ["enclave-tools"]

FROM $TARGET_BASE_IMAGE AS op-proposer-target
RUN apk add jq
COPY --from=op-proposer-builder /app/op-proposer/bin/op-proposer /usr/local/bin/
COPY --from=op-proposer-builder /app/espresso/deployment/deployer /deployer
CMD ["op-proposer"]

FROM $TARGET_BASE_IMAGE AS op-deployer-target
RUN apk add jq curl bash openssl
COPY --from=deployment-utils-builder /go/bin/dasel /usr/local/bin/
COPY --from=deployment-utils-builder /go/bin/yq /usr/local/bin/
CMD ["op-deployer"]
