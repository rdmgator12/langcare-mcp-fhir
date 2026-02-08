---
name: progress-note-writer
description: |
  Generates daily inpatient progress notes from FHIR data including overnight events, subjective status, vitals trends, I/O, labs, imaging, problem-based assessment, and plan. Use when user asks to "write a progress note", "daily note", "rounding note", "write today's note", mentions "inpatient note", "hospital day", or needs structured daily documentation for hospitalized patients. Do NOT use for admission H&P, discharge summaries, SOAP notes for outpatient visits, or procedure documentation.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: documentation
---

# Progress Note Writer

## Overview

Generate structured daily inpatient progress notes by pulling 24-hour data windows from FHIR resources. Assemble vitals trends, intake and output, laboratory results, imaging results, medication changes, and overnight events into a problem-oriented format. Include standing order status for DVT prophylaxis, GI prophylaxis, diet, activity level, and code status. Flag critical values and significant changes from prior day.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Encounter | Admission context, hospital day calculation | period.start, status, class |
| Patient | Demographics for note header | name, birthDate, gender, identifier |
| Observation | Vitals, labs, I&O, Glasgow score | code, value[x], effectiveDateTime, category |
| MedicationRequest | Active medications, new orders, changes | medicationCodeableConcept, status, authoredOn, dosageInstruction |
| MedicationAdministration | Medications actually given in 24h | medicationCodeableConcept, effectiveDateTime, dosage |
| DiagnosticReport | Imaging, pathology results | code, conclusion, effectiveDateTime, result |
| Condition | Active problem list for assessment | code, clinicalStatus |
| Procedure | Procedures in last 24h | code, performedDateTime, status |
| CarePlan | Diet, activity, DVT/GI prophylaxis | category, activity, status |

## Instructions

### Step 1: Retrieve Encounter and Calculate Hospital Day

```
Tool: fhir_read
resourceType: "Encounter"
id: "[encounter-id]"
```

Calculate hospital day: `(today - period.start) + 1`. Extract admitting diagnosis from `reasonCode`.

If encounter ID unknown:
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

### Step 3: Pull 24-Hour Vital Signs Trends

Define the 24-hour window: `[today 00:00]` to `[now]`, or `[yesterday same-time]` to `[now]`.

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=vital-signs&date=ge[24h-ago]&_sort=date&_count=100"
```

Present as trends:
- **Temperature**: Tmax, Tcurrent (flag if Tmax >= 38.0 C / 100.4 F)
- **Heart rate**: Range (low-high), current
- **Blood pressure**: Range, current (flag if MAP < 65 or SBP > 180)
- **Respiratory rate**: Range, current
- **SpO2**: Range, current, oxygen delivery method if available
- **Pain score** (LOINC 72514-3): Most recent

### Step 4: Pull Intake and Output

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=http://loinc.org|9187-6&date=ge[24h-ago]"
```

LOINC codes for I&O:
- 9187-6: Fluid intake 24h
- 9192-6: Fluid output 24h
- Also search: `code=http://loinc.org|9192-6`

Calculate net fluid balance. Flag if net positive > 1L or net negative > 2L.

If I&O Observations not available, note: "I&O data not captured in structured FHIR observations -- obtain from nursing flowsheet."

### Step 5: Pull Laboratory Results (24h)

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&date=ge[24h-ago]&_sort=-date&_count=50"
```

Group by panel. For each result:
- Compare to prior value if available (trend direction)
- Flag abnormal values with reference range
- Highlight critical values (e.g., K < 3.0 or > 6.0, Na < 125 or > 155, Hgb < 7, platelets < 50k, lactate > 4)

### Step 6: Pull Imaging and Diagnostic Reports (24h)

```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&date=ge[24h-ago]&_sort=-date"
```

Extract `conclusion` or `presentedForm` text. If results reference Observation resources in `result`, follow references with `fhir_read`.

### Step 7: Pull Active Medications and 24h Changes

**7a: Current active orders**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=100"
```

**7b: Medications administered in 24h**
```
Tool: fhir_search
resourceType: "MedicationAdministration"
queryParams: "patient=[patient-id]&effective-time=ge[24h-ago]&_count=100"
```

Identify:
- New medications started in 24h (`authoredOn` within window)
- Dose changes
- Discontinued medications (`status=stopped` with `authoredOn` in window)
- PRN medications administered with frequency

### Step 8: Pull Procedures Performed in 24h

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&date=ge[24h-ago]"
```

### Step 9: Pull Active Problem List

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

### Step 10: Check Standing Orders and Prophylaxis

Search for prophylaxis orders:

**DVT prophylaxis:**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&code=http://www.nlm.nih.gov/research/umls/rxnorm|67108"
```
RxNorm 67108 = enoxaparin. Also check: 11289 (heparin).

**GI prophylaxis:**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&code=http://www.nlm.nih.gov/research/umls/rxnorm|40790"
```
RxNorm 40790 = pantoprazole. Also check: 29046 (famotidine).

If prophylaxis not found, flag as gap.

**Diet and activity**: Check CarePlan resources:
```
Tool: fhir_search
resourceType: "CarePlan"
queryParams: "patient=[patient-id]&status=active&category=http://hl7.org/fhir/us/core/CodeSystem/careplan-category|assess-plan"
```

### Step 11: Assemble Progress Note

```
DAILY PROGRESS NOTE
====================
Patient: [name] | MRN: [mrn] | DOB: [dob] (Age: [age])
Date: [today] | Hospital Day #[n] | Admitting Dx: [diagnosis]
Service: [service] | Attending: [attending]

OVERNIGHT EVENTS
----------------
[Flag: "Obtain from nursing/overnight team if not in structured data"]
[Include any 24h procedures, code/rapid response events, new consults]

SUBJECTIVE
----------
[Flag: "Obtain from patient interview"]
Pain: [score from LOINC 72514-3 if available]

24-HOUR VITALS
--------------
Tmax: [max temp] | Tcurrent: [current]
HR: [range] ([current]) | BP: [range] ([current])
RR: [range] ([current]) | SpO2: [range]% on [O2 device]

INTAKE & OUTPUT (24h)
---------------------
Intake: [total] mL | Output: [total] mL | Net: [+/-] mL
[Flag if significantly positive or negative]

LABORATORIES
------------
[Panel groupings with trend arrows]
BMP: Na [val][trend] | K [val][trend] | Cr [val][trend] | BUN [val][trend] | Gluc [val][trend]
CBC: WBC [val][trend] | Hgb [val][trend] | Plt [val][trend]
[Additional labs]
[CRITICAL VALUES flagged]

IMAGING / STUDIES
-----------------
[From DiagnosticReport conclusions]

MEDICATIONS
-----------
Active Medications: [count]
New (24h): [list]
Changed (24h): [list]
Discontinued (24h): [list]
PRN Given: [list with frequency]

ASSESSMENT & PLAN (by problem)
------------------------------
1. [Primary problem] - [ICD-10]
   - Status: [improving/stable/worsening based on objective data]
   - Plan: [Flag: "Clinician input required"]

2. [Secondary problem] - [ICD-10]
   - Status: [assessment]
   - Plan: [Flag: "Clinician input required"]

STANDING ORDERS
---------------
DVT Prophylaxis: [medication or "NOT ORDERED - ACTION REQUIRED"]
GI Prophylaxis: [medication or "Not indicated" or "NOT ORDERED - verify"]
Diet: [from CarePlan or "Verify"]
Activity: [from CarePlan or "Verify"]
Code Status: [from Consent or "Verify"]

DISPOSITION
-----------
Estimated discharge: [if available from Encounter.hospitalization]
Barriers to discharge: [Flag: "Clinician input"]
```

## Examples

### Example 1: Medical Floor Progress Note

**User says**: "Write today's progress note for patient 67890, admitted for pneumonia."

**Actions**:
1. `fhir_search` Encounter for patient 67890, class=IMP, status=in-progress. Returns ENC-MED-300, admitted 2024-03-13, reasonCode="Community-acquired pneumonia". Today is 2024-03-15 = Hospital Day #3.
2. `fhir_read` Patient/67890. Returns: Alice Chen, DOB 1965-07-22, Female.
3. `fhir_search` Observation vital-signs 24h. Returns: Tmax 38.2C (overnight), Tcurrent 37.1C, HR 78-92 (current 82), BP 118-132/68-78, RR 16-20, SpO2 94-97% on 2L NC.
4. `fhir_search` Observation I&O. Returns: Intake 2400mL, Output 1800mL, Net +600mL.
5. `fhir_search` Observation laboratory 24h. Returns: WBC 11.2 (down from 15.8), procalcitonin 0.8 (down from 2.1), BMP normal.
6. `fhir_search` DiagnosticReport 24h. Returns: CXR - "Right lower lobe consolidation, improved from prior."
7. `fhir_search` MedicationRequest active. Returns: ceftriaxone 1g IV daily (day 3), azithromycin 500mg PO daily (day 3), enoxaparin 40mg SQ daily, pantoprazole 40mg IV daily, acetaminophen 650mg PO q6h PRN.
8. `fhir_search` Condition active. Returns: CAP (J18.9), HTN (I10), T2DM (E11.9).

**Result**:
```
DAILY PROGRESS NOTE
====================
Patient: Alice Chen | MRN: MRN-67890 | DOB: 1965-07-22 (Age: 58)
Date: 2024-03-15 | Hospital Day #3 | Admitting Dx: Community-acquired pneumonia

24-HOUR VITALS
Tmax: 38.2C (overnight) -> Tcurrent: 37.1C [IMPROVING - fever curve trending down]
HR: 78-92 (82) | BP: 118-132/68-78 | RR: 16-20 | SpO2: 94-97% on 2L NC

I&O: Intake 2400mL | Output 1800mL | Net +600mL

LABORATORIES
WBC: 11.2 [DOWN from 15.8 - improving]
Procalcitonin: 0.8 [DOWN from 2.1 - improving]
BMP: Within normal limits

IMAGING
CXR (today): RLL consolidation, improved from prior

ASSESSMENT & PLAN (by problem)
1. Community-acquired pneumonia (J18.9) - IMPROVING
   - Fever curve trending down, WBC and procalcitonin improving, CXR improved
   - Antibiotics: Ceftriaxone + Azithromycin day 3 of planned 5-day course
   - Continue 2L NC, wean O2 as tolerated

2. Hypertension (I10) - STABLE
   - BP within range on home medications

3. Type 2 diabetes (E11.9) - STABLE
   - Monitor fingerstick glucose

STANDING ORDERS
DVT Prophylaxis: Enoxaparin 40mg SQ daily [ACTIVE]
GI Prophylaxis: Pantoprazole 40mg IV daily [ACTIVE]
Diet: Verify
Activity: Verify
Code Status: Verify
```

### Example 2: Post-Surgical Progress Note

**User says**: "Progress note for patient pt-444, POD 1 from appendectomy."

**Actions**:
1. `fhir_search` Encounter for pt-444, IMP, in-progress. Returns ENC-SURG-88, admitted 2024-03-14.
2. `fhir_read` Patient/pt-444. Returns: David Kim, DOB 2000-01-15, Male.
3. `fhir_search` Procedure 24h. Returns: Laparoscopic appendectomy, completed 2024-03-14.
4. `fhir_search` Observation vital-signs 24h. Returns: Tmax 37.8C, current normal, HR 70-85, BP normal, SpO2 99% RA.
5. `fhir_search` Observation laboratory 24h. Returns: WBC 9.8 (down from 18.2 pre-op), BMP normal.
6. `fhir_search` MedicationRequest active + MedicationAdministration 24h. Returns: ketorolac 15mg IV q6h, ondansetron 4mg IV q8h PRN (given x1), enoxaparin 40mg SQ daily, cefazolin 2g IV (single dose, perioperative).

**Result**: Progress note with POD 1 framing, surgical site assessment flagged for clinician input, diet advancement plan, ambulation status, drain output if applicable, discharge criteria checklist.

## Troubleshooting

### MedicationAdministration not available on the FHIR server
- This resource is optional in many FHIR implementations. Rely on MedicationRequest with `status=active` for the medication list.
- Note in the progress note that administered medication data could not be verified. PRN usage will need to be obtained from nursing documentation.
- Check if the encounter-level medication list is available through a different resource type.

### I&O data not captured in FHIR Observations
- Intake/output is frequently not stored as discrete FHIR Observations. It may exist only in nursing flowsheets within the EHR.
- Flag as: "I&O: Not available in structured FHIR data -- obtain from nursing flowsheet."
- Do not fabricate I&O values. The note should explicitly state data is unavailable.

### DiagnosticReport returns references but no conclusion text
- Follow `DiagnosticReport.result` references to individual Observation resources using `fhir_read`.
- Check `DiagnosticReport.presentedForm` for full text attachments.
- If neither is available, note: "[Study type] performed [date] -- result pending or not available in structured data."

## Related Skills

- `history-and-physical-generator` - For the initial admission documentation
- `soap-note-generator` - For outpatient encounter documentation
- `discharge-summary-writer` - For discharge documentation when the patient is ready
- `critical-value-alert-generator` - For flagging critical lab values found during note assembly
