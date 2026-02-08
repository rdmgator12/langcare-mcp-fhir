---
name: discharge-summary-writer
description: |
  Generates comprehensive discharge summaries per CMS and Joint Commission requirements from FHIR data including admission and discharge diagnoses, hospital course, procedures, discharge medications with changes, follow-up, and patient instructions. Use when user asks to "write a discharge summary", "create DC summary", "discharge documentation", mentions "discharge paperwork", "transition of care document", or needs discharge documentation. Do NOT use for admission H&P, daily progress notes, outpatient visit notes, or transfer summaries.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: documentation
---

# Discharge Summary Writer

## Overview

Generate comprehensive discharge summaries meeting CMS Conditions of Participation (CoP) and The Joint Commission (TJC) requirements. Pull admission and discharge diagnoses, hospital course data, procedures performed, discharge medication list with changes from admission highlighted, follow-up appointments, discharge condition, and generate patient-appropriate discharge instructions. Per CMS CoP 482.24(c)(2), discharge summaries must include: reason for hospitalization, significant findings, procedures performed, treatment rendered, patient condition at discharge, patient/family instructions, and attending signature.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Encounter | Admission/discharge context, LOS | period, reasonCode, hospitalization, diagnosis |
| Patient | Demographics, language preference | name, birthDate, gender, communication |
| Condition | Admission dx, discharge dx, active problems | code, clinicalStatus, category, encounter |
| Procedure | Procedures during hospitalization | code, performedDateTime, outcome, report |
| MedicationRequest | Discharge medications vs admission meds | medicationCodeableConcept, status, intent, authoredOn, dosageInstruction |
| MedicationStatement | Pre-admission home medication list | medicationCodeableConcept, dosage, status |
| AllergyIntolerance | Allergy list for summary | code, reaction, clinicalStatus |
| DiagnosticReport | Significant findings, pathology | code, conclusion, effectiveDateTime |
| Observation | Key labs at discharge, vitals | code, value[x], effectiveDateTime |
| CarePlan | Follow-up plans, discharge instructions | activity, description, period |
| Appointment | Follow-up appointments scheduled | status, serviceType, start, participant |
| DocumentReference | Persist the discharge summary | type, content, context |

## Instructions

### Step 1: Retrieve Encounter Details

```
Tool: fhir_read
resourceType: "Encounter"
id: "[encounter-id]"
```

Extract:
- **Admission date**: `period.start`
- **Discharge date**: `period.end` (or today if discharging now)
- **Length of stay**: Calculate from period
- **Admission reason**: `reasonCode`
- **Discharge disposition**: `hospitalization.dischargeDisposition.coding[0].display`
- **Encounter diagnoses**: `diagnosis` array with `use.coding[0].code` (AD = admission, DD = discharge)

### Step 2: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract: name, DOB, age, gender, preferred language (for instruction readability), PCP from `generalPractitioner`.

### Step 3: Pull Admission and Discharge Diagnoses

**3a: Encounter-linked diagnoses**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&encounter=[encounter-id]"
```

**3b: All active conditions at discharge**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Separate into:
- **Principal discharge diagnosis**: Condition from encounter diagnosis list with rank=1 or use=DD
- **Secondary diagnoses**: Remaining encounter diagnoses
- **Chronic conditions**: Active conditions not linked to this encounter

### Step 4: Pull Procedures Performed During Hospitalization

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&encounter=[encounter-id]&status=completed"
```

For each procedure:
- Name from `code.coding[0].display`
- Date from `performedDateTime`
- Outcome from `outcome.coding[0].display` if available
- Complications from `complication` if present

### Step 5: Build Discharge Medication List with Changes

**5a: Current discharge medications**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&intent=order&_count=100"
```

**5b: Pre-admission home medications**
```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active,on-hold,completed&_count=100"
```

**5c: Stopped medications**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=stopped,cancelled&authoredon=ge[admission-date]&_count=100"
```

Compare lists and categorize each medication:
- **CONTINUE**: Same medication, same dose as home
- **NEW**: Started during hospitalization (authoredOn >= admission date, no matching home med)
- **CHANGED**: Same medication, different dose/frequency/route
- **DISCONTINUED**: On home list but stopped during admission
- **ON HOLD**: Temporarily held (e.g., metformin held for contrast)

### Step 6: Pull Significant Findings and Results

**6a: Key diagnostic reports**
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&date=ge[admission-date]&_sort=-date"
```

**6b: Discharge vitals**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=vital-signs&_sort=-date&_count=10"
```

**6c: Discharge labs (most recent)**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&_sort=-date&_count=30"
```

### Step 7: Pull Follow-Up Appointments

```
Tool: fhir_search
resourceType: "Appointment"
queryParams: "patient=[patient-id]&date=ge[discharge-date]&status=booked,proposed&_sort=date"
```

If no appointments found, flag as "No follow-up appointments scheduled -- ACTION REQUIRED."

### Step 8: Pull Care Plan and Discharge Instructions

```
Tool: fhir_search
resourceType: "CarePlan"
queryParams: "patient=[patient-id]&status=active&_sort=-date&_count=5"
```

### Step 9: Pull Allergies

```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]&clinical-status=active"
```

### Step 10: Assemble Discharge Summary

```
DISCHARGE SUMMARY
==================
Patient: [name] | MRN: [mrn] | DOB: [dob] (Age: [age]) | Sex: [gender]
Admission Date: [date] | Discharge Date: [date] | LOS: [n] days
Attending Physician: [name]
PCP: [name]

ADMISSION DIAGNOSIS
-------------------
[From encounter reasonCode or admission diagnosis]

DISCHARGE DIAGNOSES
-------------------
Principal: [diagnosis] - [ICD-10]
Secondary:
1. [diagnosis] - [ICD-10]
2. [diagnosis] - [ICD-10]

HOSPITAL COURSE
---------------
[Flag: "Hospital course narrative requires clinician input"]
[Pre-populate timeline from available data:]
- Day 1: Admitted for [reason]. Initial workup: [key labs, imaging].
- [Procedures performed with dates]
- [Significant medication changes with dates]
- [Key lab trends: admission -> discharge values]
- [Consults obtained]

PROCEDURES PERFORMED
--------------------
1. [Procedure name] - [date] - Outcome: [outcome]
2. [Procedure name] - [date] - Outcome: [outcome]

SIGNIFICANT FINDINGS
--------------------
[Key diagnostic report conclusions]
[Pathology results if applicable]

CONDITION AT DISCHARGE
----------------------
[Flag: "Clinician to specify: improved / stable / fair / guarded"]
Discharge Vitals: T [temp] | HR [hr] | BP [sys]/[dia] | RR [rr] | SpO2 [spo2]%

DISCHARGE MEDICATIONS
---------------------
[CONTINUE from home]
1. [Drug] [dose] [route] [frequency] - CONTINUE UNCHANGED

[NEW medications]
2. [Drug] [dose] [route] [frequency] - NEW: [indication]
3. [Drug] [dose] [route] [frequency] - NEW: [indication]

[CHANGED medications]
4. [Drug] [new dose] [route] [frequency] - CHANGED from [old dose]: [reason]

[DISCONTINUED medications]
5. [Drug] - DISCONTINUED: [reason]

ALLERGIES
---------
[allergy list]

FOLLOW-UP APPOINTMENTS
----------------------
1. [Provider/Specialty] - [date] [time] - [location]
2. [Provider/Specialty] - [date] [time] - [location]
[Or: "No follow-up scheduled -- SCHEDULE REQUIRED"]

PENDING RESULTS AT DISCHARGE
-----------------------------
[List any outstanding labs, pathology, cultures with expected completion]

PATIENT INSTRUCTIONS
--------------------
[Condition-specific discharge instructions]
[Activity restrictions]
[Diet instructions]
[Wound care if applicable]
[Return to ED if: warning signs]
[Flag: "Instructions should be at 5th-6th grade reading level per health literacy guidelines"]

CLINICIAN ATTESTATION
---------------------
[Flag: "Requires attending physician review and signature"]
```

### Step 11: Persist as DocumentReference

```
Tool: fhir_create
resourceType: "DocumentReference"
resource: {
  "resourceType": "DocumentReference",
  "status": "current",
  "type": {
    "coding": [{
      "system": "http://loinc.org",
      "code": "18842-5",
      "display": "Discharge summary"
    }]
  },
  "subject": {"reference": "Patient/[patient-id]"},
  "date": "[current-datetime]",
  "author": [{"reference": "Practitioner/[practitioner-id]"}],
  "content": [{
    "attachment": {
      "contentType": "text/plain",
      "data": "[base64-encoded-summary]"
    }
  }],
  "context": {
    "encounter": [{"reference": "Encounter/[encounter-id]"}],
    "period": {
      "start": "[admission-date]",
      "end": "[discharge-date]"
    }
  }
}
```

## Examples

### Example 1: CHF Exacerbation Discharge

**User says**: "Write discharge summary for patient 54321, encounter ENC-ADM-200."

**Actions**:
1. `fhir_read` Encounter/ENC-ADM-200. Returns: admitted 2024-03-15, discharging 2024-03-19, LOS 4 days, reasonCode="Acute systolic heart failure exacerbation", discharge disposition="Home".
2. `fhir_read` Patient/54321. Returns: Robert Brown, DOB 1948-11-03, Male. PCP: Dr. Williams.
3. `fhir_search` Condition for encounter. Returns: Acute on chronic systolic HF (I50.22), hyponatremia (E87.1). Active: HTN, T2DM, CKD3, afib.
4. `fhir_search` Procedure for encounter. Returns: Echocardiogram (2024-03-16), right heart catheterization (2024-03-17).
5. `fhir_search` MedicationRequest active + MedicationStatement + stopped. Comparison yields:
   - CONTINUE: apixaban 5mg BID, metoprolol 50mg BID, metformin 500mg BID
   - NEW: spironolactone 25mg daily (indication: HFrEF), sacubitril/valsartan 24/26mg BID
   - CHANGED: furosemide 40mg -> 80mg BID (from 40mg daily)
   - DISCONTINUED: lisinopril 20mg daily (replaced by sacubitril/valsartan)
6. `fhir_search` DiagnosticReport. Returns: Echo EF 30%, moderate MR. RHC: elevated filling pressures.
7. `fhir_search` Observation labs. Returns: BNP 380 (down from 1250), Na 136 (up from 133), Cr 1.5 (down from 1.8), K 4.2.
8. `fhir_search` Appointment. Returns: Cardiology follow-up 2024-03-26, PCP 2024-04-02.

**Result**:
```
DISCHARGE SUMMARY
==================
Patient: Robert Brown | MRN: MRN-54321 | DOB: 1948-11-03 (Age: 75) | Sex: Male
Admission: 2024-03-15 | Discharge: 2024-03-19 | LOS: 4 days
Attending: Dr. Carter | PCP: Dr. Williams

DISCHARGE DIAGNOSES
Principal: Acute on chronic systolic heart failure exacerbation - I50.22
Secondary:
1. Hyponatremia, resolved - E87.1
2. Hypertension - I10
3. Atrial fibrillation - I48.91
4. CKD stage 3 - N18.3
5. Type 2 diabetes - E11.9

PROCEDURES
1. Transthoracic echocardiogram - 2024-03-16
2. Right heart catheterization - 2024-03-17

SIGNIFICANT FINDINGS
- Echo: EF 30% (down from 40% in 2023), moderate mitral regurgitation
- RHC: Elevated filling pressures consistent with decompensated HF

DISCHARGE MEDICATIONS
CONTINUE: Apixaban 5mg BID, Metoprolol 50mg BID, Metformin 500mg BID
NEW: Spironolactone 25mg daily (for HFrEF), Sacubitril/valsartan 24/26mg BID (for HFrEF)
CHANGED: Furosemide 80mg PO BID (was 40mg daily - increased for volume management)
DISCONTINUED: Lisinopril 20mg daily (replaced by sacubitril/valsartan - do NOT take together)

FOLLOW-UP
1. Cardiology (Dr. Patel) - 2024-03-26 - Heart failure clinic
2. PCP (Dr. Williams) - 2024-04-02
Lab check: BMP in 1 week (monitor K and Cr on new spironolactone)
```

### Example 2: Pneumonia Discharge

**User says**: "DC summary for patient 67890, encounter ENC-MED-300."

**Actions**:
1. `fhir_read` Encounter/ENC-MED-300. Returns: admitted 2024-03-13, discharging 2024-03-17, LOS 4 days, CAP.
2. `fhir_read` Patient/67890. Returns: Alice Chen, DOB 1965-07-22, Female.
3. `fhir_search` Condition. Returns: CAP resolved, HTN, T2DM.
4. `fhir_search` MedicationRequest comparison: NEW: amoxicillin/clavulanate 875mg BID x 3 days (complete antibiotic course), azithromycin PO (transitioned from IV). CONTINUE: home meds unchanged.
5. `fhir_search` Observation discharge labs. Returns: WBC 7.2 (normalized), procalcitonin 0.15 (normalized).

**Result**: Full discharge summary with antibiotic course completion instructions, return precautions for fever/dyspnea, pneumococcal and influenza vaccination status check, PCP follow-up with repeat CXR in 6 weeks.

## Troubleshooting

### Encounter diagnosis array is empty
- Search Condition resources linked to the encounter: `fhir_search` Condition with `encounter=[encounter-id]`.
- Check `Encounter.reasonCode` as a fallback for admission diagnosis.
- Some systems populate diagnoses only in the billing system. If no structured diagnoses are found, flag: "Discharge diagnoses require clinician input -- no structured diagnosis data linked to encounter."

### MedicationStatement not available for home medication comparison
- Use the earliest MedicationRequest records (those with `authoredOn` before admission) as a proxy for the pre-admission medication list.
- Check if a medication reconciliation DocumentReference exists from admission: `fhir_search` DocumentReference with `type=http://loinc.org|11503-0` (medication reconciliation note).
- Flag medications where home vs hospital comparison could not be performed.

### No follow-up appointments found in FHIR
- Appointment resources may not be created until closer to discharge. Flag prominently: "No follow-up appointments found -- SCHEDULE REQUIRED per CMS/TJC requirements."
- Check CarePlan for planned follow-up activities that may not yet have corresponding Appointment resources.

## Related Skills

- `history-and-physical-generator` - For the admission documentation this summary references
- `progress-note-writer` - For daily notes that inform the hospital course section
- `medication-reconciliation` - For detailed medication comparison at discharge
- `transition-of-care-summary` - For C-CDA transition of care documents sent to receiving providers
