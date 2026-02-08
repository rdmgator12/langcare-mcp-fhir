# Progress Note Format and Documentation Standards

## Problem-Based vs SOAP Documentation

### SOAP Format (Traditional)

Used primarily in outpatient settings and some inpatient programs. Single unified note structure.

```
S: Subjective
O: Objective
A: Assessment
P: Plan
```

**Advantages**: Simple, familiar, quick for straightforward visits.
**Disadvantages**: Does not scale well for complex inpatient patients with multiple active problems.

### Problem-Oriented Progress Note (Recommended for Inpatient)

Organizes assessment and plan by problem. Each problem gets its own mini-assessment and plan. Objective data (vitals, labs, imaging) is shared across all problems.

```
EVENTS OVERNIGHT / INTERVAL CHANGES
SUBJECTIVE
OBJECTIVE (vitals, I&O, labs, imaging, meds)
ASSESSMENT & PLAN:
  Problem 1: assessment -> plan
  Problem 2: assessment -> plan
  Problem 3: assessment -> plan
DISPOSITION
```

**Advantages**: Clearer thinking, easier to follow for consultants and covering providers, better for handoffs, ensures each problem is addressed daily.
**Disadvantages**: Longer notes, potential for redundancy across problems.

### APSO Format (Assessment-Plan-Subjective-Objective)

Some institutions lead with assessment and plan, followed by supporting data. Used by physicians who want the "bottom line" first.

```
A/P: Assessment and plan by problem (most important information first)
S: Subjective
O: Objective (supporting data)
```

**Advantages**: Readers immediately see clinical thinking and actions. Efficient for sign-out and handoff.
**Disadvantages**: Less intuitive for learners.

## Daily Progress Note Required Elements

### Per CMS/TJC Documentation Standards

Every daily inpatient progress note must contain:

| Element | Description | Required |
|---------|-------------|----------|
| Date and time | When the note was written and when the patient was seen | Yes |
| Author identification | Name, credentials, service | Yes |
| Hospital day number | Calculated from admission date | Yes |
| Subjective interval changes | Patient-reported symptoms since last note | Yes |
| Vital signs | At minimum: T, HR, BP, RR, SpO2 | Yes |
| Pertinent objective data | Labs, imaging, exam findings relevant to active problems | Yes |
| Assessment | Clinical assessment of each active problem | Yes |
| Plan | Specific plan for each active problem | Yes |
| Response to treatment | Documentation of improvement, stability, or decline | Yes |
| Code status | Must be documented or referenced | Recommended |
| Discharge planning | Estimated date, barriers, pending items | Recommended |

### Overnight Events Documentation

Capture all significant overnight occurrences:

| Event Type | What to Document |
|-----------|-----------------|
| Nursing calls | Reason, assessment, intervention, outcome |
| Rapid response / Code Blue | Time, interventions, outcome, post-event plan |
| Fever spikes | Time, max temp, cultures obtained, empiric treatment |
| Vital sign deviations | Specific abnormality, intervention, response |
| Falls | Circumstances, injury assessment, fall prevention update |
| New symptoms | Onset time, assessment, workup initiated |
| PRN medication use | What was given, indication, frequency, response |
| Procedure complications | Type, intervention, current status |

### Vital Signs Trend Presentation

Present vitals as trends rather than single values for inpatient notes:

```
VITALS (24h):
Tmax: 38.2 C (02:00) -> Tcurrent: 37.0 C [Trending down]
HR: 72-95 (current 78) [Regular / Irregular]
BP: 110-135 / 62-78 (current 122/68) [MAP range: 78-97]
RR: 14-20 (current 16)
SpO2: 94-98% on [RA / 2L NC / HFNC 40L 50%]
Pain: [0-10 scale, current and trend]
```

### Intake and Output Documentation

```
I&O (24h):
  INTAKE:
    IV fluids: [type] at [rate] = [total] mL
    PO intake: [estimated] mL
    Blood products: [type] x [units] = [total] mL
    Medications (IV): [total] mL
    TOTAL INTAKE: [sum] mL

  OUTPUT:
    Urine: [total] mL (UOP: [mL/kg/hr])
    Drain #1 ([type]): [total] mL
    NG/OG output: [total] mL
    Stool: [number] episodes, [estimated] mL
    Emesis: [estimated] mL
    Insensible losses: [estimated] mL
    TOTAL OUTPUT: [sum] mL

  NET BALANCE: [+/-] [total] mL
  CUMULATIVE BALANCE (since admission): [+/-] [total] mL
```

**Fluid balance flags:**
- Net positive > 1L/day in CHF: concern for volume overload
- Net negative > 2L/day: risk of hypovolemia, monitor electrolytes
- UOP < 0.5 mL/kg/hr for 6+ hours: AKI concern (KDIGO criteria)

### Laboratory Results Presentation

Present labs grouped by panel with trend indicators:

```
LABS ([date]):
  BMP:
    Na: 138 [136 -> 138, up] | K: 4.2 [3.8 -> 4.2, up] | Cl: 102 | CO2: 24
    BUN: 22 [28 -> 22, improving] | Cr: 1.2 [1.5 -> 1.2, improving]
    Glucose: 145 [HIGH - fasting]
  CBC:
    WBC: 11.2 [15.8 -> 11.2, improving] | Hgb: 10.1 [10.8 -> 10.1, down]
    Plt: 198 | ANC: 8.4
  Coags:
    PT: 12.5 | INR: 1.1 | aPTT: 28
  Special:
    Procalcitonin: 0.3 [2.1 -> 0.8 -> 0.3, improving]
    Lactate: 1.2 [NORMAL]
    BNP: 450 [1250 -> 450, improving]

  PENDING: Blood cultures (day 2), urine culture
  CRITICAL VALUES: None
```

### Imaging Results Presentation

```
IMAGING ([date]):
  CXR (portable AP): [Brief conclusion]. Compared to [prior date]: [improved/stable/worsened].
  CT [body part] with/without contrast: [Key findings].
  Ultrasound [type]: [Key findings].
  PENDING: [Study] ordered [date], results expected [date/time].
```

### Medication Changes Section

Document all changes from the prior day:

```
MEDICATIONS:
  Active orders: [total count]

  STARTED (24h):
  - [Drug] [dose] [route] [frequency] - Indication: [reason]

  CHANGED (24h):
  - [Drug]: [old dose] -> [new dose] - Reason: [reason]

  DISCONTINUED (24h):
  - [Drug] - Reason: [reason]

  PRN ADMINISTERED (24h):
  - [Drug] [dose] x [count] doses - Indication: [reason]

  HIGH-ALERT MEDICATIONS:
  - [Drug] [dose] - Last level: [value] ([date]) - Next level due: [date]
```

## Problem-Based Assessment and Plan Structure

### For Each Problem

```
[#]. [Problem name] ([ICD-10]) - [Acute / Chronic / Acute-on-chronic]
    Status: [Improving / Stable / Worsening] - [supporting evidence]
    Overnight: [Any overnight events related to this problem]
    Assessment: [Clinical reasoning, response to treatment, trend interpretation]
    Plan:
      - [Specific action item 1]
      - [Specific action item 2]
      - [Monitoring plan]
      - [Contingency: if [trigger], then [action]]
```

### Problem Prioritization

Order problems by clinical urgency:
1. **Primary reason for admission** (always first)
2. **Active acute problems** requiring daily management
3. **Chronic conditions** being actively managed or affected by hospitalization
4. **Preventive/prophylactic measures** (DVT, GI, glycemic management)
5. **Disposition planning**

### Standing Orders / Prophylaxis Section

Always document the status of:

| Category | Options to Document |
|----------|-------------------|
| DVT prophylaxis | Pharmacologic (drug, dose) OR Mechanical (SCDs) OR Contraindicated (reason) |
| GI prophylaxis | PPI/H2B OR Not indicated (risk assessment) |
| Glycemic management | Insulin regimen OR oral agents OR diet-controlled OR not applicable |
| Bowel regimen | Current regimen OR not needed |
| Diet | NPO / clear liquids / full liquids / regular / cardiac / renal / diabetic / other |
| Activity | Bedrest / BRP / OOB to chair / ambulate with assist / ad lib |
| Code status | Full code / DNR / DNR-DNI / comfort measures only (document date of discussion) |
| Fall precautions | Standard / high risk (specify interventions) |
| Isolation | Type: contact / droplet / airborne / none |

## Specialty-Specific Progress Note Additions

### Surgical Progress Note (Post-Op)

Add to standard note:
```
POST-OPERATIVE DAY: #[n] from [procedure name]
Surgical site: [Clean, dry, intact / erythema / drainage (characterize) / dehiscence]
Drains: [Type] - Output: [amount] mL/24h ([character: serous/serosanguinous/sanguinous/purulent])
Wound care: [Current orders]
Diet advancement: [Current diet -> plan to advance]
Activity: [Current level -> plan to advance]
Pain management: [Regimen, pain scores, opioid taper plan]
```

### Obstetric Progress Note

Add to standard note:
```
OBSTETRIC DATA:
Gestational age: [weeks + days]
Fetal status: FHR [rate], [reactive/nonreactive], [Category I/II/III]
Contractions: [frequency, duration, pattern]
Cervical exam: [dilation]/[effacement]/[station]/[position]
Membranes: [Intact / ruptured (time, fluid color)]
GBS status: [positive / negative / unknown]
```

### Psychiatric Progress Note

Add to standard note:
```
BEHAVIORAL OBSERVATIONS (24h):
  Sleep: [hours, quality, disturbances]
  Appetite: [% meals eaten]
  Participation: [group therapy / individual / refused]
  Behavioral events: [aggression, self-harm, elopement risk events]
  1:1 observation: [if applicable, findings]

MENTAL STATUS EXAM:
  Appearance: [grooming, dress]
  Behavior: [cooperative, guarded, agitated]
  Speech: [rate, rhythm, volume]
  Mood: "[patient's words]"
  Affect: [quality, range, congruence]
  Thought process: [organized, linear / tangential / disorganized]
  Thought content: [SI/HI assessment, delusions, obsessions]
  Perceptions: [hallucinations - A/V/T]
  Cognition: [orientation, attention]
  Insight/Judgment: [good / fair / poor]

SAFETY: SI: [denied / present (plan/intent/means)] | HI: [denied / present]
```

## Note Timing and Attestation

### Attending Attestation Requirements

| Setting | Requirement |
|---------|------------|
| Teaching hospital | Attending must personally see the patient and document a note or attest to the resident's note within 24 hours |
| Non-teaching | Attending documents their own note or co-signs mid-level notes per state law |
| Critical care | Attending must document direct involvement in critical care management |

### Timeliness Standards

| Note Type | Timing Requirement |
|-----------|-------------------|
| Admission H&P | Within 24 hours of admission |
| Daily progress note | Daily, before or shortly after rounds |
| Procedure note | Immediately after procedure |
| Discharge summary | Within 30 days (per TJC); best practice: within 48 hours |
| Transfer note | Before transfer |
| Death summary | Within 30 days |
| Operative note | Immediately after surgery (brief) + dictated note within 24 hours |
