---
name: diabetes-panel-review
description: |
  Pulls and interprets diabetes-related labs (HbA1c, glucose, lipids, renal, microalbumin) with ADA Standards of Care assessment and complications screening status.
  Use when user asks to "review diabetes labs", "check A1c", "diabetes panel", "glucose control", "diabetes management review",
  "diabetic screening", or needs assessment of glycemic control and diabetes complication monitoring.
  Do NOT use for general lab review, non-diabetic renal function, or lipid-only analysis.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: lab-diagnostics
---

# Diabetes Panel Review

## Overview

Pull all diabetes-relevant laboratory Observations: HbA1c, fasting glucose, random glucose, urine microalbumin, urine albumin-to-creatinine ratio, lipid panel, eGFR, creatinine, and serum potassium. Trend HbA1c over time. Classify glycemic control per ADA Standards of Care 2024 targets. Assess complications screening status for retinopathy, nephropathy, and neuropathy. Evaluate cardiovascular risk factors. Generate a structured diabetes management summary with actionable recommendations.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Observation | Lab results (HbA1c, glucose, lipids, renal, urine) | code, valueQuantity, effectiveDateTime, interpretation, referenceRange |
| Condition | Diabetes diagnosis and complications | code, clinicalStatus, onsetDateTime |
| MedicationStatement | Diabetes medications | medicationCodeableConcept, status, dosage |
| Procedure | Screening procedures (eye exam, foot exam) | code, performedDateTime, status |
| Patient | Demographics for target individualization | birthDate, gender |
| CarePlan | Existing diabetes care plans | status, activity, period |

## Instructions

### Step 1: Confirm Diabetes Diagnosis and Type

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|44054006,http://snomed.info/sct|73211009,http://snomed.info/sct|46635009&clinical-status=active"
```

SNOMED codes:
- 44054006 = Type 2 Diabetes Mellitus
- 73211009 = Type 1 Diabetes Mellitus
- 46635009 = Type 1 Diabetes Mellitus (alternate)

If no Condition found, check for HbA1c >= 6.5% or diabetes medications as proxy evidence. Note the diagnosis gap.

Extract `onsetDateTime` to determine disease duration (affects target selection).

### Step 2: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age. Age determines A1c target individualization:
- Younger adults, short disease duration, no CVD: target < 6.5%
- Most adults: target < 7.0%
- Older adults (> 65), long disease duration, significant comorbidities: target < 8.0%
- Limited life expectancy, extensive comorbidities: avoid symptomatic hyperglycemia, less stringent targets

### Step 3: Pull HbA1c History

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=4548-4&_sort=-date&_count=10"
```

LOINC 4548-4 = Hemoglobin A1c. Retrieve up to 10 values to build a trend (approximately 2.5 years at quarterly intervals).

Classify most recent HbA1c:
- **Well-controlled**: < 7.0% (or individualized target)
- **Moderate**: 7.0% - 9.0%
- **Poorly controlled**: > 9.0%
- **Hypoglycemia risk**: < 6.0% in patients on insulin or sulfonylureas

Calculate:
- Trend direction: improving, stable, worsening
- Average A1c over past 4 readings
- Estimated average glucose (eAG) = 28.7 * A1c - 46.7 (mg/dL)

### Step 4: Pull Glucose Values

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2345-7,1558-6,2339-0&date=ge[90-days-ago]&_sort=-date&_count=20"
```

LOINC codes:
- 2345-7 = Glucose (general)
- 1558-6 = Fasting glucose
- 2339-0 = Glucose [Mass/volume] in Blood

Assess:
- Fasting glucose target: 80-130 mg/dL (ADA)
- Post-prandial target: < 180 mg/dL
- Hypoglycemia events: < 70 mg/dL (Level 1), < 54 mg/dL (Level 2), altered mental status (Level 3)
- Hyperglycemia: > 250 mg/dL warrants evaluation for DKA/HHS

### Step 5: Pull Renal Function (Nephropathy Screening)

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=33914-3,14959-1,2160-0,48642-3,69405-9&_sort=-date&_count=5"
```

LOINC codes:
- 33914-3 = Estimated GFR (CKD-EPI)
- 14959-1 = Urine microalbumin [Mass/volume]
- 2160-0 = Creatinine [Mass/volume] in Serum
- 48642-3 = Urine albumin/creatinine ratio (UACR)
- 69405-9 = Urine albumin/creatinine ratio (alternate)

Classify albuminuria (ADA/KDIGO):
- **A1 (normal)**: UACR < 30 mg/g
- **A2 (moderately increased, formerly "microalbuminuria")**: UACR 30-300 mg/g
- **A3 (severely increased, formerly "macroalbuminuria")**: UACR > 300 mg/g

Screening interval: annually for Type 2 (from diagnosis), Type 1 (after 5 years duration).

### Step 6: Pull Lipid Panel

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2093-3,18262-6,2085-9,2571-8&_sort=-date&_count=5"
```

LOINC codes:
- 2093-3 = Total Cholesterol
- 18262-6 = LDL Cholesterol (calculated)
- 2085-9 = HDL Cholesterol
- 2571-8 = Triglycerides

ADA targets for diabetes:
- LDL < 100 mg/dL (general); < 70 mg/dL (established ASCVD)
- Triglycerides < 150 mg/dL
- HDL > 40 mg/dL (male), > 50 mg/dL (female)
- Statin indicated: all diabetics age 40-75 (moderate-intensity minimum); high-intensity if ASCVD or ASCVD risk >= 20%

### Step 7: Check Complications Screening Status

**Retinopathy:**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|274795007,http://snomed.info/sct|252779009&_sort=-date&_count=1"
```
SNOMED 274795007 = Examination of retina, 252779009 = Dilated fundus examination.
Screening interval: annually (may extend to biennial if no retinopathy on two consecutive exams).

**Foot Exam:**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|401191002,http://snomed.info/sct|284473002&_sort=-date&_count=1"
```
SNOMED 401191002 = Diabetic foot examination, 284473002 = Foot examination.
Screening interval: annually (every visit if peripheral neuropathy or PAD present).

**Neuropathy assessment** -- check for monofilament testing or nerve conduction:
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|165242009&_sort=-date&_count=1"
```
SNOMED 165242009 = Peripheral nerve function test.

### Step 8: Retrieve Current Diabetes Medications

```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```

Categorize diabetes medications:
- **Metformin**: first-line, check renal function (contraindicated eGFR < 30)
- **SGLT2 inhibitors** (empagliflozin, dapagliflozin, canagliflozin): cardiovascular and renal protection; check eGFR (reduced efficacy < 30)
- **GLP-1 receptor agonists** (semaglutide, liraglutide, dulaglutide): cardiovascular benefit, weight loss
- **DPP-4 inhibitors** (sitagliptin, linagliptin, saxagliptin): weight-neutral, dose-adjust for renal
- **Sulfonylureas** (glipizide, glimepiride): hypoglycemia risk, caution in elderly/renal impairment
- **Insulin**: basal (glargine, detemir, degludec), bolus (lispro, aspart, glulisine), mixed
- **Thiazolidinediones** (pioglitazone): fluid retention, avoid in CHF

Assess medication escalation per ADA algorithm:
1. Metformin monotherapy (if A1c > 6.5%)
2. Add second agent based on patient characteristics: SGLT2i if CKD/HF, GLP-1 RA if ASCVD/obesity
3. Add third agent or insulin if A1c remains above target after 3-6 months

### Step 9: Format Output

```
DIABETES PANEL REVIEW -- [Patient Name] ([Age][Sex])
=====================================================
Diagnosis: [Type 1/Type 2] Diabetes Mellitus
Duration:  [years since onset]
A1c Target: [individualized target with rationale]

GLYCEMIC CONTROL
================
Most Recent HbA1c: [value]% ([date]) -- [WELL-CONTROLLED / MODERATE / POORLY CONTROLLED]
  Estimated Average Glucose: [eAG] mg/dL
A1c Trend:
  [date1]: [value]%
  [date2]: [value]%
  [date3]: [value]%
  [date4]: [value]%
  Direction: [improving / stable / worsening]

Recent Glucose Values:
  Fasting range: [min]-[max] mg/dL (target 80-130)
  Episodes < 70 mg/dL: [count] (hypoglycemia Level 1)
  Episodes < 54 mg/dL: [count] (hypoglycemia Level 2)

NEPHROPATHY SCREENING
=====================
eGFR:  [value] mL/min/1.73m2 ([date]) -- CKD Stage [stage]
UACR:  [value] mg/g ([date]) -- Albuminuria Category [A1/A2/A3]
Last screening: [date] ([months] ago) -- [CURRENT / OVERDUE]

LIPID PANEL
============
Total Cholesterol: [value] mg/dL
LDL:              [value] mg/dL  (target < [70 or 100])  [AT GOAL / ABOVE GOAL]
HDL:              [value] mg/dL
Triglycerides:    [value] mg/dL  (target < 150)

COMPLICATIONS SCREENING STATUS
===============================
Retinal Exam:     [date] -- [CURRENT / OVERDUE by X months]
Foot Exam:        [date] -- [CURRENT / OVERDUE by X months]
Neuropathy Screen: [date] -- [CURRENT / OVERDUE]

CURRENT MEDICATIONS
===================
[medication list with doses]

RECOMMENDATIONS
===============
1. [Glycemic management recommendation]
2. [Screening recommendation if overdue]
3. [Medication adjustment if indicated]
4. [Referral if needed]
```

## Examples

### Example 1: Well-Controlled Type 2 Diabetes

**User says:** "Review diabetes labs for patient 22190"

**Actions:**
1. `fhir_search` Condition?patient=22190&code=http://snomed.info/sct|44054006&clinical-status=active -- Type 2 DM, onset 2019
2. `fhir_read` Patient/22190 -- 58-year-old female
3. `fhir_search` Observation?patient=22190&code=4548-4&_sort=-date&_count=10 -- A1c values: 6.8%, 6.9%, 7.2%, 7.5% (improving trend)
4. `fhir_search` Observation?patient=22190&code=48642-3&_sort=-date&_count=5 -- UACR 18 mg/g (A1 normal)
5. `fhir_search` Observation?patient=22190&code=33914-3&_sort=-date&_count=5 -- eGFR 82 (G2)
6. `fhir_search` Observation?patient=22190&code=18262-6&_sort=-date&_count=5 -- LDL 88 mg/dL
7. `fhir_search` Procedure?patient=22190&code=http://snomed.info/sct|274795007&_sort=-date&_count=1 -- retinal exam 8 months ago
8. `fhir_search` MedicationStatement?patient=22190&status=active -- metformin 1000mg BID, empagliflozin 10mg daily

**Result:** Well-controlled. A1c 6.8% (target < 7.0%), improving trend from 7.5% over 18 months. eGFR 82, UACR normal. LDL at goal. Retinal exam current. Foot exam overdue (> 12 months). Continue current regimen. Schedule foot exam.

### Example 2: Poorly Controlled With Missing Screenings

**User says:** "Diabetes panel for patient def-440"

**Actions:**
1. `fhir_search` Condition?patient=def-440&code=http://snomed.info/sct|44054006&clinical-status=active -- Type 2 DM, onset 2015
2. `fhir_read` Patient/def-440 -- 67-year-old male
3. `fhir_search` Observation?patient=def-440&code=4548-4&_sort=-date&_count=10 -- A1c values: 9.4%, 9.1%, 8.8%, 8.5% (worsening)
4. `fhir_search` Observation?patient=def-440&code=48642-3&_sort=-date&_count=5 -- UACR 185 mg/g (A2, moderately increased)
5. `fhir_search` Observation?patient=def-440&code=33914-3&_sort=-date&_count=5 -- eGFR 52 (G3a)
6. `fhir_search` Observation?patient=def-440&code=18262-6&_sort=-date&_count=5 -- LDL 142 mg/dL
7. `fhir_search` Procedure?patient=def-440&code=http://snomed.info/sct|274795007&_sort=-date&_count=1 -- no retinal exam found
8. `fhir_search` MedicationStatement?patient=def-440&status=active -- metformin 500mg BID, glipizide 5mg daily

**Result:** Poorly controlled. A1c 9.4% (target < 8.0% given age/comorbidities), worsening trend. eGFR 52 with moderately increased albuminuria -- high risk of CKD progression. LDL 142, above target. Recommendations: (1) Add SGLT2 inhibitor (renal/CV protection, A1c reduction). (2) Discontinue glipizide, consider GLP-1 RA for A1c and weight. (3) Start or intensify statin -- high-intensity indicated. (4) Start ACEi/ARB for albuminuria if not already on one. (5) Urgent retinal exam referral -- never documented. (6) Schedule foot exam.

## Troubleshooting

### No HbA1c Observations Found

- Try alternate LOINC codes: 4548-4 (primary), 17856-6 (HbA1c/Hemoglobin.total in Blood), 59261-8 (HbA1c by IFCC).
- Check if `code.coding` uses a local code system rather than LOINC. Search broadly: `Observation?patient=[id]&code:text=a1c`.
- If truly no A1c exists, note as a gap and recommend ordering.

### UACR Not Available as a Calculated Ratio

- Some systems store urine microalbumin (14959-1) and urine creatinine (2161-8) as separate Observations rather than the calculated ratio. Pull both and calculate: UACR = (urine albumin mg/L) / (urine creatinine g/L) = mg/g.
- Alternatively search for the albumin-creatinine ratio panel (LOINC 9318-7).

### Diabetes Condition Resource Not Found But Patient Is Clearly Diabetic

- Patient may have been coded with ICD-10 (E11.x) instead of SNOMED. Try: `Condition?patient=[id]&code=http://hl7.org/fhir/sid/icd-10-cm|E11`.
- If no Condition at all, the presence of HbA1c >= 6.5% plus diabetes medications (metformin, insulin, SGLT2i, etc.) is sufficient clinical evidence. Note the documentation gap.

## Related Skills

- `lab-result-interpreter` -- for detailed interpretation of individual lab values
- `renal-function-dashboard` -- for comprehensive CKD staging in diabetic nephropathy
- `critical-value-alert-generator` -- if glucose or potassium hits critical thresholds
- `medication-reconciliation` -- for full diabetes medication assessment
