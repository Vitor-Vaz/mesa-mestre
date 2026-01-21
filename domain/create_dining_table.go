package domain

import "context"

type DiningTableCreatorRepository interface {
	CreateDiningTable(ctx context.Context, tableNumber, capacity int32) error
}

type DiningTableCreatorUseCase struct {
	repo DiningTableCreatorRepository
}

func NewDiningTableCreatorUseCase(repo DiningTableCreatorRepository) *DiningTableCreatorUseCase {
	return &DiningTableCreatorUseCase{repo: repo}
}

func (dtc *DiningTableCreatorUseCase) CreateDiningTables(ctx context.Context, tableNumber, capacity int32) error {
	return dtc.repo.CreateDiningTable(ctx, tableNumber, capacity)
}
