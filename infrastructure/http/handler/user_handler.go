package handler

import (
	"errors"
	"strconv"

	"github.com/cmparrela/go-clean-architecture/usecase/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	List(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	handler := &userHandler{service}
	return handler
}

func (h *userHandler) List(ctx *fiber.Ctx) error {
	users, err := h.service.List()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (h *userHandler) Find(ctx *fiber.Ctx) error {
	id, err := getIdParam(ctx)
	if err != nil {
		return err
	}

	user, err := h.service.Find(id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *userHandler) Create(ctx *fiber.Ctx) error {
	userDto := new(user.InputDto)
	if err := ctx.BodyParser(userDto); err != nil {
		return err
	}

	user, err := h.service.Create(userDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (h *userHandler) Update(ctx *fiber.Ctx) error {
	userDto := new(user.InputDto)
	if err := ctx.BodyParser(userDto); err != nil {
		return err
	}

	id, err := getIdParam(ctx)
	if err != nil {
		return err
	}

	result, err := h.service.Update(id, userDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(result)

}

func (h *userHandler) Delete(ctx *fiber.Ctx) error {
	id, err := getIdParam(ctx)
	if err != nil {
		return err
	}

	if err := h.service.Delete(id); err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)

}

func getIdParam(ctx *fiber.Ctx) (uint, error) {
	id, err := strconv.ParseUint(ctx.Params("id"), 0, 8)
	if err != nil {
		return 0, errors.New("invalid id number")
	}
	return uint(id), nil
}
