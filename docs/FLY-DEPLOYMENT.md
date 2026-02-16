# Fly.io Deployment Guide

Deploy LangCare MCP FHIR Server as a remote Streamable HTTP MCP server on Fly.io. Any MCP-compatible AI agent can connect from anywhere.

```
Any AI Agent (Claude, GPT, Gemini, custom)
        |
        |  MCP over Streamable HTTP + Bearer Token
        v
+----------------------------------+
|  langcare-mcp-dev.fly.dev        |
|                                  |
|  LangCare MCP FHIR Server (Go)  |
|  fhir_read . fhir_search        |
|  fhir_create . fhir_update      |
+----------------+-----------------+
                 |  FHIR R4 REST + OAuth2/SMART on FHIR
                 v
+----------------------------------+
|  EMR (Epic / GCP Healthcare API) |
+----------------------------------+
```

## File Layout

```
fly/
  Dockerfile               # Multi-stage Go build
  docker-entrypoint.sh     # Key materialization + config selection
  fly.dev.toml             # Fly.io dev deployment config
  config.fly.epic.yaml     # EPIC provider config
  config.fly.gcp.yaml      # GCP provider config
  README.md                # Quick start guide
.dockerignore              # Excludes keys, docs, skills (must be at repo root)
```

## Prerequisites

```bash
brew install flyctl
fly auth login
```

## Step 1: Create App

```bash
fly apps create --name langcare-mcp-dev
```

Endpoint: `https://langcare-mcp-dev.fly.dev`

## Step 2: Set CONFIG_FILE in fly.dev.toml

Edit `fly/fly.dev.toml` to select your provider:

**EPIC:**
```toml
[env]
  LOG_LEVEL = "debug"
  CONFIG_FILE = "/app/config.fly.epic.yaml"
```

**GCP:**
```toml
[env]
  LOG_LEVEL = "debug"
  CONFIG_FILE = "/app/config.fly.gcp.yaml"
```

## Step 3: Set Secrets

### EPIC Provider

```bash
fly secrets set \
  EPIC_BASE_URL="https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4" \
  EPIC_CLIENT_ID="your-epic-client-id" \
  EPIC_TOKEN_URL="https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token" \
  EPIC_PRIVATE_KEY_B64="$(base64 < keys/epic/private-key.pem)" \
  MCP_AUTH_TOKENS="your-mcp-bearer-token" \
  --app langcare-mcp-dev
```

The entrypoint script decodes `EPIC_PRIVATE_KEY_B64` to `/tmp/keys/epic-private-key.pem` at startup. The PEM never touches the image or git.

### GCP Healthcare API Provider

```bash
fly secrets set \
  GCP_PROJECT_ID="your-project-id" \
  GCP_LOCATION="us-central1" \
  GCP_DATASET_ID="your-dataset-id" \
  GCP_FHIR_STORE_ID="your-fhir-store-id" \
  GCP_CREDENTIALS_B64="$(base64 < keys/gcp/service-account.json)" \
  MCP_AUTH_TOKENS="your-mcp-bearer-token" \
  --app langcare-mcp-dev
```

The entrypoint decodes `GCP_CREDENTIALS_B64` to `/tmp/keys/gcp-credentials.json` and sets `GOOGLE_APPLICATION_CREDENTIALS`.

## Step 4: Deploy

```bash
fly deploy -c fly/fly.dev.toml --app langcare-mcp-dev
```

## Step 5: Verify

```bash
curl https://langcare-mcp-dev.fly.dev/health

fly logs --app langcare-mcp-dev
fly status --app langcare-mcp-dev
```

## Step 6: Connect AI Agents

The server uses **Streamable HTTP** transport. MCP endpoint is `/mcp`.

### Claude Desktop

`claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "langcare-fhir": {
      "url": "https://langcare-mcp-dev.fly.dev/mcp",
      "headers": {
        "Authorization": "Bearer your-mcp-bearer-token"
      }
    }
  }
}
```

For clients without native Streamable HTTP support, use `mcp-remote` as a stdio-to-HTTP bridge:

```json
{
  "mcpServers": {
    "langcare-fhir": {
      "command": "npx",
      "args": [
        "mcp-remote",
        "https://langcare-mcp-dev.fly.dev/mcp",
        "--header",
        "Authorization: Bearer your-mcp-bearer-token"
      ]
    }
  }
}
```

### Any MCP Client

```
URL:   https://langcare-mcp-dev.fly.dev/mcp
Auth:  Authorization: Bearer <your-token>
Tools: fhir_read, fhir_search, fhir_create, fhir_update (auto-discovered)
```

## Docker Local Testing

Test the Docker image locally before deploying:

```bash
# Build
docker build -t langcare-mcp:local .

# Run (EPIC — requires key file on host)
docker run -p 8080:8080 \
  -e CONFIG_FILE=/app/config.fly.epic.yaml \
  -e EPIC_BASE_URL=https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4 \
  -e EPIC_CLIENT_ID=your-client-id \
  -e EPIC_TOKEN_URL=https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token \
  -e EPIC_PRIVATE_KEY_B64="$(base64 < keys/epic/private-key.pem)" \
  -e MCP_AUTH_TOKENS=test-token-123 \
  -e LOG_LEVEL=debug \
  langcare-mcp:local

# Test
curl http://localhost:8080/health
curl -H "Authorization: Bearer test-token-123" http://localhost:8080/mcp
```

## Provider Selection

| Provider | `CONFIG_FILE` value |
|----------|---------------------|
| EPIC | `/app/config.fly.epic.yaml` |
| GCP | `/app/config.fly.gcp.yaml` |

Set in `fly/fly.dev.toml` `[env]` block.

## Secrets Summary

| Secret | Provider | Required |
|--------|----------|----------|
| `MCP_AUTH_TOKENS` | All | Yes |
| `EPIC_BASE_URL` | EPIC | Yes |
| `EPIC_CLIENT_ID` | EPIC | Yes |
| `EPIC_TOKEN_URL` | EPIC | Yes |
| `EPIC_PRIVATE_KEY_B64` | EPIC | Yes |
| `GCP_PROJECT_ID` | GCP | Yes |
| `GCP_LOCATION` | GCP | Yes |
| `GCP_DATASET_ID` | GCP | Yes |
| `GCP_FHIR_STORE_ID` | GCP | Yes |
| `GCP_CREDENTIALS_B64` | GCP | Yes |

## Security

- Private keys and credentials are **never** in the Docker image or git
- Keys are base64-encoded Fly.io secrets, decoded to `/tmp/keys/` at container startup
- `.gitignore` blocks `*.pem`, `keys/`, `.env`
- `.dockerignore` blocks `*.pem`, `keys/`, `.env`
- MCP auth tokens required for HTTP transport (config validation enforces this)
- TLS terminated by Fly.io edge (`force_https = true`)
- Non-root container user (`mcpuser`, UID 1001)

## Operational Commands

```bash
# View logs
fly logs --app langcare-mcp-dev

# SSH into running machine
fly ssh console --app langcare-mcp-dev

# Scale (add machines/regions)
fly scale count 2 --region iad --app langcare-mcp-dev

# Rotate MCP auth token
fly secrets set MCP_AUTH_TOKENS="new-token" --app langcare-mcp-dev

# List releases (for rollback)
fly releases --app langcare-mcp-dev
```

## Cost Estimate (Dev, scale-to-zero)

| | Estimate |
|--|---------|
| VM | ~$0 when stopped |
| Bandwidth | negligible |
| **Total** | **~$0-2/mo** |
