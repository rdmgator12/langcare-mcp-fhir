---
name: drug-interaction-checker
description: |
  Checks active medications for drug-drug interactions by pulling all MedicationRequests and cross-referencing known interaction pairs with severity classification and CYP450 pathways. Use when user asks to "check drug interactions", "check for interactions", "are these medications safe together", mentions "drug-drug interaction", "CYP450", or needs interaction screening. Do NOT use for medication reconciliation, adherence, or single drug information lookups.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: medication-management
---

# Drug Interaction Checker

## Overview

Pull all active medications for a patient, cross-reference each pair against known major drug-drug interaction databases, classify by severity (contraindicated, major, moderate, minor), identify CYP450 pathway conflicts, and suggest therapeutic alternatives when significant interactions are detected.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| MedicationRequest | Active prescribed medications | status, medicationCodeableConcept, dosageInstruction |
| MedicationStatement | Patient-reported medications including OTC | status, medicationCodeableConcept, dosage |
| Patient | Patient demographics, age for risk stratification | birthDate, gender |
| Condition | Comorbidities affecting interaction severity | code, clinicalStatus |
| Observation | Renal/hepatic function affecting drug metabolism | code, valueQuantity |
| AllergyIntolerance | Cross-sensitivity patterns | code, reaction |

## Instructions

### Step 1: Verify Patient and Retrieve Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age, gender, weight if available. Age and organ function affect CYP450 metabolism rates.

### Step 2: Pull All Active Medications

**2a: Prescribed medications**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

**2b: Patient-reported medications (includes OTC, supplements)**
```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

Merge both lists. Deduplicate by RxNorm code (`system: http://www.nlm.nih.gov/research/umls/rxnorm`).

### Step 3: Pull Renal and Hepatic Function

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=33914-3,48642-3,69405-9,77147-7&_sort=-date&_count=4"
```

LOINC codes:
- 33914-3: eGFR (CKD-EPI)
- 48642-3: eGFR for African American
- 69405-9: Glomerular filtration rate
- 77147-7: eGFR (CKD-EPI 2021)

Impaired renal/hepatic function reduces drug clearance and increases interaction risk.

### Step 4: Build Medication Pair Matrix

For N active medications, evaluate N*(N-1)/2 unique pairs. For each pair, check against the interaction reference (see references/major-drug-interactions.md) in the following order:

1. **Contraindicated combinations** - Never co-prescribe
2. **Major interactions** - Significant clinical risk, avoid or monitor closely
3. **Moderate interactions** - Clinical significance in certain patients
4. **Minor interactions** - Limited clinical significance

### Step 5: Evaluate CYP450 Pathway Conflicts

For each medication, identify its CYP450 profile (see references/cyp450-pathways.md):
- Is it a **substrate** (metabolized by the enzyme)?
- Is it an **inhibitor** (blocks the enzyme, increasing substrate levels)?
- Is it an **inducer** (accelerates the enzyme, decreasing substrate levels)?

Flag when:
- A strong inhibitor of CYP3A4 is combined with a CYP3A4 substrate with a narrow therapeutic index
- A strong inducer is combined with a substrate where reduced levels cause treatment failure
- Multiple inhibitors of the same pathway are present (additive inhibition)

### Step 6: Classify and Report Interactions

For each identified interaction, report:

| Field | Content |
|-------|---------|
| Drug A | Name, dose, RxNorm code |
| Drug B | Name, dose, RxNorm code |
| Severity | Contraindicated / Major / Moderate / Minor |
| Mechanism | Pharmacokinetic (CYP450, P-gp, renal) or Pharmacodynamic (additive, synergistic, antagonistic) |
| Clinical Effect | What happens (e.g., "increased bleeding risk", "QT prolongation") |
| Evidence Level | Established / Probable / Suspected / Possible |
| Management | Specific recommendation (avoid, adjust dose, monitor parameter, timing separation) |
| Alternative | Suggested substitute medication if available |

### Step 7: Generate Interaction Summary

Structure output:

1. **Contraindicated Pairs** (if any) - RED flag, require immediate prescriber attention
2. **Major Interactions** - ORANGE flag, require monitoring plan or alternative
3. **Moderate Interactions** - YELLOW flag, clinical judgment needed
4. **Minor Interactions** - Informational only
5. **CYP450 Summary** - Table of which enzymes are affected by current regimen
6. **Recommended Actions** - Prioritized list of changes

### Step 8: Document Interaction Check (if requested)

```
Tool: fhir_create
resourceType: "DetectedIssue"
resource: {
  "resourceType": "DetectedIssue",
  "status": "preliminary",
  "code": {
    "coding": [{
      "system": "http://terminology.hl7.org/CodeSystem/v3-ActCode",
      "code": "DRG",
      "display": "Drug Interaction Alert"
    }]
  },
  "severity": "high",
  "patient": {"reference": "Patient/[patient-id]"},
  "identifiedDateTime": "[current-datetime]",
  "implicated": [
    {"reference": "MedicationRequest/[drug-a-id]"},
    {"reference": "MedicationRequest/[drug-b-id]"}
  ],
  "detail": "[interaction description]",
  "mitigation": [{
    "action": {"text": "[recommended action]"},
    "date": "[date]"
  }]
}
```

## Examples

### Example 1: Checking Interactions for a New Prescription

**User says**: "Check drug interactions for patient 67890 - about to start amiodarone."

**Actions**:
1. Read Patient 67890.
2. Pull active MedicationRequests: finds warfarin 5mg daily, simvastatin 40mg daily, metformin 1000mg BID, lisinopril 20mg daily.
3. Pull renal function: eGFR 55 mL/min (Stage 3a CKD).
4. Cross-reference amiodarone against each active medication.
5. Check CYP450: amiodarone is a strong inhibitor of CYP3A4, CYP2C9, CYP2D6, and P-glycoprotein.

**Result**:
```
DRUG INTERACTION CHECK - Patient 67890
New medication: Amiodarone

CONTRAINDICATED: None

MAJOR INTERACTIONS:
1. Amiodarone + Warfarin [MAJOR]
   Mechanism: Amiodarone inhibits CYP2C9, reducing warfarin metabolism
   Effect: INR increase of 50-70%, bleeding risk
   Management: Reduce warfarin dose by 30-50%. Check INR within 3-5 days, then weekly x4.

2. Amiodarone + Simvastatin [MAJOR]
   Mechanism: Amiodarone inhibits CYP3A4, increasing simvastatin levels
   Effect: Rhabdomyolysis risk. Simvastatin max dose with amiodarone is 20mg.
   Alternative: Switch to pravastatin 40mg (not CYP3A4 metabolized) or rosuvastatin 10mg.

MODERATE INTERACTIONS:
3. Amiodarone + Metformin [MODERATE]
   Mechanism: Amiodarone may increase metformin levels via OCT transporter inhibition
   Effect: Increased metformin exposure, lactic acidosis risk (especially with eGFR 55)
   Management: Monitor renal function closely. Consider metformin dose reduction.

CYP450 IMPACT:
- CYP3A4: STRONGLY INHIBITED by amiodarone (affects simvastatin)
- CYP2C9: STRONGLY INHIBITED by amiodarone (affects warfarin)
- CYP2D6: INHIBITED by amiodarone (no current substrates affected)

RECOMMENDED ACTIONS:
1. [URGENT] Reduce warfarin dose 30-50% before starting amiodarone
2. [URGENT] Switch simvastatin to pravastatin or reduce to 20mg max
3. [MONITOR] Renal function and metformin tolerance
4. [SCHEDULE] INR check 3-5 days after amiodarone start
```

### Example 2: Routine Polypharmacy Interaction Screen

**User says**: "Screen all drug interactions for patient Jane Doe, DOB 1940-07-22."

**Actions**:
1. Search Patient by name and DOB. Confirm identity. Age 85.
2. Pull 14 active medications (polypharmacy).
3. Evaluate 91 unique pairs (14 choose 2).
4. Identify 3 major, 5 moderate, 2 minor interactions.
5. Note elderly patient - flag CYP450 metabolism decline with age.

**Result**:
```
DRUG INTERACTION SCREEN - Jane Doe (Age 85)
Total active medications: 14
Pairs evaluated: 91

MAJOR INTERACTIONS: 3
1. Clopidogrel + Omeprazole [MAJOR]
   Mechanism: Omeprazole inhibits CYP2C19, reducing clopidogrel activation
   Alternative: Switch to pantoprazole (weaker CYP2C19 inhibitor)

2. Methotrexate + Trimethoprim [MAJOR]
   Mechanism: Both cause folate depletion, additive bone marrow suppression
   Management: Avoid combination. If unavoidable, supplement with leucovorin.

3. Fluoxetine + Tramadol [MAJOR]
   Mechanism: Serotonin syndrome risk (both serotonergic). Fluoxetine inhibits CYP2D6, reducing tramadol conversion to active metabolite.
   Alternative: Switch tramadol to acetaminophen or non-serotonergic analgesic.

MODERATE INTERACTIONS: 5
[... listed with management ...]

AGE-RELATED CONSIDERATIONS:
- CYP450 activity reduced ~30% at age 85
- Renal clearance likely reduced (check eGFR)
- Polypharmacy (14 meds) increases interaction probability
- Recommend comprehensive prescription appropriateness review
```

## Troubleshooting

### Medication codes are not in RxNorm
- Use `medicationCodeableConcept.text` for name-based matching against interaction tables.
- If medication is a contained reference, resolve via `fhir_read` on the referenced Medication resource.
- Note in output that coded interaction checking was not possible for unrecognized medications.

### Patient has very large medication list (>20 medications)
- Prioritize pairs involving high-alert medications (anticoagulants, antiarrhythmics, opioids, immunosuppressants).
- Focus on contraindicated and major interactions first.
- Present interactions grouped by severity rather than listing all 190+ pairs.

### Interaction found but clinical significance is unclear
- Report the interaction with its evidence level (Established/Probable/Suspected/Possible).
- Include the mechanism so the clinician can assess applicability.
- Note patient-specific factors that modify risk (renal function, age, duration of concurrent use).

## Related Skills

- `medication-reconciliation` - Reconcile medication list before running interaction check
- `prescription-appropriateness-review` - Broader appropriateness review beyond interactions
- `opioid-risk-assessment` - Focused assessment when opioid interactions detected
- `problem-list-review` - Review conditions that affect interaction severity
