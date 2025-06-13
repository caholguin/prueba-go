package repository

import (
	"errors"
	"prueba-go/config"
	"prueba-go/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindById(id uint) (*models.User, error) {
    var user models.User
    result := config.DB.First(&user, id)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, result.Error
    }

    return &user, nil
}


func (r *userRepository) UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func (r *userRepository) DeleteUser(id *models.User) error {
	return config.DB.Delete(id).Error
}