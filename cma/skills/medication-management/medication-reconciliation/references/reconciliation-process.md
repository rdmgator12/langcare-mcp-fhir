# Medication Reconciliation Process Reference

## Joint Commission NPSG.03.06.01 Requirements

1. **Obtain** a complete and accurate list of medications the patient is currently taking (Best Possible Medication History - BPMH)
2. **Compare** the BPMH against medications ordered in the new care setting
3. **Resolve** any discrepancies with the prescriber
4. **Communicate** the updated list to the patient and next provider of care

## Reconciliation at Each Transition Point

### Admission
- Compare home medication list (MedicationStatement) against admission orders (MedicationRequest)
- Flag: home medications not ordered (intentional hold vs omission?), new medications started, dose changes

### Transfer
- Compare sending unit orders against receiving unit orders
- Flag: medications not continued, dose adjustments, new medications

### Discharge
- Compare inpatient orders against intended outpatient regimen
- Flag: new medications for patient education, discontinued medications, dose changes from pre-admission

## Discrepancy Categories

| Category | Definition | Risk Level |
|----------|-----------|------------|
| Omission | Home medication not ordered in new setting | High |
| Commission | Medication ordered with no prior history | Medium |
| Dose discrepancy | Same medication, different dose across sources | Medium |
| Frequency discrepancy | Same medication, different frequency | Medium |
| Route discrepancy | Same medication, different route | Medium |
| Therapeutic duplication | Two medications in same class | High |
| Drug-allergy conflict | Active prescription matching documented allergy | Critical |
| Drug-drug interaction | Clinically significant interaction in combined list | High |

## ISMP High-Alert Medications

These require independent double-check during reconciliation:

| Category | Examples |
|----------|----------|
| Anticoagulants | Warfarin, heparin, enoxaparin, apixaban, rivaroxaban, dabigatran, edoxaban |
| Insulin | All formulations (rapid, short, intermediate, long, mixed) |
| Opioids | Morphine, hydromorphone, fentanyl, oxycodone, methadone |
| Antiarrhythmics | Amiodarone, sotalol, dofetilide, flecainide |
| Chemotherapy | All agents |
| Concentrated electrolytes | KCl >20 mEq, NaCl >0.9%, magnesium sulfate IV |
| Neuromuscular blocking agents | Succinylcholine, rocuronium, vecuronium |
| Parenteral nutrition | TPN |

## WHO High 5s Protocol Steps

1. **Collect** BPMH using at least 2 sources (patient interview + pharmacy records/pill bottles)
2. **Verify** each medication: name, dose, route, frequency, indication
3. **Reconcile** against new orders within 24 hours of admission, within 4 hours of transfer to critical care
4. **Document** reconciliation outcome for each medication: continued, modified, discontinued, new
5. **Communicate** to patient and receiving provider in structured format
