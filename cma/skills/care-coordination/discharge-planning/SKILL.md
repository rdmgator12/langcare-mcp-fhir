---
name: langcare-discharge-planning
description: >
  Performs comprehensive discharge readiness assessment by checking pending labs,
  imaging, medication reconciliation, follow-up appointments, patient education,
  DME orders, and home health referrals. Calculates LACE readmission risk index.
  Use when asked to check discharge readiness, discharge checklist, is patient
  ready for discharge, or prepare for discharge.
---

# Discharge Planning Checklist

## When to Use This Skill
Use when a clinician needs to verify discharge readiness by checking all CMS Condition of Participation requirements and calculating readmission risk.

## Clinical Workflow
1. Use `fhir_search` to retrieve the current inpatient Encounter (status=in-progress, class=IMP)
2. Use `fhir_search` to check for pending ServiceRequest resources (active/draft labs, imaging, consults)
3. Use `fhir_search` to check for pending Observation results (status=preliminary/registered)
4. Use `fhir_search` to verify medication reconciliation: compare active MedicationRequest against MedicationStatement for discrepancies
5. Use `fhir_search` to check for scheduled follow-up Appointments (status=booked, date >= today)
6. Use `fhir_search` to check for patient education DocumentReference resources (LOINC 69981-9)
7. Use `fhir_search` to check for DME orders (DeviceRequest) and home health referrals (ServiceRequest)
8. Calculate LACE readmission risk index: L=length of stay, A=acuity of admission, C=comorbidities (Charlson), E=ED visits in 6 months (see references/lace-index.md)
9. Generate or update discharge CarePlan with pass/fail status for each requirement

## FHIR Resources
- **Encounter** -- Current admission: period, class, reasonCode, admitSource
- **ServiceRequest** -- Pending orders: status, intent, code, category
- **Observation** -- Pending lab results: status
- **MedicationRequest** / **MedicationStatement** -- Medication reconciliation
- **Appointment** -- Follow-up scheduling: status, start, serviceType
- **DocumentReference** -- Patient education: type (LOINC 69981-9)
- **DeviceRequest** -- DME orders: status, code
- **CarePlan** -- Discharge plan with activities

## FHIR Query Examples
### Check Pending Orders
```
fhir_search(resourceType="ServiceRequest", queryParams="patient=[patient-id]&status=active,draft&encounter=[encounter-id]")
```

### Check Follow-up Appointments
```
fhir_search(resourceType="Appointment", queryParams="patient=[patient-id]&status=booked,proposed&date=ge[today]")
```

### Count ED Visits (6 months) for LACE
```
fhir_search(resourceType="Encounter", queryParams="patient=[patient-id]&class=http://terminology.hl7.org/CodeSystem/v3-ActCode|EMER&date=ge[6-months-ago]")
```

## Clinical Guidelines
- CMS Conditions of Participation for Discharge Planning (42 CFR 482.43)
- Joint Commission standards for discharge planning
- LACE index for readmission prediction (van Walraven et al.)

## Interpretation Guide
- Present as pass/fail checklist: pending labs/imaging, medication reconciliation, follow-up appointments, patient education, DME orders, home health referral
- LACE score: 0-4 low risk, 5-9 moderate risk, >=10 high risk for 30-day readmission
- Flag any FAIL item as a discharge blocker with specific action needed
- For high LACE scores (>=10), recommend enhanced post-discharge interventions: nurse phone call within 48h, transitional care management visit within 7 days

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
