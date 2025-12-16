package v1

import (
	"context"

	"mesa-mestre/extension/chi"
	"mesa-mestre/extension/huma"
)

const APIPrefix = "/api/v1"

type HandlerProvider struct {
	CreateOwnerHandler func(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error)
}

func RegisterRoutes(provider HandlerProvider) chi.Router {

	r := chi.NewRouter()

	config := huma.NewConfig("Mesa Mestre API", "1.0.0")

	api := huma.NewAPI(r.C, config)

	huma.Register(api, CreateOwnerOperation(), provider.CreateOwnerHandler)

	return r
}
