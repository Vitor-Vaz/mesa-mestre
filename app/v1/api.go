package v1

import (
	"context"
	"net/http"

	"mesa-mestre/extension/chi"
	"mesa-mestre/extension/huma"
)

const APIPrefix = "/api/v1"

type HandlerProvider struct {
	CreateOwnerHandler       func(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error)
	CreateDiningTableHandler func(ctx context.Context, req *CreateDiningTableRequest) (*CreateDiningTableResponse, error)
	CreatePlateHandler       func(ctx context.Context, req *CreatePlateRequest) (*CreatePlateResponse, error)
}

func RegisterRoutes(provider HandlerProvider) chi.Router {

	r := chi.NewRouter()

	config := huma.NewConfig("Mesa Mestre API", "1.0.0")

	api := huma.NewAPI(r.C, config)

	huma.Register(api, huma.Operation{
		OperationID: "create-owner",
		Method:      http.MethodPost,
		Path:        APIPrefix + "/owners",
		Summary:     "Create a new owner",
		Tags:        []string{"Owners"},
	}, provider.CreateOwnerHandler)

	huma.Register(api, huma.Operation{
		OperationID: "create-dining-table",
		Method:      http.MethodPost,
		Path:        APIPrefix + "/dining-table",
		Summary:     "Create a new dining table",
		Tags:        []string{"Dining Tables"},
	}, provider.CreateDiningTableHandler)

	huma.Register(api, huma.Operation{
		OperationID: "create-plate",
		Method:      http.MethodPost,
		Path:        APIPrefix + "/plate",
		Summary:     "Create a new plate",
		Tags:        []string{"Plates"},
	}, provider.CreatePlateHandler)

	return r
}
