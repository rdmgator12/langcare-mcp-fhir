# Referral Requirements by Specialty

## Required Information Per Specialty Type

### All Referrals -- Minimum Required

| Element | Source | Notes |
|---------|--------|-------|
| Patient demographics | Patient resource | Name, DOB, gender, MRN, contact info |
| Insurance information | Coverage resource | Payor, plan type, subscriber ID |
| Referring provider | Practitioner resource | Name, NPI, contact info |
| Reason for referral | Condition resource | Specific diagnosis with code |
| Clinical question | Generated from context | What the specialist is being asked to evaluate/manage |
| Current medications | MedicationRequest | Full active medication list |
| Allergies | AllergyIntolerance | Drug and non-drug allergies |
| Relevant history | Condition (active) | Pertinent active conditions |

### Cardiology Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| Recent ECG | DiagnosticReport | Within 30 days |
| Recent labs: troponin, BNP, lipids, BMP | Observation | Most recent values |
| Blood pressure readings | Observation (vital-signs) | Last 3 readings |
| Prior cardiac history | Condition, Procedure | Cath, CABG, stents, valve procedures |
| Functional capacity | Clinical note | Exercise tolerance, NYHA class if HF |
| Echocardiogram (if available) | DiagnosticReport | Most recent |
| Chest X-ray | DiagnosticReport | Most recent |

### Endocrinology Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| HbA1c trend (last 3 values) | Observation | With dates |
| Current glucose log/CGM data | Observation | If available |
| Failed medication trials | MedicationRequest (stopped) | What was tried, why stopped |
| Thyroid labs | Observation | TSH, free T4, antibodies |
| BMI/weight trend | Observation (vital-signs) | Last 3 values |
| Relevant imaging | DiagnosticReport | Thyroid ultrasound, DEXA |

### Nephrology Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| Creatinine/eGFR trend | Observation | Last 4-6 values with dates |
| Urine albumin/creatinine ratio | Observation | Most recent |
| Urinalysis | Observation | Most recent |
| Electrolytes (BMP) | Observation | Most recent |
| Renal ultrasound | DiagnosticReport | If performed |
| Blood pressure history | Observation (vital-signs) | Especially if referring for renal HTN |
| Diabetes status | Condition | If diabetic nephropathy suspected |

### Gastroenterology Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| Hepatic panel (ALT, AST, ALP, bilirubin, albumin) | Observation | Most recent |
| CBC with differential | Observation | For anemia evaluation |
| Iron studies (if anemia) | Observation | Ferritin, iron, TIBC |
| Hepatitis B and C serologies | Observation | If liver disease |
| Stool studies | Observation | If diarrhea workup |
| Prior endoscopy/colonoscopy | Procedure | With pathology results |
| Imaging (CT abdomen, ultrasound) | DiagnosticReport | If performed |

### Pulmonology Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| PFTs | DiagnosticReport | If performed |
| Chest X-ray or CT | DiagnosticReport | Most recent |
| Smoking history | Observation | Pack-years |
| Oxygen saturation | Observation (vital-signs) | Resting and ambulatory if available |
| Current inhalers/respiratory meds | MedicationRequest | All respiratory medications |
| ABG | Observation | If severe disease |

### Rheumatology Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| Inflammatory markers (ESR, CRP) | Observation | Most recent |
| Autoimmune panel (ANA, RF, anti-CCP) | Observation | If ordered |
| Joint X-rays | DiagnosticReport | Of affected joints |
| CBC, CMP | Observation | Baseline labs |
| Prior joint aspirations | Procedure | Results including crystal analysis |
| Symptom duration and distribution | Clinical note | Which joints, how long, morning stiffness |

### Oncology Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| Pathology report | DiagnosticReport | REQUIRED -- type, grade, margins, receptors |
| Staging imaging | DiagnosticReport | CT, MRI, PET as appropriate |
| Tumor markers | Observation | Relevant to cancer type |
| CBC, CMP, LDH | Observation | Baseline labs |
| Performance status | Clinical note | ECOG or Karnofsky |
| Prior cancer treatment history | Procedure, MedicationRequest | Surgeries, chemo, radiation |
| Genetic testing | DiagnosticReport | If performed (BRCA, Lynch, etc.) |

### Psychiatry Referral

| Additional Required | Source | Notes |
|--------------------|--------|-------|
| Screening tool results (PHQ-9, GAD-7, etc.) | Observation | Most recent scores |
| Current psychotropic medications | MedicationRequest | All psych meds with duration |
| Failed medication trials | MedicationRequest (stopped) | What was tried, duration, reason for discontinuation |
| Substance use history | Condition | Active and past |
| Suicide risk assessment | Clinical note | Most recent assessment |
| TSH, B12 | Observation | Rule out organic causes |
| Medical comorbidities | Condition | Conditions affecting psych treatment |

---

## Prior Authorization Common Triggers

### By Insurance Plan Type

| Plan Type | Specialist Visit Auth | Procedure Auth | Imaging Auth |
|-----------|----------------------|----------------|--------------|
| HMO | Required for all | Required | Required for advanced (MRI, CT, PET) |
| PPO | Usually not required | Required for major | Required for advanced |
| POS | Required if out-of-network | Required | Required for advanced |
| EPO | Required for all | Required | Required for advanced |
| Medicare Advantage | Varies by plan | Required for major | Required for advanced |
| Traditional Medicare | Not required | Not required for most | Not required |
| Medicaid | Varies by state | Usually required | Usually required |

### Procedures Commonly Requiring Prior Authorization

#### Cardiology
- Cardiac catheterization
- Stress testing (nuclear, echo)
- Echocardiogram (transesophageal)
- Cardiac MRI
- Electrophysiology studies
- Pacemaker/ICD implantation
- Cardiac rehabilitation

#### Gastroenterology
- Upper endoscopy (EGD)
- Colonoscopy (screening may be excluded)
- ERCP
- Capsule endoscopy
- FibroScan / liver elastography
- Biologic medications (infliximab, adalimumab)

#### Oncology
- Chemotherapy regimens
- Radiation therapy
- PET scans
- Genetic testing
- Targeted therapy / immunotherapy
- Bone marrow biopsy

#### Rheumatology
- Biologic medications (all)
- JAK inhibitors
- Advanced imaging (MRI joints)
- Joint injections (some plans)

#### Neurology
- MRI brain/spine
- EEG / video EEG
- EMG / nerve conduction studies
- Botox injections
- MS medications (all DMTs)
- Sleep studies

#### General
- Advanced imaging: MRI, CT with contrast, PET, nuclear medicine
- Outpatient surgery / procedures
- DME (CPAP, braces, wheelchairs)
- Home health services
- Physical/occupational therapy (after initial visits)
- Infusion therapies
- Genetic / molecular testing

### Prior Authorization Submission Requirements

| Element | Required | Notes |
|---------|----------|-------|
| Patient demographics | Yes | Name, DOB, member ID |
| Diagnosis code (ICD-10) | Yes | Primary and supporting |
| CPT/HCPCS code | Yes | For the requested service |
| Clinical justification | Yes | Why the service is medically necessary |
| Supporting documentation | Often | Lab results, imaging, prior treatment failure |
| Referring provider NPI | Yes | |
| Specialist/facility NPI | Yes | Where service will be performed |
| Requested dates of service | Yes | Start date, duration |

---

## Referral Letter Format

### Structure

```
[Date]
[Referring Provider Name, Credentials]
[Practice Name]
[Address]
[Phone] | [Fax]

RE: Referral for [Patient Name]
DOB: [Date of Birth]
MRN: [Medical Record Number]
Insurance: [Payor] - [Plan] - [Member ID]

Dear [Specialist Name or "Colleague"],

I am referring [Patient Name], a [age]-year-old [sex], for evaluation and
management of [primary reason for referral].

CLINICAL QUESTION:
[Specific question for the specialist]

HISTORY OF PRESENT ILLNESS:
[Brief relevant history - onset, duration, severity, associated symptoms,
what has been tried, response to treatment]

RELEVANT PAST MEDICAL HISTORY:
- [Condition 1]
- [Condition 2]

CURRENT MEDICATIONS:
- [Medication 1] [dose] [frequency]
- [Medication 2] [dose] [frequency]

ALLERGIES:
- [Allergy 1]: [reaction]
- [Allergy 2]: [reaction]

RELEVANT LABS/IMAGING:
- [Lab/Study] [value] [units] ([date])
- [Lab/Study] [value] [units] ([date])

PRIOR AUTHORIZATION STATUS:
[Obtained / Pending / Not required]

Please contact our office if you need additional information.

Sincerely,
[Referring Provider Name, Credentials]
[NPI]
```

### Key Principles for Effective Referral Letters
1. State the clinical question explicitly -- do not assume the specialist will know what you want
2. Include what has already been tried and the response
3. Provide relevant data, not exhaustive data -- the specialist does not need the full chart
4. Note urgency level and preferred timeframe
5. Include callback information for questions
6. Specify if you want the specialist to co-manage or assume care
