# USPSTF Recommendation Grades Reference

## Grade Definitions

| Grade | Meaning | Clinical Action |
|-------|---------|----------------|
| **A** | High certainty of substantial net benefit | Offer/provide this service |
| **B** | High certainty of moderate benefit OR moderate certainty of moderate-to-substantial benefit | Offer/provide this service |
| **C** | Moderate certainty of small net benefit | Offer selectively based on patient context |
| **D** | Moderate-to-high certainty of no benefit or harms outweigh benefits | Discourage this service |
| **I** | Insufficient evidence | Clinical judgment needed |

## Grade A Recommendations (Must Offer)

| Screening | Population | Interval | FHIR Check |
|-----------|-----------|----------|------------|
| Cervical cancer screening | Women 21-65 | Pap q3y (21-29); Pap+HPV q5y (30-65) | Procedure 171149006 or DiagnosticReport 10524-7 |
| Colorectal cancer screening | Adults 45-75 | Per test type (colonoscopy q10y, FIT q1y) | Procedure 73761001 or Observation 29771-3 |
| HIV screening | All 15-65 | At least once (more if high risk) | Observation 7917-8 |
| Hypertension screening | Adults 18+ | Annual/per visit | Observation 85354-9 |
| Rh(D) incompatibility screening | Pregnant | First prenatal visit + 24-28 weeks | Observation 10331-7 |
| Syphilis screening | Pregnant | First prenatal visit | Observation 20507-0 |
| Tobacco use counseling | All adults | Every visit | Observation 72166-2 |

## Grade B Recommendations (Should Offer)

| Screening | Population | Interval | FHIR Check |
|-----------|-----------|----------|------------|
| Breast cancer screening | Women 40-74 | Biennial mammography | DiagnosticReport 24606-6 or Procedure 71651007 |
| Depression screening | Adults 19+ | Annual | Observation 44249-1 (PHQ-9) |
| Anxiety screening | Adults | Per clinical judgment | Observation 70274-6 (GAD-7) |
| Diabetes/prediabetes screening | 35-70, overweight/obese | Every 3 years | Observation 4548-4 or 1558-6 |
| Hepatitis B screening | Adults 15-65 | One-time | Observation 5196-1 |
| Hepatitis C screening | Adults 18-79 | One-time | Observation 16128-1 |
| Lung cancer screening (LDCT) | 50-80, 20+ pack-year hx | Annual | DiagnosticReport or Procedure LDCT |
| Lipid screening (statin decision) | Adults 40-75 | ASCVD risk calculation | Observation 2093-3, 2085-9 |
| Osteoporosis screening (DEXA) | Women 65+ | Every 2 years | Procedure or DiagnosticReport DEXA |
| STI screening (CT/GC) | Sexually active women <=24 | Annual | Observation 43304-5, 43305-2 |
| Unhealthy alcohol use (AUDIT-C) | Adults 18+ | Annual | Observation 75626-2 |
| Unhealthy drug use | Adults 18+ | Annual | Observation 82667-7 |
| AAA screening | Men 65-75, ever smoked | One-time ultrasound | Procedure or DiagnosticReport |
| Aspirin for preeclampsia | High-risk pregnant | Start 12-28 weeks | MedicationRequest aspirin |
| Falls prevention | Adults 65+ | Annual assessment | Clinical assessment |
| Folic acid supplementation | Women planning/capable of pregnancy | Daily | MedicationRequest folic acid |
| Gestational diabetes screening | Pregnant, 24-28 weeks | One-time | Observation 1504-0 (GCT) |
| Intimate partner violence | Women of reproductive age | Per clinical judgment | Clinical screening |
| Statin for CVD prevention | 40-75 with CV risk factors | Based on ASCVD risk | MedicationRequest statin |

## Age-Based Quick Reference

### Ages 18-39
- [A] BP screening, HIV (one-time), tobacco counseling
- [B] Depression screening, anxiety screening, unhealthy alcohol/drug use screening
- [B] Hepatitis B (one-time), hepatitis C (one-time age 18+)
- [B] STI screening for women <=24

### Ages 40-49
- All of above plus:
- [A] Colorectal cancer screening starting at 45
- [B] Breast cancer screening starting at 40 (shared decision 40-49; routine 50+)
- [B] Diabetes screening if overweight (starting at 35)
- [B] Lipid screening/statin decision
- [B] Lung cancer LDCT (starting at 50 if 20+ pack-year)

### Ages 50-64
- All of above plus:
- [B] Lung cancer LDCT (50-80, 20+ pack-year)
- Cervical cancer screening through 65 (if adequate prior screening)
- Colorectal cancer screening through 75

### Ages 65+
- [A] Colorectal cancer screening through 75 (selectively 76-85)
- [B] Breast cancer screening through 74
- [B] Osteoporosis screening (women 65+)
- [B] AAA screening (men 65-75 who smoked)
- [B] Falls prevention assessment
- Pneumococcal vaccination (PCV20)
- Shingrix (age 50+)
- Annual flu, COVID-19
