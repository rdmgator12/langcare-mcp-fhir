---
name: medication-adherence-assessment
description: |
  Assesses medication adherence by analyzing fill patterns from MedicationDispense records and calculating MPR and PDC metrics. Use when user asks to "check medication adherence", "assess compliance", "review refill history", mentions "medication possession ratio", "proportion of days covered", "missed refills", or needs adherence counseling support. Do NOT use for drug interactions, reconciliation, or prescribing appropriateness.
metadata:
  author: LangCare
  version: 1.0.0
  mcp-server: langcare-mcp-fhir
  category: medication-management
---

# Medication Adherence Assessment

## Overview

Analyze medication fill patterns using MedicationDispense and MedicationRequest records. Calculate Medication Possession Ratio (MPR) and Proportion of Days Covered (PDC) for each chronic medication. Flag gaps in refills greater than 7 days, early refill patterns suggesting hoarding, and polypharmacy burden exceeding 10 medications. Apply WHO 5-dimension adherence framework to identify modifiable barriers. Provide motivational interviewing prompts for adherence counseling.

## FHIR Resources Used

| Resource | Purpose | Key Fields |
|----------|---------|------------|
| MedicationDispense | Pharmacy fill records | status, medicationCodeableConcept, quantity, daysSupply, whenHandedOver |
| MedicationRequest | Prescription orders with intended supply | status, medicationCodeableConcept, dispenseRequest.numberOfRepeatsAllowed, dispenseRequest.expectedSupplyDuration |
| MedicationStatement | Patient-reported medication taking behavior | status, medicationCodeableConcept, effectivePeriod, adherence |
| Patient | Demographics for risk stratification | birthDate, gender, address |
| Condition | Chronic conditions requiring adherence | code, clinicalStatus, onsetDateTime |
| Observation | Clinical outcomes correlating with adherence | code, valueQuantity, effectiveDateTime |

## Instructions

### Step 1: Identify Patient and Active Chronic Medications

```
Tool: fhir_search
resourceType: "MedicationRequest"
queryParams: "patient=[patient-id]&status=active&intent=order&_count=100"
```

Filter for chronic medications (those with repeating dispense requests or ongoing conditions). Exclude PRN-only and short-course medications (antibiotics, acute pain).

### Step 2: Pull Dispense History for Each Medication

For each chronic medication identified:
```
Tool: fhir_search
resourceType: "MedicationDispense"
queryParams: "patient=[patient-id]&medication=[medication-code]&_sort=-whenhandedover&_count=50"
```

If medication-specific search fails, pull all dispenses:
```
Tool: fhir_search
resourceType: "MedicationDispense"
queryParams: "patient=[patient-id]&status=completed&_sort=-whenhandedover&_count=200"
```

Extract for each fill:
- `whenHandedOver` (dispense date)
- `quantity.value` (tablets/units dispensed)
- `daysSupply.value` (days supply)
- `medicationCodeableConcept` (drug identity)

### Step 3: Calculate Adherence Metrics

**Medication Possession Ratio (MPR)**:
```
MPR = (Total days supply dispensed in period) / (Number of days in period) x 100

Evaluation period: typically 12 months or since first fill
```

**Proportion of Days Covered (PDC)** (preferred measure):
```
PDC = (Number of days with medication available) / (Number of days in period) x 100

Key difference from MPR: PDC does not double-count overlapping fills.
When fills overlap, shift the later fill start to after the previous supply runs out.
```

See references/adherence-metrics.md for detailed calculation formulas and interpretation.

### Step 4: Assess Fill Pattern Anomalies

**Refill Gaps > 7 days**:
Calculate gap between each fill's expected end date (whenHandedOver + daysSupply) and the next fill date. Flag any gap exceeding 7 days.

```
Gap = next_fill_date - (current_fill_date + current_days_supply)
If gap > 7 days: FLAG as potential non-adherence
```

**Early Refills**:
If a fill occurs more than 7 days before the previous supply should have run out:
```
Early_days = (current_fill_date + current_days_supply) - next_fill_date
If early_days > 7 and pattern is recurring: FLAG as potential hoarding/diversion concern
```

**Escalating Gaps**:
If gaps between fills are progressively increasing, flag as declining adherence trend.

### Step 5: Assess Polypharmacy Burden

Count total active medications from Step 1.

| Count | Classification | Action |
|-------|---------------|--------|
| 1-4 | Low burden | Standard monitoring |
| 5-9 | Moderate burden | Assess for simplification opportunities |
| 10-14 | High burden (polypharmacy) | Flag for comprehensive review, adherence support |
| 15+ | Extreme polypharmacy | Urgent deprescribing review |

Identify opportunities to reduce pill burden:
- Combination products (e.g., lisinopril/HCTZ instead of separate)
- Once-daily formulations instead of BID/TID
- Therapeutic duplication removal

### Step 6: Pull Clinical Outcome Correlations

For key conditions, pull relevant lab values to correlate with adherence:

**Diabetes (HbA1c)**:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=4548-4&_sort=-date&_count=4"
```

**Hypertension (BP readings)**:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=85354-9&_sort=-date&_count=6"
```

**Hyperlipidemia (LDL)**:
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=18262-6&_sort=-date&_count=4"
```

Correlate: Poor adherence (PDC <80%) + uncontrolled outcomes (HbA1c >7%, BP >140/90, LDL >100) = strong evidence adherence is the barrier.

### Step 7: Apply WHO 5-Dimension Framework

Assess each dimension (see references/adherence-interventions.md):

1. **Socioeconomic**: Insurance coverage, copay burden, health literacy
2. **Health system**: Pharmacy access, provider relationship, appointment frequency
3. **Condition-related**: Symptom burden, depression, cognitive impairment
4. **Therapy-related**: Side effects, complexity, pill burden
5. **Patient-related**: Beliefs, motivation, knowledge, forgetfulness

### Step 8: Generate Adherence Report

Structure output:

```
MEDICATION ADHERENCE ASSESSMENT - [Patient Name]
Assessment Date: [date]
Evaluation Period: [start] to [end]

SUMMARY:
- Total chronic medications: [N]
- Polypharmacy status: [Low/Moderate/High/Extreme]
- Medications with PDC >= 80%: [N] of [total]
- Medications with PDC < 80%: [N] of [total]

MEDICATION-LEVEL DETAIL:
| Medication | Fills (12mo) | MPR | PDC | Max Gap | Status |
|------------|-------------|-----|-----|---------|--------|
| Metformin 1000mg | 10 | 85% | 82% | 12 days | Borderline |
| Lisinopril 20mg | 7 | 58% | 55% | 45 days | Non-adherent |
| Atorvastatin 40mg | 11 | 92% | 91% | 5 days | Adherent |

FLAGS:
- Lisinopril: PDC 55% with BP readings averaging 158/94. Adherence likely primary barrier.
- [Early refill pattern on X medication - Y occurrences in 12 months]

CLINICAL CORRELATION:
- HbA1c: 7.8% (metformin PDC 82% - borderline adherence may contribute)
- BP: 158/94 average (lisinopril PDC 55% - strong adherence-outcome gap)

RECOMMENDED INTERVENTIONS:
1. Lisinopril: Explore barriers. Consider once-daily combination product.
2. Metformin: Assess GI side effects (common cause of gaps).
3. General: [motivational interviewing prompt based on identified barriers]
```

## Examples

### Example 1: Chronic Disease Adherence Review

**User says**: "Assess medication adherence for patient 45678 who has uncontrolled diabetes and hypertension."

**Actions**:
1. Pull active MedicationRequests for patient 45678: finds metformin 1000mg BID, glipizide 10mg BID, lisinopril 40mg daily, amlodipine 10mg daily, atorvastatin 40mg daily.
2. Pull MedicationDispense for each medication over past 12 months.
3. Calculate PDC: metformin 72%, glipizide 68%, lisinopril 45%, amlodipine 81%, atorvastatin 88%.
4. Pull HbA1c (8.9%) and BP (162/98).
5. Identify: lisinopril has 3 gaps >30 days. Glipizide and metformin have intermittent 10-14 day gaps.

**Result**:
```
ADHERENCE ASSESSMENT - Patient 45678

PDC BELOW THRESHOLD (<80%):
- Lisinopril 40mg: PDC 45% [CRITICAL - 3 gaps >30 days]
- Metformin 1000mg: PDC 72% [MODERATE - intermittent gaps 10-14 days]
- Glipizide 10mg: PDC 68% [MODERATE - pattern matches metformin gaps]

CLINICAL IMPACT:
- HbA1c 8.9% (goal <7%) - metformin+glipizide non-adherence contributing
- BP 162/98 (goal <130/80) - lisinopril PDC 45% is primary driver

INTERVENTION SUGGESTIONS:
- Lisinopril: Major gap pattern suggests intentional discontinuation. Explore: side effects (cough, dizziness), cost, perceived need.
  MI prompt: "I notice your blood pressure medication has some gaps in fills. What has your experience been with this medication?"
- Metformin/Glipizide: Parallel gap pattern suggests shared barrier. Explore: pharmacy access, insurance renewal timing, forgetfulness.
  MI prompt: "Sometimes it can be hard to keep up with multiple medications. What makes it difficult for you?"
- Consider: Combine lisinopril+amlodipine into single pill. Switch metformin to extended-release (once daily, less GI upset).
```

### Example 2: Polypharmacy and Simplification

**User says**: "Review medication adherence for elderly patient Mary Johnson, MRN 99887."

**Actions**:
1. Search Patient by MRN. Find age 82.
2. Pull 16 active MedicationRequests (extreme polypharmacy).
3. Pull dispense records. Calculate PDC for each.
4. Find 7 medications with PDC <80%. 3 medications with no fills in past 90 days.

**Result**:
```
ADHERENCE ASSESSMENT - Mary Johnson (Age 82)

POLYPHARMACY: EXTREME (16 active medications)

MEDICATIONS WITH NO RECENT FILLS (>90 days):
1. Omeprazole 20mg - Last fill: 142 days ago
2. Ferrous sulfate 325mg - Last fill: 98 days ago
3. Docusate 100mg - Last fill: 112 days ago

PDC BELOW THRESHOLD:
[... 7 medications listed ...]

SIMPLIFICATION OPPORTUNITIES:
1. Combine lisinopril 20mg + HCTZ 12.5mg into lisinopril/HCTZ combo (eliminate 1 pill)
2. Switch metoprolol tartrate 50mg BID to metoprolol succinate 100mg daily (eliminate 1 daily dose)
3. Omeprazole: no active indication found. Candidate for deprescribing.
4. Docusate: limited evidence for efficacy. Consider discontinuation.
5. Net reduction potential: 16 -> 12 medications

PILL BURDEN ANALYSIS:
Current: 22 pills/day across 16 medications
After simplification: 15 pills/day across 12 medications
```

## Troubleshooting

### MedicationDispense not available for this FHIR server
- Not all EMR FHIR endpoints expose MedicationDispense. This is common with EPIC patient-facing APIs.
- Fallback: Use MedicationRequest `authoredOn` dates as proxy for fill dates, combined with `dispenseRequest.expectedSupplyDuration`.
- Alternative: Check MedicationStatement for `adherence` extension if the system populates it.
- Note in the report that dispense-based PDC was unavailable and estimates are based on prescription dates.

### Days supply not populated in MedicationDispense
- Calculate from `quantity` and the dosing instructions in the corresponding MedicationRequest.
- Example: quantity 60 tablets, dose "1 tablet BID" = 30 days supply.
- If dosing frequency is unclear, use common defaults: tablets usually 30-day supply, inhalers 30-day supply, insulin vials 28-day supply.

### Patient has medications from multiple pharmacies
- MedicationDispense records may only reflect one pharmacy's data.
- Check if MedicationStatement has additional reported medications.
- Note in the report which medications have verified fill history and which are self-reported only.

## Related Skills

- `medication-reconciliation` - Ensure medication list is accurate before adherence assessment
- `prescription-appropriateness-review` - Assess if non-adherence is to an appropriate medication
- `drug-interaction-checker` - Non-adherence to one drug in an interacting pair changes risk profile
- `clinical-summary-generator` - Include adherence metrics in clinical summaries
