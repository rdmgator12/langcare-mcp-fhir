---
name: chronic-disease-registry-query
description: |
  Builds and queries disease registries from FHIR data for diabetes, hypertension, COPD, asthma, CHF, and CKD. Generates patient counts, severity distributions, treatment patterns, and outcome trends. Use when user asks to "query registry", "disease registry", "registry report", "how many patients with [condition]", "severity breakdown", "treatment patterns", "disease outcomes", or mentions "registry management", "chronic disease tracking". Do NOT use for individual patient review, quality measure calculation (use quality-measure-dashboard), or general panel overview without disease-specific depth (use patient-panel-overview).
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: population-health
---

# Chronic Disease Registry Query

## Overview

Construct virtual disease registries by querying FHIR Condition, Observation, MedicationRequest, and Encounter resources. Support six registries: diabetes, hypertension, COPD, asthma, CHF, and CKD. For each registry, determine patient enrollment, severity/stage distribution, current treatment patterns, outcome metric trends, and intervention triggers. See references/registry-criteria.md for inclusion codes and stratification logic.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Condition | Registry inclusion, comorbidities, complications | code, clinicalStatus, severity, onsetDateTime, subject |
| Observation | Severity metrics, outcome tracking | code, valueQuantity, effectiveDateTime, component |
| MedicationRequest | Treatment pattern analysis | medicationCodeableConcept, status, dosageInstruction, authoredOn |
| Encounter | Visit frequency, hospitalization tracking | class, status, period, reasonCode, subject |
| Patient | Demographics, age/sex stratification | birthDate, gender, deceasedBoolean |
| Procedure | Surgical interventions, dialysis | code, performedDateTime, status |
| CarePlan | Care plan adherence | status, category, activity |

## Instructions

### Step 1: Select Registry

Determine which disease registry to query. Supported registries and their primary codes (see references/registry-criteria.md for full code sets):

| Registry | SNOMED Code | ICD-10 Prefix | Key Metric LOINC |
|----------|-------------|---------------|------------------|
| Diabetes | 44054006, 46635009 | E10, E11 | 4548-4 (A1c) |
| Hypertension | 38341003 | I10, I11-I13 | 85354-9 (BP panel) |
| COPD | 13645005 | J44 | 19926-5 (FEV1/FVC) |
| Asthma | 195967001 | J45 | 19926-5 (FEV1/FVC) |
| CHF | 42343007 | I50 | 42637-9 (BNP), 33762-6 (NT-proBNP) |
| CKD | 709044004 | N18 | 33914-3 (eGFR), 14959-1 (Microalbumin/Cr) |

### Step 2: Build Registry Population

Query patients with the target condition:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "code=[snomed-code]&clinical-status=active&_count=200&_include=Condition:subject"
```

If zero results, retry with ICD-10:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "code=http://hl7.org/fhir/sid/icd-10-cm|[icd10-prefix]&clinical-status=active&_count=200"
```

Collect all unique patient IDs. Fetch Patient demographics:
```
Tool: fhir_search
resourceType: "Patient"
queryParams: "_id=[id1],[id2],...&_count=50"
```

Exclude deceased patients (`deceasedBoolean` = true or `deceasedDateTime` present).

### Step 3: Determine Severity/Stage Distribution

Query condition-specific severity metrics.

**Diabetes -- A1c-based stratification:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=http://loinc.org|4548-4&date=ge[12-months-ago]&_sort=-date&_count=200"
```
Stratify: Well controlled (A1c < 7%), Moderate (7-9%), Poor (> 9%), Unknown (no recent A1c).

**Hypertension -- BP-based staging:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=http://loinc.org|85354-9&date=ge[6-months-ago]&_sort=-date&_count=200"
```
Stratify per AHA: Normal (< 120/80), Elevated (120-129/<80), Stage 1 (130-139/80-89), Stage 2 (>= 140/90), Crisis (> 180/120).

**CHF -- BNP/NYHA-based:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=http://loinc.org|42637-9&date=ge[12-months-ago]&_sort=-date&_count=200"
```
LOINC 42637-9 = BNP. Also check for NYHA classification in Condition.severity or Observation with LOINC 88020-3.

**CKD -- eGFR staging:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=http://loinc.org|33914-3&date=ge[12-months-ago]&_sort=-date&_count=200"
```
Stratify: Stage 1 (eGFR >= 90 with kidney damage), Stage 2 (60-89), Stage 3a (45-59), Stage 3b (30-44), Stage 4 (15-29), Stage 5 (< 15).

**COPD/Asthma -- Spirometry:**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=http://loinc.org|19926-5&date=ge[12-months-ago]&_sort=-date&_count=200"
```
LOINC 19926-5 = FEV1/FVC. GOLD staging for COPD: Mild (>= 80%), Moderate (50-79%), Severe (30-49%), Very Severe (< 30%).

### Step 4: Analyze Treatment Patterns

Query active medications for registry patients by therapeutic class.

**Diabetes medications:**
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&_count=50"
```

Classify medications into: Metformin, Sulfonylureas, DPP-4 inhibitors, GLP-1 agonists, SGLT2 inhibitors, Insulin (basal, bolus, premixed), Thiazolidinediones, Other. See references/registry-interventions.md for drug class mappings.

Calculate: % on monotherapy, % on dual therapy, % on triple therapy, % on insulin, % on cardioprotective agents (GLP-1 or SGLT2).

**Hypertension medications:**
Classify: ACE inhibitors, ARBs, Calcium channel blockers, Thiazide diuretics, Beta blockers, Mineralocorticoid receptor antagonists, Other.

**CHF medications:**
Track guideline-directed medical therapy (GDMT): % on ACEi/ARB/ARNI, % on beta blocker, % on MRA, % on SGLT2i, % on all four pillars.

### Step 5: Track Outcome Trends

Query historical observations for trend analysis (last 24 months, quarterly intervals):
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "code=[loinc-code]&date=ge[24-months-ago]&_sort=date&_count=500"
```

For each quarter, compute: mean value, median, % at target, % worsening. Present trend as a table.

### Step 6: Identify Comorbidity Burden

For each registry patient, query additional active conditions:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```

Calculate: mean number of comorbidities, most common comorbidity pairs (e.g., diabetes + hypertension, CHF + CKD), % with >= 3 chronic conditions.

### Step 7: Check for Complications

Query condition-specific complications. See references/registry-criteria.md for complication codes.

**Diabetes complications:**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "code=http://snomed.info/sct|422034002,http://snomed.info/sct|127013003,http://snomed.info/sct|421986006&clinical-status=active&_count=100"
```
SNOMED: 422034002 = Diabetic retinopathy, 127013003 = Diabetic renal disease, 421986006 = Diabetic peripheral neuropathy.

### Step 8: Flag Intervention Triggers

Based on references/registry-interventions.md, flag patients meeting intervention criteria:

- **Therapy escalation**: At-goal metric worsening over 2+ consecutive measurements
- **Specialist referral**: Severity exceeding primary care management threshold
- **Hospitalization risk**: High-risk combination of metrics (e.g., CHF with BNP > 500 + recent ED visit)
- **Care management enrollment**: Multiple intervention triggers active simultaneously

### Step 9: Present Registry Report

```
CHRONIC DISEASE REGISTRY: [Disease Name]
==========================================
Registry Size: [N] patients
Report Date: [today]

DEMOGRAPHICS
------------
Mean Age: [age] | Male: [%] | Female: [%]
Age Distribution: 18-40: [N] | 41-65: [N] | 66+: [N]

SEVERITY DISTRIBUTION
---------------------
[Severity Level 1]: [N] ([%])
[Severity Level 2]: [N] ([%])
[Severity Level 3]: [N] ([%])
Unknown/No Data:    [N] ([%])

TREATMENT PATTERNS
------------------
[Drug Class 1]:     [N] ([%])
[Drug Class 2]:     [N] ([%])
Monotherapy:        [N] ([%])
Dual therapy:       [N] ([%])
Triple+ therapy:    [N] ([%])

OUTCOME TRENDS (Quarterly, Last 12 Months)
-------------------------------------------
Quarter  | Mean [Metric] | % at Target | Trend
Q1 2024  | [value]       | [%]         | --
Q2 2024  | [value]       | [%]         | [up/down/stable]
Q3 2024  | [value]       | [%]         | [up/down/stable]
Q4 2024  | [value]       | [%]         | [up/down/stable]

COMPLICATIONS
-------------
[Complication 1]: [N] ([%])
[Complication 2]: [N] ([%])
[Complication 3]: [N] ([%])

INTERVENTION TRIGGERS
---------------------
Patients needing therapy escalation:  [N]
Patients needing specialist referral: [N]
High hospitalization risk:            [N]
Care management candidates:           [N]
```

## Examples

### Example 1: Diabetes Registry Query

**User says**: "Pull up the diabetes registry. Show severity breakdown and treatment patterns."

**Actions**:
1. Search Condition for SNOMED 44054006 and 46635009 (T1DM + T2DM). Returns 94 patients.
2. Fetch demographics. Exclude 2 deceased. Registry = 92 patients.
3. Query A1c (LOINC 4548-4) in past 12 months. Stratify by control level.
4. Query MedicationRequest for each patient. Classify by drug class.
5. Query diabetes complications (retinopathy, nephropathy, neuropathy).

**Result**:
```
CHRONIC DISEASE REGISTRY: Diabetes Mellitus
=============================================
Registry Size: 92 patients (89 T2DM, 3 T1DM)
Report Date: 2024-11-15

DEMOGRAPHICS
------------
Mean Age: 58.4 | Male: 48 (52.2%) | Female: 44 (47.8%)

SEVERITY (A1c-based)
---------------------
Well Controlled (< 7%):    52 (56.5%)
Moderate (7.0-8.9%):       24 (26.1%)
Poor Control (>= 9%):       9 (9.8%)
No Recent A1c:               7 (7.6%)

TREATMENT PATTERNS
------------------
Metformin:           78 (84.8%)
SGLT2 inhibitor:     31 (33.7%)
GLP-1 agonist:       22 (23.9%)
DPP-4 inhibitor:     12 (13.0%)
Sulfonylurea:        18 (19.6%)
Insulin (any):       24 (26.1%)
Monotherapy:         34 (37.0%)
Dual therapy:        38 (41.3%)
Triple+ therapy:     20 (21.7%)
Cardioprotective (GLP-1/SGLT2): 41 (44.6%)

COMPLICATIONS
-------------
Retinopathy:    11 (12.0%)
Nephropathy:    14 (15.2%)
Neuropathy:     19 (20.7%)
Any complication: 28 (30.4%)

INTERVENTION TRIGGERS
---------------------
Need therapy escalation (A1c rising): 6 patients
Need ophthalmology referral (no eye exam): 23 patients
Need nephrology referral (eGFR < 30): 3 patients
```

### Example 2: CHF Registry with GDMT Analysis

**User says**: "How many heart failure patients are on full GDMT? Show me the gaps."

**Actions**:
1. Search Condition SNOMED 42343007 (CHF). Returns 34 patients.
2. Query MedicationRequest for each patient. Check for 4 GDMT pillars: ACEi/ARB/ARNI, beta blocker, MRA, SGLT2i.
3. Query BNP (LOINC 42637-9) for severity.

**Result**:
```
CHF REGISTRY - GDMT ANALYSIS
==============================
Registry Size: 34 patients

GDMT PILLAR COVERAGE
---------------------
ACEi/ARB/ARNI:    28 (82.4%)
Beta Blocker:      30 (88.2%)
MRA:               18 (52.9%)
SGLT2 Inhibitor:   14 (41.2%)

COMBINATION ANALYSIS
--------------------
All 4 pillars:     11 (32.4%)
3 of 4 pillars:    12 (35.3%)
2 of 4 pillars:     8 (23.5%)
1 or fewer:          3 (8.8%)

LARGEST GAP: SGLT2 inhibitors (20 patients not on one)
- 8 of these have no contraindication documented
- ACTION: Review for SGLT2i initiation

PATIENTS MISSING MULTIPLE PILLARS:
1. Robert Ellis (MRN: 44102) - On beta blocker only - BNP: 890 - HIGH PRIORITY
2. Carol Fisher (MRN: 55210) - On ACEi + BB - BNP: 620 - Add MRA + SGLT2i
3. James Ward (MRN: 33098) - On BB only - BNP: 445 - Needs ACEi/ARB + MRA + SGLT2i
```

## Troubleshooting

### Registry population is much smaller than expected

- Check for alternate coding. Many FHIR servers use ICD-10 exclusively. Add ICD-10 codes to the search: `code=http://hl7.org/fhir/sid/icd-10-cm|E11,http://snomed.info/sct|44054006`.
- Conditions may be stored with `verificationStatus` = `provisional` or `unconfirmed`. Include these if the registry should capture suspected cases: `verification-status=confirmed,provisional`.
- Some systems record historical conditions with `clinical-status` = `inactive` but never update to `active` after confirmation. Consider including `inactive` with recent `onsetDateTime`.

### Medication classification fails due to non-standard codes

- Fall back to `medicationCodeableConcept.text` for string matching against drug class keywords (e.g., "metformin", "lisinopril", "carvedilol").
- If the server uses Medication references instead of CodeableConcept, resolve the reference with `fhir_read` on the Medication resource to get the code.
- Request medications with `_include=MedicationRequest:medication` to pull contained Medication resources inline.

### Observation queries return data for patients outside the registry

- Always filter returned Observations by matching `subject.reference` against the registry patient ID list.
- If the server supports it, use patient-specific queries: `patient=[id]&code=[loinc]&_sort=-date&_count=1` to get the most recent value per patient. This is slower but more precise.

## Related Skills

- `patient-panel-overview` -- lighter-weight panel aggregate without registry depth
- `quality-measure-dashboard` -- formal HEDIS/CMS measure calculation
- `care-gap-identifier` -- individual patient care gap identification
- `diabetes-panel-review` -- individual patient diabetes review
- `renal-function-dashboard` -- individual CKD patient analysis
