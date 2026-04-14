# USPSTF A and B Grade Recommendations Reference

## Cancer Screening

| Screening | Population | Interval | Grade | FHIR Check |
|-----------|-----------|----------|-------|------------|
| Breast cancer (mammography) | Women 50-74 | Every 2 years | B | DiagnosticReport code=24606-6 or Procedure code=71651007 |
| Cervical cancer (Pap/HPV) | Women 21-65 | Pap every 3 years (21-29); Pap+HPV every 5 years (30-65) | A | Procedure code=171149006 or DiagnosticReport code=10524-7 |
| Colorectal cancer | Adults 45-75 | Colonoscopy q10y OR FIT q1y OR FIT-DNA q3y OR flex sig q5y | A | Procedure code=73761001 (colonoscopy) or Observation code=29771-3 (FOBT) |
| Lung cancer (LDCT) | Adults 50-80 with 20+ pack-year hx | Annual | B | Procedure code=LDCT or DiagnosticReport |

## Cardiovascular

| Screening | Population | Interval | Grade | FHIR Check |
|-----------|-----------|----------|-------|------------|
| Hypertension screening | Adults 18+ | Annual (or per visit) | A | Observation code=85354-9 |
| Lipid screening (statin decision) | Adults 40-75 | ASCVD risk calc to guide statin | B | Observation code=2093-3, 2085-9 |
| Abdominal aortic aneurysm | Men 65-75 who ever smoked | One-time ultrasound | B | Procedure or DiagnosticReport |

## Metabolic

| Screening | Population | Interval | Grade | FHIR Check |
|-----------|-----------|----------|-------|------------|
| Prediabetes/T2DM | Adults 35-70, overweight/obese | Every 3 years | B | Observation code=4548-4 (A1c) or 1558-6 (fasting glucose) |

## Infectious Disease

| Screening | Population | Interval | Grade | FHIR Check |
|-----------|-----------|----------|-------|------------|
| Hepatitis B (HBsAg) | Adolescents/adults 15-65 | One-time | B | Observation code=5196-1 |
| Hepatitis C (anti-HCV) | Adults 18-79 | One-time | B | Observation code=16128-1 |
| HIV | Adults 15-65 | At least once; more if high risk | A | Observation code=7917-8 |
| Syphilis | Pregnant women; high-risk adults | Per risk | A | Observation code=20507-0 |
| Chlamydia/Gonorrhea | Sexually active women <=24 and high-risk | Annual | B | Observation code=43304-5, 43305-2 |

## Mental Health

| Screening | Population | Interval | Grade | FHIR Check |
|-----------|-----------|----------|-------|------------|
| Depression (PHQ-9) | Adults 19+ | Annual | B | Observation code=44249-1 (PHQ-9 total) |
| Anxiety | Adults | Unclear optimal interval | B | Observation code=70274-6 (GAD-7 total) |
| Unhealthy alcohol use (AUDIT-C) | Adults 18+ | Annual | B | Observation code=75626-2 (AUDIT-C) |
| Unhealthy drug use | Adults 18+ | Annual | B | Observation code=82667-7 |

## Pregnancy

| Screening | Population | Interval | Grade |
|-----------|-----------|----------|-------|
| Gestational diabetes | Pregnant, 24-28 weeks | One-time | B |
| Preeclampsia (aspirin prophylaxis) | High-risk pregnant | Low-dose aspirin 12-28 weeks | B |
| Rh(D) incompatibility | Pregnant, first visit + 24-28 weeks | Per pregnancy | A |

## Other

| Screening | Population | Interval | Grade | FHIR Check |
|-----------|-----------|----------|-------|------------|
| Osteoporosis (DEXA) | Women 65+ (or 50-64 with risk factors) | Every 2 years | B | Procedure or DiagnosticReport |
| Falls prevention | Adults 65+ | Annual | B | Observation or clinical assessment |
| Vision screening | Children 3-5 | One-time | B | -- |
| Tobacco use counseling | All adults | Every visit | A | Observation code=72166-2 |

## Age-Based Summary (Quick Reference)

### Ages 18-39
- BP screening, depression screening, HIV (one-time), STI screening (if applicable), tobacco screening

### Ages 40-49
- Add: diabetes screening (if overweight), hepatitis C (one-time), lung cancer CT (if 20+ pack-year smoker age 50+)
- Colorectal cancer screening starting at 45

### Ages 50-64
- Add: colorectal cancer screening, breast cancer screening (mammogram), cervical cancer screening (Pap+HPV)
- Lung cancer LDCT (if 20+ pack-year history, current or quit within 15 years)

### Ages 65-75
- Add: AAA screening (men who smoked), osteoporosis screening (women), falls prevention
- Continue: colorectal, breast (to 74), cervical (to 65 if adequate prior screening)
- Consider stopping: cervical cancer screening at 65 if adequate prior screening

### Ages 75+
- Individualize screening based on life expectancy and comorbidities
- Stop: colorectal (age 76-85 individualize, stop at 85), breast (stop at 75 per USPSTF), lung (stop at 80)
