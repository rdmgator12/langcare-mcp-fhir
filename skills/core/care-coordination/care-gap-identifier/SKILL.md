---
name: care-gap-identifier
description: |
  Screens for preventive care gaps per USPSTF A and B recommendations by checking existing Procedures, Observations,
  and Immunizations against age/sex-appropriate due dates.
  Use when user asks to "check care gaps", "preventive care due", "screenings overdue", "what is this patient due for",
  "wellness check", "annual physical prep", or mentions "quality measures" or "HEDIS gaps".
  Do NOT use for disease-specific management, acute care assessment, or medication review.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: care-coordination
---

# Care Gap Identifier

## Overview

Screen a patient for preventive care gaps based on USPSTF grade A and B recommendations. Query FHIR resources for prior screenings, immunizations, and relevant observations. Compare against age, sex, and risk-factor-adjusted intervals. Generate a prioritized list of overdue and upcoming preventive care items with patient-friendly descriptions.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Demographics (age, sex, risk factors) | birthDate, gender, extension (race/ethnicity) |
| Condition | Risk factors modifying screening criteria | code, clinicalStatus, onsetDateTime |
| Observation | Screening results (labs, vital signs) | code, valueQuantity, effectiveDateTime, status |
| Procedure | Completed screenings (colonoscopy, mammogram) | code, performedDateTime, status |
| Immunization | Vaccination history | vaccineCode, occurrenceDateTime, status |
| DiagnosticReport | Imaging results (mammogram, LDCT, DEXA) | code, effectiveDateTime, conclusion |
| MedicationStatement | Medications affecting screening (e.g., statins) | medicationCodeableConcept, status |
| FamilyMemberHistory | Family risk factors (cancer, CVD) | condition, relationship |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract and calculate:
- Age from `birthDate`
- Sex from `gender` (administrative gender; biological sex may differ)
- Smoking status: check Observation with LOINC 72166-2

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=72166-2&_sort=-date&_count=1"
```

### Step 2: Retrieve Active Conditions (Risk Factors)

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Identify risk-modifying conditions:
- Diabetes (SNOMED 44054006) -- affects statin, BP screening thresholds
- HIV (SNOMED 86406008) -- affects STI, cancer screening frequency
- Obesity (SNOMED 414916001) -- affects diabetes screening age
- Immunocompromised states -- affects immunization recommendations
- Family history of colorectal cancer -- earlier screening start
- BRCA carrier status -- different breast cancer screening

### Step 3: Retrieve Family History

```
Tool: fhir_search
resourceType: "FamilyMemberHistory"
queryParams: "patient=[patient-id]"
```

Key family history items affecting screening:
- First-degree relative with colorectal cancer before age 60 -- start screening at 40 or 10 years before youngest case
- First-degree relative with breast cancer -- discuss earlier mammography
- Family history of abdominal aortic aneurysm -- lower threshold for AAA screening

### Step 4: Check Cancer Screenings

#### Breast Cancer Screening (assigned female, age 50-74)
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&code=http://loinc.org|24606-6&_sort=-date&_count=1"
```
LOINC 24606-6 = Mammography screening. Interval: every 2 years. Also check Procedure for SNOMED 71651007 (mammography).

#### Cervical Cancer Screening (assigned female, age 21-65)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=10524-7&_sort=-date&_count=1"
```
LOINC 10524-7 = Cervical cytology (Pap smear). Interval: every 3 years (Pap alone, age 21-65) or every 5 years (Pap + HPV cotesting, age 30-65). Also check HPV test LOINC 21440-3.

#### Colorectal Cancer Screening (age 45-75)
Check multiple modalities:
```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|73761001&_sort=-date&_count=1"
```
SNOMED 73761001 = Colonoscopy. Interval: every 10 years.

Also check stool-based testing:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=57905-2&_sort=-date&_count=1"
```
LOINC 57905-2 = Fecal immunochemical test (FIT). Interval: annually.

#### Lung Cancer Screening (age 50-80, 20+ pack-year smoking history)
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&code=http://loinc.org|87278-0&_sort=-date&_count=1"
```
LOINC 87278-0 = Low-dose CT chest screening. Interval: annually. Only for current smokers or those who quit within 15 years with 20+ pack-year history.

### Step 5: Check Metabolic and Cardiovascular Screenings

#### Diabetes Screening (age 35-70, overweight/obese)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=4548-4&_sort=-date&_count=1"
```
LOINC 4548-4 = HbA1c. Also check fasting glucose (1558-6). Interval: every 3 years if normal.

#### Lipid Screening (age 40-75)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=57698-3&_sort=-date&_count=1"
```
LOINC 57698-3 = Lipid panel. Interval: every 5 years (more frequently if borderline or on treatment).

#### Hypertension Screening (all adults)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=85354-9&_sort=-date&_count=1"
```
LOINC 85354-9 = Blood pressure panel. Interval: annually.

### Step 6: Check Immunizations

```
Tool: fhir_search
resourceType: "Immunization"
queryParams: "patient=[patient-id]&status=completed"
```

Check against adult immunization schedule:
- **Influenza**: CVX 158 (IIV4), CVX 197 (adjuvanted), CVX 185 (recombinant). Annual.
- **Tdap/Td**: CVX 115 (Tdap). Every 10 years.
- **Pneumococcal**: CVX 215 (PCV20) or CVX 133 (PCV15) + CVX 33 (PPSV23). Age 65+ or high-risk.
- **Shingles**: CVX 187 (Shingrix). Age 50+, 2-dose series.
- **Hepatitis B**: CVX 43 (HepB). Universal adult recommendation if not previously vaccinated.
- **COVID-19**: CVX 213, 228, 229. Per current CDC schedule.
- **HPV**: CVX 165 (HPV9). Through age 26 (shared decision 27-45).

### Step 7: Check Depression and Mental Health Screening

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=44249-1&_sort=-date&_count=1"
```
LOINC 44249-1 = PHQ-9 total score. Interval: annually. Also check PHQ-2 (55758-7).

### Step 8: Check STI Screening

#### HIV Screening (age 15-65)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=75622-1&_sort=-date&_count=1"
```
LOINC 75622-1 = HIV 1/2 Ag+Ab. At least once; annually if high-risk.

#### Hepatitis C Screening (all adults 18-79)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=16128-1&_sort=-date&_count=1"
```
LOINC 16128-1 = Hepatitis C antibody. One-time screening.

#### Hepatitis B Screening (all adults 18+)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=5195-3&_sort=-date&_count=1"
```
LOINC 5195-3 = Hepatitis B surface antigen. One-time screening.

### Step 9: Check Additional USPSTF Screenings

#### Abdominal Aortic Aneurysm (males age 65-75, ever smoked)
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&code=http://loinc.org|24850-0&_sort=-date&_count=1"
```
LOINC 24850-0 = Abdominal ultrasound. One-time screening.

#### Osteoporosis (females age 65+, or postmenopausal with risk factors)
```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&code=http://loinc.org|38269-7&_sort=-date&_count=1"
```
LOINC 38269-7 = DEXA scan. Interval: per risk assessment (typically every 2 years if abnormal).

### Step 10: Format Output

```
PREVENTIVE CARE GAP ANALYSIS
==============================
Patient: [name] | Age: [age] | Sex: [sex]
Smoking Status: [current/former/never] | BMI: [value]

OVERDUE SCREENINGS (action needed)
-----------------------------------
1. [Screening name] -- Last: [date or Never] -- Due: [when] -- [Patient-friendly description]
2. [Screening name] -- Last: [date or Never] -- Due: [when] -- [Patient-friendly description]

UPCOMING SCREENINGS (within 6 months)
--------------------------------------
1. [Screening name] -- Due: [date] -- [Patient-friendly description]

UP TO DATE
----------
1. [Screening name] -- Last: [date] -- Next due: [date]

NOT APPLICABLE
--------------
1. [Screening name] -- Reason: [age/sex/risk factor exclusion]

IMMUNIZATION GAPS
-----------------
1. [Vaccine] -- Status: [Overdue/Never received/Series incomplete]

PATIENT SUMMARY (plain language)
---------------------------------
[2-3 sentence summary suitable for sharing with the patient, describing
what screenings are overdue and why they are important]
```

## Examples

### Example 1: Annual Physical Prep for 55-Year-Old Male

**User says:** "What preventive care is patient 54321 due for?"

**Actions:**
1. `fhir_read` Patient/54321 -- returns James Miller, DOB 1969-06-14, Male, age 55
2. `fhir_search` Observation?patient=54321&code=72166-2 -- current smoker, 30 pack-years
3. `fhir_search` Condition?patient=54321&clinical-status=active -- hypertension, obesity (BMI 32)
4. `fhir_search` Procedure?patient=54321&code=73761001 -- colonoscopy 2018 (6 years ago)
5. `fhir_search` Observation?patient=54321&code=4548-4 -- HbA1c 5.9% (2023-03-01, 10 months ago)
6. `fhir_search` Observation?patient=54321&code=57698-3 -- lipid panel 2021 (3 years ago)
7. `fhir_search` DiagnosticReport?patient=54321&code=87278-0 -- no LDCT on file
8. `fhir_search` Immunization?patient=54321&status=completed -- flu 2023, Td 2019 (no Tdap)
9. `fhir_search` Observation?patient=54321&code=44249-1 -- PHQ-9 not on file
10. `fhir_search` Observation?patient=54321&code=75622-1 -- HIV never tested

**Result:**
```
PREVENTIVE CARE GAP ANALYSIS
==============================
Patient: James Miller | Age: 55 | Sex: Male
Smoking Status: Current (30 pack-years) | BMI: 32

OVERDUE SCREENINGS (action needed)
-----------------------------------
1. Lung cancer screening (LDCT) -- Last: Never -- Due: Now -- Annual low-dose CT chest recommended for current smokers age 50-80 with 20+ pack-year history
2. Depression screening (PHQ-9) -- Last: Never -- Due: Now -- Annual screening recommended for all adults
3. HIV screening -- Last: Never -- Due: Now -- One-time screening recommended for all adults 15-65
4. Hepatitis C screening -- Last: Never -- Due: Now -- One-time screening for all adults 18-79

UPCOMING SCREENINGS (within 6 months)
--------------------------------------
1. HbA1c / diabetes screening -- Last: 2023-03-01 (prediabetes 5.9%) -- Due: 2024-03-01 -- Annual recheck recommended given prediabetes + obesity

UP TO DATE
----------
1. Colonoscopy -- Last: 2018 -- Next due: 2028 (10-year interval)
2. Blood pressure -- Monitored via HTN management
3. Influenza vaccine -- Received 2023

IMMUNIZATION GAPS
-----------------
1. Tdap -- Never received (Td 2019, but Tdap preferred). Recommend Tdap at next opportunity
2. Shingrix -- Age 50+, 2-dose series not started
3. Pneumococcal (PCV20) -- Not yet indicated (age < 65, review risk factors)

PATIENT SUMMARY
---------------------------------
You are due for a lung cancer screening CT scan because of your smoking history. You should also have a depression screening questionnaire, an HIV test, and a hepatitis C test, which are recommended for all adults. Your prediabetes should be rechecked soon with a blood sugar test.
```

### Example 2: Young Female Patient

**User says:** "Check care gaps for patient xyz-100"

**Actions:**
1. `fhir_read` Patient/xyz-100 -- returns Emma Johnson, DOB 1996-09-22, Female, age 27
2. `fhir_search` Observation?patient=xyz-100&code=72166-2 -- never smoker
3. `fhir_search` Condition?patient=xyz-100&clinical-status=active -- no active conditions
4. `fhir_search` Observation?patient=xyz-100&code=10524-7 -- Pap smear 2021 (3 years ago)
5. `fhir_search` Observation?patient=xyz-100&code=85354-9 -- BP 118/72 (2023-06-01)
6. `fhir_search` Immunization?patient=xyz-100&status=completed -- HPV series complete, flu 2022, Tdap 2020
7. `fhir_search` Observation?patient=xyz-100&code=44249-1 -- PHQ-9 score 3 (2023-06-01)
8. `fhir_search` Observation?patient=xyz-100&code=75622-1 -- HIV negative (2022)
9. `fhir_search` Observation?patient=xyz-100&code=16128-1 -- HCV negative (2022)

**Result:**
```
PREVENTIVE CARE GAP ANALYSIS
==============================
Patient: Emma Johnson | Age: 27 | Sex: Female
Smoking Status: Never | BMI: 23.1

OVERDUE SCREENINGS (action needed)
-----------------------------------
1. Cervical cancer screening (Pap) -- Last: 2021 -- Due: 2024 -- Pap smear every 3 years for ages 21-29

UPCOMING SCREENINGS (within 6 months)
--------------------------------------
(none)

UP TO DATE
----------
1. Blood pressure -- Last: 2023-06-01 (118/72, normal)
2. Depression screening -- Last: 2023-06-01 (PHQ-9: 3, normal)
3. HIV screening -- Completed 2022
4. Hepatitis C screening -- Completed 2022
5. HPV vaccination -- Series complete
6. Tdap -- Current (2020, next due 2030)

NOT APPLICABLE
--------------
1. Mammography -- Not indicated until age 40 (or earlier if high-risk family history)
2. Colorectal cancer screening -- Not indicated until age 45
3. Lung cancer screening -- Never smoker
4. Diabetes screening -- Age < 35, not overweight
5. Lipid screening -- Age < 40 (unless risk factors)
6. AAA screening -- Female, not indicated
7. Osteoporosis screening -- Premenopausal, not indicated
8. Shingrix -- Age < 50

IMMUNIZATION GAPS
-----------------
1. Influenza -- Last: 2022. Overdue for annual flu vaccine

PATIENT SUMMARY
---------------------------------
You are due for your Pap smear, which screens for cervical cancer and is recommended every 3 years. You should also get your annual flu shot. Everything else is up to date.
```

## Troubleshooting

### Screening Dates Not Found in Expected Resources
- Mammography may be stored as Procedure (SNOMED 71651007), DiagnosticReport, or Observation. Search all three resource types.
- External screenings (done at another facility) may not be in the FHIR server. If critical screenings show as "never done" for a patient who reports otherwise, note that the record may be incomplete.

### Patient Gender Does Not Match Expected Screening Criteria
- FHIR `Patient.gender` is administrative gender. Sex-specific screenings (cervical, breast, prostate) should be based on biological sex, which may be in an extension (`http://hl7.org/fhir/us/core/StructureDefinition/us-core-birthsex`). Check the extension before excluding screenings.
- If birth sex extension is unavailable, apply screenings based on administrative gender but note the limitation.

### Immunization Records Are Incomplete
- Patients vaccinated before EHR adoption may not have historical records. If Immunization search returns zero results, note "No immunization records found -- may need patient-reported history or state immunization registry query."
- Check state IIS (Immunization Information System) if FHIR server supports it.

## Related Skills

- `clinical-summary-generator` -- include care gap summary in clinical notes
- `referral-generator` -- generate referrals for identified screening needs (e.g., colonoscopy referral)
- `follow-up-task-generator` -- create tasks for overdue screenings
