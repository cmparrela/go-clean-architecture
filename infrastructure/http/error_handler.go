package http

import (
	"errors"
	"github.com/cmparrela/go-clean-architecture/infrastructure/customerror"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError
	payload := &PayloadError{Message: err.Error()}

	if e, ok := err.(*customerror.Error); ok {
		code = e.Code
	}

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errors := make([]string, len(ve))
		for i, fieldErr := range ve {
			errors[i] = fieldErr.Error()
		}

		code = http.StatusUnprocessableEntity
		payload.Message = "Validation error in the following fields"
		payload.Errors = errors

	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(code).JSON(payload)
}
