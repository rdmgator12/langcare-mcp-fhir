---
name: sepsis-screening
description: |
  Screens patients for sepsis using qSOFA, SOFA, and SIRS criteria by pulling vitals and labs from FHIR Observation resources.
  Use when user asks to "screen for sepsis", "check sepsis criteria", "qSOFA score", "SOFA score", "SIRS criteria",
  "sepsis risk", "sepsis workup", or mentions suspected infection with hemodynamic instability.
  Do NOT use for chronic infection management, antibiotic stewardship reviews, or wound infection assessment.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: clinical-decision-support
---

# Sepsis Screening

## Overview

Calculate qSOFA, SOFA, and SIRS scores from FHIR Observation and Condition resources. Pull vitals (temperature, heart rate, respiratory rate, blood pressure, SpO2) and labs (WBC, lactate, creatinine, bilirubin, platelets, PaO2/FiO2) to generate a comprehensive sepsis risk assessment. Create a ClinicalImpression resource documenting the findings. Evaluate compliance with Surviving Sepsis Campaign hour-1 bundle elements.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Observation | Vitals and lab values | code, valueQuantity, effectiveDateTime, status |
| Condition | Active infections, comorbidities | code, clinicalStatus, onsetDateTime |
| Patient | Age, demographics | birthDate, gender |
| MedicationRequest | Vasopressor use, current antibiotics | medicationCodeableConcept, status, authoredOn |
| Procedure | Mechanical ventilation status | code, status, performedPeriod |
| ClinicalImpression | Output: sepsis risk assessment | status, description, finding, summary |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract birthDate (calculate age) and gender. Age modifies baseline risk interpretation.

### Step 2: Pull Vital Signs (Last 24 Hours)

Retrieve each vital sign category using LOINC codes. Use `_sort=-date&_count=5` to get recent trends.

**Temperature:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8310-5&_sort=-date&_count=5"
```
LOINC 8310-5 = Body temperature. Flag if >38.3C (100.9F) or <36.0C (96.8F).

**Heart Rate:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8867-4&_sort=-date&_count=5"
```
LOINC 8867-4 = Heart rate. Flag if >90 bpm (SIRS) or >100 bpm.

**Respiratory Rate:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=9279-1&_sort=-date&_count=5"
```
LOINC 9279-1 = Respiratory rate. Flag if >=22 (qSOFA) or >20 (SIRS).

**Systolic Blood Pressure:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8480-6&_sort=-date&_count=5"
```
LOINC 8480-6 = Systolic BP. Flag if <=100 mmHg (qSOFA).

**SpO2:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2708-6&_sort=-date&_count=5"
```
LOINC 2708-6 = Oxygen saturation. Flag if <90%.

### Step 3: Pull Laboratory Values

**WBC:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=6690-2&_sort=-date&_count=3"
```
LOINC 6690-2 = WBC count. Flag if >12,000/mm3 or <4,000/mm3 or >10% bands.

**Lactate:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2524-7&_sort=-date&_count=3"
```
LOINC 2524-7 = Lactate (arterial). Also try 32693-4 (venous lactate). Flag if >2 mmol/L.

**Creatinine:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2160-0&_sort=-date&_count=3"
```
LOINC 2160-0 = Serum creatinine. Needed for SOFA renal component.

**Total Bilirubin:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=1975-2&_sort=-date&_count=3"
```
LOINC 1975-2 = Total bilirubin. Needed for SOFA hepatic component.

**Platelets:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=777-3&_sort=-date&_count=3"
```
LOINC 777-3 = Platelet count. Needed for SOFA coagulation component.

**PaO2:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2703-7&_sort=-date&_count=3"
```
LOINC 2703-7 = PaO2. Needed for SOFA respiration component (PaO2/FiO2 ratio).

**FiO2:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=3150-0&_sort=-date&_count=3"
```
LOINC 3150-0 = FiO2. If unavailable, assume 0.21 (room air) unless patient on supplemental O2.

### Step 4: Check Active Conditions and Infection Source

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Scan for active infection-related conditions. Relevant SNOMED codes:
- 128241005 (Inflammatory disorder)
- 40733004 (Infectious disease)
- 233604007 (Pneumonia)
- 68566005 (UTI)
- 373945007 (Pyelonephritis)
- 302809008 (Cellulitis)
- 13125006 (Abscess)

### Step 5: Check Vasopressor and Ventilation Status

**Vasopressors:**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&code=3628,3616,11149,35208"
```
RxNorm: norepinephrine (3628), epinephrine (3616), vasopressin (11149), dopamine (35208).

**Mechanical Ventilation:**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&status=in-progress&code=40617009"
```
SNOMED 40617009 = Artificial respiration.

### Step 6: Calculate Scores

Calculate all three scores using the reference tables in `references/sofa-scoring.md`.

**qSOFA (Quick SOFA):** Score 0-3.
- Respiratory rate >=22: +1
- Systolic BP <=100 mmHg: +1
- Altered mental status (GCS <15): +1
- Score >=2: high risk for poor outcome, investigate for organ dysfunction.

**SIRS Criteria:** Score 0-4.
- Temp >38.3C or <36.0C: +1
- Heart rate >90: +1
- Respiratory rate >20 or PaCO2 <32 mmHg: +1
- WBC >12,000 or <4,000 or >10% bands: +1
- Score >=2 with suspected infection: meets SIRS-based sepsis definition.

**SOFA Score:** Score 0-24 (6 organ systems, 0-4 each). See `references/sofa-scoring.md` for complete table.

### Step 7: Evaluate Sepsis Bundle Compliance

Check whether the following hour-1 bundle elements have been initiated. See `references/sepsis-bundles.md` for details.

1. **Lactate measured** -- check if lactate Observation exists within last 6 hours
2. **Blood cultures obtained** -- search for blood culture orders:
   ```
   Tool: fhir_search
   resourceType: "ServiceRequest"
   queryParams: "patient=[patient-id]&code=600-7&_sort=-date&_count=1"
   ```
   LOINC 600-7 = Blood culture.
3. **Broad-spectrum antibiotics administered** -- check active antibiotic MedicationRequests
4. **Crystalloid fluid resuscitation** -- check for IV fluid orders (30 mL/kg for hypotension or lactate >=4)
5. **Vasopressors** -- if MAP <65 after fluids, check vasopressor orders (from Step 5)

### Step 8: Create ClinicalImpression Resource

```
Tool: fhir_create
resourceType: "ClinicalImpression"
resource: {
  "resourceType": "ClinicalImpression",
  "status": "completed",
  "subject": {"reference": "Patient/[patient-id]"},
  "effectiveDateTime": "[current-datetime]",
  "description": "Sepsis screening - qSOFA: [score]/3, SIRS: [score]/4, SOFA: [score]/24",
  "finding": [
    {
      "itemCodeableConcept": {
        "coding": [{"system": "http://snomed.info/sct", "code": "91302008", "display": "Sepsis"}],
        "text": "Sepsis risk: [LOW|MODERATE|HIGH]. qSOFA=[X], SIRS=[X], SOFA=[X]. [Summary of abnormal findings]."
      }
    }
  ],
  "note": [{"text": "Bundle compliance: Lactate [Y/N], Cultures [Y/N], Antibiotics [Y/N], Fluids [Y/N], Vasopressors [Y/N/NA]"}]
}
```

### Step 9: Format Output

```
SEPSIS SCREENING RESULTS
========================
Patient: [name] | Age: [age] | MRN: [mrn]
Screening Time: [datetime]

SCORES
------
qSOFA:  [X]/3  [LOW RISK / HIGH RISK]
SIRS:   [X]/4  [MET / NOT MET]
SOFA:   [X]/24 [Organ dysfunction: YES/NO]

ABNORMAL FINDINGS
-----------------
Vitals: [list abnormal vitals with values and timestamps]
Labs:   [list abnormal labs with values and timestamps]

IDENTIFIED INFECTION SOURCE
----------------------------
[Active infection conditions or "No documented infection source"]

HOUR-1 BUNDLE COMPLIANCE
-------------------------
[x] Lactate measured: [value] at [time]
[ ] Blood cultures: NOT ORDERED
[x] Antibiotics: [drug] started at [time]
[ ] IV Fluids: 30 mL/kg NOT initiated
[N/A] Vasopressors: MAP >65, not indicated

RISK ASSESSMENT
---------------
[Summary statement with recommended actions]
```

## Examples

### Example 1: High-Risk Sepsis Screening

**User says:** "Screen patient 67890 for sepsis, she's febrile and hypotensive"

**Actions:**
1. `fhir_read` Patient/67890 -- 72F female
2. `fhir_search` Observation vitals -- Temp 39.1C, HR 118, RR 26, SBP 88, SpO2 92%
3. `fhir_search` Observation labs -- WBC 18,200, Lactate 4.2, Creatinine 2.1, Platelets 98,000, Bilirubin 1.8
4. `fhir_search` Condition active -- UTI (SNOMED 68566005) documented
5. `fhir_search` MedicationRequest vasopressors -- norepinephrine active
6. Calculate: qSOFA 3/3, SIRS 4/4, SOFA 11/24
7. `fhir_create` ClinicalImpression with high-risk assessment

**Result:**
```
SEPSIS SCREENING RESULTS
========================
Patient: Dorothy Chen | Age: 72 | MRN: MRN-67890
Screening Time: 2025-01-15T14:30:00Z

SCORES
------
qSOFA:  3/3  HIGH RISK
SIRS:   4/4  MET
SOFA:   11/24 Organ dysfunction: YES (renal, hepatic, coagulation, cardiovascular, respiratory)

ABNORMAL FINDINGS
-----------------
Vitals: Temp 39.1C, HR 118, RR 26, SBP 88, SpO2 92%
Labs:   WBC 18.2K, Lactate 4.2 mmol/L, Cr 2.1, Plt 98K, Bili 1.8

IDENTIFIED INFECTION SOURCE
----------------------------
Active: Urinary tract infection (onset 2025-01-14)

HOUR-1 BUNDLE COMPLIANCE
-------------------------
[x] Lactate measured: 4.2 mmol/L at 14:15
[x] Blood cultures: Ordered at 13:45
[x] Antibiotics: Piperacillin-tazobactam started at 14:00
[ ] IV Fluids: 30 mL/kg NOT documented (REQUIRED - lactate >=4)
[x] Vasopressors: Norepinephrine active

RISK ASSESSMENT
---------------
HIGH RISK - Septic shock. SOFA 11 with vasopressor requirement and lactate >2.
ACTION NEEDED: Ensure 30 mL/kg crystalloid initiated. Repeat lactate in 2-4 hours. Re-assess MAP target 65 mmHg.
```

### Example 2: Low-Risk Screening With SIRS Criteria Met

**User says:** "Run qSOFA and SIRS on patient abc-456"

**Actions:**
1. `fhir_read` Patient/abc-456 -- 34M
2. `fhir_search` Observation vitals -- Temp 38.5C, HR 102, RR 18, SBP 125, SpO2 98%
3. `fhir_search` Observation labs -- WBC 13,400, Lactate 1.1, Creatinine 0.9, Platelets 245,000
4. `fhir_search` Condition active -- Acute pharyngitis
5. Calculate: qSOFA 0/3, SIRS 3/4, SOFA 0/24

**Result:**
```
SCORES
------
qSOFA:  0/3  LOW RISK
SIRS:   3/4  MET (temp, HR, WBC)
SOFA:   0/24 Organ dysfunction: NO

RISK ASSESSMENT
---------------
LOW RISK - SIRS criteria met in context of acute pharyngitis but no organ dysfunction.
No bundle activation required. Monitor clinically. Repeat screening if clinical status worsens.
```

## Troubleshooting

### Lactate Observation Not Found
- Try both arterial (LOINC 2524-7) and venous (LOINC 32693-4) lactate codes. Some systems use 2519-7 (lactic acid, blood). Expand search:
  ```
  Tool: fhir_search
  resourceType: "Observation"
  queryParams: "patient=[patient-id]&code=2524-7,32693-4,2519-7&_sort=-date&_count=3"
  ```
- If no lactate exists, flag as "Lactate NOT measured" in bundle compliance and recommend ordering.

### Vital Signs Returned With Mixed Units
- Temperature may be in Celsius or Fahrenheit. Check `valueQuantity.unit`. Convert F to C: (F - 32) x 5/9.
- Blood pressure may be a single Observation with components (LOINC 85354-9) rather than separate systolic/diastolic. Check `component` array for systolic (8480-6) and diastolic (8462-4) sub-values.

### Mental Status / GCS Not Available for qSOFA
- Search for GCS score: `code=9269-2` (LOINC Glasgow coma score total). If unavailable, check nursing assessments or notes. If no GCS data exists, note "GCS unavailable -- qSOFA may underestimate risk" and score the mental status component as 0 with a caveat.

## Related Skills

- `clinical-summary-generator` -- for comprehensive patient overview before sepsis screening
- `medication-reconciliation` -- to verify current antibiotic regimen and potential interactions
- `problem-list-review` -- to identify chronic conditions affecting SOFA baseline
