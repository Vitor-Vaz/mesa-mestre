package v1

import (
	"context"
	"mesa-mestre/extension/huma"
)

type CreatePlateBody struct {
	Name             string  `json:"name" example:"Spaghetti Carbonara"`
	PlateDescription string  `json:"plate_description" example:"Classic Italian pasta dish with eggs, cheese, pancetta, and pepper."`
	Price            float64 `json:"price" example:"25.50"`
}

type CreatePlateRequest struct {
	Body CreatePlateBody `json:"body"`
}

type CreatePlateResponse struct {
}

//go:generate moq -out ./mocks/plate_creator_mock.go -pkg mocks . PlateCreator
type PlateCreator interface {
	CreatePlate(ctx context.Context, name, plateDescription string, price float64) error
}

type PlateHandler struct {
	plateCreator PlateCreator
}

func NewPlateHandler(plateCreator PlateCreator) *PlateHandler {
	return &PlateHandler{
		plateCreator: plateCreator,
	}
}

func (p *PlateHandler) CreatePlateHandler(ctx context.Context, req *CreatePlateRequest) (*CreatePlateResponse, error) {
	err := p.plateCreator.CreatePlate(ctx, req.Body.Name, req.Body.PlateDescription, req.Body.Price)
	if err != nil {
		return nil, huma.Error500InternalServerError()
	}

	return &CreatePlateResponse{}, nil
}
