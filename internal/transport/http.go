package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/langcare/langcare-mcp-fhir/internal/config"
	"github.com/langcare/langcare-mcp-fhir/internal/middleware"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// HTTPServerOptions contains options for starting the HTTP server
type HTTPServerOptions struct {
	AuthTokens []string
	RateLimit  config.RateLimitConfig
}

// StartHTTP starts the MCP server with Streamable HTTP transport
func StartHTTP(ctx context.Context, server *mcp.Server, cfg config.HTTPConfig, opts HTTPServerOptions) error {
	addr := fmt.Sprintf(":%d", cfg.Port)

	log.Printf("[HTTP] Starting MCP server with Streamable HTTP transport on %s", addr)

	// Create Streamable HTTP handler (replaces legacy SSE handler)
	handler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
		return server
	}, nil)

	// Apply middleware chain (in reverse order of execution)
	var finalHandler http.Handler = handler

	// 1. Authentication middleware (validates MCP tokens)
	if len(opts.AuthTokens) > 0 {
		log.Printf("[HTTP] MCP client authentication enabled with %d valid tokens", len(opts.AuthTokens))
		finalHandler = middleware.AuthMiddleware(opts.AuthTokens)(finalHandler)
	} else {
		log.Printf("[HTTP] WARNING: MCP client authentication disabled (no tokens configured)")
	}

	// 2. Rate limiting middleware
	if opts.RateLimit.Enabled {
		log.Printf("[HTTP] Rate limiting enabled: %d req/s, burst %d", opts.RateLimit.Rate, opts.RateLimit.Burst)
		rateLimiter := middleware.NewRateLimiter(opts.RateLimit.Rate, opts.RateLimit.Burst)
		finalHandler = rateLimiter.Middleware(finalHandler)
	}

	// 3. Security headers middleware
	finalHandler = middleware.SecurityHeadersMiddleware(finalHandler)

	// Add health check endpoint
	mux := http.NewServeMux()
	mux.Handle("/mcp", finalHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Create HTTP server
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// Handle TLS if enabled
	if cfg.TLS.Enabled {
		log.Printf("[HTTP] TLS enabled with cert: %s", cfg.TLS.CertFile)
		return httpServer.ListenAndServeTLS(cfg.TLS.CertFile, cfg.TLS.KeyFile)
	}

	log.Printf("[HTTP] Server ready on %s (no TLS)", addr)
	return httpServer.ListenAndServe()
}
