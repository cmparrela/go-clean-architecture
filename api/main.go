package main

import (
	"github.com/cmparrela/go-clean-architecture/infrastructure/config"
	"github.com/cmparrela/go-clean-architecture/infrastructure/database"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http"
	"github.com/cmparrela/go-clean-architecture/infrastructure/http/handler"
	"github.com/cmparrela/go-clean-architecture/infrastructure/identifier"
	"github.com/cmparrela/go-clean-architecture/infrastructure/repository"
	"github.com/cmparrela/go-clean-architecture/infrastructure/swagger"
	"github.com/cmparrela/go-clean-architecture/infrastructure/validator"
	"github.com/cmparrela/go-clean-architecture/usecase/book"
	"github.com/cmparrela/go-clean-architecture/usecase/user"
	validatorv10 "github.com/go-playground/validator/v10"
)

func main() {
	swagger.InitSwagger()
	envConfig := config.SetupEnvFile()

	mysql := database.InitMysql(envConfig)
	//mongo := database.InitMongo(envConfig.MongoAddress, envConfig.DatabaseName)

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	userRepository := repository.NewUserRepository(mysql)
	userService := user.NewUserService(userRepository, validator, identifier)
	userHandler := handler.NewUserHandler(userService)

	bookRepository := repository.NewBookRepository()
	bookService := book.NewService(bookRepository, validator)
	bookHandler := handler.NewBookHandler(bookService)

	http.SetupHttpServer(
		userHandler,
		bookHandler,
	)
}
