---
name: langcare-follow-up-tasks
description: >
  Generates structured follow-up tasks from encounter data including pending
  results, referral tracking, medication monitoring, and screening reminders.
  Creates FHIR Task resources for care team tracking. Use when asked to
  generate follow-up tasks, create to-do list, post-visit tasks, pending
  follow-up items, or results management tasks.
---

# Follow-Up Task Generator

## When to Use This Skill
Use when a clinician needs to generate actionable follow-up tasks after an encounter, ensuring nothing falls through the cracks.

## Clinical Workflow
1. Use `fhir_search` to pull the recent Encounter and its associated resources
2. Use `fhir_search` to identify pending ServiceRequest resources (labs ordered but not resulted, imaging ordered, referrals sent)
3. Use `fhir_search` to identify new MedicationRequest resources requiring follow-up monitoring (new statin -> recheck LFTs in 6 weeks; new ACEi -> recheck BMP in 2 weeks)
4. Use `fhir_search` to identify Condition resources with time-based follow-up needs
5. Generate Task resources for each follow-up item with: description, owner, due date, priority, and status
6. Use `fhir_create` to create Task resources in FHIR

## FHIR Resources
- **Encounter** -- Visit context and associated resources
- **ServiceRequest** -- Pending orders to track
- **MedicationRequest** -- New medications requiring monitoring
- **Condition** -- Conditions requiring follow-up
- **Task** -- Output: follow-up task items with owner, due date, priority

## FHIR Query Examples
### Pull Pending Orders
```
fhir_search(resourceType="ServiceRequest", queryParams="patient=[patient-id]&status=active&_sort=-authored&_count=50")
```

### Create Follow-Up Task
```
fhir_create(resourceType="Task", resource={"resourceType":"Task","status":"requested","intent":"order","priority":"routine","description":"[task description]","for":{"reference":"Patient/[patient-id]"},"executionPeriod":{"end":"[due-date]"}})
```

## Clinical Guidelines
- AHRQ guidelines for results management and follow-up
- Test results management: all ordered tests must have a documented review and action
- Medication monitoring timelines per drug-specific guidelines

## Interpretation Guide
- Categorize tasks by type: pending results, medication monitoring, referral tracking, screening due, patient callback
- Assign priority: urgent (results expected to change management), routine (standard monitoring), low (administrative)
- Include specific due dates based on clinical guidelines (e.g., recheck potassium 1 week after starting ACEi, recheck A1c 3 months after medication change)
- Flag overdue tasks from prior encounters that were never completed

## Safety
- Never fabricate clinical data -- only report what FHIR returns
- Flag critical/abnormal values immediately
- Scope all FHIR queries to the authenticated patient
- Use standard terminology (LOINC, SNOMED CT, RxNorm, ICD-10)
- Present data in clinician-friendly format with reference ranges
