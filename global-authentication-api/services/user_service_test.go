package services

import (
	"global-authentication/models"
	"global-authentication/utils/mocks"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FindByUsernameFunc: func(username string) (*models.User, error) {
			return nil, nil
		},
		CreateFunc: func(user *models.User) error {
			return nil
		},
	}

	userService := NewUserService(mockRepo)

	user := &models.User{Username: "testuser", Password: "password123"}
	err := userService.Register(user)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if user.Password == "password123" {
		t.Error("expected password to be hashed, but it is still the original")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password123")); err != nil {
		t.Error("expected password to be hashed, but comparison failed")
	}
}

func TestRegister_UsernameAlreadyTaken(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FindByUsernameFunc: func(username string) (*models.User, error) {
			return &models.User{Username: "testuser"}, nil
		},
	}

	userService := NewUserService(mockRepo)

	user := &models.User{Username: "testuser", Password: "password123"}
	err := userService.Register(user)

	if err == nil || err.Error() != "username already taken" {
		t.Errorf("expected 'username already taken' error, got %v", err)
	}
}

func TestLogin_Success(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FindByUsernameFunc: func(username string) (*models.User, error) {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
			return &models.User{Username: "testuser", Password: string(hashedPassword)}, nil
		},
	}

	userService := NewUserService(mockRepo)

	token, err := userService.Login("testuser", "password123")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if token == "" {
		t.Error("expected a valid token, got empty string")
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FindByUsernameFunc: func(username string) (*models.User, error) {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
			return &models.User{Username: "testuser", Password: string(hashedPassword)}, nil
		},
	}

	userService := NewUserService(mockRepo)

	_, err := userService.Login("testuser", "wrongpassword")

	if err == nil || err.Error() != "invalid credentials" {
		t.Errorf("expected 'invalid credentials' error, got %v", err)
	}
}
