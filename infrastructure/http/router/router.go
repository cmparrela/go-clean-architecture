package router

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(userHandler handler.UserHandler, app fiber.Router) {
	app.Get("/", monitor.New())

	app.Get("/users", userHandler.List)
	app.Get("/users/:id", userHandler.Find)
	app.Put("/users/:id", userHandler.Update)
	app.Post("/users", userHandler.Create)
	app.Delete("/users/:id", userHandler.Delete)
}
