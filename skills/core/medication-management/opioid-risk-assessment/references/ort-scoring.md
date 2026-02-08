# Opioid Risk Tool (ORT) Scoring Reference

Source: Webster LR, Webster RM. Predicting Aberrant Behaviors in Opioid-Treated Patients: Preliminary Validation of the Opioid Risk Tool. Pain Medicine. 2005;6(6):432-442.

## Overview

The Opioid Risk Tool (ORT) is a brief, self-report screening instrument designed to predict which patients are at risk for opioid misuse, abuse, or aberrant behavior when prescribed opioids for chronic pain. It is administered BEFORE initiating opioid therapy.

Scoring differs by gender (Male vs Female).

---

## ORT Scoring Criteria

### Family History of Substance Abuse

| Factor | Female Score | Male Score |
|--------|------------|-----------|
| Family history of alcohol abuse | 1 | 3 |
| Family history of illegal drug abuse | 2 | 3 |
| Family history of prescription drug abuse | 4 | 4 |

**FHIR Source**: FamilyMemberHistory resource with condition codes for substance use disorders. May also be documented as Condition with `category=family-history`.

**SNOMED codes for substance use disorders**:
- Alcohol abuse: 15167005
- Alcohol dependence: 7200002
- Drug abuse: 26416006
- Drug dependence: 191816009
- Prescription drug abuse: 5602001

**Note**: Family history is often incompletely documented in EHRs. If not found in FHIR data, flag as "unable to assess from records - clinician interview needed."

### Personal History of Substance Abuse

| Factor | Female Score | Male Score |
|--------|------------|-----------|
| Personal history of alcohol abuse | 3 | 3 |
| Personal history of illegal drug abuse | 4 | 4 |
| Personal history of prescription drug abuse | 5 | 5 |

**FHIR Source**: Condition resource with substance use disorder codes. Check both active and resolved/inactive conditions.

**SNOMED codes**:
- Alcohol abuse/dependence: 15167005, 7200002, 44870007
- Illicit drug use disorder: 307052004, 191816009
- Opioid abuse/dependence: 75544000, 191909007
- Cannabis use disorder: 37344009
- Cocaine use disorder: 4448006, 78267003
- Stimulant use disorder: 191889006
- Sedative use disorder: 231477003
- Prescription drug abuse: 5602001

**ICD-10 codes**:
- F10.x: Alcohol-related disorders
- F11.x: Opioid-related disorders
- F12.x: Cannabis-related disorders
- F13.x: Sedative-related disorders
- F14.x: Cocaine-related disorders
- F15.x: Stimulant-related disorders
- F19.x: Other psychoactive substance-related disorders

### Age (16-45 years)

| Factor | Female Score | Male Score |
|--------|------------|-----------|
| Age 16-45 years | 4 | 0 |

**Note**: This is a significant risk factor ONLY for females in the ORT. For males, age is not scored. Calculate from Patient.birthDate.

### Psychological Disease

| Factor | Female Score | Male Score |
|--------|------------|-----------|
| ADD/ADHD, OCD, Bipolar disorder, Schizophrenia | 2 | 3 |
| Depression | 1 | 1 |

**FHIR Source**: Condition resource

**SNOMED codes**:
- ADHD: 406506008
- OCD: 191736004
- Bipolar disorder: 13746004
- Schizophrenia: 58214004
- Major depressive disorder: 35489007
- Persistent depressive disorder (dysthymia): 78667006
- Generalized anxiety disorder: 197480006
- PTSD: 47505003
- Schizoaffective disorder: 68890003

**ICD-10 codes**:
- F90.x: ADHD
- F42.x: OCD
- F31.x: Bipolar disorder
- F20.x: Schizophrenia
- F32.x, F33.x: Depressive disorders
- F41.x: Anxiety disorders
- F43.1: PTSD

**Note**: Depression scores 1 for both genders. ADD/ADHD/OCD/Bipolar/Schizophrenia score 2 (female) or 3 (male). If a patient has both depression AND another listed psychiatric condition, both are scored.

### History of Preadolescent Sexual Abuse

| Factor | Female Score | Male Score |
|--------|------------|-----------|
| History of preadolescent sexual abuse | 3 | 0 |

**Note**: This is a significant risk factor ONLY for females. This information is rarely documented in structured FHIR data. Flag as "unable to assess from records - sensitive topic, clinician interview needed."

---

## Score Summary Table

### Female Scoring

| Risk Factor | Possible Points |
|------------|----------------|
| Family history alcohol abuse | 1 |
| Family history illegal drug abuse | 2 |
| Family history Rx drug abuse | 4 |
| Personal history alcohol abuse | 3 |
| Personal history illegal drug abuse | 4 |
| Personal history Rx drug abuse | 5 |
| Age 16-45 | 4 |
| Depression | 1 |
| ADD/ADHD/OCD/Bipolar/Schizophrenia | 2 |
| Preadolescent sexual abuse | 3 |
| **Maximum possible score** | **29** |

### Male Scoring

| Risk Factor | Possible Points |
|------------|----------------|
| Family history alcohol abuse | 3 |
| Family history illegal drug abuse | 3 |
| Family history Rx drug abuse | 4 |
| Personal history alcohol abuse | 3 |
| Personal history illegal drug abuse | 4 |
| Personal history Rx drug abuse | 5 |
| Age 16-45 | 0 |
| Depression | 1 |
| ADD/ADHD/OCD/Bipolar/Schizophrenia | 3 |
| Preadolescent sexual abuse | 0 |
| **Maximum possible score** | **26** |

---

## Risk Stratification

| ORT Score | Risk Category | Predicted Aberrant Behavior Rate | Recommended Actions |
|-----------|--------------|--------------------------------|-------------------|
| 0-3 | Low Risk | ~6% | Standard monitoring. Annual UDS. Standard prescribing intervals. |
| 4-7 | Moderate Risk | ~28% | Enhanced monitoring. UDS every 6 months. Shorter prescriptions (14-28 day supply). Pill counts. Prescription drug monitoring program (PDMP) check at each visit. |
| >= 8 | High Risk | ~91% | Intensive monitoring. UDS every 1-3 months. 7-14 day prescriptions. Pill counts. PDMP every visit. Consider pain management specialty referral. Informed consent for opioid therapy with documented discussion of risks. Consider non-opioid alternatives first. |

---

## Implementation Notes for FHIR-Based Assessment

### Data Availability Challenges

The ORT relies on patient self-report and history that may not be fully captured in structured FHIR data. Common gaps:

1. **Family history**: Rarely coded as structured FHIR resources. May appear in clinical notes (unstructured).
2. **Preadolescent sexual abuse**: Almost never in structured data. Sensitive topic requiring direct patient interview.
3. **Personal substance use history**: May be documented as Condition but could be in social history (unstructured).
4. **Age**: Always available from Patient.birthDate.
5. **Psychiatric diagnoses**: Usually well-documented as Condition resources.

### Scoring with Missing Data

When FHIR data is incomplete:

1. Score what is available from structured data.
2. Present a **minimum score** (assuming all missing factors are 0) and **maximum score** (assuming all missing factors are positive).
3. List the specific factors that could not be assessed.
4. Recommend clinician complete the full ORT via patient interview.

**Example output**:
```
ORT SCORE (from available FHIR data):
- Confirmed factors: Depression (+1), Age 32 female (+4) = 5
- Unable to assess: Family history, personal SUD history, preadolescent sexual abuse
- Minimum possible score: 5 (Moderate Risk)
- Maximum possible score: 24 (High Risk)
- Recommendation: Complete ORT via patient interview before prescribing decision.
```

### Complementary Screening Tools

The ORT should be used alongside other assessments:

| Tool | Purpose | When to Use |
|------|---------|-------------|
| CAGE-AID | Screen for current substance use problems | All patients considered for opioids |
| PHQ-9 | Depression severity | If depression identified |
| DAST-10 | Drug abuse screening | If substance use suspected |
| SOAPP-R (14-item) | Longer validated risk assessment | When ORT is moderate/high risk |
| COMM (Current Opioid Misuse Measure) | Monitor current patients on opioids | Ongoing monitoring (not pre-prescribing) |
| PEG Scale (Pain, Enjoyment, General Activity) | Functional outcomes | All patients on chronic opioids |

---

## Aberrant Behaviors to Monitor

When the ORT indicates moderate or high risk, monitor for these behaviors:

### Predictive of Addiction
- Requesting specific opioids by name
- Escalating dose without authorization
- Using opioids for non-pain purposes (mood, sleep, anxiety)
- Multiple prescribers or pharmacies ("doctor shopping")
- Obtaining opioids from non-medical sources
- Concurrent abuse of alcohol or illicit drugs
- Selling or giving away medications
- Forging or altering prescriptions
- Injecting oral formulations
- Recurrent prescription "losses"

### Less Predictive (May Reflect Undertreated Pain)
- Requesting dose increase
- Early refill requests
- Hoarding during periods of less pain
- Unsanctioned dose escalation once or twice
- Using opioids for other symptoms (insomnia, depression)
- Reluctance to taper despite side effects

### Documentation
Document all aberrant behaviors as clinical observations. Use structured notes or Observation resources when possible to maintain a longitudinal record.
