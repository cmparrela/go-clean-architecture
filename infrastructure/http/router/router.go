package router

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app fiber.Router, userHandler handler.UserHandler, bookHandler handler.BookHandler) {
	app.Get("/", monitor.New())

	app.Get("/users", userHandler.List)
	app.Get("/users/:id", userHandler.Find)
	app.Put("/users/:id", userHandler.Update)
	app.Post("/users", userHandler.Create)
	app.Delete("/users/:id", userHandler.Delete)

	app.Get("/books", bookHandler.List)
	app.Get("/books/:id", bookHandler.Find)
	app.Put("/books/:id", bookHandler.Update)
	app.Post("/books", bookHandler.Create)
	app.Delete("/books/:id", bookHandler.Delete)
}
