package domain

import "context"

type OwnerCreatorRepository interface {
	CreateOwner(ctx context.Context, name string, email string) error
}

type OwnerCreator struct {
	repo OwnerCreatorRepository
}

func NewOwnerCreator(repo OwnerCreatorRepository) *OwnerCreator {
	return &OwnerCreator{repo: repo}
}

func (oc *OwnerCreator) CreateOwner(ctx context.Context, name string, email string) error {
	return oc.repo.CreateOwner(ctx, name, email)
}
