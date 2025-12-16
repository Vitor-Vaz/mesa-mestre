package huma

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

type Operation struct {
	OperationID string
	Method      string
	Path        string
	Summary     string
	Tags        []string
}

func NewConfig(appName, appVersion string) huma.Config {
	cfg := huma.Config{
		OpenAPI: &huma.OpenAPI{
			OpenAPI: "3.1.0",
			Info: &huma.Info{
				Title:   appName,
				Version: appVersion,
			},
			Components: &huma.Components{
				Schemas: huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer),
			},
		},
		OpenAPIPath:   "/openapi",
		DocsPath:      "/docs",
		SchemasPath:   "/schemas",
		Formats:       huma.DefaultFormats,
		DefaultFormat: "application/json",
	}

	return huma.Config(cfg)
}

func NewAPI(r *chi.Mux, config huma.Config) huma.API {
	return humachi.New(r, config)
}

func Register[I any, O any](api huma.API, op Operation, handler func(ctx context.Context, input *I) (*O, error)) {
	huma.Register(api, huma.Operation{
		OperationID: op.OperationID,
		Method:      op.Method,
		Path:        op.Path,
		Summary:     op.Summary,
		Tags:        op.Tags,
	}, handler)
}

// Post is a helper function to register POST handlers. It was created help setup and endpoint in test environments.
func Post[I any, O any](api huma.API, path string, handler func(ctx context.Context, input *I) (*O, error)) {
	huma.Post(api, path, handler)
}
