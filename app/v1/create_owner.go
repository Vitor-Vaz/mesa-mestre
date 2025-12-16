package v1

import (
	"context"
	"errors"
	"mesa-mestre/domain"
	"net/http"

	adapter "mesa-mestre/extension/huma"

	"github.com/danielgtaylor/huma/v2"
)

type CreateOwnerBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateOwnerRequest struct {
	Body CreateOwnerBody `json:"body"`
}

type CreateOwnerResponse struct {
}

//go:generate moq -out ./mocks/owner_creator_mock.go -pkg mocks . OnwerCreator
type OnwerCreator interface {
	CreateOwner(ctx context.Context, name string, email string) error
}

type OwnerHandler struct {
	ownerCreator OnwerCreator
}

func NewOwnerHandler(ownerCreator OnwerCreator) *OwnerHandler {
	return &OwnerHandler{
		ownerCreator: ownerCreator,
	}
}

func (o *OwnerHandler) CreateOwnerHandler(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error) {
	err := o.ownerCreator.CreateOwner(ctx, req.Body.Name, req.Body.Email)
	switch {
	case errors.Is(err, domain.ErrConflict):
		return nil, huma.Error409Conflict("Owner already exists")
	case err != nil:
		return nil, huma.Error500InternalServerError("Internal server error")
	}

	return &CreateOwnerResponse{}, nil
}

type APIError struct {
	Title  string `json:"title" doc:"Error title" example:"Conflict"`
	Status int    `json:"status" doc:"HTTP status code" example:"409"`
	Detail string `json:"detail" doc:"Error details" example:"Owner already exists"`
}

func CreateOwnerOperation() adapter.Operation {
	return adapter.Operation{
		OperationID: "create-owner",
		Method:      http.MethodPost,
		Path:        APIPrefix + "/owners",
		Summary:     "Create a new owner",
		Tags:        []string{"Owners"},

		Responses: map[string]*huma.Response{
			"204": {
				Description: "Owner created successfully",
			},
			"409": {
				Description: "Owner already exists",
				Content: map[string]*huma.MediaType{
					"application/json": {
						Examples: map[string]*huma.Example{
							"conflict": {
								Summary: "Dono já existe",
								Value: APIError{
									Title:  "Conflict",
									Status: 409,
									Detail: "Owner already exists",
								},
							},
						},
					},
				},
			},
			"500": {
				Description: "Internal server error",
				Content: map[string]*huma.MediaType{
					"application/json": {
						Examples: map[string]*huma.Example{
							"internalServerError": {
								Summary: "Erro interno do servidor",
								Value: APIError{
									Title:  "Internal Server Error",
									Status: 500,
									Detail: "Internal server error",
								},
							},
						},
					},
				},
			},
		},
	}
}
