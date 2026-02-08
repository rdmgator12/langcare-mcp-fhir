---
name: oncology-treatment-timeline
description: |
  Maps cancer treatment history from FHIR resources into a structured timeline of diagnosis, staging,
  treatment, response assessment, and surveillance. Use when user asks to "review cancer treatment",
  "oncology timeline", "treatment history", "cancer staging", "tumor markers", "chemo regimen",
  "treatment response", mentions a specific cancer type, or needs a longitudinal oncology summary.
  Do NOT use for cancer screening (use preventive-care-compliance-report), benign tumors, or palliative care planning.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: specialty
---

# Oncology Treatment Timeline

## Overview

Construct a comprehensive cancer treatment timeline from FHIR resources. Pull cancer diagnosis Conditions with staging data, chemotherapy MedicationRequests, surgical and radiation Procedures, tumor marker Observations, and imaging DiagnosticReports. Organize chronologically into phases: diagnosis/staging, neoadjuvant treatment, definitive treatment (surgery/radiation), adjuvant treatment, response assessment, and surveillance. Track treatment response through tumor marker trends and imaging results using RECIST criteria where applicable. Identify current treatment phase and upcoming milestones.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Demographics, age | birthDate, gender |
| Condition | Cancer diagnosis, staging, comorbidities | code, stage, clinicalStatus, onsetDateTime |
| MedicationRequest | Chemotherapy regimens, supportive care | medicationCodeableConcept, status, dosageInstruction, authoredOn |
| Procedure | Surgery, radiation, biopsy | code, performedDateTime, status, outcome |
| Observation | Tumor markers, performance status, lab values | code, valueQuantity, effectiveDateTime, interpretation |
| DiagnosticReport | Pathology, imaging | code, conclusion, result, effectiveDateTime |
| CarePlan | Treatment plan documentation | category, activity, status |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract age, gender. Relevant for gender-specific cancers and age-adjusted treatment protocols.

### Step 2: Pull Cancer Diagnosis

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&category=encounter-diagnosis&code:below=363346000"
```

SNOMED 363346000 = Malignant neoplastic disease (parent concept). This captures all malignant neoplasm subtypes.

If SNOMED search fails, try ICD-10 range:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code:below=C00"
```

ICD-10 C00-C97 covers all malignant neoplasms.

For each cancer Condition, extract:
- `code.coding` -- cancer type (SNOMED or ICD-10)
- `onsetDateTime` -- date of diagnosis
- `stage.summary` -- staging information (TNM)
- `stage.type` -- staging system used
- `clinicalStatus` -- active, remission, recurrence
- `evidence` -- references to supporting diagnostic resources

### Step 3: Pull Staging Details

If `stage` is not populated in the Condition, search for staging Observations:

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=21908-9,21902-2,21914-7,21913-9&_sort=-date"
```

LOINC staging codes:
- 21908-9 = Stage group (overall stage I-IV)
- 21902-2 = Tumor primary (T stage)
- 21914-7 = Nodes regional (N stage)
- 21913-9 = Distant metastasis (M stage)

Also pull performance status:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=89247-1,89262-0&_sort=-date&_count=5"
```

- LOINC 89247-1 = ECOG performance status
- LOINC 89262-0 = Karnofsky performance status

See references/staging-systems.md for TNM staging interpretation and performance status scales.

### Step 4: Pull Chemotherapy/Systemic Therapy

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&category=outpatient,inpatient&_sort=-date&_count=100"
```

Filter for antineoplastic agents. Common chemotherapy RxNorm codes by cancer type -- see references/oncology-treatment.md.

For each MedicationRequest:
- `medicationCodeableConcept.coding` -- drug name
- `authoredOn` -- start date
- `dosageInstruction` -- dose, route, frequency
- `status` -- active, completed, stopped
- `statusReason` -- reason for stopping (toxicity, progression, completed course)
- `reasonReference` -- link to cancer Condition

Group into regimen names when multiple agents started on the same date (e.g., FOLFOX = 5-FU + leucovorin + oxaliplatin).

### Step 5: Pull Surgical and Radiation Procedures

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&_sort=-date&_count=50"
```

Filter for cancer-related procedures:
- Surgery: look for codes containing excision, resection, mastectomy, colectomy, prostatectomy, lobectomy, etc.
- Radiation: SNOMED 108290001 (radiation therapy), CPT 77401-77499 range
- Biopsy: SNOMED 86273004 (biopsy)

For each Procedure:
- `code.coding` -- procedure name
- `performedDateTime` or `performedPeriod` -- date(s)
- `outcome` -- result/outcome
- `report` -- reference to pathology report
- `reasonReference` -- link to cancer Condition

### Step 6: Pull Tumor Markers

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2857-1,10466-1,19167-7,24108-3,1834-1,21198-7&_sort=date&_count=100"
```

Common tumor marker LOINC codes:
- 2857-1 = PSA (prostate)
- 10466-1 = CEA (colorectal, others)
- 19167-7 = CA-125 (ovarian)
- 24108-3 = CA 19-9 (pancreatic, others)
- 1834-1 = AFP (liver, germ cell)
- 21198-7 = Beta-HCG (germ cell, gestational trophoblastic)

For each marker, build a time series: date, value, unit, interpretation. Calculate percent change between consecutive values. See references/oncology-treatment.md for reference values and trending interpretation.

### Step 7: Pull Imaging Results

```
Tool: fhir_search
resourceType: "DiagnosticReport"
queryParams: "patient=[patient-id]&category=imaging&_sort=-date&_count=20"
```

Extract:
- `code.coding` -- imaging type (CT, MRI, PET-CT)
- `effectiveDateTime` -- date
- `conclusion` -- radiologist impression
- `result` -- references to Observation resources with measurements

Look for RECIST response criteria in conclusions:
- Complete Response (CR): disappearance of all target lesions
- Partial Response (PR): >=30% decrease in sum of diameters
- Stable Disease (SD): neither PR nor PD criteria met
- Progressive Disease (PD): >=20% increase in sum of diameters or new lesions

### Step 8: Pull Supportive Care Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

Identify supportive care:
- Antiemetics: ondansetron, granisetron, palonosetron, aprepitant
- Growth factors: filgrastim, pegfilgrastim (G-CSF), epoetin alfa
- Pain management: opioids, gabapentin, NSAIDs
- Bone protection: zoledronic acid, denosumab
- Prophylactic antibiotics, antifungals

### Step 9: Construct Timeline

Organize all events chronologically into phases:

```
ONCOLOGY TREATMENT TIMELINE -- [Patient Name]
Diagnosis: [Cancer type] | Stage: [stage] | Diagnosed: [date]
ECOG: [score] | Current status: [active treatment / surveillance / remission]
==========================================================================

PHASE 1: DIAGNOSIS & STAGING ([date range])
  [date] Biopsy: [pathology result]
  [date] Staging: [TNM stage, grade, biomarkers]
  [date] Imaging: [baseline scan results]
  [date] Tumor markers: [baseline values]

PHASE 2: NEOADJUVANT THERAPY ([date range])  [if applicable]
  [date]-[date] [Regimen name]: [drugs, doses, # cycles]
  Response: [marker trend, imaging response]

PHASE 3: DEFINITIVE TREATMENT ([date range])
  [date] Surgery: [procedure, pathology outcome, margins]
  [date]-[date] Radiation: [dose, fractions, field]

PHASE 4: ADJUVANT THERAPY ([date range])  [if applicable]
  [date]-[date] [Regimen name]: [drugs, doses, # cycles]
  Toxicities: [noted adverse effects]
  Dose modifications: [if any]

PHASE 5: RESPONSE ASSESSMENT
  [date] Imaging: [RECIST response category]
  Tumor markers: [trend summary]
  Performance status: ECOG [score]

PHASE 6: SURVEILLANCE ([date range - ongoing])
  [date] Follow-up imaging: [result]
  [date] Tumor markers: [values and trend]
  Next scheduled: [upcoming appointments/scans]

TUMOR MARKER TREND
==================
[Marker]: [date1]=[val1] -> [date2]=[val2] -> [date3]=[val3] [trend direction]

CURRENT STATUS
==============
Phase: [active treatment / surveillance / progression]
Current regimen: [if active] or Last treatment: [if surveillance]
Next action: [scan/visit/treatment due date]
```

## Examples

### Example 1: Colorectal Cancer Treatment Review

**User says:** "Review oncology treatment history for patient ONC-3301"

**Actions:**
1. `fhir_read` Patient/ONC-3301 -- 61-year-old male
2. `fhir_search` Condition?patient=ONC-3301&category=encounter-diagnosis&code:below=363346000 -- Stage III colon adenocarcinoma (SNOMED 93761005), diagnosed 2025-06-15
3. `fhir_search` Observation?patient=ONC-3301&code=21908-9,21902-2,21914-7,21913-9&_sort=-date -- T3N1M0, Stage IIIB
4. `fhir_search` Procedure?patient=ONC-3301&_sort=-date&_count=50 -- right hemicolectomy 2025-07-10, port placement 2025-08-01
5. `fhir_search` MedicationRequest?patient=ONC-3301&_sort=-date&_count=100 -- FOLFOX regimen started 2025-08-15 (5-FU, leucovorin, oxaliplatin), completed 12 cycles 2026-01-20
6. `fhir_search` Observation?patient=ONC-3301&code=10466-1&_sort=date&_count=100 -- CEA: pre-op 12.4, post-op 2.1, nadir 1.0, current 1.2
7. `fhir_search` DiagnosticReport?patient=ONC-3301&category=imaging&_sort=-date -- CT chest/abdomen/pelvis at 3-month intervals

**Result:**
```
ONCOLOGY TREATMENT TIMELINE -- James Wilson
Diagnosis: Colon adenocarcinoma | Stage IIIB (T3N1M0) | Diagnosed: 2025-06-15
ECOG: 0 | Current status: Surveillance (adjuvant completed)
==========================================================================

PHASE 1: DIAGNOSIS & STAGING (2025-06)
  2025-06-10 Colonoscopy with biopsy: moderately differentiated adenocarcinoma, ascending colon
  2025-06-15 CT C/A/P: 4.2cm ascending colon mass, 2 enlarged pericolic lymph nodes, no distant metastasis
  2025-06-15 CEA: 12.4 ng/mL (elevated, ref <5.0)
  Staging: T3N1M0, Stage IIIB, MSI-stable, KRAS wild-type

PHASE 3: DEFINITIVE TREATMENT (2025-07)
  2025-07-10 Right hemicolectomy: R0 resection, clear margins (>5cm), 2/18 lymph nodes positive
  Post-op CEA (2025-07-25): 2.1 ng/mL (declining)

PHASE 4: ADJUVANT THERAPY (2025-08 to 2026-01)
  2025-08-15 to 2026-01-20: FOLFOX x12 cycles
    - 5-FU 400mg/m2 bolus + 2400mg/m2 46-hr infusion
    - Leucovorin 400mg/m2
    - Oxaliplatin 85mg/m2
  Toxicities: Grade 2 peripheral neuropathy (oxaliplatin), dose reduced to 65mg/m2 for cycles 10-12
  CEA nadir: 1.0 ng/mL (cycle 8)

PHASE 6: SURVEILLANCE (2026-01 - ongoing)
  2026-01-30 CT C/A/P: No evidence of recurrence
  2026-01-30 CEA: 1.2 ng/mL (stable, within normal)
  ECOG: 0

TUMOR MARKER TREND (CEA, ref <5.0 ng/mL)
  2025-06-15: 12.4 -> 2025-07-25: 2.1 -> 2025-10-15: 1.0 -> 2026-01-30: 1.2 [STABLE AT NORMAL]

NEXT SCHEDULED
  - CT C/A/P: due 2026-04 (3-month surveillance scan)
  - CEA: due 2026-04
  - Colonoscopy: due 2026-07 (1-year post-op)
```

### Example 2: Active Treatment with Progression

**User says:** "Pull cancer timeline for patient ONC-7784"

**Actions:**
1. `fhir_read` Patient/ONC-7784 -- 55-year-old female
2. `fhir_search` Condition?patient=ONC-7784&code:below=363346000 -- Stage IV NSCLC (SNOMED 254637007), diagnosed 2025-03-01, adenocarcinoma, EGFR mutation positive
3. `fhir_search` MedicationRequest?patient=ONC-7784&_sort=-date&_count=100 -- osimertinib (1st line) started 2025-04-01, stopped 2025-11-15 (progression); carboplatin + pemetrexed started 2025-12-01 (2nd line, active)
4. `fhir_search` DiagnosticReport?patient=ONC-7784&category=imaging&_sort=-date -- serial CT scans showing initial response then progression
5. `fhir_search` Observation?patient=ONC-7784&code=89247-1&_sort=-date -- ECOG 1

**Result:**
```
ONCOLOGY TREATMENT TIMELINE -- Linda Park
Diagnosis: NSCLC adenocarcinoma | Stage IV (T2N2M1a) | Diagnosed: 2025-03-01
Biomarkers: EGFR exon 19 deletion positive | PD-L1 TPS 40%
ECOG: 1 | Current status: Active treatment (2nd line)
==========================================================================

PHASE 1: DIAGNOSIS & STAGING (2025-03)
  2025-03-01 CT-guided lung biopsy: adenocarcinoma, EGFR exon 19 deletion
  2025-03-05 PET-CT: 3.5cm RUL mass, mediastinal LN, bilateral pulmonary nodules (M1a)

PHASE 2: 1ST LINE THERAPY (2025-04 to 2025-11)
  2025-04-01 to 2025-11-15: Osimertinib 80mg PO daily (EGFR-TKI)
  2025-06-15 CT: Partial response (PR) -- primary tumor 2.1cm (40% decrease)
  2025-09-15 CT: Stable disease (SD) -- 2.3cm
  2025-11-15 CT: Progressive disease (PD) -- primary 3.8cm, new liver metastasis
  Duration of response: 7.5 months

PHASE 3: 2ND LINE THERAPY (2025-12 - ongoing)
  2025-12-01: Carboplatin AUC 5 + Pemetrexed 500mg/m2 q21d
  Cycle 3 completed 2026-01-25
  Restaging CT due after cycle 4

ECOG PERFORMANCE STATUS TREND
  2025-03: ECOG 0 -> 2025-11: ECOG 1 -> 2026-01: ECOG 1

CURRENT STATUS
  Phase: Active 2nd-line treatment, cycle 3 of 4-6 planned
  Next restaging: CT C/A/P due ~2026-02-20 (after cycle 4)
  Considerations: If progression, consider immunotherapy or clinical trial
```

## Troubleshooting

### Cancer Condition Has No Stage Information

- Search for staging Observations using LOINC codes (21908-9, 21902-2, 21914-7, 21913-9).
- Check DiagnosticReport resources for pathology reports that may contain staging in `conclusion` text.
- Check Procedure resources for surgical pathology outcomes that include pathologic staging.
- If no staging found, note "staging data unavailable" and recommend querying the oncology team.

### Chemotherapy Agents Not Identifiable from MedicationRequest

- Filter MedicationRequest by `category=outpatient` and look for known antineoplastic RxNorm codes.
- Some systems store chemotherapy protocols in CarePlan resources. Search: `fhir_search` CarePlan?patient=[id]&category=assess-plan.
- Check `MedicationRequest.reasonReference` for links to the cancer Condition -- this distinguishes chemo from other medications.
- If medication names are in local codes, use `medicationCodeableConcept.text` for display name matching.

### Tumor Markers Return No Results

- Not all cancers have standard tumor markers. Some cancers are monitored by imaging alone.
- Try broader search: `fhir_search` Observation?patient=[id]&category=laboratory&_sort=-date&_count=50 and inspect for marker-like results.
- PSA may be coded as LOINC 2857-1 (total PSA) or 35741-8 (PSA in serum). Try both.

## Related Skills

- `lab-result-interpreter` -- for detailed interpretation of tumor marker panels and treatment-related labs
- `medication-reconciliation` -- for comprehensive medication review including supportive care
- `clinical-summary-generator` -- for full patient summary with oncology context
- `chronic-pain-management-review` -- for cancer pain management assessment
