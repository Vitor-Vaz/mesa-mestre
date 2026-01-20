package repositories

import (
	"context"
	"errors"
	"mesa-mestre/domain"
	"mesa-mestre/extension/telemetryfs"
	"mesa-mestre/gateway/postgres/pggen"

	"github.com/lib/pq"
)

type DiningTablesRepository struct {
	q *pggen.Queries
}

func NewDiningTablesRepository(db pggen.DBTX) *DiningTablesRepository {
	return &DiningTablesRepository{
		q: pggen.New(db),
	}
}

// CreateDiningTable inserts a new dining table into the database. tableNumber and capacity are required fields.
func (r *DiningTablesRepository) CreateDiningTable(ctx context.Context, tableNumber, capacity int32) error {
	err := r.q.InsertDiningTable(ctx, pggen.InsertDiningTableParams{
		TableNumber: tableNumber,
		Capacity:    capacity,
		TableStatus: domain.DiningTableStatusAvailable.String(),
	})

	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return domain.ErrConflict
	}

	if err != nil {
		telemetryfs.Error(ctx, "failed to create dining table due to unexpected error")
		return errors.Join(domain.ErrUnexpected, err)
	}

	return nil

}
