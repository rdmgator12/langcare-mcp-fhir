---
name: langcare-medication-adherence
description: >
  Assesses medication adherence by analyzing prescription fill patterns from
  MedicationDispense resources, calculating proportion of days covered (PDC),
  and identifying gaps in therapy. Use when asked about medication adherence,
  compliance, fill history, refill gaps, PDC, or whether a patient is taking
  their medications as prescribed.
---

# Medication Adherence Assessment

## When to Use This Skill
Use when a clinician needs to evaluate whether a patient is taking medications as prescribed by analyzing fill history, calculating adherence metrics, and identifying therapy gaps.

## Clinical Workflow
1. Use `fhir_search` to pull active MedicationRequest resources for the patient's prescribed medication regimen
2. Use `fhir_search` to pull MedicationDispense resources for fill history with dates and quantities
3. Calculate Proportion of Days Covered (PDC) for each medication: PDC = (days supply dispensed / days in period) x 100
4. Identify refill gaps: periods where days supply ran out before next fill
5. Cross-reference adherence findings with clinical outcomes (e.g., uncontrolled A1c with poor metformin PDC, elevated BP with antihypertensive gaps)
6. Present adherence report with per-medication PDC, gap analysis, and clinical impact assessment

## FHIR Resources
- **MedicationRequest** -- Active prescriptions defining the intended regimen
- **MedicationDispense** -- Fill history: medicationCodeableConcept, quantity, daysSupply, whenHandedOver, status
- **Observation** -- Clinical outcomes (A1c, BP, LDL) correlated with adherence

## FHIR Query Examples
### Pull Fill History
```
fhir_search(resourceType="MedicationDispense", queryParams="patient=[patient-id]&_sort=-whenhandedover&_count=200")
```

### Pull Active Prescriptions
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

## Clinical Guidelines
- CMS Star Ratings use PDC >= 80% as the adherence threshold for diabetes, hypertension, and cholesterol medications
- WHO defines five dimensions of adherence: social/economic, health system, condition-related, therapy-related, patient-related
- NCQA HEDIS medication adherence measures: ADM (antidepressants), SPC (statins), PBD (beta-blockers post-discharge)

## Interpretation Guide
- PDC >= 80%: adherent. PDC 50-79%: partially adherent. PDC < 50%: non-adherent
- Calculate PDC over the most recent 12 months or since therapy start, whichever is shorter
- Flag early refills (>7 days early) which may indicate stockpiling or diversion
- Flag medications with no fills despite active prescription (never-filled prescriptions)
- Correlate adherence gaps with clinical events: ED visits, hospitalizations, worsening lab values

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
