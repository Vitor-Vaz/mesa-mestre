package repositories_test

import (
	"context"
	"mesa-mestre/domain"
	"mesa-mestre/gateway/postgres/repositories"
	"testing"

	testehelpers "mesa-mestre/extension/testhelpers"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateOwner(t *testing.T) {

	db := testehelpers.SetupTestDB(t)
	repo := repositories.NewOwnersRepository(db)

	name := "Michael Scott"
	email := "michael.scott@dundlermufflinpapers.com.br"

	t.Run("should create owner sucessfully", func(t *testing.T) {
		err := repo.CreateOwner(context.Background(), name, email)
		assert.NoError(t, err)

		var dbName, dbEmail string
		err = db.QueryRow(`SELECT name, email FROM owners WHERE email = $1`, email).Scan(&dbName, &dbEmail)
		assert.NoError(t, err)

		assert.Equal(t, name, dbName)
		assert.Equal(t, email, dbEmail)
	})

	t.Run("should fail when creating owner with duplicate email", func(t *testing.T) {
		err := repo.CreateOwner(context.Background(), "Another Name", email)

		assert.ErrorIs(t, domain.ErrConflict, err)
	})

	t.Run("should fail when return an unexpected error", func(t *testing.T) {
		_, _ = db.Exec(`DROP TABLE IF EXISTS owners CASCADE;`)

		err := repo.CreateOwner(context.Background(), "Name", email)

		assert.ErrorIs(t, err, domain.ErrUnexpected)
	})

}
