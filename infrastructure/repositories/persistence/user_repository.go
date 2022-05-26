package persistence

import (
	"github.com/cmparrela/go-clean-architecture/domain/entity"
	"github.com/cmparrela/go-clean-architecture/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) repository.UserRepository {
	return &userRepository{database}
}

func (r *userRepository) List() ([]entity.User, error) {
	users := []entity.User{}
	err := r.database.Find(&users).Error
	return users, err
}

func (r *userRepository) Find(id uint) (entity.User, error) {
	user := entity.User{}
	err := r.database.First(&user, id).Error
	return user, err
}

func (r *userRepository) Create(user *entity.User) error {
	return r.database.Create(user).Error
}

func (r *userRepository) Update(user *entity.User) error {

	return r.database.Save(user).Error
}

func (r *userRepository) Delete(user *entity.User) error {
	return r.database.Delete(user).Error
}
