package mocks

import (
	"global-authentication/models"
)

type MockUserRepository struct {
	FindByUsernameFunc func(username string) (*models.User, error)
	CreateFunc         func(user *models.User) error
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	return m.FindByUsernameFunc(username)
}

func (m *MockUserRepository) Create(user *models.User) error {
	return m.CreateFunc(user)
}
