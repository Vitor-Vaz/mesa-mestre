package main

import (
	"context"
	"fmt"
	"mesa-mestre/extension/database"
	"mesa-mestre/extension/telemetryfs"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
}

func main() {

	logger, err := telemetryfs.NewLogger()
	if err != nil {
		panic(fmt.Errorf("error when creating logger: %v", err))
	}

	ctx := telemetryfs.WithLogger(context.Background(), logger)

	telemetryfs.Info(ctx, "Starting application")

	_ = godotenv.Load()

	var config Config
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	_, err = database.NewDatabase()
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %v\n", err)
		return
	}
}
