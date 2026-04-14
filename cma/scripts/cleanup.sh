#!/usr/bin/env bash
# LangCare CMA — Full Cleanup Script
# Deletes all LangCare agents, skills, environments, vaults, and credentials
# from the Anthropic workspace. Also removes local generated state files.
#
# Usage:
#   ./cleanup.sh [--yes]
#
# Without --yes, prints what would be deleted and asks for confirmation.
# With --yes, skips confirmation (for CI / scripted use).
#
# WARNING: This is irreversible. Run setup.sh to recreate everything.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CMA_DIR="$(dirname "$SCRIPT_DIR")"
AUTO_YES="${1:-}"

if [[ -z "${ANTHROPIC_API_KEY:-}" ]]; then
  echo "Error: ANTHROPIC_API_KEY is not set" >&2; exit 1
fi

AGENTS_BETA="managed-agents-2026-04-01"
AGENT_DEL_BETA="agent-api-2026-03-01"
SKILLS_BETA="skills-2025-10-02"

# Helpers — no method arg, use -X explicitly at call site
agents_get()   { curl -sS "$1" -H "x-api-key: $ANTHROPIC_API_KEY" -H "anthropic-version: 2023-06-01" -H "anthropic-beta: $AGENTS_BETA"; }
agents_delete(){ curl -sS -X DELETE "$1" -H "x-api-key: $ANTHROPIC_API_KEY" -H "anthropic-version: 2023-06-01" -H "anthropic-beta: $AGENTS_BETA"; }
agent_delete() { curl -sS -X DELETE "$1" -H "x-api-key: $ANTHROPIC_API_KEY" -H "anthropic-version: 2023-06-01" -H "anthropic-beta: $AGENT_DEL_BETA"; }
skills_get()   { curl -sS "$1" -H "x-api-key: $ANTHROPIC_API_KEY" -H "anthropic-version: 2023-06-01" -H "anthropic-beta: $SKILLS_BETA"; }
skills_delete(){ curl -sS -X DELETE "$1" -H "x-api-key: $ANTHROPIC_API_KEY" -H "anthropic-version: 2023-06-01" -H "anthropic-beta: $SKILLS_BETA"; }

# ── Discover what exists ──────────────────────────────────────────────────────

echo "Scanning workspace..."

LANGCARE_AGENTS=$(agents_get "https://api.anthropic.com/v1/agents?limit=100" | python3 -c "
import json,sys
for a in json.load(sys.stdin).get('data', []):
    if 'LangCare' in a.get('name',''):
        print(a['id'] + '|' + a['name'])
")

LANGCARE_SKILLS=$(skills_get "https://api.anthropic.com/v1/skills?source=custom&limit=100" | python3 -c "
import json,sys
for s in json.load(sys.stdin).get('data', []):
    if s.get('display_title','').startswith('langcare-'):
        print(s['id'] + '|' + s['display_title'])
")

LANGCARE_ENVS=$(agents_get "https://api.anthropic.com/v1/environments?limit=100" | python3 -c "
import json,sys
for e in json.load(sys.stdin).get('data', []):
    if 'langcare' in e.get('name','').lower():
        print(e['id'] + '|' + e['name'])
")

LANGCARE_VAULTS=$(agents_get "https://api.anthropic.com/v1/vaults?limit=100" | python3 -c "
import json,sys
for v in json.load(sys.stdin).get('data', []):
    if 'langcare' in v.get('display_name','').lower():
        print(v['id'] + '|' + v['display_name'])
")

REGISTRY_FILE="$CMA_DIR/skills-registry.json"

# ── Preview ───────────────────────────────────────────────────────────────────

echo ""
echo "========================================"
echo "  LangCare CMA — Cleanup Preview"
echo "========================================"

echo ""
echo "Agents:"
if [[ -n "$LANGCARE_AGENTS" ]]; then
  while IFS='|' read -r id name; do echo "  $name ($id)"; done <<< "$LANGCARE_AGENTS"
else echo "  (none)"; fi

echo ""
echo "Skills:"
if [[ -n "$LANGCARE_SKILLS" ]]; then
  while IFS='|' read -r id name; do echo "  $name ($id)"; done <<< "$LANGCARE_SKILLS"
else echo "  (none)"; fi

echo ""
echo "Environments:"
if [[ -n "$LANGCARE_ENVS" ]]; then
  while IFS='|' read -r id name; do echo "  $name ($id)"; done <<< "$LANGCARE_ENVS"
else echo "  (none)"; fi

echo ""
echo "Vaults (+ credentials):"
if [[ -n "$LANGCARE_VAULTS" ]]; then
  while IFS='|' read -r id name; do echo "  $name ($id)"; done <<< "$LANGCARE_VAULTS"
else echo "  (none)"; fi

echo ""
echo "Local state files:"
for f in "$CMA_DIR"/.setup-state-*.json; do [[ -f "$f" ]] && echo "  $f"; done
[[ -f "$REGISTRY_FILE" ]] && echo "  $REGISTRY_FILE"

echo ""

# ── Confirm ───────────────────────────────────────────────────────────────────

if [[ "$AUTO_YES" != "--yes" ]]; then
  read -r -p "Delete all of the above? [y/N] " CONFIRM
  [[ "$CONFIRM" =~ ^[Yy]$ ]] || { echo "Aborted."; exit 0; }
  echo ""
fi

# ── Delete Agents ─────────────────────────────────────────────────────────────

if [[ -n "$LANGCARE_AGENTS" ]]; then
  echo "Deleting agents..."
  while IFS='|' read -r id name; do
    agent_delete "https://api.anthropic.com/v1/agents/$id" >/dev/null
    echo "  deleted: $name"
  done <<< "$LANGCARE_AGENTS"
fi

# ── Delete Skills ─────────────────────────────────────────────────────────────

if [[ -n "$LANGCARE_SKILLS" ]]; then
  echo "Deleting skills..."
  while IFS='|' read -r id name; do
    # Delete all versions first (required by API)
    VERSIONS=$(skills_get "https://api.anthropic.com/v1/skills/$id/versions" | \
      python3 -c "import json,sys; print('\n'.join(v['version'] for v in json.load(sys.stdin)['data']))")
    while IFS= read -r ver; do
      [[ -z "$ver" ]] && continue
      skills_delete "https://api.anthropic.com/v1/skills/$id/versions/$ver" >/dev/null
    done <<< "$VERSIONS"
    skills_delete "https://api.anthropic.com/v1/skills/$id" >/dev/null
    echo "  deleted: $name"
  done <<< "$LANGCARE_SKILLS"
fi

# ── Delete Environments ───────────────────────────────────────────────────────

if [[ -n "$LANGCARE_ENVS" ]]; then
  echo "Deleting environments..."
  while IFS='|' read -r id name; do
    agents_delete "https://api.anthropic.com/v1/environments/$id" >/dev/null
    echo "  deleted: $name"
  done <<< "$LANGCARE_ENVS"
fi

# ── Delete Vaults + Credentials ───────────────────────────────────────────────

if [[ -n "$LANGCARE_VAULTS" ]]; then
  echo "Deleting vaults..."
  while IFS='|' read -r id name; do
    # Delete credentials first
    CREDS=$(agents_get "https://api.anthropic.com/v1/vaults/$id/credentials" | \
      python3 -c "import json,sys; print('\n'.join(c['id'] for c in json.load(sys.stdin)['data']))")
    while IFS= read -r cred_id; do
      [[ -z "$cred_id" ]] && continue
      agents_delete "https://api.anthropic.com/v1/vaults/$id/credentials/$cred_id" >/dev/null
      echo "  deleted credential: $cred_id"
    done <<< "$CREDS"
    agents_delete "https://api.anthropic.com/v1/vaults/$id" >/dev/null
    echo "  deleted: $name"
  done <<< "$LANGCARE_VAULTS"
fi

# ── Remove local state files ──────────────────────────────────────────────────

echo "Removing local state files..."
for f in "$CMA_DIR"/.setup-state-*.json; do
  [[ -f "$f" ]] && { rm -f "$f"; echo "  removed: $(basename "$f")"; }
done
[[ -f "$REGISTRY_FILE" ]] && { rm -f "$REGISTRY_FILE"; echo "  removed: skills-registry.json"; }

# ── Done ──────────────────────────────────────────────────────────────────────

echo ""
echo "Cleanup complete. Run ./setup.sh to recreate everything."
