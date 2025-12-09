package v1

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerProvider struct {
	CreateOwnerHandler func(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error)
}

func RegisterRoutes(provider HandlerProvider) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	config := huma.DefaultConfig("Mesa Mestre API", "1.0.0")

	api := humachi.New(r, config)

	huma.Register(api, huma.Operation{
		OperationID: "create-owner",
		Method:      http.MethodPost,
		Path:        "/owners",
		Summary:     "Create a new owner",
		Tags:        []string{"Owners"},
	}, provider.CreateOwnerHandler)

	return r
}
