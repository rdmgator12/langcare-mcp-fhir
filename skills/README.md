# Clinical Skills for FHIR MCP

**Skills are optional, agent-agnostic clinical workflow guides** that help AI agents perform complex healthcare tasks using the LangCare MCP FHIR Server's 4 generic FHIR tools.

## What Are Skills?

Skills are **structured clinical workflows** that describe how to use FHIR operations to accomplish real-world healthcare tasks. They are:

- **Optional** - The MCP server works perfectly without them
- **Agent-agnostic** - Work with Claude, ChatGPT, Gemini, or any AI agent
- **FHIR-focused** - Describe resource patterns, search strategies, and validation rules
- **Copy-paste ready** - Can be added to any agent's custom instructions or system prompts

## Skills vs. Agent Prompts

| Feature | Skills (This Directory) | Agent Prompts |
|---------|------------------------|---------------|
| **Purpose** | Clinical workflows using FHIR | Agent-specific configuration |
| **Location** | `skills/` directory | Agent's custom instructions |
| **Portability** | Works across all agents | Tied to specific agent |
| **Focus** | FHIR resource operations | Agent behavior and tone |
| **Required** | No, optional enhancement | Yes, for agent setup |

## Available Skills

### Core Skills

**[fhir-clinical](core/fhir-clinical/)** - General clinical workflows
- Patient chart review
- Medication reconciliation
- Lab result interpretation
- Preventive care screening
- Chronic disease management (diabetes, hypertension)
- Clinical documentation

### Community Skills

Community-contributed skills will appear in `community/`. To contribute a skill, see [CONTRIBUTING.md](../CONTRIBUTING.md).

## How to Use Skills

Skills are designed to be copied into your AI agent's custom instructions or system prompt. Each skill includes:

1. **SKILL.md** - The actual workflow guide (copy this)
2. **README.md** - Documentation about the skill

**Integration guides for specific agents:**
- **[Claude](../integrations/claude/)** - Projects, custom instructions
- **[ChatGPT](../integrations/chatgpt/)** - Custom GPTs, custom instructions
- **[Gemini](../integrations/gemini/)** - System instructions

## Skill Format

Each skill should follow this structure:

```
skills/
├── core/
│   └── your-skill-name/
│       ├── SKILL.md           # The workflow guide (agent-agnostic)
│       └── README.md          # Documentation about the skill
└── community/
    └── contributor-skill/
        ├── SKILL.md
        └── README.md
```

**SKILL.md** should contain:
- Clinical context and purpose
- Required FHIR resources
- Step-by-step workflows using fhir_search, fhir_read, fhir_create, fhir_update
- Search parameter examples
- Validation rules
- Safety considerations

**README.md** should contain:
- Skill overview
- Use cases
- Prerequisites (required OAuth scopes)
- Clinical accuracy notes
- Author/contributor info

## Contributing Skills

We welcome clinical skills from healthcare professionals, informaticists, and developers!

**What makes a good skill:**
- Addresses real clinical workflows
- Uses FHIR resources correctly
- Includes validation and safety checks
- Agent-agnostic (works with any AI)
- Clear, tested examples

**How to contribute:**
1. Create a new directory in `skills/community/your-skill-name/`
2. Write `SKILL.md` with workflows and `README.md` with documentation
3. Test with at least one AI agent (Claude, ChatGPT, or Gemini)
4. Submit PR with clinical accuracy validation

See [CONTRIBUTING.md](../CONTRIBUTING.md) for detailed guidelines.

## Skill Development Guidelines

### Clinical Accuracy
- Follow HL7 FHIR R4 specifications
- Use standard code systems (LOINC, SNOMED, RxNorm, ICD-10)
- Include reference ranges and normal values
- Validate clinical logic with healthcare professionals

### FHIR Best Practices
- Use appropriate search parameters
- Handle Bundle pagination for large result sets
- Include error handling guidance
- Respect FHIR resource invariants

### Safety First
- Always verify patient identity before updates
- Include confirmation steps for critical actions
- Document contraindications and red flags
- Note scope requirements for operations

### Agent Portability
- Don't reference specific agent features (no "use Claude's artifacts")
- Use standard FHIR tool names: fhir_read, fhir_search, fhir_create, fhir_update
- Avoid agent-specific syntax or formatting
- Test with multiple agents if possible

## Examples

**Good skill content (agent-agnostic):**
```markdown
## Medication Reconciliation Workflow

1. **Search for current medications**
   Use fhir_search with:
   - resourceType: "MedicationStatement"
   - queryParams: "patient=[id]&status=active"

2. **Validate each medication**
   - Check for duplicates (same ingredient)
   - Identify drug-drug interactions
   - Verify dosing is within therapeutic range
```

**Bad skill content (agent-specific):**
```markdown
## Medication Reconciliation (Claude-Specific)

1. Use artifacts to create a table
2. Ask the user if they want to continue
3. Use Claude's analysis feature to...
```

## License

All skills in this directory are licensed under the same license as the LangCare MCP FHIR Server. By contributing, you agree to license your skills under this license.

## Questions?

- Open an issue for skill requests or questions
- See [CONTRIBUTING.md](../CONTRIBUTING.md) for contribution guidelines
- Join discussions in GitHub Discussions
