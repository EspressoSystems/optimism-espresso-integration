FROM golang:1.23-bookworm AS builder
WORKDIR /app
COPY . .

# Build the op-deployer binary
RUN CGO_ENABLED=0 go build -o /op-deployer ./op-deployer/cmd/op-deployer

FROM debian:bookworm-20240812-slim
ENTRYPOINT ["/op-up"]
COPY op-up /op-up

# Install ca-certificates so that HTTPS requests work
RUN apt-get update && apt-get install -y ca-certificates

# Symlink onto the PATH
RUN ln -s /op-up /usr/local/bin/op-up
