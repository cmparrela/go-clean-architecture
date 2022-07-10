package domain

import (
	"time"
)

type Book struct {
	ID        string    `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
