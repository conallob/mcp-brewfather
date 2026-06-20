package tools

import (
	"context"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerRecipeTools(s *server.MCPServer, client Client) {
	s.AddTool(mcp.NewTool("list_recipes",
		mcp.WithDescription("List brewing recipes from Brewfather. Returns recipe names, styles, and key statistics (OG, FG, ABV, IBU, color)."),
		mcp.WithInteger("limit",
			mcp.Description("Maximum number of results to return (1-50, default 10)"),
			mcp.Min(1),
			mcp.Max(50),
		),
		mcp.WithString("start_after",
			mcp.Description("Pagination cursor: ID of the last item from the previous page"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		opts := brewfather.ListOptions{
			Limit:      req.GetInt("limit", 0),
			StartAfter: req.GetString("start_after", ""),
		}
		recipes, err := client.ListRecipes(ctx, opts)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(recipes)
	})

	s.AddTool(mcp.NewTool("get_recipe",
		mcp.WithDescription("Get full details of a specific brewing recipe by its ID."),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Recipe ID"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		recipe, err := client.GetRecipe(ctx, id)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(recipe)
	})
}
