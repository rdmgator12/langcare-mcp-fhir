package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

const maxResponseBytes = 100 * 1024 * 1024 // 100 MB

// Provider defines the interface for FHIR API providers
type Provider interface {
	// Read retrieves a single FHIR resource by type and ID
	Read(ctx context.Context, resourceType, id string) (json.RawMessage, error)

	// Search performs a FHIR search operation
	Search(ctx context.Context, resourceType, queryParams string) (json.RawMessage, error)

	// Create creates a new FHIR resource
	Create(ctx context.Context, resourceType string, resource json.RawMessage) (json.RawMessage, error)

	// Update updates an existing FHIR resource
	Update(ctx context.Context, resourceType, id string, resource json.RawMessage) (json.RawMessage, error)

	// GetProviderType returns the provider type (gcp, epic, cerner, generic)
	GetProviderType() string
}

// BaseProvider provides common HTTP functionality for all providers
type BaseProvider struct {
	baseURL    string
	httpClient *http.Client
	logger     *slog.Logger
}

// NewBaseProvider creates a new base provider
func NewBaseProvider(baseURL string, httpClient *http.Client, logger *slog.Logger) *BaseProvider {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &BaseProvider{
		baseURL:    baseURL,
		httpClient: httpClient,
		logger:     logger,
	}
}

// Read implements generic FHIR read operation
func (b *BaseProvider) Read(ctx context.Context, resourceType, id string) (json.RawMessage, error) {
	u := fmt.Sprintf("%s/%s/%s", b.baseURL, url.PathEscape(resourceType), url.PathEscape(id))
	return b.doRequest(ctx, "GET", u, nil)
}

// Search implements generic FHIR search operation
func (b *BaseProvider) Search(ctx context.Context, resourceType, queryParams string) (json.RawMessage, error) {
	u := fmt.Sprintf("%s/%s", b.baseURL, url.PathEscape(resourceType))
	if queryParams != "" {
		u = fmt.Sprintf("%s?%s", u, queryParams)
	}
	return b.doRequest(ctx, "GET", u, nil)
}

// Create implements generic FHIR create operation
func (b *BaseProvider) Create(ctx context.Context, resourceType string, resource json.RawMessage) (json.RawMessage, error) {
	u := fmt.Sprintf("%s/%s", b.baseURL, url.PathEscape(resourceType))
	return b.doRequest(ctx, "POST", u, resource)
}

// Update implements generic FHIR update operation
func (b *BaseProvider) Update(ctx context.Context, resourceType, id string, resource json.RawMessage) (json.RawMessage, error) {
	u := fmt.Sprintf("%s/%s/%s", b.baseURL, url.PathEscape(resourceType), url.PathEscape(id))
	return b.doRequest(ctx, "PUT", u, resource)
}

// doRequest performs an HTTP request with proper FHIR headers
func (b *BaseProvider) doRequest(ctx context.Context, method, url string, body json.RawMessage) (json.RawMessage, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader([]byte(body))
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set FHIR headers
	req.Header.Set("Accept", "application/fhir+json")
	if body != nil {
		req.Header.Set("Content-Type", "application/fhir+json")
	}

	b.logger.Debug("fhir request",
		"method", method,
		"url", url,
	)

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("FHIR request failed with status %d (body redacted for PHI safety)", resp.StatusCode)
	}

	b.logger.Debug("fhir response",
		"status", resp.StatusCode,
		"size", len(respBody),
	)

	return json.RawMessage(respBody), nil
}
