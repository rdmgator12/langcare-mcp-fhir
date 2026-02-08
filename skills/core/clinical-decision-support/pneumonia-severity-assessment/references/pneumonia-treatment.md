# Pneumonia Treatment and Management Reference

## Disposition Recommendations by Severity

### Community-Acquired Pneumonia (CAP)

| Severity Level | CURB-65 | PSI Class | Disposition | Setting |
|---------------|---------|-----------|-------------|---------|
| Mild | 0-1 | I-II | Outpatient | Home with oral antibiotics |
| Moderate | 2 | III | Observation or short inpatient stay | ED observation unit or general ward (24-72h) |
| Moderate-severe | 3 | IV | Inpatient | General ward with telemetry |
| Severe | 4-5 | V | ICU | ICU or step-down unit |
| Severe (ATS/IDSA criteria met) | Any | Any | ICU | Regardless of CURB-65/PSI score |

### Criteria for Safe Outpatient Treatment

All of the following must be met:
- CURB-65 <= 1 and PSI Class I-II
- SpO2 >= 92% on room air
- Able to tolerate oral medications
- Reliable follow-up within 48-72 hours
- No unstable comorbidities
- Adequate home support
- No complications (empyema, lung abscess, large effusion)

---

## Empiric Antibiotic Selection: CAP (IDSA/ATS 2019 Guidelines)

### Outpatient, No Comorbidities

| Regimen | Dose | Duration | Notes |
|---------|------|----------|-------|
| Amoxicillin | 1g PO TID | 5 days (minimum) | First-line per 2019 guidelines |
| Doxycycline | 100mg PO BID | 5 days | Alternative. Covers atypicals. |
| Azithromycin | 500mg day 1, then 250mg days 2-5 | 5 days | ONLY if local macrolide resistance <25%. Not first-line in most US regions. |

### Outpatient, With Comorbidities

Comorbidities: chronic heart, lung, liver, or renal disease; diabetes; alcoholism; malignancy; asplenia; immunosuppression.

| Regimen | Dose | Duration | Notes |
|---------|------|----------|-------|
| Amoxicillin-clavulanate 875/125mg PO BID + Macrolide | Amox-clav + Azithromycin or Doxycycline | 5 days | Preferred combination |
| Cefpodoxime 200mg PO BID + Macrolide | Cefpodoxime + Azithromycin or Doxycycline | 5 days | Alternative beta-lactam |
| Respiratory fluoroquinolone monotherapy | Levofloxacin 750mg PO daily OR Moxifloxacin 400mg PO daily | 5 days | Reserve for patients who cannot take beta-lactam + macrolide. Risk of C. diff, tendinopathy. |

### Inpatient, Non-Severe (General Ward)

| Regimen | Dose | Duration | Notes |
|---------|------|----------|-------|
| Ceftriaxone 2g IV daily + Azithromycin 500mg IV/PO daily | As listed | 5-7 days | First-line. Switch to oral when clinically stable x 24-48h. |
| Ceftriaxone 2g IV daily + Doxycycline 100mg IV/PO BID | As listed | 5-7 days | Alternative if macrolide-intolerant |
| Ampicillin-sulbactam 3g IV q6h + Azithromycin 500mg IV/PO daily | As listed | 5-7 days | Alternative beta-lactam |
| Levofloxacin 750mg IV daily | As listed | 5-7 days | Fluoroquinolone monotherapy (if beta-lactam allergy). |

### Inpatient, Severe (ICU)

| Regimen | Dose | Duration | Notes |
|---------|------|----------|-------|
| Ceftriaxone 2g IV daily + Azithromycin 500mg IV daily | As listed | 7 days | Standard severe CAP regimen |
| Ceftriaxone 2g IV daily + Levofloxacin 750mg IV daily | As listed | 7 days | Alternative to macrolide |

#### Add-On Coverage for Risk Factors

| Risk Factor | Additional Agent | Criteria |
|-------------|-----------------|----------|
| MRSA risk | Vancomycin 15-20mg/kg IV q8-12h OR Linezolid 600mg IV/PO q12h | Prior MRSA isolation, recent hospitalization, cavitary infiltrate, empyema, post-influenza pneumonia |
| Pseudomonas risk | Replace ceftriaxone with: Piperacillin-tazobactam 4.5g IV q6h OR Cefepime 2g IV q8h OR Meropenem 1g IV q8h | Prior Pseudomonas isolation, structural lung disease, immunocompromised, recent broad-spectrum antibiotics |

### Beta-Lactam Allergy

| Allergy Type | Outpatient | Inpatient |
|-------------|-----------|-----------|
| Non-severe (rash, GI) | Respiratory fluoroquinolone | Respiratory fluoroquinolone OR aztreonam + macrolide |
| Severe (anaphylaxis) | Respiratory fluoroquinolone | Respiratory fluoroquinolone OR aztreonam + macrolide |

---

## Atypical Pathogen Coverage Criteria

### When to Cover for Atypical Pathogens

Atypical pathogens: Mycoplasma pneumoniae, Chlamydophila pneumoniae, Legionella pneumophila.

| Setting | Cover Atypicals? | Rationale |
|---------|-----------------|-----------|
| Outpatient CAP | Yes (always) | Atypicals account for 10-30% of CAP. Macrolide or doxycycline provides coverage. |
| Inpatient, non-severe | Yes (always) | Beta-lactam + macrolide/doxycycline is standard dual therapy. |
| Inpatient, severe (ICU) | Yes (always) | Beta-lactam + macrolide or fluoroquinolone. |
| HAP | No (usually not) | Atypicals uncommon in HAP. Exception: Legionella in hospital water outbreaks. |
| VAP | No | Focus on gram-negatives and MRSA. |

### Legionella-Specific Considerations

Test for Legionella when:
- Severe CAP requiring ICU
- Hyponatremia (Na <130)
- GI symptoms (diarrhea, nausea)
- Neurologic symptoms (confusion, headache)
- Recent travel or hotel/cruise ship exposure
- Institutional/hospital outbreak suspected

Testing: Legionella urinary antigen (detects serogroup 1 only, ~80% of cases). Also send Legionella culture on respiratory specimen. LOINC for urinary antigen: 32172-9.

---

## Hospital-Acquired Pneumonia (HAP) Antibiotic Selection

HAP = pneumonia occurring >= 48 hours after hospital admission in non-intubated patients.

### Risk Assessment for MDR Pathogens

Risk factors for MDR HAP:
- Prior IV antibiotic use within 90 days
- Hospitalization >= 5 days before pneumonia onset
- Septic shock at time of HAP diagnosis
- ARDS preceding HAP
- Renal replacement therapy before HAP
- Local unit with >10-20% MRSA or >10% MDR gram-negative prevalence

### HAP Treatment

| Risk Category | Regimen | Duration |
|---------------|---------|----------|
| No MDR risk factors, not severe, not high mortality risk | Piperacillin-tazobactam 4.5g IV q6h | 7 days |
| | OR Cefepime 2g IV q8h | 7 days |
| | OR Levofloxacin 750mg IV daily | 7 days |
| MDR risk factors OR high severity | Anti-pseudomonal beta-lactam + MRSA coverage | 7 days |
| | Piperacillin-tazobactam + Vancomycin | |
| | OR Cefepime + Vancomycin (or Linezolid) | |
| | Consider adding aminoglycoside or fluoroquinolone for dual gram-negative coverage if prior antibiotics in 90 days | |

---

## Ventilator-Associated Pneumonia (VAP) Antibiotic Selection

VAP = pneumonia occurring >= 48 hours after endotracheal intubation.

### VAP Treatment

| Component | Agents | Notes |
|-----------|--------|-------|
| Anti-pseudomonal beta-lactam | Piperacillin-tazobactam 4.5g IV q6h OR Cefepime 2g IV q8h OR Meropenem 1g IV q8h | Always include one anti-pseudomonal agent |
| MRSA coverage | Vancomycin 15-20mg/kg IV q8-12h OR Linezolid 600mg IV q12h | Add if MRSA risk factors or local MRSA prevalence >10-20% |
| Second anti-pseudomonal agent | Aminoglycoside (tobramycin 5-7mg/kg IV daily) OR Fluoroquinolone (ciprofloxacin 400mg IV q8h) | Add if: prior IV antibiotics in 90 days, local MDR prevalence >10%, or septic shock |

### Duration
- 7 days for most VAP episodes
- Shorter courses (may be adequate) guided by clinical response and procalcitonin
- Longer courses may be needed for: non-fermenting gram-negatives (Pseudomonas, Acinetobacter), empyema, lung abscess, necrotizing pneumonia

---

## De-escalation Criteria

De-escalation is the narrowing of empiric broad-spectrum antibiotics based on culture results and clinical response. Should occur at 48-72 hours.

### When to De-escalate

| Criterion | Action |
|-----------|--------|
| Culture with identified pathogen and susceptibilities | Narrow to targeted therapy |
| Negative cultures at 48-72h + clinical improvement | Consider stopping MRSA and/or second gram-negative agent |
| Respiratory cultures negative + improving | Consider stopping antibiotics (especially if alternative diagnosis likely) |
| Procalcitonin declining >80% from peak or to <0.25 ng/mL | Support for shorter duration or de-escalation |

### When NOT to De-escalate

- Clinical worsening (rising temp, WBC, worsening oxygenation)
- Cultures pending or inadequate specimens
- Immunocompromised patients (may have false-negative cultures)
- Necrotizing infection or empyema
- Known MDR organisms with limited alternatives

### IV-to-Oral Switch Criteria

Switch from IV to oral antibiotics when ALL of the following are met:
- Temperature <37.8C for >=24 hours
- Heart rate <100 bpm
- Respiratory rate <24
- SBP >=90 mmHg
- SpO2 >=90% on room air or baseline O2
- Able to tolerate oral intake
- Normal GI absorption (no vomiting, ileus)
- Improving or stable WBC

Oral step-down options:
- Ceftriaxone -> Amoxicillin-clavulanate 875/125mg PO BID
- Azithromycin IV -> Azithromycin 500mg PO daily (then 250mg)
- Levofloxacin IV -> Levofloxacin 750mg PO daily
- Piperacillin-tazobactam -> Amoxicillin-clavulanate (if susceptible) or Levofloxacin PO

---

## Treatment Duration

### General Principles

| Pneumonia Type | Minimum Duration | Standard Duration | Extended Duration |
|----------------|-----------------|-------------------|-------------------|
| CAP, uncomplicated | 5 days | 5-7 days | Not needed if responding |
| CAP, severe | 7 days | 7-10 days | If slow response |
| HAP | 7 days | 7 days | If MDR pathogen or slow response |
| VAP | 7 days | 7 days | Non-fermenting GNR or complications |
| Lung abscess | 4-6 weeks | -- | Until radiographic improvement |
| Empyema | 2-4 weeks | -- | Plus drainage procedure |

### Criteria for Stopping Antibiotics

Stop antibiotics when:
- Completed minimum duration (>=5 days for CAP)
- Afebrile for >= 48 hours
- No more than 1 sign of clinical instability (see IV-to-oral criteria)
- Able to take oral intake
- No evidence of complications (empyema, abscess, metastatic infection)

Procalcitonin-guided: if available, consider stopping when PCT <0.25 ng/mL or decreased >80% from peak (supported by multiple RCTs for reducing antibiotic exposure).
