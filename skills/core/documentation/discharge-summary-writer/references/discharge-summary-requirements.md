# Discharge Summary Requirements and Medication Reconciliation Format

## CMS/TJC Required Elements

### CMS Conditions of Participation -- 42 CFR 482.24(c)(2)(viii)

A discharge summary must contain:

| Element | Description | Mandatory |
|---------|------------|-----------|
| Reason for hospitalization | Principal diagnosis or chief complaint prompting admission | Yes |
| Significant findings | Key diagnostic results, imaging, pathology, consultant conclusions | Yes |
| Procedures and treatment provided | All significant procedures, surgeries, and therapeutic interventions | Yes |
| Patient's discharge condition | Clinical status and functional status at time of discharge | Yes |
| Patient and family instructions | Activity, diet, medications, follow-up, warning signs | Yes |
| Attending physician signature | Physician attestation that the summary is complete and accurate | Yes |

### The Joint Commission PC.04.02.01

Additional TJC requirements:

| Element | Description |
|---------|------------|
| Admission diagnosis | Working diagnosis at time of admission |
| Discharge diagnosis | Final diagnosis at time of discharge (must include principal + secondary) |
| Disposition | Where the patient went (home, SNF, rehab, LTAC, hospice, AMA, deceased) |
| Medications at discharge | Complete list with dose, route, frequency |
| Pending results | Labs, cultures, pathology not yet finalized at discharge |
| Follow-up plan | Specific appointments, timeframes, providers |

### Timeliness Requirements

| Standard | Timeframe |
|----------|-----------|
| CMS | Completed "promptly" -- no specific deadline in regulation |
| TJC | Must be completed within **30 days** of discharge |
| Best practice | Within **24-48 hours** of discharge |
| Transfers | Must accompany or be available to receiving facility at time of transfer |
| Deaths | Summary/death note within 30 days |

## Discharge Diagnosis Documentation

### Principal Diagnosis Selection

Per UHDDS (Uniform Hospital Discharge Data Set):
- The **principal diagnosis** is the condition established after study to be chiefly responsible for the hospital admission
- This may differ from the admitting diagnosis
- Must be assigned an ICD-10-CM code

### Secondary Diagnoses

Document all conditions that:
- Required clinical evaluation
- Required therapeutic treatment
- Required diagnostic procedures
- Extended the length of stay
- Increased nursing care or monitoring

### Present on Admission (POA) Indicators

| Indicator | Meaning |
|-----------|---------|
| Y | Yes, present at admission |
| N | No, not present at admission (hospital-acquired) |
| U | Unknown at time of discharge |
| W | Clinically undetermined |
| Exempt | Exempt from POA reporting (certain V/Z codes) |

## Hospital Course Documentation

### Structured Hospital Course Format

```
HOSPITAL COURSE:

BRIEF SUMMARY:
[1-2 sentence overview: Patient was admitted for [reason] and treated with [key interventions].
Hospital course was [uncomplicated / complicated by X]. Patient is being discharged in [condition].]

DETAILED COURSE:

Day 1 (Admission - [date]):
  - Presenting symptoms: [brief]
  - Initial workup: [key labs, imaging ordered]
  - Initial treatment: [medications started, procedures performed]
  - Consulting services: [services consulted]

Hospital Day [n] ([date]):
  - [Significant events, changes in condition]
  - [New test results]
  - [Treatment changes]
  - [Procedure performed]

[Continue for each significant day or group routine days]

Day [n] (Discharge - [date]):
  - Clinical status: [stable for discharge criteria met]
  - Discharge planning: [completed, instructions given]
```

### Hospital Course for Common Admitting Diagnoses

#### Heart Failure Exacerbation
Key elements to document:
- Admission weight vs dry weight
- BNP/NT-proBNP trend (admission -> discharge)
- Daily weight trend and I&O
- Diuretic regimen and response
- Echocardiogram findings (EF, wall motion, valvular disease)
- Optimization of GDMT (beta-blocker, ACEi/ARB/ARNI, MRA, SGLT2i)
- Volume status at discharge (euvolemic, near-euvolemic)
- Discharge weight vs dry weight
- Precipitant identified (dietary indiscretion, medication non-compliance, arrhythmia, infection, ACS)

#### Pneumonia
Key elements to document:
- Organism identified (or no organism identified)
- Antibiotic regimen with start date and planned total duration
- Day of clinical improvement (defervescence, WBC normalization)
- Imaging trend (CXR admission vs discharge)
- Oxygen requirement trend (peak -> discharge)
- Antibiotic transition (IV to PO date)
- Remaining antibiotic course at discharge
- Vaccination status (pneumococcal, influenza)

#### Acute MI
Key elements to document:
- STEMI vs NSTEMI, culprit lesion
- Peak troponin
- Catheterization findings and intervention (PCI, stent type, CABG referral)
- Complications (arrhythmia, HF, cardiogenic shock, mechanical complication)
- EF post-event
- Dual antiplatelet therapy initiated (agents, duration plan)
- GDMT initiated (statin, beta-blocker, ACEi/ARB)
- Cardiac rehabilitation referral
- Smoking cessation if applicable

#### Surgical Admissions
Key elements to document:
- Pre-operative workup summary
- Procedure performed (include operative note reference)
- Intra-operative complications (if any)
- Post-operative course by POD
- Wound status at discharge
- Drain removal or management plan
- Activity progression
- Pain management transition (IV -> PO, opioid taper)
- VTE prophylaxis plan post-discharge

## Discharge Medication Reconciliation Format

### Required Comparison Format

```
DISCHARGE MEDICATION RECONCILIATION
====================================

MEDICATIONS CONTINUING UNCHANGED FROM HOME:
  1. [Drug name] [dose] [route] [frequency]
     Purpose: [indication]
  2. [Drug name] [dose] [route] [frequency]
     Purpose: [indication]

NEW MEDICATIONS STARTED DURING HOSPITALIZATION:
  3. [Drug name] [dose] [route] [frequency] -- NEW
     Indication: [why started]
     Duration: [indefinite / X days / until follow-up]
     Special instructions: [if any]
  4. [Drug name] [dose] [route] [frequency] -- NEW
     Indication: [why started]

MEDICATIONS WITH DOSE OR FREQUENCY CHANGES:
  5. [Drug name] [NEW dose] [route] [frequency] -- CHANGED
     Previous: [old dose] [old frequency]
     Reason for change: [reason]

MEDICATIONS DISCONTINUED:
  6. [Drug name] [old dose] -- STOPPED
     Reason: [reason for discontinuation]
     Important: [any warnings, e.g., "Do NOT resume without physician approval"]

MEDICATIONS ON HOLD (temporary):
  7. [Drug name] [dose] -- ON HOLD
     Reason: [reason]
     Resume: [when/condition for resumption]
     Follow-up required: [what needs to happen before resuming]
```

### High-Alert Medication Discharge Requirements

For each high-alert medication being sent home, document:

| Requirement | Description |
|------------|-------------|
| Indication | Why the patient needs this medication |
| Dose verification | Confirmed correct dose for patient (weight, renal function, etc.) |
| Monitoring plan | What the patient/provider needs to monitor (INR, glucose, etc.) |
| Patient education | Documented that education was provided |
| Follow-up labs | Specific lab tests and timing (e.g., "INR in 3 days") |
| Drug interactions | Key interactions reviewed with patient |
| Warning signs | When to seek medical attention |
| Prescriber contact | Who to call with questions |

### Anticoagulant-Specific Discharge Documentation

```
ANTICOAGULATION DISCHARGE PLAN:
  Agent: [warfarin / apixaban / rivaroxaban / dabigatran / edoxaban / enoxaparin]
  Indication: [DVT/PE / AFib / mechanical valve / other]
  Dose: [dose and frequency]
  Duration: [planned duration or indefinite]

  For warfarin:
    - Discharge INR: [value]
    - Target INR range: [range]
    - Next INR check: [date]
    - Anticoagulation clinic: [name, phone, first appointment]
    - Bridge therapy: [Yes/No - if yes: agent, dose, duration]

  For all anticoagulants:
    - Bleeding precautions reviewed: [Yes]
    - Drug/food interactions reviewed: [Yes]
    - When to seek emergency care: [signs of major bleeding]
    - Medical alert identification: [recommended]
    - Dental/surgical procedures: [notify provider of anticoagulant use]
```

### Insulin Discharge Documentation

```
INSULIN DISCHARGE PLAN:
  Basal insulin: [type] [dose] [timing]
  Mealtime insulin: [type] [dose] [timing]
  Correction scale: [included / not needed]

  Blood glucose monitoring:
    - Frequency: [# times/day]
    - Targets: Fasting [range], Pre-meal [range], Bedtime [range]
    - Hypoglycemia plan: If BG < 70: [instructions]
    - When to call provider: BG > [threshold] or BG < [threshold]

  Supplies prescribed:
    - [ ] Insulin pen/vials
    - [ ] Pen needles/syringes
    - [ ] Glucometer
    - [ ] Test strips
    - [ ] Lancets
    - [ ] Sharps container

  Follow-up: Endocrinology / PCP in [timeframe]
  HbA1c at discharge: [value] (next check in 3 months)
```

## Condition-Specific Discharge Instructions

### Heart Failure (CHF)

```
HEART FAILURE DISCHARGE INSTRUCTIONS:

DAILY MONITORING:
  - Weigh yourself every morning BEFORE eating, AFTER urinating
  - Record weight in a log
  - CALL YOUR DOCTOR if weight increases > 2 lbs in 1 day or > 5 lbs in 1 week

DIET:
  - Sodium restriction: < 2,000 mg/day
  - Fluid restriction: [amount] per day (if applicable)
  - Heart-healthy diet: limit saturated fats, increase fruits and vegetables

ACTIVITY:
  - [Specific restrictions]
  - Cardiac rehabilitation: [referral provided / not indicated]
  - Gradually increase activity as tolerated

MEDICATIONS: [See medication reconciliation above]

RETURN TO EMERGENCY DEPARTMENT IF:
  - Sudden weight gain > 3 lbs overnight
  - Increasing shortness of breath, especially at rest
  - Cannot lie flat due to breathing difficulty
  - New or worsening leg/ankle swelling
  - Chest pain
  - Dizziness or fainting

FOLLOW-UP:
  - Cardiology: [date]
  - PCP: [date]
  - Lab work: BMP in [X days] at [location]
```

### Pneumonia

```
PNEUMONIA DISCHARGE INSTRUCTIONS:

ANTIBIOTICS:
  - Complete the full course: [drug] [dose] for [remaining days]
  - Take [with food / on empty stomach]
  - Do NOT stop early even if feeling better

ACTIVITY:
  - Expect gradual improvement over 2-4 weeks
  - Fatigue may persist for several weeks -- this is normal
  - Gradually return to normal activities

RETURN TO EMERGENCY DEPARTMENT IF:
  - Fever > 101.5 F (38.6 C) that does not respond to Tylenol
  - Increasing shortness of breath or difficulty breathing
  - Coughing up blood
  - Chest pain, especially with breathing
  - Confusion or disorientation
  - Unable to keep fluids down

FOLLOW-UP:
  - PCP: [date] (within 1-2 weeks)
  - Repeat chest X-ray: In 6 weeks to confirm resolution
  - Pneumonia vaccine: [Due / Up to date]
  - Influenza vaccine: [Due / Up to date]
```

### Acute MI

```
POST-HEART ATTACK DISCHARGE INSTRUCTIONS:

MEDICATIONS -- CRITICAL:
  - Aspirin [dose]: Take EVERY DAY. Do NOT stop.
  - [P2Y12 inhibitor]: Take EVERY DAY for at least [12 months]. Do NOT stop without calling cardiologist.
  - Statin [dose]: Take EVERY DAY.
  - Beta-blocker [dose]: Do NOT stop suddenly.
  - ACE inhibitor/ARB [dose]: Take as prescribed.

ACTIVITY:
  - No heavy lifting > [10 lbs] for [2-4 weeks]
  - No driving for [48h / 1 week] after procedure
  - Cardiac rehabilitation: ENROLL -- referral has been placed
  - Sexual activity: May resume in [1-2 weeks] if tolerating moderate exertion

DIET:
  - Heart-healthy diet (Mediterranean, DASH)
  - Limit sodium, saturated fats, trans fats
  - Increase omega-3 fatty acids

RETURN TO EMERGENCY DEPARTMENT IMMEDIATELY IF:
  - Chest pain or pressure similar to heart attack
  - Shortness of breath at rest
  - Pain/swelling at catheterization site
  - Arm or leg becomes cold, pale, or painful
  - Fainting or severe dizziness

FOLLOW-UP:
  - Cardiologist: [date] (within 1-2 weeks)
  - Cardiac rehabilitation: [date and location]
  - PCP: [date]
```

### Post-Surgical Discharge

```
SURGICAL DISCHARGE INSTRUCTIONS:

WOUND CARE:
  - [Specific wound care instructions per procedure]
  - Keep incision [clean and dry / covered / open to air]
  - [Shower OK after X hours/days / sponge bath only for X days]
  - Steri-strips: Will fall off on their own in 7-10 days
  - Staple/suture removal: At follow-up on [date]

DRAINS (if applicable):
  - Empty drain [frequency]
  - Record output volume
  - Call if output > [threshold] or changes to [bright red / foul-smelling / cloudy]

ACTIVITY:
  - No lifting > [weight] for [duration]
  - No driving while on narcotic pain medication
  - No [swimming / bathing / submerging wound] for [duration]
  - Ambulate [frequency] -- walking is encouraged

PAIN MANAGEMENT:
  - [Non-opioid first]: Acetaminophen [dose] every [frequency] AND/OR Ibuprofen [dose] every [frequency]
  - [Opioid if needed]: [drug] [dose] every [frequency] AS NEEDED for severe pain
  - Opioid precautions: Do not drive, do not drink alcohol, may cause constipation
  - Constipation prevention: [Docusate / senna] while taking opioid pain medication

RETURN TO EMERGENCY DEPARTMENT IF:
  - Fever > 101.5 F (38.6 C)
  - Increasing redness, swelling, or warmth around incision
  - Drainage from incision (pus, foul-smelling, increasing)
  - Wound opens or separates
  - Uncontrolled pain despite medications
  - [Procedure-specific warning signs]
```

## Pending Results Documentation

```
RESULTS PENDING AT DISCHARGE:
  1. [Test] - Collected [date] - Expected result: [date]
     Action if abnormal: [specific plan, who will follow up]
  2. [Test] - Collected [date] - Expected result: [date]
     Action if abnormal: [specific plan]
  3. Blood cultures - Collected [date] - No growth to date
     Action if positive: [covering physician will contact patient, empiric plan]

RESPONSIBLE PROVIDER FOR PENDING RESULTS:
  Name: [Provider name]
  Contact: [phone/pager]
  Plan: Results will be reviewed and patient contacted within [timeframe]
```
