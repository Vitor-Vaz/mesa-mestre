package domain

import "context"

type PlateCreatorRepository interface {
	CreatePlate(ctx context.Context, name string, description string, price int64) error
}

type PlateCreatorUseCase struct {
	repo PlateCreatorRepository
}

func NewPlateCreatorUseCase(repo PlateCreatorRepository) *PlateCreatorUseCase {
	return &PlateCreatorUseCase{repo: repo}
}

func (pc *PlateCreatorUseCase) CreatePlate(ctx context.Context, name string, description string, price int64) error {
	return pc.repo.CreatePlate(ctx, name, description, price)
}
