---
name: langcare-clinical-summary
description: >
  Generates a comprehensive clinical summary (CCD-style) from FHIR resources
  including problems, medications, allergies, labs, vitals, procedures, and
  immunizations. Use when asked for chart review, clinical summary, patient
  overview, CCD, or comprehensive patient snapshot before a visit or consult.
---

# Clinical Summary Generator

## When to Use This Skill
Use when a clinician needs a consolidated clinical overview of a patient spanning active problems, medications, allergies, recent labs, vitals, procedures, immunizations, and care plans -- analogous to a Continuity of Care Document (CCD).

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (name, DOB, age, gender, MRN)
2. Use `fhir_search` to pull active Condition resources for the problem list
3. Use `fhir_search` to pull active MedicationRequest resources for the medication list
4. Use `fhir_search` to pull active AllergyIntolerance resources
5. Use `fhir_search` to pull recent Observation resources (category=laboratory, last 90 days) and vital signs (last encounter)
6. Use `fhir_search` to pull recent Procedure resources (last 12 months)
7. Use `fhir_search` to pull Immunization resources
8. Present as structured CCD sections with abnormals flagged

## FHIR Resources
- **Patient** -- Demographics, identifiers
- **Condition** -- Active problem list with ICD-10/SNOMED codes, onset dates
- **MedicationRequest** -- Active medications with dose, frequency, route
- **AllergyIntolerance** -- Allergies with reaction type and severity
- **Observation** -- Lab results (category=laboratory) and vital signs (category=vital-signs)
- **Procedure** -- Recent procedures with dates and outcomes
- **Immunization** -- Vaccination history with dates and vaccine codes
- **CarePlan** -- Active care plans with goals and activities

## FHIR Query Examples
### Pull Active Problems
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&clinical-status=active")
```

### Pull Active Medications
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

### Pull Recent Labs (90 days)
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=laboratory&date=ge[90-days-ago]&_sort=-date&_count=200")
```

### Pull Latest Vitals
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=vital-signs&_sort=-date&_count=10")
```

## Clinical Guidelines
- ONC USCDI v3 data classes define the minimum data elements for a clinical summary
- HL7 CDA CCD template defines section structure: Problems, Medications, Allergies, Results, Vitals, Procedures, Immunizations, Plan of Care
- Joint Commission requires medication reconciliation documentation at every transition of care

## Interpretation Guide
- Group lab results by panel (CBC, BMP, CMP, lipids) and flag abnormals with reference ranges
- Sort problems by clinical significance: acute conditions first, then chronic active conditions
- For medications, distinguish between chronic maintenance therapy and recent/acute prescriptions by authoredOn date
- Flag overdue preventive care items (e.g., no colonoscopy if age 45+, no mammogram if female 50+)
- Present vitals with trend arrows if multiple readings available

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
