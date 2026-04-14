---
name: langcare-diabetes-panel
description: >
  Reviews diabetes-related laboratory results against ADA Standards of Care
  including HbA1c, fasting glucose, lipid panel, renal function, and urine
  albumin. Tracks glycemic control trends and flags overdue monitoring.
  Use when asked to review diabetes labs, A1c trends, diabetes panel,
  glycemic control, or diabetic monitoring.
---

# Diabetes Panel Review

## When to Use This Skill
Use when a clinician needs a focused diabetes laboratory review with ADA guideline-based interpretation, trend analysis, and gap identification.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age for A1c target individualization)
2. Use `fhir_search` to pull HbA1c results (LOINC 4548-4) over the past 2 years for trending
3. Use `fhir_search` to pull fasting glucose, random glucose, lipid panel, renal function (creatinine, eGFR), urine albumin/creatinine ratio
4. Use `fhir_search` to pull active Condition resources for diabetes type, complications (retinopathy, nephropathy, neuropathy)
5. Use `fhir_search` to pull active MedicationRequest resources for diabetes medications
6. Evaluate against ADA Standards of Care (see references/ada-standards.md): A1c target, lipid goals, renal monitoring, eye exam status
7. Calculate time since last A1c, last lipid panel, last urine albumin -- flag if overdue
8. Present structured diabetes dashboard with trends, goal attainment, and recommended actions

## FHIR Resources
- **Observation** -- HbA1c (4548-4), glucose (2345-7), lipids (57698-3), creatinine (2160-0), eGFR (33914-3), urine albumin/Cr (14959-1)
- **Condition** -- Diabetes type, complications
- **MedicationRequest** -- Diabetes medications (metformin, insulin, GLP-1 RA, SGLT2i, sulfonylureas)
- **Procedure** -- Dilated eye exam (SNOMED 252891009)

## FHIR Query Examples
### Pull A1c Trend (2 years)
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=4548-4&date=ge[2-years-ago]&_sort=date")
```

### Pull Urine Albumin
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=14959-1&_sort=-date&_count=3")
```

## Clinical Guidelines
- ADA Standards of Medical Care in Diabetes (updated annually)
- KDIGO guidelines for diabetic kidney disease
- ACC/AHA lipid management in diabetes

## Interpretation Guide
- A1c targets: <7.0% for most adults, <8.0% for elderly/high hypoglycemia risk, <6.5% for newly diagnosed without hypoglycemia risk
- A1c monitoring frequency: every 3 months if above target, every 6 months if stable at target
- Lipid goals with diabetes: LDL <100 (no ASCVD), <70 (with ASCVD or high risk)
- Urine albumin staging: <30 mg/g normal, 30-300 moderately increased (microalbuminuria), >300 severely increased (macroalbuminuria)
- Annual screening requirements: urine albumin, lipid panel, comprehensive metabolic panel, dilated eye exam

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
