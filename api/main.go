package main

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/config"
	"github.com/cmparrela/go-clean-architecture/infrastructure/database"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http"
)

func main() {
	envConfig := config.SetupEnvFile()
	db := database.SetupDatabase(envConfig)
	http.SetupHttpServer(db)
}
