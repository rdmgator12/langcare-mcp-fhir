---
name: langcare-history-physical
description: >
  Generates a comprehensive History and Physical (H&P) document from FHIR
  resources for hospital admissions or initial consultations. Includes HPI,
  PMH, surgical history, medications, allergies, social/family history,
  ROS, physical exam, labs, assessment, and plan. Use when asked to write
  an H&P, admission note, history and physical, or initial consultation note.
---

# History and Physical Generator

## When to Use This Skill
Use when a clinician needs a comprehensive H&P document for hospital admission or initial specialist consultation.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics, emergency contacts
2. Use `fhir_search` to retrieve the admitting Encounter (reasonCode, class=IMP)
3. Use `fhir_search` to pull all Condition resources: active problems (PMH), encounter diagnoses (admitting diagnosis)
4. Use `fhir_search` to pull Procedure resources for surgical history
5. Use `fhir_search` to pull active MedicationRequest and MedicationStatement for home medications
6. Use `fhir_search` to pull AllergyIntolerance resources
7. Use `fhir_search` to pull Observation resources: social history (category=social-history), vitals, labs
8. Use `fhir_search` to pull FamilyMemberHistory resources for family history
9. Assemble H&P in standard format: CC, HPI, PMH, PSH, Medications, Allergies, Social Hx, Family Hx, ROS, Physical Exam, Labs/Studies, Assessment, Plan
10. Optionally use `fhir_create` to persist as DocumentReference (LOINC 34117-2)

## FHIR Resources
- **Encounter** -- Admission context
- **Patient** -- Demographics
- **Condition** -- Active problems (PMH) and admitting diagnoses
- **Procedure** -- Surgical history
- **MedicationRequest** / **MedicationStatement** -- Medications
- **AllergyIntolerance** -- Allergies
- **Observation** -- Social history (smoking, alcohol, drugs), vitals, labs
- **FamilyMemberHistory** -- Family medical history
- **DocumentReference** -- Output: H&P document

## FHIR Query Examples
### Pull Surgical History
```
fhir_search(resourceType="Procedure", queryParams="patient=[patient-id]&_sort=-date&_count=50")
```

### Pull Social History
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=social-history")
```

### Pull Family History
```
fhir_search(resourceType="FamilyMemberHistory", queryParams="patient=[patient-id]")
```

## Clinical Guidelines
- CMS H&P documentation requirements for hospital admissions
- Joint Commission medical record standards
- Specialty-specific H&P templates (surgical, medical, psychiatric)

## Interpretation Guide
- Structure sections in standard medical H&P order
- PMH: organize by system (cardiovascular, endocrine, neurologic, etc.)
- PSH: list procedures with approximate dates
- Social history: include tobacco (pack-years), alcohol (drinks/week), drug use, occupation, living situation, advance directives
- Family history: organize by first-degree relative with conditions and age of onset/death
- Assessment: primary admitting diagnosis first, then secondary diagnoses numbered by priority
- Plan: per-problem approach with immediate orders, workup plan, and anticipated hospital course

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
