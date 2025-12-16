package repositories

import (
	"context"
	"database/sql"
	"errors"
	"mesa-mestre/domain"
	"mesa-mestre/gateway/postgres/pggen"

	"mesa-mestre/extension/telemetryfs"

	"github.com/lib/pq"
)

type OwnersRepository struct {
	q *pggen.Queries
}

func NewOwnersRepository(db pggen.DBTX) *OwnersRepository {
	return &OwnersRepository{
		q: pggen.New(db),
	}
}

// CreateOwner inserts a new owner into the database. name and email are required fields.
func (r *OwnersRepository) CreateOwner(ctx context.Context, name string, email string) error {

	err := r.q.InsertOwner(ctx, pggen.InsertOwnerParams{
		Name:  name,
		Email: email,
	})

	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return domain.ErrConflict
	}

	if err != nil {
		telemetryfs.Error(ctx, "failed to create owner due to unexpected error")
		return errors.Join(domain.ErrUnexpected, err)
	}

	return nil
}

// FetchOwnerByEmail retrieves an owner from the database by their email.
func (r *OwnersRepository) FetchOwnerByEmail(ctx context.Context, email string) (domain.Owner, error) {
	ownerRecord, err := r.q.FetchOwnerByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Owner{}, domain.ErrNotFound
	}

	if err != nil {
		telemetryfs.Error(ctx, "failed to fetch owner by email due to unexpected error")
		return domain.Owner{}, errors.Join(domain.ErrUnexpected, err)
	}

	return domain.Owner{
		ID:        ownerRecord.ID,
		Name:      ownerRecord.Name,
		Email:     ownerRecord.Email,
		CreatedAt: ownerRecord.CreatedAt.Time,
	}, nil
}
