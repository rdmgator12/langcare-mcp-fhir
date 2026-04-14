---
name: langcare-referrals
description: >
  Generates specialist referral requests by compiling relevant clinical data
  from FHIR resources into a structured referral package. Creates ServiceRequest
  resources for the referral order. Use when asked to generate a referral,
  specialist consult request, create a referral, or compile referral data
  for a specialist.
---

# Referral Generator

## When to Use This Skill
Use when a clinician needs to create a specialist referral with supporting clinical documentation compiled from the patient's FHIR record.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics and insurance (Coverage) for referral header
2. Use `fhir_search` to pull relevant Condition resources for the referral indication
3. Use `fhir_search` to pull relevant Observation resources (labs, vitals) supporting the referral
4. Use `fhir_search` to pull active MedicationRequest resources for current treatment context
5. Use `fhir_search` to pull relevant Procedure resources for prior interventions
6. Compile referral package: demographics, insurance, referring diagnosis, clinical question, relevant labs/vitals, current medications, prior treatments attempted
7. Use `fhir_create` to create ServiceRequest resource with category=referral, reasonCode, and supportingInfo references

## FHIR Resources
- **Patient** -- Demographics for referral header
- **Coverage** -- Insurance for authorization
- **Condition** -- Referral indication and supporting diagnoses
- **Observation** -- Relevant lab results and vitals
- **MedicationRequest** -- Current medications
- **Procedure** -- Prior related procedures
- **ServiceRequest** -- Output: referral order

## FHIR Query Examples
### Create Referral Order
```
fhir_create(resourceType="ServiceRequest", resource={"resourceType":"ServiceRequest","status":"active","intent":"order","category":[{"coding":[{"system":"http://snomed.info/sct","code":"3457005","display":"Patient referral"}]}],"code":{"text":"[specialty] referral"},"subject":{"reference":"Patient/[patient-id]"},"reasonCode":[{"text":"[referral reason]"}]})
```

## Clinical Guidelines
- CMS referral and prior authorization requirements
- Specialty-specific referral criteria (e.g., nephrology at eGFR <30, cardiology for new murmur/arrhythmia)

## Interpretation Guide
- Structure the referral with: reason for referral (clinical question), relevant history, current medications, pertinent labs, prior treatments, specific request (evaluation, co-management, procedure)
- Include insurance authorization requirements if applicable
- Flag urgent referrals: conditions requiring expedited evaluation (suspected malignancy, acute decompensation, new neurologic deficit)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
