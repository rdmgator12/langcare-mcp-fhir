# ICD-10 Coding Specificity Improvements

## Purpose

Identify opportunities to improve ICD-10-CM coding specificity on the active problem list. More specific codes improve data quality, risk adjustment accuracy, quality measure performance, and reimbursement accuracy.

## General Principles

1. **Code to the highest level of specificity** supported by documentation
2. **Laterality matters** -- always specify left, right, or bilateral when applicable
3. **Type and stage** -- specify type 1 vs type 2 diabetes, CKD stage, cancer stage
4. **With vs without complications** -- document and code complications when present
5. **Combination codes** -- use a single code that captures both condition and manifestation when available

## Common Specificity Improvements by System

### Diabetes Mellitus (E08-E13)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| E11.9 | Type 2 DM without complications | E11.65 | With hyperglycemia (poorly controlled) |
| E11.9 | Type 2 DM without complications | E11.22 | With diabetic CKD (requires N18.x secondary) |
| E11.9 | Type 2 DM without complications | E11.40 | With diabetic neuropathy (unspecified type) |
| E11.9 | Type 2 DM without complications | E11.42 | With diabetic polyneuropathy |
| E11.9 | Type 2 DM without complications | E11.311 | With unspecified diabetic retinopathy with macular edema |
| E11.9 | Type 2 DM without complications | E11.319 | With unspecified diabetic retinopathy without macular edema |
| E11.9 | Type 2 DM without complications | E11.51 | With diabetic peripheral angiopathy without gangrene |
| E11.9 | Type 2 DM without complications | E11.621 | With foot ulcer |
| E11.9 | Type 2 DM without complications | E11.69 | With other specified complication |

**Key check:** If patient has both diabetes AND one of: CKD, neuropathy, retinopathy, nephropathy, peripheral vascular disease -- use the combination code instead of coding them separately.

### Hypertension (I10-I16)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| I10 | Essential hypertension | I11.0 | With heart failure (HF documented) |
| I10 | Essential hypertension | I11.9 | With heart disease without HF |
| I10 | Essential hypertension | I12.0 | With CKD stage 5 or ESRD |
| I10 | Essential hypertension | I12.9 | With CKD stage 1-4 |
| I10 | Essential hypertension | I13.0 | With heart disease and CKD |
| I10 | Essential hypertension | I13.10 | With heart failure and CKD |

**Key rule:** ICD-10 assumes a causal relationship between hypertension and CKD, and between hypertension and heart disease, unless documentation explicitly states otherwise.

### Heart Failure (I50)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| I50.9 | Heart failure, unspecified | I50.20 | Systolic (HFrEF), unspecified |
| I50.9 | Heart failure, unspecified | I50.22 | Chronic systolic (HFrEF) |
| I50.9 | Heart failure, unspecified | I50.23 | Acute on chronic systolic |
| I50.9 | Heart failure, unspecified | I50.30 | Diastolic (HFpEF), unspecified |
| I50.9 | Heart failure, unspecified | I50.32 | Chronic diastolic (HFpEF) |
| I50.9 | Heart failure, unspecified | I50.40 | Combined systolic and diastolic |
| I50.9 | Heart failure, unspecified | I50.42 | Chronic combined |
| I50.9 | Heart failure, unspecified | I50.810 | Right heart failure, unspecified |
| I50.9 | Heart failure, unspecified | I50.814 | Right heart failure due to left HF |

**Key check:** Look at echocardiogram results. EF < 40% = systolic (HFrEF). EF >= 50% with diastolic dysfunction = diastolic (HFpEF). EF 40-49% = borderline/mid-range.

### Chronic Kidney Disease (N18)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| N18.9 | CKD, unspecified | N18.1 | Stage 1 (eGFR >= 90 with kidney damage) |
| N18.9 | CKD, unspecified | N18.2 | Stage 2 (eGFR 60-89) |
| N18.9 | CKD, unspecified | N18.30 | Stage 3 unspecified |
| N18.9 | CKD, unspecified | N18.31 | Stage 3a (eGFR 45-59) |
| N18.9 | CKD, unspecified | N18.32 | Stage 3b (eGFR 30-44) |
| N18.9 | CKD, unspecified | N18.4 | Stage 4 (eGFR 15-29) |
| N18.9 | CKD, unspecified | N18.5 | Stage 5 (eGFR < 15) |
| N18.9 | CKD, unspecified | N18.6 | ESRD (on dialysis) |

**Key check:** Pull latest eGFR (LOINC 33914-3 or 48642-3) to determine stage. If both eGFR and creatinine are available, use the CKD-EPI equation result.

### Hyperlipidemia (E78)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| E78.5 | Hyperlipidemia, unspecified | E78.00 | Pure hypercholesterolemia, unspecified |
| E78.5 | Hyperlipidemia, unspecified | E78.01 | Familial hypercholesterolemia |
| E78.5 | Hyperlipidemia, unspecified | E78.1 | Pure hypertriglyceridemia |
| E78.5 | Hyperlipidemia, unspecified | E78.2 | Mixed hyperlipidemia |
| E78.5 | Hyperlipidemia, unspecified | E78.4 | Other hyperlipidemia |

**Key check:** Pull latest lipid panel. Elevated LDL only = E78.00. Elevated triglycerides only = E78.1. Both elevated = E78.2. Family history of premature CAD + very high LDL = consider E78.01.

### Asthma (J45)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| J45.909 | Unspecified asthma, uncomplicated | J45.20 | Mild intermittent |
| J45.909 | Unspecified asthma, uncomplicated | J45.30 | Mild persistent |
| J45.909 | Unspecified asthma, uncomplicated | J45.40 | Moderate persistent |
| J45.909 | Unspecified asthma, uncomplicated | J45.50 | Severe persistent |

**Classification based on treatment step:**
- Mild intermittent: SABA PRN only, symptoms <= 2 days/week
- Mild persistent: Low-dose ICS or LTRA
- Moderate persistent: Medium-dose ICS or low-dose ICS/LABA
- Severe persistent: High-dose ICS/LABA +/- add-on

Add 5th character for exacerbation status:
- 0 = uncomplicated
- 1 = with acute exacerbation
- 2 = with status asthmaticus

### Depression (F32-F33)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| F32.9 | MDD single episode, unspecified | F32.0 | Single episode, mild |
| F32.9 | MDD single episode, unspecified | F32.1 | Single episode, moderate |
| F32.9 | MDD single episode, unspecified | F32.2 | Single episode, severe without psychosis |
| F33.9 | MDD recurrent, unspecified | F33.0 | Recurrent, mild |
| F33.9 | MDD recurrent, unspecified | F33.1 | Recurrent, moderate |
| F33.9 | MDD recurrent, unspecified | F33.2 | Recurrent, severe without psychosis |
| F33.9 | MDD recurrent, unspecified | F33.40 | Recurrent, in remission, unspecified |
| F33.9 | MDD recurrent, unspecified | F33.41 | Recurrent, in partial remission |
| F33.9 | MDD recurrent, unspecified | F33.42 | Recurrent, in full remission |

**Key check:** PHQ-9 scores (LOINC 44249-1): 5-9 mild, 10-14 moderate, 15-19 moderately severe, 20-27 severe. Single vs recurrent based on episode history.

### Musculoskeletal -- Laterality

Many MSK codes require laterality. Common omissions:

| Unspecified Code | Description | Right | Left | Bilateral |
|-----------------|-------------|-------|------|-----------|
| M17.9 | Osteoarthritis of knee, unspecified | M17.11 | M17.12 | M17.0 |
| M16.9 | Osteoarthritis of hip, unspecified | M16.11 | M16.12 | M16.0 |
| M54.50 | Low back pain, unspecified | M54.51 | M54.52 | -- |
| M79.3 | Panniculitis, unspecified | -- | -- | -- |
| G56.00 | Carpal tunnel, unspecified | G56.01 | G56.02 | G56.03 |
| M75.10 | Rotator cuff syndrome, unspecified | M75.11 | M75.12 | -- |

### COPD (J44)

| Vague Code | Description | Improved Code | When to Use |
|-----------|-------------|---------------|-------------|
| J44.9 | COPD, unspecified | J44.0 | With acute lower respiratory infection |
| J44.9 | COPD, unspecified | J44.1 | With acute exacerbation |

GOLD staging is clinical, not captured in ICD-10 directly, but document the severity in the note to support medical decision-making.

## Lab Values Supporting Code Specificity

| Condition | Lab Test | LOINC | Value Thresholds |
|-----------|----------|-------|-----------------|
| CKD staging | eGFR | 33914-3 | See N18.x table above |
| Diabetes control | HbA1c | 4548-4 | < 5.7% normal, 5.7-6.4% prediabetes, >= 6.5% diabetes |
| Hyperlipidemia type | LDL | 18262-6 | > 190 consider familial |
| Hyperlipidemia type | Triglycerides | 2571-8 | > 150 elevated, > 500 severe |
| Anemia type | Ferritin | 2276-4 | < 30 iron deficiency |
| Anemia type | MCV | 787-2 | < 80 microcytic, > 100 macrocytic |
| Depression severity | PHQ-9 | 44249-1 | 5-9 mild, 10-14 moderate, 15+ severe |
| Heart failure type | EF | 10230-1 | < 40% HFrEF, 40-49% borderline, >= 50% HFpEF |

## Risk Adjustment Impact

More specific ICD-10 codes directly impact:
- **HCC (Hierarchical Condition Category) scores** for Medicare Advantage
- **CMS quality measures** that require specific codes
- **Risk-adjusted payment** under value-based care models
- **HEDIS measures** that depend on accurate problem list coding

High-impact HCC conditions where specificity matters most:
- Diabetes with complications (HCC 18) vs without (HCC 19)
- Heart failure (HCC 85)
- CKD Stage 4-5 (HCC 137-138)
- COPD (HCC 111)
- Vascular disease (HCC 107-108)
