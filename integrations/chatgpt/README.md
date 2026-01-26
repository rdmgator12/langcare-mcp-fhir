# ChatGPT Integration for LangCare MCP FHIR Server

**Connect ChatGPT Custom GPTs to LangCare MCP FHIR Server using HTTP/SSE transport.**

## Overview

ChatGPT Custom GPTs can connect to external APIs through "Actions". This integration runs the MCP server in HTTP mode and exposes it as an OpenAPI-compatible endpoint that ChatGPT can call.

**Best for:**
- Web-based multi-user access
- Shared team FHIR assistants
- Public FHIR exploration tools
- Integration with ChatGPT's web UI

## Prerequisites

- **ChatGPT Plus or Enterprise** subscription (required for Custom GPTs)
- **LangCare MCP FHIR Server** built (`make build`)
- **FHIR Server Credentials** (EPIC, Cerner, GCP, or test server)
- **Public endpoint** (ngrok, Cloud Run, or similar)

## Architecture

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   ChatGPT   │  HTTPS  │  MCP Server  │  OAuth  │  FHIR API   │
│   Custom    │────────▶│   (HTTP/SSE) │────────▶│   (EMR)     │
│    GPT      │         │   + Auth     │         │             │
└─────────────┘         └──────────────┘         └─────────────┘
```

**Key differences from Claude:**
- Uses **HTTP/SSE** transport (not stdio)
- Requires **MCP authentication** (Bearer tokens)
- Needs **public endpoint** (ChatGPT calls from OpenAI servers)
- Supports **multiple concurrent users**

## Setup Instructions

### Step 1: Build and Configure MCP Server

**Create HTTP-mode config:**

```yaml
# configs/config.http.epic.yaml
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
  stdio: false       # Must be false for HTTP mode
  http:
    enabled: true    # Enable HTTP/SSE transport
    port: 8080

security:
  mcp_auth_tokens: "your-secret-token-here,another-token"  # ChatGPT will use this
  rate_limit:
    enabled: true
    rate: 100
    burst: 200

logging:
  scrub_phi: true
```

**Build and run:**

```bash
make build
./bin/langcare-mcp-fhir -config configs/config.http.epic.yaml -http -port 8080
```

**Verify it's running:**

```bash
curl http://localhost:8080/health
# Should return: {"status":"healthy"}
```

### Step 2: Expose to Public Internet

ChatGPT needs a public HTTPS endpoint to connect to your MCP server.

#### Option A: ngrok (Development)

```bash
ngrok http 8080
```

Note the HTTPS URL: `https://abc123.ngrok.io`

**Pros:** Quick setup for testing
**Cons:** URL changes on restart, free tier limits

#### Option B: Cloud Run (Production)

**Deploy to Google Cloud Run:**

```bash
# Build Docker image
docker build -t gcr.io/your-project/langcare-mcp-fhir .

# Push to GCR
docker push gcr.io/your-project/langcare-mcp-fhir

# Deploy to Cloud Run
gcloud run deploy langcare-mcp-fhir \
  --image gcr.io/your-project/langcare-mcp-fhir \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars MCP_AUTH_TOKENS=your-secret-token
```

**Pros:** Stable HTTPS URL, auto-scaling, production-ready
**Cons:** Requires GCP account, costs money

#### Option C: Heroku, Railway, Fly.io

Similar to Cloud Run - follow platform-specific deployment guides.

### Step 3: Create Custom GPT

1. Go to https://chat.openai.com
2. Click your profile → **My GPTs**
3. Click **Create a GPT**
4. Click **Configure**

**GPT Configuration:**

**Name:** `FHIR Medical Assistant`

**Description:**
```
AI assistant with direct access to FHIR R4 electronic medical records. Can search patients, review charts, check lab results, and manage clinical data using standard FHIR operations.
```

**Instructions:**
```
You are a FHIR-enabled medical AI assistant with access to electronic medical record systems through the Model Context Protocol (MCP).

You have 4 FHIR tools available:
- fhir_search: Search for resources with query parameters
- fhir_read: Read a specific resource by type and ID
- fhir_create: Create a new FHIR resource
- fhir_update: Update an existing resource

Always verify patient identity before any operation. Use partial identifiers for searches, then confirm with the user before retrieving full records.

Present clinical data clearly:
- Labs: Include reference ranges and flag abnormal values
- Medications: Show name, dose, frequency, route
- Problems: Note onset dates and clinical status
- Encounters: Include date, type, provider, location

Follow HIPAA best practices. Never share PHI unnecessarily. Log all access appropriately.
```

**Optional: Add Clinical Skills**

To enhance clinical capabilities, copy the entire contents of `skills/core/fhir-clinical/SKILL.md` and append it to the Instructions field above.

**Conversation starters:**
```
- Search for patients named Smith
- Review chart for patient ID 12345
- What are the latest lab results for this patient?
- Record blood pressure 120/80 for patient ID 67890
```

### Step 4: Configure Actions

In the Custom GPT Configure page, scroll to **Actions** section:

**1. Click "Create new action"**

**2. Authentication:**
- Type: **Bearer Token**
- Token: `your-secret-token-here` (from config `mcp_auth_tokens`)

**3. Schema:**

Paste the following OpenAPI schema:

```yaml
openapi: 3.0.0
info:
  title: LangCare MCP FHIR Server
  version: 2.0.0
  description: Model Context Protocol server for FHIR R4 operations

servers:
  - url: https://your-ngrok-or-cloud-url.com
    description: MCP Server

paths:
  /mcp:
    post:
      operationId: callTool
      summary: Execute FHIR operations through MCP
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                method:
                  type: string
                  enum: [tools/call]
                params:
                  type: object
                  properties:
                    name:
                      type: string
                      enum: [fhir_read, fhir_search, fhir_create, fhir_update]
                    arguments:
                      type: object
      responses:
        '200':
          description: Successful response
          content:
            application/json:
              schema:
                type: object

components:
  securitySchemes:
    BearerAuth:
      type: http
      scheme: bearer

security:
  - BearerAuth: []
```

**Important:** Replace `https://your-ngrok-or-cloud-url.com` with your actual public endpoint.

**4. Privacy Policy (Optional)**

If making your GPT public, add a privacy policy URL.

**5. Save and Test**

Click **Save** then **View GPT** to test.

### Step 5: Test the Integration

**Test 1: List Tools**

```
You: What tools do you have?

GPT should respond with: fhir_read, fhir_search, fhir_create, fhir_update
```

**Test 2: Search Patients**

```
You: Search for patients with last name Lopez

GPT should use fhir_search tool and return matching patients
```

**Test 3: Read Patient**

```
You: Read patient with ID eg2bapHUHrr1geTcBYqPJVGFjV-9YEpNwI5pT5RBl6ikB

GPT should use fhir_read tool and return patient demographics
```

## Troubleshooting

### GPT Can't Connect to MCP Server

**Symptom:** "Error calling action" or timeout

**Solutions:**
1. Verify MCP server is running: `curl https://your-url.com/health`
2. Check ngrok/cloud service is active
3. Ensure Bearer token matches between config and GPT Action
4. Check CORS headers (MCP server should allow OpenAI origins)
5. Review MCP server logs for authentication errors

### Authentication Errors

**Symptom:** "401 Unauthorized"

**Solutions:**
1. Verify Bearer token is correct in GPT Action settings
2. Check `security.mcp_auth_tokens` in server config
3. Ensure token is passed in Authorization header
4. Test manually:
   ```bash
   curl -H "Authorization: Bearer your-token" https://your-url.com/mcp
   ```

### FHIR Operations Failing

**Symptom:** GPT gets errors from FHIR server

**Solutions:**
1. Check FHIR server credentials in config
2. Verify OAuth scopes for requested operations
3. Test FHIR server access manually:
   ```bash
   go run test/test_epic_token.go "client-id" "/path/to/key.pem"
   ```
4. Review MCP server stderr logs

### Rate Limiting

**Symptom:** "Too many requests" errors

**Solutions:**
1. Increase rate limit in config: `rate: 200, burst: 500`
2. Add more MCP auth tokens (rate limit per token)
3. Implement backoff in GPT instructions
4. Monitor usage and scale infrastructure

## Advanced Configuration

### Multiple FHIR Servers

Create separate Custom GPTs for different FHIR servers:

**GPT 1:** EPIC Production
- Config: `config.http.epic-prod.yaml`
- Port: 8080
- URL: `https://epic-prod.example.com`

**GPT 2:** EPIC Sandbox
- Config: `config.http.epic-sandbox.yaml`
- Port: 8081
- URL: `https://epic-sandbox.example.com`

### Team Deployment

**1. Deploy to Cloud Run / ECS / Kubernetes**

```bash
# Example: Google Cloud Run with secrets
gcloud run deploy langcare-mcp-fhir \
  --image gcr.io/project/langcare-mcp-fhir \
  --set-env-vars MCP_AUTH_TOKENS=token1,token2,token3 \
  --set-secrets EPIC_CLIENT_ID=epic-client-id:latest \
  --set-secrets EPIC_PRIVATE_KEY=epic-private-key:latest
```

**2. Create Team GPT**

- Create Custom GPT as above
- Share with team via "Share" button
- Distribute to workspace (ChatGPT Enterprise only)

**3. Monitor Usage**

- Enable Prometheus metrics (built into MCP server)
- Track: requests/sec, error rate, FHIR latency
- Alert on authentication failures or rate limit hits

### TLS / HTTPS

For production, always use HTTPS:

**Option 1: Cloud provider handles TLS**
- Cloud Run, Heroku, Railway auto-provision TLS

**Option 2: Reverse proxy**
```bash
# nginx with Let's Encrypt
server {
  listen 443 ssl;
  server_name fhir-mcp.example.com;

  ssl_certificate /etc/letsencrypt/live/fhir-mcp.example.com/fullchain.pem;
  ssl_certificate_key /etc/letsencrypt/live/fhir-mcp.example.com/privkey.pem;

  location / {
    proxy_pass http://localhost:8080;
    proxy_set_header Authorization $http_authorization;
  }
}
```

## Security Considerations

### MCP Authentication

**Critical:** Always use strong, unique Bearer tokens.

```bash
# Generate secure tokens
openssl rand -hex 32
```

Store in environment variables, never commit to git:

```bash
export MCP_AUTH_TOKENS="$(cat secrets/mcp-tokens.txt)"
./bin/langcare-mcp-fhir -config config.yaml -http
```

### Network Security

1. **Firewall:** Only allow inbound HTTPS (443) or HTTP (8080)
2. **Rate Limiting:** Prevent abuse (100 req/s default)
3. **Audit Logging:** Log all PHI access with request ID
4. **TLS:** Always use HTTPS in production

### HIPAA Compliance

For HIPAA-compliant deployments:
- ✅ Use HTTPS/TLS 1.3
- ✅ Enable audit logging (`logging.scrub_phi: true`)
- ✅ Store credentials in secrets manager (GCP Secret Manager, AWS Secrets Manager)
- ✅ Encrypt data at rest (cloud provider defaults)
- ✅ Business Associate Agreement (BAA) with cloud provider
- ✅ Access controls (MCP auth tokens, FHIR OAuth scopes)

See [SECURITY.md](../../docs/SECURITY.md) for complete checklist.

### OpenAI Data Usage

**Important:** Understand OpenAI's data usage policies.

- ChatGPT may use conversations for model training (unless opted out)
- For PHI data, use **ChatGPT Enterprise** with:
  - Data processing agreement (DPA)
  - Training opt-out
  - HIPAA compliance (available in Enterprise tier)

**Best practices:**
- Use de-identified data for testing
- Implement access controls at FHIR server level
- Audit all PHI access
- Follow institutional policies

## Examples

### Example 1: Clinical Documentation GPT

**Custom GPT Configuration:**

**Name:** Clinical Documentation Assistant

**Instructions:**
```
You are a clinical documentation assistant with access to FHIR R4 EHR data.

Your role:
1. Help clinicians review patient charts efficiently
2. Generate structured clinical summaries
3. Identify documentation gaps
4. Assist with coding and billing (ICD-10, CPT)

Always:
- Verify patient identity before accessing records
- Present data in clinical order (HPI, ROS, Exam, A&P)
- Flag missing documentation (allergies unknown, medication reconciliation not done)
- Use standard medical terminology

[Optionally append skills/core/fhir-clinical/SKILL.md here]
```

**Use case:**
- Pre-visit chart prep
- Encounter documentation
- Discharge summaries

### Example 2: Preventive Care GPT

**Name:** Preventive Care Tracker

**Instructions:**
```
You are a preventive care specialist focused on population health.

Your role:
1. Identify patients overdue for screenings (cancer, CV, immunizations)
2. Generate outreach lists for quality measures
3. Track HEDIS/MIPS compliance
4. Recommend evidence-based interventions

Use USPSTF guidelines for screening recommendations. Always note the last screening date and next due date.

[Append preventive care section from SKILL.md]
```

**Use case:**
- Annual wellness visits
- Population health campaigns
- Quality reporting

### Example 3: Public FHIR Explorer GPT

**Name:** FHIR Explorer (Public Test Server)

**Configuration:**
- FHIR Server: `https://hapi.fhir.org/baseR4` (no auth)
- No PHI (synthetic test data only)
- Public sharing enabled

**Instructions:**
```
You are a FHIR R4 education assistant using a public test server with synthetic data.

Help users:
- Learn FHIR resource structures
- Practice FHIR searches
- Understand FHIR relationships (references)
- Test FHIR queries before production

This is test data only - no real patients.
```

**Use case:**
- FHIR education
- Onboarding new developers
- Testing queries before production

## Production Deployment Checklist

- [ ] HTTPS endpoint with valid TLS certificate
- [ ] Strong Bearer tokens (32+ chars, random)
- [ ] Rate limiting enabled (`rate: 100, burst: 200`)
- [ ] PHI scrubbing enabled (`scrub_phi: true`)
- [ ] Audit logging configured
- [ ] FHIR OAuth scopes follow least privilege
- [ ] Environment variables for secrets (no hardcoded credentials)
- [ ] Monitoring and alerting (Prometheus, Datadog, etc.)
- [ ] Backup and disaster recovery plan
- [ ] Business Associate Agreement (BAA) with cloud provider
- [ ] ChatGPT Enterprise with DPA and training opt-out
- [ ] Institutional review and approval
- [ ] User training and documentation

## Getting Help

**ChatGPT/OpenAI Issues:**
- OpenAI help: https://help.openai.com
- Custom GPTs docs: https://help.openai.com/en/articles/8554397-creating-a-gpt

**MCP Server Issues:**
- GitHub Issues: https://github.com/langcare/langcare-mcp-fhir/issues
- Security guide: [SECURITY.md](../../docs/SECURITY.md)
- Local testing: [LOCAL-TESTING.md](../../docs/LOCAL-TESTING.md)

**Skills:**
- Skills README: [skills/README.md](../../skills/README.md)
- Contributing: [CONTRIBUTING.md](../../CONTRIBUTING.md)

## Next Steps

1. **Deploy to cloud** - Cloud Run, Heroku, or Railway
2. **Create Custom GPT** - Follow instructions above
3. **Add clinical skills** - Copy SKILL.md to GPT Instructions
4. **Test thoroughly** - Verify all 4 FHIR operations work
5. **Share with team** - Distribute GPT to colleagues
6. **Monitor usage** - Track requests, errors, and FHIR latency
7. **Iterate** - Refine instructions based on user feedback

**Questions?** Open an issue or discussion on GitHub!
