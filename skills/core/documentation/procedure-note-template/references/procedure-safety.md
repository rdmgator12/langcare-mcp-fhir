# Procedure Safety Documentation

## Informed Consent Requirements

### Legal Elements of Valid Informed Consent

| Element | Description |
|---------|------------|
| Capacity | Patient has decision-making capacity (or legal surrogate identified) |
| Disclosure | Nature of the procedure, risks, benefits, alternatives, and risk of no treatment explained |
| Understanding | Patient demonstrates understanding of the information |
| Voluntariness | Decision made without coercion |
| Authorization | Patient (or surrogate) agrees and signs consent |

### Consent Documentation Checklist

```
INFORMED CONSENT:
  Procedure: [Exact procedure name, laterality if applicable]
  Consented by: [Patient / Surrogate name and relationship]
  Consenting provider: [Name, credentials]
  Date/Time of consent: [datetime]

  Discussed with patient/surrogate:
  - [ ] Nature and purpose of the procedure
  - [ ] Expected benefits
  - [ ] Material risks and potential complications
  - [ ] Alternatives to the proposed procedure
  - [ ] Risks of not performing the procedure
  - [ ] Opportunity for questions provided

  Patient/surrogate verbalized understanding: [Yes]
  Written consent signed: [Yes / Verbal consent witnessed by [name] -- document reason written consent not obtained]
```

### Emergency Consent Exceptions

When informed consent cannot be obtained:

| Situation | Documentation Required |
|-----------|----------------------|
| Life-threatening emergency | Document: "Emergency procedure. Patient unable to consent due to [reason]. No surrogate immediately available. Delay would result in [death / permanent harm]. Proceeding under emergency exception." |
| Incapacitated, no surrogate | Two-physician consent documented |
| Minor without parent | Document attempts to reach parent/guardian; proceed if emergency |
| Telephone consent | Document: witness name, time of call, relationship of consenter, two-witness verification |

### Surrogate Decision-Maker Hierarchy

When patient lacks capacity (varies by state, general hierarchy):
1. Court-appointed guardian
2. Healthcare power of attorney / healthcare proxy
3. Spouse or domestic partner
4. Adult child
5. Parent
6. Adult sibling
7. Close friend or other relative

Document: "[Surrogate name], [relationship], [legal authority basis]"

## Time-Out / Universal Protocol

### The Joint Commission Universal Protocol (NPSG)

Required for ALL invasive procedures, not just operating room procedures.

#### Pre-Procedure Verification

```
PRE-PROCEDURE VERIFICATION:
  - [ ] Correct patient identified (2 identifiers: [name] + [DOB/MRN])
  - [ ] Correct procedure verified (matches consent, matches order)
  - [ ] Correct site/side verified (if applicable)
  - [ ] Site marked (if laterality, multiple structures, or levels involved)
      Site marked by: [operating/performing physician or licensed designee]
      Mark: [initials or "YES" at procedure site, visible after prep/drape]
  - [ ] Relevant documents available:
      - Signed consent: [Yes]
      - H&P or pre-procedure note: [Yes]
      - Pre-procedure labs reviewed: [Yes]
      - Imaging available: [Yes / NA]
  - [ ] Required equipment and implants verified: [Yes]
  - [ ] Blood products available (if applicable): [Yes / NA]
```

#### Time-Out (Performed Immediately Before Starting)

```
TIME-OUT PERFORMED:
  Time: [HH:MM]
  Participants: [All team members present and participating]

  Verified:
  - [ ] Correct patient identity
  - [ ] Correct procedure and site/side
  - [ ] Patient position correct
  - [ ] Relevant imaging displayed (if applicable)
  - [ ] Antibiotic prophylaxis given (if applicable): [Agent] at [time]
  - [ ] Allergies reviewed: [list or NKDA]
  - [ ] Special equipment/implants confirmed
  - [ ] Anticipated critical events discussed:
      - Surgeon/operator: anticipated steps, duration, EBL, special considerations
      - Anesthesia: patient-specific concerns (if applicable)
      - Nursing: sterility confirmed, equipment available

  All team members confirmed agreement to proceed: [Yes]
```

### Site Marking Requirements

| Procedure | Site Marking Required? |
|-----------|----------------------|
| Procedures involving laterality (right/left) | Yes |
| Multiple structures (e.g., fingers, toes) | Yes |
| Multiple levels (spine) | Yes |
| Single-organ procedures (e.g., LP at midline) | Not required but recommended |
| Bedside procedures at a single site | Institutional policy varies |

**Marking rules:**
- Mark must be at or near the procedure site
- Mark must be unambiguous (typically "YES" or operator initials)
- Mark must be visible after prep and draping
- Mark must be made by the person performing the procedure or a licensed designee
- Do NOT mark the non-operative site

## Specimen Handling Documentation

### Chain of Custody

```
SPECIMEN DOCUMENTATION:
  Specimen type: [Fluid / Tissue / Culture swab / Biopsy]
  Source: [Anatomic location]
  Collection time: [HH:MM]
  Collected by: [Name]

  Labeling:
  - [ ] Patient name on specimen container
  - [ ] Patient MRN/DOB on specimen container
  - [ ] Date and time of collection
  - [ ] Specimen source/type
  - [ ] Two-person verification at bedside (for labeled specimens)

  Sent to:
  - [ ] Laboratory: [cell count, chemistry, etc.]
  - [ ] Microbiology: [culture type, Gram stain]
  - [ ] Cytology/Pathology: [cytology, cell block, tissue analysis]
  - [ ] Special handling: [on ice / in formalin / in heparinized tube / in anaerobic transport]
```

### Specimen Handling by Procedure Type

| Procedure | Specimens | Tubes/Containers | Special Handling |
|-----------|-----------|-----------------|-----------------|
| Lumbar puncture | CSF | Tube 1-4 (sterile, numbered) | Send sequentially; glucose in fluoride tube |
| Paracentesis | Ascitic fluid | Blood culture bottles (bedside inoculation), purple top, chemistry tube | Inoculate cultures at bedside for higher sensitivity |
| Thoracentesis | Pleural fluid | Purple top, chemistry, pH in heparinized syringe, culture bottles, cytology | pH: send on ice as blood gas stat |
| Chest tube | Pleural fluid | Same as thoracentesis | |
| Wound drainage | Fluid | Culture swab or aspirate in syringe | Anaerobic transport if deep tissue/abscess |
| Biopsy | Tissue | Formalin container (pathology), sterile container (culture) | Do NOT put culture specimens in formalin |

### Laboratory Test Selection by Fluid Type

#### Cerebrospinal Fluid (CSF)

| Test | Tube | Purpose |
|------|------|---------|
| Cell count + differential | Tube 1 and Tube 4 | Infection, SAH (compare tube 1 vs 4 for traumatic vs true hemorrhage) |
| Glucose | Tube 2 | Low in bacterial meningitis (< 40 or < 2/3 serum) |
| Protein | Tube 2 | Elevated in meningitis, GBS, malignancy |
| Gram stain + culture | Tube 3 | Bacterial identification |
| HSV PCR | Tube 4 | Herpes encephalitis |
| Cryptococcal antigen | Additional | Fungal meningitis (immunocompromised) |
| VDRL | Additional | Neurosyphilis |
| Cytology | Additional | Carcinomatous meningitis |
| Oligoclonal bands, IgG index | Additional | Multiple sclerosis |
| AFB smear + culture | Additional | TB meningitis |

#### Ascitic Fluid

| Test | Purpose |
|------|---------|
| Cell count + differential | SBP: PMN > 250/mm3 |
| Albumin | SAAG calculation (serum albumin - ascitic albumin) |
| Total protein | SAAG >= 1.1 = portal hypertension |
| Culture (blood culture bottles) | Organism identification |
| Glucose | Low in SBP |
| LDH | Secondary peritonitis if very elevated |
| Amylase | Pancreatic ascites |
| Triglycerides | Chylous ascites (> 200 mg/dL) |
| Cytology | Malignant ascites |
| Bilirubin | Biliary leak (ascitic bilirubin > serum) |

#### Pleural Fluid

| Test | Purpose |
|------|---------|
| Cell count + differential | Neutrophil-predominant (bacterial) vs lymphocyte-predominant (TB, malignancy) |
| Protein | Light's criteria |
| LDH | Light's criteria |
| Albumin | Serum-effusion albumin gradient (if Light's misclassifies) |
| Glucose | Low in empyema, RA, malignancy |
| pH | < 7.2 suggests empyema requiring drainage |
| Gram stain + culture | Organism identification |
| Cytology | Malignancy |
| Triglycerides | Chylothorax (> 110 mg/dL) |
| Adenosine deaminase | TB pleuritis (> 40 U/L) |
| AFB smear + culture | TB |

## Conscious Sedation Documentation Requirements

### Pre-Sedation Assessment

```
PRE-SEDATION ASSESSMENT:
  ASA Physical Status: [I / II / III / IV / V / VI]
  Mallampati Class: [I / II / III / IV]
  NPO status: Solids [hours ago] | Liquids [hours ago]
  Relevant history:
  - Prior sedation/anesthesia complications: [None / describe]
  - Sleep apnea: [Yes / No]
  - Difficult airway history: [Yes / No]
  - Active respiratory illness: [Yes / No]
  - Allergies to sedation medications: [None / list]
  Baseline vital signs: HR [hr] | BP [sys]/[dia] | RR [rr] | SpO2 [%]
  Consent for sedation: [Obtained / included in procedure consent]
```

### ASA Physical Status Classification

| Class | Definition | Examples |
|-------|-----------|----------|
| I | Normal healthy patient | Healthy, non-smoking, minimal alcohol |
| II | Mild systemic disease | Controlled HTN, controlled DM, mild lung disease, social drinker, BMI 30-40 |
| III | Severe systemic disease | Poorly controlled DM or HTN, COPD, morbid obesity (BMI >= 40), active hepatitis, alcohol dependence, pacemaker, moderate EF reduction, ESRD on dialysis, history of MI/CVA/TIA/CAD > 3 months |
| IV | Severe systemic disease that is a constant threat to life | Recent MI/CVA/TIA (< 3 months), severe valve dysfunction, sepsis, DIC, ARDS, severe EF reduction |
| V | Moribund, not expected to survive without procedure | Ruptured AAA, massive trauma, intracranial bleed with mass effect |
| VI | Brain-dead organ donor | |

### Sedation Monitoring Requirements

```
INTRA-PROCEDURE SEDATION MONITORING:
  Monitoring modalities:
  - [ ] Continuous pulse oximetry
  - [ ] Continuous cardiac monitoring (ECG)
  - [ ] Continuous capnography / ETCO2 (recommended for moderate-deep sedation)
  - [ ] Blood pressure q [3-5] minutes
  - [ ] Level of consciousness assessed q [5] minutes

  Medications administered:
  | Time    | Drug          | Dose      | Route | Provider |
  |---------|---------------|-----------|-------|----------|
  | [HH:MM] | [drug name]   | [dose mg] | IV    | [name]   |
  | [HH:MM] | [drug name]   | [dose mg] | IV    | [name]   |

  Vital signs during sedation:
  | Time    | HR  | BP      | RR  | SpO2 | ETCO2 | Sedation Level |
  |---------|-----|---------|-----|------|-------|----------------|
  | [HH:MM] | [v] | [s]/[d] | [v] | [%]  | [v]   | [level]        |

  Level of sedation achieved: [Minimal / Moderate / Deep]
  Sedation level scale used: [Ramsay / RASS / institutional]
```

### Sedation Level Definitions

| Level | Responsiveness | Airway | Ventilation | CV Function |
|-------|---------------|--------|-------------|-------------|
| Minimal (anxiolysis) | Normal to verbal | Unaffected | Unaffected | Unaffected |
| Moderate (conscious) | Purposeful to verbal/tactile | No intervention needed | Adequate | Usually maintained |
| Deep | Purposeful to repeated/painful stimuli | May need intervention | May be inadequate | Usually maintained |
| General anesthesia | Unarousable to painful stimuli | Intervention often needed | Often inadequate | May be impaired |

### Post-Sedation Recovery Documentation

```
POST-SEDATION RECOVERY:
  Recovery monitoring start time: [HH:MM]
  Vital signs q [5-15] min during recovery:
  | Time    | HR  | BP      | RR  | SpO2 | Sedation Level |
  |---------|-----|---------|-----|------|----------------|
  | [HH:MM] | [v] | [s]/[d] | [v] | [%]  | [level]        |

  Discharge criteria met (modified Aldrete score >= 9):
  - [ ] Activity: moves 4 extremities [2] / 2 extremities [1] / 0 [0]
  - [ ] Respiration: deep breath and cough [2] / dyspnea [1] / apnea [0]
  - [ ] Circulation: BP within 20% of pre-sedation [2] / within 20-50% [1] / >50% [0]
  - [ ] Consciousness: fully awake [2] / arousable [1] / not responding [0]
  - [ ] SpO2: >92% on RA [2] / needs O2 to maintain >90% [1] / <90% with O2 [0]

  Total Aldrete score: [/10]
  Time to meet discharge criteria: [HH:MM]
  Adverse events during recovery: [None / list]
  Patient discharged from recovery at: [HH:MM]
```

## Complications by Procedure Type

### Central Venous Catheter

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Pneumothorax | 1-5% (subclavian) | Post-procedure CXR, dyspnea, chest pain | Small: observe. Large/symptomatic: chest tube |
| Arterial puncture | 3-15% | Bright red pulsatile blood, compare to ABG | Remove needle, hold pressure 10-15 min |
| Air embolism | < 1% | Acute dyspnea, hypotension, mill-wheel murmur | Trendelenburg, left lateral decubitus, aspiration |
| Arrhythmia | 10-40% (guidewire) | Monitor during insertion | Withdraw wire, usually self-limited |
| Hematoma | 2-5% | Swelling at site | Pressure, monitor, reversal agents if coagulopathic |
| Catheter malposition | 5-15% | CXR shows tip in wrong location | Reposition or replace |
| Catheter-related bloodstream infection | 1-5/1000 catheter-days | Fever, bacteremia, no other source | Remove catheter, blood cultures, antibiotics |
| Thoracic duct injury | Rare (left subclavian/IJ) | Chylothorax | Surgical consult |

### Arterial Line

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Hematoma | 10-15% | Swelling at site | Pressure, monitor distal perfusion |
| Thrombosis | 5-25% (often subclinical) | Loss of waveform, distal ischemia | Remove line, anticoagulation if symptomatic |
| Pseudoaneurysm | < 1% | Pulsatile mass | Ultrasound-guided compression, surgical repair |
| Distal ischemia | < 1% (radial), higher (brachial) | Cool/pale/painful fingers | Remove line, vascular consult |
| Infection | < 1% | Site erythema, bacteremia | Remove line, local care |

### Endotracheal Intubation

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Esophageal intubation | 2-8% | No ETCO2, no breath sounds, gastric distension | Immediately remove, re-intubate |
| Right mainstem | 5-10% | Decreased left breath sounds, CXR | Pull back to 21-23 cm at teeth |
| Dental injury | 1-12% | Direct visualization | Document, dental consult |
| Aspiration | 1-8% | Witnessed regurgitation, new infiltrate on CXR | Suction, antibiotics if clinical aspiration pneumonia |
| Desaturation | Variable | SpO2 drop | Preoxygenation, apneic oxygenation, limit attempts |
| Hypotension | Common post-intubation | BP drop after induction | Fluid bolus, vasopressor |
| Laryngeal injury | 4-13% | Post-extubation stridor, hoarseness | Steroids, ENT consult if severe |

### Lumbar Puncture

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Post-LP headache | 10-30% | Positional headache (worse upright, better supine) within 48h | Conservative (caffeine, fluids, analgesics). Blood patch if persistent > 24-48h |
| Traumatic tap | 10-20% | Bloody CSF, clearing with sequential tubes | Compare tube 1 and tube 4 RBC counts |
| Back pain | 25-35% | Local pain at insertion site | NSAIDs, ice, usually self-limited |
| Cerebral herniation | < 1% (screen with CT) | Decreased consciousness, pupil changes | Emergency neurosurgical intervention |
| Epidural hematoma | < 0.1% | Back pain, leg weakness, bladder dysfunction | MRI, neurosurgical decompression if needed |
| Infection (meningitis, abscess) | < 0.01% | Fever, headache, meningismus after LP | Blood/CSF cultures, antibiotics |

### Paracentesis

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Persistent leak | 5% | Continued ascitic fluid from puncture site | Pressure, position change, purse-string suture |
| Abdominal wall hematoma | 1-2% | Pain, swelling, dropping hemoglobin | Pressure, imaging, transfusion if needed |
| Bowel perforation | < 0.1% | Feculent fluid, peritonitis signs | Surgical consult, NPO, antibiotics |
| Hypotension (large volume) | Variable | BP drop during/after removal | Albumin replacement (6-8g per liter > 5L), IV fluids |
| Infection | < 1% | Site cellulitis | Antibiotics |

### Thoracentesis

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Pneumothorax | 5-15% (lower with US) | Dyspnea, decreased breath sounds, CXR | Small/asymptomatic: observe. Symptomatic: chest tube |
| Re-expansion pulmonary edema | < 1% | Cough, dyspnea, frothy sputum during drainage | Stop drainage, supportive care, limit to 1500 mL per session |
| Vasovagal | 5-10% | Bradycardia, hypotension, diaphoresis | Supine position, atropine if severe |
| Hemothorax | < 1% | Bloody aspirate, hemodynamic instability | Chest tube, surgical consult if ongoing |
| Infection (empyema) | < 1% | Fever, purulent fluid on repeat imaging | Chest tube drainage, antibiotics |

### Chest Tube

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Malposition | 5-10% | Inadequate drainage, CXR shows tube outside pleural space | Reposition or replace |
| Intercostal artery injury | < 1% | Significant hemorrhage from insertion site | Pressure, may need embolization or surgery |
| Lung parenchymal injury | < 1% | Air leak, hemoptysis | Usually self-limited, surgical consult if persistent |
| Subcutaneous emphysema | 5-10% | Crepitus around site and chest wall | Usually benign, ensure tube functioning |
| Re-expansion pulmonary edema | < 1% | Cough, frothy sputum after lung re-expansion | Supportive, avoid rapid large-volume drainage |
| Empyema/infection | 2-5% | Fever, purulent drainage, positive culture | Antibiotics, possible tube exchange or decortication |
| Dislodgement | Variable | Tube migrates, loss of function | Secure properly, replace if needed |

### Foley Catheter

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| UTI (CAUTI) | 3-7% per day indwelling | Fever, pyuria, bacteriuria, dysuria | Remove catheter, antibiotics if symptomatic |
| Urethral trauma | 1-3% | Resistance on insertion, bleeding, false passage | Urology consult, do not force |
| Bladder spasm | 10-20% | Suprapubic pain, bypassing around catheter | Anticholinergics, ensure catheter not obstructed |
| Urethral stricture (long-term) | Variable | Difficulty with subsequent catheterization | Urology referral |
| Encrustation/obstruction | Variable | Decreased output, catheter blockage | Irrigation, catheter change |

### NG Tube

| Complication | Incidence | Recognition | Management |
|-------------|-----------|-------------|------------|
| Epistaxis | 5-10% | Bleeding from naris during insertion | Direct pressure, vasoconstrictors, use other naris |
| Aspiration | < 1% | Cough, desaturation during insertion | Remove tube, suction, CXR |
| Tracheal placement | 1-2% | Coughing, desaturation, inability to speak | Remove immediately, confirm all tube placements with X-ray before use for feeding |
| Esophageal perforation | < 0.1% | Chest pain, subcutaneous air, pneumomediastinum | NPO, CT, surgical consult |
| Sinusitis (prolonged use) | 5-10% | Facial pain, fever, nasal discharge | Remove NG, antibiotics, consider OG route |
| Nasal alar necrosis | Variable with prolonged use | Tissue breakdown at naris | Reposition, alternate nares, nasal bridle |
