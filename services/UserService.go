package services

import (
	"prueba-go/models"
	"prueba-go/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) FindAll() ([]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *userService) FindById(id uint) (*models.User, error) {
	return s.userRepository.FindById(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id *models.User) error {
	return s.userRepository.DeleteUser(id)
}