---
name: langcare-insurance-coverage
description: >
  Retrieves and analyzes insurance coverage from FHIR Coverage and
  Organization resources including coordination of benefits, coverage
  gaps, and eligibility details. Use when asked about insurance info,
  coverage details, benefits check, coordination of benefits, coverage
  verification, or payer information.
---

# Insurance Coverage Summary

## When to Use This Skill
Use when a clinician or administrative staff needs detailed insurance coverage information including primary/secondary coordination, coverage periods, subscriber details, and payer contacts.

## Clinical Workflow
1. Use `fhir_search` to retrieve all Coverage resources for the patient (active and inactive)
2. Use `fhir_read` to retrieve referenced Organization resources (payers) for contact details
3. Organize by coverage order: primary (order=1), secondary (order=2), tertiary
4. Extract plan details from class array: plan name, group number, RX BIN/PCN
5. Identify coverage gaps: periods without active coverage, upcoming coverage end dates
6. Present coordination of benefits summary

## FHIR Resources
- **Coverage** -- Insurance entries: status, type, subscriberId, beneficiary, payor, period, class, order, network
- **Organization** -- Payer organization details: name, telecom, address
- **Patient** -- Subscriber vs dependent relationship

## FHIR Query Examples
### Pull All Coverage
```
fhir_search(resourceType="Coverage", queryParams="patient=[patient-id]")
```

### Pull Active Coverage Only
```
fhir_search(resourceType="Coverage", queryParams="patient=[patient-id]&status=active")
```

### Retrieve Payer Organization
```
fhir_read(resourceType="Organization", id="[payer-organization-id]")
```

## Clinical Guidelines
- CMS requires verification of insurance eligibility before service delivery
- Coordination of benefits rules: primary coverage pays first, secondary covers remaining eligible expenses
- Medicare Secondary Payer rules apply when patient has employer coverage and Medicare

## Interpretation Guide
- Present primary coverage first with full plan details, then secondary
- For each coverage: plan name, group number, subscriber ID, coverage period, payer name and phone
- Extract pharmacy benefits separately: RX BIN, RX PCN, RX Group from class array
- Flag: coverage ending within 30 days, no active coverage found, subscriber ID discrepancies between coverages
- Calculate coverage duration and identify any gap periods

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
