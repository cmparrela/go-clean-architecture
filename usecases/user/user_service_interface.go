package user

import "github.com/cmparrela/go-clean-architecture/entities"

type UserServiceInterface interface {
	List() ([]entities.User, error)
	Find(id uint) (*entities.User, error)
	Create(user *entities.User) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(user *entities.User) error
}
