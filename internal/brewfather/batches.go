package brewfather

import (
	"context"
	"net/url"
)

// ListBatches returns a list of batches, optionally filtered by status.
func (c *Client) ListBatches(ctx context.Context, opts ListBatchesOptions) ([]Batch, error) {
	q := buildListQuery(opts.ListOptions)
	if opts.Status != "" {
		q.Set("status", opts.Status)
	}
	if opts.Complete {
		q.Set("complete", "true")
	}

	var batches []Batch
	if err := c.get(ctx, "/batches", q, &batches); err != nil {
		return nil, err
	}
	return batches, nil
}

// GetBatch returns the details of a single batch by ID.
func (c *Client) GetBatch(ctx context.Context, id string) (*Batch, error) {
	var batch Batch
	if err := c.get(ctx, "/batches/"+url.PathEscape(id), nil, &batch); err != nil {
		return nil, err
	}
	return &batch, nil
}

// GetBatchReadings returns fermentation sensor readings for a batch.
func (c *Client) GetBatchReadings(ctx context.Context, id string, opts ListOptions) ([]BatchReading, error) {
	q := buildListQuery(opts)
	var readings []BatchReading
	if err := c.get(ctx, "/batches/"+url.PathEscape(id)+"/readings", q, &readings); err != nil {
		return nil, err
	}
	return readings, nil
}
