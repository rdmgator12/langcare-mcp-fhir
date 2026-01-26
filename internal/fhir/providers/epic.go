package providers

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// EPICConfig holds EPIC FHIR configuration
type EPICConfig struct {
	BaseURL        string
	ClientID       string
	PrivateKeyPath string
	TokenURL       string
	Scopes         []string
}

// EPICProvider implements FHIR access to EPIC with JWT Backend Services Auth
type EPICProvider struct {
	*BaseProvider
	clientID       string
	privateKey     *rsa.PrivateKey
	tokenURL       string
	scopes         []string
	tokenCache     *tokenCache
	httpClientBase *http.Client
	logger         *slog.Logger
}

// tokenCache manages OAuth2 token caching and refresh
type tokenCache struct {
	mu          sync.RWMutex
	accessToken string
	expiry      time.Time
}

// NewEPICProvider creates a new EPIC FHIR provider
func NewEPICProvider(cfg *EPICConfig, logger *slog.Logger) (*EPICProvider, error) {
	logger.Info("initializing epic provider",
		"base_url", cfg.BaseURL,
		"client_id", cfg.ClientID,
	)

	// Load RSA private key
	privateKey, err := loadPrivateKey(cfg.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	// Create base HTTP client (will add auth later)
	baseHTTPClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	provider := &EPICProvider{
		clientID:       cfg.ClientID,
		privateKey:     privateKey,
		tokenURL:       cfg.TokenURL,
		scopes:         cfg.Scopes,
		tokenCache:     &tokenCache{},
		httpClientBase: baseHTTPClient,
		logger:         logger,
	}

	// Create HTTP client with token injection
	httpClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: &epicTransport{provider: provider},
	}

	provider.BaseProvider = NewBaseProvider(cfg.BaseURL, httpClient, logger)

	// Pre-authenticate to validate credentials
	if err := provider.refreshToken(context.Background()); err != nil {
		return nil, fmt.Errorf("initial authentication failed: %w", err)
	}

	logger.Info("epic provider authenticated successfully")

	return provider, nil
}

// epicTransport adds OAuth2 bearer token to requests
type epicTransport struct {
	provider *EPICProvider
}

func (t *epicTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Get valid token
	token, err := t.provider.getValidToken(req.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token: %w", err)
	}

	// Add authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Use base transport
	return http.DefaultTransport.RoundTrip(req)
}

// getValidToken returns a valid access token, refreshing if necessary
func (e *EPICProvider) getValidToken(ctx context.Context) (string, error) {
	e.tokenCache.mu.RLock()
	token := e.tokenCache.accessToken
	expiry := e.tokenCache.expiry
	e.tokenCache.mu.RUnlock()

	// Token is valid if it expires more than 5 minutes from now
	if token != "" && time.Until(expiry) > 5*time.Minute {
		return token, nil
	}

	// Need to refresh
	e.tokenCache.mu.Lock()
	defer e.tokenCache.mu.Unlock()

	// Double-check after acquiring write lock
	if e.tokenCache.accessToken != "" && time.Until(e.tokenCache.expiry) > 5*time.Minute {
		return e.tokenCache.accessToken, nil
	}

	// Refresh token
	if err := e.refreshToken(ctx); err != nil {
		return "", err
	}

	return e.tokenCache.accessToken, nil
}

// refreshToken obtains a new access token using JWT assertion
func (e *EPICProvider) refreshToken(ctx context.Context) error {
	e.logger.Debug("refreshing epic access token")

	// Create JWT assertion
	now := time.Now()
	claims := jwt.MapClaims{
		"iss": e.clientID,
		"sub": e.clientID,
		"aud": e.tokenURL,
		"exp": now.Add(5 * time.Minute).Unix(),
		"iat": now.Unix(),
		"jti": generateJTI(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
	assertion, err := token.SignedString(e.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign JWT: %w", err)
	}

	// Exchange JWT for access token
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", assertion)
	if len(e.scopes) > 0 {
		data.Set("scope", strings.Join(e.scopes, " "))
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := e.httpClientBase.Do(req)
	if err != nil {
		return fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse token response
	var tokenResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return fmt.Errorf("failed to parse token response: %w", err)
	}

	// Cache token (note: lock already held by caller)
	e.tokenCache.accessToken = tokenResp.AccessToken
	e.tokenCache.expiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	e.logger.Info("epic access token refreshed",
		"expires_in", tokenResp.ExpiresIn,
	)

	return nil
}

// GetProviderType returns the provider type
func (e *EPICProvider) GetProviderType() string {
	return "epic"
}

// Read, Search, Create, Update delegate to BaseProvider (which uses our custom transport)
func (e *EPICProvider) Read(ctx context.Context, resourceType, id string) (json.RawMessage, error) {
	return e.BaseProvider.Read(ctx, resourceType, id)
}

func (e *EPICProvider) Search(ctx context.Context, resourceType, queryParams string) (json.RawMessage, error) {
	return e.BaseProvider.Search(ctx, resourceType, queryParams)
}

func (e *EPICProvider) Create(ctx context.Context, resourceType string, resource json.RawMessage) (json.RawMessage, error) {
	return e.BaseProvider.Create(ctx, resourceType, resource)
}

func (e *EPICProvider) Update(ctx context.Context, resourceType, id string, resource json.RawMessage) (json.RawMessage, error) {
	return e.BaseProvider.Update(ctx, resourceType, id, resource)
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
