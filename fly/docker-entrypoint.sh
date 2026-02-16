#!/bin/sh
set -e

# Materialize key files from base64-encoded Fly.io secrets.
# Fly secrets are env vars — credential files must be written to disk at startup.

mkdir -p /tmp/keys

# EPIC: decode private key PEM from base64
if [ -n "$EPIC_PRIVATE_KEY_B64" ]; then
  echo "$EPIC_PRIVATE_KEY_B64" | base64 -d > /tmp/keys/epic-private-key.pem
  chmod 600 /tmp/keys/epic-private-key.pem
fi

# GCP: decode service account JSON from base64
if [ -n "$GCP_CREDENTIALS_B64" ]; then
  echo "$GCP_CREDENTIALS_B64" | base64 -d > /tmp/keys/gcp-credentials.json
  chmod 600 /tmp/keys/gcp-credentials.json
  export GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/gcp-credentials.json
fi

# CONFIG_FILE must be set in fly.dev.toml [env] block.
#   epic: /app/config.fly.epic.yaml
#   gcp:  /app/config.fly.gcp.yaml
if [ -z "$CONFIG_FILE" ]; then
  echo "ERROR: CONFIG_FILE env var not set. Set it in fly.dev.toml [env]." >&2
  exit 1
fi

exec /app/mcp-server -config "$CONFIG_FILE" -http
