package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Configuration
	clientID := os.Getenv("EPIC_CLIENT_ID")
	privateKeyPath := os.Getenv("EPIC_PRIVATE_KEY_PATH")
	tokenURL := "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/token"

	// Comprehensive FHIR scopes - matches config.local.epic.yaml
	scopes := []string{
		// Patient Demographics
		"system/Patient.read",
		"system/Patient.write",

		// Clinical Observations (Labs, Vitals, etc.)
		"system/Observation.read",
		"system/Observation.write",

		// Conditions & Diagnoses
		"system/Condition.read",
		"system/Condition.write",

		// Medications
		"system/MedicationRequest.read",
		"system/MedicationRequest.write",
		"system/MedicationStatement.read",
		"system/MedicationStatement.write",
		"system/Medication.read",

		// Encounters & Visits
		"system/Encounter.read",
		"system/Encounter.write",

		// Procedures
		"system/Procedure.read",
		"system/Procedure.write",

		// Diagnostic Reports
		"system/DiagnosticReport.read",
		"system/DiagnosticReport.write",

		// Allergies & Intolerances
		"system/AllergyIntolerance.read",
		"system/AllergyIntolerance.write",

		// Immunizations
		"system/Immunization.read",
		"system/Immunization.write",

		// Clinical Documents & Notes
		"system/DocumentReference.read",
		"system/DocumentReference.write",
		"system/Binary.read",
		"system/Composition.read",
		"system/Composition.write",
		"system/Media.read",
		"system/Media.write",
		"system/QuestionnaireResponse.read",
		"system/QuestionnaireResponse.write",
		"system/Provenance.read",

		// Care Plans & Goals
		"system/CarePlan.read",
		"system/CarePlan.write",
		"system/Goal.read",
		"system/Goal.write",

		// Scheduling
		"system/Appointment.read",
		"system/Appointment.write",
		"system/Schedule.read",
		"system/Slot.read",

		// Coverage & Claims
		"system/Coverage.read",
		"system/Claim.read",
		"system/Claim.write",
		"system/ExplanationOfBenefit.read",

		// Providers & Organizations
		"system/Practitioner.read",
		"system/PractitionerRole.read",
		"system/Organization.read",
		"system/Location.read",

		// Family History
		"system/FamilyMemberHistory.read",
		"system/FamilyMemberHistory.write",

		// Specimens & Lab Orders
		"system/Specimen.read",
		"system/ServiceRequest.read",
		"system/ServiceRequest.write",

		// Device & DeviceUseStatement
		"system/Device.read",
		"system/DeviceUseStatement.read",

		// Additional Clinical Resources
		"system/ClinicalImpression.read",
		"system/ClinicalImpression.write",
		"system/RiskAssessment.read",
		"system/RiskAssessment.write",
		"system/Communication.read",
		"system/Communication.write",
		"system/CommunicationRequest.read",
		"system/CommunicationRequest.write",
	}

	// Allow command-line override
	if len(os.Args) > 1 {
		clientID = os.Args[1]
	}
	if len(os.Args) > 2 {
		privateKeyPath = os.Args[2]
	}

	// Validate inputs
	if clientID == "" {
		fmt.Println("Error: EPIC_CLIENT_ID not set")
		fmt.Println("\n=== EPIC Token Test - Comprehensive FHIR Scopes ===")
		fmt.Println("\nThis test requests 60+ OAuth2 scopes for comprehensive FHIR access.")
		fmt.Println("It will show which scopes were granted vs. requested.")
		fmt.Println("\nUsage:")
		fmt.Println("  export EPIC_CLIENT_ID='your-client-id'")
		fmt.Println("  export EPIC_PRIVATE_KEY_PATH='./keys/private-key.pem'")
		fmt.Println("  go run test/test_epic_token.go")
		fmt.Println("\nOr:")
		fmt.Println("  go run test/test_epic_token.go <client-id> <private-key-path>")
		fmt.Println("\nNote: Scopes must be registered in EPIC App Orchard first!")
		os.Exit(1)
	}

	if privateKeyPath == "" {
		fmt.Println("Error: EPIC_PRIVATE_KEY_PATH not set")
		os.Exit(1)
	}

	fmt.Println("=== EPIC Token Test ===")
	fmt.Printf("Client ID: %s\n", clientID)
	fmt.Printf("Private Key: %s\n", privateKeyPath)
	fmt.Printf("Token URL: %s\n", tokenURL)
	fmt.Printf("Requesting %d scopes (comprehensive FHIR access)\n", len(scopes))
	fmt.Printf("Scopes: %s\n\n", strings.Join(scopes, ", "))

	// Step 1: Load private key
	fmt.Println("Step 1: Loading private key...")
	privateKey, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		fmt.Printf("❌ Failed to load private key: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ Private key loaded successfully")

	// Step 2: Create JWT assertion
	fmt.Println("\nStep 2: Creating JWT assertion...")
	now := time.Now()
	claims := jwt.MapClaims{
		"iss": clientID,
		"sub": clientID,
		"aud": tokenURL,
		"exp": now.Add(5 * time.Minute).Unix(),
		"iat": now.Unix(),
		"jti": generateJTI(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
	assertion, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Printf("❌ Failed to sign JWT: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ JWT assertion created and signed")
	fmt.Printf("   JWT (first 50 chars): %s...\n", assertion[:50])

	// Step 3: Exchange JWT for access token
	fmt.Println("\nStep 3: Exchanging JWT for access token...")
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", assertion)
	data.Set("scope", strings.Join(scopes, " "))

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Printf("❌ Failed to create request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Token request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("❌ Token request failed with status %d\n", resp.StatusCode)
		fmt.Printf("Response: %s\n", string(body))
		os.Exit(1)
	}

	// Step 4: Parse token response
	fmt.Println("\nStep 4: Parsing token response...")
	var tokenResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	if err := json.Unmarshal(body, &tokenResp); err != nil {
		fmt.Printf("❌ Failed to parse token response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Access token obtained successfully!")
	fmt.Printf("\n=== TOKEN DETAILS ===\n")
	fmt.Printf("Token Type: %s\n", tokenResp.TokenType)
	fmt.Printf("Expires In: %d seconds (%d minutes)\n", tokenResp.ExpiresIn, tokenResp.ExpiresIn/60)

	// Analyze granted scopes
	grantedScopes := strings.Fields(tokenResp.Scope)
	fmt.Printf("\nRequested Scopes: %d\n", len(scopes))
	fmt.Printf("Granted Scopes: %d\n", len(grantedScopes))

	if len(grantedScopes) < len(scopes) {
		fmt.Printf("\n⚠️  WARNING: Not all scopes were granted!\n")
		fmt.Printf("   You may need to register additional scopes in EPIC App Orchard.\n\n")

		// Find missing scopes
		grantedMap := make(map[string]bool)
		for _, s := range grantedScopes {
			grantedMap[s] = true
		}

		missingScopes := []string{}
		for _, s := range scopes {
			if !grantedMap[s] {
				missingScopes = append(missingScopes, s)
			}
		}

		if len(missingScopes) > 0 {
			fmt.Printf("Missing Scopes (%d):\n", len(missingScopes))
			for _, s := range missingScopes {
				fmt.Printf("  - %s\n", s)
			}
			fmt.Println()
		}
	} else {
		fmt.Printf("✅ All requested scopes were granted!\n\n")
	}

	fmt.Printf("Granted Scopes:\n%s\n", tokenResp.Scope)
	fmt.Printf("\nAccess Token (first 50 chars): %s...\n", tokenResp.AccessToken[:50])
	fmt.Printf("\n=== FULL ACCESS TOKEN ===\n")
	fmt.Println(tokenResp.AccessToken)

	// Step 5: Test token with a simple FHIR request
	fmt.Println("\n=== Testing Token with FHIR API ===")
	fhirURL := "https://fhir.epic.com/interconnect-fhir-oauth/api/FHIR/R4/metadata"
	fmt.Printf("Calling: %s\n", fhirURL)

	fhirReq, err := http.NewRequest("GET", fhirURL, nil)
	if err != nil {
		fmt.Printf("❌ Failed to create FHIR request: %v\n", err)
		os.Exit(1)
	}
	fhirReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenResp.AccessToken))
	fhirReq.Header.Set("Accept", "application/fhir+json")

	fhirResp, err := client.Do(fhirReq)
	if err != nil {
		fmt.Printf("❌ FHIR request failed: %v\n", err)
		os.Exit(1)
	}
	defer fhirResp.Body.Close()

	fmt.Printf("Response Status: %d %s\n", fhirResp.StatusCode, fhirResp.Status)

	if fhirResp.StatusCode == 200 {
		fmt.Println("✅ Token works! Successfully called FHIR API")

		var metadata map[string]interface{}
		if err := json.NewDecoder(fhirResp.Body).Decode(&metadata); err == nil {
			if rt, ok := metadata["resourceType"].(string); ok {
				fmt.Printf("   Resource Type: %s\n", rt)
			}
			if version, ok := metadata["fhirVersion"].(string); ok {
				fmt.Printf("   FHIR Version: %s\n", version)
			}
		}
	} else {
		body, _ := io.ReadAll(fhirResp.Body)
		fmt.Printf("❌ FHIR request failed\n")
		fmt.Printf("Response: %s\n", string(body))
	}

	fmt.Println("\n=== TEST COMPLETE ===")
}

// loadPrivateKey loads an RSA private key from PEM file
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// Try PKCS1 format first
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try PKCS8 format
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}
		var ok bool
		privateKey, ok = key.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("key is not RSA private key")
		}
	}

	return privateKey, nil
}

// generateJTI generates a unique JWT ID
func generateJTI() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
