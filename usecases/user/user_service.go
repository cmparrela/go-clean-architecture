package user

import (
	"github.com/cmparrela/go-clean-architecture/entities"
	"github.com/cmparrela/go-clean-architecture/infrastructure/repositories/persistence"
)

type UserService struct {
	Repository persistence.UserRepositoryInterface
}

func NewUserService(userRepository persistence.UserRepositoryInterface) *UserService {
	return &UserService{userRepository}
}

func (service *UserService) List() ([]entities.User, error) {
	users, err := service.Repository.List()
	return users, err
}

func (service *UserService) Find(id uint) (*entities.User, error) {
	user, err := service.Repository.Find(id)
	return user, err
}

func (service *UserService) Create(user *entities.User) (*entities.User, error) {
	user, err := service.Repository.Create(user)
	return user, err
}

func (service *UserService) Update(user *entities.User) (*entities.User, error) {
	user, err := service.Repository.Update(user)
	return user, err
}

func (service *UserService) Delete(user *entities.User) error {
	return service.Repository.Delete(user)
}
