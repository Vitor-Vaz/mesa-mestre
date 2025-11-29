package repositories_test

import (
	"context"
	"database/sql"
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
	defer db.Close()

	repo := repositories.NewOwnersRepository(db)

	name := "Michael Scott"
	email := "michael.scott@dundlermufflinpapers.com"

	t.Run("should create owner sucessfully", func(t *testing.T) {
		err := repo.CreateOwner(context.Background(), name, email)
		assert.NoError(t, err)

		var dbName, dbEmail string
		err = db.QueryRow(`SELECT name, email FROM owners WHERE email = $1`, email).Scan(&dbName, &dbEmail)

		assert.NoError(t, err)

		assert.Equal(t, name, dbName)
		assert.Equal(t, email, dbEmail)
	})

}
