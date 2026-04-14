#!/usr/bin/env bash
# List uploaded CMA skills
# Usage: ./list-skills.sh
set -euo pipefail

curl -sS "https://api.anthropic.com/v1/skills" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: managed-agents-2026-04-01" \
  | jq -r '.data[] | "\(.id)\t\(.skill_id)\tv\(.version)"' \
  | column -t -s $'\t'
