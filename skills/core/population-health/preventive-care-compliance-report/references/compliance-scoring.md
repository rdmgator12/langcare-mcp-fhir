# Compliance Scoring Methodology

Methodology for calculating individual and practice-level preventive care compliance scores, including benchmark rates, scoring algorithms, and reporting formats.

---

## Individual Patient Compliance Score

### Scoring Algorithm

For each patient, evaluate all applicable guideline items. An item is "applicable" if the patient meets the age, sex, and risk criteria defined in the USPSTF/ACS/CDC recommendation.

**Item Status Classification:**

| Status | Definition | Score |
|--------|-----------|-------|
| Up to date | Evidence of screening/intervention within recommended interval | 1.0 |
| Due soon | Within 3 months of due date (grace period) | 0.75 |
| Overdue | Past recommended interval | 0.0 |
| Not applicable | Patient does not meet guideline criteria (age, sex, risk) | Excluded |
| Data unavailable | No FHIR data to evaluate; cannot determine status | 0.0 (flagged) |

**Patient Compliance Rate:**
```
Compliance Rate = (Sum of item scores) / (Count of applicable items) * 100
```

**Example:**
- 10 applicable items
- 7 up to date (7.0), 1 due soon (0.75), 2 overdue (0.0)
- Score = (7.0 + 0.75 + 0.0) / 10 * 100 = 77.5%

### Weighting (Optional)

If the practice prioritizes certain screenings, apply weights:

| Category | Weight | Rationale |
|----------|--------|-----------|
| Cancer screenings (mammography, colonoscopy, cervical, lung) | 1.5x | Highest morbidity/mortality impact |
| Cardiovascular prevention (BP, lipids, statin) | 1.25x | Leading cause of death |
| Metabolic screening (diabetes) | 1.0x | Standard |
| Behavioral health (depression, tobacco, alcohol) | 1.0x | Standard |
| Immunizations | 1.0x | Standard |
| Infectious disease screening (HIV, HCV, HBV) | 0.75x | One-time screenings, lower recurring impact |

Weighted score:
```
Weighted Rate = Sum(item_score * weight) / Sum(weights for applicable items) * 100
```

---

## Practice-Level Compliance Rates

### Calculation per Measure

For each screening/prevention item:
```
Eligible Population = Patients meeting age/sex/risk criteria for the item
Compliant Count = Patients with evidence of item within recommended interval
Compliance Rate = Compliant / Eligible * 100
```

### Aggregate Practice Score

```
Practice Compliance Score = Mean of all individual measure rates
```

Or weighted:
```
Practice Weighted Score = Sum(measure_rate * measure_weight) / Sum(weights)
```

### Confidence Intervals

For small practices, include 95% confidence interval:
```
CI = Rate +/- 1.96 * sqrt(Rate * (1 - Rate) / N)
```
Where N = eligible population for that measure.

Flag measures where N < 30 as "insufficient sample size for reliable rate."

---

## Benchmark Rates

### National Benchmarks (HEDIS / BRFSS 2023)

| Measure | National Average | 75th Percentile | 90th Percentile | Healthy People 2030 |
|---------|-----------------|-----------------|-----------------|---------------------|
| Mammography (50-74, biennial) | 76.8% | 83.0% | 86.0% | 80.5% |
| Cervical screening (21-65) | 83.5% | 87.0% | 90.0% | 84.3% |
| Colorectal screening (45-75) | 72.1% | 80.0% | 84.0% | 74.4% |
| Lung cancer CT (eligible) | 15.4% | -- | -- | -- |
| BP screening (annual) | 89.0% | 92.0% | 95.0% | -- |
| Diabetes screening (35-70, overweight) | 74.0% | 80.0% | 85.0% | -- |
| Lipid screening (per interval) | 86.5% | 90.0% | 93.0% | -- |
| Depression screening (PHQ) | 78.0% | 83.0% | 87.0% | -- |
| Tobacco screening | 90.0% | 93.0% | 96.0% | -- |
| Tobacco cessation (smokers) | 55.0% | 62.0% | 70.0% | -- |
| HIV screening (ever, 15-65) | 62.0% | -- | -- | -- |
| HCV screening (ever, 18-79) | 60.5% | -- | -- | -- |
| Flu vaccination (>= 18) | 48.2% | 55.0% | 65.0% | 70.0% |
| A1c control (< 8%, diabetics) | 81.0% | 85.0% | 89.0% | -- |
| BP control (< 140/90, HTN) | 74.0% | 79.0% | 84.0% | -- |

### Performance Tiers

| Tier | Percentile Range | Label | Color Code |
|------|-----------------|-------|------------|
| Excellent | >= 90th | Top performer | Green |
| Above average | 75th - 89th | Strong | Light green |
| Average | 50th - 74th | Meeting expectations | Yellow |
| Below average | 25th - 49th | Needs improvement | Orange |
| Poor | < 25th | Significant gap | Red |

---

## Reporting Formats

### Individual Patient Scorecard

```
PREVENTIVE CARE COMPLIANCE SCORECARD
======================================
Patient: [Name] | MRN: [MRN] | Age: [Age] | Sex: [Sex]
Report Date: [Date]
Overall Compliance: [XX.X%] ([N]/[M] items up to date)

CATEGORY BREAKDOWN
------------------
Cancer Screenings:        [X/Y] ([%])
Cardiovascular:           [X/Y] ([%])
Metabolic:                [X/Y] ([%])
Behavioral Health:        [X/Y] ([%])
Immunizations:            [X/Y] ([%])
Infectious Disease:       [X/Y] ([%])

ITEM DETAIL
-----------
[STATUS] [Item Name] - Last: [date] - Next due: [date] - Interval: [X years]
...

PRIORITY ACTIONS
================
1. [Most overdue item] - [days/months overdue]
2. [Second most overdue] - [days/months overdue]
...
```

### Practice Dashboard

```
PRACTICE PREVENTIVE CARE DASHBOARD
=====================================
Total Active Patients: [N]
Report Period: [Start] to [End]
Overall Practice Compliance: [XX.X%]

MEASURE-LEVEL RATES
--------------------
Measure                    | Eligible | Compliant | Rate  | Benchmark | Status
---------------------------|----------|-----------|-------|-----------|-------
Mammography                |      [N] |       [N] | [XX%] | [XX%]     | [Tier]
Colorectal Screening       |      [N] |       [N] | [XX%] | [XX%]     | [Tier]
Cervical Screening         |      [N] |       [N] | [XX%] | [XX%]     | [Tier]
...

TREND (Quarterly)
-----------------
Quarter  | Overall Rate | Change
---------|-------------|-------
Q1       | [XX%]       | --
Q2       | [XX%]       | [+/-X.X%]
Q3       | [XX%]       | [+/-X.X%]
Q4       | [XX%]       | [+/-X.X%]

TOP IMPROVEMENT OPPORTUNITIES
------------------------------
1. [Measure with largest gap below benchmark] - [N] patients need outreach
2. [Second largest gap] - [N] patients need outreach
...
```

### Provider-Level Comparison (if multi-provider practice)

```
PROVIDER COMPARISON
====================
Provider       | Panel Size | Overall Compliance | Top Gap Measure | Gap Count
---------------|------------|-------------------|-----------------|----------
Dr. Smith      |       [N]  | [XX%]             | [Measure]       | [N]
Dr. Jones      |       [N]  | [XX%]             | [Measure]       | [N]
NP Williams    |       [N]  | [XX%]             | [Measure]       | [N]
...
```

---

## Improvement Tracking

### Month-Over-Month Tracking

Track these metrics monthly:
1. **Gap closure rate**: Number of gaps closed / Total gaps at start of month
2. **New gaps created**: Patients becoming overdue during the month
3. **Net gap change**: Gaps closed - New gaps
4. **Outreach success rate**: Patients who completed screening after outreach / Total outreached

### Targets for Improvement

| Metric | Minimum Target | Stretch Target |
|--------|---------------|----------------|
| Monthly gap closure rate | >= 5% | >= 10% |
| Outreach success rate | >= 30% | >= 50% |
| Practice compliance (overall) | >= 75th percentile | >= 90th percentile |
| No measure below 50th percentile | -- | All measures >= 50th |

### Intervention Effectiveness

Track which outreach methods produce the highest screening completion rates:

| Method | Typical Completion Rate | Cost per Completion |
|--------|------------------------|-------------------|
| Phone call from care team | 25-35% | $15-25 |
| Patient portal message | 15-25% | $2-5 |
| Mailed letter | 10-20% | $3-7 |
| Mailed FIT kit (colorectal) | 20-30% | $10-15 |
| Text message reminder | 12-18% | $1-3 |
| In-visit nursing prompt (huddle) | 40-60% | $0 (built into workflow) |
| Standing order (automatic) | 60-80% | $0 (built into workflow) |

**Most effective combination:** Pre-visit huddle (identify due items) + in-visit nursing prompt + standing orders for labs/immunizations.

---

## Data Quality Considerations

### Common Data Issues

| Issue | Impact | Mitigation |
|-------|--------|-----------|
| External screenings not in FHIR | Undercount compliant patients | Query HIE; accept patient self-report with documentation |
| Procedure codes vary by system | Miss screenings coded differently | Search multiple code systems (SNOMED, CPT, LOINC) |
| Historical data incomplete | Overcount gaps for established patients | Apply lookback only from patient's first visit date |
| No FamilyMemberHistory | Cannot determine high-risk status | Default to average-risk screening; flag for history collection |
| Condition vs Observation for risk factors | Miss smoking, obesity if not in Conditions | Check both Condition and Observation for smoking status, BMI |

### Minimum Data Completeness Thresholds

| Requirement | Threshold | Action if Not Met |
|-------------|-----------|-------------------|
| Active patient panel identified | 100% | Cannot generate report |
| Patient demographics (age, sex) | >= 99% | Exclude patients without demographics |
| At least 1 screening data source | >= 80% | Flag data quality issue; rates may be understated |
| Family history documented | >= 50% | Note: high-risk screening may be underidentified |
| Smoking status documented | >= 85% | Lung cancer screening eligibility uncertain for undocumented |
