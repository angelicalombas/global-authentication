package repositories

import (
	"errors"
	"global-authentication/models"
	"global-authentication/utils/mocks"
	"testing"
)

func TestCreate_Success(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		CreateFunc: func(user *models.User) error {
			return nil
		},
	}

	user := &models.User{Username: "testuser", Password: "hashedpassword"}
	err := mockRepo.Create(user)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCreate_Error(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		CreateFunc: func(user *models.User) error {
			return errors.New("database error")
		},
	}

	user := &models.User{Username: "testuser", Password: "hashedpassword"}
	err := mockRepo.Create(user)

	if err == nil || err.Error() != "database error" {
		t.Errorf("expected 'database error', got %v", err)
	}
}

func TestFindByUsername_Success(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FindByUsernameFunc: func(username string) (*models.User, error) {
			return &models.User{Username: "testuser", Password: "hashedpassword"}, nil
		},
	}

	user, err := mockRepo.FindByUsername("testuser")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if user == nil || user.Username != "testuser" {
		t.Errorf("expected user 'testuser', got %v", user)
	}
}

func TestFindByUsername_NotFound(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FindByUsernameFunc: func(username string) (*models.User, error) {
			return nil, nil
		},
	}

	user, err := mockRepo.FindByUsername("nonexistentuser")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if user != nil {
		t.Errorf("expected user to be nil, got %v", user)
	}
}
