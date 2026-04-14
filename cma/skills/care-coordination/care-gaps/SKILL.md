---
name: langcare-care-gaps
description: >
  Identifies care gaps for an individual patient by checking overdue preventive
  screenings, missing chronic disease monitoring, and unmet quality measure
  criteria from FHIR data. Use when asked about care gaps, overdue screenings,
  missing preventive care, what is this patient due for, or patient care
  compliance check.
---

# Care Gap Identifier

## When to Use This Skill
Use when a clinician needs to identify overdue preventive screenings, missing chronic disease monitoring labs, and unmet quality measure criteria for an individual patient.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender for screening eligibility)
2. Use `fhir_search` to pull active Condition resources for chronic disease monitoring requirements
3. Use `fhir_search` to pull Observation resources (labs, screenings) to check completion dates
4. Use `fhir_search` to pull Procedure and DiagnosticReport resources for cancer screenings (mammogram, colonoscopy, Pap)
5. Use `fhir_search` to pull Immunization resources for vaccination status
6. Apply USPSTF/ACS/ACIP screening schedules based on age, gender, and risk factors (see references/uspstf-recommendations.md)
7. For each applicable screening/monitoring item: determine if it was completed within the required interval; flag as current, due soon, or overdue
8. Present prioritized list of care gaps with recommended actions

## FHIR Resources
- **Patient** -- Age, gender for screening eligibility
- **Condition** -- Chronic diseases requiring monitoring (DM, HTN, CKD)
- **Observation** -- Lab results, screening scores (PHQ-9, AUDIT-C)
- **Procedure** -- Screening procedures (colonoscopy, mammogram)
- **DiagnosticReport** -- Screening results
- **Immunization** -- Vaccination history

## FHIR Query Examples
### Pull Recent Screenings
```
fhir_search(resourceType="Procedure", queryParams="patient=[patient-id]&date=ge[3-years-ago]&_sort=-date")
```

### Check Mammogram Status
```
fhir_search(resourceType="DiagnosticReport", queryParams="patient=[patient-id]&code=24606-6&date=ge[27-months-ago]")
```

## Clinical Guidelines
- USPSTF A and B grade recommendations
- ACS cancer screening guidelines
- ACIP immunization schedule
- ADA diabetes monitoring standards
- ACC/AHA lipid and BP monitoring guidelines

## Interpretation Guide
- Categorize gaps as: Overdue (past recommended interval), Due Now (within recommended window), Upcoming (due within 3 months)
- Prioritize by clinical impact: cancer screenings, chronic disease monitoring, immunizations, general wellness
- For each gap: state the screening/test, when it was last performed (or never), the recommended interval, and specific ordering action
- Include age-specific screening eligibility cutoffs (e.g., colonoscopy 45-75, mammogram 50-74, lung cancer screening 50-80 with 20+ pack-year history)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
