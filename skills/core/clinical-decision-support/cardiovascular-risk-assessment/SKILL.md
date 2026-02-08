---
name: cardiovascular-risk-assessment
description: |
  Calculates cardiovascular risk scores including CHA2DS2-VASc, HEART Score, Framingham, ASCVD Pooled Cohort Equations, and HAS-BLED.
  Use when user asks to "assess cardiac risk", "stroke risk in AFib", "HEART score", "CHA2DS2-VASc", "ASCVD risk",
  "statin candidacy", "bleeding risk", "chest pain workup", or needs cardiovascular risk stratification.
  Do NOT use for acute STEMI management, heart failure staging, or valvular disease assessment.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: clinical-decision-support
---

# Cardiovascular Risk Assessment

## Overview

Calculate multiple cardiovascular risk scores from FHIR Patient, Observation, Condition, and MedicationRequest resources. Supports CHA2DS2-VASc (stroke risk in atrial fibrillation), HEART Score (acute chest pain evaluation), Framingham Risk Score (10-year CHD risk), ASCVD Pooled Cohort Equations (10-year ASCVD risk), and HAS-BLED (bleeding risk on anticoagulation). Generate a RiskAssessment FHIR resource with findings and treatment-threshold recommendations.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Age, gender, race | birthDate, gender, extension (us-core-race) |
| Condition | CHF, HTN, DM, stroke, vascular disease, AF | code, clinicalStatus |
| Observation | BP, cholesterol, HDL, glucose, HbA1c, INR | code, valueQuantity, effectiveDateTime |
| MedicationRequest | Antiplatelets, anticoagulants, NSAIDs, antihypertensives | medicationCodeableConcept, status |
| RiskAssessment | Output: risk scores | method, prediction, basis |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract: age (from birthDate), gender, race (from US Core extension for ASCVD PCE).

### Step 2: Retrieve Active Conditions

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Identify conditions relevant to scoring. Key SNOMED codes:
- 49436004: Atrial fibrillation
- 42343007: Congestive heart failure
- 38341003: Hypertension
- 73211009: Diabetes mellitus
- 230690007: Cerebrovascular accident (stroke)
- 266257000: Transient ischemic attack (TIA)
- 27550009: Vascular disease (peripheral, aortic plaque, prior MI)
- 22298006: Myocardial infarction
- 400047006: Peripheral vascular disease
- 235856003: Hepatic disease
- 709044004: Chronic kidney disease
- 414545008: Ischemic heart disease

Also search historical conditions for stroke/TIA/MI history:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=230690007,266257000,22298006"
```

### Step 3: Retrieve Relevant Observations

**Blood Pressure:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=85354-9&_sort=-date&_count=3"
```
LOINC 85354-9 = Blood pressure panel. Extract systolic (component code 8480-6) and diastolic (8462-4).

**Total Cholesterol:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2093-3&_sort=-date&_count=1"
```
LOINC 2093-3 = Total cholesterol.

**HDL Cholesterol:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2085-9&_sort=-date&_count=1"
```
LOINC 2085-9 = HDL cholesterol.

**LDL Cholesterol:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=13457-7&_sort=-date&_count=1"
```
LOINC 13457-7 = LDL cholesterol (calculated).

**HbA1c:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=4548-4&_sort=-date&_count=1"
```
LOINC 4548-4 = HbA1c.

**Troponin (for HEART Score):**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=6598-7,89579-7&_sort=-date&_count=3"
```
LOINC 6598-7 = Troponin T; 89579-7 = Troponin I high sensitivity.

**INR (for HAS-BLED):**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=6301-6&_sort=-date&_count=3"
```
LOINC 6301-6 = INR.

**Creatinine:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2160-0&_sort=-date&_count=1"
```
LOINC 2160-0 = Serum creatinine. For renal function assessment.

### Step 4: Retrieve Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

Identify:
- Antihypertensives (for HEART Score, CHA2DS2-VASc HTN criterion)
- Statins (current therapy)
- Anticoagulants: warfarin, apixaban, rivaroxaban, edoxaban, dabigatran
- Antiplatelets: aspirin, clopidogrel, prasugrel, ticagrelor
- NSAIDs (for HAS-BLED)
- Antidepressants/SSRIs (bleeding risk)

### Step 5: Determine Which Scores to Calculate

Select based on clinical context:
- **Atrial fibrillation present** -> CHA2DS2-VASc + HAS-BLED
- **Chest pain presentation** -> HEART Score
- **Primary prevention / statin candidacy** -> ASCVD PCE (age 40-79) and/or Framingham
- **If unclear**, calculate all applicable scores

Refer to `references/cv-scoring-systems.md` for complete scoring criteria.

### Step 6: Calculate Scores and Generate Recommendations

Apply scoring criteria from reference tables. Map scores to treatment thresholds per `references/acc-aha-guidelines.md`.

### Step 7: Create RiskAssessment Resource

```
Tool: fhir_create
resourceType: "RiskAssessment"
resource: {
  "resourceType": "RiskAssessment",
  "status": "final",
  "subject": {"reference": "Patient/[patient-id]"},
  "occurrenceDateTime": "[current-datetime]",
  "method": {
    "coding": [{"system": "http://snomed.info/sct", "code": "225338004", "display": "Risk assessment"}]
  },
  "prediction": [
    {
      "outcome": {"text": "[Score name]: [value] - [risk level]"},
      "probabilityDecimal": [probability if applicable],
      "qualitativeRisk": {"coding": [{"system": "http://terminology.hl7.org/CodeSystem/risk-probability", "code": "[low|moderate|high]"}]}
    }
  ],
  "basis": [
    {"reference": "Condition/[relevant-condition-id]"},
    {"reference": "Observation/[relevant-observation-id]"}
  ],
  "note": [{"text": "[Treatment recommendations based on risk level]"}]
}
```

### Step 8: Format Output

```
CARDIOVASCULAR RISK ASSESSMENT
===============================
Patient: [name] | Age: [age] | Sex: [sex]
Assessment Date: [datetime]

SCORES
------
[Score Name]: [value] - [risk level]
  Components: [list scored items]
  Recommendation: [guideline-based recommendation]

[Repeat for each calculated score]

CURRENT THERAPY
---------------
Anticoagulation: [current or none]
Antiplatelets: [current or none]
Statin: [current or none]
Antihypertensives: [current or none]

RECOMMENDATIONS
---------------
[Prioritized list of guideline-based recommendations]
```

## Examples

### Example 1: AFib Patient -- Stroke and Bleeding Risk

**User says:** "What's the stroke risk for patient 54321? She has AFib."

**Actions:**
1. `fhir_read` Patient/54321 -- 74F
2. `fhir_search` Condition active -- AF, HTN, DM, prior TIA
3. `fhir_search` Observation BP -- 142/88
4. `fhir_search` Observation INR -- 2.8 (on warfarin)
5. `fhir_search` MedicationRequest active -- warfarin, metformin, lisinopril
6. Calculate: CHA2DS2-VASc = 6 (age 74 = 2, female = 1, HTN = 1, DM = 1, TIA = 2, minus 1 for overlap... recalculate: age 65-74 = 1, female = 1, HTN = 1, DM = 1, stroke/TIA = 2 = total 6)
7. Calculate: HAS-BLED = 2 (HTN = 1, age >65 = 1)
8. `fhir_create` RiskAssessment

**Result:**
```
CARDIOVASCULAR RISK ASSESSMENT
===============================
Patient: Eleanor Voss | Age: 74 | Sex: Female

SCORES
------
CHA2DS2-VASc: 6/9 - HIGH RISK
  CHF: 0, HTN: 1, Age 65-74: 1, DM: 1, Stroke/TIA: 2, Vascular: 0, Sex (female): 1
  Annual stroke risk: ~9.8%
  Recommendation: Oral anticoagulation strongly indicated

HAS-BLED: 2/9 - MODERATE BLEEDING RISK
  HTN (uncontrolled): 1, Renal/Liver: 0, Stroke: 0, Bleeding hx: 0,
  Labile INR: 0, Age >65: 1, Drugs/Alcohol: 0
  Recommendation: Anticoagulation benefits outweigh bleeding risk. Monitor INR closely.

CURRENT THERAPY
---------------
Anticoagulation: Warfarin (active)
Statin: None

RECOMMENDATIONS
---------------
1. Continue anticoagulation -- CHA2DS2-VASc 6 strongly supports OAC
2. Consider DOAC switch (apixaban/rivaroxaban) -- lower intracranial hemorrhage risk vs warfarin
3. Optimize BP control -- SBP 142 exceeds target
4. Evaluate statin candidacy -- age 74 with DM and HTN
```

### Example 2: Chest Pain Evaluation with HEART Score

**User says:** "Calculate HEART score for patient 99887, presenting with chest pain"

**Actions:**
1. `fhir_read` Patient/99887 -- 58M
2. `fhir_search` Condition active -- HTN, hyperlipidemia, prior PCI (2019)
3. `fhir_search` Observation troponin -- high-sensitivity troponin I: 45 ng/L (upper limit 26)
4. `fhir_search` Observation ECG interpretation -- no specific ST changes
5. Calculate: HEART = 7 (History moderately suspicious = 1, ECG non-diagnostic = 1, Age 55-64 = 1, Risk factors >=3 = 2, Troponin 1-3x ULN = 1... recalculate with prior PCI as significant risk = 2)

**Result:**
```
HEART Score: 7/10 - HIGH RISK
  History: 1 (moderately suspicious)
  ECG: 1 (non-specific repolarization disturbance)
  Age: 1 (45-64)
  Risk factors: 2 (HTN, hyperlipidemia, prior PCI)
  Troponin: 2 (>3x normal limit -- rechecked: 45/26 = 1.7x = 1 point)
  Corrected total: 6/10 - HIGH RISK

Recommendation: HEART >= 7: admit, cardiology consult, early invasive strategy.
HEART 4-6: observation unit, serial troponins, stress testing or CCTA.
This patient at 6: admit for observation, serial troponins q3h, consider non-invasive testing.
```

## Troubleshooting

### Cholesterol Values Not Found for ASCVD Calculation
- Check alternate LOINC codes: 2093-3 (cholesterol, total), 48620-9 (cholesterol total, calculated). For HDL: 2085-9 or 14646-4 (HDL direct). Expand search window with `&date=ge[date-1-year]` to find most recent lipid panel.
- If no lipid values exist, ASCVD PCE cannot be calculated. Recommend ordering fasting lipid panel.

### Race Data Not Available for ASCVD PCE
- ASCVD PCE uses race-specific coefficients (White, African American). Check US Core race extension on Patient resource. If race is not documented or is other than White/African American, use White coefficients and note the limitation. The equations have not been validated in other racial/ethnic groups.

### Multiple Conditions Match Same Criterion
- For CHA2DS2-VASc, each criterion scores once regardless of how many matching conditions exist. Prior stroke + prior TIA = still 2 points (single stroke/TIA category), not 4 points. Deduplicate condition matches per scoring criterion.

## Related Skills

- `clinical-summary-generator` -- for full patient context before risk assessment
- `medication-reconciliation` -- to verify current anticoagulant/statin therapy
- `drug-interaction-checker` -- when recommending new anticoagulant or statin therapy
