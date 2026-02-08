---
name: mental-health-screening
description: |
  Implements validated mental health screening tools (PHQ-9, GAD-7, AUDIT-C, C-SSRS, MDQ, PC-PTSD-5)
  with scoring, interpretation, and safety planning. Use when user asks to "screen for depression",
  "anxiety screening", "PHQ-9 score", "GAD-7", "suicide risk assessment", "AUDIT-C", "bipolar screening",
  "PTSD screen", "mental health assessment", or needs behavioral health screening documentation.
  Do NOT use for psychiatric medication management, therapy notes, or detailed psychiatric evaluations.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: specialty
---

# Mental Health Screening

## Overview

Administer, score, and document validated mental health screening instruments using FHIR QuestionnaireResponse and Observation resources. Supported tools: PHQ-2 (brief depression), PHQ-9 (depression severity), GAD-7 (generalized anxiety), AUDIT-C (alcohol use), Columbia Suicide Severity Rating Scale (C-SSRS), Mood Disorder Questionnaire (MDQ, bipolar), and PC-PTSD-5 (PTSD). Calculate scores, classify severity, generate clinical recommendations by score range, and trigger safety assessment workflows for positive suicidality screens. Record results as FHIR Observations with standardized LOINC codes.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Demographics, age | birthDate, gender |
| QuestionnaireResponse | Screening instrument responses | questionnaire, item, answer, authored |
| Observation | Screening scores and interpretations | code, valueQuantity, interpretation, effectiveDateTime |
| Condition | Existing mental health diagnoses | code, clinicalStatus, onsetDateTime |
| MedicationRequest | Current psychotropic medications | medicationCodeableConcept, status, dosageInstruction |
| RiskAssessment | Suicide risk documentation | prediction, method, mitigation |

## Instructions

### Step 1: Retrieve Patient Demographics and History

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age (screening tools may have age-specific versions) and gender.

Pull existing mental health conditions:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&category=encounter-diagnosis&code=35489007,197480006,724689000,66590003,47505003,313182004&clinical-status=active"
```

SNOMED codes:
- 35489007 = Depressive disorder
- 197480006 = Anxiety disorder
- 724689000 = Alcohol use disorder
- 66590003 = Bipolar disorder
- 47505003 = PTSD
- 313182004 = Suicidal ideation

### Step 2: Pull Prior Screening Results

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=44249-1,69737-5,75626-2,77564-3,89206-7&_sort=-date&_count=10"
```

LOINC codes for screening scores:
- 44249-1 = PHQ-9 total score
- 69737-5 = PHQ-2 total score
- 75626-2 = GAD-7 total score
- 77564-3 = AUDIT-C total score
- 89206-7 = C-SSRS (Columbia suicide severity)

Extract prior scores and dates for trend analysis.

### Step 3: Pull Current Psychotropic Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

Identify psychotropic medications:
- SSRIs: sertraline, fluoxetine, escitalopram, citalopram, paroxetine
- SNRIs: venlafaxine, duloxetine, desvenlafaxine
- Benzodiazepines: alprazolam, lorazepam, clonazepam, diazepam
- Mood stabilizers: lithium, valproate, lamotrigine, carbamazepine
- Antipsychotics: quetiapine, aripiprazole, olanzapine, risperidone
- Stimulants: methylphenidate, amphetamine salts
- Bupropion, mirtazapine, trazodone, buspirone

Note current medications to inform treatment recommendations.

### Step 4: Determine Which Screening Tool to Administer

Selection logic:
- **Universal screening**: PHQ-2 first; if score >=3, proceed to PHQ-9
- **Depression concern**: PHQ-9 directly
- **Anxiety concern**: GAD-7
- **Alcohol use concern**: AUDIT-C
- **Suicidal ideation**: C-SSRS (always if PHQ-9 item 9 > 0)
- **Bipolar concern**: MDQ (before starting antidepressant, if mood cycling history)
- **Trauma history**: PC-PTSD-5

If user does not specify, default to PHQ-2 + GAD-7 as standard primary care screening.

### Step 5: Score the Screening Instrument

See references/screening-tools.md for complete item-level scoring. Summary scoring:

**PHQ-2** (items 1-2 of PHQ-9, range 0-6):
- Score >=3: positive screen, proceed to full PHQ-9

**PHQ-9** (9 items, range 0-27):
- 0-4: minimal/none
- 5-9: mild
- 10-14: moderate
- 15-19: moderately severe
- 20-27: severe
- Item 9 (suicidal ideation) > 0: trigger C-SSRS safety assessment

**GAD-7** (7 items, range 0-21):
- 0-4: minimal
- 5-9: mild
- 10-14: moderate
- 15-21: severe

**AUDIT-C** (3 items, range 0-12):
- Men: score >=4 positive
- Women: score >=3 positive

**C-SSRS** (severity scale 1-5):
- 1: Wish to be dead
- 2: Non-specific active suicidal thoughts
- 3: Suicidal ideation with method (no plan or intent)
- 4: Suicidal ideation with intent (no specific plan)
- 5: Suicidal ideation with specific plan and intent

**MDQ** (13 yes/no items + 2 supplemental):
- Positive screen: >=7 "yes" items AND items occurred at same time AND caused moderate/serious problems

**PC-PTSD-5** (5 yes/no items, range 0-5):
- Score >=3: positive screen, proceed to full PTSD evaluation

### Step 6: Record Screening Score as Observation

```
Tool: fhir_create
resourceType: "Observation"
resource: {
  "resourceType": "Observation",
  "status": "final",
  "category": [{"coding": [{"system": "http://terminology.hl7.org/CodeSystem/observation-category", "code": "survey", "display": "Survey"}]}],
  "code": {"coding": [{"system": "http://loinc.org", "code": "[LOINC-code]", "display": "[instrument-name] total score"}]},
  "subject": {"reference": "Patient/[patient-id]"},
  "effectiveDateTime": "[current-datetime]",
  "valueQuantity": {"value": [score], "unit": "{score}"},
  "interpretation": [{"coding": [{"system": "http://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation", "code": "[N/A/H]", "display": "[severity]"}]}]
}
```

### Step 7: Record QuestionnaireResponse (if item-level data available)

```
Tool: fhir_create
resourceType: "QuestionnaireResponse"
resource: {
  "resourceType": "QuestionnaireResponse",
  "status": "completed",
  "subject": {"reference": "Patient/[patient-id]"},
  "authored": "[current-datetime]",
  "item": [
    {"linkId": "1", "text": "[question text]", "answer": [{"valueInteger": [score]}]},
    ...
  ]
}
```

### Step 8: Safety Assessment for Positive Suicidality Screen

If PHQ-9 item 9 > 0 or C-SSRS positive:

1. Determine imminent risk level:
   - **Imminent risk** (C-SSRS 4-5, or active plan/intent): immediate psychiatric evaluation, do not leave patient unattended, contact crisis team
   - **Elevated risk** (C-SSRS 2-3, PHQ-9 item 9 = 2-3): same-day mental health evaluation, safety plan
   - **Low risk** (C-SSRS 1, PHQ-9 item 9 = 1): outpatient mental health referral within 1 week, safety plan

2. Document safety assessment:
```
Tool: fhir_create
resourceType: "Observation"
resource: {
  "resourceType": "Observation",
  "status": "final",
  "code": {"coding": [{"system": "http://loinc.org", "code": "89206-7", "display": "Columbia suicide severity rating scale"}]},
  "subject": {"reference": "Patient/[patient-id]"},
  "effectiveDateTime": "[current-datetime]",
  "valueQuantity": {"value": [severity-level], "unit": "{score}"},
  "note": [{"text": "Safety assessment completed. [Plan details]."}]
}
```

See references/mental-health-management.md for safety planning template and crisis resources.

### Step 9: Generate Recommendations by Score

Present structured output:

```
MENTAL HEALTH SCREENING RESULTS -- [Patient Name] -- [Date]
=============================================================

SCREENING INSTRUMENT: [Name]
Score: [value] / [max] -- Severity: [classification]
Prior score: [value] on [date] -- Trend: [improved/stable/worsened]

CURRENT PSYCHOTROPIC MEDICATIONS
  [list or "None"]

INTERPRETATION
  [Clinical interpretation based on score and context]

RECOMMENDATIONS
  [Specific actions per severity level -- see references/screening-tools.md]

SAFETY STATUS
  Suicidal ideation screen: [Negative / Positive -- details]
  [If positive: safety plan status, risk level, disposition]
```

## Examples

### Example 1: Routine Depression and Anxiety Screening

**User says:** "Run depression and anxiety screening for patient MH-1120"

**Actions:**
1. `fhir_read` Patient/MH-1120 -- 34-year-old female
2. `fhir_search` Condition?patient=MH-1120&category=encounter-diagnosis&code=35489007,197480006&clinical-status=active -- generalized anxiety disorder (active since 2024)
3. `fhir_search` Observation?patient=MH-1120&code=44249-1,75626-2&_sort=-date&_count=10 -- prior PHQ-9: 12 (moderate, 3 months ago), prior GAD-7: 14 (moderate, 3 months ago)
4. `fhir_search` MedicationRequest?patient=MH-1120&status=active -- sertraline 100mg daily
5. Score current PHQ-9 = 8 (mild), GAD-7 = 10 (moderate). PHQ-9 item 9 = 0 (no suicidal ideation).
6. `fhir_create` Observation for PHQ-9 score (LOINC 44249-1, value 8)
7. `fhir_create` Observation for GAD-7 score (LOINC 75626-2, value 10)

**Result:**
```
MENTAL HEALTH SCREENING RESULTS -- Sarah Mitchell -- 2026-02-07
================================================================

PHQ-9: 8/27 -- Mild depression
  Prior: 12 (moderate) on 2025-11-07 -- IMPROVED
  Item 9 (suicidal ideation): 0 -- NEGATIVE

GAD-7: 10/21 -- Moderate anxiety
  Prior: 14 (moderate) on 2025-11-07 -- IMPROVED (borderline)

CURRENT MEDICATIONS: sertraline 100mg daily (started 2024)

INTERPRETATION
  Depression improving from moderate to mild on current SSRI.
  Anxiety remains moderate but trending downward.
  No safety concerns.

RECOMMENDATIONS
  - Continue sertraline 100mg. Consider dose optimization for residual anxiety.
  - Rescreen in 4-6 weeks.
  - Consider CBT referral for anxiety management.
  - Discuss sleep hygiene, exercise, stress management.
```

### Example 2: Positive Suicidality Screen with Safety Assessment

**User says:** "PHQ-9 screening for patient MH-4450, he endorsed suicidal thoughts"

**Actions:**
1. `fhir_read` Patient/MH-4450 -- 52-year-old male
2. `fhir_search` Condition?patient=MH-4450&clinical-status=active -- major depressive disorder, alcohol use disorder
3. `fhir_search` MedicationRequest?patient=MH-4450&status=active -- bupropion 300mg daily
4. Score PHQ-9 = 22 (severe). Item 9 = 2 (more than half the days).
5. Administer C-SSRS: severity level 3 (suicidal ideation with method, no plan or intent).
6. `fhir_create` Observation for PHQ-9 score (value 22, interpretation HH)
7. `fhir_create` Observation for C-SSRS (value 3)

**Result:**
```
MENTAL HEALTH SCREENING RESULTS -- Robert Johnson -- 2026-02-07
================================================================

[SAFETY ALERT] POSITIVE SUICIDALITY SCREEN

PHQ-9: 22/27 -- SEVERE depression
  Item 9 (suicidal ideation): 2 -- endorsed "more than half the days"

C-SSRS: Severity Level 3 -- Suicidal ideation WITH METHOD (no plan or intent)
  Patient reports thinking about overdose but denies specific plan or intent.

RISK FACTORS IDENTIFIED
  - Severe depression (PHQ-9 = 22)
  - Active alcohol use disorder
  - Male sex, age >45
  - Has identified method

PROTECTIVE FACTORS
  - Denies specific plan or intent
  - Currently engaged in treatment
  - On antidepressant (bupropion)

IMMEDIATE ACTIONS REQUIRED
  1. Do not leave patient unattended until safety plan completed
  2. Same-day psychiatric evaluation recommended
  3. Lethal means counseling -- assess access to medications, firearms
  4. Complete safety plan (see references/mental-health-management.md)
  5. Provide crisis resources: 988 Suicide and Crisis Lifeline

TREATMENT RECOMMENDATIONS
  - Current bupropion may be insufficient -- psychiatric consultation for medication adjustment
  - Screen for alcohol use severity (AUDIT-C recommended)
  - Intensive outpatient or partial hospitalization may be appropriate
  - Follow-up within 48-72 hours if discharged
```

## Troubleshooting

### QuestionnaireResponse Resource Not Supported by Server

- Fall back to creating Observation resources only with the total score and LOINC code.
- Document individual item responses in `Observation.note[].text` as structured text.
- Some servers accept QuestionnaireResponse but require a canonical Questionnaire URL. Use standard URLs: `http://loinc.org/vs/LL3342-0` for PHQ-9, `http://loinc.org/vs/LL3318-0` for GAD-7.

### Prior Screening Scores Not Found

- Scores may be stored in different Observation categories. Try searching without category filter: `patient=[id]&code=44249-1&_sort=-date`.
- Some systems store screening results as part of encounter notes (DocumentReference) rather than discrete Observations. Note that trend analysis is unavailable.
- Check if DiagnosticReport contains embedded screening results.

### RiskAssessment Resource Not Supported

- Document suicide risk assessment entirely within Observation resources using LOINC 89206-7 (C-SSRS).
- Add clinical detail in `Observation.note[]` including risk level, protective factors, and disposition plan.
- If no structured resource is available, clearly present the safety assessment in the output for clinical documentation.

## Related Skills

- `medication-reconciliation` -- for reviewing psychotropic medication lists
- `clinical-summary-generator` -- for comprehensive summary including behavioral health
- `follow-up-task-generator` -- for scheduling follow-up screening and safety check-ins
- `referral-generator` -- for generating psychiatry or therapy referrals
