package tools

import (
	"context"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerInventoryTools(s *server.MCPServer, client Client) {
	listOpts := func(req mcp.CallToolRequest) brewfather.ListOptions {
		return brewfather.ListOptions{
			Limit:      req.GetInt("limit", 0),
			StartAfter: req.GetString("start_after", ""),
		}
	}

	paginationParams := []mcp.ToolOption{
		mcp.WithInteger("limit",
			mcp.Description("Maximum number of results to return (1-50, default 10)"),
			mcp.Min(1),
			mcp.Max(50),
		),
		mcp.WithString("start_after",
			mcp.Description("Pagination cursor: ID of the last item from the previous page"),
		),
	}

	s.AddTool(mcp.NewTool("list_fermentables",
		append([]mcp.ToolOption{
			mcp.WithDescription("List fermentable ingredients (grains, extracts, sugars, etc.) from Brewfather inventory."),
		}, paginationParams...)...,
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		items, err := client.ListFermentables(ctx, listOpts(req))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(items)
	})

	s.AddTool(mcp.NewTool("list_hops",
		append([]mcp.ToolOption{
			mcp.WithDescription("List hop ingredients from Brewfather inventory, including alpha acid percentages and amounts."),
		}, paginationParams...)...,
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		items, err := client.ListHops(ctx, listOpts(req))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(items)
	})

	s.AddTool(mcp.NewTool("list_yeasts",
		append([]mcp.ToolOption{
			mcp.WithDescription("List yeast ingredients from Brewfather inventory, including laboratory, product ID, and attenuation."),
		}, paginationParams...)...,
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		items, err := client.ListYeasts(ctx, listOpts(req))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(items)
	})

	s.AddTool(mcp.NewTool("list_miscs",
		append([]mcp.ToolOption{
			mcp.WithDescription("List miscellaneous ingredients (finings, spices, water agents, etc.) from Brewfather inventory."),
		}, paginationParams...)...,
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		items, err := client.ListMiscs(ctx, listOpts(req))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(items)
	})
}
