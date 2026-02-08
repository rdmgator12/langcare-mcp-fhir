# ACS Cancer Screening Guidelines

Source: American Cancer Society (ACS) cancer screening guidelines. ACS recommendations may differ from USPSTF in start ages, intervals, and modalities. This reference provides the ACS-specific guidance for use when a practice follows ACS recommendations or when ACS and USPSTF both inform clinical decisions.

---

## Breast Cancer

### Average Risk (ACS 2024)

| Age | Recommendation | Frequency |
|-----|---------------|-----------|
| 40-44 | May begin annual mammography (optional) | Annual if chosen |
| 45-54 | Should have annual mammography | Annual |
| >= 55 | May switch to biennial or continue annual | Biennial or annual |
| No upper age limit | Continue as long as life expectancy >= 10 years | -- |

**Key differences from USPSTF:**
- ACS recommends annual for 45-54 (USPSTF says biennial starting at 40)
- ACS has no upper age limit (USPSTF stops at 74)

### High Risk (ACS)

**High-risk criteria (lifetime risk >= 20%):**
- Known BRCA1/BRCA2 mutation carrier
- First-degree relative of BRCA carrier (untested)
- Chest radiation between ages 10-30 (e.g., Hodgkin lymphoma treatment)
- Li-Fraumeni syndrome, Cowden syndrome, Bannayan-Riley-Ruvalcaba syndrome
- Lifetime risk >= 20% by validated risk model (Tyrer-Cuzick, BRCAPRO, etc.)

**Screening for high-risk:**
- Annual mammography AND annual breast MRI, starting at age 30 (or 8 years after chest radiation, whichever is later)
- Do not start before age 25

**FHIR identification of high-risk:**
- FamilyMemberHistory with breast cancer in first-degree relative
- Condition: SNOMED 254837009 (BRCA mutation carrier), 716284001 (BRCA1), 716285000 (BRCA2)
- History of chest radiation: Procedure with SNOMED 108290001 (radiation therapy)

---

## Colorectal Cancer

### Average Risk (ACS 2018)

**Start age:** 45 (ACS aligned with USPSTF since 2018)
**End age:** 75 (individual decision 76-85)

| Modality | Interval | Notes |
|----------|----------|-------|
| Colonoscopy | 10 years | Preferred structural exam |
| CT colonography | 5 years | Structural exam |
| Flexible sigmoidoscopy | 5 years | Structural exam |
| FIT (immunochemical) | 1 year | Stool-based, preferred stool test |
| gFOBT (guaiac) | 1 year | Stool-based |
| FIT-DNA (Cologuard) | 3 years | Stool-based |

**ACS preference:** ACS prefers stool-based testing with FIT annually for resource-limited settings. For structural exams, colonoscopy is preferred.

### High Risk -- Earlier/More Frequent Screening

| Risk Factor | Start Age | Interval | Modality |
|-------------|-----------|----------|----------|
| First-degree relative with CRC < 60 or adenoma < 60 | Age 40 or 10 years before youngest case | Every 5 years | Colonoscopy |
| First-degree relative with CRC >= 60 | Age 40 | Per average risk | Per average risk |
| 2+ first-degree relatives with CRC (any age) | Age 40 or 10 years before youngest case | Every 5 years | Colonoscopy |
| Personal history of adenomatous polyps | Per surveillance guidelines | 3-5 years | Colonoscopy |
| Personal history of CRC | 1 year after resection | Then per findings | Colonoscopy |
| Inflammatory bowel disease (UC/Crohn's colitis) | 8 years after symptom onset | Every 1-2 years | Colonoscopy with biopsies |
| Lynch syndrome (HNPCC) | Age 20-25 or 2-5 years before earliest family case | Every 1-2 years | Colonoscopy |
| FAP | Age 10-12 | Annual until colectomy | Flexible sigmoidoscopy |

**FHIR identification of high-risk:**
- FamilyMemberHistory: condition code for colorectal cancer (SNOMED 363406005, ICD-10 C18-C20) in first-degree relative
- Condition: SNOMED 64766004 (Ulcerative colitis), 34000006 (Crohn's disease)
- Condition: SNOMED 716318002 (Lynch syndrome), 81528003 (FAP)

---

## Cervical Cancer

### ACS 2020 Guideline

| Age | Primary Test | Interval |
|-----|-------------|----------|
| < 25 | No screening | -- |
| 25-65 | HPV primary test (preferred) | Every 5 years |
| 25-65 | HPV + cytology co-test (acceptable) | Every 5 years |
| 25-65 | Cytology alone (acceptable if HPV unavailable) | Every 3 years |
| > 65 | No screening if adequate prior | -- |

**Key difference from USPSTF:** ACS starts at 25 (USPSTF starts Pap at 21). ACS prefers HPV primary testing.

**FHIR codes:**
- LOINC 21440-3 (HPV DNA)
- LOINC 77399-4 (HPV high-risk)
- LOINC 10524-7 (Pap smear)
- LOINC 19762-4 (Pap thin prep)

### Exceptions (continue screening)

- History of CIN2/CIN3/AIS: Continue screening for 25 years after treatment
- Immunocompromised (HIV, organ transplant): Per ASCCP guidelines (more frequent)
- DES exposure in utero: Continue screening per provider judgment

---

## Lung Cancer

### ACS 2024

**Population:** Adults 50-80 with >= 20 pack-year history who currently smoke or quit within past 15 years
**Test:** Annual low-dose CT (LDCT)
**Notes:** Aligns with USPSTF 2021. ACS emphasizes shared decision-making including discussion of benefits, limitations, harms (false positives, radiation exposure).

**FHIR codes:**
- LOINC 24604-1 (CT chest low dose)
- CPT 71271
- Smoking: Observation LOINC 72166-2

---

## Prostate Cancer

### ACS 2024

**Population:** Men with >= 10 year life expectancy

| Risk Group | Discussion Start Age |
|------------|---------------------|
| Average risk | 50 |
| High risk (African American, first-degree relative with prostate cancer < 65) | 45 |
| Very high risk (multiple first-degree relatives with prostate cancer < 65) | 40 |

**Test:** PSA blood test with or without digital rectal exam (DRE)
**Interval:** If PSA < 2.5: rescreen every 2 years. If PSA >= 2.5: rescreen annually.

**FHIR codes:**
- LOINC 2857-1 (PSA)
- Procedure: SNOMED 410006001 (DRE)

**Note:** USPSTF gives Grade C (shared decision for ages 55-69). ACS provides more specific age-based guidance by risk group.

---

## Endometrial Cancer

ACS does not recommend routine screening for endometrial cancer in average-risk women. However:
- At menopause, inform all women about risks/symptoms of endometrial cancer
- Report unexpected vaginal bleeding promptly

**High risk (Lynch syndrome):** Consider annual endometrial biopsy starting at age 35.

---

## Screening Matrix by Age and Sex

### Females

| Age | Breast | Cervical | Colorectal | Lung |
|-----|--------|----------|------------|------|
| 21-24 | -- | Pap q3y (USPSTF) | -- | -- |
| 25-39 | -- | HPV q5y (ACS) / Pap q3y (USPSTF) | -- | -- |
| 40-44 | Mammo optional annual (ACS) / Biennial (USPSTF) | HPV/Pap per above | -- | -- |
| 45-49 | Mammo annual (ACS) / Biennial (USPSTF) | HPV/Pap per above | Colonoscopy or FIT (start 45) | -- |
| 50-54 | Mammo annual (ACS) / Biennial (USPSTF) | HPV/Pap per above | Per interval | LDCT if smoking hx |
| 55-64 | Mammo biennial or annual (ACS) / Biennial (USPSTF) | HPV/Pap per above | Per interval | LDCT if smoking hx |
| 65-74 | Mammo biennial or annual | Stop if adequate prior (ACS/USPSTF) | Per interval | LDCT if smoking hx |
| 75-80 | Continue if life exp >= 10y (ACS) / Stop at 74 (USPSTF) | -- | Individual decision 76-85 | LDCT if smoking hx (to 80) |

### Males

| Age | Prostate | Colorectal | Lung |
|-----|----------|------------|------|
| 40-44 | PSA at 40 if very high risk | -- | -- |
| 45-49 | PSA at 45 if high risk | Colonoscopy or FIT (start 45) | -- |
| 50-54 | PSA at 50 if average risk | Per interval | LDCT if smoking hx |
| 55-69 | PSA per interval (ACS) / Shared decision (USPSTF) | Per interval | LDCT if smoking hx |
| 70-74 | Individual decision | Per interval | LDCT if smoking hx |
| 75-80 | Generally stop | Individual decision 76-85 | LDCT if smoking hx (to 80) |

---

## Special Populations

### Immunocompromised (HIV, transplant)

- Cervical: Screen more frequently (cytology annually, co-test may be appropriate)
- Anal cancer: Consider anal Pap for HIV+ MSM and women with cervical dysplasia (no formal ACS recommendation, but ANCHOR trial supports screening)
- Skin cancer: Annual dermatologic exam for transplant recipients

### Prior Cancer Survivors

- Survivors of childhood cancer with chest radiation: Breast MRI + mammography starting 8 years post-radiation or age 25 (whichever is later)
- Prior colorectal cancer: Surveillance colonoscopy per USMSTF guidelines (1 year, then 3 years, then 5 years)
- Prior cervical precancer (CIN2+): Continue screening for 25 years post-treatment

### Family History Impact on Screening

| Family History | Affected Screening | Modification |
|---------------|-------------------|-------------|
| First-degree relative with breast cancer | Breast | Consider starting 10 years before relative's age at diagnosis (not before 30) |
| First-degree relative with CRC < 60 | Colorectal | Start at 40 or 10 years before case, every 5 years |
| First-degree relative with prostate cancer < 65 | Prostate (shared decision) | Discuss PSA at 45 |
| Lynch syndrome | Colorectal, endometrial | Colonoscopy q1-2y from 20-25; endometrial biopsy q1y from 35 |
| BRCA1/2 carrier | Breast, ovarian | MRI + mammo annually from 25-30; discuss risk-reducing surgery |
