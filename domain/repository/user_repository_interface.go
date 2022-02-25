package repository

import "github.com/cmparrela/go-clean-architecture/domain/entity"

type UserRepositoryInterface interface {
	List() ([]entity.User, error)
	Find(id uint) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(user *entity.User) error
}
