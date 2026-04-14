---
name: langcare-critical-values
description: >
  Detects and alerts on critical laboratory values requiring immediate clinical
  action per CAP/CLIA thresholds. Generates structured critical value
  notifications with recommended interventions. Use when asked to check for
  critical labs, critical value alerts, panic values, stat lab review, or
  when monitoring for dangerous lab results.
---

# Critical Value Alert Generator

## When to Use This Skill
Use when a clinician needs to identify laboratory results that fall outside critical (panic) value thresholds requiring immediate notification and intervention.

## Clinical Workflow
1. Use `fhir_search` to pull recent Observation resources (category=laboratory) with interpretation codes HH or LL, or status=preliminary for pending results
2. Compare each result against CAP/CLIA critical value thresholds (see references/cap-clia-thresholds.md)
3. For each critical value: identify the analyte, current value, threshold breached, and timestamp
4. Use `fhir_search` to pull active MedicationRequest resources to identify contributing medications
5. Use `fhir_search` to pull active Condition resources for clinical context
6. Generate structured alert with: critical value details, probable cause, recommended immediate action, and monitoring plan

## FHIR Resources
- **Observation** -- Lab results with interpretation codes (HH, LL, H, L, A) and referenceRange
- **MedicationRequest** -- Active medications that may contribute to the critical value
- **Condition** -- Active conditions providing clinical context

## FHIR Query Examples
### Pull Labs with Critical Interpretation
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=laboratory&_sort=-date&_count=200")
```

### Pull Pending Results
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=laboratory&status=preliminary,registered")
```

## Clinical Guidelines
- CAP (College of American Pathologists) critical value reporting requirements
- CLIA '88 regulations for immediate notification of life-threatening results
- Joint Commission NPSG.02.03.01: Report critical results in a timely manner

## Interpretation Guide
- Critical values require immediate clinician notification and read-back confirmation
- For each critical value: state the value, the critical threshold, the clinical risk, and recommended immediate intervention
- Common critical values and immediate actions: K+ <2.5 or >6.5 (ECG, replacement/Kayexalate), Na <120 or >160 (fluid management, correct slowly), glucose <40 or >500 (D50/insulin), Hgb <7 (transfusion threshold), platelets <20 (bleeding risk, transfusion), INR >5 (hold warfarin, vitamin K), troponin elevated (ACS protocol)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
