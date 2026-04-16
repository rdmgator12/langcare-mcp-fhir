package middleware

import (
	"crypto/subtle"
	"log"
	"net/http"
	"strings"

	"github.com/langcare/langcare-mcp-fhir/internal/audit"
)

func constantTimeContains(tokens []string, candidate string) bool {
	for _, t := range tokens {
		if subtle.ConstantTimeCompare([]byte(t), []byte(candidate)) == 1 {
			return true
		}
	}
	return false
}

// AuthMiddleware validates MCP client authentication using Bearer tokens
func AuthMiddleware(validTokens []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for health check endpoints
			if r.URL.Path == "/health" || r.URL.Path == "/healthz" {
				next.ServeHTTP(w, r)
				return
			}

			// Extract Bearer token
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Printf("[AUTH] Missing Authorization header from %s", r.RemoteAddr)
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				log.Printf("[AUTH] Invalid Authorization header format from %s", r.RemoteAddr)
				http.Error(w, "Invalid Authorization header format. Expected: Bearer <token>", http.StatusUnauthorized)
				return
			}

			token := parts[1]
			if !constantTimeContains(validTokens, token) {
				log.Printf("[AUTH] Unauthorized access attempt from %s to %s", r.RemoteAddr, r.URL.Path)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			log.Printf("[AUTH] Authorized MCP access from %s to %s", r.RemoteAddr, r.URL.Path)

			ctx := audit.WithRequestContext(r.Context(), audit.HashToken(token), r.RemoteAddr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
