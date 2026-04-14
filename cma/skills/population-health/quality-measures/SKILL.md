---
name: langcare-quality-measures
description: >
  Calculates HEDIS-style quality measures from FHIR data including
  denominator/numerator/exclusion logic, measure rates, gap-to-goal analysis,
  and non-compliant patient identification. Use when asked to calculate
  quality measures, HEDIS rates, quality dashboard, star rating, measure
  compliance, or CMS quality scores.
---

# Quality Measure Dashboard

## When to Use This Skill
Use when a clinician or quality team needs formal HEDIS/CMS quality measure calculations with denominator/numerator logic, benchmark comparison, and actionable gap lists.

## Clinical Workflow
1. Select measures to calculate (default core set: A1c control, BP control, breast/colorectal cancer screening, depression screening, statin therapy)
2. For each measure, use `fhir_search` to build the denominator population from Condition/Patient resources per measure specification
3. Use `fhir_search` to identify and remove excluded patients (hospice, ESRD, bilateral mastectomy, etc.)
4. Use `fhir_search` to query numerator evidence: Observation (lab results), Procedure (screenings), DiagnosticReport (imaging)
5. Compute rates: Rate = Numerator / (Denominator - Exclusions) x 100
6. Compare against benchmarks (see references/hedis-measures.md)
7. Identify non-compliant patients (in denominator but not in numerator) as care gaps
8. Present dashboard with rates, targets, gap counts, and patient-level gap lists

## FHIR Resources
- **Condition** -- Denominator identification (diabetes, HTN diagnoses)
- **Patient** -- Demographics for age/sex stratification
- **Observation** -- Lab results (A1c, BP), screening scores (PHQ-9)
- **Procedure** -- Screening procedures (colonoscopy, mammogram)
- **DiagnosticReport** -- Screening results
- **MedicationRequest** -- Medication-based measures (statin therapy)
- **Encounter** -- Visit-based eligibility, hospice exclusions
- **Immunization** -- Immunization measures

## FHIR Query Examples
### Pull Diabetes Patients (Denominator)
```
fhir_search(resourceType="Condition", queryParams="code=http://snomed.info/sct|44054006,http://snomed.info/sct|46635009&clinical-status=active&_count=500")
```

### Pull A1c Results (Numerator)
```
fhir_search(resourceType="Observation", queryParams="code=http://loinc.org|4548-4&date=ge[measurement-year-start]&_sort=-date&_count=500")
```

### Pull Mammograms (Numerator)
```
fhir_search(resourceType="DiagnosticReport", queryParams="code=http://loinc.org|24606-6&date=ge[27-months-ago]&_count=500")
```

## Clinical Guidelines
- NCQA HEDIS Technical Specifications
- CMS Star Ratings methodology
- CMS Quality Payment Program (MIPS) measures

## Interpretation Guide
- Present measures in dashboard format: measure name, denominator, exclusions, numerator, rate, target, gap-to-goal
- Star rating estimate based on measure performance vs. CMS cut points
- Non-compliant patient lists sorted by: measures with most gaps first, then by longest time since last compliant event
- For each gap patient: patient ID, last relevant result (value and date), recommended action (order test, schedule screening, intensify therapy)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
