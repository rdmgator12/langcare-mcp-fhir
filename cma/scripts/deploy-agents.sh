#!/usr/bin/env bash
# Deploy all CMA agents. Creates new agents or updates existing ones by name.
# Usage: ./deploy-agents.sh [tier]
set -euo pipefail

TIER="${1:-dev}"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
AGENTS_DIR="$(dirname "$SCRIPT_DIR")/agents"

# Fetch existing agents so we can update instead of duplicate
echo "Fetching existing agents..."
EXISTING=$(curl -sS "https://api.anthropic.com/v1/agents?limit=100" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: managed-agents-2026-04-01")

echo "Deploying all agents (tier: $TIER)"
echo "=================================="

FAILED=0
for agent_file in "$AGENTS_DIR"/*.yaml; do
  AGENT_NAME=$(python3 -c "import yaml,sys; d=yaml.safe_load(open(sys.argv[1])); print(d['name'])" "$agent_file")
  AGENT_ID=$(echo "$EXISTING" | jq -r --arg name "$AGENT_NAME" '.data[] | select(.name == $name) | .id' | head -1)

  if [[ -n "$AGENT_ID" ]]; then
    # Agent exists — update it
    if ! "$SCRIPT_DIR/update-agent.sh" "$AGENT_ID" "$agent_file" "$TIER"; then
      echo "FAILED: $agent_file" >&2
      FAILED=$((FAILED + 1))
    fi
  else
    # New agent — create it
    if ! "$SCRIPT_DIR/deploy-agent.sh" "$agent_file" "$TIER"; then
      echo "FAILED: $agent_file" >&2
      FAILED=$((FAILED + 1))
    fi
  fi
  echo ""
done

if [[ $FAILED -gt 0 ]]; then
  echo "$FAILED agent(s) failed." >&2
  exit 1
fi

echo "All agents deployed successfully."
