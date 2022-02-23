package router

import (
	"app/api/handler"
	"app/infrastructure/repositories/persistence"
	"app/usecases/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, database *gorm.DB) {
	app.Get("/", monitor.New())

	userRepository := persistence.NewUserRepository(database)
	userService := user.NewUserService(*userRepository)
	handler.NewUserHandler(app, *userService)
}
