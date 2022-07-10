package validator

import (
	validatorv10 "github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(obj interface{}) error
}

type validator struct {
	validatorv10 *validatorv10.Validate
}

func NewValidator(validatorv10 *validatorv10.Validate) Validator {
	return &validator{validatorv10}
}

func (v *validator) Validate(obj interface{}) error {
	return v.validatorv10.Struct(obj)
}
