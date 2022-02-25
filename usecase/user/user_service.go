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

func (service *UserService) Find(id uint) (entity.User, error) {
	user, err := service.Repository.Find(id)
	return user, err
}

func (service *UserService) Create(input *UserInputDto) (*entity.User, error) {
	user := entity.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := user.Validate(); err != nil {
		return &user, err
	}

	err := service.Repository.Create(&user)
	return &user, err
}

func (service *UserService) Update(id uint, input *UserInputDto) (*entity.User, error) {
	user, err := service.Repository.Find(id)
	if err != nil {
		return &entity.User{}, err
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := service.Repository.Update(&user); err != nil {
		return &entity.User{}, err
	}
	return &user, err
}

func (service *UserService) Delete(id uint) error {
	user, err := service.Repository.Find(id)
	if err != nil {
		return err
	}

	return service.Repository.Delete(&user)
}
