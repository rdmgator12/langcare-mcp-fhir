# Claude Integration for LangCare MCP FHIR Server

**Connect Claude Desktop to LangCare MCP FHIR Server using stdio transport.**

## Overview

Claude Desktop (by Anthropic) supports MCP servers natively through stdio transport. This is the **simplest integration** - no HTTP servers, no authentication middleware, just a local process-to-process connection.

**Best for:**
- Local development and testing
- Single-user deployments
- Personal FHIR exploration
- Privacy-sensitive workflows (no network exposure)

## Prerequisites

- **Claude Desktop** installed (Mac/Windows/Linux)
- **LangCare MCP FHIR Server** built (`make build`)
- **FHIR Server Credentials** (EPIC, Cerner, GCP, or test server)
- **Config file** for your FHIR provider (see `configs/` directory)

## Setup Instructions

### Step 1: Build the MCP Server

```bash
cd /path/to/langcare-mcp-fhir
make build
```

This creates the binary at `bin/langcare-mcp-fhir`.

### Step 2: Create Configuration File

Copy the example config for your FHIR provider:

**For EPIC:**
```bash
cp configs/config.epic.example.yaml configs/config.local.epic.yaml
```

**Edit the config:**
```yaml
fhir_server:
  provider: "epic"
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"
  epic:
    client_id: "your-client-id"
    private_key_path: "/path/to/private-key.pem"
    scopes:
      - "system/Patient.read"
      - "system/Observation.read"
      # ... add more scopes as needed

transport:
  stdio: true        # IMPORTANT: Must be true for Claude Desktop
  http:
    enabled: false   # Must be false for stdio mode
```

**For testing (no auth):**
```bash
cp configs/config.yaml configs/config.local.yaml
```

Uses public HAPI FHIR test server - works immediately, no credentials needed.

### Step 3: Configure Claude Desktop

**Location of config file:**
- **Mac:** `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows:** `%APPDATA%\Claude\claude_desktop_config.json`
- **Linux:** `~/.config/Claude/claude_desktop_config.json`

**Add MCP server entry:**
```json
{
  "mcpServers": {
    "langcare-mcp-fhir": {
      "command": "/absolute/path/to/langcare-mcp-fhir/bin/langcare-mcp-fhir",
      "args": [
        "-config",
        "/absolute/path/to/langcare-mcp-fhir/configs/config.local.epic.yaml"
      ]
    }
  }
}
```

**Important:**
- Use **absolute paths** (no `~` or relative paths)
- Ensure `transport.stdio: true` in your config file
- The config file must exist and be valid YAML

### Step 4: Restart Claude Desktop

Fully quit and restart Claude Desktop to load the new MCP server configuration.

### Step 5: Verify Connection

In Claude Desktop, check that the MCP server is connected:
- Look for the tool icon or MCP server indicator
- Type: "What tools do you have access to?"
- Claude should list: `fhir_read`, `fhir_search`, `fhir_create`, `fhir_update`

## Testing the Integration

### Basic Test: Search for Patients

```
You: Search for patients with last name Smith

Expected: Claude uses fhir_search tool with:
- resourceType: "Patient"
- queryParams: "family=Smith"

Result: List of matching patients
```

### Read Test: Get Patient Details

```
You: Read patient with ID 12345

Expected: Claude uses fhir_read tool with:
- resourceType: "Patient"
- id: "12345"

Result: Complete patient resource with demographics
```

### Create Test: Record Vital Signs

```
You: Record blood pressure 120/80 for patient ID 12345

Expected: Claude uses fhir_create tool with:
- resourceType: "Observation"
- resource: { BP observation with components }

Result: Confirmation of created resource
```

## Adding Clinical Skills (Optional)

Skills are **optional enhancements** that teach Claude structured clinical workflows.

### Option 1: Claude Projects (Recommended)

1. Create a new **Project** in Claude Desktop
2. Add **Project Knowledge** → Upload file
3. Upload `skills/core/fhir-clinical/SKILL.md` from this repository
4. Claude will now reference clinical workflows automatically in this project

**Pros:**
- Isolated to specific project
- Easy to enable/disable
- Can mix multiple skills

### Option 2: Custom Instructions (Global)

1. Open Claude Desktop **Settings**
2. Navigate to **Custom Instructions**
3. Copy contents of `skills/core/fhir-clinical/SKILL.md`
4. Paste into Custom Instructions field
5. Save

**Pros:**
- Available in all conversations
- Persistent across sessions

**Cons:**
- Character limit (~10,000 characters)
- Applies globally (can't disable per-conversation)

### Testing with Skills

If you added the fhir-clinical skill, test a complete workflow:

```
You: Review the chart for patient Derrick Lin

Expected: Claude performs comprehensive chart review:
1. Searches for patient by name
2. Retrieves active conditions
3. Gets current medications
4. Reviews recent lab results
5. Checks recent encounters
6. Lists allergies

Result: Structured clinical summary with all sections
```

## Troubleshooting

### MCP Server Not Connecting

**Symptom:** Claude doesn't show FHIR tools

**Solutions:**
1. Check Claude Desktop logs:
   - Mac: `~/Library/Logs/Claude/`
   - Windows: `%APPDATA%\Claude\logs\`
   - Linux: `~/.config/Claude/logs/`

2. Verify config file syntax:
   ```bash
   cat ~/.config/Claude/claude_desktop_config.json | python -m json.tool
   ```

3. Test MCP server manually:
   ```bash
   ./bin/langcare-mcp-fhir -config configs/config.local.epic.yaml
   # Should start without errors
   ```

4. Ensure paths are absolute (no `~`, `/Users/you/...` not `~/...`)

### FHIR Authentication Errors

**Symptom:** "Authentication failed" or "401 Unauthorized"

**Solutions:**
1. Verify FHIR server credentials in config file
2. Check OAuth scopes match required operations
3. Test credentials directly:
   ```bash
   go run test/test_epic_token.go "your-client-id" "/path/to/private-key.pem"
   ```
4. See [EPIC-APP-SECURITY.md](../../docs/EPIC-APP-SECURITY.md) for EPIC setup

### Tool Execution Errors

**Symptom:** Claude tries to use tool but gets an error

**Solutions:**
1. Check MCP server logs (stderr output)
2. Verify resource type exists in FHIR server
3. Ensure search parameters are valid for that resource
4. Test query manually with curl or Postman

### Skills Not Working

**Symptom:** Claude doesn't follow clinical workflows

**Solutions:**
1. Verify skill was added to Project Knowledge or Custom Instructions
2. Check character limits (Custom Instructions max ~10k chars)
3. Try explicitly mentioning the workflow: "Use the patient chart review workflow"
4. Ensure MCP server connection works first (skills require tools)

## Configuration Examples

### Example 1: EPIC with Full Clinical Scopes

```yaml
fhir_server:
  provider: "epic"
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"
  epic:
    client_id: "085e800e-401a-4303-9613-0dabec0f84c5"
    private_key_path: "/Users/you/keys/epic-private-key.pem"
    scopes:
      - "system/Patient.read"
      - "system/Patient.write"
      - "system/Observation.read"
      - "system/Observation.write"
      - "system/Condition.read"
      - "system/MedicationStatement.read"
      - "system/Encounter.read"
      - "system/AllergyIntolerance.read"

transport:
  stdio: true
  http:
    enabled: false

logging:
  scrub_phi: true
```

### Example 2: HAPI Test Server (No Auth)

```yaml
fhir_server:
  provider: "base"
  base_url: "https://hapi.fhir.org/baseR4"
  auth:
    type: "none"

transport:
  stdio: true
  http:
    enabled: false

logging:
  scrub_phi: true
```

### Example 3: Multiple MCP Servers

```json
{
  "mcpServers": {
    "langcare-fhir-epic": {
      "command": "/path/to/bin/langcare-mcp-fhir",
      "args": ["-config", "/path/to/config.epic.yaml"]
    },
    "langcare-fhir-test": {
      "command": "/path/to/bin/langcare-mcp-fhir",
      "args": ["-config", "/path/to/config.test.yaml"]
    }
  }
}
```

Claude will have access to both FHIR servers simultaneously (different tool prefixes).

## Advanced Usage

### Environment Variables

Override config values with environment variables:

```json
{
  "mcpServers": {
    "langcare-mcp-fhir": {
      "command": "/path/to/bin/langcare-mcp-fhir",
      "args": ["-config", "/path/to/config.yaml"],
      "env": {
        "FHIR_BASE_URL": "https://fhir.example.com/R4",
        "EPIC_CLIENT_ID": "your-client-id",
        "EPIC_PRIVATE_KEY_PATH": "/path/to/key.pem"
      }
    }
  }
}
```

### Debug Mode

Enable verbose logging:

```json
{
  "mcpServers": {
    "langcare-mcp-fhir": {
      "command": "/path/to/bin/langcare-mcp-fhir",
      "args": ["-config", "/path/to/config.yaml"],
      "env": {
        "LOG_LEVEL": "debug"
      }
    }
  }
}
```

Check logs in `~/Library/Logs/Claude/mcp-langcare-mcp-fhir.log` (Mac).

### Custom Skills

Create your own clinical skill:

1. Write `my-custom-skill.md` with FHIR workflows
2. Add to Claude Project Knowledge
3. Test workflows with patients

See [skills/README.md](../../skills/README.md) for skill development guidelines.

## Security Considerations

### Stdio Transport Security

**Process-level isolation:**
- MCP server runs as subprocess of Claude Desktop
- No network exposure
- Trust boundary at OS process level

**Credential management:**
- Store private keys in secure location (`~/keys/`, `~/.ssh/`)
- Set file permissions: `chmod 600 private-key.pem`
- Never commit credentials to git

**PHI protection:**
- Enable `scrub_phi: true` in config
- Audit logs hash resource IDs
- No persistent PHI storage (stateless proxy)

### Best Practices

1. **Separate configs for prod/dev**
   - Use test server for development
   - Production config separate file
   - Never share production credentials

2. **OAuth scopes principle of least privilege**
   - Only request scopes you need
   - Use read-only scopes when possible
   - Review [EPIC-SCOPES.md](../../docs/EPIC-SCOPES.md)

3. **HIPAA compliance**
   - Ensure Claude Desktop on encrypted disk
   - Use organization-approved devices
   - Follow institutional policies
   - See [SECURITY.md](../../docs/SECURITY.md)

## Production Considerations

**Claude Desktop is designed for single-user development.**

For **multi-user production** deployments:
- Use HTTP/SSE transport instead of stdio
- Deploy MCP server to cloud (Kubernetes)
- Implement MCP client authentication
- Enable rate limiting and monitoring
- See [integrations/chatgpt/](../chatgpt/) for HTTP-based setup pattern

## Getting Help

**Claude Desktop Issues:**
- Anthropic support: https://support.anthropic.com
- Claude Desktop docs: https://docs.anthropic.com/claude/docs

**MCP Server Issues:**
- GitHub Issues: https://github.com/langcare/langcare-mcp-fhir/issues
- Local testing guide: [LOCAL-TESTING.md](../../docs/LOCAL-TESTING.md)
- Security guide: [SECURITY.md](../../docs/SECURITY.md)

**Skill Development:**
- Skills README: [skills/README.md](../../skills/README.md)
- Contributing guide: [CONTRIBUTING.md](../../CONTRIBUTING.md)

## Examples

### Complete Example: EPIC Patient Chart Review

**1. Configure MCP server with EPIC credentials**

```yaml
# config.local.epic.yaml
fhir_server:
  provider: "epic"
  base_url: "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4"
  epic:
    client_id: "your-client-id"
    private_key_path: "/Users/you/keys/epic-key.pem"
    scopes:
      - "system/Patient.read"
      - "system/Condition.read"
      - "system/Observation.read"
      - "system/MedicationStatement.read"
      - "system/Encounter.read"
      - "system/AllergyIntolerance.read"
transport:
  stdio: true
  http:
    enabled: false
```

**2. Add to Claude Desktop config**

```json
{
  "mcpServers": {
    "langcare-epic": {
      "command": "/Users/you/langcare-mcp-fhir/bin/langcare-mcp-fhir",
      "args": ["-config", "/Users/you/langcare-mcp-fhir/configs/config.local.epic.yaml"]
    }
  }
}
```

**3. Add clinical skill to Claude Project**

- Create new Project: "EPIC Clinical Review"
- Add Project Knowledge: Upload `skills/core/fhir-clinical/SKILL.md`

**4. Test workflow**

```
You: Review the chart for patient Derrick Lin

Claude:
I'll perform a comprehensive chart review for patient Derrick Lin.

[Uses fhir_search for Patient with name="Derrick Lin"]
[Uses fhir_search for Condition, MedicationStatement, Observation, Encounter, AllergyIntolerance]

**Chart Review for Derrick Lin**

**Demographics:**
- Name: Derrick Lin
- DOB: 1940-10-01 (83 years old)
- MRN: 12345
- Sex: Male

**Active Problems:**
1. Type 2 Diabetes Mellitus (Onset: 2010-03-15)
2. Essential Hypertension (Onset: 2008-07-20)
3. Hyperlipidemia (Onset: 2012-11-10)

**Current Medications:**
1. Metformin 1000mg PO BID
2. Lisinopril 20mg PO Daily
3. Atorvastatin 40mg PO Daily

**Allergies:**
- Penicillin (Rash)

**Recent Lab Results (Last 30 days):**
- HbA1c: 7.2% (2024-01-10) - Above target
- LDL: 95 mg/dL (2024-01-10) - At goal
- Creatinine: 1.1 mg/dL (2024-01-10) - Normal

**Recent Encounters:**
- 2024-01-10: Office Visit - Dr. Smith (Routine diabetes follow-up)
- 2023-10-15: Office Visit - Dr. Smith (Annual wellness visit)
```

## Next Steps

1. **Test basic functionality** - Search, read operations
2. **Add clinical skills** - Copy SKILL.md to Project Knowledge
3. **Test clinical workflows** - Patient chart review, med reconciliation
4. **Explore advanced features** - Multiple FHIR servers, custom skills
5. **Contribute** - Share your skills or improvements with the community

**Questions?** Open an issue or discussion on GitHub!
