# Fly.io Deployment

Deploy LangCare MCP FHIR Server to Fly.io as a remote MCP server using **Streamable HTTP** transport. Any MCP-compatible AI agent (Claude, GPT, Gemini, custom) can connect from anywhere via HTTPS.

## Files

```
fly/
  Dockerfile               # Multi-stage Go build (~15MB image)
  docker-entrypoint.sh     # Decodes key files from secrets, starts server
  fly.dev.toml             # Fly.io deployment config (VM, scaling, health checks)
  config.fly.epic.yaml     # EPIC provider config (env vars expanded at startup)
  config.fly.cerner.yaml   # Cerner provider config (env vars expanded at startup)
  config.fly.gcp.yaml      # GCP provider config (env vars expanded at startup)
```

`.dockerignore` lives at repo root (Docker requirement).

## How It Works

```
fly secrets set ...          → Encrypted env vars stored on Fly.io
                               (credentials, project IDs, base64-encoded keys)
        |
        v
fly deploy                   → Builds Docker image, deploys container
        |
        v
Dockerfile                   → Copies Go binary + config YAMLs + entrypoint into image
        |                      (NO keys or credentials — only ${VAR} placeholders)
        |
        v  container starts
docker-entrypoint.sh         → 1. Reads base64-encoded secrets from Fly env vars
                               2. Decodes them to files in /tmp/keys/
                               3. Reads CONFIG_FILE env var (set in fly.dev.toml)
                               4. Starts: mcp-server -config <CONFIG_FILE> -http
        |
        v
config.Load()                → Reads the YAML config file
                               Runs os.ExpandEnv() — replaces ${GCP_PROJECT_ID},
                               ${EPIC_CLIENT_ID}, etc. with Fly secret values
        |
        v
Provider init                → EPIC:   loads PEM from /tmp/keys/epic-private-key.pem
                               GCP:    loads JSON from /tmp/keys/gcp-credentials.json
                               Cerner: uses client_id + client_secret directly
                               Authenticates with EMR backend
        |
        v
Server on :8080              → Fly.io routes https://langcare-mcp-dev.fly.dev → :8080
                               Health check at /health, MCP endpoint at /mcp
```

## Quick Start

### 1. Install Fly CLI

```bash
brew install flyctl
fly auth login
```

### 2. Create App

```bash
fly apps create --name langcare-mcp-dev
```

### 3. Choose Provider and Configure

---

#### EPIC

**Set provider in `fly/fly.dev.toml`:**

```toml
[env]
  LOG_LEVEL = "debug"
  CONFIG_FILE = "/app/config.fly.epic.yaml"
```

**Base64-encode your private key and set secrets:**

```bash
fly secrets set \
  EPIC_BASE_URL="https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4" \
  EPIC_CLIENT_ID="your-epic-client-id" \
  EPIC_TOKEN_URL="https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token" \
  EPIC_PRIVATE_KEY_B64="$(base64 < keys/epic/private-key.pem)" \
  MCP_AUTH_TOKENS="pick-a-bearer-token" \
  --app langcare-mcp-dev
```

**What happens at container startup:**

1. `docker-entrypoint.sh` sees `EPIC_PRIVATE_KEY_B64` env var (from Fly secrets)
2. Decodes it → writes `/tmp/keys/epic-private-key.pem` (exists only in container memory)
3. Reads `CONFIG_FILE` env var → `/app/config.fly.epic.yaml`
4. Starts `mcp-server -config /app/config.fly.epic.yaml -http`
5. `config.Load()` reads the YAML, `os.ExpandEnv()` replaces `${EPIC_CLIENT_ID}`, `${EPIC_BASE_URL}`, `${EPIC_TOKEN_URL}` with Fly secret values
6. EPIC provider loads PEM from `/tmp/keys/epic-private-key.pem`, signs JWT, authenticates with EPIC
7. Server listens on `:8080`, Fly routes `https://langcare-mcp-dev.fly.dev` to it

---

#### GCP Healthcare API

**Set provider in `fly/fly.dev.toml`:**

```toml
[env]
  LOG_LEVEL = "debug"
  CONFIG_FILE = "/app/config.fly.gcp.yaml"
```

**Base64-encode your service account JSON and set secrets:**

```bash
fly secrets set \
  GCP_PROJECT_ID="your-project-id" \
  GCP_LOCATION="us-central1" \
  GCP_DATASET_ID="your-dataset-id" \
  GCP_FHIR_STORE_ID="your-fhir-store-id" \
  GCP_CREDENTIALS_B64="$(base64 < keys/gcp/service-account.json)" \
  MCP_AUTH_TOKENS="pick-a-bearer-token" \
  --app langcare-mcp-dev
```

**What happens at container startup:**

1. `docker-entrypoint.sh` sees `GCP_CREDENTIALS_B64` env var (from Fly secrets)
2. Decodes it → writes `/tmp/keys/gcp-credentials.json` (exists only in container memory)
3. Sets `GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/gcp-credentials.json` as fallback
4. Reads `CONFIG_FILE` env var → `/app/config.fly.gcp.yaml`
5. Starts `mcp-server -config /app/config.fly.gcp.yaml -http`
6. `config.Load()` reads the YAML, `os.ExpandEnv()` replaces `${GCP_PROJECT_ID}`, `${GCP_LOCATION}`, `${GCP_DATASET_ID}`, `${GCP_FHIR_STORE_ID}` with Fly secret values
7. GCP provider reads `/tmp/keys/gcp-credentials.json`, authenticates with Google Cloud Healthcare API
8. Server listens on `:8080`, Fly routes `https://langcare-mcp-dev.fly.dev` to it

---

#### Cerner (Oracle Health)

**Set provider in `fly/fly.dev.toml`:**

```toml
[env]
  LOG_LEVEL = "debug"
  CONFIG_FILE = "/app/config.fly.cerner.yaml"
```

**Set secrets:**

```bash
fly secrets set \
  CERNER_BASE_URL="https://fhir-ehr-code.cerner.com/r4/<your-tenant-id>" \
  CERNER_CLIENT_ID="your-cerner-client-id" \
  CERNER_CLIENT_SECRET="your-cerner-client-secret" \
  CERNER_TOKEN_URL="https://authorization.cerner.com/tenants/<your-tenant-id>/protocols/oauth2/profiles/smart-v1/token" \
  MCP_AUTH_TOKENS="pick-a-bearer-token" \
  --app langcare-mcp-dev
```

**What happens at container startup:**

1. Reads `CONFIG_FILE` env var → `/app/config.fly.cerner.yaml`
2. Starts `mcp-server -config /app/config.fly.cerner.yaml -http`
3. `config.Load()` reads the YAML, `os.ExpandEnv()` replaces `${CERNER_CLIENT_ID}`, `${CERNER_CLIENT_SECRET}`, `${CERNER_BASE_URL}`, `${CERNER_TOKEN_URL}` with Fly secret values
4. Cerner provider uses OAuth2 Client Credentials flow (client_id + client_secret) — no key files needed
5. Server listens on `:8080`, Fly routes `https://langcare-mcp-dev.fly.dev` to it

---

### 4. Deploy

```bash
fly deploy -c fly/fly.dev.toml --app langcare-mcp-dev
```

This builds the Docker image on Fly's builders (your keys never leave your machine — only the base64-encoded secrets set in step 3 are sent, encrypted).

### 5. Verify

```bash
curl https://langcare-mcp-dev.fly.dev/health
fly logs --app langcare-mcp-dev
fly status --app langcare-mcp-dev
```

### 6. Connect

The server uses **Streamable HTTP** transport on `/mcp`. Any MCP client that supports remote HTTP servers can connect directly.

**Any MCP client:**

```
URL:   https://langcare-mcp-dev.fly.dev/mcp
Auth:  Authorization: Bearer pick-a-bearer-token
Tools: fhir_read, fhir_search, fhir_create, fhir_update (auto-discovered)
```

**Claude Desktop** (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "langcare-fhir": {
      "url": "https://langcare-mcp-dev.fly.dev/mcp",
      "headers": {
        "Authorization": "Bearer pick-a-bearer-token"
      }
    }
  }
}
```

**Clients without native Streamable HTTP support** — use `mcp-remote` as a stdio-to-HTTP bridge:

```json
{
  "mcpServers": {
    "langcare-fhir": {
      "command": "npx",
      "args": [
        "mcp-remote",
        "https://langcare-mcp-dev.fly.dev/mcp",
        "--header",
        "Authorization: Bearer pick-a-bearer-token"
      ]
    }
  }
}
```

## Security

- Private keys and credentials **never** enter git or the Docker image
- Keys are base64-encoded Fly secrets, decoded to `/tmp/keys/` at container startup only
- `.gitignore` blocks `*.pem`, `keys/`, `.env`
- `.dockerignore` blocks `*.pem`, `keys/`, `.env`
- TLS terminated by Fly.io edge (`force_https = true`)
- Non-root container user (`mcpuser`, UID 1001)

## Operations

```bash
fly logs --app langcare-mcp-dev           # View logs
fly ssh console --app langcare-mcp-dev    # SSH into container
fly status --app langcare-mcp-dev         # App status
fly secrets list --app langcare-mcp-dev   # List set secrets
fly releases --app langcare-mcp-dev       # List releases (rollback)

# Rotate MCP auth token
fly secrets set MCP_AUTH_TOKENS="new-token" --app langcare-mcp-dev
```

## Secrets Reference

| Secret | Provider | Description |
|--------|----------|-------------|
| `MCP_AUTH_TOKENS` | All | Bearer token for MCP client auth |
| `EPIC_BASE_URL` | EPIC | FHIR R4 endpoint URL |
| `EPIC_CLIENT_ID` | EPIC | App Orchard client ID |
| `EPIC_TOKEN_URL` | EPIC | OAuth2 token endpoint |
| `EPIC_PRIVATE_KEY_B64` | EPIC | Base64-encoded RSA private key PEM |
| `GCP_PROJECT_ID` | GCP | Google Cloud project ID |
| `GCP_LOCATION` | GCP | Dataset region (e.g. us-central1) |
| `GCP_DATASET_ID` | GCP | Healthcare dataset ID |
| `GCP_FHIR_STORE_ID` | GCP | FHIR store ID |
| `GCP_CREDENTIALS_B64` | GCP | Base64-encoded service account JSON |
| `CERNER_BASE_URL` | Cerner | FHIR R4 endpoint URL (includes tenant ID) |
| `CERNER_CLIENT_ID` | Cerner | OAuth2 client ID |
| `CERNER_CLIENT_SECRET` | Cerner | OAuth2 client secret |
| `CERNER_TOKEN_URL` | Cerner | OAuth2 token endpoint (includes tenant ID) |

## Key Storage Security on Fly.io

Each Fly.io container gets its own isolated filesystem. `/tmp` is ephemeral — it's destroyed when the machine stops or redeploys. No other machine or user can access it.

The entrypoint writes decoded keys to `/tmp/keys/` with `chmod 600` and runs as `mcpuser` (non-root, UID 1001).

`fly secrets` is Fly.io's built-in secrets manager. Secrets are encrypted at rest, decrypted only when the machine starts, and injected as env vars. They are never visible in logs, deploy output, or `fly.dev.toml`. `fly secrets list` shows names but never values. Rotating is just `fly secrets set KEY="new-value"` which triggers a redeploy.

### Attack Surface

| Risk | Status |
|------|--------|
| Keys in git | Blocked by `.gitignore` |
| Keys in Docker image | Blocked by `.dockerignore` + not in COPY |
| Keys on container disk | `/tmp` is ephemeral, destroyed on stop |
| Other processes reading keys | `chmod 600` + single-user container |
| Keys surviving redeploy | No — new container, new `/tmp` |

For higher-compliance environments, a `tmpfs` mount guarantees keys live in RAM only and never touch disk, even if the container's filesystem is backed by a disk volume.
