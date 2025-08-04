# OP Geth Dockerfile

FROM us-docker.pkg.dev/oplabs-tools-artifacts/images/op-geth:v1.101503.2-rc.3

# Add a cache-busting layer
RUN echo "Cache bust: $(date)" > /cache-bust.txt

# For healtcheck and JSON operations.
RUN apk add curl jq

# Include the initialization script.
COPY op-geth-init.sh /op-geth-init.sh

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
