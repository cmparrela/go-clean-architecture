package persistence

import (
	"github.com/cmparrela/go-clean-architecture/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database}
}

func (repository *UserRepository) List() ([]entities.User, error) {
	users := []entities.User{}
	err := repository.DB.Find(&users).Error
	return users, err
}

func (repository *UserRepository) Find(id uint) (*entities.User, error) {
	user := &entities.User{}
	err := repository.DB.First(&user, id).Error
	return user, err
}

func (repository *UserRepository) Create(user *entities.User) (*entities.User, error) {
	err := repository.DB.Create(user).Error
	return user, err
}

func (repository *UserRepository) Update(user *entities.User) (*entities.User, error) {
	err := repository.DB.Save(user).Error
	return user, err
}

func (repository *UserRepository) Delete(user *entities.User) error {
	err := repository.DB.Delete(user).Error
	return err
}
