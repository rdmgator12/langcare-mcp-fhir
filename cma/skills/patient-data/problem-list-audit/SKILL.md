---
name: langcare-problem-list-audit
description: >
  Audits the patient problem list for accuracy by cross-referencing active
  Conditions against medications, labs, and encounter diagnoses. Identifies
  missing diagnoses, resolved conditions still marked active, and coding
  discrepancies. Use when asked to review problem list, audit diagnoses,
  clean up problem list, or verify active conditions.
---

# Problem List Audit

## When to Use This Skill
Use when a clinician needs to validate the accuracy and completeness of a patient's active problem list by cross-referencing against medications, lab results, and encounter history.

## Clinical Workflow
1. Use `fhir_search` to pull all Condition resources (active, resolved, inactive) for the patient
2. Use `fhir_search` to pull active MedicationRequest resources -- infer expected diagnoses from medications (e.g., metformin implies diabetes, levothyroxine implies hypothyroidism)
3. Use `fhir_search` to pull recent Observation (laboratory) results -- infer conditions from abnormal lab patterns (e.g., elevated HbA1c implies diabetes, elevated TSH implies hypothyroidism)
4. Use `fhir_search` to pull recent Encounter resources and their associated Condition resources (encounter diagnoses) to find conditions documented in visits but missing from the problem list
5. Cross-reference all sources: flag medications without a supporting active problem, lab patterns without a corresponding diagnosis, encounter diagnoses not on the active problem list, and active problems with no supporting evidence in recent encounters/medications/labs

## FHIR Resources
- **Condition** -- Problem list entries with clinicalStatus, verificationStatus, code (ICD-10/SNOMED), onset, abatement
- **MedicationRequest** -- Active medications implying diagnoses
- **Observation** -- Lab results with patterns suggesting undiagnosed conditions
- **Encounter** -- Recent visits with encounter-specific diagnoses

## FHIR Query Examples
### Pull All Conditions
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&_count=200")
```

### Pull Active Medications
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

### Pull Recent Encounter Diagnoses
```
fhir_search(resourceType="Encounter", queryParams="patient=[patient-id]&date=ge[12-months-ago]&_sort=-date&_count=50")
```

## Clinical Guidelines
- CMS requires accurate problem list maintenance for quality reporting and risk adjustment (HCC coding)
- ICD-10 coding specificity requirements: prefer specific codes (E11.65 T2DM with hyperglycemia) over unspecified (E11.9)
- SNOMED CT preferred for clinical documentation; ICD-10-CM required for billing

## Interpretation Guide
- Categorize findings as: Missing Diagnoses (medication/lab evidence but no problem list entry), Potentially Resolved (active condition with no supporting evidence in 12+ months), Coding Gaps (non-specific codes that should be updated), and Duplicate Entries (same condition coded multiple ways)
- Rank findings by clinical impact: HCC-relevant conditions first, then safety-relevant, then administrative
- For each finding, provide the evidence source and recommended action (add, resolve, update code, merge duplicates)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
