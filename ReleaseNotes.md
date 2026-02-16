# LangCare MCP FHIR Server - Release Highlights

## v2.2.0 — Remote HTTP Server Deployment

Deploy LangCare MCP FHIR Server as a **remote Streamable HTTP MCP server** accessible by any AI agent from anywhere. Includes a reference implementation for Fly.io with full provider support for EPIC, Cerner, and GCP Healthcare API.

### Remote Streamable HTTP Transport

- Switched HTTP transport from legacy SSE to **Streamable HTTP** (`NewStreamableHTTPHandler`) — standard HTTP POST/GET on `/mcp`, no long-lived connections, works reliably behind reverse proxies and edge networks
- Any MCP-compatible client (Claude Desktop, GPT, Gemini, custom agents) can connect via HTTPS with Bearer token auth
- stdio transport unchanged — local Claude Desktop usage continues to work as before

### Fly.io Reference Deployment

Complete deployment setup under `fly/` — Docker multi-stage build, config-per-provider, secrets management:

- **Multi-stage Dockerfile** — ~15MB Alpine image with Go binary, config YAMLs, and entrypoint
- **Provider configs** — `config.fly.epic.yaml`, `config.fly.cerner.yaml`, `config.fly.gcp.yaml` with `${ENV_VAR}` placeholders expanded at startup via `os.ExpandEnv()`
- **Secrets management** — credentials stored as encrypted Fly secrets, base64-encoded key files (EPIC PEM, GCP JSON) decoded to ephemeral `/tmp/keys/` at container startup, never in git or Docker image
- **Scale-to-zero** — `auto_stop_machines = "stop"`, ~$0-2/mo for dev
- **Security** — TLS at Fly edge, non-root container user (UID 1001), `chmod 600` on key files, health checks at `/health`

### Provider Support

| Provider | Auth Flow | Key Files |
|----------|-----------|-----------|
| EPIC | JWT assertion (RSA private key) | PEM decoded from `EPIC_PRIVATE_KEY_B64` |
| Cerner | OAuth2 Client Credentials | None — client_id + client_secret as env vars |
| GCP | Service account credentials | JSON decoded from `GCP_CREDENTIALS_B64` |

### Cerner Configuration

- Added `configs/config.cerner.example.yaml` — example config for Cerner (Oracle Health) with OAuth2 Client Credentials flow
- Added `fly/config.fly.cerner.yaml` — Fly.io deployment config for Cerner

### Connecting MCP Clients

Claude Desktop with native Streamable HTTP:

```json
{
  "mcpServers": {
    "langcare-fhir": {
      "url": "https://langcare-mcp-dev.fly.dev/mcp",
      "headers": {
        "Authorization": "Bearer your-token"
      }
    }
  }
}
```

Fallback via `mcp-remote` for clients without native support:

```json
{
  "mcpServers": {
    "langcare-fhir": {
      "command": "npx",
      "args": ["mcp-remote", "https://langcare-mcp-dev.fly.dev/mcp", "--header", "Authorization: Bearer your-token"]
    }
  }
}
```

**Full guide:** [fly/README.md](fly/README.md) | [docs/FLY-DEPLOYMENT.md](docs/FLY-DEPLOYMENT.md)

---

## 🚀 Enterprise-Grade MCP FHIR Server for Agentic AI Platforms

**Fully rewritten in Go** - A production-ready MCP Server that connects AI agents to FHIR-based EMRs with enterprise-grade security and extensible architecture.

### ✨ Key Highlights

- **Generic FHIR Operations** - 4 universal tools (read, search, create, update) work with any FHIR R4 resource type (Patient, Observation, Medication, etc.) - no resource-specific code needed
- **Enterprise Security** - Two-layer authentication (MCP client + FHIR backend), OAuth2 with token refresh, TLS 1.3, PHI scrubbing, HIPAA-compliant audit logging
- **Multi-Backend Support** - EPIC, Cerner, GCP Healthcare API, and any FHIR R4 server with pluggable authentication
- **Stateless Proxy Design** - No persistent storage, minimal attack surface, compliant with healthcare data protection requirements
- **Production Ready** - Streamable HTTP and stdio transports, rate limiting, error sanitization, comprehensive security documentation
- **40+ Clinical Skills Library** - Agent-agnostic workflow guides covering medication management, lab interpretation, clinical decision support, documentation, and more

### 🏗️ Architecture

- **Built entirely in Go** using the official `modelcontextprotocol/go-sdk` (Anthropic/Google maintained)
- **Stateless proxy** between AI agents and FHIR servers - lightweight and scalable
- **Pluggable providers** for EPIC, Cerner, GCP, and custom FHIR backends
- **Kubernetes-ready** with security manifests and environment-based credential management

### 📋 What's Included

- **4 Generic FHIR Tools** - Works with any resource type out of the box
- **Multiple Auth Methods** - Bearer, OAuth2, SMART on FHIR, Basic Auth
- **40 Clinical Skills** - Evidence-based workflow guides across 8 categories (see below)
- **Complete Documentation** - Setup guides, security architecture, EPIC/Cerner integration, local testing
- **Public Test Server** - HAPI FHIR for immediate testing without setup

### 🧠 Clinical Skills Library (New in 2.1.1)

40 agent-agnostic clinical workflow guides organized into 8 categories. Each skill teaches AI agents how to perform complex healthcare tasks using the server's 4 FHIR tools. Copy a skill's `SKILL.md` into your agent's system prompt — works with Claude, ChatGPT, Gemini, or any AI agent.

| Category | Skills | Highlights |
|----------|--------|------------|
| **Patient Data & Summary** | 5 | Demographics, CCD-style clinical summary, problem list audit, allergy review, insurance coverage |
| **Medication Management** | 5 | Med reconciliation (Joint Commission), drug interactions (CYP450), adherence (MPR/PDC), Beers Criteria, opioid risk (ORT/MME) |
| **Lab & Diagnostics** | 5 | Lab interpretation, critical values (CAP/CLIA), pre-op labs, diabetes panel (ADA), renal function (KDIGO) |
| **Clinical Decision Support** | 5 | Sepsis (qSOFA/SOFA), cardiovascular risk (ASCVD/HEART), VTE (Wells/Caprini), fall risk (Morse), pneumonia (CURB-65) |
| **Care Coordination** | 5 | Discharge planning (LACE), referrals, care gaps (USPSTF), transitions of care (I-PASS), follow-up tasks |
| **Documentation** | 5 | SOAP notes, H&P, progress notes, discharge summaries, procedure notes |
| **Population Health** | 5 | Panel overview, quality measures (HEDIS), chronic disease registries, immunization status (CDC), preventive care compliance |
| **Specialty** | 5 | Prenatal (ACOG), pediatric growth (WHO/CDC), mental health (PHQ-9/GAD-7), oncology (TNM/RECIST), chronic pain |

Each skill includes detailed `references/` files with scoring criteria, code tables (LOINC, SNOMED, RxNorm, ICD-10), and society guideline citations.

**Full catalog:** [skills/README.md](https://github.com/langcare/langcare-mcp-fhir/blob/main/skills/README.md)

### 🔐 Security Features

✅ TLS 1.3 encryption
✅ PHI scrubbing in logs
✅ OAuth 2.0 with automatic token refresh
✅ HIPAA-compliant audit logging
✅ mTLS support for Kubernetes
✅ No persistent PHI storage

### 🎯 Perfect For

- Healthcare teams building AI-powered clinical workflows
- Agentic platforms integrating with EMRs
- Organizations needing HIPAA-compliant FHIR connectivity
- Teams requiring extensible, resource-agnostic solutions

---

**Get Started:** See [README](https://github.com/langcare/langcare-mcp-fhir#quick-start) for build instructions, [Local Testing Guide](docs/LOCAL-TESTING.md) for setup, and [Security Docs](docs/SECURITY.md) for production deployment.
