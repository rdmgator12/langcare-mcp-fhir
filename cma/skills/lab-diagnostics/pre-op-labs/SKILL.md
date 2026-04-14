---
name: langcare-pre-op-labs
description: >
  Evaluates preoperative laboratory readiness by checking required labs based
  on procedure type, patient age, comorbidities, and ASA classification.
  Identifies missing or expired labs and flags abnormal results that may
  delay surgery. Use when asked for pre-op labs, surgical clearance labs,
  preoperative lab checklist, or readiness for surgery assessment.
---

# Preoperative Lab Checklist

## When to Use This Skill
Use when a clinician needs to verify that all required preoperative laboratory tests have been completed, are within acceptable timeframes, and have acceptable results for surgical clearance.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender)
2. Use `fhir_search` to pull active Condition resources for comorbidity-based lab requirements
3. Use `fhir_search` to pull active MedicationRequest resources for medication-based lab requirements (e.g., anticoagulants require coag studies, digoxin requires level and electrolytes)
4. Use `fhir_search` to pull scheduled Procedure or ServiceRequest for the planned surgery to determine procedure type
5. Determine required labs based on: patient age (>45: ECG, BMP; >65: CBC, CMP, ECG), comorbidities (diabetes: glucose/A1c; cardiac: BNP/troponin; renal: BMP/CrCl; liver: LFTs/coags; anemia: CBC/iron), procedure type (blood loss risk: T&S/CBC; cardiac surgery: full panel), medications (anticoagulants: PT/INR; digoxin: level)
6. Use `fhir_search` to pull recent Observation resources (laboratory) and check if required labs are present and within acceptable timeframe (typically 30 days, 7 days for coags on anticoagulants)
7. Flag: missing labs, expired labs (older than required timeframe), abnormal results that may delay surgery

## FHIR Resources
- **Patient** -- Age for age-based requirements
- **Condition** -- Comorbidities driving lab requirements
- **MedicationRequest** -- Medications requiring monitoring
- **Observation** -- Lab results to verify completeness and values
- **ServiceRequest** -- Planned procedure for procedure-specific requirements

## FHIR Query Examples
### Pull Recent Labs (30 days)
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=laboratory&date=ge[30-days-ago]&_sort=-date&_count=200")
```

### Pull Planned Procedure
```
fhir_search(resourceType="ServiceRequest", queryParams="patient=[patient-id]&status=active&intent=order&category=387713003")
```

## Clinical Guidelines
- ASA Practice Advisory for Preanesthesia Evaluation (2012, reaffirmed)
- ACS NSQIP Surgical Risk Calculator inputs
- ACC/AHA perioperative cardiovascular evaluation guidelines

## Interpretation Guide
- Present as a checklist: required lab, status (completed/missing/expired), result, date, and pass/fail
- Flag surgical hold criteria: K+ <3.0 or >5.5, Hgb <8 (varies by procedure), platelets <50K for most surgeries, INR >1.5 (unless cardiac surgery), glucose >250, positive pregnancy test
- Include pre-op ECG requirement status for patients >45 or with cardiac history
- Calculate estimated blood loss risk by procedure type and verify type & screen/crossmatch availability

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
