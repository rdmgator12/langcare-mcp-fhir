# LangCare MCP FHIR Server

[![GitHub Stars](https://img.shields.io/github/stars/langcare/langcare-mcp-fhir?style=social)](https://github.com/langcare/langcare-mcp-fhir)
[![Contributors](https://img.shields.io/github/contributors/langcare/langcare-mcp-fhir)](https://github.com/langcare/langcare-mcp-fhir/graphs/contributors)
[![License](https://img.shields.io/github/license/langcare/langcare-mcp-fhir)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/langcare/langcare-mcp-fhir)](go.mod)

Enterprise-grade MCP Server for FHIR-based EMRs, designed for robust deployments in agentic AI platforms.

## Architecture

This MCP server acts as an intelligent proxy between AI agents and FHIR R4 servers. It exposes 4 generic FHIR operations through the Model Context Protocol (MCP), enabling AI-powered workflows for any FHIR resource type.

**Key Design:**
- **MCP SDK:** Official `github.com/modelcontextprotocol/go-sdk` (Anthropic/Google maintained)
- **FHIR Client:** Generic HTTP client working with any FHIR R4 server
- **Transport:** stdio and HTTP/SSE
- **Backend:** Proxy to existing FHIR server (no database)

## Security Architecture

LangCare MCP FHIR Server implements a **two-layer security model** for HIPAA-compliant healthcare data access:

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   Claude    │ Auth1   │  MCP Server  │ Auth2   │  FHIR API   │
│   Client    │────────▶│   (Go)       │────────▶│   (EMR)     │
└─────────────┘         └──────────────┘         └─────────────┘

Auth1: MCP Client Authentication (Bearer Token/API Key)
Auth2: FHIR Backend Authentication (Bearer/OAuth2/SMART on FHIR)
```

### Security Layers

**Layer 1: MCP Client Authentication**
- Bearer token authentication for AI agent clients
- Token-based access control
- Rate limiting per client token
- Configurable via `MCP_AUTH_TOKENS` environment variable

**Layer 2: FHIR Backend Authentication**
- Supports multiple auth methods: Bearer, OAuth2, Basic, None
- Automatic token refresh for OAuth2
- Credential management via environment variables
- SMART on FHIR support for EPIC/Cerner

### Key Security Features

- ✅ **TLS 1.3** encryption for HTTP transport
- ✅ **PHI Scrubbing** in logs (enabled by default)
- ✅ **HIPAA-compliant** audit logging
- ✅ **No persistent PHI storage** (stateless proxy)
- ✅ **Secrets via environment variables** (never in config files)
- ✅ **OAuth 2.0** with automatic token refresh
- ✅ **mTLS** support for service-to-service communication

### Deployment Models

**Development (stdio):**
- Local, single-user
- Process-level isolation
- Credentials in environment variables

**Production (HTTP/SSE):**
- Multi-user, centralized
- TLS encryption required
- Kubernetes secrets for credential management
- Service mesh (Istio/Linkerd) for mTLS

**🔐 For complete security documentation, see [docs/SECURITY.md](docs/SECURITY.md):**
- HIPAA compliance checklist
- OAuth configuration for EPIC/Cerner/GCP
- Kubernetes security manifests
- Credential management procedures
- Network security architecture
- Audit logging implementation
- Incident response procedures

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
│   ├── config.gcp.example.yaml              # Example configuration for GCP
│   └── claude_desktop_config.json           # Claude Desktop setup example
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

## Agent Usage

AI agents use LangCare MCP FHIR Server to help healthcare professionals access and manage patient health records through 4 FHIR tools. The server handles EMR authentication, allowing agents to focus on clinical workflows while maintaining strict privacy and accuracy standards.

**Agent role and capabilities:**
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

**📖 Complete system prompt: [docs/AGENT_PROMPT.md](docs/AGENT_PROMPT.md)**

This guide provides:
- **System prompt** - Role definition, operational guidelines, response formats
- **Tool reference** - Detailed examples for fhir_search, fhir_read, fhir_create, fhir_update
- **Workflow examples** - Patient lookup, lab review, vitals recording, status updates
- **Search parameters** - Resource-specific query patterns and date handling
- **Error handling** - Common scenarios and recovery strategies
- **Quick reference** - Task → Tool mapping table

## Clinical Skills (Optional)

**Skills are optional, agent-agnostic clinical workflow guides** that enhance AI agents' ability to perform complex healthcare tasks using the MCP server's FHIR tools.

### What Are Skills?

Skills provide structured clinical workflows - like patient chart review, medication reconciliation, and preventive care screening - that work across all AI agents (Claude, ChatGPT, Gemini). They're copy-paste ready and focus on FHIR resource operations, not agent-specific features.

**Key points:**
- ✅ **Optional** - MCP server works perfectly without them
- ✅ **Portable** - Work with any AI agent (Claude, ChatGPT, Gemini)
- ✅ **Clinical Focus** - Based on evidence-based guidelines (USPSTF, ADA, ACC/AHA)
- ✅ **Community-Driven** - Contributions welcome from healthcare professionals

### Available Skills

**Core Skills:**
- **[fhir-clinical](skills/core/fhir-clinical/)** - Comprehensive clinical workflows including:
  - Patient chart review
  - Medication reconciliation
  - Lab result interpretation
  - Preventive care screening
  - Diabetes & hypertension management
  - Clinical documentation
  - Vital signs tracking

**Community Skills:**
- Browse community-contributed skills in [skills/community/](skills/community/)
- Contribute your own workflows following [skills/README.md](skills/README.md)

### How to Use Skills

**1. Choose a skill** from [skills/](skills/) directory
**2. Copy SKILL.md** to your AI agent's custom instructions or system prompt
**3. Follow integration guide** for your agent:
   - **[Claude](integrations/claude/)** - Projects, custom instructions
   - **[ChatGPT](integrations/chatgpt/)** - Custom GPTs, instructions
   - **[Gemini](integrations/gemini/)** - System instructions, function calling

**Example workflow with skills:**
```
User: "Review chart for patient John Doe, DOB 1965-03-15"

Agent (with fhir-clinical skill):
1. Searches for patient by name and DOB
2. Retrieves active conditions, medications, allergies
3. Gets recent labs with reference ranges
4. Reviews recent encounters
5. Presents structured clinical summary

Without skill: Agent would need step-by-step guidance for each operation
With skill: Agent follows learned workflow automatically
```

**📖 Learn more:** [skills/README.md](skills/README.md) | [integrations/README.md](integrations/README.md)

## Quick Start

### Build

```bash
make build
```

### Run (stdio mode)

```bash
make run
# or
./bin/langcare-mcp-fhir
```

### Run (HTTP mode)

```bash
make run-http
# or
./bin/langcare-mcp-fhir -http -port 8080
```

### Local Testing with EPIC

For step-by-step instructions on setting up EPIC credentials and testing locally, see:
**[📖 Local Testing Guide](docs/LOCAL-TESTING.md)**

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
See [test/README.md](test/README.md) for details.

### Configuration
See [config/config.epic.example.yaml](config/config.epic.example.yaml) for EPIC FHIR config.
See [config/config.gcp.example.yaml](config/config.gcp.example.yaml) for Google Cloud Healthcare API config.

security:
  mcp_auth_tokens: "token1,token2,token3"
  rate_limit:
    enabled: true
    rate: 100  # requests per second
    burst: 200 # burst size

```

## Development

### Build
```bash
make build
```

### Test
```bash
make test
```

### Lint
```bash
make lint
```

### Clean
```bash
make clean
```

## Authentication

Supports multiple auth types:

- **Bearer Token:** `type: "bearer"`
- **OAuth2:** `type: "oauth2"`
- **Basic Auth:** `type: "basic"`
- **None:** `type: "none"`

Set credentials via environment variables or directly in config.

## HIPAA Compliance

- PHI scrubbing enabled by default
- Never logs patient identifiers
- TLS support for HTTP transport
- Proper error sanitization
- Audit logging ready

## Testing

### Public Test Server

Default configuration uses HAPI FHIR public test server (`https://hapi.fhir.org/baseR4`) for immediate testing without setup.

### Testing the MCP Server

- **[📖 Local Development & Testing Guide](docs/LOCAL-TESTING.md)** - Complete guide for setup, testing with Claude Desktop, MCP Inspector, and automation
- **[🔐 EPIC Security Setup](docs/EPIC-APP-SECURITY.md)** - Detailed EPIC authentication guide
- **[🛡️ Security Documentation](docs/SECURITY.md)** - Production deployment and security

## Documentation

### Getting Started
- **[📖 Local Development & Testing Guide](docs/LOCAL-TESTING.md)** - Complete guide for local setup and testing
- **[🚀 Quick Start](#quick-start)** - Build and run in 5 minutes

### Agent Integration
- **[🤖 Agent Prompt Guide](docs/AGENT_PROMPT.md)** - Complete guide for AI agents using LangCare MCP FHIR (tool examples, workflows, best practices)

### Security & Authentication
- **[🛡️ Security Documentation](docs/SECURITY.md)** - Complete security architecture and HIPAA compliance
- **[🔐 EPIC Setup Guide](docs/EPIC-APP-SECURITY.md)** - JWT authentication, key generation, and JWKS registration
- **[📋 EPIC Scopes Reference](docs/EPIC-SCOPES.md)** - Complete OAuth2 scopes guide for FHIR resources
- **[🔑 Authentication](#authentication)** - Supported auth methods

### Development & Testing
- **[🧪 Testing Methods](docs/LOCAL-TESTING.md#testing-with-claude-desktop)** - Claude Desktop, MCP Inspector, manual testing, and automation
- **[📦 Project Structure](#project-structure)** - Directory layout and architecture
- **[🔧 Build Commands](#development)** - Development workflow

## Dependencies

- `github.com/modelcontextprotocol/go-sdk` - Official MCP SDK
- `gopkg.in/yaml.v3` - Configuration parsing
- `golang.org/x/oauth2` - OAuth2 client library
- `github.com/golang-jwt/jwt/v5` - JWT signing and verification
- Go 1.21+

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

**Get started:** Read [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines, code standards, and submission process.

**Recognition:** Contributors are credited in README, release notes, and skill/integration author credits. Outstanding contributors may be invited as maintainers.

**Questions?** Open a [GitHub Discussion](https://github.com/langcare/langcare-mcp-fhir/discussions) or issue!

## Community

- **GitHub Discussions** - Ask questions, share ideas: https://github.com/langcare/langcare-mcp-fhir/discussions
- **GitHub Issues** - Report bugs, request features: https://github.com/langcare/langcare-mcp-fhir/issues
- **Contributing Guide** - How to contribute: [CONTRIBUTING.md](CONTRIBUTING.md)
- **Skills** - Clinical workflows: [skills/README.md](skills/README.md)
- **Integrations** - Agent setup guides: [integrations/README.md](integrations/README.md)

## License

See LICENSE file.

---

**Built with ❤️ by the LangCare team and contributors.**

*Improving healthcare through better AI infrastructure.*