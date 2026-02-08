---
name: renal-function-dashboard
description: |
  Pulls renal function labs (BUN, creatinine, eGFR, cystatin C, urine protein, electrolytes) and stages CKD per KDIGO 2024 guidelines with trajectory analysis.
  Use when user asks to "check kidney function", "renal function", "CKD staging", "eGFR trend", "kidney labs",
  "nephrology dashboard", "renal panel", or needs assessment of chronic kidney disease status and progression.
  Do NOT use for acute kidney injury diagnosis alone, urinalysis interpretation, or dialysis management.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: lab-diagnostics
---

# Renal Function Dashboard

## Overview

Pull BUN, creatinine, eGFR, cystatin C, urine protein, urine albumin-to-creatinine ratio, and electrolytes (potassium, calcium, phosphorus, magnesium, bicarbonate). Stage CKD per KDIGO 2024 guidelines using GFR categories (G1-G5) and albuminuria categories (A1-A3). Track eGFR trajectory over time, calculate annual rate of decline, and flag rapid progression. Detect AKI superimposed on CKD. Assess electrolyte management adequacy by CKD stage. Flag medication dose adjustments needed and nephrotoxic drugs.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Observation | Renal labs, electrolytes, urinalysis | code, valueQuantity, effectiveDateTime, interpretation |
| Condition | CKD diagnosis, comorbidities | code, clinicalStatus, onsetDateTime |
| MedicationStatement | Medications needing renal adjustment | medicationCodeableConcept, status, dosage |
| Patient | Demographics for eGFR calculation | birthDate, gender, extension (race) |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract: age, gender, race (from US Core race extension). These are inputs to the CKD-EPI 2021 equation (which removed the race variable) or CKD-EPI cystatin C equation if available.

### Step 2: Pull Core Renal Function Labs

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2160-0,3094-0,33914-3,82810-3,76484-0,48642-3,69405-9,14959-1&_sort=-date&_count=50"
```

LOINC codes:
- 2160-0 = Serum Creatinine
- 3094-0 = BUN (Blood Urea Nitrogen)
- 33914-3 = eGFR by CKD-EPI (creatinine-based)
- 82810-3 = eGFR by CKD-EPI 2021 (without race)
- 76484-0 = eGFR by CKD-EPI cystatin C
- 48642-3 = Urine albumin/creatinine ratio (UACR)
- 69405-9 = Urine albumin/creatinine ratio (alternate)
- 14959-1 = Urine microalbumin

If eGFR not directly available, calculate from creatinine using CKD-EPI 2021:
- eGFR = 142 * min(Scr/kappa, 1)^alpha * max(Scr/kappa, 1)^(-1.200) * 0.9938^age * (1.012 if female)
- kappa = 0.7 (female), 0.9 (male); alpha = -0.241 (female), -0.302 (male)

### Step 3: Pull Electrolytes

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2823-3,17861-6,2777-1,19123-9,1963-8,1920-8&_sort=-date&_count=30"
```

LOINC codes:
- 2823-3 = Potassium
- 17861-6 = Calcium (total)
- 2777-1 = Phosphorus
- 19123-9 = Magnesium
- 1963-8 = Bicarbonate (CO2)
- 1920-8 = AST (for hepatorenal context)

Also pull:
- PTH (2731-8) -- secondary hyperparathyroidism in CKD
- Vitamin D 25-OH (1989-3) -- deficiency common in CKD
- Hemoglobin (718-7) -- anemia of CKD

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=2731-8,1989-3,718-7&_sort=-date&_count=10"
```

### Step 4: Stage CKD per KDIGO 2024

**GFR Categories:**
- **G1**: eGFR >= 90 (normal or high)
- **G2**: eGFR 60-89 (mildly decreased)
- **G3a**: eGFR 45-59 (mildly to moderately decreased)
- **G3b**: eGFR 30-44 (moderately to severely decreased)
- **G4**: eGFR 15-29 (severely decreased)
- **G5**: eGFR < 15 (kidney failure)

**Albuminuria Categories:**
- **A1**: UACR < 30 mg/g (normal to mildly increased)
- **A2**: UACR 30-300 mg/g (moderately increased)
- **A3**: UACR > 300 mg/g (severely increased)

**Prognosis Risk (KDIGO Heat Map):**
- Low risk (green): G1A1, G2A1
- Moderately increased risk (yellow): G1A2, G2A2, G3aA1
- High risk (orange): G1A3, G2A3, G3aA2, G3bA1
- Very high risk (red): G3aA3, G3bA2, G3bA3, G4A1, G4A2, G4A3, G5A1, G5A2, G5A3

CKD diagnosis requires GFR < 60 OR albuminuria (UACR >= 30) OR other markers of kidney damage persisting for > 3 months.

### Step 5: Track eGFR Trajectory

Pull historical eGFR values (or creatinine to calculate):

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=33914-3,82810-3&_sort=date&_count=20"
```

Calculate:
- **Annual rate of decline**: linear regression slope of eGFR over time, expressed as mL/min/1.73m2 per year
- **Normal age-related decline**: approximately 1 mL/min/1.73m2 per year after age 40
- **Rapid decline**: > 5 mL/min/1.73m2 per year (KDIGO definition)
- **Projected time to eGFR < 15**: extrapolate using current trajectory (triggers dialysis planning)

### Step 6: Detect AKI Superimposed on CKD

Compare most recent creatinine to baseline (lowest value in past 90 days or known stable value):

KDIGO AKI criteria:
- **Stage 1**: Creatinine increase >= 0.3 mg/dL within 48 hours OR 1.5-1.9x baseline within 7 days
- **Stage 2**: Creatinine 2.0-2.9x baseline
- **Stage 3**: Creatinine >= 3.0x baseline OR creatinine >= 4.0 mg/dL OR initiation of renal replacement therapy

If AKI detected, flag as `[AKI ALERT]` with stage and recommend: hold nephrotoxic medications, IV fluid assessment, urine output monitoring, nephrology consultation if Stage 2+.

### Step 7: Assess Electrolyte Management by CKD Stage

**Potassium:**
- G1-G3a: target 3.5-5.0 mEq/L (standard)
- G3b-G5: target 3.5-5.5 mEq/L; monitor more frequently
- Flag: > 5.5 (hold ACEi/ARB/spironolactone, dietary counseling); > 6.0 (urgent)

**Phosphorus:**
- G1-G3a: target 2.5-4.5 mg/dL
- G3b-G5: target 2.5-4.5 mg/dL (maintain within normal range; KDIGO de-emphasizes strict targets)
- Flag: > 4.5 in G3b+: consider phosphate binder, dietary restriction

**Calcium:**
- Target: 8.5-10.5 mg/dL across all stages
- Correct for albumin: Corrected Ca = total Ca + 0.8 * (4.0 - albumin)
- Flag: < 8.5 with elevated PTH: secondary hyperparathyroidism

**Bicarbonate:**
- Target: >= 22 mEq/L
- Flag: < 22 in CKD G3-G5: start oral sodium bicarbonate supplementation (slows CKD progression per KDIGO)

**PTH:**
- G3a-G3b: target within normal range
- G4-G5: target 2-9x upper normal limit (KDIGO)
- Flag: elevated PTH + low vitamin D: correct vitamin D first

### Step 8: Check Medications for Renal Dosing

```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```

Flag medications needing dose adjustment or contraindicated at current eGFR:
- **Metformin**: reduce dose at eGFR 30-45, contraindicated < 30
- **SGLT2 inhibitors**: reduced glycemic efficacy < 30 but renal protection continues (per KDIGO/ADA); stop if starting dialysis
- **NSAIDs**: avoid in eGFR < 30 (relative contraindication in all CKD)
- **Gabapentin/pregabalin**: dose reduce by eGFR
- **Allopurinol**: start low dose in CKD
- **ACEi/ARB**: continue for proteinuria benefit but monitor K+ and creatinine (up to 30% rise in Cr acceptable after initiation)
- **Contrast dye**: flag if recent or upcoming (hold metformin, hydration protocol)
- **Aminoglycosides**: avoid or monitor levels closely
- **Lithium**: narrow therapeutic index, accumulates in renal impairment

See references/renal-medication-dosing.md for complete dosing table.

### Step 9: Assess Referral Criteria

Per KDIGO, refer to nephrology if:
- eGFR < 30 (G4-G5)
- UACR > 300 mg/g (A3) persistent
- Rapid decline (> 5 mL/min/1.73m2 per year)
- AKI or unexplained acute decline
- Refractory hypertension (>= 4 agents)
- Persistent electrolyte abnormalities (hyperkalemia, metabolic acidosis)
- Hereditary kidney disease suspected

Dialysis planning discussion recommended when eGFR < 20 (projected to reach < 15 within 12 months).

### Step 10: Format Output

```
RENAL FUNCTION DASHBOARD -- [Patient Name] ([Age][Sex])
========================================================
CKD STAGING (KDIGO 2024)
========================
GFR Category:         G[stage] (eGFR [value] mL/min/1.73m2)
Albuminuria Category: A[stage] (UACR [value] mg/g)
Overall Risk:         [Low / Moderate / High / Very High]
AKI Status:           [None detected / AKI Stage X]

eGFR TRAJECTORY
===============
[date1]: [eGFR1]
[date2]: [eGFR2]
[date3]: [eGFR3]
[date4]: [eGFR4]
Annual rate of decline: [value] mL/min/1.73m2/year [STABLE / RAPID DECLINE]
Projected eGFR < 15:   [estimated date or "Not projected within 5 years"]

CORE RENAL LABS
===============
Creatinine:   [value] mg/dL  ([date])  Baseline: [value]
BUN:          [value] mg/dL  ([date])
BUN/Cr ratio: [calculated]
eGFR:         [value]        ([date])
Cystatin C:   [value] mg/L   ([date])  [if available]
UACR:         [value] mg/g   ([date])

ELECTROLYTES & MINERAL BONE
============================
Potassium:    [value] mEq/L  [target for CKD stage]  [NORMAL / FLAG]
Calcium:      [value] mg/dL  (corrected: [value])     [NORMAL / FLAG]
Phosphorus:   [value] mg/dL  [target for CKD stage]  [NORMAL / FLAG]
Magnesium:    [value] mg/dL                           [NORMAL / FLAG]
Bicarbonate:  [value] mEq/L  (target >= 22)          [NORMAL / LOW]
PTH:          [value] pg/mL  [target for CKD stage]  [NORMAL / ELEVATED]
Vitamin D:    [value] ng/mL                           [NORMAL / DEFICIENT]
Hemoglobin:   [value] g/dL                            [NORMAL / ANEMIA]

MEDICATION REVIEW
=================
[drug]: [current dose] -- [appropriate / ADJUST to X / DISCONTINUE]
[drug]: [current dose] -- [appropriate / ADJUST to X / DISCONTINUE]

NEPHROTOXIC MEDICATIONS: [list or "None identified"]

RECOMMENDATIONS
===============
1. [CKD management recommendation]
2. [Electrolyte management]
3. [Medication adjustment]
4. [Referral if indicated]
5. [Next monitoring interval]
```

## Examples

### Example 1: Stable CKD Stage 3a

**User says:** "Check kidney function for patient 9920"

**Actions:**
1. `fhir_read` Patient/9920 -- 64-year-old female
2. `fhir_search` Observation?patient=9920&code=2160-0,3094-0,33914-3,48642-3&_sort=-date&_count=50 -- Cr 1.3, BUN 24, eGFR 52, UACR 45 mg/g
3. `fhir_search` Observation?patient=9920&code=33914-3,82810-3&_sort=date&_count=20 -- eGFR over 2 years: 58, 56, 54, 53, 52
4. `fhir_search` Observation?patient=9920&code=2823-3,17861-6,2777-1,19123-9,1963-8&_sort=-date&_count=30 -- K 4.6, Ca 9.2, Phos 3.8, Mg 2.0, Bicarb 23
5. `fhir_search` Observation?patient=9920&code=2731-8,1989-3,718-7&_sort=-date&_count=10 -- PTH 68, Vit D 28, Hgb 12.1
6. `fhir_search` MedicationStatement?patient=9920&status=active -- lisinopril, metformin 1000mg BID, atorvastatin
7. `fhir_search` Condition?patient=9920&clinical-status=active -- Type 2 DM, HTN, CKD stage 3

**Result:** CKD G3a/A2, high risk (orange on KDIGO heat map). eGFR declining at 3 mL/min/year (not yet rapid but above age-related). UACR 45 (moderately increased). Electrolytes well-managed. PTH mildly elevated but Vitamin D borderline -- supplement to > 30 ng/mL. Metformin dose appropriate at eGFR 52 (reduce if eGFR drops below 45). ACEi appropriate for albuminuria. Consider adding SGLT2 inhibitor for renal protection. Next monitoring: 3-6 months.

### Example 2: Rapid Decline With AKI

**User says:** "Renal panel for patient r-5501, creatinine has been going up"

**Actions:**
1. `fhir_read` Patient/r-5501 -- 73-year-old male
2. `fhir_search` Observation?patient=r-5501&code=2160-0&_sort=-date&_count=10 -- Cr values: 3.8 (today), 2.1 (1 week ago), 1.8 (1 month ago), 1.6 (3 months ago), 1.5 (6 months ago)
3. `fhir_search` Observation?patient=r-5501&code=33914-3&_sort=date&_count=20 -- eGFR: 42, 38, 34, 28, 14 (calculated from Cr)
4. `fhir_search` Observation?patient=r-5501&code=2823-3&_sort=-date&_count=5 -- K 5.9 mEq/L
5. `fhir_search` MedicationStatement?patient=r-5501&status=active -- ibuprofen 600mg TID, lisinopril, metformin
6. `fhir_search` Condition?patient=r-5501&clinical-status=active -- CKD stage 3b, DM2, OA

**Result:**
```
[AKI ALERT] Stage 2 superimposed on CKD
Creatinine 3.8 mg/dL = 2.1x baseline (1.8), rise of 1.7 mg/dL in 1 week
eGFR dropped from 34 to 14 -- now G5 territory
Baseline trajectory was already rapid: 8 mL/min/year decline over 6 months

URGENT ACTIONS:
1. STOP ibuprofen immediately (nephrotoxic, primary suspect for acute decline)
2. HOLD metformin (eGFR < 30, lactic acidosis risk)
3. HOLD lisinopril (hyperkalemia K 5.9, AKI)
4. Potassium 5.9: ECG stat, consider calcium gluconate, kayexalate
5. Nephrology consult STAT
6. IV fluid assessment, urine output monitoring
7. Renal ultrasound to rule out obstruction
```

## Troubleshooting

### eGFR Not Available as a Separate Observation

- Calculate from serum creatinine using CKD-EPI 2021 equation if `Observation` for LOINC 33914-3 or 82810-3 is absent. Requires: creatinine value, age, sex.
- Some systems report eGFR as a component of a renal panel Observation. Check `component` array on BMP/CMP panel Observations.

### UACR Reported as Separate Urine Albumin and Urine Creatinine

- Pull urine albumin (14959-1) and urine creatinine (2161-8) separately. Calculate UACR = urine albumin (mg/L) / urine creatinine (g/L). Ensure unit conversion: if urine creatinine is in mg/dL, convert to g/L by dividing by 100.
- If only urine protein (2888-6) is available (not albumin-specific), use protein/creatinine ratio as a surrogate. UPCR > 0.5 g/g approximates UACR > 300 mg/g.

### Historical eGFR Values Use Different Equations

- Older labs may use MDRD instead of CKD-EPI, or may include the race coefficient. For trend analysis, recalculate all values using CKD-EPI 2021 (without race) from stored creatinine values for consistency.
- Note in the trajectory: "eGFR values recalculated using CKD-EPI 2021 for consistency."

## Related Skills

- `lab-result-interpreter` -- for comprehensive interpretation of all lab results beyond renal
- `diabetes-panel-review` -- for diabetic nephropathy in the context of full diabetes management
- `critical-value-alert-generator` -- for critical electrolyte values in renal failure
- `preoperative-lab-checklist` -- for renal function assessment before surgery
