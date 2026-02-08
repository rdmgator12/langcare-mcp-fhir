# Oncology Treatment Reference

## Common Chemotherapy Regimens by Cancer Type

### Breast Cancer

| Regimen | Components | RxNorm Codes | Setting |
|---------|-----------|-------------|---------|
| AC | Doxorubicin + Cyclophosphamide | 3639, 3002 | Adjuvant |
| AC-T | AC followed by Paclitaxel | 3639, 3002, 56946 | Adjuvant |
| TC | Docetaxel + Cyclophosphamide | 72962, 3002 | Adjuvant |
| TCH | Docetaxel + Carboplatin + Trastuzumab | 72962, 40048, 224905 | HER2+ adjuvant |
| AC-THP | AC then Paclitaxel + Trastuzumab + Pertuzumab | 3639, 3002, 56946, 224905, 1298944 | HER2+ neoadjuvant |
| CMF | Cyclophosphamide + Methotrexate + 5-FU | 3002, 6851, 4492 | Adjuvant (older) |
| Capecitabine | Single agent | 194000 | Adjuvant (post-neoadjuvant residual) |
| T-DXd | Trastuzumab deruxtecan | 2377291 | HER2+ metastatic |
| CDK4/6i + AI | Palbociclib/Ribociclib/Abemaciclib + Letrozole/Anastrozole | 1601380/1873983/1946818 + 6572/1600 | HR+/HER2- metastatic |

### Lung Cancer (NSCLC)

| Regimen | Components | RxNorm | Setting |
|---------|-----------|--------|---------|
| Cisplatin + Pemetrexed | Cisplatin + Pemetrexed | 2555, 258494 | Non-squamous 1st line |
| Carboplatin + Paclitaxel | Carboplatin + Paclitaxel | 40048, 56946 | 1st line |
| Carboplatin + Pemetrexed + Pembrolizumab | 3 agents | 40048, 258494, 1547220 | Non-squamous 1st line |
| Osimertinib | Single agent TKI | 1860459 | EGFR-mutant 1st line |
| Alectinib | Single agent TKI | 1727455 | ALK-rearranged 1st line |
| Pembrolizumab | Single agent | 1547220 | PD-L1 >= 50% 1st line |
| Nivolumab + Ipilimumab | Immunotherapy combination | 1597876, 1094833 | 1st line (certain PD-L1) |

### Colorectal Cancer

| Regimen | Components | RxNorm | Setting |
|---------|-----------|--------|---------|
| FOLFOX | 5-FU + Leucovorin + Oxaliplatin | 4492, 6313, 258487 | Adjuvant Stage III, 1st line metastatic |
| FOLFIRI | 5-FU + Leucovorin + Irinotecan | 4492, 6313, 253455 | 2nd line or 1st line metastatic |
| FOLFOXIRI | 5-FU + Leucovorin + Oxaliplatin + Irinotecan | 4492, 6313, 258487, 253455 | Aggressive 1st line |
| CAPOX (XELOX) | Capecitabine + Oxaliplatin | 194000, 258487 | Adjuvant or 1st line |
| + Bevacizumab | Add to backbone regimen | 337535 | Metastatic (anti-VEGF) |
| + Cetuximab | Add to backbone regimen | 318341 | KRAS/NRAS wild-type metastatic |
| + Panitumumab | Add to backbone regimen | 583218 | KRAS/NRAS wild-type metastatic |
| Pembrolizumab | Single agent | 1547220 | MSI-H/dMMR 1st line metastatic |

### Prostate Cancer

| Regimen | Components | RxNorm | Setting |
|---------|-----------|--------|---------|
| ADT | Leuprolide or Goserelin | 6414, 4750 | All stages requiring hormonal therapy |
| Enzalutamide | Single agent | 1313988 | mCRPC |
| Abiraterone + Prednisone | Combination | 1100072, 8640 | mCRPC |
| Docetaxel + Prednisone | Combination | 72962, 8640 | mCRPC, mCSPC (high volume) |
| Cabazitaxel | Single agent | 996610 | mCRPC post-docetaxel |
| Darolutamide | Single agent | 2167236 | nmCRPC, mCSPC |
| Radium-223 | Radiopharmaceutical | 1371046 | mCRPC with bone metastases |

### Pancreatic Cancer

| Regimen | Components | RxNorm | Setting |
|---------|-----------|--------|---------|
| FOLFIRINOX | 5-FU + Leucovorin + Irinotecan + Oxaliplatin | 4492, 6313, 253455, 258487 | Adjuvant, 1st line metastatic (good PS) |
| mFOLFIRINOX | Modified doses | Same as above | Better tolerated adjuvant |
| Gemcitabine + nab-Paclitaxel | Combination | 72261, 583214 | 1st line metastatic |
| Gemcitabine | Single agent | 72261 | Adjuvant (if cannot tolerate FOLFIRINOX) |

## Tumor Marker Reference Values and Interpretation

### PSA (Prostate-Specific Antigen)

- LOINC: 2857-1 (total PSA)
- Normal: < 4.0 ng/mL (age-adjusted ranges vary)
- Gray zone: 4.0-10.0 ng/mL (biopsy consideration)
- Elevated: > 10.0 ng/mL (higher cancer probability)

**Post-treatment monitoring:**
- Post-prostatectomy: should be undetectable (< 0.1 ng/mL). Biochemical recurrence = PSA >= 0.2 ng/mL confirmed.
- Post-radiation: nadir + 2.0 ng/mL = biochemical recurrence (Phoenix definition)
- PSA doubling time < 3 months: aggressive biology, consider systemic therapy
- PSA velocity > 0.75 ng/mL/year: concerning

### CEA (Carcinoembryonic Antigen)

- LOINC: 10466-1
- Normal: < 5.0 ng/mL (non-smokers < 3.0 ng/mL)
- Used for: colorectal, gastric, pancreatic, lung, breast, ovarian

**Post-treatment monitoring (colorectal):**
- Check every 3 months for 2 years, then every 6 months for 3 more years
- Rising CEA after surgery: evaluate for recurrence (CT, PET-CT)
- Doubling pattern: concerning for recurrent disease
- Can be falsely elevated in smokers, liver disease, inflammatory bowel disease, hypothyroidism

### CA-125 (Cancer Antigen 125)

- LOINC: 19167-7
- Normal: < 35 U/mL
- Used for: ovarian, fallopian tube, peritoneal cancer

**Monitoring:**
- Nadir after treatment: lower nadir correlates with better prognosis
- Rising CA-125 (> 2x upper limit of normal): suggests recurrence
- Can be elevated in: endometriosis, PID, liver disease, pregnancy, menstruation, other cancers

### CA 19-9 (Cancer Antigen 19-9)

- LOINC: 24108-3
- Normal: < 37 U/mL
- Used for: pancreatic, biliary, gastric cancer

**Monitoring:**
- Baseline before treatment: prognostic value (>1000 U/mL suggests unresectable/metastatic)
- Post-surgery: should normalize; persistent elevation suggests residual disease
- Rising trend: suggests progression
- ~5-10% of population lacks Lewis antigen and cannot produce CA 19-9 (false negative)

### AFP (Alpha-Fetoprotein)

- LOINC: 1834-1
- Normal: < 10 ng/mL (adults)
- Used for: hepatocellular carcinoma, germ cell tumors

**Monitoring:**
- HCC surveillance: AFP + ultrasound every 6 months in at-risk patients (cirrhosis, chronic HBV)
- Germ cell tumors: AFP half-life ~5 days. Should normalize after treatment. Persistent elevation = residual tumor.
- Can be elevated in: pregnancy, hepatitis, cirrhosis (non-malignant)

### Beta-HCG (Human Chorionic Gonadotropin)

- LOINC: 21198-7
- Normal: < 5 mIU/mL (non-pregnant)
- Used for: germ cell tumors (testicular, ovarian), gestational trophoblastic disease

**Monitoring:**
- Half-life: 24-36 hours. Should normalize post-treatment.
- Persistent or rising: residual or recurrent disease
- Must exclude pregnancy before attributing to malignancy

### LDH (Lactate Dehydrogenase)

- LOINC: 2532-0
- Normal: 120-246 U/L (varies by lab)
- Used for: lymphoma, melanoma, germ cell tumors, general tumor burden marker

**Monitoring:**
- Non-specific but correlates with tumor burden
- Part of International Prognostic Index (IPI) for lymphoma
- Elevated in testicular cancer staging workup

## Treatment Response Criteria (RECIST 1.1)

### Target Lesion Measurement Rules

- Measure up to 5 target lesions total (max 2 per organ)
- Minimum size: 10mm longest diameter (15mm for lymph nodes short axis)
- Measure on CT or MRI with consistent technique
- Sum of longest diameters (SLD) for non-nodal targets + short axis for nodal targets

### Response Categories

| Category | Abbreviation | Definition |
|----------|-------------|------------|
| Complete Response | CR | Disappearance of all target lesions. All pathologic lymph nodes short axis < 10mm. |
| Partial Response | PR | >= 30% decrease in sum of diameters of target lesions (from baseline) |
| Progressive Disease | PD | >= 20% increase in sum of diameters from nadir AND absolute increase >= 5mm, OR appearance of new lesion(s) |
| Stable Disease | SD | Neither PR nor PD criteria met |

### Best Overall Response

| Time Point 1 | Time Point 2 | Best Overall Response |
|-------------|-------------|---------------------|
| CR | CR | CR |
| CR | PR | PR |
| PR | PR | PR |
| SD | SD | SD |
| SD | PD | SD |
| PD | Any | PD |

### Non-Target Lesion Assessment

| Response | Criteria |
|----------|---------|
| CR | Disappearance of all non-target lesions and normalization of tumor markers |
| Non-CR/Non-PD | Persistence of one or more non-target lesions and/or maintenance of tumor marker levels above normal |
| PD | Unequivocal progression of existing non-target lesions or appearance of new lesions |

## Supportive Care Requirements

### Antiemetic Protocols by Emetogenic Risk

| Emetogenic Risk | Regimens | Antiemetic Protocol |
|----------------|----------|-------------------|
| High (>90%) | Cisplatin, AC, FOLFIRINOX | NK1 antagonist (aprepitant/fosaprepitant) + 5-HT3 antagonist (ondansetron/palonosetron) + dexamethasone +/- olanzapine |
| Moderate (30-90%) | Carboplatin, oxaliplatin, irinotecan, doxorubicin | 5-HT3 antagonist + dexamethasone +/- NK1 antagonist |
| Low (10-30%) | Docetaxel, paclitaxel, gemcitabine, pemetrexed | Dexamethasone or 5-HT3 antagonist alone |
| Minimal (<10%) | Bevacizumab, cetuximab, immunotherapy | As needed only |

### Growth Factor Support

**G-CSF (filgrastim/pegfilgrastim) indications:**
- Primary prophylaxis: if febrile neutropenia risk >= 20% for the regimen
- Secondary prophylaxis: after an episode of febrile neutropenia
- Dose-dense regimens (e.g., dose-dense AC-T)
- Filgrastim (RxNorm 27400): 5 mcg/kg/day starting 24-72 hours after chemo
- Pegfilgrastim (RxNorm 286898): 6mg SC once per cycle, 24 hours after chemo

**ESA (erythropoiesis-stimulating agents):**
- Only for chemotherapy-associated anemia, NOT for curative-intent treatment
- Target Hgb: 10-12 g/dL (do not exceed 12)
- Check iron stores before starting

### Common Chemotherapy Toxicities to Monitor

| Toxicity | Agents | Monitoring |
|---------|--------|-----------|
| Cardiotoxicity | Doxorubicin, trastuzumab | Baseline and periodic echocardiogram (LVEF). Cumulative doxorubicin limit: 450-550 mg/m2 |
| Nephrotoxicity | Cisplatin | BMP before each cycle. Aggressive hydration. |
| Neurotoxicity | Oxaliplatin, paclitaxel, vincristine | Neuropathy assessment each cycle. Dose modify or stop if Grade 3+. |
| Pulmonary toxicity | Bleomycin, gemcitabine, immunotherapy | PFTs at baseline if bleomycin. Chest imaging if new dyspnea. |
| Hepatotoxicity | Methotrexate, immunotherapy | LFTs regularly. |
| Bone marrow suppression | Most cytotoxic agents | CBC before each cycle. Nadir typically day 7-14. |
| Immune-related adverse events | Checkpoint inhibitors (pembro, nivo, ipi) | TSH, LFTs, lipase, glucose at each visit. Monitor for colitis, pneumonitis, dermatitis, endocrinopathies. |
