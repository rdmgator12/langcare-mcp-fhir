# Fall Risk Scoring Systems Reference

## Morse Fall Scale (MFS)

Most widely used fall risk assessment tool in acute care settings. Score range: 0-125.

### Scoring Criteria

| Item | Criterion | Points | How to Assess from FHIR |
|------|-----------|--------|------------------------|
| 1 | **History of falling** (within 3 months, or current fall event) | 0 = No, 25 = Yes | Condition: SNOMED 161898004, 217082002. Check onsetDateTime within 90 days. |
| 2 | **Secondary diagnosis** (>=2 medical diagnoses documented) | 0 = No, 15 = Yes | Count active Condition resources. >=2 active conditions = Yes. |
| 3 | **Ambulatory aid** | 0 = None / bed rest / nurse assist | Check Procedure/DeviceRequest for assistive devices. |
|   |   | 15 = Crutches / cane / walker | SNOMED 183301007 (walking aid), 465565008 (walking frame), 36700001 (cane) |
|   |   | 30 = Furniture walking (holds onto furniture to ambulate) | Clinical observation -- may not be in structured data |
| 4 | **IV therapy / heparin lock** | 0 = No, 20 = Yes | Check active MedicationAdministration for IV route, or active IV access |
| 5 | **Gait** | 0 = Normal / bed rest / immobile | Clinical assessment or documented mobility status |
|   |   | 10 = Weak (stooped but can lift head, poor balance, needs support) | Condition: gait abnormality SNOMED 22325002, balance impairment |
|   |   | 20 = Impaired (cannot rise without assistance, very poor balance, short steps, shuffling, lurching) | SNOMED 250043000 (shuffling gait), Parkinson gait, ataxic gait |
| 6 | **Mental status** | 0 = Oriented to own ability (knows limitations) | Check for cognitive conditions |
|   |   | 15 = Overestimates or forgets limitations | SNOMED 386807006 (cognitive impairment), 26929004 (Alzheimer), 386806002 (impaired cognition) |

### Risk Level Interpretation

| Score | Risk Level | Action |
|-------|-----------|--------|
| 0-24 | Low risk | Standard fall precautions |
| 25-44 | Moderate risk | Implement fall prevention interventions |
| >= 45 | High risk | Implement high-risk fall prevention protocol |

### Standard Fall Precautions (All Patients)
- Orient to room and call light
- Bed in lowest position with brakes locked
- Non-slip footwear
- Personal items within reach
- Adequate lighting

### High-Risk Interventions (Score >= 45)
- Fall risk identification (wristband, door sign, chart flag)
- Bed alarm / chair alarm
- 1:1 sitter consideration if score >75 or cognitively impaired
- Toileting schedule (every 2 hours)
- Hourly rounding
- PT/OT consultation
- Medication review for fall-risk medications
- Low bed or floor mat
- Remove clutter and obstacles
- Hip protectors if osteoporosis

---

## Hendrich II Fall Risk Model (HIIFRM)

Validated in acute care. Faster to administer than Morse. Score range: 0-16.

### Scoring Criteria

| Risk Factor | Points | How to Assess from FHIR |
|-------------|--------|------------------------|
| **Confusion / Disorientation / Impulsivity** | 4 | Condition: SNOMED 386807006 (cognitive impairment), 130987000 (acute confusion), 26929004 (Alzheimer). MedicationRequest for antipsychotics may indicate impulsivity/agitation. |
| **Symptomatic depression** | 2 | Condition: SNOMED 35489007 (depressive disorder), 370143000 (major depression). MedicationRequest for antidepressants. PHQ-9 score documented as Observation. |
| **Altered elimination** (urgency, frequency, incontinence, nocturia) | 1 | Condition: SNOMED 111516008 (urinary incontinence), 162116003 (urinary frequency), 139394000 (nocturia). MedicationRequest for overactive bladder meds. |
| **Dizziness / vertigo** | 1 | Condition: SNOMED 404640003 (dizziness), 399153001 (vertigo). Active medications with dizziness as common side effect. |
| **Male gender** | 1 | Patient.gender = "male" |
| **Antiepileptics administered** | 2 | MedicationRequest: phenytoin, carbamazepine, valproate, levetiracetam, lamotrigine, topiramate, gabapentin, pregabalin, lacosamide. Active status. |
| **Benzodiazepines administered** | 1 | MedicationRequest: diazepam, lorazepam, alprazolam, clonazepam, midazolam, temazepam, triazolam. Active status. |
| **Get Up and Go test** | 0 = Able to rise in single movement (no loss of balance) | Clinical assessment |
|   | 1 = Pushes up, successful in one attempt | -- |
|   | 3 = Multiple attempts but successful | -- |
|   | 4 = Unable to rise without assistance | -- |

### Risk Level Interpretation

| Score | Risk Level | Action |
|-------|-----------|--------|
| 0-4 | Low risk | Standard precautions |
| >= 5 | High risk | Fall prevention protocol. Every additional point = 10% increased fall risk. |

### Advantages Over Morse
- Includes medication risk factors directly in scoring
- Get Up and Go component directly assesses functional mobility
- No subjective "mental status" criterion -- uses specific conditions
- Faster to calculate (7 items vs 6 items with multi-level scoring)

---

## Timed Up and Go (TUG) Test

Functional mobility assessment. Measures time (in seconds) to:
1. Stand up from standard arm chair
2. Walk 3 meters (10 feet) at normal pace
3. Turn around
4. Walk back to chair
5. Sit down

### Requirements
- Standard arm chair (seat height ~46 cm / 18 inches)
- Measured 3-meter (10-foot) distance marked on floor
- Stopwatch
- Patient wears regular footwear, uses usual walking aid if applicable

### Interpretation

| Time | Risk Level | Interpretation |
|------|-----------|---------------|
| < 10 seconds | Normal | Freely mobile, low fall risk |
| 10-12 seconds | Borderline | Mostly independent, monitor |
| 12-20 seconds | Moderate risk | Some mobility impairment, may benefit from intervention |
| 20-30 seconds | High risk | Significant impairment, needs assistive device and fall prevention |
| > 30 seconds | Very high risk | Severely impaired, dependent mobility, consider wheelchair assessment |

### LOINC Code
- 54614-3: Timed Up and Go test result

### Clinical Context
- Best used in community-dwelling elderly and outpatient settings
- Less applicable in acute inpatient (patients may be at non-baseline function)
- Serial measurements useful for tracking mobility changes over time
- Validated in: community elderly, Parkinson disease, stroke rehabilitation, hip fracture recovery

### Combined Scoring Approach

For comprehensive assessment, use all three tools:
1. **Morse Fall Scale**: Best for acute care inpatients (captures IV, ambulatory aid specifics)
2. **Hendrich II**: Best for medication-related risk and cognitive factors
3. **TUG**: Best for functional mobility and community-dwelling patients

If any one score indicates HIGH risk, classify the patient as HIGH risk overall.

---

## Age-Related Fall Risk Thresholds

| Age Group | Annual Fall Rate | Fall-Related Injury Rate | Notes |
|-----------|-----------------|-------------------------|-------|
| 65-74 | ~28-35% | 10-15% | Screen annually |
| 75-84 | ~35-45% | 15-20% | Screen at every visit |
| >= 85 | ~50%+ | 20-30% | Highest risk, most likely to result in serious injury |

### Screening Triggers (When to Perform Full Assessment)
- Age >= 65 at annual visit
- Any reported fall in past 12 months
- Gait or balance concern expressed by patient, family, or clinician
- New prescription of high-risk medication
- After acute illness or hospitalization
- Change in functional status
- New diagnosis affecting mobility (stroke, Parkinson, fracture)
