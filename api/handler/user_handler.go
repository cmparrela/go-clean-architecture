package handler

import (
	"errors"
	"strconv"

	"github.com/cmparrela/go-clean-architecture/entities"
	"github.com/cmparrela/go-clean-architecture/usecases/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service user.UserServiceInterface
}

func NewUserHandler(app *fiber.App, userService user.UserServiceInterface) *UserHandler {
	handler := &UserHandler{userService}
	app.Get("/users", handler.List)
	app.Get("/users/:id", handler.Find)
	app.Put("/users/:id", handler.Update)
	app.Post("/users", handler.Create)
	app.Delete("/users/:id", handler.Delete)
	return handler
}

func (handler *UserHandler) List(context *fiber.Ctx) error {
	users, err := handler.Service.List()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}
	return context.Status(fiber.StatusOK).JSON(users)
}

func (handler *UserHandler) Find(context *fiber.Ctx) error {
	id, err := getIdParam(context)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(PayloadError{err.Error()})
	}

	user, err := handler.Service.Find(id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}
	return context.Status(fiber.StatusOK).JSON(user)
}

func (handler *UserHandler) Create(context *fiber.Ctx) error {
	user := new(entities.User)
	if err := context.BodyParser(user); err != nil {
		return err
	}

	if err := user.Validate(); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(PayloadError{err.Error()})
	}

	user, err := handler.Service.Create(user)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}
	return context.Status(fiber.StatusCreated).JSON(user)
}

func (handler *UserHandler) Update(context *fiber.Ctx) error {
	id, err := getIdParam(context)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(PayloadError{err.Error()})
	}

	user, err := handler.Service.Find(id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}

	if err := context.BodyParser(user); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}

	if err := user.Validate(); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(PayloadError{err.Error()})
	}

	result, err := handler.Service.Update(user)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}
	return context.Status(fiber.StatusOK).JSON(result)

}

func (handler *UserHandler) Delete(context *fiber.Ctx) error {
	id, err := getIdParam(context)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(PayloadError{err.Error()})
	}

	user, err := handler.Service.Find(id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}

	if err := handler.Service.Delete(user); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(PayloadError{err.Error()})
	}
	return context.SendStatus(fiber.StatusNoContent)

}

func getIdParam(context *fiber.Ctx) (uint, error) {
	id, err := strconv.ParseUint(context.Params("id"), 0, 8)
	if err != nil {
		return 0, errors.New("invalid id number")
	}
	return uint(id), nil
}
