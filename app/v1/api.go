package v1

import (
	"context"
	"net/http"

	"mesa-mestre/extension/chi"
	"mesa-mestre/extension/huma"
)

type HandlerProvider struct {
	CreateOwnerHandler func(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error)
}

func RegisterRoutes(provider HandlerProvider) chi.Router {

	r := chi.NewRouter()

	config := huma.NewConfig("Mesa Mestre API", "1.0.0")

	api := huma.NewAPI(r.C, config)

	huma.Register(api, huma.Operation{
		OperationID: "create-owner",
		Method:      http.MethodPost,
		Path:        "/owners",
		Summary:     "Create a new owner",
		Tags:        []string{"Owners"},
	}, provider.CreateOwnerHandler)

	return r
}
