# FHIR Patient Resource Field Reference

## Core Demographics

| Field | Path | Description |
|-------|------|-------------|
| MRN | `identifier` where `type.coding.code` = "MR" | Medical Record Number |
| SSN | `identifier` where `type.coding.code` = "SS" | Display last 4 only |
| Official Name | `name` where `use` = "official" | Primary display name |
| Maiden Name | `name` where `use` = "maiden" | Previous surname |
| Birth Date | `birthDate` | YYYY-MM-DD format |
| Gender | `gender` | administrative gender: male, female, other, unknown |
| Marital Status | `maritalStatus` | CodeableConcept |
| Deceased | `deceasedBoolean` or `deceasedDateTime` | Check both forms |

## US Core Extensions

| Extension | URL | Extraction |
|-----------|-----|------------|
| Race | `http://hl7.org/fhir/us/core/StructureDefinition/us-core-race` | `extension.valueCoding` within `ombCategory` and `detailed` sub-extensions |
| Ethnicity | `http://hl7.org/fhir/us/core/StructureDefinition/us-core-ethnicity` | `extension.valueCoding` within `ombCategory` sub-extension |
| Birth Sex | `http://hl7.org/fhir/us/core/StructureDefinition/us-core-birthsex` | `extension.valueCode` (M, F, UNK) |
| Gender Identity | `http://hl7.org/fhir/us/core/StructureDefinition/us-core-genderIdentity` | `extension.valueCodeableConcept` |

## Contact Information

| Field | Path | Values |
|-------|------|--------|
| Home Phone | `telecom` where `system` = "phone" and `use` = "home" | |
| Mobile Phone | `telecom` where `system` = "phone" and `use` = "mobile" | |
| Work Phone | `telecom` where `system` = "phone" and `use` = "work" | |
| Email | `telecom` where `system` = "email" | |
| Fax | `telecom` where `system` = "fax" | |
| Home Address | `address` where `use` = "home" | line, city, state, postalCode, country |
| Mailing Address | `address` where `use` = "billing" or `type` = "postal" | |

## Communication Preferences

| Field | Path | Notes |
|-------|------|-------|
| Preferred Language | `communication` where `preferred` = true | `communication.language.coding.display` |
| All Languages | `communication[].language` | Full list of spoken languages |

## RelatedPerson Relationship Codes

| Code | Display | Common Use |
|------|---------|------------|
| C | Emergency Contact | Primary emergency contact |
| N | Next-of-Kin | Next of kin |
| E | Employer | Employer contact |
| F | Federal Agency | Government entity |
| I | Insurance Company | Insurer |
| S | State Agency | State entity |
| U | Unknown | Unspecified |
| GUARD | Guardian | Legal guardian |
| PRN | Parent | Parent |
| SPS | Spouse | Spouse/partner |
| CHILD | Child | Child |
| SIB | Sibling | Sibling |
| POWATT | Power of Attorney | Healthcare POA |

## Coverage Class Types

| Code | Display | Purpose |
|------|---------|---------|
| group | Group | Employer group number |
| plan | Plan | Plan name/identifier |
| subplan | SubPlan | Sub-plan identifier |
| class | Class | Coverage class |
| subclass | SubClass | Sub-class |
| rxbin | RX BIN | Pharmacy BIN |
| rxpcn | RX PCN | Pharmacy PCN |
| rxid | RX ID | Pharmacy member ID |
| rxgroup | RX Group | Pharmacy group |

## Advance Directive LOINC Codes

| LOINC | Description |
|-------|-------------|
| 64298-3 | Power of attorney and advance directive |
| 75320-2 | Advance directive |
| 81334-5 | Patient goals, preferences, and priorities |
| 42348-3 | Advance directives |
