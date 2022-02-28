package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user User) Validate() error {
	validator := validator.New()
	return validator.Struct(user)
}
