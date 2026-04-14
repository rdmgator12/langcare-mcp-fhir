---
name: langcare-discharge-summaries
description: >
  Generates comprehensive discharge summary documents from FHIR data including
  admission diagnosis, hospital course, procedures, discharge medications,
  follow-up instructions, and pending results. Use when asked to write a
  discharge summary, generate discharge documentation, or create a hospital
  discharge note.
---

# Discharge Summary Writer

## When to Use This Skill
Use when a clinician needs a comprehensive discharge summary document compiled from the patient's hospital course in FHIR.

## Clinical Workflow
1. Use `fhir_search` to retrieve the Encounter (admission date, discharge date, reason, attending)
2. Use `fhir_read` to retrieve Patient demographics
3. Use `fhir_search` to pull Condition resources: admitting diagnosis, discharge diagnoses, complications
4. Use `fhir_search` to pull Procedure resources performed during the hospitalization
5. Use `fhir_search` to pull active MedicationRequest for discharge medications; compare against admission medications for reconciliation summary
6. Use `fhir_search` to pull key Observation results (admission labs, discharge labs, significant findings)
7. Use `fhir_search` to pull ServiceRequest resources for pending results and referrals
8. Use `fhir_search` to pull scheduled Appointments for follow-up
9. Assemble discharge summary: demographics, admission/discharge dates, admitting diagnosis, hospital course, procedures, discharge diagnoses, discharge medications, follow-up, pending results, patient instructions
10. Use `fhir_create` to persist as DocumentReference (LOINC 18842-5)

## FHIR Resources
- **Encounter** -- Admission/discharge dates, reason, attending
- **Patient** -- Demographics
- **Condition** -- Admitting and discharge diagnoses
- **Procedure** -- Inpatient procedures
- **MedicationRequest** -- Discharge medications
- **Observation** -- Key lab results
- **Appointment** -- Follow-up appointments
- **ServiceRequest** -- Pending results, referrals
- **DocumentReference** -- Output: discharge summary

## FHIR Query Examples
### Pull Encounter
```
fhir_read(resourceType="Encounter", id="[encounter-id]")
```

### Pull Procedures During Admission
```
fhir_search(resourceType="Procedure", queryParams="patient=[patient-id]&encounter=[encounter-id]")
```

## Clinical Guidelines
- CMS Conditions of Participation require discharge summary within 30 days
- Joint Commission: discharge summary must include reason for admission, significant findings, procedures, discharge condition, discharge medications, follow-up instructions, pending results
- TJC NPSG.02.03.01: communicate pending test results to outpatient provider

## Interpretation Guide
- Hospital course: narrative summary of key events, procedures, complications, and clinical response to treatment
- Discharge medications: full list with dose, frequency, route; highlight NEW medications and CHANGED medications vs pre-admission
- Pending results: list any results not yet available at discharge with responsible provider for follow-up
- Follow-up: specific provider names, dates, and purpose
- Patient instructions: diagnosis-specific education, activity restrictions, warning signs for ED return, diet modifications

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
