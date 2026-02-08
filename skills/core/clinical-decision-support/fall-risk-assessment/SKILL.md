---
name: fall-risk-assessment
description: |
  Assesses fall risk using Morse Fall Scale, Hendrich II, and Timed Up and Go from FHIR resources and generates a fall prevention CarePlan.
  Use when user asks to "assess fall risk", "Morse fall scale", "fall prevention", "Hendrich score",
  "TUG test", "fall history", or mentions patient falls, gait instability, or balance concerns.
  Do NOT use for injury severity scoring, trauma assessment, or post-fall diagnostic imaging decisions.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: clinical-decision-support
---

# Fall Risk Assessment

## Overview

Calculate fall risk scores from FHIR Patient, Condition, Observation, and MedicationRequest resources using Morse Fall Scale, Hendrich II Fall Risk Model, and Timed Up and Go (TUG) test. Identify high-risk medications, mobility deficits, and environmental risk factors. Generate a CarePlan FHIR resource with individualized fall prevention interventions.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Age, demographics | birthDate, gender |
| Condition | Fall history, diagnoses affecting balance/gait | code, clinicalStatus, onsetDateTime |
| Observation | Vitals (orthostatic BP), mobility assessments, vision | code, valueQuantity |
| MedicationRequest | High-risk medications (sedatives, opioids, antihypertensives) | medicationCodeableConcept, status |
| Procedure | Assistive device use, recent surgeries | code, status |
| CarePlan | Output: fall prevention plan | status, category, activity, goal |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age and gender. Age >=65 is a primary fall risk factor.

### Step 2: Retrieve Fall History and Relevant Conditions

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active,recurrence,remission"
```

Key SNOMED codes:
- 161898004: Falls (history of falling)
- 217082002: Accidental fall
- 129839007: At risk for falls
- 44695005: Paralysis
- 22253000: Pain (chronic pain affecting mobility)
- 386807006: Impaired memory / cognitive impairment
- 40917007: Visual impairment
- 15188001: Hearing loss
- 85600001: Peripheral neuropathy
- 69896004: Rheumatoid arthritis
- 396275006: Osteoarthritis
- 64859006: Osteoporosis
- 230690007: Stroke (with residual deficit)
- 49049000: Parkinson disease
- 26929004: Alzheimer disease
- 84757009: Epilepsy / seizure disorder
- 271327008: Syncope
- 111516008: Urinary incontinence (urgency increases fall risk)
- 73211009: Diabetes mellitus (neuropathy risk)
- 38341003: Hypertension (orthostatic hypotension risk)
- 35489007: Depressive disorder

Also search specifically for fall history:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=161898004,217082002"
```

### Step 3: Retrieve Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

Flag medications that increase fall risk. See `references/fall-prevention.md` for comprehensive list. Key categories:
- **Sedatives/Hypnotics**: benzodiazepines (diazepam, lorazepam, alprazolam), zolpidem, zaleplon
- **Opioids**: oxycodone, hydrocodone, morphine, fentanyl, tramadol
- **Antihypertensives**: alpha-blockers (doxazosin, prazosin, terazosin), beta-blockers, diuretics, ACE inhibitors
- **Antipsychotics**: haloperidol, quetiapine, risperidone, olanzapine
- **Antidepressants**: SSRIs, SNRIs, tricyclics, trazodone
- **Anticonvulsants**: gabapentin, pregabalin, phenytoin, carbamazepine
- **Anticholinergics**: diphenhydramine, oxybutynin, promethazine
- **Muscle relaxants**: cyclobenzaprine, methocarbamol, baclofen

Count total high-risk medications. Polypharmacy (>=4 medications total or >=2 high-risk medications) is an independent fall risk factor.

### Step 4: Retrieve Orthostatic Vital Signs

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8480-6&_sort=-date&_count=5"
```

Check for orthostatic blood pressure measurements (supine, sitting, standing). Orthostatic hypotension: SBP drop >=20 mmHg or DBP drop >=10 mmHg within 3 minutes of standing.

Also retrieve heart rate:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8867-4&_sort=-date&_count=3"
```

### Step 5: Check for Assistive Device Use

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=183301007,228869008"
```
SNOMED 183301007 = Use of walking aid. 228869008 = Wheelchair use. Also check Condition or DeviceRequest resources.

### Step 6: Retrieve Functional Assessment Data

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=54614-3&_sort=-date&_count=1"
```
LOINC 54614-3 = Timed Up and Go test (TUG). If formal TUG score is documented as Observation.

Also check for gait/balance assessments:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=52732-7&_sort=-date&_count=1"
```
LOINC 52732-7 = Fall risk assessment total score (if previously documented).

### Step 7: Calculate Scores

Refer to `references/fall-scoring.md` for complete criteria.

**Morse Fall Scale (Score 0-125):**
- History of falling (within 3 months): 0 or 25
- Secondary diagnosis (>=2 medical diagnoses): 0 or 15
- Ambulatory aid: none (0), furniture/crutch/cane/walker (15), or impaired (30)
- IV/heparin lock: 0 or 20
- Gait: normal (0), weak (10), impaired (20)
- Mental status: oriented (0), overestimates/forgets limitations (15)

**Hendrich II Fall Risk Model (Score 0-16):**
- Confusion/disorientation/impulsivity: 4
- Symptomatic depression: 2
- Altered elimination: 1
- Dizziness/vertigo: 1
- Male gender: 1
- Antiepileptics administered: 2
- Benzodiazepines administered: 1
- Get Up and Go test: able to rise in single movement (0), pushes up successfully (1), multiple attempts (3), unable (4)

**Timed Up and Go (TUG):**
- Measures time to stand from chair, walk 3 meters, turn, walk back, sit
- < 12 seconds: normal mobility
- 12-20 seconds: moderate fall risk
- > 20 seconds: high fall risk, need for assistive device
- > 30 seconds: severely impaired, dependent mobility

### Step 8: Create CarePlan Resource

```
Tool: fhir_create
resourceType: "CarePlan"
resource: {
  "resourceType": "CarePlan",
  "status": "active",
  "intent": "plan",
  "category": [
    {
      "coding": [{"system": "http://snomed.info/sct", "code": "710971006", "display": "Fall prevention care plan"}]
    }
  ],
  "subject": {"reference": "Patient/[patient-id]"},
  "created": "[current-date]",
  "description": "Fall risk assessment: Morse [X]/125, Hendrich II [X]/16, TUG [X] sec. Risk level: [LOW/MODERATE/HIGH].",
  "activity": [
    {
      "detail": {
        "code": {"coding": [{"system": "http://snomed.info/sct", "code": "390994007", "display": "Fall prevention intervention"}]},
        "status": "not-started",
        "description": "[Specific intervention based on identified risk factors]"
      }
    }
  ],
  "goal": [{"reference": "Goal/[fall-prevention-goal-id]"}],
  "note": [{"text": "High-risk medications: [list]. Modifiable risk factors: [list]."}]
}
```

### Step 9: Format Output

```
FALL RISK ASSESSMENT
====================
Patient: [name] | Age: [age] | Sex: [sex]
Assessment Date: [datetime]

SCORES
------
Morse Fall Scale:  [X]/125  [LOW/MODERATE/HIGH]
Hendrich II:       [X]/16   [LOW/HIGH]
TUG:               [X] sec  [NORMAL/MODERATE/HIGH/SEVERE]

RISK FACTORS IDENTIFIED
-----------------------
Medical: [conditions affecting balance/gait/cognition]
Medications: [high-risk medications with count]
Functional: [mobility limitations, assistive device use]
Orthostatic: [BP changes if measured]

HIGH-RISK MEDICATIONS ([count])
-------------------------------
[List each with class and specific fall risk mechanism]

FALL PREVENTION CARE PLAN
=========================
1. [Intervention] -- [rationale]
2. [Intervention] -- [rationale]
3. [Intervention] -- [rationale]
...

MEDICATION REVIEW RECOMMENDATIONS
----------------------------------
[Deprescribing or substitution recommendations for high-risk medications]
```

## Examples

### Example 1: High-Risk Elderly Patient

**User says:** "Assess fall risk for patient 33221, she fell last week"

**Actions:**
1. `fhir_read` Patient/33221 -- 81F
2. `fhir_search` Condition -- history of falling (recent), Parkinson disease, osteoporosis, urinary incontinence, depression
3. `fhir_search` MedicationRequest active -- carbidopa-levodopa, sertraline, oxybutynin, lorazepam 0.5mg qHS, amlodipine, gabapentin
4. `fhir_search` Observation BP -- sitting 138/82, standing 112/68 (orthostatic drop 26 mmHg systolic)
5. Calculate: Morse = 80 (fall hx 25, secondary dx 15, walker 15, no IV 0, impaired gait 20, oriented 0 + depression adjustment 15), Hendrich II = 11
6. `fhir_create` CarePlan with fall prevention interventions

**Result:**
```
FALL RISK ASSESSMENT
====================
Patient: Margaret Liu | Age: 81 | Sex: Female

SCORES
------
Morse Fall Scale:  80/125   HIGH RISK (>=45)
Hendrich II:       11/16    HIGH RISK (>=5)
TUG:               Not documented

RISK FACTORS IDENTIFIED
-----------------------
Medical: Parkinson disease, osteoporosis, urinary incontinence, depression, recent fall
Medications: 4 high-risk medications identified
Functional: Uses walker, impaired gait (Parkinsonian), orthostatic hypotension
Orthostatic: SBP drop 26 mmHg (sitting 138 -> standing 112)

HIGH-RISK MEDICATIONS (4)
-------------------------
1. Lorazepam 0.5mg (benzodiazepine) -- sedation, impaired balance, cognitive slowing
2. Gabapentin (anticonvulsant) -- dizziness, somnolence, ataxia
3. Oxybutynin (anticholinergic) -- confusion, dizziness, blurred vision
4. Sertraline (SSRI) -- hyponatremia risk, dizziness, orthostatic hypotension

FALL PREVENTION CARE PLAN
=========================
1. Medication review -- recommend taper lorazepam, switch oxybutynin to mirabegron (non-anticholinergic)
2. Orthostatic hypotension management -- rise slowly, compression stockings, review amlodipine dose
3. Physical therapy referral -- balance and strength training, Parkinson-specific gait program
4. Home safety assessment -- remove throw rugs, install grab bars, adequate lighting
5. Hip protectors -- osteoporosis with high fall risk
6. Toileting schedule -- address incontinence-related urgency falls
7. Vitamin D supplementation -- 800-1000 IU daily if not already taking
8. Fall alarm/alert system -- wearable call device
```

### Example 2: Moderate-Risk Hospitalized Patient

**User says:** "Check fall risk on patient 55443, admitted for pneumonia"

**Actions:**
1. `fhir_read` Patient/55443 -- 68M
2. `fhir_search` Condition -- pneumonia, type 2 DM, peripheral neuropathy, BPH
3. `fhir_search` MedicationRequest active -- tamsulosin, metformin, levofloxacin, acetaminophen, IV fluids (heparin lock)
4. `fhir_search` Observation BP -- no orthostatic drop
5. Calculate: Morse = 45 (no fall hx 0, secondary dx 15, no aid 0, IV 20, weak gait 10, oriented 0), Hendrich II = 2

**Result:**
```
SCORES
------
Morse Fall Scale:  45/125   MODERATE RISK (25-44: low, >=45: high -- borderline)
Hendrich II:       2/16     LOW RISK (<5)

FALL PREVENTION CARE PLAN
=========================
1. Non-slip footwear during ambulation
2. Call light within reach, bed in lowest position
3. Assist with ambulation (weakness from acute illness)
4. Tamsulosin timing -- take at bedtime to minimize orthostatic effect
5. Night light in bathroom pathway
6. Reassess after acute illness resolves
```

## Troubleshooting

### Fall History Not Documented as Condition
- Falls may be documented in clinical notes rather than the problem list. Search for fall-related encounters:
  ```
  Tool: fhir_search
  resourceType: "Encounter"
  queryParams: "patient=[patient-id]&reason-code=217082002"
  ```
  SNOMED 217082002 = Accidental fall. Also check ICD-10 codes W00-W19 (falls). If no structured fall history exists, note "Fall history unable to be verified from structured data" and recommend asking the patient directly.

### Medication Classification Uncertain
- Some medications may not clearly map to fall-risk categories. When in doubt, check the medication class. Use the ATC classification or RxNorm hierarchy. If a medication has documented side effects of dizziness, somnolence, or orthostatic hypotension in its labeling, classify it as fall-risk. Total count of >=5 medications (any type) is an independent risk factor for falls in elderly patients (polypharmacy).

### Orthostatic Vital Signs Not Documented
- Many FHIR systems do not differentiate supine/sitting/standing BP measurements. If multiple BP readings at similar times exist, they may represent positional measurements -- check the timestamps (measurements 1-3 minutes apart). If orthostatic vitals are unavailable, recommend ordering them and flag "Orthostatic assessment pending" in the output.

## Related Skills

- `medication-reconciliation` -- for comprehensive medication review and deprescribing candidates
- `clinical-summary-generator` -- for full patient context
- `problem-list-review` -- to identify all conditions contributing to fall risk
