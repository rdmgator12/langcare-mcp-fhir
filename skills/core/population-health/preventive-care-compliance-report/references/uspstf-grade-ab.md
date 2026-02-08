# USPSTF Grade A and B Recommendations

Source: U.S. Preventive Services Task Force (USPSTF) current recommendations. Grade A = high certainty of substantial net benefit. Grade B = high certainty of moderate benefit or moderate certainty of substantial benefit.

---

## Cancer Screenings

### Breast Cancer Screening (Grade B)

**Population:** Women aged 40-74
**Recommendation:** Biennial screening mammography
**Interval:** Every 2 years
**Start age:** 40 (updated 2024 -- previously 50)
**Stop age:** 74 (insufficient evidence beyond 74)

**FHIR Evidence:**
- Procedure: SNOMED 71651007, 24623002
- DiagnosticReport: LOINC 24606-6
- CPT: 77067 (bilateral screening mammogram), 77063 (tomosynthesis)

**Exclusions from screening:**
- Prior bilateral mastectomy (SNOMED 27865001, ICD-10 Z90.13)
- BRCA carriers may follow different protocol (more frequent, include MRI)

### Cervical Cancer Screening (Grade A)

**Population:** Women aged 21-65

| Age Group | Test | Interval | LOINC |
|-----------|------|----------|-------|
| 21-29 | Pap smear alone | Every 3 years | 10524-7 (Pap), 19762-4 (Pap) |
| 30-65 | Pap + HPV co-test | Every 5 years | 10524-7 + 21440-3 (HPV) |
| 30-65 | HPV primary testing | Every 5 years | 21440-3, 77399-4 (HPV HR) |
| 30-65 | Pap alone | Every 3 years | 10524-7 |

**Stop screening:** Age > 65 with adequate prior screening (3 consecutive normal Paps or 2 consecutive normal co-tests in past 10 years).
**Not applicable:** Women with hysterectomy + cervix removal for non-cancer indication (SNOMED 236886002).

### Colorectal Cancer Screening (Grade A)

**Population:** Adults aged 45-75
**Shared decision (Grade C):** Ages 76-85

| Modality | Interval | FHIR Code |
|----------|----------|-----------|
| Colonoscopy | 10 years | SNOMED 73761001, CPT 45378 |
| FIT/gFOBT | 1 year | LOINC 29771-3, 57905-2, 27396-1 |
| FIT-DNA (Cologuard) | 3 years | LOINC 77353-1, CPT 81528 |
| CT Colonography | 5 years | SNOMED 418714002, CPT 74263 |
| Flexible Sigmoidoscopy | 5 years | SNOMED 44441009, CPT 45330 |
| Flex Sig + FIT | Sig every 10y + FIT annually | Combined |

### Lung Cancer Screening (Grade B)

**Population:** Adults aged 50-80 with >= 20 pack-year smoking history AND currently smoke or quit within past 15 years
**Test:** Annual low-dose CT (LDCT)
**LOINC:** 24604-1 (CT chest low dose)
**SNOMED:** 28163009 (CT of chest)
**CPT:** 71271

**FHIR Strategy:** Identify eligible population by:
1. Age 50-80
2. Smoking status: Observation LOINC 72166-2 (current or former smoker)
3. Pack-year history: Observation LOINC 8663-7 (cigarettes/day) or documented in Condition/social history

**Discontinue:** When person has not smoked for 15+ years or has limited life expectancy.

### Prostate Cancer Screening (Grade C -- Shared Decision)

**Population:** Men aged 55-69 (shared decision, not universal recommendation)
**Note:** NOT Grade A or B. Include only if practice tracks shared decision measures.

---

## Cardiovascular Prevention

### Hypertension Screening (Grade A)

**Population:** Adults >= 18 years
**Test:** Office blood pressure measurement
**Interval:** Annual (for normal BP); more frequent if elevated
**LOINC:** 85354-9 (BP panel), 8480-6 (systolic), 8462-4 (diastolic)
**Confirmation:** Elevated office BP should be confirmed with ambulatory or home monitoring before diagnosis.

### Statin Use for CVD Prevention (Grade B)

**Population:** Adults 40-75 with >= 1 CVD risk factor (dyslipidemia, diabetes, hypertension, smoking) AND estimated 10-year CVD risk >= 10%

**Evidence of compliance:**
- MedicationRequest for statin (active status)
- RxNorm: atorvastatin (83367), rosuvastatin (301542), simvastatin (36567), pravastatin (42463), lovastatin (6472)

**Risk factors (FHIR):**
- Dyslipidemia: Condition SNOMED 370992007; Observation LOINC 13457-7 (LDL >= 130) or 2093-3 (total cholesterol >= 200)
- Diabetes: Condition SNOMED 44054006
- Hypertension: Condition SNOMED 38341003
- Smoking: Observation LOINC 72166-2 (current smoker)

### Aspirin Use for CVD Prevention (Grade C -- Shared Decision)

**Note:** NOT Grade A or B for primary prevention. Downgraded in 2022. Only include if practice specifically tracks this.

### Abdominal Aortic Aneurysm Screening (Grade B)

**Population:** Men aged 65-75 who have ever smoked
**Test:** One-time ultrasound screening
**SNOMED:** 241462001 (Ultrasound of abdominal aorta)
**CPT:** 76706
**LOINC:** 24856-7 (Abdominal aorta US)

---

## Metabolic Screening

### Diabetes Screening (Grade B) -- Prediabetes and Type 2

**Population:** Adults 35-70 who are overweight or obese (BMI >= 25, or >= 23 for Asian Americans)
**Test:** Fasting glucose, HbA1c, or oral glucose tolerance test
**Interval:** Every 3 years if normal

**LOINC:**
- 4548-4 (HbA1c)
- 1558-6 (Fasting glucose)
- 1518-0 (2-hour glucose tolerance)

**Prediabetes thresholds:**
- A1c 5.7-6.4%
- Fasting glucose 100-125 mg/dL
- 2-hour glucose 140-199 mg/dL

### Lipid Disorders Screening

**Note:** USPSTF does not have a standalone lipid screening recommendation with a letter grade for the general adult population. However, lipid testing is integral to CVD risk assessment for statin therapy recommendations.

**Practical approach:**
- Men >= 35: Screen lipids every 5 years
- Women >= 45: Screen lipids every 5 years (sooner if risk factors)
- Men 20-35 and Women 20-44: Screen if increased CVD risk

**LOINC:** 2093-3 (Total cholesterol), 13457-7 (LDL-C), 2085-9 (HDL-C), 2571-8 (Triglycerides)

---

## Infectious Disease Screening

### HIV Screening (Grade A)

**Population:** Adolescents and adults 15-65
**Test:** At least once in lifetime; repeat screening for those at increased risk
**Interval:** At least once; annually or more for high-risk
**LOINC:** 75622-1 (HIV 1+2 Ag+Ab), 7918-6 (HIV-1 Ab), 80387-4 (HIV 1/2 combo)

**High-risk:** MSM, injection drug use, exchange sex for money/drugs, new sex partner with unknown HIV status.

### Hepatitis B Screening (Grade B)

**Population:** Adolescents and adults >= 15 years (universal screening, updated 2023)
**Test:** At least once; HBsAg, anti-HBs, and anti-HBc
**LOINC:** 5196-1 (HBsAg), 22322-2 (anti-HBs), 16933-4 (anti-HBc)

### Hepatitis C Screening (Grade B)

**Population:** Adults 18-79
**Test:** At least once; HCV antibody followed by confirmatory RNA if positive
**LOINC:** 16128-1 (HCV Ab), 11259-9 (HCV RNA)

### STI Screening (Grade B)

**Chlamydia and Gonorrhea:**
- Population: Sexually active women <= 24 and older women at increased risk
- LOINC: 43304-5 (Chlamydia NAAT), 43305-2 (Gonorrhea NAAT)
- Interval: Annually

**Syphilis:**
- Population: Persons at increased risk
- Grade A for pregnant women
- LOINC: 20507-0 (RPR), 22461-8 (Treponema Ab)

---

## Behavioral Health Screening

### Depression Screening (Grade B)

**Population:** Adults >= 18 years (Grade B); Adolescents 12-18 (Grade B)
**Tool:** PHQ-2 (screen) then PHQ-9 (if PHQ-2 positive)
**Interval:** Annual
**LOINC:** 55758-7 (PHQ-2), 44261-6 (PHQ-9), 89206-7 (PHQ-A for adolescents)

**Positive thresholds:**
- PHQ-2 >= 3: proceed to PHQ-9
- PHQ-9 >= 10: moderate depression, follow-up required

### Unhealthy Alcohol Use Screening (Grade B)

**Population:** Adults >= 18 years
**Tool:** AUDIT-C (abbreviated) or full AUDIT
**LOINC:** 75626-2 (AUDIT-C total), 75624-7 (AUDIT total)
**Interval:** Annual

### Unhealthy Drug Use Screening (Grade B)

**Population:** Adults >= 18 years
**Tool:** Single-question screening or NIDA Quick Screen
**Interval:** Annual or per clinical judgment

### Tobacco Use Screening and Cessation (Grade A)

**Population:** All adults
**Action:** Ask about tobacco use at every visit. For users, provide cessation interventions.
**LOINC:** 72166-2 (Tobacco smoking status)
**Cessation evidence:**
- Procedure: SNOMED 225323000 (Smoking cessation education)
- Medication: Varenicline (RxNorm 637190), Bupropion (RxNorm 42347), NRT (various)

---

## Other Screenings

### Obesity Screening (Grade B)

**Population:** Adults >= 18 years
**Measure:** BMI
**LOINC:** 39156-5 (BMI)
**Action:** Screen all adults. Offer or refer adults with BMI >= 30 to intensive behavioral interventions.

### Vision Screening -- Children (Grade B)

**Population:** Children 3-5 years
**Test:** Visual acuity testing
**LOINC:** 79880-1 (Visual acuity)

### Hearing Screening -- Newborns (Grade B)

**Population:** All newborns
**Test:** Before 1 month of age

### Preeclampsia Screening (Grade B)

**Population:** Pregnant women
**Test:** Blood pressure measurement at each prenatal visit

### Gestational Diabetes Screening (Grade B)

**Population:** Pregnant women at 24 weeks gestation (asymptomatic)
**Test:** 1-hour glucose challenge or 2-hour 75g OGTT

---

## Immunization-Related USPSTF Items

The USPSTF does not directly recommend vaccines (ACIP does), but preventive care compliance reports should include immunization status. Cross-reference with `immunization-status-checker` skill for detailed schedule.

Key immunizations to include in preventive care audit:
- Influenza (annual)
- COVID-19 (annual updated)
- Tdap/Td (every 10 years)
- Shingrix (>= 50 years)
- Pneumococcal (>= 65 or risk-based)
- HPV (through age 26)
