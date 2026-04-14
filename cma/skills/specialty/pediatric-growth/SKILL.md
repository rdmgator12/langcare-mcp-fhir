---
name: langcare-pediatric-growth
description: >
  Assesses pediatric growth by plotting weight, height/length, head
  circumference, and BMI against WHO (0-2 years) and CDC (2-20 years)
  growth chart percentiles. Flags growth faltering, obesity, and failure
  to thrive. Use when asked about pediatric growth, growth chart, growth
  percentiles, FTT, childhood obesity, or weight-for-age.
---

# Pediatric Growth Assessment

## When to Use This Skill
Use when a clinician needs to evaluate a pediatric patient's growth trajectory against age-appropriate percentile charts.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (birthDate for age calculation, gender for sex-specific charts)
2. Use `fhir_search` to pull Observation resources for weight (LOINC 29463-7), length/height (LOINC 8302-2 height, 8306-3 body length), head circumference (LOINC 9843-4), and BMI (LOINC 39156-5) over time
3. Determine chart set: WHO (0-2 years), CDC (2-20 years) -- see references/who-cdc-percentiles.md
4. Plot current and historical measurements against sex-specific percentile curves
5. Calculate percentiles and Z-scores for current measurements
6. Identify growth concerns: weight-for-age <3rd percentile (failure to thrive), BMI >85th (overweight), BMI >95th (obese), crossing >2 major percentile lines (growth faltering or acceleration), head circumference <3rd or >97th
7. Use `fhir_search` to pull Condition resources for underlying diagnoses affecting growth
8. Present growth dashboard with current percentiles, trajectory, and flagged concerns

## FHIR Resources
- **Patient** -- Age and sex for chart selection
- **Observation** -- Weight (29463-7), height (8302-2), length (8306-3), head circumference (9843-4), BMI (39156-5)
- **Condition** -- Growth-related diagnoses (FTT, obesity, short stature, endocrine conditions)

## FHIR Query Examples
### Pull Growth Measurements
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=29463-7,8302-2,8306-3,9843-4,39156-5&_sort=date&_count=100")
```

## Clinical Guidelines
- WHO growth standards (0-2 years): international reference for breastfed infants
- CDC growth charts (2-20 years): US reference population
- AAP guidelines for failure to thrive evaluation
- AAP Expert Committee recommendations for childhood obesity

## Interpretation Guide
- Use WHO charts for 0-2 years (based on breastfed infant norms), CDC charts for 2-20 years
- Percentile thresholds: <3rd concerning for undernutrition, 3rd-85th normal, 85th-95th overweight, >95th obese
- Track trajectory: crossing 2+ major percentile lines in either direction warrants evaluation
- Head circumference: <3rd may indicate microcephaly, >97th may indicate macrocephaly or hydrocephalus (monitor closely in first 18 months)
- For premature infants: use corrected gestational age until 24 months (corrected = chronological age minus weeks premature)
- BMI interpretation for children: use BMI-for-age percentile, not adult BMI categories

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
