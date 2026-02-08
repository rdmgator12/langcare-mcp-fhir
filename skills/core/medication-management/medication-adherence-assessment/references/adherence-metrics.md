# Adherence Metrics Reference

## Medication Possession Ratio (MPR)

### Formula

```
MPR = (Sum of days supply for all fills in the measurement period) / (Number of days in the measurement period) x 100
```

### Calculation Steps

1. Define the measurement period (typically 12 months from first fill or a fixed calendar year).
2. Identify all fills for the medication within the period.
3. Sum the `daysSupply` from each MedicationDispense.
4. Divide by the number of days in the period.
5. Multiply by 100 for percentage.

### Example Calculation

Patient fills lisinopril 20mg:
- Fill 1: Jan 1, 30 days supply
- Fill 2: Feb 5, 30 days supply (5-day gap)
- Fill 3: Mar 10, 30 days supply (3-day gap)
- Fill 4: May 1, 30 days supply (22-day gap)
- Fill 5: Jun 2, 30 days supply (2-day gap)
- Fill 6: Jul 1, 30 days supply (no gap)
- Fill 7: Aug 3, 30 days supply (3-day gap)
- Fill 8: Sep 5, 30 days supply (3-day gap)
- Fill 9: Nov 1, 30 days supply (27-day gap)
- Fill 10: Dec 1, 30 days supply (no gap)

Total days supply = 10 x 30 = 300 days
Measurement period = 365 days (Jan 1 - Dec 31)
MPR = 300 / 365 x 100 = **82.2%**

### Limitations of MPR

- Can exceed 100% if fills overlap (early refills counted at face value)
- Does not account for stockpiling
- Assumes all dispensed medication was taken
- Does not distinguish between intentional and unintentional non-adherence

---

## Proportion of Days Covered (PDC)

### Formula

```
PDC = (Number of days in the period "covered" by at least one fill) / (Number of days in the measurement period) x 100
```

### Calculation Steps

1. Define the measurement period.
2. For each day in the period, determine if the patient had medication available based on fill dates and days supply.
3. When fills overlap, do NOT double-count. Shift the start of the overlapping fill to the day after the previous supply ends.
4. Count total days covered.
5. Divide by period length. PDC is capped at 100%.

### Example Calculation (same fills as above)

```
Jan 1-30: Covered (Fill 1, 30 days) = 30 days
Jan 31-Feb 4: NOT covered = 5 days gap
Feb 5-Mar 6: Covered (Fill 2, 30 days) = 30 days
Mar 7-Mar 9: NOT covered = 3 days gap
Mar 10-Apr 8: Covered (Fill 3, 30 days) = 30 days
Apr 9-Apr 30: NOT covered = 22 days gap
May 1-May 30: Covered (Fill 4, 30 days) = 30 days
May 31-Jun 1: NOT covered = 2 days gap
Jun 2-Jul 1: Covered (Fill 5, 30 days) = 30 days
Jul 1-Jul 30: Covered (overlaps, shift to Jul 2-Jul 31) = 30 days
Aug 1-Aug 2: NOT covered = 2 days gap (shifted)
Aug 3-Sep 1: Covered (Fill 7, 30 days) = 30 days
Sep 2-Sep 4: NOT covered = 3 days gap
Sep 5-Oct 4: Covered (Fill 8, 30 days) = 30 days
Oct 5-Oct 31: NOT covered = 27 days gap
Nov 1-Nov 30: Covered (Fill 9, 30 days) = 30 days
Dec 1-Dec 30: Covered (Fill 10, 30 days) = 30 days

Total covered days = 300 days
Total gap days = 64 days (5+3+22+2+2+3+27 = 64)
Wait: 300 + 64 = 364. Period = 365. Dec 31 uncovered = 1 day.
Total covered = 300 days
PDC = 300 / 365 x 100 = 82.2%
```

In this case MPR and PDC are similar. They diverge when fills overlap.

### PDC with Overlapping Fills

If Fill 6 was picked up on Jun 28 instead of Jul 1 (3 days early):
- Fill 5 covers Jun 2 - Jul 1
- Fill 6 picked up Jun 28, but medication from Fill 5 lasts until Jul 1
- Shift Fill 6 start to Jul 2, covers Jul 2 - Jul 31
- MPR stays the same (still 300 days supply)
- PDC also stays the same in this case (no overlap gap)

But if the patient picks up 15 fills in 12 months (450 days supply):
- MPR = 450/365 = 123% (misleadingly high)
- PDC = capped at 100% (accurate representation)

PDC is preferred by CMS and PQA (Pharmacy Quality Alliance).

---

## Interpretation Thresholds

| PDC / MPR | Classification | Clinical Implication |
|-----------|---------------|---------------------|
| >= 80% | Adherent | Adequate medication exposure for therapeutic effect |
| 60-79% | Partially adherent | Suboptimal. Likely contributing to poor outcomes. |
| 40-59% | Non-adherent | Significant gaps. Medication unlikely to achieve full effect. |
| < 40% | Severely non-adherent | Minimal medication exposure. Equivalent to untreated in many cases. |

### Condition-Specific Thresholds

| Condition | Adherence Measure | Threshold | Source |
|-----------|------------------|-----------|--------|
| Diabetes (oral agents) | PDC | >= 80% | CMS Star Rating |
| Hypertension (RASA) | PDC | >= 80% | CMS Star Rating |
| Hyperlipidemia (statins) | PDC | >= 80% | CMS Star Rating |
| HIV antiretrovirals | PDC | >= 95% | DHHS Guidelines |
| Transplant immunosuppressants | PDC | >= 95% | Transplant societies |
| Epilepsy (antiepileptics) | PDC | >= 80% | AAN Guidelines |
| Asthma (controller inhalers) | PDC | >= 75% | NAEPP |
| Depression (antidepressants) | PDC (acute phase) | >= 80% for 84 days | HEDIS |
| Depression (continuation) | PDC (continuation) | >= 80% for 180 days | HEDIS |

---

## WHO 5 Dimensions of Adherence

The World Health Organization identifies 5 interacting dimensions that affect medication adherence. No single factor is sufficient to explain non-adherence.

### 1. Socioeconomic Factors
- Income and insurance coverage
- Medication cost and copay burden
- Health literacy level
- Cultural and language barriers
- Social support system
- Housing stability
- Employment and work schedule conflicts

### 2. Health System / Healthcare Team Factors
- Patient-provider relationship quality
- Appointment availability and wait times
- Pharmacy accessibility (distance, hours)
- Care coordination across providers
- Prescription complexity and fragmentation
- Prior authorization barriers
- Follow-up and monitoring frequency

### 3. Condition-Related Factors
- Symptom severity and fluctuation
- Disease understanding and perceived seriousness
- Comorbid depression (reduces adherence by 20-30%)
- Comorbid cognitive impairment
- Comorbid substance use disorder
- Rate of disease progression
- Presence vs absence of symptoms (asymptomatic conditions have lower adherence)

### 4. Therapy-Related Factors
- Medication side effects (real or anticipated)
- Dosing complexity (frequency, timing, food requirements)
- Pill burden (number of medications)
- Treatment duration (adherence declines over time)
- Route of administration
- Medication taste, size, or formulation
- Time to perceive benefit
- Prior treatment failures

### 5. Patient-Related Factors
- Health beliefs and medication attitudes
- Self-efficacy (confidence in ability to adhere)
- Perceived necessity vs perceived harm of medication
- Forgetfulness
- Motivation and readiness for change
- Fear of side effects or dependency
- Perceived stigma of the condition
- Competing priorities

---

## CMS Star Rating Adherence Measures

CMS Medicare Part D Star Ratings include three medication adherence measures, all using PDC with an 80% threshold:

### D12: Medication Adherence for Diabetes Medications
- Includes: biguanides, sulfonylureas, thiazolidinediones, DPP-4 inhibitors, incretin mimetics, meglitinides, SGLT2 inhibitors
- Excludes: insulin (not measured via PDC due to variable dosing)
- Measurement period: calendar year
- Denominator: patients with >= 2 fills

### D13: Medication Adherence for Hypertension (RASA)
- Includes: ACE inhibitors, ARBs, direct renin inhibitors
- Measurement period: calendar year
- Denominator: patients with >= 2 fills

### D14: Medication Adherence for Cholesterol (Statins)
- Includes: all statin medications (single and combination products)
- Measurement period: calendar year
- Denominator: patients with >= 2 fills

### Scoring
- 5 Stars: >= 89% of eligible patients adherent
- 4 Stars: 84-88%
- 3 Stars: 78-83%
- 2 Stars: 72-77%
- 1 Star: < 72%

---

## Gap Analysis Methods

### Simple Gap Calculation
```
For each consecutive fill pair:
  gap_days = next_fill_date - (current_fill_date + current_days_supply)
  if gap_days > 0: record gap
  if gap_days > 7: flag as clinically significant
  if gap_days > 30: flag as critical
```

### Permissible Gap
Some measures allow a "permissible gap" (typically 1-7 days) to account for pharmacy processing time and minor delays. Gaps within the permissible window are not counted as non-adherent days.

### Seasonality Assessment
Look for patterns in gap timing:
- Consistent gaps at end of calendar year (deductible reset)
- Gaps around holidays (pharmacy closure)
- Gaps during summer (travel, routine change)
- Gaps during illness (hospitalizations may interrupt outpatient fills)
