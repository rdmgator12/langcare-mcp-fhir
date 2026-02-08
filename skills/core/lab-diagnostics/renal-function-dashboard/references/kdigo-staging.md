# KDIGO CKD Staging and Prognosis

Based on KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease.

## CKD Definition

CKD is defined as abnormalities of kidney structure or function, present for > 3 months, with implications for health.

Criteria (one or more for > 3 months):
- GFR < 60 mL/min/1.73m2
- Albuminuria (UACR >= 30 mg/g)
- Urine sediment abnormalities
- Electrolyte or other abnormalities due to tubular disorders
- Abnormalities detected by histology
- Structural abnormalities detected by imaging
- History of kidney transplantation

## GFR Categories

| Category | GFR Range (mL/min/1.73m2) | Description |
|----------|---------------------------|-------------|
| G1 | >= 90 | Normal or high |
| G2 | 60-89 | Mildly decreased |
| G3a | 45-59 | Mildly to moderately decreased |
| G3b | 30-44 | Moderately to severely decreased |
| G4 | 15-29 | Severely decreased |
| G5 | < 15 | Kidney failure |

Note: G1 and G2 require additional evidence of kidney damage (albuminuria, structural abnormality, etc.) to qualify as CKD. GFR < 60 alone = CKD.

## Albuminuria Categories

| Category | UACR (mg/g) | AER (mg/24h) | Description |
|----------|-------------|-------------|-------------|
| A1 | < 30 | < 30 | Normal to mildly increased |
| A2 | 30-300 | 30-300 | Moderately increased |
| A3 | > 300 | > 300 | Severely increased |

Subdivisions of A3:
- A3a: 300-2000 mg/g (nephrotic range begins ~3500 mg/g)
- A3b: > 2000 mg/g (approaching nephrotic)

## KDIGO Risk Heat Map (Prognosis of CKD)

Risk of adverse outcomes (all-cause mortality, cardiovascular events, kidney failure, AKI):

```
                  Albuminuria Categories
                  A1 (<30)    A2 (30-300)   A3 (>300)
GFR Categories
G1 (>=90)         LOW         MODERATE      HIGH
G2 (60-89)        LOW         MODERATE      HIGH
G3a (45-59)       MODERATE    HIGH          VERY HIGH
G3b (30-44)       HIGH        VERY HIGH     VERY HIGH
G4 (15-29)        VERY HIGH   VERY HIGH     VERY HIGH
G5 (<15)          VERY HIGH   VERY HIGH     VERY HIGH
```

Color coding (for display purposes):
- LOW = Green -- annual monitoring
- MODERATE = Yellow -- annual monitoring, assess progression
- HIGH = Orange -- monitoring every 6 months, nephrology referral if appropriate
- VERY HIGH = Red -- monitoring every 3-4 months, nephrology referral

## Monitoring Frequency by Risk

| Risk Level | GFR Monitoring | UACR Monitoring | Notes |
|-----------|---------------|-----------------|-------|
| Low (green) | Annual | Annual | -- |
| Moderate (yellow) | Annual | Annual | Assess for progression annually |
| High (orange) | Every 6 months | Every 6 months | Consider nephrology referral |
| Very high (red) | Every 3-4 months | Every 3-4 months | Nephrology referral recommended |

## eGFR Equations

### CKD-EPI 2021 (Recommended -- without race variable)

eGFR = 142 * min(Scr/kappa, 1)^alpha * max(Scr/kappa, 1)^(-1.200) * 0.9938^age * (1.012 if female)

Where:
- Scr = serum creatinine (mg/dL)
- kappa = 0.7 (female), 0.9 (male)
- alpha = -0.241 (female), -0.302 (male)
- min = minimum of Scr/kappa or 1
- max = maximum of Scr/kappa or 1

### CKD-EPI Cystatin C (2021)

eGFR = 135 * min(Scys/0.8, 1)^(-0.544) * max(Scys/0.8, 1)^(-0.323) * 0.9961^age * (0.963 if female)

Where Scys = serum cystatin C (mg/L)

### CKD-EPI Creatinine-Cystatin C (2021)

Combined equation using both biomarkers. More accurate than either alone. Recommended when confirmation of GFR category is needed.

### When to Use Cystatin C

- Confirm GFR category when creatinine-based eGFR is near a decision threshold (e.g., eGFR 55 -- confirm G3a vs G2)
- When creatinine may be unreliable: extremes of muscle mass (bodybuilders, amputees, sarcopenia), high-protein diet, creatine supplements
- Discordance between creatinine-based eGFR and clinical picture

## Rate of GFR Decline

### Normal Aging
- After age 40: approximately 0.7-1.0 mL/min/1.73m2 per year
- This is physiologic and does not necessarily indicate CKD

### Pathologic Decline
- Moderate decline: 2-5 mL/min/1.73m2 per year
- Rapid decline: > 5 mL/min/1.73m2 per year (KDIGO definition)
- Very rapid decline: > 10 mL/min/1.73m2 per year

### Calculating Rate of Decline

Use linear regression of eGFR values over time:
1. Collect at least 3-4 eGFR values over >= 12 months
2. Calculate slope: (last eGFR - first eGFR) / time in years
3. Negative slope = declining; magnitude = rate per year

Alternatively, use the KFRE (Kidney Failure Risk Equation) for 2-year and 5-year risk of kidney failure requiring dialysis:
- 4-variable model: age, sex, eGFR, UACR
- 8-variable model adds: calcium, phosphate, bicarbonate, albumin
- Available at kidneyfailurerisk.com

## AKI Superimposed on CKD

### KDIGO AKI Staging

| Stage | Creatinine Criteria | Urine Output Criteria |
|-------|--------------------|--------------------- |
| 1 | Increase >= 0.3 mg/dL within 48h OR 1.5-1.9x baseline within 7 days | < 0.5 mL/kg/h for 6-12 hours |
| 2 | 2.0-2.9x baseline | < 0.5 mL/kg/h for >= 12 hours |
| 3 | >= 3.0x baseline OR creatinine >= 4.0 mg/dL OR initiation of RRT | < 0.3 mL/kg/h for >= 24h OR anuria for >= 12h |

### Baseline Creatinine Determination
- Best: lowest stable creatinine in past 3-6 months
- If no baseline available: estimate from CKD-EPI equation assuming eGFR = 75 mL/min/1.73m2 (ADQI recommendation for unknown baseline)

### AKI on CKD Red Flags
- Creatinine rise > 0.3 mg/dL over any 48-hour window
- eGFR drop > 25% from recent stable baseline
- New oliguria (< 400 mL/24h)
- New electrolyte derangement (hyperkalemia, metabolic acidosis)
- New medication known to cause AKI (NSAIDs, aminoglycosides, contrast, ACEi/ARB initiation with > 30% Cr rise)

## CKD Management by Stage

### General Measures (All Stages)

| Intervention | Target/Action | Evidence |
|-------------|--------------|---------|
| Blood pressure | < 120 mmHg systolic (KDIGO 2021; if tolerated) | SPRINT trial extrapolation |
| ACEi or ARB | Start/maximize if UACR >= 30 mg/g OR proteinuria | Nephroprotection, reduce proteinuria |
| SGLT2 inhibitor | Start if eGFR >= 20 (dapagliflozin, empagliflozin) | DAPA-CKD, EMPA-KIDNEY: reduced CKD progression and CV events |
| Finerenone | Consider if Type 2 DM + CKD with UACR >= 30 (on max ACEi/ARB, K+ < 5.0) | FIDELIO-DKD, FIGARO-DKD |
| Glycemic control | Individualized A1c target; avoid hypoglycemia | ADA/KDIGO consensus |
| Dietary sodium | < 2g/day | Reduces proteinuria, BP |
| Dietary protein | 0.8 g/kg/day (moderate restriction) for G3-G5 | May slow progression; avoid < 0.6 g/kg/day (malnutrition risk) |
| Smoking cessation | Absolute | Smoking accelerates CKD progression |
| Avoid nephrotoxins | NSAIDs, aminoglycosides, iodinated contrast (use with caution, pre-hydrate) | Prevent AKI episodes |

### Stage-Specific Management

| CKD Stage | Additional Considerations |
|-----------|--------------------------|
| G1-G2 | Focus on underlying cause (DM, HTN). Monitor annually if A1 albuminuria. |
| G3a | Start checking: calcium, phosphorus, PTH, vitamin D, CBC (q6-12 months). Medication dose review. |
| G3b | All G3a measures. Consider nephrology referral if progressing. Avoid metformin > 1000mg/day. Hepatitis B vaccination. |
| G4 | Nephrology referral required. Discuss renal replacement options (dialysis modality education, transplant evaluation). Monitor q3 months. Erythropoiesis-stimulating agents if Hgb < 10 (target 10-11.5). |
| G5 (non-dialysis) | Active dialysis planning when eGFR < 15-20. AV fistula creation 6+ months before anticipated dialysis start. Transplant listing if eligible. Dietary K+ and phos restriction. |
| G5D (dialysis) | Dialysis adequacy monitoring (Kt/V). Vascular access care. Monthly labs (CBC, CMP, Ca, Phos, PTH q3m). |

## Dialysis Initiation Criteria

Absolute indications for dialysis (regardless of eGFR):
- Refractory hyperkalemia (K+ > 6.5 despite medical management)
- Refractory fluid overload (pulmonary edema not responding to diuretics)
- Uremic pericarditis
- Uremic encephalopathy
- Severe metabolic acidosis (pH < 7.1 refractory to bicarbonate)
- Refractory nausea/vomiting from uremia
- GFR < 5-7 with symptoms

Relative indication:
- eGFR 5-10 with progressive symptoms (fatigue, anorexia, nausea, pruritus, cognitive decline)
- Note: IDEAL trial showed no benefit to initiating dialysis at eGFR 10-14 vs 5-7 based on symptoms

## LOINC Codes for Renal Function Monitoring

| Lab Test | LOINC Code | Display Name |
|----------|-----------|--------------|
| Serum creatinine | 2160-0 | Creatinine [Mass/volume] in Serum or Plasma |
| BUN | 3094-0 | Urea nitrogen [Mass/volume] in Serum or Plasma |
| eGFR (CKD-EPI creatinine) | 33914-3 | GFR/1.73 sq M.predicted [Volume Rate/Area] in Serum, Plasma or Blood |
| eGFR (CKD-EPI 2021) | 82810-3 | GFR/1.73 sq M.predicted by CKD-EPI 2021 Creatinine |
| eGFR (CKD-EPI cystatin C) | 76484-0 | GFR/1.73 sq M.predicted by Cystatin C |
| Cystatin C | 33863-2 | Cystatin C [Mass/volume] in Serum or Plasma |
| UACR | 48642-3 | Urine albumin/creatinine [Mass Ratio] |
| Urine microalbumin | 14959-1 | Microalbumin [Mass/volume] in Urine |
| Urine creatinine | 2161-8 | Creatinine [Mass/volume] in Urine |
| Urine protein | 2888-6 | Protein [Mass/volume] in Urine |
| Potassium | 2823-3 | Potassium [Moles/volume] in Serum or Plasma |
| Calcium (total) | 17861-6 | Calcium [Mass/volume] in Serum or Plasma |
| Phosphorus | 2777-1 | Phosphate [Mass/volume] in Serum or Plasma |
| Magnesium | 19123-9 | Magnesium [Mass/volume] in Serum or Plasma |
| Bicarbonate | 1963-8 | Bicarbonate [Moles/volume] in Serum or Plasma |
| PTH (intact) | 2731-8 | Parathyrin.intact [Mass/volume] in Serum or Plasma |
| Vitamin D 25-OH | 1989-3 | Calcifediol [Mass/volume] in Serum or Plasma |
| Hemoglobin | 718-7 | Hemoglobin [Mass/volume] in Blood |
| Albumin | 1751-7 | Albumin [Mass/volume] in Serum or Plasma |
| Iron | 2498-4 | Iron [Mass/volume] in Serum or Plasma |
| Ferritin | 2276-4 | Ferritin [Mass/volume] in Serum or Plasma |
| TSAT | 2502-3 | Iron saturation [Mass Fraction] in Serum or Plasma |
