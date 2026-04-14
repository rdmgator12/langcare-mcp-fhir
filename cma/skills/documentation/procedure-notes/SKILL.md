---
name: langcare-procedure-notes
description: >
  Generates structured procedure note templates from FHIR data including
  pre-procedure assessment, procedure details, findings, complications,
  and post-procedure plan. Use when asked to write a procedure note,
  create an operative note, document a procedure, or generate a
  procedural documentation template.
---

# Procedure Note Template

## When to Use This Skill
Use when a clinician needs a structured procedure note populated with patient context from FHIR resources.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics and allergies
2. Use `fhir_search` to pull the Procedure resource for the completed procedure (code, performedDateTime, outcome)
3. Use `fhir_search` to pull pre-procedure Observation resources (vitals, relevant labs, coagulation studies)
4. Use `fhir_search` to pull active Condition resources for relevant comorbidities and procedure indication
5. Use `fhir_search` to pull MedicationRequest for anesthesia/sedation medications administered
6. Assemble procedure note: patient identification, procedure name/CPT, indication, consent status, anesthesia type, pre-procedure assessment, procedure description (technique, findings), specimens collected, complications, estimated blood loss, post-procedure condition, and plan
7. Optionally use `fhir_create` to persist as DocumentReference (LOINC 28570-0 for procedure note)

## FHIR Resources
- **Patient** -- Demographics, allergies
- **Procedure** -- Procedure details: code (CPT/SNOMED), status, performedDateTime, outcome, complication
- **Observation** -- Pre-procedure vitals and labs
- **Condition** -- Procedure indication and comorbidities
- **MedicationRequest** / **MedicationAdministration** -- Anesthesia, sedation, antibiotics
- **DocumentReference** -- Output: procedure note

## FHIR Query Examples
### Pull Procedure Details
```
fhir_read(resourceType="Procedure", id="[procedure-id]")
```

### Pull Pre-Procedure Labs
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=laboratory&date=ge[procedure-date-minus-7d]&date=le[procedure-date]&_sort=-date")
```

## Clinical Guidelines
- CMS operative report documentation requirements
- Joint Commission immediate post-procedure documentation standards
- Specialty-specific procedure note elements (ACS, ASGE, ACC)

## Interpretation Guide
- Structure: Date/Time, Patient ID, Procedure (name + CPT code), Surgeon/Proceduralist, Assistant(s), Anesthesia (type + provider), Indication, Consent (confirmed), Pre-procedure assessment (vitals, relevant labs), Technique (step-by-step), Findings, Specimens (sent to pathology), Complications (none or describe), EBL, Post-procedure condition (stable/ICU/etc.), Post-procedure orders, Follow-up plan
- For endoscopic procedures: include scope type, insertion/withdrawal times, visualization quality, biopsy sites
- For surgical procedures: include incision type, exposure, key operative findings, closure method, drain placement

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
