package fhir

import (
	"github.com/langcare/langcare-mcp-fhir/internal/fhir/providers"
)

// Client is an alias for the Provider interface
// This maintains backwards compatibility with existing code
type Client = providers.Provider
