# OP Geth Dockerfile

FROM us-docker.pkg.dev/oplabs-tools-artifacts/images/op-geth:v1.101503.2-rc.3

# Copy config files into the image
COPY ./l2-genesis-devnet.json /l2-genesis-devnet.json
COPY jwt.txt /jwt.txt

# Create config directory
RUN mkdir -p /config
