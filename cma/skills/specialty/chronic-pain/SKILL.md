---
name: langcare-chronic-pain
description: >
  Performs comprehensive chronic pain management review including pain
  assessment scores, current analgesic regimen, functional status, opioid
  risk evaluation, and multimodal therapy assessment per CDC 2022 guidelines.
  Use when asked about chronic pain review, pain management assessment,
  non-opioid pain therapy, multimodal pain plan, or pain clinic evaluation.
---

# Chronic Pain Management Review

## When to Use This Skill
Use when a clinician needs a comprehensive evaluation of a chronic pain patient's current management including pain scores, medication regimen, functional status, and guideline compliance.

## Clinical Workflow
1. Use `fhir_read` to retrieve Patient demographics
2. Use `fhir_search` to pull Condition resources for pain diagnoses (chronic pain, neuropathy, fibromyalgia, low back pain, osteoarthritis)
3. Use `fhir_search` to pull Observation resources for pain scores (LOINC 72514-3) and functional assessments over time
4. Use `fhir_search` to pull active MedicationRequest for current pain regimen: opioids, NSAIDs, acetaminophen, gabapentinoids, antidepressants (duloxetine, TCAs), topicals, muscle relaxants
5. If opioids present, calculate total daily MME and apply CDC 2022 thresholds (see references/cdc-pain-guidelines.md)
6. Use `fhir_search` to pull Procedure resources for interventional pain procedures (injections, nerve blocks, spinal cord stimulation)
7. Assess multimodal therapy: physical therapy referrals, CBT/psychology referrals, complementary medicine
8. Present comprehensive pain management dashboard with pain trend, current regimen, MME (if applicable), functional status, and recommended optimizations

## FHIR Resources
- **Condition** -- Pain diagnoses with ICD-10/SNOMED codes
- **Observation** -- Pain scores (LOINC 72514-3), functional assessments, UDS
- **MedicationRequest** -- Current pain medications
- **Procedure** -- Interventional procedures, physical therapy
- **ServiceRequest** -- Referrals (pain clinic, PT, psychology, complementary)

## FHIR Query Examples
### Pull Pain Scores Over Time
```
fhir_search(resourceType="Observation", queryParams="patient=[patient-id]&code=72514-3&_sort=date&_count=20")
```

### Pull Pain Diagnoses
```
fhir_search(resourceType="Condition", queryParams="patient=[patient-id]&clinical-status=active")
```

### Pull Pain Medications
```
fhir_search(resourceType="MedicationRequest", queryParams="patient=[patient-id]&status=active&_count=100")
```

## Clinical Guidelines
- CDC 2022 Clinical Practice Guideline for Prescribing Opioids for Chronic Pain
- VA/DoD Clinical Practice Guideline for Opioid Therapy for Chronic Pain
- ACR/ACP guidelines for specific pain conditions (OA, low back pain, fibromyalgia)
- IASP multimodal pain management framework

## Interpretation Guide
- Pain score trend: improving (decreasing scores), stable, worsening (increasing scores); correlate with functional goals rather than pain score alone
- Multimodal assessment: document which modalities have been tried (pharmacologic: non-opioid analgesics, adjuvants; interventional: injections, nerve blocks; rehabilitative: PT, OT; psychological: CBT, ACT, mindfulness; complementary: acupuncture, massage)
- CDC 2022 key recommendations: non-opioid therapy preferred for chronic pain; if opioids used, start lowest effective dose, reassess at 1-4 weeks; continue only if meaningful improvement in pain AND function
- Functional goals: use PEG scale (Pain, Enjoyment, General Activity 0-10) or similar functional outcome measure rather than pain score alone
- Flag for optimization: opioid-only regimen without non-opioid adjuncts, no PT trial, no behavioral health referral for pain coping, MME >50 without documented benefit

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
