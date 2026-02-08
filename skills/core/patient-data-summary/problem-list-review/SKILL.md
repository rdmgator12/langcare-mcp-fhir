---
name: problem-list-review
description: |
  Reviews the active problem list by cross-referencing Conditions with medications, encounters, and labs to identify
  discrepancies. Use when user asks to "review the problem list", "audit diagnoses", "check conditions",
  "problem list reconciliation", "coding review", "are there missing diagnoses", or "what conditions are untreated".
  Do NOT use for allergy review, medication reconciliation, or generating full clinical summaries.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: patient-data
---

# Problem List Review

## Overview

Perform a structured audit of the active problem list by pulling Condition resources and cross-referencing with MedicationRequests, recent Encounters, and Observations. Identify clinical discrepancies: conditions without active treatment, resolved conditions still listed as active, medications implying undocumented conditions, and ICD-10 coding specificity opportunities. Output actionable findings for clinical review.

## FHIR Resources Used

| Resource | Purpose | Search Parameters |
|----------|---------|-------------------|
| Condition | Active problem list | patient, clinical-status |
| MedicationRequest | Cross-reference treatments | patient, status=active |
| Encounter | Recent visit context | patient, date |
| Observation | Lab correlation | patient, category=laboratory, code |

## Instructions

### Step 1: Pull All Conditions

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]"
```

Retrieve all conditions regardless of status. Categorize each by:
- `clinicalStatus`: active, recurrence, relapse, inactive, remission, resolved
- `verificationStatus`: confirmed, unconfirmed, provisional, differential, refuted, entered-in-error
- `category`: problem-list-item, encounter-diagnosis, health-concern
- `code.coding`: Extract SNOMED CT code and ICD-10-CM code if present

Build two lists:
1. **Active conditions** (clinicalStatus = active, recurrence, or relapse)
2. **Inactive/resolved conditions** (clinicalStatus = inactive, remission, or resolved)

### Step 2: Pull Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

For each medication, extract:
- Drug name and RxNorm code from `medicationCodeableConcept`
- Drug class (infer from name suffix or RxNorm hierarchy)
- Dosage from `dosageInstruction`

### Step 3: Pull Recent Encounters (Last 12 Months)

```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&date=ge[12-months-ago]&_sort=-date"
```

Extract encounter diagnoses from `reasonCode` or associated `Condition` references with `category` = encounter-diagnosis.

### Step 4: Cross-Reference Medications Against Problems

For each active medication, check if a corresponding condition exists on the active problem list. Use the medication-condition mapping reference.

**Flag Type A -- Medication Without Matching Condition (Implied Diagnosis Missing):**

| Medication / Class | Implied Condition | Check Problem List For |
|--------------------|-------------------|----------------------|
| Metformin, glipizide, insulin | Type 2 Diabetes Mellitus | SNOMED 44054006 or ICD-10 E11.x |
| Lisinopril, losartan, amlodipine | Hypertension | SNOMED 59621000 or ICD-10 I10 |
| Atorvastatin, rosuvastatin | Hyperlipidemia | SNOMED 55822004 or ICD-10 E78.x |
| Levothyroxine | Hypothyroidism | SNOMED 40930008 or ICD-10 E03.x |
| Albuterol, fluticasone inhaler | Asthma or COPD | SNOMED 195967001 or 13645005 |
| Sertraline, fluoxetine, escitalopram | Depression or Anxiety | SNOMED 35489007 or 197480006 |
| Omeprazole, pantoprazole | GERD | SNOMED 235595009 or ICD-10 K21.0 |
| Warfarin, apixaban, rivaroxaban | A-fib, DVT/PE, or valve condition | SNOMED 49436004 |
| Gabapentin, pregabalin | Neuropathy, seizure disorder, or chronic pain | SNOMED 386033004 |
| Methotrexate | Rheumatoid arthritis or psoriasis | SNOMED 69896004 |

If medication is present but no matching condition exists on the active problem list, flag as: "Patient is on [medication] but [condition] is not on the active problem list."

### Step 5: Check for Conditions Without Active Treatment

For each active condition, check if there is at least one active medication or recent encounter addressing it.

**Flag Type B -- Untreated Active Condition:**
- Active condition with no medication from the expected drug class
- No encounter in the last 12 months with the condition as a reason code
- Exception: conditions that may not require pharmacotherapy (e.g., obesity managed with lifestyle, resolved fracture under monitoring)

### Step 6: Identify Resolved Conditions Still Listed as Active

**Flag Type C -- Potentially Resolved Conditions:**
- Condition has `clinicalStatus` = active but `abatementDateTime` is in the past
- Condition is an acute diagnosis (e.g., pneumonia, UTI, fracture) with onset > 6 months ago and no recent encounter referencing it
- Condition was treated with a medication that has been stopped

Verify with lab data where applicable:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=[relevant-LOINC]&_sort=-date&_count=1"
```

Examples:
- "Diabetes" active but latest HbA1c (LOINC 4548-4) is consistently < 5.7% for 2+ years and no diabetes medications active
- "Hypothyroidism" active but latest TSH (LOINC 3016-3) normal and no levothyroxine prescribed
- "Iron deficiency anemia" active but latest ferritin (LOINC 2276-4) and hemoglobin (LOINC 718-7) both normal

### Step 7: Assess ICD-10 Coding Specificity

For each active condition, check if the ICD-10 code can be made more specific:

**Flag Type D -- Coding Specificity Opportunities:**

| Current Code | Issue | Suggested Improvement |
|-------------|-------|----------------------|
| E11.9 (Type 2 DM without complications) | Patient has documented neuropathy or nephropathy | E11.40 (with neuropathy) or E11.22 (with CKD) |
| I10 (Essential hypertension) | Adequate but check for hypertensive heart/kidney disease | I11.x, I12.x, I13.x if applicable |
| E78.5 (Hyperlipidemia, unspecified) | Can specify type | E78.00 (pure hypercholesterolemia), E78.1 (hypertriglyceridemia) |
| J45.20 (Mild intermittent asthma, uncomplicated) | Patient on daily controller | J45.30 (mild persistent) or higher |
| M17.9 (Osteoarthritis of knee, unspecified) | Laterality missing | M17.11 (right), M17.12 (left) |
| F32.9 (Major depressive disorder, unspecified) | Can specify episode/severity | F32.0 (mild), F32.1 (moderate), F32.2 (severe) |

### Step 8: Format Output

```
PROBLEM LIST REVIEW
===================
Patient: [name] ([patient-id])
Reviewed: [timestamp]
Active Conditions: [count]
Inactive/Resolved: [count]
Active Medications: [count]

ACTIVE PROBLEM LIST
-------------------
1. [Condition] - SNOMED: [code] / ICD-10: [code] - Onset: [date]
   Treatment: [matching medication(s)]
2. ...

FINDINGS
--------

[A] MISSING DIAGNOSES (Medications imply undocumented conditions)
  1. On Metformin 1000mg BID but Type 2 Diabetes not on problem list
     -> Recommend adding: E11.9 Type 2 diabetes mellitus without complications
  2. ...

[B] UNTREATED CONDITIONS
  1. Essential hypertension (I10) - Active since 2020 - No antihypertensive medication found
     -> Verify: Is patient on lifestyle management only? If pharmacotherapy indicated, address.
  2. ...

[C] POTENTIALLY RESOLVED CONDITIONS
  1. Acute bronchitis (J20.9) - Onset: 2025-03-15 - Still listed as active
     -> Recommend: Update clinicalStatus to "resolved"
  2. ...

[D] CODING SPECIFICITY OPPORTUNITIES
  1. E11.9 (DM2 without complications) - Patient has documented diabetic neuropathy
     -> Recommend: E11.40 (Type 2 DM with diabetic neuropathy)
  2. ...

SUMMARY
-------
- [N] conditions reviewed
- [N] missing diagnoses identified
- [N] conditions potentially untreated
- [N] conditions potentially resolved
- [N] coding improvements suggested
```

## Examples

### Example 1: Routine Problem List Audit

**User says:** "Review the problem list for patient 44321"

**Actions:**
1. `fhir_search` Condition?patient=44321 -- returns 8 conditions (6 active, 2 resolved)
2. `fhir_search` MedicationRequest?patient=44321&status=active -- returns 7 active medications
3. `fhir_search` Encounter?patient=44321&date=ge2025-02-07&_sort=-date -- returns 4 encounters

**Cross-reference findings:**
- Metformin 1000mg and Glipizide 5mg on med list, "Type 2 Diabetes" IS on problem list. OK.
- Lisinopril 20mg on med list, "Hypertension" IS on problem list. OK.
- Omeprazole 20mg on med list, but NO GERD or peptic ulcer on problem list. FLAG A.
- Sertraline 100mg on med list, but NO depression or anxiety on problem list. FLAG A.
- "Acute sinusitis" (J01.90) onset 2025-01-10, still active, no recent encounter. FLAG C.
- Diabetes coded as E11.9 but patient also has documented CKD Stage 3. FLAG D: suggest E11.22.

**Result:** Formatted report with 2 missing diagnoses, 1 potentially resolved condition, 1 coding improvement.

### Example 2: Post-Discharge Problem List Reconciliation

**User says:** "Audit the problem list for patient 77890 after their hospital stay"

**Actions:**
1. `fhir_search` Condition?patient=77890 -- returns 12 conditions
2. `fhir_search` MedicationRequest?patient=77890&status=active -- returns 10 medications
3. `fhir_search` Encounter?patient=77890&date=ge2025-02-07&_sort=-date -- returns 6 encounters including inpatient stay

**Findings:**
- New medications from discharge (apixaban, metoprolol) without corresponding new diagnoses added (atrial fibrillation). FLAG A.
- Hospital-acquired pneumonia on problem list from admission, patient discharged 3 weeks ago with completed antibiotic course. FLAG C.
- "Heart failure" (I50.9, unspecified) could be I50.22 (chronic systolic, compensated) based on echo results. FLAG D.

## Troubleshooting

### Condition Resources Use Only ICD-10 Codes Without SNOMED
- Cross-reference is still possible. Use the ICD-10 code directly for matching.
- Note that some medication-condition mappings are more reliable with SNOMED codes. Use the display text as a fallback for matching.

### MedicationRequest Uses medicationReference Instead of medicationCodeableConcept
- The medication details are in a separate Medication resource. Resolve the reference:
  ```
  Tool: fhir_read
  resourceType: "Medication"
  id: "[medication-reference-id]"
  ```
- Extract the drug name from `Medication.code.coding`.

### Too Many Conditions Returned (Including Historical)
- Filter client-side by `category` = "problem-list-item" to exclude encounter-only diagnoses.
- Sort by `clinicalStatus` to prioritize active conditions.
- If the list exceeds 50 items, paginate with `_count=50` and follow `Bundle.link` where `relation` = "next".

## Related Skills

- `clinical-summary-generator` -- for full chart summary including problems in context
- `allergy-adverse-reaction-summary` -- for allergy-specific review
- `patient-demographics-summary` -- when problem list review reveals missing demographic data
