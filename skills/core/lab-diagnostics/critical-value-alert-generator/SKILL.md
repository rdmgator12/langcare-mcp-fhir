---
name: critical-value-alert-generator
description: |
  Scans recent laboratory Observations against CAP/CLIA critical value thresholds and generates structured alerts with clinical actions.
  Use when user asks to "check for critical labs", "critical value alert", "flag dangerous labs", "any critical results",
  "stat labs alert", "panic values", or needs automated screening of lab results for values requiring immediate clinical action.
  Do NOT use for routine lab interpretation, trending, or non-laboratory observations like vital signs.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: lab-diagnostics
---

# Critical Value Alert Generator

## Overview

Scan recent laboratory Observation resources for values exceeding critical thresholds defined by CAP (College of American Pathologists) and CLIA guidelines. Generate structured alerts containing: the critical value, threshold exceeded, clinical significance, and recommended immediate action. Support read-back verification workflow for clinical notification compliance. Prioritize alerts by clinical urgency.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Observation | Lab results to evaluate | code, valueQuantity, interpretation, effectiveDateTime, status |
| Patient | Demographics for context | birthDate, gender, name |
| MedicationStatement | Drug context for critical values | medicationCodeableConcept, status |
| Condition | Clinical context | code, clinicalStatus |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract name (for alert addressee), age (pediatric thresholds differ), gender.

### Step 2: Pull Recent Laboratory Observations

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&date=ge[24-hours-ago]&_sort=-date&_count=200"
```

Default scan window: 24 hours. Expand to 48 or 72 hours if user requests or if investigating missed alerts.

### Step 3: Screen Against Critical Value Thresholds

For each Observation, compare `valueQuantity.value` against critical thresholds.

**Priority 1 -- Immediately Life-Threatening:**

| Analyte | LOINC | Critical Low | Critical High | Unit |
|---------|-------|-------------|---------------|------|
| Potassium | 2823-3 | < 2.5 | > 6.5 | mEq/L |
| Sodium | 2951-2 | < 120 | > 160 | mEq/L |
| Glucose | 2345-7 | < 40 | > 500 | mg/dL |
| Calcium (total) | 17861-6 | < 6.0 | > 13.0 | mg/dL |
| Ionized Calcium | 1994-3 | < 0.78 | > 1.58 | mmol/L |
| pH (arterial) | 2744-1 | < 7.20 | > 7.60 | -- |
| pCO2 | 2019-8 | < 20 | > 70 | mmHg |
| pO2 | 2703-7 | < 40 | -- | mmHg |
| Lactate | 2524-7 | -- | > 4.0 | mmol/L |
| Troponin I | 10839-9 | -- | > 0.04 | ng/mL |
| Troponin T | 6598-7 | -- | > 0.10 | ng/mL |

**Priority 2 -- Urgent (within 1 hour):**

| Analyte | LOINC | Critical Low | Critical High | Unit |
|---------|-------|-------------|---------------|------|
| Hemoglobin | 718-7 | < 7.0 | > 20.0 | g/dL |
| WBC | 26464-8 | < 1.5 | > 30.0 | x10^3/uL |
| Platelets | 777-3 | < 20 | > 1000 | x10^3/uL |
| INR | 6301-6 | -- | > 5.0 | -- |
| aPTT | 3173-2 | -- | > 100 | seconds |
| Fibrinogen | 3255-7 | < 100 | -- | mg/dL |
| Magnesium | 19123-9 | < 1.0 | > 4.0 | mg/dL |
| Phosphorus | 2777-1 | < 1.0 | > 8.0 | mg/dL |
| Ammonia | 1925-7 | -- | > 100 | umol/L |
| Bilirubin (neonatal) | 58941-6 | -- | > 15.0 | mg/dL |

**Priority 3 -- Urgent Drug Levels:**

| Drug | LOINC | Critical Low | Critical High | Unit |
|------|-------|-------------|---------------|------|
| Digoxin | 10535-3 | -- | > 2.0 | ng/mL |
| Lithium | 14334-7 | -- | > 1.5 | mEq/L |
| Phenytoin | 3968-5 | -- | > 30.0 | ug/mL |
| Vancomycin (trough) | 4090-7 | -- | > 30.0 | ug/mL |
| Theophylline | 4049-3 | -- | > 20.0 | ug/mL |
| Carbamazepine | 3432-2 | -- | > 15.0 | ug/mL |
| Valproic Acid | 4086-5 | -- | > 120.0 | ug/mL |
| Methotrexate (24h) | 14836-1 | -- | > 5.0 | umol/L |

See references/critical-value-table.md for complete thresholds with pediatric adjustments.

Also check the `interpretation` field: codes `HH` (critical high) or `LL` (critical low) from the FHIR server itself may flag values the threshold table does not cover.

### Step 4: Retrieve Clinical Context for Each Critical Value

For each critical value found:

```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Correlate:
- Critical potassium + ACE inhibitor/ARB/spironolactone = drug-induced hyperkalemia
- Critical potassium low + loop diuretic = drug-induced hypokalemia
- Critical INR + warfarin = supratherapeutic anticoagulation
- Critical glucose low + insulin/sulfonylurea = drug-induced hypoglycemia
- Critical lithium + renal impairment = accumulation risk
- Critical troponin + known CAD = acute coronary syndrome context

### Step 5: Check for Prior Value (Delta)

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=[loinc-code]&_sort=-date&_count=2"
```

Compare current critical value to the most recent prior result. Note whether the value is newly critical or has been trending toward critical.

### Step 6: Generate Structured Alert

For each critical value, produce an alert block:

```
=====================================================
CRITICAL VALUE ALERT -- PRIORITY [1/2/3]
=====================================================
Patient:    [Name] (MRN: [id])
Test:       [Analyte name]
Value:      [result] [unit]
Threshold:  [critical low/high threshold]
Collected:  [effectiveDateTime]
Prior:      [prior value] on [prior date] (delta: [change])

CLINICAL SIGNIFICANCE:
[1-2 sentence explanation of the clinical risk]

RECOMMENDED IMMEDIATE ACTIONS:
1. [First action -- most urgent]
2. [Second action]
3. [Third action]

MEDICATION CONTEXT:
- [Relevant active medication and its relationship to the critical value]

READ-BACK VERIFICATION:
"The critical value for [Patient Name] is [analyte] of [value] [unit],
collected at [time]. The critical threshold is [threshold]. Please
read back the patient name, test name, and value to confirm receipt."
=====================================================
```

### Step 7: Prioritize and Order Alerts

Sort alerts by:
1. Priority level (1 before 2 before 3)
2. Within same priority, by clinical urgency (e.g., potassium 2.0 before potassium 2.4)
3. Timestamp (most recent first within same analyte)

If no critical values found, report: "No critical values detected in laboratory results from [date range]. [N] total observations reviewed."

## Examples

### Example 1: Multiple Critical Values Detected

**User says:** "Check for critical labs on patient 33201"

**Actions:**
1. `fhir_read` Patient/33201 -- returns 58-year-old male, John Martinez
2. `fhir_search` Observation?patient=33201&category=laboratory&date=ge[24-hours-ago]&_sort=-date&_count=200 -- returns 12 results; potassium 6.8 mEq/L (HH), glucose 38 mg/dL (LL), hemoglobin 6.2 g/dL (L)
3. `fhir_search` MedicationStatement?patient=33201&status=active -- returns spironolactone, insulin glargine, aspirin
4. `fhir_search` Condition?patient=33201&clinical-status=active -- returns CKD stage 4, Type 2 DM, CHF
5. `fhir_search` Observation?patient=33201&code=2823-3&_sort=-date&_count=2 -- prior potassium 5.8 two days ago
6. `fhir_search` Observation?patient=33201&code=2345-7&_sort=-date&_count=2 -- prior glucose 95 yesterday
7. `fhir_search` Observation?patient=33201&code=718-7&_sort=-date&_count=2 -- prior hemoglobin 7.1 one week ago

**Result:** Three alerts generated:
- PRIORITY 1: Potassium 6.8 (threshold > 6.5), trending up from 5.8. Spironolactone + CKD4 causative. Action: ECG stat, calcium gluconate, hold spironolactone, kayexalate.
- PRIORITY 1: Glucose 38 (threshold < 40), prior was 95. Insulin glargine causative. Action: D50 IV push, recheck in 15 min, hold insulin pending evaluation.
- PRIORITY 2: Hemoglobin 6.2 (threshold < 7.0), dropping from 7.1. Action: Type and screen, transfuse 1 unit pRBC, evaluate for active bleeding.

### Example 2: No Critical Values Found

**User says:** "Any critical labs for patient abc-100 in the last 48 hours?"

**Actions:**
1. `fhir_read` Patient/abc-100 -- returns 34-year-old female
2. `fhir_search` Observation?patient=abc-100&category=laboratory&date=ge[48-hours-ago]&_sort=-date&_count=200 -- returns 8 results; all within normal limits or mildly abnormal (e.g., LDL 142 mg/dL flagged H)

**Result:** "No critical values detected in laboratory results from the past 48 hours. 8 total observations reviewed. Note: 1 non-critical abnormality -- LDL 142 mg/dL (above reference 0-130). No immediate action required."

## Troubleshooting

### FHIR Server Does Not Use interpretation Coding

- Some servers omit the `interpretation` field entirely. Rely on threshold comparison against `valueQuantity.value` using the critical value table.
- If `valueQuantity` is also absent, check `valueString` -- some results are reported as text (e.g., "> 500" for troponin). Parse the numeric value from the string.

### Observation Uses Different Units Than Expected

- Potassium may be reported as mmol/L instead of mEq/L (1:1 conversion for monovalent ions, so thresholds are the same).
- Glucose may be mmol/L instead of mg/dL. Convert: mmol/L * 18 = mg/dL. Adjust thresholds: < 2.2 mmol/L (critical low), > 27.8 mmol/L (critical high).
- Hemoglobin may be g/L instead of g/dL. Convert: g/L / 10 = g/dL.
- Always check `valueQuantity.unit` or `valueQuantity.code` before comparing to thresholds.

### Multiple Specimens for Same Analyte in Scan Window

- A repeated critical value may indicate the first was already addressed and this is a follow-up. Compare timestamps.
- If the second value is improving (moving toward normal), note "trending toward correction" in the alert.
- If the second value is worsening, escalate priority and note "worsening despite intervention."

## Related Skills

- `lab-result-interpreter` -- for comprehensive lab interpretation beyond critical values
- `renal-function-dashboard` -- for critical electrolytes in the context of renal function staging
- `diabetes-panel-review` -- for critical glucose values in the context of diabetes management
