---
name: langcare-mental-health
description: >
  Performs mental health screening using PHQ-9 (depression), GAD-7 (anxiety),
  AUDIT-C (alcohol use), and Columbia Suicide Severity Rating Scale from FHIR
  Observation data. Tracks symptom trends and treatment response. Use when
  asked for mental health screening, depression assessment, PHQ-9 score,
  GAD-7 score, anxiety assessment, or behavioral health evaluation.
---

# Mental Health Screening

## When to Use This Skill
Use when a clinician needs to evaluate mental health screening scores, track symptom trends over time, and assess treatment response.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics
2. Use `fhir_search` to pull screening Observation resources: PHQ-9 (LOINC 44249-1), GAD-7 (LOINC 70274-6), AUDIT-C (LOINC 75626-2), PHQ-2 (LOINC 55758-7), Columbia CSSRS (LOINC 93373-2)
3. Pull historical screening scores for trend analysis
4. Use `fhir_search` to pull active Condition resources for mental health diagnoses
5. Use `fhir_search` to pull active MedicationRequest for psychotropic medications (antidepressants, anxiolytics, mood stabilizers, antipsychotics)
6. Score interpretation per validated scales (see references/phq9-gad7-scoring.md)
7. Assess treatment response: compare current scores to baseline and post-medication initiation
8. Flag urgent findings: PHQ-9 item 9 (suicidal ideation) positive, severe scores, worsening trajectory

## FHIR Resources
- **Observation** -- PHQ-9 (44249-1), GAD-7 (70274-6), AUDIT-C (75626-2), PHQ-2 (55758-7), CSSRS (93373-2)
- **Condition** -- Mental health diagnoses (MDD, GAD, PTSD, bipolar, SUD)
- **MedicationRequest** -- Psychotropic medications
- **Patient** -- Demographics

## FHIR Query Examples
### Pull PHQ-9 Trend
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=44249-1&_sort=date&_count=20")
```

### Pull GAD-7 Scores
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=70274-6&_sort=date&_count=20")
```

### Pull AUDIT-C
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=75626-2&_sort=-date&_count=5")
```

## Clinical Guidelines
- USPSTF Grade B: depression screening for all adults
- USPSTF Grade B: anxiety screening for adults
- USPSTF Grade B: unhealthy alcohol use screening
- APA Practice Guidelines for Major Depressive Disorder
- SAMHSA screening and brief intervention guidelines

## Interpretation Guide
- PHQ-9: 0-4 minimal, 5-9 mild, 10-14 moderate, 15-19 moderately severe, 20-27 severe. Treatment threshold: >=10
- GAD-7: 0-4 minimal, 5-9 mild, 10-14 moderate, 15-21 severe. Treatment threshold: >=10
- AUDIT-C: positive screen >=3 (women), >=4 (men)
- PHQ-9 Item 9 (suicidal ideation): ANY positive response requires immediate safety assessment
- Treatment response: >=50% reduction in PHQ-9 = response; score <5 = remission
- If PHQ-9 not improving after 4-8 weeks of adequate therapy, consider dose adjustment, augmentation, or referral

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
