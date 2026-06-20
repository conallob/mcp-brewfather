package tools_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
)

func TestListFermentablesTool(t *testing.T) {
	client := &mockClient{
		listFermentables: func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Fermentable, error) {
			return []brewfather.Fermentable{
				{ID: "f1", Name: "Pale Malt (2-Row)", Type: "Grain", Amount: 5.0, Unit: "kg", PPG: 37},
				{ID: "f2", Name: "Crystal 60L", Type: "Grain", Amount: 0.5, Unit: "kg"},
			}, nil
		},
	}

	result := callTool(t, client, "list_fermentables", map[string]any{})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var items []brewfather.Fermentable
	if err := json.Unmarshal([]byte(textContent(result)), &items); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 fermentables, got %d", len(items))
	}
	if items[0].Name != "Pale Malt (2-Row)" {
		t.Errorf("unexpected name: %q", items[0].Name)
	}
}

func TestListHopsTool(t *testing.T) {
	client := &mockClient{
		listHops: func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Hop, error) {
			return []brewfather.Hop{
				{ID: "h1", Name: "Citra", Alpha: 12.5, Amount: 100, Unit: "g", Origin: "US"},
				{ID: "h2", Name: "Mosaic", Alpha: 11.5, Amount: 50, Unit: "g"},
			}, nil
		},
	}

	result := callTool(t, client, "list_hops", map[string]any{})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var items []brewfather.Hop
	if err := json.Unmarshal([]byte(textContent(result)), &items); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 hops, got %d", len(items))
	}
	if items[0].Alpha != 12.5 {
		t.Errorf("expected alpha=12.5, got %f", items[0].Alpha)
	}
}

func TestListYeastsTool(t *testing.T) {
	client := &mockClient{
		listYeasts: func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Yeast, error) {
			return []brewfather.Yeast{
				{ID: "y1", Name: "US-05", Laboratory: "Fermentis", ProductID: "US-05", Attenuation: 81},
			}, nil
		},
	}

	result := callTool(t, client, "list_yeasts", map[string]any{})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var items []brewfather.Yeast
	if err := json.Unmarshal([]byte(textContent(result)), &items); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if len(items) != 1 || items[0].Laboratory != "Fermentis" {
		t.Errorf("unexpected yeasts: %v", items)
	}
}

func TestListMiscsTool(t *testing.T) {
	client := &mockClient{
		listMiscs: func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Misc, error) {
			return []brewfather.Misc{
				{ID: "m1", Name: "Irish Moss", Type: "Fining", Amount: 5, Unit: "g"},
				{ID: "m2", Name: "Gypsum", Type: "Water Agent", Amount: 3, Unit: "g"},
			}, nil
		},
	}

	result := callTool(t, client, "list_miscs", map[string]any{})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var items []brewfather.Misc
	if err := json.Unmarshal([]byte(textContent(result)), &items); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 miscs, got %d", len(items))
	}
}

func TestListInventory_Pagination(t *testing.T) {
	client := &mockClient{
		listHops: func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Hop, error) {
			if opts.Limit != 10 {
				t.Errorf("expected limit=10, got %d", opts.Limit)
			}
			if opts.StartAfter != "lasthop" {
				t.Errorf("expected start_after=lasthop, got %q", opts.StartAfter)
			}
			return []brewfather.Hop{}, nil
		},
	}

	result := callTool(t, client, "list_hops", map[string]any{
		"limit":       10,
		"start_after": "lasthop",
	})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}
}
