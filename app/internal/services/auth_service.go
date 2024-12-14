package services

import (
	"app/internal/models"
	r "app/internal/repositories"
	"app/pkg/errors"
	"app/pkg/utils"
	"time"
)

type AuthService interface {
	Register(user *models.User) error
}

type authService struct {
	UserRepo r.UserRepository
}

func NewAuthService(userRepo r.UserRepository) AuthService {
	return &authService{
		UserRepo: userRepo,
	}
}

func (s *authService) Register(user *models.User) error {
	emailExist, err := s.UserRepo.IsEmailExist(user.Email)
	if err != nil {
		return err
	}
	if emailExist {
		return errors.ErrEmailExist
	}

	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return err
	}

	user = &models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}