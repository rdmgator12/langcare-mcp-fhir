---
name: soap-note-generator
description: |
  Generates structured SOAP notes from FHIR encounter data including chief complaint, vitals, labs, medications, conditions, and procedures. Use when user asks to "write a SOAP note", "generate encounter note", "document this visit", "create a clinic note", mentions "SOAP", "visit note", or needs structured clinical documentation for a completed or in-progress encounter. Do NOT use for inpatient progress notes, H&P documents, discharge summaries, or procedure notes.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: documentation
---

# SOAP Note Generator

## Overview

Generate structured SOAP (Subjective, Objective, Assessment, Plan) notes from FHIR encounter data. Pull chief complaint, history of present illness context, vital signs, laboratory results, active medications, active conditions, and procedures performed during the encounter. Support specialty-specific formatting for primary care, emergency department, and inpatient encounters. Optionally create a DocumentReference resource to persist the generated note.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Encounter | Visit context, chief complaint, type | status, class, type, reasonCode, period, participant |
| Condition | Active problems, encounter diagnoses | code, clinicalStatus, verificationStatus, encounter |
| Observation | Vitals, labs, clinical findings | code, value[x], effectiveDateTime, category |
| MedicationRequest | Current medications, new prescriptions | medicationCodeableConcept, status, authoredOn, dosageInstruction |
| Procedure | Procedures performed during encounter | code, status, performedDateTime, outcome |
| AllergyIntolerance | Allergy list for note context | code, clinicalStatus, reaction |
| Patient | Demographics for note header | name, birthDate, gender, identifier |
| DocumentReference | Persist the generated note | type, content, context, date |

## Instructions

### Step 1: Retrieve Encounter Context

```
Tool: fhir_read
resourceType: "Encounter"
id: "[encounter-id]"
```

Extract:
- **Chief complaint**: `reasonCode[0].text` or `reasonCode[0].coding[0].display`
- **Visit type**: `type[0].coding[0].display` (e.g., Office Visit, ED Visit, Follow-up)
- **Encounter class**: `class.code` (AMB = ambulatory, EMER = emergency, IMP = inpatient)
- **Period**: `period.start` and `period.end`
- **Attending**: `participant` where `type.coding.code` = "ATND"

If no encounter ID is provided, search by patient and date:
```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&date=[YYYY-MM-DD]&_sort=-date&_count=1"
```

### Step 2: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract name, DOB (calculate age), gender, MRN for note header.

### Step 3: Pull Vital Signs

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=vital-signs&date=ge[encounter-start]&date=le[encounter-end]&_sort=-date"
```

Map vital signs by LOINC code:
- 8310-5: Body temperature
- 8867-4: Heart rate
- 9279-1: Respiratory rate
- 85354-9: Blood pressure panel (components: 8480-6 systolic, 8462-4 diastolic)
- 29463-7: Body weight
- 8302-2: Body height
- 59408-5: SpO2 (pulse oximetry)
- 39156-5: BMI

### Step 4: Pull Laboratory Results

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&date=ge[encounter-start]&_sort=-date&_count=50"
```

Group lab results by panel. Present abnormal values with reference ranges. Flag critical values.

### Step 5: Pull Active Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

Separate into:
- **Pre-existing medications**: `authoredOn` before encounter start
- **New prescriptions**: `authoredOn` during encounter period
- **Changed medications**: status changes during encounter

### Step 6: Pull Encounter Diagnoses and Active Conditions

**6a: Encounter-specific diagnoses**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&encounter=[encounter-id]"
```

**6b: Active problem list**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

### Step 7: Pull Procedures Performed

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&encounter=[encounter-id]"
```

### Step 8: Pull Allergies

```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]&clinical-status=active"
```

### Step 9: Assemble SOAP Note

Structure the note based on encounter class:

**For ambulatory (AMB) encounters:**

```
SOAP NOTE
=========
Patient: [name] | MRN: [mrn] | DOB: [dob] (Age: [age]) | Sex: [gender]
Date: [encounter-date] | Provider: [attending]
Visit Type: [type] | Encounter: [encounter-id]

SUBJECTIVE
----------
Chief Complaint: [reasonCode text]
HPI: [Synthesize from encounter reason, active conditions relevant to visit]
ROS: [Derive from Observation resources with category=survey if available]
Current Medications: [List with dose/frequency]
Allergies: [List with reaction type]

OBJECTIVE
---------
Vitals: T [temp] | HR [hr] | RR [rr] | BP [sys]/[dia] | SpO2 [spo2]% | Wt [wt] | Ht [ht] | BMI [bmi]
[Physical exam findings from Observation category=exam if available]
Labs: [Grouped by panel, abnormals flagged]

ASSESSMENT
----------
[Numbered problem list with ICD-10 codes from Condition.code]
1. [Primary diagnosis] - [ICD-10]
2. [Secondary diagnosis] - [ICD-10]

PLAN
----
[Per problem:]
1. [Diagnosis]: [Actions - new meds, referrals, follow-up]
2. [Diagnosis]: [Actions]
Procedures Performed: [List from Step 7]
Follow-up: [From CarePlan if available]
```

**For emergency (EMER) encounters**, add: triage acuity, time-stamped reassessments, disposition.

**For inpatient (IMP) encounters**, add: admitting diagnosis, hospital day number, consults.

### Step 10: Persist Note as DocumentReference (if authorized)

```
Tool: fhir_create
resourceType: "DocumentReference"
resource: {
  "resourceType": "DocumentReference",
  "status": "current",
  "type": {
    "coding": [{
      "system": "http://loinc.org",
      "code": "11506-3",
      "display": "Progress note"
    }]
  },
  "subject": {"reference": "Patient/[patient-id]"},
  "date": "[current-datetime]",
  "author": [{"reference": "Practitioner/[practitioner-id]"}],
  "content": [{
    "attachment": {
      "contentType": "text/plain",
      "data": "[base64-encoded-note]"
    }
  }],
  "context": {
    "encounter": [{"reference": "Encounter/[encounter-id]"}]
  }
}
```

Use LOINC codes for document type based on setting:
- 11506-3: Progress note
- 34111-5: Emergency department note
- 34117-2: History and physical note
- 18842-5: Discharge summary

## Examples

### Example 1: Primary Care Office Visit

**User says**: "Write a SOAP note for encounter ENC-5678 for patient 12345."

**Actions**:
1. `fhir_read` Encounter/ENC-5678. Returns: AMB class, type=Office Visit, reasonCode="Annual physical", period 2024-03-15T09:00 to 2024-03-15T09:30, attending=Dr. Smith.
2. `fhir_read` Patient/12345. Returns: John Doe, DOB 1970-05-20, Male, MRN-12345.
3. `fhir_search` Observation with category=vital-signs for encounter window. Returns: T 98.6, HR 72, RR 16, BP 142/88, SpO2 99%, Wt 210 lbs, Ht 70 in, BMI 30.1.
4. `fhir_search` Observation with category=laboratory. Returns: HbA1c 7.2%, lipid panel (TC 240, LDL 155, HDL 38, TG 235).
5. `fhir_search` MedicationRequest active. Returns: metformin 1000mg BID, lisinopril 10mg daily.
6. `fhir_search` Condition for encounter + active. Returns: encounter dx: HTN uncontrolled, T2DM. Active: HTN, T2DM, obesity.
7. `fhir_search` AllergyIntolerance. Returns: penicillin (rash).

**Result**:
```
SOAP NOTE
=========
Patient: John Doe | MRN: MRN-12345 | DOB: 1970-05-20 (Age: 53) | Sex: Male
Date: 2024-03-15 | Provider: Dr. Smith
Visit Type: Office Visit | Encounter: ENC-5678

SUBJECTIVE
----------
Chief Complaint: Annual physical
Current Medications: Metformin 1000mg BID, Lisinopril 10mg daily
Allergies: Penicillin (rash)

OBJECTIVE
---------
Vitals: T 98.6 | HR 72 | RR 16 | BP 142/88 | SpO2 99% | Wt 210 lbs | Ht 70 in | BMI 30.1
Labs:
  HbA1c: 7.2% [HIGH - goal <7%]
  Total Cholesterol: 240 [HIGH]
  LDL: 155 [HIGH]
  HDL: 38 [LOW]
  Triglycerides: 235 [HIGH]

ASSESSMENT
----------
1. Hypertension, uncontrolled - I10
2. Type 2 diabetes mellitus, above goal - E11.65
3. Hyperlipidemia, mixed - E78.2
4. Obesity, BMI 30.1 - E66.01

PLAN
----
1. HTN: Increase lisinopril to 20mg daily. Recheck BP in 2 weeks.
2. T2DM: Continue metformin. Reinforce dietary counseling. Recheck HbA1c in 3 months.
3. Hyperlipidemia: Start atorvastatin 20mg daily. Fasting lipid panel in 6 weeks.
4. Obesity: Discussed diet and exercise. Referral to nutrition.
Follow-up: 3 months or sooner if BP not improved at 2-week check.
```

### Example 2: Emergency Department Visit

**User says**: "Generate SOAP note for the ED visit, patient ID pt-9999, today's encounter."

**Actions**:
1. `fhir_search` Encounter for patient pt-9999, today, class=EMER. Returns: encounter ENC-ED-100, reasonCode="Chest pain", period started 2024-03-15T14:22.
2. `fhir_read` Patient/pt-9999. Returns: Mary Johnson, DOB 1952-08-10, Female.
3. `fhir_search` Observation vital-signs during encounter. Returns: T 98.2, HR 92, RR 20, BP 168/95, SpO2 97%.
4. `fhir_search` Observation laboratory. Returns: Troponin I 0.02 (normal), BMP normal, CBC normal, D-dimer 0.3 (normal).
5. `fhir_search` Procedure for encounter. Returns: 12-lead ECG performed.
6. `fhir_search` Condition for encounter. Returns: Chest pain NOS (R07.9), HTN (I10).

**Result**:
```
SOAP NOTE - EMERGENCY DEPARTMENT
=================================
Patient: Mary Johnson | MRN: MRN-9999 | DOB: 1952-08-10 (Age: 71) | Sex: Female
Date: 2024-03-15 14:22 | Encounter: ENC-ED-100

SUBJECTIVE
----------
Chief Complaint: Chest pain
HPI: 71-year-old female presenting with substernal chest pain onset 2 hours prior.

OBJECTIVE
---------
Vitals: T 98.2 | HR 92 | RR 20 | BP 168/95 | SpO2 97%
Labs:
  Troponin I: 0.02 ng/mL [NORMAL]
  BMP: Within normal limits
  CBC: Within normal limits
  D-dimer: 0.3 mcg/mL [NORMAL]
Procedures: 12-lead ECG performed

ASSESSMENT
----------
1. Chest pain, unspecified - R07.9
2. Hypertension - I10

PLAN
----
1. Chest pain: Serial troponins q3h. Cardiology consult if troponin rises. Chest pain observation protocol.
2. HTN: Elevated in setting of pain/anxiety. Monitor.
Disposition: Observation pending serial troponins.
```

## Troubleshooting

### Encounter reasonCode is empty or missing chief complaint
- Check `Encounter.type` for visit type context.
- Search Condition resources linked to the encounter: `fhir_search` Condition with `encounter=[encounter-id]`. Use the primary diagnosis as a proxy for chief complaint.
- Some systems store chief complaint in an Observation with LOINC code 8661-1 (Chief complaint - Reported). Search: `fhir_search` Observation with `patient=[id]&code=8661-1&encounter=[encounter-id]`.

### Vital signs search returns no results for encounter window
- Widen the date range by 1 hour on each side of the encounter period.
- Try searching without date filter and take the most recent set: `patient=[id]&category=vital-signs&_sort=-date&_count=10`.
- Some systems store vitals under the encounter reference directly: try `encounter=[encounter-id]` as a search parameter on Observation.

### Lab results return as referenced Observation groups rather than individual values
- If an Observation has `hasMember` references, follow each reference with `fhir_read` to retrieve individual result values.
- Panel-level Observations (e.g., CBC panel LOINC 58410-2) contain component results; check `Observation.component` array for individual values.

## Related Skills

- `history-and-physical-generator` - For comprehensive H&P documentation on admission
- `progress-note-writer` - For daily inpatient progress notes
- `clinical-summary-generator` - For broader clinical summary beyond a single encounter
- `problem-list-review` - For validating the assessment problem list
