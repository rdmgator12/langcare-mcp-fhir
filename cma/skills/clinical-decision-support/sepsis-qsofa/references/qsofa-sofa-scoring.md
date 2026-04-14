# qSOFA, SOFA, and SIRS Scoring Reference

## qSOFA (Quick SOFA) -- Bedside Screening

| Criterion | Threshold | Points |
|-----------|-----------|--------|
| Respiratory rate | >= 22 breaths/min | +1 |
| Systolic blood pressure | <= 100 mmHg | +1 |
| Altered mental status | GCS < 15 | +1 |

**Interpretation:**
- Score 0-1: Low risk
- Score >= 2: High risk for poor outcome; investigate for organ dysfunction with full SOFA

## SIRS Criteria

| Criterion | Threshold | Points |
|-----------|-----------|--------|
| Temperature | > 38.3C (100.9F) OR < 36.0C (96.8F) | +1 |
| Heart rate | > 90 bpm | +1 |
| Respiratory rate | > 20 breaths/min OR PaCO2 < 32 mmHg | +1 |
| WBC | > 12,000/mm3 OR < 4,000/mm3 OR > 10% bands | +1 |

**Interpretation:**
- Score >= 2 with suspected/confirmed infection = Sepsis (traditional definition)
- SIRS is sensitive but not specific; many non-infectious conditions trigger SIRS

## SOFA Score (Sequential Organ Failure Assessment)

### Respiration: PaO2/FiO2 Ratio

| Score | PaO2/FiO2 | Condition |
|-------|-----------|-----------|
| 0 | >= 400 | Normal |
| 1 | 300-399 | Mild impairment |
| 2 | 200-299 | Moderate (may need O2) |
| 3 | 100-199 with respiratory support | Severe |
| 4 | < 100 with respiratory support | Critical |

### Coagulation: Platelets

| Score | Platelets (x10^3/uL) |
|-------|----------------------|
| 0 | >= 150 |
| 1 | 100-149 |
| 2 | 50-99 |
| 3 | 20-49 |
| 4 | < 20 |

### Liver: Bilirubin

| Score | Total Bilirubin (mg/dL) |
|-------|------------------------|
| 0 | < 1.2 |
| 1 | 1.2-1.9 |
| 2 | 2.0-5.9 |
| 3 | 6.0-11.9 |
| 4 | >= 12.0 |

### Cardiovascular: Hypotension

| Score | Condition |
|-------|-----------|
| 0 | MAP >= 70 mmHg |
| 1 | MAP < 70 mmHg |
| 2 | Dopamine <= 5 mcg/kg/min OR dobutamine (any dose) |
| 3 | Dopamine > 5 mcg/kg/min OR epinephrine <= 0.1 mcg/kg/min OR norepinephrine <= 0.1 mcg/kg/min |
| 4 | Dopamine > 15 mcg/kg/min OR epinephrine > 0.1 mcg/kg/min OR norepinephrine > 0.1 mcg/kg/min |

### Central Nervous System: Glasgow Coma Scale

| Score | GCS |
|-------|-----|
| 0 | 15 |
| 1 | 13-14 |
| 2 | 10-12 |
| 3 | 6-9 |
| 4 | < 6 |

### Renal: Creatinine or Urine Output

| Score | Creatinine (mg/dL) | Urine Output |
|-------|-------------------|--------------|
| 0 | < 1.2 | -- |
| 1 | 1.2-1.9 | -- |
| 2 | 2.0-3.4 | < 500 mL/day |
| 3 | 3.5-4.9 | < 200 mL/day |
| 4 | >= 5.0 | < 200 mL/day |

**SOFA Interpretation:**
- Baseline SOFA assumed 0 for patients not known to have pre-existing organ dysfunction
- Acute increase >= 2 from baseline = organ dysfunction = **Sepsis** (Sepsis-3 definition)
- Total SOFA > 11: mortality > 95%

## Sepsis Definitions Summary

| Term | Definition |
|------|-----------|
| **Sepsis** (Sepsis-3) | Life-threatening organ dysfunction caused by dysregulated host response to infection. Operationalized as suspected infection + SOFA increase >= 2 |
| **Septic Shock** | Sepsis + vasopressor requirement to maintain MAP >= 65 + lactate > 2 mmol/L despite adequate fluid resuscitation |

## Surviving Sepsis Campaign Hour-1 Bundle

| Element | Target | FHIR Check |
|---------|--------|------------|
| Measure lactate | Within 1 hour | Observation: code=2524-7 or 32693-4 |
| Obtain blood cultures before antibiotics | Within 1 hour | ServiceRequest: code=600-7 |
| Administer broad-spectrum antibiotics | Within 1 hour | MedicationRequest: active antibiotics |
| Begin rapid fluid resuscitation | 30 mL/kg crystalloid for hypotension or lactate >= 4 | MedicationAdministration or order for LR/NS |
| Apply vasopressors | If MAP < 65 after fluid resuscitation | MedicationRequest: norepinephrine, vasopressin |
| Re-measure lactate | If initial lactate > 2, repeat within 2-4 hours | Observation: second lactate result |

## LOINC Codes for Sepsis Screening

| Parameter | LOINC | Notes |
|-----------|-------|-------|
| Temperature | 8310-5 | Body temperature |
| Heart rate | 8867-4 | Heart rate |
| Respiratory rate | 9279-1 | Respiratory rate |
| Systolic BP | 8480-6 | Systolic blood pressure |
| SpO2 | 2708-6 | Oxygen saturation |
| WBC | 6690-2 | Leukocytes |
| Lactate (arterial) | 2524-7 | Lactic acid |
| Lactate (venous) | 32693-4 | Venous lactic acid |
| Creatinine | 2160-0 | Serum creatinine |
| Bilirubin (total) | 1975-2 | Total bilirubin |
| Platelets | 777-3 | Platelet count |
| PaO2 | 2703-7 | Arterial O2 partial pressure |
| FiO2 | 3150-0 | Inspired O2 fraction |
| GCS total | 9269-2 | Glasgow coma score |
| Blood culture | 600-7 | Blood culture order |
