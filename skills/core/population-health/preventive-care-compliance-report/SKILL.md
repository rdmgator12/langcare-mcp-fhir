---
name: preventive-care-compliance-report
description: |
  Generates a comprehensive preventive care compliance audit per USPSTF, ACS, and CDC guidelines. Pulls procedures, observations, immunizations, demographics, and risk factors to build a compliance scorecard showing up-to-date, overdue, and not-applicable items. Calculates practice-level compliance rates. Use when user asks for "preventive care audit", "screening compliance", "wellness check", "what screenings are due", "annual physical prep", "health maintenance", "prevention scorecard", or mentions "USPSTF", "cancer screening". Do NOT use for single disease management, immunization-only review (use immunization-status-checker), or quality measure reporting (use quality-measure-dashboard).
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: population-health
---

# Preventive Care Compliance Report

## Overview

Audit a patient's or practice's compliance with evidence-based preventive care guidelines from USPSTF (Grade A and B), ACS cancer screening, and CDC immunization schedules. Query Procedure, Observation, Immunization, DiagnosticReport, and Patient resources to determine which screenings, labs, and immunizations are current, overdue, or not applicable based on age, sex, and risk factors. Generate a compliance scorecard and calculate practice-level rates. See references/uspstf-grade-ab.md, references/cancer-screening-guidelines.md, and references/compliance-scoring.md for specifications.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Age, sex, demographics for guideline applicability | birthDate, gender |
| Condition | Risk factors modifying screening recommendations | code, clinicalStatus |
| Observation | Screening results, vitals, lab values | code, valueQuantity, effectiveDateTime, status |
| Procedure | Screening procedures (colonoscopy, mammogram) | code, performedDateTime, status |
| DiagnosticReport | Imaging and pathology results | code, effectiveDateTime, status |
| Immunization | Vaccination compliance | vaccineCode, occurrenceDateTime, status |
| MedicationRequest | Statin/aspirin therapy compliance | medicationCodeableConcept, status |
| FamilyMemberHistory | Family risk factors for screening age adjustment | condition, relationship |

## Instructions

### Step 1: Retrieve Patient Demographics and Risk Profile

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Calculate age, note gender. These determine which guidelines apply.

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Identify risk factors that modify screening recommendations:
- Smoking history (SNOMED 77176002 = current smoker, 8517006 = former smoker)
- Obesity (SNOMED 414916001, or BMI observation)
- Diabetes (SNOMED 44054006)
- Family history of cancer (query FamilyMemberHistory)
- HIV (SNOMED 86406008)
- Immunocompromised status

```
Tool: fhir_search
resourceType: "FamilyMemberHistory"
queryParams: "patient=[patient-id]&_count=50"
```

Check for family history of: breast cancer, colorectal cancer, lung cancer, ovarian cancer, prostate cancer.

### Step 2: Query Screening Procedures and Results

**Cancer screenings:**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|73761001,http://snomed.info/sct|71651007,http://snomed.info/sct|28163009&_sort=-date&_count=20"
```
SNOMED: 73761001 = Colonoscopy, 71651007 = Mammography, 28163009 = Lung CT.

Also check DiagnosticReport:
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&code=http://loinc.org|24606-6,http://loinc.org|24604-1&_sort=-date&_count=20"
```
LOINC: 24606-6 = Mammogram screening, 24604-1 = CT chest low dose (lung cancer screening).

**Lab-based screenings:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=http://loinc.org|4548-4,http://loinc.org|2093-3,http://loinc.org|10839-9,http://loinc.org|21440-3,http://loinc.org|44261-6,http://loinc.org|82810-3&_sort=-date&_count=50"
```
LOINC codes: 4548-4 (A1c), 2093-3 (Total cholesterol), 10839-9 (FIT/FOBT), 21440-3 (HPV), 44261-6 (PHQ-9 depression screen), 82810-3 (Pregnancy status).

**Vitals:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=http://loinc.org|85354-9,http://loinc.org|39156-5&_sort=-date&_count=10"
```
LOINC: 85354-9 (Blood pressure panel), 39156-5 (BMI).

### Step 3: Query Immunization Status

```
Tool: fhir_search
resourceType: "Immunization"
queryParams: "patient=[patient-id]&status=completed&_sort=-date&_count=50"
```

Match against CDC schedule relevant vaccines. See references/uspstf-grade-ab.md for immunization items.

### Step 4: Query Medication-Based Preventive Measures

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=50"
```

Check for:
- Statin therapy (for cardiovascular risk reduction, ages 40-75 with risk factors)
- Low-dose aspirin (shared decision for CVD prevention in select populations)
- Tobacco cessation medications (if active smoker)

### Step 5: Evaluate Each Guideline Item

For each applicable guideline from references/uspstf-grade-ab.md and references/cancer-screening-guidelines.md:

1. **Determine applicability**: Does the patient meet age/sex/risk criteria?
2. **Check compliance**: Is there evidence of the screening/intervention within the recommended interval?
3. **Classify status**:
   - **Up to date**: Evidence within recommended interval
   - **Overdue**: Past the recommended interval or never done when indicated
   - **Due soon**: Within 3 months of recommended date
   - **Not applicable**: Patient outside the guideline criteria (age, sex, risk)

Key guideline items to evaluate (see references for full list):

| Screening | Population | Interval | Evidence Source |
|-----------|-----------|----------|-----------------|
| Mammography | Female 50-74 (40-49 shared decision) | Every 2 years | Procedure/DiagnosticReport |
| Colonoscopy | Age 45-75 | Every 10 years (or FIT annually) | Procedure |
| Cervical cancer (Pap/HPV) | Female 21-65 | Pap every 3y (21-29), Pap+HPV every 5y (30-65) | Observation/Procedure |
| Lung cancer CT | Age 50-80, >= 20 pack-year smoking | Annual | DiagnosticReport |
| Diabetes screening (A1c) | Age 35-70, overweight/obese | Every 3 years | Observation |
| Lipid panel | Males >= 40, Females >= 50 (or 20+ with risk) | Every 5 years | Observation |
| Blood pressure | Age >= 18 | Annual | Observation |
| Depression screening (PHQ-9) | Age >= 12 | Annual | Observation |
| HIV screening | Age 15-65 | At least once; repeat if at risk | Observation |
| Hepatitis C screening | Age 18-79 | At least once | Observation |
| STI screening | Sexually active, risk-based | Varies | Observation |
| Abdominal aortic aneurysm | Males 65-75 who ever smoked | Once | Procedure/DiagnosticReport |
| Statin therapy | Age 40-75 with CVD risk >= 10% | Ongoing | MedicationRequest |
| Tobacco cessation | Current smokers | Every visit | Observation/Procedure |

### Step 6: Build Compliance Scorecard

```
PREVENTIVE CARE COMPLIANCE SCORECARD
======================================
Patient: [name] | Age: [age] | Sex: [sex]
Risk Factors: [list]
Report Date: [today]

CANCER SCREENINGS
-----------------
[UP TO DATE] Breast Cancer (Mammography) - Last: 2023-09-15 - Next due: 2025-09
[OVERDUE]    Colorectal Cancer - No colonoscopy or FIT on record - Due since age 45
[UP TO DATE] Cervical Cancer (Pap + HPV) - Last: 2022-03-01 - Next due: 2027-03
[N/A]        Lung Cancer - Non-smoker, does not meet criteria
[N/A]        Prostate Cancer - Female patient

CARDIOVASCULAR PREVENTION
--------------------------
[UP TO DATE] Blood Pressure - Last: 128/82 on 2024-10-01
[OVERDUE]    Lipid Panel - Last: 2019-05-20 - Overdue (> 5 years)
[UP TO DATE] Statin Therapy - Active prescription: atorvastatin 20mg
[N/A]        AAA Screening - Female patient

METABOLIC SCREENING
-------------------
[UP TO DATE] Diabetes (A1c) - Last: 5.6% on 2024-06-15 - Next due: 2027-06

BEHAVIORAL HEALTH
-----------------
[OVERDUE]    Depression Screening (PHQ-9) - No screening on record
[N/A]        Tobacco Cessation - Non-smoker
[UP TO DATE] Alcohol Screening (AUDIT-C) - Last: 2024-01-15

INFECTIOUS DISEASE
------------------
[UP TO DATE] HIV Screening - Negative 2023-02-01
[UP TO DATE] Hepatitis C Screening - Negative 2022-11-15
[N/A]        STI Screening - Not in risk group

IMMUNIZATIONS
-------------
[UP TO DATE] Influenza 2024-25 - Given 2024-10-01
[OVERDUE]    Shingrix - Age >= 50, not received
[UP TO DATE] Tdap - Given 2020-06-15, next Td 2030
[DUE SOON]   COVID-19 Updated Booster - Last dose > 12 months ago

COMPLIANCE SUMMARY
==================
Total Applicable Items: 14
Up to Date:             9 (64.3%)
Overdue:                3 (21.4%)
Due Soon:               1 (7.1%)
Not Applicable:         5 (excluded from score)

PRIORITY ACTIONS
================
1. [OVERDUE] Schedule colonoscopy or order FIT test
2. [OVERDUE] Order lipid panel
3. [OVERDUE] Administer PHQ-9 depression screening
4. [DUE SOON] Schedule COVID-19 updated booster
```

### Step 7: Practice-Level Compliance (if requested)

If the user requests practice-level rates, iterate across the patient panel:

```
Tool: fhir_search
resourceType: "Patient"
queryParams: "active=true&_count=200"
```

For each patient, run Steps 1-6 (or a simplified version querying key screenings). Aggregate:

```
PRACTICE-LEVEL PREVENTIVE CARE COMPLIANCE
===========================================
Total Active Patients: [N]
Report Date: [today]

Screening                    | Eligible | Compliant | Rate   | Benchmark
-----------------------------|----------|-----------|--------|----------
Mammography                  |       65 |        51 | 78.5%  | 82.0%
Colorectal Cancer Screening  |       98 |        72 | 73.5%  | 78.0%
Cervical Cancer Screening    |       82 |        68 | 82.9%  | 85.0%
Depression Screening         |      310 |       248 | 80.0%  | 83.0%
BP Screening (annual)        |      320 |       295 | 92.2%  | 90.0%
Diabetes Screening           |      145 |       112 | 77.2%  | 80.0%
```

See references/compliance-scoring.md for benchmark sources and scoring methodology.

## Examples

### Example 1: Individual Patient Wellness Check Prep

**User says**: "Prep a wellness check for patient 67890. What screenings are due?"

**Actions**:
1. Read Patient/67890. Female, age 55, DOB 1969-04-22.
2. Search Conditions. Active: hypertension, former smoker (quit 2018).
3. Search FamilyMemberHistory. Mother had breast cancer at age 60.
4. Search Procedures/DiagnosticReports for screening history.
5. Search Observations for lab screenings, vitals.
6. Search Immunizations. Search MedicationRequests.
7. Evaluate each guideline. Family history of breast cancer may warrant earlier/more frequent mammography per ACS.

**Result**:
```
PREVENTIVE CARE - WELLNESS CHECK PREP
=======================================
Patient: Sandra Lee | Age: 55 | Female
Risk Factors: Hypertension, former smoker (quit 2018), family hx breast CA

OVERDUE ITEMS:
1. Colonoscopy - Never done. Due since age 45. SCHEDULE.
2. Lung Cancer CT - Former smoker, 22 pack-years, quit < 15 years ago. Eligible. ORDER.
3. Shingrix - Age >= 50. Not received. ADMINISTER TODAY.

DUE SOON:
4. Mammography - Last 2023-01-15 (23 months ago). Due by March 2025.
   NOTE: Family hx breast CA -- ACS may recommend annual. Discuss.

UP TO DATE:
5. Cervical (Pap + HPV) - 2022-08-01. Next due 2027.
6. Blood pressure - 134/86 on 2024-09-15. Controlled.
7. Lipid panel - 2023-06-01. Next due 2028.
8. A1c (DM screening) - 5.8% on 2024-03-01. Prediabetic range. Recheck in 1 year.
9. PHQ-9 - Score 4 on 2024-09-15. Minimal depression. Annual.
10. Influenza 2024-25 - Given 2024-10-05.
```

### Example 2: Practice-Level Compliance Dashboard

**User says**: "What are our practice-level screening compliance rates?"

**Actions**:
1. Search all active patients. Returns 320 patients.
2. For each major screening, query eligible population and evidence of compliance.
3. Aggregate rates. Compare against national benchmarks.

**Result**:
```
PRACTICE PREVENTIVE CARE COMPLIANCE
=====================================
320 Active Patients | Report Date: 2024-11-15

Screening                    | Eligible | Compliant | Rate  | National Avg | Status
-----------------------------|----------|-----------|-------|--------------|-------
Mammography (F 50-74)        |       65 |        51 | 78.5% | 76.8%        | ABOVE
Colorectal (45-75)           |       98 |        72 | 73.5% | 72.1%        | ABOVE
Cervical (F 21-65)           |       82 |        68 | 82.9% | 83.5%        | AT
Lung CT (50-80, smokers)     |       18 |         6 | 33.3% | 15.4%        | ABOVE
Depression (PHQ-9)           |      310 |       248 | 80.0% | 78.0%        | ABOVE
BP Screening                 |      320 |       295 | 92.2% | 89.0%        | ABOVE
Diabetes (A1c, overweight)   |      145 |       112 | 77.2% | 74.0%        | ABOVE
HIV (18-79, once)            |      290 |       198 | 68.3% | 62.0%        | ABOVE
Hep C (18-79, once)          |      290 |       185 | 63.8% | 60.5%        | ABOVE

LOWEST COMPLIANCE: Hepatitis C screening (63.8%) -- recommend universal screening outreach.
BIGGEST GAP: Lung cancer CT -- only 18 eligible but 12 not screened. Flag for smoking cessation + CT orders.
```

## Troubleshooting

### Screening procedures not found despite being performed

- Screening mammograms may be stored under different SNOMED codes: 71651007 (mammography), 24623002 (screening mammography), or LOINC 24606-6. Try multiple codes.
- Colonoscopy may be coded as 73761001 (SNOMED), 44388 (CPT via `http://www.ama-assn.org/go/cpt`), or stored as DiagnosticReport rather than Procedure.
- Some systems store screening events in Observation resources with result values rather than Procedure resources. Check both.

### Patient has risk factors that change screening recommendations but conditions are not documented

- Query Observation for smoking status (LOINC 72166-2) if Condition for smoking is absent.
- Check BMI observation (LOINC 39156-5) for obesity if no obesity Condition exists.
- If FamilyMemberHistory is empty, note in the report that family history has not been documented and screening recommendations assume average risk.

### Practice-level queries are slow or timeout

- Use `_summary=count` first to estimate volume before pulling full resources.
- For practice-level, query screenings directly rather than per-patient: `fhir_search` Procedure with code filter, no patient filter, date range. Then match to patient demographics.
- Limit to the most impactful screenings (mammography, colonoscopy, cervical, depression) rather than running all simultaneously.

## Related Skills

- `immunization-status-checker` -- deep-dive immunization review with catch-up schedules
- `quality-measure-dashboard` -- formal HEDIS/CMS quality measures (overlaps with some screening measures)
- `patient-panel-overview` -- panel-level chronic disease metrics
- `care-gap-identifier` -- individual patient care gap analysis
- `cancer-screening-guidelines` -- detailed cancer screening reference (references/ file)
