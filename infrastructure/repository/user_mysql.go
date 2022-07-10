package repository

import (
	"github.com/cmparrela/go-clean-architecture/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	List() ([]domain.User, error)
	Find(id string) (domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(user *domain.User) error
}

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{database}
}

func (r *userRepository) List() ([]domain.User, error) {
	users := []domain.User{}
	err := r.database.Find(&users).Error
	return users, err
}

func (r *userRepository) Find(id string) (domain.User, error) {
	user := domain.User{ID: id}
	err := r.database.First(&user).Error
	return user, err
}

func (r *userRepository) Create(user *domain.User) error {
	return r.database.Create(user).Error
}

func (r *userRepository) Update(user *domain.User) error {

	return r.database.Save(user).Error
}

func (r *userRepository) Delete(user *domain.User) error {
	return r.database.Delete(user).Error
}
