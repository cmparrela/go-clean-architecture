package persistence

import (
	"github.com/cmparrela/go-clean-architecture/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database}
}

func (repository *UserRepository) List() ([]entity.User, error) {
	users := []entity.User{}
	err := repository.DB.Find(&users).Error
	return users, err
}

func (repository *UserRepository) Find(id uint) (*entity.User, error) {
	user := &entity.User{}
	err := repository.DB.First(&user, id).Error
	return user, err
}

func (repository *UserRepository) Create(user *entity.User) (*entity.User, error) {
	err := repository.DB.Create(user).Error
	return user, err
}

func (repository *UserRepository) Update(user *entity.User) (*entity.User, error) {
	err := repository.DB.Save(user).Error
	return user, err
}

func (repository *UserRepository) Delete(user *entity.User) error {
	err := repository.DB.Delete(user).Error
	return err
}
