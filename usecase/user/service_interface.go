package user

import "github.com/cmparrela/go-clean-architecture/domain/entity"

type Service interface {
	List() ([]entity.User, error)
	Find(id uint) (entity.User, error)
	Create(input *InputDto) (*entity.User, error)
	Update(id uint, input *InputDto) (*entity.User, error)
	Delete(id uint) error
}
