#!/usr/bin/env bash
# Deploy a single CMA agent, resolving skill name placeholders from skills-registry.json
# Usage: ./deploy-agent.sh <agent-yaml-file> [dev|staging|prod]
set -euo pipefail

AGENT_FILE="${1:?Usage: deploy-agent.sh <agent-yaml-file> [tier]}"
TIER="${2:-dev}"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CMA_DIR="$(dirname "$SCRIPT_DIR")"
REGISTRY_FILE="$CMA_DIR/skills-registry.json"

if [[ ! -f "$AGENT_FILE" ]]; then
  echo "Error: Agent file not found: $AGENT_FILE" >&2; exit 1
fi

if [[ ! -f "$REGISTRY_FILE" ]]; then
  echo "Error: skills-registry.json not found — run upload-skills.sh first" >&2; exit 1
fi

case "$TIER" in
  dev)     MCP_URL="https://langcare-mcp-dev.fly.dev/mcp" ;;
  staging) MCP_URL="https://langcare-mcp-staging.fly.dev/mcp" ;;
  prod)    MCP_URL="https://langcare-mcp.fly.dev/mcp" ;;
  *)       echo "Error: Unknown tier: $TIER" >&2; exit 1 ;;
esac

# Write payload to temp file to avoid bash variable control-character corruption
TMPFILE=$(mktemp /tmp/agent_payload_XXXXXX.json)
trap "rm -f $TMPFILE" EXIT

# Use Python to parse YAML → JSON, resolve skill IDs, set MCP URL
python3 - "$AGENT_FILE" "$REGISTRY_FILE" "$MCP_URL" > "$TMPFILE" <<'PYEOF'
import sys, json, yaml

agent_file, registry_file, mcp_url = sys.argv[1], sys.argv[2], sys.argv[3]

with open(agent_file) as f:
    agent = yaml.safe_load(f)

with open(registry_file) as f:
    registry = json.load(f)

# Resolve skill name placeholders → real API skill IDs
resolved_skills = []
for skill in agent.get("skills", []):
    name = skill.get("skill_id", "")
    real_id = registry.get(name, {}).get("skill_id")
    if real_id:
        skill = dict(skill, skill_id=real_id)
    else:
        print(f"  WARN: skill '{name}' not found in registry", file=sys.stderr)
    resolved_skills.append(skill)
agent["skills"] = resolved_skills

# Set MCP URL for the langcare server
for srv in agent.get("mcp_servers", []):
    if srv.get("name") == "langcare":
        srv["url"] = mcp_url

# Remove metadata (not accepted by API)
agent.pop("metadata", None)

print(json.dumps(agent))
PYEOF

AGENT_NAME=$(jq -r '.name' "$TMPFILE")
echo "Deploying: $AGENT_NAME (tier: $TIER)"

RESPONSE=$(curl -sS --fail-with-body "https://api.anthropic.com/v1/agents" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: managed-agents-2026-04-01" \
  -H "content-type: application/json" \
  -d @"$TMPFILE")

AGENT_ID=$(echo "$RESPONSE" | jq -r '.id')
VERSION=$(echo "$RESPONSE" | jq -r '.version')
echo "  → $AGENT_ID v$VERSION"
