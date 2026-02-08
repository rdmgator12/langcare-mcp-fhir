# Pneumonia Scoring Systems Reference

## CURB-65

Simple 5-point score for community-acquired pneumonia severity. Validated for mortality prediction and disposition decisions.

### Scoring Criteria

| Criterion | Points | Threshold | FHIR Source |
|-----------|--------|-----------|-------------|
| **C** - Confusion | 1 | New mental confusion (AMT <=8 or disorientation to person, place, or time) | Condition: SNOMED 130987000, 419284004. Observation: GCS (LOINC 9269-2) <15. |
| **U** - Urea (BUN) | 1 | BUN > 20 mg/dL (>7 mmol/L) | Observation: LOINC 3094-0 (BUN) or 3091-6 (Urea) |
| **R** - Respiratory rate | 1 | >= 30 breaths/min | Observation: LOINC 9279-1 |
| **B** - Blood pressure | 1 | SBP < 90 mmHg OR DBP <= 60 mmHg | Observation: LOINC 8480-6 (SBP), 8462-4 (DBP) |
| **65** - Age | 1 | >= 65 years | Patient.birthDate |

### Risk Stratification

| Score | 30-Day Mortality | Risk Group | Disposition |
|-------|-----------------|------------|-------------|
| 0 | 0.6% | Low | Outpatient treatment |
| 1 | 2.7% | Low | Outpatient (consider observation if borderline) |
| 2 | 6.8% | Moderate | Short inpatient stay or observation |
| 3 | 14% | High | Inpatient admission, consider ICU |
| 4 | 27.8% | Very high | ICU admission |
| 5 | 57% | Extremely high | ICU admission |

### CRB-65 (No Lab Variant)

For settings where BUN is not immediately available (primary care, triage). Remove U criterion. Score 0-4.

| Score | Disposition |
|-------|-------------|
| 0 | Outpatient |
| 1-2 | Consider admission |
| 3-4 | Urgent admission, likely ICU |

---

## PSI/PORT Score (Pneumonia Severity Index)

More comprehensive than CURB-65. Uses demographics, comorbidities, physical exam, and labs to classify into Risk Classes I-V. Validated for 30-day mortality prediction.

### Step 1: Determine if Risk Class I (No Scoring Needed)

Patient is Risk Class I if ALL of the following are true:
- Age < 50
- No comorbidities (neoplasm, liver disease, CHF, cerebrovascular disease, renal disease)
- Normal mental status
- Normal vital signs (HR <125, RR <30, SBP >=90, Temp >=35C and <40C)

If any criterion is NOT met, proceed to point calculation.

### Step 2: Calculate Points

#### Demographic Factors

| Factor | Points |
|--------|--------|
| Age (years) | Age in years (men) |
| Age (years) | Age in years - 10 (women) |
| Nursing home resident | +10 |

#### Comorbidities

| Comorbidity | Points | SNOMED Code |
|-------------|--------|-------------|
| Neoplastic disease | +30 | 363346000 (active within 1 year, excluding basal/squamous skin cancer) |
| Liver disease | +20 | 235856003 (cirrhosis, chronic active hepatitis) |
| Congestive heart failure | +10 | 42343007 |
| Cerebrovascular disease | +10 | 230690007 (stroke, TIA) |
| Renal disease | +10 | 709044004 (chronic kidney disease, Cr >1.5 historically, on dialysis) |

#### Physical Exam Findings

| Finding | Points | Source |
|---------|--------|--------|
| Altered mental status | +20 | Disorientation, stupor, coma |
| Respiratory rate >= 30 | +20 | LOINC 9279-1 |
| Systolic BP < 90 mmHg | +20 | LOINC 8480-6 |
| Temperature < 35C or >= 40C | +15 | LOINC 8310-5 |
| Heart rate >= 125 | +10 | LOINC 8867-4 |

#### Laboratory/Radiology Findings

| Finding | Points | Source |
|---------|--------|--------|
| Arterial pH < 7.35 | +30 | LOINC 2744-1 |
| BUN >= 30 mg/dL (>= 10.7 mmol/L) | +20 | LOINC 3094-0 |
| Sodium < 130 mEq/L | +20 | LOINC 2951-2 |
| Glucose >= 250 mg/dL (>= 13.9 mmol/L) | +10 | LOINC 2345-7 |
| Hematocrit < 30% | +10 | LOINC 4544-3 |
| PaO2 < 60 mmHg (or SpO2 < 90%) | +10 | LOINC 2703-7 (PaO2), 2708-6 (SpO2) |
| Pleural effusion | +10 | Condition: SNOMED 60046008, or imaging report |

### Risk Class Assignment

| Risk Class | Point Range | 30-Day Mortality | Disposition |
|------------|-------------|-----------------|-------------|
| I | Algorithm (age <50, no comorbidities, normal vitals) | 0.1% | Outpatient |
| II | <= 70 | 0.6% | Outpatient |
| III | 71-90 | 0.9-2.8% | Outpatient or brief observation |
| IV | 91-130 | 8.2-9.3% | Inpatient |
| V | > 130 | 27-31.1% | Inpatient, likely ICU |

### PSI Limitations
- May underestimate severity in young patients without comorbidities (age-weighted)
- Does not capture hypoxemia severity well (only binary <60 mmHg threshold)
- Does not include multilobar disease
- Should be used alongside clinical judgment, not as sole disposition tool

---

## A-DROP Score

Japanese Respiratory Society pneumonia severity score. Validated in Asian populations. Increasingly used internationally. Simpler than PSI with comparable predictive value.

### Scoring Criteria

| Criterion | Points | Threshold | FHIR Source |
|-----------|--------|-----------|-------------|
| **A** - Age | 1 | Male >= 70 years OR Female >= 75 years | Patient.birthDate, Patient.gender |
| **D** - Dehydration | 1 | BUN >= 21 mg/dL (>= 7.5 mmol/L) | Observation: LOINC 3094-0 |
| **R** - Respiration | 1 | SpO2 <= 90% OR PaO2 <= 60 mmHg | Observation: LOINC 2708-6, 2703-7 |
| **O** - Orientation | 1 | Confusion / disorientation (new onset) | Condition: SNOMED 130987000, GCS <15 |
| **P** - Pressure | 1 | Systolic BP <= 90 mmHg | Observation: LOINC 8480-6 |

### Risk Stratification

| Score | Severity | Mortality | Disposition |
|-------|----------|-----------|-------------|
| 0 | Mild | 0-2.1% | Outpatient |
| 1-2 | Moderate | 3.2-7.5% | Inpatient (general ward) |
| 3 | Severe | 14.3% | Inpatient, consider ICU |
| 4-5 | Extremely severe | 28.6-42.9% | ICU admission |

### A-DROP vs CURB-65 Comparison

| Feature | A-DROP | CURB-65 |
|---------|--------|---------|
| Age threshold | M>=70, F>=75 (sex-specific) | >=65 (same for both) |
| BUN threshold | >=21 mg/dL | >20 mg/dL |
| Oxygenation | SpO2 <=90% or PaO2 <=60 | Not included |
| BP threshold | SBP <=90 | SBP <90 OR DBP <=60 |
| Advantage | Includes oxygenation | Includes DBP |

---

## ATS/IDSA Severe CAP Criteria (2007/2019)

Used to identify patients requiring ICU admission regardless of CURB-65 or PSI score.

### Major Criteria (Either one = ICU admission)

| Criterion | Definition |
|-----------|-----------|
| Septic shock requiring vasopressors | Fluid-refractory hypotension requiring vasopressor support |
| Respiratory failure requiring mechanical ventilation | Intubation and mechanical ventilation for respiratory failure |

### Minor Criteria (>= 3 = ICU admission recommended)

| Criterion | Threshold | FHIR Source |
|-----------|-----------|-------------|
| Respiratory rate >= 30 | >= 30 breaths/min | LOINC 9279-1 |
| PaO2/FiO2 ratio <= 250 | PaO2/FiO2 <= 250 | LOINC 2703-7 / 3150-0 |
| Multilobar infiltrates | >= 2 lobes on imaging | DiagnosticReport or Condition |
| Confusion/disorientation | New onset | SNOMED 130987000 |
| BUN >= 20 mg/dL | >= 20 mg/dL | LOINC 3094-0 |
| WBC < 4,000 cells/mm3 | < 4,000 | LOINC 6690-2 |
| Platelets < 100,000 cells/mm3 | < 100,000 | LOINC 777-3 |
| Core temperature < 36C | < 36C (hypothermia) | LOINC 8310-5 |
| Hypotension requiring aggressive fluid resuscitation | SBP <90 requiring >2L crystalloid | Clinical assessment |

### Interpretation

| Major Criteria | Minor Criteria | Recommendation |
|---------------|---------------|----------------|
| >= 1 | Any | ICU admission (mandatory) |
| 0 | >= 3 | ICU admission (strongly recommended) |
| 0 | 1-2 | General ward, close monitoring, reassess q4-6h |
| 0 | 0 | General ward or observation based on CURB-65/PSI |

---

## Score Concordance and Clinical Decision-Making

When scores disagree on disposition, use the following hierarchy:

1. **ATS/IDSA severe criteria override all other scores.** If major criteria met or >=3 minor criteria: ICU regardless of CURB-65 or PSI.
2. **PSI tends to overtriage elderly** (age-weighted). A healthy 75-year-old with mild pneumonia may score PSI Class IV solely due to age. Use CURB-65 as counter-check.
3. **CURB-65 may undertriage young patients** with severe physiologic derangement. Use PSI and ATS/IDSA criteria to catch these.
4. **A-DROP adds oxygenation assessment** missing from CURB-65. Consider A-DROP when SpO2/PaO2 is borderline.
5. **Clinical judgment always supersedes scores.** Social factors (homelessness, inability to take oral medications, unreliable follow-up) may warrant admission even with low scores.
