# Local Development and Testing Guide

Complete guide for setting up, running, and testing the LangCare MCP FHIR Server locally with EPIC credentials.

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Setup Steps](#setup-steps)
3. [Configuration Options](#configuration-options)
4. [Running Locally](#running-locally)
5. [Testing with Claude Desktop](#testing-with-claude-desktop)
6. [Testing with MCP Inspector](#testing-with-mcp-inspector)
7. [Test Cases](#test-cases)
8. [Troubleshooting](#troubleshooting)
9. [Test Automation](#test-automation)
10. [Performance Testing](#performance-testing)

---

## Prerequisites

Before you start, you need:

1. **Claude Desktop installed** (for testing)
2. **EPIC Client ID** - Get from EPIC App Orchard
3. **Private Key** - Generated RSA private key (2048-bit)
4. **Public Key Registered** - Your public key must be registered with EPIC (see [EPIC-APP-SECURITY.md](EPIC-APP-SECURITY.md))

---

## Setup Steps

### Step 1: Generate Keys (First Time Only)

If you haven't generated keys yet:

```bash
# Create keys directory
mkdir -p keys

# Generate private key
openssl genrsa -out keys/private-key.pem 2048

# Generate public key
openssl rsa -in keys/private-key.pem -pubout -out keys/public-key.pem

# Set proper permissions (important!)
chmod 600 keys/private-key.pem
chmod 644 keys/public-key.pem
```

### Step 2: Create JWKS and Register with EPIC

```bash
# Generate JWKS from public key
cd scripts
./create_jwks.sh

# This creates jwks.json - upload to EPIC App Orchard
```

**Register with EPIC**:
1. Go to EPIC App Orchard: https://apporchard.epic.com/
2. Navigate to your app → Authentication
3. Upload `jwks.json` or provide JWKS URL
4. Note your **Client ID**

### Step 3: Configure Application

Edit `configs/config.yaml`:

```yaml
fhir_server:
  # Set provider to "epic"
  provider: "epic"

  # EPIC FHIR base URL
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"

  # EPIC OAuth2 configuration
  epic:
    # Your client ID from EPIC App Orchard
    client_id: "abc123-your-client-id"

    # Path to your private key (relative or absolute)
    private_key_path: "./keys/private-key.pem"

    # EPIC token endpoint
    token_url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"

    # Scopes you registered with EPIC
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"
      - "system/MedicationRequest.read"
      - "system/Condition.read"
      - "system/Encounter.read"

  timeout: 30s

# Keep stdio for local testing (no MCP auth needed)
transport:
  stdio: true
  http:
    enabled: false

logging:
  level: "debug"  # Use debug level for local testing
  format: "json"
  scrub_phi: true
```

### Step 4: Directory Structure

Your project should look like:

```
langcare-mcp-fhir/
├── configs/
│   └── config.yaml          # ← EPIC config here
├── keys/                    # ← Create this directory
│   ├── private-key.pem      # ← Your private key (600 permissions)
│   └── public-key.pem       # ← Your public key (for reference)
│   └── jwks.json            # ← Generated JWKS (upload to EPIC)
├── scripts/     
│   └── create_jwks.sh
└── bin/
    └── langcare-mcp-fhir        # ← Compiled binary
```

---

## Running Locally

### Build the Application

```bash
# Build
go build -o bin/langcare-mcp-fhir ./cmd/server

# Or use make
make build
```

### Start the Server

```bash
# Run with stdio transport (for local testing)
./bin/langcare-mcp-fhir

# Or specify config file explicitly
./bin/langcare-mcp-fhir -config configs/config.local.yaml
```

### Expected Output

```
=== LangCare MCP FHIR Server ===
Name: LangCare MCP FHIR Server
Version: 2.0.0
FHIR Server: https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4
Auth Type: epic
PHI Scrubbing: true
============================

2026/01/17 16:45:00 INFO creating fhir client provider=epic base_url=https://fhir.epic.com/...
2026/01/17 16:45:00 INFO initializing epic provider base_url=https://fhir.epic.com/...
2026/01/17 16:45:00 INFO refreshing epic access token
2026/01/17 16:45:01 INFO epic access token refreshed expires_in=3600
2026/01/17 16:45:01 INFO epic provider authenticated successfully
[FHIR] Provider: epic
[FHIR] Base URL: https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4
[AUDIT] Audit logging enabled (PHI scrubbing: true)
[TOOLS] Registered 4 tools
  - fhir_read: Read a FHIR resource by type and ID
  - fhir_search: Search FHIR resources with query parameters
  - fhir_create: Create a new FHIR resource
  - fhir_update: Update an existing FHIR resource
[MCP] Server initialized

[TRANSPORT] Starting stdio transport...
[SECURITY] Stdio mode: No MCP authentication (process-level isolation only)

Starting MCP server with stdio transport...
```

---

## Configuration Options

### Option 1: Hardcode in config.yaml (Local Only)

**✅ Good for**: Local development, testing

```yaml
epic:
  client_id: "abc123-your-client-id"
  private_key_path: "./keys/private-key.pem"
```

**⚠️ Warning**: Never commit credentials to git!

### Option 2: Environment Variables (Better)

**✅ Good for**: Local development with better security

```yaml
epic:
  client_id: "${EPIC_CLIENT_ID}"
  private_key_path: "${EPIC_PRIVATE_KEY_PATH}"
```

Then set environment variables:

```bash
export EPIC_CLIENT_ID="abc123-your-client-id"
export EPIC_PRIVATE_KEY_PATH="./keys/private-key.pem"

./bin/langcare-mcp-fhir
```

### Option 3: Separate Config File (Best for Local)

Create `configs/config.local.yaml` (add to .gitignore):

```yaml
fhir_server:
  provider: "epic"
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"
  epic:
    client_id: "abc123-your-client-id"
    private_key_path: "./keys/private-key.pem"
    token_url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"

transport:
  stdio: true

logging:
  level: "debug"
```

Run with:

```bash
./bin/langcare-mcp-fhir -config configs/config.local.yaml
```

---

## Testing with Claude Desktop

### Step 1: Configure Claude Desktop

**Location**:
- macOS/Linux: `~/.config/claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

**Configuration**:

```json
{
  "mcpServers": {
    "fhir-epic": {
      "command": "/Users/yourname/langcare-mcp-fhir/bin/langcare-mcp-fhir",
      "args": ["-config", "/Users/yourname/langcare-mcp-fhir/configs/config.local.yaml"]
    }
  }
}
```

**⚠️ Important Notes**:
- Use **absolute paths**, not relative paths
- Server name (`fhir-epic`) can be anything you want
- No `env` needed for stdio transport (server inherits environment)

### Step 2: Start Claude Desktop

1. Save the config file
2. **Restart Claude Desktop** (important!)
3. Look for the 🔌 (plug) icon in the bottom-right
4. Click it to see connected MCP servers
5. Verify `fhir-epic` appears in the list

### Step 3: Test FHIR Operations

#### Test 1: List Available Tools

```
What FHIR tools do you have available?
```

**Expected Response**:
```
I have access to 4 FHIR tools:
1. fhir_read - Read a FHIR resource by type and ID
2. fhir_search - Search FHIR resources with query parameters
3. fhir_create - Create a new FHIR resource
4. fhir_update - Update an existing FHIR resource
```

#### Test 2: Search for Patients

```
Can you search for patients with last name "Smith"?
```

**What Claude will do**:
1. Call `fhir_search` tool with:
   ```json
   {
     "resourceType": "Patient",
     "queryParams": "family=Smith"
   }
   ```
2. Parse the Bundle response
3. Summarize the results

#### Test 3: Read Specific Patient

```
Can you read Patient resource with ID "12345"?
```

**What Claude will do**:
1. Call `fhir_read` tool with:
   ```json
   {
     "resourceType": "Patient",
     "id": "12345"
   }
   ```
2. Display patient information

#### Test 4: Search Observations

```
Can you find vital signs observations for patient "12345"
from the last 7 days?
```

**What Claude will do**:
1. Calculate date range (7 days ago to now)
2. Call `fhir_search` tool with:
   ```json
   {
     "resourceType": "Observation",
     "queryParams": "patient=12345&category=vital-signs&date=ge2024-01-10"
   }
   ```
3. Parse and present the observations

### Step 4: Verify Logs

In your terminal where you can view logs:

```bash
# Check Claude Desktop logs
# macOS
tail -f ~/Library/Logs/Claude/mcp*.log

# Linux
tail -f ~/.config/Claude/logs/mcp*.log
```

**Expected Log Entries**:
```
[AUDIT] PHI_ACCESS timestamp=2026-01-17T... operation=search resource_type=Patient status=success
[AUDIT] PHI_ACCESS timestamp=2026-01-17T... operation=read resource_type=Patient resource_id=12345 status=success
```

### Debugging Claude Desktop

**If server doesn't appear**:

1. **Check config syntax**:
   ```bash
   # Validate JSON
   cat ~/.config/claude/claude_desktop_config.json | jq .
   ```

2. **Check paths are absolute**:
   ```bash
   # Test command manually
   /full/path/to/bin/langcare-mcp-fhir -config /full/path/to/configs/config.yaml
   ```

3. **Check Claude Desktop logs**:
   ```bash
   # macOS
   tail -f ~/Library/Logs/Claude/mcp*.log

   # Linux
   tail -f ~/.config/Claude/logs/mcp*.log
   ```

4. **Restart Claude Desktop completely**:
   - Quit Claude Desktop (Cmd+Q on Mac)
   - Wait 5 seconds
   - Reopen Claude Desktop

---

## Testing with MCP Inspector

The MCP Inspector is the official testing tool from Anthropic for debugging MCP servers. It provides a visual interface for testing MCP tools and is the **recommended approach** for HTTP transport testing.

### Install Inspector

```bash
# Clear cache and run latest version
npm cache clean --force
npx @modelcontextprotocol/inspector@latest
```

### Test stdio Transport

For local development with stdio transport:

```bash
npx @modelcontextprotocol/inspector \
  /path/to/langcare-mcp-fhir/bin/langcare-mcp-fhir \
  -config /path/to/langcare-mcp-fhir/configs/config.local.yaml
```

**What you'll see**:
- A web browser opens with the Inspector UI
- Your 4 FHIR tools are listed
- You can click each tool to see its schema

### Test HTTP Transport with EPIC

For testing HTTP/SSE transport with EPIC FHIR server:

#### Step 1: Configure HTTP Transport

Ensure your `config.local.yaml` has HTTP enabled:

```yaml
fhir_server:
  provider: "epic"
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"
  epic:
    client_id: "your-client-id"
    private_key_path: "./keys/private-key.pem"
    token_url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
    scopes:
      - "system/Patient.read"
      # ... additional scopes

transport:
  stdio: false
  http:
    enabled: true
    port: 8080

security:
  mcp_auth_tokens: "token1"  # Or use environment variable
```

#### Step 2: Start Server with Authentication

```bash
# Set auth token via environment variable
export MCP_AUTH_TOKENS="token1,token2,token3"

# Start server
./bin/langcare-mcp-fhir -config ./configs/config.local.yaml
```

#### Step 3: Connect MCP Inspector

1. **Launch Inspector**:
   ```bash
   npm cache clean --force
   npx @modelcontextprotocol/inspector@latest
   ```

2. **Configure Connection** in the Inspector UI:
   - **Transport Type**: SSE
   - **URL**: `http://localhost:8080/mcp`
   - **Connection Type**: Via proxy
   - **Authentication Headers**:
     - Header Name: `Authorization`
     - Header Value: `Bearer token1`

3. **Click Connect**

#### Step 4: Test FHIR Operations

Once connected, you can test the FHIR tools:

1. **List Tools**: Click the "Tools" tab to see all 4 available tools

2. **Test Patient Search**:
   - Select the `fhir_search` tool
   - Fill in parameters:
     ```json
     {
       "resourceType": "Patient",
       "queryParams": "family=Lin&given=Derrick&birthdate=1973-06-03"
     }
     ```
   - Click "Call Tool"
   - View the response with patient data

3. **Test Patient Read**:
   - Select the `fhir_read` tool
   - Fill in parameters:
     ```json
     {
       "resourceType": "Patient",
       "id": "erXuFYUfucBZaryVksYEcMg3"
     }
     ```
   - Click "Call Tool"

### Using the Inspector UI

1. **View Server Info**:
   - Server name and version
   - Available capabilities
   - Connection status

2. **List Tools**:
   - Click "Tools" tab
   - See all 4 FHIR tools with complete schemas
   - View required/optional parameters

3. **Test Tool Calls**:
   - Select any tool (fhir_read, fhir_search, fhir_create, fhir_update)
   - Fill in parameters using the visual form
   - Click "Call Tool"
   - See formatted request and response
   - View request/response times

4. **View Logs**:
   - All tool calls are logged in the Inspector
   - Debug errors with full stack traces
   - Monitor performance metrics

### Health Check

To verify the server is running, you can use a simple curl command:

```bash
curl http://localhost:8080/health
```

**Expected**:
```
OK
```

### Inspector Advantages

✅ **Visual tool testing** - No need to write JSON manually
✅ **Schema validation** - Shows required/optional fields with type checking
✅ **Request/response inspection** - Full protocol debugging capabilities
✅ **Authentication support** - Easy header configuration
✅ **Works with all transports** - stdio and HTTP/SSE
✅ **Real-time testing** - Immediate feedback on tool calls

---

## Test Cases

### Basic Functionality Tests

#### Test Case 1: List Tools
**Expected**: 4 tools (fhir_read, fhir_search, fhir_create, fhir_update)

#### Test Case 2: Read Patient
**Input**:
```json
{
  "resourceType": "Patient",
  "id": "example"
}
```
**Expected**: Patient resource JSON

#### Test Case 3: Search with Parameters
**Input**:
```json
{
  "resourceType": "Patient",
  "queryParams": "birthdate=gt1990-01-01&gender=male"
}
```
**Expected**: Bundle with matching patients

#### Test Case 4: Create Resource
**Input**:
```json
{
  "resourceType": "Patient",
  "resource": {
    "resourceType": "Patient",
    "name": [{"family": "Test", "given": ["John"]}]
  }
}
```
**Expected**: Created patient with ID

#### Test Case 5: Update Resource
**Input**:
```json
{
  "resourceType": "Patient",
  "id": "12345",
  "resource": {
    "resourceType": "Patient",
    "id": "12345",
    "name": [{"family": "Updated"}]
  }
}
```
**Expected**: Updated patient resource

### Error Handling Tests

#### Test Case 6: Invalid Resource Type
**Input**:
```json
{
  "resourceType": "InvalidType",
  "id": "12345"
}
```
**Expected**: 404 or OperationOutcome error

#### Test Case 7: Missing Required Fields
**Input**:
```json
{
  "resourceType": "Patient"
}
```
**Expected**: Error about missing "id" field

#### Test Case 8: Invalid Authentication
**Test**: Call API without auth token
**Expected**: 401 Unauthorized

#### Test Case 9: Rate Limiting
**Test**: Send 300 requests rapidly
**Expected**: Some 429 Too Many Requests responses

### EPIC-Specific Tests

#### Test Case 10: Token Refresh
1. Start server with EPIC config
2. Wait 55 minutes
3. Make FHIR request
**Expected**: Token automatically refreshes, request succeeds

#### Test Case 11: Invalid JWT
**Test**: Modify private key, restart server
**Expected**: "Invalid client assertion signature" error

### Audit Logging Tests

#### Test Case 12: PHI Access Logging
**Test**: Read Patient/12345
**Expected Log**:
```
[AUDIT] PHI_ACCESS timestamp=... operation=read resource_type=Patient resource_id=12345 status=success
```

#### Test Case 13: Failed Request Logging
**Test**: Read invalid Patient ID
**Expected Log**:
```
[AUDIT] PHI_ACCESS timestamp=... operation=read resource_type=Patient resource_id=invalid status=error
```

---

## Troubleshooting

### Claude Desktop Issues

#### Server Not Listed
**Symptoms**: No MCP servers appear in Claude Desktop

**Solutions**:
1. Check config file location and syntax
2. Use absolute paths (not relative)
3. Verify binary exists and is executable
4. Restart Claude Desktop completely
5. Check Claude logs: `~/Library/Logs/Claude/mcp*.log`

#### Server Crashes on Startup
**Symptoms**: Server appears briefly then disappears

**Solutions**:
1. Test command manually in terminal
2. Check FHIR server connectivity
3. Verify EPIC credentials (if using)
4. Check for port conflicts (HTTP mode)

#### Tools Not Working
**Symptoms**: Claude says tools are unavailable

**Solutions**:
1. Verify server is running (check logs)
2. Test with MCP Inspector
3. Check FHIR server is accessible
4. Verify authentication works

### MCP Inspector Issues

#### Cannot Connect to Server
**Solutions**:
1. Verify server is running: `curl http://localhost:8080/health`
2. Check port is correct
3. Verify auth token is correct
4. Check firewall settings

#### Tools Show But Don't Work
**Solutions**:
1. Check FHIR server connectivity
2. Verify EPIC/Cerner credentials
3. Check audit logs for errors
4. Test directly with curl

### EPIC Authentication Issues

#### "Failed to load private key"
**Check**:
```bash
# Verify file exists
ls -la keys/private-key.pem

# Verify it's a valid key
openssl rsa -in keys/private-key.pem -text -noout

# Check permissions
chmod 600 keys/private-key.pem
```

#### "Invalid client assertion signature"
**Causes**:
1. Public key not registered with EPIC
2. Wrong client ID in config
3. Private/public key mismatch

**Solutions**:
1. Verify keys match:
   ```bash
   openssl rsa -in keys/private-key.pem -noout -modulus | openssl md5
   openssl rsa -pubin -in keys/public-key.pem -noout -modulus | openssl md5
   # These should match!
   ```
2. Regenerate JWKS and re-upload to EPIC
3. Verify client ID is correct

#### "Invalid audience"
**Cause**: Wrong token URL in config

**Solution**: Must be EPIC's token endpoint:
```yaml
token_url: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"
# NOT the FHIR base URL!
```

#### "JWT expired"
**Cause**: Clock skew or expired JWT

**Solutions**:
1. Sync system clock: `sudo ntpdate -u time.apple.com`
2. Check JWT expiry in logs
3. Verify token auto-refresh is working

### Enable Debug Logging

```yaml
logging:
  level: "debug"  # Shows JWT creation and token requests
  format: "text"  # Easier to read locally
```

---

## Test Automation

### Shell Script for Basic Tests

```bash
#!/bin/bash
# test_mcp.sh - Basic MCP server tests

set -e

SERVER_URL="http://localhost:8080"
AUTH_TOKEN="test-token-123"

echo "Testing MCP Server..."

# Test 1: Health check
echo "1. Health check..."
curl -s "$SERVER_URL/health" | grep -q "OK" && echo "✅ Pass" || echo "❌ Fail"

# Test 2: Authentication required
echo "2. Authentication required..."
STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$SERVER_URL/mcp")
[ "$STATUS" = "401" ] && echo "✅ Pass" || echo "❌ Fail"

# Test 3: Valid auth works
echo "3. Valid authentication..."
STATUS=$(curl -s -o /dev/null -w "%{http_code}" \
  -H "Authorization: Bearer $AUTH_TOKEN" \
  "$SERVER_URL/mcp")
[ "$STATUS" = "200" ] && echo "✅ Pass" || echo "❌ Fail"

# Test 4: List tools
echo "4. List tools..."
RESULT=$(curl -s -X POST "$SERVER_URL/mcp" \
  -H "Authorization: Bearer $AUTH_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}')
echo "$RESULT" | grep -q "fhir_read" && echo "✅ Pass" || echo "❌ Fail"

echo "Tests complete!"
```

### Run Tests

```bash
# Make executable
chmod +x test_mcp.sh

# Start server
export MCP_AUTH_TOKENS="test-token-123"
./bin/langcare-mcp-fhir -http -port 8080 &
SERVER_PID=$!

# Wait for startup
sleep 2

# Run tests
./test_mcp.sh

# Cleanup
kill $SERVER_PID
```

---

## Performance Testing

### Load Testing with Apache Bench

```bash
# Install apache bench
# macOS: brew install httpd
# Ubuntu: apt-get install apache2-utils

# Test health endpoint
ab -n 1000 -c 10 http://localhost:8080/health

# Test with authentication
ab -n 1000 -c 10 \
  -H "Authorization: Bearer test-token-123" \
  http://localhost:8080/mcp
```

### Monitor Performance

```bash
# CPU and memory usage
top -pid $(pgrep langcare-mcp-fhir)

# Request rate
watch -n 1 'grep "\[AUDIT\]" logs.txt | wc -l'

# Token refresh monitoring
grep "refreshing epic access token" logs.txt
```

---

## Security Checklist for Local Testing

- [ ] Private key has 600 permissions (`chmod 600 keys/private-key.pem`)
- [ ] Keys directory is in `.gitignore`
- [ ] Local config file is in `.gitignore` (if using config.local.yaml)
- [ ] Never commit `config.yaml` with credentials
- [ ] Use separate keys for dev/staging/production
- [ ] Rotate test keys every 90 days

---

## Quick Reference

### Common Commands

```bash
# Build
make build

# Run with default config
./bin/langcare-mcp-fhir

# Run with custom config
./bin/langcare-mcp-fhir -config configs/config.local.yaml

# Run with environment variables
export EPIC_CLIENT_ID="..."
export EPIC_PRIVATE_KEY_PATH="./keys/private-key.pem"
./bin/langcare-mcp-fhir

# Check logs for errors
./bin/langcare-mcp-fhir 2>&1 | grep -i error

# Debug mode
./bin/langcare-mcp-fhir 2>&1 | tee debug.log
```

### File Locations

| What | Where |
|------|-------|
| **Config** | `configs/config.local.yaml` |
| **Private Key** | `keys/private-key.pem` |
| **Public Key** | `keys/public-key.pem` |
| **JWKS** | `keys/jwks.json` |
| **Binary** | `bin/langcare-mcp-fhir` |
| **Logs** | stdout/stderr |

---

## Next Steps

After successful local testing:

1. ✅ Test all FHIR operations (read, search, create, update)
2. ✅ Verify audit logging
3. ✅ Test token refresh (wait 55 minutes)
4. ✅ Run automated test suite
5. ✅ Perform load testing
6. 📋 Deploy to staging environment
7. 📋 Configure production credentials
8. 📋 Set up monitoring and alerting
9. 📋 Perform security audit
10. 📋 Deploy to production (see [SECURITY.md](../SECURITY.md))

---

## Additional Resources

- [EPIC App Security Setup](EPIC-APP-SECURITY.md) - Complete EPIC authentication guide
- [Main Security Documentation](../SECURITY.md) - Production deployment and HIPAA compliance
- [MCP Inspector Docs](https://github.com/modelcontextprotocol/inspector) - Official Anthropic testing tool
- [EPIC Documentation](https://fhir.epic.com/Documentation?docId=oauth2) - Official EPIC OAuth2 docs
