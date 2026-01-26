package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GCPConfig holds GCP Healthcare API configuration
type GCPConfig struct {
	ProjectID     string
	Location      string
	DatasetID     string
	FHIRStoreID   string
	CredentialPath string // Path to service account JSON key
}

// GCPProvider implements FHIR access to Google Cloud Healthcare API
type GCPProvider struct {
	*BaseProvider
	projectID   string
	location    string
	datasetID   string
	fhirStoreID string
}

// NewGCPProvider creates a new GCP Healthcare API provider
func NewGCPProvider(cfg *GCPConfig, logger *slog.Logger) (*GCPProvider, error) {
	// Build FHIR store URL
	baseURL := fmt.Sprintf(
		"https://healthcare.googleapis.com/v1/projects/%s/locations/%s/datasets/%s/fhirStores/%s/fhir",
		cfg.ProjectID, cfg.Location, cfg.DatasetID, cfg.FHIRStoreID,
	)

	logger.Info("initializing gcp provider",
		"project", cfg.ProjectID,
		"location", cfg.Location,
		"dataset", cfg.DatasetID,
		"fhir_store", cfg.FHIRStoreID,
	)

	// Create OAuth2 token source from service account
	ctx := context.Background()
	var tokenSource oauth2.TokenSource

	if cfg.CredentialPath != "" {
		// Use explicit service account key file
		jsonData, err := os.ReadFile(cfg.CredentialPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read credential file: %w", err)
		}

		creds, credErr := google.CredentialsFromJSON(ctx, jsonData,
			"https://www.googleapis.com/auth/cloud-healthcare",
		)
		if credErr != nil {
			return nil, fmt.Errorf("failed to load credentials from file: %w", credErr)
		}
		tokenSource = creds.TokenSource
	} else {
		// Use Application Default Credentials (ADC)
		// This supports: GOOGLE_APPLICATION_CREDENTIALS env var, GCE metadata, etc.
		creds, credErr := google.FindDefaultCredentials(ctx,
			"https://www.googleapis.com/auth/cloud-healthcare",
		)
		if credErr != nil {
			return nil, fmt.Errorf("failed to get default credentials: %w", credErr)
		}
		tokenSource = creds.TokenSource
	}

	// Create HTTP client with auto-refreshing OAuth2 tokens
	httpClient := oauth2.NewClient(ctx, tokenSource)

	return &GCPProvider{
		BaseProvider: NewBaseProvider(baseURL, httpClient, logger),
		projectID:    cfg.ProjectID,
		location:     cfg.Location,
		datasetID:    cfg.DatasetID,
		fhirStoreID:  cfg.FHIRStoreID,
	}, nil
}

// GetProviderType returns the provider type
func (g *GCPProvider) GetProviderType() string {
	return "gcp"
}

// Read retrieves a single FHIR resource
func (g *GCPProvider) Read(ctx context.Context, resourceType, id string) (json.RawMessage, error) {
	return g.BaseProvider.Read(ctx, resourceType, id)
}

// Search performs a FHIR search operation
func (g *GCPProvider) Search(ctx context.Context, resourceType, queryParams string) (json.RawMessage, error) {
	return g.BaseProvider.Search(ctx, resourceType, queryParams)
}

// Create creates a new FHIR resource
func (g *GCPProvider) Create(ctx context.Context, resourceType string, resource json.RawMessage) (json.RawMessage, error) {
	return g.BaseProvider.Create(ctx, resourceType, resource)
}

// Update updates an existing FHIR resource
func (g *GCPProvider) Update(ctx context.Context, resourceType, id string, resource json.RawMessage) (json.RawMessage, error) {
	return g.BaseProvider.Update(ctx, resourceType, id, resource)
}
