package brewfather

import (
	"context"
	"net/http"
	"testing"
)

func TestListRecipes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, []Recipe{
			{ID: "r1", Name: "West Coast IPA", StyleName: "American IPA", OG: 1.065, IBU: 60},
			{ID: "r2", Name: "Oatmeal Stout", StyleName: "Oatmeal Stout", OG: 1.060, IBU: 30},
		})
	})
	c := newTestClient(t, mux)
	recipes, err := c.ListRecipes(context.Background(), ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(recipes) != 2 {
		t.Fatalf("expected 2 recipes, got %d", len(recipes))
	}
	if recipes[0].ID != "r1" {
		t.Errorf("expected r1, got %q", recipes[0].ID)
	}
}

func TestGetRecipe(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/recipes/recipe1", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, Recipe{
			ID:        "recipe1",
			Name:      "Hoppy Lager",
			StyleName: "German Pilsner",
			OG:        1.048,
			FG:        1.010,
			ABV:       5.0,
			IBU:       25,
		})
	})
	c := newTestClient(t, mux)
	recipe, err := c.GetRecipe(context.Background(), "recipe1")
	if err != nil {
		t.Fatal(err)
	}
	if recipe.Name != "Hoppy Lager" {
		t.Errorf("expected 'Hoppy Lager', got %q", recipe.Name)
	}
	if recipe.ABV != 5.0 {
		t.Errorf("expected ABV 5.0, got %f", recipe.ABV)
	}
}

func TestGetRecipe_NotFound(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/recipes/notexist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"not found"}`))
	})
	c := newTestClient(t, mux)
	_, err := c.GetRecipe(context.Background(), "notexist")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("expected 404, got %d", apiErr.StatusCode)
	}
}
