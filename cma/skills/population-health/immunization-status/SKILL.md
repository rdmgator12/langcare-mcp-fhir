---
name: langcare-immunization-status
description: >
  Checks patient immunization status against CDC/ACIP recommended schedules
  by age group and risk factors. Identifies missing or overdue vaccinations
  and generates recommendations. Use when asked to check immunization status,
  vaccine history, immunizations due, vaccination compliance, or what vaccines
  does this patient need.
---

# Immunization Status Checker

## When to Use This Skill
Use when a clinician needs to evaluate a patient's vaccination history against current CDC/ACIP recommendations and identify gaps.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender for age-based schedule)
2. Use `fhir_search` to pull all Immunization resources (status=completed)
3. Use `fhir_search` to pull active Condition resources for risk-based vaccine recommendations (immunocompromised, asplenia, pregnancy, chronic conditions)
4. Map completed immunizations against CDC/ACIP schedule (see references/cdc-schedule.md) by age group and risk factors
5. Identify: completed (up to date), due now (within recommended window), overdue (past recommended date), contraindicated (based on conditions/allergies)
6. Present immunization status checklist with specific vaccine recommendations

## FHIR Resources
- **Patient** -- Age for schedule determination
- **Immunization** -- Vaccination history: vaccineCode (CVX), occurrenceDateTime, status, doseNumberPositiveInt
- **Condition** -- Risk factors for additional vaccine recommendations
- **AllergyIntolerance** -- Vaccine contraindications (egg allergy, gelatin, neomycin)

## FHIR Query Examples
### Pull Immunization History
```
fhir_search(resourceType="Immunization", queryParams="patient=[patient-id]&status=completed")
```

### Pull Risk Conditions
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&clinical-status=active")
```

## Clinical Guidelines
- CDC/ACIP Recommended Immunization Schedules (updated annually)
- CDC catch-up immunization schedule
- ACOG recommendations for immunizations in pregnancy

## Interpretation Guide
- Organize by vaccine: name, series status (dose X of Y completed), date of last dose, next dose due, status (current/due/overdue/contraindicated)
- Age-specific schedule: childhood (0-18), adult (19-64), older adult (65+)
- Risk-based additions: pneumococcal (PCV20 or PCV15+PPSV23) for age 65+ or immunocompromised, hepatitis A/B for risk factors, meningococcal for asplenia
- Pregnancy-specific: Tdap each pregnancy 27-36 weeks, flu in flu season, avoid live vaccines (MMR, varicella)
- Flag high-priority gaps: no flu vaccine in season, no COVID primary series, no Tdap in 10+ years

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
