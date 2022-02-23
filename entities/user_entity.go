package entities

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name" validate:"required"`
}

func (user User) Validate() error {
	validator := validator.New()
	return validator.Struct(user)
}
