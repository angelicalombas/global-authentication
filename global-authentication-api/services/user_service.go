package services

import (
	"errors"
	"global-authentication/models"
	"global-authentication/repositories"
	"global-authentication/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.User) error
	Login(username, password string) (string, error)
	FindByUsername(username string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(user *models.User) error {
	existingUser, err := s.userRepo.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(user)
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) FindByUsername(username string) (*models.User, error) {
	return s.userRepo.FindByUsername(username)
}
