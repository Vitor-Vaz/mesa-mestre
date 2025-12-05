package repositories_test

import (
	"context"
	"database/sql"
	"fmt"
	"mesa-mestre/domain"
	"mesa-mestre/gateway/postgres/repositories"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateOwner(t *testing.T) {

	db, err := sql.Open("postgres", "postgres://myuser:mypassword@localhost:5432/mesa-mestre?sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	repo := repositories.NewOwnersRepository(tx)

	name := "Michael Scott"
	email := "michael.scott@dundlermufflinpapers.com.br"

	t.Run("should create owner sucessfully", func(t *testing.T) {
		err := repo.CreateOwner(context.Background(), name, email)
		assert.NoError(t, err)

		var dbName, dbEmail string
		err = tx.QueryRow(`SELECT name, email FROM owners WHERE email = $1`, email).Scan(&dbName, &dbEmail)
		assert.NoError(t, err)

		assert.Equal(t, name, dbName)
		assert.Equal(t, email, dbEmail)
	})

	t.Run("should fail when creating owner with duplicate email", func(t *testing.T) {
		err := repo.CreateOwner(context.Background(), "Another Name", email)
		fmt.Println(err)
		assert.ErrorIs(t, domain.ErrConflict, err)
	})

}
