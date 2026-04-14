---
name: langcare-drug-interactions
description: >
  Checks for clinically significant drug-drug interactions across active
  medications using CYP450 metabolism pathways, pharmacodynamic interactions,
  and contraindicated combinations. Use when asked to check drug interactions,
  medication safety check, CYP interactions, or before adding a new medication
  to an existing regimen.
---

# Drug Interaction Checker

## When to Use This Skill
Use when a clinician needs to evaluate the active medication list for drug-drug interactions, particularly before prescribing a new medication or during medication reconciliation.

## Clinical Workflow
1. Use `fhir_search` to pull all active MedicationRequest and MedicationStatement resources
2. Use `fhir_search` to pull active AllergyIntolerance for drug allergy cross-check
3. Use `fhir_search` to pull relevant Observation resources (renal function, hepatic function) for dose adjustment context
4. Map each medication to its CYP450 substrate/inhibitor/inducer profile (see references/cyp450-table.md)
5. Identify pharmacokinetic interactions (CYP inhibition/induction affecting drug levels) and pharmacodynamic interactions (additive/synergistic effects: QT prolongation, serotonin syndrome, bleeding risk, CNS depression)
6. Classify interactions by severity: contraindicated, major, moderate, minor
7. Present findings with mechanism, clinical significance, and recommended action

## FHIR Resources
- **MedicationRequest** -- Active prescriptions with medication codes
- **MedicationStatement** -- Patient-reported medications including OTC and supplements
- **AllergyIntolerance** -- Drug allergy cross-reference
- **Observation** -- Renal function (CrCl/eGFR), hepatic function (ALT, AST) for dose adjustment context

## FHIR Query Examples
### Pull All Active Medications
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

### Pull Patient-Reported Medications (OTC, supplements)
```
fhir_search(resourceType="MedicationStatement", queryParams="patient=[patient-id]&status=active&_count=100")
```

### Pull Renal Function for Dose Context
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=33914-3&_sort=-date&_count=1")
```

## Clinical Guidelines
- FDA drug interaction guidance for labeling
- AHA/ACC guidelines on QT-prolonging drug combinations
- Beers Criteria for potentially inappropriate drug combinations in elderly

## Interpretation Guide
- Classify by severity: Contraindicated (avoid combination), Major (modify therapy), Moderate (monitor closely), Minor (be aware)
- For CYP interactions, specify the pathway (CYP3A4, CYP2D6, etc.), whether the interacting drug is an inhibitor or inducer, and the expected effect on the substrate drug levels
- For pharmacodynamic interactions, specify the shared mechanism (QT prolongation, serotonin, bleeding, CNS depression, anticholinergic burden)
- Provide actionable recommendations: alternative medications, dose adjustments, monitoring parameters

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
