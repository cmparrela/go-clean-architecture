package http

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/customerror"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/handler"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupHttpServer(
	userHandler handler.UserHandler,
	bookHandler handler.BookHandler,
) {
	app := fiber.New(fiber.Config{
		ErrorHandler: customerror.ErrorHandler,
	})

	router.SetupRoutes(
		app,
		userHandler,
		bookHandler,
	)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
