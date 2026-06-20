package tools_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
	"github.com/conallob/mcp-brewfather/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// callTool registers all tools on a fresh server and invokes the named tool directly.
func callTool(t *testing.T, client tools.Client, toolName string, args map[string]any) *mcp.CallToolResult {
	t.Helper()
	s := server.NewMCPServer("test", "1.0.0", server.WithToolCapabilities(false))
	tools.Register(s, client)

	tool := s.GetTool(toolName)
	if tool == nil {
		t.Fatalf("tool %q not registered", toolName)
	}

	var req mcp.CallToolRequest
	req.Params.Name = toolName
	req.Params.Arguments = args

	result, err := tool.Handler(context.Background(), req)
	if err != nil {
		t.Fatalf("handler returned unexpected error: %v", err)
	}
	return result
}

func textContent(r *mcp.CallToolResult) string {
	var sb strings.Builder
	for _, c := range r.Content {
		if tc, ok := c.(mcp.TextContent); ok {
			sb.WriteString(tc.Text)
		}
	}
	return sb.String()
}

func TestListBatchesTool(t *testing.T) {
	client := &mockClient{
		listBatches: func(ctx context.Context, opts brewfather.ListBatchesOptions) ([]brewfather.Batch, error) {
			if opts.Status != "Fermenting" {
				t.Errorf("expected status=Fermenting, got %q", opts.Status)
			}
			if opts.Limit != 5 {
				t.Errorf("expected limit=5, got %d", opts.Limit)
			}
			return []brewfather.Batch{
				{ID: "b1", Name: "West Coast IPA", Status: "Fermenting", BatchNo: 1},
				{ID: "b2", Name: "Oatmeal Stout", Status: "Fermenting", BatchNo: 2},
			}, nil
		},
	}

	result := callTool(t, client, "list_batches", map[string]any{
		"status": "Fermenting",
		"limit":  5,
	})

	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	text := textContent(result)
	if !strings.Contains(text, "West Coast IPA") {
		t.Errorf("expected 'West Coast IPA' in result, got: %s", text)
	}

	var batches []brewfather.Batch
	if err := json.Unmarshal([]byte(text), &batches); err != nil {
		t.Fatalf("result is not valid JSON: %v\ncontent: %s", err, text)
	}
	if len(batches) != 2 {
		t.Errorf("expected 2 batches, got %d", len(batches))
	}
}

func TestListBatchesTool_NoFilter(t *testing.T) {
	client := &mockClient{
		listBatches: func(ctx context.Context, opts brewfather.ListBatchesOptions) ([]brewfather.Batch, error) {
			if opts.Status != "" {
				t.Errorf("expected no status filter, got %q", opts.Status)
			}
			return []brewfather.Batch{}, nil
		},
	}

	result := callTool(t, client, "list_batches", map[string]any{})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var batches []brewfather.Batch
	if err := json.Unmarshal([]byte(textContent(result)), &batches); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
}

func TestGetBatchTool(t *testing.T) {
	client := &mockClient{
		getBatch: func(ctx context.Context, id string) (*brewfather.Batch, error) {
			if id != "abc123" {
				t.Errorf("expected id=abc123, got %q", id)
			}
			return &brewfather.Batch{ID: "abc123", Name: "Stout", Status: "Completed", BatchNo: 5}, nil
		},
	}

	result := callTool(t, client, "get_batch", map[string]any{"id": "abc123"})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var batch brewfather.Batch
	if err := json.Unmarshal([]byte(textContent(result)), &batch); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if batch.ID != "abc123" || batch.Name != "Stout" {
		t.Errorf("unexpected batch: %+v", batch)
	}
}

func TestGetBatchTool_MissingID(t *testing.T) {
	result := callTool(t, &mockClient{}, "get_batch", map[string]any{})
	if !result.IsError {
		t.Error("expected error result for missing id param")
	}
}

func TestGetBatchTool_APIError(t *testing.T) {
	client := &mockClient{
		getBatch: func(ctx context.Context, id string) (*brewfather.Batch, error) {
			return nil, &brewfather.APIError{StatusCode: 404, Body: "not found"}
		},
	}

	result := callTool(t, client, "get_batch", map[string]any{"id": "notexist"})
	if !result.IsError {
		t.Error("expected error result for API 404")
	}
	if !strings.Contains(textContent(result), "404") {
		t.Errorf("expected 404 in error message, got: %s", textContent(result))
	}
}

func TestGetBatchReadingsTool(t *testing.T) {
	client := &mockClient{
		getBatchReadings: func(ctx context.Context, id string, opts brewfather.ListOptions) ([]brewfather.BatchReading, error) {
			if id != "batch1" {
				t.Errorf("expected id=batch1, got %q", id)
			}
			return []brewfather.BatchReading{
				{ID: "r1", Gravity: 1.060, Temperature: 20.5},
				{ID: "r2", Gravity: 1.020, Temperature: 21.0},
			}, nil
		},
	}

	result := callTool(t, client, "get_batch_readings", map[string]any{"id": "batch1"})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var readings []brewfather.BatchReading
	if err := json.Unmarshal([]byte(textContent(result)), &readings); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if len(readings) != 2 {
		t.Errorf("expected 2 readings, got %d", len(readings))
	}
	if readings[0].Gravity != 1.060 {
		t.Errorf("expected gravity 1.060, got %f", readings[0].Gravity)
	}
}

func TestGetBatchReadingsTool_MissingID(t *testing.T) {
	result := callTool(t, &mockClient{}, "get_batch_readings", map[string]any{})
	if !result.IsError {
		t.Error("expected error result for missing id param")
	}
}
