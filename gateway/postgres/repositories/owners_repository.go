package repositories

import (
	"context"
	"database/sql"
	"mesa-mestre/gateway/postgres/pggen"
)

type OwnersRepository struct {
	db *sql.DB
	q  *pggen.Queries
}

func NewOwnersRepository(db *sql.DB) *OwnersRepository {
	return &OwnersRepository{
		db: db,
		q:  pggen.New(db),
	}
}

// CreateOwner inserts a new owner into the database. name and email are required fields.
func (r *OwnersRepository) CreateOwner(ctx context.Context, name string, email string) error {
	return r.q.InsertOwner(ctx, pggen.InsertOwnerParams{
		Name:  name,
		Email: email,
	})
}
