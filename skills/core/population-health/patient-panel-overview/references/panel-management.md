# Panel Management Reference

## Quality Measure Definitions for Common Chronic Diseases

### Diabetes Mellitus (Type 2)

**Registry Inclusion Codes:**
- SNOMED: 44054006 (Type 2 diabetes mellitus), 46635009 (Type 1 diabetes mellitus)
- ICD-10: E11 (Type 2), E10 (Type 1), E13 (Other specified diabetes)

**Key Metrics and Targets:**

| Metric | LOINC Code | Target (General) | Target (Elderly/Complex) | Source |
|--------|-----------|-------------------|--------------------------|--------|
| HbA1c | 4548-4 | < 7.0% | < 8.0% | ADA 2024 |
| Fasting Glucose | 1558-6 | 80-130 mg/dL | 100-180 mg/dL | ADA 2024 |
| LDL Cholesterol | 13457-7 | < 100 mg/dL | < 100 mg/dL | ADA 2024 |
| Blood Pressure (systolic) | 8480-6 | < 130 mmHg | < 140 mmHg | ADA/AHA |
| Blood Pressure (diastolic) | 8462-4 | < 80 mmHg | < 90 mmHg | ADA/AHA |
| Urine Albumin/Creatinine | 14959-1 | < 30 mg/g | < 30 mg/g | KDIGO |
| eGFR | 33914-3 | >= 60 mL/min | >= 60 mL/min | KDIGO |

**Annual Monitoring Requirements:**
- HbA1c: Every 3-6 months (quarterly if not at goal, biannually if stable)
- Lipid panel: Annually
- Urine albumin/creatinine ratio: Annually
- eGFR: Annually (more frequently if declining)
- Dilated eye exam: Annually (every 2 years if no retinopathy and well controlled)
- Foot exam: Annually (comprehensive)
- Blood pressure: Every visit

### Hypertension

**Registry Inclusion Codes:**
- SNOMED: 38341003 (Hypertensive disorder)
- ICD-10: I10 (Essential hypertension), I11 (Hypertensive heart disease), I12 (Hypertensive CKD), I13 (Hypertensive heart and CKD)

**Key Metrics and Targets:**

| Metric | LOINC Code | Target (General) | Target (Age >= 65) | Target (DM/CKD) |
|--------|-----------|-------------------|---------------------|------------------|
| Systolic BP | 8480-6 | < 130 mmHg | < 130 mmHg | < 130 mmHg |
| Diastolic BP | 8462-4 | < 80 mmHg | < 80 mmHg | < 80 mmHg |
| Serum Creatinine | 2160-0 | Baseline-dependent | Baseline-dependent | Monitor trend |
| Serum Potassium | 2823-3 | 3.5-5.0 mEq/L | 3.5-5.0 mEq/L | 3.5-5.0 mEq/L |
| Urine Albumin/Cr | 14959-1 | < 30 mg/g | < 30 mg/g | < 30 mg/g |

**Monitoring Frequency:**
- BP: Every 3-6 months until controlled, then every 6-12 months
- BMP (potassium, creatinine): 2-4 weeks after medication change, then annually
- Urine albumin: Annually if on ACEi/ARB or with diabetes

### Heart Failure (CHF)

**Registry Inclusion Codes:**
- SNOMED: 42343007 (Congestive heart failure), 84114007 (Heart failure), 441481004 (HFrEF), 446221000 (HFpEF)
- ICD-10: I50 (Heart failure), I50.2 (Systolic), I50.3 (Diastolic), I50.4 (Combined)

**Key Metrics and Targets:**

| Metric | LOINC Code | Significance | Action Threshold |
|--------|-----------|--------------|------------------|
| BNP | 42637-9 | Volume status | > 100 pg/mL suggests CHF; > 500 = decompensation |
| NT-proBNP | 33762-6 | Volume status | > 300 pg/mL suggests CHF; > 900 = decompensation |
| LVEF | 10230-1 | Systolic function | < 40% = HFrEF; 40-49% = HFmrEF; >= 50% = HFpEF |
| Weight | 29463-7 | Fluid retention | > 2 lb gain in 1 day or > 5 lb in 1 week |
| Serum Sodium | 2951-2 | Dilutional state | < 135 mEq/L concerning |
| Serum Potassium | 2823-3 | Electrolyte balance | Monitor closely on ACEi/ARB/MRA |
| Serum Creatinine | 2160-0 | Renal function | Rising creatinine may limit GDMT titration |

**GDMT Monitoring:**
- Vitals + weight: Every visit
- BMP: 1-2 weeks after dose change, then every 3-6 months
- BNP/NT-proBNP: At diagnosis, with status change, every 6-12 months if stable

### COPD

**Registry Inclusion Codes:**
- SNOMED: 13645005 (Chronic obstructive pulmonary disease)
- ICD-10: J44 (Other COPD), J44.0 (with acute lower resp infection), J44.1 (with acute exacerbation)

**Key Metrics:**

| Metric | LOINC Code | Staging (GOLD) |
|--------|-----------|----------------|
| FEV1/FVC ratio | 19926-5 | < 0.70 confirms COPD |
| FEV1 % predicted | 20150-9 | >= 80% Mild, 50-79% Moderate, 30-49% Severe, < 30% Very Severe |
| SpO2 | 2708-6 | < 88% at rest = evaluate for O2 |
| CAT Score | 89555-7 | < 10 low impact, >= 10 high impact |
| mMRC Dyspnea | 89557-3 | 0-1 less symptoms, 2-4 more symptoms |

### CKD

**Registry Inclusion Codes:**
- SNOMED: 709044004 (Chronic kidney disease), 431855005 (CKD stage 1-5 codes)
- ICD-10: N18 (Chronic kidney disease), N18.1-N18.6 (stages)

**Key Metrics:**

| Metric | LOINC Code | CKD Staging |
|--------|-----------|-------------|
| eGFR | 33914-3 | G1: >= 90, G2: 60-89, G3a: 45-59, G3b: 30-44, G4: 15-29, G5: < 15 |
| Urine Albumin/Cr | 14959-1 | A1: < 30, A2: 30-300, A3: > 300 mg/g |
| Serum Creatinine | 2160-0 | Used for eGFR calculation |
| Serum Phosphorus | 2777-1 | Target 2.5-4.5 (CKD 3-4), 3.5-5.5 (CKD 5) |
| Serum Calcium | 17861-6 | Target 8.4-10.2 mg/dL |
| PTH | 2731-8 | Stage-dependent target |
| Hemoglobin | 718-7 | < 10 g/dL = evaluate for ESA |

---

## Risk Stratification Criteria

### General Risk Tier Framework

**High Risk:**
- Primary metric severely out of range (e.g., A1c > 9%, BP > 160/100, eGFR < 30, BNP > 500)
- >= 3 active chronic conditions
- ED visit or hospitalization in past 90 days for the target condition
- No clinic visit in > 6 months AND metric out of range
- Documented medication non-adherence
- Recent decline (metric worsening over 2+ consecutive measurements)

**Medium Risk:**
- Primary metric moderately out of range (e.g., A1c 7-9%, BP 140-159/90-99, eGFR 30-59)
- 1-2 additional chronic conditions
- Clinic visit within past 6 months
- Stable or slowly worsening trend

**Low Risk:**
- At target for all primary metrics
- Clinic visit within past 6 months
- Stable or improving trend
- Adherent to medications and follow-up

### Disease-Specific High-Risk Flags

| Disease | High-Risk Indicator | FHIR Query Strategy |
|---------|--------------------|--------------------|
| Diabetes | A1c > 9% | Observation code=4548-4, value > 9 |
| Diabetes | DKA admission | Encounter class=EMER/IMP, reasonCode=E10.1/E11.1 |
| Hypertension | BP > 180/120 | Observation code=85354-9, component value |
| CHF | BNP > 500 | Observation code=42637-9, value > 500 |
| CHF | Weight gain > 5 lb/week | Serial Observation code=29463-7 |
| COPD | 2+ exacerbations/year | Condition code=J44.1 count in 12 months |
| CKD | eGFR decline > 5/year | Serial Observation code=33914-3 trend |
| CKD | New proteinuria | Observation code=14959-1, value > 300 |

---

## Panel Management Best Practices

### Visit Cadence by Risk Tier

| Risk Tier | Visit Frequency | Lab Frequency | Phone/Portal Check-in |
|-----------|----------------|---------------|-----------------------|
| High | Every 1-2 months | Per protocol + PRN | Weekly |
| Medium | Every 3-4 months | Per protocol | Monthly |
| Low | Every 6-12 months | Per protocol | Quarterly |

### Outreach Prioritization Algorithm

Priority score = sum of weighted factors:

1. **Metric severity** (0-3 points): 0 = at goal, 1 = mildly above, 2 = moderately above, 3 = severely above
2. **Time since last visit** (0-3 points): 0 = < 3 months, 1 = 3-6 months, 2 = 6-12 months, 3 = > 12 months
3. **Recent acute event** (0-2 points): 0 = none, 1 = ED visit, 2 = hospitalization in past 90 days
4. **Comorbidity burden** (0-2 points): 0 = target condition only, 1 = 1-2 additional, 2 = 3+ additional
5. **Missing data** (0-1 point): 1 = no lab in measurement period

Patients with score >= 7: Immediate outreach (phone call, schedule within 1 week)
Patients with score 4-6: Standard outreach (letter/portal message, schedule within 1 month)
Patients with score < 4: Routine (schedule at next planned interval)

### Panel Size Benchmarks

| Provider Type | Recommended Panel Size | Adjusted for Chronic Disease |
|---------------|----------------------|------------------------------|
| Primary Care Physician | 1,500-2,000 | Reduce by 20% if > 30% chronic disease |
| Nurse Practitioner | 1,000-1,500 | Same adjustment |
| Care Team (physician + NP + MA) | 2,500-3,500 | Distributed by acuity |

### Data Completeness Thresholds

For reliable panel metrics, require:
- >= 80% of panel patients with a primary metric in the measurement period
- >= 90% of panel patients with an office visit in the past 12 months
- >= 95% of panel patients with verified demographics

If thresholds not met, note data completeness percentage in the report and flag for data quality improvement.
