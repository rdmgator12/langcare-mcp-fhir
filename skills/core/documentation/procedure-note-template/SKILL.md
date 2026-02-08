---
name: procedure-note-template
description: |
  Generates procedure documentation templates pre-populated from FHIR data including patient demographics, indication, relevant labs, allergies, pre-procedure verification, sedation documentation, and complications checklist. Use when user asks to "write a procedure note", "document a procedure", "procedure template", "create procedure documentation", mentions "central line note", "intubation note", "LP note", or needs structured procedure documentation. Do NOT use for surgical operative reports, SOAP notes, progress notes, or discharge documentation.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: documentation
---

# Procedure Note Template

## Overview

Generate pre-populated procedure documentation templates from FHIR data. Pull patient demographics, procedure indication from active conditions, relevant pre-procedure labs (coagulation studies, platelets, hemoglobin), allergy list, and current anticoagulant status. Include required elements: informed consent verification, time-out documentation, procedure details, specimen handling, complications, and post-procedure orders. Support common bedside procedures: central venous catheter, arterial line, intubation, lumbar puncture, paracentesis, thoracentesis, chest tube, foley catheter, and NG tube.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Demographics for note header | name, birthDate, gender, identifier |
| Condition | Procedure indication | code, clinicalStatus |
| Observation | Pre-procedure labs (coags, CBC), vitals | code, value[x], effectiveDateTime |
| AllergyIntolerance | Allergy check (esp. latex, iodine, lidocaine) | code, reaction, clinicalStatus |
| MedicationRequest | Anticoagulant status, sedation orders | medicationCodeableConcept, status, dosageInstruction |
| MedicationAdministration | Sedation medications given | medicationCodeableConcept, dosage, effectiveDateTime |
| Consent | Informed consent status | status, scope, dateTime |
| Procedure | Create procedure record | code, status, performedDateTime, outcome, complication |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract: name, DOB, age, gender, MRN for procedure note header and patient identification band verification.

### Step 2: Identify Procedure Indication

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Match the stated procedure to an active condition as the indication. Common mappings:
- Central line: difficult IV access, need for vasopressors, TPN, prolonged IV antibiotics
- Arterial line: hemodynamic instability, frequent ABG monitoring
- Intubation: respiratory failure, airway protection
- Lumbar puncture: meningitis workup, subarachnoid hemorrhage evaluation
- Paracentesis: ascites (tense, diagnostic)
- Thoracentesis: pleural effusion (diagnostic or therapeutic)
- Chest tube: pneumothorax, hemothorax, empyema
- Foley catheter: urinary retention, strict I&O monitoring, perioperative
- NG tube: bowel obstruction, GI decompression, medication administration

### Step 3: Pull Pre-Procedure Labs

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&code=http://loinc.org|5902-2,http://loinc.org|6301-6,http://loinc.org|777-3,http://loinc.org|718-7,http://loinc.org|3173-2&_sort=-date&_count=20"
```

Critical pre-procedure LOINC codes:
- 5902-2: PT (Prothrombin time)
- 6301-6: INR
- 3173-2: aPTT (Activated partial thromboplastin time)
- 777-3: Platelet count
- 718-7: Hemoglobin
- 4544-3: Hematocrit

Flag if:
- INR > 1.5 (relative contraindication for most invasive procedures)
- Platelets < 50,000 (increased bleeding risk)
- Platelets < 20,000 (contraindication without transfusion)
- aPTT > 1.5x control
- Hemoglobin < 7 (consider transfusion before elective procedure)
- Labs > 24 hours old (recommend recheck)

### Step 4: Check Allergies

```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Flag procedure-relevant allergies:
- **Latex**: Use non-latex gloves, equipment
- **Iodine/Betadine**: Use chlorhexidine for skin prep
- **Chlorhexidine**: Use betadine for skin prep
- **Lidocaine/local anesthetics**: Use alternative anesthetic, allergy consult
- **Adhesive/tape**: Use alternative securement
- **Heparin** (HIT): Avoid heparin-coated catheters and flushes

### Step 5: Check Anticoagulant Status

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&category=http://terminology.hl7.org/CodeSystem/medicationrequest-category|inpatient"
```

Check active medications for anticoagulants and antiplatelets:
- Heparin drip: Check if held, last aPTT value
- Enoxaparin: Timing of last dose (hold 12h for prophylactic, 24h for therapeutic)
- Warfarin: Current INR
- DOACs (apixaban, rivarelbán, edoxaban): Timing of last dose (hold 24-48h)
- Clopidogrel, prasugrel, ticagrelor: Document if held
- Aspirin: Generally continued for most bedside procedures

### Step 6: Check Consent Status

```
Tool: fhir_search
resourceType: "Consent"
queryParams: "patient=[patient-id]&status=active&scope=treatment"
```

If no procedure-specific consent found, flag: "INFORMED CONSENT: NOT DOCUMENTED -- obtain before proceeding."

### Step 7: Pull Pre-Procedure Vitals

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=vital-signs&_sort=-date&_count=10"
```

Document baseline vitals before procedure.

### Step 8: Assemble Procedure Note Template

```
PROCEDURE NOTE
===============
Patient: [name] | MRN: [mrn] | DOB: [dob] (Age: [age]) | Sex: [gender]
Date/Time: [procedure datetime]
Procedure: [procedure name]
Operator: [Flag: "Enter operator name and credentials"]
Supervising Physician: [if applicable]
Service: [service]

INDICATION
----------
[Condition from Step 2 with ICD-10 code]

INFORMED CONSENT
----------------
[Consent status from Step 6]
Risks, benefits, and alternatives discussed with: [patient / surrogate]
Consent signed: [date/time or "REQUIRED"]

ALLERGIES
---------
[List with procedure-relevant flags]

PRE-PROCEDURE VERIFICATION (TIME-OUT)
--------------------------------------
- [ ] Correct patient (two-identifier verification)
- [ ] Correct procedure confirmed
- [ ] Correct site/laterality marked (if applicable)
- [ ] Informed consent obtained
- [ ] Relevant labs reviewed:
      PT/INR: [value] ([date]) [FLAG if abnormal]
      Platelets: [value] ([date]) [FLAG if abnormal]
      Hemoglobin: [value] ([date]) [FLAG if abnormal]
      aPTT: [value] ([date]) [FLAG if abnormal]
- [ ] Anticoagulant status: [status from Step 5]
- [ ] Allergies reviewed: [summary]
- [ ] Equipment and supplies verified

PRE-PROCEDURE VITALS
---------------------
HR: [hr] | BP: [sys]/[dia] | RR: [rr] | SpO2: [spo2]% on [O2]

SEDATION / ANESTHESIA
----------------------
[Flag: "Complete if conscious sedation used"]
Sedation type: [none / local only / moderate sedation / deep sedation]
Medications administered:
  - [Drug] [dose] [route] [time] [Flag: "Enter"]
  - [Drug] [dose] [route] [time] [Flag: "Enter"]
Pre-sedation assessment: ASA class [I-V], Mallampati [I-IV], NPO status [hours]
Monitoring: Continuous pulse oximetry, cardiac monitor, ETCO2 (if applicable)

PROCEDURE DETAILS
-----------------
[Flag: "Operator to complete procedure details"]
Position: [supine / lateral decubitus / sitting / Trendelenburg]
Skin prep: [chlorhexidine / betadine] [Note allergy-based selection]
Draping: Sterile draping applied
Anesthesia: [lidocaine X% / bupivacaine X%] [volume] mL infiltrated to [site]
Technique: [Description of procedure steps]
Site: [anatomical location, laterality]

[Procedure-specific fields -- see references/procedure-documentation.md]

SPECIMENS
---------
[If applicable]
Type: [fluid / tissue / culture]
Sent to: [lab / microbiology / cytology / pathology]
Tests ordered: [cell count, culture, protein, glucose, LDH, cytology, etc.]
Labeled: [Yes -- two-identifier verification]

ESTIMATED BLOOD LOSS
--------------------
[volume] mL

COMPLICATIONS
-------------
[None / describe]
[Procedure-specific complication checklist -- see references/procedure-safety.md]

POST-PROCEDURE
--------------
Patient tolerated procedure: [well / with complications]
Post-procedure vitals: HR [hr] | BP [sys]/[dia] | SpO2 [spo2]%
Post-procedure imaging ordered: [CXR for central line/chest tube / none]
Post-procedure orders:
  - [Site check q[interval]]
  - [Dressing change instructions]
  - [Activity restrictions]
  - [Lab follow-up]

DISPOSITION
-----------
Patient returned to: [floor / ICU / recovery]
Attending notified: [Yes/No]
```

### Step 9: Create Procedure Resource in FHIR

```
Tool: fhir_create
resourceType: "Procedure"
resource: {
  "resourceType": "Procedure",
  "status": "completed",
  "code": {
    "coding": [{
      "system": "http://www.ama-assn.org/go/cpt",
      "code": "[CPT-code]",
      "display": "[procedure-name]"
    }]
  },
  "subject": {"reference": "Patient/[patient-id]"},
  "encounter": {"reference": "Encounter/[encounter-id]"},
  "performedDateTime": "[procedure-datetime]",
  "performer": [{
    "actor": {"reference": "Practitioner/[practitioner-id]"}
  }],
  "reasonReference": [{"reference": "Condition/[indication-condition-id]"}],
  "outcome": {
    "coding": [{
      "system": "http://snomed.info/sct",
      "code": "385669000",
      "display": "Successful"
    }]
  },
  "note": [{"text": "[brief procedure summary]"}]
}
```

Common CPT codes:
- 36556: Central venous catheter insertion (non-tunneled)
- 36620: Arterial line insertion
- 31500: Intubation, endotracheal
- 62270: Lumbar puncture
- 49083: Paracentesis
- 32555: Thoracentesis
- 32551: Chest tube insertion
- 51702: Foley catheter insertion
- 43752: NG tube insertion

## Examples

### Example 1: Central Line Placement

**User says**: "Procedure note for central line placement on patient 11111."

**Actions**:
1. `fhir_read` Patient/11111. Returns: James Torres, DOB 1955-06-30, Male, MRN-11111.
2. `fhir_search` Condition active. Returns: Septic shock (R65.21), pneumonia (J18.9), T2DM, CKD4.
3. `fhir_search` Observation labs (coags, CBC). Returns: INR 1.2, platelets 188k, Hgb 9.8, aPTT 28.
4. `fhir_search` AllergyIntolerance. Returns: Latex allergy (urticaria). Flag: USE NON-LATEX EQUIPMENT.
5. `fhir_search` MedicationRequest anticoagulants. Returns: Heparin drip active -- held 2 hours ago, aPTT at hold was 55.
6. `fhir_search` Observation vitals. Returns: HR 105, BP 88/52 on norepinephrine, SpO2 96% on 4L NC.

**Result**:
```
PROCEDURE NOTE
===============
Patient: James Torres | MRN: MRN-11111 | DOB: 1955-06-30 (Age: 68) | Sex: Male
Procedure: Central venous catheter insertion (non-tunneled)

INDICATION: Septic shock requiring vasopressor administration (R65.21)

ALLERGIES
** LATEX ALLERGY (urticaria) -- USE NON-LATEX GLOVES AND EQUIPMENT **

PRE-PROCEDURE VERIFICATION
- Labs: INR 1.2 [OK] | Platelets 188k [OK] | Hgb 9.8 [OK] | aPTT 28 [OK]
- Anticoagulant: Heparin drip HELD 2h ago, aPTT at hold: 55
- Consent: [VERIFY]

PRE-PROCEDURE VITALS
HR: 105 | BP: 88/52 (on norepinephrine) | SpO2: 96% on 4L NC

[Procedure details: operator to complete -- site, technique, number of attempts,
 catheter type/size, line placement confirmation method, post-procedure CXR ordered]
```

### Example 2: Lumbar Puncture

**User says**: "Generate LP procedure note template for patient pt-222, meningitis workup."

**Actions**:
1. `fhir_read` Patient/pt-222. Returns: Emily Park, DOB 1990-03-22, Female.
2. `fhir_search` Condition. Returns: Fever of unknown origin (R50.9), headache (R51.9), nuchal rigidity (R29.1).
3. `fhir_search` Observation labs. Returns: INR 1.0, platelets 245k, Hgb 12.8, WBC 18.5.
4. `fhir_search` AllergyIntolerance. Returns: NKDA.
5. `fhir_search` MedicationRequest. Returns: No anticoagulants active.
6. `fhir_search` Observation vitals. Returns: T 39.2C, HR 110, BP 128/78, SpO2 99% RA.

**Result**: Pre-populated LP template with indication (meningitis workup), normal coags confirmed, no allergy concerns, specimen handling section pre-filled (tube 1: cell count/diff, tube 2: glucose/protein, tube 3: Gram stain/culture, tube 4: hold for additional studies), opening pressure documentation field, post-LP instructions (flat 1-2 hours, monitor for headache).

## Troubleshooting

### Pre-procedure labs are older than 24 hours
- Flag prominently: "Labs dated [date] -- [X] hours old. Consider recheck before procedure if clinically indicated."
- For INR and platelets, 24-48 hours is generally acceptable if no interval events (bleeding, transfusion, new anticoagulation).
- For hemoglobin in actively bleeding patients, recommend point-of-care testing.

### Consent resource not found in FHIR
- Consent resources are not universally implemented in FHIR servers. Many systems store consent in paper or scanned documents.
- Search DocumentReference for scanned consent: `fhir_search` DocumentReference with `patient=[id]&type=http://loinc.org|59284-0` (LOINC 59284-0 = Consent document).
- If not found, prominently flag: "INFORMED CONSENT STATUS: UNABLE TO VERIFY IN ELECTRONIC RECORD -- confirm paper consent before proceeding."

### Procedure-specific CPT code not in standard list
- Use SNOMED CT coding as an alternative: `system: "http://snomed.info/sct"`.
- Common SNOMED codes: 233573008 (central line), 52765003 (intubation), 277762005 (lumbar puncture), 86088003 (paracentesis), 91602002 (thoracentesis).
- If no standard code matches, use `code.text` with the procedure name as free text.

## Related Skills

- `soap-note-generator` - For documenting the encounter containing the procedure
- `progress-note-writer` - For post-procedure daily documentation
- `lab-result-interpreter` - For interpreting pre-procedure lab values
- `preoperative-lab-checklist` - For verifying all required pre-procedure labs are current
