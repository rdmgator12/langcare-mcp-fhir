---
name: chronic-pain-management-review
description: |
  Performs comprehensive chronic pain management assessment with opioid safety review, MME calculation,
  and multimodal treatment evaluation. Use when user asks to "review pain management", "check MME",
  "opioid review", "pain assessment", "pain medication review", "morphine equivalent", mentions
  "chronic pain", "pain score trends", "urine drug screen", or needs pain regimen safety evaluation.
  Do NOT use for acute post-surgical pain, cancer pain (use oncology-treatment-timeline), or headache-only evaluations.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: specialty
---

# Chronic Pain Management Review

## Overview

Pull and analyze all pain-related clinical data for a comprehensive chronic pain management review. Retrieve pain-related Conditions, current analgesic MedicationRequests (opioid and non-opioid), pain score Observations over time, functional status assessments, and urine drug screen results. Calculate current daily morphine milligram equivalents (MME). Evaluate adherence to CDC 2022 opioid prescribing guidelines. Flag high-risk patterns: elevated MME, concurrent benzodiazepines, lack of multimodal approach, missing urine drug screens, absent naloxone prescription. Present structured review with risk mitigation recommendations.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Demographics, age | birthDate, gender |
| Condition | Pain diagnoses, comorbidities | code, clinicalStatus, onsetDateTime, bodySite |
| MedicationRequest | Analgesics, adjuvants, naloxone | medicationCodeableConcept, status, dosageInstruction, authoredOn |
| Observation | Pain scores, UDS, functional status | code, valueQuantity, effectiveDateTime, interpretation |
| Procedure | Pain interventions (injections, nerve blocks) | code, performedDateTime, status |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age, gender. Age is relevant for Beers criteria (>=65 years: avoid opioids as first-line per AGS).

### Step 2: Pull Pain-Related Conditions

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code:below=22253000&clinical-status=active"
```

SNOMED 22253000 = Pain (finding). This captures chronic pain subtypes.

Also search for specific chronic pain conditions:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=82423001,279039007,203082005,431855005,73211009,724637000&clinical-status=active"
```

SNOMED codes:
- 82423001 = Chronic pain
- 279039007 = Low back pain
- 203082005 = Fibromyalgia
- 431855005 = Chronic widespread pain
- 73211009 = Diabetic neuropathy
- 724637000 = Complex regional pain syndrome

For each Condition:
- Extract `code` (pain type), `bodySite`, `onsetDateTime` (duration of chronic pain)
- Note if neuropathic vs nociceptive vs mixed (affects treatment approach -- see references/pain-management-framework.md)

### Step 3: Pull All Current Analgesic Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

Categorize medications into:

**Opioids** (calculate MME for each):
- Codeine (RxNorm 2670): MME factor 0.15
- Hydrocodone (RxNorm 5489): MME factor 1
- Hydromorphone (RxNorm 3423): MME factor 4
- Methadone (RxNorm 6813): MME factor varies by dose (see references/opioid-safety.md)
- Morphine (RxNorm 7052): MME factor 1 (reference)
- Oxycodone (RxNorm 7804): MME factor 1.5
- Oxymorphone (RxNorm 7814): MME factor 3
- Tramadol (RxNorm 10689): MME factor 0.1
- Fentanyl transdermal (RxNorm 4337): patch mcg/hr * 2.4 = daily MME
- Tapentadol (RxNorm 787390): MME factor 0.4
- Buprenorphine (RxNorm 1819): for pain, MME factor 12.6 (transdermal) or 10 (sublingual)

**Non-Opioid Analgesics**:
- NSAIDs: ibuprofen, naproxen, meloxicam, celecoxib, diclofenac
- Acetaminophen
- Muscle relaxants: cyclobenzaprine, baclofen, tizanidine, methocarbamol

**Adjuvant Analgesics**:
- Gabapentinoids: gabapentin (RxNorm 25480), pregabalin (RxNorm 190823)
- Antidepressants: duloxetine (RxNorm 72625), amitriptyline (RxNorm 704), nortriptyline (RxNorm 7531), venlafaxine (RxNorm 39786)
- Topicals: lidocaine patch, capsaicin, diclofenac gel
- Muscle relaxants for spasticity: baclofen, tizanidine

**High-Risk Combinations to Flag**:
- Benzodiazepines: alprazolam, lorazepam, clonazepam, diazepam (RxNorm 596, 6470, 2598, 3322)
- Carisoprodol (Schedule IV muscle relaxant with abuse potential)
- Z-drugs: zolpidem, eszopiclone (sedation stacking)

### Step 4: Calculate Total Daily MME

For each opioid MedicationRequest:
1. Extract dose from `dosageInstruction[0].doseAndRate[0].doseQuantity`
2. Extract frequency from `dosageInstruction[0].timing.repeat` (period, periodUnit, frequency)
3. Calculate daily dose: single dose * doses per day
4. Multiply by MME conversion factor
5. Sum all opioid MMEs for total daily MME

CDC 2022 thresholds:
- <50 MME/day: standard risk
- 50-89 MME/day: increased risk, reassess benefit vs harm
- >=90 MME/day: high risk, avoid or justify with documented rationale

See references/opioid-safety.md for complete MME conversion table and methadone variable conversion.

### Step 5: Pull Pain Score Observations

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=72514-3,38208-5,77565-0&_sort=-date&_count=50"
```

LOINC codes:
- 72514-3 = Pain severity (0-10 NRS)
- 38208-5 = Pain intensity (0-10 NRS)
- 77565-0 = PEG pain scale (Pain, Enjoyment, General activity)

Build pain score time series. Calculate:
- Current pain score
- Average pain score over last 3 months
- Pain score trend (improving, stable, worsening)
- Correlation with medication changes

### Step 6: Pull Functional Status Assessments

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=77565-0,89555-7&_sort=-date&_count=10"
```

- LOINC 77565-0 = PEG scale (Pain, Enjoyment of life, General activity -- 3 items, 0-10 each)
- LOINC 89555-7 = Brief Pain Inventory interference

Functional improvement is a key metric -- pain reduction without functional improvement suggests suboptimal management.

### Step 7: Pull Urine Drug Screen Results

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=3426-4,19659-2,12300-3&_sort=-date&_count=10"
```

LOINC codes (common UDS panel):
- 3426-4 = Opiates screen
- 19659-2 = Benzodiazepines screen
- 12300-3 = Cocaine metabolite
- 19270-8 = Amphetamines screen
- 3416-5 = THC (cannabinoids) screen
- 3773-9 = Methadone screen
- 16369-5 = Fentanyl screen
- 19585-9 = Buprenorphine screen

Expected UDS findings for chronic opioid therapy:
- Prescribed opioid: should be POSITIVE (adherence)
- Non-prescribed substances: should be NEGATIVE
- Unexpected positive: investigate (recreational use, diversion concern)
- Unexpected negative for prescribed opioid: investigate (non-adherence, diversion)

Flag:
- No UDS in past 12 months (should be at least annual per CDC)
- Unexpected positive results
- Missing expected positive for prescribed opioid

### Step 8: Pull Pain Interventional Procedures

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=231253001,399097000,18946005&_sort=-date"
```

SNOMED codes:
- 231253001 = Nerve block
- 399097000 = Epidural injection
- 18946005 = Trigger point injection
- 14490009 = Joint injection
- SNOMED 261551009 = Radiofrequency ablation
- CPT 64493-64495 = Facet joint injections

Also search for physical therapy, chiropractic, acupuncture:
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=91251008,44868003,231785002&_sort=-date"
```

SNOMED: 91251008 = Physical therapy, 44868003 = Chiropractic, 231785002 = Acupuncture.

### Step 9: Check Naloxone Prescription

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&code=7242&status=active"
```

RxNorm 7242 = naloxone. CDC recommends naloxone co-prescription if MME >=50, history of overdose, concurrent benzodiazepines, or substance use disorder history.

### Step 10: Check PDMP-Related Observations

Search for controlled substance agreement and PDMP check documentation:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=PDMP&_sort=-date&_count=5"
```

Note: PDMP data may not be available via FHIR. Flag if no documentation of PDMP check exists.

### Step 11: Generate Pain Management Review

```
CHRONIC PAIN MANAGEMENT REVIEW -- [Patient Name] -- [Date]
=============================================================

PAIN DIAGNOSES
  1. [Condition] -- onset [date] -- duration [years/months]
  Pain type: [nociceptive / neuropathic / mixed / nociplastic]
  Body site: [location]

CURRENT PAIN REGIMEN
  Opioids:
    [Drug] [dose] [frequency] = [daily dose] = [MME] MME/day
    ...
  Total Daily MME: [sum] MME/day [<50 / 50-89 / >=90 risk level]

  Non-Opioid Analgesics:
    [Drug] [dose] [frequency]
    ...

  Adjuvant Medications:
    [Drug] [dose] [frequency]
    ...

MULTIMODAL ASSESSMENT
  [x] Opioid therapy
  [x/o] Non-opioid pharmacotherapy
  [x/o] Adjuvant analgesics (neuropathic agents)
  [x/o] Interventional procedures
  [x/o] Physical therapy / rehabilitation
  [x/o] Behavioral/psychological therapy
  [x/o] Complementary approaches (acupuncture, massage)

PAIN SCORE TREND
  [Date]: [score] -> [Date]: [score] -> [Date]: [score]
  Average (3 months): [value]/10
  Trend: [improving / stable / worsening]

FUNCTIONAL STATUS
  PEG Score: [value]/30 (Pain [x], Enjoyment [x], General activity [x])
  Trend: [improving / stable / declining]

RISK MITIGATION STATUS
  [x/o] Controlled substance agreement on file
  [x/o] PDMP checked within last [X] months
  [x/o] Urine drug screen: [date] -- [results summary]
  [x/o] Naloxone prescribed [required if MME >=50 or concurrent BZD]

SAFETY FLAGS
  [!] [Description of any safety concern]
  ...

RECOMMENDATIONS
  [Specific actions based on review findings]
```

## Examples

### Example 1: Stable Chronic Pain Patient on Opioids

**User says:** "Review pain management for patient PAIN-2201"

**Actions:**
1. `fhir_read` Patient/PAIN-2201 -- 54-year-old male
2. `fhir_search` Condition?patient=PAIN-2201&code:below=22253000&clinical-status=active -- chronic low back pain (2019), lumbar degenerative disc disease
3. `fhir_search` MedicationRequest?patient=PAIN-2201&status=active -- oxycodone 10mg q8h, gabapentin 600mg TID, meloxicam 15mg daily, naloxone nasal spray
4. `fhir_search` Observation?patient=PAIN-2201&code=72514-3&_sort=-date&_count=50 -- pain scores: 7, 6, 5, 5, 5 over 6 months
5. `fhir_search` Observation?patient=PAIN-2201&code=3426-4,19659-2&_sort=-date&_count=10 -- UDS 2 months ago: positive opiates (expected), negative benzodiazepines, negative cocaine, negative amphetamines
6. `fhir_search` Procedure?patient=PAIN-2201&_sort=-date -- lumbar epidural 3 months ago, PT completed 6 months ago

**Result:**
```
CHRONIC PAIN MANAGEMENT REVIEW -- David Thompson -- 2026-02-07
================================================================

PAIN DIAGNOSES
  1. Chronic low back pain -- onset 2019 -- 7 years
  2. Lumbar degenerative disc disease -- L4-L5, L5-S1
  Pain type: Mixed (nociceptive + neuropathic component)

CURRENT PAIN REGIMEN
  Opioids:
    Oxycodone 10mg q8h (30mg/day) = 45 MME/day
  Total Daily MME: 45 MME/day [BELOW 50 threshold -- standard risk]

  Non-Opioid: Meloxicam 15mg daily
  Adjuvant: Gabapentin 600mg TID (neuropathic component)

MULTIMODAL ASSESSMENT
  [x] Opioid therapy
  [x] Non-opioid pharmacotherapy (NSAID)
  [x] Adjuvant analgesic (gabapentin)
  [x] Interventional (epidural steroid injection, 3 months ago)
  [x] Physical therapy (completed 6-month course)
  [o] Behavioral/psychological therapy -- NOT DOCUMENTED
  [o] Complementary approaches -- NOT DOCUMENTED

PAIN SCORE TREND
  2025-08: 7 -> 2025-10: 6 -> 2025-12: 5 -> 2026-01: 5 -> 2026-02: 5
  Average (3 months): 5.0/10
  Trend: IMPROVED then STABLE

RISK MITIGATION STATUS
  [x] UDS: 2025-12-10 -- expected opioid positive, all others negative [CONCORDANT]
  [x] Naloxone prescribed (nasal spray on file)
  [?] Controlled substance agreement -- not found in FHIR records
  [?] PDMP check -- not documented in FHIR records

ASSESSMENT: Well-managed chronic pain regimen with appropriate multimodal approach.
  MME below 50 threshold. UDS concordant. Pain improving/stable.

RECOMMENDATIONS
  - Continue current regimen
  - Consider CBT referral for chronic pain (behavioral therapy gap)
  - Document PDMP check and controlled substance agreement
  - Next UDS due by 2026-06 (annual minimum)
  - Re-assess epidural benefit at next visit
```

### Example 2: High-Risk Opioid Prescribing Pattern

**User says:** "Pain management review for patient PAIN-5590, high MME concern"

**Actions:**
1. `fhir_read` Patient/PAIN-5590 -- 42-year-old female
2. `fhir_search` Condition?patient=PAIN-5590&code:below=22253000&clinical-status=active -- chronic pain syndrome, fibromyalgia, anxiety disorder
3. `fhir_search` MedicationRequest?patient=PAIN-5590&status=active -- oxycodone ER 40mg q12h, oxycodone IR 10mg q4h PRN, alprazolam 1mg TID, acetaminophen 500mg PRN
4. Calculate MME: oxycodone ER 80mg/day * 1.5 = 120 MME + oxycodone IR up to 60mg/day * 1.5 = 90 MME = up to 210 MME/day
5. `fhir_search` Observation?patient=PAIN-5590&code=72514-3&_sort=-date&_count=50 -- pain scores stable at 7-8/10 despite high MME
6. `fhir_search` Observation?patient=PAIN-5590&code=3426-4,19659-2&_sort=-date -- last UDS was 14 months ago
7. `fhir_search` MedicationRequest?patient=PAIN-5590&code=7242 -- NO naloxone on file

**Result:**
```
CHRONIC PAIN MANAGEMENT REVIEW -- Jennifer Adams -- 2026-02-07
================================================================

[!] MULTIPLE SAFETY FLAGS IDENTIFIED

PAIN DIAGNOSES
  1. Chronic pain syndrome -- onset 2020 -- 6 years
  2. Fibromyalgia (SNOMED 203082005)
  3. Generalized anxiety disorder (comorbid)
  Pain type: Nociplastic (fibromyalgia) + nociceptive

CURRENT PAIN REGIMEN
  Opioids:
    Oxycodone ER 40mg q12h (80mg/day) = 120 MME/day
    Oxycodone IR 10mg q4h PRN (up to 60mg/day) = up to 90 MME/day
  Total Daily MME: 120-210 MME/day [EXCEEDS 90 MME THRESHOLD -- HIGH RISK]

  Concurrent sedative: Alprazolam 1mg TID [CONCURRENT BENZODIAZEPINE]
  Non-Opioid: Acetaminophen 500mg PRN (inadequate as sole non-opioid)

SAFETY FLAGS
  [!] MME >= 90: Total 120-210 MME/day -- significantly exceeds CDC threshold
  [!] CONCURRENT BENZODIAZEPINE: Alprazolam 1mg TID with high-dose opioids -- CDC strongly advises against
  [!] NO NALOXONE PRESCRIBED: Required at >=50 MME and concurrent benzodiazepine
  [!] UDS OVERDUE: Last UDS 14 months ago (should be at minimum annual, more frequent with risk factors)
  [!] PAIN NOT CONTROLLED: Scores 7-8/10 despite 120+ MME -- suggests opioid-induced hyperalgesia or inadequate approach
  [!] NO ADJUVANT THERAPY: Fibromyalgia responds poorly to opioids; duloxetine/pregabalin not prescribed
  [!] NO INTERVENTIONAL OR NON-PHARMACOLOGIC THERAPY documented

RECOMMENDATIONS
  1. IMMEDIATE: Prescribe naloxone nasal spray and provide overdose education
  2. IMMEDIATE: Obtain UDS and PDMP check
  3. URGENT: Develop benzodiazepine taper plan -- concurrent use with high-dose opioids is highest-risk combination
  4. Add fibromyalgia-specific therapy: duloxetine or pregabalin (first-line for fibromyalgia per ACR)
  5. Develop opioid taper plan targeting <90 MME initially, with goal of further reduction
  6. Refer to pain management specialist for multimodal approach
  7. Refer to behavioral health for anxiety management (alternative to benzodiazepine) and chronic pain CBT
  8. Consider physical therapy and exercise program referral
```

## Troubleshooting

### MME Calculation Unclear Due to PRN Dosing

- For PRN medications, calculate maximum daily MME based on prescribed frequency (e.g., q4h = up to 6 doses/day).
- Report as range: "[minimum scheduled MME] to [maximum with all PRN doses] MME/day."
- Check `dosageInstruction.maxDosePerDay` if populated for prescriber-defined daily limits.
- If frequency is `asNeeded` without timing, estimate conservatively at 4 doses/day and note the assumption.

### Urine Drug Screen Results Not Found Under Standard LOINC Codes

- UDS may be coded under panel LOINC 49569-7 (Drug abuse screening) or as individual drug-class Observations.
- Try broader search: `fhir_search` Observation?patient=[id]&category=laboratory&code:text=drug+screen&_sort=-date.
- Some systems store UDS as DiagnosticReport with component results. Search DiagnosticReport?patient=[id]&code:text=drug+screen.
- If no UDS found, flag as "No urine drug screen documented" rather than assuming one was never done.

### Pain Scores Stored Under Different Codes

- Some systems use LOINC 38208-5 (pain intensity) instead of 72514-3 (pain severity). Try both.
- Pain scores may appear in vital signs category rather than survey. Try: `category=vital-signs&code=72514-3`.
- If no discrete pain Observations exist, note "No structured pain scores available -- recommend implementing standardized pain documentation."

## Related Skills

- `opioid-risk-assessment` -- for detailed opioid risk stratification (ORT, SOAPP-R)
- `medication-reconciliation` -- for complete medication review
- `lab-result-interpreter` -- for UDS result interpretation and hepatic/renal function monitoring
- `mental-health-screening` -- for depression/anxiety screening in chronic pain patients
