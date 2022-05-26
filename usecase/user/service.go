package user

import (
	"github.com/cmparrela/go-clean-architecture/domain/adapter"
	"github.com/cmparrela/go-clean-architecture/domain/entity"
	"github.com/cmparrela/go-clean-architecture/domain/repository"
)

type service struct {
	repository repository.UserRepository
	validator  adapter.Validator
}

func NewUserService(repository repository.UserRepository, validator adapter.Validator) Service {
	return &service{repository, validator}
}

func (s *service) List() ([]entity.User, error) {
	users, err := s.repository.List()
	return users, err
}

func (s *service) Find(id uint) (entity.User, error) {
	user, err := s.repository.Find(id)
	return user, err
}

func (s *service) Create(input *InputDto) (*entity.User, error) {
	user := entity.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := s.validator.Validate(user); err != nil {
		return &user, err
	}

	err := s.repository.Create(&user)
	return &user, err
}

func (s *service) Update(id uint, input *InputDto) (*entity.User, error) {
	user, err := s.repository.Find(id)
	if err != nil {
		return &entity.User{}, err
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := s.repository.Update(&user); err != nil {
		return &entity.User{}, err
	}
	return &user, err
}

func (s *service) Delete(id uint) error {
	user, err := s.repository.Find(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(&user)
}
