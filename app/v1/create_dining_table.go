package v1

import (
	"context"
	"errors"
	"mesa-mestre/domain"
	"mesa-mestre/extension/huma"
)

type CreateDiningTableBody struct {
	Number int32 `json:"number" example:"10"`
	Seats  int32 `json:"seats" example:"4"`
}

type CreateDiningTableRequest struct {
	Body CreateDiningTableBody `json:"body"`
}

type CreateDiningTableResponse struct {
}

//go:generate moq -out ./mocks/dining_table_creator_mock.go -pkg mocks . DiningTableCreator
type DiningTableCreator interface {
	CreateDiningTable(ctx context.Context, number, seats int32) error
}

type DiningTableHandler struct {
	diningTableCreator DiningTableCreator
}

func NewDiningTableHandler(diningTableCreator DiningTableCreator) *DiningTableHandler {
	return &DiningTableHandler{
		diningTableCreator: diningTableCreator,
	}
}

func (d *DiningTableHandler) CreateDiningTableHandler(ctx context.Context, req *CreateDiningTableRequest) (*CreateDiningTableResponse, error) {
	err := d.diningTableCreator.CreateDiningTable(ctx, req.Body.Number, req.Body.Seats)
	switch {
	case errors.Is(err, domain.ErrConflict):
		return nil, huma.Error409Conflict("Dining table already exists")
	case err != nil:
		return nil, huma.Error500InternalServerError()
	}

	return &CreateDiningTableResponse{}, nil
}
