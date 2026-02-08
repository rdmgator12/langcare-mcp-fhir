# Cancer Staging Systems Reference

## TNM Staging System (AJCC 8th Edition)

### Overview

The TNM system classifies cancer based on three components:
- **T** (Tumor): Size and extent of the primary tumor
- **N** (Nodes): Regional lymph node involvement
- **M** (Metastasis): Presence of distant metastasis

### T Stage (Primary Tumor)

| Stage | Description |
|-------|-------------|
| TX | Primary tumor cannot be assessed |
| T0 | No evidence of primary tumor |
| Tis | Carcinoma in situ (pre-invasive) |
| T1 | Small tumor, limited to organ of origin (typically <= 2 cm, varies by site) |
| T2 | Larger tumor or minimal extension (typically 2-5 cm, varies by site) |
| T3 | Larger tumor or extension into surrounding structures |
| T4 | Very large tumor or invasion of adjacent organs |
| T4a | Resectable T4 tumor |
| T4b | Unresectable T4 tumor |

### N Stage (Regional Lymph Nodes)

| Stage | Description |
|-------|-------------|
| NX | Regional lymph nodes cannot be assessed |
| N0 | No regional lymph node metastasis |
| N1 | Metastasis in 1-3 regional lymph nodes (varies by site) |
| N2 | Metastasis in 4-9 regional lymph nodes (varies by site) |
| N2a | Metastasis in 4-6 nodes |
| N2b | Metastasis in 7-9 nodes |
| N3 | Metastasis in >= 10 regional lymph nodes |

### M Stage (Distant Metastasis)

| Stage | Description |
|-------|-------------|
| M0 | No distant metastasis |
| M1 | Distant metastasis present |
| M1a | Metastasis to single distant organ |
| M1b | Metastasis to multiple distant organs |
| M1c | Peritoneal metastasis (for some cancers) |

### Overall Stage Grouping (General)

| Stage | TNM | 5-Year Survival (general) |
|-------|-----|--------------------------|
| 0 | Tis N0 M0 | >95% |
| I | T1-T2 N0 M0 | 70-90% |
| II | T2-T3 N0-N1 M0 | 50-80% |
| III | T3-T4 N1-N2 M0 | 20-60% |
| IV | Any T, Any N, M1 | 5-30% |

Survival varies significantly by cancer type.

## Site-Specific Staging (Top 10 Cancers)

### Breast Cancer

| Stage | TNM | Description |
|-------|-----|-------------|
| IA | T1 N0 M0 | Tumor <= 2cm, no nodes |
| IB | T0-T1 N1mi M0 | Micrometastasis in nodes |
| IIA | T0-T1 N1 M0 or T2 N0 M0 | Small tumor with nodes or medium tumor without |
| IIB | T2 N1 M0 or T3 N0 M0 | |
| IIIA | T0-T3 N2 M0 or T3 N1 M0 | |
| IIIB | T4 N0-N2 M0 | Chest wall or skin involvement |
| IIIC | Any T N3 M0 | Extensive nodal involvement |
| IV | Any T Any N M1 | Distant metastasis |

Biomarkers affecting staging and treatment:
- ER/PR status (hormone receptor)
- HER2 status (amplification)
- Ki-67 (proliferation index)
- Oncotype DX recurrence score (for ER+/HER2- node-negative)
- BRCA1/2 mutation status

### Lung Cancer (NSCLC)

| Stage | TNM | Description |
|-------|-----|-------------|
| IA1 | T1a N0 M0 | <= 1 cm |
| IA2 | T1b N0 M0 | >1 to <= 2 cm |
| IA3 | T1c N0 M0 | >2 to <= 3 cm |
| IB | T2a N0 M0 | >3 to <= 4 cm |
| IIA | T2b N0 M0 | >4 to <= 5 cm |
| IIB | T1-T2 N1 M0 or T3 N0 M0 | |
| IIIA | T1-T2 N2 M0 or T3-T4 N1 M0 | |
| IIIB | T1-T2 N3 M0 or T3-T4 N2 M0 | |
| IIIC | T3-T4 N3 M0 | |
| IVA | Any T Any N M1a-M1b | |
| IVB | Any T Any N M1c | |

Key biomarkers: EGFR, ALK, ROS1, BRAF V600E, PD-L1 TPS, KRAS G12C, MET, RET, NTRK, HER2

### Colorectal Cancer

| Stage | TNM | Description |
|-------|-----|-------------|
| 0 | Tis N0 M0 | In situ |
| I | T1-T2 N0 M0 | Into submucosa or muscularis |
| IIA | T3 N0 M0 | Through muscularis into subserosa |
| IIB | T4a N0 M0 | Penetrates visceral peritoneum |
| IIC | T4b N0 M0 | Invades adjacent organs |
| IIIA | T1-T2 N1 M0 or T1 N2a M0 | |
| IIIB | T3-T4a N1 M0 or T2-T3 N2a M0 or T1-T2 N2b M0 | |
| IIIC | T4a N2a M0 or T3-T4a N2b M0 or T4b N1-N2 M0 | |
| IVA | Any T Any N M1a | Single organ metastasis without peritoneal |
| IVB | Any T Any N M1b | >= 2 organs without peritoneal |
| IVC | Any T Any N M1c | Peritoneal metastasis |

Key biomarkers: MSI/MMR status, KRAS/NRAS, BRAF V600E, HER2

### Prostate Cancer

| Stage Group | TNM | Grade Group | PSA |
|------------|-----|-------------|-----|
| I | T1-T2a N0 M0 | 1 (Gleason <= 6) | < 10 |
| IIA | T1-T2 N0 M0 | 1 | >= 10 < 20 |
| IIB | T1-T2 N0 M0 | 2 (Gleason 3+4=7) | < 20 |
| IIC | T1-T2 N0 M0 | 3-4 (Gleason 4+3=7 or 8) | < 20 |
| IIIA | T1-T2 N0 M0 | 1-4 | >= 20 |
| IIIB | T3-T4 N0 M0 | Any | Any |
| IIIC | Any T N0 M0 | 5 (Gleason 9-10) | Any |
| IVA | Any T N1 M0 | Any | Any |
| IVB | Any T Any N M1 | Any | Any |

## ECOG Performance Status Scale

Source: Eastern Cooperative Oncology Group. LOINC: 89247-1.

| Grade | Description | Functional Capacity |
|-------|-------------|-------------------|
| 0 | Fully active, able to carry on all pre-disease performance without restriction | 100% |
| 1 | Restricted in physically strenuous activity but ambulatory and able to carry out work of a light or sedentary nature (e.g., light housework, office work) | ~80% |
| 2 | Ambulatory and capable of all self-care but unable to carry out any work activities; up and about more than 50% of waking hours | ~60% |
| 3 | Capable of only limited self-care; confined to bed or chair more than 50% of waking hours | ~40% |
| 4 | Completely disabled; cannot carry on any self-care; totally confined to bed or chair | ~20% |
| 5 | Dead | 0% |

### Clinical Significance

- ECOG 0-1: Generally eligible for most clinical trials and aggressive treatment
- ECOG 2: May be eligible for some clinical trials; treatment decisions individualized
- ECOG 3-4: Aggressive chemotherapy generally not recommended; focus on palliative care and symptom management
- ECOG is a strong independent prognostic factor across most cancer types

## Karnofsky Performance Status (KPS)

Source: Karnofsky DA et al. (1948). LOINC: 89262-0.

| Score | Description | Category |
|-------|-------------|----------|
| 100 | Normal, no complaints, no evidence of disease | Able to carry on normal activity; no special care needed |
| 90 | Able to carry on normal activity; minor signs or symptoms of disease | |
| 80 | Normal activity with effort; some signs or symptoms of disease | |
| 70 | Cares for self but unable to carry on normal activity or do active work | Unable to work; able to live at home; cares for most personal needs |
| 60 | Requires occasional assistance but is able to care for most personal needs | |
| 50 | Requires considerable assistance and frequent medical care | |
| 40 | Disabled; requires special care and assistance | Unable to care for self; requires institutional or hospital care |
| 30 | Severely disabled; hospitalization is indicated, although death not imminent | |
| 20 | Very sick; hospitalization and active supportive treatment necessary | |
| 10 | Moribund; fatal processes progressing rapidly | |
| 0 | Dead | |

### ECOG to KPS Approximate Conversion

| ECOG | KPS Range |
|------|-----------|
| 0 | 90-100 |
| 1 | 70-80 |
| 2 | 50-60 |
| 3 | 30-40 |
| 4 | 10-20 |

## Tumor Grading

### General Histologic Grade

| Grade | Differentiation | Description |
|-------|----------------|-------------|
| GX | Cannot be assessed | |
| G1 | Well differentiated | Cells look nearly normal; slow-growing |
| G2 | Moderately differentiated | Cells look somewhat abnormal |
| G3 | Poorly differentiated | Cells look very abnormal; faster-growing |
| G4 | Undifferentiated | Cells look nothing like normal tissue; most aggressive |

### Gleason Score (Prostate Cancer)

| Grade Group | Gleason Score | Pattern | Risk |
|------------|--------------|---------|------|
| 1 | <= 6 (3+3) | Well differentiated | Low |
| 2 | 7 (3+4) | Predominantly well differentiated | Favorable intermediate |
| 3 | 7 (4+3) | Predominantly poorly differentiated | Unfavorable intermediate |
| 4 | 8 (4+4, 3+5, 5+3) | Poorly differentiated | High |
| 5 | 9-10 (4+5, 5+4, 5+5) | Poorly/undifferentiated | Very high |

### Breast Cancer Grading (Nottingham/SBR)

| Grade | Score | Description |
|-------|-------|-------------|
| 1 (Low) | 3-5 | Well differentiated, slow-growing |
| 2 (Intermediate) | 6-7 | Moderately differentiated |
| 3 (High) | 8-9 | Poorly differentiated, aggressive |

Based on: tubule formation + nuclear pleomorphism + mitotic count (each scored 1-3).

## FHIR Staging LOINC Codes

| Concept | LOINC | Description |
|---------|-------|-------------|
| Stage group | 21908-9 | Overall stage (I, II, III, IV) |
| T category | 21905-5 | Clinical T stage |
| N category | 21906-3 | Clinical N stage |
| M category | 21907-1 | Clinical M stage |
| Pathologic T | 21899-0 | Pathologic T stage |
| Pathologic N | 21900-6 | Pathologic N stage |
| Pathologic M | 21901-4 | Pathologic M stage |
| Pathologic stage group | 21902-2 | Overall pathologic stage |
| ECOG | 89247-1 | ECOG performance status |
| KPS | 89262-0 | Karnofsky performance status |
