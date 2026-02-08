---
name: transition-of-care-summary
description: |
  Generates a C-CDA compliant transition of care document by pulling active problems, medications, allergies,
  recent procedures, pending results, code status, and follow-up needs from FHIR resources.
  Use when user asks to "generate TOC", "create transfer summary", "discharge summary", "handoff document",
  "transition of care", "care summary for transfer", or needs documentation for "sending patient to another facility".
  Do NOT use for discharge readiness checks, referral generation, or routine clinical summaries.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: care-coordination
---

# Transition of Care Summary

## Overview

Generate a comprehensive transition of care document containing all required C-CDA sections for patient handoff. Pull all relevant FHIR resources to create a complete clinical picture. Output a structured summary organized by Joint Commission TOC requirements. Create a DocumentReference FHIR resource linking to the generated summary.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Demographics and identifiers | name, birthDate, gender, identifier, address, telecom |
| Encounter | Current/recent encounter details | status, class, period, reasonCode, participant, hospitalization |
| Condition | Active problems and hospital diagnoses | code, clinicalStatus, verificationStatus, onsetDateTime, category |
| MedicationRequest | Discharge/current prescriptions | medicationCodeableConcept, dosageInstruction, status, intent |
| MedicationStatement | Home medication list | medicationCodeableConcept, status, dosage |
| AllergyIntolerance | Allergies and adverse reactions | code, reaction, clinicalStatus, criticality |
| Procedure | Procedures performed during stay | code, performedDateTime, status, outcome |
| Observation | Labs, vitals, functional status | code, valueQuantity, effectiveDateTime, interpretation |
| DiagnosticReport | Imaging and pathology results | code, conclusion, effectiveDateTime, status |
| Immunization | Immunization history | vaccineCode, occurrenceDateTime, status |
| CarePlan | Active care plans | status, category, activity, goal |
| Consent | Advance directives, code status | category, status, provision |
| DocumentReference | Store the generated TOC | type, content, context |
| ServiceRequest | Pending orders and referrals | status, code, intent |
| Goal | Treatment goals | lifecycleStatus, description, target |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract: full legal name, DOB, age, gender, MRN, address, phone, preferred language, emergency contacts (from `contact` array or RelatedPerson search).

### Step 2: Retrieve Encounter Information

```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&_sort=-date&_count=1"
```

Extract: encounter type (inpatient, observation, ED), admission date, discharge date (if available), reason for admission from `reasonCode`, attending provider from `participant`, discharge disposition from `hospitalization.dischargeDisposition`.

### Step 3: Retrieve Active Problem List

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Organize by:
- **Principal diagnosis**: The primary reason for the encounter
- **Active problems**: All other active conditions
- **Hospital-acquired conditions**: Conditions with onset during the encounter period
- Include SNOMED or ICD-10 codes, onset dates, and verification status

### Step 4: Retrieve Medication Lists

Discharge medications:
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&intent=order"
```

Home medication list for comparison:
```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```

For each medication, extract: name, dose, route, frequency, prescriber, start date. Flag:
- **NEW**: Medications started during encounter
- **CHANGED**: Medications with dose or frequency modifications
- **DISCONTINUED**: Medications intentionally stopped (search MedicationRequest with `status=stopped` or `status=cancelled`)
- **CONTINUED**: Medications unchanged from pre-admission

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=stopped,cancelled&encounter=[encounter-id]"
```

### Step 5: Retrieve Allergies

```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]"
```

Include: allergen, reaction type, severity, criticality. If zero results, document as "No Known Allergies (NKA)" or "Allergy status not reviewed" (these are clinically different).

### Step 6: Retrieve Procedures Performed

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&date=ge=[encounter-start-date]&status=completed"
```

Include: procedure name with code, date performed, performer, outcome/findings. For surgical procedures, include anesthesia type and complications if documented.

### Step 7: Retrieve Recent Results

Lab results:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&date=ge=[encounter-start-date]&_sort=-date"
```

Imaging results:
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&date=ge=[encounter-start-date]&_sort=-date"
```

Flag any results with status `preliminary` or `registered` as PENDING. Include interpretation flags for abnormal values.

### Step 8: Retrieve Most Recent Vitals

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=vital-signs&_sort=-date&_count=10"
```

Extract the most recent value for: BP (85354-9), HR (8867-4), RR (9279-1), Temp (8310-5), SpO2 (2708-6), Weight (29463-7).

### Step 9: Check Advance Directives and Code Status

```
Tool: fhir_search
resourceType: "Consent"
queryParams: "patient=[patient-id]&category=http://terminology.hl7.org/CodeSystem/consentcategorycodes|acd"
```

Also check for resuscitation status:
```
Tool: fhir_search
resourceType: "Consent"
queryParams: "patient=[patient-id]&category=http://terminology.hl7.org/CodeSystem/consentcategorycodes|dnr"
```

If no Consent resources found, check Condition for code status documentation (some systems store as Observation or flag).

### Step 10: Retrieve Pending Orders and Follow-up Needs

```
Tool: fhir_search
resourceType: "ServiceRequest"
queryParams: "patient=[patient-id]&status=active,draft"
```

```
Tool: fhir_search
resourceType: "CarePlan"
queryParams: "patient=[patient-id]&status=active"
```

Identify: pending labs awaiting results, scheduled follow-up appointments, home health orders, DME orders, referrals in progress.

### Step 11: Retrieve Immunization History

```
Tool: fhir_search
resourceType: "Immunization"
queryParams: "patient=[patient-id]&status=completed&_sort=-date"
```

Include recent immunizations (administered during encounter) and relevant historical immunizations (pneumococcal, influenza, COVID-19, tetanus).

### Step 12: Retrieve Functional Status

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=survey&_sort=-date&_count=5"
```

Look for: ADL assessments, mobility status, cognitive assessments (MMSE, MoCA), fall risk, Braden scale.

### Step 13: Create DocumentReference

```
Tool: fhir_create
resourceType: "DocumentReference"
resource: {
  "resourceType": "DocumentReference",
  "status": "current",
  "type": { "coding": [{ "system": "http://loinc.org", "code": "18761-7", "display": "Transfer summary note" }] },
  "category": [{ "coding": [{ "system": "http://loinc.org", "code": "18842-5", "display": "Discharge summary" }] }],
  "subject": { "reference": "Patient/[patient-id]" },
  "date": "[current-timestamp]",
  "author": [{ "reference": "Practitioner/[author-id]" }],
  "description": "Transition of Care Summary",
  "content": [{ "attachment": { "contentType": "text/plain", "data": "[base64-encoded-summary]" } }],
  "context": { "encounter": [{ "reference": "Encounter/[encounter-id]" }] }
}
```

### Step 14: Format Output

Use the I-PASS structure for the summary (see `references/handoff-frameworks.md`):

```
TRANSITION OF CARE SUMMARY
============================
Generated: [timestamp]
DocumentReference: DocumentReference/[id]

PATIENT INFORMATION
-------------------
Name: [full name] | DOB: [date] (Age: [age]) | Sex: [gender]
MRN: [mrn] | Language: [preferred language]
Emergency Contact: [name] - [relationship] - [phone]

ENCOUNTER DETAILS
-----------------
Type: [Inpatient/Observation/ED]
Admitted: [date] | Discharged: [date]
Attending: [provider name]
Admit Reason: [reason]
Discharge Disposition: [home/SNF/rehab/etc.]

I - ILLNESS SEVERITY
---------------------
Principal Diagnosis: [diagnosis] ([code])
Active Problems:
  1. [condition] - onset [date] ([code])
  2. [condition] - onset [date] ([code])

P - PATIENT SUMMARY
--------------------
Hospital Course:
[Brief narrative of what happened during the stay]

Procedures Performed:
  1. [procedure] - [date] - [outcome]
  2. [procedure] - [date] - [outcome]

Key Results:
  Labs:
    - [lab]: [value] [units] ([date]) [flag]
    - [lab]: [value] [units] ([date]) [flag]
  Imaging:
    - [study]: [key finding] ([date])

Most Recent Vitals:
  BP: [value] | HR: [value] | RR: [value] | Temp: [value] | SpO2: [value]

A - ACTION LIST
---------------
Pending Results:
  - [test]: ordered [date], result pending

Follow-up Appointments:
  - [specialty]: [date] with [provider]
  - [specialty]: [date] with [provider]

Pending Referrals:
  - [referral type]: [status]

S - SITUATION AWARENESS
-----------------------
Code Status: [Full code / DNR / DNR-DNI / POLST on file]
Advance Directive: [On file / Not on file]

Medications (Discharge):
  NEW:
    - [medication] [dose] [route] [frequency] - Reason: [indication]
  CHANGED:
    - [medication] [old dose] -> [new dose] - Reason: [why changed]
  DISCONTINUED:
    - [medication] - Reason: [why stopped]
  CONTINUED:
    - [medication] [dose] [route] [frequency]

Allergies:
  - [allergen]: [reaction] (Severity: [severity])

S - SYNTHESIS BY RECEIVER
--------------------------
Contingency Plans:
  - If [scenario], then [action]
  - If [scenario], then [action]

Diet: [restrictions]
Activity: [restrictions]
Weight: [daily weights? fluid restriction?]
Warning Signs: [symptoms requiring medical attention]

IMMUNIZATION STATUS
-------------------
  - [vaccine]: [date] [status]

FUNCTIONAL STATUS
-----------------
  Mobility: [status]
  ADLs: [status]
  Cognition: [status]
```

## Examples

### Example 1: Post-MI Transfer to Cardiac Rehab

**User says:** "Generate a transition of care summary for patient 77777 being transferred to cardiac rehab"

**Actions:**
1. `fhir_read` Patient/77777 -- returns William Thompson, 71yo Male
2. `fhir_search` Encounter (most recent) -- inpatient E-300, admitted 5 days ago for STEMI
3. `fhir_search` Condition?clinical-status=active -- STEMI, HTN, HLD, T2DM
4. `fhir_search` MedicationRequest?status=active -- aspirin, clopidogrel, atorvastatin 80mg (NEW), metoprolol 25mg (NEW), lisinopril (CONTINUED), metformin (CONTINUED)
5. `fhir_search` MedicationRequest?status=stopped -- amlodipine (DISCONTINUED, BP well controlled on new regimen)
6. `fhir_search` AllergyIntolerance -- PCN (hives), sulfa (rash)
7. `fhir_search` Procedure -- PCI with DES to LAD (day 1), echocardiogram (day 2)
8. `fhir_search` Observation?category=laboratory -- troponin peaked at 12.4 (now trending down), Cr 1.1, HbA1c 7.2%
9. `fhir_search` Observation?category=vital-signs -- BP 128/78, HR 68, SpO2 97%
10. `fhir_search` Consent -- Full code
11. `fhir_search` ServiceRequest?status=active -- cardiac rehab referral active
12. `fhir_create` DocumentReference -- TOC document created

**Result:**
Summary includes principal diagnosis (STEMI), procedure (PCI with DES to LAD), medication changes (new DAPT, statin, beta-blocker), discharge to cardiac rehab, follow-up with cardiology in 2 weeks, contingency plans for chest pain recurrence.

### Example 2: SNF Transfer for Elderly Patient Post Hip Fracture

**User says:** "Create transfer summary for patient abc-888 going to skilled nursing"

**Actions:**
1. `fhir_read` Patient/abc-888 -- returns Margaret O'Brien, 84yo Female
2. `fhir_search` Encounter -- inpatient E-450, admitted 8 days ago for hip fracture
3. `fhir_search` Condition -- right hip fracture, osteoporosis, dementia (mild), CHF, hypothyroidism
4. `fhir_search` MedicationRequest?status=active -- 12 active medications including enoxaparin (NEW), calcium/vitamin D (NEW), acetaminophen PRN (NEW)
5. `fhir_search` MedicationRequest?status=stopped -- ibuprofen (DISCONTINUED due to surgical risk + CKD)
6. `fhir_search` AllergyIntolerance -- codeine (nausea), latex
7. `fhir_search` Procedure -- right hip ORIF (day 1)
8. `fhir_search` Observation?category=laboratory -- Hgb 9.8, Cr 1.4, INR 1.0
9. `fhir_search` Observation?category=vital-signs -- stable vitals
10. `fhir_search` Consent -- DNR, advance directive on file with daughter as HCP
11. `fhir_search` Observation?category=survey -- Braden 16 (mild risk), AMT 7/10, requires 2-person assist for transfers
12. `fhir_create` DocumentReference

**Result:**
Summary includes surgical details (ORIF), VTE prophylaxis plan (enoxaparin x28 days), weight-bearing restrictions (TDWB right lower extremity), fall prevention with dementia precautions, PT/OT goals, code status (DNR), HCP contact information, wound care instructions, and SNF-specific contingency plans.

## Troubleshooting

### Encounter Resource Missing Key Fields
- Not all FHIR servers populate `hospitalization.dischargeDisposition` or `reasonCode`. Check the Condition list for the principal diagnosis and ask the user for discharge disposition if not in the Encounter resource.
- If `participant` (attending provider) is empty, check ServiceRequest or Procedure resources for a performer reference.

### Medication Reconciliation Discrepancies Between MedicationRequest and MedicationStatement
- MedicationStatement represents what the patient reports taking. MedicationRequest represents what was prescribed. Discrepancies are expected and clinically significant.
- Present both lists and flag discrepancies. The TOC should clearly indicate which list is "discharge medications" (MedicationRequest with `intent=order`, `status=active`) vs "home medications prior to admission" (MedicationStatement).

### No Advance Directive or Code Status Found
- Code status may be stored differently across systems: as Consent, as a flag on the Patient resource, as an Observation, or only in clinical notes.
- If not found in structured data, note "Code status not documented in structured data -- verify with clinical team before transfer."
- This is a Joint Commission requirement for TOC. Flag it prominently if missing.

## Related Skills

- `discharge-planning-checklist` -- run before generating TOC to ensure readiness
- `medication-reconciliation` -- detailed medication comparison and reconciliation
- `follow-up-task-generator` -- create structured follow-up tasks from the TOC action items
