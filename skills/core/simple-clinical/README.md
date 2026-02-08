# FHIR Clinical Workflows Skill

**Agent-agnostic clinical workflows for AI-powered healthcare using FHIR R4 resources.**

## Overview

This skill provides structured workflows for common clinical tasks using the LangCare MCP FHIR Server's 4 generic tools:
- `fhir_search` - Find resources with query parameters
- `fhir_read` - Retrieve specific resource by ID
- `fhir_create` - Create new resource
- `fhir_update` - Update existing resource

## Workflows Included

1. **Patient Chart Review** - Comprehensive review of clinical record (problems, meds, labs, encounters)
2. **Medication Reconciliation** - Validate medication list, check interactions, identify gaps
3. **Lab Result Interpretation** - Review and interpret laboratory results with reference ranges
4. **Preventive Care Screening** - Identify overdue screenings (cancer, CV, immunizations)
5. **Diabetes Management** - Monitor glucose control, complications, and adherence to guidelines
6. **Hypertension Management** - Blood pressure monitoring and antihypertensive optimization
7. **Clinical Documentation** - Create structured clinical notes using Composition resources
8. **Vital Signs Tracking** - Record and trend vital signs over time

## Use Cases

### Primary Care
- Annual wellness visits
- Chronic disease management (diabetes, hypertension, COPD)
- Preventive care tracking
- Medication review and renewal

### Hospital Medicine
- Admission history and physical
- Daily progress notes
- Medication reconciliation at admission/discharge
- Discharge summaries

### Specialty Care
- Pre-visit chart review
- Focused assessments (cardiology, endocrinology, nephrology)
- Procedure documentation
- Consultation notes

### Population Health
- Identify patients overdue for screenings
- Quality measure reporting (HEDIS, MIPS)
- Care gap closure
- Risk stratification

## Prerequisites

### Required OAuth Scopes

The workflows in this skill require appropriate FHIR OAuth2 scopes for the EMR system:

**Read Operations (Required for all workflows):**
- `system/Patient.read` - Patient demographics
- `system/Condition.read` - Diagnoses and problems
- `system/Observation.read` - Labs, vitals, assessments
- `system/MedicationStatement.read` - Current medications
- `system/Encounter.read` - Visits and hospitalizations
- `system/AllergyIntolerance.read` - Allergies
- `system/Procedure.read` - Procedures and screenings

**Write Operations (For documentation workflows):**
- `system/Observation.write` - Record vitals
- `system/Composition.write` - Create clinical notes
- `system/CarePlan.write` - Document care plans
- `system/MedicationRequest.write` - Prescribe medications

**Extended Resources (For comprehensive workflows):**
- `system/DiagnosticReport.read` - Lab/imaging reports
- `system/DocumentReference.read` - Clinical documents
- `system/Immunization.read` - Vaccine history
- `system/CarePlan.read` - Treatment plans

See [EPIC-SCOPES.md](../../../docs/EPIC-SCOPES.md) for detailed scope configuration for EPIC FHIR.

### FHIR Server Compatibility

This skill is designed for **FHIR R4** servers. It has been tested with:
- EPIC FHIR API (Interconnect)
- Cerner FHIR API
- Google Cloud Healthcare API
- HAPI FHIR Server (public test server)

**Server Requirements:**
- FHIR R4 (HL7 FHIR version 4.0.1)
- Support for standard search parameters
- OAuth 2.0 or SMART on FHIR authentication
- JSON response format

## How to Use

1. **Copy SKILL.md** to your AI agent's custom instructions or system prompt
2. **Configure MCP server** with appropriate FHIR server and OAuth scopes
3. **Test workflows** with non-critical patients first
4. **Validate clinical accuracy** - Always review AI-generated recommendations

### Integration by Agent

**Claude (Anthropic):**
- Copy [SKILL.md](SKILL.md) into Project Knowledge or custom instructions
- See [integrations/claude/](../../../integrations/claude/)

**ChatGPT (OpenAI):**
- Create Custom GPT with SKILL.md as instructions
- See [integrations/chatgpt/](../../../integrations/chatgpt/)

**Gemini (Google):**
- Add SKILL.md to system instructions
- See [integrations/gemini/](../../../integrations/gemini/)

## Clinical Accuracy Notes

### Evidence-Based Guidelines

Workflows in this skill follow current clinical practice guidelines from:
- **USPSTF** - US Preventive Services Task Force (screening recommendations)
- **ADA** - American Diabetes Association (diabetes management)
- **ACC/AHA** - American College of Cardiology / American Heart Association (hypertension, lipids)
- **CDC** - Centers for Disease Control (immunizations)
- **HL7 FHIR** - Fast Healthcare Interoperability Resources (data standards)

### Standard Code Systems

This skill uses internationally recognized medical code systems:
- **LOINC** - Logical Observation Identifiers Names and Codes (labs, vitals)
- **SNOMED CT** - Systematized Nomenclature of Medicine (diagnoses, procedures)
- **RxNorm** - Normalized drug names (medications)
- **ICD-10-CM** - International Classification of Diseases (billing diagnoses)
- **CPT** - Current Procedural Terminology (procedures, E&M codes)

### Known Limitations

1. **Guideline Variations** - Recommendations may differ between organizations (e.g., USPSTF vs ACS for colorectal screening start age)
2. **Individualization Required** - Guidelines are population-based; clinical judgment must account for individual patient factors
3. **Data Quality** - Workflows assume complete, accurate FHIR data; missing or incorrect data may lead to incomplete assessments
4. **AI Limitations** - AI agents may misinterpret complex clinical scenarios; human review is essential
5. **Regulatory Compliance** - Clinical use requires appropriate licensure, supervision, and adherence to institutional policies

### Safety Warnings

**This skill is for informational purposes only. It does not constitute medical advice.**

- Always verify patient identity before any read, create, or update operation
- Confirm AI-generated recommendations with clinical guidelines and patient context
- Never rely solely on AI for critical clinical decisions (e.g., medication dosing, diagnoses)
- Ensure appropriate clinical oversight for all AI-assisted workflows
- Comply with HIPAA, institutional policies, and local regulations
- Report errors or safety concerns immediately

## Examples

### Example 1: Patient Chart Review

**User:** "Review chart for patient John Doe, DOB 1965-03-15"

**Agent uses:**
1. `fhir_search` - Find patient by name and DOB
2. `fhir_search` - Get active conditions
3. `fhir_search` - Get active medications
4. `fhir_search` - Get recent labs (last 30 days)
5. `fhir_search` - Get recent encounters (last 90 days)
6. `fhir_search` - Get active allergies

**Output:** Structured summary with demographics, problems, meds, labs (with flags), encounters, and allergies

### Example 2: Identify Overdue Preventive Care

**User:** "What preventive care screenings is patient Jane Smith overdue for?"

**Agent uses:**
1. `fhir_read` - Get patient demographics (age, sex)
2. `fhir_search` - Get active conditions (diabetes, smoking)
3. `fhir_search` - Get previous screenings (colonoscopy, mammography, etc.)

**Output:** List of overdue screenings with last completion date and recommendation

### Example 3: Record Blood Pressure

**User:** "Record BP 142/88 for patient ID 12345"

**Agent uses:**
1. `fhir_read` - Verify patient identity
2. `fhir_create` - Create Observation resource with BP panel (systolic + diastolic components)

**Output:** Confirmation of recorded BP with interpretation (Stage 1 Hypertension)

## Testing

### Test Environment
Use a FHIR test server for initial testing:
- **HAPI FHIR:** https://hapi.fhir.org/baseR4 (public test server)
- **EPIC Sandbox:** https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4 (requires app registration)

### Test Patients
EPIC provides synthetic test patients:
- **Derrick Lin:** Patient ID `Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB`
- **Camila Lopez:** Patient ID `eg2bapHUHrr1geTcBYqPJVGFjV-9YEpNwI5pT5RBl6ikB`

Search by name to discover test patients in your FHIR server.

### Validation Checklist
- [ ] Searches return expected results
- [ ] Read operations retrieve complete resources
- [ ] Created resources appear in subsequent searches
- [ ] Updated resources reflect changes
- [ ] OAuth scopes are sufficient for all operations
- [ ] Error messages are handled gracefully
- [ ] Clinical logic produces accurate recommendations

## Contributing

We welcome contributions to improve this skill!

**How to contribute:**
1. **Clinical Enhancements** - Add new workflows, update guidelines, improve accuracy
2. **Code System Updates** - Add LOINC/SNOMED/RxNorm codes for additional resources
3. **Error Handling** - Improve guidance for common errors
4. **Testing** - Validate workflows with additional FHIR servers
5. **Documentation** - Clarify instructions, add examples

**Contribution requirements:**
- Follow HL7 FHIR R4 specifications
- Use evidence-based clinical guidelines (cite sources)
- Test with at least one FHIR server
- Include examples with expected inputs/outputs
- Maintain agent-agnostic language (no agent-specific features)

See [CONTRIBUTING.md](../../../CONTRIBUTING.md) for detailed guidelines.

## Support

**Issues:**
- Report bugs or inaccuracies: https://github.com/langcare/langcare-mcp-fhir/issues
- Include FHIR server type, resource type, and error message

**Questions:**
- GitHub Discussions: https://github.com/langcare/langcare-mcp-fhir/discussions
- Email: support@langcare.ai (for private inquiries)

**Updates:**
- Check this README for new workflows and guideline updates
- Follow repository for notifications of major changes

## License

This skill is licensed under the same license as the LangCare MCP FHIR Server. See [LICENSE](../../../LICENSE) for details.

## Authors

- **LangCare Team** - Initial workflows
- **Contributors** - See [CONTRIBUTORS.md](../../../CONTRIBUTORS.md)

## Version History

- **v1.0.0** (2026-01-25) - Initial release with 8 core clinical workflows
