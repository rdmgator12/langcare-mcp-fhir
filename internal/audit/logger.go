package audit

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type contextKey string

const (
	ctxKeyTokenHash  contextKey = "auditTokenHash"
	ctxKeyRemoteAddr contextKey = "auditRemoteAddr"
	ctxKeyRequestID  contextKey = "auditRequestID"
)

func WithRequestContext(ctx context.Context, tokenHash, remoteAddr string) context.Context {
	requestID := generateRequestID()
	ctx = context.WithValue(ctx, ctxKeyTokenHash, tokenHash)
	ctx = context.WithValue(ctx, ctxKeyRemoteAddr, remoteAddr)
	ctx = context.WithValue(ctx, ctxKeyRequestID, requestID)
	return ctx
}

func generateRequestID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("req_%s", hex.EncodeToString(b))
}

func getContextString(ctx context.Context, key contextKey) string {
	if v, ok := ctx.Value(key).(string); ok {
		return v
	}
	return ""
}

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

// LogPHIAccess logs a PHI access event, enriching from request context.
func (a *AuditLogger) LogPHIAccess(ctx context.Context, event PHIAccessEvent) {
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now().UTC()
	}
	if event.TokenHash == "" {
		event.TokenHash = getContextString(ctx, ctxKeyTokenHash)
	}
	if event.RemoteAddr == "" {
		event.RemoteAddr = getContextString(ctx, ctxKeyRemoteAddr)
	}
	if event.RequestID == "" {
		event.RequestID = getContextString(ctx, ctxKeyRequestID)
	}

	if a.scrubPHI {
		event.ResourceID = a.scrubResourceID(event.ResourceID)
		event.ErrorMessage = scrubErrorMessage(event.ErrorMessage)
	}

	log.Printf("[AUDIT] PHI_ACCESS timestamp=%s request_id=%s operation=%s resource_type=%s resource_id=%s token_hash=%s remote_addr=%s status=%s error=%s",
		event.Timestamp.Format(time.RFC3339),
		event.RequestID,
		event.Operation,
		event.ResourceType,
		event.ResourceID,
		event.TokenHash,
		event.RemoteAddr,
		event.Status,
		event.ErrorMessage,
	)
}

func scrubErrorMessage(msg string) string {
	if len(msg) > 200 {
		return msg[:200] + "...[truncated]"
	}
	return msg
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
