# LangCare MCP FHIR Server

[![GitHub Stars](https://img.shields.io/github/stars/langcare/langcare-mcp-fhir?style=social)](https://github.com/langcare/langcare-mcp-fhir)
[![Contributors](https://img.shields.io/github/contributors/langcare/langcare-mcp-fhir)](https://github.com/langcare/langcare-mcp-fhir/graphs/contributors)
[![License](https://img.shields.io/github/license/langcare/langcare-mcp-fhir)](https://github.com/langcare/langcare-mcp-fhir/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/langcare/langcare-mcp-fhir)](https://github.com/langcare/langcare-mcp-fhir/blob/main/go.mod)

Enterprise-grade MCP Server for FHIR-based EMRs, designed for robust deployments in agentic AI platforms. Fully written in Go with enterprise-grade security and generic FHIR operations that work with any FHIR R4 resource type.

## Installation

Install via npm:

```bash
npm install -g @langcare/langcare-mcp-fhir
```

Or use directly without installation:

```bash
npx @langcare/langcare-mcp-fhir -config /path/to/config.yaml
```

## Quick Configuration

LangCare MCP FHIR connects Claude to your FHIR-based EMR system. You need a YAML configuration file pointing to your backend.

### 1. Get a Config Template

Choose your backend:

- **EPIC:** [config.epic.example.yaml](https://github.com/langcare/langcare-mcp-fhir/blob/main/configs/config.epic.example.yaml)
- **Cerner:** [config.cerner.example.yaml](https://github.com/langcare/langcare-mcp-fhir/blob/main/configs/config.cerner.example.yaml)
- **GCP Healthcare API:** [config.gcp.example.yaml](https://github.com/langcare/langcare-mcp-fhir/blob/main/configs/config.gcp.example.yaml)
- **Any FHIR R4 Server:** [config.base.example.yaml](https://github.com/langcare/langcare-mcp-fhir/blob/main/configs/config.base.example.yaml)

### 2. Configure Claude Desktop

Add to your Claude Desktop config file (`~/.config/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "langcare-mcp-fhir": {
      "command": "langcare-mcp-fhir",
      "args": ["-config", "/path/to/your/config.yaml"]
    }
  }
}
```

On macOS, the config is typically at:
```
~/Library/Application\ Support/Claude/claude_desktop_config.json
```

### 3. Restart Claude Desktop

Close and reopen Claude Desktop. The FHIR tools will now be available.

**Need detailed setup help?** See the [Local Testing Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/LOCAL-TESTING.md).

## Architecture

This MCP server acts as an intelligent proxy between AI agents and FHIR R4 servers. It exposes 4 generic FHIR operations through the Model Context Protocol (MCP), enabling AI-powered workflows for any FHIR resource type.

**Key Design:**
- **MCP SDK:** Official `github.com/modelcontextprotocol/go-sdk` (Anthropic/Google maintained)
- **FHIR Client:** Generic HTTP client working with any FHIR R4 server
- **Transport:** stdio and HTTP/SSE
- **Backend:** Proxy to existing FHIR server (no database)
- **Language:** 100% Go for high performance and reliability

## 4 Generic MCP Tools

All tools work with **any FHIR resource type** (Patient, Observation, Medication, etc.):

### 1. fhir_read
Read a FHIR resource by type and ID.

```json
{
  "resourceType": "Patient",
  "id": "example-123"
}
```

### 2. fhir_search
Search FHIR resources with query parameters.

```json
{
  "resourceType": "Patient",
  "queryParams": "name=John&birthdate=gt1990-01-01"
}
```

### 3. fhir_create
Create a new FHIR resource.

```json
{
  "resourceType": "Observation",
  "resource": {
    "resourceType": "Observation",
    "status": "final",
    "code": { ... },
    "subject": { "reference": "Patient/123" }
  }
}
```

### 4. fhir_update
Update an existing FHIR resource.

```json
{
  "resourceType": "Patient",
  "id": "example-123",
  "resource": {
    "resourceType": "Patient",
    "id": "example-123",
    "name": [{ "family": "Smith" }]
  }
}
```

## Security Architecture

LangCare MCP FHIR implements a **two-layer security model** for HIPAA-compliant healthcare data access:

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   Claude    │ Auth1   │  MCP Server  │ Auth2   │  FHIR API   │
│   Client    │────────▶│   (Go)       │────────▶│   (EMR)     │
└─────────────┘         └──────────────┘         └─────────────┘

Auth1: MCP Client Authentication (Bearer Token/API Key)
Auth2: FHIR Backend Authentication (Bearer/OAuth2/SMART on FHIR)
```

### Security Features

- ✅ **TLS 1.3** encryption for HTTP transport
- ✅ **PHI Scrubbing** in logs (enabled by default)
- ✅ **HIPAA-compliant** audit logging
- ✅ **No persistent PHI storage** (stateless proxy)
- ✅ **Secrets via environment variables** (never in config files)
- ✅ **OAuth 2.0** with automatic token refresh
- ✅ **mTLS** support for service-to-service communication
- ✅ **Rate limiting** per client

### Supported Authentication Methods

- **Bearer Token** - Simple API key authentication
- **OAuth2** - Full OAuth2 flow with token refresh
- **SMART on FHIR** - EPIC, Cerner, and other EMR standards
- **Basic Auth** - Username/password authentication
- **Custom** - Extensible for additional auth methods

**For complete security documentation, see [Security Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/SECURITY.md):**
- HIPAA compliance checklist
- OAuth configuration for EPIC/Cerner/GCP
- Kubernetes security manifests
- Credential management procedures
- Audit logging implementation

## Agent Usage

AI agents use LangCare MCP FHIR Server to help healthcare professionals access and manage patient health records through 4 FHIR tools. The server handles EMR authentication, allowing agents to focus on clinical workflows while maintaining strict privacy and accuracy standards.

**Agent capabilities:**
- **Search, Read, Create, Update** - Any FHIR R4 resource (Patient, Observation, Medication, etc.)
- **Patient privacy** - Use partial identifiers, confirm identity before updates
- **Clinical accuracy** - Verify data, use standard codes (LOINC, SNOMED, RxNorm)
- **Professional communication** - Structure responses with context, findings, and next steps

**Common workflows:**
- **Patient lookup:** Search by name/DOB → verify identity → read full details
- **Clinical review:** Retrieve labs, vitals, medications → present with reference ranges
- **Documentation:** Extract structured data → map to FHIR resources → confirm → create
- **Updates:** Verify existing resource → modify → confirm changes → update

**System support:**
- Works with any FHIR R4 resource type (60+ types including DocumentReference, Binary, Media)
- Automatic authentication and token refresh to EPIC, Cerner, GCP Healthcare API
- HIPAA-compliant PHI handling with audit logging
- Comprehensive OAuth2 scopes for clinical data access

**📖 Complete guide:** [Agent Prompt Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/AGENT_PROMPT.md) - System prompt, tool examples, workflows, and error handling

## Clinical Skills Library (Optional)

**40+ agent-agnostic clinical workflow guides** that teach AI agents how to perform complex healthcare tasks using the MCP server's 4 FHIR tools (`fhir_search`, `fhir_read`, `fhir_create`, `fhir_update`).

- **Optional** - The MCP server works without them
- **Portable** - Work with Claude, ChatGPT, Gemini, or any AI agent
- **Evidence-based** - Built on USPSTF, ADA, ACC/AHA, CDC, ACOG, KDIGO, and other society guidelines
- **Copy-paste ready** - Add a skill's `SKILL.md` to your agent's system prompt or custom instructions

### Skill Categories (40 Skills)

| Category | Skills | Examples |
|----------|--------|----------|
| **Patient Data & Summary** | 5 | Demographics, clinical summary (CCD-style), problem list audit, allergy review, insurance coverage |
| **Medication Management** | 5 | Med reconciliation, drug interactions (CYP450), adherence (MPR/PDC), Beers Criteria, opioid risk (ORT/MME) |
| **Lab & Diagnostics** | 5 | Lab interpretation, critical values (CAP/CLIA), pre-op labs, diabetes panel (ADA), renal function (KDIGO) |
| **Clinical Decision Support** | 5 | Sepsis (qSOFA/SOFA), cardiovascular risk (ASCVD/HEART), VTE (Wells/Caprini), fall risk (Morse), pneumonia (CURB-65) |
| **Care Coordination** | 5 | Discharge planning (LACE), referrals, care gaps (USPSTF), transitions of care (I-PASS), follow-up tasks |
| **Documentation** | 5 | SOAP notes, H&P, progress notes, discharge summaries, procedure notes |
| **Population Health** | 5 | Panel overview, quality measures (HEDIS), chronic disease registries, immunization status (CDC), preventive care compliance |
| **Specialty** | 5 | Prenatal (ACOG), pediatric growth (WHO/CDC), mental health (PHQ-9/GAD-7), oncology (TNM/RECIST), chronic pain |

**Full catalog with links:** [skills/README.md](skills/README.md)

### How to Use Skills

1. **Browse** the [skills/core/](skills/core/) directory and pick a skill
2. **Copy** the skill's `SKILL.md` content into your AI agent's system prompt or custom instructions
3. **Reference files** in each skill's `references/` subdirectory contain detailed clinical knowledge (scoring criteria, code tables, thresholds) that can optionally be included for deeper clinical accuracy

```
# Example: Add medication-reconciliation skill to your agent
skills/core/medication-management/medication-reconciliation/
├── SKILL.md              # Copy this into agent instructions
└── references/
    ├── reconciliation-process.md   # Joint Commission standards
    └── high-risk-medications.md    # ISMP high-alert drug list
```

**Integration guides:** [Claude](integrations/claude/) | [ChatGPT](integrations/chatgpt/) | [Gemini](integrations/gemini/)

**Community contributions welcome** - see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Development & Testing

### Build from Source

```bash
make build
```

### Run Locally (stdio mode)

```bash
make run
# or
./bin/langcare-mcp-fhir -config configs/config.local.yaml
```

### Run in HTTP Mode (Production)

```bash
make run-http
# or
./bin/langcare-mcp-fhir -http -port 8080 -config configs/config.yaml
```

### Run Tests

```bash
make test
```

### Lint Code

```bash
make lint
```

### Local Testing with EPIC

For step-by-step instructions on setting up EPIC credentials and testing locally:

**[📖 Local Testing Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/LOCAL-TESTING.md)**

This guide covers:
- Generating RSA keys and JWKS
- Configuring EPIC credentials
- Running the server locally
- Testing with Claude Desktop
- Troubleshooting common issues

**Quick credential test:**
```bash
# Test your EPIC credentials before running the server
go run test/test_epic_token.go "your-client-id" "/path/to/private-key.pem"
```

## Project Structure

```
langcare-mcp-fhir/
├── cmd/
│   └── server/
│       └── main.go                          # Entry point
├── internal/
│   ├── audit/
│   │   └── logger.go                        # HIPAA audit logging
│   ├── config/
│   │   └── config.go                        # YAML configuration loading
│   ├── fhir/
│   │   ├── client.go                        # FHIR HTTP client interface
│   │   ├── types.go                         # FHIR client types
│   │   └── providers/                       # Backend implementations
│   │       ├── base.go                      # Base HTTP provider
│   │       ├── epic.go                      # EPIC OAuth2 provider
│   │       ├── cerner.go                    # Cerner OAuth2 provider
│   │       └── gcp.go                       # GCP Healthcare API provider
│   ├── mcp/
│   │   └── server.go                        # MCP server initialization
│   ├── middleware/
│   │   ├── auth.go                          # MCP authentication
│   │   └── rate_limit.go                    # Rate limiting
│   ├── tools/                               # MCP tool implementations
│   │   ├── registry.go                      # Tool registry
│   │   ├── fhir_read.go                     # Read FHIR resource
│   │   ├── fhir_search.go                   # Search FHIR resources
│   │   ├── fhir_create.go                   # Create FHIR resource
│   │   └── fhir_update.go                   # Update FHIR resource
│   └── transport/
│       ├── stdio.go                         # stdio transport (Claude Desktop)
│       └── http.go                          # HTTP/SSE transport (production)
├── pkg/
│   └── types/
│       └── errors.go                        # Custom error types
├── configs/
│   ├── config.epic.example.yaml             # Example configuration for EPIC
│   ├── config.cerner.example.yaml           # Example configuration for Cerner
│   ├── config.gcp.example.yaml              # Example configuration for GCP
│   └── config.base.example.yaml             # Example configuration for any FHIR R4 server
├── docs/
│   ├── AGENT_PROMPT.md                      # AI agent system prompt
│   ├── EPIC-APP-SECURITY.md                 # EPIC authentication setup
│   ├── EPIC-SCOPES.md                       # OAuth2 scopes reference
│   ├── LOCAL-TESTING.md                     # Local development guide
│   └── SECURITY.md                          # Production security guide
├── test/
│   ├── README.md                            # Test documentation
│   └── test_epic_token.go                   # EPIC OAuth2 token tester
├── scripts/
│   └── create_jwks.sh                       # Generate JWKS from public key (used for EPIC FHIR)
├── bin/                                     # Build output (gitignored)
│   └── langcare-mcp-fhir                    # Compiled binary
├── go.mod                                   # Go module definition
├── go.sum                                   # Go module checksums
├── Makefile                                 # Build commands
└── README.md                                # This file
```

**Note:** The following are gitignored and not committed:
- `keys/` - Private keys and credentials
- `config.local.*.yaml` - Local configuration files
- `bin/` - Compiled binaries
- `.env` - Environment variables

## Documentation

### Getting Started
- **[📖 Local Development & Testing Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/LOCAL-TESTING.md)** - Complete guide for local setup and testing
- **[🚀 Installation & Configuration](#installation)** - Quick setup guide above

### Agent Integration
- **[🤖 Agent Prompt Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/AGENT_PROMPT.md)** - Complete guide for AI agents using LangCare MCP FHIR (tool examples, workflows, best practices)

### Security & Authentication
- **[🛡️ Security Documentation](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/SECURITY.md)** - Complete security architecture and HIPAA compliance
- **[🔐 EPIC Setup Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/EPIC-APP-SECURITY.md)** - JWT authentication, key generation, and JWKS registration
- **[📋 EPIC Scopes Reference](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/EPIC-SCOPES.md)** - Complete OAuth2 scopes guide for FHIR resources
- **[🔑 Authentication Methods](#supported-authentication-methods)** - Supported auth methods

### Development & Testing
- **[🧪 Testing Methods](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/LOCAL-TESTING.md#testing-with-claude-desktop)** - Claude Desktop, MCP Inspector, manual testing, and automation
- **[📦 Project Structure](#project-structure)** - Directory layout and architecture
- **[🔧 Build Commands](#development--testing)** - Development workflow

## Dependencies

- `github.com/modelcontextprotocol/go-sdk` - Official MCP SDK
- `gopkg.in/yaml.v3` - Configuration parsing
- `golang.org/x/oauth2` - OAuth2 client library
- `github.com/golang-jwt/jwt/v5` - JWT signing and verification
- Go 1.21+

## HIPAA Compliance

- PHI scrubbing enabled by default
- Never logs patient identifiers
- TLS support for HTTP transport
- Proper error sanitization
- Audit logging ready
- Stateless proxy design (no persistent storage)

## Testing

### Public Test Server

Default configuration uses HAPI FHIR public test server (`https://hapi.fhir.org/baseR4`) for immediate testing without setup.

### Test Your Setup

- **[📖 Local Development & Testing Guide](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/LOCAL-TESTING.md)** - Complete guide for setup, testing with Claude Desktop, MCP Inspector, and automation
- **[🔐 EPIC Security Setup](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/EPIC-APP-SECURITY.md)** - Detailed EPIC authentication guide
- **[🛡️ Security Documentation](https://github.com/langcare/langcare-mcp-fhir/blob/main/docs/SECURITY.md)** - Production deployment and security

## Contributing

**We welcome contributions from healthcare professionals, developers, and informaticists!**

There are three main ways to contribute:

### 1. Core MCP Server (Go Development)
- Bug fixes and performance improvements
- New FHIR provider implementations (AllScripts, Athenahealth, etc.)
- Security enhancements and observability features
- Testing and CI/CD improvements

### 2. Clinical Skills (Healthcare Workflows)
- Evidence-based clinical workflows using FHIR
- Specialty-specific protocols (cardiology, oncology, etc.)
- Population health and quality measure workflows
- Clinical decision support algorithms

Skills are agent-agnostic workflow guides that work across Claude, ChatGPT, and Gemini. No coding required - just clinical expertise and FHIR knowledge!

### 3. Agent Integrations (Platform Setup)
- Setup guides for new AI platforms
- Deployment examples (Docker, Kubernetes, cloud)
- Monitoring and observability setups
- CI/CD pipelines

**Get started:** Read [CONTRIBUTING.md](https://github.com/langcare/langcare-mcp-fhir/blob/main/CONTRIBUTING.md) for detailed guidelines, code standards, and submission process.

**Recognition:** Contributors are credited in README, release notes, and skill/integration author credits. Outstanding contributors may be invited as maintainers.

**Questions?** Open a [GitHub Discussion](https://github.com/langcare/langcare-mcp-fhir/discussions) or [issue](https://github.com/langcare/langcare-mcp-fhir/issues)!

## Community

- **GitHub Discussions** - Ask questions, share ideas: https://github.com/langcare/langcare-mcp-fhir/discussions
- **GitHub Issues** - Report bugs, request features: https://github.com/langcare/langcare-mcp-fhir/issues
- **Contributing Guide** - How to contribute: https://github.com/langcare/langcare-mcp-fhir/blob/main/CONTRIBUTING.md
- **Skills** - Clinical workflows: https://github.com/langcare/langcare-mcp-fhir/blob/main/skills/README.md

## License

See [LICENSE](https://github.com/langcare/langcare-mcp-fhir/blob/main/LICENSE) file.

---

**Built with ❤️ by the LangCare team and contributors.**

*Improving healthcare through better AI infrastructure.*