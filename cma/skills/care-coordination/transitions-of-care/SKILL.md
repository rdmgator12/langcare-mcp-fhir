---
name: langcare-transitions-of-care
description: >
  Generates a structured transition-of-care summary using the I-PASS framework
  for safe patient handoffs between care settings. Compiles illness severity,
  patient summary, action list, situation awareness, and synthesis from FHIR
  data. Use when asked for transition of care summary, handoff, transfer
  summary, or I-PASS handoff.
---

# Transitions of Care Summary

## When to Use This Skill
Use when a clinician needs a structured care transition document for transfers between inpatient units, facilities, or inpatient-to-outpatient handoffs.

## Clinical Workflow
1. Use `fhir_search` to retrieve the current Encounter with admission details
2. Use `fhir_read` to retrieve Patient demographics
3. Use `fhir_search` to pull active Conditions for illness severity and problem list
4. Use `fhir_search` to pull active MedicationRequest for current medication list
5. Use `fhir_search` to pull recent Observation resources (vitals, labs) for clinical status
6. Use `fhir_search` to pull pending ServiceRequest resources for action items
7. Use `fhir_search` to pull active CarePlan for ongoing care goals
8. Structure output using I-PASS framework (see references/ipass-framework.md): Illness severity, Patient summary, Action list, Situation awareness, Synthesis by receiver
9. Use `fhir_create` to create DocumentReference with the transition summary

## FHIR Resources
- **Encounter** -- Current admission context, reason, period
- **Patient** -- Demographics
- **Condition** -- Active problems, admission diagnosis
- **MedicationRequest** -- Current medications
- **Observation** -- Recent vitals and labs
- **ServiceRequest** -- Pending orders and action items
- **CarePlan** -- Active care plans and goals
- **DocumentReference** -- Output: transition summary document

## FHIR Query Examples
### Pull Current Encounter
```
fhir_search(resourceType="Encounter", queryParams="patient=[patient-id]&status=in-progress&class=http://terminology.hl7.org/CodeSystem/v3-ActCode|IMP")
```

### Pull Pending Actions
```
fhir_search(resourceType="ServiceRequest", queryParams="patient=[patient-id]&status=active,draft&encounter=[encounter-id]")
```

## Clinical Guidelines
- I-PASS Study Group (Starmer et al., NEJM 2014): Standardized handoff reduces medical errors
- Joint Commission NPSG.02.05.01: Standardize approach to hand-off communications
- CMS Conditions of Participation for discharge/transfer documentation

## Interpretation Guide
- I-PASS structure: I = Illness severity (stable/watcher/unstable); P = Patient summary (one-liner, key events, ongoing issues); A = Action list (to-do items with responsible party and timeline); S = Situation awareness (what to watch for, contingency plans); S = Synthesis (receiver read-back confirmation)
- Include code status, allergies, isolation precautions, and fall risk in every handoff
- Flag anticipatory guidance: potential complications, if-then plans, escalation criteria

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
