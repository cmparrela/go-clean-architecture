package user

import "github.com/cmparrela/go-clean-architecture/domain/entity"

type UserServiceInterface interface {
	List() ([]entity.User, error)
	Find(id uint) (entity.User, error)
	Create(input *UserInputDto) (*entity.User, error)
	Update(id uint, input *UserInputDto) (*entity.User, error)
	Delete(id uint) error
}
