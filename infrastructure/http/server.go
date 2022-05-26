package http

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/adapter"
	"github.com/cmparrela/go-clean-architecture/infrastructure/customerror"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/handler"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/router"
	"github.com/cmparrela/go-clean-architecture/infrastructure/repositories/persistence"
	"github.com/cmparrela/go-clean-architecture/usecase/user"
	pvalidator "github.com/go-playground/validator/v10"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupHttpServer(database *gorm.DB) {
	app := fiber.New(fiber.Config{
		ErrorHandler: customerror.ErrorHandler,
	})

	validator := adapter.NewValidator(pvalidator.New())
	userRepository := persistence.NewUserRepository(database)
	userService := user.NewUserService(userRepository, validator)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRoutes(userHandler, app)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
