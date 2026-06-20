package brewfather

import (
	"context"
	"net/http"
	"testing"
)

func TestListBatches(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/batches", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("status") != "Fermenting" {
			t.Errorf("expected status=Fermenting, got %q", r.URL.Query().Get("status"))
		}
		if r.URL.Query().Get("limit") != "5" {
			t.Errorf("expected limit=5, got %q", r.URL.Query().Get("limit"))
		}
		writeJSON(w, []Batch{
			{ID: "batch1", Name: "IPA #1", Status: "Fermenting", BatchNo: 1},
			{ID: "batch2", Name: "Stout #2", Status: "Fermenting", BatchNo: 2},
		})
	})
	c := newTestClient(t, mux)
	batches, err := c.ListBatches(context.Background(), ListBatchesOptions{
		Status:      "Fermenting",
		ListOptions: ListOptions{Limit: 5},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(batches) != 2 {
		t.Fatalf("expected 2 batches, got %d", len(batches))
	}
	if batches[0].ID != "batch1" {
		t.Errorf("expected batch1, got %q", batches[0].ID)
	}
}

func TestGetBatch(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/batches/abc123", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, Batch{ID: "abc123", Name: "My IPA", Status: "Completed", BatchNo: 42})
	})
	c := newTestClient(t, mux)
	batch, err := c.GetBatch(context.Background(), "abc123")
	if err != nil {
		t.Fatal(err)
	}
	if batch.ID != "abc123" {
		t.Errorf("expected abc123, got %q", batch.ID)
	}
	if batch.Name != "My IPA" {
		t.Errorf("expected 'My IPA', got %q", batch.Name)
	}
}

func TestGetBatchReadings(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/batches/abc123/readings", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, []BatchReading{
			{ID: "r1", Gravity: 1.060, Temperature: 20.0},
			{ID: "r2", Gravity: 1.015, Temperature: 21.0},
		})
	})
	c := newTestClient(t, mux)
	readings, err := c.GetBatchReadings(context.Background(), "abc123", ListOptions{Limit: 10})
	if err != nil {
		t.Fatal(err)
	}
	if len(readings) != 2 {
		t.Fatalf("expected 2 readings, got %d", len(readings))
	}
	if readings[0].Gravity != 1.060 {
		t.Errorf("expected gravity 1.060, got %f", readings[0].Gravity)
	}
}

func TestListBatches_Empty(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/batches", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, []Batch{})
	})
	c := newTestClient(t, mux)
	batches, err := c.ListBatches(context.Background(), ListBatchesOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(batches) != 0 {
		t.Errorf("expected 0 batches, got %d", len(batches))
	}
}

func TestListBatches_Pagination(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/batches", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("start_after") != "lastid" {
			t.Errorf("expected start_after=lastid, got %q", r.URL.Query().Get("start_after"))
		}
		writeJSON(w, []Batch{})
	})
	c := newTestClient(t, mux)
	_, err := c.ListBatches(context.Background(), ListBatchesOptions{
		ListOptions: ListOptions{StartAfter: "lastid"},
	})
	if err != nil {
		t.Fatal(err)
	}
}
