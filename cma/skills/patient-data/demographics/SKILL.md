---
name: langcare-demographics
description: >
  Retrieves and formats a complete patient demographic summary from FHIR Patient,
  RelatedPerson, and Coverage resources. Use when asked to pull demographics, get
  patient info, show patient details, who is this patient, emergency contacts, or
  insurance info. Flags missing critical data elements that may affect care delivery.
---

# Patient Demographics Summary

## When to Use This Skill
Use when a clinician needs a consolidated view of non-clinical patient data including identifiers, contact information, emergency contacts, insurance coverage, preferred language, and advance directive status.

## Clinical Workflow
1. Use `fhir_read` to retrieve the Patient resource and extract identifiers, name, DOB, gender, race/ethnicity (US Core extensions), address, telecom, preferred language, and marital status
2. Use `fhir_search` to find RelatedPerson resources linked to the patient for emergency contacts, next of kin, and guarantors
3. Use `fhir_search` to retrieve active Coverage resources for insurance information including plan name, group, subscriber ID, payor, and coverage order
4. Use `fhir_search` to check for advance directives via Consent (category=acd) or DocumentReference (LOINC 64298-3)
5. Flag missing critical data: no emergency contact, no active insurance, no preferred language, no phone/email, no address, no advance directive on file

## FHIR Resources
- **Patient** -- Core demographics: name, birthDate, gender, address, telecom, communication, identifier, extension (us-core-race, us-core-ethnicity)
- **RelatedPerson** -- Emergency contacts, next of kin, guarantors: name, telecom, relationship, period
- **Coverage** -- Insurance information: status, type, subscriberId, payor, period, class, order
- **Consent** -- Advance directives: category, status, dateTime
- **DocumentReference** -- Advance directive documents: type (LOINC 64298-3), status

## FHIR Query Examples
### Retrieve Patient Resource
```
fhir_read(resourceType="Patient", id="[patient-id]")
```

### Find Emergency Contacts
```
fhir_search(resourceType="RelatedPerson", queryParams="patient=[patient-id]")
```

### Retrieve Active Insurance
```
fhir_search(resourceType="Coverage", queryParams="patient=[patient-id]&status=active")
```

### Check Advance Directives
```
fhir_search(resourceType="Consent", queryParams="patient=[patient-id]&category=http://terminology.hl7.org/CodeSystem/consentcategorycodes|acd")
```

## Clinical Guidelines
- Joint Commission standards require documented emergency contact and advance directive status
- CMS Conditions of Participation require insurance verification before service delivery
- ONC USCDI v3 mandates race, ethnicity, preferred language, and sexual orientation/gender identity data capture

## Interpretation Guide
- Present identifiers with MRN prominently; mask SSN to last 4 digits
- Calculate age from birthDate and display alongside DOB
- Use `name.use` = "official" as primary display name; note maiden names separately
- Classify coverage by order field: 1 = primary, 2 = secondary
- Flag all missing data elements in a dedicated section with clinical impact notes
- For Patient.contact array (EPIC-style embedded contacts), extract relationship, name, and telecom when RelatedPerson search returns empty

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
