package data

import (
	"context"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	tools "githum.com/ragaoua/mcpg/internal/data/tools"
)

func Run(db_url string) error {
	mcp_server := server.NewMCPServer(
		"mcpg_data",
		"0.1",
		server.WithToolCapabilities(true),
	)

	list_all_roles_tool := mcp.NewTool(
		"list_all_roles",
		mcp.WithDescription("List all roles in the cluster"),
	)

	mcp_server.AddTool(
		list_all_roles_tool,
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return tools.ListAllRolesHandler(ctx, request, db_url)
		},
	)

	log.Println("Starting StreamableHTTP server on :8080")
	httpServer := server.NewStreamableHTTPServer(mcp_server)
	err := httpServer.Start(":8080")
	return err
}
