package handler

import (
	"github.com/cmparrela/go-clean-architecture/usecase/book"
	"github.com/gofiber/fiber/v2"
)

type BookHandler interface {
	List(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) BookHandler {
	return &bookHandler{service: service}
}

// List godoc
// @Summary      List book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array} domain.Book
// @Router       /books [get]
func (b *bookHandler) List(ctx *fiber.Ctx) error {
	books, err := b.service.List()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(books)
}

// Find godoc
// @Summary      Find book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200  {object} domain.Book
// @Router       /books/{id} [get]
func (b *bookHandler) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book, err := b.service.Find(id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(book)
}

// Create godoc
// @Summary      Create book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} domain.Book
// @Router       /books [post]
func (b *bookHandler) Create(ctx *fiber.Ctx) error {
	bookDto := new(book.CreateDto)
	if err := ctx.BodyParser(bookDto); err != nil {
		return err
	}

	book, err := b.service.Create(bookDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(book)
}

// Update godoc
// @Summary      Update book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} domain.Book
// @Router       /books/{id} [put]
func (b *bookHandler) Update(ctx *fiber.Ctx) error {
	bookDto := new(book.UpdateDto)
	if err := ctx.BodyParser(bookDto); err != nil {
		return err
	}

	id := ctx.Params("id")

	result, err := b.service.Update(id, bookDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(result)

}

// Delete godoc
// @Summary      Delete book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200
// @Router       /books/{id} [delete]
func (b *bookHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := b.service.Delete(id); err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)

}
