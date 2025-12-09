package v1

import (
	"context"
	"fmt"
	"net/http"

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

func (h *OwnerHandler) CreateOwnerHandler(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error) {
	// Handler logic to create an owner

	fmt.Println("to criando um dono")

	return &CreateOwnerResponse{}, nil
}

func CreateOwnerOperation() huma.Operation {
	return huma.Operation{
		OperationID: "create-owner",
		Method:      http.MethodPost,
		Path:        "/owners",
		Summary:     "Create a new owner",
		Tags:        []string{"Owners"},
	}
}
