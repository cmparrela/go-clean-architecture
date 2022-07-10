package handler

import (
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
	return &userHandler{service}
}

func (h *userHandler) List(ctx *fiber.Ctx) error {
	users, err := h.service.List()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (h *userHandler) Find(ctx *fiber.Ctx) error {
	user, err := h.service.Find(ctx.Params("id"))
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *userHandler) Create(ctx *fiber.Ctx) error {
	userDto := new(user.CreateDto)
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
	userDto := new(user.UpdateDto)
	if err := ctx.BodyParser(userDto); err != nil {
		return err
	}

	result, err := h.service.Update(ctx.Params("id"), userDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(result)

}

func (h *userHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Params("id")); err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)

}
