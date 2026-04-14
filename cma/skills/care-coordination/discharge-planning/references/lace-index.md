# LACE Readmission Risk Index Reference

## LACE Components

### L -- Length of Stay

| LOS (days) | Points |
|-----------|--------|
| 1 | 1 |
| 2 | 2 |
| 3 | 3 |
| 4-6 | 4 |
| 7-13 | 5 |
| >= 14 | 7 |

### A -- Acuity of Admission

| Admission Type | Points |
|---------------|--------|
| Elective / Planned | 0 |
| Urgent / Emergent (via ED) | 3 |

FHIR: Check `Encounter.hospitalization.admitSource` or if the encounter was preceded by an ED encounter.

### C -- Comorbidity (Charlson Comorbidity Index)

| Charlson Score | LACE Points |
|---------------|-------------|
| 0 | 0 |
| 1 | 1 |
| 2 | 2 |
| 3 | 3 |
| >= 4 | 5 |

#### Charlson Comorbidity Index Conditions

| Condition | Points | SNOMED/ICD-10 |
|-----------|--------|---------------|
| Myocardial infarction | 1 | I21-I22 / 22298006 |
| Congestive heart failure | 1 | I50 / 42343007 |
| Peripheral vascular disease | 1 | I73 / 400047006 |
| Cerebrovascular disease | 1 | I60-I69 / 230690007 |
| Dementia | 1 | F00-F03 / 52448006 |
| Chronic pulmonary disease | 1 | J40-J47 / 13645005 |
| Connective tissue disease | 1 | M30-M36 / 105969002 |
| Peptic ulcer disease | 1 | K25-K28 / 13200003 |
| Mild liver disease | 1 | K70-K73 / 235856003 |
| Diabetes (uncomplicated) | 1 | E10-E11 (without complications) |
| Diabetes (with complications) | 2 | E10-E11 (with .2-.5) |
| Hemiplegia | 2 | G81 / 50582007 |
| Moderate/severe renal disease | 2 | N18.3-N18.5 / 709044004 |
| Any malignancy (including lymphoma/leukemia) | 2 | C00-C97 / 363346000 |
| Moderate/severe liver disease | 3 | K72-K74 (with portal HTN) |
| Metastatic solid tumor | 6 | C77-C80 / 128462008 |
| AIDS/HIV | 6 | B20-B24 / 62479008 |

### E -- Emergency Department Visits (6 months prior)

| ED Visits | Points |
|-----------|--------|
| 0 | 0 |
| 1 | 1 |
| 2 | 2 |
| 3 | 3 |
| >= 4 | 4 |

FHIR: Count Encounter resources with `class=EMER` in the 6 months preceding admission.

## LACE Score Interpretation

| Total Score | Risk Level | 30-Day Readmission Risk | Action |
|-------------|-----------|------------------------|--------|
| 0-4 | Low | ~5% | Standard discharge |
| 5-9 | Moderate | ~10-15% | Enhanced discharge planning |
| 10-12 | High | ~20-25% | Intensive post-discharge interventions |
| >= 13 | Very High | ~30%+ | TCM visit within 7 days, nurse callback within 48h, care coordination |

## Post-Discharge Interventions by Risk

### Low Risk (LACE 0-4)
- Standard discharge instructions
- PCP follow-up within 14 days

### Moderate Risk (LACE 5-9)
- PCP follow-up within 7 days
- Medication reconciliation phone call within 72 hours
- Ensure pharmacy access before discharge

### High Risk (LACE 10-12)
- PCP follow-up within 3-7 days
- Transitional Care Management (TCM) visit
- RN phone call within 24-48 hours
- Pharmacist medication review
- Home health consideration

### Very High Risk (LACE >= 13)
- All high-risk interventions
- Consider care coordination/case management
- Consider post-acute facility vs direct home
- Community health worker engagement
- Ensure social determinants addressed (transport, food, medication access)
