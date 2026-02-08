# Chronic Disease Registry Inclusion Criteria

Registry inclusion codes, key metrics, severity stratification, and complication tracking per disease.

---

## Diabetes Registry

### Inclusion Criteria

**Primary Codes:**

| System | Code | Display |
|--------|------|---------|
| SNOMED | 44054006 | Type 2 diabetes mellitus |
| SNOMED | 46635009 | Type 1 diabetes mellitus |
| SNOMED | 73211009 | Diabetes mellitus |
| SNOMED | 237599002 | Insulin-treated type 2 diabetes |
| SNOMED | 11530004 | Brittle diabetes mellitus |
| ICD-10 | E10 | Type 1 diabetes mellitus |
| ICD-10 | E11 | Type 2 diabetes mellitus |
| ICD-10 | E13 | Other specified diabetes mellitus |

**Exclusion from Registry:**
- Gestational diabetes only (SNOMED 11687002, ICD-10 O24) -- exclude unless persistent postpartum
- Prediabetes (SNOMED 714628002, ICD-10 R73.03) -- separate tracking, not full registry
- Steroid-induced diabetes (SNOMED 190416008) -- include only if persistent after steroid discontinuation

### Severity Stratification (A1c-based)

| Tier | A1c Range | Label | Action Level |
|------|-----------|-------|-------------|
| 1 | < 7.0% | Well controlled | Maintenance |
| 2 | 7.0 - 7.9% | Above target | Optimize therapy |
| 3 | 8.0 - 8.9% | Moderately uncontrolled | Intensify therapy |
| 4 | 9.0 - 11.9% | Poorly controlled | Urgent intervention |
| 5 | >= 12.0% | Severely uncontrolled | Immediate intervention + specialist referral |
| U | No A1c in 12 months | Unknown | Order labs urgently |

### Complication Codes

| Complication | SNOMED | ICD-10 | Key Metric |
|--------------|--------|--------|------------|
| Diabetic retinopathy | 422034002 | E11.3x, H36.0 | Dilated eye exam date |
| Diabetic nephropathy | 127013003 | E11.2x, N08.3 | eGFR, UACR |
| Diabetic neuropathy (peripheral) | 421986006 | E11.4x, G63.2 | Monofilament exam date |
| Diabetic neuropathy (autonomic) | 230572002 | E11.43 | Symptom assessment |
| Diabetic foot ulcer | 280137006 | E11.621 | Wound assessment |
| Diabetic ketoacidosis | 420422005 | E11.10, E11.11 | ED/admission count |
| Hypoglycemia | 302866003 | E11.649 | ED/admission count |
| Peripheral vascular disease | 399957001 | E11.51, I73.9 | ABI result |
| Cardiovascular disease | 56265001 | I25.x, I63.x | Cardiac risk score |

### Key Registry Metrics

| Metric | LOINC | Frequency | Target |
|--------|-------|-----------|--------|
| HbA1c | 4548-4 | Every 3-6 months | < 7.0% (general), < 8.0% (elderly/complex) |
| Fasting glucose | 1558-6 | Per visit PRN | 80-130 mg/dL |
| LDL cholesterol | 13457-7 | Annually | < 100 mg/dL (< 70 if ASCVD) |
| Triglycerides | 2571-8 | Annually | < 150 mg/dL |
| eGFR | 33914-3 | Annually | >= 60 mL/min |
| Urine albumin/creatinine | 14959-1 | Annually | < 30 mg/g |
| Blood pressure (panel) | 85354-9 | Every visit | < 130/80 |
| BMI | 39156-5 | Every visit | < 30 kg/m2 |

---

## Hypertension Registry

### Inclusion Criteria

| System | Code | Display |
|--------|------|---------|
| SNOMED | 38341003 | Hypertensive disorder |
| SNOMED | 59621000 | Essential hypertension |
| SNOMED | 1201005 | Benign essential hypertension |
| SNOMED | 78975002 | Malignant essential hypertension |
| ICD-10 | I10 | Essential (primary) hypertension |
| ICD-10 | I11 | Hypertensive heart disease |
| ICD-10 | I12 | Hypertensive chronic kidney disease |
| ICD-10 | I13 | Hypertensive heart and CKD |

### Severity Stratification (AHA 2017 Guidelines)

| Stage | Systolic | Diastolic | Label |
|-------|----------|-----------|-------|
| Normal | < 120 | < 80 | Well controlled on therapy |
| Elevated | 120-129 | < 80 | Elevated, monitor |
| Stage 1 | 130-139 | 80-89 | Above target |
| Stage 2 | >= 140 | >= 90 | Uncontrolled |
| Crisis | > 180 | > 120 | Hypertensive crisis |

### Complication Codes

| Complication | SNOMED | ICD-10 |
|--------------|--------|--------|
| Hypertensive heart disease | 64715009 | I11.x |
| Hypertensive CKD | 194774006 | I12.x |
| Hypertensive retinopathy | 36225005 | H35.03x |
| Hypertensive emergency | 706882009 | I16.1 |
| Stroke (CVA) | 230690007 | I63.x, I61.x |
| Myocardial infarction | 22298006 | I21.x |
| Aortic aneurysm | 233985008 | I71.x |

---

## COPD Registry

### Inclusion Criteria

| System | Code | Display |
|--------|------|---------|
| SNOMED | 13645005 | Chronic obstructive lung disease |
| SNOMED | 185086009 | Chronic obstructive bronchitis |
| SNOMED | 87433001 | Pulmonary emphysema |
| ICD-10 | J44.0 | COPD with acute lower respiratory infection |
| ICD-10 | J44.1 | COPD with acute exacerbation |
| ICD-10 | J44.9 | COPD, unspecified |
| ICD-10 | J43 | Emphysema |

### Severity Stratification (GOLD 2024)

| GOLD Stage | FEV1 % Predicted | Label |
|------------|-----------------|-------|
| GOLD 1 | >= 80% | Mild |
| GOLD 2 | 50-79% | Moderate |
| GOLD 3 | 30-49% | Severe |
| GOLD 4 | < 30% | Very Severe |

**LOINC for FEV1 % predicted:** 20150-9
**LOINC for FEV1/FVC ratio:** 19926-5

**ABE Group (symptoms + exacerbation risk):**
- Group A: Low symptoms (mMRC 0-1 or CAT < 10) + 0-1 moderate exacerbations/year
- Group B: High symptoms (mMRC >= 2 or CAT >= 10) + 0-1 moderate exacerbations/year
- Group E: Any symptoms + >= 2 moderate exacerbations OR >= 1 hospitalization/year

### Complication/Event Codes

| Event | SNOMED | ICD-10 | Tracking |
|-------|--------|--------|----------|
| Acute exacerbation | 195951007 | J44.1 | Count per 12 months |
| Pneumonia | 233604007 | J18.x | Count per 12 months |
| Respiratory failure | 65710008 | J96.x | Count per 12 months |
| Cor pulmonale | 49584005 | I27.81 | Presence |

---

## Asthma Registry

### Inclusion Criteria

| System | Code | Display |
|--------|------|---------|
| SNOMED | 195967001 | Asthma |
| SNOMED | 233678006 | Childhood asthma |
| SNOMED | 304527002 | Acute asthma |
| ICD-10 | J45.2x | Mild intermittent asthma |
| ICD-10 | J45.3x | Mild persistent asthma |
| ICD-10 | J45.4x | Moderate persistent asthma |
| ICD-10 | J45.5x | Severe persistent asthma |

### Severity Classification (NAEPP EPR-4)

| Severity | Symptoms | Nighttime | FEV1 % Predicted | Treatment Step |
|----------|----------|-----------|-------------------|----------------|
| Intermittent | <= 2 days/week | <= 2x/month | > 80% | Step 1 |
| Mild persistent | > 2 days/week | 3-4x/month | >= 80% | Step 2 |
| Moderate persistent | Daily | > 1x/week | 60-80% | Step 3-4 |
| Severe persistent | Throughout day | Often nightly | < 60% | Step 5-6 |

---

## CHF Registry

### Inclusion Criteria

| System | Code | Display |
|--------|------|---------|
| SNOMED | 42343007 | Congestive heart failure |
| SNOMED | 84114007 | Heart failure |
| SNOMED | 441481004 | Heart failure with reduced EF |
| SNOMED | 446221000 | Heart failure with preserved EF |
| ICD-10 | I50.1 | Left ventricular failure |
| ICD-10 | I50.2x | Systolic (HFrEF) |
| ICD-10 | I50.3x | Diastolic (HFpEF) |
| ICD-10 | I50.4x | Combined systolic and diastolic |
| ICD-10 | I50.9 | Heart failure, unspecified |

### Classification

**By Ejection Fraction (LOINC 10230-1):**
| Type | LVEF | Label |
|------|------|-------|
| HFrEF | < 40% | Reduced |
| HFmrEF | 40-49% | Mildly reduced |
| HFpEF | >= 50% | Preserved |

**NYHA Functional Class:**
| Class | Limitation | LOINC 88020-3 value |
|-------|-----------|---------------------|
| I | No limitation | None |
| II | Slight limitation | Mild |
| III | Marked limitation | Moderate |
| IV | Severe limitation | Severe |

### Key Metrics

| Metric | LOINC | Significance |
|--------|-------|-------------|
| BNP | 42637-9 | > 100 suggests CHF; > 500 decompensation |
| NT-proBNP | 33762-6 | Age-adjusted: > 300 (< 50y), > 900 (50-75y), > 1800 (> 75y) |
| LVEF | 10230-1 | Classification and GDMT eligibility |
| Weight | 29463-7 | Daily monitoring for fluid status |
| Serum sodium | 2951-2 | < 135 dilutional hyponatremia |

---

## CKD Registry

### Inclusion Criteria

| System | Code | Display |
|--------|------|---------|
| SNOMED | 709044004 | Chronic kidney disease |
| SNOMED | 431855005 | CKD stage 1 |
| SNOMED | 431856006 | CKD stage 2 |
| SNOMED | 433144002 | CKD stage 3 |
| SNOMED | 431857002 | CKD stage 4 |
| SNOMED | 433146000 | CKD stage 5 |
| ICD-10 | N18.1 | CKD stage 1 |
| ICD-10 | N18.2 | CKD stage 2 |
| ICD-10 | N18.30-N18.32 | CKD stage 3 (3a, 3b) |
| ICD-10 | N18.4 | CKD stage 4 |
| ICD-10 | N18.5 | CKD stage 5 |
| ICD-10 | N18.6 | ESRD |

### Staging (KDIGO)

**GFR Categories (LOINC 33914-3):**

| Stage | eGFR (mL/min/1.73m2) | Label |
|-------|----------------------|-------|
| G1 | >= 90 | Normal or high (with kidney damage markers) |
| G2 | 60-89 | Mildly decreased |
| G3a | 45-59 | Mildly to moderately decreased |
| G3b | 30-44 | Moderately to severely decreased |
| G4 | 15-29 | Severely decreased |
| G5 | < 15 | Kidney failure |

**Albuminuria Categories (LOINC 14959-1):**

| Category | UACR (mg/g) | Label |
|----------|-------------|-------|
| A1 | < 30 | Normal to mildly increased |
| A2 | 30-300 | Moderately increased |
| A3 | > 300 | Severely increased |

### Complication Tracking

| Complication | SNOMED | ICD-10 | Trigger |
|--------------|--------|--------|---------|
| Anemia of CKD | 691421000119108 | D63.1 | Hemoglobin < 10 |
| Secondary hyperparathyroidism | 66999008 | N25.81 | PTH > 2x upper normal |
| Metabolic acidosis | 59455009 | E87.2 | Bicarb < 22 |
| Hyperkalemia | 14140009 | E87.5 | K > 5.5 |
| Mineral bone disease | 700457001 | N25.0 | Phosphorus > 4.5, Calcium abnl |
| Volume overload | 276514003 | E87.70 | Clinical assessment |
