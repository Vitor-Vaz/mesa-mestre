package main

import (
	"context"
	"fmt"
	v1 "mesa-mestre/app/v1"
	"mesa-mestre/extension/database"
	"mesa-mestre/extension/telemetryfs"
	"mesa-mestre/gateway/postgres/repositories"
	"net/http"
	"time"

	"github.com/caarlos0/env/v10"
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

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Printf("Error when connecting to database: %v", err)
		return
	}

	// setup repositories
	ownerRepo := repositories.NewOwnersRepository(db)

	// setup domain handlers
	createUserHandler := v1.NewOwnerHandler(ownerRepo)

	// setup useCases
	HandlerProvider := v1.HandlerProvider{
		CreateOwnerHandler: createUserHandler.CreateOwnerHandler,
	}

	// Register routes
	routes := v1.RegisterRoutes(HandlerProvider)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      routes.C,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server
	err = server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil && err != http.ErrServerClosed {
		telemetryfs.Error(ctx, "Failed to start server: %s", zap.String(err.Error(), "error"))
	}
	// test
}
