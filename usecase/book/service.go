package book

import (
	"github.com/cmparrela/go-clean-architecture/domain"
	"github.com/cmparrela/go-clean-architecture/infrastructure/repository"
	"github.com/cmparrela/go-clean-architecture/infrastructure/validator"
)

type Service interface {
	List() ([]*domain.Book, error)
	Find(id string) (*domain.Book, error)
	Create(input *CreateDto) (*domain.Book, error)
	Update(id string, input *UpdateDto) (*domain.Book, error)
	Delete(id string) error
}

type service struct {
	repository repository.BookRepository
	validator  validator.Validator
}

func NewService(repository repository.BookRepository, validator validator.Validator) Service {
	return &service{repository: repository, validator: validator}
}

func (s service) List() ([]*domain.Book, error) {
	return s.repository.List()
}

func (s service) Find(id string) (*domain.Book, error) {
	return s.repository.Find(id)
}

func (s service) Create(input *CreateDto) (*domain.Book, error) {
	book := domain.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	if err := s.validator.Validate(book); err != nil {
		return &book, err
	}

	err := s.repository.Create(&book)
	return &book, err
}

func (s service) Update(id string, input *UpdateDto) (*domain.Book, error) {
	book, err := s.repository.Find(id)
	if err != nil {
		return nil, err
	}

	book.Title = input.Title
	book.Author = input.Author

	if err := s.repository.Update(book); err != nil {
		return nil, err
	}
	return book, err
}

func (s service) Delete(id string) error {
	book, err := s.repository.Find(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(book)
}
