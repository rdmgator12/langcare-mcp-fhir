# I-PASS Handoff Framework Reference

## I-PASS Components

### I -- Illness Severity

| Level | Definition | Examples |
|-------|-----------|----------|
| **Stable** | Patient is in routine status; no anticipated changes in management | Awaiting discharge, stable chronic conditions |
| **Watcher** | Patient has potential for clinical deterioration; close monitoring needed | New medication titration, post-procedure observation, borderline vitals |
| **Unstable** | Patient is actively decompensating or requires urgent/emergent intervention | Active sepsis, respiratory failure, hemodynamic instability, new acute event |

### P -- Patient Summary

Structure as a brief narrative:
1. **One-liner**: [Age] [Sex] with [primary problem] admitted for [reason] on [date]
2. **Hospital course**: Key events, procedures, medication changes
3. **Ongoing issues**: Active problems requiring continued management
4. **Baseline status**: Functional status, code status, allergies, isolation

Example:
> 72F with CHF (EF 25%), CKD 3b, AFib admitted for CHF exacerbation on 3/12. Diuresed 4L with IV furosemide, converted to oral 40mg BID. Creatinine stable at 1.8. On telemetry, rate-controlled on metoprolol. DNR/DNI. NKDA.

### A -- Action List

| Format | Description |
|--------|------------|
| Task | What needs to be done |
| Owner | Who is responsible |
| Timeline | When it needs to happen |
| Contingency | What to do if the expected result does not occur |

Example Actions:
- [ ] Recheck BMP in AM (owner: night intern) -- if K+ > 5.5, hold spironolactone and recheck in 4h
- [ ] Daily weight in AM (owner: nursing) -- if weight up >1kg, give IV furosemide 40mg
- [ ] Cardiology to see patient in AM for device evaluation (owner: day team to follow up)
- [ ] Repeat CXR in AM if still dyspneic (owner: night intern)

### S -- Situation Awareness and Contingency Planning

| What to Watch For | If This Happens | Then Do This |
|-------------------|-----------------|--------------|
| Increasing dyspnea or SpO2 <90% | Respiratory decompensation | Give IV furosemide 40mg, CXR, ABG, call attending |
| Heart rate >130 or new irregular rhythm | Rapid AFib or new arrhythmia | 12-lead ECG, check K+/Mg, consider metoprolol IV, call cardiology |
| UOP <0.5 mL/kg/hr x 6 hours | Acute kidney injury from diuresis | Hold diuretics, check BMP, volume assessment |
| New confusion or agitation | Delirium | Check vitals, glucose, UOP, review medications, avoid benzos |
| Temperature >38.5C | New fever | Blood cultures x2, UA+Cx, CXR, CBC, lactate |

### S -- Synthesis by Receiver

The receiving clinician:
1. **Reads back** key elements (illness severity, main problem, action items)
2. **Asks questions** to clarify any ambiguity
3. **Confirms** understanding of contingency plans
4. **Documents** receipt of handoff (time, participants)

## Essential Handoff Data Elements

### Always Include
- [ ] Patient name, MRN, location, age/sex
- [ ] Primary team and attending
- [ ] Illness severity (stable/watcher/unstable)
- [ ] Admission diagnosis and date
- [ ] Code status (Full code / DNR / DNI / DNR-DNI / COMFORT)
- [ ] Allergies (with reaction type)
- [ ] Isolation precautions (contact, droplet, airborne, none)
- [ ] Fall risk level
- [ ] Active medication list
- [ ] Recent vital signs (last set)
- [ ] Key labs (today's results and any pending)
- [ ] Action list with timelines
- [ ] Contingency plans for anticipated problems
- [ ] Lines, drains, devices (IV access, Foley, O2, etc.)

### Situation-Specific
- [ ] Pending consults and their status
- [ ] Pending imaging or procedures
- [ ] Anticipated discharge date
- [ ] Family/caregiver communication needs
- [ ] Diet and activity orders
- [ ] DVT prophylaxis status

## Evidence Base

Starmer AJ et al. "Changes in Medical Errors after Implementation of a Handoff Program." NEJM 2014;371:1803-12.

Key findings:
- 23% reduction in medical errors
- 30% reduction in preventable adverse events
- No significant change in handoff duration when structured
