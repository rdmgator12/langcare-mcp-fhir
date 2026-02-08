# Sepsis Bundles and Management Reference

## Surviving Sepsis Campaign 2021: Hour-1 Bundle

All elements should be initiated within 1 hour of sepsis recognition. "Time zero" is the moment of sepsis identification (triage time in ED, time of clinical deterioration on floor).

### Bundle Elements

| Element | Action | Target | How to Verify in FHIR |
|---------|--------|--------|----------------------|
| 1. Measure lactate | Stat lactate level | Result within 1 hour | Observation with LOINC 2524-7 or 32693-4, effectiveDateTime within 1 hour of recognition |
| 2. Blood cultures before antibiotics | At least 2 sets (aerobic + anaerobic) from 2 sites | Cultures drawn before first antibiotic dose | ServiceRequest with LOINC 600-7, authoredOn before antibiotic MedicationAdministration |
| 3. Broad-spectrum antibiotics | Empiric IV antibiotics | Administered within 1 hour | MedicationAdministration with effectiveDateTime within 1 hour |
| 4. Crystalloid resuscitation | 30 mL/kg IV crystalloid | For hypotension (SBP <90, MAP <65) or lactate >=4 mmol/L | MedicationAdministration for IV fluids or custom fluid balance Observation |
| 5. Vasopressors | If hypotensive during or after fluid resuscitation | MAP >= 65 mmHg | MedicationRequest for vasopressor agents, active status |

### Lactate Reassessment

If initial lactate > 2 mmol/L, reassess within 2-4 hours. Target: lactate clearance >= 10% or normalization. Search for repeat lactate Observation with effectiveDateTime 2-4 hours after initial measurement.

---

## Empiric Antibiotic Selection by Suspected Source

### Community-Acquired Pneumonia (CAP)

| Severity | Regimen | Notes |
|----------|---------|-------|
| Non-severe, no MRSA risk | Ceftriaxone 2g IV + Azithromycin 500mg IV | Or respiratory fluoroquinolone monotherapy |
| Severe (ICU admission) | Ceftriaxone 2g IV + Azithromycin 500mg IV | Add vancomycin if MRSA risk factors |
| Pseudomonas risk | Piperacillin-tazobactam 4.5g IV + Azithromycin 500mg IV | Or cefepime + azithromycin |
| Aspiration suspected | Ampicillin-sulbactam 3g IV | Or clindamycin if penicillin allergy |

### Urinary Tract Infection / Urosepsis

| Setting | Regimen | Notes |
|---------|---------|-------|
| Uncomplicated | Ceftriaxone 2g IV | Or ciprofloxacin 400mg IV |
| Complicated / obstruction | Piperacillin-tazobactam 4.5g IV | Urology consult for obstruction |
| ESBL risk | Meropenem 1g IV q8h | Prior ESBL organism, recent hospitalization, recent fluoroquinolone use |
| Catheter-associated | Ceftriaxone 2g IV + Vancomycin (if enterococcal risk) | Remove/replace catheter |

### Intra-Abdominal Infection

| Setting | Regimen | Notes |
|---------|---------|-------|
| Community-acquired, mild-moderate | Ceftriaxone 2g IV + Metronidazole 500mg IV | Or ertapenem 1g IV |
| Community-acquired, severe | Piperacillin-tazobactam 4.5g IV | Or meropenem 1g IV q8h |
| Healthcare-associated | Meropenem 1g IV q8h + Vancomycin | Source control essential |
| Biliary source | Ceftriaxone 2g IV + Metronidazole 500mg IV | ERCP/surgical consult |

### Skin and Soft Tissue Infection

| Type | Regimen | Notes |
|------|---------|-------|
| Necrotizing fasciitis suspected | Vancomycin + Piperacillin-tazobactam + Clindamycin | Surgical emergency -- immediate surgical consult |
| Cellulitis with sepsis | Cefazolin 2g IV q8h | Add vancomycin if MRSA risk |
| Diabetic foot | Piperacillin-tazobactam 4.5g IV | Or ertapenem + vancomycin |

### Unknown Source

| Risk Level | Regimen | Notes |
|------------|---------|-------|
| No MRSA/Pseudomonas risk | Ceftriaxone 2g IV + Metronidazole 500mg IV | Broad gram-negative + anaerobic coverage |
| MRSA risk | Add Vancomycin 25-30 mg/kg IV loading dose | Recent MRSA, healthcare exposure, IV drug use |
| Pseudomonas risk | Piperacillin-tazobactam 4.5g IV or Cefepime 2g IV | Recent Pseudomonas, structural lung disease, immunocompromised |
| Immunocompromised | Meropenem 1g IV q8h + Vancomycin | Broad coverage, consider antifungal |

### MRSA Risk Factors
- Prior MRSA colonization or infection
- Recent hospitalization (within 90 days)
- Hemodialysis
- Residence in long-term care facility
- IV drug use
- Indwelling catheter or device

### Pseudomonas Risk Factors
- Prior Pseudomonas isolation
- Structural lung disease (bronchiectasis, cystic fibrosis)
- Immunocompromised (neutropenia, transplant, high-dose steroids)
- Recent broad-spectrum antibiotic use (within 90 days)
- Chronic wound

---

## Fluid Resuscitation

### Initial Resuscitation

| Parameter | Target | Notes |
|-----------|--------|-------|
| Volume | 30 mL/kg ideal body weight crystalloid | Administer within first 1-3 hours |
| Fluid type | Balanced crystalloid preferred (Lactated Ringer's) | Normal saline acceptable; avoid hydroxyethyl starch |
| Rate | As rapidly as tolerated | May bolus 500 mL over 15-30 min |

### Assessment of Fluid Responsiveness

After initial 30 mL/kg, assess before giving additional fluids:
- **Passive leg raise test**: Raise legs to 45 degrees, measure cardiac output change. Positive if CO increases >10%.
- **Pulse pressure variation**: >13% variation suggests fluid responsiveness (mechanically ventilated patients only).
- **IVC ultrasound**: >50% collapsibility suggests fluid responsive (spontaneously breathing).
- **Lactate clearance**: Repeat lactate q2-4h. Clearance >10% suggests adequate resuscitation.

### When to Stop Fluids

- MAP >= 65 mmHg sustained
- Lactate normalizing or clearing
- Urine output >= 0.5 mL/kg/hr
- Signs of fluid overload: rising CVP, pulmonary edema, worsening oxygenation
- No further improvement in hemodynamic parameters with fluid challenge

---

## Vasopressor Algorithm

### First-Line: Norepinephrine

- Starting dose: 0.01-0.05 mcg/kg/min
- Titrate to MAP >= 65 mmHg
- Max dose before adding second agent: 0.5-1.0 mcg/kg/min (institutional variation)
- RxNorm: 3628

### Second-Line: Vasopressin

- Fixed dose: 0.03-0.04 units/min (NOT titrated)
- Add when norepinephrine dose >= 0.25-0.5 mcg/kg/min
- Purpose: reduce norepinephrine requirement, different mechanism (V1 receptor)
- RxNorm: 11149

### Third-Line: Epinephrine

- Starting dose: 0.01-0.05 mcg/kg/min
- Add when MAP target not met with norepinephrine + vasopressin
- Note: increases lactate (aerobic glycolysis) -- lactate less reliable for monitoring
- RxNorm: 3616

### Refractory Shock Considerations

- **Corticosteroids**: Hydrocortisone 200mg/day IV (50mg q6h) if norepinephrine >= 0.25 mcg/kg/min for >= 4 hours
- **Angiotensin II**: For catecholamine-resistant vasodilatory shock
- **Phenylephrine**: Avoid as first-line (reduces cardiac output); consider only if significant tachyarrhythmia limits norepinephrine use

### Vasopressor Goals

| Parameter | Target |
|-----------|--------|
| MAP | >= 65 mmHg (higher targets not beneficial in most patients) |
| Lactate | Trending down, target normalization (<2 mmol/L) |
| Urine output | >= 0.5 mL/kg/hr |
| Mental status | Improving or stable |
| Skin perfusion | Capillary refill < 3 seconds |

---

## Sepsis-3 Definitions

### Sepsis
Life-threatening organ dysfunction caused by a dysregulated host response to infection. Operationalized as suspected infection + SOFA score increase >= 2 from baseline.

### Septic Shock
Subset of sepsis with circulatory and cellular/metabolic dysfunction associated with higher mortality. Operationalized as sepsis + need for vasopressors to maintain MAP >= 65 mmHg + serum lactate > 2 mmol/L despite adequate fluid resuscitation. Hospital mortality > 40%.

### Key Timepoints to Document

| Timepoint | Significance | FHIR Documentation |
|-----------|-------------|-------------------|
| Time zero (sepsis recognition) | Starts all bundle clocks | ClinicalImpression.effectiveDateTime |
| Lactate drawn | Bundle element 1 | Observation.effectiveDateTime |
| Cultures collected | Bundle element 2 | ServiceRequest.authoredOn / Specimen.collection.collectedDateTime |
| First antibiotic dose | Bundle element 3 | MedicationAdministration.effectiveDateTime |
| Fluid resuscitation start | Bundle element 4 | MedicationAdministration or procedure note |
| Vasopressor start | Bundle element 5 | MedicationAdministration.effectiveDateTime |
| Repeat lactate | Reassessment | Observation.effectiveDateTime (2-4h after initial) |
