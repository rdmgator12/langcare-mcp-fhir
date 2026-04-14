# Cardiovascular Risk Scoring Reference

## CHA2DS2-VASc Score (Stroke Risk in Atrial Fibrillation)

| Factor | Points |
|--------|--------|
| **C** - Congestive heart failure (or LVEF <=40%) | +1 |
| **H** - Hypertension | +1 |
| **A2** - Age >= 75 | +2 |
| **D** - Diabetes mellitus | +1 |
| **S2** - Stroke/TIA/thromboembolism history | +2 |
| **V** - Vascular disease (prior MI, PAD, aortic plaque) | +1 |
| **A** - Age 65-74 | +1 |
| **Sc** - Sex category (female) | +1 |

**Maximum score: 9**

### Annual Stroke Risk Estimates

| Score | Annual Stroke Risk | Recommendation |
|-------|-------------------|----------------|
| 0 (male) | ~0% | No anticoagulation |
| 1 (male), 1 (female only) | ~1.3% | Consider anticoagulation |
| 2 | ~2.2% | Anticoagulation recommended |
| 3 | ~3.2% | Anticoagulation recommended |
| 4 | ~4.0% | Anticoagulation recommended |
| 5 | ~6.7% | Anticoagulation recommended |
| 6 | ~9.8% | Anticoagulation recommended |
| 7 | ~9.6% | Anticoagulation recommended |
| 8 | ~12.5% | Anticoagulation recommended |
| 9 | ~15.2% | Anticoagulation recommended |

Note: Female sex scores 1 point but is not counted in isolation (CHA2DS2-VASc 1 from sex alone does not trigger anticoagulation).

## HAS-BLED Score (Bleeding Risk on Anticoagulation)

| Factor | Points | Definition |
|--------|--------|------------|
| **H** - Hypertension | +1 | Uncontrolled SBP > 160 mmHg |
| **A** - Abnormal renal/liver function | +1 each (max 2) | Renal: dialysis, transplant, Cr > 2.3. Liver: cirrhosis, bilirubin >2x ULN, AST/ALT/ALP >3x ULN |
| **S** - Stroke history | +1 | Prior stroke |
| **B** - Bleeding history/predisposition | +1 | Prior major bleed, anemia, bleeding diathesis |
| **L** - Labile INR | +1 | TTR < 60% on warfarin |
| **E** - Elderly | +1 | Age > 65 |
| **D** - Drugs or alcohol | +1 each (max 2) | Antiplatelets, NSAIDs; alcohol >= 8 drinks/week |

**Maximum score: 9**

| Score | Bleeding Risk | Action |
|-------|-------------|--------|
| 0-2 | Low-Moderate | Anticoagulation benefits typically outweigh risks |
| >= 3 | High | Caution; address modifiable risk factors; does NOT mean avoid anticoagulation |

## HEART Score (Acute Chest Pain Evaluation)

| Component | 0 Points | 1 Point | 2 Points |
|-----------|----------|---------|----------|
| **H** - History | Slightly suspicious | Moderately suspicious | Highly suspicious |
| **E** - ECG | Normal | Non-specific ST/repol changes | Significant ST deviation |
| **A** - Age | < 45 | 45-64 | >= 65 |
| **R** - Risk factors | No known risk factors | 1-2 risk factors | >= 3 risk factors or history of atherosclerotic disease |
| **T** - Troponin | <= normal limit | 1-3x normal limit | > 3x normal limit |

Risk factors: HTN, hyperlipidemia, DM, obesity (BMI >30), smoking, family hx of CAD, PAD, prior stroke.

| Score | Risk | 6-Week MACE Rate | Action |
|-------|------|-------------------|--------|
| 0-3 | Low | 0.9-1.7% | Discharge with outpatient follow-up |
| 4-6 | Moderate | 12-16.6% | Observation, serial troponins, stress test or CCTA |
| 7-10 | High | 50-65% | Admit, cardiology consult, early invasive strategy |

## ASCVD Pooled Cohort Equations (10-Year ASCVD Risk)

### Eligible Population
- Age 40-79 years
- Without known ASCVD

### Required Inputs
| Parameter | Source |
|-----------|--------|
| Age | Patient.birthDate |
| Sex | Patient.gender |
| Race | Patient extension (us-core-race): White or African American (equations validated only for these groups) |
| Total cholesterol | Observation LOINC 2093-3 |
| HDL cholesterol | Observation LOINC 2085-9 |
| Systolic BP | Observation LOINC 8480-6 |
| On BP treatment | MedicationRequest: antihypertensives active |
| Diabetes | Condition: DM active |
| Current smoker | Observation: smoking status (LOINC 72166-2) |

### Risk Categories and Statin Recommendations (ACC/AHA 2019)

| 10-Year Risk | Category | Statin Recommendation |
|-------------|----------|----------------------|
| < 5% | Low | Lifestyle, no statin |
| 5-7.4% | Borderline | Consider if risk enhancers present |
| 7.5-19.9% | Intermediate | Moderate-intensity statin; if uncertain, obtain CAC score |
| >= 20% | High | High-intensity statin |

### Risk-Enhancing Factors (favor statin in borderline/intermediate)
- Family history of premature ASCVD (male <55, female <65)
- LDL-C >= 160 mg/dL
- Metabolic syndrome
- CKD
- Chronic inflammatory conditions (RA, lupus, psoriasis, HIV)
- South Asian ancestry
- Triglycerides >= 175 mg/dL (persistently)
- ApoB >= 130 mg/dL
- hs-CRP >= 2.0 mg/L
- Lp(a) >= 50 mg/dL
- Ankle-brachial index < 0.9
- Coronary artery calcium (CAC) score >= 100 Agatston units

### Statin Intensity

| Intensity | LDL Reduction | Examples |
|-----------|-------------|----------|
| High | >= 50% | Atorvastatin 40-80mg, Rosuvastatin 20-40mg |
| Moderate | 30-49% | Atorvastatin 10-20mg, Rosuvastatin 5-10mg, Simvastatin 20-40mg, Pravastatin 40-80mg |
| Low | < 30% | Simvastatin 10mg, Pravastatin 10-20mg, Fluvastatin 20-40mg |
