---
name: prenatal-visit-workflow
description: |
  Performs structured prenatal visit assessment organized by trimester with ACOG guideline alignment.
  Use when user asks to "review prenatal visit", "prenatal assessment", "OB visit", "pregnancy checkup",
  "trimester labs", "prenatal labs", mentions "gestational age", "fundal height", "prenatal care", or needs
  pregnancy-related clinical data organized by visit schedule.
  Do NOT use for labor and delivery, postpartum care, gynecologic exams, or non-obstetric visits.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: specialty
---

# Prenatal Visit Workflow

## Overview

Pull and organize prenatal clinical data by trimester and gestational age following ACOG prenatal visit schedule. Retrieve pregnancy-related Conditions, maternal vital signs, fetal Observations, laboratory results, and screening outcomes. Track weight gain trajectory, blood pressure trends, and glucose screening results. Flag preeclampsia risk factors, gestational diabetes indicators, preterm labor risk, and deviations from expected visit schedule. Present a structured visit summary with outstanding orders and upcoming milestones.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Maternal demographics, age, gravidity | birthDate, gender, identifier |
| Condition | Pregnancy diagnosis, complications | code, clinicalStatus, onsetDateTime, category |
| Observation | Vitals, fundal height, FHR, weight, urine, labs | code, valueQuantity, effectiveDateTime, component, interpretation |
| MedicationRequest | Prenatal vitamins, aspirin, Rhogam, GBS prophylaxis | medicationCodeableConcept, status, dosageInstruction |
| Procedure | Ultrasounds, amniocentesis, cervical cerclage | code, performedDateTime, status |
| DiagnosticReport | Ultrasound reports, lab panels | code, result, effectiveDateTime |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract `birthDate` (calculate maternal age -- flag if >=35 for advanced maternal age), `gender` (confirm female).

### Step 2: Pull Pregnancy Condition

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=77386006&clinical-status=active"
```

SNOMED 77386006 = Pregnancy (finding). If not found, try:
- SNOMED 72892002 = Normal pregnancy
- ICD-10 Z34 (Encounter for supervision of normal pregnancy)
- ICD-10 O09 (Supervision of high-risk pregnancy)

Extract `onsetDateTime` to determine estimated gestational age (EGA). If EGA is not derivable from onset, search for EDD Observation (LOINC 11778-8 = Delivery date estimated).

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=11778-8&_sort=-date&_count=1"
```

Calculate current gestational age from EDD: GA weeks = 40 - (weeks remaining until EDD).

### Step 3: Determine Trimester and Visit Context

Based on gestational age:
- **First trimester**: 0-13 weeks 6 days
- **Second trimester**: 14-27 weeks 6 days
- **Third trimester**: 28-40+ weeks

Set expected visit frequency:
- Weeks 4-28: every 4 weeks
- Weeks 28-36: every 2 weeks
- Weeks 36-delivery: weekly

### Step 4: Pull Prenatal Vital Signs

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=vital-signs&date=ge[pregnancy-onset]&_sort=-date&_count=100"
```

Extract and trend:
- Blood pressure (LOINC 85354-9): systolic 8480-6, diastolic 8462-4
- Weight (LOINC 29463-7): calculate weight gain from prepregnancy baseline
- BMI (LOINC 39156-5): classify prepregnancy BMI for weight gain targets

Weight gain targets by prepregnancy BMI (IOM guidelines):
- Underweight (BMI <18.5): 28-40 lbs total
- Normal weight (BMI 18.5-24.9): 25-35 lbs total
- Overweight (BMI 25.0-29.9): 15-25 lbs total
- Obese (BMI >=30.0): 11-20 lbs total

### Step 5: Pull Fetal Observations

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=11996-6,11948-7,249043002&date=ge[pregnancy-onset]&_sort=-date&_count=50"
```

LOINC codes:
- 11996-6 = Fetal heart rate
- 11948-7 = Fundal height
- SNOMED 249043002 = Fetal movement (alternatively LOINC 57088-7)

Flag: FHR outside 110-160 bpm, fundal height discrepancy >3 cm from expected (GA in weeks = expected fundal height in cm at 20-36 weeks).

### Step 6: Pull Prenatal Laboratory Results

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&date=ge[pregnancy-onset]&_sort=-date&_count=200"
```

Organize by trimester milestone. See references/acog-prenatal-schedule.md for complete schedule. Key labs to extract:

**First trimester (initial visit)**:
- Blood type and Rh (LOINC 883-9, 10331-7)
- Antibody screen (LOINC 890-4)
- CBC (LOINC 58410-2)
- Rubella IgG (LOINC 5334-8)
- Hepatitis B surface antigen (LOINC 5196-1)
- HIV (LOINC 7917-8)
- RPR/VDRL syphilis (LOINC 20507-0)
- Urine culture (LOINC 630-4)
- Chlamydia/Gonorrhea (LOINC 43304-5, 43305-2)
- Pap smear if due
- Urine protein (LOINC 5804-0)

**Second trimester (24-28 weeks)**:
- Glucose challenge test 1-hour (LOINC 1504-0) -- flag if >=130 or >=140 mg/dL per institutional cutoff
- 3-hour GTT if 1-hour abnormal (LOINC 1518-0, 1530-5, 1547-9, 1507-3)
- Repeat Rh antibody screen if Rh-negative (LOINC 890-4)
- Quad screen or cell-free DNA results if applicable

**Third trimester (35-37 weeks)**:
- GBS culture (LOINC 76685-3)
- Repeat CBC
- Repeat HIV, syphilis in high-risk
- Repeat Chlamydia/Gonorrhea if high-risk

### Step 7: Assess Preeclampsia Risk

Pull BP trend from Step 4. Flag preeclampsia risk if:
- Systolic >=140 mmHg OR diastolic >=90 mmHg on 2 occasions 4+ hours apart after 20 weeks
- Systolic >=160 mmHg OR diastolic >=110 mmHg (severe features)
- Urine protein >=300 mg/24hr or protein/creatinine ratio >=0.3

Check for aspirin prophylaxis if high-risk factors present:
```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&code=1191&status=active"
```
RxNorm 1191 = aspirin. ACOG recommends low-dose aspirin (81mg) starting 12-28 weeks for high-risk patients.

### Step 8: Pull Prenatal Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active"
```

Verify:
- Prenatal vitamin with folic acid (RxNorm concept for prenatal vitamins)
- Iron supplementation if anemic
- Aspirin 81mg if high-risk for preeclampsia
- Rhogam (RhD immune globulin) if Rh-negative at 28 weeks and after delivery
- Progesterone if history of preterm birth
- GBS prophylaxis plan if GBS-positive

### Step 9: Pull Prenatal Procedures/Imaging

```
Tool: fhir_search
resourceType: "Procedure"
queryParams: "patient=[patient-id]&date=ge[pregnancy-onset]&_sort=-date"
```

Expected ultrasounds per ACOG:
- Dating/viability US: 6-9 weeks
- NT scan: 11-14 weeks (if elected)
- Anatomy scan: 18-22 weeks
- Growth scan: 28-32 weeks (if indicated)

### Step 10: Generate Visit Summary

```
PRENATAL VISIT SUMMARY -- [Patient Name] -- GA [X] weeks [X] days
Trimester: [1st/2nd/3rd] | EDD: [date] | Visit #: [N]
==================================================================

GRAVIDITY/PARITY: G[X]P[X]-[X]-[X]-[X] (GTPAL)

VITALS THIS VISIT
  BP: [systolic]/[diastolic] mmHg [trend arrow] [flag if elevated]
  Weight: [current] [unit] | Gain to date: [X] lbs [on track / above / below target]
  Urine protein: [neg/trace/+]

FETAL ASSESSMENT
  FHR: [value] bpm [normal / flag]
  Fundal height: [value] cm [concordant / discordant with GA]
  Fetal movement: [reported / not assessed]

LAB RESULTS DUE THIS VISIT
  [Lab name]: [result if available] | [PENDING if not yet resulted] | [OVERDUE if past window]
  ...

RISK ASSESSMENT
  Preeclampsia risk: [low/moderate/high] -- [basis]
  GDM screening: [completed/pending/positive/negative]
  Preterm labor risk: [low/elevated] -- [basis]

UPCOMING MILESTONES
  [Next GA milestone]: [what is due]
  Next visit: [date] ([interval])

MEDICATIONS VERIFIED
  1. [Medication] [dose] [status]
  ...
```

## Examples

### Example 1: Routine Second Trimester Visit

**User says:** "Pull prenatal visit data for patient OB-4421, she's at 26 weeks"

**Actions:**
1. `fhir_read` Patient/OB-4421 -- 29-year-old female
2. `fhir_search` Condition?patient=OB-4421&code=77386006&clinical-status=active -- confirms pregnancy
3. `fhir_search` Observation?patient=OB-4421&category=vital-signs&date=ge[pregnancy-onset]&_sort=-date&_count=100 -- BP 118/72, weight 158 lbs (prepreg 142 lbs, gain 16 lbs at 26 weeks, on track for normal BMI)
4. `fhir_search` Observation?patient=OB-4421&code=11996-6,11948-7&date=ge[pregnancy-onset]&_sort=-date -- FHR 148 bpm, fundal height 26 cm
5. `fhir_search` Observation?patient=OB-4421&category=laboratory&date=ge[pregnancy-onset]&_sort=-date&_count=200 -- first trimester labs all normal, no glucose challenge yet
6. `fhir_search` MedicationRequest?patient=OB-4421&status=active -- prenatal vitamin active

**Result:**
```
PRENATAL VISIT SUMMARY -- Jane Doe -- GA 26w0d
Trimester: 2nd | EDD: 2026-05-16 | Visit #: 6
==================================================

VITALS THIS VISIT
  BP: 118/72 mmHg (stable, all readings <130/80)
  Weight: 158 lbs | Gain to date: 16 lbs [ON TRACK for normal BMI target 25-35 lbs]

FETAL ASSESSMENT
  FHR: 148 bpm [NORMAL]
  Fundal height: 26 cm [CONCORDANT with 26 weeks GA]

UPCOMING -- DUE AT NEXT VISIT (28 weeks):
  - 1-hour glucose challenge test (GCT) [NOT YET COMPLETED]
  - Repeat Rh antibody screen (patient is Rh-negative)
  - Rhogam administration if antibody screen negative
  - Switch to biweekly visits

MEDICATIONS VERIFIED
  1. Prenatal vitamin daily -- active
```

### Example 2: High-Risk Third Trimester with Complications

**User says:** "Prenatal assessment for patient OB-8830, 34 weeks, history of gestational diabetes"

**Actions:**
1. `fhir_read` Patient/OB-8830 -- 37-year-old female (advanced maternal age)
2. `fhir_search` Condition?patient=OB-8830&clinical-status=active -- pregnancy, gestational diabetes (SNOMED 11687002), chronic hypertension (SNOMED 38341003)
3. `fhir_search` Observation?patient=OB-8830&category=vital-signs&date=ge[pregnancy-onset]&_sort=-date&_count=100 -- BP trending up: 128/82, 134/86, 138/88
4. `fhir_search` Observation?patient=OB-8830&code=11996-6,11948-7&_sort=-date -- FHR 142, fundal height 36 cm (2 cm > expected)
5. `fhir_search` Observation?patient=OB-8830&category=laboratory&date=ge[pregnancy-onset]&_sort=-date&_count=200 -- 3-hour GTT positive, fasting glucose logs available, HbA1c 5.9%
6. `fhir_search` MedicationRequest?patient=OB-8830&status=active -- insulin glargine, prenatal vitamin, aspirin 81mg
7. `fhir_search` Procedure?patient=OB-8830&date=ge[pregnancy-onset]&_sort=-date -- growth scan at 32 weeks showing EFW >90th percentile

**Result:**
```
PRENATAL VISIT SUMMARY -- Maria Garcia -- GA 34w0d
Trimester: 3rd | EDD: 2026-03-14 | Visit #: 12
HIGH RISK: Advanced maternal age, GDM, chronic HTN
==================================================

VITALS THIS VISIT
  BP: 138/88 mmHg [ELEVATED -- trending up from 128/82 at 30 weeks]
  Weight: 192 lbs | Gain to date: 32 lbs [ABOVE TARGET for overweight BMI, target 15-25 lbs]

FETAL ASSESSMENT
  FHR: 142 bpm [NORMAL]
  Fundal height: 36 cm [LARGE FOR DATES -- 2 cm above expected]
  Growth scan (32w): EFW >90th percentile -- macrosomia risk with GDM

RISK FLAGS
  [!] Preeclampsia risk: HIGH -- chronic HTN with rising BP, advanced maternal age
      Action: Check urine protein/creatinine ratio, monitor for symptoms
  [!] GDM control: Fasting glucose logs show intermittent readings >95 mg/dL
      Action: Review insulin dose adjustment
  [!] Macrosomia: EFW >90th percentile
      Action: Discuss delivery timing and mode

UPCOMING (35-37 weeks):
  - GBS culture [DUE at 36 weeks]
  - Weekly NST/BPP (per high-risk protocol)
  - Antenatal corticosteroids if preterm delivery anticipated
  - Repeat growth scan at 36 weeks

MEDICATIONS VERIFIED
  1. Insulin glargine 14 units QHS -- active
  2. Prenatal vitamin daily -- active
  3. Aspirin 81mg daily -- active (preeclampsia prophylaxis)
```

## Troubleshooting

### Pregnancy Condition Not Found with SNOMED 77386006

- Try alternate codes: SNOMED 72892002 (normal pregnancy), ICD-10 Z34.x, O09.x.
- Search broadly: `fhir_search` Condition?patient=[id]&clinical-status=active and manually inspect for pregnancy-related conditions.
- Some EHR systems encode pregnancy as an Observation or Flag rather than a Condition. Search Observation with LOINC 82810-3 (Pregnancy status).

### Gestational Age Cannot Be Determined

- Search for EDD Observation (LOINC 11778-8).
- Search for LMP Observation (LOINC 8665-2 = Last menstrual period start date). Calculate GA: current date - LMP date.
- Check Procedure resources for dating ultrasound with GA assessment.
- If no GA source found, ask the user to provide GA or EDD manually.

### Fetal Observations Return Empty

- Some systems record fetal data under the mother's patient ID; others create a separate fetal patient record. Try searching without code filter and inspect available Observation categories.
- Fundal height may be coded as SNOMED 249016007 instead of LOINC 11948-7.
- FHR may appear as LOINC 55283-6 (fetal heart rate by US) instead of 11996-6.

## Related Skills

- `lab-result-interpreter` -- for detailed interpretation of prenatal lab panels
- `medication-reconciliation` -- for comprehensive prenatal medication verification
- `clinical-summary-generator` -- for full chart summary including pregnancy context
- `vital-signs-trend-analyzer` -- for detailed BP and weight trending
