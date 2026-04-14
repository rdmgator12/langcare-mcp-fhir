---
name: langcare-opioid-risk
description: >
  Performs opioid risk assessment by calculating morphine milligram equivalents
  (MME), applying Opioid Risk Tool scoring, and evaluating CDC 2022 guideline
  thresholds. Flags concurrent benzodiazepines, missing naloxone, and high-MME
  prescriptions. Use when asked to assess opioid risk, calculate MME, check
  opioid safety, or review opioid prescriptions.
---

# Opioid Risk Assessment

## When to Use This Skill
Use when a clinician needs to evaluate opioid prescribing safety including total daily MME calculation, concurrent CNS depressant screening, ORT scoring, and CDC guideline compliance.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender for ORT scoring)
2. Use `fhir_search` to pull all active MedicationRequest resources; classify as opioids, benzodiazepines, gabapentinoids, muscle relaxants, Z-drugs, naloxone
3. Calculate total daily MME for each opioid using dose x frequency x MME conversion factor (see references/mme-conversion-table.md)
4. Evaluate against CDC 2022 thresholds: <20 lower risk, 20-49 moderate, 50-89 caution, >=90 high risk
5. Use `fhir_search` to pull Condition resources for ORT scoring: substance use history, psychiatric diagnoses
6. Apply Opioid Risk Tool scoring (gender-specific criteria)
7. Use `fhir_search` to check for recent urine drug screen (LOINC 19295-5) and pain assessments (LOINC 72514-3)
8. Assess naloxone co-prescription status; recommend if MME >= 50 or concurrent CNS depressants

## FHIR Resources
- **MedicationRequest** -- Active opioid and concurrent prescriptions with dosing details
- **MedicationDispense** -- Fill history for pattern analysis
- **Patient** -- Age and gender for ORT scoring
- **Condition** -- Psychiatric diagnoses, substance use history for ORT
- **Observation** -- Urine drug screen results, pain scores

## FHIR Query Examples
### Pull Active Medications
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

### Pull Substance Use History
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&_count=100")
```

### Pull Urine Drug Screen
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=19295-5,3426-4&_sort=-date&_count=5")
```

## Clinical Guidelines
- CDC 2022 Clinical Practice Guideline for Prescribing Opioids for Pain
- FDA REMS for extended-release/long-acting opioids
- PDMP (Prescription Drug Monitoring Program) check recommended before each opioid prescription

## Interpretation Guide
- Total daily MME thresholds: <50 standard monitoring, 50-89 prescribe naloxone and reassess, >=90 avoid or taper
- Concurrent benzodiazepine + opioid = FDA Boxed Warning for respiratory depression
- ORT scores: 0-3 low risk, 4-7 moderate risk, >=8 high risk (scores differ by gender)
- UDS consistency: presence of prescribed opioid confirms adherence; absence suggests diversion; non-prescribed substances flag polysubstance risk
- Include CAGE-AID screening questions for clinician to administer

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
