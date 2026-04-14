---
name: langcare-renal-function
description: >
  Assesses renal function using KDIGO staging criteria from creatinine, eGFR,
  BUN, urine albumin, and electrolytes. Tracks CKD progression and flags
  medication dose adjustments needed for renal impairment. Use when asked
  about renal function, kidney function, CKD staging, eGFR trends, or
  nephrology assessment.
---

# Renal Function Dashboard

## When to Use This Skill
Use when a clinician needs a comprehensive renal function assessment including CKD staging, progression trending, electrolyte monitoring, and medication dose adjustment recommendations.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender, race for eGFR calculation context)
2. Use `fhir_search` to pull creatinine (LOINC 2160-0) and eGFR (LOINC 33914-3) trends over 12+ months
3. Use `fhir_search` to pull urine albumin/creatinine ratio (LOINC 14959-1) for albuminuria staging
4. Use `fhir_search` to pull electrolytes (K+, Ca++, phosphate, bicarb) and CBC (anemia of CKD)
5. Use `fhir_search` to pull active Condition resources for CKD staging documentation and comorbidities
6. Stage CKD per KDIGO criteria (see references/kdigo-staging.md): GFR category (G1-G5) and albuminuria category (A1-A3)
7. Use `fhir_search` to pull active MedicationRequest resources; flag medications requiring renal dose adjustment
8. Calculate rate of eGFR decline (mL/min/year); flag rapid decline (>5 mL/min/year)

## FHIR Resources
- **Observation** -- Creatinine (2160-0), eGFR (33914-3), BUN (3094-0), urine albumin/Cr (14959-1), electrolytes, CBC
- **Condition** -- CKD staging, diabetes, hypertension
- **MedicationRequest** -- Medications requiring renal dose adjustment
- **Patient** -- Demographics for eGFR context

## FHIR Query Examples
### Pull eGFR Trend
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=33914-3&_sort=date&_count=20")
```

### Pull Urine Albumin/Creatinine Ratio
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=14959-1&_sort=-date&_count=5")
```

## Clinical Guidelines
- KDIGO 2024 Clinical Practice Guideline for CKD Evaluation and Management
- ADA Standards for diabetic kidney disease
- ACC/AHA guidelines for CKD and cardiovascular risk

## Interpretation Guide
- CKD staging uses both GFR and albuminuria: G1-G5 x A1-A3 grid (see references/kdigo-staging.md)
- Rate of decline: calculate from serial eGFR values; >5 mL/min/year is rapid progression warranting nephrology referral
- Medication adjustments: flag metformin (contraindicated if eGFR <30), NSAIDs (avoid in CKD), ACEi/ARB (renal-protective but monitor K+ and Cr), gabapentin/pregabalin (dose reduce), DOACs (dose adjust per eGFR)
- Anemia of CKD: check Hgb, iron studies; consider EPO if Hgb <10 and iron replete
- Electrolyte monitoring: hyperkalemia (common with ACEi/ARB + CKD), hyperphosphatemia (CKD 4-5), metabolic acidosis (bicarb <22)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
