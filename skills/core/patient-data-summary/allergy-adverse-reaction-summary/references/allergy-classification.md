# WHO Adverse Reaction Classification and Severity Grading

## Purpose

Standardized classification framework for adverse reactions, aligned with WHO-UMC (Uppsala Monitoring Centre) and FHIR AllergyIntolerance resource semantics. Use for consistent categorization and severity assessment.

## WHO Adverse Drug Reaction Classification

### Type A: Augmented (Dose-Dependent)
- **Mechanism:** Exaggerated pharmacological effect
- **Predictable:** Yes
- **Dose-Related:** Yes
- **Incidence:** Common (80% of ADRs)
- **Examples:** Hypoglycemia from insulin, bleeding from warfarin, sedation from benzodiazepines
- **FHIR Mapping:** `type` = "intolerance" (generally)

### Type B: Bizarre (Dose-Independent)
- **Mechanism:** Immunological or idiosyncratic
- **Predictable:** No
- **Dose-Related:** No
- **Incidence:** Uncommon
- **Examples:** Penicillin anaphylaxis, malignant hyperthermia from anesthetics, Stevens-Johnson syndrome
- **FHIR Mapping:** `type` = "allergy"

### Type C: Chronic (Dose and Time-Related)
- **Mechanism:** Cumulative dose effect
- **Predictable:** Sometimes
- **Examples:** Osteoporosis from chronic corticosteroids, nephrotoxicity from aminoglycosides
- **FHIR Mapping:** Generally not documented as AllergyIntolerance; tracked via Condition or AdverseEvent

### Type D: Delayed
- **Mechanism:** Delayed onset (teratogenicity, carcinogenicity)
- **Predictable:** Sometimes
- **Examples:** Vaginal clear cell carcinoma from in-utero DES exposure
- **FHIR Mapping:** Not typically AllergyIntolerance

### Type E: End-of-Treatment (Withdrawal)
- **Mechanism:** Withdrawal effect
- **Examples:** Adrenal crisis from abrupt corticosteroid discontinuation, opioid withdrawal
- **FHIR Mapping:** Not typically AllergyIntolerance

### Type F: Failure
- **Mechanism:** Unexpected treatment failure (often drug interactions)
- **Examples:** OCP failure with concurrent rifampin
- **FHIR Mapping:** Not typically AllergyIntolerance

**For AllergyIntolerance documentation, focus on Type A (intolerances) and Type B (true allergies).**

## Immune-Mediated Reaction Classification (Gell and Coombs)

### Type I: Immediate Hypersensitivity (IgE-Mediated)
- **Onset:** Minutes to 1 hour
- **Mechanism:** IgE antibodies, mast cell degranulation
- **Manifestations:** Urticaria, angioedema, bronchospasm, anaphylaxis
- **Examples:** Penicillin anaphylaxis, bee sting anaphylaxis, peanut allergy
- **FHIR Severity:** Usually moderate to severe
- **FHIR Criticality:** Usually high

### Type II: Cytotoxic (IgG/IgM-Mediated)
- **Onset:** Hours to days
- **Mechanism:** Antibodies directed against cell-surface antigens
- **Manifestations:** Hemolytic anemia, thrombocytopenia, neutropenia
- **Examples:** Drug-induced hemolytic anemia (methyldopa, penicillin), heparin-induced thrombocytopenia
- **FHIR Severity:** Moderate to severe
- **FHIR Criticality:** Varies

### Type III: Immune Complex
- **Onset:** Days to weeks
- **Mechanism:** Antigen-antibody complex deposition
- **Manifestations:** Serum sickness, vasculitis, drug fever
- **Examples:** Serum sickness-like reaction to cefaclor, lupus-like syndrome from hydralazine
- **FHIR Severity:** Mild to moderate
- **FHIR Criticality:** Usually low

### Type IV: Delayed Hypersensitivity (T-Cell-Mediated)
- **Onset:** 24-72 hours (sometimes longer)
- **Mechanism:** T-cell activation
- **Manifestations:** Contact dermatitis, maculopapular rash, SJS/TEN, DRESS
- **Examples:** Contact dermatitis to nickel, maculopapular rash from amoxicillin, SJS from carbamazepine
- **FHIR Severity:** Mild (contact dermatitis) to severe (SJS/TEN)
- **FHIR Criticality:** Low (simple rash) to high (SJS/TEN history)

## Severe Cutaneous Adverse Reactions (SCARs)

These require special documentation and are high-criticality by definition:

### Stevens-Johnson Syndrome (SJS)
- **Definition:** < 10% body surface area detachment
- **SNOMED:** 73442001
- **Criticality:** Always HIGH
- **Common culprits:** Allopurinol, carbamazepine, lamotrigine, phenytoin, sulfonamides, nevirapine
- **Re-exposure:** ABSOLUTE CONTRAINDICATION to causative drug

### Toxic Epidermal Necrolysis (TEN)
- **Definition:** > 30% body surface area detachment
- **SNOMED:** 768962006
- **Criticality:** Always HIGH
- **Mortality:** 25-35%
- **Re-exposure:** ABSOLUTE CONTRAINDICATION

### DRESS Syndrome (Drug Reaction with Eosinophilia and Systemic Symptoms)
- **SNOMED:** 293104008 (Drug-induced hypersensitivity syndrome)
- **Onset:** 2-8 weeks after drug initiation
- **Criticality:** Always HIGH
- **Common culprits:** Allopurinol, anticonvulsants, dapsone, minocycline, sulfonamides
- **Re-exposure:** ABSOLUTE CONTRAINDICATION

### Acute Generalized Exanthematous Pustulosis (AGEP)
- **Onset:** < 48 hours
- **Criticality:** HIGH
- **Common culprits:** Aminopenicillins, quinolones, hydroxychloroquine, diltiazem

## FHIR AllergyIntolerance Severity Grading

### Mapping Clinical Assessment to FHIR

| Clinical Presentation | FHIR reaction.severity | FHIR criticality |
|----------------------|----------------------|-------------------|
| Localized rash, mild GI symptoms | mild | low |
| Urticaria (generalized), significant GI symptoms, non-life-threatening angioedema | moderate | low or high |
| Anaphylaxis, bronchospasm, hypotension, airway compromise | severe | high |
| SJS/TEN/DRESS | severe | high |
| Nausea/itching from opioids (expected side effect) | mild | low |
| Headache from nitrates (expected side effect) | mild | low |

### Criticality Assessment Guide

**High Criticality -- Always assign for:**
- History of anaphylaxis to the substance
- History of SJS/TEN/DRESS
- Allergy to NMBAs (perioperative risk)
- Severe bronchospasm history
- Drug-specific: methotrexate hypersensitivity, chemotherapy anaphylaxis

**Low Criticality -- Assign for:**
- Mild rash without systemic symptoms
- GI intolerance (nausea, diarrhea)
- Mild local reactions (injection site)
- Environmental allergies (rhinitis, mild conjunctivitis)
- Food intolerances (lactose, gluten sensitivity without celiac)

**Unable to Assess -- Assign when:**
- Historical report with no details on reaction type
- Patient unable to describe reaction
- Allergy reported by family member without details
- Documentation states "allergic" with no further information

## SNOMED CT Codes for Common Manifestations

| Manifestation | SNOMED Code | Display |
|---------------|-------------|---------|
| Anaphylaxis | 39579001 | Anaphylaxis |
| Urticaria (hives) | 126485001 | Urticaria |
| Angioedema | 41291007 | Angioedema |
| Rash (maculopapular) | 271807003 | Eruption of skin |
| Bronchospasm | 4386001 | Bronchospasm |
| Dyspnea | 267036007 | Dyspnea |
| Nausea | 422587007 | Nausea |
| Vomiting | 422400008 | Vomiting |
| Diarrhea | 62315008 | Diarrhea |
| Pruritus (itching) | 418290006 | Itching |
| Contact dermatitis | 40275004 | Contact dermatitis |
| Hypotension | 45007003 | Low blood pressure |
| Tachycardia | 3424008 | Tachycardia |
| Serum sickness | 234660009 | Serum sickness |
| Erythema multiforme | 36715001 | Erythema multiforme |
| Drug fever | 405543007 | Drug fever |
| Hemolytic anemia | 61261009 | Hemolytic anemia |
| Thrombocytopenia | 302215000 | Thrombocytopenia |
| Agranulocytosis | 417672002 | Agranulocytosis |
| Hepatotoxicity | 235856003 | Hepatotoxicity |
| Nephrotoxicity | 55074005 | Nephrotoxicity |
| Ototoxicity | 88151004 | Ototoxicity |
| QT prolongation | 111975006 | Prolonged QT interval |

## Allergy vs Intolerance Decision Tree

```
Did the reaction involve the immune system?
  |
  +-- YES (or likely) --> type = "allergy"
  |   |
  |   +-- Was it IgE-mediated (immediate: urticaria, angioedema, anaphylaxis)?
  |   |   +-- YES --> Type I hypersensitivity, criticality likely HIGH
  |   |
  |   +-- Was it T-cell mediated (delayed: rash, SJS/TEN)?
  |       +-- YES --> Type IV hypersensitivity
  |
  +-- NO (pharmacological or unknown) --> type = "intolerance"
  |   |
  |   +-- Was it a known/expected side effect?
  |   |   +-- YES --> Consider NOT documenting as AllergyIntolerance
  |   |   |          (e.g., opioid nausea, statin myalgia)
  |   |
  |   +-- Was it dose-related?
  |       +-- YES --> Type A adverse reaction, low criticality
  |
  +-- UNKNOWN --> type = "allergy" (err on side of caution)
                  verificationStatus = "unconfirmed"
                  criticality = "unable-to-assess"
```

## Documentation Quality Metrics

An allergy entry is considered COMPLETE when it has:
1. Coded substance (RxNorm or SNOMED) -- not just free text
2. `type` specified (allergy vs intolerance)
3. `category` specified (food, medication, environment, biologic)
4. `clinicalStatus` = active or resolved
5. `verificationStatus` = confirmed (or appropriately set)
6. `criticality` assigned
7. At least one `reaction` with:
   - Coded `manifestation` (SNOMED)
   - `severity` (mild, moderate, severe)
8. `recorder` and `recordedDate` populated

An allergy entry is INCOMPLETE (flag for review) when:
- Free text only, no coded substance
- No reaction details
- verificationStatus = "unconfirmed" for > 12 months
- Criticality = "unable-to-assess" for > 12 months
- Side effects documented as allergies (opioid nausea, statin myalgia)
