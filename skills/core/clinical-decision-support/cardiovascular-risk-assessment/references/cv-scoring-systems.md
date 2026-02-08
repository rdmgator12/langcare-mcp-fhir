# Cardiovascular Scoring Systems Reference

## CHA2DS2-VASc Score (Stroke Risk in Atrial Fibrillation)

Estimates annual stroke risk in patients with non-valvular atrial fibrillation. Used to guide anticoagulation decisions.

### Scoring Criteria

| Criterion | Points | SNOMED / Notes |
|-----------|--------|----------------|
| **C** - Congestive heart failure (or LVEF <=40%) | 1 | SNOMED 42343007 |
| **H** - Hypertension (or on antihypertensive) | 1 | SNOMED 38341003 |
| **A2** - Age >= 75 years | 2 | From Patient.birthDate |
| **D** - Diabetes mellitus | 1 | SNOMED 73211009 |
| **S2** - Stroke / TIA / thromboembolism history | 2 | SNOMED 230690007 (CVA), 266257000 (TIA) |
| **V** - Vascular disease (prior MI, PAD, aortic plaque) | 1 | SNOMED 22298006 (MI), 400047006 (PVD) |
| **Asc** - Age 65-74 years | 1 | From Patient.birthDate (mutually exclusive with A2) |
| **Sc** - Sex category (female) | 1 | From Patient.gender |

Maximum score: 9

### Risk Stratification and Recommendations

| Score | Annual Stroke Risk | Recommendation |
|-------|-------------------|----------------|
| 0 (male) / 1 (female, sex point only) | ~0.2-0.7% | No anticoagulation recommended |
| 1 (male) | ~0.6-1.3% | Consider OAC (oral anticoagulation) -- shared decision-making |
| >= 2 | 2.2-15.2% | OAC recommended (Class I recommendation) |

Annual stroke risk by score:
- 0: 0.2%
- 1: 0.6%
- 2: 2.2%
- 3: 3.2%
- 4: 4.8%
- 5: 7.2%
- 6: 9.7%
- 7: 11.2%
- 8: 10.8%
- 9: 12.2%

---

## HAS-BLED Score (Bleeding Risk on Anticoagulation)

Estimates annual major bleeding risk in anticoagulated patients. High score does NOT contraindicate anticoagulation -- it identifies modifiable risk factors.

### Scoring Criteria

| Criterion | Points | Definition |
|-----------|--------|------------|
| **H** - Hypertension (uncontrolled) | 1 | SBP > 160 mmHg |
| **A** - Abnormal renal/liver function | 1-2 | Renal: dialysis, transplant, Cr >2.26 mg/dL (1 pt). Liver: cirrhosis, bilirubin >2x ULN + AST/ALT >3x ULN (1 pt) |
| **S** - Stroke history | 1 | Prior ischemic or hemorrhagic stroke |
| **B** - Bleeding history or predisposition | 1 | Prior major bleed, anemia, thrombocytopenia |
| **L** - Labile INR | 1 | TTR (time in therapeutic range) < 60%. Only for warfarin patients. |
| **E** - Elderly (age > 65) | 1 | From Patient.birthDate |
| **D** - Drugs or alcohol | 1-2 | Antiplatelet/NSAID concomitant use (1 pt). Alcohol excess >=8 drinks/week (1 pt) |

Maximum score: 9

### Interpretation

| Score | Annual Major Bleed Risk | Action |
|-------|------------------------|--------|
| 0 | 0.9% | Low risk |
| 1 | 3.4% | Low-moderate risk |
| 2 | 4.1% | Moderate risk -- address modifiable factors |
| >= 3 | 5.8-12.5% | High risk -- address ALL modifiable factors, closer monitoring |

Modifiable factors: uncontrolled HTN, labile INR (switch to DOAC), concomitant NSAIDs/antiplatelets, alcohol use.

---

## HEART Score (Chest Pain Evaluation)

Risk stratification for acute chest pain in the emergency department. Predicts 6-week MACE (major adverse cardiac events).

### Scoring Criteria

| Component | 0 Points | 1 Point | 2 Points |
|-----------|----------|---------|----------|
| **H** - History | Slightly suspicious | Moderately suspicious | Highly suspicious |
| **E** - ECG | Normal | Non-specific repolarization disturbance | Significant ST deviation |
| **A** - Age | < 45 | 45-64 | >= 65 |
| **R** - Risk factors | No known risk factors | 1-2 risk factors | >= 3 risk factors or known atherosclerotic disease |
| **T** - Troponin | <= normal limit | 1-3x normal limit | > 3x normal limit |

Risk factors for R component: HTN, hyperlipidemia, DM, obesity (BMI >=30), smoking (current or quit <3mo), family hx premature CAD (M<55, F<65), known atherosclerotic disease.

### History Scoring Guide

- **0 (slightly suspicious)**: Non-specific symptoms, atypical presentation, pleuritic, positional, reproducible with palpation
- **1 (moderately suspicious)**: Chest pain with some typical features but not classic; moderate clinical concern
- **2 (highly suspicious)**: Pressure/squeezing, substernal, radiation to jaw/arm, diaphoresis, associated with exertion, similar to prior ACS

### Interpretation and Disposition

| Score | 6-Week MACE Risk | Disposition |
|-------|-------------------|-------------|
| 0-3 | 0.9-1.7% | Discharge with outpatient follow-up. Consider accelerated diagnostic protocol. |
| 4-6 | 12-16.6% | Observation, serial troponins, non-invasive testing (stress test or CCTA) |
| 7-10 | 50-65% | Admit, cardiology consult, early invasive strategy |

---

## Framingham Risk Score (10-Year CHD Risk)

Estimates 10-year risk of coronary heart disease events. Age range: 30-74 years.

### Variables

| Variable | Men Points | Women Points |
|----------|-----------|--------------|
| Age 30-34 | 0 | 0 |
| Age 35-39 | 2 | 2 |
| Age 40-44 | 5 | 4 |
| Age 45-49 | 6 | 5 |
| Age 50-54 | 8 | 7 |
| Age 55-59 | 10 | 8 |
| Age 60-64 | 11 | 9 |
| Age 65-69 | 12 | 10 |
| Age 70-74 | 14 | 11 |
| Total cholesterol <160 | 0 | 0 |
| Total cholesterol 160-199 | 1 | 1 |
| Total cholesterol 200-239 | 2 | 3 |
| Total cholesterol 240-279 | 3 | 4 |
| Total cholesterol >=280 | 4 | 5 |
| HDL >=60 | -2 | -2 |
| HDL 50-59 | -1 | -1 |
| HDL 40-49 | 0 | 0 |
| HDL <40 | 1 | 2 |
| SBP <120 untreated | -2 | -3 |
| SBP 120-129 untreated | 0 | 0 |
| SBP 130-139 untreated | 1 | 1 |
| SBP 140-159 untreated | 2 | 2 |
| SBP >=160 untreated | 3 | 3 |
| SBP <120 treated | 0 | -1 |
| SBP 120-129 treated | 2 | 2 |
| SBP 130-139 treated | 3 | 3 |
| SBP 140-159 treated | 4 | 5 |
| SBP >=160 treated | 5 | 6 |
| Smoker | 4 | 3 |
| Non-smoker | 0 | 0 |
| Diabetes | 3 | 4 |
| No diabetes | 0 | 0 |

### Interpretation

| 10-Year Risk | Category |
|-------------|----------|
| < 10% | Low risk |
| 10-20% | Intermediate risk |
| > 20% | High risk (CHD equivalent) |

---

## ASCVD Pooled Cohort Equations (10-Year ASCVD Risk)

ACC/AHA recommended calculator for 10-year atherosclerotic cardiovascular disease risk. Valid for ages 40-79 without prior ASCVD events.

### Required Variables

| Variable | Source | LOINC/SNOMED |
|----------|--------|-------------|
| Age | Patient.birthDate | -- |
| Sex | Patient.gender | -- |
| Race (White or African American) | Patient US Core race extension | -- |
| Total cholesterol (mg/dL) | Observation | 2093-3 |
| HDL cholesterol (mg/dL) | Observation | 2085-9 |
| Systolic BP (mmHg) | Observation | 8480-6 |
| On BP treatment | MedicationRequest (antihypertensive active) | -- |
| Diabetes | Condition | SNOMED 73211009 |
| Current smoker | Observation (smoking status) | 72166-2 |

LOINC 72166-2 = Tobacco smoking status. Values: current every day smoker, current some day smoker, former smoker, never smoker.

### Race-Specific Coefficients

The equations use different coefficients for:
1. White women
2. White men
3. African American women
4. African American men

If race is not White or African American, use White coefficients with notation that equations are not validated for other groups.

### Interpretation and Treatment Thresholds

| 10-Year ASCVD Risk | Category | Statin Recommendation |
|--------------------|----------|----------------------|
| < 5% | Low risk | Lifestyle modifications |
| 5-7.4% | Borderline risk | Consider moderate-intensity statin if risk enhancers present |
| 7.5-19.9% | Intermediate risk | Moderate-intensity statin recommended. If uncertain, consider CAC scoring. |
| >= 20% | High risk | High-intensity statin recommended |

### Risk Enhancers (for borderline/intermediate risk)

- Family history premature ASCVD (men <55, women <65)
- LDL-C 160-189 mg/dL
- Metabolic syndrome
- Chronic kidney disease (eGFR 15-59)
- Chronic inflammatory conditions (RA, psoriasis, HIV)
- Premature menopause (before age 40)
- Preeclampsia history
- South Asian ancestry
- hs-CRP >= 2.0 mg/L
- Lp(a) >= 50 mg/dL or >= 125 nmol/L
- ABI < 0.9
- Elevated apoB >= 130 mg/dL

---

## Score Selection Logic

| Clinical Scenario | Primary Score | Secondary Score |
|-------------------|--------------|-----------------|
| Atrial fibrillation -- anticoagulation decision | CHA2DS2-VASc | HAS-BLED |
| Acute chest pain in ED | HEART Score | -- |
| Primary prevention, age 40-79 | ASCVD PCE | Framingham (comparison) |
| Primary prevention, age <40 or >79 | Framingham | -- |
| Patient on anticoagulation -- bleeding risk | HAS-BLED | -- |
| Multiple scenarios apply | Calculate all applicable | -- |
