# CDC/ACIP Immunization Schedule Reference

## Adult Immunization Schedule (19+ Years)

### Recommended for All Adults

| Vaccine | Age 19-26 | Age 27-49 | Age 50-64 | Age 65+ | Series/Interval |
|---------|----------|----------|----------|---------|-----------------|
| Influenza | Annual | Annual | Annual | Annual | 1 dose annually (HD or adjuvanted for 65+) |
| COVID-19 | Per current CDC guidance | Per current CDC guidance | Per current CDC guidance | Per current CDC guidance | Updated vaccine annually |
| Td/Tdap | 1 dose Tdap, then Td every 10 years | 1 dose Tdap if not received, then Td q10y | Td/Tdap q10y | Td/Tdap q10y | Tdap once, then Td booster q10y |
| MMR | 1-2 doses if no evidence of immunity | 1 dose if no evidence of immunity | -- | -- | Born before 1957 generally immune |
| Varicella | 2 doses if no evidence of immunity | 2 doses if no evidence of immunity | -- | -- | 4-8 weeks apart |
| HPV | 2-3 doses through age 26 | Shared decision 27-45 | -- | -- | 2 doses if started <15; 3 doses if 15+ |
| Zoster (Shingrix) | -- | -- | 50+ | 2 doses | 2 doses, 2-6 months apart |
| Pneumococcal | Risk-based | Risk-based | Risk-based | PCV20 or PCV15+PPSV23 | See pneumococcal section |
| Hepatitis B | 2-3 doses if not vaccinated | 2-3 doses if not vaccinated | 2-3 doses if not vaccinated | Risk-based | 3-dose (0, 1, 6 months) or 2-dose (Heplisav-B) |

### Pneumococcal Vaccine Schedule (2023+)

| Population | Recommendation |
|-----------|----------------|
| Age 65+ (PCV-naive) | 1 dose PCV20 **OR** 1 dose PCV15 followed by PPSV23 (1+ year later) |
| Age 19-64 with risk factors | Same as above |
| Previously received PPSV23 only | 1 dose PCV20 (at least 1 year after PPSV23) |

Risk factors: immunocompromising conditions, CSF leak, cochlear implant, asplenia, chronic heart/lung/liver disease, diabetes, alcoholism, smoking.

### Risk-Based Vaccines

| Vaccine | Risk Conditions | CVX Code |
|---------|----------------|----------|
| Hepatitis A | Chronic liver disease, travel, MSM, drug use, homelessness | 83 |
| Hepatitis B | CKD, diabetes (19-59), HIV, hepatitis C, multiple partners, MSM | 43 |
| Meningococcal ACWY | Asplenia, complement deficiency, HIV, travel, college freshmen in dorms | 114 |
| Meningococcal B | Asplenia, complement deficiency, outbreak | 162 |
| Hib | Asplenia, HSCT | 17 |

## Pregnancy-Specific Recommendations

| Vaccine | Timing | Notes |
|---------|--------|-------|
| Tdap | 27-36 weeks each pregnancy | Passive immunity transfer to newborn |
| Influenza | Any trimester during flu season | Inactivated only |
| COVID-19 | Any trimester | Updated vaccine |
| RSV (Abrysvo) | 32-36 weeks (Sep-Jan) | OR infant receives nirsevimab instead |
| **AVOID** | All live vaccines: MMR, varicella, LAIV, yellow fever | Contraindicated in pregnancy |

## Common CVX Codes for FHIR Immunization.vaccineCode

| Vaccine | CVX Code | Display |
|---------|----------|---------|
| Influenza (IIV4) | 197 | Influenza, inactivated, quadrivalent |
| Influenza (HD) | 153 | Influenza, high-dose |
| Tdap | 115 | Tetanus, diphtheria, acellular pertussis |
| Td | 139 | Td (adult) |
| MMR | 03 | Measles, mumps, rubella |
| Varicella | 21 | Varicella |
| HPV (9-valent) | 165 | HPV9 |
| Shingrix | 187 | Zoster recombinant |
| PCV20 | 216 | Pneumococcal conjugate PCV20 |
| PCV15 | 215 | Pneumococcal conjugate PCV15 |
| PPSV23 | 33 | Pneumococcal polysaccharide PPV23 |
| Hepatitis A | 83 | Hep A, adult |
| Hepatitis B (3-dose) | 43 | Hep B, adult |
| Hepatitis B (Heplisav-B) | 189 | HepB, CpG adjuvant |
| Meningococcal ACWY | 114 | MenACWY |
| Meningococcal B | 162 | MenB |
| COVID-19 mRNA | 228 | COVID-19, mRNA, updated |
| RSV (Abrysvo) | 305 | RSV, maternal |

## Catch-Up Schedule Notes

- Adults without documentation of childhood vaccines should be evaluated as unvaccinated
- MMR: 1 dose for adults born 1957+; 2 doses for healthcare workers, students, international travelers
- Varicella: 2 doses for adults without evidence of immunity (born in US after 1980 without hx of disease)
- Hepatitis B: universal screening recommended (HBsAg, anti-HBs, anti-HBc); vaccinate if non-immune
