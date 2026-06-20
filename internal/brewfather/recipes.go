package brewfather

import (
	"context"
	"net/url"
)

// ListRecipes returns a list of recipes.
func (c *Client) ListRecipes(ctx context.Context, opts ListOptions) ([]Recipe, error) {
	q := buildListQuery(opts)
	var recipes []Recipe
	if err := c.get(ctx, "/recipes", q, &recipes); err != nil {
		return nil, err
	}
	return recipes, nil
}

// GetRecipe returns the details of a single recipe by ID.
func (c *Client) GetRecipe(ctx context.Context, id string) (*Recipe, error) {
	var recipe Recipe
	if err := c.get(ctx, "/recipes/"+url.PathEscape(id), nil, &recipe); err != nil {
		return nil, err
	}
	return &recipe, nil
}
