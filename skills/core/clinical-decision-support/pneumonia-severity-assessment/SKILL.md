---
name: pneumonia-severity-assessment
description: |
  Assesses pneumonia severity using CURB-65, PSI/PORT Score, and A-DROP criteria from FHIR resources with disposition and treatment recommendations.
  Use when user asks to "assess pneumonia severity", "CURB-65 score", "PORT score", "PSI score", "pneumonia risk class",
  "pneumonia disposition", "CAP severity", or needs help determining inpatient vs outpatient pneumonia treatment.
  Do NOT use for pneumonia diagnosis, chest X-ray interpretation, or chronic lung disease management.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: clinical-decision-support
---

# Pneumonia Severity Assessment

## Overview

Calculate pneumonia severity scores from FHIR Patient, Observation, and Condition resources using CURB-65, PSI/PORT Score (Pneumonia Severity Index), and A-DROP criteria. Differentiate community-acquired pneumonia (CAP) from hospital-acquired pneumonia (HAP) and ventilator-associated pneumonia (VAP). Generate disposition recommendations (outpatient, observation, inpatient, ICU) and empiric antibiotic guidance based on severity level and guidelines. Create a ClinicalImpression resource documenting findings.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Age, gender, nursing home residence | birthDate, gender, address |
| Condition | Comorbidities, pneumonia type, pleural effusion | code, clinicalStatus, onsetDateTime |
| Observation | Vitals (BP, HR, RR, temp, SpO2), labs (BUN, glucose, Na, Hct, PaO2, pH) | code, valueQuantity, effectiveDateTime |
| Encounter | Admission source, current encounter type | class, period, hospitalization |
| MedicationRequest | Current antibiotics, immunosuppressants | medicationCodeableConcept, status |
| ClinicalImpression | Output: severity assessment | status, description, finding |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age (exact years), gender (male = additional PSI points). Check address for nursing home/long-term care facility (affects CAP vs HCAP classification).

### Step 2: Determine CAP vs HAP/VAP

```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&status=in-progress"
```

Classification logic:
- **CAP**: Pneumonia onset before or within 48 hours of hospital admission. Includes nursing home patients (previously "HCAP" -- per 2019 ATS/IDSA guidelines, HCAP category eliminated; treat as CAP unless HAP/VAP risk factors present).
- **HAP**: Pneumonia onset >=48 hours after hospital admission, not intubated at time of onset.
- **VAP**: Pneumonia onset >=48 hours after endotracheal intubation.

Check Encounter.period.start against pneumonia Condition.onsetDateTime. Check for active Procedure with SNOMED 40617009 (mechanical ventilation).

### Step 3: Retrieve Active Conditions and Comorbidities

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active,recurrence,remission"
```

PSI-relevant conditions (SNOMED codes):
- 363346000: Neoplastic disease (malignancy)
- 235856003: Hepatic disease (liver disease)
- 42343007: Congestive heart failure
- 230690007: Cerebrovascular disease
- 709044004: Chronic kidney disease / renal disease
- 233604007: Pneumonia (the current diagnosis)
- 60046008: Pleural effusion

Also search for pleural effusion specifically:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=60046008"
```

### Step 4: Retrieve Vital Signs

**Respiratory Rate:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=9279-1&_sort=-date&_count=3"
```
LOINC 9279-1. CURB-65: >=30. PSI: >=30. A-DROP: >=30.

**Systolic Blood Pressure:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8480-6&_sort=-date&_count=3"
```
LOINC 8480-6. CURB-65: SBP <90 or DBP <=60. PSI: SBP <90.

**Diastolic Blood Pressure:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8462-4&_sort=-date&_count=3"
```
LOINC 8462-4. CURB-65: DBP <=60.

**Heart Rate:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8867-4&_sort=-date&_count=3"
```
LOINC 8867-4. PSI: >=125. A-DROP does not use HR directly.

**Temperature:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8310-5&_sort=-date&_count=3"
```
LOINC 8310-5. PSI: <35C or >=40C.

**SpO2:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2708-6&_sort=-date&_count=3"
```
LOINC 2708-6. A-DROP: SpO2 <=90% or PaO2 <=60 mmHg. ATS/IDSA severe CAP criterion.

### Step 5: Retrieve Laboratory Values

**BUN (Blood Urea Nitrogen):**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=3094-0&_sort=-date&_count=1"
```
LOINC 3094-0. CURB-65: BUN >20 mg/dL (>7 mmol/L). PSI: BUN >=30 mg/dL. A-DROP: BUN >=21 mg/dL.

**Serum Sodium:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2951-2&_sort=-date&_count=1"
```
LOINC 2951-2. PSI: Na <130 mEq/L.

**Serum Glucose:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2345-7&_sort=-date&_count=1"
```
LOINC 2345-7. PSI: glucose >=250 mg/dL.

**Hematocrit:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=4544-3&_sort=-date&_count=1"
```
LOINC 4544-3. PSI: Hct <30%.

**Arterial pH:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2744-1&_sort=-date&_count=1"
```
LOINC 2744-1. PSI: pH <7.35.

**PaO2:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2703-7&_sort=-date&_count=1"
```
LOINC 2703-7. PSI: PaO2 <60 mmHg. A-DROP: PaO2 <=60 mmHg.

### Step 6: Assess Mental Status

Search for documented confusion or altered mental status:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=9269-2&_sort=-date&_count=1"
```
LOINC 9269-2 = GCS total. Also check for documented Condition of confusion (SNOMED 130987000) or altered mental status (SNOMED 419284004).

CURB-65: new mental confusion (defined as AMT <=8 or new disorientation to person, place, or time).
PSI: altered mental status (disorientation, stupor, coma).
A-DROP: confusion (new onset).

### Step 7: Calculate Scores

Refer to `references/pneumonia-scoring.md` for complete criteria.

**CURB-65 (Score 0-5):**
- C: Confusion (new onset)
- U: Urea (BUN) >20 mg/dL (>7 mmol/L)
- R: Respiratory rate >=30
- B: Blood pressure SBP <90 or DBP <=60
- 65: Age >=65

**PSI/PORT Score:** Demographics + comorbidities + exam findings + labs = point total mapped to Risk Class I-V.

**A-DROP (Score 0-5):**
- A: Age (male >=70, female >=75)
- D: Dehydration (BUN >=21 mg/dL)
- R: Respiration (SpO2 <=90% or PaO2 <=60 mmHg)
- O: Orientation (confusion)
- P: Pressure (SBP <=90 mmHg)

### Step 8: Evaluate ATS/IDSA Severe CAP Criteria

Check for major criteria (either = ICU admission):
1. Septic shock requiring vasopressors
2. Respiratory failure requiring mechanical ventilation

Check for minor criteria (>=3 = ICU admission):
1. RR >= 30
2. PaO2/FiO2 <= 250
3. Multilobar infiltrates
4. Confusion/disorientation
5. BUN >= 20 mg/dL
6. WBC < 4,000
7. Platelets < 100,000
8. Temperature < 36C
9. Hypotension requiring aggressive fluid resuscitation

### Step 9: Create ClinicalImpression Resource

```
Tool: fhir_create
resourceType: "ClinicalImpression"
resource: {
  "resourceType": "ClinicalImpression",
  "status": "completed",
  "subject": {"reference": "Patient/[patient-id]"},
  "effectiveDateTime": "[current-datetime]",
  "description": "Pneumonia severity: CURB-65 [X]/5, PSI Class [I-V] ([X] pts), A-DROP [X]/5. Type: [CAP/HAP/VAP].",
  "finding": [
    {
      "itemCodeableConcept": {
        "coding": [{"system": "http://snomed.info/sct", "code": "233604007", "display": "Pneumonia"}],
        "text": "Pneumonia type: [CAP/HAP/VAP]. Severity: [MILD/MODERATE/SEVERE]. Disposition: [outpatient/observation/inpatient/ICU]."
      }
    }
  ],
  "note": [{"text": "Recommended empiric therapy: [antibiotic regimen]. ATS/IDSA severe criteria: [X] minor criteria met."}]
}
```

### Step 10: Format Output

```
PNEUMONIA SEVERITY ASSESSMENT
==============================
Patient: [name] | Age: [age] | Sex: [sex]
Assessment Date: [datetime]
Pneumonia Type: [CAP / HAP / VAP]

SCORES
------
CURB-65:  [X]/5  - [disposition recommendation]
  C (Confusion): [Y/N]
  U (BUN >20): [Y/N] ([value])
  R (RR >=30): [Y/N] ([value])
  B (BP low): [Y/N] (SBP [value], DBP [value])
  65 (Age >=65): [Y/N]

PSI/PORT: Class [I-V] ([X] points) - [mortality risk]
  Demographics: [points breakdown]
  Comorbidities: [points breakdown]
  Exam: [points breakdown]
  Labs: [points breakdown]

A-DROP:   [X]/5  - [severity level]

ATS/IDSA SEVERE CAP CRITERIA
-----------------------------
Major criteria: [X]/2  Minor criteria: [X]/9
ICU recommended: [YES/NO]

DISPOSITION RECOMMENDATION
--------------------------
[Outpatient / Observation / General ward / ICU]
Rationale: [based on scoring concordance]

EMPIRIC ANTIBIOTIC RECOMMENDATION
----------------------------------
[Regimen based on severity, setting, and risk factors]
```

## Examples

### Example 1: Moderate Severity CAP

**User says:** "Score pneumonia severity for patient 22334"

**Actions:**
1. `fhir_read` Patient/22334 -- 72M
2. `fhir_search` Encounter -- admitted 6 hours ago via ED
3. `fhir_search` Condition -- CAP, COPD, diabetes. No liver/renal/heart failure/cancer.
4. `fhir_search` Observation vitals -- RR 24, SBP 105, DBP 62, HR 98, Temp 38.9C, SpO2 93%
5. `fhir_search` Observation labs -- BUN 28, Na 136, Glucose 210, Hct 38%, pH 7.38, PaO2 65
6. Mental status: oriented, no confusion
7. Calculate: CURB-65 = 2 (BUN >20 = 1, Age >=65 = 1), PSI = 102 (age 72 + male 0 adjustment... age 72, COPD +10, DM +10, BUN >=30 not met at 28, actually BUN <30 = 0 for PSI... recalculate: age 72, COPD +10, DM +10 = 92 base + exam/lab points), A-DROP = 2 (Age M>=70, BUN >=21)
8. ATS/IDSA: 0 major, 1 minor (BUN >=20)

**Result:**
```
PNEUMONIA SEVERITY ASSESSMENT
==============================
Patient: Harold Kim | Age: 72 | Sex: Male
Pneumonia Type: CAP (admitted via ED, onset pre-admission)

SCORES
------
CURB-65:  2/5  - Consider hospital admission (short stay or observation)
PSI/PORT: Class IV (92 points) - Mortality 8.2%
A-DROP:   2/5  - Moderate severity

ATS/IDSA SEVERE CAP CRITERIA
-----------------------------
Major criteria: 0/2  Minor criteria: 1/9
ICU recommended: NO

DISPOSITION RECOMMENDATION
--------------------------
General ward admission (medical floor with telemetry)
Rationale: CURB-65 2 supports admission. PSI Class IV supports inpatient. No ICU criteria met. SpO2 93% requires supplemental O2.

EMPIRIC ANTIBIOTIC RECOMMENDATION
----------------------------------
Non-severe inpatient CAP (no Pseudomonas risk):
- Ceftriaxone 2g IV daily + Azithromycin 500mg IV daily
- OR Respiratory fluoroquinolone: Levofloxacin 750mg IV daily (if beta-lactam allergy)
- Reassess at 48-72 hours for clinical improvement and step-down to oral.
```

### Example 2: Severe CAP Requiring ICU

**User says:** "CURB-65 and PORT score for patient 88776, she's in the ED with bad pneumonia"

**Actions:**
1. `fhir_read` Patient/88776 -- 68F
2. `fhir_search` Condition -- CAP (multilobar), CHF, CKD stage 3
3. `fhir_search` Observation vitals -- RR 34, SBP 78, DBP 45, HR 122, Temp 35.2C, SpO2 84%
4. `fhir_search` Observation labs -- BUN 42, Na 128, Glucose 180, Hct 28%, pH 7.28, PaO2 52
5. Mental status: confused, disoriented
6. Calculate: CURB-65 = 5/5, PSI Class V (>130 pts), A-DROP = 5/5
7. ATS/IDSA: 0 major (not yet intubated/on vasopressors), 7 minor (RR >=30, PaO2/FiO2 <250, multilobar, confusion, BUN >=20, Plt check, Temp <36)

**Result:**
```
SCORES
------
CURB-65:  5/5  - HIGH RISK (mortality ~57%, ICU admission)
PSI/PORT: Class V (168 points) - Mortality 27-31%
A-DROP:   5/5  - EXTREMELY SEVERE

ATS/IDSA SEVERE CAP CRITERIA
-----------------------------
Major criteria: 0/2  Minor criteria: 7/9
ICU recommended: YES (>=3 minor criteria met)

DISPOSITION RECOMMENDATION
--------------------------
ICU ADMISSION
Rationale: All scores at maximum severity. 7 ATS/IDSA minor criteria met. Impending respiratory failure (SpO2 84%, PaO2 52). Hemodynamically unstable (SBP 78). Prepare for possible intubation and vasopressors.

EMPIRIC ANTIBIOTIC RECOMMENDATION
----------------------------------
Severe CAP, ICU admission:
- Ceftriaxone 2g IV daily + Azithromycin 500mg IV daily + Vancomycin 25mg/kg IV (MRSA coverage given severity)
- If Pseudomonas risk: Piperacillin-tazobactam 4.5g IV q6h + Levofloxacin 750mg IV daily
- Obtain sputum culture, blood cultures x2, Legionella/Pneumococcal urinary antigens before antibiotics if possible.
```

## Troubleshooting

### BUN Available Only as Urea (Non-US Systems)
- Some FHIR servers report urea in mmol/L (LOINC 3091-6) instead of BUN in mg/dL (LOINC 3094-0). Conversion: BUN (mg/dL) = Urea (mmol/L) x 2.8. CURB-65 threshold: BUN >20 mg/dL = Urea >7 mmol/L. Search both:
  ```
  Tool: fhir_search
  resourceType: "Observation"
  queryParams: "patient=[patient-id]&code=3094-0,3091-6&_sort=-date&_count=1"
  ```

### Arterial Blood Gas Not Available
- If PaO2 and pH are unavailable (no ABG drawn), use SpO2 as surrogate for oxygenation assessment. SpO2 <=90% on room air approximates PaO2 <=60 mmHg. Note "ABG not available -- oxygenation assessed by SpO2 only" in output. pH component of PSI cannot be scored without ABG -- note as "pH unavailable, PSI score may underestimate severity."

### Nursing Home Status Not in Patient Resource
- Check Patient.address for long-term care facility indicators. Also check Encounter.hospitalization.origin for transfer from skilled nursing facility. If unclear, flag "Nursing home/LTCF status unknown -- assumed community-dwelling for scoring purposes."

## Related Skills

- `sepsis-screening` -- pneumonia is a leading cause of sepsis; run sepsis screening if CURB-65 >=3 or hemodynamic instability
- `clinical-summary-generator` -- for full patient context
- `medication-reconciliation` -- to verify current antibiotic therapy and potential interactions
