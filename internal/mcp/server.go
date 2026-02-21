package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/langcare/langcare-mcp-fhir/internal/apps"
	"github.com/langcare/langcare-mcp-fhir/internal/config"
	"github.com/langcare/langcare-mcp-fhir/internal/fhir"
	"github.com/langcare/langcare-mcp-fhir/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Server represents the MCP server
type Server struct {
	mcpServer    *mcp.Server
	toolRegistry *tools.Registry
	fhirClient   fhir.Client
	config       *config.Config
}

// NewServer creates a new MCP server with the official SDK
func NewServer(fhirClient fhir.Client, toolRegistry *tools.Registry, cfg *config.Config) (*Server, error) {
	// Create MCP server with implementation info
	impl := &mcp.Implementation{
		Name:    cfg.MCP.Name,
		Version: cfg.MCP.Version,
	}
	mcpServer := mcp.NewServer(impl, nil)

	server := &Server{
		mcpServer:    mcpServer,
		toolRegistry: toolRegistry,
		fhirClient:   fhirClient,
		config:       cfg,
	}

	// Register all tools from registry
	if err := server.registerTools(); err != nil {
		return nil, fmt.Errorf("failed to register tools: %w", err)
	}

	// Register MCP App tools and resources
	server.registerApps()

	return server, nil
}

// GetMCPServer returns the underlying MCP server
func (s *Server) GetMCPServer() *mcp.Server {
	return s.mcpServer
}

// registerTools registers all tools from the registry with the MCP server
func (s *Server) registerTools() error {
	toolInfos := s.toolRegistry.List()

	for _, toolInfo := range toolInfos {
		// Create MCP tool definition
		mcpTool := mcp.Tool{
			Name:        toolInfo.Name,
			Description: toolInfo.Description,
			InputSchema: convertToJSONSchema(toolInfo.InputSchema),
		}

		// Capture tool name for closure
		toolName := toolInfo.Name

		// Add tool using the generic AddTool function
		mcp.AddTool(s.mcpServer, &mcpTool, func(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, any, error) {
			// Get tool from registry
			tool, ok := s.toolRegistry.Get(toolName)
			if !ok {
				return nil, nil, fmt.Errorf("tool not found: %s", toolName)
			}

			// Execute tool
			result, err := tool.Execute(ctx, args)
			if err != nil {
				return &mcp.CallToolResult{
					Content: []mcp.Content{
						&mcp.TextContent{
							Text: fmt.Sprintf("Error: %v", err),
						},
					},
					IsError: true,
				}, nil, nil
			}

			// Format result as JSON
			resultJSON, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return nil, nil, fmt.Errorf("failed to marshal result: %w", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: string(resultJSON),
					},
				},
			}, nil, nil
		})
	}

	return nil
}

// registerApps registers MCP App resources and their dedicated tools.
// Each app gets a UI resource and a wrapper tool linked via _meta.ui.resourceUri.
// The View calls generic FHIR tools (fhir_search, fhir_read, etc.) via app.callServerTool().
func (s *Server) registerApps() {
	for _, appConfig := range apps.DefaultApps {
		appConfig := appConfig // capture loop variable

		html, err := apps.LoadAppHTML(appConfig.Name)
		if err != nil {
			log.Printf("Warning: could not load app %s: %v", appConfig.Name, err)
			continue
		}

		// Register the UI resource
		s.mcpServer.AddResource(
			&mcp.Resource{
				URI:         appConfig.ResourceURI,
				Name:        appConfig.Name,
				Description: appConfig.Description,
				MIMEType:    "text/html;profile=mcp-app",
			},
			func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
				return &mcp.ReadResourceResult{
					Contents: []*mcp.ResourceContents{
						{
							URI:      appConfig.ResourceURI,
							MIMEType: "text/html;profile=mcp-app",
							Text:     html,
						},
					},
				}, nil
			},
		)

		// Register the dedicated app tool with _meta.ui linking to the resource
		s.mcpServer.AddTool(
			&mcp.Tool{
				Name:        appConfig.ToolName,
				Description: appConfig.ToolDesc,
				InputSchema: json.RawMessage(`{"type": "object"}`),
				Meta: mcp.Meta{
					"ui": map[string]interface{}{
						"resourceUri": appConfig.ResourceURI,
					},
				},
			},
			func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
				return &mcp.CallToolResult{
					Content: []mcp.Content{
						&mcp.TextContent{
							Text: fmt.Sprintf("%s — open in an MCP Apps-capable host to see the interactive UI.", appConfig.Description),
						},
					},
				}, nil
			},
		)

		log.Printf("Registered MCP App: %s (tool: %s)", appConfig.Name, appConfig.ToolName)
	}
}

// convertToJSONSchema converts a map to JSON schema format
func convertToJSONSchema(schema map[string]interface{}) json.RawMessage {
	data, err := json.Marshal(schema)
	if err != nil {
		// Return empty schema on error
		return json.RawMessage(`{"type": "object"}`)
	}
	return json.RawMessage(data)
}
