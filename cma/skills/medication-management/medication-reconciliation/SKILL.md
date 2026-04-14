---
name: langcare-medication-reconciliation
description: >
  Performs comprehensive medication reconciliation across care transitions by
  comparing inpatient, outpatient, and patient-reported medication lists.
  Identifies discrepancies including duplicates, therapeutic duplications,
  dose changes, and discontinued medications. Use when asked to reconcile
  medications, compare med lists, admission meds, or discharge meds.
---

# Medication Reconciliation

## When to Use This Skill
Use when a clinician needs to reconcile medications across care settings per Joint Commission NPSG.03.06.01 -- admission, discharge, or transfer.

## Clinical Workflow
1. Use `fhir_search` to pull active MedicationRequest resources (prescribed orders) with `status=active,on-hold`
2. Use `fhir_search` to pull active MedicationStatement resources (patient-reported home medications)
3. Use `fhir_search` to pull MedicationAdministration resources (inpatient administered medications) filtered by admission date
4. Use `fhir_search` to pull active AllergyIntolerance and Condition resources for safety cross-reference
5. Normalize all medications to RxNorm codes; build unified comparison table with columns: Medication, Dose, Frequency, Route, Source, Status
6. Identify discrepancies: duplicates, therapeutic duplications (same drug class), dose changes across sources, discontinued meds still listed active, medications with no matching indication, allergy conflicts
7. Flag ISMP high-alert medications (anticoagulants, insulin, opioids, chemotherapy) for extra scrutiny
8. Present reconciliation report and use `fhir_update` to update MedicationStatement/MedicationRequest status if authorized

## FHIR Resources
- **MedicationRequest** -- Prescribed orders: status, intent, medicationCodeableConcept, dosageInstruction, authoredOn
- **MedicationStatement** -- Patient-reported medications: status, medicationCodeableConcept, dosage, informationSource
- **MedicationAdministration** -- Inpatient administrations: status, medicationCodeableConcept, dosage, effectiveDateTime
- **AllergyIntolerance** -- Drug allergy cross-check
- **Condition** -- Validate medication indications

## FHIR Query Examples
### Pull Active Prescriptions
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active,on-hold&_include=MedicationRequest:medication&_count=100")
```

### Pull Patient-Reported Medications
```
fhir_search(resourceType="MedicationStatement", queryParams="patient=[patient-id]&status=active,intended,on-hold&_count=100")
```

### Pull Inpatient Administrations
```
fhir_search(resourceType="MedicationAdministration", queryParams="patient=[patient-id]&effective-time=ge[admission-date]&_count=100")
```

## Clinical Guidelines
- Joint Commission NPSG.03.06.01: Reconcile medications at every transition of care
- WHO High 5s Medication Reconciliation protocol
- ISMP high-alert medication list for enhanced verification

## Interpretation Guide
- Use RxNorm coding (system `http://www.nlm.nih.gov/research/umls/rxnorm`) for medication matching; fall back to text matching when codes unavailable
- Present findings in categories: Verified Medications (matched across sources), Discrepancies Found (categorized by type), High-Alert Medications (highlighted), Allergy Conflicts, and Recommendations
- For each discrepancy, show both values side-by-side with source attribution

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
