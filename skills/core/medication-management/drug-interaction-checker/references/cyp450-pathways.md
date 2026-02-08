# CYP450 Enzyme Pathways Reference

Cytochrome P450 (CYP450) enzymes are responsible for metabolizing approximately 75% of all drugs. Understanding substrate, inhibitor, and inducer relationships is critical for predicting pharmacokinetic drug-drug interactions.

## Key Concepts

**Substrate**: Drug metabolized by the enzyme. Inhibition of the enzyme increases substrate levels. Induction decreases substrate levels.

**Inhibitor**: Drug that blocks the enzyme, causing substrates to accumulate. Effects are typically rapid (within 1-3 days for strong inhibitors).

**Inducer**: Drug that increases enzyme expression, causing substrates to be metabolized faster. Effects are gradual (7-14 days to reach full effect) and persist after discontinuation (7-14 days to resolve).

**Narrow Therapeutic Index (NTI)**: Drugs where small changes in levels cause toxicity or treatment failure. CYP450 interactions with NTI drugs are clinically most significant.

---

## CYP3A4

**Responsible for**: ~50% of drug metabolism. Most clinically important enzyme.

### Substrates (drugs metabolized by CYP3A4)

#### Narrow Therapeutic Index
| Drug | Category | Interaction Concern |
|------|----------|-------------------|
| Cyclosporine | Immunosuppressant | Levels increase with inhibitors -> nephrotoxicity |
| Tacrolimus | Immunosuppressant | Levels increase with inhibitors -> nephrotoxicity |
| Sirolimus | Immunosuppressant | Levels increase with inhibitors -> myelosuppression |
| Fentanyl | Opioid | Levels increase with inhibitors -> respiratory depression |
| Alfentanil | Opioid | Levels increase with inhibitors -> respiratory depression |
| Ergotamine | Antimigraine | Levels increase with inhibitors -> ergotism, vasospasm |
| Pimozide | Antipsychotic | Levels increase with inhibitors -> QT prolongation |
| Quinidine | Antiarrhythmic | Levels increase with inhibitors -> QT prolongation |

#### Other Important Substrates
| Drug | Category |
|------|----------|
| Simvastatin | Statin (rhabdomyolysis risk with inhibitors) |
| Lovastatin | Statin (rhabdomyolysis risk with inhibitors) |
| Atorvastatin | Statin (moderate risk) |
| Apixaban | Anticoagulant |
| Rivaroxaban | Anticoagulant |
| Midazolam | Benzodiazepine |
| Triazolam | Benzodiazepine |
| Alprazolam | Benzodiazepine |
| Nifedipine | Calcium channel blocker |
| Felodipine | Calcium channel blocker |
| Amlodipine | Calcium channel blocker (less affected) |
| Sildenafil | PDE5 inhibitor |
| Buspirone | Anxiolytic |
| Colchicine | Anti-gout |
| Oxycodone | Opioid |
| Methadone | Opioid (also CYP2B6) |

### Strong Inhibitors (avoid with NTI substrates)
| Drug | Category | Onset |
|------|----------|-------|
| Ketoconazole | Antifungal | 1-2 days |
| Itraconazole | Antifungal | 1-2 days |
| Voriconazole | Antifungal | 1-2 days |
| Posaconazole | Antifungal | 1-2 days |
| Clarithromycin | Macrolide antibiotic | 1-2 days |
| Ritonavir | HIV protease inhibitor | 1-2 days |
| Cobicistat | Pharmacokinetic booster | 1-2 days |
| Nefazodone | Antidepressant | 3-5 days |

### Moderate Inhibitors (use with caution)
| Drug | Category |
|------|----------|
| Erythromycin | Macrolide antibiotic |
| Fluconazole | Antifungal |
| Diltiazem | Calcium channel blocker |
| Verapamil | Calcium channel blocker |
| Amiodarone | Antiarrhythmic |
| Grapefruit juice | Food (variable) |
| Aprepitant | Antiemetic |
| Crizotinib | Antineoplastic |

### Strong Inducers (may cause treatment failure of substrates)
| Drug | Category | Time to Full Effect |
|------|----------|-------------------|
| Rifampin | Antimycobacterial | 7-10 days |
| Carbamazepine | Antiepileptic | 7-14 days |
| Phenytoin | Antiepileptic | 7-14 days |
| Phenobarbital | Antiepileptic | 14-21 days |
| St. John's Wort | Herbal supplement | 7-14 days |
| Enzalutamide | Antineoplastic | 7-14 days |
| Mitotane | Antineoplastic | 14+ days |

### Moderate Inducers
| Drug | Category |
|------|----------|
| Efavirenz | Antiretroviral |
| Bosentan | Endothelin receptor antagonist |
| Modafinil | Wakefulness agent |
| Nafcillin | Antibiotic |

---

## CYP2D6

**Responsible for**: ~25% of drug metabolism. Highly polymorphic (genetic variation affects activity).

**Genetic Variation**:
- Poor metabolizers (PM): 5-10% of Caucasians. Substrates accumulate as if inhibited.
- Ultrarapid metabolizers (UM): 1-10% (varies by ethnicity). Substrates cleared faster; prodrugs activated faster.
- Cannot be induced (unlike other CYPs).

### Substrates
| Drug | Category | Clinical Concern |
|------|----------|-----------------|
| Codeine | Opioid prodrug | PMs: no analgesia. UMs: toxicity (avoid in children). |
| Tramadol | Opioid | PMs: reduced efficacy. UMs: increased active metabolite. |
| Hydrocodone | Opioid | Partially activated by CYP2D6. |
| Oxycodone | Opioid | Minor CYP2D6 pathway (CYP3A4 primary). |
| Tamoxifen | Antineoplastic | PMs: reduced active metabolite (endoxifen), reduced efficacy. |
| Metoprolol | Beta-blocker | PMs: increased levels, bradycardia, hypotension. |
| Carvedilol | Beta-blocker | PMs: increased levels. |
| Propranolol | Beta-blocker | Increased levels with inhibitors. |
| Flecainide | Antiarrhythmic (NTI) | PMs: increased levels, proarrhythmia. |
| Thioridazine | Antipsychotic | PMs: QT prolongation. Contraindicated with inhibitors. |
| Risperidone | Antipsychotic | Increased levels with inhibitors. |
| Aripiprazole | Antipsychotic | Reduce dose 50% with strong inhibitors. |
| Atomoxetine | ADHD | PMs: 10x higher levels. Reduce dose. |
| Dextromethorphan | Antitussive | PMs: increased sedation. Used as phenotyping probe. |
| Venlafaxine | SNRI | CYP2D6 converts to active metabolite (desvenlafaxine). |
| Nortriptyline | TCA (NTI) | PMs: toxicity. Monitor levels. |
| Amitriptyline | TCA | Partially metabolized by CYP2D6. |

### Strong Inhibitors
| Drug | Category |
|------|----------|
| Fluoxetine | SSRI (long half-life: inhibition persists weeks after stopping) |
| Paroxetine | SSRI |
| Bupropion | Antidepressant/smoking cessation |
| Quinidine | Antiarrhythmic |
| Cinacalcet | Calcimimetic |
| Terbinafine | Antifungal |

### Moderate Inhibitors
| Drug | Category |
|------|----------|
| Duloxetine | SNRI |
| Sertraline (at high doses) | SSRI |
| Amiodarone | Antiarrhythmic |
| Diphenhydramine | Antihistamine |
| Celecoxib | NSAID |

### Inducers
CYP2D6 is generally NOT inducible. This is a key distinguishing feature.

---

## CYP2C9

**Responsible for**: Metabolism of warfarin (S-enantiomer), phenytoin, many NSAIDs.

**Genetic Variation**: CYP2C9*2 and *3 alleles reduce activity. Affects warfarin dosing.

### Substrates
| Drug | Category | NTI? |
|------|----------|------|
| Warfarin (S-enantiomer) | Anticoagulant | Yes |
| Phenytoin | Antiepileptic | Yes |
| Celecoxib | NSAID | No |
| Losartan | ARB (prodrug) | No |
| Irbesartan | ARB | No |
| Glipizide | Sulfonylurea | No |
| Glimepiride | Sulfonylurea | No |
| Tolbutamide | Sulfonylurea | No |
| Ibuprofen | NSAID | No |
| Naproxen | NSAID | No |
| Diclofenac | NSAID | No |

### Strong Inhibitors
| Drug | Category |
|------|----------|
| Fluconazole | Antifungal |
| Amiodarone | Antiarrhythmic |
| Miconazole (systemic) | Antifungal |

### Moderate Inhibitors
| Drug | Category |
|------|----------|
| Fluoxetine | SSRI |
| Fluvoxamine | SSRI |
| Voriconazole | Antifungal |
| Trimethoprim-sulfamethoxazole | Antibiotic |
| Metronidazole | Antibiotic |

### Inducers
| Drug | Category |
|------|----------|
| Rifampin | Antimycobacterial |
| Carbamazepine | Antiepileptic |
| Phenobarbital | Antiepileptic |
| St. John's Wort | Herbal |

---

## CYP2C19

**Responsible for**: Clopidogrel activation, PPI metabolism, some antidepressants.

**Genetic Variation**:
- PMs: 2-5% Caucasians, 15-25% East Asians. Clopidogrel may not work.
- UMs: 5-30% (varies by ethnicity). PPIs metabolized faster, may need higher doses.

### Substrates
| Drug | Category | Clinical Concern |
|------|----------|-----------------|
| Clopidogrel | Antiplatelet (prodrug) | PMs: no activation, increased CV events. Consider prasugrel/ticagrelor. |
| Omeprazole | PPI | PMs: higher levels (usually not harmful). |
| Esomeprazole | PPI | Same as omeprazole. |
| Lansoprazole | PPI | Less affected than omeprazole. |
| Pantoprazole | PPI | Least CYP2C19 dependent. |
| Voriconazole | Antifungal | PMs: increased levels. UMs: subtherapeutic. |
| Citalopram | SSRI | PMs: increased levels, QT prolongation risk. Max 20mg in PMs. |
| Escitalopram | SSRI | PMs: increased levels. Max 10mg in PMs. |
| Diazepam | Benzodiazepine | PMs: prolonged sedation. |
| Phenytoin (partially) | Antiepileptic | PMs: increased levels. |

### Strong Inhibitors
| Drug | Category | Key Interaction |
|------|----------|----------------|
| Fluoxetine | SSRI | Reduces clopidogrel activation |
| Fluvoxamine | SSRI | Reduces clopidogrel activation |
| Omeprazole | PPI | Reduces clopidogrel activation (FDA warning) |
| Esomeprazole | PPI | Same as omeprazole |
| Fluconazole | Antifungal | Reduces clopidogrel activation |
| Ticlopidine | Antiplatelet | Paradoxically inhibits 2C19 |

### Moderate Inhibitors
| Drug | Category |
|------|----------|
| Cimetidine | H2-blocker |
| Ketoconazole | Antifungal |
| Voriconazole | Antifungal |

### Inducers
| Drug | Category |
|------|----------|
| Rifampin | Antimycobacterial |
| Carbamazepine | Antiepileptic |
| St. John's Wort | Herbal |

**Clinical Pearl**: Pantoprazole is preferred over omeprazole in patients on clopidogrel due to less CYP2C19 inhibition.

---

## CYP1A2

**Responsible for**: Theophylline, caffeine, some antipsychotics.

**Unique Feature**: Induced by smoking (polycyclic aromatic hydrocarbons, not nicotine). Dose adjustments needed when patients start or stop smoking.

### Substrates
| Drug | Category | NTI? |
|------|----------|------|
| Theophylline | Bronchodilator | Yes |
| Caffeine | Stimulant | No |
| Clozapine | Antipsychotic | Yes |
| Olanzapine | Antipsychotic | No |
| Tizanidine | Muscle relaxant | No (but very sensitive) |
| Ramelteon | Hypnotic | No |
| Alosetron | IBS medication | No |
| Duloxetine | SNRI | No |
| Melatonin | Supplement | No |

### Strong Inhibitors
| Drug | Category | Key Interaction |
|------|----------|----------------|
| Fluvoxamine | SSRI | Contraindicated with tizanidine. Increases clozapine/theophylline significantly. |
| Ciprofloxacin | Fluoroquinolone | Increases theophylline levels 30-80%. Use levofloxacin or moxifloxacin instead. |

### Moderate Inhibitors
| Drug | Category |
|------|----------|
| Mexiletine | Antiarrhythmic |
| Oral contraceptives | Hormonal |
| Cimetidine | H2-blocker |
| Vemurafenib | Antineoplastic |
| Acyclovir (high dose) | Antiviral |

### Inducers
| Inducer | Effect | Clinical Impact |
|---------|--------|----------------|
| Smoking (tobacco) | Strong induction | Smokers need higher doses of clozapine, olanzapine, theophylline. |
| Smoking cessation | Removal of induction over 3-7 days | Reduce clozapine dose 30-50% when patient stops smoking. |
| Cruciferous vegetables (large amounts) | Mild induction | Usually not clinically significant. |
| Charbroiled foods | Mild induction | Usually not clinically significant. |
| Rifampin | Strong induction | Decreases substrate levels. |
| Carbamazepine | Moderate induction | Decreases substrate levels. |
| Phenytoin | Moderate induction | Decreases substrate levels. |

---

## P-glycoprotein (P-gp) / MDR1

Not a CYP enzyme but a critical drug efflux transporter. Often overlaps with CYP3A4 substrates/inhibitors.

### Substrates
| Drug | Category |
|------|----------|
| Digoxin | Cardiac glycoside (NTI) |
| Dabigatran | Anticoagulant |
| Edoxaban | Anticoagulant |
| Cyclosporine | Immunosuppressant |
| Tacrolimus | Immunosuppressant |
| Colchicine | Anti-gout |
| Loperamide | Antidiarrheal (keeps it out of CNS) |

### Inhibitors (increase substrate levels)
| Drug | Category |
|------|----------|
| Amiodarone | Antiarrhythmic |
| Verapamil | Calcium channel blocker |
| Dronedarone | Antiarrhythmic |
| Cyclosporine | Immunosuppressant |
| Ketoconazole | Antifungal |
| Ritonavir | HIV protease inhibitor |
| Clarithromycin | Macrolide |
| Quinidine | Antiarrhythmic |

### Inducers (decrease substrate levels)
| Drug | Category |
|------|----------|
| Rifampin | Antimycobacterial |
| Carbamazepine | Antiepileptic |
| Phenytoin | Antiepileptic |
| St. John's Wort | Herbal |

---

## Quick Decision Matrix

When evaluating a drug pair for CYP450 interaction:

1. Identify the metabolic pathway(s) of each drug
2. Check if one drug inhibits/induces the pathway of the other
3. Assess the clinical significance:
   - Is the substrate an NTI drug? -> High significance
   - Is the inhibitor strong or moderate? -> Strong = clinically significant; Moderate = monitor
   - Is the substrate primarily dependent on this pathway? -> Sole pathway = highest risk
4. Check patient factors:
   - Known CYP2D6/CYP2C19 poor metabolizer? -> Already reduced activity
   - Renal/hepatic impairment? -> Compounds the effect
   - Age >75? -> Generally reduced CYP activity
