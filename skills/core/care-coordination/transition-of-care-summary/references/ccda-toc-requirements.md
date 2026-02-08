# C-CDA Transition of Care Requirements

## Overview

The Consolidated Clinical Document Architecture (C-CDA) is an HL7 standard for clinical document exchange. Transition of Care (TOC) documents must contain specific sections to meet Meaningful Use, CMS, and Joint Commission requirements.

## Required C-CDA Sections for Transition of Care

### Section 1: Patient Demographics (Header)

| Element | C-CDA Field | FHIR Mapping |
|---------|-------------|--------------|
| Patient name | recordTarget/patientRole/patient/name | Patient.name |
| Date of birth | recordTarget/patientRole/patient/birthTime | Patient.birthDate |
| Sex | recordTarget/patientRole/patient/administrativeGenderCode | Patient.gender |
| Race | recordTarget/patientRole/patient/raceCode | Patient.extension (us-core-race) |
| Ethnicity | recordTarget/patientRole/patient/ethnicGroupCode | Patient.extension (us-core-ethnicity) |
| Preferred language | recordTarget/patientRole/patient/languageCommunication | Patient.communication |
| Address | recordTarget/patientRole/addr | Patient.address |
| Phone | recordTarget/patientRole/telecom | Patient.telecom |
| MRN | recordTarget/patientRole/id | Patient.identifier |

### Section 2: Allergies and Adverse Reactions (REQUIRED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Allergies and Intolerances Section | 48765-2 |
| Entry | Allergy Concern Act | -- |
| Entry | Allergy Observation | -- |

**FHIR Source:** AllergyIntolerance
- Must include: substance, reaction, severity, status
- "No Known Allergies" must be explicitly documented (not just empty)
- Include onset date if known
- Distinguish between allergy (immune-mediated) and intolerance (non-immune)

### Section 3: Medications (REQUIRED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Medications Section | 10160-0 |
| Entry | Medication Activity | -- |

**FHIR Source:** MedicationRequest, MedicationStatement
- Must include: medication name, dose, route, frequency, status
- Indicate which medications are new, changed, continued, or discontinued
- Include indication (reason for medication) where available
- Include prescriber information
- Reconciliation status should be documented

### Section 4: Problem List (REQUIRED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Problem Section | 11450-4 |
| Entry | Problem Concern Act | -- |
| Entry | Problem Observation | -- |

**FHIR Source:** Condition
- Must include: diagnosis/problem, onset date, status (active/resolved/inactive)
- Use ICD-10 or SNOMED codes
- Identify principal diagnosis separately
- Include verification status (confirmed, provisional, differential)
- Hospital-acquired conditions should be flagged

### Section 5: Procedures (REQUIRED for Discharge Summary)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Procedures Section | 47519-4 |
| Entry | Procedure Activity Procedure | -- |

**FHIR Source:** Procedure
- Must include: procedure name (with CPT/SNOMED code), date, performer
- Include complications if any
- Include pathology findings for biopsies
- Include implant information if applicable (UDI)

### Section 6: Results (REQUIRED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Results Section | 30954-2 |
| Entry | Result Organizer | -- |
| Entry | Result Observation | -- |

**FHIR Source:** Observation (category=laboratory), DiagnosticReport
- Must include: test name (LOINC code), value, units, reference range, interpretation
- Include abnormal flags
- Include pending results with "pending" status
- Group by panel where applicable (CBC, BMP, etc.)

### Section 7: Vital Signs (REQUIRED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Vital Signs Section | 8716-3 |
| Entry | Vital Signs Organizer | -- |

**FHIR Source:** Observation (category=vital-signs)
- Must include most recent: BP, HR, RR, Temp, SpO2, Weight, Height
- Include trends if clinically significant (e.g., declining BP)

### Section 8: Encounter Diagnosis (REQUIRED for Discharge Summary)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Encounter Diagnosis Section | 29308-4 |
| Entry | Encounter Diagnosis | -- |

**FHIR Source:** Encounter.reasonCode, Condition (encounter-diagnosis category)
- Principal diagnosis listed first
- Include all encounter-related diagnoses with codes

### Section 9: Plan of Treatment (REQUIRED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Plan of Treatment Section | 18776-5 |
| Entry | Planned Act / Observation / Procedure | -- |

**FHIR Source:** CarePlan, ServiceRequest, Goal
- Follow-up appointments (scheduled dates, providers)
- Pending orders and expected results
- Referrals
- Diet and activity instructions
- Self-care instructions
- Goals of care

### Section 10: Discharge Instructions (REQUIRED for Discharge Summary)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Instructions Section | 69730-0 |

**FHIR Source:** DocumentReference (type=patient education), CarePlan
- Activity restrictions
- Dietary instructions
- Medication instructions
- Warning signs (when to return to ED)
- Follow-up instructions
- Equipment/supply instructions

### Section 11: Immunizations (RECOMMENDED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Immunizations Section | 11369-6 |
| Entry | Immunization Activity | -- |

**FHIR Source:** Immunization
- Include immunizations given during encounter
- Include relevant historical immunizations (flu, pneumococcal, tetanus, COVID)

### Section 12: Advance Directives (RECOMMENDED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Advance Directives Section | 42348-3 |
| Entry | Advance Directive Observation | -- |

**FHIR Source:** Consent
- Code status (Full code, DNR, DNR-DNI, comfort measures)
- Healthcare proxy / power of attorney
- POLST/MOLST if applicable
- Date of most recent review

### Section 13: Functional Status (RECOMMENDED)

| Element | C-CDA Template | LOINC Section Code |
|---------|---------------|-------------------|
| Section | Functional Status Section | 47420-5 |

**FHIR Source:** Observation (category=survey)
- ADL assessment
- Mobility status
- Cognitive status
- Assistive devices needed
- Fall risk level

---

## Joint Commission Transition of Care Requirements

### National Patient Safety Goal (NPSG) 02.03.01

"Improve the effectiveness of communication among caregivers."

#### Required Elements for Handoff Communication
1. Interactive communication allowing opportunity to question
2. Up-to-date information regarding patient's condition, care, treatment, medications, services, and any recent or anticipated changes
3. Process for verification of received information including repeat-back or read-back
4. Opportunity for receiver to review relevant historical data (labs, vitals, trends)
5. Interruptions during handoffs are limited

#### Required Documentation Elements
| Element | Required | Notes |
|---------|----------|-------|
| Patient identification | Yes | Name, DOB, MRN minimum 2 identifiers |
| Diagnosis/problem list | Yes | Principal + active |
| Medication list (reconciled) | Yes | With changes noted |
| Allergy list | Yes | Including "no known allergies" |
| Code status/advance directives | Yes | Must be current |
| Pending tests/results | Yes | What is outstanding |
| Follow-up needs | Yes | Appointments, referrals |
| Clinical summary | Yes | Brief hospital course |
| Contingency plans | Yes | "If...then..." scenarios |
| Responsible provider | Yes | Who to contact for questions |

### PC.04.01.01 -- Planning for Care, Treatment, and Services

Applies specifically to discharge/transfer:
- Discharge plan addresses patient's identified needs at the time of discharge
- Patient and family are involved in discharge planning
- Written discharge instructions provided
- Discharge summary completed within 30 days (but best practice is within 24 hours)

### PC.04.02.01 -- When a Patient is Discharged or Transferred

The organization provides the patient or the receiving organization with:
- Reason for the care, treatment, or services provided
- Significant findings
- Procedures and treatment provided
- Patient's condition at discharge/transfer
- Any patient education provided
- Community resources or referrals made

---

## CMS Discharge Summary Requirements (42 CFR 482.24(c)(2))

A discharge summary must be completed within 30 days of discharge and include:
1. Reason for hospitalization
2. Significant findings
3. Procedures and treatment provided
4. Patient's condition and disposition at discharge
5. Patient/family education and instructions provided
6. Attending practitioner's signature

## Timing Requirements

| Document Type | Timeline | Regulation |
|---------------|----------|------------|
| Discharge summary | Within 30 days of discharge | CMS CoP |
| Discharge summary (best practice) | Within 24-48 hours | Joint Commission |
| TOC to receiving facility | Before or at time of transfer | Joint Commission |
| TOC to PCP | Within 24 hours of discharge | CMS/Meaningful Use |
| Medication reconciliation | At each transition point | Joint Commission NPSG |
| Patient discharge instructions | Before patient leaves | CMS CoP |

## C-CDA Document Types (LOINC Codes)

| Document Type | LOINC Code | Use Case |
|---------------|-----------|----------|
| Discharge summary | 18842-5 | Hospital discharge |
| Transfer summary | 18761-7 | Transfer to another facility |
| Referral note | 57133-1 | Outpatient referral |
| Progress note | 11506-3 | Ongoing care |
| Consultation note | 11488-4 | Specialist consultation |
| History and Physical | 34117-2 | Admission H&P |
| Continuity of Care Document (CCD) | 34133-9 | Ambulatory summary |
| Care Plan | 18776-5 | Treatment plan |
