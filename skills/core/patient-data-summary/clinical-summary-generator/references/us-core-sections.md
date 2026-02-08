# US Core Required Sections for Clinical Summaries

## US Core Profile Overview

US Core Implementation Guide (IG) defines the minimum mandatory data elements, extensions, terminologies, and value sets for interoperability. Version 6.1.0 is the current standard referenced by ONC and CMS.

## Required Data Classes (USCDI v3)

These are the data classes and elements mandated for certified health IT:

### 1. Patient Demographics / Header Information

| Element | US Core Profile | FHIR Resource |
|---------|----------------|---------------|
| Patient Name | US Core Patient | Patient.name |
| Date of Birth | US Core Patient | Patient.birthDate |
| Sex (Administrative) | US Core Patient | Patient.gender |
| Race | US Core Patient | Patient.extension (us-core-race) |
| Ethnicity | US Core Patient | Patient.extension (us-core-ethnicity) |
| Preferred Language | US Core Patient | Patient.communication |
| Address | US Core Patient | Patient.address |
| Phone/Email | US Core Patient | Patient.telecom |
| Birth Sex | US Core Patient | Patient.extension (us-core-birthsex) |
| Gender Identity | US Core Patient | Patient.extension (us-core-genderIdentity) |

### 2. Allergies and Intolerances

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Substance (Drug) | US Core AllergyIntolerance | AllergyIntolerance.code (RxNorm) |
| Substance (Non-Drug) | US Core AllergyIntolerance | AllergyIntolerance.code (SNOMED CT) |
| Reaction | US Core AllergyIntolerance | AllergyIntolerance.reaction.manifestation |
| Severity | US Core AllergyIntolerance | AllergyIntolerance.reaction.severity |
| Status | US Core AllergyIntolerance | AllergyIntolerance.clinicalStatus |

**Required Terminology:**
- Drug allergens: RxNorm (`http://www.nlm.nih.gov/research/umls/rxnorm`)
- Non-drug allergens: SNOMED CT (`http://snomed.info/sct`)
- Reaction manifestations: SNOMED CT
- Clinical status: `http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical`

### 3. Problems / Health Concerns

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Problem/Diagnosis | US Core Condition Problems and Health Concerns | Condition.code |
| Date of Diagnosis | US Core Condition | Condition.onsetDateTime |
| Clinical Status | US Core Condition | Condition.clinicalStatus |
| Verification Status | US Core Condition | Condition.verificationStatus |

**Required Terminology:**
- Condition codes: SNOMED CT (preferred), ICD-10-CM
- Clinical status: active, recurrence, relapse, inactive, remission, resolved
- Category must include `problem-list-item` or `health-concern`

### 4. Medications

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Medication Name | US Core MedicationRequest | MedicationRequest.medicationCodeableConcept |
| Dose | US Core MedicationRequest | MedicationRequest.dosageInstruction |
| Frequency | US Core MedicationRequest | MedicationRequest.dosageInstruction.timing |
| Route | US Core MedicationRequest | MedicationRequest.dosageInstruction.route |
| Status | US Core MedicationRequest | MedicationRequest.status |

**Required Terminology:**
- Medication codes: RxNorm
- Route: SNOMED CT

### 5. Laboratory Results

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Test Name | US Core Laboratory Result Observation | Observation.code (LOINC) |
| Value | US Core Laboratory Result Observation | Observation.value[x] |
| Units | US Core Laboratory Result Observation | Observation.valueQuantity.unit (UCUM) |
| Reference Range | US Core Laboratory Result Observation | Observation.referenceRange |
| Interpretation | US Core Laboratory Result Observation | Observation.interpretation |
| Date | US Core Laboratory Result Observation | Observation.effectiveDateTime |

**Required Terminology:**
- Test codes: LOINC (`http://loinc.org`)
- Units: UCUM (`http://unitsofmeasure.org`)

### 6. Vital Signs

| Vital Sign | LOINC Code | US Core Profile |
|------------|------------|-----------------|
| Blood Pressure Panel | 85354-9 | US Core Blood Pressure |
| Systolic BP | 8480-6 | (component) |
| Diastolic BP | 8462-4 | (component) |
| Heart Rate | 8867-4 | US Core Heart Rate |
| Respiratory Rate | 9279-1 | US Core Respiratory Rate |
| Body Temperature | 8310-5 | US Core Body Temperature |
| Body Weight | 29463-7 | US Core Body Weight |
| Body Height | 8302-2 | US Core Body Height |
| BMI | 39156-5 | US Core BMI |
| O2 Saturation | 2708-6 | US Core Pulse Oximetry |
| Head Circumference | 9843-4 | US Core Head Circumference |

### 7. Procedures

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Procedure Name | US Core Procedure | Procedure.code |
| Date Performed | US Core Procedure | Procedure.performed[x] |
| Status | US Core Procedure | Procedure.status |
| Body Site | US Core Procedure | Procedure.bodySite |

**Required Terminology:**
- Procedure codes: CPT, SNOMED CT, HCPCS, ICD-10-PCS, CDT

### 8. Immunizations

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Vaccine Administered | US Core Immunization | Immunization.vaccineCode (CVX) |
| Date Administered | US Core Immunization | Immunization.occurrenceDateTime |
| Status | US Core Immunization | Immunization.status |
| Dose Number | US Core Immunization | Immunization.protocolApplied.doseNumber |

**Required Terminology:**
- Vaccine codes: CVX (`http://hl7.org/fhir/sid/cvx`)
- NDC as secondary (`http://hl7.org/fhir/sid/ndc`)

### 9. Clinical Notes

| Note Type | LOINC Code | US Core Profile |
|-----------|------------|-----------------|
| Consultation Note | 11488-4 | US Core DocumentReference |
| Discharge Summary | 18842-5 | US Core DocumentReference |
| History & Physical | 34117-2 | US Core DocumentReference |
| Progress Note | 11506-3 | US Core DocumentReference |
| Procedures Note | 28570-0 | US Core DocumentReference |
| Imaging Narrative | 18748-4 | US Core DiagnosticReport |
| Laboratory Report | 11502-2 | US Core DiagnosticReport |
| Pathology Report | 11526-1 | US Core DiagnosticReport |

### 10. Care Team / Provenance

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Care Team Members | US Core CareTeam | CareTeam.participant |
| Care Team Status | US Core CareTeam | CareTeam.status |
| Author / Transmitter | US Core Provenance | Provenance.agent |
| Timestamp | US Core Provenance | Provenance.recorded |

### 11. Goals

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Goal Description | US Core Goal | Goal.description |
| Goal Status | US Core Goal | Goal.lifecycleStatus |
| Target Date | US Core Goal | Goal.target.dueDate |

### 12. Assessment and Plan of Treatment

| Element | US Core Profile | FHIR Path |
|---------|----------------|-----------|
| Care Plan Narrative | US Core CarePlan | CarePlan.text |
| Care Plan Status | US Core CarePlan | CarePlan.status |
| Category | US Core CarePlan | CarePlan.category |

## Section Ordering for CCD Output

Follow this order when assembling a clinical summary (mirrors C-CDA section ordering):

1. Header (Patient demographics)
2. Allergies and Intolerances
3. Active Problems / Health Concerns
4. Medications
5. Vital Signs (most recent)
6. Laboratory Results
7. Procedures
8. Immunizations
9. Encounters
10. Care Plan / Assessment and Plan
11. Goals
12. Care Team
13. Clinical Notes (if included)

## Must-Support Elements

US Core "Must Support" means:
- The sender SHALL populate the element if data is available
- The receiver SHALL be capable of processing the element
- If data is absent, the sender is not required to send the element (but must not send empty values)

For clinical summary generation, treat Must Support elements as: include if present in the FHIR response, note as "Not documented" if absent for clinically significant fields.
