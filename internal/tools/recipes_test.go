package tools_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
)

func TestListRecipesTool(t *testing.T) {
	client := &mockClient{
		listRecipes: func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Recipe, error) {
			if opts.Limit != 20 {
				t.Errorf("expected limit=20, got %d", opts.Limit)
			}
			return []brewfather.Recipe{
				{ID: "r1", Name: "West Coast IPA", StyleName: "American IPA", OG: 1.065, ABV: 6.5, IBU: 65},
				{ID: "r2", Name: "Oatmeal Stout", StyleName: "Oatmeal Stout", OG: 1.060, ABV: 5.5, IBU: 30},
			}, nil
		},
	}

	result := callTool(t, client, "list_recipes", map[string]any{"limit": 20})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var recipes []brewfather.Recipe
	if err := json.Unmarshal([]byte(textContent(result)), &recipes); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if len(recipes) != 2 {
		t.Errorf("expected 2 recipes, got %d", len(recipes))
	}
	if recipes[0].Name != "West Coast IPA" {
		t.Errorf("expected 'West Coast IPA', got %q", recipes[0].Name)
	}
}

func TestGetRecipeTool(t *testing.T) {
	client := &mockClient{
		getRecipe: func(ctx context.Context, id string) (*brewfather.Recipe, error) {
			if id != "recipe1" {
				t.Errorf("expected id=recipe1, got %q", id)
			}
			return &brewfather.Recipe{
				ID:        "recipe1",
				Name:      "Hoppy Lager",
				StyleName: "German Pilsner",
				OG:        1.048,
				FG:        1.010,
				ABV:       5.0,
				IBU:       25,
			}, nil
		},
	}

	result := callTool(t, client, "get_recipe", map[string]any{"id": "recipe1"})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}

	var recipe brewfather.Recipe
	if err := json.Unmarshal([]byte(textContent(result)), &recipe); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if recipe.Name != "Hoppy Lager" {
		t.Errorf("expected 'Hoppy Lager', got %q", recipe.Name)
	}
	if recipe.ABV != 5.0 {
		t.Errorf("expected ABV=5.0, got %f", recipe.ABV)
	}
}

func TestGetRecipeTool_MissingID(t *testing.T) {
	result := callTool(t, &mockClient{}, "get_recipe", map[string]any{})
	if !result.IsError {
		t.Error("expected error for missing id")
	}
}

func TestListRecipesTool_Pagination(t *testing.T) {
	client := &mockClient{
		listRecipes: func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Recipe, error) {
			if opts.StartAfter != "lastrecipeid" {
				t.Errorf("expected start_after=lastrecipeid, got %q", opts.StartAfter)
			}
			return []brewfather.Recipe{}, nil
		},
	}

	result := callTool(t, client, "list_recipes", map[string]any{"start_after": "lastrecipeid"})
	if result.IsError {
		t.Fatalf("unexpected error: %s", textContent(result))
	}
	text := textContent(result)
	if !strings.Contains(text, "[]") && !strings.Contains(text, "null") && text != "[]" {
		// OK if empty array or null
	}
}
