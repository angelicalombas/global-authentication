package mocks

import (
	"global-authentication/models"
)

type MockUserService struct {
	RegisterFunc func(user *models.User) error
	LoginFunc    func(username, password string) (string, error)
}

func (m *MockUserService) Register(user *models.User) error {
	return m.RegisterFunc(user)
}

func (m *MockUserService) Login(username, password string) (string, error) {
	return m.LoginFunc(username, password)
}

func (m *MockUserService) FindByUsername(username string) (*models.User, error) {
	return nil, nil
}
