package tools

import (
	"context"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
)

// Client defines the Brewfather API operations used by MCP tools.
type Client interface {
	ListBatches(ctx context.Context, opts brewfather.ListBatchesOptions) ([]brewfather.Batch, error)
	GetBatch(ctx context.Context, id string) (*brewfather.Batch, error)
	GetBatchReadings(ctx context.Context, id string, opts brewfather.ListOptions) ([]brewfather.BatchReading, error)
	ListRecipes(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Recipe, error)
	GetRecipe(ctx context.Context, id string) (*brewfather.Recipe, error)
	ListFermentables(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Fermentable, error)
	ListHops(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Hop, error)
	ListYeasts(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Yeast, error)
	ListMiscs(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Misc, error)
}
