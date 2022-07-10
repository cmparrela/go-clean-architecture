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

// List godoc
// @Summary      List user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array} domain.User
// @Router       /users [get]
func (h *userHandler) List(ctx *fiber.Ctx) error {
	users, err := h.service.List()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

// Find godoc
// @Summary      Find user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200  {object} domain.User
// @Router       /users/{id} [get]
func (h *userHandler) Find(ctx *fiber.Ctx) error {
	user, err := h.service.Find(ctx.Params("id"))
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// Create godoc
// @Summary      Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} domain.User
// @Router       /users [post]
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

// Update godoc
// @Summary      Update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} domain.User
// @Router       /users/{id} [put]
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

// Delete godoc
// @Summary      Delete user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200
// @Router       /users/{id} [delete]
func (h *userHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Params("id")); err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)

}
