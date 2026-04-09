package providers

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// OpenEMRConfig holds OpenEMR FHIR configuration
type OpenEMRConfig struct {
	BaseURL        string
	ClientID       string
	PrivateKeyPath string
	TokenURL       string
	KeyID          string
	Scopes         []string
}

// OpenEMRProvider implements FHIR access to OpenEMR with SMART on FHIR
// Backend Services Auth (private_key_jwt + RS384).
type OpenEMRProvider struct {
	*BaseProvider
	clientID       string
	privateKey     *rsa.PrivateKey
	tokenURL       string
	keyID          string
	scopes         []string
	tokenCache     *tokenCache
	tokenMu        sync.Mutex
	httpClientBase *http.Client
	logger         *slog.Logger
}

// NewOpenEMRProvider creates a new OpenEMR FHIR provider
func NewOpenEMRProvider(cfg *OpenEMRConfig, logger *slog.Logger) (*OpenEMRProvider, error) {
	logger.Info("initializing openemr provider",
		"base_url", cfg.BaseURL,
		"client_id", cfg.ClientID,
	)

	privateKey, err := loadPrivateKey(cfg.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	baseHTTPClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	provider := &OpenEMRProvider{
		clientID:       cfg.ClientID,
		privateKey:     privateKey,
		tokenURL:       cfg.TokenURL,
		keyID:          cfg.KeyID,
		scopes:         cfg.Scopes,
		tokenCache:     &tokenCache{},
		httpClientBase: baseHTTPClient,
		logger:         logger,
	}

	httpClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: &openemrTransport{provider: provider},
	}

	provider.BaseProvider = NewBaseProvider(cfg.BaseURL, httpClient, logger)

	// Pre-authenticate to validate credentials
	if _, err := provider.getValidToken(context.Background()); err != nil {
		return nil, fmt.Errorf("initial authentication failed: %w", err)
	}

	logger.Info("openemr provider authenticated successfully")

	return provider, nil
}

// openemrTransport adds OAuth2 bearer token to requests
type openemrTransport struct {
	provider *OpenEMRProvider
}

func (t *openemrTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.provider.getValidToken(req.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return http.DefaultTransport.RoundTrip(req)
}

// getValidToken returns a valid access token, refreshing if necessary.
// OpenEMR access tokens are very short-lived (60 seconds) with no refresh
// tokens, so we refresh whenever fewer than 10 seconds remain.
func (o *OpenEMRProvider) getValidToken(ctx context.Context) (string, error) {
	o.tokenCache.mu.RLock()
	token := o.tokenCache.accessToken
	expiry := o.tokenCache.expiry
	o.tokenCache.mu.RUnlock()

	if token != "" && time.Until(expiry) > 10*time.Second {
		return token, nil
	}

	// Serialize refreshes so concurrent callers don't all hit the token endpoint.
	o.tokenMu.Lock()
	defer o.tokenMu.Unlock()

	o.tokenCache.mu.RLock()
	token = o.tokenCache.accessToken
	expiry = o.tokenCache.expiry
	o.tokenCache.mu.RUnlock()
	if token != "" && time.Until(expiry) > 10*time.Second {
		return token, nil
	}

	if err := o.refreshToken(ctx); err != nil {
		return "", err
	}

	o.tokenCache.mu.RLock()
	defer o.tokenCache.mu.RUnlock()
	return o.tokenCache.accessToken, nil
}

// refreshToken obtains a new access token using JWT assertion
func (o *OpenEMRProvider) refreshToken(ctx context.Context) error {
	o.logger.Debug("refreshing openemr access token")

	now := time.Now()
	claims := jwt.MapClaims{
		"iss": o.clientID,
		"sub": o.clientID,
		"aud": o.tokenURL,
		"exp": now.Add(5 * time.Minute).Unix(),
		"iat": now.Unix(),
		"jti": generateJTI(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
	if o.keyID != "" {
		token.Header["kid"] = o.keyID
	}
	token.Header["typ"] = "JWT"

	assertion, err := token.SignedString(o.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign JWT: %w", err)
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", assertion)
	if len(o.scopes) > 0 {
		data.Set("scope", strings.Join(o.scopes, " "))
	}

	req, err := http.NewRequestWithContext(ctx, "POST", o.tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := o.httpClientBase.Do(req)
	if err != nil {
		return fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return fmt.Errorf("failed to parse token response: %w", err)
	}

	o.tokenCache.mu.Lock()
	o.tokenCache.accessToken = tokenResp.AccessToken
	o.tokenCache.expiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	o.tokenCache.mu.Unlock()

	o.logger.Info("openemr access token refreshed",
		"expires_in", tokenResp.ExpiresIn,
	)

	return nil
}

// GetProviderType returns the provider type
func (o *OpenEMRProvider) GetProviderType() string {
	return "openemr"
}

// Read, Search, Create, Update delegate to BaseProvider (which uses our custom transport)
func (o *OpenEMRProvider) Read(ctx context.Context, resourceType, id string) (json.RawMessage, error) {
	return o.BaseProvider.Read(ctx, resourceType, id)
}

func (o *OpenEMRProvider) Search(ctx context.Context, resourceType, queryParams string) (json.RawMessage, error) {
	return o.BaseProvider.Search(ctx, resourceType, queryParams)
}

func (o *OpenEMRProvider) Create(ctx context.Context, resourceType string, resource json.RawMessage) (json.RawMessage, error) {
	return o.BaseProvider.Create(ctx, resourceType, resource)
}

func (o *OpenEMRProvider) Update(ctx context.Context, resourceType, id string, resource json.RawMessage) (json.RawMessage, error) {
	return o.BaseProvider.Update(ctx, resourceType, id, resource)
}
