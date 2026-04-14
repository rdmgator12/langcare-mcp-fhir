---
name: langcare-soap-notes
description: >
  Generates structured SOAP notes from FHIR encounter data including chief
  complaint, vitals, labs, medications, conditions, and procedures. Supports
  ambulatory, ED, and inpatient formats. Use when asked to write a SOAP note,
  generate encounter note, document this visit, create a clinic note, or
  mentions SOAP format.
---

# SOAP Note Generator

## When to Use This Skill
Use when a clinician needs a structured Subjective-Objective-Assessment-Plan note generated from FHIR encounter data.

## Clinical Workflow
1. Use `fhir_read` or `fhir_search` to retrieve the Encounter (reasonCode for chief complaint, class for setting, period, participant for provider)
2. Use `fhir_read` to retrieve Patient demographics for note header
3. Use `fhir_search` to pull vital signs (Observation category=vital-signs) during the encounter window
4. Use `fhir_search` to pull laboratory results (Observation category=laboratory) from encounter date
5. Use `fhir_search` to pull active MedicationRequest resources; separate pre-existing from new prescriptions by authoredOn
6. Use `fhir_search` to pull Condition resources for encounter diagnoses and active problem list
7. Use `fhir_search` to pull Procedure resources performed during the encounter
8. Use `fhir_search` to pull active AllergyIntolerance for allergy list
9. Assemble SOAP note with setting-appropriate format (AMB, EMER, IMP)
10. Optionally use `fhir_create` to persist as DocumentReference (LOINC 11506-3 for progress note)

## FHIR Resources
- **Encounter** -- Visit context: class, type, reasonCode, period, participant
- **Patient** -- Demographics for note header
- **Observation** -- Vitals (category=vital-signs) and labs (category=laboratory)
- **MedicationRequest** -- Current medications and new prescriptions
- **Condition** -- Encounter diagnoses and active problems (ICD-10 codes)
- **Procedure** -- Procedures performed during visit
- **AllergyIntolerance** -- Allergy list
- **DocumentReference** -- Output: persisted note

## FHIR Query Examples
### Pull Encounter
```
fhir_search(resourceType="Encounter", queryParams="patient=[patient-id]&date=[YYYY-MM-DD]&_sort=-date&_count=1")
```

### Pull Encounter Vitals
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=vital-signs&date=ge[encounter-start]&date=le[encounter-end]&_sort=-date")
```

### Pull Encounter Diagnoses
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&encounter=[encounter-id]")
```

## Clinical Guidelines
- CMS E/M documentation guidelines (2021 revision: focus on medical decision-making)
- Specialty-specific documentation requirements
- LOINC document type codes: 11506-3 (progress note), 34111-5 (ED note), 34117-2 (H&P note)

## Interpretation Guide
- Subjective: chief complaint from Encounter.reasonCode, current medications, allergies. HPI synthesized from encounter reason and active conditions relevant to the visit
- Objective: vitals in standard format (T/HR/RR/BP/SpO2/Wt/Ht/BMI), labs grouped by panel with abnormals flagged, physical exam findings from Observations with category=exam if available
- Assessment: numbered problem list with ICD-10 codes from encounter Conditions
- Plan: per-problem actions including medication changes, referrals, follow-up timeline, procedures performed
- For ED encounters: add triage acuity, time-stamped reassessments, disposition
- For inpatient encounters: add hospital day number, consults, anticipated discharge

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
