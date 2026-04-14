---
name: langcare-fall-risk
description: >
  Assesses fall risk using the Morse Fall Scale and medication-related fall risk
  factors from FHIR data. Identifies high-risk medications, environmental factors,
  and mobility impairments. Use when asked to assess fall risk, Morse fall scale,
  fall prevention, fall risk medications, or patient safety assessment for falls.
---

# Fall Risk Assessment

## When to Use This Skill
Use when a clinician needs to evaluate fall risk for hospitalized or ambulatory patients using validated scoring tools and medication review.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age -- risk increases >65)
2. Use `fhir_search` to pull Condition resources: fall history, gait/mobility disorders, cognitive impairment, orthostatic hypotension, visual impairment, neuropathy
3. Use `fhir_search` to pull active MedicationRequest resources; flag fall-risk medications (see references/morse-fall-scale.md): benzodiazepines, opioids, antihypertensives, antipsychotics, antidepressants, sedative-hypnotics, anticonvulsants, anticholinergics
4. Use `fhir_search` to pull Observation resources: recent vital signs (orthostatic BP), mobility assessments, cognitive screening scores
5. Apply Morse Fall Scale scoring from available data
6. Present risk stratification with specific medication and non-pharmacologic intervention recommendations

## FHIR Resources
- **Patient** -- Age
- **Condition** -- Fall history, gait disorders, cognitive impairment, visual impairment, neuropathy
- **MedicationRequest** -- Fall-risk medications
- **Observation** -- Orthostatic vitals, mobility assessments, cognitive scores, visual acuity

## FHIR Query Examples
### Pull Fall-Related Conditions
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&clinical-status=active")
```

### Pull Active Medications
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

## Clinical Guidelines
- CDC STEADI (Stopping Elderly Accidents, Deaths & Injuries) Initiative
- Morse Fall Scale for hospitalized patients
- AGS/BGS Clinical Practice Guideline for Fall Prevention (2010, updated)
- NICE CG161 Falls in Older People

## Interpretation Guide
- Morse Fall Scale: 0-24 low risk (standard precautions), 25-50 moderate risk (implement fall prevention protocol), >50 high risk (high-risk interventions)
- Count fall-risk medications: >=4 fall-risk medications is an independent risk factor
- Orthostatic hypotension: SBP drop >=20 mmHg or DBP drop >=10 mmHg on standing; correlate with symptoms
- For each high-risk medication, recommend: taper if possible, reduce dose, schedule review, or provide specific rationale for continuing

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
