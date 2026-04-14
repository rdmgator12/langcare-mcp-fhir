#!/usr/bin/env bash
# Validate all CMA skills and agent YAML configs
# Usage: ./validate.sh
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CMA_DIR="$(dirname "$SCRIPT_DIR")"
ERRORS=0

echo "Validating CMA configuration"
echo "============================="

# Validate agent YAMLs
echo ""
echo "Agents:"
for f in "$CMA_DIR"/agents/*.yaml; do
  name=$(basename "$f")
  if ! yq e '.' "$f" > /dev/null 2>&1; then
    echo "  FAIL: $name (invalid YAML)"
    ERRORS=$((ERRORS + 1))
    continue
  fi

  # Check required fields
  for field in name model.id description system mcp_servers tools skills metadata; do
    if [[ "$(yq e ".$field" "$f")" == "null" ]]; then
      echo "  FAIL: $name (missing .$field)"
      ERRORS=$((ERRORS + 1))
    fi
  done

  SKILL_COUNT=$(yq e '.skills | length' "$f")
  if [[ "$SKILL_COUNT" -lt 1 ]]; then
    echo "  FAIL: $name (no skills defined)"
    ERRORS=$((ERRORS + 1))
  else
    echo "  OK:   $name ($SKILL_COUNT skills)"
  fi
done

# Validate environment YAMLs
echo ""
echo "Environments:"
for f in "$CMA_DIR"/environments/*.yaml; do
  name=$(basename "$f")
  if ! yq e '.' "$f" > /dev/null 2>&1; then
    echo "  FAIL: $name (invalid YAML)"
    ERRORS=$((ERRORS + 1))
    continue
  fi
  echo "  OK:   $name"
done

# Validate vault YAMLs
echo ""
echo "Vaults:"
for f in "$CMA_DIR"/vaults/*.yaml; do
  name=$(basename "$f")
  if ! yq e '.' "$f" > /dev/null 2>&1; then
    echo "  FAIL: $name (invalid YAML)"
    ERRORS=$((ERRORS + 1))
    continue
  fi
  echo "  OK:   $name"
done

# Validate MCP server registry
echo ""
echo "MCP Servers:"
REGISTRY="$CMA_DIR/mcp-servers/registry.yaml"
if [[ -f "$REGISTRY" ]]; then
  if ! yq e '.' "$REGISTRY" > /dev/null 2>&1; then
    echo "  FAIL: registry.yaml (invalid YAML)"
    ERRORS=$((ERRORS + 1))
  else
    echo "  OK:   registry.yaml"
  fi
else
  echo "  FAIL: registry.yaml not found"
  ERRORS=$((ERRORS + 1))
fi

echo ""
echo "============================="
if [[ $ERRORS -gt 0 ]]; then
  echo "FAILED: $ERRORS error(s) found."
  exit 1
fi
echo "PASSED: All configs valid."
