---
name: langcare-lab-interpretation
description: >
  Retrieves, organizes, and interprets laboratory results with clinical context
  including delta checks, abnormal pattern recognition, and drug-lab correlations.
  Use when asked to interpret labs, review lab results, explain bloodwork, check
  labs, lab trends, or abnormal labs. Flags critical values requiring immediate
  action.
---

# Lab Result Interpreter

## When to Use This Skill
Use when a clinician needs laboratory results organized by panel, interpreted with clinical context, flagged for abnormals and critical values, and correlated with active medications and conditions.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age, gender for reference range selection)
2. Use `fhir_search` to pull recent Observation resources (category=laboratory) sorted by date
3. Organize results by panel: CBC, BMP/CMP, liver, lipids, thyroid, coagulation, urinalysis
4. Flag abnormal values using interpretation field and referenceRange; classify as normal, abnormal, or critical
5. Calculate delta checks: pull prior values for each abnormal result, compute absolute and percent change
6. Identify abnormal patterns (e.g., microcytic anemia, hepatocellular injury, cholestatic pattern, anion gap)
7. Use `fhir_search` to pull active Conditions and MedicationStatements for clinical correlation (drug-lab interactions)
8. Present structured report with trending, pattern analysis, and recommended actions

## FHIR Resources
- **Observation** -- Lab results: code (LOINC), valueQuantity, interpretation, referenceRange, effectiveDateTime, status, component
- **Patient** -- Age/sex for reference range selection
- **Condition** -- Clinical correlation with active diagnoses
- **MedicationStatement** -- Drug-lab interaction correlation

## FHIR Query Examples
### Pull Recent Labs (30 days)
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&category=laboratory&date=ge[30-days-ago]&_sort=-date&_count=200")
```

### Pull Specific Panel (CBC)
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=58410-2&_sort=-date&_count=5")
```

### Pull Prior Values for Delta Check
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=[loinc-code]&_sort=-date&_count=5")
```

## Clinical Guidelines
- CAP/CLIA critical value reporting requirements
- KDIGO guidelines for renal function interpretation
- ADA standards for diabetes lab monitoring
- ACC/AHA guidelines for lipid interpretation

## Interpretation Guide
- Group by panel and present in tabular format with value, reference range, and interpretation flag
- Significant delta thresholds: Hgb drop >2 g/dL, K+ change >1.0 mEq/L, Cr increase >0.3 mg/dL or >50% (AKI criteria), platelet drop >50% (HIT concern), Na change >8 mEq/L in 24h
- Abnormal patterns: elevated AST+ALT with AST/ALT >2 (alcoholic hepatitis), low MCV + low ferritin (iron deficiency), high MCV + low B12 (megaloblastic anemia), pancytopenia (bone marrow failure workup), elevated anion gap (MUDPILES differential)
- Common drug-lab interactions: metformin + elevated lactate, statins + elevated CK/ALT, ACEi + elevated K+/Cr, warfarin + INR, heparin + platelet drop (HIT)

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
