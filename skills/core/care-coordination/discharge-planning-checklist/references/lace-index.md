# LACE Index for Readmission Risk

## Overview

The LACE index is a validated scoring tool that predicts the risk of death or unplanned readmission within 30 days of hospital discharge. Originally validated by van Walraven et al. (CMAJ 2010;182(6):551-557).

## Components

### L - Length of Stay (days)

| LOS (days) | Score |
|------------|-------|
| 1          | 1     |
| 2          | 2     |
| 3          | 3     |
| 4-6        | 4     |
| 7-13       | 5     |
| >= 14      | 7     |

**FHIR Calculation:**
- `Encounter.period.start` to current date (or `Encounter.period.end` if discharged)
- Use calendar days, not hours

### A - Acuity of Admission

| Admission Type | Score |
|----------------|-------|
| Elective / Scheduled | 0 |
| Emergent / Urgent (via ED) | 3 |

**FHIR Calculation:**
- Check `Encounter.hospitalization.admitSource`
  - Code "emd" (Emergency Department) = 3
  - Code "outp" (Outpatient) = 0
  - Code "hosp-trans" (Hospital Transfer) = 3
- Alternative: Check if an ED Encounter (`class` = "EMER") exists within 24 hours prior to inpatient Encounter start
  ```
  fhir_search Encounter?patient=[id]&class=EMER&date=ge[admission-date-minus-1-day]&date=le[admission-date]
  ```
  - If found: score = 3
  - If not found: score = 0

### C - Comorbidities (Charlson Comorbidity Index)

Score based on the Charlson Comorbidity Index, which assigns weights to 17 conditions.

#### Weight = 1
| Condition | SNOMED Code | ICD-10 Range |
|-----------|-------------|--------------|
| Myocardial infarction | 22298006 | I21-I22 |
| Congestive heart failure | 42343007 | I50 |
| Peripheral vascular disease | 400047006 | I70-I79 |
| Cerebrovascular disease | 62914000 | I60-I69, G45-G46 |
| Dementia | 52448006 | F00-F03, G30 |
| Chronic pulmonary disease | 413839001 | J40-J47 |
| Connective tissue disease | 105969002 | M05-M06, M32-M34 |
| Peptic ulcer disease | 13200003 | K25-K28 |
| Mild liver disease | 235856003 | K70, K73-K74 |
| Diabetes without complications | 44054006 | E10-E11 (.0-.1, .6, .8-.9) |

#### Weight = 2
| Condition | SNOMED Code | ICD-10 Range |
|-----------|-------------|--------------|
| Hemiplegia | 50582007 | G81-G83 |
| Moderate-severe renal disease | 709044004 | N18.3-N18.5 |
| Diabetes with complications | 74627003 | E10-E11 (.2-.5) |
| Malignancy (non-metastatic) | 363346000 | C00-C75 |
| Leukemia | 93143009 | C91-C95 |
| Lymphoma | 118600007 | C81-C85, C88, C96 |

#### Weight = 3
| Condition | SNOMED Code | ICD-10 Range |
|-----------|-------------|--------------|
| Moderate-severe liver disease | 19943007 | K72, K76.6-K76.7 |

#### Weight = 6
| Condition | SNOMED Code | ICD-10 Range |
|-----------|-------------|--------------|
| Metastatic solid tumor | 128462008 | C77-C80 |
| AIDS | 62479008 | B20-B24 |

**FHIR Calculation:**
```
fhir_search Condition?patient=[id]&clinical-status=active
```
- Match each active Condition against the SNOMED codes above
- Sum the weights for the Charlson score
- Map Charlson score to LACE C component:

| Charlson Score | LACE C Score |
|----------------|--------------|
| 0              | 0            |
| 1              | 1            |
| 2              | 2            |
| 3              | 3            |
| >= 4           | 5            |

### E - Emergency Department Visits (Prior 6 Months)

| ED Visits (past 6 months) | Score |
|---------------------------|-------|
| 0                         | 0     |
| 1                         | 1     |
| 2                         | 2     |
| 3                         | 3     |
| >= 4                      | 4     |

**FHIR Calculation:**
```
fhir_search Encounter?patient=[id]&class=http://terminology.hl7.org/CodeSystem/v3-ActCode|EMER&date=ge[6-months-ago]
```
- Count the number of entries in the Bundle
- Exclude the current encounter if it originated from the ED
- Cap score at 4

## Total Score Calculation

**LACE Score = L + A + C + E**

| Total LACE Score | Risk Level | Expected 30-Day Readmission Rate |
|------------------|------------|----------------------------------|
| 0-4              | Low        | ~2-5%                            |
| 5-9              | Moderate   | ~10-15%                          |
| 10-12            | High       | ~20-25%                          |
| 13+              | Very High  | ~30%+                            |

### Score Range
- Minimum: 0 (elective admission, 1 day stay, no comorbidities, no prior ED visits)
- Maximum: 19 (emergent, 14+ day stay, Charlson >= 4, 4+ ED visits)

## Clinical Action by Risk Level

### Low Risk (0-4)
- Standard discharge planning
- PCP follow-up within 14 days
- Written discharge instructions
- Medication reconciliation

### Moderate Risk (5-9)
- Enhanced discharge planning
- PCP follow-up within 7 days
- Post-discharge phone call within 72 hours
- Medication reconciliation with pharmacist review
- Ensure outpatient pharmacy access confirmed

### High Risk (10-12)
- Intensive discharge planning
- PCP follow-up within 3-7 days
- Post-discharge phone call within 48 hours
- Transitional care management enrollment
- Home health referral evaluation
- Social work consultation
- Consider 30-day readmission reduction program

### Very High Risk (13+)
- All high-risk interventions plus:
- Transitional care nurse assignment
- Post-discharge home visit within 48-72 hours
- Weekly phone monitoring for 30 days
- Multidisciplinary case conference
- Consider extended observation or step-down facility
- Advance care planning discussion if appropriate

## Limitations

- LACE was validated on a general medical population; may be less predictive for:
  - Surgical patients (consider HOSPITAL score instead)
  - Psychiatric admissions
  - Obstetric admissions
  - Pediatric patients
- Does not account for social determinants of health
- Does not include medication complexity or polypharmacy
- Does not differentiate planned vs unplanned readmissions in the ED visit count
- Should be used as one component of a comprehensive discharge assessment, not as sole decision-making tool

## Alternative Readmission Risk Scores

### HOSPITAL Score
- Hemoglobin at discharge (< 12 g/dL)
- Oncology service discharge
- Sodium at discharge (< 135 mEq/L)
- Procedure during admission
- Index admission type (urgent/emergent)
- Number of admissions in prior year
- Length of stay >= 5 days

### BOOST 8P Tool
- Problem medications (high-risk meds)
- Psychological (depression, cognitive impairment)
- Principal diagnosis (cancer, stroke, DM, COPD, HF)
- Polypharmacy (>= 5 medications)
- Poor health literacy
- Patient support (inadequate social support)
- Prior hospitalization (within 6 months)
- Palliative care needs
