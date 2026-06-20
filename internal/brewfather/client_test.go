package brewfather

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestClient(t *testing.T, mux *http.ServeMux) *Client {
	t.Helper()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return newClientWithBaseURL("testuser", "testkey", srv.URL)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func TestNewClient(t *testing.T) {
	c := NewClient("user1", "key1")
	if c.userID != "user1" {
		t.Errorf("expected userID=user1, got %q", c.userID)
	}
	if c.apiKey != "key1" {
		t.Errorf("expected apiKey=key1, got %q", c.apiKey)
	}
	if c.baseURL != defaultBaseURL {
		t.Errorf("expected baseURL=%q, got %q", defaultBaseURL, c.baseURL)
	}
}

func TestClient_AuthHeader(t *testing.T) {
	var gotUser, gotPass string
	mux := http.NewServeMux()
	mux.HandleFunc("/batches", func(w http.ResponseWriter, r *http.Request) {
		gotUser, gotPass, _ = r.BasicAuth()
		writeJSON(w, []Batch{})
	})
	c := newTestClient(t, mux)
	_, err := c.ListBatches(context.Background(), ListBatchesOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if gotUser != "testuser" {
		t.Errorf("expected Basic Auth user=testuser, got %q", gotUser)
	}
	if gotPass != "testkey" {
		t.Errorf("expected Basic Auth pass=testkey, got %q", gotPass)
	}
}

func TestClient_APIError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/batches", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"unauthorized"}`))
	})
	c := newTestClient(t, mux)
	_, err := c.ListBatches(context.Background(), ListBatchesOptions{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", apiErr.StatusCode)
	}
}
