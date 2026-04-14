#!/usr/bin/env bash
# LangCare CMA — Full Setup Script
# Uploads skills, creates environment + vault, deploys agents.
# Safe to re-run: skips already-uploaded skills and existing resources.
#
# Usage:
#   export ANTHROPIC_API_KEY=sk-ant-...
#   export LANGCARE_MCP_URL=https://langcare-mcp-dev.fly.dev/mcp
#   export LANGCARE_MCP_TOKEN=your-bearer-token
#   ./setup.sh [dev|staging|prod]
#
# After running, note the ENVIRONMENT_ID and VAULT_ID printed at the end.
# Use them in run-session.sh to start agent sessions.

set -euo pipefail

TIER="${1:-dev}"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CMA_DIR="$(dirname "$SCRIPT_DIR")"

# ── Validation ────────────────────────────────────────────────────────────────

if [[ -z "${ANTHROPIC_API_KEY:-}" ]]; then
  echo "Error: ANTHROPIC_API_KEY is not set" >&2; exit 1
fi
if [[ -z "${LANGCARE_MCP_URL:-}" ]]; then
  echo "Error: LANGCARE_MCP_URL is not set (e.g. https://langcare-mcp-dev.fly.dev/mcp)" >&2; exit 1
fi
if [[ -z "${LANGCARE_MCP_TOKEN:-}" ]]; then
  LANGCARE_MCP_TOKEN="changeme"
  echo "Warning: LANGCARE_MCP_TOKEN not set — using placeholder 'changeme'. Update via the portal before running sessions."
fi

command -v python3 >/dev/null || { echo "Error: python3 required" >&2; exit 1; }
command -v jq     >/dev/null || { echo "Error: jq required" >&2; exit 1; }
python3 -c "import yaml" 2>/dev/null || { echo "Error: PyYAML required — run: pip install pyyaml" >&2; exit 1; }

BETA="managed-agents-2026-04-01"
API="https://api.anthropic.com"

api() {
  local method="$1" path="$2" body="${3:-}"
  if [[ -n "$body" ]]; then
    curl -sS --fail-with-body -X "$method" "$API$path" \
      -H "x-api-key: $ANTHROPIC_API_KEY" \
      -H "anthropic-version: 2023-06-01" \
      -H "anthropic-beta: $BETA" \
      -H "content-type: application/json" \
      -d "$body"
  else
    curl -sS --fail-with-body -X "$method" "$API$path" \
      -H "x-api-key: $ANTHROPIC_API_KEY" \
      -H "anthropic-version: 2023-06-01" \
      -H "anthropic-beta: $BETA"
  fi
}

STATE_FILE="$CMA_DIR/.setup-state-$TIER.json"
# Load existing state if present
if [[ -f "$STATE_FILE" ]]; then
  ENVIRONMENT_ID=$(python3 -c "import json; d=json.load(open('$STATE_FILE')); print(d.get('environment_id',''))" 2>/dev/null || echo "")
  VAULT_ID=$(python3 -c "import json; d=json.load(open('$STATE_FILE')); print(d.get('vault_id',''))" 2>/dev/null || echo "")
else
  ENVIRONMENT_ID=""
  VAULT_ID=""
fi

echo "========================================"
echo "  LangCare CMA Setup — tier: $TIER"
echo "========================================"
echo ""

# ── Step 1: Upload Skills ─────────────────────────────────────────────────────

echo "Step 1/4: Uploading skills..."
"$SCRIPT_DIR/upload-skills.sh"
echo ""

# ── Step 2: Create Environment ────────────────────────────────────────────────

echo "Step 2/4: Environment..."

if [[ -n "$ENVIRONMENT_ID" ]]; then
  echo "  Using existing environment: $ENVIRONMENT_ID"
else
  MCP_HOST=$(python3 -c "from urllib.parse import urlparse; print(urlparse('$LANGCARE_MCP_URL').netloc)")
  ENV_BODY=$(jq -n \
    --arg name "langcare-$TIER" \
    --arg host "$MCP_HOST" \
    '{name: $name, config: {type: "cloud", networking: {type: "limited", allowed_hosts: [$host], allow_mcp_servers: true}}}')
  RESPONSE=$(api POST "/v1/environments" "$ENV_BODY")
  ENVIRONMENT_ID=$(echo "$RESPONSE" | python3 -c "import json,sys; print(json.load(sys.stdin)['id'])")
  echo "  Created environment: $ENVIRONMENT_ID"
fi

# ── Step 3: Create Vault + Credential ────────────────────────────────────────

echo "Step 3/4: Vault..."

if [[ -n "$VAULT_ID" ]]; then
  echo "  Using existing vault: $VAULT_ID"
else
  VAULT_BODY=$(jq -n --arg name "LangCare $TIER" '{display_name: $name}')
  VAULT_RESP=$(api POST "/v1/vaults" "$VAULT_BODY")
  VAULT_ID=$(echo "$VAULT_RESP" | python3 -c "import json,sys; print(json.load(sys.stdin)['id'])")
  echo "  Created vault: $VAULT_ID"

  CRED_BODY=$(jq -n \
    --arg name "LangCare MCP ($TIER)" \
    --arg url "$LANGCARE_MCP_URL" \
    --arg token "$LANGCARE_MCP_TOKEN" \
    '{display_name: $name, auth: {type: "static_bearer", mcp_server_url: $url, token: $token}}')
  CRED_RESP=$(api POST "/v1/vaults/$VAULT_ID/credentials" "$CRED_BODY")
  CRED_ID=$(echo "$CRED_RESP" | python3 -c "import json,sys; print(json.load(sys.stdin)['id'])")
  echo "  Created credential: $CRED_ID"
fi

# Save state
jq -n \
  --arg env "$ENVIRONMENT_ID" \
  --arg vlt "$VAULT_ID" \
  '{environment_id: $env, vault_id: $vlt}' > "$STATE_FILE"

# ── Step 4: Deploy Agents ─────────────────────────────────────────────────────

echo "Step 4/4: Deploying agents..."
"$SCRIPT_DIR/deploy-agents.sh" "$TIER"

# ── Done ──────────────────────────────────────────────────────────────────────

echo ""
echo "========================================"
echo "  Setup complete!"
echo "========================================"
echo ""
echo "  Environment ID : $ENVIRONMENT_ID"
echo "  Vault ID       : $VAULT_ID"
echo ""
echo "  State saved to : $STATE_FILE"
echo ""
echo "  Run a session:"
echo "  ./run-session.sh <agent-id> $ENVIRONMENT_ID $VAULT_ID \"your prompt\""
echo ""
echo "  List agents:"
echo "  ./list-agents.sh"
