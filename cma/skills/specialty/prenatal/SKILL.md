---
name: langcare-prenatal
description: >
  Performs structured prenatal visit assessment organized by trimester with
  ACOG guideline alignment. Tracks gestational age, weight gain, BP trends,
  fetal assessment, trimester-specific labs, and risk factors. Use when asked
  to review prenatal visit, prenatal assessment, OB visit, pregnancy checkup,
  trimester labs, or gestational age evaluation.
---

# Prenatal Visit Workflow

## When to Use This Skill
Use when a clinician needs a structured prenatal visit assessment organized by current gestational age and trimester with ACOG-aligned lab tracking and risk assessment.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (maternal age -- flag if >=35 for AMA)
2. Use `fhir_search` to pull pregnancy Condition (SNOMED 77386006) and determine gestational age from onset or EDD Observation (LOINC 11778-8)
3. Use `fhir_search` to pull vital signs trend: BP (preeclampsia monitoring), weight (gain trajectory per IOM guidelines)
4. Use `fhir_search` to pull fetal Observations: fetal heart rate (LOINC 11996-6), fundal height (LOINC 11948-7)
5. Use `fhir_search` to pull prenatal laboratory results organized by trimester milestone (see references/acog-schedule.md)
6. Assess preeclampsia risk: BP trend, urine protein, aspirin prophylaxis status
7. Use `fhir_search` to pull active MedicationRequest: verify prenatal vitamin, aspirin if indicated, Rhogam timing
8. Use `fhir_search` to pull Procedure resources for ultrasound history
9. Present structured visit summary with current trimester context, risk flags, and upcoming milestones

## FHIR Resources
- **Patient** -- Maternal age
- **Condition** -- Pregnancy diagnosis, complications (GDM, preeclampsia, preterm risk)
- **Observation** -- Vitals (BP, weight), fetal (FHR, fundal height), labs (by trimester), EDD
- **MedicationRequest** -- Prenatal vitamins, aspirin, Rhogam, insulin
- **Procedure** -- Ultrasounds, amniocentesis
- **DiagnosticReport** -- Ultrasound reports, lab panels

## FHIR Query Examples
### Pull Pregnancy Condition
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&code=77386006&clinical-status=active")
```

### Pull EDD
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=11778-8&_sort=-date&_count=1")
```

### Pull Fetal Heart Rate
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=11996-6&_sort=-date&_count=5")
```

## Clinical Guidelines
- ACOG prenatal visit schedule and lab timing (see references/acog-schedule.md)
- IOM weight gain recommendations by pre-pregnancy BMI
- USPSTF: aspirin prophylaxis for preeclampsia in high-risk patients (12-28 weeks)
- GBS prophylaxis per CDC guidelines

## Interpretation Guide
- Calculate gestational age from EDD: GA weeks = 40 - (weeks until EDD)
- Trimester determination: 1st (0-13w6d), 2nd (14-27w6d), 3rd (28-40+)
- Weight gain targets by pre-pregnancy BMI: underweight 28-40 lbs, normal 25-35 lbs, overweight 15-25 lbs, obese 11-20 lbs
- FHR normal: 110-160 bpm. Fundal height concordance: cm = GA in weeks at 20-36 weeks (discrepancy >3 cm warrants evaluation)
- Preeclampsia flags: SBP >=140 or DBP >=90 after 20 weeks on 2 occasions, proteinuria >=300 mg/24h
- GDM screening: 1-hour GCT at 24-28 weeks, positive if >=130-140 mg/dL (per institutional cutoff)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
