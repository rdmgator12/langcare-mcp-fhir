# Diabetes Complications Screening

Screening intervals, target values, referral criteria, and FHIR search parameters for diabetic complications per ADA Standards of Care 2024.

## Screening Overview

| Complication | Screening Test | Type 1 DM Start | Type 2 DM Start | Interval | FHIR Resource Type |
|-------------|---------------|-----------------|-----------------|----------|-------------------|
| Retinopathy | Dilated eye exam or retinal photography | 5 years after diagnosis | At diagnosis | Annual (biennial if no retinopathy x2) | Procedure |
| Nephropathy | UACR + eGFR | 5 years after diagnosis | At diagnosis | Annual | Observation |
| Neuropathy | Monofilament + vibration sense | 5 years after diagnosis | At diagnosis | Annual | Procedure |
| Cardiovascular | Lipid panel, BP, ASCVD risk | At diagnosis | At diagnosis | Annual (lipids); every visit (BP) | Observation |
| Foot exam | Comprehensive foot examination | 5 years after diagnosis | At diagnosis | Annual (more often if neuropathy/PAD) | Procedure |

## Diabetic Retinopathy

### Screening Parameters

**FHIR Search:**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|274795007,http://snomed.info/sct|252779009,http://snomed.info/sct|410451008&_sort=-date&_count=3"
```

SNOMED codes:
- 274795007 = Examination of retina
- 252779009 = Dilated fundus examination
- 410451008 = Diabetic retinal screening

### Classification (International Clinical Classification)

| Stage | Findings | Management |
|-------|----------|------------|
| No apparent retinopathy | No abnormalities | Routine screening per interval |
| Mild NPDR | Microaneurysms only | Annual eye exam. Optimize glycemic control. |
| Moderate NPDR | More than microaneurysms but less than severe | Eye exam every 6-9 months. Optimize glycemia and BP. |
| Severe NPDR | Any of: > 20 intraretinal hemorrhages in each quadrant, venous beading in 2+ quadrants, IRMA in 1+ quadrant | Referral to retinal specialist. Consider anti-VEGF or PRP. |
| Proliferative DR (PDR) | Neovascularization, vitreous/preretinal hemorrhage | Urgent retinal specialist referral. PRP laser and/or anti-VEGF injections (ranibizumab, aflibercept, bevacizumab). |
| Diabetic macular edema (DME) | Retinal thickening or hard exudates at/near macula | Anti-VEGF injections (first-line). Focal laser if not responding. |

### Referral Criteria (immediate ophthalmology referral)
- Any proliferative changes
- Suspected diabetic macular edema (vision changes, macular thickening)
- Severe NPDR
- Sudden vision loss
- New floaters or flashes (retinal detachment concern)

### Screening Interval Adjustment

| Finding | Next Exam |
|---------|-----------|
| No retinopathy, well-controlled A1c, two consecutive normal exams | May extend to every 2 years |
| No retinopathy but poorly controlled A1c | Annual |
| Mild NPDR | 9-12 months |
| Moderate NPDR | 6-9 months |
| Severe NPDR or PDR | Per retinal specialist (typically 2-4 months) |
| Pregnancy (preexisting diabetes) | Each trimester + 1 year postpartum |

## Diabetic Nephropathy

### Screening Parameters

**FHIR Searches:**
```
Observation: code=48642-3 (UACR) -- annual
Observation: code=33914-3 or 82810-3 (eGFR) -- annual
Observation: code=2160-0 (Creatinine) -- with every BMP
```

### Classification (KDIGO)

**Albuminuria Categories:**

| Category | UACR (mg/g) | Terminology | Significance |
|----------|-------------|-------------|-------------|
| A1 | < 30 | Normal to mildly increased | No diabetic nephropathy |
| A2 | 30-300 | Moderately increased | Early diabetic nephropathy; intervention can slow/reverse |
| A3 | > 300 | Severely increased | Overt diabetic nephropathy; progressive CKD risk |

**Confirmation:** UACR should be confirmed on 2 of 3 specimens collected within 3-6 months (transient albuminuria common with exercise, UTI, fever, CHF, hyperglycemia).

### Target Values

| Parameter | Target | Intervention |
|-----------|--------|--------------|
| UACR | < 30 mg/g (prevent progression) | ACEi or ARB if UACR 30-300; maximize dose |
| eGFR | Slow decline to < 2-3 mL/min/year | ACEi/ARB, SGLT2i, glycemic control, BP control |
| Blood pressure | < 130/80 mmHg | ACEi/ARB first-line if albuminuria present |
| HbA1c | Individualized target | Avoid hypoglycemia (renal clearance of insulin reduced in CKD) |

### Medication Considerations in Diabetic Nephropathy

| eGFR Range | Metformin | SGLT2i | GLP-1 RA | Insulin |
|-----------|-----------|--------|----------|---------|
| > 45 | Full dose | Full dose | No adjustment | Reduce dose as eGFR declines (reduced insulin clearance) |
| 30-45 | Reduce to 1000mg/day max | May initiate (cardiorenal benefit); reduced glycemic effect | No adjustment | Further dose reduction |
| < 30 | Contraindicated | Do not initiate (may continue if already on per KDIGO) | No adjustment | Primary agent; dose reduce significantly |
| Dialysis | Contraindicated | Discontinue | Limited data; caution | Mainstay of therapy |

### Referral Criteria (Nephrology)

- eGFR < 30 (CKD Stage 4-5)
- Persistent UACR > 300 despite ACEi/ARB optimization
- eGFR decline > 5 mL/min/year (rapid progression)
- Refractory hyperkalemia
- Unexplained anemia (Hgb < 10 with CKD)
- Refractory hypertension (4+ agents)
- Suspected non-diabetic kidney disease (active urine sediment, rapid GFR decline, nephrotic-range proteinuria without retinopathy)

## Diabetic Neuropathy

### Screening Parameters

**FHIR Search:**
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|165242009,http://snomed.info/sct|275902005,http://snomed.info/sct|60027008&_sort=-date&_count=3"
```

SNOMED codes:
- 165242009 = Peripheral nerve function test
- 275902005 = Examination of foot
- 60027008 = Monofilament sensory test

### Classification

| Type | Symptoms | Screening Test | Key Findings |
|------|----------|---------------|--------------|
| Distal symmetric polyneuropathy (DSPN) | Numbness, tingling, burning pain (stocking-glove distribution) | 10g monofilament, vibration (128 Hz tuning fork), ankle reflexes, pinprick | Loss of monofilament sensation = loss of protective sensation |
| Autonomic neuropathy | Orthostatic hypotension, resting tachycardia, gastroparesis, erectile dysfunction, neurogenic bladder | Orthostatic BP, resting HR, GI motility assessment | HR variability < 15 bpm with deep breathing; orthostatic drop > 20/10 mmHg |
| Focal neuropathies | Cranial nerve palsies (CN III, IV, VI, VII), mononeuropathies, entrapment (carpal tunnel) | Clinical exam, NCS/EMG if needed | Acute onset, usually self-limited |

### Comprehensive Foot Examination Components

1. **Inspection**: skin integrity, calluses, ulcers, deformities (Charcot, hammer toes, bunions), nail condition
2. **Vascular**: pedal pulses (dorsalis pedis, posterior tibial), capillary refill, ABI if pulses diminished
3. **Neurologic**: 10g monofilament (test 4 sites per foot: great toe plantar, 1st/3rd/5th metatarsal heads), vibration (128 Hz tuning fork at great toe), ankle reflexes
4. **Musculoskeletal**: range of motion, deformity assessment

### Risk Stratification for Ulceration

| Risk Category | Findings | Foot Exam Interval | Footwear |
|--------------|----------|-------------------|----------|
| 0 (low) | No LOPS, no PAD, no deformity | Annual | Standard |
| 1 (moderate) | LOPS OR PAD OR deformity | Every 3-6 months | Therapeutic shoes if deformity |
| 2 (high) | LOPS + PAD, OR LOPS + deformity, OR PAD + deformity | Every 1-3 months | Therapeutic shoes, custom insoles |
| 3 (very high) | History of ulcer or amputation | Every 1-3 months | Therapeutic shoes, custom insoles, consider podiatry referral |

LOPS = Loss of Protective Sensation; PAD = Peripheral Arterial Disease

### Referral Criteria

- Loss of protective sensation: podiatry referral for preventive foot care
- Active foot ulcer: wound care / podiatry / vascular surgery (urgent)
- Suspected Charcot arthropathy: orthopedic/podiatric urgent referral (non-weight-bearing, immobilization)
- PAD with ABI < 0.9: vascular surgery referral
- Intractable neuropathic pain: pain management or neurology

## Cardiovascular Risk in Diabetes

### Screening Parameters

**FHIR Searches:**
```
Observation: code=18262-6 (LDL) -- annual lipid panel
Observation: code=85354-9 (BP panel) -- every visit
Observation: code=2093-3 (Total cholesterol) -- annual
Observation: code=2571-8 (Triglycerides) -- annual
```

### Statin Therapy (ADA + ACC/AHA)

| Age | ASCVD Risk | Statin Intensity | Examples |
|-----|-----------|-----------------|---------|
| < 40, no ASCVD risk factors | Low | None or moderate | -- |
| < 40, with ASCVD risk factors | Moderate-High | Moderate or high-intensity | Atorvastatin 20-40mg, rosuvastatin 10-20mg |
| 40-75, no ASCVD | Moderate | Moderate-intensity | Atorvastatin 10-20mg, rosuvastatin 5-10mg, simvastatin 20-40mg |
| 40-75, with ASCVD risk factors (LDL >= 100, HTN, smoking, CKD, albuminuria, family hx) | High | High-intensity | Atorvastatin 40-80mg, rosuvastatin 20-40mg |
| Any age, established ASCVD | Very high | High-intensity + consider ezetimibe or PCSK9i if LDL > 70 | Atorvastatin 80mg +/- ezetimibe 10mg +/- evolocumab/alirocumab |
| > 75 | Variable | Continue if already on; initiate based on risk discussion | -- |

### Blood Pressure Targets

| Population | Target | First-Line Agent |
|-----------|--------|-----------------|
| Diabetes without albuminuria | < 130/80 mmHg | ACEi, ARB, CCB, or thiazide |
| Diabetes with albuminuria (UACR >= 30) | < 130/80 mmHg | ACEi or ARB (preferred for renal protection) |
| Diabetes with CKD | < 130/80 mmHg | ACEi or ARB |
| Elderly (> 65) with diabetes | < 130/80 (if tolerated) | Individualize; avoid orthostatic hypotension |

### Aspirin

| Scenario | Recommendation |
|----------|---------------|
| Secondary prevention (established ASCVD) | Aspirin 75-162mg daily (clopidogrel if aspirin allergy) |
| Primary prevention, 10-year ASCVD risk >= 10% | May consider aspirin 75-162mg after risk/benefit discussion (not routinely recommended for all diabetics) |
| Primary prevention, 10-year ASCVD risk < 10% | Not recommended |

## Annual Diabetes Care Checklist

| Assessment | Frequency | FHIR Resource | LOINC/SNOMED |
|-----------|-----------|---------------|-------------|
| HbA1c | Every 3-6 months (q3m if not at goal; q6m if stable) | Observation | 4548-4 |
| Lipid panel | Annual (more often if not at goal) | Observation | 57698-3 (panel) |
| UACR | Annual | Observation | 48642-3 |
| eGFR/Creatinine | Annual (more often if CKD) | Observation | 33914-3, 2160-0 |
| Dilated eye exam | Annual (biennial if stable) | Procedure | SNOMED 274795007 |
| Comprehensive foot exam | Annual (more often if high risk) | Procedure | SNOMED 275902005 |
| Blood pressure | Every visit | Observation | 85354-9 |
| Weight / BMI | Every visit | Observation | 29463-7, 39156-5 |
| Smoking status | Every visit | Observation | LOINC 72166-2 |
| Depression screening | Annual | Observation | PHQ-9 |
| Dental exam | Annual | Procedure | Periodontal disease more common in DM |
| Influenza vaccine | Annual | Immunization | -- |
| Pneumococcal vaccine | Per guidelines (PPSV23 + PCV20) | Immunization | -- |
| Hepatitis B vaccine | If unvaccinated, age 19-59 (or > 60 with risk factors) | Immunization | -- |
