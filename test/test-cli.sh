#!/bin/bash
# Usage: ./test.sh your-mc-server-url your-token "name=Williams&birthdate=1958-11-08"

curl -s -X POST "$1/mcp" \
  -H "Authorization: Bearer $2" \
  -H "Mcp-Session-Id: 5CN3XECPWLDDEIABSVEO3UM3KN" \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"method\": \"tools/call\",
    \"params\": {
      \"name\": \"fhir_search\",
      \"arguments\": {
        \"resourceType\": \"Patient\",
        \"queryParams\": \"$3\"
      }
    },
    \"id\": 1
  }" 