package server

import (
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/ragaoua/mcpg/internal/tools"
)

func Run(db_url string) error {
	mcp_server := server.NewMCPServer(
		"mcpg_data",
		"0.1",
		server.WithToolCapabilities(true),
	)

	for _, tool := range tools.GetTools(db_url) {
		mcp_server.AddTool(
			tool.Spec,
			tool.Handler,
		)
	}

	log.Println("Starting StreamableHTTP server on :8080")
	httpServer := server.NewStreamableHTTPServer(mcp_server)
	err := httpServer.Start(":8080")
	return err
}
