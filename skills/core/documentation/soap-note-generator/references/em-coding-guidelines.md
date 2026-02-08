# E&M Coding Guidelines (2021 Revisions)

## Overview

The 2021 E&M coding guidelines eliminated the 1995/1997 documentation-based system for office/outpatient visits (99202-99215). Visit level selection is now based on either medical decision-making (MDM) complexity OR total time on the date of encounter.

## Medical Decision-Making (MDM) Framework

MDM has three components. The visit level is determined by the **two of three** rule: meet or exceed the criteria in at least 2 of the 3 elements.

### Element 1: Number and Complexity of Problems Addressed

| Level | Criteria | Examples |
|-------|----------|----------|
| Minimal | 1 self-limited or minor problem | Insect bite, cold sore, mild URI |
| Low | 2+ self-limited problems OR 1 acute uncomplicated illness/injury | Cystitis, allergic rhinitis, sprain |
| Moderate | 1+ chronic illness with mild exacerbation OR 1 undiagnosed new problem with uncertain prognosis OR 1 acute illness with systemic symptoms | COPD exacerbation, new breast mass, pneumonia |
| High | 1+ chronic illness with severe exacerbation OR 1 acute/chronic illness posing threat to life or bodily function | DKA, MI, acute stroke, sepsis, pulmonary embolism |

**Key definitions:**
- **Self-limited**: Expected to resolve on its own, minimal risk
- **Acute uncomplicated**: Short duration, low risk of morbidity, full recovery expected
- **Chronic illness with mild exacerbation**: Short-term increase in symptoms
- **Chronic illness with severe exacerbation**: Requires urgent or emergent intervention, significant risk

### Element 2: Amount and/or Complexity of Data Reviewed and Analyzed

| Level | Criteria |
|-------|----------|
| Minimal / None | Not required |
| Limited | Review/order of: 2+ categories of tests (labs, imaging, etc.) OR Review of prior external notes/records OR Ordering of tests |
| Moderate | Must meet 1 of 3 categories: (A) 3+ categories of tests OR (B) Independent interpretation of test (not reported separately) OR (C) Discussion of management with external physician/QHP with documented decision |
| Extensive | Must meet 2 of 3 categories: (A) 3+ categories of tests OR (B) Independent interpretation of test OR (C) Discussion of management with external physician/QHP |

**Categories of tests:**
1. Lab tests (e.g., CBC, BMP, UA)
2. Imaging (e.g., X-ray, CT, MRI, US)
3. Medicine-section tests (e.g., PFTs, EKG interpretation, sleep study)
4. Tests in the surgical pathology section

**Independent interpretation**: Provider personally reviews the raw data (e.g., reading the actual ECG tracing, reviewing imaging). Must document the interpretation.

**External physician discussion**: Must document who was consulted, what was discussed, and the clinical decision that resulted. Internal team discussions do not count.

### Element 3: Risk of Complications, Morbidity, and/or Mortality

| Level | Risk Examples |
|-------|-------------|
| Minimal | Band-aid, OTC medication recommendation |
| Low | OTC medications, minor surgery with no risk factors, PT/OT order |
| Moderate | Prescription drug management, decision about minor surgery with risk factors, decision about elective major surgery without risk factors, diagnosis or treatment significantly limited by social determinants |
| High | Drug therapy requiring intensive monitoring (e.g., immunosuppressives, anticoagulation), decision about emergency major surgery, decision about hospitalization, decision to not resuscitate or de-escalate care |

**Key risk considerations:**
- **Prescription drug management**: Any new or ongoing prescription counts as moderate risk
- **Parenteral controlled substances**: High risk
- **Decision about hospitalization**: Always high risk (includes decisions to observe, admission, and decisions to discharge with significant risk)

## Visit Level Selection Table

| CPT (New) | CPT (Est.) | MDM Level | Time (New) | Time (Est.) |
|-----------|-----------|-----------|------------|-------------|
| 99202 | 99212 | Straightforward | 15-29 min | 10-19 min |
| 99203 | 99213 | Low | 30-44 min | 20-29 min |
| 99204 | 99214 | Moderate | 45-59 min | 30-39 min |
| 99205 | 99215 | High | 60-74 min | 40-54 min |

## MDM Determination Examples

### Level 2 (Straightforward) -- 99212/99202
- **Problem**: 1 self-limited (e.g., follow-up wart, suture removal)
- **Data**: Minimal or none
- **Risk**: Minimal (OTC recommendation, minor procedure without risk)

### Level 3 (Low) -- 99213/99203
- **Problem**: 2 self-limited problems OR 1 acute uncomplicated (e.g., UTI, ankle sprain)
- **Data**: Limited (ordered UA + culture, reviewed X-ray report)
- **Risk**: Low (prescribed antibiotics or NSAIDs)

### Level 4 (Moderate) -- 99214/99204
- **Problem**: 1 chronic with mild exacerbation (e.g., COPD exacerbation) OR 1 new undiagnosed problem (e.g., new breast lump)
- **Data**: Moderate (ordered labs + imaging + discussed with radiologist)
- **Risk**: Moderate (new prescription, decision about surgery)

### Level 5 (High) -- 99215/99205
- **Problem**: Threat to life (e.g., chest pain r/o MI, new seizure, sepsis)
- **Data**: Extensive (ordered multiple test categories + independently interpreted ECG + discussed with cardiologist)
- **Risk**: High (decision to hospitalize, parenteral controlled substance, intensive drug monitoring)

## Inpatient / Observation E&M (99221-99223, 99231-99233)

These codes still use the traditional framework with all three key components for initial visits:
- **History**: Problem-focused to comprehensive
- **Exam**: Problem-focused to comprehensive
- **MDM**: Straightforward to high complexity

### Initial Inpatient (99221-99223)

| Code | History | Exam | MDM |
|------|---------|------|-----|
| 99221 | Detailed or comprehensive | Detailed or comprehensive | Low or moderate |
| 99222 | Comprehensive | Comprehensive | Moderate |
| 99223 | Comprehensive | Comprehensive | High |

### Subsequent Inpatient (99231-99233)

| Code | History | Exam | MDM |
|------|---------|------|-----|
| 99231 | Problem-focused interval | Problem-focused | Low |
| 99232 | Expanded problem-focused interval | Expanded problem-focused | Moderate |
| 99233 | Detailed interval | Detailed | High |

## Emergency Department E&M (99281-99285)

ED visits use all three key components (history, exam, MDM).

| Code | History | Exam | MDM | Example |
|------|---------|------|-----|---------|
| 99281 | Problem-focused | Problem-focused | Straightforward | Suture removal, medication refill |
| 99282 | Expanded problem-focused | Expanded problem-focused | Low | Simple laceration, UTI |
| 99283 | Expanded problem-focused | Expanded problem-focused | Moderate | Extremity fracture, asthma exacerbation |
| 99284 | Detailed | Detailed | Moderate | Abdominal pain workup, chest pain low-risk |
| 99285 | Comprehensive | Comprehensive | High | MI, stroke, sepsis, trauma activation |

## Critical Care (99291-99292)

- **99291**: First 30-74 minutes of critical care
- **99292**: Each additional 30 minutes
- Must document critical illness/injury and direct personal management
- Time-based: document total critical care time
- Critical care time includes: direct patient care, reviewing labs/imaging, discussing care with family, coordinating critical care delivery
- Does NOT include: time for separately billable procedures, teaching time

## Add-On Codes

| Code | Description | Requirement |
|------|------------|-------------|
| +99417 | Prolonged office visit (each 15 min) | Must exceed max time for 99205/99215 |
| +99418 | Prolonged inpatient (each 15 min) | Beyond time for 99223/99233 |
| 99354 | Prolonged service, first hour | Direct face-to-face time |

## Documentation Tips for AI-Assisted Note Generation

1. **Link assessment to data**: Each problem in the assessment should reference supporting objective data (lab values, imaging findings, exam findings).

2. **Document complexity explicitly**: When MDM is moderate or high, explicitly state why:
   - "New undiagnosed problem with uncertain prognosis"
   - "Chronic illness with severe exacerbation requiring..."
   - "Decision to hospitalize based on..."

3. **Data review documentation**: When claiming data reviewed, be specific:
   - "Independently reviewed ECG showing ST elevations in leads II, III, aVF"
   - "Discussed with Dr. [name], radiologist, regarding CT findings; agreed to..."

4. **Risk documentation**: Document the specific risk factor:
   - "Started warfarin -- requires INR monitoring (intensive drug monitoring)"
   - "Decision made to admit for observation given risk of..."

5. **Time documentation**: If billing on time, state:
   - "Total time on date of encounter: XX minutes"
   - Include: chart review, face-to-face, care coordination, documentation
