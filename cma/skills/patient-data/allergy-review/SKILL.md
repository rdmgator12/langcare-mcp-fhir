---
name: langcare-allergy-review
description: >
  Retrieves and organizes allergy and adverse reaction data from FHIR
  AllergyIntolerance resources. Cross-references against active medications
  for contraindications. Use when asked to review allergies, check for drug
  allergies, allergy summary, adverse reactions, or verify allergy
  documentation before prescribing.
---

# Allergy and Adverse Reaction Review

## When to Use This Skill
Use when a clinician needs a comprehensive view of documented allergies and adverse reactions, cross-referenced against current medications for safety validation.

## Clinical Workflow
1. Use `fhir_search` to pull all AllergyIntolerance resources for the patient (active, inactive, resolved)
2. Use `fhir_search` to pull active MedicationRequest resources for cross-reference
3. Classify allergies by type: drug, food, environmental, biologic
4. For drug allergies, cross-reference against active medications and flag any contraindicated prescriptions
5. Identify allergy documentation gaps: no known allergies (NKA) vs not assessed vs documented allergies
6. Use `fhir_update` to correct allergy records if authorized (e.g., mark resolved, update reaction details)

## FHIR Resources
- **AllergyIntolerance** -- Allergy entries: clinicalStatus, type (allergy vs intolerance), category (food, medication, environment, biologic), code, reaction (substance, manifestation, severity), criticality
- **MedicationRequest** -- Active prescriptions for contraindication checking
- **Patient** -- Demographics for allergy context

## FHIR Query Examples
### Pull All Allergies
```
fhir_search(resourceType="AllergyIntolerance", queryParams="patient=[patient-id]")
```

### Pull Active Medications for Cross-Reference
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

## Clinical Guidelines
- Joint Commission NPSG.03.06.01 requires allergy documentation reconciliation at transitions of care
- Cross-reactivity groups: penicillins (ampicillin, amoxicillin, piperacillin), cephalosporins (~2% cross-reactivity with penicillin), sulfonamides, NSAIDs
- Distinguish allergy (immune-mediated: anaphylaxis, urticaria, angioedema) from intolerance (non-immune: GI upset, headache)
- Criticality: high = life-threatening potential, low = non-life-threatening, unable-to-assess

## Interpretation Guide
- Present drug allergies first (highest prescribing safety impact), then food, then environmental
- For each allergy: substance, reaction type, severity, criticality, date recorded, verificationStatus
- Flag high-criticality allergies prominently
- When cross-referencing medications, check both exact matches and drug class cross-reactivity
- If AllergyIntolerance list is empty, distinguish NKA (code for "no known allergy" present) from "not assessed" (no AllergyIntolerance resources at all)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
