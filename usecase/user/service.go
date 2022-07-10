package user

import (
	"github.com/cmparrela/go-clean-architecture/domain"
	"github.com/cmparrela/go-clean-architecture/infrastructure/identifier"
	"github.com/cmparrela/go-clean-architecture/infrastructure/repository"
	"github.com/cmparrela/go-clean-architecture/infrastructure/validator"
)

type Service interface {
	List() ([]domain.User, error)
	Find(id string) (domain.User, error)
	Create(input *CreateDto) (*domain.User, error)
	Update(id string, input *UpdateDto) (*domain.User, error)
	Delete(id string) error
}

type service struct {
	repository repository.UserRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewUserService(
	repository repository.UserRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) List() ([]domain.User, error) {
	return s.repository.List()
}

func (s *service) Find(id string) (domain.User, error) {
	return s.repository.Find(id)
}

func (s *service) Create(input *CreateDto) (*domain.User, error) {
	user := domain.User{
		ID:    s.identifier.NewUuid(),
		Name:  input.Name,
		Email: input.Email,
	}

	if err := s.validator.Validate(user); err != nil {
		return &user, err
	}

	err := s.repository.Create(&user)
	return &user, err
}

func (s *service) Update(id string, input *UpdateDto) (*domain.User, error) {
	user, err := s.repository.Find(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := s.repository.Update(&user); err != nil {
		return nil, err
	}
	return &user, err
}

func (s *service) Delete(id string) error {
	user, err := s.repository.Find(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(&user)
}
