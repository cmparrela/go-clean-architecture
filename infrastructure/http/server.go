package http

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/customerror"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/handler"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/router"
	"github.com/cmparrela/go-clean-architecture/infrastructure/repository"
	"github.com/cmparrela/go-clean-architecture/infrastructure/validator"
	"github.com/cmparrela/go-clean-architecture/usecase/book"
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

	validator := validator.NewValidator(pvalidator.New())
	userRepository := repository.NewUserRepository(database)
	userService := user.NewUserService(userRepository, validator)
	userHandler := handler.NewUserHandler(userService)

	bookRepository := repository.NewBookRepository()
	bookService := book.NewService(bookRepository, validator)
	bookHandler := handler.NewBookHandler(bookService)

	router.SetupRoutes(
		app,
		userHandler,
		bookHandler,
	)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
