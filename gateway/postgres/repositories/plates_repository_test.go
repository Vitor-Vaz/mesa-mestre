package repositories_test

import (
	"context"
	"mesa-mestre/domain"
	"mesa-mestre/extension/testhelpers"
	"mesa-mestre/gateway/postgres/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlate(t *testing.T) {

	db := testhelpers.SetupTestDB(t)
	repo := repositories.NewPlatesRepository(db)

	name := "Spaghetti Carbonara"
	description := "Classic Italian pasta dish with eggs, cheese, pancetta, and pepper."
	price := 25.50

	t.Run("should create plate successfully", func(t *testing.T) {
		err := repo.CreatePlate(context.Background(), name, description, price)
		assert.NoError(t, err)

		var dbName, dbDescription string
		var dbPrice float64
		err = db.QueryRow(`SELECT plate_name, plate_description, price FROM plates WHERE plate_name = $1`, name).Scan(&dbName, &dbDescription, &dbPrice)
		assert.NoError(t, err)

		assert.Equal(t, name, dbName)
		assert.Equal(t, description, dbDescription)
		assert.Equal(t, price, dbPrice)
	})

	t.Run("should fail when return an unexpected error", func(t *testing.T) {
		_, _ = db.Exec(`DROP TABLE IF EXISTS plates CASCADE;`)

		err := repo.CreatePlate(context.Background(), "New Plate", "Description", 20.00)

		assert.ErrorIs(t, err, domain.ErrUnexpected)
	})

}
