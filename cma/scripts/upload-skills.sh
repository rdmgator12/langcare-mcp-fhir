#!/usr/bin/env bash
# Upload all CMA skills and generate skills-registry.json
# Skips skills already uploaded (by display_title). Safe to re-run.
# Usage: ./upload-skills.sh
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CMA_DIR="$(dirname "$SCRIPT_DIR")"
SKILLS_DIR="$CMA_DIR/skills"
REGISTRY="$CMA_DIR/skills-registry.json"

if [[ ! -d "$SKILLS_DIR" ]]; then
  echo "Error: Skills directory not found: $SKILLS_DIR" >&2
  exit 1
fi

# Pre-populate registry from skills already in the workspace
echo "Fetching existing skills from API..."
EXISTING=$(curl -sS "https://api.anthropic.com/v1/skills?source=custom&limit=100" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: skills-2025-10-02")

# Build initial registry from existing skills: {display_title: {skill_id: id}}
echo "$EXISTING" | jq '[.data[] | {(.display_title): {skill_id: .id}}] | add // {}' > "$REGISTRY"
EXISTING_COUNT=$(echo "$EXISTING" | jq '.data | length')
echo "  Found $EXISTING_COUNT existing skill(s) in workspace"

echo ""
echo "Uploading skills from: $SKILLS_DIR"
UPLOADED=0
SKIPPED=0

find "$SKILLS_DIR" -name "SKILL.md" -type f | sort | while read -r skill_md; do
  skill_dir=$(dirname "$skill_md")
  skill_name=$(awk '/^---$/{n++; next} n==1 && /^name:/{sub(/^name: */, ""); print; exit}' "$skill_md")

  # Skip if already in registry
  existing_id=$(jq -r --arg n "$skill_name" '.[$n].skill_id // empty' "$REGISTRY")
  if [[ -n "$existing_id" ]]; then
    echo "  SKIP: $skill_name ($existing_id)"
    SKIPPED=$((SKIPPED + 1))
    continue
  fi

  echo "  Uploading: $skill_name"
  result=$("$SCRIPT_DIR/upload-skill.sh" "$skill_dir")
  IFS='|' read -r name skill_id <<< "$result"
  echo "    → $skill_id"

  tmp=$(mktemp)
  jq --arg n "$name" --arg id "$skill_id" '. + {($n): {skill_id: $id}}' "$REGISTRY" > "$tmp"
  mv "$tmp" "$REGISTRY"
  UPLOADED=$((UPLOADED + 1))
done

echo ""
echo "Done. Registry: $REGISTRY"
jq 'keys | length' "$REGISTRY" | xargs -I{} echo "  {} skills registered"
