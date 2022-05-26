package customerror

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrUnprocessableEntity = fiber.ErrUnprocessableEntity
	ErrNotFound            = fiber.ErrNotFound
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	payload := &PayloadError{Message: err.Error()}

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errors := make([]string, len(ve))
		for i, fieldErr := range ve {
			errors[i] = fieldErr.Error()
		}

		code = fiber.StatusUnprocessableEntity
		payload.Message = "Validation error in the following fields"
		payload.Errors = errors

	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(code).JSON(payload)
}
