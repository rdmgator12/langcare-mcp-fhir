---
name: langcare-preventive-care-compliance
description: >
  Generates a comprehensive preventive care compliance report for a patient
  or population based on USPSTF A/B grade recommendations, ACS screening
  guidelines, and ACIP immunization schedules. Identifies overdue screenings
  with age/sex/risk-appropriate criteria. Use when asked about preventive care
  compliance, wellness checkup, annual screening status, or health maintenance.
---

# Preventive Care Compliance Report

## When to Use This Skill
Use when a clinician needs a comprehensive preventive care status check covering cancer screenings, metabolic screenings, immunizations, and behavioral health assessments per guideline-recommended intervals.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender, smoking status for eligibility)
2. Use `fhir_search` to pull Condition resources for risk factors modifying screening eligibility
3. Use `fhir_search` to pull Procedure and DiagnosticReport resources for screening completion (mammogram, colonoscopy, LDCT, DEXA)
4. Use `fhir_search` to pull Observation resources for lab screenings (A1c, lipids, hepatitis, HIV, PHQ-9, AUDIT-C)
5. Use `fhir_search` to pull Immunization resources for vaccination status
6. Apply USPSTF grade A/B recommendations by age, sex, and risk factors (see references/uspstf-grades.md)
7. Present compliance checklist: screening name, eligibility, last performed, interval, status (current/due/overdue/not applicable), next due date

## FHIR Resources
- **Patient** -- Age, gender for screening eligibility
- **Condition** -- Risk factors modifying eligibility (cancer history, family history, smoking)
- **Procedure** / **DiagnosticReport** -- Screening procedures
- **Observation** -- Lab-based screenings and behavioral assessments
- **Immunization** -- Vaccination history

## FHIR Query Examples
### Pull Recent Mammogram
```
fhir_search(resourceType="DiagnosticReport", queryParams="patient=[patient-id]&code=24606-6&_sort=-date&_count=1")
```

### Pull Smoking Status
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=72166-2&_sort=-date&_count=1")
```

## Clinical Guidelines
- USPSTF A and B grade recommendations (see references/uspstf-grades.md)
- ACS cancer screening guidelines
- ACIP immunization schedules
- ADA diabetes screening guidelines

## Interpretation Guide
- Color-code status: current (green), due within 3 months (yellow), overdue (red), not applicable (gray)
- Group by category: cancer screenings, metabolic/cardiovascular, infectious disease, behavioral health, immunizations
- For each overdue item: state last date performed (or never), recommended interval, and ordering action
- Include age-specific transitions: when to start/stop each screening
- Flag risk-based modifications: enhanced screening for family history, smoking history, prior abnormal results

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
