# VTE Scoring Systems Reference

## Wells Criteria for DVT

Clinical prediction rule for pre-test probability of deep vein thrombosis. Score range: -2 to 9.

### Scoring Criteria

| Criterion | Points | Notes |
|-----------|--------|-------|
| Active cancer (treatment within 6 months or palliative) | +1 | SNOMED 363346000 |
| Paralysis, paresis, or recent plaster immobilization of lower extremity | +1 | SNOMED 309585006, 397169001 |
| Recently bedridden >3 days or major surgery within 12 weeks | +1 | Check Procedure resources, Encounter data |
| Localized tenderness along distribution of deep venous system | +1 | Clinical finding, may not be in FHIR |
| Entire leg swollen | +1 | Clinical finding |
| Calf swelling >3 cm compared to asymptomatic side (measured 10cm below tibial tuberosity) | +1 | Clinical finding |
| Pitting edema confined to symptomatic leg | +1 | Clinical finding |
| Collateral superficial veins (non-varicose) | +1 | Clinical finding |
| Previously documented DVT | +1 | SNOMED 128053003 (check historical conditions) |
| Alternative diagnosis at least as likely as DVT | -2 | Clinical judgment -- cellulitis, Baker's cyst, muscle strain, post-phlebitic syndrome |

### Interpretation

| Score | Category | Pre-Test Probability | Prevalence of DVT |
|-------|----------|---------------------|-------------------|
| <= 0 | Low (DVT unlikely) | ~5% | 5% |
| 1-2 | Moderate (DVT unlikely) | ~17% | 17% |
| >= 3 | High (DVT likely) | ~53% | 53% |

**Two-level model (preferred):**

| Score | Category | Action |
|-------|----------|--------|
| <= 1 | DVT unlikely | D-dimer testing. If negative: DVT excluded. If positive: compression ultrasound. |
| >= 2 | DVT likely | Compression ultrasound directly. D-dimer not useful for exclusion. |

---

## Wells Criteria for Pulmonary Embolism

Clinical prediction rule for pre-test probability of PE. Score range: 0 to 12.5.

### Scoring Criteria

| Criterion | Points | Notes |
|-----------|--------|-------|
| Clinical signs/symptoms of DVT (leg swelling, pain with palpation) | +3 | Clinical finding or Condition with SNOMED 128053003 |
| PE is #1 diagnosis or equally likely | +3 | Clinical judgment |
| Heart rate > 100 bpm | +1.5 | LOINC 8867-4. Check most recent Observation. |
| Immobilization (>=3 days) or surgery in previous 4 weeks | +1.5 | Check Procedure resources, Encounter |
| Previous DVT/PE | +1.5 | SNOMED 128053003 (DVT), 59282003 (PE) |
| Hemoptysis | +1 | SNOMED 66857006 |
| Malignancy (treatment within 6 months or palliative) | +1 | SNOMED 363346000 |

### Interpretation (Three-Level)

| Score | Category | Pre-Test Probability |
|-------|----------|---------------------|
| 0-1 | Low | ~1.3% |
| 2-6 | Moderate | ~16.2% |
| > 6 | High | ~40.6% |

### Interpretation (Two-Level, Preferred)

| Score | Category | Action |
|-------|----------|--------|
| <= 4 | PE unlikely | D-dimer testing. If negative: PE excluded. If positive: CTPA. |
| > 4 | PE likely | CTPA directly. D-dimer not useful for exclusion. |

### PERC Rule (Pulmonary Embolism Rule-Out Criteria)

For LOW-risk patients (Wells <=4) with low clinical suspicion, apply PERC before D-dimer. If ALL criteria met, PE effectively ruled out without D-dimer:
- Age < 50
- Heart rate < 100
- SpO2 >= 95%
- No hemoptysis
- No estrogen use
- No prior DVT/PE
- No unilateral leg swelling
- No surgery/trauma requiring hospitalization in past 4 weeks

If any PERC criterion is NOT met, proceed to D-dimer.

---

## Revised Geneva Score for PE

Alternative to Wells for PE probability. Advantage: no subjective criterion ("PE is #1 diagnosis"). Score range: 0-25.

### Scoring Criteria

| Criterion | Points | Source |
|-----------|--------|--------|
| Age > 65 | +1 | Patient.birthDate |
| Previous DVT or PE | +3 | Condition: SNOMED 128053003, 59282003 |
| Surgery under general anesthesia or lower limb fracture within 1 month | +2 | Procedure resources |
| Active malignancy (solid or hematologic, active or cured <1 year) | +2 | Condition: SNOMED 363346000 |
| Unilateral lower limb pain | +3 | Clinical finding |
| Hemoptysis | +2 | SNOMED 66857006 or clinical report |
| Heart rate 75-94 bpm | +3 | Observation: LOINC 8867-4 |
| Heart rate >= 95 bpm | +5 | Observation: LOINC 8867-4 |
| Pain on lower limb deep vein palpation AND unilateral edema | +4 | Clinical finding |

Note: Heart rate categories are mutually exclusive. Score 75-94 OR >=95, not both.

### Interpretation

| Score | Category | Pre-Test Probability |
|-------|----------|---------------------|
| 0-3 | Low | ~8% |
| 4-10 | Intermediate | ~29% |
| >= 11 | High | ~74% |

### Two-Level Interpretation

| Score | Category | Action |
|-------|----------|--------|
| 0-5 | PE unlikely | D-dimer -> if negative, PE excluded |
| >= 6 | PE likely | CTPA directly |

---

## Caprini Score (VTE Risk in Surgical Patients)

Comprehensive VTE risk assessment for surgical and hospitalized patients. Used to determine prophylaxis intensity and duration.

### 1-Point Factors

| Factor | Source |
|--------|--------|
| Age 41-60 | Patient.birthDate |
| Minor surgery planned | Procedure |
| History of prior major surgery (<1 month) | Procedure |
| Varicose veins | Condition: SNOMED 70691001 |
| Inflammatory bowel disease | Condition: SNOMED 128613002 (Crohn's 34000006, UC 64766004) |
| Swollen legs (current) | Clinical finding |
| Obesity (BMI > 25) | Observation: LOINC 39156-5 |
| Acute MI | Condition: SNOMED 22298006 |
| CHF (<1 month) | Condition: SNOMED 42343007 |
| Sepsis (<1 month) | Condition: SNOMED 91302008 |
| Serious lung disease including pneumonia (<1 month) | Condition |
| Abnormal pulmonary function (COPD) | Condition: SNOMED 13645005 |
| Medical patient currently on bed rest | Encounter/clinical status |
| Oral contraceptives or HRT | MedicationRequest |

### 2-Point Factors

| Factor | Source |
|--------|--------|
| Age 61-74 | Patient.birthDate |
| Arthroscopic surgery | Procedure: SNOMED 71106006 |
| Malignancy (present or previous) | Condition: SNOMED 363346000 |
| Major surgery (>45 minutes) | Procedure |
| Laparoscopic surgery (>45 minutes) | Procedure: SNOMED 108191006 |
| Patient confined to bed (>72 hours) | Clinical status |
| Immobilizing plaster cast | Procedure/Condition |
| Central venous access | Procedure: SNOMED 233527006 |

### 3-Point Factors

| Factor | Source |
|--------|--------|
| Age >= 75 | Patient.birthDate |
| History of DVT/PE | Condition: SNOMED 128053003, 59282003 |
| Family history of thrombosis | FamilyMemberHistory |
| Factor V Leiden positive | Condition or DiagnosticReport |
| Prothrombin 20210A mutation | Condition or DiagnosticReport |
| Lupus anticoagulant positive | Observation |
| Anticardiolipin antibodies elevated | Observation |
| Elevated serum homocysteine | Observation: LOINC 2028-9 |
| Heparin-induced thrombocytopenia (HIT) | Condition: SNOMED 73641003 |
| Other congenital or acquired thrombophilia | Condition: SNOMED 128599005 |

### 5-Point Factors

| Factor | Source |
|--------|--------|
| Stroke (<1 month) | Condition: SNOMED 230690007 |
| Multiple trauma (<1 month) | Condition |
| Acute spinal cord injury (<1 month) | Condition: SNOMED 90584004 |
| Hip, pelvis, or leg fracture (<1 month) | Condition |
| Hip or knee arthroplasty | Procedure: SNOMED 179344006, 179406003 |

### Risk Stratification

| Total Score | Risk Level | VTE Incidence (without prophylaxis) |
|-------------|-----------|--------------------------------------|
| 0 | Lowest risk | < 0.5% |
| 1-2 | Low risk | ~1.5% |
| 3-4 | Moderate risk | ~3% |
| >= 5 | High risk | ~6% |
| >= 9 | Highest risk | ~11% |
