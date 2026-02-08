---
name: preoperative-lab-checklist
description: |
  Determines required preoperative labs based on planned procedure, patient conditions, and ASA classification, then checks existing results and generates orders for missing labs.
  Use when user asks to "check pre-op labs", "preoperative workup", "surgery lab clearance", "pre-surgical labs",
  "OR readiness", "clearance labs", or needs to verify lab requirements before a scheduled procedure.
  Do NOT use for postoperative lab monitoring, routine lab review, or intraoperative lab management.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: lab-diagnostics
---

# Preoperative Lab Checklist

## Overview

Evaluate preoperative laboratory readiness for a scheduled surgical procedure. Determine required labs based on procedure type (low/medium/high risk), patient comorbidities, ASA physical status classification, and current medications. Query existing Observation resources to identify which required labs are already available within acceptable timeframes. Generate ServiceRequest resources for missing labs. Assess anticoagulation status and bridging needs.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Procedure | Planned surgery details | code, status, performedDateTime |
| Condition | Patient comorbidities for risk stratification | code, clinicalStatus, category |
| Observation | Existing lab results | code, valueQuantity, effectiveDateTime, status |
| MedicationStatement | Current medications (anticoagulants, etc.) | medicationCodeableConcept, status, dosage |
| Patient | Demographics, age | birthDate, gender |
| ServiceRequest | Lab orders to create | code, intent, priority, subject |

## Instructions

### Step 1: Retrieve Patient Demographics and Planned Procedure

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age and gender. Age > 65 triggers additional requirements. Pregnancy-capable patients may need HCG.

Retrieve the planned procedure:

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&status=preparation&_sort=-date&_count=5"
```

If no Procedure with status `preparation`, check ServiceRequest for surgical orders:

```
Tool: fhir_search
resourceType: "ServiceRequest"
queryParams: "patient=[patient-id]&category=http://snomed.info/sct|387713003&status=active"
```

SNOMED 387713003 = Surgical procedure. Extract procedure type and planned date.

### Step 2: Classify Surgical Risk

Based on the procedure code, classify into bleeding risk categories:

**Low Risk (minimal blood loss expected):**
- Cataract surgery, minor dermatologic procedures, endoscopy without biopsy, breast biopsy, hernia repair (laparoscopic)

**Medium Risk (moderate blood loss possible):**
- Laparoscopic cholecystectomy, arthroscopy, hysterectomy (laparoscopic), spine surgery (1-2 levels), ENT procedures

**High Risk (significant blood loss expected):**
- Cardiac surgery (CABG, valve), major vascular, major orthopedic (joint replacement, spine fusion > 2 levels), organ transplant, intracranial surgery, hepatic resection

### Step 3: Retrieve Active Conditions and Classify ASA Status

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Use conditions to estimate ASA Physical Status Classification:
- **ASA I**: Healthy, no comorbidities
- **ASA II**: Mild systemic disease (controlled HTN, controlled DM, BMI 30-40, current smoker, social alcohol)
- **ASA III**: Severe systemic disease (poorly controlled DM/HTN, COPD, BMI > 40, active hepatitis, ESRD on dialysis, CHF, stable angina, history of MI/CVA > 3 months)
- **ASA IV**: Severe systemic disease that is a constant threat to life (recent MI/CVA < 3 months, severe CHF EF < 25%, sepsis, DIC, ARDS)

### Step 4: Determine Required Labs

Apply requirements matrix based on surgery risk + ASA class + specific conditions:

**All patients (medium/high risk surgery or ASA >= II):**
- CBC (LOINC 58410-2) -- acceptable within 30 days
- BMP (LOINC 51990-0) -- acceptable within 30 days

**Age > 65 OR cardiac/pulmonary disease:**
- ECG (not a lab, but flag as needed)
- BMP including glucose and renal function

**Diabetes (SNOMED 44054006, 73211009):**
- HbA1c (LOINC 4548-4) -- acceptable within 90 days
- BMP (glucose, potassium, creatinine) -- within 30 days

**Renal disease (SNOMED 709044004) or Cr > 1.5:**
- CMP (LOINC 24323-8) -- within 14 days
- Magnesium (LOINC 19123-9) -- within 14 days
- Phosphorus (LOINC 2777-1) -- within 14 days

**Liver disease (SNOMED 235856003) or hepatic surgery:**
- Liver panel (LOINC 24325-3) -- within 14 days
- Coagulation: PT/INR (LOINC 5902-2, 6301-6), aPTT (LOINC 3173-2) -- within 14 days
- Albumin (LOINC 1751-7) -- within 14 days

**High-risk surgery (expected blood loss > 500mL):**
- Type and Screen (LOINC 34532-2) -- within 72 hours
- CBC (within 7 days)
- Coagulation panel -- within 14 days

**Anticoagulant use:**
- PT/INR for warfarin -- within 24-48 hours pre-op
- Anti-Xa for LMWH/DOACs -- within 24 hours pre-op if applicable
- CBC (platelet count) -- within 7 days

**Pregnancy-capable patients (female, age 12-55, unless documented otherwise):**
- Urine or serum HCG (LOINC 2106-3 or 19080-1) -- within 7 days

**Cardiac history or major vascular surgery:**
- Troponin baseline (LOINC 10839-9) -- within 24 hours pre-op
- BNP or NT-proBNP (LOINC 42637-9 or 33762-6) -- within 30 days

See references/preop-requirements.md for complete requirements matrix.

### Step 5: Check Existing Lab Results

For each required lab, search for existing results within the acceptable timeframe:

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=[loinc-code]&date=ge[acceptable-date]&status=final&_sort=-date&_count=1"
```

Classify each required lab as:
- **Available**: Result exists, within acceptable timeframe, value within surgical safety limits
- **Expired**: Result exists but older than acceptable timeframe
- **Abnormal**: Result exists, within timeframe, but value may require correction before surgery (e.g., INR > 1.5, Hgb < 8, K+ > 5.5 or < 3.0, glucose > 250, platelets < 50)
- **Missing**: No result found

### Step 6: Assess Anticoagulation Status

```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```

Check for anticoagulants and antiplatelets:
- **Warfarin**: Hold 5 days pre-op, target INR < 1.5. Bridge with LMWH if high thromboembolic risk (mechanical valve, recent VTE < 3 months, CHADS2-VASc >= 7)
- **Apixaban/Rivaroxaban**: Hold 48-72 hours (high-risk surgery) or 24-48 hours (low-risk). No bridging needed.
- **Dabigatran**: Hold 48-72 hours (CrCl > 50) or 96+ hours (CrCl 30-50). No bridging.
- **Aspirin**: Continue for cardiac patients. Hold 7 days for non-cardiac if high bleeding risk.
- **Clopidogrel**: Hold 5-7 days. Hold 5 days for ticagrelor. Hold 7 days for prasugrel.
- **Enoxaparin (therapeutic)**: Hold 24 hours. Prophylactic dose: hold 12 hours.

See references/anticoagulation-bridging.md for complete protocols.

### Step 7: Generate Orders for Missing Labs

For each missing or expired required lab, create a ServiceRequest:

```
Tool: fhir_create
resourceType: "ServiceRequest"
resource: {
  "status": "active",
  "intent": "order",
  "priority": "[routine|urgent]",
  "code": {
    "coding": [{
      "system": "http://loinc.org",
      "code": "[loinc-code]",
      "display": "[test-name]"
    }]
  },
  "subject": { "reference": "Patient/[patient-id]" },
  "reasonCode": [{
    "coding": [{
      "system": "http://snomed.info/sct",
      "code": "387713003",
      "display": "Surgical procedure"
    }],
    "text": "Preoperative laboratory evaluation for [procedure name]"
  }],
  "note": [{ "text": "Pre-op lab -- surgery scheduled [date]" }]
}
```

Set priority to `urgent` if surgery is within 48 hours.

### Step 8: Format Output

```
PREOPERATIVE LAB CHECKLIST -- [Patient Name] ([Age][Sex])
=========================================================
Planned Procedure: [procedure name]
Scheduled Date:    [date]
Surgery Risk:      [Low/Medium/High]
ASA Class:         [I-IV]
Relevant Conditions: [list]
Anticoagulation:   [medication or "None"]

REQUIRED LABS STATUS
====================
Test                  Status       Result          Date         Acceptable By
----                  ------       ------          ----         -------------
CBC                   Available    WBC 7.2, Hgb 13 2025-01-10  2025-01-25 (30d)
BMP                   EXPIRED      Na 140, K 4.1    2024-11-15  2025-01-25 (30d)
PT/INR                MISSING      --               --          [pre-op date]
Type & Screen         MISSING      --               --          [72h pre-op]
HbA1c                 Available    7.2%             2024-12-01  2025-02-01 (90d)

ANTICOAGULATION PLAN
====================
Warfarin 5mg daily: HOLD starting [5 days pre-op]
  Target INR < 1.5 on day of surgery
  Bridge: Enoxaparin 1mg/kg BID starting [3 days pre-op], last dose [24h pre-op]
  Rationale: Mechanical aortic valve -- high thromboembolic risk

ORDERS GENERATED
================
1. BMP (LOINC 51990-0) -- routine
2. PT/INR (LOINC 5902-2) -- urgent (draw within 48h of surgery)
3. Type & Screen (LOINC 34532-2) -- urgent (draw within 72h of surgery)

ABNORMAL VALUES REQUIRING ATTENTION
====================================
[None / or list values that may delay surgery]
```

## Examples

### Example 1: Straightforward Pre-Op Clearance

**User says:** "Check pre-op labs for patient 8812, scheduled for laparoscopic cholecystectomy"

**Actions:**
1. `fhir_read` Patient/8812 -- returns 52-year-old female
2. `fhir_search` Condition?patient=8812&clinical-status=active -- returns hypertension (controlled), GERD
3. `fhir_search` MedicationStatement?patient=8812&status=active -- returns lisinopril, omeprazole (no anticoagulants)
4. `fhir_search` Observation?patient=8812&code=58410-2&date=ge[30-days-ago]&status=final&_sort=-date&_count=1 -- CBC from 2 weeks ago, normal
5. `fhir_search` Observation?patient=8812&code=51990-0&date=ge[30-days-ago]&status=final&_sort=-date&_count=1 -- BMP from 2 weeks ago, normal
6. `fhir_search` Observation?patient=8812&code=2106-3&date=ge[7-days-ago]&status=final&_sort=-date&_count=1 -- no HCG result

**Result:** Medium-risk surgery, ASA II. CBC and BMP available and current. Missing: urine HCG (pregnancy-capable patient). Order generated for HCG. No anticoagulation concerns. Patient is cleared pending HCG result.

### Example 2: Complex Patient With Multiple Requirements

**User says:** "Pre-op workup for patient 4456, total knee replacement next week"

**Actions:**
1. `fhir_read` Patient/4456 -- returns 72-year-old male
2. `fhir_search` Condition?patient=4456&clinical-status=active -- returns CKD stage 3, Type 2 DM, atrial fibrillation, osteoarthritis
3. `fhir_search` MedicationStatement?patient=4456&status=active -- returns warfarin, metformin, amlodipine
4. Multiple Observation searches for: CBC, CMP, HbA1c, PT/INR, Type & Screen, Mg, Phos, BNP
5. `fhir_search` Observation?patient=4456&code=6301-6&_sort=-date&_count=1 -- INR 2.8 from 3 days ago

**Result:** High-risk surgery, ASA III. Required: CBC, CMP (renal disease), HbA1c, PT/INR (on warfarin), Type & Screen, Mg, Phos, BNP (age > 65 + cardiac). Anticoagulation plan: hold warfarin 5 days pre-op, bridge with enoxaparin (afib with CHA2DS2-VASc likely >= 4), recheck INR day before surgery, target < 1.5. Hold metformin 48h pre-op (renal, contrast risk). 4 labs missing, orders generated.

## Troubleshooting

### No Procedure Resource Found for Planned Surgery

- Surgery may not yet be entered as a FHIR Procedure resource. Ask the user for the procedure name and planned date, then determine requirements manually from the surgery risk classification.
- Check ServiceRequest with `intent=plan` or `intent=order` and `category` matching surgical SNOMED codes.

### Lab Results Are Panel-Level Without Individual Analytes

- A BMP panel Observation (LOINC 51990-0) may contain individual results in `component` array rather than as separate Observation resources. Check `component[].code` and `component[].valueQuantity`.
- If the panel Observation exists but individual values like potassium are needed, extract from `component` where `code.coding.code` = "2823-3" (potassium).

### Patient Is on a Novel Anticoagulant Not in the Protocol

- Default to holding the medication for 48-72 hours pre-op for high-bleeding-risk procedures and 24-48 hours for low-bleeding-risk.
- Flag for anesthesia/hematology consultation: "Anticoagulation management requires specialist input for [drug name]."

## Related Skills

- `lab-result-interpreter` -- for detailed interpretation of pre-op lab results
- `critical-value-alert-generator` -- if pre-op labs reveal critical values
- `renal-function-dashboard` -- for detailed renal assessment when CKD is present
