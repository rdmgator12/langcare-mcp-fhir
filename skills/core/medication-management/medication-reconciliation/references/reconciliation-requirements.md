# Joint Commission NPSG.03.06.01 - Medication Reconciliation Requirements

## National Patient Safety Goal NPSG.03.06.01

**Statement**: "Maintain and communicate an accurate patient medication list."

**Rationale**: Medication discrepancies at care transitions account for up to 50% of all medication errors in hospitals and contribute to approximately 20% of adverse drug events. The Joint Commission requires organizations to reconcile medications across the continuum of care.

## Required Reconciliation Points

### 1. Admission
- Obtain and document a complete list of current medications upon admission
- Compare the admission list with medications ordered at admission
- Resolve discrepancies with the prescriber

### 2. Transfer (Internal)
- Compare medications from the sending unit with those ordered for the receiving unit
- Resolve discrepancies before or within a specified time frame after transfer

### 3. Discharge
- Compare the admission medication list, current inpatient orders, and planned discharge medications
- Provide the patient/caregiver with the reconciled discharge medication list
- Explain any changes (new, discontinued, changed dose)

### 4. Outpatient Encounters
- Verify and update the medication list at each visit
- Compare with any new prescriptions or changes

## Reconciliation Process Steps

### Step 1: Verification (Best Possible Medication History - BPMH)
**Goal**: Obtain the most accurate list of all medications the patient is currently taking.

**Sources** (in order of reliability):
1. Patient/caregiver interview
2. Community pharmacy records (MedicationDispense)
3. Prior medical records (MedicationRequest, MedicationStatement)
4. Medication vials/packages brought by patient
5. Prescription monitoring program data

**Required Data Points per Medication**:
- Drug name (generic preferred)
- Dose
- Route
- Frequency
- Last dose taken (date/time)
- Prescriber
- Indication (when available)
- Duration of therapy
- Adherence assessment (taking as prescribed?)

### Step 2: Clarification
**Goal**: Ensure each medication and dose is appropriate for the patient.

**Checks**:
- Is the drug/dose/route/frequency correct?
- Are there allergies or intolerances?
- Are there drug-drug interactions?
- Are there drug-disease contraindications?
- Is renal/hepatic dose adjustment needed?
- Is the patient actually taking the medication (adherence)?

### Step 3: Reconciliation
**Goal**: Compare the verified list against the current orders and resolve all discrepancies.

**Discrepancy Categories**:

| Category | Definition | Action Required |
|----------|-----------|-----------------|
| Intentional documented | Prescriber deliberately changed and documented reason | No further action |
| Intentional undocumented | Prescriber deliberately changed but did not document reason | Document rationale |
| Unintentional omission | Medication not ordered but should have been | Contact prescriber to add |
| Unintentional commission | Medication ordered but should not have been | Contact prescriber to remove |
| Unintentional dose change | Dose differs without explanation | Contact prescriber to clarify |
| Unintentional duplication | Therapeutic duplication without rationale | Contact prescriber to resolve |

### Step 4: Communication
**Goal**: Communicate the reconciled list to the patient and the next provider of care.

**Documentation Requirements**:
- Final reconciled medication list
- All discrepancies identified and how they were resolved
- Person who performed the reconciliation
- Date and time of reconciliation
- Provider who approved the final list

## FHIR Mapping for Reconciliation Documentation

### MedicationStatement for Reconciled List
```json
{
  "resourceType": "MedicationStatement",
  "status": "active",
  "statusReason": [{"text": "Reconciled"}],
  "category": {
    "coding": [{
      "system": "http://terminology.hl7.org/CodeSystem/medication-statement-category",
      "code": "patientspecified",
      "display": "Patient Specified"
    }]
  },
  "medicationCodeableConcept": {
    "coding": [{
      "system": "http://www.nlm.nih.gov/research/umls/rxnorm",
      "code": "[RxNorm-code]",
      "display": "[medication-name]"
    }]
  },
  "subject": {"reference": "Patient/[id]"},
  "dateAsserted": "[reconciliation-date]",
  "informationSource": {"reference": "Practitioner/[reconciler-id]"},
  "dosage": [{
    "text": "[dose] [route] [frequency]",
    "timing": {"code": {"text": "[frequency]"}},
    "route": {"coding": [{"system": "http://snomed.info/sct", "code": "[route-code]"}]},
    "doseAndRate": [{"doseQuantity": {"value": "[dose]", "unit": "[unit]"}}]
  }]
}
```

### DetectedIssue for Discrepancies
```json
{
  "resourceType": "DetectedIssue",
  "status": "preliminary",
  "code": {
    "coding": [{
      "system": "http://terminology.hl7.org/CodeSystem/v3-ActCode",
      "code": "DUPTHPY",
      "display": "Duplicate Therapy Alert"
    }]
  },
  "patient": {"reference": "Patient/[id]"},
  "identifiedDateTime": "[date]",
  "implicated": [
    {"reference": "MedicationRequest/[id1]"},
    {"reference": "MedicationStatement/[id2]"}
  ],
  "detail": "Dose discrepancy: home lisinopril 10mg vs ordered lisinopril 20mg",
  "mitigation": [{
    "action": {"text": "Confirmed dose increase is intentional per prescriber"},
    "date": "[date]",
    "author": {"reference": "Practitioner/[id]"}
  }]
}
```

## Performance Measures

### Joint Commission Survey Focus Areas
1. Process for obtaining medication history is standardized
2. Reconciliation occurs at all required transition points
3. Discrepancies are documented and resolved
4. Patient/caregiver receives written discharge medication list
5. Responsibility for reconciliation is clearly assigned

### Quality Metrics
- **Reconciliation completion rate**: % of admissions/transfers/discharges with documented reconciliation
- **Discrepancy identification rate**: Average number of discrepancies found per reconciliation
- **Time to reconciliation**: Hours from admission to completed reconciliation
- **Unresolved discrepancy rate**: % of discrepancies not resolved within 24 hours

## Common Medication Discrepancy Patterns

### High-Frequency Omissions at Admission
1. OTC medications (vitamins, supplements, antacids)
2. PRN medications (pain, sleep, allergy)
3. Topical medications (creams, eye drops, inhalers)
4. Medications prescribed by other providers
5. Recently changed medications

### High-Risk Discrepancies
1. Anticoagulant dose changes without documentation
2. Insulin regimen differences (basal vs home regimen)
3. Opioid dose/frequency changes
4. Antiepileptic medication omissions
5. Immunosuppressant dose changes

### Common Root Causes
1. Incomplete medication history at admission
2. Multiple prescribers with overlapping formularies
3. Patient not aware of medication changes
4. Pharmacy records not available
5. Generic vs brand name confusion
6. Medications held pre-procedure not restarted
