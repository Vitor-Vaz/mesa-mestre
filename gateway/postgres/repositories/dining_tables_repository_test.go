package repositories_test

import (
	"context"
	"mesa-mestre/domain"
	"mesa-mestre/extension/testhelpers"
	"mesa-mestre/gateway/postgres/repositories"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDiningTables(t *testing.T) {

	db := testhelpers.SetupTestDB(t)
	repo := repositories.NewDiningTablesRepository(db)

	t.Run("should create dining table successfully", func(t *testing.T) {
		err := repo.CreateDiningTable(context.Background(), 1, 4)
		assert.NoError(t, err)

		var tableNumber, capacity int
		err = db.QueryRow(`SELECT table_number, capacity FROM dining_tables WHERE table_number = $1`, 1).Scan(&tableNumber, &capacity)
		assert.NoError(t, err)
		assert.Equal(t, 1, tableNumber)
		assert.Equal(t, 4, capacity)

	})

	t.Run("should fail when creating dining table with duplicate table number", func(t *testing.T) {
		err := repo.CreateDiningTable(context.Background(), 1, 6)

		assert.ErrorIs(t, err, domain.ErrConflict)
	})

	t.Run("should fail when return an unexpected error", func(t *testing.T) {
		_, _ = db.Exec(`DROP TABLE IF EXISTS dining_tables CASCADE;`)

		err := repo.CreateDiningTable(context.Background(), 2, 4)

		assert.ErrorIs(t, err, domain.ErrUnexpected)
	})

}
