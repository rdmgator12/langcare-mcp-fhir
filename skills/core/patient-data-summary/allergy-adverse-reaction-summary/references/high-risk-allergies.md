# High-Risk Allergy Categories and Cross-Reactivity Patterns

## Purpose

Reference for identifying allergies that require elevated clinical vigilance, specialized protocols, or environmental precautions. Use when reviewing AllergyIntolerance resources to flag entries that demand immediate attention.

## Anaphylaxis Risk Indicators

### FHIR Detection

Anaphylaxis is identified in AllergyIntolerance resources by:
- `reaction[].manifestation[].coding[].code` = SNOMED 39579001 (Anaphylaxis)
- `reaction[].manifestation[].coding[].display` containing "anaphylaxis" or "anaphylactic"
- `reaction[].severity` = "severe" combined with `criticality` = "high"
- `reaction[].manifestation` including multiple system involvement: urticaria + dyspnea + hypotension

### Clinical Protocol When Anaphylaxis History Present

1. Epinephrine auto-injector must be available
2. Document allergen prominently in chart
3. Patient should wear medical alert identification
4. Pre-procedure protocols required (anesthesia, contrast, blood products)
5. Cross-reactivity evaluation mandatory before prescribing related substances

## High-Risk Drug Allergy Categories

### 1. Beta-Lactam Antibiotics (Penicillin Family)

**Substances:** Penicillin, amoxicillin, ampicillin, piperacillin, nafcillin, dicloxacillin

**Cross-Reactivity Matrix:**

| Patient Allergic To | Risk With Cephalosporins | Risk With Carbapenems | Risk With Monobactams |
|---------------------|-------------------------|----------------------|----------------------|
| Penicillin (anaphylaxis) | ~2% (avoid without testing) | ~1% (use with caution) | Negligible (aztreonam safe) |
| Penicillin (rash only) | < 1% (generally safe, esp. 3rd/4th gen) | < 1% (generally safe) | Negligible |
| Amoxicillin specifically | Higher with cephalosporins sharing R1 side chain (cefadroxil, cefprozil) | Low | Negligible |
| Ampicillin specifically | Higher with cephalosporins sharing R1 side chain (cefaclor, cephalexin) | Low | Negligible |

**Side-Chain Cross-Reactivity (Current Evidence):**
- Cross-reactivity is driven by R1 side chain similarity, not the beta-lactam ring
- Aminopenicillins (amoxicillin, ampicillin) cross-react with aminocephalosporins (cephalexin, cefaclor, cefadroxil, cefprozil)
- 3rd/4th generation cephalosporins (ceftriaxone, cefepime) have very low cross-reactivity with penicillins

**Clinical Decision Guide:**
- Penicillin allergy with ANAPHYLAXIS: Avoid all beta-lactams. Refer for allergy testing. Use aztreonam, fluoroquinolones, or other alternatives.
- Penicillin allergy with RASH ONLY: 2nd+ generation cephalosporins generally safe. Monitor for 1 hour after first dose.
- Penicillin allergy UNVERIFIED: >90% of patients labeled "penicillin allergic" are NOT truly allergic. Recommend penicillin skin testing when possible.

### 2. Sulfonamide Antibiotics

**Substances:** Sulfamethoxazole (in TMP-SMX/Bactrim), sulfasalazine, silver sulfadiazine

**Cross-Reactivity Clarification:**
- Sulfonamide ANTIBIOTIC allergy does NOT contraindicate non-antibiotic sulfonamides
- Non-antibiotic sulfonamides (generally safe in sulfa-allergic patients):
  - Furosemide (Lasix)
  - Hydrochlorothiazide (HCTZ)
  - Celecoxib (Celebrex)
  - Sumatriptan (Imitrex)
  - Sulfasalazine (different arylamine structure)
- The historical teaching of broad "sulfa allergy" cross-reactivity is outdated
- True cross-reactivity between antibiotic and non-antibiotic sulfonamides is not supported by current evidence

### 3. NSAIDs

**Substances:** Aspirin, ibuprofen, naproxen, ketorolac, diclofenac, meloxicam, indomethacin, piroxicam

**Types of NSAID Reactions:**

| Type | Mechanism | Scope | Alternative |
|------|-----------|-------|-------------|
| NSAID-exacerbated respiratory disease (AERD/Samter's) | COX-1 inhibition | All NSAIDs including aspirin | Acetaminophen (< 1000mg), COX-2 selective (celecoxib) with caution |
| NSAID-induced urticaria/angioedema | COX-1 inhibition | All NSAIDs | Acetaminophen, COX-2 selective |
| Single NSAID-induced anaphylaxis | IgE-mediated | Specific drug only | Other NSAIDs may be safe |

**Samter's Triad Detection:**
- NSAID allergy/intolerance + Asthma (check Condition for SNOMED 195967001) + Nasal polyps (check Condition for SNOMED 52684005)
- If all three present, flag as AERD

### 4. Contrast Dye (Iodinated)

**Substances:** Iohexol, iopamidol, iodixanol, iopromide, diatrizoate

**Risk Classification:**

| Previous Reaction | Risk of Repeat Reaction | Protocol |
|-------------------|------------------------|----------|
| Mild (nausea, urticaria, flushing) | 10-35% without premedication | Premedication protocol |
| Moderate (bronchospasm, significant hypotension) | Higher | Premedication + alternative agent |
| Severe (anaphylaxis, cardiac arrest) | Highest | Avoid if possible; if essential, full premedication + alternative agent + resuscitation team standby |

**Standard Premedication Protocol:**
- Prednisone 50mg PO at 13h, 7h, and 1h before procedure
- Diphenhydramine 50mg PO/IV/IM 1h before procedure
- Use non-ionic low-osmolar or iso-osmolar contrast agent

**Note:** Shellfish allergy and iodine allergy are NOT risk factors for contrast reactions. This is a persistent myth. Contrast reactions are related to the contrast molecule, not iodine content.

### 5. Latex

**Substances:** Natural rubber latex (SNOMED 111088007)

**High-Risk Populations:**
- Healthcare workers with repeated exposure
- Patients with spina bifida (up to 68% prevalence)
- Patients with multiple prior surgeries
- Patients with atopic dermatitis

**Cross-Reactive Foods (Latex-Fruit Syndrome):**

| High Cross-Reactivity (> 30%) | Moderate Cross-Reactivity (10-30%) | Low Cross-Reactivity (< 10%) |
|-------------------------------|-------------------------------------|------------------------------|
| Banana | Apple | Cherry |
| Avocado | Carrot | Coconut |
| Chestnut | Celery | Fig |
| Kiwi | Papaya | Grape |
| | Potato | Pear |
| | Tomato | Plum |
| | Melon | Strawberry |

**Environmental Precautions:**
- Non-latex gloves (nitrile, vinyl)
- Latex-free IV tubing, catheters, tourniquets
- Latex-free surgical equipment
- First case of the day in OR (minimize airborne latex particles)

### 6. Opioid Reactions

**True Allergy vs Side Effects:**

| Reaction | Classification | Clinical Significance |
|----------|---------------|----------------------|
| Nausea, vomiting | Side effect (NOT allergy) | Do NOT document as allergy |
| Constipation | Side effect | Do NOT document as allergy |
| Pruritus (no rash) | Histamine release (NOT IgE-mediated) | Usually not true allergy |
| Urticaria, angioedema | Possible allergy | Document and cross-reference |
| Anaphylaxis | True allergy | Document with high criticality |

**Opioid Cross-Reactivity by Structure:**

| Class | Examples | Cross-Reactivity |
|-------|----------|-----------------|
| Phenanthrenes | Morphine, codeine, hydromorphone, oxycodone, hydrocodone | Cross-react with each other |
| Phenylpiperidines | Fentanyl, meperidine, alfentanil | Low cross-reactivity with phenanthrenes |
| Diphenylheptanes | Methadone | Low cross-reactivity with others |

If allergic to morphine/codeine: fentanyl and methadone are structurally different and usually safe.

### 7. Fluoroquinolone Reactions

**Substances:** Ciprofloxacin, levofloxacin, moxifloxacin, ofloxacin

**Cross-Reactivity:** High within the class (~10% cross-reactivity between fluoroquinolones). If allergy to one fluoroquinolone, avoid all unless skin testing performed.

**Special Concerns:**
- QT prolongation risk (especially with moxifloxacin)
- Tendon rupture risk (especially in patients > 60, on corticosteroids, or with renal impairment)
- Aortic aneurysm/dissection risk (FDA black box warning)

### 8. Local Anesthetic Reactions

**Substances:** Lidocaine, bupivacaine, mepivacaine, prilocaine, procaine

**Two Chemical Groups:**

| Group | Examples | Cross-Reactivity |
|-------|----------|-----------------|
| Esters | Procaine, chloroprocaine, benzocaine, tetracaine | Cross-react within ester group; metabolized to PABA |
| Amides | Lidocaine, bupivacaine, mepivacaine, ropivacaine, prilocaine | Cross-react within amide group (rare) |

- Ester and amide groups do NOT cross-react
- If allergic to an ester local anesthetic, amide alternatives are safe
- Most reported "local anesthetic allergies" are vasovagal reactions or epinephrine side effects, not true allergies

### 9. Neuromuscular Blocking Agents (NMBAs)

**Substances:** Succinylcholine, rocuronium, vecuronium, cisatracurium, atracurium

**Clinical Significance:**
- NMBAs are the most common cause of perioperative anaphylaxis
- Cross-reactivity between NMBAs is approximately 60-70%
- If allergy to one NMBA: allergy testing for all agents recommended before future anesthesia
- Document prominently for any planned surgeries

## Verification Status Decision Guide

| Scenario | Appropriate verificationStatus |
|----------|-------------------------------|
| Patient reports allergy, reaction details known | confirmed |
| Patient reports allergy, no reaction details | unconfirmed |
| Allergy testing performed, positive | confirmed |
| Allergy testing performed, negative | refuted |
| Patient tolerated the substance subsequently | refuted or resolved |
| Family member reported, patient unsure | unconfirmed |
| Historical record, unable to verify | unconfirmed |
| Entered by mistake | entered-in-error |

## Severity vs Criticality

These are distinct concepts in FHIR:

- **Severity** (`reaction.severity`): How bad the reaction WAS (mild, moderate, severe) -- describes a past event
- **Criticality** (`AllergyIntolerance.criticality`): How bad the reaction COULD BE on re-exposure (low, high, unable-to-assess) -- predicts future risk

A patient can have a mild past reaction (severity=mild) but high criticality if the substance class is known to cause escalating reactions on re-exposure.
