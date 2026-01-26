# EPIC OAuth2 Scopes Reference

Complete reference for EPIC Backend Services (system-level) OAuth2 scopes.

---

## Scope Format

**Backend Services Auth (System-Level Access):**
- `system/ResourceType.read` - Read access to resource type
- `system/ResourceType.write` - Create and update access (includes read)

**Important Notes:**
- All scopes must be registered in EPIC App Orchard before use
- System-level scopes are for backend services (server-to-server)
- User-level scopes (`user/`) are for patient-facing apps (not used here)
- Launch scopes (`launch/`) are for EHR-launched apps (not used here)

---

## Common FHIR Resources

### Core Patient Data

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Patient.read` | Read patient demographics | Patient lookup, demographics display |
| `system/Patient.write` | Create/update patient records | New patient registration, demographic updates |

### Clinical Observations

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Observation.read` | Read labs, vitals, measurements | Lab results, vital signs, clinical observations |
| `system/Observation.write` | Create/update observations | Recording vitals, entering lab results |

### Conditions & Diagnoses

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Condition.read` | Read patient conditions/diagnoses | Problem list, diagnosis history |
| `system/Condition.write` | Create/update conditions | Adding diagnoses, updating problem list |

### Medications

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/MedicationRequest.read` | Read medication orders | Active medications, medication history |
| `system/MedicationRequest.write` | Create/update med orders | Prescribing medications, modifying orders |
| `system/MedicationStatement.read` | Read patient-reported medications | Patient medication list, reconciliation |
| `system/MedicationStatement.write` | Record patient medications | Med reconciliation, patient-reported meds |
| `system/Medication.read` | Read medication definitions | Drug information, formulary lookup |

### Encounters & Visits

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Encounter.read` | Read patient visits/encounters | Visit history, encounter details |
| `system/Encounter.write` | Create/update encounters | Documenting visits, updating encounter info |

### Procedures

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Procedure.read` | Read procedures performed | Surgical history, procedure documentation |
| `system/Procedure.write` | Document procedures | Recording procedures, updating procedure notes |

### Diagnostic Reports

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/DiagnosticReport.read` | Read diagnostic reports | Lab reports, imaging results, pathology |
| `system/DiagnosticReport.write` | Create diagnostic reports | Entering results, completing reports |

### Allergies & Intolerances

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/AllergyIntolerance.read` | Read patient allergies | Allergy checking, medication safety |
| `system/AllergyIntolerance.write` | Document allergies | Recording new allergies, updating severity |

### Immunizations

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Immunization.read` | Read vaccination history | Immunization records, compliance checking |
| `system/Immunization.write` | Record immunizations | Documenting vaccines, updating immunization record |

---

## Documents & Care Coordination

### Clinical Documentation

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/DocumentReference.read` | Read document metadata | Progress notes, discharge summaries, clinical reports |
| `system/DocumentReference.write` | Create document references | Attaching external records, uploading documents |
| `system/Binary.read` | Read document content (raw bytes) | PDF files, images, scanned documents |
| `system/Composition.read` | Read structured clinical documents | C-CDA documents, clinical notes, discharge summaries |
| `system/Composition.write` | Create structured documents | Generating C-CDA, creating clinical notes |
| `system/Media.read` | Read media attachments | Photos, videos, audio recordings, images |
| `system/Media.write` | Upload media | Uploading photos, videos, diagnostic images |
| `system/QuestionnaireResponse.read` | Read structured assessments | Intake forms, patient questionnaires, clinical assessments |
| `system/QuestionnaireResponse.write` | Record questionnaire responses | Documenting assessments, patient-reported outcomes |
| `system/Provenance.read` | Read document provenance | Audit trails, digital signatures, document history |

### Care Management

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/CarePlan.read` | Read care plans | Treatment plans, care coordination |
| `system/CarePlan.write` | Create/update care plans | Care planning, treatment protocols |
| `system/Goal.read` | Read patient goals | Care goals, treatment objectives |
| `system/Goal.write` | Set patient goals | Goal setting, care planning |

### Communication

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Communication.read` | Read communications | Messages, alerts, notifications |
| `system/Communication.write` | Send communications | Messaging, alerts, care team communication |
| `system/CommunicationRequest.read` | Read communication requests | Consult requests, communication tasks |
| `system/CommunicationRequest.write` | Create communication requests | Requesting consults, care team coordination |

---

## Scheduling & Appointments

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Appointment.read` | Read appointments | Appointment history, scheduling |
| `system/Appointment.write` | Schedule appointments | Booking appointments, rescheduling |
| `system/Schedule.read` | Read provider schedules | Availability checking, schedule viewing |
| `system/Slot.read` | Read available time slots | Appointment booking, availability |

---

## Billing & Coverage

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Coverage.read` | Read insurance coverage | Insurance verification, eligibility |
| `system/Claim.read` | Read claims | Billing history, claim status |
| `system/Claim.write` | Submit claims | Claims submission, billing |
| `system/ExplanationOfBenefit.read` | Read EOBs | Payment information, claim adjudication |

---

## Providers & Organizations

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Practitioner.read` | Read provider information | Provider lookup, directory |
| `system/PractitionerRole.read` | Read provider roles | Provider specialties, practice locations |
| `system/Organization.read` | Read organization data | Facility information, organizational structure |
| `system/Location.read` | Read location information | Facility locations, departments |

---

## Advanced Clinical Resources

### Orders & Specimens

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/ServiceRequest.read` | Read service/lab orders | Order tracking, lab orders |
| `system/ServiceRequest.write` | Create service orders | Ordering labs, requesting services |
| `system/Specimen.read` | Read specimen information | Specimen tracking, lab workflow |

### Devices

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/Device.read` | Read device information | Implants, medical devices |
| `system/DeviceUseStatement.read` | Read device usage | Device tracking, implant registry |

### Clinical Assessment

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/ClinicalImpression.read` | Read clinical assessments | Clinical summaries, assessments |
| `system/ClinicalImpression.write` | Document assessments | Clinical documentation, assessment notes |
| `system/RiskAssessment.read` | Read risk assessments | Risk scoring, clinical decision support |
| `system/RiskAssessment.write` | Document risk assessments | Risk stratification, screening |

### Family History

| Scope | Description | Use Cases |
|-------|-------------|-----------|
| `system/FamilyMemberHistory.read` | Read family history | Genetic risk, family medical history |
| `system/FamilyMemberHistory.write` | Document family history | Recording family history, genetic counseling |

---

## Clinical Notes & Document Access Guide

### Understanding Document-Related Resources

EPIC uses several FHIR resources for clinical documentation:

**DocumentReference** - Metadata about documents
- Points to the actual document content
- Contains document type, date, author, patient
- Links to Binary resource for content

**Binary** - Raw document content
- Actual PDF, image, or file bytes
- Referenced by DocumentReference
- Content-Type header indicates file type

**Composition** - Structured clinical documents
- C-CDA documents
- Clinical notes in FHIR format
- Structured sections with narrative text

**Media** - Photos and multimedia
- Wound photos
- X-rays, scans
- Video or audio recordings

**QuestionnaireResponse** - Structured forms
- Patient intake forms
- Clinical assessments
- PRO (Patient-Reported Outcomes)

### Example: Reading Clinical Notes

**Step 1: Search for DocumentReference**
```
Scope needed: system/DocumentReference.read

GET /DocumentReference?patient=12345&category=clinical-note
```

**Step 2: Get Document Content**
```
Scope needed: system/Binary.read

GET /Binary/doc-content-id
```

**Step 3: Read Structured Document**
```
Scope needed: system/Composition.read

GET /Composition/note-composition-id
```

### Common Document Types in EPIC

| Document Type | FHIR Resource | Scope Required |
|--------------|---------------|----------------|
| Progress Notes | DocumentReference + Binary | DocumentReference.read + Binary.read |
| Discharge Summary | Composition | Composition.read |
| Lab Report | DiagnosticReport | DiagnosticReport.read |
| Imaging Report | DiagnosticReport + Binary | DiagnosticReport.read + Binary.read |
| Scanned Document | DocumentReference + Binary | DocumentReference.read + Binary.read |
| Patient Photos | Media | Media.read |
| C-CDA Document | Composition | Composition.read |
| Questionnaire | QuestionnaireResponse | QuestionnaireResponse.read |

### Scopes for Complete Clinical Note Access

```yaml
scopes:
  # Minimum for reading clinical notes
  - "system/DocumentReference.read"
  - "system/Binary.read"

  # Add for structured documents
  - "system/Composition.read"

  # Add for diagnostic reports
  - "system/DiagnosticReport.read"

  # Add for images/media
  - "system/Media.read"

  # Add for assessments
  - "system/QuestionnaireResponse.read"

  # Add for audit trails
  - "system/Provenance.read"
```

### Writing Clinical Notes

To create or update clinical notes:

```yaml
scopes:
  - "system/DocumentReference.write"  # Create note metadata
  - "system/Composition.write"        # Create structured notes
  - "system/Media.write"              # Upload images
  - "system/QuestionnaireResponse.write"  # Submit assessments
```

**Important:** Binary resources are typically created automatically when you upload documents through DocumentReference, so `Binary.write` is usually not needed.

---

## Registering Scopes in EPIC App Orchard

### Step 1: Login to App Orchard

1. Go to https://apporchard.epic.com/
2. Navigate to your application
3. Click on "Authentication" or "OAuth2"

### Step 2: Select Scopes

1. Choose "Backend Services (System-Level)"
2. Select the resource types you need
3. For each resource, choose:
   - **Read** - View data only
   - **Write** - Create and update data (includes read)

### Step 3: Justification

EPIC requires a business justification for each scope:

**Examples:**
- **Patient.read/write**: "Application needs to access and update patient demographics for care coordination"
- **Observation.read**: "Application retrieves lab results and vital signs for clinical decision support"
- **MedicationRequest.read/write**: "Application manages medication orders for prescription workflow"

### Step 4: Save and Test

1. Save your scope selections
2. Note that it may take a few minutes for changes to take effect
3. Test with the token test script:
   ```bash
   go run scripts/test_epic_token.go
   ```

---

## Scope Selection Best Practices

### Start Minimal

Begin with only the scopes you absolutely need:
```yaml
scopes:
  - "system/Patient.read"
  - "system/Observation.read"
  - "system/Condition.read"
```

### Expand as Needed

Add scopes incrementally as features are implemented:
```yaml
scopes:
  # Phase 1: Read-only patient data
  - "system/Patient.read"
  - "system/Observation.read"

  # Phase 2: Add write capabilities
  - "system/Observation.write"
  - "system/Condition.write"

  # Phase 3: Add medications
  - "system/MedicationRequest.read"
  - "system/MedicationRequest.write"
```

### Production Checklist

Before deploying to production:

- [ ] Only request scopes your application actually uses
- [ ] Provide clear justification for each scope
- [ ] Test all scopes in sandbox environment
- [ ] Document which features require which scopes
- [ ] Plan for scope changes (requires EPIC re-approval)
- [ ] Monitor scope usage in production

---

## Common Scope Combinations

### Patient Summary Application

```yaml
scopes:
  - "system/Patient.read"
  - "system/Observation.read"
  - "system/Condition.read"
  - "system/MedicationRequest.read"
  - "system/AllergyIntolerance.read"
  - "system/Immunization.read"
  - "system/Encounter.read"
```

### Care Coordination Platform

```yaml
scopes:
  - "system/Patient.read"
  - "system/CarePlan.read"
  - "system/CarePlan.write"
  - "system/Goal.read"
  - "system/Goal.write"
  - "system/Appointment.read"
  - "system/Communication.read"
  - "system/Communication.write"
```

### Clinical Decision Support

```yaml
scopes:
  - "system/Patient.read"
  - "system/Observation.read"
  - "system/Condition.read"
  - "system/MedicationRequest.read"
  - "system/AllergyIntolerance.read"
  - "system/RiskAssessment.read"
  - "system/RiskAssessment.write"
```

### Medication Management

```yaml
scopes:
  - "system/Patient.read"
  - "system/MedicationRequest.read"
  - "system/MedicationRequest.write"
  - "system/MedicationStatement.read"
  - "system/MedicationStatement.write"
  - "system/AllergyIntolerance.read"
```

### Clinical Documentation & Notes Access

```yaml
scopes:
  - "system/Patient.read"
  - "system/DocumentReference.read"     # Note metadata
  - "system/Binary.read"                # Note content (PDFs, etc.)
  - "system/Composition.read"           # Structured clinical documents
  - "system/DiagnosticReport.read"      # Lab/imaging reports
  - "system/Media.read"                 # Photos, images
  - "system/QuestionnaireResponse.read" # Clinical assessments
  - "system/Provenance.read"            # Audit trails
```

### Document Upload & Management

```yaml
scopes:
  - "system/Patient.read"
  - "system/DocumentReference.read"
  - "system/DocumentReference.write"    # Upload documents
  - "system/Binary.read"                # Read document content
  - "system/Composition.write"          # Create structured notes
  - "system/Media.write"                # Upload photos/media
  - "system/QuestionnaireResponse.write" # Submit forms
```

---

## Troubleshooting Scope Issues

### Error: "Invalid scope"

**Cause**: Requesting a scope not registered in EPIC App Orchard

**Fix**:
1. Check EPIC App Orchard settings
2. Ensure scope is exactly as registered
3. Wait a few minutes after making changes in App Orchard

### Error: "Insufficient scope"

**Cause**: Token doesn't have permission for the requested operation

**Fix**:
1. Verify scope was requested during authentication
2. Check token includes the scope: decode JWT at https://jwt.io
3. Refresh token after updating scopes in App Orchard

### Error: "Access denied"

**Cause**: Scope is registered but access is restricted

**Fix**:
1. Verify your client app has proper approvals in EPIC
2. Check if scope requires additional Epic configuration
3. Contact EPIC support for clarification

---

## Additional Resources

- [EPIC FHIR Documentation](https://fhir.epic.com/)
- [EPIC App Orchard](https://apporchard.epic.com/)
- [SMART Backend Services](http://hl7.org/fhir/smart-app-launch/backend-services.html)
- [Local Testing Guide](LOCAL-TESTING.md)
- [EPIC Security Setup](EPIC-APP-SECURITY.md)
