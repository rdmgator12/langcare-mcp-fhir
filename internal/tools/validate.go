package tools

import (
	"fmt"
	"net/url"
	"strings"
)

var validResourceTypes = map[string]bool{
	"Account": true, "ActivityDefinition": true, "AdverseEvent": true,
	"AllergyIntolerance": true, "Appointment": true, "AppointmentResponse": true,
	"AuditEvent": true, "Basic": true, "Binary": true, "BiologicallyDerivedProduct": true,
	"BodyStructure": true, "Bundle": true, "CapabilityStatement": true,
	"CarePlan": true, "CareTeam": true, "ChargeItem": true,
	"Claim": true, "ClaimResponse": true, "ClinicalImpression": true,
	"CodeSystem": true, "Communication": true, "CommunicationRequest": true,
	"Composition": true, "ConceptMap": true, "Condition": true, "Consent": true,
	"Contract": true, "Coverage": true, "CoverageEligibilityRequest": true,
	"CoverageEligibilityResponse": true, "DetectedIssue": true, "Device": true,
	"DeviceDefinition": true, "DeviceMetric": true, "DeviceRequest": true,
	"DeviceUseStatement": true, "DiagnosticReport": true, "DocumentManifest": true,
	"DocumentReference": true, "Encounter": true, "Endpoint": true,
	"EnrollmentRequest": true, "EnrollmentResponse": true,
	"EpisodeOfCare": true, "ExplanationOfBenefit": true,
	"FamilyMemberHistory": true, "Flag": true, "Goal": true, "Group": true,
	"GuidanceResponse": true, "HealthcareService": true,
	"ImagingStudy": true, "Immunization": true, "ImmunizationEvaluation": true,
	"ImmunizationRecommendation": true, "InsurancePlan": true,
	"Library": true, "List": true, "Location": true,
	"Measure": true, "MeasureReport": true, "Media": true, "Medication": true,
	"MedicationAdministration": true, "MedicationDispense": true,
	"MedicationKnowledge": true, "MedicationRequest": true,
	"MedicationStatement": true, "MessageDefinition": true, "MessageHeader": true,
	"NamingSystem": true, "NutritionOrder": true, "Observation": true,
	"OperationDefinition": true, "OperationOutcome": true, "Organization": true,
	"OrganizationAffiliation": true, "Patient": true, "PaymentNotice": true,
	"PaymentReconciliation": true, "Person": true, "PlanDefinition": true,
	"Practitioner": true, "PractitionerRole": true, "Procedure": true,
	"Provenance": true, "Questionnaire": true, "QuestionnaireResponse": true,
	"RelatedPerson": true, "RequestGroup": true, "ResearchStudy": true,
	"ResearchSubject": true, "RiskAssessment": true, "Schedule": true,
	"SearchParameter": true, "ServiceRequest": true, "Slot": true,
	"Specimen": true, "StructureDefinition": true, "Subscription": true,
	"Substance": true, "SupplyDelivery": true, "SupplyRequest": true,
	"Task": true, "ValueSet": true, "VisionPrescription": true,
}

// Patient compartment: resources that MUST include a patient scope in searches.
// Per FHIR R4 patient compartment definition.
var patientCompartmentResources = map[string]bool{
	"AllergyIntolerance": true, "Appointment": true, "CarePlan": true,
	"CareTeam": true, "ClinicalImpression": true, "Communication": true,
	"Composition": true, "Condition": true, "Consent": true, "Coverage": true,
	"DetectedIssue": true, "DeviceRequest": true, "DeviceUseStatement": true,
	"DiagnosticReport": true, "DocumentReference": true, "Encounter": true,
	"EpisodeOfCare": true, "ExplanationOfBenefit": true,
	"FamilyMemberHistory": true, "Flag": true, "Goal": true,
	"ImagingStudy": true, "Immunization": true, "ImmunizationRecommendation": true,
	"MedicationAdministration": true, "MedicationDispense": true,
	"MedicationRequest": true, "MedicationStatement": true,
	"NutritionOrder": true, "Observation": true, "Procedure": true,
	"Provenance": true, "QuestionnaireResponse": true, "RelatedPerson": true,
	"RiskAssessment": true, "ServiceRequest": true, "Specimen": true,
	"SupplyDelivery": true, "SupplyRequest": true, "VisionPrescription": true,
}

func ValidateResourceType(rt string) error {
	if rt == "" {
		return fmt.Errorf("resourceType is required")
	}
	if !validResourceTypes[rt] {
		return fmt.Errorf("invalid FHIR resource type: %q", rt)
	}
	return nil
}

func ValidateResourceID(id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	if strings.ContainsAny(id, "/?#&=") {
		return fmt.Errorf("resource id contains illegal characters")
	}
	if strings.Contains(id, "..") {
		return fmt.Errorf("resource id contains path traversal sequence")
	}
	if len(id) > 64 {
		return fmt.Errorf("resource id exceeds maximum length")
	}
	return nil
}

func IsPatientCompartment(rt string) bool {
	return patientCompartmentResources[rt]
}

// SanitizeQueryParams parses a raw query string and re-encodes it safely.
// Returns the sanitized query string (without leading '?').
func SanitizeQueryParams(raw string) (string, error) {
	if raw == "" {
		return "", nil
	}
	values, err := url.ParseQuery(raw)
	if err != nil {
		return "", fmt.Errorf("invalid query parameters: %w", err)
	}
	return values.Encode(), nil
}

// EnforcePatientScope ensures patient-compartment searches include a patient param.
// Returns an error if the resource type requires patient scoping but none is present.
func EnforcePatientScope(resourceType, queryParams string) error {
	if !IsPatientCompartment(resourceType) {
		return nil
	}
	values, err := url.ParseQuery(queryParams)
	if err != nil {
		return fmt.Errorf("invalid query parameters: %w", err)
	}
	if values.Get("patient") == "" && values.Get("subject") == "" {
		return fmt.Errorf("%s is a patient-compartment resource: 'patient' or 'subject' parameter required", resourceType)
	}
	return nil
}
