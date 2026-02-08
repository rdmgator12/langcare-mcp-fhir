---
name: immunization-status-checker
description: |
  Checks patient immunization status against CDC recommended schedules for adults and pediatrics, identifies overdue and upcoming vaccines, flags contraindications based on patient conditions, and generates ImmunizationRecommendation FHIR resources. Use when user asks to "check immunizations", "vaccine status", "what shots are due", "immunization review", "vaccination history", "overdue vaccines", mentions "immunization schedule", "catch-up vaccines", or needs vaccine recommendations. Do NOT use for immunization population reporting across a panel (use patient-panel-overview) or allergy-only queries.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: population-health
---

# Immunization Status Checker

## Overview

Retrieve a patient's Immunization history from FHIR, compare against CDC-recommended immunization schedules (adult and pediatric), identify overdue and upcoming vaccinations, evaluate contraindications based on documented conditions and allergies, and optionally generate an ImmunizationRecommendation FHIR resource. See references/adult-immunization-schedule.md and references/pediatric-immunization-schedule.md for full schedule details, CVX codes, and contraindication criteria.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Immunization | Vaccination history | vaccineCode, occurrenceDateTime, status, patient, protocolApplied |
| ImmunizationRecommendation | Generated vaccine recommendations | recommendation[].vaccineCode, dateCriterion, forecastStatus |
| Patient | Age, sex for schedule determination | birthDate, gender |
| Condition | Contraindications, risk group identification | code, clinicalStatus |
| AllergyIntolerance | Vaccine component allergies | code, reaction, clinicalStatus |
| Observation | Pregnancy status, immune status labs | code, valueQuantity, effectiveDateTime |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract birthDate (calculate age in years, months, days for pediatric precision), gender. Determine schedule type: pediatric (age < 19) or adult (age >= 19).

### Step 2: Pull Immunization History

```
Tool: fhir_search
resourceType: "Immunization"
queryParams: "patient=[patient-id]&status=completed&_sort=-date&_count=100"
```

For each Immunization resource, extract:
- `vaccineCode.coding` -- match CVX code (system `http://hl7.org/fhir/sid/cvx`)
- `occurrenceDateTime` -- administration date
- `protocolApplied[].doseNumberPositiveInt` -- dose number in series
- `protocolApplied[].seriesString` or `seriesDosesPositiveInt` -- series name and total doses

Build a vaccine history table: Vaccine Name | CVX Code | Date Given | Dose # | Series.

### Step 3: Pull Conditions for Contraindication and Risk Assessment

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Flag conditions that affect immunization recommendations:
- **Immunocompromised**: HIV (SNOMED 86406008), organ transplant (SNOMED 77465005), active chemotherapy, chronic immunosuppressive therapy
- **Pregnancy**: SNOMED 77386006 (check if current via onset/abatement dates)
- **Asplenia**: SNOMED 707147002
- **Chronic conditions**: Diabetes (44054006), chronic heart disease (128238001), chronic lung disease (413839001), chronic liver disease (328383001), CKD (709044004)
- **Healthcare worker**: Check Observation or Condition for occupational codes

### Step 4: Check Allergies for Vaccine Contraindications

```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Flag allergies to vaccine components:
- Egg allergy (SNOMED 91930004) -- affects influenza (some formulations), yellow fever
- Gelatin allergy (SNOMED 294847001) -- affects MMR, varicella, zoster, influenza (some)
- Neomycin allergy -- affects MMR, varicella, IPV
- Yeast allergy (SNOMED 294529003) -- affects Hepatitis B, HPV
- Latex allergy -- affects vaccines with latex-containing vial stoppers
- Previous severe allergic reaction to a specific vaccine (anaphylaxis)

### Step 5: Check Pregnancy and Immune Status

**Pregnancy (for females of childbearing age):**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=http://loinc.org|82810-3&_sort=-date&_count=1"
```
LOINC 82810-3 = Pregnancy status. If positive, apply pregnancy-specific schedule modifications.

**Immune status (for immunocompromised patients):**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=http://loinc.org|24467-3&_sort=-date&_count=1"
```
LOINC 24467-3 = CD4 count. If CD4 < 200, avoid live vaccines.

### Step 6: Compare Against Recommended Schedule

Use the appropriate schedule from references/ based on patient age:

**Adult (>= 19 years):** references/adult-immunization-schedule.md
**Pediatric (< 19 years):** references/pediatric-immunization-schedule.md

For each recommended vaccine series:
1. Check if the patient has completed the series (all required doses at appropriate intervals)
2. Check if any doses are due now (based on age and interval since last dose)
3. Check if any doses are overdue (past the recommended date)
4. Check for contraindications based on conditions and allergies
5. Apply risk-based recommendations (e.g., pneumococcal for diabetics, Hep B for CKD)

Classify each vaccine as:
- **Complete**: All doses received at appropriate intervals
- **Up to date**: On schedule, next dose not yet due
- **Due now**: Next dose is currently recommended
- **Overdue**: Past the recommended date for next dose
- **Contraindicated**: Patient has documented contraindication
- **Not applicable**: Outside recommended age/sex/risk group

### Step 7: Generate Immunization Status Report

```
IMMUNIZATION STATUS REPORT
============================
Patient: [name] | DOB: [date] | Age: [age]
Schedule Type: [Adult/Pediatric]
Risk Groups: [list any applicable -- immunocompromised, pregnant, healthcare worker, etc.]
Report Date: [today]

VACCINATION HISTORY
-------------------
Vaccine              | CVX | Doses Received | Dates
---------------------|-----|----------------|------
Influenza            | 158 | 1 (this season)| 2024-10-01
Tdap                 | 115 | 1              | 2020-06-15
COVID-19 (Pfizer)    | 208 | 4              | 2021-04-01, 2021-04-22, 2021-11-01, 2023-10-15
...

RECOMMENDATIONS
---------------
STATUS    | Vaccine               | Action           | Due Date   | Notes
----------|-----------------------|------------------|------------|------
OVERDUE   | Shingles (Shingrix)   | Administer dose 1| 2024-01-01 | Age >= 50, recommended
DUE NOW   | Influenza (2024-25)   | Administer       | Now        | Annual, any formulation
DUE NOW   | COVID-19 (2024-25)    | Administer       | Now        | Updated booster recommended
COMPLETE  | Tdap/Td               | None             | --         | Tdap 2020, next Td 2030
COMPLETE  | MMR                   | None             | --         | 2 doses documented

CONTRAINDICATIONS
-----------------
[Vaccine]: [Reason] (e.g., "Live zoster vaccine: Contraindicated due to immunosuppressive therapy")

MISSING DATA
------------
[List any vaccines with no history that may have been given before EHR adoption]
```

### Step 8: Generate ImmunizationRecommendation Resource (if requested)

For each vaccine due or overdue, create an ImmunizationRecommendation:
```
Tool: fhir_create
resourceType: "ImmunizationRecommendation"
resource: {
  "resourceType": "ImmunizationRecommendation",
  "patient": {"reference": "Patient/[patient-id]"},
  "date": "[today]",
  "authority": {"reference": "Organization/cdc"},
  "recommendation": [
    {
      "vaccineCode": [
        {
          "coding": [{
            "system": "http://hl7.org/fhir/sid/cvx",
            "code": "[cvx-code]",
            "display": "[vaccine-name]"
          }]
        }
      ],
      "forecastStatus": {
        "coding": [{
          "system": "http://terminology.hl7.org/CodeSystem/immunization-recommendation-status",
          "code": "[due|overdue|contraindicated]"
        }]
      },
      "dateCriterion": [
        {
          "code": {
            "coding": [{
              "system": "http://loinc.org",
              "code": "30980-7",
              "display": "Date vaccine due"
            }]
          },
          "value": "[recommended-date]"
        }
      ],
      "doseNumberPositiveInt": [next-dose-number],
      "seriesDosesPositiveInt": [total-doses-in-series]
    }
  ]
}
```

## Examples

### Example 1: Adult Immunization Review

**User says**: "Check immunization status for patient 12345. What vaccines are they due for?"

**Actions**:
1. Read Patient/12345. Age 62, male.
2. Search Immunization for patient 12345. Returns: Influenza (Oct 2024), Tdap (2019), COVID (4 doses, last Oct 2023), Pneumovax23 (2021).
3. Search Condition -- active T2DM, no immunocompromised state.
4. Search AllergyIntolerance -- no vaccine-related allergies.
5. Compare against adult schedule. Age 62 with diabetes = risk group for pneumococcal, Hep B.

**Result**:
```
IMMUNIZATION STATUS REPORT
============================
Patient: Michael Torres | DOB: 1962-05-20 | Age: 62
Schedule Type: Adult
Risk Groups: Diabetes (pneumococcal, Hep B recommended)

RECOMMENDATIONS
---------------
OVERDUE   | Shingrix (RZV)    | Dose 1 of 2     | Was due at age 50 | CVX 187
DUE NOW   | PCV20             | Administer       | Now               | CVX 216, DM risk group
DUE NOW   | Hep B             | Dose 1 of 3      | Now               | CVX 08, DM + age 19-59 or risk
UP TO DATE| Influenza 2024-25 | None             | Next Oct 2025     | CVX 158, given Oct 2024
UP TO DATE| COVID-19 2024-25  | Updated booster  | When available     | CVX 308
COMPLETE  | Tdap              | Next Td in 2029  |                   | CVX 115, given 2019

4 vaccines require action. Shingrix is highest priority (2+ years overdue).
```

### Example 2: Pediatric Catch-Up Schedule

**User says**: "8-year-old patient abc-456 just transferred. Check what vaccines they need."

**Actions**:
1. Read Patient/abc-456. Age 8 years 3 months, female.
2. Search Immunization. Returns only: DTaP x3 (2,4,6 months), IPV x2 (2,4 months), Hep B x3 (birth,1mo,6mo). Missing many childhood vaccines.
3. No relevant conditions or allergies.
4. Apply pediatric catch-up schedule.

**Result**:
```
IMMUNIZATION STATUS REPORT
============================
Patient: Emily Nakamura | DOB: 2016-08-15 | Age: 8y 3m
Schedule Type: Pediatric (catch-up needed)

CATCH-UP VACCINES NEEDED
-------------------------
OVERDUE | DTaP dose 4     | Administer now | Min interval from dose 3: 6 months (met) | CVX 20
OVERDUE | DTaP dose 5     | 6 months after dose 4 | CVX 20
OVERDUE | IPV dose 3      | Administer now | CVX 10
OVERDUE | IPV dose 4      | 4 weeks after dose 3 if < age 4; final dose at >= 4y | CVX 10
OVERDUE | MMR dose 1      | Administer now | CVX 03
OVERDUE | MMR dose 2      | 4 weeks after dose 1 | CVX 03
OVERDUE | Varicella dose 1| Administer now | CVX 21
OVERDUE | Varicella dose 2| 3 months after dose 1 | CVX 21
OVERDUE | Hep A dose 1    | Administer now | CVX 83
OVERDUE | Hep A dose 2    | 6 months after dose 1 | CVX 83
DUE NOW | Influenza 2024-25 | Administer | Annual | CVX 158

SUGGESTED VISIT 1 (today): DTaP + IPV + MMR + Varicella + Hep A + Influenza
SUGGESTED VISIT 2 (4 weeks): MMR dose 2
SUGGESTED VISIT 3 (3 months): Varicella dose 2
SUGGESTED VISIT 4 (6 months): DTaP dose 5 + Hep A dose 2 + IPV dose 4
```

## Troubleshooting

### Immunization search returns zero results but patient has vaccination history

- Some systems store historical immunizations with `status` = `not-done` or use a different status for imported records. Try without the status filter: `patient=[id]&_sort=-date&_count=100`.
- Check if vaccinations are stored as Observation resources instead of Immunization resources (non-standard but occurs in some implementations).
- Pre-EHR vaccinations may not be digitized. Note this in the report as "vaccination history may be incomplete -- records prior to [EHR go-live date] not available."

### CVX codes not present in Immunization.vaccineCode

- Fall back to `vaccineCode.text` for display name matching against known vaccine names.
- Some servers use NDC codes or proprietary codes. Map text descriptions to CVX codes using references/adult-immunization-schedule.md and references/pediatric-immunization-schedule.md which include common brand names.
- If only a generic code like "influenza vaccine" is present without specifying formulation, treat as the standard dose formulation for the patient's age group.

### Contraindication logic disagrees with documented immunizations

- Historical administration overrides theoretical contraindication (the vaccine was already safely given). Note the discrepancy but do not retroactively flag past doses.
- Egg allergy severity matters: mild egg allergy is not a contraindication for standard influenza vaccine. Only severe anaphylaxis to egg requires special formulation (cell-based or recombinant). Check `AllergyIntolerance.reaction.severity`.

## Related Skills

- `preventive-care-compliance-report` -- includes immunization status as part of comprehensive preventive audit
- `patient-demographics-summary` -- patient demographic context for immunization decisions
- `allergy-adverse-reaction-summary` -- detailed allergy review for contraindication assessment
