package main

import (
	"log"

	"github.com/cmparrela/go-clean-architecture/api/router"
	"github.com/cmparrela/go-clean-architecture/infrastructure/adapters/env"
	"github.com/cmparrela/go-clean-architecture/infrastructure/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	env.SetupEnvFile()
	database := database.SetupDatabase()
	router.SetupRoutes(app, database)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
