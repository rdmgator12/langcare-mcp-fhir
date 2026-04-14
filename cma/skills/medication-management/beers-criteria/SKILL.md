---
name: langcare-beers-criteria
description: >
  Screens medications against the AGS Beers Criteria for potentially
  inappropriate medication use in older adults (age 65+). Identifies
  medications to avoid, dose adjustments needed, and drug-disease
  interactions specific to geriatric patients. Use when asked about
  Beers criteria, inappropriate medications in elderly, geriatric
  medication review, or STOPP-START criteria.
---

# Beers Criteria Medication Review

## When to Use This Skill
Use when reviewing medications for patients age 65 and older to identify potentially inappropriate prescriptions per AGS Beers Criteria 2023.

## Clinical Workflow
1. Use `fhir_read` to confirm patient age >= 65 from Patient.birthDate
2. Use `fhir_search` to pull all active MedicationRequest resources
3. Use `fhir_search` to pull active Condition resources for drug-disease interaction screening
4. Use `fhir_search` to pull recent Observation resources (renal function for dose adjustments)
5. Screen each medication against Beers Criteria categories: medications to avoid regardless of diagnosis, medications to avoid with specific conditions, medications to use with caution, drug-drug interactions to avoid, dose adjustments based on renal function
6. Present findings with Beers category, quality of evidence, strength of recommendation, and safer alternatives

## FHIR Resources
- **Patient** -- Age verification (must be >= 65)
- **MedicationRequest** -- Active medications to screen
- **Condition** -- Active diagnoses for drug-disease interaction screening
- **Observation** -- Renal function (eGFR/CrCl) for dose-based criteria

## FHIR Query Examples
### Pull Active Medications
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

### Pull Active Conditions
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&clinical-status=active")
```

### Pull Renal Function
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=33914-3&_sort=-date&_count=1")
```

## Clinical Guidelines
- AGS Beers Criteria 2023 (American Geriatrics Society)
- STOPP/START Criteria v2 (Screening Tool of Older Persons' Prescriptions / Screening Tool to Alert to Right Treatment)
- CMS Part D medication therapy management requirements

## Interpretation Guide
- Classify findings by Beers category: Table 2 (avoid regardless), Table 3 (avoid with specific conditions), Table 4 (use with caution), Table 5 (drug-drug interactions), Table 6 (dose adjustment for renal function)
- Rate evidence quality: High, Moderate, Low, Very Low
- Rate recommendation strength: Strong Avoid, Conditional Avoid, Use with Caution
- For each flagged medication, provide: the Beers concern, the clinical rationale, safer alternative options, and whether discontinuation requires a taper

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
