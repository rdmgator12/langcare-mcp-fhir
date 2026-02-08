# H&P Requirements, Risk Factor Documentation, and Specialty Templates

## Regulatory Requirements for Admission H&P

### CMS Conditions of Participation (CoP) -- 42 CFR 482.24(c)(2)

The medical record must contain:
1. A medical history and physical examination completed and documented **no more than 30 days before** or **within 24 hours after** admission
2. If the H&P was completed within 30 days before admission, an **update** must be completed within 24 hours after admission documenting any changes

**Required H&P elements per CMS:**
- Chief complaint or reason for admission
- History of present illness
- Relevant past medical, surgical, and family history
- Current medications
- Allergies
- Review of systems relevant to the presenting problem
- Physical examination
- Conclusions and assessment
- Plan of care

### The Joint Commission (TJC) Standards -- PC.01.02.03

- H&P must be completed **within 24 hours of admission**
- Must be performed by a **physician, oral/maxillofacial surgeon, or qualified practitioner** per medical staff privileges
- If performed prior to admission (within 30 days), an **update assessment** documenting any changes must be completed within 24 hours
- Update must include: examination for changes, note of any interval changes, indication that the H&P was reviewed

### State-Specific Variations

Some states have stricter requirements:
- **California**: H&P within 24 hours, no pre-admission allowance for some facility types
- **New York**: H&P within 24 hours; surgical patients require H&P before anesthesia
- **Florida**: H&P within 24 hours of admission; must include functional assessment for patients 60+

### Pre-Operative H&P Requirements

- Must be completed before **any procedure requiring anesthesia**
- If performed > 24 hours before surgery: requires update/addendum
- Must include: **airway assessment**, **anesthesia risk** (ASA class), **NPO status**, **medication review** (anticoagulants, insulin, etc.)
- Surgical site must be **marked** per Universal Protocol

## Risk Factor Documentation Requirements

### Cardiovascular Risk Factors

Document the following for any patient with or at risk for cardiovascular disease:

| Risk Factor | Documentation Elements | FHIR Source |
|------------|----------------------|-------------|
| Hypertension | Current BP, history, medications, control status | Observation (vital-signs), Condition, MedicationRequest |
| Diabetes | Type, HbA1c, glucose management, complications | Condition, Observation (laboratory) |
| Hyperlipidemia | Most recent lipid panel, statin therapy | Observation (laboratory), MedicationRequest |
| Tobacco use | Current/former/never, pack-years, quit date | Observation (social-history, LOINC 72166-2) |
| Family history of premature CAD | First-degree relative with CAD <55 (M) or <65 (F) | FamilyMemberHistory |
| Obesity | BMI, waist circumference | Observation (vital-signs, LOINC 39156-5) |
| Physical inactivity | Exercise frequency/type | Observation (social-history) |
| Prior CVD events | MI, stroke, PAD, revascularization | Condition, Procedure |

### Fall Risk Factors

Required for all hospitalized patients (TJC NPSG.09.02.01):

| Factor | Assessment |
|--------|-----------|
| Age > 65 | From Patient.birthDate |
| History of falls | Prior fall-related Conditions or Procedures |
| Gait/balance impairment | Observation (exam findings) |
| Medications (sedatives, antihypertensives, opioids, psychotropics) | MedicationRequest review |
| Cognitive impairment | Condition (dementia), Observation (cognitive assessment) |
| Visual impairment | Condition |
| Urinary frequency/incontinence | Condition |
| Orthostatic hypotension | Observation (vital-signs) |

### VTE Risk Factors

Required documentation per CMS core measure:

| Factor | Source |
|--------|--------|
| Active cancer | Condition (malignancy codes) |
| Prior VTE | Condition (I82.x, I26.x) |
| Immobility > 3 days | CarePlan (activity orders) |
| Known thrombophilia | Condition |
| Recent surgery (< 30 days) | Procedure |
| Age > 40 | Patient.birthDate |
| Obesity (BMI > 30) | Observation (vital-signs) |
| Hormone therapy / OCP | MedicationRequest |
| Central venous catheter | Procedure |
| Pregnancy / postpartum | Condition, Observation (LOINC 82810-3) |

### Infection Risk Documentation

| Factor | Assessment |
|--------|-----------|
| Immunocompromised | Condition (HIV, transplant), MedicationRequest (immunosuppressants, steroids, chemotherapy) |
| Indwelling devices | Procedure (central line, foley, trach) with active dates |
| Recent hospitalization | Encounter history (< 90 days) |
| Recent antibiotic exposure | MedicationRequest (antibiotics in last 90 days) |
| MDRO history | Condition or Observation (prior positive cultures for MRSA, VRE, ESBL, CRE) |
| Surgical wounds | Procedure (recent surgical procedures with wound classification) |

## Specialty-Specific H&P Templates

### Cardiology Admission H&P

**Additional required elements beyond standard H&P:**

| Section | Specific Elements |
|---------|------------------|
| HPI | NYHA functional class, exercise tolerance, orthopnea (number of pillows), PND, weight gain timeline, dietary indiscretions, medication compliance |
| PMH | Prior catheterizations, interventions (PCI, CABG, valve replacements), device history (pacemaker, ICD, CRT), prior echocardiographic data (EF trend) |
| Cardiac Review | Chest pain characterization (typical/atypical), palpitations, presyncope/syncope, claudication, Raynaud's |
| Physical Exam | JVP measurement (cm H2O), hepatojugular reflux, cardiac auscultation (all valve areas, patient positioning), S3/S4, murmur characterization (grade, radiation, dynamic maneuvers), peripheral pulses (all four extremities), edema grading, ABI if applicable |
| Assessment | NYHA class assignment, ACC/AHA heart failure stage (A-D), hemodynamic profile (warm/cold, wet/dry) |

### Pulmonary Admission H&P

| Section | Specific Elements |
|---------|------------------|
| HPI | Dyspnea severity (mMRC scale 0-4), O2 requirement (liters, device), sputum production (quantity, color, consistency), exercise limitation, prior intubations |
| PMH | PFT results (FEV1, FVC, DLCO), sleep study results (AHI), home O2 prescription, CPAP/BiPAP settings, prior exacerbation history and frequency |
| Exposures | Smoking (pack-years calculation), occupational exposures, environmental exposures, travel history |
| Physical Exam | Respiratory rate and pattern, accessory muscle use, pursed-lip breathing, chest configuration (barrel chest, kyphosis), auscultation (wheeze character, crackle type -- fine vs coarse, location), tactile fremitus, percussion, diaphragmatic excursion |

### Surgical Admission H&P

| Section | Specific Elements |
|---------|------------------|
| Pre-op Assessment | Procedure planned, indication, laterality/site, prior related surgeries |
| Anesthesia Risk | ASA classification (I-VI), Mallampati score (I-IV), prior anesthesia complications, family history of malignant hyperthermia |
| Bleeding Risk | Anticoagulant/antiplatelet status and hold timing, coagulation labs, bleeding history, blood type and screen |
| Functional Status | Baseline functional capacity (METs), ADL independence, mobility status |
| NPO Status | Last oral intake (solids, liquids), medication timing |
| Consents | Surgical consent, anesthesia consent, blood product consent |
| Physical Exam Additions | Airway assessment (mouth opening, neck mobility, thyromental distance, dentition), surgical site examination, vascular access assessment |

### Psychiatric Admission H&P

| Section | Specific Elements |
|---------|------------------|
| Psychiatric History | Prior diagnoses, hospitalizations, suicide attempts (method, lethality, circumstances), self-harm history, violence history |
| Current Symptoms | PHQ-9 or GAD-7 scores, psychotic symptoms (hallucinations - type, content; delusions - type; disorganized thought), manic symptoms (YMRS if applicable) |
| Substance Use | Detailed history per substance (type, route, amount, frequency, last use, withdrawal history, treatment history), CAGE/AUDIT scores |
| Safety Assessment | Current SI (plan, intent, means, timeline), HI (target, plan), access to firearms/lethal means, protective factors |
| Legal | Voluntary vs involuntary status, guardian/conservator, court orders |
| Mental Status Exam | Full MSE including appearance, behavior, speech, mood, affect, thought process, thought content, perceptions, cognition (orientation, attention, memory, abstraction), insight, judgment |
| Social Assessment | Housing (stable/unstable/homeless), employment, relationships, support system, legal issues, financial stressors |

### Obstetric Admission H&P

| Section | Specific Elements |
|---------|------------------|
| OB History | G_P_ (gravida, para, term, preterm, abortions, living), EDD and dating criteria, GBS status, blood type and antibody screen, prenatal labs summary |
| Current Pregnancy | Complications (GDM, preeclampsia, IUGR, preterm labor), fetal status, ultrasound findings |
| Labor Assessment | Contraction pattern, membrane status (intact/SROM/AROM, time, fluid color), cervical exam (dilation, effacement, station, position), Bishop score |
| Fetal Assessment | FHR baseline, variability, accelerations, decelerations, category (I/II/III) |
| Risk Factors | Prior cesarean (type of incision), placenta location, fetal presentation, maternal comorbidities |

### Pediatric Admission H&P

| Section | Specific Elements |
|---------|------------------|
| Birth History | Gestational age, birth weight, APGAR scores, delivery complications, NICU stay |
| Development | Developmental milestones (gross motor, fine motor, language, social), school performance, behavioral concerns |
| Growth | Weight, height, head circumference (plot on percentile curves), growth trajectory |
| Immunizations | Current immunization status, any missed vaccines |
| Social History | Family structure, daycare/school, child protective concerns, exposure history |
| Physical Exam | Age-specific exam (fontanelle in infants, Tanner staging in adolescents, developmental assessment) |

## H&P Update Requirements

When a pre-admission H&P exists (completed within 30 days), the update must include:

1. **Statement of review**: "Pre-admission H&P dated [date] reviewed."
2. **Interval changes**: Document ANY changes in:
   - Symptoms
   - Medications
   - Allergies
   - Physical examination findings
   - Lab results
   - New diagnoses or procedures since the original H&P
3. **Current examination**: At minimum, examine systems relevant to the reason for admission
4. **If no changes**: Document "No interval changes since H&P dated [date]." Still requires a current examination note.

## Documentation Completeness Checklist

| Element | Required | Source |
|---------|----------|--------|
| Patient identification (2 identifiers) | Yes | Patient resource |
| Date and time of exam | Yes | Current datetime |
| Examiner name and credentials | Yes | Practitioner |
| Chief complaint | Yes | Encounter.reasonCode |
| HPI | Yes | Narrative + structured data |
| Past medical history | Yes | Condition resources |
| Past surgical history | Yes | Procedure resources |
| Medications | Yes | MedicationRequest |
| Allergies (or NKDA) | Yes | AllergyIntolerance |
| Family history | Yes | FamilyMemberHistory |
| Social history | Yes | Observation (social-history) |
| Review of systems | Yes | Narrative input required |
| Physical examination | Yes | Narrative input required + vitals from Observation |
| Assessment | Yes | Narrative + Condition codes |
| Plan | Yes | Narrative + CarePlan |
| Attestation/Signature | Yes | Practitioner signature required |
