# Contributing to LangCare MCP FHIR Server

We welcome contributions from healthcare professionals, informaticists, developers, and anyone passionate about improving healthcare AI infrastructure!

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
  - [Contributing to Core MCP Server](#contributing-to-core-mcp-server)
  - [Contributing Clinical Skills](#contributing-clinical-skills)
  - [Contributing Agent Integrations](#contributing-agent-integrations)
- [Development Setup](#development-setup)
- [Pull Request Process](#pull-request-process)
- [Code Standards](#code-standards)
- [Testing Requirements](#testing-requirements)
- [Documentation](#documentation)
- [Community](#community)

## Code of Conduct

### Our Pledge

We are committed to providing a welcoming, inclusive, and harassment-free experience for everyone, regardless of age, body size, disability, ethnicity, gender identity and expression, level of experience, nationality, personal appearance, race, religion, or sexual identity and orientation.

### Expected Behavior

- Use welcoming and inclusive language
- Be respectful of differing viewpoints and experiences
- Gracefully accept constructive criticism
- Focus on what is best for the community and patients
- Show empathy towards other community members
- Prioritize patient safety and clinical accuracy in all contributions

### Unacceptable Behavior

- Trolling, insulting/derogatory comments, and personal attacks
- Public or private harassment
- Publishing others' private information without permission
- Contributions that compromise patient safety or data security
- Other conduct which could reasonably be considered inappropriate in a professional healthcare setting

### Enforcement

Violations of the Code of Conduct may be reported to the project team at conduct@langcare.ai. All complaints will be reviewed and investigated promptly and fairly. Project maintainers have the right and responsibility to remove, edit, or reject contributions that do not align with this Code of Conduct.

## How to Contribute

There are three main areas where contributions are welcome:

1. **Core MCP Server** - Go code, FHIR operations, infrastructure
2. **Clinical Skills** - Agent-agnostic clinical workflows using FHIR
3. **Agent Integrations** - Setup guides for different AI platforms

Choose the area that matches your expertise!

## Contributing to Core MCP Server

**What we're looking for:**

- Bug fixes and performance improvements
- New FHIR provider implementations (e.g., AllScripts, Athenahealth)
- Enhanced security features
- Better error handling and logging
- Improved testing and CI/CD
- Documentation improvements

### Areas of Focus

**1. FHIR Provider Support**

Add support for additional EMR systems:

- Implement `internal/fhir/providers/yourprovider.go`
- Follow existing patterns from `epic.go`, `cerner.go`, `gcp.go`
- Support OAuth 2.0 or SMART on FHIR authentication
- Include provider-specific configuration examples
- Document required OAuth scopes

**2. MCP Protocol Features**

- Resource prompts (content from FHIR resources)
- Sampling (server-initiated requests)
- Logging (structured MCP protocol logs)
- Advanced tool schemas (conditional parameters)

**3. Performance & Scalability**

- Connection pooling for FHIR HTTP clients
- Caching layer for frequently accessed resources
- Request batching and optimization
- Load testing and benchmarks

**4. Security Enhancements**

- mTLS support for service-to-service communication
- Advanced rate limiting (per-resource, per-operation)
- Audit log enhancements (PII detection, SIEM integration)
- Secrets management (HashiCorp Vault, AWS Secrets Manager)

**5. Observability**

- Prometheus metrics for FHIR operations
- OpenTelemetry tracing
- Structured logging with correlation IDs
- Health check improvements

### Development Process

**1. Fork and Clone**

```bash
git clone https://github.com/YOUR-USERNAME/langcare-mcp-fhir.git
cd langcare-mcp-fhir
```

**2. Create Feature Branch**

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/bug-description
```

**3. Make Changes**

- Follow Go best practices and idioms
- Add tests for new functionality (>80% coverage)
- Update documentation (inline comments, README)
- Ensure code passes linting: `make lint`

**4. Test Thoroughly**

```bash
# Unit tests
make test

# Integration tests (requires FHIR server)
go test -tags=integration ./...

# Manual testing
make build
./bin/langcare-mcp-fhir -config configs/config.test.yaml
```

**5. Commit and Push**

```bash
git add .
git commit -m "feat: add AllScripts FHIR provider support"
git push origin feature/your-feature-name
```

**Commit message format:**
- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation only
- `test:` - Adding/fixing tests
- `refactor:` - Code refactoring
- `perf:` - Performance improvements
- `chore:` - Maintenance tasks

**6. Open Pull Request**

- Describe changes clearly
- Link related issues
- Include testing instructions
- Add screenshots/demos if applicable

## Contributing Clinical Skills

**What we're looking for:**

- Structured clinical workflows using FHIR resources
- Evidence-based guidelines (USPSTF, ADA, ACC/AHA, etc.)
- Real-world use cases from clinical practice
- Specialty-specific workflows (cardiology, endocrinology, oncology)
- Population health and quality measure workflows
- Clinical decision support algorithms

### Skill Development Guidelines

**1. Clinical Accuracy**

- Follow current clinical practice guidelines (cite sources)
- Use standard code systems (LOINC, SNOMED CT, RxNorm, ICD-10)
- Include reference ranges and normal values
- Note contraindications and safety warnings
- Validate with healthcare professionals before submitting

**2. FHIR Best Practices**

- Use appropriate FHIR resource types
- Follow HL7 FHIR R4 specifications
- Include proper search parameters
- Handle Bundle pagination for large result sets
- Document required OAuth scopes for operations

**3. Agent Portability**

Skills must work across all AI agents (Claude, ChatGPT, Gemini):

- ✅ Describe FHIR workflows generically
- ✅ Use standard tool names: `fhir_read`, `fhir_search`, `fhir_create`, `fhir_update`
- ✅ Focus on clinical logic, not agent-specific syntax
- ❌ Don't reference agent features ("use Claude's artifacts")
- ❌ Don't use agent-specific formatting
- ❌ Don't assume specific agent capabilities

**4. Structure and Format**

Each skill should include:

```
skills/
├── core/                    # Curated, maintained skills
│   └── your-skill/
│       ├── SKILL.md         # The workflow guide
│       └── README.md        # Documentation
└── community/               # Community contributions
    └── your-skill/
        ├── SKILL.md
        └── README.md
```

**SKILL.md template:**

```markdown
# Skill Name

Brief description of the clinical workflow and purpose.

## Workflow 1: Descriptive Name

**Purpose:** What this workflow accomplishes

### Step 1: Action Description
```
Tool: fhir_search (or fhir_read, fhir_create, fhir_update)
resourceType: "Patient"
queryParams: "name=Smith"
```

- Explain what to look for
- Note validation rules
- Document error handling

### Step 2: Next Action
...

### Clinical Interpretation

- How to interpret results
- Reference ranges or thresholds
- Red flags or critical values

## Workflow 2: Another Workflow
...

## Safety Considerations

- Patient verification requirements
- Critical values requiring immediate action
- Documentation requirements
- Compliance notes
```

**README.md template:**

```markdown
# Skill Name

## Overview
Brief description of skill purpose and use cases.

## Use Cases
- Specific scenario 1
- Specific scenario 2

## Prerequisites

### Required OAuth Scopes
- `system/Patient.read`
- `system/Observation.read`
- ...

### FHIR Server Compatibility
- FHIR R4 required
- Tested with: EPIC, Cerner, GCP Healthcare API

## Clinical Accuracy Notes
- Guidelines followed (cite sources)
- Code systems used
- Known limitations

## Testing
- How to test this skill
- Example queries
- Expected outcomes

## Contributing
How others can improve this skill

## Authors
- Your Name / Organization
- Contributors
```

### Skill Submission Process

**1. Create Skill Files**

```bash
mkdir -p skills/community/your-skill-name
cd skills/community/your-skill-name
# Create SKILL.md and README.md
```

**2. Test with Multiple Agents**

Test your skill with at least **one** AI agent (Claude, ChatGPT, or Gemini):

- Copy SKILL.md into agent's instructions
- Test all workflows with real or synthetic data
- Verify clinical accuracy of outputs
- Document any issues or limitations

**3. Submit Pull Request**

```bash
git checkout -b skill/your-skill-name
git add skills/community/your-skill-name/
git commit -m "skill: add [skill name] clinical workflow"
git push origin skill/your-skill-name
```

**In PR description:**
- Describe clinical use cases
- Cite guidelines followed
- Note testing performed (which agent, what data)
- Declare any conflicts of interest
- Confirm clinical accuracy review

### Skill Review Criteria

Submissions will be reviewed for:

- ✅ **Clinical Accuracy** - Follows evidence-based guidelines
- ✅ **Safety** - Includes appropriate warnings and validations
- ✅ **Portability** - Works across agents (not agent-specific)
- ✅ **Completeness** - Includes both SKILL.md and README.md
- ✅ **Code Systems** - Uses standard terminologies (LOINC, SNOMED, RxNorm)
- ✅ **Testing** - Demonstrated to work with at least one agent
- ✅ **Documentation** - Clear, well-structured, includes examples

Skills meeting criteria may be promoted from `community/` to `core/`.

## Contributing Agent Integrations

**What we're looking for:**

- Setup guides for additional AI platforms
- Improved integration patterns
- Deployment examples (Docker, Kubernetes, cloud platforms)
- CI/CD pipelines
- Monitoring and observability setups

### Integration Checklist

A complete integration guide should include:

**Required:**
- [ ] MCP server connection setup (stdio or HTTP/SSE)
- [ ] Authentication configuration
- [ ] Tool availability verification
- [ ] Skill integration instructions
- [ ] Example configuration files
- [ ] Testing section with sample queries
- [ ] Troubleshooting common issues

**Recommended:**
- [ ] Screenshots or diagrams
- [ ] Video walkthrough
- [ ] Advanced configuration options
- [ ] Production deployment considerations
- [ ] Security best practices
- [ ] Cost optimization tips (for cloud deployments)

### Integration Submission Process

**1. Create Integration Directory**

```bash
mkdir -p integrations/your-agent-name
cd integrations/your-agent-name
```

**2. Write README.md**

Follow the template from existing integrations:
- Overview and architecture
- Prerequisites
- Step-by-step setup
- Testing instructions
- Troubleshooting
- Examples

**3. Test Thoroughly**

- Verify all 4 FHIR tools work (read, search, create, update)
- Test with real FHIR server (HAPI test server or EMR sandbox)
- Test with optional clinical skills
- Document any platform-specific quirks

**4. Submit Pull Request**

```bash
git checkout -b integration/your-agent-name
git add integrations/your-agent-name/
git commit -m "integration: add [agent name] setup guide"
git push origin integration/your-agent-name
```

## Development Setup

### Prerequisites

- **Go 1.21+** - https://golang.org/dl/
- **Make** - Standard on Mac/Linux, install on Windows
- **Docker** (optional) - For containerized testing
- **FHIR Server Access** - Use HAPI test server or EMR sandbox

### Initial Setup

```bash
# Clone repository
git clone https://github.com/langcare/langcare-mcp-fhir.git
cd langcare-mcp-fhir

# Install dependencies
go mod download

# Build
make build

# Run tests
make test

# Run linter
make lint
```

### Running Locally

**With HAPI test server (no auth):**

```bash
./bin/langcare-mcp-fhir -config configs/config.yaml
```

**With EPIC sandbox:**

1. Register app at https://fhir.epic.com/Developer/Apps
2. Generate RSA key pair: `openssl genrsa -out private-key.pem 2048`
3. Create `configs/config.local.epic.yaml` with credentials
4. Run: `./bin/langcare-mcp-fhir -config configs/config.local.epic.yaml`

### Project Structure

```
langcare-mcp-fhir/
├── cmd/server/           # Main application entry point
├── internal/             # Private application code
│   ├── audit/           # HIPAA audit logging
│   ├── config/          # Configuration management
│   ├── fhir/            # FHIR client implementations
│   ├── mcp/             # MCP server logic
│   ├── middleware/      # HTTP middleware (auth, rate limit)
│   ├── tools/           # MCP tool implementations
│   └── transport/       # stdio and HTTP transports
├── pkg/                 # Public library code
├── configs/             # Example configurations
├── docs/                # Documentation
├── skills/              # Clinical workflow skills
├── integrations/        # Agent setup guides
├── test/                # Integration tests
├── scripts/             # Build and deployment scripts
└── deployments/         # Kubernetes manifests (planned)
```

## Pull Request Process

### Before Submitting

1. **Run tests:** `make test` (all tests pass)
2. **Run linter:** `make lint` (no warnings)
3. **Update docs:** README, inline comments, CHANGELOG
4. **Test manually:** Build and run with your changes
5. **Self-review:** Check diff for unintended changes

### PR Description Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Clinical skill
- [ ] Integration guide
- [ ] Documentation
- [ ] Refactoring
- [ ] Performance improvement

## Motivation
Why is this change needed? What problem does it solve?

## Changes Made
- Change 1
- Change 2

## Testing
How was this tested?
- [ ] Unit tests added/updated
- [ ] Integration tests run
- [ ] Manual testing completed
- [ ] Tested with [FHIR server type]
- [ ] Tested with [AI agent] (for skills/integrations)

## Checklist
- [ ] Code follows project style
- [ ] Tests pass locally
- [ ] Documentation updated
- [ ] No breaking changes (or documented)
- [ ] Commit messages follow convention

## Screenshots/Demos (if applicable)
```

### Review Process

1. **Automated Checks** - CI runs tests and linting
2. **Maintainer Review** - Code quality, design, security
3. **Clinical Review** (for skills) - Healthcare professional validates accuracy
4. **Community Feedback** - Discussion and iteration
5. **Approval** - Maintainer approves and merges

**Typical timeline:** 3-7 days for initial review, longer for complex changes

## Code Standards

### Go Style Guide

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting (automatic with `make lint`)
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Write clear, idiomatic Go code

### Key Principles

**1. Simplicity**
- Prefer simple solutions over clever ones
- Avoid premature abstraction
- Clear code > concise code

**2. Error Handling**
- Always check errors
- Wrap errors with context: `fmt.Errorf("operation failed: %w", err)`
- Return errors, don't panic (except in init/main)

**3. Security**
- Never log raw tokens or PHI
- Use constant-time comparison for tokens
- Validate all input (resource types, IDs, params)
- Sanitize errors before returning to client

**4. Testing**
- Aim for >80% test coverage
- Use table-driven tests
- Mock external dependencies (FHIR servers)
- Include integration tests for critical paths

### Code Review Checklist

- [ ] Code is clear and well-documented
- [ ] Error handling is comprehensive
- [ ] Security best practices followed
- [ ] Tests cover new functionality
- [ ] No performance regressions
- [ ] Backward compatible (or documented)
- [ ] Logging is appropriate (not excessive)
- [ ] Configuration is validated

## Testing Requirements

### Unit Tests

**Coverage target:** >80%

```bash
# Run tests
make test

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**Test structure:**

```go
func TestFHIRReadTool_Execute(t *testing.T) {
    tests := []struct {
        name    string
        args    map[string]interface{}
        want    interface{}
        wantErr bool
    }{
        {
            name: "successful read",
            args: map[string]interface{}{
                "resourceType": "Patient",
                "id":           "12345",
            },
            wantErr: false,
        },
        {
            name: "missing resource type",
            args: map[string]interface{}{
                "id": "12345",
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Integration Tests

Test against real FHIR servers:

```bash
# Tag integration tests
// +build integration

# Run with tag
go test -tags=integration ./...
```

**Test FHIR servers:**
- HAPI FHIR test server: https://hapi.fhir.org/baseR4
- EPIC sandbox: https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4
- Local FHIR server (Docker): `docker run -p 8080:8080 hapiproject/hapi:latest`

### Manual Testing

```bash
# Build and run
make build
./bin/langcare-mcp-fhir -config configs/config.test.yaml

# In another terminal, test MCP protocol
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | ./bin/langcare-mcp-fhir
```

## Documentation

### Inline Documentation

- Add GoDoc comments for exported functions, types, and packages
- Explain *why* not *what* (code shows what)
- Document non-obvious behavior
- Include examples for complex functions

```go
// TokenManager handles OAuth token lifecycle with automatic refresh.
// It caches tokens and refreshes them proactively before expiry
// to prevent authentication failures during FHIR operations.
type TokenManager struct {
    strategy Strategy
    cache    *oauth2.Token
    mu       sync.RWMutex
}
```

### README Updates

When adding features, update:
- Main README.md (if user-facing feature)
- Relevant docs/ files
- Example configs
- Integration guides (if applicable)

### Changelog

Add entry to CHANGELOG.md (if exists) or PR description:

```markdown
## [Unreleased]

### Added
- AllScripts FHIR provider support (#123)
- Prometheus metrics for tool execution (#125)

### Fixed
- OAuth token refresh race condition (#124)

### Changed
- Improved error messages for authentication failures (#126)
```

## Community

### Getting Help

- **GitHub Discussions** - Questions, ideas, showcase: https://github.com/langcare/langcare-mcp-fhir/discussions
- **GitHub Issues** - Bugs, feature requests: https://github.com/langcare/langcare-mcp-fhir/issues
- **Email** - Private inquiries: support@langcare.ai

### Communication Channels

- **GitHub** - Primary communication (issues, discussions, PRs)
- **Slack** (planned) - Real-time chat for contributors
- **Monthly Calls** (planned) - Community meetings for major decisions

### Recognition

Contributors are recognized in:
- README.md Contributors section
- CONTRIBUTORS.md (planned)
- Release notes
- Skill/integration author credits

Outstanding contributors may be invited to become maintainers.

## Maintainers

Current maintainers:
- **LangCare Team** - @langcare-ai

Maintainers are responsible for:
- Reviewing and merging PRs
- Triaging issues
- Maintaining code quality and security
- Coordinating releases
- Enforcing Code of Conduct

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (see LICENSE file).

For clinical skills, you also agree that:
- Content is medically accurate to the best of your knowledge
- You have rights to share the content
- The skill may be freely used and modified by the community

---

**Thank you for contributing to LangCare MCP FHIR Server!**

Your contributions help improve healthcare AI infrastructure and ultimately benefit patient care. Whether you're fixing a bug, adding a clinical workflow, or writing documentation, every contribution matters.

Questions? Open a discussion or issue on GitHub!
