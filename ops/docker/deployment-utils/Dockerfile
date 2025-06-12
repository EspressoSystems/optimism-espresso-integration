FROM golang:1.22.7-bookworm AS go-base

RUN go install github.com/tomwright/dasel/v2/cmd/dasel@master

FROM debian:12.7-slim AS base

SHELL ["/bin/bash", "-c"]

ENV PATH=/root/.cargo/bin:/root/.foundry/bin:$PATH
ENV DEBIAN_FRONTEND=noninteractive
ENV SHELL=/bin/bash

RUN apt-get update && apt-get install -y curl git jq build-essential

RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs > rustup.sh && \
  chmod +x ./rustup.sh && \
  sh rustup.sh -y

RUN source $HOME/.profile && rustup update nightly
RUN curl -L https://foundry.paradigm.xyz | bash
RUN foundryup

FROM debian:12.7-slim

ENV PATH=/root/.cargo/bin:/root/.foundry/bin:$PATH
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y bash curl jq

SHELL ["/bin/bash", "-c"]

COPY --from=base /root/.foundry/bin/forge /usr/local/bin/forge
COPY --from=base /root/.foundry/bin/cast /usr/local/bin/cast
COPY --from=base /root/.foundry/bin/anvil /usr/local/bin/anvil
COPY --from=go-base /go/bin/dasel /usr/local/bin/dasel

# Install geth
RUN curl -L https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.15.11-36b2371c.tar.gz \
    -o geth.tar.gz \
 && echo "a14a4285daedf75ea04a7a298e6caa48d566a2786c93fc5e86ec2c5998c92455  geth.tar.gz" | sha256sum -c - \
 && tar -xvf geth.tar.gz \
 && mv geth-linux-amd64-1.15.11-36b2371c/geth /usr/local/bin/geth \
 && rm -rf geth.tar.gz geth-linux-amd64-1.15.11-36b2371c
