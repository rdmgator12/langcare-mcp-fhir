package transport

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// StartStdio starts the MCP server with stdio transport
func StartStdio(ctx context.Context, server *mcp.Server) error {
	fmt.Fprintln(os.Stderr, "Starting MCP server with stdio transport...")

	// Run server with stdio transport
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		return fmt.Errorf("stdio transport error: %w", err)
	}

	return nil
}
