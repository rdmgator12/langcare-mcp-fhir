# USPSTF Grade A and B Recommendations

## Overview

The U.S. Preventive Services Task Force (USPSTF) assigns letter grades to its recommendations based on the strength of evidence and the balance of benefits and harms. Grade A (high certainty of substantial net benefit) and Grade B (high certainty of moderate net benefit or moderate certainty of substantial net benefit) recommendations should be offered to eligible patients.

## Cancer Screening

### Breast Cancer Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Assigned female at birth, age 40-74 |
| Intervention | Screening mammography |
| Frequency | Biennial (every 2 years) |
| FHIR Resource | DiagnosticReport or Procedure |
| LOINC Code | 24606-6 (Mammography screening) |
| SNOMED Code | 71651007 (Mammography) |
| CPT Code | 77067 (bilateral screening mammography) |
| Risk Modifiers | Start earlier if BRCA carrier or first-degree relative with breast cancer. Consider annual if high risk. |
| Stop Criteria | Age 75+ (insufficient evidence, I statement) |

### Cervical Cancer Screening

| Parameter | Detail |
|-----------|--------|
| Grade | A |
| Population | Assigned female at birth with a cervix, age 21-65 |
| Intervention | Pap smear alone (21-29), Pap + HPV cotesting or HPV alone (30-65) |
| Frequency | Every 3 years (Pap alone, 21-65), every 5 years (HPV alone or cotesting, 30-65) |
| FHIR Resource | Observation or DiagnosticReport |
| LOINC Codes | 10524-7 (Pap smear), 21440-3 (HPV DNA), 71431-1 (HPV combined panel) |
| SNOMED Code | 171149006 (Cervical smear) |
| CPT Codes | 88141-88143 (Pap), 87624-87625 (HPV) |
| Risk Modifiers | More frequent if abnormal results, HIV positive, immunosuppressed |
| Stop Criteria | Age 65+ with adequate prior screening, or after total hysterectomy with cervix removal for benign indication |

### Colorectal Cancer Screening

| Parameter | Detail |
|-----------|--------|
| Grade | A (age 45-75), B (age 76-85 selective) |
| Population | All adults age 45-75 |
| Intervention | Multiple options (see below) |
| FHIR Resource | Procedure (colonoscopy), Observation (FIT, mt-sDNA) |
| Stop Criteria | Age 86+ |

#### Screening Modality Options

| Modality | Interval | LOINC/SNOMED | CPT |
|----------|----------|--------------|-----|
| Colonoscopy | 10 years | SNOMED 73761001 | 45378 |
| FIT (fecal immunochemical test) | Annual | LOINC 57905-2 | 82274 |
| mt-sDNA (Cologuard) | 3 years | LOINC 77353-1 | 81528 |
| Flexible sigmoidoscopy | 5 years | SNOMED 44441009 | 45330 |
| CT colonography | 5 years | LOINC 72142-3 | 74263 |
| FIT + sigmoidoscopy | FIT annual + sig every 10 years | Combined | Combined |

#### Risk-Based Modifications
- First-degree relative with CRC before age 60: Start at age 40 or 10 years before youngest case, whichever is earlier. Colonoscopy every 5 years.
- Two or more first-degree relatives with CRC at any age: Start at age 40, colonoscopy every 5 years.
- Personal history of adenomatous polyps: Surveillance colonoscopy per gastroenterology guidelines (not USPSTF scope).
- Lynch syndrome: Colonoscopy every 1-2 years starting age 20-25.

### Lung Cancer Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Adults age 50-80 with 20+ pack-year smoking history who currently smoke or quit within past 15 years |
| Intervention | Annual low-dose computed tomography (LDCT) |
| Frequency | Annual |
| FHIR Resource | DiagnosticReport |
| LOINC Code | 87278-0 (LDCT chest screening) |
| SNOMED Code | 241541005 (CT of chest) |
| CPT Code | 71271 (LDCT lung screening) |
| Risk Modifiers | Pack-year calculation: (packs per day) x (years smoked) |
| Stop Criteria | Age 81+, quit smoking > 15 years ago, or develops health problem limiting life expectancy or willingness to have curative surgery |

### Skin Cancer -- Insufficient Evidence
USPSTF has an I statement (insufficient evidence) for skin cancer screening in the general population. Not recommended as routine screening.

---

## Cardiovascular Screening

### Hypertension Screening

| Parameter | Detail |
|-----------|--------|
| Grade | A |
| Population | All adults age 18+ |
| Intervention | Office blood pressure measurement |
| Frequency | Annual (more frequently if elevated or risk factors) |
| FHIR Resource | Observation |
| LOINC Code | 85354-9 (BP panel), 8480-6 (systolic), 8462-4 (diastolic) |
| CPT Code | Part of routine visit |
| Confirmation | Ambulatory BP monitoring (ABPM) or home BP monitoring before diagnosing hypertension |

### Statin Use for CVD Prevention

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Adults age 40-75 with 1+ CVD risk factor (dyslipidemia, diabetes, hypertension, smoking) and estimated 10-year CVD risk >= 10% |
| Intervention | Low-to-moderate dose statin |
| Frequency | Ongoing; reassess risk periodically |
| FHIR Resource | MedicationRequest or MedicationStatement |
| Risk Calculation | Pooled Cohort Equations (PCE) for 10-year ASCVD risk |
| Required Inputs | Age, sex, race, total cholesterol, HDL, systolic BP, BP treatment status, diabetes status, smoking status |

### Abdominal Aortic Aneurysm (AAA) Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Males age 65-75 who have ever smoked |
| Intervention | One-time abdominal ultrasound |
| Frequency | One time |
| FHIR Resource | DiagnosticReport or Procedure |
| LOINC Code | 24850-0 (Abdominal ultrasound) |
| SNOMED Code | 241462003 (Ultrasound of abdominal aorta) |
| CPT Code | 76706 |
| Risk Modifiers | Selective screening for males who never smoked. Not routinely recommended for females. |

---

## Metabolic Screening

### Prediabetes / Type 2 Diabetes Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Adults age 35-70 who are overweight or obese (BMI >= 25, or >= 23 for Asian Americans) |
| Intervention | Fasting glucose, HbA1c, or oral glucose tolerance test |
| Frequency | Every 3 years if normal |
| FHIR Resource | Observation |
| LOINC Codes | 4548-4 (HbA1c), 1558-6 (fasting glucose), 1554-5 (glucose tolerance) |
| CPT Codes | 83036 (HbA1c), 82947 (glucose) |
| Thresholds | Prediabetes: HbA1c 5.7-6.4%, FG 100-125 mg/dL. Diabetes: HbA1c >= 6.5%, FG >= 126 mg/dL |
| Risk Modifiers | Screen earlier if family history, gestational diabetes, PCOS, certain racial/ethnic groups (Black, Hispanic, Asian, Native American) |

### Lipid Disorder Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Adults age 40-75 (for CVD risk assessment) |
| Intervention | Lipid panel |
| Frequency | Every 5 years (more frequently if borderline or on therapy) |
| FHIR Resource | Observation |
| LOINC Codes | 57698-3 (Lipid panel), 2093-3 (Total cholesterol), 18262-6 (LDL), 2085-9 (HDL), 2571-8 (Triglycerides) |
| CPT Code | 80061 (Lipid panel) |

---

## Infectious Disease Screening

### HIV Screening

| Parameter | Detail |
|-----------|--------|
| Grade | A |
| Population | All adolescents and adults age 15-65; all pregnant persons |
| Intervention | HIV 1/2 antigen/antibody combination test |
| Frequency | At least once; annually if high-risk |
| FHIR Resource | Observation |
| LOINC Code | 75622-1 (HIV 1/2 Ag+Ab) |
| CPT Code | 87389 |
| High-risk criteria | MSM, IVDU, persons with STI diagnosis, persons with multiple partners |

### Hepatitis C Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | All adults age 18-79 |
| Intervention | HCV antibody test |
| Frequency | One-time (more frequently if ongoing risk) |
| FHIR Resource | Observation |
| LOINC Code | 16128-1 (HCV Ab) |
| CPT Code | 86803 |

### Hepatitis B Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | All adolescents and adults age 15+ |
| Intervention | Triple panel (HBsAg, anti-HBs, anti-HBc) |
| Frequency | One-time (more frequently if ongoing risk) |
| FHIR Resource | Observation |
| LOINC Codes | 5195-3 (HBsAg), 5193-8 (anti-HBs), 16933-4 (anti-HBc) |
| CPT Codes | 87340, 86706, 86704 |

### Chlamydia and Gonorrhea Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Sexually active females age 24 and younger; older females at increased risk |
| Intervention | NAAT testing |
| Frequency | Annual (or as indicated by risk) |
| FHIR Resource | Observation |
| LOINC Codes | 43304-5 (Chlamydia NAAT), 43305-2 (Gonorrhea NAAT) |
| CPT Codes | 87491 (Chlamydia), 87591 (Gonorrhea) |

### Syphilis Screening

| Parameter | Detail |
|-----------|--------|
| Grade | A |
| Population | All persons at increased risk; all pregnant persons |
| Intervention | Serologic testing (RPR/VDRL with confirmatory) |
| Frequency | Per risk assessment |
| FHIR Resource | Observation |
| LOINC Code | 20507-0 (RPR) |

---

## Mental Health Screening

### Depression Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | All adults age 18+ (including pregnant and postpartum) |
| Intervention | PHQ-9 or other validated instrument |
| Frequency | Annual (more frequently for high-risk or during/after pregnancy) |
| FHIR Resource | Observation |
| LOINC Codes | 44249-1 (PHQ-9 total), 55758-7 (PHQ-2 total) |
| CPT Codes | 96127 (screening), G0444 (Medicare depression screening) |
| Scoring | PHQ-9: 0-4 minimal, 5-9 mild, 10-14 moderate, 15-19 moderately severe, 20-27 severe |

### Anxiety Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | All adults age 18-64 (including pregnant and postpartum) |
| Intervention | GAD-7 or other validated instrument |
| Frequency | Annual (when screening systems are in place) |
| FHIR Resource | Observation |
| LOINC Codes | 69737-5 (GAD-7 total) |
| Scoring | GAD-7: 0-4 minimal, 5-9 mild, 10-14 moderate, 15-21 severe |

### Unhealthy Drug Use Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | All adults age 18+ |
| Intervention | Single-question screening or validated tool (DAST-10, NIDA Quick Screen) |
| Frequency | Annual |
| FHIR Resource | Observation |
| LOINC Code | 82667-7 (Substance use screening) |

### Unhealthy Alcohol Use Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | All adults age 18+ |
| Intervention | AUDIT-C or single-question screening |
| Frequency | Annual |
| FHIR Resource | Observation |
| LOINC Code | 75626-2 (AUDIT-C total) |
| Scoring | AUDIT-C: >= 4 (men) or >= 3 (women) indicates unhealthy use |

---

## Musculoskeletal Screening

### Osteoporosis Screening

| Parameter | Detail |
|-----------|--------|
| Grade | B |
| Population | Females age 65+ (or postmenopausal females < 65 with risk factors per FRAX) |
| Intervention | DEXA scan (bone density measurement) |
| Frequency | Screening interval uncertain; typically every 2 years if osteopenia |
| FHIR Resource | DiagnosticReport |
| LOINC Code | 38269-7 (DEXA) |
| SNOMED Code | 312681000 (Bone density scan) |
| CPT Code | 77080 (DEXA central) |
| Risk Factors | Low body weight, parental hip fracture, smoking, excessive alcohol, steroid use, rheumatoid arthritis |

---

## Other USPSTF Grade B Recommendations

### Falls Prevention (Age 65+)
- Exercise interventions to prevent falls
- Not a screening test but a preventive intervention
- Document as CarePlan activity

### Folic Acid Supplementation
- All persons planning or capable of pregnancy: 0.4-0.8 mg daily
- Document as MedicationStatement

### Aspirin Use in Pregnancy
- Low-dose aspirin (81 mg) after 12 weeks for preeclampsia prevention in high-risk pregnancies
- Document as MedicationRequest

### Gestational Diabetes Screening
- All pregnant persons at 24 weeks
- LOINC 1554-5 (Glucose tolerance test)
