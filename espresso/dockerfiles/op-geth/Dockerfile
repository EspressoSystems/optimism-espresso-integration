# OP Geth Dockerfile

FROM us-docker.pkg.dev/oplabs-tools-artifacts/images/op-geth:v1.101503.2-rc.3

# Include the genesis file, the secret file, and the initialization script.
COPY config/l2-genesis-devnet.json /l2-genesis-devnet.json
COPY config/jwt.txt /jwt.txt
COPY dockerfiles/l1-geth/op-geth-init.sh /usr/local/bin/op-geth-init.sh

# Create config directory
RUN mkdir -p /config

# Run the initialization script.
RUN chmod +x /usr/local/bin/l1-geth-init.sh

# Use the initialization script as the entrypoint.
ENTRYPOINT ["/usr/local/bin/op-geth-init.sh"]
