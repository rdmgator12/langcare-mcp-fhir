#!/usr/bin/env bash
# Upload a single CMA skill directory to the Skills API
# Usage: ./upload-skill.sh <skill-directory>
# Output: name|skill_id  (pipe-delimited, for use by upload-skills.sh)
set -euo pipefail

SKILL_DIR="$1"

if [[ ! -f "$SKILL_DIR/SKILL.md" ]]; then
  echo "Error: No SKILL.md in $SKILL_DIR" >&2; exit 1
fi

# Extract skill name from YAML frontmatter
SKILL_NAME=$(awk '/^---$/{n++; next} n==1 && /^name:/{sub(/^name: */, ""); print; exit}' "$SKILL_DIR/SKILL.md")
if [[ -z "$SKILL_NAME" ]]; then
  echo "Error: No name in SKILL.md frontmatter" >&2; exit 1
fi

# Root dir in uploaded filenames must match the skill name from frontmatter
ROOT_DIR="$SKILL_NAME"

# Build -F file args for every file in the skill directory
FILE_ARGS=()
while IFS= read -r -d '' file; do
  rel_path="$ROOT_DIR/${file#$SKILL_DIR/}"
  FILE_ARGS+=(-F "files[]=@$file;filename=$rel_path")
done < <(find "$SKILL_DIR" -type f -print0)

RESPONSE=$(curl -sS --fail-with-body -X POST "https://api.anthropic.com/v1/skills" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: skills-2025-10-02" \
  -F "display_title=$SKILL_NAME" \
  "${FILE_ARGS[@]}")

SKILL_ID=$(echo "$RESPONSE" | jq -r '.id')
echo "$SKILL_NAME|$SKILL_ID"
