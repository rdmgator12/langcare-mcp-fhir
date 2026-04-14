---
name: langcare-pneumonia-curb65
description: >
  Assesses pneumonia severity using CURB-65 and PSI/PORT scores to guide
  disposition (outpatient vs inpatient vs ICU). Pulls vitals, labs, and
  imaging from FHIR data. Use when asked about pneumonia severity, CURB-65
  score, pneumonia disposition, community-acquired pneumonia management,
  or PORT score calculation.
---

# Pneumonia Severity Assessment (CURB-65)

## When to Use This Skill
Use when a clinician needs to determine pneumonia severity and appropriate level of care using validated scoring tools.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age for scoring)
2. Use `fhir_search` to pull vital signs: BP (8480-6), respiratory rate (9279-1), temperature (8310-5), heart rate (8867-4), SpO2 (2708-6)
3. Use `fhir_search` to pull labs: BUN (3094-0), WBC (6690-2), creatinine (2160-0), glucose (2345-7), pH (2744-1), PaO2 (2703-7), albumin (1751-7), sodium (2951-2), hematocrit (4544-3)
4. Use `fhir_search` to pull active Condition resources for comorbidities and confirmed pneumonia diagnosis
5. Use `fhir_search` to pull DiagnosticReport for chest imaging results
6. Calculate CURB-65 score (see references/curb65-scoring.md)
7. If data sufficient, also calculate PSI/PORT score for more granular risk stratification
8. Present severity classification with disposition recommendation and antibiotic guidance

## FHIR Resources
- **Patient** -- Age, gender
- **Observation** -- Vitals and labs for scoring
- **Condition** -- Pneumonia diagnosis, comorbidities (CHF, liver disease, renal disease, cerebrovascular disease, malignancy)
- **DiagnosticReport** -- Chest X-ray or CT findings (pleural effusion, multilobar involvement)

## FHIR Query Examples
### Pull Vitals
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=vital-signs&_sort=-date&_count=10")
```

### Pull BUN
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=3094-0&_sort=-date&_count=1")
```

## Clinical Guidelines
- ATS/IDSA 2019 Community-Acquired Pneumonia Guidelines
- BTS Guidelines for CAP (CURB-65 origin)
- IDSA/ATS criteria for severe CAP requiring ICU admission

## Interpretation Guide
- CURB-65: 0-1 outpatient treatment, 2 consider short inpatient or supervised outpatient, 3-5 inpatient (4-5 consider ICU)
- PSI/PORT: Class I-III outpatient, Class IV inpatient, Class V ICU consideration
- ATS/IDSA major criteria for ICU: septic shock requiring vasopressors, mechanical ventilation
- ATS/IDSA minor criteria (>=3 = ICU): RR >=30, PaO2/FiO2 <=250, multilobar infiltrates, confusion, BUN >=20, WBC <4000, platelets <100K, temperature <36C, hypotension requiring aggressive fluids

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
