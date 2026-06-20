package brewfather

import "context"

// ListFermentables returns fermentable ingredients from inventory.
func (c *Client) ListFermentables(ctx context.Context, opts ListOptions) ([]Fermentable, error) {
	q := buildListQuery(opts)
	var items []Fermentable
	if err := c.get(ctx, "/inventory/fermentables", q, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// ListHops returns hop ingredients from inventory.
func (c *Client) ListHops(ctx context.Context, opts ListOptions) ([]Hop, error) {
	q := buildListQuery(opts)
	var items []Hop
	if err := c.get(ctx, "/inventory/hops", q, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// ListYeasts returns yeast ingredients from inventory.
func (c *Client) ListYeasts(ctx context.Context, opts ListOptions) ([]Yeast, error) {
	q := buildListQuery(opts)
	var items []Yeast
	if err := c.get(ctx, "/inventory/yeasts", q, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// ListMiscs returns miscellaneous ingredients from inventory.
func (c *Client) ListMiscs(ctx context.Context, opts ListOptions) ([]Misc, error) {
	q := buildListQuery(opts)
	var items []Misc
	if err := c.get(ctx, "/inventory/miscs", q, &items); err != nil {
		return nil, err
	}
	return items, nil
}
