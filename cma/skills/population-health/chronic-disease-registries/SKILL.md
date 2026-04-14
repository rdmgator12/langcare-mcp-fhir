---
name: langcare-chronic-disease-registries
description: >
  Queries and reports on chronic disease registries including diabetes, HTN,
  CHF, COPD, CKD, and asthma populations with severity distribution, control
  rates, and patients needing intervention. Use when asked about disease
  registries, chronic disease population, diabetes registry, HTN registry,
  disease-specific panel report, or population with specific condition.
---

# Chronic Disease Registry Query

## When to Use This Skill
Use when a clinician or care manager needs a disease-specific population report with severity stratification, control metrics, and patients needing escalation.

## Clinical Workflow
1. Use `fhir_search` to pull all patients with the target Condition (e.g., diabetes, HTN, CHF) active
2. For each patient, use `fhir_search` to pull key outcome Observations (A1c for diabetes, BP for HTN, EF for CHF, FEV1 for COPD, eGFR for CKD)
3. Stratify by severity/control: well-controlled, moderately controlled, poorly controlled, unknown (no recent data)
4. Use `fhir_search` to pull active MedicationRequest for each patient to assess therapy appropriateness
5. Identify patients needing intervention: above target despite therapy, no recent monitoring, on suboptimal regimen
6. Present registry report with population counts, severity distribution, and actionable patient lists

## FHIR Resources
- **Condition** -- Registry population identification
- **Patient** -- Demographics for stratification
- **Observation** -- Disease-specific outcome measures
- **MedicationRequest** -- Current therapy assessment

## FHIR Query Examples
### Pull Diabetes Registry
```
fhir_search(resourceType="Condition", queryParams="code=http://snomed.info/sct|44054006&clinical-status=active&_count=500")
```

### Pull A1c Values for Registry Patients
```
fhir_search(resourceType="Observation", queryParams="code=http://loinc.org|4548-4&date=ge[6-months-ago]&_sort=-date&_count=500")
```

## Clinical Guidelines
- PCMH population management standards
- Disease-specific guidelines (ADA for diabetes, AHA/ACC for CHF, GOLD for COPD, KDIGO for CKD)
- CMS Chronic Care Management requirements

## Interpretation Guide
- Registry dashboard format: total patients, severity tiers (by control metric), therapy distribution, patients needing action
- For diabetes registry: stratify by A1c (<7, 7-8, 8-9, >9, no recent A1c); report statin and ACEi/ARB rates
- For HTN registry: stratify by last BP (<130/80, 130-140/80-90, >140/90, no recent BP); report medication class distribution
- For CHF registry: stratify by EF (HFrEF <40%, HFmrEF 40-49%, HFpEF >=50%); report GDMT utilization (ACEi/ARB/ARNI, beta-blocker, MRA, SGLT2i)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
