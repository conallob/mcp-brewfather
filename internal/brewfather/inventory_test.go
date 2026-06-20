package brewfather

import (
	"context"
	"net/http"
	"testing"
)

func TestListFermentables(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/inventory/fermentables", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, []Fermentable{
			{ID: "f1", Name: "Pale Malt", Type: "Grain", Amount: 5.0, Unit: "kg"},
			{ID: "f2", Name: "Crystal 60L", Type: "Grain", Amount: 0.5, Unit: "kg"},
		})
	})
	c := newTestClient(t, mux)
	items, err := c.ListFermentables(context.Background(), ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 fermentables, got %d", len(items))
	}
	if items[0].Name != "Pale Malt" {
		t.Errorf("expected 'Pale Malt', got %q", items[0].Name)
	}
}

func TestListHops(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/inventory/hops", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, []Hop{
			{ID: "h1", Name: "Citra", Alpha: 12.0, Amount: 100, Unit: "g"},
		})
	})
	c := newTestClient(t, mux)
	items, err := c.ListHops(context.Background(), ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 hop, got %d", len(items))
	}
	if items[0].Alpha != 12.0 {
		t.Errorf("expected alpha 12.0, got %f", items[0].Alpha)
	}
}

func TestListYeasts(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/inventory/yeasts", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, []Yeast{
			{ID: "y1", Name: "US-05", Laboratory: "Fermentis", Attenuation: 81},
		})
	})
	c := newTestClient(t, mux)
	items, err := c.ListYeasts(context.Background(), ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 yeast, got %d", len(items))
	}
	if items[0].Laboratory != "Fermentis" {
		t.Errorf("expected 'Fermentis', got %q", items[0].Laboratory)
	}
}

func TestListMiscs(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/inventory/miscs", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, []Misc{
			{ID: "m1", Name: "Irish Moss", Type: "Fining", Amount: 5, Unit: "g"},
		})
	})
	c := newTestClient(t, mux)
	items, err := c.ListMiscs(context.Background(), ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 misc, got %d", len(items))
	}
	if items[0].Type != "Fining" {
		t.Errorf("expected 'Fining', got %q", items[0].Type)
	}
}

func TestListInventory_Limit(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/inventory/fermentables", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("limit") != "10" {
			t.Errorf("expected limit=10, got %q", r.URL.Query().Get("limit"))
		}
		writeJSON(w, []Fermentable{})
	})
	c := newTestClient(t, mux)
	_, err := c.ListFermentables(context.Background(), ListOptions{Limit: 10})
	if err != nil {
		t.Fatal(err)
	}
}
