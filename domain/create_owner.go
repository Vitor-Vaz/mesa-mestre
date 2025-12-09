package domain

import "context"

type OwnerCreatorRepository interface {
	CreateOwner(ctx context.Context, name string, email string) error
}

type OwnerCreatorUseCase struct {
	repo OwnerCreatorRepository
}

func NewOwnerCreatorUseCase(repo OwnerCreatorRepository) *OwnerCreatorUseCase {
	return &OwnerCreatorUseCase{repo: repo}
}

func (oc *OwnerCreatorUseCase) CreateOwner(ctx context.Context, name string, email string) error {
	return oc.repo.CreateOwner(ctx, name, email)
}
