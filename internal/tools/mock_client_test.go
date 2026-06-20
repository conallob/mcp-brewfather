package tools_test

import (
	"context"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
)

type mockClient struct {
	listBatches      func(ctx context.Context, opts brewfather.ListBatchesOptions) ([]brewfather.Batch, error)
	getBatch         func(ctx context.Context, id string) (*brewfather.Batch, error)
	getBatchReadings func(ctx context.Context, id string, opts brewfather.ListOptions) ([]brewfather.BatchReading, error)
	listRecipes      func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Recipe, error)
	getRecipe        func(ctx context.Context, id string) (*brewfather.Recipe, error)
	listFermentables func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Fermentable, error)
	listHops         func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Hop, error)
	listYeasts       func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Yeast, error)
	listMiscs        func(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Misc, error)
}

func (m *mockClient) ListBatches(ctx context.Context, opts brewfather.ListBatchesOptions) ([]brewfather.Batch, error) {
	return m.listBatches(ctx, opts)
}

func (m *mockClient) GetBatch(ctx context.Context, id string) (*brewfather.Batch, error) {
	return m.getBatch(ctx, id)
}

func (m *mockClient) GetBatchReadings(ctx context.Context, id string, opts brewfather.ListOptions) ([]brewfather.BatchReading, error) {
	return m.getBatchReadings(ctx, id, opts)
}

func (m *mockClient) ListRecipes(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Recipe, error) {
	return m.listRecipes(ctx, opts)
}

func (m *mockClient) GetRecipe(ctx context.Context, id string) (*brewfather.Recipe, error) {
	return m.getRecipe(ctx, id)
}

func (m *mockClient) ListFermentables(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Fermentable, error) {
	return m.listFermentables(ctx, opts)
}

func (m *mockClient) ListHops(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Hop, error) {
	return m.listHops(ctx, opts)
}

func (m *mockClient) ListYeasts(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Yeast, error) {
	return m.listYeasts(ctx, opts)
}

func (m *mockClient) ListMiscs(ctx context.Context, opts brewfather.ListOptions) ([]brewfather.Misc, error) {
	return m.listMiscs(ctx, opts)
}
