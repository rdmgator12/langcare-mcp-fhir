# Medication-to-Condition Implied Mappings

## Purpose

When a patient is prescribed a medication, it often implies the presence of a specific condition. Use this reference to cross-reference active medications against the problem list and identify undocumented diagnoses.

## Mapping Strategy

1. Match by drug class (preferred) or individual drug name
2. Check if any condition in the expected SNOMED/ICD-10 code set exists on the active problem list
3. Flag as "implied diagnosis missing" only if no matching condition is found

## Comprehensive Medication-Condition Mappings

### Endocrine / Metabolic

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| Biguanides | Metformin | Type 2 Diabetes Mellitus | 44054006 | E11.x |
| Sulfonylureas | Glipizide, glyburide, glimepiride | Type 2 Diabetes Mellitus | 44054006 | E11.x |
| SGLT2 Inhibitors | Empagliflozin, dapagliflozin, canagliflozin | Type 2 Diabetes Mellitus (also HF, CKD) | 44054006 | E11.x |
| GLP-1 Agonists | Semaglutide, liraglutide, dulaglutide | Type 2 Diabetes Mellitus or Obesity | 44054006 / 414916001 | E11.x / E66.x |
| DPP-4 Inhibitors | Sitagliptin, saxagliptin, linagliptin | Type 2 Diabetes Mellitus | 44054006 | E11.x |
| Thiazolidinediones | Pioglitazone, rosiglitazone | Type 2 Diabetes Mellitus | 44054006 | E11.x |
| Insulin (all forms) | Glargine, lispro, aspart, NPH, regular | Diabetes Mellitus (Type 1 or 2) | 73211009 / 44054006 | E10.x / E11.x |
| Levothyroxine | Synthroid, Levoxyl | Hypothyroidism | 40930008 | E03.x |
| Methimazole, PTU | Tapazole | Hyperthyroidism | 34486009 | E05.x |
| Alendronate, risedronate | Fosamax, Actonel | Osteoporosis | 64859006 | M81.x |
| Testosterone | Various formulations | Hypogonadism | 48130008 | E29.1 |
| Estradiol, conjugated estrogens | Premarin, Estrace | Menopausal symptoms / HRT | 161712005 | N95.1 |

### Cardiovascular

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| ACE Inhibitors | Lisinopril, enalapril, ramipril, benazepril | Hypertension (also HF, CKD, post-MI) | 59621000 | I10 |
| ARBs | Losartan, valsartan, irbesartan, olmesartan | Hypertension (also HF, CKD) | 59621000 | I10 |
| Calcium Channel Blockers | Amlodipine, nifedipine, diltiazem, verapamil | Hypertension (also angina, arrhythmia) | 59621000 | I10 |
| Thiazide Diuretics | HCTZ, chlorthalidone, indapamide | Hypertension | 59621000 | I10 |
| Loop Diuretics | Furosemide, bumetanide, torsemide | Heart Failure or Edema | 84114007 | I50.x |
| Beta-Blockers | Metoprolol, atenolol, carvedilol, bisoprolol | Hypertension, HF, A-fib, or post-MI | 59621000 / 84114007 / 49436004 | I10 / I50.x / I48.x |
| Statins | Atorvastatin, rosuvastatin, simvastatin, pravastatin | Hyperlipidemia | 55822004 | E78.x |
| Ezetimibe | Zetia | Hyperlipidemia | 55822004 | E78.x |
| PCSK9 Inhibitors | Evolocumab, alirocumab | Hyperlipidemia (usually familial) | 398036000 | E78.01 |
| Warfarin | Coumadin | A-fib, DVT/PE, mechanical valve | 49436004 / 128053003 | I48.x / I26.x |
| DOACs | Apixaban, rivaroxaban, edoxaban, dabigatran | A-fib, DVT/PE | 49436004 / 128053003 | I48.x / I26.x |
| Antiplatelet (dual) | Clopidogrel, ticagrelor, prasugrel | CAD, post-stent, post-MI | 53741008 | I25.x |
| Nitrates | Nitroglycerin, isosorbide mononitrate | Angina / CAD | 194828000 | I20.x / I25.x |
| Digoxin | Lanoxin | Heart failure or A-fib | 84114007 / 49436004 | I50.x / I48.x |
| Amiodarone | Cordarone | Atrial fibrillation or ventricular arrhythmia | 49436004 | I48.x / I49.x |
| Hydralazine + Isosorbide dinitrate | BiDil | Heart failure (especially in Black patients) | 84114007 | I50.x |
| Spironolactone, eplerenone | Aldactone, Inspra | Heart failure or hyperaldosteronism | 84114007 | I50.x |

### Respiratory

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| SABA (rescue inhaler) | Albuterol, levalbuterol | Asthma or COPD | 195967001 / 13645005 | J45.x / J44.x |
| ICS | Fluticasone, budesonide, beclomethasone | Asthma or COPD | 195967001 / 13645005 | J45.x / J44.x |
| ICS/LABA | Fluticasone/salmeterol, budesonide/formoterol | Asthma or COPD | 195967001 / 13645005 | J45.x / J44.x |
| LAMA | Tiotropium, umeclidinium | COPD | 13645005 | J44.x |
| Montelukast | Singulair | Asthma or allergic rhinitis | 195967001 / 61582004 | J45.x / J30.x |
| Roflumilast | Daliresp | COPD (severe) | 13645005 | J44.1 |

### Gastrointestinal

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| PPIs | Omeprazole, pantoprazole, esomeprazole, lansoprazole | GERD or peptic ulcer | 235595009 / 13200003 | K21.0 / K25.x-K27.x |
| H2 Blockers | Famotidine, ranitidine | GERD or peptic ulcer | 235595009 | K21.0 |
| Mesalamine | Asacol, Lialda | Ulcerative colitis or Crohn's | 64766004 / 34000006 | K51.x / K50.x |
| Lactulose, rifaximin | Kristalose, Xifaxan | Hepatic encephalopathy / cirrhosis | 19943007 | K72.x / K74.x |
| Ursodiol | Actigall | Primary biliary cholangitis or gallstones | 31712002 / 235919008 | K74.3 / K80.x |
| Pancrelipase | Creon | Exocrine pancreatic insufficiency | 47367009 | K86.81 |

### Psychiatric / Neurological

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| SSRIs | Sertraline, fluoxetine, escitalopram, citalopram, paroxetine | Depression or Anxiety | 35489007 / 197480006 | F32.x / F41.x |
| SNRIs | Venlafaxine, duloxetine, desvenlafaxine | Depression, Anxiety, or Neuropathic pain | 35489007 / 197480006 | F32.x / F41.x |
| TCAs | Amitriptyline, nortriptyline | Depression, Neuropathic pain, or Migraine prophylaxis | 35489007 | F32.x |
| Bupropion | Wellbutrin | Depression or Smoking cessation | 35489007 | F32.x |
| Lithium | Lithobid | Bipolar disorder | 13746004 | F31.x |
| Valproate, lamotrigine, carbamazepine | Depakote, Lamictal, Tegretol | Epilepsy or Bipolar disorder | 84757009 / 13746004 | G40.x / F31.x |
| Levetiracetam, topiramate, phenytoin | Keppra, Topamax, Dilantin | Epilepsy / Seizure disorder | 84757009 | G40.x |
| Gabapentin, pregabalin | Neurontin, Lyrica | Neuropathy, Seizure disorder, or Chronic pain | 386033004 / 84757009 | G62.x / G40.x |
| Donepezil, rivastigmine, memantine | Aricept, Exelon, Namenda | Alzheimer's / Dementia | 26929004 | G30.x / F03.x |
| Levodopa/carbidopa | Sinemet | Parkinson's disease | 49049000 | G20 |
| Sumatriptan, rizatriptan | Imitrex, Maxalt | Migraine | 37796009 | G43.x |
| Stimulants | Methylphenidate, amphetamine salts | ADHD | 406506008 | F90.x |
| Antipsychotics | Quetiapine, risperidone, olanzapine, aripiprazole | Schizophrenia, Bipolar, or psychotic features | 58214004 / 13746004 | F20.x / F31.x |

### Rheumatologic / Immunologic

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| Methotrexate | Trexall | Rheumatoid arthritis, Psoriasis, or Psoriatic arthritis | 69896004 / 9014002 | M06.x / L40.x |
| Hydroxychloroquine | Plaquenil | SLE or Rheumatoid arthritis | 55464009 / 69896004 | M32.x / M06.x |
| Allopurinol, febuxostat | Zyloprim, Uloric | Gout | 90560007 | M10.x |
| Colchicine | Colcrys | Gout or Familial Mediterranean Fever | 90560007 | M10.x |
| TNF inhibitors | Adalimumab, etanercept, infliximab | RA, Crohn's, UC, Psoriasis, AS | 69896004 / 34000006 | M06.x / K50.x |
| JAK inhibitors | Tofacitinib, baricitinib, upadacitinib | Rheumatoid arthritis | 69896004 | M06.x |

### Renal

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| Phosphate binders | Sevelamer, calcium acetate | CKD (advanced) | 709044004 | N18.x |
| Erythropoietin stimulating agents | Epoetin alfa, darbepoetin | CKD with anemia | 709044004 | N18.x + D63.1 |
| Sodium bicarbonate (chronic) | Oral bicarbonate | CKD with metabolic acidosis | 709044004 | N18.x |
| Calcitriol | Rocaltrol | CKD with secondary hyperparathyroidism | 709044004 | N18.x + E21.1 |

### Hematologic

| Medication / Class | Examples | Implied Condition | SNOMED CT | ICD-10 |
|--------------------|----------|-------------------|-----------|--------|
| Iron supplements (prescription) | Ferrous sulfate, IV iron | Iron deficiency anemia | 87522002 | D50.x |
| B12 injections | Cyanocobalamin | B12 deficiency / Pernicious anemia | 190606006 | D51.x |
| Folic acid (high dose) | 1mg+ daily | Folate deficiency or Methotrexate adjunct | 190604008 | D52.x |
| Hydroxyurea | Droxia | Sickle cell disease or CML | 417357006 | D57.x |

## Important Caveats

1. **Off-label use is common.** Gabapentin for anxiety, topiramate for migraine, amitriptyline for IBS -- do not assume the most common indication is the only possibility.
2. **Multi-indication drugs.** SGLT2 inhibitors may be prescribed for diabetes, heart failure, or CKD independently. Check all three conditions before flagging.
3. **Dose-dependent indications.** Low-dose aspirin (81mg) implies cardiovascular prophylaxis. High-dose aspirin (650mg+) may indicate pain management.
4. **Combination medications.** Lisinopril/HCTZ implies both hypertension and diuretic need. Count as one indication (hypertension).
5. **Discontinued medications.** Only flag based on ACTIVE medications. A stopped statin does not imply current hyperlipidemia.
6. **Supplements vs prescriptions.** OTC vitamin D or calcium is not a strong indicator of a specific condition. Prescription-strength formulations are more meaningful.
