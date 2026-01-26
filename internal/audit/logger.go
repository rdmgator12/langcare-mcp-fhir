package audit

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"
)

// PHIAccessEvent represents a PHI access event for audit logging
type PHIAccessEvent struct {
	Timestamp    time.Time
	UserID       string
	TokenHash    string // SHA256 hash of token, not plaintext
	Operation    string // "read", "search", "create", "update"
	ResourceType string
	ResourceID   string
	QueryParams  string
	Status       string // "success", "error"
	ErrorMessage string
	RemoteAddr   string
	RequestID    string
}

// AuditLogger handles PHI access audit logging
type AuditLogger struct {
	scrubPHI bool
}

// NewAuditLogger creates a new audit logger
func NewAuditLogger(scrubPHI bool) *AuditLogger {
	return &AuditLogger{
		scrubPHI: scrubPHI,
	}
}

// LogPHIAccess logs a PHI access event
func (a *AuditLogger) LogPHIAccess(ctx context.Context, event PHIAccessEvent) {
	// Always set timestamp
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now().UTC()
	}

	// Scrub sensitive data if enabled
	if a.scrubPHI {
		event.ResourceID = a.scrubResourceID(event.ResourceID)
	}

	// Log the event
	if event.Status == "success" {
		log.Printf("[AUDIT] PHI_ACCESS timestamp=%s operation=%s resource_type=%s resource_id=%s user=%s token_hash=%s remote_addr=%s status=%s",
			event.Timestamp.Format(time.RFC3339),
			event.Operation,
			event.ResourceType,
			event.ResourceID,
			event.UserID,
			event.TokenHash,
			event.RemoteAddr,
			event.Status,
		)
	} else {
		log.Printf("[AUDIT] PHI_ACCESS timestamp=%s operation=%s resource_type=%s resource_id=%s user=%s token_hash=%s remote_addr=%s status=%s error=%s",
			event.Timestamp.Format(time.RFC3339),
			event.Operation,
			event.ResourceType,
			event.ResourceID,
			event.UserID,
			event.TokenHash,
			event.RemoteAddr,
			event.Status,
			event.ErrorMessage,
		)
	}
}

// scrubResourceID scrubs resource IDs to prevent PHI in logs
func (a *AuditLogger) scrubResourceID(id string) string {
	if id == "" {
		return ""
	}
	// Return first 8 chars + "..." for audit trail while protecting PHI
	if len(id) <= 8 {
		return id
	}
	return id[:8] + "..."
}

// HashToken creates a SHA256 hash of a token for audit logging
func HashToken(token string) string {
	if token == "" {
		return ""
	}
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])[:16] // First 16 chars of hash
}

// GetStatusFromError returns "success" or "error" based on error
func GetStatusFromError(err error) string {
	if err == nil {
		return "success"
	}
	return "error"
}

// GetErrorMessage returns error message or empty string
func GetErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
