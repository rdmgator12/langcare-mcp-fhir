package middleware

import (
	"log"
	"net/http"
	"strings"
)

// AuthMiddleware validates MCP client authentication using Bearer tokens
func AuthMiddleware(validTokens []string) func(http.Handler) http.Handler {
	// Convert to map for O(1) lookup
	tokenMap := make(map[string]bool)
	for _, token := range validTokens {
		if token != "" { // Skip empty tokens
			tokenMap[token] = true
		}
	}

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
			if !tokenMap[token] {
				// Log unauthorized access attempt
				log.Printf("[AUTH] Unauthorized access attempt from %s to %s", r.RemoteAddr, r.URL.Path)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Log authorized access
			log.Printf("[AUTH] Authorized MCP access from %s to %s", r.RemoteAddr, r.URL.Path)

			// Token is valid, continue
			next.ServeHTTP(w, r)
		})
	}
}
