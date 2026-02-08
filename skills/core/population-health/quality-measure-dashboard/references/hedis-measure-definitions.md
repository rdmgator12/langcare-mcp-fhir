# HEDIS / CMS Quality Measure Definitions

Detailed measure specifications for quality measure calculation. Each measure includes denominator criteria, numerator logic, exclusions, and FHIR query strategies.

---

## 1. Comprehensive Diabetes Care -- HbA1c Control (HBD)

### HbA1c Good Control (< 8.0%)

**Measure ID:** HEDIS HBD / NQF 0059 / CMS 122v12

**Denominator:**
- Age 18-75 as of December 31 of the measurement year
- At least one diagnosis of diabetes (Type 1 or Type 2) on or before December 31
- At least one outpatient, telehealth, or ED encounter during the measurement year or the year prior

**Denominator Codes:**
- Condition: SNOMED 44054006 (T2DM), 46635009 (T1DM), 73211009 (DM NOS)
- Condition: ICD-10 E10.x, E11.x, E13.x
- Encounter types: office visit (CPT 99201-99215), telehealth (CPT 99441-99443), ED (CPT 99281-99285)

**Numerator:**
- Most recent HbA1c result during the measurement year < 8.0%
- LOINC: 4548-4 (Hemoglobin A1c/Hemoglobin.total in Blood)
- Also accept: 4549-2 (A1c in blood by HPLC), 17856-6 (A1c/Hgb.total in blood by HPLC)

**Exclusions:**
- Hospice encounter (SNOMED 385763009, Encounter type)
- ESRD (ICD-10 N18.6, SNOMED 46177005) -- any time during or prior to measurement year
- Kidney transplant (SNOMED 70536003, CPT 50360-50365) -- any time
- Dialysis (CPT 90935-90999, SNOMED 108241001) during measurement year
- Deceased during measurement year

### HbA1c Poor Control (> 9.0%) -- Inverse Measure

**Numerator:** Most recent A1c > 9.0% OR no A1c recorded in measurement year.
- LOWER rate = BETTER performance
- Patients with no A1c are counted as noncompliant (in numerator)

---

## 2. Controlling High Blood Pressure (CBP)

**Measure ID:** HEDIS CBP / NQF 0018 / CMS 165v12

**Denominator:**
- Age 18-85 as of December 31
- Diagnosis of essential hypertension on or before June 30 of the measurement year
- At least one outpatient encounter during the measurement year

**Denominator Codes:**
- Condition: SNOMED 38341003 (Hypertensive disorder), 59621000 (Essential hypertension)
- ICD-10: I10

**Numerator:**
- Most recent BP during the measurement year: systolic < 140 AND diastolic < 90
- BP Observation: LOINC 85354-9 (BP panel)
- Components: LOINC 8480-6 (Systolic), 8462-4 (Diastolic)
- If multiple BPs on the same date, use the lowest systolic and lowest diastolic (may be from different readings)

**Exclusions:**
- Hospice
- ESRD / dialysis
- Pregnancy during measurement year (SNOMED 77386006, ICD-10 O00-O9A)
- Non-acute inpatient stay > 2 days (long-term care, SNF)
- Deceased

**FHIR Query Strategy:**
```
# Denominator: hypertension patients with encounter
fhir_search Condition code=38341003|I10, clinical-status=active, onset <= Jun 30
fhir_search Encounter date=ge[year-start], class=AMB, status=finished

# Numerator: most recent BP
fhir_search Observation code=85354-9, date=ge[year-start], _sort=-date
# Extract component[0] (systolic) and component[1] (diastolic) per patient
```

---

## 3. Breast Cancer Screening (BCS)

**Measure ID:** HEDIS BCS / NQF 2372 / CMS 125v12

**Denominator:**
- Female
- Age 52-74 as of December 31

**Numerator:**
- Mammogram performed October 1 two years prior through December 31 of the measurement year (27-month lookback)
- Procedure: SNOMED 71651007, 24623002 (Screening mammography)
- DiagnosticReport: LOINC 24606-6 (Screening mammography)
- CPT: 77067 (Bilateral screening mammogram), 77063 (digital breast tomosynthesis)

**Exclusions:**
- Bilateral mastectomy (SNOMED 27865001, CPT 19303 with bilateral modifier)
- Unilateral mastectomy x2 (two separate procedures with different laterality)
- History of bilateral mastectomy (ICD-10 Z90.13)
- Hospice
- Advanced illness with frailty (age >= 66 with advanced illness AND frailty)
- Deceased

**FHIR Query Notes:**
- Some FHIR servers store mammograms only as DiagnosticReport, not Procedure
- Check both resource types
- For bilateral mastectomy exclusion, check Procedure + Condition history

---

## 4. Colorectal Cancer Screening (COL)

**Measure ID:** HEDIS COL / NQF 0034 / CMS 130v12

**Denominator:**
- Age 46-75 as of December 31
- At least one outpatient encounter during the measurement year

**Numerator (any one):**

| Test | Lookback | SNOMED | CPT | LOINC |
|------|----------|--------|-----|-------|
| Colonoscopy | 10 years | 73761001 | 45378 | -- |
| Flexible Sigmoidoscopy | 5 years | 44441009 | 45330 | -- |
| FIT/FOBT | 1 year | -- | 82274 | 29771-3, 57905-2 |
| CT Colonography | 5 years | 418714002 | 74263 | -- |
| FIT-DNA (Cologuard) | 3 years | -- | 81528 | 77353-1 |

**Exclusions:**
- Colorectal cancer (SNOMED 363406005, ICD-10 C18.x, C19, C20)
- Total colectomy (SNOMED 26390003, CPT 44150-44160)
- Hospice
- Advanced illness with frailty (age >= 66)
- Deceased

---

## 5. Depression Screening and Follow-Up (DSF)

**Measure ID:** CMS 2v13 / NQF 0418

**Denominator:**
- Age >= 12 as of beginning of measurement year
- At least one qualifying encounter during measurement year

**Numerator (two parts):**
1. Screened for depression using a standardized instrument
2. If screen positive, follow-up plan documented on date of positive screen

**Screening Tools (LOINC):**

| Tool | LOINC | Positive Threshold |
|------|-------|--------------------|
| PHQ-9 | 44261-6 | >= 10 |
| PHQ-2 | 55758-7 | >= 3 |
| PHQ-A (Adolescent) | 89206-7 | >= 10 |
| Edinburgh Postnatal | 71354-5 | >= 10 |
| Beck Depression Inventory | 89208-3 | >= 20 |
| CES-D | 71956-7 | >= 16 |
| PROMIS Depression | 71969-0 | T-score >= 60 |
| Geriatric Depression Scale (GDS) | 48545-8 | >= 5 (short form) |
| Columbia Suicide Severity | 73831-0 | Any positive |

**Exclusions:**
- Active diagnosis of depression or bipolar disorder at time of encounter (already being managed)
- Hospice
- Deceased

**Follow-Up Plan (if positive screen):**
Must be documented on the same date as positive screen. Includes:
- Referral to behavioral health provider
- Pharmacotherapy ordered (antidepressant)
- Additional evaluation with validated tool
- Suicide risk assessment
- Follow-up appointment scheduled within 30 days

---

## 6. Statin Therapy for Patients with Diabetes (SPD)

**Measure ID:** HEDIS SPD / NQF 0541 subset

**Denominator:**
- Age 40-75
- Active diabetes diagnosis (SNOMED 44054006, 46635009; ICD-10 E10.x, E11.x)

**Numerator:**
- Active statin prescription OR statin dispensing event during measurement year

**Statin Medications (RxNorm ingredient):**

| Statin | RxNorm | Intensity |
|--------|--------|-----------|
| Atorvastatin 10-20mg | 83367 | Moderate |
| Atorvastatin 40-80mg | 83367 | High |
| Rosuvastatin 5-10mg | 301542 | Moderate |
| Rosuvastatin 20-40mg | 301542 | High |
| Simvastatin 20-40mg | 36567 | Moderate |
| Pravastatin 40-80mg | 42463 | Moderate |
| Lovastatin 40mg | 6472 | Moderate |
| Fluvastatin 40mg BID / 80mg XL | 41127 | Moderate |
| Pitavastatin 1-4mg | 861634 | Moderate |

**Exclusions:**
- Documented statin adverse reaction/intolerance (AllergyIntolerance with statin code)
- Pregnancy or potential pregnancy (SNOMED 77386006, Observation LOINC 82810-3)
- ESRD/dialysis
- Hospice
- Cirrhosis or hepatic failure (SNOMED 19943007, ICD-10 K74.x)
- Myopathy (SNOMED 129565002)
- Rhabdomyolysis (SNOMED 240131006)

---

## 7. Tobacco Cessation Counseling (TCC)

**Denominator:**
- Age >= 18
- At least one outpatient encounter during the measurement year

**Numerator (two components):**
1. Tobacco use screening (asked about tobacco use)
   - LOINC: 72166-2 (Tobacco smoking status)
   - SNOMED: 365981007 (Finding of tobacco use and exposure)
2. For current users: cessation counseling or pharmacotherapy
   - Procedure: SNOMED 225323000 (Smoking cessation education), 710081004 (Smoking cessation therapy)
   - Medication: Varenicline (RxNorm 637190), Bupropion (RxNorm 42347), NRT (various)

---

## Measure Rate Calculation Formula

```
Eligible Population = Denominator - Exclusions
Measure Rate = (Numerator / Eligible Population) * 100
Gap Count = Eligible Population - Numerator
Gap-to-Goal = Target Rate - Current Rate
```

## Star Rating Methodology (Simplified)

CMS Star Ratings use cut points that shift annually. Approximate thresholds (2024):

| Stars | Percentile Range | Interpretation |
|-------|-----------------|----------------|
| 5 | >= 90th | Excellent |
| 4 | 75th - 89th | Above average |
| 3 | 50th - 74th | Average |
| 2 | 25th - 49th | Below average |
| 1 | < 25th | Poor |

Overall star rating = weighted average across all reported measures. Clinical measures weighted more heavily than patient experience or operational measures.
