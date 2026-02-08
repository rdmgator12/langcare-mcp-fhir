---
name: discharge-planning-checklist
description: |
  Performs comprehensive discharge readiness assessment by checking pending labs, imaging, medication reconciliation,
  follow-up appointments, patient education, DME orders, and home health referrals against FHIR resources.
  Use when user asks to "check discharge readiness", "discharge checklist", "is patient ready for discharge",
  "prepare for discharge", "discharge planning", or mentions "sending patient home".
  Do NOT use for transfer-of-care summaries, transition documentation, or post-discharge follow-up generation.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: care-coordination
---

# Discharge Planning Checklist

## Overview

Assess discharge readiness by querying FHIR resources for pending orders, incomplete care plan tasks, unscheduled follow-ups, and unreconciled medications. Generate a structured checklist with pass/fail status for each CMS Condition of Participation discharge requirement. Calculate LACE readmission risk index. Create or update a discharge CarePlan resource with outstanding tasks.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| ServiceRequest | Pending labs, imaging, consults | status, intent, code, authoredOn |
| CarePlan | Discharge plan with tasks | status, intent, activity, category |
| Appointment | Scheduled follow-ups | status, serviceType, start, participant |
| MedicationRequest | Active prescriptions for reconciliation | status, medicationCodeableConcept, dosageInstruction |
| MedicationStatement | Patient-reported medications | status, medicationCodeableConcept |
| Encounter | Current admission details | status, class, period, reasonCode |
| Condition | Active problems for discharge summary | clinicalStatus, code, onsetDateTime |
| Procedure | Completed procedures during stay | status, code, performedDateTime |
| Observation | Pending lab results | status, code, valueQuantity |
| DocumentReference | Patient education materials | type, status, content |
| DeviceRequest | DME orders | status, codeCodeableConcept, intent |

## Instructions

### Step 1: Retrieve Current Encounter

```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&status=in-progress&class=http://terminology.hl7.org/CodeSystem/v3-ActCode|IMP"
```

Extract: admission date from `period.start`, reason for admission from `reasonCode`, attending provider from `participant`. If no in-progress inpatient encounter found, search for `status=finished` with most recent date.

### Step 2: Check Pending ServiceRequests

```
Tool: fhir_search
resourceType: "ServiceRequest"
queryParams: "patient=[patient-id]&status=active,draft&encounter=[encounter-id]"
```

Categorize pending requests:
- **Lab orders**: `category.coding.code` = "108252007" (SNOMED: Laboratory procedure)
- **Imaging orders**: `category.coding.code` = "363679005" (SNOMED: Imaging)
- **Consult requests**: `category.coding.code` = "11429006" (SNOMED: Consultation)
- **Referrals**: `category.coding.code` = "3457005" (SNOMED: Patient referral)

Flag any active or draft ServiceRequest as a discharge blocker.

### Step 3: Check Pending Lab Results

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&status=preliminary,registered&category=laboratory"
```

Any Observation with status `preliminary` or `registered` indicates pending results. Flag as discharge blocker if ordered during current encounter.

### Step 4: Verify Medication Reconciliation

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&encounter=[encounter-id]"
```

Also retrieve:
```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```

Compare MedicationRequest (inpatient orders) against MedicationStatement (home medications):
- Identify new medications added during admission
- Identify home medications held or discontinued
- Flag unreconciled discrepancies (medication in one list but not addressed in the other)
- Check that discharge prescriptions exist for all intended outpatient medications

### Step 5: Check Follow-up Appointments

```
Tool: fhir_search
resourceType: "Appointment"
queryParams: "patient=[patient-id]&status=booked,proposed&date=ge[today]"
```

Verify at minimum:
- PCP follow-up within 7 days (14 days acceptable for low-risk)
- Specialist follow-up if condition warrants (e.g., cardiology after MI, surgery after procedure)
- Lab recheck appointment if indicated

Flag as incomplete if no future appointments found.

### Step 6: Review CarePlan for Discharge Tasks

```
Tool: fhir_search
resourceType: "CarePlan"
queryParams: "patient=[patient-id]&status=active&category=http://snomed.info/sct|58000006"
```

SNOMED 58000006 = Discharge planning. Check `activity` array for incomplete tasks. If no discharge CarePlan exists, create one in Step 9.

### Step 7: Check Patient Education Documentation

```
Tool: fhir_search
resourceType: "DocumentReference"
queryParams: "patient=[patient-id]&type=http://loinc.org|69981-9&date=ge=[admission-date]"
```

LOINC 69981-9 = Patient education note. Verify education documented for:
- Primary diagnosis
- New medications (purpose, dosing, side effects)
- Activity restrictions
- Warning signs requiring return to ED
- Follow-up instructions

### Step 8: Check DME and Home Health Orders

```
Tool: fhir_search
resourceType: "DeviceRequest"
queryParams: "patient=[patient-id]&status=active,draft&encounter=[encounter-id]"
```

Also check for home health referrals:
```
Tool: fhir_search
resourceType: "ServiceRequest"
queryParams: "patient=[patient-id]&category=http://snomed.info/sct|385763009&status=active,draft"
```

SNOMED 385763009 = Home health care. Verify all DME and home health orders have been placed, not just drafted.

### Step 9: Calculate LACE Readmission Risk Index

Gather data for LACE score calculation:
- **L** (Length of stay): Calculate from `Encounter.period.start` to today
- **A** (Acuity of admission): Check if admission was via ED (`Encounter.hospitalization.admitSource`)
- **C** (Comorbidities): Count via Charlson comorbidity conditions from active Conditions
- **E** (ED visits): Count Encounter resources with `class` = "EMER" in prior 6 months

```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&class=http://terminology.hl7.org/CodeSystem/v3-ActCode|EMER&date=ge[6-months-ago]"
```

Score interpretation: 0-4 Low risk, 5-9 Moderate risk, 10+ High risk. See `references/lace-index.md` for detailed scoring.

### Step 10: Generate or Update Discharge CarePlan

If CarePlan exists from Step 6:
```
Tool: fhir_update
resourceType: "CarePlan"
id: "[careplan-id]"
resource: {
  "status": "active",
  "intent": "plan",
  "subject": { "reference": "Patient/[patient-id]" },
  "encounter": { "reference": "Encounter/[encounter-id]" },
  "category": [{ "coding": [{ "system": "http://snomed.info/sct", "code": "58000006", "display": "Discharge planning" }] }],
  "activity": [
    { "detail": { "status": "[completed|in-progress|not-started]", "description": "Pending lab results reviewed" } },
    { "detail": { "status": "[completed|in-progress|not-started]", "description": "Medication reconciliation completed" } },
    { "detail": { "status": "[completed|in-progress|not-started]", "description": "Follow-up appointments scheduled" } },
    { "detail": { "status": "[completed|in-progress|not-started]", "description": "Patient education completed" } },
    { "detail": { "status": "[completed|in-progress|not-started]", "description": "DME ordered" } },
    { "detail": { "status": "[completed|in-progress|not-started]", "description": "Home health referral placed" } }
  ]
}
```

If no CarePlan exists, use `fhir_create` with the same payload (omit `id`).

### Step 11: Format Output

```
DISCHARGE READINESS CHECKLIST
==============================
Patient: [name] | MRN: [mrn] | Admission: [date] | LOS: [days] days
LACE Score: [score] ([Low/Moderate/High] readmission risk)

[PASS] / [FAIL] Pending Labs/Imaging
  - [list pending items or "All results finalized"]

[PASS] / [FAIL] Medication Reconciliation
  - New: [count] medications added
  - Discontinued: [count] medications stopped
  - Unreconciled: [count] discrepancies

[PASS] / [FAIL] Follow-up Appointments
  - PCP: [date] with [provider] or "NOT SCHEDULED"
  - Specialist: [date] with [provider] or "NOT SCHEDULED"

[PASS] / [FAIL] Patient Education
  - [list topics documented or "NOT DOCUMENTED"]

[PASS] / [FAIL] DME Orders
  - [list items or "No DME required" or "PENDING"]

[PASS] / [FAIL] Home Health Referral
  - [status or "Not indicated" or "PENDING"]

OVERALL: [READY / NOT READY - X items require attention]
```

## Examples

### Example 1: Patient Ready for Discharge

**User says:** "Is patient 67890 ready for discharge?"

**Actions:**
1. `fhir_search` Encounter?patient=67890&status=in-progress&class=IMP -- returns encounter E-100, admitted 3 days ago for pneumonia
2. `fhir_search` ServiceRequest?patient=67890&status=active,draft&encounter=E-100 -- returns 0 (all orders completed)
3. `fhir_search` Observation?patient=67890&status=preliminary,registered&category=laboratory -- returns 0 (all results final)
4. `fhir_search` MedicationRequest?patient=67890&status=active&encounter=E-100 -- returns 3 inpatient meds
5. `fhir_search` MedicationStatement?patient=67890&status=active -- returns 5 home meds, all reconciled
6. `fhir_search` Appointment?patient=67890&status=booked&date=ge2024-01-15 -- returns PCP in 5 days, pulmonology in 14 days
7. `fhir_search` CarePlan?patient=67890&status=active&category=58000006 -- returns active discharge plan, all activities completed
8. `fhir_search` DocumentReference?patient=67890&type=69981-9 -- returns education note for pneumonia, antibiotics
9. Calculate LACE: L=3, A=1 (ED admit), C=2, E=0 -- Score: 6 (Moderate)

**Result:**
```
DISCHARGE READINESS CHECKLIST
==============================
Patient: Robert Chen | MRN: 67890 | Admission: 2024-01-12 | LOS: 3 days
LACE Score: 6 (Moderate readmission risk)

[PASS] Pending Labs/Imaging - All results finalized
[PASS] Medication Reconciliation - 3 new (levofloxacin, guaifenesin, albuterol), 0 discontinued, 0 unreconciled
[PASS] Follow-up Appointments - PCP: 2024-01-20, Pulmonology: 2024-01-29
[PASS] Patient Education - Pneumonia care, antibiotics, return precautions documented
[PASS] DME Orders - No DME required
[PASS] Home Health Referral - Not indicated

OVERALL: READY - All discharge criteria met
```

### Example 2: Patient Not Ready -- Multiple Blockers

**User says:** "Run discharge checklist for patient abc-222"

**Actions:**
1. `fhir_search` Encounter -- returns E-500, admitted 7 days ago for CHF exacerbation
2. `fhir_search` ServiceRequest?status=active,draft -- returns 2: pending echocardiogram, pending cardiology consult
3. `fhir_search` Observation?status=preliminary -- returns 1: BNP result pending
4. `fhir_search` MedicationRequest -- returns 6 inpatient meds
5. `fhir_search` MedicationStatement -- returns 4 home meds, 2 unreconciled (new carvedilol, dose change on furosemide)
6. `fhir_search` Appointment?status=booked -- returns 0 future appointments
7. `fhir_search` CarePlan -- returns plan with 2/6 activities completed
8. `fhir_search` DocumentReference -- returns 0 education documents
9. LACE: L=7, A=1, C=4 (CHF, DM, CKD, HTN), E=2 -- Score: 14 (High)

**Result:**
```
DISCHARGE READINESS CHECKLIST
==============================
Patient: Dorothy Williams | MRN: abc-222 | Admission: 2024-01-08 | LOS: 7 days
LACE Score: 14 (High readmission risk)

[FAIL] Pending Labs/Imaging
  - Echocardiogram: ordered, not completed
  - BNP (LOINC 42637-9): result pending
  - Cardiology consult: requested, not completed

[FAIL] Medication Reconciliation
  - New: carvedilol 12.5mg BID (not yet reconciled with home list)
  - Changed: furosemide 20mg -> 40mg BID (dose change not documented)
  - 2 unreconciled discrepancies

[FAIL] Follow-up Appointments
  - PCP: NOT SCHEDULED
  - Cardiology: NOT SCHEDULED
  - Lab recheck (BMP for electrolytes): NOT SCHEDULED

[FAIL] Patient Education
  - NOT DOCUMENTED (CHF self-management, daily weights, sodium restriction, medication changes)

[PASS] DME Orders - Home scale for daily weights ordered
[FAIL] Home Health Referral - PENDING (draft order exists, not finalized)

OVERALL: NOT READY - 5 items require attention
```

## Troubleshooting

### No Inpatient Encounter Found
- Patient may be in observation status. Search with `class=http://terminology.hl7.org/CodeSystem/v3-ActCode|OBSENC` for observation encounters.
- Some systems use `status=planned` for encounters not yet started. Broaden search: `status=in-progress,planned,arrived`.

### ServiceRequest Search Returns Orders From Prior Encounters
- Always filter by `encounter=[encounter-id]` to scope to current admission.
- If encounter ID is unavailable, filter by `authored=ge[admission-date]` to approximate.

### Medication Reconciliation Shows False Discrepancies
- Some systems use MedicationRequest for both inpatient and outpatient orders. Check `intent` field: `order` = active prescription, `plan` = intended but not yet ordered.
- Brand vs generic name mismatches are common. Compare by RxNorm code (`medicationCodeableConcept.coding` where `system` = "http://www.nlm.nih.gov/research/umls/rxnorm") rather than display text.

## Related Skills

- `transition-of-care-summary` -- generate the actual discharge/transfer document after checklist passes
- `medication-reconciliation` -- detailed medication reconciliation workflow
- `follow-up-task-generator` -- create Task resources for post-discharge follow-up items
