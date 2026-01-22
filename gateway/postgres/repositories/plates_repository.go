package repositories

import (
	"context"
	"errors"
	"mesa-mestre/domain"
	"mesa-mestre/extension/telemetryfs"
	"mesa-mestre/gateway/postgres/pggen"
)

type PlatesRepository struct {
	q *pggen.Queries
}

func NewPlatesRepository(db pggen.DBTX) *PlatesRepository {
	return &PlatesRepository{
		q: pggen.New(db),
	}
}

// CreatePlate inserts a new plate into the database. name and price are required fields.
func (r *PlatesRepository) CreatePlate(ctx context.Context, name, description string, price float64) error {

	err := r.q.InsertPlate(ctx, pggen.InsertPlateParams{
		PlateName:        name,
		PlateDescription: description,
		Price:            price,
	})

	if err != nil {
		telemetryfs.Error(ctx, "failed to create plate due to unexpected error")
		return errors.Join(domain.ErrUnexpected, err)
	}

	return nil
}
