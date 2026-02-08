# CMS Quality Reporting Reference

## Overview

Quality reporting for healthcare organizations encompasses multiple programs: HEDIS (managed care), MIPS (physician practices), Medicare Star Ratings (Medicare Advantage), and state-specific programs. This reference covers reporting requirements, benchmark data, and improvement strategies relevant to FHIR-based quality measure calculation.

---

## CMS Quality Programs

### Medicare Star Ratings (Medicare Advantage)

**Program:** Annual rating of Medicare Advantage plans on a 1-5 star scale.

**Clinical Measures Included:**

| Domain | Weight | Example Measures |
|--------|--------|-----------------|
| Staying Healthy: Screenings, Tests, Vaccines | 2x | BCS, COL, Flu vaccination, BMI screening |
| Managing Chronic Conditions | 2x | HBD (A1c), CBP, SPD, OMW (osteoporosis) |
| Member Experience (CAHPS) | 1.5x | Getting needed care, timely appointments |
| Complaints and Appeals | 1x | Complaint rate, appeals upheld |
| Health Plan Customer Service | 1x | Call center, TTY availability |

**Improvement Bonus:** Plans that improve by >= 3 percentage points year-over-year on a measure can receive an improvement bonus toward star calculation.

### MIPS (Merit-based Incentive Payment System)

**Program:** Payment adjustment for clinicians based on quality performance.

**Quality Performance Category (30% of MIPS score):**
- Report on 6 quality measures (at least 1 outcome measure)
- Measured against benchmarks derived from historic performance
- Decile-based scoring: 3-10 points per measure

**Relevant FHIR-Calculable MIPS Measures:**

| MIPS ID | Measure | HEDIS Equivalent |
|---------|---------|------------------|
| 001 | Diabetes: A1c Poor Control (> 9%) | HBD (inverse) |
| 236 | Controlling High Blood Pressure | CBP |
| 226 | Tobacco Screening and Cessation | TCC |
| 134 | Depression Screening | DSF |
| 438 | Statin Therapy (ASCVD) | Superset of SPD |
| 112 | Breast Cancer Screening | BCS |
| 113 | Colorectal Cancer Screening | COL |

---

## Benchmark Data

### HEDIS Benchmarks (Measurement Year 2023)

**Commercial HMO:**

| Measure | 10th | 25th | 50th | 75th | 90th | Mean |
|---------|------|------|------|------|------|------|
| HBD A1c < 8% | 68.0 | 76.0 | 81.0 | 85.0 | 89.0 | 80.0 |
| HBD A1c > 9% (lower=better) | 16.0 | 10.0 | 7.0 | 5.0 | 3.0 | 8.5 |
| HBD Eye Exam | 45.0 | 52.0 | 58.0 | 64.0 | 70.0 | 57.0 |
| CBP < 140/90 | 58.0 | 68.0 | 74.0 | 79.0 | 84.0 | 73.0 |
| BCS Mammography | 70.0 | 76.0 | 80.0 | 83.0 | 86.0 | 79.0 |
| COL Screening | 64.0 | 72.0 | 76.0 | 80.0 | 84.0 | 75.0 |
| DSF Depression Screen | 62.0 | 73.0 | 78.0 | 83.0 | 87.0 | 76.0 |
| SPD Statin Therapy | 70.0 | 78.0 | 82.0 | 86.0 | 90.0 | 81.0 |

**Medicare HMO:**

| Measure | 10th | 25th | 50th | 75th | 90th | Mean |
|---------|------|------|------|------|------|------|
| HBD A1c < 8% | 72.0 | 79.0 | 84.0 | 88.0 | 91.0 | 83.0 |
| CBP < 140/90 | 60.0 | 68.0 | 74.0 | 80.0 | 85.0 | 73.0 |
| BCS Mammography | 62.0 | 70.0 | 76.0 | 80.0 | 84.0 | 75.0 |
| COL Screening | 60.0 | 68.0 | 74.0 | 79.0 | 83.0 | 73.0 |

### National Averages (CDC/BRFSS 2023)

Population-level rates for context:

| Screening | National Average | Healthy People 2030 Target |
|-----------|-----------------|---------------------------|
| Mammography (women 50-74, past 2 yrs) | 76.8% | 80.5% |
| Colorectal screening (50-75) | 72.1% | 74.4% |
| Cervical cancer screening (21-65) | 83.5% | 84.3% |
| A1c testing (diabetics, past year) | 91.0% | 92.0% |
| BP check (past year) | 89.0% | -- |
| Cholesterol check (ever, adults) | 86.5% | -- |
| Flu vaccination (>= 18) | 48.2% | 70.0% |

---

## Measure Calculation Timing

### Measurement Year

Standard HEDIS measurement year: January 1 - December 31.

| Event | Timeline |
|-------|----------|
| Measurement Year Start | January 1 |
| Measurement Year End | December 31 |
| Data Collection | January - March (following year) |
| Submission Deadline | June 15 (following year) |
| Results Released | October (following year) |

### Lookback Periods by Measure

| Measure | Lookback from End of MY |
|---------|-------------------------|
| A1c (HBD) | 1 year (during MY) |
| BP (CBP) | 1 year (during MY) |
| Mammography (BCS) | 27 months |
| Colonoscopy (COL) | 10 years |
| FIT/FOBT (COL) | 1 year |
| FIT-DNA (COL) | 3 years |
| Flex Sig (COL) | 5 years |
| CT Colonography (COL) | 5 years |
| Depression Screen (DSF) | 1 year (during MY) |
| Statin (SPD) | 1 year (during MY) |

---

## Gap Closure Strategies

### High-Impact Interventions by Measure

**A1c Control (HBD):**
- Pre-visit planning: identify patients with A1c > 8% before scheduled visit
- Standing lab orders: auto-generate A1c order 1 week before appointment
- Pharmacist-led medication titration protocols
- Care management enrollment for A1c > 9%
- Endocrinology e-consult for complex insulin management

**Blood Pressure (CBP):**
- Home BP monitoring programs with remote data transmission
- Pharmacist-led titration clinics
- Nurse visit for BP-only recheck within 4 weeks of medication change
- Automated BP alert for readings > 180/120

**Breast Cancer Screening (BCS):**
- Bulk outreach: annual letter/portal message to due patients in January
- Mammography van partnerships
- Order mammogram at current visit for patients due within 6 months
- Reminder call 1 week before scheduled mammogram

**Colorectal Cancer Screening (COL):**
- FIT kit mailed to home (highest impact for non-adherent patients)
- Colonoscopy scheduling navigation support
- Pre-signed colonoscopy referral at wellness visit
- FIT-DNA (Cologuard) as alternative for colonoscopy-averse patients

**Depression Screening (DSF):**
- PHQ-2 on intake form for all visits
- Automatic escalation to PHQ-9 if PHQ-2 >= 3
- Integrated behavioral health in primary care
- Warm handoff protocol for positive screens

### ROI of Gap Closure

| Measure | Estimated Revenue per Gap Closed | Method |
|---------|----------------------------------|--------|
| HBD (1 point star improvement) | $50-100K per plan | Higher star bonus payments |
| BCS (per patient) | $15-30 per outreach | Quality bonus recapture |
| COL (per patient) | $15-30 per outreach | Quality bonus recapture |

---

## Reporting Format

### Standard Quality Dashboard Layout

```
QUALITY MEASURE DASHBOARD
==========================
Organization: [Name]
Measurement Period: [Start] to [End]
Report Generated: [Date]
Product Line: [Commercial / Medicare / Medicaid]

SUMMARY
-------
Measures Reported: [N]
Above 75th Percentile: [N] measures
Below 50th Percentile: [N] measures
Estimated Star Rating: [X.X] stars

DETAIL
------
[Table with: Measure | Denom | Excl | Numer | Rate | 50th | 75th | Status]

TREND
-----
[Table with: Measure | Prior Year Rate | Current Rate | Change | Direction]

GAPS
----
[Table with: Measure | Gap Count | Top 10 Patients | Recommended Action]
```

### Trend Tracking

Track monthly/quarterly rates to monitor progress before year-end:
- January rate = baseline (carry-forward from prior year for lookback measures)
- Quarterly check: project year-end rate assuming current trajectory
- September checkpoint: final push for gap closure before year-end

### Attribution

Patients attributed to a provider/practice for quality measurement:
- **Visit-based attribution:** Patient assigned to provider with plurality of E&M visits in past 24 months
- **Panel assignment:** Explicit PCP assignment in the health plan or EMR
- **FHIR approach:** Use `Patient.generalPractitioner` reference or Encounter-based attribution
