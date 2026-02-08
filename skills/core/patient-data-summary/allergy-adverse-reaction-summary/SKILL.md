---
name: allergy-adverse-reaction-summary
description: |
  Retrieves, categorizes, and analyzes AllergyIntolerance resources with severity grading, cross-reactivity flags,
  and medication conflict detection. Use when user asks to "check allergies", "allergy summary", "drug allergies",
  "review adverse reactions", "allergy list", "is the patient allergic to", or "cross-reactivity check".
  Do NOT use for full clinical summaries or medication reconciliation -- use domain-specific skills instead.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: patient-data
---

# Allergy & Adverse Reaction Summary

## Overview

Pull all AllergyIntolerance resources for a patient. Categorize by type (drug, food, environmental). Grade severity and criticality. Flag high-risk allergies that require special clinical attention (anaphylaxis history, contrast dye, latex, NSAID sensitivity). Cross-reference against current medications to detect active conflicts. Output a structured summary with actionable safety flags.

## FHIR Resources Used

| Resource | Purpose | Search Parameters |
|----------|---------|-------------------|
| AllergyIntolerance | Allergy and intolerance records | patient, clinical-status |
| MedicationRequest | Current medications for conflict detection | patient, status=active |
| Observation | IgE levels if available | patient, code (LOINC) |

## Instructions

### Step 1: Pull All AllergyIntolerance Resources

```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]"
```

Retrieve all records regardless of clinical status to capture full history. For each entry extract:
- `code.coding`: Substance (RxNorm for drugs, SNOMED CT for non-drugs)
- `clinicalStatus`: active, inactive, resolved
- `verificationStatus`: confirmed, unconfirmed, refuted, entered-in-error
- `type`: allergy vs intolerance
- `category`: food, medication, environment, biologic
- `criticality`: low, high, unable-to-assess
- `reaction[]`: Each reaction event
  - `manifestation[].coding[].display`: What happened (hives, anaphylaxis, rash, etc.)
  - `severity`: mild, moderate, severe
  - `onset`: When the reaction occurred
  - `substance`: Specific substance if different from the top-level code
- `onsetDateTime`: When allergy was first identified
- `recorder`: Who documented it
- `note`: Free-text notes

### Step 2: Categorize Allergies

Group into four categories:

**Drug Allergies** (`category` = "medication"):
- Identify drug class from the substance code or name
- Check for class-level allergies vs specific drug allergies
- Note if the allergy is to a drug class (e.g., "penicillins") vs specific drug (e.g., "amoxicillin")

**Food Allergies** (`category` = "food"):
- Common: peanuts, tree nuts, shellfish, milk, eggs, wheat, soy, fish
- Note severity -- food allergies with anaphylaxis history are high-risk

**Environmental Allergies** (`category` = "environment"):
- Pollen, dust mites, mold, animal dander, insect venom
- Usually lower clinical urgency unless anaphylaxis to insect stings

**Biologic/Other** (`category` = "biologic" or uncategorized):
- Latex, contrast dye, blood products, vaccines

### Step 3: Flag High-Risk Allergies

Scan the allergy list for the following high-risk patterns:

**Anaphylaxis History:**
- Any allergy where `reaction[].manifestation` includes SNOMED 39579001 (Anaphylaxis) or display text contains "anaphylaxis" or "anaphylactic"
- Flag with: "ANAPHYLAXIS RISK -- Ensure epinephrine availability"

**Contrast Dye Allergy:**
- Substance code or display contains "iodine", "contrast", "iodinated contrast"
- Flag with: "CONTRAST ALLERGY -- Premedication protocol required for CT with contrast"

**Latex Allergy:**
- Substance = SNOMED 111088007 (Latex) or display contains "latex"
- Flag with: "LATEX ALLERGY -- Use non-latex gloves and equipment"
- Check for cross-reactive foods: banana, avocado, kiwi, chestnut

**NSAID Sensitivity:**
- Allergy to aspirin, ibuprofen, naproxen, or "NSAIDs" as a class
- Flag with: "NSAID SENSITIVITY -- Avoid all COX inhibitors. Consider acetaminophen as alternative."
- Check for aspirin-exacerbated respiratory disease (Samter's triad): NSAID allergy + asthma + nasal polyps

**Sulfonamide Allergy:**
- Allergy to "sulfa", "sulfamethoxazole", "trimethoprim-sulfamethoxazole", or sulfonamide antibiotics
- Note: Sulfonamide antibiotic allergy does NOT contraindicate non-antibiotic sulfonamides (furosemide, thiazides, celecoxib) -- cross-reactivity is negligible but document awareness

**Opioid Allergy:**
- If allergy to one opioid, check if it is a true allergy (e.g., anaphylaxis) or a side effect misclassified as allergy (nausea, itching are common opioid side effects, not allergies)
- Flag specific opioid vs opioid class

### Step 4: Check Cross-Reactivity Patterns

Apply known cross-reactivity rules:

**Penicillin Cross-Reactivity:**
- Penicillin allergy -> ~2% cross-reactivity with cephalosporins (1st gen > 2nd gen > 3rd gen)
- Penicillin allergy -> ~1% cross-reactivity with carbapenems
- If penicillin allergy with anaphylaxis: avoid all beta-lactams pending allergy testing
- If penicillin allergy with mild rash only: cephalosporins (2nd gen+) generally safe

**Cephalosporin Cross-Reactivity:**
- 1st gen cephalosporins (cephalexin, cefazolin) have highest cross-reactivity with penicillin
- Cross-reactivity is side-chain dependent, not ring-structure dependent (current evidence)

**ACE Inhibitor Angioedema:**
- ACE inhibitor-induced angioedema -> contraindication to all ACE inhibitors
- ARBs are generally safe alternatives (low cross-reactivity)
- Document as a separate allergy entry if not already present

**Statin Myopathy:**
- Myopathy/rhabdomyolysis with one statin does not necessarily contraindicate all statins
- Hydrophilic statins (rosuvastatin, pravastatin) may be tolerated

### Step 5: Cross-Reference with Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

For each active medication, check against the allergy list:
- Direct match: medication name/code matches allergy substance
- Class match: medication belongs to same drug class as allergy substance
- Cross-reactivity match: medication has known cross-reactivity with allergy substance

**Flag any active conflict as:**
```
ACTIVE CONFLICT: Patient is on [medication] but has documented allergy to [substance]
Reaction history: [manifestation] - Severity: [severity]
Action required: Verify with prescriber. This may be a monitored override or a documentation error.
```

### Step 6: Assess Documentation Quality

Flag documentation issues:
- Allergy with no reaction documented (`reaction` array empty)
- Allergy with no severity specified
- `verificationStatus` = "unconfirmed" for allergies older than 1 year (should be confirmed or refuted)
- `type` not specified (allergy vs intolerance distinction missing)
- No allergies documented AND no "No Known Allergies" entry (allergy status unknown vs NKDA)

### Step 7: Format Output

```
ALLERGY & ADVERSE REACTION SUMMARY
====================================
Patient: [name] ([patient-id])
Reviewed: [timestamp]
Total Entries: [count]
Allergy Status: [Reviewed / Not Reviewed / NKDA]

HIGH-RISK ALERTS
================
[!] ANAPHYLAXIS RISK: [substance] - ensure epinephrine access
[!] CONTRAST ALLERGY: Premedication required
[!] LATEX ALLERGY: Non-latex equipment required
[!] ACTIVE CONFLICT: On [medication] with allergy to [substance]

DRUG ALLERGIES
==============
1. [Substance] (RxNorm: [code])
   Type: Allergy | Status: Active | Criticality: High
   Reaction: [manifestation] | Severity: Severe
   Cross-reactivity note: [if applicable]
   Documented: [date] by [recorder]
2. ...

FOOD ALLERGIES
==============
1. [Substance] - Reaction: [manifestation] - Severity: [severity]
2. ...

ENVIRONMENTAL ALLERGIES
=======================
1. [Substance] - Reaction: [manifestation] - Severity: [severity]
2. ...

INTOLERANCES (Non-Allergic)
============================
1. [Substance] - Type: Intolerance - Reaction: [manifestation]
2. ...

DOCUMENTATION QUALITY
=====================
- [N] allergies missing reaction details
- [N] allergies with unconfirmed verification status
- [Notes on documentation gaps]

CROSS-REACTIVITY NOTES
=======================
- [Relevant cross-reactivity information based on documented allergies]
```

## Examples

### Example 1: Patient With Multiple Drug Allergies

**User says:** "Check allergies for patient 33221"

**Actions:**
1. `fhir_search` AllergyIntolerance?patient=33221 -- returns 5 entries
2. `fhir_search` MedicationRequest?patient=33221&status=active -- returns 8 medications

**Result:**
```
ALLERGY & ADVERSE REACTION SUMMARY
====================================
Patient: David Park (33221)
Reviewed: 2026-02-07T10:00:00Z
Total Entries: 5
Allergy Status: Reviewed

HIGH-RISK ALERTS
================
[!] ANAPHYLAXIS RISK: Penicillin - ensure epinephrine access
[!] ACTIVE CONFLICT: On cephalexin (Keflex) with documented penicillin allergy
    (anaphylaxis history) - IMMEDIATE REVIEW REQUIRED

DRUG ALLERGIES
==============
1. Penicillin (RxNorm: 7980)
   Type: Allergy | Status: Active | Criticality: High
   Reaction: Anaphylaxis | Severity: Severe
   Cross-reactivity: 2% risk with cephalosporins. Given anaphylaxis history,
   avoid all beta-lactams pending formal allergy testing.
   Documented: 2019-05-12

2. Codeine (RxNorm: 2670)
   Type: Intolerance | Status: Active | Criticality: Low
   Reaction: Nausea, vomiting | Severity: Mild
   Note: Likely opioid side effect, not true allergy. Other opioids
   may be tolerated. Consider reclassifying as intolerance if not already.
   Documented: 2020-08-03

FOOD ALLERGIES
==============
1. Shellfish (SNOMED: 278840001)
   Reaction: Urticaria, throat tightness | Severity: Moderate
   Note: Monitor for iodinated contrast cross-reactivity (historically
   cited but evidence is weak; shellfish allergy is protein-based,
   not iodine-based)

ENVIRONMENTAL ALLERGIES
=======================
1. Dust mites (SNOMED: 260147004)
   Reaction: Rhinitis | Severity: Mild

INTOLERANCES
============
1. Lactose (SNOMED: 782415009)
   Type: Intolerance | Reaction: Bloating, diarrhea | Severity: Mild

DOCUMENTATION QUALITY
=====================
- Codeine entry classified as "allergy" but reaction pattern suggests intolerance
- Dust mite allergy missing severity in FHIR record (inferred as mild from manifestation)
```

### Example 2: No Allergies Documented -- Ambiguous Status

**User says:** "Does patient 99001 have any allergies?"

**Actions:**
1. `fhir_search` AllergyIntolerance?patient=99001 -- returns 0 entries

**Result:**
```
ALLERGY & ADVERSE REACTION SUMMARY
====================================
Patient: Lisa Chen (99001)
Reviewed: 2026-02-07T10:15:00Z
Total Entries: 0
Allergy Status: ** NOT REVIEWED **

HIGH-RISK ALERTS
================
[!] ALLERGY STATUS UNKNOWN: No allergy records found. This does NOT mean
    "No Known Allergies" -- it means allergies have never been documented.
    Allergy history MUST be collected before prescribing.

RECOMMENDATION
==============
Ask the patient about:
1. Drug allergies (especially antibiotics, NSAIDs, anesthetics)
2. Food allergies
3. Environmental allergies
4. Latex sensitivity
5. Contrast dye reactions
6. Previous adverse drug reactions

If patient reports no allergies, document a "No Known Allergies" entry:
  Tool: fhir_create
  resourceType: "AllergyIntolerance"
  resource: {
    "clinicalStatus": {"coding": [{"system": "http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical", "code": "active"}]},
    "verificationStatus": {"coding": [{"system": "http://terminology.hl7.org/CodeSystem/allergyintolerance-verification", "code": "confirmed"}]},
    "code": {"coding": [{"system": "http://snomed.info/sct", "code": "716186003", "display": "No known allergy"}]},
    "patient": {"reference": "Patient/99001"}
  }
```

## Troubleshooting

### AllergyIntolerance Uses Free Text Instead of Coded Substances
- Some EHR systems populate `code.text` but not `code.coding`. Use the display text for matching.
- Attempt fuzzy matching against known drug names and classes for cross-reference purposes.
- Flag these entries for coding improvement: "Allergy to [free text] should be coded with RxNorm/SNOMED for interoperability."

### Medication Conflict Detected But May Be Intentional Override
- Clinicians sometimes prescribe medications despite documented allergies after risk-benefit analysis (e.g., penicillin allergy but documented tolerance to amoxicillin after skin testing).
- Check for `note` or extension fields that indicate a deliberate override.
- Present the conflict but note: "This may be a monitored override. Verify with the prescribing clinician before recommending discontinuation."

### AllergyIntolerance Resources Have verificationStatus = "entered-in-error"
- Exclude these from the active summary entirely.
- If the user specifically asks about historical or retracted allergies, include them in a separate "Retracted Entries" section.

## Related Skills

- `patient-demographics-summary` -- for complete patient overview including insurance and contacts
- `problem-list-review` -- for cross-referencing allergies with conditions
- `clinical-summary-generator` -- for allergies in the context of a full CCD summary
