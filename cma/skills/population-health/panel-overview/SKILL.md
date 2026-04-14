---
name: langcare-panel-overview
description: >
  Generates a patient panel overview with aggregate statistics including panel
  size, demographics, disease prevalence, risk stratification, and utilization
  patterns from FHIR resources. Use when asked for panel overview, patient
  panel stats, practice summary, panel demographics, or caseload analysis.
---

# Patient Panel Overview

## When to Use This Skill
Use when a clinician or practice manager needs aggregate statistics about a patient panel including size, demographics, disease prevalence, and risk distribution.

## Clinical Workflow
1. Use `fhir_search` to pull Patient resources for the panel (filtered by managing organization, practitioner, or other criteria)
2. Aggregate demographics: age distribution, gender distribution, insurance mix
3. Use `fhir_search` to pull Condition resources across the panel for disease prevalence (diabetes, HTN, CHF, COPD, CKD, depression)
4. Use `fhir_search` to pull recent Encounter resources for utilization patterns (visit frequency, ED visits, hospitalizations)
5. Use `fhir_search` to pull recent Observation resources for key outcome metrics (A1c distribution, BP control rates)
6. Present dashboard with panel demographics, top diagnoses by prevalence, risk stratification tiers, and utilization summary

## FHIR Resources
- **Patient** -- Panel demographics: birthDate, gender, managing organization
- **Condition** -- Disease prevalence across panel
- **Encounter** -- Visit utilization patterns
- **Observation** -- Outcome metrics (A1c, BP)

## FHIR Query Examples
### Pull Panel Patients
```
fhir_search(resourceType="Patient", queryParams="general-practitioner=[practitioner-id]&_count=500")
```

### Pull Diabetes Prevalence
```
fhir_search(resourceType="Condition", queryParams="code=http://snomed.info/sct|44054006&clinical-status=active&_count=500")
```

## Clinical Guidelines
- PCMH (Patient-Centered Medical Home) panel management standards
- CMS TCM and Chronic Care Management documentation
- Population health management best practices

## Interpretation Guide
- Present demographics as counts and percentages: age bands (0-17, 18-39, 40-64, 65+), gender, insurance type (commercial, Medicare, Medicaid, uninsured)
- Disease prevalence as percentage of panel with active diagnosis
- Risk stratification: low (0-1 chronic conditions), moderate (2-3), high (4+), rising risk (2+ ED visits in 6 months)
- Utilization: average visits per patient per year, ED visit rate, hospitalization rate, no-show rate

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
