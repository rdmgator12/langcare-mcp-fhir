---
name: langcare-vte-risk
description: >
  Assesses venous thromboembolism risk using Wells Criteria (DVT and PE)
  and Caprini Score for surgical patients. Recommends prophylaxis based on
  risk stratification. Use when asked about VTE risk, DVT risk, PE risk,
  Wells score, Caprini score, thromboprophylaxis, or blood clot prevention.
---

# VTE Risk Assessment

## When to Use This Skill
Use when a clinician needs to stratify VTE risk for hospitalized medical or surgical patients and determine appropriate thromboprophylaxis.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender)
2. Use `fhir_search` to pull active Condition resources for risk factors: cancer, prior VTE, recent surgery, immobilization, thrombophilia
3. Use `fhir_search` to pull recent Observation resources: vitals (HR, SpO2), D-dimer, hemoglobin
4. Use `fhir_search` to pull recent Procedure and Encounter resources for surgical/immobilization context
5. Use `fhir_search` to pull active MedicationRequest for current anticoagulation status
6. Select appropriate tool: Wells DVT or Wells PE for diagnostic workup; Caprini for surgical prophylaxis; Padua for medical prophylaxis
7. Calculate scores per references/wells-caprini-scoring.md
8. Present risk stratification with prophylaxis recommendations

## FHIR Resources
- **Patient** -- Age, gender
- **Condition** -- Cancer, prior VTE, thrombophilia, HF, stroke, acute infection
- **Observation** -- Heart rate, SpO2, D-dimer, hemoglobin
- **Procedure** -- Recent surgery type and date
- **Encounter** -- Admission context, immobilization
- **MedicationRequest** -- Current anticoagulation, hormonal therapy, chemotherapy

## FHIR Query Examples
### Pull Risk Factor Conditions
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&clinical-status=active")
```

### Pull D-dimer
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=48066-5&_sort=-date&_count=1")
```

## Clinical Guidelines
- ACCP (CHEST) Guidelines for VTE Prevention in Hospitalized Patients
- ASH 2019 VTE Guidelines
- Caprini Risk Assessment Model for surgical patients
- Padua Prediction Score for medical patients

## Interpretation Guide
- Wells DVT: <=0 low probability (D-dimer to rule out), 1-2 moderate, >=3 high (ultrasound)
- Wells PE: <=4 PE unlikely (D-dimer), >4 PE likely (CTPA)
- Caprini: 0-1 very low (early ambulation), 2 low (consider prophylaxis), 3-4 moderate (prophylaxis recommended), >=5 high (prophylaxis + mechanical)
- Padua: <4 low risk (no pharmacologic prophylaxis), >=4 high risk (anticoagulant prophylaxis)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
