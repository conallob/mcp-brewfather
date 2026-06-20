package tools

import "github.com/mark3labs/mcp-go/server"

// Register adds all Brewfather MCP tools to the given server.
func Register(s *server.MCPServer, client Client) {
	registerBatchTools(s, client)
	registerRecipeTools(s, client)
	registerInventoryTools(s, client)
}
