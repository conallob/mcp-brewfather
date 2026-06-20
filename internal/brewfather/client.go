package brewfather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const defaultBaseURL = "https://api.brewfather.app/v2"

// Client is an HTTP client for the Brewfather API.
type Client struct {
	baseURL    string
	userID     string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Brewfather API client authenticated with the given credentials.
func NewClient(userID, apiKey string) *Client {
	return &Client{
		baseURL: defaultBaseURL,
		userID:  userID,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// newClientWithBaseURL creates a client with a custom base URL (used in tests).
func newClientWithBaseURL(userID, apiKey, baseURL string) *Client {
	c := NewClient(userID, apiKey)
	c.baseURL = baseURL
	return c
}

func (c *Client) get(ctx context.Context, path string, query url.Values, result interface{}) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("parse URL: %w", err)
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.SetBasicAuth(c.userID, c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func buildListQuery(opts ListOptions) url.Values {
	q := url.Values{}
	if opts.Limit > 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.StartAfter != "" {
		q.Set("start_after", opts.StartAfter)
	}
	return q
}
