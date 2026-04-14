---
name: langcare-oncology
description: >
  Compiles oncology treatment timeline from FHIR data including cancer
  staging (TNM), treatment history (chemo, radiation, surgery), lab trends
  (tumor markers, CBC), and response assessment (RECIST criteria). Use when
  asked about oncology timeline, cancer treatment history, tumor staging,
  treatment response, or cancer care summary.
---

# Oncology Treatment Timeline

## When to Use This Skill
Use when a clinician needs a comprehensive cancer care timeline including staging, treatment history, lab trends, and response assessment.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics
2. Use `fhir_search` to pull cancer Condition resources with staging information (TNM, clinical stage, pathologic stage)
3. Use `fhir_search` to pull Procedure resources for surgical interventions (biopsy, resection, radiation)
4. Use `fhir_search` to pull MedicationRequest/MedicationAdministration for chemotherapy, immunotherapy, targeted therapy, hormonal therapy
5. Use `fhir_search` to pull Observation resources for tumor markers (PSA, CEA, CA-125, CA 19-9, AFP, HCG) and lab trends (CBC for myelosuppression, LFTs, renal function)
6. Use `fhir_search` to pull DiagnosticReport for pathology and imaging reports
7. Organize into chronological treatment timeline with staging at diagnosis and response assessments (see references/tnm-recist-staging.md)
8. Present: diagnosis summary, staging, treatment phases with dates, current regimen, tumor marker trends, key lab monitoring, and response status

## FHIR Resources
- **Condition** -- Cancer diagnosis with staging: code (ICD-10/SNOMED), stage (TNM), onsetDateTime
- **Procedure** -- Surgeries, biopsies, radiation therapy
- **MedicationRequest** / **MedicationAdministration** -- Chemotherapy, immunotherapy, targeted therapy
- **Observation** -- Tumor markers, CBC, metabolic panel
- **DiagnosticReport** -- Pathology reports, imaging (CT, PET, MRI)

## FHIR Query Examples
### Pull Cancer Conditions
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&category=encounter-diagnosis&code=http://snomed.info/sct|363346000")
```

### Pull Tumor Markers (PSA example)
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=2857-1&_sort=date&_count=20")
```

### Pull Treatment History
```
fhir_search(resourceType="MedicationAdministration", queryParams="patient=[patient-id]&_sort=date&_count=200")
```

## Clinical Guidelines
- AJCC TNM Staging System (8th Edition)
- RECIST 1.1 for solid tumor response assessment
- NCCN Clinical Practice Guidelines (disease-specific)
- ASCO Quality Oncology Practice Initiative (QOPI)

## Interpretation Guide
- TNM staging: T (tumor size/extent), N (nodal involvement), M (metastasis). Map to overall stage I-IV
- RECIST 1.1 response: CR (complete response), PR (partial response >=30% decrease), SD (stable disease), PD (progressive disease >=20% increase or new lesions)
- Treatment timeline: organize by phase (neoadjuvant, surgical, adjuvant, maintenance, palliative)
- Lab monitoring during chemo: CBC nadir timing (7-14 days post-cycle), neutropenia thresholds (ANC <500 = severe), tumor lysis labs (K+, uric acid, phosphate, LDH)
- Tumor marker interpretation: trend is more important than single value; rising markers with stable imaging may indicate early progression

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
