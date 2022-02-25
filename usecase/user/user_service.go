package user

import (
	"github.com/cmparrela/go-clean-architecture/domain/entity"
	"github.com/cmparrela/go-clean-architecture/domain/repository"
)

type UserService struct {
	Repository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{userRepository}
}

func (service *UserService) List() ([]entity.User, error) {
	users, err := service.Repository.List()
	return users, err
}

func (service *UserService) Find(id uint) (*entity.User, error) {
	user, err := service.Repository.Find(id)
	return user, err
}

func (service *UserService) Create(user *entity.User) (*entity.User, error) {
	user, err := service.Repository.Create(user)
	return user, err
}

func (service *UserService) Update(user *entity.User) (*entity.User, error) {
	user, err := service.Repository.Update(user)
	return user, err
}

func (service *UserService) Delete(user *entity.User) error {
	return service.Repository.Delete(user)
}
