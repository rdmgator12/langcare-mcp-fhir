# HEDIS Quality Measure Specifications Reference

## Core HEDIS Measures

### HBD -- Hemoglobin A1c Control for Patients with Diabetes

| Component | Definition |
|-----------|-----------|
| Denominator | Patients 18-75 with Type 1 or Type 2 diabetes (SNOMED 44054006, 46635009; ICD-10 E10, E11) |
| Exclusion | Hospice, ESRD (eGFR <15 or dialysis) |
| Numerator | Most recent A1c (LOINC 4548-4) < 8.0% in measurement year |
| Sub-measure | A1c > 9.0% (poor control) -- inverse measure, lower is better |
| Lookback | Measurement year (Jan 1 - Dec 31) |

### CBP -- Controlling High Blood Pressure

| Component | Definition |
|-----------|-----------|
| Denominator | Patients 18-85 with hypertension (SNOMED 38341003; ICD-10 I10) |
| Exclusion | Hospice, ESRD, pregnancy, non-acute inpatient |
| Numerator | Most recent BP (LOINC 85354-9) with systolic < 140 AND diastolic < 90 |
| Lookback | Measurement year |

### BCS -- Breast Cancer Screening

| Component | Definition |
|-----------|-----------|
| Denominator | Women 52-74 (age as of Dec 31) |
| Exclusion | Bilateral mastectomy (SNOMED 27865001), hospice |
| Numerator | Mammogram (LOINC 24606-6 or SNOMED 71651007) within 27 months of measurement year end |
| Lookback | 27 months ending Dec 31 |

### COL -- Colorectal Cancer Screening

| Component | Definition |
|-----------|-----------|
| Denominator | Adults 45-75 |
| Exclusion | Colorectal cancer, total colectomy, hospice |
| Numerator | Any of: colonoscopy in 10 years, flex sig in 5 years, FIT/FOBT in 1 year, FIT-DNA in 3 years, CT colonography in 5 years |
| Lookback | Variable by test type |

### DSF -- Depression Screening and Follow-Up

| Component | Definition |
|-----------|-----------|
| Denominator | Patients 12+ with an eligible encounter in measurement year |
| Exclusion | Bipolar, depression diagnosis already active |
| Numerator | PHQ-9 or equivalent (LOINC 44249-1) documented in measurement year; if positive (>=10), follow-up plan documented |
| Lookback | Measurement year |

### SPD -- Statin Therapy for Patients with Diabetes

| Component | Definition |
|-----------|-----------|
| Denominator | Patients 40-75 with diabetes |
| Exclusion | Statin allergy/intolerance, hospice, ESRD, pregnancy, rhabdomyolysis history |
| Numerator | Active statin prescription (MedicationRequest with statin RxNorm codes, status=active) |
| Lookback | Active at any point during measurement year |

## CMS Star Rating Cut Points (Approximate)

| Measure | 2-Star | 3-Star | 4-Star | 5-Star |
|---------|--------|--------|--------|--------|
| A1c Control (<8%) | <67% | 67-77% | 77-85% | >85% |
| BP Control | <53% | 53-63% | 63-73% | >73% |
| Breast Cancer Screening | <60% | 60-70% | 70-80% | >80% |
| Colorectal Cancer Screening | <50% | 50-62% | 62-74% | >74% |
| Statin Therapy (DM) | <72% | 72-80% | 80-86% | >86% |

Note: Cut points are approximate and vary by year.

## FHIR Query Patterns for Measures

### Denominator Pattern
```
fhir_search(resourceType="Condition", queryParams="code=[dx-codes]&clinical-status=active&_count=500")
```
Then filter by Patient age using Patient.birthDate.

### Exclusion Pattern
```
fhir_search(resourceType="Condition", queryParams="code=[exclusion-dx-codes]&clinical-status=active&_count=500")
fhir_search(resourceType="Procedure", queryParams="code=[exclusion-procedure-codes]&_count=500")
```

### Numerator Pattern
```
fhir_search(resourceType="Observation", queryParams="code=[loinc-code]&date=ge[start]&_sort=-date&_count=500")
```
Take most recent per patient, check if value meets threshold.

## Statin RxNorm Codes (Common)

| Statin | RxNorm |
|--------|--------|
| Atorvastatin | 83367 |
| Rosuvastatin | 301542 |
| Simvastatin | 36567 |
| Pravastatin | 42463 |
| Lovastatin | 6472 |
| Fluvastatin | 41127 |
| Pitavastatin | 861634 |
