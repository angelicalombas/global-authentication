package mocks

import (
	"net/http"
)

type MockUserController struct {
	RegisterFunc func(w http.ResponseWriter, r *http.Request)
	LoginFunc    func(w http.ResponseWriter, r *http.Request)
	HomeFunc     func(w http.ResponseWriter, r *http.Request)
}

func (m *MockUserController) Register(w http.ResponseWriter, r *http.Request) {
	if m.RegisterFunc != nil {
		m.RegisterFunc(w, r)
	}
}

func (m *MockUserController) Login(w http.ResponseWriter, r *http.Request) {
	if m.LoginFunc != nil {
		m.LoginFunc(w, r)
	}
}

func (m *MockUserController) Home(w http.ResponseWriter, r *http.Request) {
	if m.HomeFunc != nil {
		m.HomeFunc(w, r)
	}
}
