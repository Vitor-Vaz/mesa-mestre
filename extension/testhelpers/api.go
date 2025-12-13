package testhelpers

import (
	"context"
	"mesa-mestre/extension/chi"
	"mesa-mestre/extension/huma"
)

// CreateApiRouter is a helper function to create an API router with a POST endpoint for testing purposes.
func CreatePostApiRouter[I any, O any](path string, handlerFunc func(ctx context.Context, input *I) (*O, error)) chi.Router {
	r := chi.NewRouter()

	cfg := huma.NewConfig("Mesa Mestre API", "1.0.0")
	cfg.Transformers = nil

	api := huma.NewAPI(r.C, cfg)

	huma.Post(api, path, handlerFunc)

	return r
}
