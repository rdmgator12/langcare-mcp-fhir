package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/langcare/langcare-mcp-fhir/internal/audit"
	"github.com/langcare/langcare-mcp-fhir/internal/config"
	"github.com/langcare/langcare-mcp-fhir/internal/fhir"
	"github.com/langcare/langcare-mcp-fhir/internal/mcp"
	"github.com/langcare/langcare-mcp-fhir/internal/tools"
	"github.com/langcare/langcare-mcp-fhir/internal/transport"
)

var (
	configPath = flag.String("config", "configs/config.yaml", "Path to configuration file")
	stdio      = flag.Bool("stdio", false, "Use stdio transport (overrides config)")
	httpMode   = flag.Bool("http", false, "Use HTTP transport (overrides config)")
	port       = flag.Int("port", 0, "HTTP port (overrides config)")
)

func main() {
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Override transport settings from flags
	if *stdio {
		cfg.Transport.Stdio = true
		cfg.Transport.HTTP.Enabled = false
	}
	if *httpMode {
		cfg.Transport.HTTP.Enabled = true
		cfg.Transport.Stdio = false
	}
	if *port > 0 {
		cfg.Transport.HTTP.Port = *port
	}

	fmt.Fprintf(os.Stderr, "=== LangCare MCP FHIR Server ===\n")
	fmt.Fprintf(os.Stderr, "Name: %s\n", cfg.MCP.Name)
	fmt.Fprintf(os.Stderr, "Version: %s\n", cfg.MCP.Version)
	fmt.Fprintf(os.Stderr, "FHIR Server: %s\n", cfg.FHIRServer.BaseURL)
	fmt.Fprintf(os.Stderr, "Auth Type: %s\n", cfg.FHIRServer.Auth.Type)
	fmt.Fprintf(os.Stderr, "PHI Scrubbing: %v\n", cfg.Logging.ScrubPHI)
	fmt.Fprintf(os.Stderr, "============================\n\n")

	// Create logger for FHIR client
	logger := slog.Default()

	// Create FHIR client with provider architecture
	fhirClient, err := fhir.NewClient(cfg.FHIRServer, logger)
	if err != nil {
		log.Fatalf("Failed to create FHIR client: %v", err)
	}
	fmt.Fprintf(os.Stderr, "[FHIR] Provider: %s\n", fhirClient.GetProviderType())
	if cfg.FHIRServer.BaseURL != "" {
		fmt.Fprintf(os.Stderr, "[FHIR] Base URL: %s\n", cfg.FHIRServer.BaseURL)
	}

	// Create audit logger
	auditLogger := audit.NewAuditLogger(cfg.Logging.ScrubPHI)
	fmt.Fprintf(os.Stderr, "[AUDIT] Audit logging enabled (PHI scrubbing: %v)\n", cfg.Logging.ScrubPHI)

	// Create tool registry
	toolRegistry := tools.NewRegistry()

	// Register FHIR tools with audit logging
	if err := toolRegistry.Register(tools.NewFHIRReadTool(fhirClient, auditLogger)); err != nil {
		log.Fatalf("Failed to register fhir_read tool: %v", err)
	}
	if err := toolRegistry.Register(tools.NewFHIRSearchTool(fhirClient, auditLogger)); err != nil {
		log.Fatalf("Failed to register fhir_search tool: %v", err)
	}
	if err := toolRegistry.Register(tools.NewFHIRCreateTool(fhirClient, auditLogger)); err != nil {
		log.Fatalf("Failed to register fhir_create tool: %v", err)
	}
	if err := toolRegistry.Register(tools.NewFHIRUpdateTool(fhirClient, auditLogger)); err != nil {
		log.Fatalf("Failed to register fhir_update tool: %v", err)
	}

	fmt.Fprintf(os.Stderr, "[TOOLS] Registered %d tools\n", len(toolRegistry.List()))
	for _, tool := range toolRegistry.List() {
		fmt.Fprintf(os.Stderr, "  - %s: %s\n", tool.Name, tool.Description)
	}

	// Create MCP server
	mcpServer, err := mcp.NewServer(fhirClient, toolRegistry, cfg)
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}

	fmt.Fprintf(os.Stderr, "[MCP] Server initialized\n\n")

	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Fprintln(os.Stderr, "\n[SHUTDOWN] Received shutdown signal, cleaning up...")
		cancel()
	}()

	// Start appropriate transport
	if cfg.Transport.Stdio {
		fmt.Fprintln(os.Stderr, "[TRANSPORT] Starting stdio transport...")
		fmt.Fprintln(os.Stderr, "[SECURITY] Stdio mode: No MCP authentication (process-level isolation only)")
		fmt.Fprintln(os.Stderr, "")
		if err := transport.StartStdio(ctx, mcpServer.GetMCPServer()); err != nil {
			log.Fatalf("Stdio transport failed: %v", err)
		}
	} else if cfg.Transport.HTTP.Enabled {
		// Get MCP auth tokens
		authTokens := cfg.GetMCPAuthTokens()

		if len(authTokens) == 0 {
			fmt.Fprintln(os.Stderr, "[SECURITY] WARNING: No MCP auth tokens configured!")
			fmt.Fprintln(os.Stderr, "[SECURITY] Set MCP_AUTH_TOKENS environment variable or security.mcp_auth_tokens in config")
			fmt.Fprintln(os.Stderr, "[SECURITY] Example: export MCP_AUTH_TOKENS='token1,token2,token3'")
			fmt.Fprintln(os.Stderr, "")
		} else {
			fmt.Fprintf(os.Stderr, "[SECURITY] MCP client authentication enabled (%d tokens)\n", len(authTokens))
		}

		if cfg.Security.RateLimit.Enabled {
			fmt.Fprintf(os.Stderr, "[SECURITY] Rate limiting: %d req/s, burst %d\n", cfg.Security.RateLimit.Rate, cfg.Security.RateLimit.Burst)
		}

		fmt.Fprintf(os.Stderr, "[TRANSPORT] Starting HTTP/SSE transport on port %d...\n", cfg.Transport.HTTP.Port)
		fmt.Fprintf(os.Stderr, "[TRANSPORT] Endpoints:\n")
		fmt.Fprintf(os.Stderr, "  - http://localhost:%d/mcp (SSE endpoint)\n", cfg.Transport.HTTP.Port)
		fmt.Fprintf(os.Stderr, "  - http://localhost:%d/health (health check)\n", cfg.Transport.HTTP.Port)
		fmt.Fprintln(os.Stderr, "")

		opts := transport.HTTPServerOptions{
			AuthTokens: authTokens,
			RateLimit:  cfg.Security.RateLimit,
		}

		if err := transport.StartHTTP(ctx, mcpServer.GetMCPServer(), cfg.Transport.HTTP, opts); err != nil {
			log.Fatalf("HTTP transport failed: %v", err)
		}
	} else {
		log.Fatal("No transport enabled. Enable stdio or HTTP in configuration.")
	}
}
