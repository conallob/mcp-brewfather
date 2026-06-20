package tools

import (
	"context"
	"encoding/json"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerBatchTools(s *server.MCPServer, client Client) {
	s.AddTool(mcp.NewTool("list_batches",
		mcp.WithDescription("List brewing batches from Brewfather. Returns batch names, IDs, status, batch numbers, and measured values."),
		mcp.WithString("status",
			mcp.Description("Filter by batch status"),
			mcp.Enum("Planning", "Brewing", "Fermenting", "Conditioning", "Completed", "Archived"),
		),
		mcp.WithInteger("limit",
			mcp.Description("Maximum number of results to return (1-50, default 10)"),
			mcp.Min(1),
			mcp.Max(50),
		),
		mcp.WithString("start_after",
			mcp.Description("Pagination cursor: ID of the last item from the previous page"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		opts := brewfather.ListBatchesOptions{
			Status: req.GetString("status", ""),
			ListOptions: brewfather.ListOptions{
				Limit:      req.GetInt("limit", 0),
				StartAfter: req.GetString("start_after", ""),
			},
		}
		batches, err := client.ListBatches(ctx, opts)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(batches)
	})

	s.AddTool(mcp.NewTool("get_batch",
		mcp.WithDescription("Get full details of a specific brewing batch by its ID."),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Batch ID"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		batch, err := client.GetBatch(ctx, id)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(batch)
	})

	s.AddTool(mcp.NewTool("get_batch_readings",
		mcp.WithDescription("Get fermentation sensor readings (gravity, temperature, pressure) for a specific batch."),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Batch ID"),
		),
		mcp.WithInteger("limit",
			mcp.Description("Maximum number of readings to return (1-50, default 10)"),
			mcp.Min(1),
			mcp.Max(50),
		),
		mcp.WithString("start_after",
			mcp.Description("Pagination cursor: ID of the last reading from the previous page"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		opts := brewfather.ListOptions{
			Limit:      req.GetInt("limit", 0),
			StartAfter: req.GetString("start_after", ""),
		}
		readings, err := client.GetBatchReadings(ctx, id, opts)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return jsonResult(readings)
	})
}

func jsonResult(v interface{}) (*mcp.CallToolResult, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return mcp.NewToolResultError("failed to marshal response: " + err.Error()), nil
	}
	return mcp.NewToolResultText(string(data)), nil
}
