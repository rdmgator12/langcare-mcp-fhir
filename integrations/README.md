# Agent Integrations for LangCare MCP FHIR Server

**Agent-specific setup guides for using clinical skills with different AI platforms.**

## Overview

This directory contains integration guides for connecting AI agents to the LangCare MCP FHIR Server and using optional clinical skills. Each guide is tailored to a specific agent platform's configuration format and features.

## What's in an Integration Guide?

Each integration guide provides:

1. **MCP Server Setup** - How to connect the agent to LangCare MCP FHIR Server
2. **Skill Integration** - How to add optional clinical workflows to the agent
3. **Configuration Examples** - Copy-paste ready config files
4. **Testing Instructions** - How to verify the integration works
5. **Troubleshooting** - Common issues and solutions

## Available Integrations

### [Claude (Anthropic)](claude/)
- **Claude Desktop** - Local MCP server via stdio transport
- **Claude Projects** - Add skills to Project Knowledge
- **Custom Instructions** - System-level clinical workflows

### [ChatGPT (OpenAI)](chatgpt/)
- **Custom GPTs** - Create FHIR-enabled medical assistant GPT
- **Actions** - Connect to HTTP/SSE MCP server endpoint
- **Instructions** - Add clinical skills to GPT instructions

### [Gemini (Google)](gemini/)
- **Function Calling** - Map MCP tools to Gemini functions
- **System Instructions** - Add clinical workflows
- **Grounding** - Use FHIR data for health-specific responses

## Integration Comparison

| Feature | Claude Desktop | ChatGPT GPT | Gemini |
|---------|---------------|-------------|---------|
| **Transport** | stdio (local) | HTTP/SSE | HTTP/SSE |
| **Setup Complexity** | Low | Medium | Medium |
| **Skill Integration** | Project Knowledge | GPT Instructions | System Instructions |
| **Authentication** | Process-level | Bearer Token | API Key |
| **Best For** | Development, single-user | Multi-user, web access | Google Cloud integration |

## Quick Start

**1. Choose Your Agent**
- **Local development?** → Start with [Claude Desktop](claude/)
- **Web-based GPT?** → Use [ChatGPT Custom GPT](chatgpt/)
- **Google Cloud?** → Try [Gemini](gemini/)

**2. Follow Agent-Specific Guide**
Each directory contains a README.md with complete setup instructions.

**3. Add Optional Skills** (Optional)
Copy clinical workflows from [skills/](../skills/) to enhance your agent's clinical capabilities.

**4. Test the Integration**
Use provided test queries to verify FHIR operations work correctly.

## How Skills Work with Agents

**Skills are optional, agent-agnostic clinical workflows** that live in the [skills/](../skills/) directory. Integrations show you how to add them to your specific agent.

**Example Flow:**
1. Developer selects a skill (e.g., `skills/core/fhir-clinical/SKILL.md`)
2. Developer follows agent integration guide (e.g., `integrations/claude/README.md`)
3. Developer copies skill content into agent's instructions (Project Knowledge, GPT Instructions, etc.)
4. Agent now understands clinical workflows and can use MCP tools effectively

**Skills → Agent-Agnostic (portable across all agents)**
**Integrations → Agent-Specific (tailored to each platform)**

## Contributing Integrations

We welcome integration guides for additional AI agents and platforms!

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

**Nice to Have:**
- [ ] Screenshots or diagrams
- [ ] Video walkthrough
- [ ] Advanced configuration (rate limiting, multiple FHIR servers)
- [ ] Production deployment considerations
- [ ] CI/CD integration examples

### Contribution Process

1. **Create Directory** - `integrations/your-agent-name/`
2. **Write README.md** - Follow template from existing integrations
3. **Add Config Examples** - Include working configuration files
4. **Test Thoroughly** - Verify all 4 FHIR tools work (read, search, create, update)
5. **Document Gotchas** - Note platform-specific quirks or limitations
6. **Submit PR** - Include screenshots if possible

See [CONTRIBUTING.md](../CONTRIBUTING.md) for detailed guidelines.

## Supported Agent Platforms

### Currently Supported
✅ **Claude (Anthropic)** - Desktop app with stdio transport
✅ **ChatGPT (OpenAI)** - Custom GPTs with Actions
✅ **Gemini (Google)** - Function calling and system instructions

### Community Requested
📋 **LangChain** - MCP tool integration
📋 **AutoGPT** - Autonomous agent with FHIR access
📋 **Microsoft Copilot** - Healthcare plugins
📋 **Amazon Bedrock** - Claude/Llama with MCP
📋 **Hugging Face Agents** - Open-source LLM integration

**Want to see your favorite agent supported?** [Request it here](https://github.com/langcare/langcare-mcp-fhir/issues) or contribute an integration guide!

## Security Considerations

### Stdio Transport (Local Development)
- **Process-level isolation** - No network exposure
- **Single-user** - Direct access to FHIR credentials
- **No MCP auth** - Trust boundary at OS process level
- **Best for:** Local development, personal use

### HTTP/SSE Transport (Production)
- **Network-accessible** - Requires proper security
- **Multi-user** - Needs MCP client authentication (Bearer tokens)
- **TLS required** - Encrypt all communication
- **Rate limiting** - Prevent abuse
- **Best for:** Production deployments, team use

**Security Checklist:**
- [ ] Use HTTPS/TLS for HTTP transport
- [ ] Configure MCP authentication tokens (HTTP mode)
- [ ] Enable rate limiting in production
- [ ] Use environment variables for secrets (never commit credentials)
- [ ] Implement audit logging for PHI access
- [ ] Review [SECURITY.md](../docs/SECURITY.md) for HIPAA compliance

## Testing Your Integration

### Basic Functionality Test

**1. List Available Tools**
The agent should discover 4 FHIR tools:
- `fhir_read`
- `fhir_search`
- `fhir_create`
- `fhir_update`

**2. Test Tool Execution**
Run a simple search:
```
User: "Search for patients named Smith"
Agent should use: fhir_search with resourceType="Patient", queryParams="name=Smith"
```

**3. Verify FHIR Server Connection**
Check that responses contain valid FHIR resources (JSON with resourceType field).

**4. Test Error Handling**
Try an invalid resource ID and verify agent handles the error gracefully.

### Clinical Workflow Test (with Skills)

If you've added clinical skills, test a complete workflow:

**Patient Chart Review:**
```
User: "Review the chart for patient John Doe, DOB 1965-03-15"
Expected: Agent searches for patient, retrieves conditions, medications, labs, encounters, allergies
Output: Structured summary of clinical information
```

**Preventive Care:**
```
User: "What preventive care is this patient overdue for?"
Expected: Agent checks demographics, calculates age, searches for previous screenings
Output: List of overdue screenings with recommendations
```

## Getting Help

**Integration Issues:**
- Check agent-specific README in subdirectory
- Review troubleshooting section
- Search existing issues: https://github.com/langcare/langcare-mcp-fhir/issues

**MCP Server Issues:**
- See [LOCAL-TESTING.md](../docs/LOCAL-TESTING.md)
- Check server logs for errors
- Verify FHIR server credentials

**Skill Issues:**
- See [skills/README.md](../skills/README.md)
- Ensure OAuth scopes are sufficient
- Test FHIR server access manually

**Ask Questions:**
- GitHub Discussions: https://github.com/langcare/langcare-mcp-fhir/discussions
- Email: support@langcare.ai

## Examples

### Example 1: Claude Desktop (Simplest)

```bash
# 1. Build MCP server
make build

# 2. Add to Claude Desktop config
# ~/.config/Claude/claude_desktop_config.json

# 3. Restart Claude Desktop

# 4. Test: "Search for patients named Smith"
```

### Example 2: ChatGPT Custom GPT (Web-Based)

```bash
# 1. Run MCP server in HTTP mode
./bin/langcare-mcp-fhir -http -port 8080

# 2. Expose with ngrok or deploy to cloud
ngrok http 8080

# 3. Create Custom GPT in ChatGPT
# Add Action pointing to your SSE endpoint

# 4. Test in ChatGPT web interface
```

### Example 3: Gemini with Skills

```bash
# 1. Deploy MCP server to Google Cloud Run

# 2. Configure Gemini function calling
# Map fhir_* tools to Gemini functions

# 3. Add clinical skills to system instructions

# 4. Test via Gemini API
```

## Roadmap

**Upcoming Integrations:**
- [ ] LangChain integration example
- [ ] Slack bot with FHIR access
- [ ] Microsoft Teams integration
- [ ] Jupyter notebook examples
- [ ] REST API documentation for custom integrations

**Improvements:**
- [ ] Docker Compose setup for quick testing
- [ ] Integration test suite
- [ ] Video tutorials for each platform
- [ ] Monitoring and observability examples

**Want to contribute?** See [CONTRIBUTING.md](../CONTRIBUTING.md)

## License

All integration guides are licensed under the same license as the LangCare MCP FHIR Server. See [LICENSE](../LICENSE) for details.
