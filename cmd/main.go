package main

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/cmd"
	"github.com/cmparrela/go-clean-architecture/infrastructure/config"
	"github.com/cmparrela/go-clean-architecture/infrastructure/database"
)

func main() {
	envConfig := config.SetupEnvFile()
	db := database.InitMysql(envConfig)
	cmd.Execute(db)
}
