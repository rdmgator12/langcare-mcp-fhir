# Task Prioritization and Assignment Logic

## Urgency Classification Criteria

### STAT Priority (Same-Day Action)

FHIR Task priority: `stat`

#### Critical Lab Values Requiring Immediate Action

| Lab | Critical Value | LOINC | Immediate Action |
|-----|---------------|-------|-----------------|
| Potassium | > 6.5 mEq/L | 2823-3 | ECG, calcium gluconate, insulin/dextrose, recheck in 1-2 hrs |
| Potassium | < 2.5 mEq/L | 2823-3 | IV potassium, cardiac monitoring, recheck in 2-4 hrs |
| Sodium | < 120 mEq/L | 2951-2 | Fluid restriction or hypertonic saline, recheck q4-6h |
| Sodium | > 160 mEq/L | 2951-2 | Free water replacement, recheck q4-6h |
| Glucose | < 40 mg/dL | 2339-0 | Dextrose IV or glucagon, recheck in 15 min |
| Glucose | > 600 mg/dL | 2339-0 | Insulin, IV fluids, recheck hourly |
| Hemoglobin | < 7.0 g/dL (acute drop) | 718-7 | Type and cross, transfusion, find source |
| Platelets | < 20,000/uL | 777-3 | Hematology consult, transfusion consideration |
| WBC | < 500/uL (ANC) | 751-8 | Neutropenic precautions, infection workup |
| INR | > 9.0 | 6301-6 | Hold anticoagulant, vitamin K, consider PCC |
| Troponin | New positive | 6598-7 | Cardiology consult, serial monitoring, ECG |
| Lactate | > 4.0 mmol/L | 2524-7 | Sepsis workup, fluid resuscitation |
| pH (arterial) | < 7.20 or > 7.60 | 2744-1 | ABG, targeted treatment |
| pO2 (arterial) | < 60 mmHg | 2703-7 | Supplemental oxygen, respiratory assessment |
| Calcium (corrected) | > 14.0 mg/dL | 17861-6 | IV fluids, calcitonin |
| Calcium (corrected) | < 6.5 mg/dL | 17861-6 | IV calcium, cardiac monitoring |
| Creatinine | > 2x baseline (acute) | 2160-0 | Fluid assessment, hold nephrotoxics, renal consult |

#### Critical Clinical Findings
- New positive blood culture
- New pulmonary embolism on imaging
- New stroke finding on imaging
- Pathology showing unexpected malignancy
- Imaging showing surgical emergency (free air, dissection, torsion)

### ASAP Priority (Within 48 Hours)

FHIR Task priority: `asap`

| Trigger | Action Required | Due |
|---------|----------------|-----|
| Abnormal results requiring treatment change | Provider review and order modification | 24-48 hours |
| Post-discharge critical medication monitoring | Lab draw (e.g., potassium on new spironolactone) | 48 hours |
| Positive culture with sensitivity results | Antibiotic adjustment | 24 hours |
| Pending results from urgently ordered tests | Review when available | When resulted |
| Post-procedure complication concern | Re-evaluation | 24-48 hours |
| Missed critical medication refill | Prescription renewal | 24 hours |

### Urgent Priority (Within 1 Week)

FHIR Task priority: `urgent`

| Trigger | Action Required | Due |
|---------|----------------|-----|
| Significant lab abnormality (see "Significant" table in follow-up-intervals.md) | Lab recheck | 3-7 days |
| New diagnosis requiring education (diabetes, CHF) | Education referral | Within 1 week |
| Post-discharge follow-up for high-risk patient | PCP visit | 3-7 days |
| Medication requiring early monitoring lab | Lab draw | Per medication schedule |
| Specialist referral for significant finding | Referral placement | Within 1 week |
| Biopsy results follow-up | Provider review | When resulted (typically 5-7 days) |
| Post-ED visit follow-up | PCP visit | 2-3 days |

### Routine Priority (Within 1 Month)

FHIR Task priority: `routine`

| Trigger | Action Required | Due |
|---------|----------------|-----|
| Mild lab abnormality | Lab recheck | 2-4 weeks |
| New statin -- lipid recheck | Lab draw | 6-8 weeks |
| New thyroid medication -- TSH recheck | Lab draw | 6-8 weeks |
| HbA1c recheck (stable diabetes) | Lab draw | 3 months |
| Annual screening due | Schedule screening | Within recommended window |
| Preventive care gap | Schedule appointment | Per screening interval |
| Medication refill due | Prescription renewal | Before current supply runs out |
| Chronic disease monitoring labs | Lab draw | Per disease interval |

---

## Responsible Party Assignment Logic

### Decision Tree

```
START
  |
  Is this a critical/STAT result?
  |-- YES --> Ordering provider (or covering provider if after hours)
  |-- NO
      |
      Is this a medication monitoring task?
      |-- YES --> Prescribing provider
      |-- NO
          |
          Is this a post-procedure follow-up?
          |-- YES --> Performing provider / surgeon
          |-- NO
              |
              Is this a specialist-managed condition?
              |-- YES --> Managing specialist
              |-- NO
                  |
                  Is the patient's PCP identifiable?
                  |-- YES --> PCP
                  |-- NO --> Assign to encounter attending / flag for manual assignment
```

### Provider Role Identification from FHIR

| Role | FHIR Source | Identifier |
|------|------------|-----------|
| Ordering provider | ServiceRequest.requester, Observation.performer | Reference to Practitioner |
| Prescribing provider | MedicationRequest.requester | Reference to Practitioner |
| Performing provider | Procedure.performer.actor | Reference to Practitioner |
| Attending physician | Encounter.participant (type = "ATND") | Reference to Practitioner |
| PCP | Patient.generalPractitioner | Reference to Practitioner/Organization |
| Consulting specialist | Encounter.participant (type = "CON") | Reference to Practitioner |
| Covering provider | Not typically in FHIR | May need external lookup |

### Fallback Assignment Rules

| Scenario | Fallback Assignment | Notes |
|----------|-------------------|-------|
| Ordering provider unavailable | Encounter attending | Most common fallback |
| PCP not identified | Encounter attending | Flag for manual assignment |
| Specialist not in FHIR | Create task without owner, flag for assignment | Note specialty needed |
| After-hours critical result | Covering provider (may not be in FHIR) | Note urgency in task description |
| Patient has no assigned providers | Flag as unassigned | Requires manual triage |

---

## Escalation Pathways

### Lab Value Escalation

| Level | Criteria | Notification | Timeline |
|-------|----------|-------------|----------|
| Level 1 (Critical) | Values in STAT table above | Direct provider call, verbal read-back | Immediate (within 30 minutes of result) |
| Level 2 (Significant) | Markedly abnormal, treatment-affecting | Secure message or callback request | Within 4 hours |
| Level 3 (Mild) | Outside reference range, monitoring needed | Task creation, inbox message | Within 24 hours |
| Level 4 (Informational) | Trending, no immediate action | Include in next visit summary | Next encounter |

### Task Overdue Escalation

| Time Past Due | Escalation Action |
|--------------|-------------------|
| STAT task > 2 hours overdue | Escalate to department chief / medical director |
| ASAP task > 48 hours overdue | Re-notify assigned provider + supervisor |
| Urgent task > 7 days overdue | Notify assigned provider + PCP + quality team |
| Routine task > 30 days overdue | Notify assigned provider, consider reassignment |

### Clinical Deterioration Escalation

| Indicator | Action | Notify |
|-----------|--------|--------|
| Repeat critical lab value after correction | Consider ICU transfer, subspecialty consult | Attending + specialist |
| New critical value on floor patient | Rapid response team activation | Bedside nurse + charge nurse + attending |
| Post-discharge patient with critical result | Direct patient call, may need ED referral | PCP + discharging provider |
| Failure to improve on current treatment | Reassess diagnosis and plan | Attending, consider additional consults |

---

## Task Status Lifecycle

### FHIR Task Status Flow

```
requested --> accepted --> in-progress --> completed
    |             |            |
    +-> rejected  +-> on-hold  +-> failed
    |
    +-> cancelled
```

### Status Definitions

| Status | Definition | When to Use |
|--------|-----------|-------------|
| requested | Task has been created and is awaiting acceptance | Initial state when auto-generated |
| accepted | Provider has acknowledged the task | Provider reviews and accepts |
| in-progress | Work on the task has begun | Lab order placed, appointment scheduled |
| completed | Task has been finished | Result reviewed, appointment attended |
| rejected | Provider has declined the task | Wrong provider, not clinically indicated |
| cancelled | Task is no longer needed | Patient transferred, condition resolved |
| on-hold | Task is paused | Awaiting prerequisite, patient preference |
| failed | Task could not be completed | Unable to draw labs, patient no-show |

### Auto-Status Updates
- When a referenced lab result (Observation) moves from `registered` to `final`, the associated "review result" Task can auto-update to `in-progress`
- When a referenced Appointment status changes to `fulfilled`, the associated "follow-up visit" Task can auto-update to `completed`
- When a patient is admitted or transferred, outpatient Tasks may be placed `on-hold`

---

## Task Deduplication Rules

Before creating a new Task, check for existing Tasks to avoid duplicates:

```
Tool: fhir_search
resourceType: "Task"
queryParams: "patient=[patient-id]&status=requested,accepted,in-progress&code=[task-type]"
```

### Deduplication Criteria

| Scenario | Rule |
|----------|------|
| Same lab recheck already requested | Do not duplicate if existing Task has same focus (Observation code) and due date within 7 days |
| Same appointment type already scheduled | Do not duplicate if existing Task references same specialty |
| Same referral already in progress | Do not duplicate if ServiceRequest already exists for same specialty |
| Prior Task was completed | Create new Task (this is a new cycle) |
| Prior Task was cancelled or rejected | Create new Task with updated context |
