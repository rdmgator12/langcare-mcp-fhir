package providers

import (
	"context"
	"encoding/json"
	"log/slog"

	"golang.org/x/oauth2/clientcredentials"
)

// CernerConfig holds Cerner FHIR configuration
type CernerConfig struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	TokenURL     string
	Scopes       []string
}

// CernerProvider implements FHIR access to Cerner with OAuth2 Client Credentials
type CernerProvider struct {
	*BaseProvider
	clientID     string
	clientSecret string
	tokenURL     string
	scopes       []string
}

// NewCernerProvider creates a new Cerner FHIR provider
func NewCernerProvider(cfg *CernerConfig, logger *slog.Logger) (*CernerProvider, error) {
	logger.Info("initializing cerner provider",
		"base_url", cfg.BaseURL,
		"client_id", cfg.ClientID,
	)

	// Configure OAuth2 client credentials flow
	oauthConfig := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     cfg.TokenURL,
		Scopes:       cfg.Scopes,
	}

	// Create HTTP client with auto-refreshing OAuth2 tokens
	// This client automatically handles token refresh before expiry
	httpClient := oauthConfig.Client(context.Background())

	provider := &CernerProvider{
		BaseProvider: NewBaseProvider(cfg.BaseURL, httpClient, logger),
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		tokenURL:     cfg.TokenURL,
		scopes:       cfg.Scopes,
	}

	logger.Info("cerner provider initialized successfully")

	return provider, nil
}

// GetProviderType returns the provider type
func (c *CernerProvider) GetProviderType() string {
	return "cerner"
}

// Read retrieves a single FHIR resource
func (c *CernerProvider) Read(ctx context.Context, resourceType, id string) (json.RawMessage, error) {
	return c.BaseProvider.Read(ctx, resourceType, id)
}

// Search performs a FHIR search operation
func (c *CernerProvider) Search(ctx context.Context, resourceType, queryParams string) (json.RawMessage, error) {
	return c.BaseProvider.Search(ctx, resourceType, queryParams)
}

// Create creates a new FHIR resource
func (c *CernerProvider) Create(ctx context.Context, resourceType string, resource json.RawMessage) (json.RawMessage, error) {
	return c.BaseProvider.Create(ctx, resourceType, resource)
}

// Update updates an existing FHIR resource
func (c *CernerProvider) Update(ctx context.Context, resourceType, id string, resource json.RawMessage) (json.RawMessage, error) {
	return c.BaseProvider.Update(ctx, resourceType, id, resource)
}
