# SOFA, qSOFA, and SIRS Scoring Reference

## SOFA Score (Sequential Organ Failure Assessment)

Total score range: 0-24. Six organ systems scored 0-4 each. A SOFA score increase of >=2 points from baseline indicates organ dysfunction consistent with sepsis (Sepsis-3 definition).

### Respiration: PaO2/FiO2 Ratio (mmHg)

| Score | PaO2/FiO2 | Notes |
|-------|-----------|-------|
| 0 | >= 400 | Normal |
| 1 | 300 - 399 | Mild impairment |
| 2 | 200 - 299 | Moderate impairment |
| 3 | 100 - 199 WITH respiratory support | Severe, on ventilator |
| 4 | < 100 WITH respiratory support | Critical, on ventilator |

Calculation: PaO2 (mmHg) / FiO2 (as decimal, e.g., 0.40 for 40%). If PaO2 unavailable and patient on room air with SpO2 available, use SpO2/FiO2 ratio as surrogate (less validated). If FiO2 unknown and patient on room air, use FiO2 = 0.21.

### Coagulation: Platelets (x10^3/uL)

| Score | Platelets | Notes |
|-------|-----------|-------|
| 0 | >= 150 | Normal |
| 1 | 100 - 149 | Mild thrombocytopenia |
| 2 | 50 - 99 | Moderate thrombocytopenia |
| 3 | 20 - 49 | Severe thrombocytopenia |
| 4 | < 20 | Critical |

LOINC: 777-3 (Platelet count). Units may vary: 10^3/uL, K/uL, x10^9/L (equivalent to 10^3/uL).

### Hepatic: Bilirubin (mg/dL)

| Score | Bilirubin (mg/dL) | Bilirubin (umol/L) | Notes |
|-------|-------------------|---------------------|-------|
| 0 | < 1.2 | < 20 | Normal |
| 1 | 1.2 - 1.9 | 20 - 32 | Mild elevation |
| 2 | 2.0 - 5.9 | 33 - 101 | Moderate elevation |
| 3 | 6.0 - 11.9 | 102 - 204 | Severe elevation |
| 4 | >= 12.0 | > 204 | Critical |

LOINC: 1975-2 (Total bilirubin, serum). Check units in `valueQuantity.unit` -- mg/dL vs umol/L.

### Cardiovascular: Mean Arterial Pressure (MAP) and Vasopressors

| Score | Criteria |
|-------|----------|
| 0 | MAP >= 70 mmHg, no vasopressors |
| 1 | MAP < 70 mmHg, no vasopressors |
| 2 | Dopamine <= 5 mcg/kg/min OR any dobutamine |
| 3 | Dopamine > 5 mcg/kg/min OR epinephrine <= 0.1 mcg/kg/min OR norepinephrine <= 0.1 mcg/kg/min |
| 4 | Dopamine > 15 mcg/kg/min OR epinephrine > 0.1 mcg/kg/min OR norepinephrine > 0.1 mcg/kg/min |

MAP calculation: MAP = (SBP + 2 x DBP) / 3. If MAP not directly available, calculate from systolic (LOINC 8480-6) and diastolic (LOINC 8462-4).

Vasopressor RxNorm codes:
- Norepinephrine: 3628
- Epinephrine: 3616
- Dopamine: 35208
- Dobutamine: 3616 (check medication name to differentiate from epinephrine)
- Vasopressin: 11149
- Phenylephrine: 8163

### Central Nervous System: Glasgow Coma Scale (GCS)

| Score | GCS | Notes |
|-------|-----|-------|
| 0 | 15 | Normal |
| 1 | 13 - 14 | Mild impairment |
| 2 | 10 - 12 | Moderate impairment |
| 3 | 6 - 9 | Severe impairment |
| 4 | < 6 | Critical |

LOINC codes:
- 9269-2: Glasgow coma score total
- 9267-6: GCS Eye opening
- 9270-0: GCS Verbal response
- 9268-4: GCS Motor response

If total GCS unavailable, sum component scores. If sedated, use pre-sedation GCS or note as confounded.

### Renal: Creatinine (mg/dL) or Urine Output

| Score | Creatinine (mg/dL) | Urine Output | Notes |
|-------|---------------------|--------------|-------|
| 0 | < 1.2 | >= 500 mL/day | Normal |
| 1 | 1.2 - 1.9 | -- | Mild elevation |
| 2 | 2.0 - 3.4 | < 500 mL/day | Oliguria |
| 3 | 3.5 - 4.9 | < 200 mL/day | Severe oliguria |
| 4 | >= 5.0 | < 200 mL/day | Anuria/renal failure |

LOINC: 2160-0 (Serum creatinine). Urine output: LOINC 9187-6 (Urine output 24h). Use creatinine if urine output unavailable.

### SOFA Score Interpretation

| Total SOFA | Mortality Risk | Clinical Significance |
|------------|---------------|----------------------|
| 0-1 | < 3% | Minimal organ dysfunction |
| 2-3 | ~5% | Mild organ dysfunction |
| 4-5 | ~10% | Moderate organ dysfunction |
| 6-7 | ~15-20% | Significant organ dysfunction |
| 8-9 | ~20-30% | Severe organ dysfunction |
| 10-11 | ~30-40% | Very severe |
| 12-14 | ~50% | Critical |
| >= 15 | > 80% | Extremely high mortality |

Key threshold: SOFA increase >=2 from baseline = sepsis (Sepsis-3). If baseline unknown, assume SOFA baseline = 0 for previously healthy patients. For patients with chronic organ dysfunction, estimate baseline from prior records.

---

## qSOFA (Quick SOFA)

Bedside screening tool. Does NOT require lab values. Score 0-3.

| Criterion | Points | Threshold |
|-----------|--------|-----------|
| Respiratory rate >= 22 breaths/min | 1 | LOINC 9279-1 |
| Systolic blood pressure <= 100 mmHg | 1 | LOINC 8480-6 |
| Altered mental status (GCS < 15) | 1 | LOINC 9269-2 |

### Interpretation

| qSOFA Score | Interpretation | Action |
|-------------|---------------|--------|
| 0 | Low risk | Continue monitoring |
| 1 | Low-intermediate risk | Increase monitoring frequency |
| >= 2 | High risk for poor outcome | Investigate for organ dysfunction, consider ICU evaluation, calculate full SOFA |

Limitations:
- Sensitivity ~50% -- a low qSOFA does NOT rule out sepsis
- Specificity ~85% -- useful for identifying high-risk patients
- Should prompt further investigation, not used as sole diagnostic criterion
- Does not replace clinical judgment or full SOFA scoring

---

## SIRS Criteria (Systemic Inflammatory Response Syndrome)

Legacy criteria. Still widely used in clinical practice despite Sepsis-3 recommendation for SOFA. Score 0-4.

| Criterion | Points | Threshold | LOINC |
|-----------|--------|-----------|-------|
| Temperature > 38.3C (100.9F) or < 36.0C (96.8F) | 1 | 8310-5 | Body temperature |
| Heart rate > 90 bpm | 1 | 8867-4 | Heart rate |
| Respiratory rate > 20 or PaCO2 < 32 mmHg | 1 | 9279-1 / 2019-8 | RR / PaCO2 |
| WBC > 12,000/mm3 or < 4,000/mm3 or > 10% bands | 1 | 6690-2 / 764-1 | WBC / Band neutrophils |

### Interpretation

| SIRS Score | With Suspected Infection | Without Suspected Infection |
|------------|-------------------------|-----------------------------|
| 0-1 | SIRS not met | SIRS not met |
| >= 2 | Sepsis (SIRS-based definition) | SIRS without sepsis (consider non-infectious causes) |
| >= 2 + organ dysfunction | Severe sepsis | Investigate for infection |
| >= 2 + refractory hypotension | Septic shock (SIRS-based) | Investigate for infection + other shock etiologies |

Non-infectious SIRS causes: trauma, burns, pancreatitis, major surgery, adrenal insufficiency, pulmonary embolism, drug reactions, transfusion reactions.

Limitations:
- Low specificity -- many non-septic conditions trigger SIRS
- High sensitivity (~90%) -- useful for ruling out sepsis when SIRS < 2
- Sepsis-3 (2016) recommended replacing SIRS-based definition with SOFA-based definition
- Both SIRS and qSOFA/SOFA should be calculated for comprehensive assessment
