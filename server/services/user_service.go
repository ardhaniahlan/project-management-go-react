package services

import (
	"errors"
	"project-management-be/models"
	"project-management-be/repositories"
	"project-management-be/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	GetByID(id uint64) (*models.User, error)
	GetByPublicID(publicID string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	_, err := s.repo.FindByEmail(user.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err == nil {
		return errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashed)
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.ComparePassword(user.Password, password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *userService) GetByID(id uint64) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetByPublicID(publicID string) (*models.User, error) {
	return s.repo.FindByPublicID(publicID)
}