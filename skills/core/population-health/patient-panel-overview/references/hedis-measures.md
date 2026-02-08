# HEDIS Measures Reference for Panel Management

Source: NCQA HEDIS (Healthcare Effectiveness Data and Information Set) measure specifications, adapted for panel management context.

## A1c Control -- Comprehensive Diabetes Care (CDC/HBD)

### Hemoglobin A1c Control (< 8%) -- HBD

**Denominator:** Patients 18-75 with diabetes (Type 1 or Type 2) and at least one outpatient visit during the measurement year.

**Denominator Identification:**
- SNOMED: 44054006 (T2DM), 46635009 (T1DM), 73211009 (DM)
- ICD-10: E10.x (T1DM), E11.x (T2DM), E13.x (Other specified DM)
- LOINC for A1c presence: 4548-4 (confirms diabetes monitoring)

**Numerator:** Most recent HbA1c during the measurement year is < 8.0%.
- LOINC: 4548-4 (Hemoglobin A1c/Hemoglobin.total in Blood)
- Take the most recent result if multiple exist

**Exclusions:**
- Hospice care (SNOMED 385763009) during measurement year
- Patients who died during the measurement year
- ESRD or dialysis (SNOMED 46177005, ICD-10 N18.6, CPT 90935-90999)

**Panel Management Implications:**
- Patients with A1c >= 9% should be flagged as "poor control" and prioritized for outreach
- Patients with no A1c in the measurement year are noncompliant with monitoring -- order labs
- Track quarterly trend for patients recently started on new therapy

### HbA1c Poor Control (> 9%) -- HBD (Inverse Measure)

**Note:** This is an inverse measure -- a LOWER rate is better.

**Numerator:** Most recent A1c > 9.0% OR no A1c recorded in the measurement year.

**Panel Action:** Every patient in this numerator needs active intervention -- either medication adjustment, adherence support, or endocrinology referral.

---

## Blood Pressure Control -- Controlling High Blood Pressure (CBP)

**Denominator:** Patients 18-85 with a diagnosis of hypertension and at least one outpatient visit during the measurement year.

**Denominator Identification:**
- SNOMED: 38341003 (Hypertensive disorder)
- ICD-10: I10 (Essential hypertension), I11.x-I13.x (Hypertensive heart/kidney disease)

**Numerator:** Most recent blood pressure during the measurement year has systolic < 140 mmHg AND diastolic < 90 mmHg.
- LOINC: 85354-9 (Blood pressure panel), with components:
  - 8480-6 (Systolic blood pressure)
  - 8462-4 (Diastolic blood pressure)

**Exclusions:**
- Hospice care during measurement year
- ESRD or dialysis
- Pregnancy during measurement year (SNOMED 77386006)
- Patients who died during the measurement year

**Panel Management Notes:**
- Use the most recent BP, not an average
- Office BP preferred; ambulatory BP monitoring (ABPM) acceptable
- Patients with BP >= 180/120 require same-day evaluation (hypertensive crisis)
- Stage 2 hypertension (>= 160/100) warrants visit within 1 month

---

## Statin Therapy for Patients with Diabetes -- SPD

**Denominator:** Patients 40-75 with diabetes (any type) who do NOT have a documented statin contraindication or adverse reaction.

**Numerator:** Patient received a statin dispensing event during the measurement year OR has an active statin prescription.

**Statin Identification (RxNorm ingredient codes):**
- Atorvastatin: 83367
- Rosuvastatin: 301542
- Simvastatin: 36567
- Pravastatin: 42463
- Lovastatin: 6472
- Fluvastatin: 41127
- Pitavastatin: 861634

**Exclusions:**
- Documented statin intolerance (myopathy, rhabdomyolysis)
- Pregnancy or planned pregnancy
- Hospice care
- ESRD/dialysis

**Panel Management Notes:**
- Cross-reference MedicationRequest with statin RxNorm codes
- If no active statin and no exclusion, flag for prescriber review
- High-intensity statins (atorvastatin 40-80mg, rosuvastatin 20-40mg) preferred for ASCVD risk >= 7.5%

---

## Breast Cancer Screening -- BCS

**Denominator:** Women 52-74 as of December 31 of the measurement year.

**Numerator:** Mammogram performed during the measurement year or the year prior (27-month lookback from end of measurement year).
- LOINC: 24606-6 (Mammography screening)
- SNOMED: 71651007 (Mammography), 24623002 (Screening mammography)
- CPT: 77067 (Bilateral screening mammography)

**Exclusions:**
- Bilateral mastectomy (SNOMED 27865001, CPT 19303 bilateral)
- History of bilateral mastectomy

**Panel Management Notes:**
- 27-month lookback accommodates biennial screening
- Query both Procedure and DiagnosticReport resources
- Patients age 40-51 are eligible for shared decision-making per USPSTF but not required for HEDIS

---

## Colorectal Cancer Screening -- COL

**Denominator:** Patients 46-75 as of December 31 of the measurement year.

**Numerator (any one of the following):**
- Colonoscopy in past 10 years (SNOMED 73761001, CPT 45378)
- Flexible sigmoidoscopy in past 5 years (SNOMED 44441009, CPT 45330)
- FIT/FOBT in past 1 year (LOINC 29771-3, 57905-2, 27396-1)
- CT colonography in past 5 years (SNOMED 418714002, CPT 74263)
- FIT-DNA (Cologuard) in past 3 years (LOINC 77353-1, CPT 81528)

**Exclusions:**
- Total colectomy (SNOMED 26390003, CPT 44150-44160)
- Colorectal cancer (SNOMED 363406005, ICD-10 C18-C20)
- Hospice care

---

## Depression Screening and Follow-Up -- DSF

**Denominator:** Patients >= 12 years with at least one eligible encounter during the measurement year.

**Numerator:** Patient screened for depression using a standardized instrument AND, if positive, a follow-up plan documented on the date of the positive screen.

**Screening Instruments (LOINC):**
- PHQ-9: 44261-6
- PHQ-2: 55758-7
- PHQ-A (adolescent): 89206-7
- Edinburgh Postnatal: 71354-5
- PROMIS Depression: 71969-0

**Positive Screen Thresholds:**
- PHQ-9 >= 10 (moderate depression)
- PHQ-2 >= 3 (positive screen, requires PHQ-9 follow-up)

**Follow-Up Actions (if positive):**
- Referral to behavioral health
- Pharmacotherapy initiation
- Additional evaluation ordered
- Suicide risk assessment if PHQ-9 item 9 > 0

**Panel Management Notes:**
- Annual screening required
- Track PHQ-9 scores over time for patients in treatment (target: 50% reduction or score < 5)
- Integrate with visit workflows -- administer PHQ-2/PHQ-9 at check-in

---

## Benchmarks (HEDIS MY 2023, Commercial HMO 50th Percentile)

| Measure | 25th %ile | 50th %ile | 75th %ile | 90th %ile |
|---------|-----------|-----------|-----------|-----------|
| HBD A1c < 8% | 76.0% | 81.0% | 85.0% | 89.0% |
| HBD A1c > 9% (lower is better) | 10.0% | 7.0% | 5.0% | 3.0% |
| CBP < 140/90 | 68.0% | 74.0% | 79.0% | 84.0% |
| BCS Mammography | 76.0% | 80.0% | 83.0% | 86.0% |
| COL Screening | 72.0% | 76.0% | 80.0% | 84.0% |
| DSF Screening | 73.0% | 78.0% | 83.0% | 87.0% |
| SPD Statin Use | 78.0% | 82.0% | 86.0% | 90.0% |

These benchmarks are used for Star Rating estimation. 4-star typically requires >= 75th percentile. 5-star requires >= 90th percentile.
