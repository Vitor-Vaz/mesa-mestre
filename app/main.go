package main

import (
	"fmt"
	"mesa-mestre/extension"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
}

func main() {

	_ = godotenv.Load()

	var config Config
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	_, err := extension.NewDatabase()
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %v\n", err)
		return
	}
}
