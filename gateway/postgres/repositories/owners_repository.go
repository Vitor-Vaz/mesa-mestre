package repositories

import (
	"context"
	"errors"
	"mesa-mestre/domain"
	"mesa-mestre/gateway/postgres/pggen"

	"mesa-mestre/extension/telemetryfs"

	"github.com/lib/pq"
	"go.uber.org/zap"
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
		telemetryfs.Error(ctx, "failed to create owner due to unexpected error", zap.Error(err))
		return errors.Join(domain.ErrUnexpected, err)
	}

	return nil
}
