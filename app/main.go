package main

import (
	"context"
	"fmt"
	v1 "mesa-mestre/app/v1"
	"mesa-mestre/extension/database"
	"mesa-mestre/extension/telemetryfs"
	"net/http"

	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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

	r := chi.NewRouter()

	// Register routes
	v1.RegisterRoutes(r)

	// Start the server
	err = http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", r)
	if err != nil {
		telemetryfs.Error(ctx, "Failed to start server: %s", zap.String(err.Error(), "error"))

	}

}
