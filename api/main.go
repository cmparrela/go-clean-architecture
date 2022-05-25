package main

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/adapters/env"
	"github.com/cmparrela/go-clean-architecture/infrastructure/database"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http"
)

func main() {
	envConfig := env.SetupEnvFile()
	database := database.SetupDatabase(envConfig)
	http.SetupHttpServer(database)
}
