package http

import (
	"log"

	"github.com/cmparrela/go-clean-architecture/infrastructure/http/router"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupHttpServer(database *gorm.DB) {
	app := fiber.New()
	router.SetupRoutes(app, database)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
