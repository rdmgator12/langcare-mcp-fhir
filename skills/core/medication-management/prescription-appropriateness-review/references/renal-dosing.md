# Renal Dosing Adjustments Reference

Common medications requiring dose adjustment based on estimated glomerular filtration rate (eGFR) in mL/min/1.73m2. Values derived from manufacturer prescribing information and clinical guidelines.

## eGFR Estimation

### CKD-EPI 2021 Equation (Race-Free)
LOINC 77147-7

```
eGFR = 142 x min(SCr/kappa, 1)^alpha x max(SCr/kappa, 1)^-1.200 x 0.9938^age x (1.012 if female)

Where:
  SCr = serum creatinine (mg/dL)
  kappa = 0.7 (female), 0.9 (male)
  alpha = -0.241 (female), -0.302 (male)
  min = minimum of SCr/kappa or 1
  max = maximum of SCr/kappa or 1
```

### CKD Stages

| Stage | eGFR (mL/min) | Description |
|-------|---------------|-------------|
| 1 | >= 90 | Normal or high |
| 2 | 60-89 | Mildly decreased |
| 3a | 45-59 | Mildly to moderately decreased |
| 3b | 30-44 | Moderately to severely decreased |
| 4 | 15-29 | Severely decreased |
| 5 | < 15 | Kidney failure |

---

## Cardiovascular Medications

### ACE Inhibitors

| Drug | Normal Dose | eGFR 30-59 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|------------|----------|
| Lisinopril | 10-40mg daily | No adjustment | Start 2.5-5mg, titrate cautiously | Start 2.5mg, monitor closely |
| Enalapril | 5-40mg daily | No adjustment | Start 2.5mg daily | Start 2.5mg daily |
| Ramipril | 2.5-20mg daily | No adjustment | Start 1.25mg daily | Start 1.25mg daily |
| Benazepril | 10-40mg daily | Start 5mg, max 40mg | Start 5mg, max 40mg | Not recommended |
| Captopril | 25-150mg TID | Reduce dose 25% | Reduce dose 50% | Reduce dose 75% |

### ARBs

| Drug | Normal Dose | eGFR 30-59 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|------------|----------|
| Losartan | 25-100mg daily | No adjustment | No adjustment | No adjustment |
| Valsartan | 80-320mg daily | No adjustment | No adjustment | Use with caution |
| Irbesartan | 150-300mg daily | No adjustment | No adjustment | No adjustment |
| Candesartan | 8-32mg daily | No adjustment | No adjustment | Use with caution |

### Anticoagulants

| Drug | Normal Dose | eGFR 30-49 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|------------|----------|
| Apixaban | 5mg BID | No adjustment (unless age >=80, weight <=60kg, or SCr >=1.5: reduce to 2.5mg BID) | 2.5mg BID | Avoid (limited data) |
| Rivaroxaban (AFib) | 20mg daily with food | 15mg daily with food | 15mg daily (use with caution) | Avoid |
| Dabigatran | 150mg BID | 150mg BID (some guidelines: 110mg BID) | Avoid | Contraindicated |
| Edoxaban | 60mg daily | 30mg daily (if CrCl 15-50) | 30mg daily | Avoid |
| Enoxaparin (treatment) | 1mg/kg BID | 1mg/kg BID (monitor anti-Xa) | 1mg/kg ONCE daily | Avoid |
| Fondaparinux | 2.5-10mg daily | Use with caution | Avoid | Contraindicated |

### Digoxin

| eGFR | Recommended Dose | Monitoring |
|------|-----------------|------------|
| >= 60 | 0.125-0.25mg daily | Trough level 0.5-0.9 ng/mL |
| 30-59 | 0.0625-0.125mg daily | Monitor levels closely |
| 15-29 | 0.0625mg daily or every other day | Monitor levels at each visit |
| < 15 | 0.0625mg every other day | Monitor levels, consider avoiding |

### Beta-Blockers

Most beta-blockers do not require renal dose adjustment. Exceptions:
| Drug | eGFR Threshold | Adjustment |
|------|---------------|------------|
| Atenolol | <35 | 25-50mg daily (max) |
| Nadolol | <30 | Increase dosing interval |
| Sotalol | <60: reduce dose; <40: extend interval; <10: avoid |
| Bisoprolol | <20 | Max 10mg daily |

---

## Diabetes Medications

| Drug | Normal Dose | eGFR 30-44 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|------------|----------|
| Metformin | 500-2000mg daily | Max 1000mg/day. Do not initiate if eGFR <30. | Discontinue | Contraindicated |
| Glipizide | 5-40mg daily | No adjustment | No adjustment | No adjustment (preferred SU) |
| Glyburide | 2.5-20mg daily | Avoid if eGFR <60 (active metabolites accumulate) | Avoid | Avoid |
| Sitagliptin | 100mg daily | 50mg daily | 25mg daily | 25mg daily |
| Saxagliptin | 5mg daily | 2.5mg daily | 2.5mg daily | 2.5mg daily |
| Linagliptin | 5mg daily | No adjustment | No adjustment | No adjustment |
| Empagliflozin | 10-25mg daily | 10mg if eGFR 20-44 (for CKD/HF benefit, not glycemic) | Discontinue if eGFR <20 | Contraindicated |
| Dapagliflozin | 5-10mg daily | 10mg if eGFR 20-44 (for CKD/HF benefit) | Discontinue if eGFR <20 | Contraindicated |
| Canagliflozin | 100-300mg daily | 100mg daily | Discontinue if eGFR <20 | Contraindicated |
| Exenatide (extended release) | 2mg weekly | Not recommended if eGFR <45 | Avoid | Avoid |
| Liraglutide | 0.6-1.8mg daily | No adjustment | No adjustment (limited data) | Not recommended |
| Dulaglutide | 0.75-4.5mg weekly | No adjustment | No adjustment | Not recommended |
| Semaglutide (oral) | 3-14mg daily | No adjustment | No adjustment | Not recommended |
| Insulin | Per protocol | Reduce dose 25% when eGFR <30 (decreased clearance, hypoglycemia risk) | Reduce dose 25-50% | Reduce dose 50%, monitor closely |

---

## Analgesics

### Opioids

| Drug | eGFR 30-59 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|----------|
| Morphine | Reduce dose 25-50%, extend interval | Reduce dose 50-75%, active metabolites accumulate | Avoid (M6G accumulation) |
| Hydromorphone | Reduce dose 25-50% | Reduce dose 50%, extend interval | Avoid or use with extreme caution |
| Oxycodone | Reduce dose 25-50%, start low | Reduce dose 50%, start 2.5mg | Avoid or use 2.5mg with monitoring |
| Hydrocodone | Reduce initial dose | Reduce dose 50% | Avoid |
| Fentanyl | No renal adjustment (hepatic metabolism) | No renal adjustment | No renal adjustment |
| Methadone | No renal adjustment (hepatic metabolism) | Use with caution | Use with caution |
| Tramadol IR | No adjustment | Max 100mg BID | Max 50mg BID |
| Tramadol ER | No adjustment | Avoid | Avoid |
| Codeine | Reduce dose 25% | Avoid (active metabolites) | Avoid |
| Meperidine | Avoid at all CKD stages | Avoid (normeperidine neurotoxicity) | Avoid |
| Tapentadol | No adjustment | Avoid | Avoid |

### NSAIDs

| eGFR | Recommendation |
|------|---------------|
| >= 60 | Use with caution, shortest duration, lowest dose |
| 30-59 | Avoid if possible. If necessary, short course (<5 days) with monitoring |
| < 30 | Contraindicated |

### Acetaminophen

No dose adjustment required for renal impairment. Preferred analgesic in CKD.

---

## Antimicrobials

| Drug | eGFR 30-59 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|----------|
| Amoxicillin | No adjustment | 250-500mg q12h | 250-500mg q24h |
| Amoxicillin-clavulanate | No adjustment (use 500/125 instead of 875/125 if eGFR <30) | 250-500/125mg q12h | 250-500/125mg q24h |
| Cephalexin | No adjustment | 250-500mg q8-12h | 250mg q12-24h |
| Ciprofloxacin | No adjustment | 250-500mg q12h | 250-500mg q18-24h |
| Levofloxacin | 750mg q48h or 500mg q24h | 500mg load then 250mg q24h | 500mg load then 250mg q48h |
| Trimethoprim-sulfamethoxazole | No adjustment | 50% dose | Avoid |
| Nitrofurantoin | Avoid if eGFR <30 (ineffective + neurotoxicity) | Contraindicated | Contraindicated |
| Vancomycin (oral) | No adjustment (minimal absorption) | No adjustment | No adjustment |
| Vancomycin (IV) | Dose per AUC-guided monitoring | Dose per AUC, extend interval | Dose per AUC |
| Acyclovir/Valacyclovir | Adjust per indication | Reduce dose significantly | Reduce dose significantly |
| Fluconazole | No adjustment | 50% dose | 50% dose |

---

## CNS / Psychiatric Medications

| Drug | eGFR 30-59 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|----------|
| Gabapentin | 200-700mg BID | 100-300mg daily | 100-300mg daily (post-HD dose on dialysis days) |
| Pregabalin | 75mg BID max | 25-75mg daily | 25-75mg daily |
| Levetiracetam | 500-1000mg BID | 250-500mg BID | 250-500mg daily |
| Lithium | Reduce dose, monitor levels q2-4wk | Reduce dose 50-75%, monitor closely | Generally avoid |
| Baclofen | Reduce dose | Avoid if possible (seizure risk) | Avoid |
| Duloxetine | No adjustment (if eGFR >30) | Avoid | Avoid |
| Memantine | No adjustment | 5mg BID | 5mg BID |

---

## Other Medications

| Drug | eGFR 30-59 | eGFR 15-29 | eGFR <15 |
|------|------------|------------|----------|
| Allopurinol | Start 100mg, max 200mg | Start 100mg, max 100-200mg | Start 50-100mg daily |
| Colchicine | 0.3mg daily or 0.6mg every other day | 0.3mg every other day | Avoid |
| Methotrexate | Reduce dose 50% if eGFR 30-59 | Avoid | Contraindicated |
| Spironolactone | Use with caution, monitor K+ | Avoid (hyperkalemia risk) | Contraindicated |
| Ranitidine/Famotidine | 50% dose | 25% dose or extend interval | 25% dose |

---

## Monitoring Requirements in CKD

For all renally adjusted medications, monitor:
1. **Serum creatinine and eGFR**: At baseline, 1 week after dose change, then every 3-6 months
2. **Serum potassium**: For ACEi/ARB/MRA - at baseline, 1 week, then every 3-6 months
3. **Drug levels**: For vancomycin (AUC), lithium, digoxin, aminoglycosides
4. **Signs of toxicity**: Specific to each drug class
5. **eGFR trajectory**: If declining, reassess all renally-cleared medications
