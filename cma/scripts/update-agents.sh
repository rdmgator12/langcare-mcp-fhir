#!/usr/bin/env bash
# Update all 9 deployed CMA agents to the latest YAML definition
# Agent IDs are fetched live from the API by matching agent names
# Usage: ./update-agents.sh [dev|staging|prod]
set -euo pipefail

TIER="${1:-dev}"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
AGENTS_DIR="$(dirname "$SCRIPT_DIR")/agents"

echo "Fetching current agent list from API..."
EXISTING=$(curl -sS "https://api.anthropic.com/v1/agents?limit=100" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: managed-agents-2026-04-01")

echo "Updating all agents (tier: $TIER)"
echo "=================================="

FAILED=0
for agent_file in "$AGENTS_DIR"/*.yaml; do
  # Extract agent name from YAML
  AGENT_NAME=$(python3 -c "import yaml,sys; d=yaml.safe_load(open(sys.argv[1])); print(d['name'])" "$agent_file")

  # Look up existing agent ID by name
  AGENT_ID=$(echo "$EXISTING" | jq -r --arg name "$AGENT_NAME" '.data[] | select(.name == $name) | .id' | head -1)

  if [[ -z "$AGENT_ID" ]]; then
    echo "WARN: No existing agent found for '$AGENT_NAME' — deploying as new"
    if ! "$SCRIPT_DIR/deploy-agent.sh" "$agent_file" "$TIER"; then
      echo "FAILED: $agent_file" >&2
      FAILED=$((FAILED + 1))
    fi
  else
    if ! "$SCRIPT_DIR/update-agent.sh" "$AGENT_ID" "$agent_file" "$TIER"; then
      echo "FAILED: $agent_file ($AGENT_ID)" >&2
      FAILED=$((FAILED + 1))
    fi
  fi
  echo ""
done

if [[ $FAILED -gt 0 ]]; then
  echo "$FAILED agent(s) failed to update." >&2
  exit 1
fi

echo "All agents updated successfully."
