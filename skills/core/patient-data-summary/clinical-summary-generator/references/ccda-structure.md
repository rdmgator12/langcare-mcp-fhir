# C-CDA Document Structure and Required Elements

## Overview

Consolidated Clinical Document Architecture (C-CDA) is the HL7 standard for clinical document exchange. C-CDA Release 2.1 (with Companion Guide R4) is the current standard. When generating FHIR-based clinical summaries, follow C-CDA section structure for consistency with existing clinical workflows.

## Document Types

### Continuity of Care Document (CCD)
- **Template OID:** 2.16.840.1.113883.10.20.22.1.2
- **Use:** Ongoing care summary, referrals, care transitions
- **Required Sections:** Allergies, Medications, Problems, Procedures, Results, Encounters, Immunizations

### Discharge Summary
- **Template OID:** 2.16.840.1.113883.10.20.22.1.8
- **Use:** Hospital discharge
- **Additional Required:** Hospital Course, Discharge Diagnosis, Discharge Medications, Discharge Instructions

### Referral Note
- **Template OID:** 2.16.840.1.113883.10.20.22.1.14
- **Use:** Specialist referral
- **Additional Required:** Reason for Referral, History of Present Illness

### Progress Note
- **Template OID:** 2.16.840.1.113883.10.20.22.1.9
- **Use:** Ongoing visit documentation
- **Additional Required:** Assessment, Plan of Treatment, Subjective, Objective

## CCD Required Sections

### 1. Allergies and Intolerances Section
**LOINC:** 48765-2
**Template OID:** 2.16.840.1.113883.10.20.22.2.6.1

Required entries per allergy:
- Substance (coded: RxNorm for drugs, UNII for non-drugs, NDF-RT)
- Status (active, inactive, resolved)
- Severity (mild, moderate, severe, life-threatening)
- Reaction (coded: SNOMED CT)
- Criticality (low, high)

FHIR mapping:
```
AllergyIntolerance.code            -> Substance
AllergyIntolerance.clinicalStatus  -> Status
AllergyIntolerance.reaction.severity -> Severity
AllergyIntolerance.reaction.manifestation -> Reaction
AllergyIntolerance.criticality     -> Criticality
```

Special values:
- "No Known Allergies" = AllergyIntolerance.code = SNOMED 716186003
- "No Known Drug Allergies" = SNOMED 409137002

### 2. Medications Section
**LOINC:** 10160-0
**Template OID:** 2.16.840.1.113883.10.20.22.2.1.1

Required entries per medication:
- Drug name (RxNorm coded)
- Dose quantity and unit
- Route (SNOMED CT or NCI)
- Frequency / timing
- Start date
- Status (active, completed, on-hold, stopped)
- Prescriber

FHIR mapping:
```
MedicationRequest.medicationCodeableConcept -> Drug name
MedicationRequest.dosageInstruction.doseAndRate -> Dose
MedicationRequest.dosageInstruction.route -> Route
MedicationRequest.dosageInstruction.timing -> Frequency
MedicationRequest.authoredOn -> Start date
MedicationRequest.status -> Status
MedicationRequest.requester -> Prescriber
```

### 3. Problem List Section
**LOINC:** 11450-4
**Template OID:** 2.16.840.1.113883.10.20.22.2.5.1

Required entries per problem:
- Problem name (SNOMED CT coded, ICD-10-CM secondary)
- Status (active, inactive, resolved)
- Date of onset
- Date resolved (if applicable)
- Verification status

FHIR mapping:
```
Condition.code -> Problem name
Condition.clinicalStatus -> Status
Condition.onsetDateTime -> Date of onset
Condition.abatementDateTime -> Date resolved
Condition.verificationStatus -> Verification
```

### 4. Procedures Section
**LOINC:** 47519-4
**Template OID:** 2.16.840.1.113883.10.20.22.2.7.1

Required entries per procedure:
- Procedure name (CPT, SNOMED CT, HCPCS, ICD-10-PCS)
- Date performed
- Status (completed, in-progress, aborted)
- Body site (if applicable)
- Performer

FHIR mapping:
```
Procedure.code -> Procedure name
Procedure.performedDateTime -> Date
Procedure.status -> Status
Procedure.bodySite -> Body site
Procedure.performer.actor -> Performer
```

### 5. Results Section
**LOINC:** 30954-2
**Template OID:** 2.16.840.1.113883.10.20.22.2.3.1

Organized as Result Organizer (panel) containing Result Observations (individual tests).

Required per result:
- Test name (LOINC coded)
- Value and units (UCUM)
- Reference range
- Interpretation (Normal, Abnormal, High, Low, Critical)
- Date/time of result
- Status (final, preliminary, corrected)

FHIR mapping:
```
Observation.code -> Test name (LOINC)
Observation.valueQuantity -> Value + units
Observation.referenceRange -> Reference range
Observation.interpretation -> Interpretation
Observation.effectiveDateTime -> Date/time
Observation.status -> Status
```

### 6. Encounters Section
**LOINC:** 46240-8
**Template OID:** 2.16.840.1.113883.10.20.22.2.22.1

Required per encounter:
- Encounter type (CPT or SNOMED CT)
- Date/time
- Location/facility
- Provider
- Diagnosis (encounter diagnosis)

FHIR mapping:
```
Encounter.type -> Encounter type
Encounter.period -> Date/time
Encounter.location -> Location
Encounter.participant -> Provider
Encounter.reasonCode -> Diagnosis
```

### 7. Immunizations Section
**LOINC:** 11369-6
**Template OID:** 2.16.840.1.113883.10.20.22.2.2.1

Required per immunization:
- Vaccine (CVX coded)
- Date administered
- Dose number
- Route
- Site
- Manufacturer (if known)
- Lot number (if known)

FHIR mapping:
```
Immunization.vaccineCode -> Vaccine (CVX)
Immunization.occurrenceDateTime -> Date
Immunization.protocolApplied.doseNumber -> Dose number
Immunization.route -> Route
Immunization.site -> Site
Immunization.manufacturer -> Manufacturer
Immunization.lotNumber -> Lot number
```

### 8. Vital Signs Section
**LOINC:** 8716-3
**Template OID:** 2.16.840.1.113883.10.20.22.2.4.1

Follows same structure as Results section but filtered to vital-signs category observations.

### 9. Social History Section
**LOINC:** 29762-2
**Template OID:** 2.16.840.1.113883.10.20.22.2.17

Key elements:
- Smoking status (SNOMED CT: 72166-2 LOINC)
- Alcohol use
- Drug use
- Occupation
- Education level

### 10. Plan of Treatment Section
**LOINC:** 18776-5
**Template OID:** 2.16.840.1.113883.10.20.22.2.10

FHIR mapping:
```
CarePlan.text -> Narrative plan
CarePlan.activity -> Planned activities
CarePlan.goal -> Treatment goals
```

### 11. Goals Section
**LOINC:** 61146-7
**Template OID:** 2.16.840.1.113883.10.20.22.2.60

FHIR mapping:
```
Goal.description -> Goal text
Goal.lifecycleStatus -> Status
Goal.target.dueDate -> Target date
Goal.achievementStatus -> Progress
```

## Narrative Requirements

Each C-CDA section requires a human-readable narrative (`text` element). When generating summaries from FHIR data, construct a readable text block for each section.

Rules:
- Narrative must be consistent with structured entries
- Use tables for multi-row data (meds, labs, problems)
- Include all coded data in display form
- Flag missing required elements as "Not available" rather than omitting

## Null Flavor Handling

When data is absent in the FHIR response:

| Situation | C-CDA Null Flavor | Summary Display |
|-----------|-------------------|-----------------|
| Not asked | NI (No Information) | "Not documented" |
| Asked, unknown | UNK (Unknown) | "Unknown" |
| Masked/restricted | MSK (Masked) | "Information restricted" |
| Not applicable | NA (Not Applicable) | "N/A" |

For clinical summary output, always indicate why data is missing rather than silently omitting sections.

## Document Metadata

Every clinical summary should include:
- **Generation timestamp**: When the summary was created
- **Data source**: FHIR server / EHR system
- **Author**: System-generated (note: AI-assisted)
- **Confidentiality**: Normal (N), Restricted (R), or Very Restricted (V)
- **Data freshness**: Note the date range of included data
