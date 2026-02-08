---
name: opioid-risk-assessment
description: |
  Performs opioid risk assessment by calculating morphine milligram equivalents (MME) from active prescriptions, applying Opioid Risk Tool scoring, and checking CDC 2022 guideline thresholds. Use when user asks to "assess opioid risk", "calculate MME", "check opioid safety", "review opioid prescriptions", mentions "morphine equivalents", "naloxone", "opioid tapering", or needs opioid prescribing guidance. Do NOT use for general drug interactions or non-opioid medication reviews.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: medication-management
---

# Opioid Risk Assessment

## Overview

Perform a comprehensive opioid risk assessment by pulling all active opioid prescriptions, calculating total daily morphine milligram equivalents (MME), applying Opioid Risk Tool (ORT) scoring criteria, and evaluating against CDC 2022 Clinical Practice Guideline thresholds. Flag MME >50 (caution), MME >90 (high risk), concurrent benzodiazepine use, concurrent gabapentinoid use, and absence of naloxone co-prescription. Include CAGE-AID screening questions for substance use risk assessment.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| MedicationRequest | Active opioid prescriptions with dosing | status, medicationCodeableConcept, dosageInstruction, dispenseRequest |
| MedicationStatement | Patient-reported opioid and concurrent substance use | status, medicationCodeableConcept, dosage |
| MedicationDispense | Opioid fill history for pattern analysis | status, medicationCodeableConcept, quantity, daysSupply, whenHandedOver |
| Patient | Demographics for ORT scoring (age, gender) | birthDate, gender |
| Condition | Psychiatric diagnoses, substance use history for ORT | code, clinicalStatus |
| Observation | Urine drug screen results, pain scores | code, valueCodeableConcept, valueQuantity |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age and gender. Both are required for ORT scoring (scoring criteria differ by gender).

### Step 2: Pull All Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

Classify each active medication:
- **Opioids**: Identify by RxNorm code or drug name (see references/mme-conversion-table.md)
- **Benzodiazepines**: alprazolam, clonazepam, diazepam, lorazepam, temazepam, triazolam, chlordiazepoxide, clorazepate, oxazepam, midazolam
- **Gabapentinoids**: gabapentin, pregabalin
- **Z-drugs**: zolpidem, zaleplon, eszopiclone
- **Muscle relaxants**: carisoprodol, cyclobenzaprine, methocarbamol, tizanidine
- **Naloxone**: Check if co-prescribed (nasal naloxone, injectable naloxone)

### Step 3: Calculate Total Daily MME

For each active opioid prescription:

1. Determine the opioid agent and formulation
2. Extract daily dose from `dosageInstruction`:
   - `doseAndRate[0].doseQuantity.value` = single dose amount
   - `timing.repeat.frequency` / `timing.repeat.period` = doses per day
   - `timing.code.text` for text-based frequency (e.g., "every 6 hours" = 4x/day)
3. Multiply daily dose by MME conversion factor (see references/mme-conversion-table.md)
4. Sum all opioid MME values for total daily MME

```
Total Daily MME = SUM(each opioid's daily dose x MME conversion factor)
```

### Step 4: Evaluate CDC 2022 Thresholds

| Total Daily MME | Risk Level | Action |
|----------------|------------|--------|
| < 20 | Lower risk | Standard monitoring |
| 20-49 | Moderate risk | Reassess benefits vs risks |
| 50-89 | CAUTION | Increased overdose risk. Implement risk mitigation. Prescribe naloxone. |
| >= 90 | HIGH RISK | Substantial overdose risk. Avoid increasing to this level. If already at this level, consider tapering. Prescribe naloxone. |

### Step 5: Check Concurrent High-Risk Medications

**Concurrent Benzodiazepines** (FDA Boxed Warning):
- Flag any active benzodiazepine prescription with active opioid
- Risk: respiratory depression, sedation, coma, death
- CDC recommendation: Avoid concurrent prescribing whenever possible

**Concurrent Gabapentinoids**:
- Flag gabapentin or pregabalin with active opioid
- Risk: respiratory depression (even without benzodiazepines)
- CDC recommendation: Use caution, consider lower opioid doses

**Concurrent Muscle Relaxants**:
- Flag concurrent CNS-depressant muscle relaxants
- Additive sedation and respiratory depression risk

**Concurrent Z-drugs / Sedative-Hypnotics**:
- Flag concurrent sleep medications
- Additive CNS depression

### Step 6: Apply Opioid Risk Tool (ORT)

Pull relevant history from Conditions:

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&_count=100"
```

Look for (by SNOMED code or text):
- Alcohol abuse/dependence: SNOMED 7200002, 15167005
- Illegal drug use: SNOMED 307052004
- Prescription drug abuse: SNOMED 5602001
- Depression/mental health: SNOMED 35489007, 197480006
- ADD/ADHD/OCD/bipolar/schizophrenia: multiple SNOMED codes

Apply ORT scoring (see references/ort-scoring.md). Scores differ by gender.

**ORT Interpretation**:
| Score | Risk Category | Action |
|-------|--------------|--------|
| 0-3 | Low risk | Standard monitoring |
| 4-7 | Moderate risk | Enhanced monitoring, consider UDS, shorter prescriptions |
| >= 8 | High risk | Intensive monitoring, frequent UDS, specialty referral consideration |

### Step 7: Check Urine Drug Screen (UDS) History

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=19295-5,3426-4&_sort=-date&_count=5"
```

LOINC codes:
- 19295-5: Drug screen panel (urine)
- 3426-4: Drugs of abuse screen (urine)

Check for:
- Date of most recent UDS
- Presence of prescribed opioid (confirms patient is taking medication)
- Absence of prescribed opioid (possible diversion or non-adherence)
- Presence of non-prescribed substances (concerning for polysubstance use)
- Presence of non-prescribed opioids (use of external opioid sources)

### Step 8: Check Pain Assessment Documentation

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=72514-3,38208-5&_sort=-date&_count=5"
```

LOINC codes:
- 72514-3: Pain severity (0-10 numeric rating scale)
- 38208-5: Pain assessment

Evaluate: Is pain severity being regularly documented? Are functional outcomes being tracked alongside pain scores?

### Step 9: Assess Naloxone Co-Prescription

Check if naloxone is in the active medication list (from Step 2).

**Naloxone should be co-prescribed when**:
- Total daily MME >= 50
- Concurrent benzodiazepine or other CNS depressant
- History of overdose
- History of substance use disorder
- Moderate or high ORT score
- Patient request

If naloxone is indicated but not prescribed, include recommendation.

### Step 10: CAGE-AID Screening Questions

Include these screening questions in the output for the clinician to ask during the visit:

```
CAGE-AID (Adapted to Include Drugs):
1. Have you ever felt you should CUT DOWN on your drinking or drug use?
2. Have people ANNOYED you by criticizing your drinking or drug use?
3. Have you ever felt bad or GUILTY about your drinking or drug use?
4. Have you ever had a drink or used drugs first thing in the morning to steady
   your nerves or get rid of a hangover (EYE-OPENER)?

Scoring: >= 2 "yes" answers = positive screen for substance use disorder risk
```

### Step 11: Generate Opioid Risk Assessment Report

```
OPIOID RISK ASSESSMENT - [Patient Name] (Age [X], [Gender])
Assessment Date: [date]

ACTIVE OPIOID PRESCRIPTIONS:
| Opioid | Dose | Frequency | Daily Dose | MME Factor | Daily MME |
|--------|------|-----------|-----------|------------|-----------|
| [name] | [dose] | [freq] | [daily] | [factor] | [mme] |

TOTAL DAILY MME: [sum] [RISK LEVEL: caution/high risk]

CONCURRENT HIGH-RISK MEDICATIONS:
- Benzodiazepines: [Yes/No - list if yes] [FDA BOXED WARNING]
- Gabapentinoids: [Yes/No - list if yes]
- Muscle relaxants: [Yes/No - list if yes]
- Sedative-hypnotics: [Yes/No - list if yes]

OPIOID RISK TOOL SCORE: [score] / [max] - [Low/Moderate/High Risk]
[Breakdown of scoring factors]

URINE DRUG SCREEN:
- Last UDS: [date] or NOT FOUND
- Results: [consistent/inconsistent with prescribed medications]

NALOXONE STATUS: [Prescribed / NOT PRESCRIBED - RECOMMEND]

CDC GUIDELINE COMPLIANCE:
- MME threshold: [compliant / exceeds caution / exceeds high risk]
- Concurrent CNS depressants: [compliant / non-compliant]
- UDS monitoring: [adequate / overdue / not performed]
- Naloxone co-prescribing: [compliant / non-compliant]

CAGE-AID SCREENING: [Include questions for clinician to administer]

RECOMMENDATIONS:
1. [Specific actionable recommendations based on findings]
```

## Examples

### Example 1: High-MME Patient with Concurrent Benzodiazepines

**User says**: "Assess opioid risk for patient 11223."

**Actions**:
1. Read Patient 11223: age 54, male.
2. Pull active medications: oxycodone 30mg QID (RxNorm: 7804), fentanyl patch 50mcg/hr q72h (RxNorm: 4337), lorazepam 1mg TID (RxNorm: 6470), gabapentin 600mg TID (RxNorm: 25480).
3. Calculate MME:
   - Oxycodone: 30mg x 4 = 120mg/day x 1.5 = 180 MME
   - Fentanyl patch 50mcg/hr: 50 x 2.4 = 120 MME
   - Total: 300 MME
4. Pull Conditions: chronic low back pain, lumbar radiculopathy, anxiety disorder. No documented SUD history.
5. Pull UDS: last performed 8 months ago.
6. No naloxone on medication list.
7. ORT score: Male, age 16-45 (0 for age >45), depression/anxiety (+1), no SUD history (0) = Low-Moderate.

**Result**:
```
OPIOID RISK ASSESSMENT - Patient 11223 (Age 54, Male)

ACTIVE OPIOID PRESCRIPTIONS:
| Opioid | Dose | Frequency | Daily Dose | MME Factor | Daily MME |
|--------|------|-----------|-----------|------------|-----------|
| Oxycodone IR | 30mg | QID | 120mg | 1.5 | 180 |
| Fentanyl patch | 50mcg/hr | q72h | continuous | 2.4 | 120 |

TOTAL DAILY MME: 300 [CRITICAL - EXCEEDS 90 MME THRESHOLD BY 3.3x]

CONCURRENT HIGH-RISK MEDICATIONS:
- Benzodiazepines: YES - Lorazepam 1mg TID [FDA BOXED WARNING]
- Gabapentinoids: YES - Gabapentin 600mg TID

OPIOID RISK TOOL SCORE: 2 / 26 - Low Risk (based on available diagnoses)
- Anxiety disorder: +1
- No documented SUD history: 0
- Age >45: 0

URINE DRUG SCREEN:
- Last UDS: 8 months ago [OVERDUE - recommend at least every 6 months at this MME]

NALOXONE STATUS: NOT PRESCRIBED [URGENT - RECOMMEND IMMEDIATE CO-PRESCRIPTION]

RECOMMENDATIONS:
1. [CRITICAL] Total MME 300 far exceeds CDC caution threshold (50) and high-risk threshold (90).
   Discuss gradual taper plan. Target reduction of 10% per month.
2. [CRITICAL] Concurrent opioid + benzodiazepine. FDA boxed warning for respiratory depression.
   Develop plan to taper lorazepam. Consider non-benzodiazepine anxiolytics.
3. [URGENT] Prescribe naloxone (Narcan nasal spray 4mg). Educate patient and household contacts.
4. [HIGH] Obtain urine drug screen at next visit.
5. [HIGH] Reassess functional goals. Opioid therapy at this level should demonstrate
   meaningful functional improvement.
6. [MODERATE] Consider pain management specialty referral for multimodal approach.
7. [MODERATE] Consider gabapentin taper given concurrent opioid risk.
```

### Example 2: New Opioid Prescription Risk Screening

**User says**: "Patient 55667, age 38 female, is being considered for opioid therapy for chronic pain. Run opioid risk assessment."

**Actions**:
1. Read Patient 55667: age 38, female.
2. Pull active medications: no current opioids. Acetaminophen 1000mg TID, cyclobenzaprine 10mg TID, escitalopram 20mg daily.
3. Pull Conditions: fibromyalgia, major depressive disorder, generalized anxiety disorder, history of alcohol dependence (in remission).
4. No UDS on file.
5. Calculate ORT for female:
   - Family history: unknown (0, but recommend asking)
   - Personal history alcohol dependence: +3
   - Depression: +1
   - Age 25-44: +4 (female scoring)
   - Total: 8 = HIGH RISK

**Result**:
```
OPIOID RISK ASSESSMENT - Patient 55667 (Age 38, Female)
PRE-PRESCRIBING ASSESSMENT

CURRENT OPIOID USE: None

OPIOID RISK TOOL SCORE: 8 / 26 - HIGH RISK
- Personal history of alcohol dependence: +3
- Depression/anxiety: +1
- Age 25-44 (female): +4
NOTE: Family history of substance abuse not documented - if positive, score increases further.

CONCURRENT MEDICATIONS OF CONCERN:
- Cyclobenzaprine (muscle relaxant) - would add CNS depression if opioid started
- Escitalopram (SSRI) - serotonin syndrome risk with tramadol if considered

CAGE-AID SCREENING QUESTIONS (administer before prescribing decision):
1. Have you ever felt you should CUT DOWN on your drinking or drug use?
2. Have people ANNOYED you by criticizing your drinking or drug use?
3. Have you ever felt bad or GUILTY about your drinking or drug use?
4. Have you ever had a drink or used drugs first thing in the morning (EYE-OPENER)?

RECOMMENDATIONS:
1. [HIGH] ORT score of 8 indicates HIGH RISK for opioid misuse. If opioids are prescribed:
   - Start at lowest effective dose
   - Use immediate-release only initially
   - Short supply (7-14 day prescriptions)
   - Baseline UDS before prescribing
   - UDS every 1-3 months
   - PDMP check before each prescription
   - Prescribe naloxone from the start
2. [HIGH] Administer CAGE-AID screening before decision.
3. [HIGH] Consider non-opioid alternatives first:
   - Duloxetine (FDA-approved for fibromyalgia, also treats depression/anxiety)
   - Pregabalin (FDA-approved for fibromyalgia)
   - Physical therapy
   - Cognitive behavioral therapy for chronic pain
4. [MODERATE] Obtain UDS baseline if opioids will be prescribed.
5. [MODERATE] Document informed consent discussion including addiction risk.
6. [MODERATE] If opioids started, do not exceed 50 MME without reassessment.
```

## Troubleshooting

### Opioid medication uses a combination product code (e.g., hydrocodone/acetaminophen)
- For combination products, the opioid component determines the MME. Extract the opioid strength from the medication name or dosage text.
- Example: "Hydrocodone/acetaminophen 10/325mg" - use the hydrocodone 10mg for MME calculation (10mg x 1.0 = 10 MME per dose).
- If the specific opioid strength cannot be extracted from the code, search for it in `dosageInstruction.text` or `medicationCodeableConcept.text`.

### Multiple opioid formulations make MME calculation complex
- Handle each formulation separately with its specific MME factor.
- For transdermal fentanyl, use mcg/hr x 2.4 for daily MME (this is already a per-day conversion).
- For methadone, the conversion factor varies by total daily dose (see references/mme-conversion-table.md). Do not use a single fixed factor.
- For buprenorphine transdermal or sublingual for pain, use the appropriate factor. Buprenorphine for OUD (Suboxone) is NOT included in MME calculations.

### Condition history for ORT scoring is incomplete
- Document which ORT factors could not be assessed from available FHIR data.
- Recommend the clinician complete the assessment through patient interview.
- Present the minimum and maximum possible ORT scores based on available vs missing data.

## Related Skills

- `drug-interaction-checker` - Check opioid interactions with full medication list
- `prescription-appropriateness-review` - Broader review including Beers criteria for elderly on opioids
- `medication-adherence-assessment` - Assess opioid fill patterns for early refills or gaps
- `medication-reconciliation` - Ensure all opioid sources are captured before MME calculation
