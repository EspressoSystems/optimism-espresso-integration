# OP Geth Dockerfile

FROM us-docker.pkg.dev/oplabs-tools-artifacts/images/op-geth:v1.101503.2-rc.3

# Include the genesis and the secret files
COPY config/l2-genesis-devnet.json /l2-genesis-devnet.json
COPY config/jwt.txt /jwt.txt

# Create config directory
RUN mkdir -p /config
