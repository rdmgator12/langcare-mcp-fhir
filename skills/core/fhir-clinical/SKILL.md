# FHIR Clinical Workflows

Structured clinical workflows using LangCare MCP FHIR Server's 4 generic tools: `fhir_search`, `fhir_read`, `fhir_create`, `fhir_update`.

---

## 1. Patient Chart Review

**Purpose:** Comprehensive review of patient's clinical record.

### Workflow

**Step 1: Verify Patient Identity**
```
Tool: fhir_search
resourceType: "Patient"
queryParams: "family=[lastname]&given=[firstname]&birthdate=[YYYY-MM-DD]"
```
- Confirm demographics match (name, DOB, MRN)
- If multiple matches, ask user to specify

**Step 2: Retrieve Active Problems**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```
- Group by category (problem, diagnosis, symptom)
- Note onset dates and verification status

**Step 3: Get Current Medications**
```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```
- Include medication name, dose, frequency, route
- Check for high-risk medications (anticoagulants, insulin, opioids)

**Step 4: Review Recent Lab Results**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&date=ge[30-days-ago]"
```
- Present with reference ranges
- Flag abnormal values (High/Low indicators)

**Step 5: Check Recent Encounters**
```
Tool: fhir_search
resourceType: "Encounter"
queryParams: "patient=[patient-id]&date=ge[90-days-ago]"
```
- Include type, date, provider, location
- Note hospital admissions vs outpatient visits

**Step 6: Review Allergies**
```
Tool: fhir_search
resourceType: "AllergyIntolerance"
queryParams: "patient=[patient-id]&clinical-status=active"
```
- Include substance, reaction, severity
- Distinguish true allergies from intolerances

### Summary Format
Present findings in clinical order:
1. **Demographics** - Name, DOB, MRN
2. **Active Problems** - Conditions with dates
3. **Medications** - Name, dose, frequency
4. **Allergies** - Substance and reaction
5. **Recent Labs** - Results with flags
6. **Recent Visits** - Date, type, provider

---

## 2. Medication Reconciliation

**Purpose:** Ensure accurate, complete medication list during care transitions.

### Workflow

**Step 1: Retrieve All Medication Sources**
```
Tool: fhir_search (run in parallel)
1. resourceType: "MedicationStatement" (patient-reported)
2. resourceType: "MedicationRequest" (prescribed)
3. resourceType: "MedicationAdministration" (administered, if available)
```

**Step 2: Identify Discrepancies**
- Compare lists for duplicates (same ingredient, different name)
- Check for conflicting doses
- Note discontinued medications still listed as active

**Step 3: Validate Each Medication**
- **Indication:** Does it match active conditions?
- **Dose:** Within therapeutic range?
- **Drug-Drug Interactions:** Check high-risk combinations:
  - Warfarin + NSAIDs (bleeding risk)
  - ACE inhibitors + K+ supplements (hyperkalemia)
  - MAOIs + SSRIs (serotonin syndrome)
- **Contraindications:** Check against allergies and conditions

**Step 4: Check for Gaps**
- Is there treatment for all active chronic conditions?
- Are guideline-recommended medications present (e.g., statin for diabetes)?

**Step 5: Document Reconciled List**
```
Tool: fhir_update (for each medication)
resourceType: "MedicationStatement"
id: "[medication-id]"
resource: { confirmed status, last_reviewed timestamp }
```

### Safety Checks
- **High-Alert Medications:** Insulin, heparin, warfarin, opioids, chemotherapy
- **Renal Dosing:** Adjust for CrCl if available
- **Geriatric Concerns:** BEERS criteria for patients 65+
- **Pregnancy:** Category X medications in women of childbearing age

---

## 3. Lab Result Interpretation

**Purpose:** Review and interpret laboratory results in clinical context.

### Workflow

**Step 1: Retrieve Lab Results**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&category=laboratory&date=ge[date]"
```

**Step 2: Organize by Category**
- **CBC:** WBC, Hemoglobin, Hematocrit, Platelets
- **BMP/CMP:** Sodium, Potassium, Creatinine, Glucose, Calcium
- **Liver Panel:** ALT, AST, Alkaline Phosphatase, Bilirubin
- **Lipids:** Total Cholesterol, LDL, HDL, Triglycerides
- **Other:** TSH, HbA1c, Vitamin D, etc.

**Step 3: Flag Abnormal Values**
- Parse `interpretation` field (High, Low, Critical)
- Compare to `referenceRange` if present
- Calculate delta from previous results (trend analysis)

**Step 4: Clinical Correlation**
- Link to relevant conditions (e.g., HbA1c → Diabetes)
- Note medication effects (e.g., Creatinine elevation on ACE inhibitor)
- Identify critical values requiring immediate action

**Step 5: Trend Analysis** (if historical data available)
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=[LOINC-code]&date=ge[6-months-ago]"
```
- Plot trend: improving, worsening, stable
- Note response to interventions

### Common LOINC Codes
- **Glucose:** 2339-0 (mg/dL), 2345-7 (mmol/L)
- **Creatinine:** 2160-0 (mg/dL)
- **HbA1c:** 4548-4 (%)
- **TSH:** 3016-3 (mIU/L)
- **Total Cholesterol:** 2093-3 (mg/dL)
- **LDL:** 18262-6 (mg/dL)

### Interpretation Guidelines
- **Critical Values:** WBC <1.5 or >30, Platelets <20, K+ <2.5 or >6.5, Glucose <40 or >600
- **Reference Ranges:** Use provided ranges; may vary by lab
- **Clinical Context:** Always interpret in context of symptoms and history

---

## 4. Preventive Care Screening

**Purpose:** Identify overdue preventive care based on guidelines (USPSTF, ADA, ACC/AHA).

### Workflow

**Step 1: Get Patient Demographics**
```
Tool: fhir_read
resourceType: "Patient"
id: "[patient-id]"
```
Extract: age, sex, smoking status

**Step 2: Check Active Conditions**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&clinical-status=active"
```
Look for: Diabetes, Hypertension, Hyperlipidemia, CKD

**Step 3: Review Previous Screenings**
```
Tool: fhir_search (run for each screening type)
resourceType: "Observation" or "Procedure"
queryParams: "patient=[patient-id]&code=[screening-code]"
```

### Screening Recommendations

**Cancer Screening:**
- **Colorectal (50-75 years):** Colonoscopy q10y, FIT stool test annually
- **Breast (50-74 years, female):** Mammography q2y
- **Cervical (21-65 years, female):** Pap smear q3y or HPV q5y
- **Lung (55-80 years, 30 pack-year smoking):** Low-dose CT annually
- **Prostate (discuss 55-69 years, male):** Shared decision-making on PSA

**Cardiovascular:**
- **Blood Pressure:** All adults annually
- **Lipids:** Age 40-75, calculate 10-year ASCVD risk
- **Diabetes:** Age 35-70 with overweight/obesity, q3y

**Immunizations:**
- **Influenza:** Annual, all adults
- **Pneumococcal:** Age 65+, or high-risk
- **Shingles:** Age 50+, Shingrix x2
- **Tdap:** Every 10 years
- **COVID-19:** Per current CDC guidance

**Step 4: Identify Overdue Screenings**
- Calculate time since last screening
- Compare to guideline intervals
- Prioritize by risk and time overdue

**Step 5: Generate Recommendations**
Present as structured list:
```
Overdue Screenings:
1. Colorectal Cancer Screening - Last: 2015 (8 years ago) - Recommend: Colonoscopy or FIT test
2. Mammography - Last: 2021 (2 years ago) - Recommend: Schedule mammogram
3. Lipid Panel - Last: 2019 (4 years ago) - Recommend: Check cholesterol
```

---

## 5. Diabetes Management

**Purpose:** Monitor and optimize diabetes care using FHIR resources.

### Workflow

**Step 1: Confirm Diabetes Diagnosis**
```
Tool: fhir_search
resourceType: "Condition"
queryParams: "patient=[patient-id]&code=http://snomed.info/sct|44054006"
```
SNOMED code 44054006 = Diabetes Mellitus Type 2

**Step 2: Retrieve Glucose Control Metrics**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=4548-4&date=ge[3-months-ago]"
```
LOINC 4548-4 = Hemoglobin A1c

**Target:** HbA1c <7% (individualize to <6.5% or <8% based on comorbidities)

**Step 3: Get Current Diabetes Medications**
```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```
Filter for:
- Metformin (first-line)
- SGLT2 inhibitors (cardiovascular/renal protection)
- GLP-1 agonists (weight loss, CV benefit)
- DPP-4 inhibitors
- Sulfonylureas (hypoglycemia risk)
- Insulin (basal, bolus, mixed)

**Step 4: Check for Complications Screening**
```
Tool: fhir_search (run in parallel)
1. Retinopathy: resourceType="Procedure", code=[eye-exam-code]
2. Nephropathy: resourceType="Observation", code=[urine-albumin-code]
3. Neuropathy: resourceType="Procedure", code=[foot-exam-code]
```

**Screening Intervals:**
- **Retinal exam:** Annually (can extend to q2y if no retinopathy)
- **Urine albumin/creatinine ratio:** Annually
- **Foot exam:** Annually, more often if neuropathy present

**Step 5: Cardiovascular Risk Assessment**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=18262-6" (LDL)
```
- **Statin:** Indicated for all diabetics 40-75 years
- **Blood pressure:** Target <130/80 mmHg
- **Aspirin:** Consider if ASCVD risk ≥10%

**Step 6: Generate Care Plan**
```
Tool: fhir_create
resourceType: "CarePlan"
resource: {
  status: "active",
  intent: "plan",
  subject: { reference: "Patient/[id]" },
  category: [{ coding: [{ system: "http://snomed.info/sct", code: "698358001", display: "Diabetes self-management plan" }] }],
  activity: [
    { detail: { description: "Check HbA1c every 3 months until at goal" } },
    { detail: { description: "Annual dilated eye exam" } },
    { detail: { description: "Check urine albumin annually" } }
  ]
}
```

### ADA Quality Measures
- **HbA1c <7%:** Glycemic control
- **BP <140/90 mmHg:** Hypertension control
- **LDL <100 mg/dL:** Lipid control
- **On ACE/ARB:** If hypertensive or albuminuric
- **On statin:** If age 40-75

---

## 6. Hypertension Management

**Purpose:** Monitor blood pressure and optimize antihypertensive therapy.

### Workflow

**Step 1: Retrieve Blood Pressure Readings**
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=85354-9&date=ge[6-months-ago]"
```
LOINC 85354-9 = Blood Pressure Panel
- Extract systolic (8480-6) and diastolic (8462-4) components
- Calculate average of most recent 3 readings

**Step 2: Classify Hypertension Stage**
- **Normal:** <120/80 mmHg
- **Elevated:** 120-129/<80 mmHg
- **Stage 1:** 130-139/80-89 mmHg
- **Stage 2:** ≥140/90 mmHg
- **Hypertensive Crisis:** ≥180/120 mmHg (urgent evaluation)

**Step 3: Review Current Antihypertensives**
```
Tool: fhir_search
resourceType: "MedicationStatement"
queryParams: "patient=[patient-id]&status=active"
```
Medication classes:
- **ACE Inhibitors:** -pril (lisinopril, enalapril)
- **ARBs:** -sartan (losartan, valsartan)
- **CCBs:** -dipine (amlodipine, nifedipine)
- **Diuretics:** HCTZ, chlorthalidone, furosemide
- **Beta-blockers:** -olol (metoprolol, atenolol)

**Step 4: Assess for Target Organ Damage**
```
Tool: fhir_search (parallel)
1. Creatinine/eGFR (renal)
2. Echocardiogram (LVH)
3. Retinal exam (hypertensive retinopathy)
4. Urine albumin (proteinuria)
```

**Step 5: Optimize Therapy**
- **Lifestyle:** Weight loss, DASH diet, sodium <2g/day, exercise
- **First-line:** ACE/ARB, CCB, or thiazide diuretic
- **Combination therapy:** If Stage 2, start 2 drugs
- **Monitoring:** BP recheck in 1 month after change

### Special Populations
- **CKD:** Prefer ACE/ARB (nephroprotection)
- **Heart Failure:** ACE/ARB + beta-blocker + diuretic
- **Post-MI:** Beta-blocker + ACE/ARB
- **Diabetes:** Target <130/80, use ACE/ARB
- **Elderly:** Individualize target (may accept <150/90)

---

## 7. Clinical Documentation

**Purpose:** Create structured clinical notes using FHIR resources.

### Workflow

**Step 1: Gather Visit Information**
```
Tool: fhir_read
resourceType: "Encounter"
id: "[encounter-id]"
```
Extract: date, type, location, provider

**Step 2: Create Composition Resource**
```
Tool: fhir_create
resourceType: "Composition"
resource: {
  status: "final",
  type: { coding: [{ system: "http://loinc.org", code: "11506-3", display: "Progress note" }] },
  subject: { reference: "Patient/[id]" },
  encounter: { reference: "Encounter/[id]" },
  date: "[timestamp]",
  author: [{ reference: "Practitioner/[id]" }],
  title: "Progress Note",
  section: [
    { title: "Chief Complaint", text: { status: "generated", div: "<div>...</div>" } },
    { title: "History of Present Illness", ... },
    { title: "Assessment", ... },
    { title: "Plan", ... }
  ]
}
```

**Step 3: Link Supporting Resources**
- Observations (labs, vitals)
- Conditions (diagnoses)
- Procedures performed
- MedicationRequests (prescriptions)

**Step 4: Add Provenance**
```
Tool: fhir_create
resourceType: "Provenance"
resource: {
  target: [{ reference: "Composition/[id]" }],
  recorded: "[timestamp]",
  agent: [{
    who: { reference: "Practitioner/[id]" },
    role: [{ coding: [{ code: "author" }] }]
  }],
  signature: [{ type: [{ code: "1.2.840.10065.1.12.1.1" }], when: "[timestamp]", who: { reference: "Practitioner/[id]" } }]
}
```

### Common Note Types (LOINC)
- **Progress Note:** 11506-3
- **History & Physical:** 34117-2
- **Discharge Summary:** 18842-5
- **Consultation Note:** 11488-4
- **Operative Note:** 11504-8

---

## 8. Vital Signs Tracking

**Purpose:** Record and monitor vital signs over time.

### Create Vital Sign Observation

**Blood Pressure:**
```
Tool: fhir_create
resourceType: "Observation"
resource: {
  status: "final",
  category: [{ coding: [{ system: "http://terminology.hl7.org/CodeSystem/observation-category", code: "vital-signs" }] }],
  code: { coding: [{ system: "http://loinc.org", code: "85354-9", display: "Blood pressure panel" }] },
  subject: { reference: "Patient/[id]" },
  effectiveDateTime: "[timestamp]",
  component: [
    { code: { coding: [{ system: "http://loinc.org", code: "8480-6", display: "Systolic blood pressure" }] }, valueQuantity: { value: 120, unit: "mmHg" } },
    { code: { coding: [{ system: "http://loinc.org", code: "8462-4", display: "Diastolic blood pressure" }] }, valueQuantity: { value: 80, unit: "mmHg" } }
  ]
}
```

**Other Vital Signs:**
- **Heart Rate:** LOINC 8867-4 (beats/min)
- **Respiratory Rate:** LOINC 9279-1 (breaths/min)
- **Temperature:** LOINC 8310-5 (°C or °F)
- **Oxygen Saturation:** LOINC 2708-6 (%)
- **Weight:** LOINC 29463-7 (kg or lb)
- **Height:** LOINC 8302-2 (cm or in)
- **BMI:** LOINC 39156-5 (kg/m²)

### Trending
```
Tool: fhir_search
resourceType: "Observation"
queryParams: "patient=[patient-id]&code=8480-6&date=ge[6-months-ago]&_sort=date"
```
- Plot trend over time
- Identify patterns (e.g., morning vs evening BP)
- Assess treatment response

---

## Safety Considerations

### Patient Identity Verification
- Always confirm with at least 2 identifiers (name + DOB, name + MRN)
- Ask user to verify before any create/update operations
- Use partial matches cautiously (e.g., "John Smith" → multiple results)

### Data Validation
- Check for required FHIR fields before creating resources
- Validate codes against proper code systems (LOINC, SNOMED, RxNorm)
- Ensure dates are in ISO 8601 format (YYYY-MM-DD)
- Verify references point to existing resources

### Clinical Accuracy
- Always present reference ranges with lab values
- Note data source and recency (e.g., "Last measured 2 days ago")
- Distinguish between patient-reported and clinician-verified data
- Flag missing critical information (e.g., allergy status unknown)

### Error Handling
- If search returns 0 results, suggest alternative searches
- If authentication fails, check OAuth scopes
- If resource not found, verify ID and resource type
- If create/update fails, check for validation errors in response

### Documentation
- Log all PHI access in audit logs
- Note data source for clinical decisions
- Include timestamps for all retrievals
- Document reasoning for clinical recommendations

---

## Appendix: Common Search Parameters

### Patient
- `name` - Family and/or given name
- `birthdate` - Date of birth (YYYY-MM-DD)
- `identifier` - MRN or other identifier
- `gender` - male, female, other, unknown

### Observation (Labs/Vitals)
- `patient` - Patient reference
- `code` - LOINC code
- `category` - laboratory, vital-signs, survey
- `date` - Date range (ge, le, eq)
- `status` - final, preliminary, amended

### Condition (Diagnoses)
- `patient` - Patient reference
- `code` - SNOMED CT or ICD-10 code
- `clinical-status` - active, inactive, resolved
- `onset-date` - When condition began

### MedicationStatement/MedicationRequest
- `patient` - Patient reference
- `status` - active, completed, stopped
- `code` - RxNorm code
- `effective` - Date range

### Encounter
- `patient` - Patient reference
- `date` - Date range
- `type` - Visit type code
- `status` - planned, arrived, in-progress, finished

### Date Modifiers
- `ge` - Greater than or equal (≥)
- `le` - Less than or equal (≤)
- `gt` - Greater than (>)
- `lt` - Less than (<)
- `eq` - Equal (=)

**Example:** `date=ge2023-01-01&date=le2023-12-31` (year 2023)
