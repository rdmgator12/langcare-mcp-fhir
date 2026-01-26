package providers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// GenericConfig holds configuration for generic FHIR server with simple auth
type GenericConfig struct {
	BaseURL  string
	AuthType string // "bearer", "basic", "none"
	Token    string // For bearer auth
	Username string // For basic auth
	Password string // For basic auth
}

// GenericProvider implements FHIR access with simple authentication
// This is for backwards compatibility and simple FHIR servers
type GenericProvider struct {
	*BaseProvider
	authType string
	token    string
	username string
	password string
}

// NewGenericProvider creates a new generic FHIR provider
func NewGenericProvider(cfg *GenericConfig, logger *slog.Logger) (*GenericProvider, error) {
	logger.Info("initializing generic provider",
		"base_url", cfg.BaseURL,
		"auth_type", cfg.AuthType,
	)

	provider := &GenericProvider{
		authType: cfg.AuthType,
		token:    cfg.Token,
		username: cfg.Username,
		password: cfg.Password,
	}

	// Create HTTP client with auth injection
	httpClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: &genericTransport{provider: provider},
	}

	provider.BaseProvider = NewBaseProvider(cfg.BaseURL, httpClient, logger)

	logger.Info("generic provider initialized successfully")

	return provider, nil
}

// genericTransport adds authentication to requests
type genericTransport struct {
	provider *GenericProvider
}

func (t *genericTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	switch t.provider.authType {
	case "bearer":
		if t.provider.token != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.provider.token))
		}
	case "basic":
		if t.provider.username != "" && t.provider.password != "" {
			auth := base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%s:%s", t.provider.username, t.provider.password)),
			)
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))
		}
	case "none":
		// No authentication
	default:
		return nil, fmt.Errorf("unknown auth type: %s", t.provider.authType)
	}

	return http.DefaultTransport.RoundTrip(req)
}

// GetProviderType returns the provider type
func (g *GenericProvider) GetProviderType() string {
	return "generic"
}

// Read retrieves a single FHIR resource
func (g *GenericProvider) Read(ctx context.Context, resourceType, id string) (json.RawMessage, error) {
	return g.BaseProvider.Read(ctx, resourceType, id)
}

// Search performs a FHIR search operation
func (g *GenericProvider) Search(ctx context.Context, resourceType, queryParams string) (json.RawMessage, error) {
	return g.BaseProvider.Search(ctx, resourceType, queryParams)
}

// Create creates a new FHIR resource
func (g *GenericProvider) Create(ctx context.Context, resourceType string, resource json.RawMessage) (json.RawMessage, error) {
	return g.BaseProvider.Create(ctx, resourceType, resource)
}

// Update updates an existing FHIR resource
func (g *GenericProvider) Update(ctx context.Context, resourceType, id string, resource json.RawMessage) (json.RawMessage, error) {
	return g.BaseProvider.Update(ctx, resourceType, id, resource)
}
