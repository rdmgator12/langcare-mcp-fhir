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

**[simple-clinical](core/simple-clinical/)** - General clinical workflows (chart review, vitals, documentation)

#### Patient Data & Summary (`core/patient-data-summary/`)
- **[patient-demographics-summary](core/patient-data-summary/patient-demographics-summary/)** - Demographics, emergency contacts, insurance, advance directives
- **[clinical-summary-generator](core/patient-data-summary/clinical-summary-generator/)** - CCD-style comprehensive clinical summary
- **[problem-list-review](core/patient-data-summary/problem-list-review/)** - Problem list audit with medication cross-referencing
- **[allergy-adverse-reaction-summary](core/patient-data-summary/allergy-adverse-reaction-summary/)** - Allergy categorization, cross-reactivity, medication conflicts
- **[insurance-coverage-summary](core/patient-data-summary/insurance-coverage-summary/)** - Coverage hierarchy, coordination of benefits, validation

#### Medication Management (`core/medication-management/`)
- **[medication-reconciliation](core/medication-management/medication-reconciliation/)** - Inpatient/outpatient med list comparison per Joint Commission
- **[drug-interaction-checker](core/medication-management/drug-interaction-checker/)** - CYP450 interaction analysis with severity classification
- **[medication-adherence-assessment](core/medication-management/medication-adherence-assessment/)** - MPR/PDC calculation, fill pattern analysis
- **[prescription-appropriateness-review](core/medication-management/prescription-appropriateness-review/)** - Beers Criteria, STOPP/START, renal dosing
- **[opioid-risk-assessment](core/medication-management/opioid-risk-assessment/)** - ORT scoring, MME calculation, CDC guideline thresholds

#### Lab & Diagnostics (`core/lab-diagnostics/`)
- **[lab-result-interpreter](core/lab-diagnostics/lab-result-interpreter/)** - Lab interpretation with delta checks and pattern recognition
- **[critical-value-alert-generator](core/lab-diagnostics/critical-value-alert-generator/)** - Critical value detection per CAP/CLIA thresholds
- **[preoperative-lab-checklist](core/lab-diagnostics/preoperative-lab-checklist/)** - Pre-op labs by surgery risk and ASA class
- **[diabetes-panel-review](core/lab-diagnostics/diabetes-panel-review/)** - HbA1c trending, ADA targets, complications screening
- **[renal-function-dashboard](core/lab-diagnostics/renal-function-dashboard/)** - KDIGO CKD staging, eGFR trajectory, electrolyte management

#### Clinical Decision Support (`core/clinical-decision-support/`)
- **[sepsis-screening](core/clinical-decision-support/sepsis-screening/)** - qSOFA, SOFA, SIRS scoring with sepsis bundle checklist
- **[cardiovascular-risk-assessment](core/clinical-decision-support/cardiovascular-risk-assessment/)** - CHA2DS2-VASc, HEART, Framingham, ASCVD, HAS-BLED
- **[vte-risk-assessment](core/clinical-decision-support/vte-risk-assessment/)** - Wells DVT/PE, Geneva, Caprini with prophylaxis recommendations
- **[fall-risk-assessment](core/clinical-decision-support/fall-risk-assessment/)** - Morse Fall Scale, Hendrich II, TUG with CarePlan generation
- **[pneumonia-severity-assessment](core/clinical-decision-support/pneumonia-severity-assessment/)** - CURB-65, PSI/PORT, CAP vs HAP/VAP differentiation

#### Care Coordination & Workflow (`core/care-coordination/`)
- **[discharge-planning-checklist](core/care-coordination/discharge-planning-checklist/)** - Discharge readiness with LACE readmission risk
- **[referral-generator](core/care-coordination/referral-generator/)** - Specialist referral ServiceRequest with prior auth awareness
- **[care-gap-identifier](core/care-coordination/care-gap-identifier/)** - USPSTF preventive care gap screening
- **[transition-of-care-summary](core/care-coordination/transition-of-care-summary/)** - C-CDA compliant TOC with I-PASS handoff
- **[follow-up-task-generator](core/care-coordination/follow-up-task-generator/)** - FHIR Task generation with priority and assignment logic

#### Documentation (`core/documentation/`)
- **[soap-note-generator](core/documentation/soap-note-generator/)** - Structured SOAP notes with E&M coding guidance
- **[history-and-physical-generator](core/documentation/history-and-physical-generator/)** - Comprehensive H&P with OLDCARTS, 14-system ROS
- **[progress-note-writer](core/documentation/progress-note-writer/)** - Daily inpatient progress notes with ICU additions
- **[discharge-summary-writer](core/documentation/discharge-summary-writer/)** - CMS/TJC compliant discharge summaries
- **[procedure-note-template](core/documentation/procedure-note-template/)** - Procedure documentation with safety checklists

#### Population Health & Analytics (`core/population-health/`)
- **[patient-panel-overview](core/population-health/patient-panel-overview/)** - Cohort querying, metric aggregation, risk stratification
- **[quality-measure-dashboard](core/population-health/quality-measure-dashboard/)** - HEDIS/CMS quality measure calculation
- **[chronic-disease-registry-query](core/population-health/chronic-disease-registry-query/)** - Disease registry building for 6 chronic conditions
- **[immunization-status-checker](core/population-health/immunization-status-checker/)** - CDC adult/pediatric schedule comparison
- **[preventive-care-compliance-report](core/population-health/preventive-care-compliance-report/)** - USPSTF/ACS/CDC compliance scorecard

#### Specialty-Specific (`core/specialty/`)
- **[prenatal-visit-workflow](core/specialty/prenatal-visit-workflow/)** - ACOG prenatal visits by trimester with complication screening
- **[pediatric-growth-assessment](core/specialty/pediatric-growth-assessment/)** - WHO/CDC growth charts, percentiles, FTT detection
- **[mental-health-screening](core/specialty/mental-health-screening/)** - PHQ-9, GAD-7, AUDIT-C, C-SSRS, MDQ, PC-PTSD-5
- **[oncology-treatment-timeline](core/specialty/oncology-treatment-timeline/)** - Cancer treatment mapping with TNM staging and RECIST
- **[chronic-pain-management-review](core/specialty/chronic-pain-management-review/)** - Multimodal pain assessment with MME and opioid safety

### Community Skills

Community-contributed skills will appear in `community/`. To contribute a skill, see [CONTRIBUTING.md](../CONTRIBUTING.md).

## How to Use Skills

Skills are designed to be copied into your AI agent's custom instructions or system prompt. Each skill includes:

1. **SKILL.md** - The workflow guide with YAML frontmatter and step-by-step instructions
2. **references/** - Detailed clinical knowledge (scoring criteria, drug tables, guidelines)

**Integration guides for specific agents:**
- **[Claude](../integrations/claude/)** - Projects, custom instructions
- **[ChatGPT](../integrations/chatgpt/)** - Custom GPTs, custom instructions
- **[Gemini](../integrations/gemini/)** - System instructions

## Skill Format

Each skill should follow this structure:

```
skills/
├── core/
│   ├── fhir-clinical/                    # General clinical workflows
│   ├── patient-data-summary/             # Patient Data & Summary
│   │   └── patient-demographics-summary/
│   │       ├── SKILL.md                  # The workflow guide (agent-agnostic)
│   │       └── references/               # Detailed clinical knowledge
│   │           ├── scoring-criteria.md
│   │           └── code-tables.md
│   ├── medication-management/            # Medication Management
│   ├── lab-diagnostics/                  # Lab & Diagnostics
│   ├── clinical-decision-support/        # Clinical Decision Support
│   ├── care-coordination/                # Care Coordination & Workflow
│   ├── documentation/                    # Documentation
│   ├── population-health/                # Population Health & Analytics
│   └── specialty/                        # Specialty-Specific
└── community/
    └── contributor-skill/
        ├── SKILL.md
        └── references/
```

**SKILL.md** should contain:
- YAML frontmatter (name, description with triggers, metadata)
- Overview and FHIR Resources Used
- Step-by-step instructions using fhir_search, fhir_read, fhir_create, fhir_update
- At least 2 examples with "User says -> Actions -> Result" format
- At least 2 troubleshooting entries
- Related Skills section

**references/** should contain:
- Detailed clinical knowledge (scoring criteria, thresholds, diagnostic criteria)
- Society guideline references (AHA, IDSA, ADA, ACOG, etc.)
- LOINC/SNOMED/RxNorm code tables
- Drug interaction tables, dosing references, protocol details

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
