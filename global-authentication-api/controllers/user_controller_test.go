package controllers

import (
	"bytes"
	"encoding/json"
	"global-authentication/models"
	"global-authentication/utils/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister_Success(t *testing.T) {
	mockService := &mocks.MockUserService{
		RegisterFunc: func(user *models.User) error {
			return nil
		},
	}

	controller := NewUserController(mockService)

	user := models.User{Username: "testuser", Password: "password123"}
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if response["message"] != "User created successfully." {
		t.Errorf("unexpected response: got %v want %v", response["message"], "User created successfully.")
	}
}

func TestRegister_ValidationError(t *testing.T) {
	mockService := &mocks.MockUserService{}

	controller := NewUserController(mockService)

	user := models.User{Username: "testuser"}
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestLogin_Success(t *testing.T) {
	mockService := &mocks.MockUserService{
		LoginFunc: func(username, password string) (string, error) {
			return "valid-token", nil
		},
	}

	controller := NewUserController(mockService)

	user := models.User{Username: "testuser", Password: "password123"}
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	token := rr.Body.String()
	if token != "valid-token" {
		t.Errorf("unexpected token: got %v want %v", token, "valid-token")
	}
}
