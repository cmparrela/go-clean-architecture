package main

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/adapters/env"
	"github.com/cmparrela/go-clean-architecture/infrastructure/cmd"
	"github.com/cmparrela/go-clean-architecture/infrastructure/database"
)

func main() {
	env.SetupEnvFile()
	database := database.SetupDatabase()
	cmd.Execute(database)
}
