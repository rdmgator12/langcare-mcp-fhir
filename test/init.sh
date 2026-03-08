#!/bin/bash
# Usage: ./init.sh your-mc-server-url your-token

  curl -s -D - -X POST $1/mcp \
  -H "Authorization: Bearer $2" \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "clientInfo": {"name": "langcare-cli", "version": "1.0.0"},
      "capabilities": {}
    },
    "id": 1
  }'