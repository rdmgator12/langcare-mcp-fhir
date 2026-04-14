---
name: langcare-progress-notes
description: >
  Generates daily inpatient progress notes from FHIR data including overnight
  events, current vitals, labs, I/O, medication changes, and updated
  assessment/plan. Use when asked to write a progress note, daily note,
  inpatient rounding note, or hospital day update.
---

# Progress Note Writer

## When to Use This Skill
Use when a clinician needs a daily inpatient progress note summarizing the patient's hospital course, overnight events, current clinical status, and updated plan.

## Clinical Workflow
1. Use `fhir_search` to retrieve the current inpatient Encounter (calculate hospital day number)
2. Use `fhir_read` to retrieve Patient demographics for note header
3. Use `fhir_search` to pull Observation resources (vitals from last 24h with trends, overnight extremes)
4. Use `fhir_search` to pull Observation resources (labs from last 24h)
5. Use `fhir_search` to pull active MedicationRequest with changes from last 24h (compare authoredOn)
6. Use `fhir_search` to pull active Condition resources for problem list
7. Use `fhir_search` to pull pending ServiceRequest resources for outstanding orders
8. Structure as: Date/HD#, Subjective (overnight events, patient-reported symptoms), Objective (vitals, I/O, labs, imaging), Assessment (updated problem list), Plan (per-problem with changes from prior day)
9. Optionally persist as DocumentReference (LOINC 11506-3)

## FHIR Resources
- **Encounter** -- Hospital day calculation from period.start
- **Patient** -- Demographics
- **Observation** -- Vitals (24h with min/max), labs (new results), I/O if available
- **MedicationRequest** -- Current medications with recent changes
- **Condition** -- Active problem list
- **ServiceRequest** -- Pending consults, tests

## FHIR Query Examples
### Pull 24h Vitals
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=vital-signs&date=ge[24-hours-ago]&_sort=-date&_count=50")
```

### Pull Today's Labs
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=laboratory&date=ge[today]&_sort=-date")
```

## Clinical Guidelines
- CMS inpatient documentation requirements
- Attending physician must document assessment/plan at least daily
- Problem-based daily progress notes preferred for complex patients

## Interpretation Guide
- Hospital Day format: "HD #X" calculated from admission date
- Subjective: overnight events (nursing calls, code events, new symptoms), patient-reported status, pain score, sleep quality, diet tolerance
- Objective: vitals (current + 24h range for T, HR, RR, BP, SpO2), labs (highlight new abnormals and improving/worsening trends), I/O balance (if available), imaging results
- Assessment: numbered active problem list with brief status update per problem (improving, stable, worsening)
- Plan: per-problem updates with specific orders, medication changes, pending results, and anticipated next steps

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
