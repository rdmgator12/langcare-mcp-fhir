package fhir

import (
	"fmt"
	"log/slog"

	"github.com/langcare/langcare-mcp-fhir/internal/config"
	"github.com/langcare/langcare-mcp-fhir/internal/fhir/providers"
)

// NewClient creates a new FHIR client based on the configuration
// This factory function selects the appropriate provider implementation
func NewClient(cfg config.FHIRServerConfig, logger *slog.Logger) (providers.Provider, error) {
	logger.Info("creating fhir client",
		"provider", cfg.Provider,
		"base_url", cfg.BaseURL,
	)

	switch cfg.Provider {
	case "gcp":
		return providers.NewGCPProvider(&providers.GCPConfig{
			ProjectID:      cfg.GCP.ProjectID,
			Location:       cfg.GCP.Location,
			DatasetID:      cfg.GCP.DatasetID,
			FHIRStoreID:    cfg.GCP.FHIRStoreID,
			CredentialPath: cfg.GCP.CredentialPath,
		}, logger)

	case "epic":
		return providers.NewEPICProvider(&providers.EPICConfig{
			BaseURL:        cfg.BaseURL,
			ClientID:       cfg.EPIC.ClientID,
			PrivateKeyPath: cfg.EPIC.PrivateKeyPath,
			TokenURL:       cfg.EPIC.TokenURL,
			Scopes:         cfg.EPIC.Scopes,
		}, logger)

	case "cerner":
		return providers.NewCernerProvider(&providers.CernerConfig{
			BaseURL:      cfg.BaseURL,
			ClientID:     cfg.Cerner.ClientID,
			ClientSecret: cfg.Cerner.ClientSecret,
			TokenURL:     cfg.Cerner.TokenURL,
			Scopes:       cfg.Cerner.Scopes,
		}, logger)

	case "generic":
		return providers.NewGenericProvider(&providers.GenericConfig{
			BaseURL:  cfg.BaseURL,
			AuthType: cfg.Auth.Type,
			Token:    cfg.Auth.Token,
			Username: cfg.Auth.Username,
			Password: cfg.Auth.Password,
		}, logger)

	default:
		return nil, fmt.Errorf("unknown FHIR provider: %s", cfg.Provider)
	}
}
