---
name: quality-measure-dashboard
description: |
  Calculates HEDIS-style quality measures from FHIR data, reports measure rates with gap-to-goal analysis, and identifies non-compliant patients. Use when user asks to "calculate quality measures", "HEDIS rates", "quality dashboard", "star rating", "measure compliance", "gap to goal", "quality reporting", or mentions "CMS measures", "HEDIS", "quality scores". Do NOT use for individual patient chart review, panel-level aggregate statistics without formal measure logic (use patient-panel-overview), or preventive care audits (use preventive-care-compliance-report).
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: population-health
---

# Quality Measure Dashboard

## Overview

Calculate formal HEDIS and CMS quality measures using FHIR R4 data. Apply numerator/denominator/exclusion logic per measure specification. Report measure rates, compare against benchmarks, identify patients falling out of compliance, and track trends. Supports measures for diabetes care (A1c, eye exam, nephropathy), hypertension control, breast cancer screening, colorectal cancer screening, depression screening, statin therapy, and tobacco cessation. See references/hedis-measure-definitions.md for full specifications.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Condition | Denominator identification (diagnosis) | code, clinicalStatus, onsetDateTime, subject |
| Patient | Demographics for age/sex stratification | birthDate, gender |
| Observation | Lab results, vitals, screening scores | code, valueQuantity, effectiveDateTime, status |
| Procedure | Screening procedures, surgeries (exclusions) | code, performedDateTime, status |
| Encounter | Visit-based measure eligibility | class, period, type, status |
| MedicationRequest | Medication-based measures (statin therapy) | medicationCodeableConcept, status, authoredOn |
| DiagnosticReport | Screening results (mammogram, colonoscopy) | code, effectiveDateTime, status, conclusion |
| Immunization | Immunization-related measures | vaccineCode, occurrenceDateTime, status |

## Instructions

### Step 1: Select Measures to Calculate

Determine which measures the user wants. If unspecified, calculate the core set:
1. **HBD** - Hemoglobin A1c Control for Patients with Diabetes (A1c < 8%)
2. **CBP** - Controlling High Blood Pressure (BP < 140/90)
3. **BCS** - Breast Cancer Screening (mammogram in 27 months)
4. **COL** - Colorectal Cancer Screening (per USPSTF intervals)
5. **DSF** - Depression Screening and Follow-Up
6. **SPD** - Statin Therapy for Patients with Diabetes
7. **TCC** - Tobacco Cessation Counseling

### Step 2: Build Denominator Population

For each measure, query the eligible population per specification. See references/hedis-measure-definitions.md for detailed criteria.

**HBD - Diabetes A1c Control:**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "code=http://snomed.info/sct|44054006,http://snomed.info/sct|46635009&clinical-status=active&_count=200"
```
SNOMED 44054006 = Type 2 DM, 46635009 = Type 1 DM. Denominator = patients 18-75 with active diabetes.

**CBP - Blood Pressure Control:**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "code=http://snomed.info/sct|38341003&clinical-status=active&_count=200"
```
Denominator = patients 18-85 with active hypertension diagnosis.

**BCS - Breast Cancer Screening:**
```
Tool: fhir_search
resourceType: "Patient"
queryParams: "gender=female&birthdate=ge[73-years-ago]&birthdate=le[52-years-ago]&_count=200"
```
Denominator = females age 52-74.

Filter each denominator by age using Patient.birthDate.

### Step 3: Apply Exclusions

For each measure, query and remove excluded patients.

**HBD exclusions (hospice, ESRD):**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "code=http://snomed.info/sct|46177005&clinical-status=active&_count=100"
```
SNOMED 46177005 = End-stage renal disease.

```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "type=http://snomed.info/sct|385763009&date=ge[measurement-year-start]&_count=100"
```
SNOMED 385763009 = Hospice care.

**BCS exclusions (bilateral mastectomy):**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "code=http://snomed.info/sct|27865001&_count=100"
```
SNOMED 27865001 = Bilateral mastectomy.

Remove excluded patients from the denominator.

### Step 4: Calculate Numerator

Query evidence of compliance for denominator patients.

**HBD numerator (A1c < 8%):**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=http://loinc.org|4548-4&date=ge[measurement-year-start]&_sort=-date&_count=200"
```
LOINC 4548-4 = HbA1c. Take most recent per patient. Numerator = patients with A1c < 8.0%.

**CBP numerator (BP < 140/90):**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=http://loinc.org|85354-9&date=ge[measurement-year-start]&_sort=-date&_count=200"
```
LOINC 85354-9 = Blood pressure panel. Numerator = most recent systolic < 140 AND diastolic < 90.

**BCS numerator (mammogram in 27 months):**
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "code=http://loinc.org|24606-6&date=ge[27-months-ago]&_count=200"
```
LOINC 24606-6 = Breast screening mammogram. If no DiagnosticReport, try Procedure:
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "code=http://snomed.info/sct|71651007&date=ge[27-months-ago]&_count=200"
```

**COL numerator:**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "code=http://snomed.info/sct|73761001&date=ge[10-years-ago]&_count=200"
```
SNOMED 73761001 = Colonoscopy (within 10 years). Also check FIT/FOBT within 1 year (LOINC 29771-3) and flexible sigmoidoscopy within 5 years (SNOMED 44441009).

### Step 5: Compute Measure Rates

For each measure:
- **Rate** = Numerator / (Denominator - Exclusions) * 100
- **Gap count** = Denominator - Exclusions - Numerator
- **Gap-to-goal** = Target rate - Current rate

Use benchmarks from references/quality-reporting.md.

### Step 6: Identify Non-Compliant Patients

For each measure, list patients in the denominator but NOT in the numerator (after exclusions). These are the "gaps in care."

Sort by:
1. Measures with most gaps first
2. Within each measure, patients with longest time since last compliant event

### Step 7: Present Dashboard

```
QUALITY MEASURE DASHBOARD
==========================
Measurement Period: [start] to [end]
Report Date: [today]

MEASURE SUMMARY
---------------
Measure                      | Denominator | Exclusions | Numerator | Rate   | Target | Gap
-----------------------------|-------------|------------|-----------|--------|--------|----
A1c Control (< 8%)           |          87 |          3 |        68 | 81.0%  | 85.0%  | -4.0%
BP Control (< 140/90)        |         142 |          5 |       102 | 74.5%  | 80.0%  | -5.5%
Breast Cancer Screening      |          65 |          2 |        51 | 81.0%  | 82.0%  | -1.0%
Colorectal Cancer Screening  |          98 |          1 |        72 | 74.2%  | 78.0%  | -3.8%
Depression Screening          |         310 |          0 |       248 | 80.0%  | 83.0%  | -3.0%
Statin Therapy (Diabetes)    |          87 |         12 |        63 | 84.0%  | 85.0%  | -1.0%

STAR RATING ESTIMATE: 3.5 stars (see references/quality-reporting.md for methodology)

PATIENTS WITH CARE GAPS
------------------------
A1c Control (16 patients with gaps):
  1. [Name] (MRN: [mrn]) - Last A1c: 8.4% on [date] - ACTION: Intensify therapy
  2. [Name] (MRN: [mrn]) - No A1c in measurement year - ACTION: Order A1c
  ...

BP Control (35 patients with gaps):
  1. [Name] (MRN: [mrn]) - Last BP: 148/92 on [date] - ACTION: Medication adjustment
  ...
```

## Examples

### Example 1: Full Quality Dashboard

**User says**: "Run my quality dashboard. Show me all HEDIS measures."

**Actions**:
1. For each measure (HBD, CBP, BCS, COL, DSF, SPD), query denominator conditions/demographics.
2. Apply exclusions per measure specification.
3. Query numerator evidence (labs, procedures, diagnostic reports).
4. Compute rates and gaps.
5. Identify non-compliant patients per measure.

**Result**:
```
QUALITY MEASURE DASHBOARD
==========================
Measurement Period: 2024-01-01 to 2024-12-31
Report Date: 2024-11-15

MEASURE SUMMARY
---------------
Measure                      | Denom | Excl | Numer | Rate  | Target | Gap
-----------------------------|-------|------|-------|-------|--------|------
A1c Control (< 8%)           |    87 |    3 |    68 | 81.0% | 85.0%  | -4.0%
BP Control (< 140/90)        |   142 |    5 |   102 | 74.5% | 80.0%  | -5.5%
Breast Cancer Screening      |    65 |    2 |    51 | 81.0% | 82.0%  | -1.0%
Colorectal Cancer Screening  |    98 |    1 |    72 | 74.2% | 78.0%  | -3.8%
Depression Screening          |   310 |    0 |   248 | 80.0% | 83.0%  | -3.0%
Statin Therapy (DM)          |    87 |   12 |    63 | 84.0% | 85.0%  | -1.0%

PRIORITY ACTIONS:
- BP Control has largest absolute gap (35 patients). Focus medication titration.
- Colorectal Cancer Screening: 25 patients need outreach for colonoscopy/FIT scheduling.
- Depression Screening: 62 patients missing PHQ-9 this year.
```

### Example 2: Single Measure Deep Dive

**User says**: "How are we doing on A1c control? Show me who's not at goal."

**Actions**:
1. Query diabetes Conditions (active T1DM + T2DM). Returns 87 patients.
2. Filter ages 18-75. Remove hospice/ESRD exclusions (3 excluded).
3. Query A1c observations in measurement year. Take most recent per patient.
4. Classify: < 7% (well controlled), 7-8% (at HEDIS goal), 8-9% (above goal), > 9% (poorly controlled), no data.

**Result**:
```
A1c CONTROL MEASURE - DEEP DIVE
================================
Denominator: 87 | Exclusions: 3 | Eligible: 84
Numerator (A1c < 8%): 68 | Rate: 81.0% | Target: 85.0%

A1c DISTRIBUTION
----------------
< 7.0% (well controlled):    52 (61.9%)
7.0-7.9% (at HEDIS goal):    16 (19.0%)
8.0-8.9% (above goal):        7 (8.3%)
>= 9.0% (poorly controlled):  5 (6.0%)
No A1c this year:              4 (4.8%)

NON-COMPLIANT PATIENTS (16):
1. Marcus Hall (MRN: 55102) - A1c: 11.8% (2024-03-15) - On metformin only - RECOMMEND: Add second agent
2. Nancy Chen (MRN: 44290) - A1c: 10.2% (2024-06-01) - On metformin + glipizide - RECOMMEND: Switch to GLP-1/SGLT2
3. Thomas Brown (MRN: 33871) - A1c: 9.6% (2024-08-20) - Not on any diabetes meds - RECOMMEND: Start therapy
4. Susan Park (MRN: 22150) - No A1c since 2023-09-01 - RECOMMEND: Order stat A1c
...

CLOSING THE GAP:
- 4 patients need A1c ordered (quick win)
- 5 patients with A1c > 9% need urgent therapy intensification
- 7 patients with A1c 8.0-8.9% need follow-up within 3 months
```

## Troubleshooting

### Measure denominator is unexpectedly small or zero

- The FHIR server may not use SNOMED codes. Try ICD-10: `code=http://hl7.org/fhir/sid/icd-10-cm|E11` for diabetes, `I10` for hypertension.
- Some servers require exact code match. For diabetes, also include: E11.0, E11.1, E11.2, etc. Use prefix matching if supported: `code=http://hl7.org/fhir/sid/icd-10-cm|E11`.
- Check that `clinical-status` is populated. Some systems leave it blank for active conditions.

### DiagnosticReport not available for screening measures

- Not all FHIR servers expose DiagnosticReport. Fall back to Procedure searches for the screening event itself.
- Check if the server stores mammograms as Observation resources with imaging LOINC codes.
- For colonoscopy, also search Procedure with CPT code: `code=http://www.ama-assn.org/go/cpt|45378`.

### Numerator count seems too low despite patients being treated

- Verify date range. HEDIS measurement year is typically calendar year. Ensure `date=ge[year-start]` covers the correct period.
- Some measures accept lookback beyond the measurement year (e.g., colonoscopy 10 years). Adjust date filter accordingly.
- Lab results may be stored with `effectiveDateTime` outside the search range if the specimen was collected near year boundaries.

## Related Skills

- `patient-panel-overview` -- aggregate panel metrics without formal measure logic
- `chronic-disease-registry-query` -- disease-specific registry queries with severity data
- `care-gap-identifier` -- individual patient care gap analysis
- `preventive-care-compliance-report` -- USPSTF/ACS preventive care compliance
