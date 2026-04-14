---
name: langcare-cardiovascular-risk
description: >
  Calculates cardiovascular risk scores including CHA2DS2-VASc, HEART Score,
  ASCVD Pooled Cohort Equations, and HAS-BLED from FHIR data. Generates
  treatment-threshold recommendations per ACC/AHA guidelines. Use when asked
  to assess cardiac risk, stroke risk in AFib, HEART score, CHA2DS2-VASc,
  ASCVD risk, or statin candidacy.
---

# Cardiovascular Risk Assessment

## When to Use This Skill
Use when a clinician needs cardiovascular risk stratification for treatment decisions including anticoagulation for AFib, statin therapy initiation, or acute chest pain evaluation.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender, race for ASCVD PCE)
2. Use `fhir_search` to pull active Condition resources for scoring: AFib, CHF, HTN, DM, stroke/TIA, MI, PVD
3. Use `fhir_search` to pull Observation resources: BP, cholesterol, HDL, LDL, HbA1c, troponin, INR, creatinine
4. Use `fhir_search` to pull active MedicationRequest: anticoagulants, antiplatelets, statins, antihypertensives, NSAIDs
5. Select scores based on context: AFib -> CHA2DS2-VASc + HAS-BLED; chest pain -> HEART Score; primary prevention -> ASCVD PCE
6. Calculate scores using reference tables (see references/ascvd-heart-scoring.md)
7. Use `fhir_create` to create RiskAssessment resource with findings and predictions
8. Present scores with treatment-threshold recommendations per ACC/AHA guidelines

## FHIR Resources
- **Patient** -- Age, gender, race (US Core extension for ASCVD PCE)
- **Condition** -- CHF, HTN, DM, stroke/TIA, MI, AFib, PVD, CKD
- **Observation** -- BP (85354-9), cholesterol (2093-3), HDL (2085-9), LDL (18262-6), troponin (6598-7, 89579-7), INR (6301-6)
- **MedicationRequest** -- Anticoagulants, antiplatelets, statins, antihypertensives
- **RiskAssessment** -- Output: risk scores and predictions

## FHIR Query Examples
### Pull Active Conditions
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&clinical-status=active")
```

### Pull Blood Pressure
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=85354-9&_sort=-date&_count=3")
```

### Pull Lipids
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=2093-3,2085-9,18262-6&_sort=-date&_count=3")
```

## Clinical Guidelines
- ACC/AHA 2019 Primary Prevention of Cardiovascular Disease
- ACC/AHA/HRS 2019 AFib Management (CHA2DS2-VASc thresholds)
- AHA 2021 Chest Pain Guideline (HEART Score)
- ASCVD PCE risk calculator (40-79 years)

## Interpretation Guide
- CHA2DS2-VASc: 0 (male) or 1 (female) = no anticoagulation; 1 (male) = consider; >=2 = anticoagulate
- HAS-BLED >=3 = high bleeding risk, but does NOT mean withhold anticoagulation; optimize modifiable risk factors
- HEART Score: 0-3 = low risk (discharge), 4-6 = moderate (observation), 7-10 = high (admit, invasive strategy)
- ASCVD 10-year risk: <5% low, 5-7.5% borderline, 7.5-20% intermediate (statin benefit), >=20% high (statin strongly recommended)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
