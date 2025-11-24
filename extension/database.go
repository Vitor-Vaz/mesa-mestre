package extension

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DatabaseHost string `env:"DATABASE_HOST" required:"true"`
	DatabasePort string `env:"DATABASE_PORT" required:"true"`
	DatabaseUser string `env:"DATABASE_USER" required:"true"`
	DatabasePass string `env:"DATABASE_PASS" required:"true"`
	DatabaseName string `env:"DATABASE_NAME" required:"true"`
}

func NewDatabase() (*sql.DB, error) {

	_ = godotenv.Load()

	var config Config
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseUser,
		config.DatabasePass,
		config.DatabaseName,
	)

	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao PostgreSQL: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao testar conexão com PostgreSQL: %v", err)
		return nil, err
	}

	log.Println("✅ Conectado ao PostgreSQL com sucesso!")
	return db, nil
}
