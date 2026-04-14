#!/usr/bin/env bash
# Delete all skills whose display_title matches a given prefix
# Usage: ./delete-skills.sh <prefix>
# Example: ./delete-skills.sh lc-
set -euo pipefail

PREFIX="${1:?Usage: delete-skills.sh <prefix>}"

echo "Fetching skills with prefix: $PREFIX"
SKILLS=$(curl -sS "https://api.anthropic.com/v1/skills?source=custom&limit=100" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -H "anthropic-beta: skills-2025-10-02")

SKILL_IDS=$(echo "$SKILLS" | python3 -c "
import json, sys
d = json.load(sys.stdin)
matched = [s['id'] for s in d['data'] if s['display_title'].startswith('$PREFIX')]
print('\n'.join(matched))
")

COUNT=$(echo "$SKILL_IDS" | grep -c . || true)
echo "Found $COUNT skill(s) to delete"
echo ""

while IFS= read -r SKILL_ID; do
  [[ -z "$SKILL_ID" ]] && continue

  TITLE=$(echo "$SKILLS" | python3 -c "
import json,sys
d=json.load(sys.stdin)
s=next(x for x in d['data'] if x['id']=='$SKILL_ID')
print(s['display_title'])
")

  echo "Deleting: $TITLE ($SKILL_ID)"

  # Get all version timestamps
  VERSIONS=$(curl -sS "https://api.anthropic.com/v1/skills/$SKILL_ID/versions" \
    -H "x-api-key: $ANTHROPIC_API_KEY" \
    -H "anthropic-version: 2023-06-01" \
    -H "anthropic-beta: skills-2025-10-02" | \
    python3 -c "import json,sys; print('\n'.join(v['version'] for v in json.load(sys.stdin)['data']))")

  # Delete each version
  while IFS= read -r VERSION; do
    [[ -z "$VERSION" ]] && continue
    curl -sS -X DELETE "https://api.anthropic.com/v1/skills/$SKILL_ID/versions/$VERSION" \
      -H "x-api-key: $ANTHROPIC_API_KEY" \
      -H "anthropic-version: 2023-06-01" \
      -H "anthropic-beta: skills-2025-10-02" >/dev/null
    echo "  deleted version $VERSION"
  done <<< "$VERSIONS"

  # Delete the skill
  curl -sS -X DELETE "https://api.anthropic.com/v1/skills/$SKILL_ID" \
    -H "x-api-key: $ANTHROPIC_API_KEY" \
    -H "anthropic-version: 2023-06-01" \
    -H "anthropic-beta: skills-2025-10-02" >/dev/null
  echo "  deleted skill"

done <<< "$SKILL_IDS"

echo ""
echo "Done."
