package repositories_test

import (
	"context"
	"mesa-mestre/domain"
	"mesa-mestre/gateway/postgres/repositories"
	"testing"

	"mesa-mestre/extension/testhelpers"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateOwner(t *testing.T) {

	db := testhelpers.SetupTestDB(t)
	repo := repositories.NewOwnersRepository(db)

	name := "Michael Scott"
	email := "michael.scott@dundlermufflinpapers.com.br"

	ctx := context.Background()

	t.Run("should create owner sucessfully", func(t *testing.T) {
		err := repo.CreateOwner(ctx, name, email)
		assert.NoError(t, err)

		var dbName, dbEmail string
		err = db.QueryRow(`SELECT name, email FROM owners WHERE email = $1`, email).Scan(&dbName, &dbEmail)
		assert.NoError(t, err)

		assert.Equal(t, name, dbName)
		assert.Equal(t, email, dbEmail)
	})

	t.Run("should fail when creating owner with duplicate email", func(t *testing.T) {
		err := repo.CreateOwner(ctx, "Another Name", email)

		assert.ErrorIs(t, domain.ErrConflict, err)
	})

	t.Run("should fail when return an unexpected error", func(t *testing.T) {
		_, _ = db.Exec(`DROP TABLE IF EXISTS owners CASCADE;`)

		err := repo.CreateOwner(ctx, "Name", email)

		assert.ErrorIs(t, err, domain.ErrUnexpected)
	})

}

func TestFetchOWnerByEmail(t *testing.T) {

	db := testhelpers.SetupTestDB(t)
	repo := repositories.NewOwnersRepository(db)

	name := "Pam Beesly"
	email := "pam.beesly@dundlermufflinpapers.com.br"

	ctx := context.Background()

	t.Run("should fetch owner by email successfully", func(t *testing.T) {
		err := repo.CreateOwner(ctx, name, email)
		assert.NoError(t, err)

		owner, err := repo.FetchOwnerByEmail(ctx, email)
		assert.NoError(t, err)

		assert.Equal(t, name, owner.Name)
		assert.Equal(t, email, owner.Email)

	})

	t.Run("should return not found error when owner does not exist", func(t *testing.T) {
		_, err := repo.FetchOwnerByEmail(ctx, "any.email@gmail.com")

		assert.ErrorIs(t, err, domain.ErrNotFound)
	})

	t.Run("should return unexpected error when db fails", func(t *testing.T) {
		_, _ = db.Exec(`DROP TABLE IF EXISTS owners CASCADE;`)

		_, err := repo.FetchOwnerByEmail(ctx, email)

		assert.ErrorIs(t, err, domain.ErrUnexpected)
	})

}
