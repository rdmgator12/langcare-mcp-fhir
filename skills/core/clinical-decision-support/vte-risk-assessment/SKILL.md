---
name: vte-risk-assessment
description: |
  Assesses venous thromboembolism risk using Wells DVT, Wells PE, Geneva Score, and Caprini Score from FHIR resources.
  Use when user asks to "assess VTE risk", "Wells score", "DVT risk", "PE probability", "Caprini score",
  "clot risk", "thrombosis risk", "VTE prophylaxis", or needs pre-operative VTE assessment.
  Do NOT use for active anticoagulation management, HIT evaluation, or arterial thrombosis assessment.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: clinical-decision-support
---

# VTE Risk Assessment

## Overview

Calculate VTE risk scores from FHIR Patient, Condition, Observation, Procedure, and MedicationRequest resources. Supports Wells Criteria for DVT, Wells Criteria for PE, Revised Geneva Score for PE, and Caprini Score for surgical VTE prophylaxis. Generate prophylaxis and diagnostic recommendations based on risk stratification. Create a RiskAssessment FHIR resource documenting findings.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Age, gender | birthDate, gender |
| Condition | Active diagnoses, cancer, prior VTE, immobilization | code, clinicalStatus, onsetDateTime |
| Observation | Vitals (HR, SpO2), D-dimer, labs | code, valueQuantity, effectiveDateTime |
| Procedure | Recent surgeries, immobilization, central lines | code, status, performedDateTime |
| MedicationRequest | Hormonal therapy, anticoagulants, contraceptives | medicationCodeableConcept, status |
| RiskAssessment | Output: VTE risk scores | method, prediction, basis |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age and gender. Age >40 and >60 have incremental Caprini points.

### Step 2: Retrieve Active and Historical Conditions

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active,recurrence,remission"
```

Key SNOMED codes for scoring:
- 128053003: Deep vein thrombosis
- 59282003: Pulmonary embolism
- 363346000: Malignant neoplasm (active cancer)
- 128599005: Thrombophilia (Factor V Leiden, protein C/S deficiency, antithrombin III deficiency, antiphospholipid syndrome)
- 40733004: Infectious disease (active infection)
- 22298006: Myocardial infarction (recent)
- 42343007: Congestive heart failure
- 84114007: Heart failure (NYHA class III-IV)
- 128613002: Inflammatory bowel disease
- 13645005: COPD
- 44054006: Diabetes mellitus type 2
- 414916001: Obesity
- 70691001: Varicose veins
- 309585006: Immobility
- 371128001: Leg swelling (current symptom for Wells DVT)
- 57676002: Hemoptysis (for Wells PE)

### Step 3: Retrieve Recent Procedures

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&date=ge[30-days-ago]&status=completed"
```

Identify:
- Major orthopedic surgery (hip/knee replacement): SNOMED 179344006, 179406003
- Major abdominal/pelvic surgery: SNOMED 108034002
- Neurosurgery: SNOMED 392244008
- Laparoscopic surgery: SNOMED 108191006
- Central venous catheterization: SNOMED 233527006
- Arthroscopic surgery: SNOMED 71106006

### Step 4: Retrieve Vital Signs

**Heart Rate:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8867-4&_sort=-date&_count=3"
```
LOINC 8867-4 = Heart rate. Used in Revised Geneva Score (75-94 bpm = 3 pts, >=95 = 5 pts).

**SpO2:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2708-6&_sort=-date&_count=1"
```

### Step 5: Retrieve D-dimer (If Available)

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=48066-5&_sort=-date&_count=1"
```
LOINC 48066-5 = D-dimer (FEU). Also check 48065-7 (DDU units). Used in diagnostic pathway after clinical probability assessment.

### Step 6: Retrieve Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

Identify:
- Oral contraceptives / HRT (estrogen-containing): increases VTE risk
- Current anticoagulants (already on prophylaxis/treatment)
- Chemotherapy agents (increased VTE risk)
- Tamoxifen/raloxifene (SERMs)

### Step 7: Determine Which Scores to Calculate

| Clinical Scenario | Score(s) |
|-------------------|----------|
| Suspected DVT (leg swelling, pain) | Wells DVT |
| Suspected PE (dyspnea, chest pain, tachycardia) | Wells PE + Revised Geneva |
| Pre-operative VTE prophylaxis | Caprini |
| Hospitalized medical patient | Caprini or Padua (use Caprini) |
| Post-operative patient | Caprini |

Refer to `references/vte-scoring.md` for complete criteria.

### Step 8: Calculate Scores and Generate Recommendations

Apply scoring from reference tables. Map to diagnostic pathways and prophylaxis recommendations per `references/vte-management.md`.

### Step 9: Create RiskAssessment Resource

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
      "qualitativeRisk": {"coding": [{"system": "http://terminology.hl7.org/CodeSystem/risk-probability", "code": "[low|moderate|high]"}]}
    }
  ],
  "note": [{"text": "[Diagnostic/prophylaxis recommendations]"}]
}
```

### Step 10: Format Output

```
VTE RISK ASSESSMENT
====================
Patient: [name] | Age: [age] | Sex: [sex]
Assessment Date: [datetime]
Clinical Context: [suspected DVT / suspected PE / pre-operative / hospitalized medical]

SCORES
------
[Score Name]: [value] - [risk level]
  Components: [list scored items with individual points]

DIAGNOSTIC PATHWAY
------------------
[Based on score: D-dimer indicated / Imaging indicated / Alternative diagnosis likely]

PROPHYLAXIS RECOMMENDATIONS
----------------------------
[If applicable: pharmacologic / mechanical / combined]

CURRENT ANTICOAGULATION
-----------------------
[Current therapy or none]

RISK FACTORS IDENTIFIED
-----------------------
[List all VTE risk factors found in patient data]
```

## Examples

### Example 1: Suspected DVT

**User says:** "Assess DVT risk for patient 44556, has leg swelling"

**Actions:**
1. `fhir_read` Patient/44556 -- 62M
2. `fhir_search` Condition -- active cancer (colon), prior DVT (2022), no immobilization documented
3. `fhir_search` Procedure recent -- none in last 30 days
4. `fhir_search` Observation D-dimer -- 1.8 mcg/mL FEU (elevated)
5. `fhir_search` MedicationRequest -- chemotherapy active (oxaliplatin), no anticoagulant
6. Calculate: Wells DVT = 4 (active cancer = 1, prior DVT = 1, entire leg swollen = 1, calf >3cm = 1)

**Result:**
```
VTE RISK ASSESSMENT
====================
Patient: Richard Cole | Age: 62 | Sex: Male
Clinical Context: Suspected DVT

SCORES
------
Wells DVT: 4 - HIGH PROBABILITY (DVT likely)
  Active cancer (+1): Colon cancer, on chemotherapy
  Prior documented DVT (+1): DVT 2022
  Entire leg swollen (+1): Per clinical assessment
  Calf swelling >3cm (+1): Per clinical assessment
  Paralysis/paresis (0), Recently bedridden (0), Localized tenderness (0)
  Pitting edema (0), Collateral veins (0), Alternative diagnosis (-2): NOT applied

DIAGNOSTIC PATHWAY
------------------
Wells >= 2 (DVT likely): Proceed directly to compression ultrasound.
D-dimer (1.8 mcg/mL FEU): Elevated, but D-dimer not useful in "likely" category -- imaging required regardless.
If ultrasound negative but high suspicion: repeat in 5-7 days or consider venography.

RISK FACTORS IDENTIFIED
-----------------------
- Active malignancy (colon cancer)
- Prior VTE history
- Active chemotherapy
- Age >60
```

### Example 2: Pre-Operative VTE Prophylaxis

**User says:** "Caprini score for patient 78901, scheduled for hip replacement"

**Actions:**
1. `fhir_read` Patient/78901 -- 71F, BMI 34
2. `fhir_search` Condition -- obesity, varicose veins, HTN, no prior VTE, no cancer
3. `fhir_search` Procedure -- elective total hip arthroplasty scheduled
4. `fhir_search` MedicationRequest -- HRT (conjugated estrogens) active
5. Calculate: Caprini = 9 (age 61-74 = 2, major surgery >45min = 2, BMI >25 = 1, varicose veins = 1, hip arthroplasty = 5... note: some items overlap)

**Result:**
```
VTE RISK ASSESSMENT
====================
Patient: Patricia Wong | Age: 71 | Sex: Female
Clinical Context: Pre-operative (total hip arthroplasty)

SCORES
------
Caprini: 9 - HIGHEST RISK
  Age 61-74 (+2), Female (+1 if applicable in scoring), BMI >25 (+1),
  Varicose veins (+1), Major surgery >45min (+2),
  Hip arthroplasty (+5)
  Note: Score may vary by institutional Caprini version.

PROPHYLAXIS RECOMMENDATIONS
----------------------------
Caprini >= 5 (Highest Risk):
- Pharmacologic: LMWH (enoxaparin 40mg SQ daily) or rivaroxaban 10mg daily
- Start: 12 hours pre-op or 6-12 hours post-op per protocol
- Duration: Extended prophylaxis 35 days for hip arthroplasty
- Mechanical: Intermittent pneumatic compression devices intra-op and post-op
- Discontinue HRT: Recommend stopping conjugated estrogens >=4 weeks pre-op if elective
- Early mobilization protocol
```

## Troubleshooting

### Prior VTE History Not Found in Conditions
- VTE may be documented as resolved/inactive. Search without clinical-status filter:
  ```
  Tool: fhir_search
  resourceType: "Condition"
  queryParams: "patient=[patient-id]&code=128053003,59282003"
  ```
- Check for ICD-10 mapped conditions: I26.x (PE), I82.x (DVT). Some FHIR servers use ICD-10 coding system `http://hl7.org/fhir/sid/icd-10-cm`.

### D-dimer Units Vary Between Labs
- FEU (fibrinogen equivalent units): normal <0.5 mcg/mL or <500 ng/mL
- DDU (D-dimer units): normal <0.25 mcg/mL or <250 ng/mL
- FEU = approximately 2x DDU
- Check `valueQuantity.unit` to determine which unit system. LOINC 48066-5 = FEU, 48065-7 = DDU.
- Age-adjusted D-dimer cutoff: age x 10 ng/mL FEU for patients >50 years (e.g., 70-year-old: cutoff = 700 ng/mL FEU).

### Cancer Status Unclear
- Search for any malignant neoplasm conditions regardless of clinical status. Active cancer includes: currently receiving treatment, diagnosed within 6 months, or receiving palliative care. Conditions in remission for >6 months without active treatment are generally not scored as "active cancer" for Wells criteria.

## Related Skills

- `clinical-summary-generator` -- for full patient context before VTE assessment
- `medication-reconciliation` -- to verify anticoagulation therapy and identify hormonal risk factors
- `drug-interaction-checker` -- when initiating anticoagulant therapy
