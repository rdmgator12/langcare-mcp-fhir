---
name: prescription-appropriateness-review
description: |
  Reviews prescription appropriateness using Beers Criteria (AGS 2023), STOPP/START criteria, and renal/hepatic dose adjustments. Use when user asks to "review prescriptions for appropriateness", "check Beers criteria", "check STOPP START", "review medications for elderly", mentions "anticholinergic burden", "PPI overuse", "renal dosing", or needs geriatric medication review. Do NOT use for drug interactions, adherence, or opioid-specific assessments.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: medication-management
---

# Prescription Appropriateness Review

## Overview

Perform a comprehensive prescription appropriateness review by applying AGS 2023 Beers Criteria for patients aged 65 and older, STOPP/START criteria for potentially inappropriate prescribing and prescribing omissions, renal dose adjustments based on eGFR from laboratory Observations, and hepatic dose adjustments. Flag anticholinergic burden, CNS-active polypharmacy, and prolonged PPI use.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Age determination, demographics | birthDate, gender |
| MedicationRequest | Active prescriptions to review | status, medicationCodeableConcept, dosageInstruction, reasonCode |
| MedicationStatement | Patient-reported medications including OTC | status, medicationCodeableConcept, dosage |
| Condition | Active diagnoses for drug-disease interaction checks | code, clinicalStatus, onsetDateTime |
| Observation | Lab values for renal/hepatic function | code, valueQuantity, effectiveDateTime |
| AllergyIntolerance | Cross-reference for contraindicated meds | code, clinicalStatus |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Calculate age from `birthDate`. Beers Criteria apply to patients >= 65 years. STOPP/START criteria apply to patients >= 65 years. Renal/hepatic dose adjustments apply to all ages.

### Step 2: Pull All Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

Merge and deduplicate by RxNorm code. For each medication, extract: name, dose, frequency, route, indication (if available via `reasonCode`).

### Step 3: Pull Active Conditions

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active&_count=100"
```

Extract SNOMED CT or ICD-10 codes. Key conditions to identify:
- Dementia (SNOMED: 52448006)
- Heart failure (SNOMED: 84114007)
- CKD (SNOMED: 709044004)
- Falls history (SNOMED: 217082002)
- Delirium history (SNOMED: 2776000)
- Parkinson disease (SNOMED: 49049000)
- Gastric/duodenal ulcer (SNOMED: 13200003)
- COPD (SNOMED: 13645005)
- Urinary incontinence (SNOMED: 165232002)
- Constipation (SNOMED: 14760008)
- Benign prostatic hyperplasia (SNOMED: 266569009)
- Depression (SNOMED: 35489007)
- Insomnia (SNOMED: 193462001)
- Epilepsy (SNOMED: 84757009)
- Gout (SNOMED: 90560007)

### Step 4: Pull Renal Function

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=33914-3,77147-7,48642-3&_sort=-date&_count=1"
```

LOINC codes for eGFR:
- 33914-3: eGFR (MDRD or CKD-EPI)
- 77147-7: eGFR (CKD-EPI 2021, race-free)
- 48642-3: eGFR (African American)

Also pull serum creatinine:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2160-0&_sort=-date&_count=1"
```

LOINC 2160-0: Serum creatinine (mg/dL)

### Step 5: Pull Hepatic Function (if applicable)

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=1742-6,1920-8,1975-2,14631-6&_sort=-date&_count=4"
```

LOINC codes:
- 1742-6: ALT (U/L)
- 1920-8: AST (U/L)
- 1975-2: Total bilirubin (mg/dL)
- 14631-6: Albumin (g/dL)

If ALT/AST > 3x upper limit of normal or bilirubin elevated, apply hepatic dose adjustments.

### Step 6: Apply Beers Criteria (if age >= 65)

Cross-reference each active medication against the AGS 2023 Beers Criteria (see references/beers-criteria.md). Evaluate in order:

**Table 1: Medications to Potentially Avoid in Older Adults**
Flag any medication on the Beers avoid list regardless of diagnosis.

**Table 2: Medications to Potentially Avoid Due to Drug-Disease/Drug-Syndrome Interactions**
Cross-reference active conditions against medications. Flag any condition-medication pair.

**Table 3: Medications to Use with Caution**
Flag medications requiring extra monitoring or dose limits.

**Table 4: Drug-Drug Interactions to Avoid**
Check medication pairs against Beers interaction table.

**Table 5: Medications to Avoid or Reduce Based on Kidney Function**
Apply eGFR-based restrictions.

### Step 7: Apply STOPP/START Criteria (if age >= 65)

**STOPP (Screening Tool of Older Persons' Prescriptions)**:
Cross-reference active medications and conditions against STOPP criteria (see references/stopp-start-criteria.md). Identify potentially inappropriate prescriptions.

**START (Screening Tool to Alert to Right Treatment)**:
Cross-reference active conditions against START criteria. Identify potential prescribing omissions where evidence-based medications are not prescribed.

### Step 8: Calculate Anticholinergic Burden

Assign anticholinergic burden score to each medication using the Anticholinergic Cognitive Burden (ACB) scale:

| Score | Level | Examples |
|-------|-------|---------|
| 1 | Possible anticholinergic | Furosemide, metoprolol, ranitidine, warfarin |
| 2 | Definite, clinically relevant | Amantadine, carbamazepine, cetirizine |
| 3 | Definite, strong | Amitriptyline, diphenhydramine, oxybutynin, paroxetine |

**Total ACB Score Interpretation**:
- 0-2: Acceptable
- 3-5: Moderate risk of cognitive impairment. Review for alternatives.
- 6+: High risk. Strong recommendation to reduce anticholinergic burden.

### Step 9: Assess CNS-Active Polypharmacy

Count concurrent CNS-active medications:
- Opioids
- Benzodiazepines
- Z-drugs (zolpidem, zaleplon, eszopiclone)
- Antipsychotics
- Antidepressants (SSRIs, SNRIs, TCAs)
- Antiepileptics
- Skeletal muscle relaxants
- Gabapentinoids (gabapentin, pregabalin)

**Flag if >= 3 concurrent CNS-active medications**. Associated with increased risk of falls, fractures, cognitive impairment, and mortality in older adults.

### Step 10: Check PPI Overuse

If patient is on a proton pump inhibitor (omeprazole, esomeprazole, lansoprazole, pantoprazole, rabeprazole):

1. Check duration: Pull earliest MedicationRequest or MedicationDispense date.
2. If PPI use > 8 weeks without clear indication (GERD erosive esophagitis, Barrett esophagus, Zollinger-Ellison, chronic NSAID use with risk factors), flag for deprescribing.
3. Long-term PPI risks: C. difficile infection, bone fracture, hypomagnesemia, vitamin B12 deficiency, CKD progression.

### Step 11: Apply Renal Dose Adjustments

Cross-reference each medication against renal dosing requirements (see references/renal-dosing.md). For each medication requiring adjustment:

- Compare current dose against recommended dose for patient's eGFR
- Flag if dose exceeds recommendation
- Flag if medication is contraindicated at patient's eGFR level

### Step 12: Generate Appropriateness Report

```
PRESCRIPTION APPROPRIATENESS REVIEW - [Patient Name] (Age [X])
Review Date: [date]
eGFR: [value] mL/min/1.73m2 (Date: [date])

BEERS CRITERIA FLAGS:
1. [Medication] - [Beers category] - [Rationale] - [Recommendation]
   Quality of Evidence: [High/Moderate/Low] | Strength: [Strong/Weak]

STOPP CRITERIA FLAGS:
1. [Medication] - [STOPP criterion] - [Rationale] - [Recommendation]

START CRITERIA FLAGS (Potential Omissions):
1. [Condition present] - [Recommended medication not prescribed] - [Evidence]

ANTICHOLINERGIC BURDEN:
Total ACB Score: [N]
Contributing medications: [list with individual scores]
Recommendation: [if score >= 3]

CNS-ACTIVE POLYPHARMACY:
Count: [N] concurrent CNS-active medications
Medications: [list]
Recommendation: [if >= 3]

PPI ASSESSMENT:
Duration: [months/years]
Indication documented: [Yes/No]
Recommendation: [continue/taper/discontinue]

RENAL DOSE ADJUSTMENTS NEEDED:
1. [Medication] - Current: [dose] - Recommended for eGFR [X]: [adjusted dose]

HEPATIC CONSIDERATIONS:
[if applicable]

SUMMARY:
- Total flags: [N]
- Critical (requires immediate action): [N]
- Moderate (review at next visit): [N]
- Informational: [N]
```

## Examples

### Example 1: Geriatric Appropriateness Review

**User says**: "Review prescription appropriateness for patient 78901, age 82, with dementia."

**Actions**:
1. Read Patient 78901. Confirm age 82.
2. Pull active medications: donepezil 10mg, metoprolol 50mg BID, diphenhydramine 25mg QHS, oxybutynin 5mg BID, omeprazole 20mg daily, tramadol 50mg TID, zolpidem 10mg QHS, gabapentin 300mg TID, escitalopram 20mg daily, aspirin 81mg daily.
3. Pull active conditions: Alzheimer dementia, overactive bladder, GERD, chronic low back pain, insomnia, depression.
4. Pull eGFR: 42 mL/min (Stage 3b CKD).

**Result**:
```
PRESCRIPTION APPROPRIATENESS REVIEW - Patient 78901 (Age 82)
eGFR: 42 mL/min/1.73m2

BEERS CRITERIA FLAGS:
1. Diphenhydramine 25mg - TABLE 1: Avoid - Highly anticholinergic, cognitive impairment risk
   [Strong recommendation, Moderate evidence]
2. Oxybutynin 5mg BID - TABLE 2: Avoid with dementia - Worsens cognitive impairment
   [Strong recommendation, Moderate evidence]
3. Zolpidem 10mg - TABLE 1: Avoid - Falls, delirium, fractures. If used, max 5mg.
   [Strong recommendation, Moderate evidence]
4. Tramadol 50mg TID - TABLE 1: Use with caution - CNS effects, falls risk, seizure threshold
5. Omeprazole - TABLE 3: Use with caution >8 weeks without indication review

STOPP CRITERIA FLAGS:
1. Oxybutynin with dementia (STOPP B7) - Anticholinergic worsens cognition
2. Benzodiazepine-receptor agonist (zolpidem) with falls risk (STOPP K1)
3. First-generation antihistamine (diphenhydramine) long-term (STOPP B1)

ANTICHOLINERGIC BURDEN:
Total ACB Score: 8 [CRITICAL]
- Diphenhydramine: 3
- Oxybutynin: 3
- Escitalopram: 1
- Metoprolol: 1
Recommendation: URGENT reduction. Target ACB <= 2. Discontinue diphenhydramine and oxybutynin.

CNS-ACTIVE POLYPHARMACY:
Count: 5 [CRITICAL]
- Tramadol, Zolpidem, Gabapentin, Escitalopram, Diphenhydramine
Recommendation: Reduce to <= 2 CNS-active medications.

RENAL DOSE ADJUSTMENTS:
1. Gabapentin: Current 300mg TID. For eGFR 42: Max 300mg BID. Reduce.
2. Tramadol: Extended-release contraindicated if eGFR <30. Currently ok but monitor.

RECOMMENDED ACTIONS:
1. [URGENT] Discontinue diphenhydramine. Use melatonin 1-3mg for sleep if needed.
2. [URGENT] Discontinue oxybutynin. Consider mirabegron (beta-3 agonist, non-anticholinergic).
3. [HIGH] Taper zolpidem. If continued, reduce to 5mg max.
4. [HIGH] Reduce gabapentin to 300mg BID for CKD.
5. [MODERATE] Review tramadol necessity. Consider topical alternatives.
6. [MODERATE] Reassess omeprazole indication for deprescribing.
```

### Example 2: START Criteria - Identifying Missing Medications

**User says**: "Check if patient 34567, age 72 with diabetes and heart failure, is missing any recommended medications."

**Actions**:
1. Read Patient 34567. Age 72.
2. Pull active medications: metformin 1000mg BID, glipizide 10mg BID, furosemide 40mg daily, carvedilol 12.5mg BID, aspirin 81mg daily.
3. Pull active conditions: Type 2 diabetes (SNOMED: 44054006), heart failure with reduced EF (SNOMED: 84114007), hypertension, hyperlipidemia.
4. Pull eGFR: 58 mL/min, HbA1c: 7.8%, LDL: 135 mg/dL, EF: 35% (if available).

**Result**:
```
PRESCRIPTION APPROPRIATENESS REVIEW - Patient 34567 (Age 72)

START CRITERIA FLAGS (Potential Omissions):
1. ACE inhibitor or ARB - Not prescribed. START A3: ACE/ARB indicated for HFrEF.
   Also START A5: ACE/ARB for diabetic nephropathy (check urine albumin).
2. Statin - Not prescribed. START A3: Statin indicated for diabetes age 40-75.
   LDL 135 mg/dL is above goal (<70 for diabetes + HF).
3. SGLT2 inhibitor - Not prescribed. START: SGLT2i (dapagliflozin, empagliflozin) indicated
   for HFrEF regardless of diabetes status. Dual benefit here.
4. Mineralocorticoid receptor antagonist (spironolactone/eplerenone) - Not prescribed.
   Indicated for HFrEF with EF <=35% (if EF confirmed at 35%).

BEERS CRITERIA:
1. Glipizide - TABLE 1: Avoid long-acting sulfonylureas in elderly. Hypoglycemia risk.
   Consider: Switch to short-acting or SGLT2i/GLP-1 RA.

RECOMMENDED ACTIONS:
1. [HIGH] Add ACE inhibitor (e.g., lisinopril 5mg, titrate) or ARB. Guideline-directed for HFrEF.
2. [HIGH] Add high-intensity statin (atorvastatin 40-80mg or rosuvastatin 20-40mg).
3. [HIGH] Add SGLT2 inhibitor (empagliflozin 10mg or dapagliflozin 10mg). Check eGFR >= 20.
4. [MODERATE] Add spironolactone 25mg if K+ <5.0. Monitor K+ in 1 week.
5. [MODERATE] Consider switching glipizide to DPP-4 inhibitor or GLP-1 RA (lower hypoglycemia risk).
```

## Troubleshooting

### Patient age is under 65 - Beers and STOPP/START do not apply
- Skip Beers and STOPP/START sections. Still apply renal/hepatic dose adjustments, anticholinergic burden assessment, CNS-active polypharmacy check, and PPI overuse assessment. These are age-independent safety checks.
- Note in the report that geriatric criteria were not applied due to age.

### eGFR not available in Observation records
- Check for serum creatinine (LOINC: 2160-0) and calculate eGFR using CKD-EPI 2021 formula:
  - eGFR = 142 x min(SCr/kappa, 1)^alpha x max(SCr/kappa, 1)^-1.200 x 0.9938^age x (1.012 if female)
  - kappa = 0.7 (female), 0.9 (male); alpha = -0.241 (female), -0.302 (male)
- If no creatinine available, note that renal dose adjustments could not be assessed and recommend obtaining baseline labs.

### Medications use local codes instead of RxNorm
- Fall back to text-based matching against Beers/STOPP criteria drug names.
- Search for medication class indicators in the text (e.g., "-pril" for ACE inhibitors, "-sartan" for ARBs, "-olol" for beta-blockers).
- Note reduced confidence in the assessment when coded matching is unavailable.

## Related Skills

- `drug-interaction-checker` - Complementary interaction analysis beyond Beers Table 4
- `medication-reconciliation` - Ensure medication list is complete before appropriateness review
- `opioid-risk-assessment` - Detailed opioid assessment when opioids flagged by Beers
- `medication-adherence-assessment` - Assess if patient is actually taking the flagged medications
