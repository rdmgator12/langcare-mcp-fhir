# CCD (Continuity of Care Document) Section Reference

## Required CCD Sections per HL7 CDA R2

| Section | LOINC Code | Description |
|---------|------------|-------------|
| Problems | 11450-4 | Active problem list with ICD-10/SNOMED codes |
| Medications | 10160-0 | Active medication list with dosing |
| Allergies | 48765-2 | Allergy and adverse reaction list |
| Results | 30954-2 | Laboratory and diagnostic results |
| Vital Signs | 8716-3 | Most recent vital sign set |
| Procedures | 47519-4 | Procedures performed |
| Immunizations | 11369-6 | Immunization history |
| Plan of Care | 18776-5 | Active care plans and goals |

## Optional CCD Sections

| Section | LOINC Code | Description |
|---------|------------|-------------|
| Encounters | 46240-8 | Recent encounter history |
| Social History | 29762-2 | Smoking, alcohol, substance use |
| Family History | 10157-6 | Family medical history |
| Functional Status | 47420-5 | ADL/IADL assessments |
| Advance Directives | 42348-3 | Living will, POA status |
| Medical Equipment | 46264-8 | DME and implanted devices |
| Payers | 48768-6 | Insurance coverage |
| Reason for Referral | 42349-1 | When summary is for a referral |

## FHIR Resource to CCD Section Mapping

| CCD Section | Primary FHIR Resource | Query Strategy |
|-------------|----------------------|----------------|
| Problems | Condition | `clinical-status=active` |
| Medications | MedicationRequest | `status=active` |
| Allergies | AllergyIntolerance | `clinical-status=active` |
| Results | Observation | `category=laboratory&date=ge[90-days-ago]` |
| Vital Signs | Observation | `category=vital-signs&_sort=-date&_count=10` |
| Procedures | Procedure | `date=ge[12-months-ago]` |
| Immunizations | Immunization | `status=completed` |
| Plan of Care | CarePlan | `status=active` |
| Social History | Observation | `category=social-history` |
| Encounters | Encounter | `date=ge[12-months-ago]&_sort=-date` |

## Common Lab Panel LOINC Codes

| Panel | LOINC | Key Components |
|-------|-------|----------------|
| CBC | 58410-2 | WBC (26464-8), Hgb (718-7), Hct (4544-3), Plt (777-3) |
| BMP | 51990-0 | Na (2951-2), K (2823-3), Cl (2075-0), CO2 (1963-8), BUN (3094-0), Cr (2160-0), Glu (2345-7) |
| CMP | 24323-8 | BMP + Albumin (1751-7), ALP (6768-6), ALT (1742-6), AST (1920-8), Bilirubin (1975-2), Total Protein (2885-2) |
| Lipid Panel | 57698-3 | TC (2093-3), LDL (18262-6), HDL (2085-9), TG (2571-8) |
| HbA1c | 4548-4 | Standalone |
| TSH | 3016-3 | Often with Free T4 (3024-7) |

## Vital Sign LOINC Codes

| Vital | LOINC | Units |
|-------|-------|-------|
| Temperature | 8310-5 | Cel or [degF] |
| Heart Rate | 8867-4 | /min |
| Respiratory Rate | 9279-1 | /min |
| Blood Pressure | 85354-9 | Panel: systolic 8480-6, diastolic 8462-4 (mmHg) |
| SpO2 | 59408-5 | % |
| Weight | 29463-7 | kg or [lb_av] |
| Height | 8302-2 | cm or [in_i] |
| BMI | 39156-5 | kg/m2 |
