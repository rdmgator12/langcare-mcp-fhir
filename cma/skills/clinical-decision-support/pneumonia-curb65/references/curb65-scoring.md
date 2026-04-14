# CURB-65 and PSI/PORT Scoring Reference

## CURB-65 Score

| Criterion | Definition | Points |
|-----------|-----------|--------|
| **C** - Confusion | New mental confusion (AMT <=8 or GCS <15 or new disorientation) | +1 |
| **U** - Urea (BUN) | BUN > 19 mg/dL (>7 mmol/L urea) | +1 |
| **R** - Respiratory rate | >= 30 breaths/min | +1 |
| **B** - Blood pressure | Systolic < 90 mmHg OR Diastolic <= 60 mmHg | +1 |
| **65** - Age | >= 65 years | +1 |

### CURB-65 Interpretation

| Score | 30-Day Mortality | Disposition |
|-------|-----------------|-------------|
| 0 | 0.7% | Outpatient treatment |
| 1 | 2.1% | Outpatient treatment (consider observation if borderline) |
| 2 | 9.2% | Short inpatient stay or hospital-supervised outpatient |
| 3 | 14.5% | Hospitalize (consider ICU) |
| 4 | 40% | Hospitalize, ICU consideration |
| 5 | 57% | Hospitalize, ICU admission |

### CRB-65 (No Lab Variant -- For Office/Triage Use)
Same as CURB-65 but without BUN (max score 4):
- 0: Very low risk, consider outpatient
- 1-2: Increased risk, consider referral to hospital
- 3-4: High risk, urgent hospitalization

## PSI/PORT Score (Pneumonia Severity Index)

### Step 1: Check for Low-Risk Criteria (Class I)
If ALL of the following, patient is Class I (no scoring needed):
- Age < 50
- No comorbidities (neoplastic disease, liver disease, CHF, cerebrovascular disease, renal disease)
- Normal vital signs (HR <125, RR <30, SBP >=90, Temp >=35C and <40C)
- Normal mental status

### Step 2: Point Calculation (if not Class I)

**Demographics:**
| Factor | Points |
|--------|--------|
| Age (male) | Age in years |
| Age (female) | Age - 10 |
| Nursing home resident | +10 |

**Comorbidities:**
| Factor | Points |
|--------|--------|
| Neoplastic disease | +30 |
| Liver disease | +20 |
| Congestive heart failure | +10 |
| Cerebrovascular disease | +10 |
| Renal disease | +10 |

**Physical Exam:**
| Factor | Points |
|--------|--------|
| Altered mental status | +20 |
| Respiratory rate >= 30 | +20 |
| Systolic BP < 90 mmHg | +20 |
| Temperature < 35C or >= 40C | +15 |
| Heart rate >= 125 bpm | +10 |

**Laboratory/Imaging:**
| Factor | Points |
|--------|--------|
| Arterial pH < 7.35 | +30 |
| BUN >= 30 mg/dL | +20 |
| Sodium < 130 mEq/L | +20 |
| Glucose >= 250 mg/dL | +10 |
| Hematocrit < 30% | +10 |
| PaO2 < 60 mmHg or SpO2 < 90% | +10 |
| Pleural effusion | +10 |

### PSI/PORT Risk Classes

| Class | Points | 30-Day Mortality | Disposition |
|-------|--------|-----------------|-------------|
| I | Algorithm (see Step 1) | 0.1% | Outpatient |
| II | <= 70 | 0.6% | Outpatient |
| III | 71-90 | 0.9% | Outpatient or brief observation |
| IV | 91-130 | 9.3% | Inpatient |
| V | > 130 | 27% | Inpatient, consider ICU |

## ATS/IDSA Criteria for Severe CAP (ICU Admission)

### Major Criteria (1 = ICU)
- Septic shock requiring vasopressors
- Respiratory failure requiring mechanical ventilation

### Minor Criteria (>=3 = ICU)
- Respiratory rate >= 30 breaths/min
- PaO2/FiO2 ratio <= 250
- Multilobar infiltrates
- Confusion/disorientation
- BUN >= 20 mg/dL
- WBC < 4,000 cells/uL
- Platelet count < 100,000/uL
- Core temperature < 36C
- Hypotension requiring aggressive fluid resuscitation

## Empiric Antibiotic Therapy (ATS/IDSA 2019)

### Outpatient (No Comorbidities)
- Amoxicillin 1g TID **OR**
- Doxycycline 100mg BID **OR**
- Macrolide (azithromycin 500mg then 250mg daily x4) if local resistance <25%

### Outpatient (With Comorbidities)
- Amoxicillin/clavulanate 875/125mg BID + macrolide **OR**
- Respiratory fluoroquinolone (levofloxacin 750mg daily or moxifloxacin 400mg daily)

### Inpatient (Non-ICU)
- Beta-lactam (ampicillin/sulbactam, ceftriaxone, or cefotaxime) + macrolide **OR**
- Respiratory fluoroquinolone monotherapy

### Inpatient (ICU)
- Beta-lactam (ceftriaxone, ampicillin/sulbactam, or cefotaxime) + macrolide **OR**
- Beta-lactam + respiratory fluoroquinolone
- Add vancomycin or linezolid if MRSA risk factors
- Add antipseudomonal beta-lactam if Pseudomonas risk factors

## LOINC Codes for Pneumonia Assessment

| Parameter | LOINC |
|-----------|-------|
| BUN | 3094-0 |
| Respiratory rate | 9279-1 |
| Systolic BP | 8480-6 |
| Diastolic BP | 8462-4 |
| Heart rate | 8867-4 |
| Temperature | 8310-5 |
| SpO2 | 59408-5 |
| WBC | 6690-2 |
| Platelets | 777-3 |
| Sodium | 2951-2 |
| Glucose | 2345-7 |
| Hematocrit | 4544-3 |
| Arterial pH | 2744-1 |
| PaO2 | 2703-7 |
| GCS | 9269-2 |
