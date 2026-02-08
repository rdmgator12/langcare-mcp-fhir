# Mental Health Management Reference

## Safety Planning Template for Suicidal Ideation

Source: Stanley-Brown Safety Planning Intervention (SPI). Evidence-based intervention shown to reduce suicide attempts by ~50%.

### Safety Plan Structure

Complete collaboratively with the patient. Document in the EHR and provide a copy to the patient.

**Step 1: Warning Signs**
Identify thoughts, feelings, behaviors, or situations that precede a suicidal crisis:
- Thoughts: "I'm a burden", "No one cares", "Things will never get better"
- Feelings: hopelessness, overwhelming anxiety, anger, emptiness
- Behaviors: isolating, increased substance use, giving away possessions, researching methods
- Situations: anniversaries, interpersonal conflict, financial stress, job loss

**Step 2: Internal Coping Strategies**
Activities the patient can do alone to take their mind off the crisis (without contacting another person):
- Physical activity (walking, exercise)
- Relaxation techniques (deep breathing, progressive muscle relaxation)
- Distraction activities (music, reading, gaming, puzzles)
- Mindfulness or meditation
- Journaling

**Step 3: Social Contacts and Settings for Distraction**
People and places that provide healthy distraction (patient does not need to disclose the crisis):
- Name: [person 1] / Phone: [number]
- Name: [person 2] / Phone: [number]
- Places: coffee shop, gym, library, place of worship

**Step 4: People to Contact for Help**
Trusted individuals the patient can tell they are in crisis:
- Name: [person 1] / Phone: [number] / Relationship: [friend/family/sponsor]
- Name: [person 2] / Phone: [number] / Relationship: [friend/family/clergy]
- Clinician: [name] / Phone: [number] / After-hours: [number]

**Step 5: Professional and Agency Contacts**
- Therapist: [name] / Phone: [number]
- Psychiatrist: [name] / Phone: [number]
- 988 Suicide and Crisis Lifeline: call or text 988
- Crisis Text Line: text HOME to 741741
- Veterans Crisis Line: 988, press 1
- Local crisis center: [name] / Phone: [number]
- Nearest emergency department: [name] / Address: [address]

**Step 6: Lethal Means Restriction**
Actions to reduce access to lethal means:
- Firearms: remove from home, lock in safe with someone else holding the key, store at gun shop
- Medications: lock up, reduce supply (7-day prescriptions), give to trusted person
- Other means: secure knives, remove ligature points if relevant
- Document specific actions taken: ________________________________

### Safety Plan FHIR Documentation

Record as CarePlan resource:
```
Tool: fhir_create
resourceType: "CarePlan"
resource: {
  "resourceType": "CarePlan",
  "status": "active",
  "intent": "plan",
  "category": [{"coding": [{"system": "http://snomed.info/sct", "code": "736377005", "display": "Mental health safety plan"}]}],
  "subject": {"reference": "Patient/[patient-id]"},
  "created": "[date]",
  "activity": [
    {"detail": {"description": "Safety plan created. Warning signs: [list]. Coping strategies: [list]. Emergency contacts documented. Lethal means counseling completed."}}
  ]
}
```

## Referral Decision Matrix: Primary Care vs Psychiatry

### Manage in Primary Care

Appropriate when:
- Mild to moderate depression (PHQ-9 5-14) without suicidal ideation
- Mild to moderate anxiety (GAD-7 5-14) without panic attacks or severe avoidance
- First episode, uncomplicated presentation
- Patient preference for primary care management
- Good response to initial SSRI/SNRI trial
- No psychotic features
- No bipolar features (negative MDQ)
- No active substance use disorder complicating treatment

### Refer to Psychiatry

Refer when:
- Severe depression (PHQ-9 >= 20) or treatment-resistant depression (failed 2+ adequate medication trials)
- Active suicidal ideation with plan or intent (C-SSRS 4-5)
- Positive MDQ (bipolar concern) -- do not start antidepressant monotherapy
- Psychotic features (hallucinations, delusions, paranoia)
- Comorbid severe substance use disorder
- Complex medication regimen (multiple psychotropics, mood stabilizers)
- Diagnostic uncertainty (is it bipolar, schizoaffective, personality disorder?)
- Pregnancy/postpartum with moderate-severe psychiatric symptoms
- Prior psychiatric hospitalization
- Medication augmentation needed (lithium, atypical antipsychotics)
- ECT or TMS evaluation
- Age < 18 with suicidal ideation (refer to child/adolescent psychiatry)

### Refer to Psychotherapy (CBT, IPT, DBT, EMDR)

Refer when:
- Patient preference for non-pharmacologic treatment
- Mild-moderate depression or anxiety
- Trauma history (EMDR or CPT for PTSD)
- Personality disorder features (DBT for borderline)
- Chronic pain with psychological component
- Interpersonal difficulties contributing to symptoms
- Augmentation of medication management
- Relapse prevention after medication stabilization

## Medication Selection Guidance

### First-Line Antidepressants (SSRI/SNRI)

| Medication | Starting Dose | Therapeutic Dose | Key Considerations |
|-----------|--------------|-----------------|-------------------|
| Sertraline | 25-50 mg/day | 50-200 mg/day | First choice for most. GI side effects common initially. Safe in pregnancy (relative). |
| Escitalopram | 5-10 mg/day | 10-20 mg/day | Most selective SSRI. Fewest drug interactions. QTc prolongation at high doses. |
| Fluoxetine | 10-20 mg/day | 20-80 mg/day | Long half-life (good for adherence issues). More activating. Significant CYP2D6 inhibition. |
| Duloxetine (SNRI) | 30 mg/day | 60-120 mg/day | Dual action. Good for comorbid pain (neuropathic, fibromyalgia, chronic musculoskeletal). |
| Venlafaxine (SNRI) | 37.5-75 mg/day | 75-225 mg/day | Dual action. Good for comorbid anxiety. Monitor BP at higher doses. Difficult to taper. |

### Second-Line and Adjunctive Options

| Medication | Indication | Key Notes |
|-----------|-----------|-----------|
| Bupropion | Depression with fatigue/low motivation; smoking cessation; SSRI-induced sexual dysfunction | Activating. Contraindicated in seizure disorder, bulimia. Weight-neutral. No sexual side effects. |
| Mirtazapine | Depression with insomnia and weight loss | Sedating (use at bedtime). Weight gain. Increases appetite. Good for elderly with poor appetite. |
| Buspirone | GAD (adjunct or monotherapy) | Non-addictive. Takes 2-4 weeks for effect. No withdrawal. Good adjunct to SSRI for residual anxiety. |
| Hydroxyzine | Acute anxiety (PRN) | Antihistamine. Non-addictive alternative to benzodiazepines. Sedating. |
| Trazodone | Insomnia (low dose) | 25-100 mg at bedtime. Not effective as antidepressant at low doses. Priapism risk (rare). |

### Medications to Avoid or Use with Caution

| Situation | Avoid | Reason |
|-----------|-------|--------|
| Positive MDQ / suspected bipolar | Antidepressant monotherapy | Risk of mania switch. Use mood stabilizer first. |
| Elderly (>= 65) | Paroxetine, amitriptyline, TCAs | Anticholinergic burden (Beers criteria). |
| Elderly | Benzodiazepines | Fall risk, cognitive impairment, paradoxical agitation. |
| Substance use disorder | Benzodiazepines | High addiction potential. Use non-addictive alternatives. |
| Pregnancy | Paroxetine (1st trimester) | Cardiac malformation risk. Sertraline preferred. |
| Suicidal ideation | Large TCA supply | Lethal in overdose. Limit supply to 7-day prescriptions. |
| QTc prolongation risk | High-dose citalopram (>40mg), escitalopram (>20mg) | QTc prolongation. Obtain ECG if concerns. |

### Adequate Trial Definition

Before declaring treatment failure:
- Minimum 4-6 weeks at therapeutic dose
- Confirmed adherence (ask directly, check pharmacy fills)
- Adequate dose titration (not still on starting dose)
- If partial response at 4-6 weeks: increase dose or augment before switching

### Treatment-Resistant Depression Criteria

- Failed >= 2 adequate antidepressant trials from different classes
- Each trial: adequate dose for >= 6 weeks with confirmed adherence
- Augmentation strategies before declaring resistance: add bupropion, add buspirone, add atypical antipsychotic (aripiprazole, quetiapine), add lithium
- Consider: esketamine nasal spray, TMS, ECT

## Crisis Resources

### National Resources (United States)

| Resource | Contact | Hours |
|----------|---------|-------|
| 988 Suicide and Crisis Lifeline | Call or text 988 | 24/7 |
| Crisis Text Line | Text HOME to 741741 | 24/7 |
| Veterans Crisis Line | 988 then press 1, or text 838255 | 24/7 |
| SAMHSA National Helpline | 1-800-662-4357 | 24/7 |
| National Domestic Violence Hotline | 1-800-799-7233 | 24/7 |
| Trevor Project (LGBTQ+ youth) | 1-866-488-7386 or text START to 678-678 | 24/7 |
| National Alliance on Mental Illness (NAMI) | 1-800-950-6264 | M-F 10am-10pm ET |
| Trans Lifeline | 1-877-565-8860 | 24/7 |
| Childhelp National Child Abuse Hotline | 1-800-422-4453 | 24/7 |

### Emergency Criteria for Psychiatric Hospitalization

Involuntary hold criteria (vary by state, general framework):
- Imminent danger to self (active suicidal plan with intent and means)
- Imminent danger to others (active homicidal ideation with plan and intent)
- Gravely disabled (unable to provide for basic needs due to mental illness)

Voluntary admission recommended when:
- Patient acknowledges need for higher level of care
- Outpatient safety plan insufficient
- Acute psychosis or mania
- Need for medication stabilization in monitored setting
- Failed outpatient management with escalating risk

## Monitoring Parameters During Psychotropic Treatment

### SSRI/SNRI Monitoring

| Parameter | Timing |
|-----------|--------|
| Suicidal ideation (PHQ-9 item 9, C-SSRS) | Every visit, especially first 4 weeks and dose changes (FDA black box for age < 25) |
| Side effects | Every visit for first 3 months |
| Symptom improvement (PHQ-9, GAD-7) | Every 2-4 weeks until stable, then every 3 months |
| Sexual function | Ask at each visit (common reason for non-adherence) |
| Weight | Quarterly |
| Sodium level | Consider in elderly, especially with diuretics (SIADH risk) |

### Lithium Monitoring

| Parameter | Timing |
|-----------|--------|
| Lithium level | 5-7 days after initiation or dose change; then every 3-6 months |
| BMP (creatinine, electrolytes) | Baseline, then every 3-6 months |
| TSH | Baseline, then every 6-12 months (hypothyroidism risk) |
| Calcium | Baseline, then annually (hyperparathyroidism risk) |
| Urinalysis | Baseline, then annually (nephrogenic DI risk) |
| ECG | Baseline if > 40 years old |
| Pregnancy test | Before initiation in women of childbearing age (teratogenic -- Ebstein anomaly) |

### Valproate Monitoring

| Parameter | Timing |
|-----------|--------|
| Valproate level | 5-7 days after dose change; target 50-125 mcg/mL |
| CBC with differential | Baseline, then every 6 months (thrombocytopenia risk) |
| LFTs | Baseline, then every 6 months (hepatotoxicity risk) |
| Ammonia level | If mental status changes (hyperammonemia risk) |
| Pregnancy test | Before initiation (highly teratogenic -- neural tube defects) |
| Weight | Quarterly (weight gain common) |

### Atypical Antipsychotic Monitoring

| Parameter | Timing |
|-----------|--------|
| Fasting glucose + lipids | Baseline, 12 weeks, then annually (metabolic syndrome risk) |
| Weight/BMI | Monthly for first 3 months, then quarterly |
| Blood pressure | Every visit |
| HbA1c | Annually if on metabolically active agents (olanzapine, clozapine, quetiapine) |
| Prolactin | If symptoms of hyperprolactinemia (galactorrhea, amenorrhea, sexual dysfunction) |
| ECG | If QTc-prolonging agent (ziprasidone) or patient risk factors |
| AIMS (tardive dyskinesia screen) | Every 6 months (typical antipsychotics) or annually (atypicals) |
