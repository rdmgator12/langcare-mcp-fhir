# Critical Care Documentation Additions

## ICU Progress Note Additional Requirements

Critical care progress notes require documentation beyond standard inpatient notes per CMS billing requirements (CPT 99291-99292). The note must demonstrate:
1. **Critical illness/injury**: Life-threatening condition with high probability of significant deterioration
2. **High-complexity decision-making**: Time spent directly managing critical illness
3. **Personal and direct involvement**: Attending personally directing care

## Ventilator Documentation

### Mechanical Ventilation Parameters

Document all ventilator settings and changes:

```
VENTILATOR:
  Mode: [AC/VC | AC/PC | SIMV | PSV | PRVC | APRV | HFOV]
  Settings:
    FiO2: [%]
    PEEP: [cm H2O]
    Tidal Volume (set): [mL] ([mL/kg IBW])
    Respiratory Rate (set): [breaths/min]
    Pressure Support: [cm H2O] (if applicable)
    Inspiratory Pressure: [cm H2O] (if PC mode)
    I:E Ratio: [ratio]
  Measured:
    Tidal Volume (delivered): [mL]
    Respiratory Rate (total): [breaths/min]
    Minute Ventilation: [L/min]
    Peak Inspiratory Pressure: [cm H2O]
    Plateau Pressure: [cm H2O]
    Mean Airway Pressure: [cm H2O]
    Auto-PEEP: [cm H2O]
    Static Compliance: [mL/cm H2O]
    Driving Pressure: [Pplat - PEEP] cm H2O
  ABG: pH [x] / pCO2 [x] / pO2 [x] / HCO3 [x] / BE [x]
  P/F Ratio: [PaO2/FiO2] -> ARDS severity: [mild >200 / moderate 100-200 / severe <100]
  SpO2: [%]
```

### Ventilator Change Log

```
VENT CHANGES (24h):
  [Time]: [Parameter changed] from [old] to [new] - Reason: [reason]
  [Time]: [Parameter changed] from [old] to [new] - Reason: [reason]
```

### Weaning Assessment (Daily)

Document readiness for spontaneous breathing trial (SBT):

| Criterion | Status | Threshold |
|-----------|--------|-----------|
| FiO2 | [value] | <= 40% |
| PEEP | [value] | <= 5-8 cm H2O |
| P/F ratio | [value] | >= 150-200 |
| Hemodynamic stability | [Yes/No] | No vasopressors or low-dose only |
| Adequate mental status | [Yes/No] | Following commands |
| Cough / gag reflex | [Present/Absent] | Present |
| Minimal secretions | [Yes/No] | Suctioning < q2h |
| Rapid Shallow Breathing Index | [value] | < 105 |
| **SBT performed** | [Yes/No] | |
| **SBT result** | [Pass/Fail] | Tolerated 30-120 min without distress |
| **Extubation plan** | [description] | |

### Non-Invasive Ventilation

```
NIV:
  Device: [BiPAP / CPAP / HFNC]
  Settings:
    BiPAP: IPAP [cm H2O] / EPAP [cm H2O] / FiO2 [%] / Rate [if applicable]
    CPAP: Pressure [cm H2O] / FiO2 [%]
    HFNC: Flow [L/min] / FiO2 [%]
  Duration: [hours/day or continuous]
  Tolerance: [Tolerating well / intermittent breaks / poorly tolerated]
  ABG on current settings: [values]
```

## Vasoactive Medication (Drip) Documentation

### Vasopressor/Inotrope Documentation

```
VASOACTIVE MEDICATIONS:
  Norepinephrine: [dose] mcg/min ([mcg/kg/min]) - [Stable / Weaning / Escalating]
    Range (24h): [min] - [max] mcg/min
    Current MAP: [value] mmHg (target: [value])
  Vasopressin: [units/min] - [Fixed dose / Titrating]
  Epinephrine: [mcg/min] - [indication]
  Phenylephrine: [mcg/min] - [indication]
  Dobutamine: [mcg/kg/min] - [indication: cardiogenic shock / low CO]
  Milrinone: [mcg/kg/min] - [indication]

  HEMODYNAMICS:
    MAP: [range] (target: [goal])
    HR: [range]
    CVP: [value] (if monitored)
    ScvO2: [value]% (if available)
    Lactate trend: [values with times]
    Urine output: [mL/kg/hr over last 6h]
    Vasopressor-free hours (24h): [count]
```

### Sedation and Analgesia Documentation

```
SEDATION/ANALGESIA:
  Sedation:
    Agent: [Propofol / Dexmedetomidine / Midazolam / Ketamine]
    Rate: [dose/unit]
    Range (24h): [min] - [max]
    RASS target: [value] (e.g., -2 to 0)
    RASS current: [value]
    RASS range (24h): [min] to [max]
    Daily sedation interruption: [Performed / Not performed - reason]

  Analgesia:
    Agent: [Fentanyl / Hydromorphone / Morphine / Ketamine]
    Rate: [dose/unit] continuous + PRN
    Pain assessment (CPOT/BPS for intubated): [score]
    PRN doses (24h): [count]

  Neuromuscular Blockade (if applicable):
    Agent: [Cisatracurium / Rocuronium / Vecuronium]
    Rate: [dose/unit]
    Train-of-four: [0-4 twitches] (target: 1-2)
    Indication: [ARDS/refractory hypoxemia / status asthmaticus / other]
    Duration: [hours]
    BIS monitoring: [value] (if available)
```

### Richmond Agitation-Sedation Scale (RASS)

| Score | Term | Description |
|-------|------|-------------|
| +4 | Combative | Overtly combative, violent, immediate danger to staff |
| +3 | Very agitated | Pulls/removes tubes/catheters, aggressive |
| +2 | Agitated | Frequent non-purposeful movement, fights ventilator |
| +1 | Restless | Anxious but movements not aggressive |
| 0 | Alert and calm | |
| -1 | Drowsy | Not fully alert, sustained awakening to voice (>10 sec) |
| -2 | Light sedation | Briefly awakens to voice (eye contact <10 sec) |
| -3 | Moderate sedation | Movement or eye opening to voice, no eye contact |
| -4 | Deep sedation | No response to voice, movement to physical stimulation |
| -5 | Unarousable | No response to voice or physical stimulation |

## Lines, Tubes, and Drains Documentation

```
LINES/TUBES/DRAINS:
  Central Lines:
    [Type] ([site]) - Day [#] - Dressing: [CDI / Tegaderm] - Site: [Clean / Erythema / Drainage]
    [Type] ([site]) - Day [#] - Dressing: [type] - Site: [assessment]
    Daily necessity assessment: [Still indicated / Consider removal]

  Arterial Line:
    [Site] - Day [#] - Functioning: [Yes / Dampened / Non-functional]
    Waveform: [Normal / Dampened / Over-dampened]

  Endotracheal Tube:
    Size: [mm] - Position: [cm at lip/teeth] - Cuff pressure: [cm H2O]
    Last chest X-ray confirmation: [date]

  Tracheostomy:
    Size: [mm] - Type: [Cuffed / Uncuffed / Fenestrated] - Day [#] post-placement
    Cuff: [Inflated / Deflated] - Speaking valve: [Yes / No]

  Nasogastric/Orogastric Tube:
    Size: [Fr] - Position: [confirmed by X-ray / pH] - To: [suction / gravity / feeding]
    Output (24h): [volume] mL ([character])

  Foley Catheter:
    Day [#] - Daily necessity assessment: [Still indicated / Remove today]
    Urine output: [mL/24h] ([mL/kg/hr])
    Urine character: [Clear yellow / Concentrated / Hematuria / Sediment]

  Chest Tubes:
    [Side] [Size Fr] - Day [#] - To: [suction -20 cm H2O / water seal / clamped]
    Output (24h): [volume] mL ([character: serous / serosanguinous / bloody])
    Air leak: [None / Small / Moderate / Large]

  Drains:
    [Type] ([location]) - Output (24h): [volume] mL ([character])
```

### Line Day Tracking and CLABSI Prevention

| Metric | Documentation |
|--------|-------------|
| Central line day | Increment daily, flag if > 7 days |
| Daily necessity | "Line still needed: [Yes - reason / No - plan for removal]" |
| Insertion site | Daily assessment for erythema, drainage, tenderness |
| Dressing status | Intact, clean, dated |
| CHG bathing | Daily chlorhexidine bathing documented |
| Catheter care bundle | Compliance: [Yes / Missed element] |

## Nursing Assessment Integration Points

### Data Points to Pull from Nursing Documentation

These elements are typically documented by nursing but critical for progress notes. In FHIR, they may appear as Observation resources or may require obtaining from the nursing team.

| Data Point | FHIR Potential Source | Commonly in FHIR |
|-----------|----------------------|-------------------|
| Neurological checks (GCS, pupil response) | Observation | Sometimes |
| Skin assessment (Braden score, wounds) | Observation, Condition | Rarely |
| Fall risk assessment (Morse score) | Observation | Sometimes |
| Restraint use and assessment | Procedure, Observation | Rarely |
| Diet intake (% meals eaten) | Observation | Rarely |
| Mobility status (ambulation distance) | Observation | Rarely |
| Pain assessment (0-10, CPOT, FLACC) | Observation (LOINC 72514-3) | Often |
| Blood glucose (fingerstick) | Observation (LOINC 2345-7) | Often |

### Glasgow Coma Scale Documentation

```
NEUROLOGICAL:
  GCS: [total]/15
    Eye: [1-4] ([description])
    Verbal: [1-5] ([description]) [or NT if intubated]
    Motor: [1-6] ([description])
  Pupils: [size] mm / [size] mm - [Reactive / Sluggish / Fixed]
  Lateralizing signs: [None / describe]
  Sedation: [On / Off - must document off-sedation exam when possible]
```

| Component | Score | Response |
|-----------|-------|----------|
| Eye Opening | 4 | Spontaneous |
| | 3 | To voice |
| | 2 | To pressure |
| | 1 | None |
| Verbal | 5 | Oriented |
| | 4 | Confused |
| | 3 | Inappropriate words |
| | 2 | Incomprehensible sounds |
| | 1 | None |
| Motor | 6 | Obeys commands |
| | 5 | Localizes pain |
| | 4 | Withdrawal (flexion) |
| | 3 | Abnormal flexion (decorticate) |
| | 2 | Extension (decerebrate) |
| | 1 | None |

## ICU-Specific Documentation Checklists

### Daily ICU Checklist (ABCDEF Bundle)

| Element | Description | Status |
|---------|-------------|--------|
| **A** - Assess/manage pain | CPOT < 3, NRS < 4 | [Met / Not met] |
| **B** - Both SAT and SBT | Sedation vacation + breathing trial | [Performed / Not eligible - reason] |
| **C** - Choice of sedation | Lightest sedation to target RASS | [At target / Over-sedated / Under-sedated] |
| **D** - Delirium monitoring | CAM-ICU assessment | [Negative / Positive / Unable to assess] |
| **E** - Early mobility | Highest level of mobility achieved | [Bed exercises / Chair / Standing / Walking] |
| **F** - Family engagement | Family updated, involved in care | [Yes / No family available] |

### Delirium Assessment (CAM-ICU)

Document at least once per shift:

```
DELIRIUM (CAM-ICU):
  Feature 1 - Acute onset/fluctuating course: [Present / Absent]
  Feature 2 - Inattention: [Present / Absent]
  Feature 3 - Altered level of consciousness: [Present / Absent] (RASS != 0)
  Feature 4 - Disorganized thinking: [Present / Absent]
  Result: [CAM-ICU Positive (Features 1+2 plus 3 or 4) / CAM-ICU Negative]
  If positive: Type: [Hyperactive / Hypoactive / Mixed]
  Interventions: [Reorientation, sleep hygiene, medication review, mobility]
```

## Critical Care Time Documentation

For billing CPT 99291-99292, document:

```
CRITICAL CARE TIME:
  Total critical care time: [XX] minutes
  Activities included:
    - Direct bedside care and examination: [XX] min
    - Review and interpretation of data (labs, imaging, hemodynamics): [XX] min
    - Medical decision-making and care coordination: [XX] min
    - Family discussion regarding goals of care: [XX] min
    - [Other qualifying activities]: [XX] min

  Separately billable procedures performed (time excluded):
    - [Procedure] ([CPT code]) - [XX] min
    - [Procedure] ([CPT code]) - [XX] min

  Critical illness: [Specific life-threatening condition requiring direct management]
```

**Qualifying critical care activities:**
- Direct patient care in the ICU
- Reviewing test results and imaging
- Discussing management with other physicians
- Documenting critical care services
- Time spent on the floor/unit making critical decisions for the patient

**Non-qualifying activities:**
- Separately billable procedures (endotracheal intubation, central line, chest tube, etc.)
- Time spent teaching residents (non-billable portion)
- Time when not actively managing the critical illness

## Organ System Review for ICU Notes

### Structured System-by-System ICU Assessment

```
NEURO: GCS [score], RASS [score], CAM-ICU [+/-], pupils [size/reactivity], [exam findings]
CV: HR [rate/rhythm], BP [range] MAP [range], [vasopressors if any], [rhythm: NSR/AFib/etc], [cardiac output if measured]
PULM: Vent mode/settings, ABG, P/F ratio, lung exam, [secretions], SBT status
GI: Diet [type], bowel sounds, abdominal exam, [TPN if applicable], bowel regimen, LFTs if relevant
RENAL: UOP [mL/kg/hr], Cr [value/trend], BUN, electrolytes, fluid balance, [RRT if applicable]
HEME: Hgb/Plt/INR, [anticoagulation status], [transfusion needs], DVT prophylaxis
ID: Tmax, WBC trend, cultures [pending/results], antibiotics [day #/planned duration], procalcitonin
ENDO: Glucose range, insulin regimen, [thyroid if applicable], [adrenal if applicable]
SKIN: Wounds, pressure injuries (Braden score), line sites
SOCIAL: Family communication, goals of care status, code status, [barriers to care]
DISPO: ICU day #, anticipated ICU course, barriers to step-down
```
