---
name: history-and-physical-generator
description: |
  Generates comprehensive History and Physical (H&P) documents from FHIR data including CC, HPI, PMH, PSH, medications, allergies, family history, social history, ROS, physical exam, and assessment/plan. Use when user asks to "write an H&P", "generate history and physical", "admission documentation", "create H&P", mentions "admission note", "history and physical", or needs comprehensive intake documentation. Do NOT use for SOAP notes, progress notes, discharge summaries, or procedure documentation.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: documentation
---

# History and Physical Generator

## Overview

Generate a comprehensive History and Physical document from FHIR resources. Auto-populate all available structured data including past medical history, past surgical history, medication list, allergy list, family history, social history, and current vital signs/labs. Flag data gaps requiring clinician input. Support regulatory requirements for admission H&P per CMS CoP and TJC standards (must be completed within 24 hours of admission, updated within 30 days if pre-admission H&P exists).

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Demographics, social determinants | name, birthDate, gender, communication, extension |
| Encounter | Admission context, reason for admission | class, reasonCode, period, hospitalization |
| Condition | PMH, current diagnoses | code, clinicalStatus, onsetDateTime, category |
| Procedure | Past surgical history | code, performedDateTime, status |
| MedicationRequest | Current and home medications | medicationCodeableConcept, dosageInstruction, status |
| AllergyIntolerance | Allergies and adverse reactions | code, reaction, clinicalStatus, criticality |
| FamilyMemberHistory | Family medical history | relationship, condition, deceasedBoolean |
| Observation | Vitals, labs, social history | code, value[x], category, effectiveDateTime |
| DocumentReference | Prior H&P documents | type, content, date |
| CarePlan | Existing care plans | status, intent, category, activity |

## Instructions

### Step 1: Retrieve Admission Encounter

```
Tool: fhir_read
resourceType: "Encounter"
id: "[encounter-id]"
```

Extract:
- **Admitting diagnosis**: `reasonCode` array
- **Admission date/time**: `period.start`
- **Encounter class**: `class.code` (IMP = inpatient, OBSENC = observation)
- **Admitting physician**: `participant` where `type.coding.code` = "ADM"
- **Source of admission**: `hospitalization.admitSource`

If no encounter ID provided:
```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&class=IMP&status=in-progress&_sort=-date&_count=1"
```

### Step 2: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract: name, DOB (age), gender, preferred language, race/ethnicity from US Core extensions.

### Step 3: Pull Past Medical History (PMH)

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active,recurrence,remission&_count=100&_sort=-onset-date"
```

Categorize conditions:
- **Active chronic conditions**: clinicalStatus=active with onset > 30 days ago
- **Active acute conditions**: clinicalStatus=active with onset <= 30 days
- **Conditions in remission**: clinicalStatus=remission (e.g., cancer history)
- Include onset date for each if available from `onsetDateTime`

### Step 4: Pull Past Surgical History (PSH)

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&status=completed&_count=100&_sort=-date"
```

Filter to surgical procedures. Present chronologically with procedure date. Include:
- Procedure name from `code.coding[0].display`
- Date performed from `performedDateTime` or `performedPeriod.start`
- Relevant notes from `note`

### Step 5: Pull Current Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active,on-hold&_count=100"
```

For each medication extract:
- Drug name (prefer RxNorm display)
- Dose and frequency from `dosageInstruction`
- Route from `dosageInstruction.route`
- Prescriber from `requester`
- Intent: distinguish `order` (active prescription) from `plan` (intended)

### Step 6: Pull Allergies

```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]&clinical-status=active"
```

For each allergy:
- Substance from `code.coding[0].display`
- Reaction type from `reaction[0].manifestation[0].coding[0].display`
- Severity from `reaction[0].severity` (mild, moderate, severe)
- Criticality from `criticality` (low, high, unable-to-assess)

If zero results, explicitly state "NKDA (No Known Drug Allergies)" but flag that allergy review status should be confirmed.

### Step 7: Pull Family History

```
Tool: fhir_search
resourceType: "FamilyMemberHistory"
queryParams: "patient=[patient-id]&status=completed"
```

Organize by relationship:
- Parents: conditions, age of onset, deceased status
- Siblings: conditions
- Grandparents: conditions (if available)
- Flag hereditary risk factors (CAD, diabetes, cancer syndromes)

If no FamilyMemberHistory resources exist, flag as "Family history not documented in structured data."

### Step 8: Pull Social History

Search for social history Observations:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=social-history"
```

Key LOINC codes for social history:
- 72166-2: Tobacco smoking status
- 11367-0: History of tobacco use
- 74013-4: Alcoholic drinks per day
- 96842-0: Substance use
- 63512-8: Employment status
- 76690-7: Sexual orientation
- 71802-3: Housing status
- 82810-3: Pregnancy status (if applicable)

### Step 9: Pull Current Vital Signs

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=vital-signs&_sort=-date&_count=10"
```

Use most recent set. Map by LOINC (same as SOAP note: 8310-5, 8867-4, 9279-1, 85354-9, 29463-7, 8302-2, 59408-5, 39156-5).

### Step 10: Pull Recent Labs

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&_sort=-date&_count=50"
```

Group by panel. Flag abnormal values. Include reference ranges from `referenceRange`.

### Step 11: Check for Prior H&P

```
Tool: fhir_search
resourceType: "DocumentReference"
queryParams: "patient=[patient-id]&type=http://loinc.org|34117-2&_sort=-date&_count=1"
```

LOINC 34117-2 = History and physical note. If found within 30 days, note it as available for update rather than full rewrite (per CMS/TJC).

### Step 12: Assemble H&P Document

```
HISTORY AND PHYSICAL
====================
Patient: [name] | MRN: [mrn] | DOB: [dob] (Age: [age]) | Sex: [gender]
Admission Date: [date] | Admitting Physician: [physician]
Admitting Diagnosis: [reasonCode]

CHIEF COMPLAINT
---------------
[Encounter reasonCode text]

HISTORY OF PRESENT ILLNESS
---------------------------
[Synthesize from encounter reason, recent conditions, recent observations.
Flag: "HPI requires clinician narrative input" -- structured data provides context only.]

PAST MEDICAL HISTORY
--------------------
[Numbered list from Step 3 with onset dates]
1. [Condition] (onset [date])
2. [Condition] (onset [date])

PAST SURGICAL HISTORY
---------------------
[Numbered list from Step 4 with dates]
1. [Procedure] ([date])
2. [Procedure] ([date])

MEDICATIONS
-----------
[Numbered list from Step 5]
1. [Drug] [dose] [route] [frequency]

ALLERGIES
---------
[List from Step 6 with reactions]
1. [Substance] - [Reaction] (Severity: [severity])

FAMILY HISTORY
--------------
[Organized by relationship from Step 7]
Father: [conditions]
Mother: [conditions]

SOCIAL HISTORY
--------------
Tobacco: [status]
Alcohol: [status]
Substances: [status]
Occupation: [status]
Living Situation: [status]

REVIEW OF SYSTEMS
-----------------
[Flag: "ROS requires clinician input -- not reliably captured in structured FHIR data"]
[Pre-populate any available Observation category=survey data]

PHYSICAL EXAMINATION
--------------------
Vitals: T [temp] | HR [hr] | RR [rr] | BP [sys]/[dia] | SpO2 [spo2]% | Wt [wt] | Ht [ht] | BMI [bmi]
[Flag: "Physical exam requires clinician input"]
[Pre-populate any Observation category=exam data]

LABORATORY DATA
---------------
[From Step 10, abnormals flagged]

IMAGING
-------
[Search DiagnosticReport if available, otherwise flag as "No imaging data in structured records"]

ASSESSMENT
----------
[Primary admission diagnosis with ICD-10]
[Active problem list from Step 3]

PLAN
----
[Flag: "Plan requires clinician input"]
[Pre-populate from CarePlan if available]

DATA GAPS REQUIRING CLINICIAN INPUT
------------------------------------
- [ ] HPI narrative
- [ ] Review of Systems
- [ ] Physical Examination findings
- [ ] Assessment and Plan details
[Additional gaps identified during assembly]
```

### Step 13: Persist as DocumentReference (if authorized)

```
Tool: fhir_create
resourceType: "DocumentReference"
resource: {
  "resourceType": "DocumentReference",
  "status": "current",
  "type": {
    "coding": [{
      "system": "http://loinc.org",
      "code": "34117-2",
      "display": "History and physical note"
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

## Examples

### Example 1: Admission H&P for CHF Exacerbation

**User says**: "Generate an H&P for patient 54321, admitted today for CHF exacerbation."

**Actions**:
1. `fhir_search` Encounter for patient 54321, class=IMP, status=in-progress. Returns ENC-ADM-200, reasonCode="Acute on chronic systolic heart failure exacerbation".
2. `fhir_read` Patient/54321. Returns: Robert Brown, DOB 1948-11-03 (age 75), Male.
3. `fhir_search` Condition active. Returns: CHF (I50.22), HTN (I10), T2DM (E11.9), CKD stage 3 (N18.3), atrial fibrillation (I48.91).
4. `fhir_search` Procedure completed. Returns: CABG x3 (2015), pacemaker insertion (2020).
5. `fhir_search` MedicationRequest active. Returns: furosemide 40mg daily, lisinopril 20mg daily, metoprolol 50mg BID, apixaban 5mg BID, metformin 500mg BID.
6. `fhir_search` AllergyIntolerance. Returns: ACE inhibitor cough (lisinopril documented as tolerated despite class allergy).
7. `fhir_search` FamilyMemberHistory. Returns: Father - MI age 58 (deceased), Mother - T2DM, HTN.
8. `fhir_search` Observation social-history. Returns: Former smoker (quit 2010), alcohol 1-2 drinks/week.
9. `fhir_search` Observation vital-signs. Returns: T 98.4, HR 88 irregular, RR 22, BP 158/92, SpO2 91% on RA, Wt 230 lbs (up 15 lbs from last visit).
10. `fhir_search` Observation laboratory. Returns: BNP 1250, Cr 1.8 (baseline 1.4), K 5.1, Na 133.

**Result**:
```
HISTORY AND PHYSICAL
====================
Patient: Robert Brown | MRN: MRN-54321 | DOB: 1948-11-03 (Age: 75) | Sex: Male
Admission Date: 2024-03-15 | Admitting Diagnosis: Acute on chronic systolic HF exacerbation

CHIEF COMPLAINT: Heart failure exacerbation

PAST MEDICAL HISTORY
1. Heart failure, chronic systolic (I50.22)
2. Hypertension (I10)
3. Type 2 diabetes mellitus (E11.9)
4. Chronic kidney disease, stage 3 (N18.3)
5. Atrial fibrillation, persistent (I48.91)

PAST SURGICAL HISTORY
1. CABG x3 (2015)
2. Pacemaker insertion (2020)

MEDICATIONS
1. Furosemide 40mg PO daily
2. Lisinopril 20mg PO daily
3. Metoprolol succinate 50mg PO BID
4. Apixaban 5mg PO BID
5. Metformin 500mg PO BID

ALLERGIES
1. ACE inhibitors - Cough (note: lisinopril currently tolerated)

FAMILY HISTORY
Father: MI at age 58 (deceased)
Mother: Type 2 diabetes, hypertension

SOCIAL HISTORY
Tobacco: Former smoker, quit 2010
Alcohol: 1-2 drinks per week

LABORATORY DATA
BNP: 1250 pg/mL [CRITICAL HIGH]
Creatinine: 1.8 mg/dL [HIGH - baseline 1.4]
Potassium: 5.1 mEq/L [HIGH]
Sodium: 133 mEq/L [LOW]

DATA GAPS REQUIRING CLINICIAN INPUT
- [ ] HPI narrative (onset, exacerbating factors, orthopnea, PND, edema)
- [ ] Review of Systems
- [ ] Physical Examination (JVD, lung sounds, edema grading, S3/S4)
- [ ] Assessment and detailed Plan
```

### Example 2: Elective Surgical Admission

**User says**: "Create H&P for patient pt-777, pre-op admission for total knee replacement."

**Actions**:
1. `fhir_search` Encounter for pt-777, class=IMP, today. Returns ENC-SURG-50, reasonCode="Right total knee arthroplasty".
2. `fhir_read` Patient/pt-777. Returns: Susan Lee, DOB 1960-04-12 (age 63), Female.
3. `fhir_search` Condition active. Returns: Right knee osteoarthritis (M17.11), HTN controlled (I10), hypothyroidism (E03.9).
4. `fhir_search` Procedure completed. Returns: Left TKA (2021), cholecystectomy (2005), C-section x2.
5. `fhir_search` MedicationRequest active. Returns: amlodipine 5mg daily, levothyroxine 75mcg daily, vitamin D 2000 IU daily.
6. `fhir_search` AllergyIntolerance. Returns: codeine (nausea/vomiting).
7. `fhir_search` Observation vital-signs. Returns: normal vitals.
8. `fhir_search` Observation laboratory. Returns: CBC normal, BMP normal, PT/INR 1.0, type & screen done.

**Result**: Full H&P with pre-populated PMH, PSH, meds, allergies, labs. Flags: "Physical exam requires clinician input", "Surgical consent verification needed", "Anesthesia clearance status: check DocumentReference".

## Troubleshooting

### FamilyMemberHistory returns zero resources
- Family history is often not captured as structured FHIR data. Flag as "Family history not available in structured format -- obtain from patient interview."
- Some systems store family history in Observation resources with category=social-history. Check Observation search results for family-related LOINC codes (54114-4 = Family member health history).

### Condition resources lack onset dates
- Use `recordedDate` as a fallback for when the condition was first documented.
- If neither `onsetDateTime` nor `recordedDate` exists, present the condition without a date and note "onset unknown."
- Check `Condition.evidence` for linked Observations that may have dates.

### Social history Observations are missing or sparse
- Social history is variably captured across EMR systems. Flag each missing element explicitly in the DATA GAPS section.
- Check for SDOH (Social Determinants of Health) data in Observation with category=sdoh if the FHIR server supports US Core 5.0+.

## Related Skills

- `soap-note-generator` - For outpatient visit documentation
- `progress-note-writer` - For subsequent daily inpatient notes after the H&P
- `discharge-summary-writer` - For discharge documentation referencing the admission H&P
- `medication-reconciliation` - For detailed medication comparison at admission
