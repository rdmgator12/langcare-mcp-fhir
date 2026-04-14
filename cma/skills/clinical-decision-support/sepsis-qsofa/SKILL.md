---
name: langcare-sepsis-qsofa
description: >
  Screens patients for sepsis using qSOFA, SOFA, and SIRS criteria by pulling
  vitals and labs from FHIR Observation resources. Evaluates Surviving Sepsis
  Campaign hour-1 bundle compliance. Use when asked to screen for sepsis, check
  qSOFA score, SOFA score, SIRS criteria, sepsis risk, or when suspected
  infection with hemodynamic instability is present.
---

# Sepsis Screening (qSOFA/SOFA/SIRS)

## When to Use This Skill
Use when a clinician needs to calculate sepsis risk scores from vital signs and laboratory data, with bundle compliance assessment.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics (age for baseline risk)
2. Use `fhir_search` to pull vital signs (last 24h): temperature (8310-5), heart rate (8867-4), respiratory rate (9279-1), systolic BP (8480-6), SpO2 (2708-6)
3. Use `fhir_search` to pull labs: WBC (6690-2), lactate (2524-7), creatinine (2160-0), bilirubin (1975-2), platelets (777-3), PaO2 (2703-7), FiO2 (3150-0)
4. Use `fhir_search` to pull active Condition resources for infection source identification
5. Use `fhir_search` to pull active MedicationRequest for vasopressor status and current antibiotics
6. Calculate qSOFA (0-3), SIRS (0-4), and SOFA (0-24) scores (see references/qsofa-sofa-scoring.md)
7. Evaluate hour-1 bundle compliance: lactate measured, blood cultures ordered, antibiotics administered, IV fluids initiated, vasopressors if MAP <65
8. Use `fhir_create` to create ClinicalImpression resource documenting the assessment

## FHIR Resources
- **Observation** -- Vitals (LOINC codes for temp, HR, RR, BP, SpO2) and labs (WBC, lactate, Cr, bilirubin, platelets, PaO2/FiO2)
- **Condition** -- Active infection source identification
- **MedicationRequest** -- Vasopressors (norepinephrine, epinephrine, vasopressin, dopamine) and antibiotics
- **Procedure** -- Mechanical ventilation status (SNOMED 40617009)
- **ServiceRequest** -- Blood culture orders (LOINC 600-7)
- **ClinicalImpression** -- Output: sepsis risk documentation

## FHIR Query Examples
### Pull Recent Vitals
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=8310-5&_sort=-date&_count=5")
```

### Pull Lactate
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=2524-7,32693-4,2519-7&_sort=-date&_count=3")
```

### Check Vasopressor Orders
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&code=3628,3616,11149,35208")
```

## Clinical Guidelines
- Surviving Sepsis Campaign 2021 International Guidelines
- Sepsis-3 Definitions (JAMA 2016): qSOFA for bedside screening, SOFA for organ dysfunction
- CMS SEP-1 Severe Sepsis and Septic Shock Management Bundle

## Interpretation Guide
- qSOFA >=2: high risk for poor outcome, investigate organ dysfunction. Components: RR >=22 (+1), SBP <=100 (+1), altered mental status/GCS <15 (+1)
- SIRS >=2 with suspected infection: meets SIRS-based sepsis definition. Components: temp >38.3 or <36 (+1), HR >90 (+1), RR >20 or PaCO2 <32 (+1), WBC >12K or <4K or >10% bands (+1)
- SOFA increase >=2: organ dysfunction (sepsis per Sepsis-3). Score 0-24 across 6 organ systems
- Septic shock: sepsis + vasopressor requirement + lactate >2 despite fluid resuscitation
- Bundle compliance: all elements should be initiated within 1 hour of sepsis recognition

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
