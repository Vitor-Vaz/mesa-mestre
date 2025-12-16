package v1

import (
	"context"
	"errors"
	"mesa-mestre/domain"

	"mesa-mestre/extension/huma"
)

type CreateOwnerBody struct {
	Name  string `json:"name" example:"Michael Scott"`
	Email string `json:"email" example:"michael.scott@dundermifflin.com"`
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
		return nil, huma.Error409Conflict("owner already exists")
	case err != nil:
		return nil, huma.Error500InternalServerError()
	}

	return &CreateOwnerResponse{}, nil
}
