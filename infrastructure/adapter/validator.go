package adapter

import (
	"github.com/cmparrela/go-clean-architecture/domain/adapter"
	pvalidator "github.com/go-playground/validator/v10"
)

type validatorAdapter struct {
	validator *pvalidator.Validate
}

func NewValidator(validator *pvalidator.Validate) adapter.Validator {
	return &validatorAdapter{validator}
}

func (v *validatorAdapter) Validate(obj interface{}) error {
	return v.validator.Struct(obj)
}
