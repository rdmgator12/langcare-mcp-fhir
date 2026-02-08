---
name: pediatric-growth-assessment
description: |
  Assesses pediatric growth using WHO and CDC growth charts with percentile and z-score calculations.
  Use when user asks to "plot growth", "check growth chart", "growth assessment", "growth percentiles",
  "is the child growing normally", "failure to thrive", "FTT workup", "pediatric BMI", mentions
  "growth velocity", "weight for age", "height for age", or needs growth trend analysis for a child.
  Do NOT use for adult BMI, adult weight management, or prenatal growth (use prenatal-visit-workflow).
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: specialty
---

# Pediatric Growth Assessment

## Overview

Retrieve longitudinal height, weight, and head circumference Observations for a pediatric patient. Calculate age- and sex-specific percentiles and z-scores using WHO standards (0-2 years) and CDC growth charts (2-20 years). Assess growth velocity, BMI-for-age, and weight-for-length (under 2 years). Flag percentile crossing (shift of 2+ major percentile lines), failure to thrive criteria, obesity, and short stature triggers. Generate a structured growth assessment with referral recommendations when indicated.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| Patient | Age, sex, birth date | birthDate, gender |
| Observation | Height, weight, head circumference, BMI | code, valueQuantity, effectiveDateTime |
| Condition | Known growth-related diagnoses | code, clinicalStatus, onsetDateTime |
| MedicationRequest | Growth hormone, nutritional supplements | medicationCodeableConcept, status |

## Instructions

### Step 1: Retrieve Patient Demographics

```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```

Extract `birthDate` (calculate exact age in years, months, days), `gender` (required for sex-specific charts). Confirm age is 0-20 years. If older than 20, this skill does not apply.

### Step 2: Pull All Growth Observations

```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=29463-7,8302-2,9843-4,8287-5,39156-5&_sort=date&_count=200"
```

LOINC codes:
- 29463-7 = Body weight
- 8302-2 = Body height (standing)
- 8306-3 = Body height lying (recumbent length, used <2 years)
- 9843-4 = Head occipitofrontal circumference
- 8287-5 = Head circumference percentile
- 39156-5 = BMI

Also search for recumbent length separately if child is under 2:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8306-3&_sort=date&_count=50"
```

For each Observation, extract:
- `effectiveDateTime` -- calculate age at measurement
- `valueQuantity.value` and `valueQuantity.unit` -- normalize to metric (kg, cm)
- Convert lbs to kg (* 0.453592), inches to cm (* 2.54) if needed

### Step 3: Pull Growth-Related Conditions

```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=36857005,414916001,238131007,237837007,15628003&clinical-status=active"
```

SNOMED codes:
- 36857005 = Failure to thrive
- 414916001 = Obesity
- 238131007 = Overweight
- 237837007 = Short stature
- 15628003 = Growth hormone deficiency

### Step 4: Select Appropriate Growth Chart

Based on age:
- **Birth to 23 months**: WHO growth standards
  - Weight-for-age
  - Length-for-age (recumbent)
  - Weight-for-length
  - Head circumference-for-age
- **2 to 20 years**: CDC growth charts
  - Weight-for-age
  - Stature-for-age (standing height)
  - BMI-for-age
- **Preterm infants**: Use corrected gestational age until 24 months for weight, 40 months for head circumference

See references/growth-chart-data.md for percentile lookup tables and z-score formulas.

### Step 5: Calculate Percentiles and Z-Scores

For each measurement at each time point, calculate:

**Z-score formula (LMS method)**:
- Z = ((X/M)^L - 1) / (L * S) when L != 0
- Z = ln(X/M) / S when L = 0

Where L, M, S are age- and sex-specific parameters from growth chart data tables.

**Percentile from z-score**: Use standard normal CDF approximation.

**Classification by percentile**:
- <1st percentile (z < -2.33): severely underweight/short
- 1st-3rd percentile (z -2.33 to -1.88): underweight/short -- monitor closely
- 3rd-5th percentile (z -1.88 to -1.65): low normal -- monitor
- 5th-85th percentile: normal range
- 85th-95th percentile (z 1.04 to 1.65): overweight (BMI-for-age)
- >=95th percentile (z >= 1.65): obese (BMI-for-age)
- >=99th percentile (z >= 2.33): severe obesity

### Step 6: Assess Growth Velocity

Calculate growth velocity between consecutive measurements:

**Height velocity** (cm/year):
- (height2 - height1) / (age2 - age1 in years)

Compare to age-specific norms (see references/growth-abnormalities.md):
- Infancy (0-1 year): 23-27 cm/year
- Toddler (1-2 years): 10-14 cm/year
- Childhood (2-puberty): 5-7 cm/year
- Pubertal growth spurt: 8-14 cm/year (girls peak ~12y, boys peak ~14y)

**Weight velocity** (grams/day for infants, kg/year for children):
- Newborn: 20-30 g/day (first 3 months)
- Infant 3-6 months: 15-20 g/day
- Infant 6-12 months: 10-15 g/day

Flag growth velocity below 25th percentile for age.

### Step 7: Detect Percentile Crossing

Compare percentiles across time points. Flag if:
- Weight-for-age crosses downward through 2 or more major percentile lines (5th, 10th, 25th, 50th, 75th, 90th, 95th)
- Height-for-age crosses downward through 2 or more major percentile lines
- Weight-for-age drops below 5th percentile
- Head circumference crosses 2+ percentile lines in either direction (macrocephaly or microcephaly)
- BMI-for-age crosses upward through 85th or 95th percentile

### Step 8: Evaluate Failure to Thrive Criteria

FTT is flagged if any of:
- Weight < 5th percentile for age on more than one occasion
- Weight deceleration crossing 2+ major percentile lines
- Weight-for-length < 5th percentile (under 2 years)
- BMI-for-age < 5th percentile (2-20 years)
- Weight velocity < 50% of expected for age

If FTT criteria met, check for:
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```
Look for organic causes: GERD, celiac disease, cystic fibrosis, congenital heart disease, metabolic disorders.

### Step 9: Calculate BMI-for-Age (Children >= 2 Years)

BMI = weight (kg) / height (m)^2

Classify using CDC BMI-for-age percentiles:
- <5th percentile: underweight
- 5th to <85th: healthy weight
- 85th to <95th: overweight
- >=95th: obese
- >=120% of 95th percentile: severe obesity

### Step 10: Format Output

```
PEDIATRIC GROWTH ASSESSMENT -- [Patient Name]
DOB: [date] | Age: [Y]y [M]m | Sex: [M/F]
Growth Chart: [WHO 0-2 / CDC 2-20]
=============================================

CURRENT MEASUREMENTS ([date])
  Weight: [value] kg ([percentile]th %ile, z=[z-score])
  Height: [value] cm ([percentile]th %ile, z=[z-score])
  Head circ: [value] cm ([percentile]th %ile, z=[z-score])  [if <3 years]
  BMI: [value] kg/m2 ([percentile]th %ile, z=[z-score])     [if >=2 years]

GROWTH TRAJECTORY
  [Date1]: Wt [val] ([%ile]), Ht [val] ([%ile]), HC [val] ([%ile])
  [Date2]: Wt [val] ([%ile]), Ht [val] ([%ile]), HC [val] ([%ile])
  ...
  Trend: [stable / upward crossing / downward crossing / decelerating]

GROWTH VELOCITY
  Height velocity: [value] cm/year [normal / below expected]
  Weight velocity: [value] g/day or kg/year [normal / below expected]

FLAGS
  [!] [Description of any concerning finding]
  ...

ASSESSMENT
  [Overall growth status: normal / FTT / overweight / obese / short stature]

RECOMMENDATIONS
  [Specific actions: monitor, dietary evaluation, labs, referral]
```

## Examples

### Example 1: Normal Well-Child Growth Check

**User says:** "Check growth for patient PED-2201, 15-month well child visit"

**Actions:**
1. `fhir_read` Patient/PED-2201 -- 15-month-old male
2. `fhir_search` Observation?patient=PED-2201&code=29463-7,8302-2,8306-3,9843-4&_sort=date&_count=200 -- returns measurements from birth, 2mo, 4mo, 6mo, 9mo, 12mo, 15mo visits
3. `fhir_search` Condition?patient=PED-2201&clinical-status=active -- no growth-related conditions

**Result:**
```
PEDIATRIC GROWTH ASSESSMENT -- Liam Chen
DOB: 2024-11-07 | Age: 1y 3m | Sex: Male
Growth Chart: WHO 0-2
=============================================

CURRENT MEASUREMENTS (2026-02-07)
  Weight: 10.8 kg (55th %ile, z=0.13)
  Length: 79.5 cm (60th %ile, z=0.25)
  Head circ: 47.2 cm (50th %ile, z=0.00)

GROWTH TRAJECTORY
  Birth:  3.4 kg (45th), 50 cm (42nd), HC 34.5 cm (48th)
  6mo:    7.8 kg (50th), 67 cm (48th), HC 43.5 cm (50th)
  12mo:   9.9 kg (50th), 76 cm (55th), HC 46.2 cm (48th)
  15mo:  10.8 kg (55th), 79.5 cm (60th), HC 47.2 cm (50th)
  Trend: STABLE -- tracking along 50th-60th percentile corridor

GROWTH VELOCITY
  Height: 26 cm/year (first year), 14 cm/year (12-15mo) [NORMAL]
  Weight: 24 g/day (first 3mo) [NORMAL]

ASSESSMENT: Normal growth pattern. All parameters tracking appropriately.
```

### Example 2: Failure to Thrive Evaluation

**User says:** "Growth assessment for patient PED-5590, parents concerned about weight gain"

**Actions:**
1. `fhir_read` Patient/PED-5590 -- 10-month-old female
2. `fhir_search` Observation?patient=PED-5590&code=29463-7,8302-2,8306-3,9843-4&_sort=date&_count=200 -- shows weight percentile dropping: birth 50th, 4mo 35th, 6mo 15th, 9mo 5th, 10mo 3rd
3. `fhir_search` Condition?patient=PED-5590&clinical-status=active -- GERD documented

**Result:**
```
PEDIATRIC GROWTH ASSESSMENT -- Sophia Martinez
DOB: 2025-04-07 | Age: 0y 10m | Sex: Female
Growth Chart: WHO 0-2
=============================================

CURRENT MEASUREMENTS (2026-02-07)
  Weight: 7.1 kg (3rd %ile, z=-1.88)
  Length: 70.5 cm (25th %ile, z=-0.67)
  Head circ: 44.8 cm (45th %ile, z=-0.13)
  Weight-for-length: 3rd %ile (z=-1.88)

GROWTH TRAJECTORY
  Birth:  3.3 kg (50th), 49 cm (45th), HC 34 cm (50th)
  4mo:    5.8 kg (35th), 62 cm (35th), HC 40 cm (42nd)
  6mo:    6.4 kg (15th), 65 cm (28th), HC 42 cm (45th)
  9mo:    6.8 kg (5th),  68 cm (22nd), HC 44 cm (44th)
  10mo:   7.1 kg (3rd),  70.5 cm (25th), HC 44.8 cm (45th)
  Trend: DOWNWARD CROSSING -- weight crossed from 50th to 3rd percentile

FLAGS
  [!] FAILURE TO THRIVE CRITERIA MET
      - Weight crossed downward through 3 major percentile lines (50th->35th->15th->3rd)
      - Weight-for-length at 3rd percentile
      - Weight velocity: 8 g/day (months 6-10) -- below expected 10-15 g/day
  [!] Known GERD -- potential organic contributor to poor weight gain
  [!] Length sparing: height at 25th percentile while weight at 3rd -- suggests nutritional FTT

RECOMMENDATIONS
  - Dietary assessment: caloric intake evaluation, feeding observation
  - Laboratory workup: CBC, CMP, celiac panel (tTG-IgA), thyroid, urinalysis
  - Consider increasing caloric density of feeds
  - GERD management optimization -- assess current treatment adequacy
  - Follow-up weight check in 2 weeks
  - Refer to pediatric GI if no improvement in 4 weeks
```

## Troubleshooting

### Growth Observations Use Non-Standard LOINC Codes

- Some EHR systems use LOINC 3141-9 (body weight measured) instead of 29463-7. Try both.
- Recumbent length may be coded as LOINC 8306-3 or 89269-5. Search without code filter and inspect available Observation codes.
- Head circumference may use LOINC 8287-5 (percentile) instead of 9843-4 (actual measurement).

### Measurements in Non-Metric Units

- Convert on the fly: lbs to kg (* 0.453592), inches to cm (* 2.54), oz to g (* 28.3495).
- Check `valueQuantity.unit` for `[lb_av]`, `[in_i]`, `lb`, `in`, `oz`.
- WHO/CDC chart data uses metric. All calculations must be in kg and cm.

### Insufficient Data Points for Velocity Calculation

- Need at least 2 measurements separated by >=1 month for meaningful velocity.
- If only one measurement exists, calculate percentile and z-score but note velocity cannot be assessed.
- Recommend establishing baseline with follow-up measurement.

## Related Skills

- `lab-result-interpreter` -- for FTT laboratory workup interpretation
- `immunization-status-checker` -- often combined with well-child growth checks
- `clinical-summary-generator` -- for comprehensive pediatric summary
- `preventive-care-compliance-report` -- for well-child visit schedule adherence
