package main

import (
	"app/api/router"
	"app/infrastructure/adapters/env"
	"app/infrastructure/database"
	"log"

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
